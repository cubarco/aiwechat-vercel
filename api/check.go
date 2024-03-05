package api

import (
	"fmt"
	"github.com/cubarco/aiwechat-vercel/config"
	"net/http"
)

func Check(w http.ResponseWriter, req *http.Request) {
	botType, err := config.CheckBotConfig()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "BOT [%v] config check passed", botType)
}
