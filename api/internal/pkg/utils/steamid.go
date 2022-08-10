package utils

import (
	"regexp"
)

const (
	reProfileBase = `(?:(?:(?:(?:https?)?:\/\/)?(?:www\.)?steamcommunity\.com)?)?\/?`
	reCommunityID = `^(\d{17})$`
	reSteamID2    = `^(STEAM_\d+:\d+:\d+)$`
	reSteamID3    = `^(\[U:\d+:\d+\])$`
	reProfileURL  = reProfileBase + `profiles\/(\d{17})`
	reProfileID   = reProfileBase + `id\/([a-z0-9_-]{2,32})`
)

// GetSteamID returns the username or steamid from a community URL
func GetSteamID(value string) string {
	var matchs []string
	// community id match, ex. 76561198378422474
	r, _ := regexp.Compile(reCommunityID)
	matchs = r.FindStringSubmatch(value)
	if len(matchs) > 0 {
		return matchs[1]
	}

	// url, https://steamcommunity.com/profiles/76561198378422474
	r, _ = regexp.Compile(reProfileURL)
	matchs = r.FindStringSubmatch(value)
	if len(matchs) > 0 {
		return matchs[1]
	}

	// vanity id, https://steamcommunity.com/id/xDim
	r, _ = regexp.Compile(reProfileID)
	matchs = r.FindStringSubmatch(value)
	if len(matchs) > 0 {
		return matchs[1]
	}

	return ""
}

func IsSteamID(value string) bool {
	// community id match, ex. 76561198378422474
	m, _ := regexp.Match(reCommunityID, []byte(value))

	return m
}
