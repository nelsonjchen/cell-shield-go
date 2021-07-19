package shieldinformation

import (
	"fmt"
	"testing"
)

func TestGrabShieldInformation(t *testing.T) {
	shieldInfo, err := GrabShieldInformation("1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo", "Typical")

	if err != nil {
		t.Errorf("Error grabbing shield: %s\n", err)
	}

	fmt.Printf("%s\n", *shieldInfo)

	if shieldInfo.ShieldLabel != "Value" {
		t.Errorf("Shield Label is incorrect: %s\n", shieldInfo.ShieldLabel)
	}
	if shieldInfo.ShieldMessage != "9001" {
		t.Errorf("Message is incorrect: %s\n", shieldInfo.ShieldMessage)
	}
	if shieldInfo.ColorHex != "93c47d" {
		t.Errorf("ColorHex is incorrect: %s\n", shieldInfo.ColorHex)
	}
}

func TestGrabShieldInformationWithTitle(t *testing.T) {
	shieldInfo, err := GrabShieldInformation("1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo", "Typical")

	if err != nil {
		t.Errorf("Error grabbing shield: %s\n", err)
	}

	fmt.Printf("%s\n", *shieldInfo)

	if shieldInfo.ShieldLabel != "Value" {
		t.Errorf("Shield Label is incorrect: %s\n", shieldInfo.ShieldLabel)
	}
	if shieldInfo.ShieldMessage != "9001" {
		t.Errorf("Message is incorrect: %s\n", shieldInfo.ShieldMessage)
	}
	if shieldInfo.ColorHex != "93c47d" {
		t.Errorf("ColorHex is incorrect: %s\n", shieldInfo.ColorHex)
	}
}

func TestGrabShieldInformationWithNoTitleButColonInValue(t *testing.T) {
	shieldInfo, err := GrabShieldInformation("1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo", "B3")

	if err != nil {
		t.Errorf("Error grabbing shield: %s\n", err)
	}

	fmt.Printf("%s\n", *shieldInfo)

	if shieldInfo.ShieldLabel != "Money Count" {
		t.Errorf("Shield Label is incorrect: %s\n", shieldInfo.ShieldLabel)
	}
	if shieldInfo.ShieldMessage != "9001" {
		t.Errorf("Message is incorrect: %s\n", shieldInfo.ShieldMessage)
	}
	if shieldInfo.ColorHex != "ffffff" {
		t.Errorf("ColorHex is incorrect: %s\n", shieldInfo.ColorHex)
	}
}
