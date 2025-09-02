package types

import (
	"Vk-api-go/Group"
	"Vk-api-go/account"
	"net/url"
)

// VkResponse Структура, хранящая в себе ответы от Vk API
type VkResponse struct {
	Response interface{} `json:"response"`
	Error    *VkError    `json:"error,omitempty"`
}

// VkRequest Структура для запросов к VK API
type VkRequest struct {
	Method string
	Params url.Values
}

// VkError Структура, хранящая ошибки, полученные от VK API
type VkError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

// VkConfig представляет конфигурацию для работы с VK API.
// Структура предназначена для хранения учетных данных и идентификаторов
// пользователей и групп, участвующих в работе приложения.
type VkConfig struct {
	// PrimaryAccount представляет основной аккаунт пользователя.
	// Используется как аккаунт по умолчанию для операций, требующих
	// авторизации, или как управляющий аккаунт в сценариях с множеством аккаунтов.
	PrimaryAccount account.VkAccount `json:"primary_account"`

	// SecondaryAccounts содержит дополнительные аккаунты пользователей.
	// Используется для сценариев, требующих одновременной работы с несколькими аккаунтами,
	// таких как массовые операции или управление несколькими профилями.
	SecondaryAccounts []account.VkAccount `json:"secondary_accounts"`

	// PrimaryGroup представляет основную группу/сообщество.
	// Используется как группа по умолчанию для публикации контента
	// или управления сообществом.
	PrimaryGroup Group.VkGroup `json:"primary_group"`

	// SecondaryGroups содержит дополнительные группы/сообщества.
	// Используется для управления несколькими сообществами или
	// кросс-постинга между различными группами.
	SecondaryGroups []Group.VkGroup `json:"secondary_groups"`
}
