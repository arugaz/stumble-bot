package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ArugaZ/stumble-bot/types"
	"github.com/ArugaZ/stumble-bot/vars"
	"golang.org/x/net/context"
)

func Run(auth *vars.Vars) {
	urls := fmt.Sprintf(auth.Url, auth.Round)
	auths := auth.Auth
	for {
		data := new(types.StumbleResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// http request
		tr := &http.Transport{}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequestWithContext(ctx, "GET", urls, nil)
		if err != nil {
			fmt.Printf("%s[errors] %s%v\n", vars.ColorRed, vars.ColorReset, err)
			continue
		}
		req.Header.Set("Authorization", auths)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("%s[errors] %s%v\n", vars.ColorRed, vars.ColorReset, err)
			continue
		}
		defer resp.Body.Close()

		// result parser
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s[errors] %s%v\n", vars.ColorRed, vars.ColorReset, err)
			continue
		}
		html := string(body)
		if html == "" {
			continue
		}
		if strings.Contains(html, "BANNED") {
			fmt.Printf("%s[errors] %syour account got%v\n", vars.ColorRed, vars.ColorReset, err)
			break
		}
		if strings.Contains(html, "SERVER_ERROR") {
			continue
		}
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Printf("%s[errors] %s%v\n", vars.ColorRed, vars.ColorReset, err)
		}

		// output
		fmt.Printf("%s[success]%s Id:%s %d %s|%s Username:%s %s %s|%s Country:%s %s %s|%s Trophy:%s %d %s|%s Crown:%s %d\n", vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.ID, vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.Username, vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.Country, vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.HiddenRating, vars.ColorGreen, vars.ColorCyan, vars.ColorReset, data.User.Crowns)
		continue
	}
}
