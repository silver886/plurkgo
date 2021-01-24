package types

// User is the data of a user responsed by Plurk API 2.0.
type User struct {
	ID       int    `json:"id"`
	Premium  bool   `json:"premium"`
	Name     string `json:"display_name"`
	Color    string `json:"name_color"`
	Avatar   int    `json:"avatar"`
	Birthday Time   `json:"date_of_birth"`
	NickName string `json:"nick_name"`
}

// Post is the data of a post responsed by Plurk API 2.0.
type Post struct {
	ID        int      `json:"plurk_id"`
	LimitedTo IntSlice `json:"limited_to"`

	Anonymous   bool   `json:"anonymous"`
	Bookmark    bool   `json:"bookmark"`
	Coins       int    `json:"coins"`
	ContentHTML string `json:"content"`
	ContentRaw  string `json:"content_raw"`
	//  null,  `json:"excluded"`
	Favorite      bool   `json:"favorite"`
	FavoriteCount int    `json:"favorite_count"`
	Gift          bool   `json:"has_gift"`
	Unread        int    `json:"is_unread"`
	Lang          string `json:"lang"`
	//  null,  `json:"last_edited"`
	Mentioned           int    `json:"mentioned"`
	NoComments          int    `json:"no_comments"`
	Owner               int    `json:"owner_id"`
	Type                int    `json:"plurk_type"`
	Porn                bool   `json:"porn"`
	Posted              Time   `json:"posted"`
	PublishToFollowers  bool   `json:"publish_to_followers"`
	Qualifier           string `json:"qualifier"`
	QualifierTranslated string `json:"qualifier_translated"`
	Replurkable         bool   `json:"replurkable"`
	Replurked           bool   `json:"replurked"`
	//  null,  `json:"replurker_id"`
	ReplurkersCount int `json:"replurkers_count"`
	Responded       int `json:"responded"`
	ResponseCount   int `json:"response_count"`
	UserID          int `json:"user_id"`
}
