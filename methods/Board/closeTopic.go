package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// CloseTopic Метод закрывает обсуждение в сообществе.
// Примечание. В закрытом обсуждении невозможно оставить новые сообщения.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
func (am *BoardMethods) CloseTopic(ctx context.Context, groupId string, topicId string) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", groupId)
	params.Set("topic_id", topicId)

	VkRequest := types.VkRequest{
		Method: "board.closeTopic",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
