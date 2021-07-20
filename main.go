package main

import (
	"cell-shield-go/shieldinformation"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/ReneKroon/ttlcache/v2"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

type ShieldsIoJson struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
}

//go:embed static/*
var content embed.FS

func main() {
	serverRoot, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}

	var cache ttlcache.SimpleCache = ttlcache.NewCache()
	err = cache.SetTTL(time.Duration(30 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(serverRoot)))

	http.HandleFunc("/gs", func(w http.ResponseWriter, r *http.Request) {
		spreadSheetId := r.URL.Query().Get("spreadSheetId")
		if spreadSheetId == "" {
			err := json.NewEncoder(w).Encode(ShieldsIoJson{
				SchemaVersion: 1,
				Label:         "Error",
				Message:       "spreadSheetId is missing!",
				Color:         "red",
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
				Label:         "Error",
				Message:       "cellRange is missing!",
				Color:         "red",
			})
			if err != nil {
				return
			}
			log.Println("cellRange missing from:", r.URL.RawQuery)
			return
		}

		cacheKey := "Sheet " + spreadSheetId + " | Range " + cellRange
		shieldInformationRetrieve, err := cache.Get(cacheKey)
		if err == ttlcache.ErrNotFound {
			log.Println("Cache was not hit for", cacheKey)
			freshShieldInformation, shErr := shieldinformation.GrabShieldInformation(spreadSheetId, cellRange)
			shieldInformationRetrieve = freshShieldInformation
			if shErr != nil {
				err := json.NewEncoder(w).Encode(ShieldsIoJson{
					SchemaVersion: 1,
					Label:         "Error",
					Message:       fmt.Sprintf("Unable to retrieve shield information | %s", shErr),
					Color:         "red",
				})
				if err != nil {
					return
				}
				log.Println("Unable to retrieve shield information for", r.URL.RawQuery, shErr)
				return
			}
			chErr := cache.Set(cacheKey, shieldInformationRetrieve)
			if chErr != nil {
				// If we can't set the cache, we don't care oh well.
				log.Println("Unable to set", cacheKey)
			}
		} else {
			log.Println("Cache was hit for", cacheKey)
		}
		shieldInformation := shieldInformationRetrieve.(*shieldinformation.ShieldInformation)

		err = json.NewEncoder(w).Encode(ShieldsIoJson{
			SchemaVersion: 1,
			Label:         shieldInformation.ShieldLabel,
			Message:       shieldInformation.ShieldMessage,
			Color:         shieldInformation.ColorHex,
		})

		if err != nil {
			err := json.NewEncoder(w).Encode(ShieldsIoJson{
				SchemaVersion: 1,
				Label:         "Error",
				Message:       fmt.Sprintf("Unable to generate shield information | %s", err),
				Color:         "red",
			})
			if err != nil {
				return
			}
		}
		log.Println("Sent shield information for", r.URL.RawQuery)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
