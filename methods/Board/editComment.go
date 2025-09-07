package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type EditCommentOptions struct {
	groupId     uint   //Идентификатор сообщества, в котором размещено обсуждение.
	topicId     uint   //Идентификатор обсуждения.
	commentId   string //Идентификатор комментария в обсуждении.
	message     string //Новый текст комментария (является обязательным, если не задан параметр attachments).
	attachments string //Список объектов, приложенных к записи, в формате:
	//<type><owner_id>_<media_id>,<type><owner_id>_<media_id>
	//<type> — тип медиа-приложения:
	//• photo — фотография;
	//• video — видеозапись ;
	//• audio — аудиозапись;
	//• doc — документ.
	//<owner_id> — идентификатор владельца медиа-приложения. <media_id> — идентификатор медиа-приложения.
	//Например:
	//photo100172_166443618,photo66748_265827614
	//Параметр является обязательным, если не задан параметр text.
}

type EditCommentOption func(*EditCommentOptions)

func EditCommentWithMessage(message string) EditCommentOption {
	return func(o *EditCommentOptions) {
		o.message = message
	}
}
func EditCommentWithAttachments(attachments string) EditCommentOption {
	return func(o *EditCommentOptions) {
		o.attachments = attachments
	}
}

// EditComment Редактирует одно из сообщений в обсуждении сообщества.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
func (am *BoardMethods) EditComment(ctx context.Context, groupId uint, topicId uint, commentId uint, opts ...EditCommentOption) (types.VkResponse, error) {

	options := &EditCommentOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))
	params.Set("comment_id", strconv.Itoa(int(commentId)))

	if options.message != "" {
		params.Set("message", options.message)
	}
	if options.attachments != "" {
		params.Set("attachments", options.attachments)
	}

	VkRequest := types.VkRequest{
		Method: "board.editComment",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
