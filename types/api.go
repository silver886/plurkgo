package types

// TimelineGetPlurks is the data responsed by Plurk API 2.0 (/APP/Timeline/getPlurks).
type TimelineGetPlurks struct {
	Users map[int]User `json:"plurk_users"`
	Posts []Post       `json:"plurks"`
}

// TimelineGetPlurk is the data responsed by Plurk API 2.0 (/APP/Timeline/getPlurk).
type TimelineGetPlurk struct {
	Post  Post         `json:"plurk"`
	User  User         `json:"user"`
	Users map[int]User `json:"plurk_users"`
}
