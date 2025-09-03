package Ads

import (
	"Vk-api-go/types"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// CreateCampaignsSpecification описывает структуру администратора
type CreateCampaignsSpecification struct {
	client_id   int    `json:"client_id"` //Только для рекламных агентств. id клиента, в рекламном кабинете которого будет создаваться кампания.
	companyType string `json:"type"`      //Тип кампании:
	//normal — обычная кампания, в которой можно создавать любые объявления, кроме описанных в следующих пунктах;
	//promoted_posts — кампания, в которой можно рекламировать только записи в сообществе;
	//adaptive_ads — кампания, в которой можно рекламировать только объявления адаптивного формата.
	name       string `json:"name"`       //Название рекламной кампании. Строка длиной от 3 до 60 символов.
	day_limit  uint   `json:"day_limit"`  //Дневной лимит в рублях. Положительное число.
	all_limit  uint   `json:"all_limit"`  //Общий лимит в рублях. Положительное число.
	start_time uint   `json:"start_time"` //Время запуска кампании в формате unixtime. Положительное число.
	stop_time  uint   `json:"stop_time"`  //Время остановки кампании в формате unixtime. Положительное число.
	status     bool   `json:"status"`     //Статус рекламной кампании (1 — запущена, 0 — остановлена).
}

// CreateCampaignsRequest структура для передачи в функцию
type CreateCampaignsRequest struct {
	CreateCampaignsSpecification []CreateCampaignsSpecification `json:"campaign_specification"`
}

// CreateCampaignsSerialize сериализует массив CreateCampaignsSpecification в JSON
func CreateCampaignsSerialize(CreateCampaignsSpecification []CreateCampaignsSpecification) (string, error) {
	request := CreateCampaignsRequest{
		CreateCampaignsSpecification: CreateCampaignsSpecification,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации: %w", err)
	}

	return string(jsonData), nil
}

// CreateCampaigns Создает рекламные кампании.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) CreateCampaigns(ctx context.Context, accountID int, data string) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("data", data)
	VkRequest := types.VkRequest{
		Method: "ads.createCampaigns",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
