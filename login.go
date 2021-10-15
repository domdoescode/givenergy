package givenergy

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type LoginResponse struct {
	Success     bool   `json:"success"`
	MessageCode int    `json:"msgCode"`
	Role        string `json:"role"`
	Inverters   []struct {
		SerialNumber string `json:"serialNum"`
	} `json:"inverters"`
}

func Login(account, password string) (*LoginResponse, error) {
	var response LoginResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	jar, err := cookiejar.New(nil)
	if err != nil {
		return &response, err
	}

	client.Jar = jar

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("login"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("account", account)
	q.Add("password", password)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return &response, err
	}

	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response, err
	}

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}
