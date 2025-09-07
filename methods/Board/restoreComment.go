package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// RestoreComment Метод восстанавливает в сообществе удалённое из обсуждения сообщение.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
// •ключ доступа сообщества
func (am *BoardMethods) RestoreComment(ctx context.Context, groupId, topicId, commentId uint) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))
	params.Set("comment_id", strconv.Itoa(int(commentId)))

	VkRequest := types.VkRequest{
		Method: "board.restoreComment",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
