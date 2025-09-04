package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
	"strings"
)

type UpdateTargetGroupOptions struct {
	accountId     int    //Идентификатор рекламного кабинета.
	clientId      int    //Только для рекламных агентств. Id клиента, в рекламном кабинете которого будет создаваться аудитория.
	targetGroupId int    //Идентификатор аудитории.
	name          string //Название аудитории ретаргетинга — строка до 64 символов.
	domain        string //Домен сайта, на котором будет размещен код учета пользователей.
	//Устаревший параметр. Используется только для старых групп ретаргетинга,
	//которые пополнялись без помощи пикселя. Для новых аудиторий его следует опускать, иначе будет возвращена ошибка.
	lifetime         int      //Количество дней, через которое пользователи, добавляемые в аудиторию, будут автоматически исключены из неё.
	targetPixelId    int      //Идентификатор пикселя, если требуется собирать аудиторию с веб-сайта.
	targetPixelRules []string //Массив правил для пополнения аудитории из пикселя. Имеет вид:
	/*
		[
		{"type": args},
		{"type": args},
		...
		{"type": args}
		](
		{"type": args},
		{"type": args},
		...
		{"type": args}
		)
	*/
	//Подробнее в официальной документации Vk API https://dev.vk.com/ru/method/ads.UpdateTargetGroup
}
type UpdateTargetGroupOption func(*UpdateTargetGroupOptions)

func UpdateTargetGroupWithClientId(clientId int) UpdateTargetGroupOption {
	return func(o *UpdateTargetGroupOptions) {
		o.clientId = clientId
	}
}

func UpdateTargetGroupWithTargetPixelId(targetPixelId int) UpdateTargetGroupOption {
	return func(o *UpdateTargetGroupOptions) {
		o.targetPixelId = targetPixelId
	}
}

func UpdateTargetGroupWithTargetPixelRules(targetPixelRules []string) UpdateTargetGroupOption {
	return func(o *UpdateTargetGroupOptions) {
		o.targetPixelRules = targetPixelRules
	}
}

func UpdateTargetGroupWithDomain(domain string) UpdateTargetGroupOption {
	return func(o *UpdateTargetGroupOptions) {
		o.domain = domain
	}
}

func UpdateTargetGroupWithTargetGroupId(targetGroupId int) UpdateTargetGroupOption {
	return func(o *UpdateTargetGroupOptions) {
		o.targetGroupId = targetGroupId
	}
}

// UpdateTargetGroup Создаёт аудиторию для ретаргетинга рекламных объявлений на пользователей,
// которые посетили сайт рекламодателя (просмотрели информации о товаре, зарегистрировались и т.д.).
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) UpdateTargetGroup(ctx context.Context, accountID int, name string, lifetime int, opts ...UpdateTargetGroupOption) (types.VkResponse, error) {

	options := &UpdateTargetGroupOptions{
		accountId: accountID,
		name:      name,
		lifetime:  lifetime,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(options.accountId))
	params.Set("name", options.name)
	params.Set("lifetime", strconv.Itoa(options.lifetime))

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}

	if options.targetPixelId != 0 {
		params.Set("target_pixel_id", strconv.Itoa(options.targetPixelId))
	}

	if options.targetPixelRules != nil {
		params.Set("target_pixel_rules", strings.Join(options.targetPixelRules, ","))
	}
	if options.domain != "" {
		params.Set("domain", options.domain)
	}
	if options.targetGroupId != 0 {
		params.Set("target_group_id", strconv.Itoa(options.targetGroupId))
	}

	VkRequest := types.VkRequest{
		Method: "ads.updateTargetGroup",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
