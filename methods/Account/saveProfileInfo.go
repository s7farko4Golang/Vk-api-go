package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type SaveProfileInfoOptions struct {
	firstName       string //Имя пользователя. Обязательно с большой буквы.
	lastName        string //Фамилия пользователя. Обязательно с большой буквы.
	maidenName      string //Девичья фамилия пользователя (только для женского пола).
	screenName      string //Короткое имя страницы.
	cancelRequestId uint   //Идентификатор заявки на смену имени, которую необходимо отменить.
	// Если передан этот параметр, все остальные параметры игнорируются.
	sex      uint //Пол пользователя. Возможные значения: 1 — женский; 2 — мужской.
	relation uint //Семейное положение пользователя. Возможные значения:
	//1 — не женат/не замужем;
	//2 — есть друг/есть подруга;
	//3 — помолвлен/помолвлена;
	//4 — женат/замужем;
	//5 — всё сложно;
	//6 — в активном поиске;
	//7 — влюблён/влюблена;
	//8 — в гражданском браке;
	//0 — не указано.
	relationPartnerId int    //Идентификатор пользователя, с которым связано семейное положение.
	bdate             string //Дата рождения пользователя в формате DD.MM.YYYY, например 15.11.1984.
	bdateVisibility   uint   //Видимость даты рождения. Возможные значения:
	//1 — показывать дату рождения;
	//2 — показывать только месяц и день;
	//0 — не показывать дату рождения.
	homeTown  string //Родной город пользователя.
	countryId uint   //Идентификатор страны пользователя.
	cityId    uint   //Идентификатор города пользователя.
	status    string //Статус пользователя, который также может быть изменен методом status.set.
}
type SaveProfileInfoOption func(*SaveProfileInfoOptions)

func SaveProfileInfoWithFirstName(firstName string) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.firstName = firstName
	}
}

func SaveProfileInfoWithLastName(lastName string) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.lastName = lastName
	}
}

func SaveProfileInfoWithMaidenName(maidenName string) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.maidenName = maidenName
	}
}

func SaveProfileInfoWithScreenName(screenName string) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.screenName = screenName
	}
}

func SaveProfileInfoWithCancelRequestId(cancelRequestId uint) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.cancelRequestId = cancelRequestId
	}
}

func SaveProfileInfoWithSex(sex uint) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.sex = sex
	}
}

func SaveProfileInfoWithRelation(relation uint, relationPartnerId int) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.relation = relation
	}
}

func SaveProfileInfoWithRelationPartnerId(relationPartnerId int) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.relationPartnerId = relationPartnerId
	}
}

func SaveProfileInfoWithBdate(bdate string) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.bdate = bdate
	}
}

func SaveProfileInfoWithBdateVisibility(bdateVisibility uint) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.bdateVisibility = bdateVisibility
	}
}

func SaveProfileInfoWithHomeTown(homeTown string) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.homeTown = homeTown
	}
}

func SaveProfileInfoWithCountryId(countryId uint) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.countryId = countryId
	}
}

func SaveProfileInfoWithCityId(cityId uint) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.cityId = cityId
	}
}

func SaveProfileInfoWithStatus(status string) SaveProfileInfoOption {
	return func(o *SaveProfileInfoOptions) {
		o.status = status
	}
}

// SaveProfileInfo Редактирует информацию текущего профиля.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) SaveProfileInfo(ctx context.Context, opts ...SaveProfileInfoOption) (types.VkResponse, error) {

	options := &SaveProfileInfoOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.firstName != "" {
		params.Set("first_name", options.firstName)
	}

	if options.lastName != "" {
		params.Set("last_name", options.lastName)
	}

	if options.maidenName != "" {
		params.Set("maiden_name", options.maidenName)
	}

	if options.screenName != "" {
		params.Set("screen_name", options.screenName)
	}

	if options.cancelRequestId != 0 {
		params.Set("cancel_request_id", strconv.FormatUint(uint64(options.cancelRequestId), 10))
	}

	if options.sex != 0 {
		params.Set("sex", strconv.FormatUint(uint64(options.sex), 10))
	}

	if options.relation != 0 {
		params.Set("relation", strconv.FormatUint(uint64(options.relation), 10))
	}

	if options.relationPartnerId != 0 {
		params.Set("relation_partner_id", strconv.FormatUint(uint64(options.relationPartnerId), 10))
	}

	if options.bdate != "" {
		params.Set("bdate", options.bdate)
	}

	if options.bdateVisibility != 0 {
		params.Set("bdate_visibility", strconv.FormatUint(uint64(options.bdateVisibility), 10))
	}

	if options.homeTown != "" {
		params.Set("home_town", options.homeTown)
	}

	if options.countryId != 0 {
		params.Set("country_id", strconv.FormatUint(uint64(options.countryId), 10))
	}

	if options.cityId != 0 {
		params.Set("city_id", strconv.FormatUint(uint64(options.cityId), 10))
	}

	if options.status != "" {
		params.Set("status", options.status)
	}

	VkRequest := types.VkRequest{
		Method: "account.saveProfileInfo",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
