package cmds

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Channel struct {
	Id string `json:"id"`
}

func clear() {
	fmt.Print("\033[2J\033[H")
}


func Die(Serverid int, Token string) {
	client := resty.New()
	chns := []Channel{}

	url := fmt.Sprintf("https://discord.com/api/v9/guilds/%d/channels", Serverid)

	
	resp, err := client.R().
	SetHeader("Authorization", Token).
	SetResult(&chns).
	Get(url)		

	if err != nil || resp == nil {
		fmt.Println("\033[33m[!]\033[0m | erro na requisição. | > ", err)
	}

	fmt.Println("[/] A pegar canais.")
	time.Sleep(2)

	for _, chn := range chns {
		delurl := fmt.Sprintf("https://discord.com/api/v9/channels/%s", chn.Id)

		resp, err := client.R().
			SetHeader("Authorization", Token).
			Delete(delurl)

			if err != nil || resp == nil {
				fmt.Println("\033[33m[-]\033[0m erro: ", err)
				time.Sleep(1)
			}

			status := resp.StatusCode()

			switch status {
			case 200:
				fmt.Println("\033[33m[+]\033[0m canal [", chn.Id ,"] apagado com sucesso.")
			case 204:
					fmt.Println("\033[33m[+]\033[0m canal [", chn.Id ,"] apagado com sucesso.")
			case 202:
				fmt.Println("\033[33m[+]\033[0m canal [", chn.Id ,"] apagado com sucesso.")
			case 401:
				fmt.Println("\033[33m[-]\033[0m token nao esta valido")
			case 403:
				fmt.Println("\033[33m[-]\033[0m sem permissao para apagar o canal")
			case 429:
				fmt.Println("\033[33m[-]\033[0m limite de requests excedido")
			}
	}
	fmt.Printf("\033[33m[/]\033[0m a sair")
	time.Sleep(2)
}
