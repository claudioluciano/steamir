package utils

import (
	"fmt"
	"testing"
)

func TestGetSteamID(t *testing.T) {
	value := GetSteamID("76561197962949921")
	fmt.Println(value)

	value = GetSteamID("https://steamcommunity.com/profiles/76561197962949921")
	fmt.Println(value)

	value = GetSteamID("https://steamcommunity.com/id/shelwn/")
	fmt.Println(value)
}

func TestIsSteamID(t *testing.T) {
	value := IsSteamID("76561197962949921")
	fmt.Println(value)

	value = IsSteamID("https://steamcommunity.com/profiles/76561197962949921")
	fmt.Println(value)

	value = IsSteamID("https://steamcommunity.com/id/shelwn/")
	fmt.Println(value)
}
