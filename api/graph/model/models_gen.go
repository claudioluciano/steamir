// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Game struct {
	Appid           int    `json:"appid"`
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtime_forever"`
	ImageURL        string `json:"image_url"`
}

type SteamUser struct {
	Steamid    string `json:"steamid"`
	Name       string `json:"name"`
	ProfileURL string `json:"profile_url"`
	AvatarURL  string `json:"avatar_url"`
}
