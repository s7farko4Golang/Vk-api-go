package Ads

import (
	"Vk-api-go/types"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// AddOfficeUserSpecification описывает структуру администратора
type AddOfficeUserSpecification struct {
	UserID int `json:"user_id"` //Обязательный параметр, int (числовое значение). Идентификатор пользователя,
	// добавляемого как администратор или наблюдатель.
	Role string `json:"role"` //Обязательный параметр, строка. Флаг, описывающий тип полномочий:
	//reports — наблюдатель;
	//manager — администратор.
	ClientID   int `json:"client_id"`             //Обязательный параметр, int (числовое значение). Идентификатор клиента.
	ViewBudget int `json:"view_budget,omitempty"` // Показывать ли бюджет пользователю. Необязательный параметр, [1, 0].
}

// AddOfficeAddAdminsRequest структура для передачи в функцию
type AddOfficeUserRequest struct {
	AddOfficeUserSpecification []AddOfficeUserSpecification `json:"user_specification"`
}

// AddOfficeUsersSerialize сериализует массив AddOfficeUserSpecification в JSON
func AddOfficeUsersSerialize(AddOfficeUserSpecification []AddOfficeUserSpecification) (string, error) {
	request := AddOfficeUserRequest{
		AddOfficeUserSpecification: AddOfficeUserSpecification,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации: %w", err)
	}

	return string(jsonData), nil
}

// AddOfficeUsers Добавляет администраторов и/или наблюдателей в рекламный кабинет.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) AddOfficeUsers(ctx context.Context, accountID int, data string) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("data", data)
	VkRequest := types.VkRequest{
		Method: "ads.addOfficeUsers",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
