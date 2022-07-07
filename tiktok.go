package tiktokgo

import (
	tiktokConfig "github.com/wjpxxx/tiktokgo/config"
	"github.com/wjpxxx/tiktokgo/oauth"
)

//TikToker
type TikToker interface {
	AuthorizationURL(state string) string
}

//TikTok
type TikTok struct {
	oauth.OAuth
}

//NewApi
func NewApi(cfg *tiktokConfig.Config) TikToker {
	return &TikTok{
		oauth.OAuth{Config: cfg},
	}
}
