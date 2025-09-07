package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// OpenTopic Метод открывает ранее закрытое обсуждение. После этого в обсуждении снова можно оставлять сообщения.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
func (am *BoardMethods) OpenTopic(ctx context.Context, grouipId, topicId uint) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(grouipId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))

	VkRequest := types.VkRequest{
		Method: "board.openTopic",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
