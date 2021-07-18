package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
)

func main() {
	ctx := context.Background()
	b, err := ioutil.ReadFile("key.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	//config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(context.Background())

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo"
	titleRange := "Typical"
	resp, err := srv.Spreadsheets.Get(spreadsheetId).Ranges(titleRange).IncludeGridData(true).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	if len(resp.Sheets) > 0 {
		cellValue := resp.Sheets[0].Data[0].RowData[0].Values[0]
		formattedValue := cellValue.FormattedValue
		backgroundColor := cellValue.EffectiveFormat.BackgroundColor
		fmt.Println("cellValue", formattedValue)
		fmt.Printf("%02x ", int(backgroundColor.Red*255))
		fmt.Printf("%02x ", int(backgroundColor.Green*255))
		fmt.Printf("%02x ", int(backgroundColor.Blue*255))
	} else {
		fmt.Println("No data found.", resp)
	}

}
