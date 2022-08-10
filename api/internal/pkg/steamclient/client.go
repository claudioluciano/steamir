package steamclient

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	httpClient "github.com/claudioluciano/simple_httpclient/pkg"
)

const (
	steam_base_url = "https://api.steampowered.com"
)

type SteamClient struct {
	http     *httpClient.Client
	steamKey string
}

func New(steamKey string) *SteamClient {
	c := httpClient.New(&httpClient.NewClientOptions{
		BaseURL:            steam_base_url,
		DefaultContentType: "application/json",
		Timeout:            30 * time.Second,
		Attemps:            3,
	})

	return &SteamClient{
		steamKey: steamKey,
		http:     c,
	}
}

func (s *SteamClient) ResolveVanityURL(ctx context.Context, value string) (*ResolveVanityURLResponse, error) {
	res, err := s.http.Do(ctx, &httpClient.DoOptions{
		Request: &httpClient.Request{
			URL:    "/ISteamUser/ResolveVanityURL/v1/",
			Method: httpClient.GET,
			Query: map[string]string{
				"key":       s.steamKey,
				"vanityurl": value,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var r ResolveVanityURLResponse
	err = json.Unmarshal([]byte(res.Body), &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (s *SteamClient) GetPlayerSummaries(ctx context.Context, steamids []string) (*GetPlayerSummariesResponse, error) {
	res, err := s.http.Do(ctx, &httpClient.DoOptions{
		Request: &httpClient.Request{
			URL:    "/ISteamUser/GetPlayerSummaries/v2/",
			Method: httpClient.GET,
			Query: map[string]string{
				"key":      s.steamKey,
				"steamids": strings.Join(steamids, ","),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var r GetPlayerSummariesResponse
	err = json.Unmarshal([]byte(res.Body), &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (s *SteamClient) GetOwnedGames(ctx context.Context, steamid string) (*GetOwnedGames, error) {
	res, err := s.http.Do(ctx, &httpClient.DoOptions{
		Request: &httpClient.Request{
			URL:    "/IPlayerService/GetOwnedGames/v1/",
			Method: httpClient.GET,
			Query: map[string]string{
				"key":                       s.steamKey,
				"steamid":                   steamid,
				"include_appinfo":           "true",
				"include_played_free_games": "true",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var r GetOwnedGames
	err = json.Unmarshal([]byte(res.Body), &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (s *SteamClient) GetFriendList(ctx context.Context, steamid string) (*GetFriendList, error) {
	res, err := s.http.Do(ctx, &httpClient.DoOptions{
		Request: &httpClient.Request{
			URL:    "/ISteamUser/GetFriendList/v1/",
			Method: httpClient.GET,
			Query: map[string]string{
				"key":     s.steamKey,
				"steamid": steamid,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var r GetFriendList
	err = json.Unmarshal([]byte(res.Body), &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
