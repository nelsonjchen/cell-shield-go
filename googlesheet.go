package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/sheets/v4"
	"strings"
)

type ShieldInformation struct {
	ShieldTitle    string
	FormattedValue string
	ColorHex       string
}

func GrabShieldInformation(shieldTitle string, spreadsheetId string, cellRange string) (*ShieldInformation, error) {
	ctx := context.Background()

	srv, err := sheets.NewService(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := srv.Spreadsheets.Get(spreadsheetId).Ranges(cellRange).IncludeGridData(true).Do()
	if err != nil {
		return nil, err
	}
	if len(resp.Sheets) == 0 {
		return nil, errors.New("Range specified but somehow no data was returned")

	}

	cellValue := resp.Sheets[0].Data[0].RowData[0].Values[0]
	formattedValue := cellValue.FormattedValue
	backgroundColor := cellValue.EffectiveFormat.BackgroundColor

	if shieldTitle == "" {
		titleAndValue := strings.Split(formattedValue, ":")
		// Trim Spaces after split
		for i := range titleAndValue {
			titleAndValue[i] = strings.TrimSpace(titleAndValue[i])
		}
		if len(titleAndValue) > 1 {
			shieldTitle = titleAndValue[0]
			formattedValue = titleAndValue[1]
		} else {
			shieldTitle = "Value"
		}
	}

	shieldInformation := ShieldInformation{
		ShieldTitle:    shieldTitle,
		FormattedValue: formattedValue,
		ColorHex: fmt.Sprintf(
			"%02x%02x%02x",
			int(backgroundColor.Red*255),
			int(backgroundColor.Green*255),
			int(backgroundColor.Blue*255),
		),
	}
	return &shieldInformation, nil
}
