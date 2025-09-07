package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type CreateCommentOptions struct {
	groupId     uint   //Идентификатор сообщества, в котором находится обсуждение.
	topicId     uint   //Идентификатор темы, в которой необходимо оставить комментарий.
	message     string //Текст комментария. Обязательный параметр, если не передано значение attachments.
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
	fromGroup bool   //1 — сообщение будет опубликовано от имени группы, 0 — сообщение будет опубликовано от имени пользователя (по умолчанию).
	stickerId uint   //Идентификатор стикера.
	guid      string //Уникальный идентификатор, предназначенный для предотвращения повторной отправки одинакового комментария.
}

type CreateCommentOption func(*CreateCommentOptions)

func CreateCommentWithMessage(message string) CreateCommentOption {
	return func(o *CreateCommentOptions) {
		o.message = message
	}
}
func CreateCommentWithAttachments(attachments string) CreateCommentOption {
	return func(o *CreateCommentOptions) {
		o.attachments = attachments
	}
}
func CreateCommentWithFromGroup(fromGroup bool) CreateCommentOption {
	return func(o *CreateCommentOptions) {
		o.fromGroup = fromGroup
	}
}
func CreateCommentWithSticker(stickerId uint) CreateCommentOption {
	return func(o *CreateCommentOptions) {
		o.stickerId = stickerId
	}
}
func CreateCommentWithGuid(guid string) CreateCommentOption {
	return func(o *CreateCommentOptions) {
		o.guid = guid
	}
}

// CreateComment Добавляет новый комментарий в обсуждении.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
func (am *BoardMethods) CreateComment(ctx context.Context, groupID uint, topicId uint, opts ...CreateCommentOption) (types.VkResponse, error) {

	options := &CreateCommentOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("group_id", strconv.FormatUint(uint64(groupID), 10))
	params.Set("topic_id", strconv.FormatUint(uint64(topicId), 10))

	if options.message != "" {
		params.Set("message", options.message)
	}
	if options.attachments != "" {
		params.Set("attachments", options.attachments)
	}
	if options.fromGroup {
		params.Set("from_group", "1")
	}
	if options.stickerId != 0 {
		params.Set("sticker_id", strconv.FormatUint(uint64(options.stickerId), 10))
	}
	if options.guid != "" {
		params.Set("guid", options.guid)
	}

	VkRequest := types.VkRequest{
		Method: "board.createComment",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
