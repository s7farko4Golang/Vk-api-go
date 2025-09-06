package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetCatalogOptions struct {
	sort string //Способ сортировки приложений. Возможные значения:
	//• popular_today — популярные за день;
	//• visitors — по посещаемости;
	//• create_date — по дате создания приложения;
	//• growth_rate — по скорости роста;
	//• popular_week — популярные за неделю.
	//По умолчанию: popular_today.
	offset   uint   //Смещение, необходимое для выборки определенного подмножества приложений.
	count    uint   //Количество приложений, информацию о которых необходимо вернуть.
	platform string //Платформа, для которой необходимо вернуть приложения. Возможные значения:
	//• ios — iOS;
	//• android — Android;
	//• winphone — Windows Phone;
	//• web — приложения на vk.com;
	//• html5 — Direct Games.
	//По умолчанию: web.
	extended      bool   //1 — возвращать дополнительные поля приложений. Если указан extended – count не должен быть больше 100.
	returnFriends bool   //1 – возвращать список друзей, установивших это приложение. По умолчанию: 0.
	fields        string //Необязательный параметр. Список дополнительных полей, которые необходимо вернуть для профилей пользователей. Возможные значения:
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
	//Параметр учитывается только при return_friends = 1.
	nameCase string //Падеж для склонения имени и фамилии пользователей. Возможные значения:
	//• именительный – nom,
	//• родительный – gen,
	//• дательный – dat,
	//• винительный – acc,
	//• творительный – ins,
	//• предложный – abl.
	//По умолчанию: nom.
	//Параметр учитывается только при return_friends = 1.
	q       string //Поисковая строка для поиска по каталогу приложений.
	genreId uint   //Идентификатор жанра.
	filter  string //installed — возвращает список установленных приложений (только для мобильных приложений),
	// featured — возвращает список приложений, установленных в «Выбор редакции» (только для мобильных приложений).
}

type GetCatalogOption func(*GetCatalogOptions)

func GetCatalogWithSort(sort string) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.sort = sort
	}
}
func GetCatalogWithOffset(offset uint) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.offset = offset
	}
}
func GetCatalogWithCount(count uint) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.count = count
	}
}
func GetCatalogWithPlatform(platform string, extended bool) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.extended = extended
	}
}
func GetCatalogWithExtended(extended bool) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.extended = extended
	}
}
func GetCatalogWithReturnFriends(returnFriends bool) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.returnFriends = returnFriends
	}
}
func GetCatalogWithFields(fields string) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.fields = fields
	}
}
func GetCatalogWithNameCase(nameCase string) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.nameCase = nameCase
	}
}
func GetCatalogWithQ(q string) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.q = q
	}
}
func GetCatalogWithGenreId(genreId uint) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.genreId = genreId
	}
}
func GetCatalogWithFilter(filter string) GetCatalogOption {
	return func(o *GetCatalogOptions) {
		o.filter = filter
	}
}

// GetCatalog Возвращает список приложений, доступных для пользователей сайта через каталог приложений.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
// •сервисный ключ доступа
func (am *AppMethods) GetCatalog(ctx context.Context, count uint, opts ...GetCatalogOption) (types.VkResponse, error) {

	options := &GetCatalogOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("count", strconv.Itoa(int(count)))

	if options.sort != "" {
		params.Set("sort", options.sort)
	}
	if options.offset != 0 {
		params.Set("offset", strconv.Itoa(int(options.offset)))
	}
	if options.count != 0 {
		params.Set("count", strconv.Itoa(int(options.count)))
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
	if options.q != "" {
		params.Set("q", options.q)
	}
	if options.genreId != 0 {
		params.Set("genre_id", strconv.Itoa(int(options.genreId)))
	}
	if options.filter != "" {
		params.Set("filter", options.filter)
	}

	VkRequest := types.VkRequest{
		Method: "apps.getCatalog",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
