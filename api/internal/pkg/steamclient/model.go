package steamclient

type ResolveVanityURLResponse struct {
	Response Response `json:"response"`
}

type GetPlayerSummariesResponse struct {
	Response Response `json:"response"`
}

type Players struct {
	Steamid                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	PersonaName              string `json:"personaname"`
	CommentPermission        int    `json:"commentpermission"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	AvatarHash               string `json:"avatarhash"`
	LastLogoff               int    `json:"lastlogoff"`
	PersonaState             int    `json:"personastate"`
	PrimaryClanID            string `json:"primaryclanid"`
	TimeCreated              int    `json:"timecreated"`
	PersonaStateFlags        int    `json:"personastateflags"`
}

type GetOwnedGames struct {
	Response Response `json:"response"`
}

type Games struct {
	Appid                    int    `json:"appid"`
	Name                     string `json:"name"`
	PlayTimeForever          int    `json:"playtime_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats,omitempty"`
	PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
	PlaytimeMacForever       int    `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
}

type Response struct {
	Players   []Players `json:"players"`
	Steamid   string    `json:"steamid"`
	Success   int       `json:"success"`
	GameCount int       `json:"game_count"`
	Games     []Games   `json:"games"`
	Message   string    `json:"message"`
}

type GetFriendList struct {
	Friendslist Friendslist `json:"friendslist"`
}

type Friend struct {
	Steamid      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

type Friendslist struct {
	Friends []Friend `json:"friends"`
}
