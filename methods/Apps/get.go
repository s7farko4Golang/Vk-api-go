package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetOptions struct {
	appId uint //Необязательный параметр. Идентификатор приложения, данные которого необходимо получить. Если этот параметр
	// и параметр appIds не указаны, возвращается идентификатор приложения, через которое выдан ключ доступа (access_token).
	appIds   string //Необязательный параметр. Список идентификаторов приложений, данные которых необходимо получить, перечисленный через запятую.
	platform string //Необязательный параметр. Платформа, для которой необходимо вернуть данные. Возможные значения:
	//• ios — мобильное приложение для iOS.
	//• android — мобильное приложение для Android.
	//• winphone — мобильное приложение для Windows Phone.
	//• web — десктопная версия сайта (vk.com). Значение используется по умолчанию.
	extended bool //Необязательный параметр. Информация о том, вернуть ли дополнительные поля. Возможные значения:
	//• 1 — вернуть дополнительные поля.
	//• 0 — не возвращать дополнительные поля. Значение используется по умолчанию.
	returnFriends bool //Примечание. Параметр учитывается, только если передан параметр access_token.
	//Необязательный параметр. Информация о том, вернуть ли список друзей, установивших это приложение. Возможные значения:
	//• 1 – вернуть список друзей.
	//• 0 — не возвращать список друзей. Значение используется по умолчанию.
	fields string //Примечание. Параметр учитывается, только если параметр returnFriends передан со значением 1.
	//Необязательный параметр. Список дополнительных полей, которые необходимо вернуть для профилей пользователей. Возможные значения:
	//• bdate
	//• can_post
	//• can_see_all_posts
	//• can_see_audio
	//• can_write_private_message
	//• city
	//• common_count
	//• connections
	//• contacts
	//• counters
	//• country
	//• domain
	//• education
	//• has_mobile
	//• last_seen
	//• lists
	//• online
	//• online_mobile
	//• photo_100
	//• photo_200
	//• photo_200_orig
	//• photo_400_orig
	//• photo_50
	//• photo_max
	//• photo_max_orig
	//• relation
	//• relatives
	//• schools
	//• screen_name
	//• sex
	//• site
	//• status
	//• timezone
	//• universities
	nameCase string //Примечание. Параметр учитывается, только если параметр returnFriends передан со значением 1.
	//Необязательный параметр. Падеж для склонения имени и фамилии пользователей. Возможные значения:
	//• nom — именительный. Значение используется по умолчанию.
	//• gen — родительный.
	//• dat — дательный.
	//• acc — винительный.
	//• ins — творительный.
	//• abl — предложный.
	appFields string //Необязательный параметр. Список названий полей из метода apps.get, которые должны вернуться. Поля id, type, title вернутся всегда.
}

type GetOption func(*GetOptions)

func GetWithAppId(appId uint) GetOption {
	return func(o *GetOptions) {
		o.appId = appId
	}
}
func GetWithAppIds(appIds string) GetOption {
	return func(o *GetOptions) {
		o.appIds = appIds
	}
}
func GetWithPlatform(platform string) GetOption {
	return func(o *GetOptions) {
		o.platform = platform
	}
}
func GetWithExtended(extended bool) GetOption {
	return func(o *GetOptions) {
		o.extended = extended
	}
}
func GetWithReturnFriends(returnFriends bool) GetOption {
	return func(o *GetOptions) {
		o.returnFriends = returnFriends
	}
}
func GetWithFields(fields string) GetOption {
	return func(o *GetOptions) {
		o.fields = fields
	}
}
func GetWithNameCase(name string) GetOption {
	return func(o *GetOptions) {
		o.nameCase = name
	}
}
func GetWithAppFields(appFields string) GetOption {
	return func(o *GetOptions) {
		o.appFields = appFields
	}
}

// Get Метод возвращает данные о приложениях.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
// •сервисный ключ доступа
func (am *AppMethods) Get(ctx context.Context, opts ...GetOption) (types.VkResponse, error) {

	options := &GetOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.appId != 0 {
		params.Set("app_id", strconv.FormatUint(uint64(options.appId), 10))
	}
	if options.appIds != "" {
		params.Set("app_ids", options.appIds)
	}
	if options.platform != "" {
		params.Set("platform", options.platform)
	}
	if options.extended {
		params.Set("extended", "1")
	}
	if options.returnFriends {
		params.Set("return_friends", "1")
	}
	if options.fields != "" {
		params.Set("fields", options.fields)
	}
	if options.nameCase != "" {
		params.Set("name_case", options.nameCase)
	}
	if options.appFields != "" {
		params.Set("app_fields", options.appFields)
	}

	VkRequest := types.VkRequest{
		Method: "apps.get",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
