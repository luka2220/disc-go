package models

type GithubUser struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Url       string `json:"url"`
	Blog      string `json:"blog"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Followers int64  `json:"followers"`
}
