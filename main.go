package main

import (
	"cell-shield-go/shieldinformation"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

type ShieldsIoJson struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label   string `json:"label"`
	Message string `json:"message"`
	Color   string `json:"color"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			return
		}
	})
	http.HandleFunc("/gs", func(w http.ResponseWriter, r *http.Request) {
		spreadSheetId := r.URL.Query().Get("spreadSheetId")
		if spreadSheetId == "" {
			err := json.NewEncoder(w).Encode(ShieldsIoJson{
				SchemaVersion: 1,
				Label:   "Error",
				Message: "spreadSheetId is missing!",
				Color:   "red",
			})
			if err != nil {
				return
			}
			log.Println("spreadSheetId missing from:", r.URL.RawQuery)
			return
		}
		cellRange := r.URL.Query().Get("cellRange")
		if cellRange == "" {
			err := json.NewEncoder(w).Encode(ShieldsIoJson{
				SchemaVersion: 1,
				Label:   "Error",
				Message: "cellRange is missing!",
				Color:   "red",
			})
			if err != nil {
				return
			}
			log.Println("cellRange missing from:", r.URL.RawQuery)
			return
		}
		shieldLabel := r.URL.Query().Get("shieldLabel")

		shieldInformation, err := shieldinformation.GrabShieldInformation(shieldLabel, spreadSheetId, cellRange)
		if err != nil {
			err := json.NewEncoder(w).Encode(ShieldsIoJson{
				SchemaVersion: 1,
				Label:   "Error",
				Message: fmt.Sprintf("Unable to retrieve shield information | %s", err),
				Color:   "red",
			})
			if err != nil {
				return
			}
			log.Println("Unable to retrieve shield information for", r.URL.RawQuery)
			return
		}

		err = json.NewEncoder(w).Encode(ShieldsIoJson{
			SchemaVersion: 1,
			Label:   shieldInformation.ShieldLabel,
			Message: shieldInformation.ShieldMessage,
			Color:   shieldInformation.ColorHex,
		})

		if err != nil {
			err := json.NewEncoder(w).Encode(ShieldsIoJson{
				SchemaVersion: 1,
				Label:   "Error",
				Message: fmt.Sprintf("Unable to generate shield information | %s", err),
				Color:   "red",
			})
			if err != nil {
				return
			}
		}
		log.Println("Sent shield information for", r.URL.RawQuery)
	})

	log.Println("Beginning to Listen")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
