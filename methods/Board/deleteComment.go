package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// DeleteComment Удаляет сообщение темы в обсуждениях сообщества.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
// •ключ доступа сообщества
func (am *BoardMethods) DeleteComment(ctx context.Context, groupId string, topicId string, commentId string) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", groupId)
	params.Set("topic_id", topicId)
	params.Set("comment_id", commentId)

	VkRequest := types.VkRequest{
		Method: "board.deleteComment",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
