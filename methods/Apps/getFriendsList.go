package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetFriendsListOptions struct {
	extended bool //Параметр, определяющий необходимость возвращать расширенную информацию о пользователях. Возможные значения:
	//• 0 — возвращаются только идентификаторы;
	//• 1 — будут дополнительно возвращены имя и фамилия.
	//По умолчанию: 0.
	count              uint   //Количество пользователей в создаваемом списке.
	offset             uint   //Смещение относительно первого пользователя для выборки определенного подмножества.
	getFriendsListType string //Тип создаваемого списка друзей. Возможные значения:
	//• invite — доступные для приглашения (не играют в игру);
	//• request — доступные для отправки запроса (уже играют).
	//По умолчанию: invite.
	fields string //Список дополнительных полей профилей, которые необходимо вернуть.
}

type GetFriendsListOption func(*GetFriendsListOptions)

func GetFriendsListWithExtended(extended bool) GetFriendsListOption {
	return func(o *GetFriendsListOptions) {
		o.extended = extended
	}
}
func GetFriendsListWithCount(count uint) GetFriendsListOption {
	return func(o *GetFriendsListOptions) {
		o.count = count
	}
}
func GetFriendsListWithOffset(offset uint) GetFriendsListOption {
	return func(o *GetFriendsListOptions) {
		o.offset = offset
	}
}
func GetFriendsListWithGetFriendsListType(getFriendsListType string) GetFriendsListOption {
	return func(o *GetFriendsListOptions) {
		o.getFriendsListType = getFriendsListType
	}
}
func GetFriendsListWithFields(fields string) GetFriendsListOption {
	return func(o *GetFriendsListOptions) {
		o.fields = fields
	}
}

// GetFriendsList Создает список друзей, который будет использоваться при отправке пользователем приглашений в приложение и игровых запросов.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
func (am *AppMethods) GetFriendsList(ctx context.Context, opts ...GetFriendsListOption) (types.VkResponse, error) {

	options := &GetFriendsListOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	// Add parameters to values based on options
	if options.extended {
		params.Set("extended", "1")
	}
	if options.count != 0 {
		params.Set("count", strconv.Itoa(int(options.count)))
	}
	if options.offset != 0 {
		params.Set("offset", strconv.Itoa(int(options.offset)))
	}
	if options.getFriendsListType != "" {
		params.Set("type", options.getFriendsListType)
	}
	if options.fields != "" {
		params.Set("fields", options.fields)
	}

	VkRequest := types.VkRequest{
		Method: "apps.getFriendsList",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
