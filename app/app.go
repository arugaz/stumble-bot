package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ArugaZ/stumble-bot/types"
	"github.com/ArugaZ/stumble-bot/vars"
)

func httpRequest(url string, Auth string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", Auth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func decResponse(resp *http.Response) (*types.StumbleResponse, error) {
	data := new(types.StumbleResponse)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	html := string(body)
	if html == "" {
		return nil, errors.New("check your auth")
	}
	if strings.Contains(html, "BANNED") {
		return nil, errors.New("your account got banned")
	}
	if strings.Contains(html, "SERVER_ERROR") {
		return nil, errors.New("server error")
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func Run(auth *vars.Vars) {
	url := fmt.Sprintf(auth.Url, auth.Round)
	auths := auth.Auth

	for {
		resp, err := httpRequest(url, auths)
		if err != nil {
			fmt.Printf("%s[errors] %s%v\n", vars.ColorRed, vars.ColorReset, err)
			continue
		}
		data, err := decResponse(resp)
		if err != nil {
			fmt.Printf("%s[errors] %s%v\n", vars.ColorRed, vars.ColorReset, err)
			continue
		}
		fmt.Printf("%s[success]%s Id:%s %d %s|%s Username:%s %s %s|%s Country:%s %s %s|%s Trophy:%s %d %s|%s Crown:%s %d\n", vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.ID, vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.Username, vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.Country, vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.HiddenRating, vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.Crowns)
	}
}
