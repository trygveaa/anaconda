package anaconda

type Tweet struct {
	Source                    string
	Id                        int64
	Retweeted                 bool
	Favorited                 bool
	User                      TwitterUser
	Truncated                 bool
	Text                      string
	Retweet_count             int64
	Id_str                    string
	Created_at                string
	Entities                  TwitterEntities
	In_reply_to_status_id     *int64
	In_reply_to_status_id_str *string
}
