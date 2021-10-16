package config

type AuthLog struct {
	Date         string `json:"date"`
	Server       string `json:"server"`
	IdentityName string `json:"identity_name"`
	Desc         string `json:"desc"`
}
