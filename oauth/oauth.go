package oauth

import (
	"fmt"

	tiktokConfig "github.com/wjpxxx/tiktokgo/config"
)

//OAuth
type OAuth struct {
	Config *tiktokConfig.Config
}

//AuthorizationURL
func (a *OAuth) AuthorizationURL(state string) string {
	return fmt.Sprintf("%s/oauth/authorize?app_key=%s&state=%s", a.Config.BaseURL, a.Config.AppKey, state)
}
