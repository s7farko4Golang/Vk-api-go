package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// DeleteComment Удаляет сообщение темы в обсуждениях сообщества.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
// •ключ доступа сообщества
func (am *BoardMethods) DeleteComment(ctx context.Context, groupId uint, topicId uint, commentId string) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))
	params.Set("comment_id", commentId)

	VkRequest := types.VkRequest{
		Method: "board.deleteComment",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
