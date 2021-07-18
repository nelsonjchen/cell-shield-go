package main

import (
	"fmt"
	"testing"
)

func TestGrabShieldInformation(t *testing.T) {
	shieldInfo, err := GrabShieldInformation(
		"1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo", "Typical")

	if err != nil {
		t.Errorf("Error grabbing shield: %s", err)
	}

	fmt.Printf("%s", *shieldInfo)

	if shieldInfo.FormattedValue != "9001" {
		t.Errorf("Formatted Value is incorrect")
	}
	if shieldInfo.ColorHex != "93c47d" {
		t.Errorf("ColorHex is incorrect")
	}
}
