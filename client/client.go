package client

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Client struct {
	BaseUrl    string
	HttpClient *http.Client
	Language   string
	logger     slog.Logger
}

func NewClient() *Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Client{
		BaseUrl:    os.Getenv("VK_CLIENT_BASE_URL"),
		HttpClient: http.DefaultClient,
		Language:   os.Getenv("VK_LANG"),
	}
}
