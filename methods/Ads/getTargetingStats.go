package Ads

import (
	"Vk-api-go/types"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type GetTargetingStatsCriteria struct {
	Sex int `json:"sex"` //Пол. Возможные значения:
	//• 0 — любой;
	//• 1 — женский;
	//• 2 — мужской.
	AgeFrom  int `json:"age_from"` //Нижняя граница возраста (0 — не задано).
	AgeTo    int `json:"age_to"`   //Верхняя граница возраста (0 — не задано)
	Birthday int `json:"birthday"` //День рождения. Задаётся в виде суммы флагов:
	//• +1 — сегодня;
	//• +2 — завтра;
	//• +4 — в течение недели.
	Country   int    `json:"country"`    //Идентификатор страны (0 — не задано).
	Cities    string `json:"cities"`     //Идентификаторы городов и регионов. Идентификаторы регионов необходимо указывать со знаком «минус».
	CitiesNot string `json:"cities_not"` //Идентификаторы городов и регионов, которые следует исключить из таргетинга.
	// Идентификаторы регионов необходимо указывать со знаком «минус».
	GeoNear string `json:"geo_near"` // Места для гео-таргетинга. Список точек, разделённых «;».
	// Каждая точка задаётся строкой вида "<широта>,<долгота>,<радиус>[,название места]".
	// Широта и долгота задаются в десятичной записи градусов, радиус в метрах из списка допустимых:
	//• 500,
	//• 1000,
	//• 1500,
	//• 2000,
	//• 2500,
	//• 3000,
	//• 3500,
	//• 4000,
	//• 4500,
	//• 5000,
	//• 6000,
	//• 7000,
	//• 8000,
	//• 9000,
	//• 10000,
	//• 11000,
	//• 12000,
	//• 13000,
	//• 14000,
	//• 15000,
	//• 20000,
	//• 25000,
	//• 30000,
	//• 35000,
	//• 40000,
	//• 45000,
	//• 50000,
	//• 55000,
	//• 60000,
	//• 65000,
	//• 70000,
	//• 75000,
	//• 80000,
	//• 85000,
	//• 90000,
	//• 95000,
	//• 100000
	//Использование этого критерия несовместимо с критериями country, cities и cities_not.
	GeoPointType string `json:"geo_point_type"` //Тип мест гео-таргетинга. Возможные значения:
	//• regular — регулярно бывает;
	//• home — дом;
	//• work — работа или учёба;
	//• online — находится сейчас. Ограничивает максимальное значение радиуса в 5000м.
	//Используется только вместе с критерием geo_near.
	Statuses string `json:"statuses"` //Семейное положение (значения, перечисленные через запятую). Возможные значения:
	//• 1 — не женат или не замужем;
	//• 2 — есть подруга или есть друг;
	//• 3 — полмолвлен(а);
	//• 4 — женат или замужем;
	//• 5 — все сложно;
	//• 6 — в активном поиске;
	//• 7 — влюблен(а);
	//• 8 — в гражданском браке.
	Groups               string `json:"groups"`                 //Идентификаторы сообществ, разделенные запятой.
	GroupsNot            string `json:"groups_not"`             //Идентификаторы сообществ, разделенные запятой, которые следует исключить из таргетинга.
	Apps                 string `json:"apps"`                   //Идентификаторы приложений, разделенные запятой.
	AppsNot              string `json:"apps_not"`               //Идентификаторы приложений, разделенные запятой, которые следует исключить из таргетинга.
	Districts            string `json:"districts"`              //Идентификаторы районов, разделенные запятой.
	Stations             string `json:"stations"`               //Идентификаторы станций метро, разделенные запятой.
	Streets              string `json:"streets"`                //Идентификаторы улиц, разделенные запятой.
	Schools              string `json:"schools"`                //Идентификаторы учебных заведений.
	Positions            string `json:"positions"`              //Идентификаторы должностей.
	Religions            string `json:"religions"`              //Идентификаторы религиозных взглядов.
	InterestCategories   string `json:"interest_categories"`    //Категории интересов.
	Interests            string `json:"interests"`              //Интересы.
	UserDevices          string `json:"user_devices"`           //Устройства.
	UserOs               string `json:"user_os"`                // Операционные системы.
	UserBrowsers         string `json:"user_browsers"`          //Интернет-браузеры.
	RetargetingGroups    string `json:"retargeting_groups"`     //Идентификаторы групп ретаргетинга.
	RetargetingGroupsNot string `json:"retargeting_groups_not"` //Идентификаторы групп ретаргетинга, которые следует исключить из таргетинга.
	Paying               int    `json:"paying"`                 //Платежи. Возможные значения:
	//• 1 — не использовали голоса ВКонтакте;
	//• 2 — использовали голоса ВКонтакте.
	Travellers int `json:"travellers"`  //Только путешественники (1 — включить фильтр).
	SchoolFrom int `json:"school_from"` //Нижняя граница года окончания школы (0 — не задано).
	SchoolTo   int `json:"school_to"`   //Верхняя граница года окончания школы (0 — не задано).
	Uni_from   int `json:"uni_from"`    //Нижняя граница года окончания вуза (0 — не задано).
	UniTo      int `json:"uni_to"`      //Верхняя граница года окончания вуза (0 — не задано).
}

// GetTargetingStatsRequest структура для передачи в функцию
type GetTargetingStatsRequest struct {
	GetTargetingStatsCriteria []GetTargetingStatsCriteria `json:"criteria"`
}

// GetTargetingStatsSerialize сериализует массив GetTargetingStatsCriteria в JSON
func GetTargetingStatsSerialize(GetTargetingStatsCriteria []GetTargetingStatsCriteria) (string, error) {
	request := GetTargetingStatsRequest{
		GetTargetingStatsCriteria: GetTargetingStatsCriteria,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации: %w", err)
	}

	return string(jsonData), nil
}

type GetTargetingStatsOptions struct {
	accountId              int
	clientId               int
	criteria               string
	adId                   int
	adFormat               int
	adPlatform             string
	adPlatformNoWall       string
	adPlatformNoAdNetwork  string
	linkUrl                string
	linkDomain             string
	needPrecise            bool
	impressionsLimitPeriod int
}

type GetTargetingStatsOption func(*GetTargetingStatsOptions)

func GetTargetingStatsWithClientId(clientId int) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.clientId = clientId
	}
}

func GetTargetingStatsWithAccountId(accountId int) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.accountId = accountId
	}
}

func GetTargetingStatsWithAdId(adId int) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.adId = adId
	}
}

func GetTargetingStatsWithAdFormat(adFormat int) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.adFormat = adFormat
	}
}

func GetTargetingStatsWithAdPlatform(adPlatform string) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.adPlatform = adPlatform
	}
}

func GetTargetingStatsWithAdPlatformNoWall(adPlatformNoWall string) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.adPlatform = adPlatformNoWall
	}
}

func GetTargetingStatsWithAdPlatformNoAdNetwork(adPlatformNoAdNetwork string) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.adPlatformNoAdNetwork = adPlatformNoAdNetwork
	}
}

func GetTargetingStatsWithLinkDomain(linkDomain string) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.linkDomain = linkDomain
	}
}

func GetTargetingStatsWithNeedPrecise(needPrecise bool) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.needPrecise = needPrecise
	}
}

func GetTargetingStatsWithImpressionsLimitPeriod(impressionsLimitPeriod int) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.impressionsLimitPeriod = impressionsLimitPeriod
	}
}

func GetTargetingStatsWithCriteria(criteria string) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.criteria = criteria
	}
}

// GetTargetingStats Возвращает список пикселей ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetTargetingStats(ctx context.Context, linkUrl string, opts ...GetTargetingStatsOption) (types.VkResponse, error) {

	options := &GetTargetingStatsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("link_url", linkUrl)

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}
	if options.accountId != 0 {
		params.Set("account_id", strconv.Itoa(options.accountId))
	}
	if options.criteria != "" {
		params.Set("criteria", options.criteria)
	}
	if options.adId != 0 {
		params.Set("ad_id", strconv.Itoa(options.adId))
	}
	if options.adFormat != 0 {
		params.Set("ad_format", strconv.Itoa(options.adFormat))
	}
	if options.adPlatform != "" {
		params.Set("ad_platform", options.adPlatform)
	}
	if options.adPlatformNoWall != "" {
		params.Set("ad_platform_no_wall", options.adPlatformNoWall)
	}
	if options.adPlatformNoAdNetwork != "" {
		params.Set("ad_platform_no_ad_network", options.adPlatformNoAdNetwork)
	}
	if options.linkDomain != "" {
		params.Set("link_domain", options.linkDomain)
	}
	if options.needPrecise {
		params.Set("need_precise", "1")
	}
	if options.impressionsLimitPeriod != 0 {
		params.Set("impressions_limit_period", strconv.Itoa(options.impressionsLimitPeriod))
	}

	VkRequest := types.VkRequest{
		Method: "ads.getTargetingStats",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
