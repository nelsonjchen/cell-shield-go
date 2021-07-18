package main

import (
	"fmt"
	"testing"
)

func TestGrabShieldInformation(t *testing.T) {
	shieldInfo, err := GrabShieldInformation("", "1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo", "Typical")

	if err != nil {
		t.Errorf("Error grabbing shield: %s", err)
	}

	fmt.Printf("%s", *shieldInfo)

	if shieldInfo.ShieldTitle != "Value" {
		t.Errorf("Shield Title is incorrect")
	}
	if shieldInfo.FormattedValue != "9001" {
		t.Errorf("Formatted Value is incorrect")
	}
	if shieldInfo.ColorHex != "93c47d" {
		t.Errorf("ColorHex is incorrect")
	}
}

func TestGrabShieldInformationWithTitle(t *testing.T) {
	shieldInfo, err := GrabShieldInformation("Title", "1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo", "Typical")

	if err != nil {
		t.Errorf("Error grabbing shield: %s", err)
	}

	fmt.Printf("%s", *shieldInfo)

	if shieldInfo.ShieldTitle != "Title" {
		t.Errorf("Shield Title is incorrect")
	}
	if shieldInfo.FormattedValue != "9001" {
		t.Errorf("Formatted Value is incorrect")
	}
	if shieldInfo.ColorHex != "93c47d" {
		t.Errorf("ColorHex is incorrect")
	}
}

func TestGrabShieldInformationWithNoTitleButColonInValue(t *testing.T) {
	shieldInfo, err := GrabShieldInformation("", "1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo", "B3")

	if err != nil {
		t.Errorf("Error grabbing shield: %s", err)
	}

	fmt.Printf("%s", *shieldInfo)

	if shieldInfo.ShieldTitle != "Money Count" {
		t.Errorf("Shield Title is incorrect: %s", shieldInfo.ShieldTitle)
	}
	if shieldInfo.FormattedValue != "9001" {
		t.Errorf("Formatted Value is incorrect: %s", shieldInfo.FormattedValue)
	}
	if shieldInfo.ColorHex != "ffffff" {
		t.Errorf("ColorHex is incorrect: %s", shieldInfo.ColorHex)
	}
}
