package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// EditTopic Метод изменяет заголовок обсуждения в сообществе.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
func (am *BoardMethods) EditTopic(ctx context.Context, groupId uint, topicId uint, title string) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))
	params.Set("title", title)

	VkRequest := types.VkRequest{
		Method: "board.editTopic",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
