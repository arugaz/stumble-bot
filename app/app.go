package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

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
	if string(body) == "" {
		return nil, errors.New("check your auth")
	}
	if string(body) == "BANNED" {
		return nil, errors.New("your account got banned")
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func Run(auth *vars.Vars) {
	url := fmt.Sprintf(auth.Url, auth.Round)
	auths := auth.Auth
	wg := sync.WaitGroup{}

	for {
		wg.Add(1)
		go func(url, auths string) {
			defer wg.Done()
			resp, err := httpRequest(url, auths)
			if err != nil {
				fmt.Printf("%serrors: %s%v\n", vars.ColorRed, vars.ColorReset, err)
				return
			}
			data, err := decResponse(resp)
			if err != nil {
				fmt.Printf("%serrors: %s%v\n", vars.ColorRed, vars.ColorReset, err)
				return
			}
			fmt.Printf("%ssuccess: %s%v\n", vars.ColorGreen, vars.ColorReset, data.User.Crowns)
		}(url, auths)
		time.Sleep(50 * time.Millisecond)
	}
}
