package givenergy

import (
	"fmt"
	"net/http"
)

var (
	baseURL = "https://www.givenergy.cloud/GivManage/api/"

	client = http.DefaultClient

	cookieJar http.CookieJar
)

func getURL(url string) string {
	return fmt.Sprintf("%s%s", baseURL, url)
}
