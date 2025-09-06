package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetScoreOptions struct {
	userId int //Идентификатор пользователя.
}

type GetScoreOption func(*GetScoreOptions)

func GetScoreWithUserId(userId int) GetScoreOption {
	return func(o *GetScoreOptions) {
		o.userId = userId
	}
}

// GetScore Метод возвращает количество очков пользователя в этой игре.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
func (am *AppMethods) GetScore(ctx context.Context, opts ...GetScoreOption) (types.VkResponse, error) {

	options := &GetScoreOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.userId != 0 {
		params.Set("user_id", strconv.Itoa(options.userId))
	}

	VkRequest := types.VkRequest{
		Method: "apps.getScore",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
