package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
	"strings"
)

type CreateTargetGroupOptions struct {
	accountId        int      //Идентификатор рекламного кабинета.
	clientId         int      //Только для рекламных агентств. Id клиента, в рекламном кабинете которого будет создаваться аудитория.
	name             string   //Название аудитории ретаргетинга — строка до 64 символов.
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
	//Подробнее в официальной документации Vk API https://dev.vk.com/ru/method/ads.createTargetGroup
}
type CreateTargetGroupOption func(*CreateTargetGroupOptions)

func CreateTargetGroupWithClientId(clientId int) CreateTargetGroupOption {
	return func(o *CreateTargetGroupOptions) {
		o.clientId = clientId
	}
}

func CreateTargetGroupWithTargetPixelId(targetPixelId int) CreateTargetGroupOption {
	return func(o *CreateTargetGroupOptions) {
		o.targetPixelId = targetPixelId
	}
}

func CreateTargetGroupWithTargetPixelRules(targetPixelRules []string) CreateTargetGroupOption {
	return func(o *CreateTargetGroupOptions) {
		o.targetPixelRules = targetPixelRules
	}
}

// CreateTargetGroup Создаёт аудиторию для ретаргетинга рекламных объявлений на пользователей,
// которые посетили сайт рекламодателя (просмотрели информации о товаре, зарегистрировались и т.д.).
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) CreateTargetGroup(ctx context.Context, accountID int, name string, lifetime int, opts ...CreateTargetGroupOption) (types.VkResponse, error) {

	options := &CreateTargetGroupOptions{
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

	VkRequest := types.VkRequest{
		Method: "ads.createTargetGroup",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
