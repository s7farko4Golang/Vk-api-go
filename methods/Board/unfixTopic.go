package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// UnfixTopic Отменяет прикрепление темы в списке обсуждений группы (тема будет выводиться согласно выбранной сортировке).
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
func (am *BoardMethods) UnfixTopic(ctx context.Context, groupId, topicId uint) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))

	VkRequest := types.VkRequest{
		Method: "board.unfixTopic",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
