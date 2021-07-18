package shieldinformation

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/sheets/v4"
	"strings"
)

type ShieldInformation struct {
	ShieldLabel   string
	ShieldMessage string
	ColorHex      string
}

func GrabShieldInformation(shieldLabel string, spreadsheetId string, cellRange string) (*ShieldInformation, error) {
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
		return nil, errors.New("range specified but somehow no data was returned")
	}

	if len(resp.Sheets[0].Data[0].RowData) == 0 {
		return nil, errors.New("no data in cell")
	}

	cellValue := resp.Sheets[0].Data[0].RowData[0].Values[0]
	formattedValue := cellValue.FormattedValue
	backgroundColor := cellValue.EffectiveFormat.BackgroundColor

	if shieldLabel == "" {
		labelAndValue := strings.Split(formattedValue, ":")
		// Trim Spaces after split
		for i := range labelAndValue {
			labelAndValue[i] = strings.TrimSpace(labelAndValue[i])
		}
		if len(labelAndValue) > 1 {
			shieldLabel = labelAndValue[0]
			formattedValue = labelAndValue[1]
		} else {
			shieldLabel = "Value"
		}
	}

	shieldInformation := ShieldInformation{
		ShieldLabel:   shieldLabel,
		ShieldMessage: formattedValue,
		ColorHex: fmt.Sprintf(
			"%02x%02x%02x",
			int(backgroundColor.Red*255),
			int(backgroundColor.Green*255),
			int(backgroundColor.Blue*255),
		),
	}
	return &shieldInformation, nil
}
