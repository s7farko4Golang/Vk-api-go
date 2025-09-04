package Ads

import (
	"Vk-api-go/types"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// UpdateCampaignsSpecification описывает структуру администратора
type UpdateCampaignsSpecification struct {
	ClientId    int    `json:"client_id"` //Только для рекламных агентств. id клиента, в рекламном кабинете которого будет создаваться кампания.
	CompanyType string `json:"type"`      //Тип кампании:
	//normal — обычная кампания, в которой можно создавать любые объявления, кроме описанных в следующих пунктах;
	//promoted_posts — кампания, в которой можно рекламировать только записи в сообществе;
	//adaptive_ads — кампания, в которой можно рекламировать только объявления адаптивного формата.
	Name      string `json:"name"`       //Название рекламной кампании. Строка длиной от 3 до 60 символов.
	DayLimit  uint   `json:"day_limit"`  //Дневной лимит в рублях. Положительное число.
	AllLimit  uint   `json:"all_limit"`  //Общий лимит в рублях. Положительное число.
	StartTime uint   `json:"start_time"` //Время запуска кампании в формате unixtime. Положительное число.
	StopTime  uint   `json:"stop_time"`  //Время остановки кампании в формате unixtime. Положительное число.
	Status    bool   `json:"status"`     //Статус рекламной кампании (1 — запущена, 0 — остановлена).
}

// UpdateCampaignsRequest структура для передачи в функцию
type UpdateCampaignsRequest struct {
	UpdateCampaignsSpecification []UpdateCampaignsSpecification `json:"campaign_specification"`
}

// UpdateCampaignsSerialize сериализует массив UpdateCampaignsSpecification в JSON
func UpdateCampaignsSerialize(UpdateCampaignsSpecification []UpdateCampaignsSpecification) (string, error) {
	request := UpdateCampaignsRequest{
		UpdateCampaignsSpecification: UpdateCampaignsSpecification,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации: %w", err)
	}

	return string(jsonData), nil
}

// UpdateCampaigns Создает рекламные кампании.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) UpdateCampaigns(ctx context.Context, accountID int, data string) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("data", data)
	VkRequest := types.VkRequest{
		Method: "ads.updateCampaigns",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
