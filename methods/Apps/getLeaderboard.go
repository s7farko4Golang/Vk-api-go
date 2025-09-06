package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type GetLeaderboardOptions struct {
	getLeaderboardType string //Тип турнирной таблицы. Возможные значения:
	//• level — по уровням;
	//• points — по баллам, начисленным за выполнение миссий;
	//• score — по очкам, начисленным напрямую (apps.getScore).
	global bool //Тип рейтинга. Возможные значения:
	//• 1 — глобальный рейтинг по всем игрокам (возвращается не более 20 результатов);
	//• 0 — рейтинг по друзьям пользователя.
	extended bool //1 — возвращать дополнительную информацию о пользователе.
}

type GetLeaderboardOption func(*GetLeaderboardOptions)

func GetLeaderboardWithGlobal(global bool) GetLeaderboardOption {
	return func(o *GetLeaderboardOptions) {
		o.global = true
	}
}
func GetLeaderboardWithExtended(global bool) GetLeaderboardOption {
	return func(o *GetLeaderboardOptions) {
		o.extended = true
	}
}

// GetLeaderboard Возвращает рейтинг пользователей в игре.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
func (am *AppMethods) GetLeaderboard(ctx context.Context, getLeaderboardType string, opts ...GetLeaderboardOption) (types.VkResponse, error) {

	options := &GetLeaderboardOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("type", getLeaderboardType)
	if options.global {
		params.Set("global", "1")
	}
	if options.extended {
		params.Set("extended", "1")
	}

	VkRequest := types.VkRequest{
		Method: "apps.getLeaderboard",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
