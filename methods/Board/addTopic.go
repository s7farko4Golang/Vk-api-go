package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type AddTopicOptions struct {
	groupId     string //Идентификатор сообщества.
	title       string //Название обсуждения.
	text        string //Текст первого сообщения в обсуждении.
	fromGroup   bool   //1 — тема будет создана от имени группы, 0 — тема будет создана от имени пользователя (по умолчанию).
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

type AddTopicOption func(*AddTopicOptions)

func AddTopicWithText(text string) AddTopicOption {
	return func(o *AddTopicOptions) {
		o.text = text
	}
}
func AddTopicWithFromGroup(fromGroup bool) AddTopicOption {
	return func(o *AddTopicOptions) {
		o.fromGroup = fromGroup
	}
}
func AddTopicWithAttachments(attachments string) AddTopicOption {
	return func(o *AddTopicOptions) {
		o.attachments = attachments
	}
}

// AddTopic Создает новую тему в списке обсуждений группы.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученн
func (am *BoardMethods) AddTopic(ctx context.Context, groupId string, title string, opts ...AddTopicOption) (types.VkResponse, error) {

	options := &AddTopicOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("group_id", groupId)
	params.Set("title", title)
	// Add parameters to values based on options
	if options.text != "" {
		params.Set("text", options.text)
	}
	if options.fromGroup {
		params.Set("from_group", "1")
	}
	if options.attachments != "" {
		params.Set("attachments", options.attachments)
	}

	VkRequest := types.VkRequest{
		Method: "board.addTopic",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
