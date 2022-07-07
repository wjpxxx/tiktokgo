package tiktokgo

import (
	tiktokConfig "github.com/wjpxxx/tiktokgo/config"
	"testing"
	"fmt"
)

func TestConfig(t *testing.T) {
	ap:=NewApi(&tiktokConfig.Config{
		BaseURL:   "https://auth.tiktok-shops.com/",
		AppKey:    "65ad0lqi7v7al",
		AppSecret: "d2d657d1a8f45af0906566ee21baf3aedf336915",
	})
	r:=ap.AuthorizationURL("123123")
	fmt.Println(r)
}
