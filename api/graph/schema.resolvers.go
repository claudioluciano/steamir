package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/claudioluciano/steamir/api/graph/generated"
	"github.com/claudioluciano/steamir/api/graph/model"
	"github.com/claudioluciano/steamir/api/internal/pkg/utils"
)

// Player is the resolver for the player field.
func (r *queryResolver) Player(ctx context.Context, steamid string) (model.PlayerResult, error) {
	steamid = utils.GetSteamID(steamid)

	if !utils.IsSteamID(steamid) {
		res, err := r.SteamClient.ResolveVanityURL(ctx, steamid)
		if err != nil {
			return nil, err
		}
		if res.Response.Success == 42 {
			return model.NotFoundError{
				Message: "User not found",
			}, nil
		}

		steamid = res.Response.Steamid
	}

	res, err := r.SteamClient.GetPlayerSummaries(ctx, []string{steamid})
	if err != nil {
		return nil, err
	}

	p := res.Response.Players[0]

	return &model.SteamUser{
		Steamid:    p.Steamid,
		Name:       p.PersonaName,
		ProfileURL: p.ProfileUrl,
		AvatarURL:  p.AvatarFull,
	}, nil
}

// Friends is the resolver for the Friends field.
func (r *queryResolver) Friends(ctx context.Context, steamid string) (model.PlayersResult, error) {
	resFriends, err := r.SteamClient.GetFriendList(ctx, steamid)
	if err != nil {
		return nil, err
	}

	var steamdis []string
	for _, f := range resFriends.Friendslist.Friends {
		steamdis = append(steamdis, f.Steamid)
	}

	resPlayers, err := r.SteamClient.GetPlayerSummaries(ctx, steamdis)
	if err != nil {
		return nil, err
	}

	var us []*model.SteamUser
	for _, p := range resPlayers.Response.Players {
		us = append(us, &model.SteamUser{
			Steamid:    p.Steamid,
			Name:       p.PersonaName,
			ProfileURL: p.ProfileUrl,
			AvatarURL:  p.AvatarFull,
		})
	}

	return model.SteamUserList{
		Users: us,
		Count: len(us),
	}, nil
}

// Players is the resolver for the Players field.
func (r *queryResolver) Players(ctx context.Context, steamids []string) (model.PlayersResult, error) {
	res, err := r.SteamClient.GetPlayerSummaries(ctx, steamids)
	if err != nil {
		return nil, err
	}

	var us []*model.SteamUser
	for _, p := range res.Response.Players {
		us = append(us, &model.SteamUser{
			Steamid:    p.Steamid,
			Name:       p.PersonaName,
			ProfileURL: p.ProfileUrl,
			AvatarURL:  p.AvatarFull,
		})
	}

	return model.SteamUserList{
		Users: us,
		Count: len(us),
	}, nil
}

// OwnedGames is the resolver for the OwnedGames field.
func (r *queryResolver) OwnedGames(ctx context.Context, steamid string) (model.GameResult, error) {
	res, err := r.SteamClient.GetOwnedGames(ctx, steamid)
	if err != nil {
		return nil, err
	}

	var gs []*model.Game
	for _, g := range res.Response.Games {
		gs = append(gs, &model.Game{
			Appid:           g.Appid,
			Name:            g.Name,
			PlaytimeForever: g.PlayTimeForever,
			ImageURL:        fmt.Sprintf("https://steamcdn-a.akamaihd.net/steam/apps/%v/capsule_231x87.jpg", g.Appid),
		})
	}

	return model.GameList{
		Games: gs,
		Count: len(gs),
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
