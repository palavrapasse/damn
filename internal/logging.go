package internal

import (
	"os"

	as "github.com/palavrapasse/aspirador/pkg"
)

const (
	telegramBotTokenEnvKey = "telegram_token"
	telegramChatIdEnvKey   = "telegram_chat_id"
	logsFilePathEnvKey     = "logging_absolute_filepath"
)

var (
	telegramBotToken = os.Getenv(telegramBotTokenEnvKey)
	telegramChatId   = os.Getenv(telegramChatIdEnvKey)
	logsFilePath     = os.Getenv(logsFilePathEnvKey)
)

var Aspirador as.Aspirador

func init() {

	consoleClient := as.NewConsoleClient()

	telegramClient := as.NewTelegramClient(telegramBotToken, telegramChatId)

	clients := []as.Client{consoleClient, telegramClient}

	fileClient, err := as.NewFileClient(logsFilePath)

	if err == nil {
		clients = append(clients, fileClient)
	}

	Aspirador = as.WithClients(clients)
}
