package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"os"
	"log"
	"regexp"
	"strconv"
	"math/big"
	"fmt"
)

var pattern = regexp.MustCompile(`(\d+)\!`)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Args[1])

	if err != nil {
		log.Println(err)
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		txt := update.Message.Text
		a := pattern.FindStringSubmatch(txt)

		if len(a) < 2 {
			continue
		}

		val, err := strconv.Atoi(a[1])
		if err != nil {
			continue
		}

		input := int64(val)
		if input > 100000 {
			input = 100000
		}

		x := new(big.Int)
		go func() {
			res := new(big.Float)
			factorial := x.MulRange(1, input)
			res.SetInt(factorial)
			m := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%d! = %s", input, res.Text(byte('e'), 2)))
			bot.Send(m)
		}()
	}
}
