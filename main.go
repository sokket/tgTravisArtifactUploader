package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: ./tcn filename")
		return
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	usersStr := os.Getenv("TG_BOT_USERS")
	users := strings.Split(usersStr, ":")
	ids := make([]int64, 0)
	for i := range users {
		id, err := strconv.Atoi(users[i])
		if err != nil {
			fmt.Println("Atoi err: " + users[i])
			continue
		}
		ids = append(ids, int64(id))
	}

	for i := range ids {
		file := tgbotapi.NewDocumentUpload(ids[i], os.Args[1])
		file.Caption =
			os.Getenv("TRAVIS_REPO_SLUG") + "/" + os.Getenv("TRAVIS_BRANCH") + "\n" +
				os.Getenv("TRAVIS_COMMIT") + "\n" +
				os.Getenv("TRAVIS_COMMIT_MESSAGE") + "\n" +
				os.Getenv("TRAVIS_BUILD_WEB_URL")
		message, err := bot.Send(file)
		if err != nil {
			fmt.Println("Can't send: ", message)
		}
	}

}
