package cmds

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)


func Chn(Serverid int, Token string) {

	payload := map[string]interface{}{
		"name": "real-ordem",
		"type": 0,
	}

	client := resty.New()
	url := fmt.Sprintf("https://discord.com/api/v9/guilds/%d/channels", Serverid)
	for i := 0; i <= 12; i++ {
		resp, err := client.R().
			SetHeader("Authorization", Token).
			SetHeader("Content-Type", "application/json").
			SetBody(payload).
			Post(url)

		if err != nil {
			fmt.Println("\033[33m[*]\033[0m error")
		}

		status := resp.StatusCode()

		if status == 200 {
			fmt.Println("[\033[33m+\033[0m] | [200]")
		} else if status == 204 {
			fmt.Println("[\033[33m+\033[0m] | [204]")
		} else if status == 404 {
			fmt.Println("[\033[33m-\033[0m] | [404] | Not found")
		} else {
			msg := fmt.Sprintf("[\033[33m-\033[0m] | [%d]", status)
			fmt.Println(msg)
		}

	}
} 
