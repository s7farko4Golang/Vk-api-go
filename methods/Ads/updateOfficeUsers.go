package Ads

import (
	"Vk-api-go/types"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// UpdateOfficeUserSpecification описывает структуру администратора
type UpdateOfficeUserSpecification struct {
	UserID int `json:"user_id"` //Обязательный параметр, int (числовое значение). Идентификатор пользователя,
	// добавляемого как администратор или наблюдатель.
	Role string `json:"role"` //Обязательный параметр, строка. Флаг, описывающий тип полномочий:
	//reports — наблюдатель;
	//manager — администратор.
	ClientID                int  `json:"client_id"`                   //Обязательный параметр, int (числовое значение). Идентификатор клиента.
	ViewBudget              int  `json:"view_budget,omitempty"`       // Показывать ли бюджет пользователю. Необязательный параметр, [1, 0].
	GrantAccessToAllClients bool `json:"grant_access_to_all_clients"` //(true или false) Доступ ко всем текущим и новым клиентам этого кабинета.
}

// UpdateOfficeUserRequest структура для передачи в функцию
type UpdateOfficeUserRequest struct {
	UpdateOfficeUserSpecification []UpdateOfficeUserSpecification `json:"user_specification"`
}

// UpdateOfficeUsersSerialize сериализует массив UpdateOfficeUserSpecification в JSON
func UpdateOfficeUsersSerialize(UpdateOfficeUserSpecification []UpdateOfficeUserSpecification) (string, error) {
	request := UpdateOfficeUserRequest{
		UpdateOfficeUserSpecification: UpdateOfficeUserSpecification,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации: %w", err)
	}

	return string(jsonData), nil
}

// UpdateOfficeUsers Добавляет администраторов и/или наблюдателей в рекламный кабинет.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) UpdateOfficeUsers(ctx context.Context, accountID int, data string) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("data", data)
	VkRequest := types.VkRequest{
		Method: "ads.updateOfficeUsers",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
