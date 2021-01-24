package types

// TimelineGetPlurks is the data responsed by Plurk API (/APP/Timeline/getPlurks).
type TimelineGetPlurks struct {
	Users map[int]User `json:"plurk_users"`
	Posts []Post       `json:"plurks"`
}
