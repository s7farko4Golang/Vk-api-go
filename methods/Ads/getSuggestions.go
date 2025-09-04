package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetSuggestionsOptions struct {
	section string //Раздел, по которому запрашиваются подсказки. Поддерживаются следующие значения:
	//• countries — страны.
	//	• q — необязательный параметр;
	//		• 1 — возвращается полный список стран;
	//		• не задан — возвращается краткий список стран.
	//• regions — регионы.
	//	• country — обязательный параметр;
	//	• q — обязательный параметр.
	//• cities — города.
	//	• country — обязательный параметр;
	//	• q — необязательный параметр.
	//• districts — районы.
	//	• cities — обязательный параметр.
	//• stations — станции метро.
	//	• cities — обязательный параметр.
	//• streets — улицы.
	//	• cities — обязательный параметр,
	//	• q — обязательный параметр.
	//• schools — учебные заведения: школы, университеты, факультеты, кафедры.
	//	• cities — обязательный параметр;
	//	• q — обязательный параметр.
	//• interest_categories_v2 — категории интересов.
	//• interest_categories — устаревшие категории интересов.
	//• positions — должности (профессии).
	//	• q — необязательный параметр.
	//• religions — религиозные взгляды.
	//• user_devices — устройства.
	//• user_os — операционные системы.
	//• user_browsers — интернет-браузеры.
	ids string //ID объектов, разделённые запятыми. Служит для расшифровки ID, возвращаемых в методе ads.getAdsTargeting.
	// Если задан этот параметр, то параметры q, country, cities не должны передаваться, таким образом отменяется их
	//обязательность для конкретного раздела. Объекты возвращаются в том же порядке, в каком они были заданы в этом параметре.
	q       string //Строка-фильтр запроса.
	country int    //Id страны, в которой ищутся объекты.
	cities  string //Id городов, в которых ищутся объекты, разделенные запятыми.
	lang    string //Язык возвращаемых строковых значений. Поддерживаемые языки:
	//• ru — русский;
	//• ua — украинский;
	//• en — английский.
}

type GetSuggestionsOption func(*GetSuggestionsOptions)

func GetSuggestionsWithIds(ids string) GetSuggestionsOption {
	return func(o *GetSuggestionsOptions) {
		o.ids = ids
	}
}

func GetSuggestionsWithQ(q string) GetSuggestionsOption {
	return func(o *GetSuggestionsOptions) {
		o.q = q
	}
}

func GetSuggestionsWithCountry(country int) GetSuggestionsOption {
	return func(o *GetSuggestionsOptions) {
		o.country = country
	}
}

func GetSuggestionsWithCities(cities string) GetSuggestionsOption {
	return func(o *GetSuggestionsOptions) {
		o.cities = cities
	}
}

func GetSuggestionsWithLang(lang string) GetSuggestionsOption {
	return func(o *GetSuggestionsOptions) {
		o.lang = lang
	}
}

// GetSuggestions Возвращает список аудиторий ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetSuggestions(ctx context.Context, section string, opts ...GetSuggestionsOption) (types.VkResponse, error) {

	options := &GetSuggestionsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("section", section)

	if options.ids != "" {
		params.Set("ids", options.ids)
	}
	if options.q != "" {
		params.Set("q", options.q)
	}
	if options.country != 0 {
		params.Set("country", strconv.Itoa(options.country))
	}
	if options.cities != "" {
		params.Set("cities", options.cities)
	}
	if options.lang != "" {
		params.Set("lang", options.lang)
	}

	VkRequest := types.VkRequest{
		Method: "ads.getSuggestions",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
