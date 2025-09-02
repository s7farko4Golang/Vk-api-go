package config

import (
	"Vk-api-go/Group"
	"Vk-api-go/account"
	"Vk-api-go/types"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadConfigFromEnv загружает конфигурацию VK из переменных среды.
// Возвращает структуру VkConfig с заполненными данными.
func LoadConfigFromEnv() (*types.VkConfig, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден. Будут использованы системные переменные окружения.")
	}
	config := &types.VkConfig{}

	config.PrimaryAccount = account.VkAccount{
		AccessToken: os.Getenv("VK_PRIMARY_ACCOUNT_TOKEN"),
		UserID:      os.Getenv("VK_PRIMARY_ACCOUNT_ID"),
	}

	config.PrimaryGroup = Group.VkGroup{
		AccessToken: os.Getenv("VK_PRIMARY_GROUP_TOKEN"),
		GroupID:     os.Getenv("VK_PRIMARY_GROUP_ID"),
	}

	config.SecondaryAccounts = loadSecondaryAccounts()

	config.SecondaryGroups = loadSecondaryGroups()

	return config, nil
}

// loadSecondaryAccounts загружает дополнительные аккаунты из переменных среды.
func loadSecondaryAccounts() []account.VkAccount {
	var accounts []account.VkAccount

	for i := 1; i <= 20; i++ {
		prefix := "VK_ACCOUNT_" + strconv.Itoa(i)
		token := os.Getenv(prefix + "_TOKEN")

		if token == "" {
			break
		}

		account := account.VkAccount{
			AccessToken: token,
			UserID:      os.Getenv(prefix + "_ID"),
		}

		accounts = append(accounts, account)
	}

	return accounts
}

// loadSecondaryGroups загружает дополнительные группы из переменных среды.
func loadSecondaryGroups() []Group.VkGroup {
	var groups []Group.VkGroup
	for i := 1; i <= 20; i++ {
		prefix := "VK_GROUP_" + strconv.Itoa(i)
		token := os.Getenv(prefix + "_TOKEN")

		if token == "" {
			break
		}

		group := Group.VkGroup{
			AccessToken: token,
			GroupID:     os.Getenv(prefix + "_ID"),
		}

		groups = append(groups, group)
	}

	return groups
}
