package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetTopicsOptions struct {
	groupId uint //Идентификатор сообщества, информацию об обсуждениях которого необходимо получить.
	// Если сообщество закрытое или частное, для вызова метода потребуется право доступа groups.
	topicId int //Список идентификаторов тем, которые необходимо получить (не более 100).
	// По умолчанию возвращаются все темы. Если указан данный параметр, игнорируются параметры order, offset и count
	// (возвращаются все запрошенные темы в указанном порядке).
	order int // Порядок, в котором необходимо вернуть список тем. Возможные значения:
	//• 1 — по убыванию даты обновления;
	//• 2 — по убыванию даты создания;
	//• -1 — по возрастанию даты обновления;
	//• -2 — по возрастанию даты создания.
	// По умолчанию темы возвращаются в порядке, установленном администратором группы.
	// «Прилепленные» темы при любой сортировке возвращаются первыми в списке.
	offset   uint //Смещение, необходимое для выборки определенного подмножества тем.
	count    uint //Количество тем, которое необходимо получить (но не более 100). По умолчанию — 40.
	extended bool //Если указать в качестве этого параметра 1, то будет возвращена информация о пользователях,
	// являющихся создателями тем или оставившими в них последнее сообщение. По умолчанию 0.
	preview int //Набор флагов, определяющий,
	// необходимо ли вернуть вместе с информацией о темах текст первых и последних сообщений в них. Является суммой флагов:
	//• 1 — вернуть первое сообщение в каждой теме (поле first_comment);
	//• 2 — вернуть последнее сообщение в каждой теме (поле last_comment). По умолчанию — 0 (не возвращать текст сообщений)`
	previewLength uint //Количество символов, по которому нужно обрезать первое и последнее сообщение.
	// Укажите 0, если Вы не хотите обрезать сообщение. (по умолчанию — 90).
}

type GetTopicsOption func(*GetTopicsOptions)

func GetTopicsWithTopicId(topicId int) GetTopicsOption {
	return func(o *GetTopicsOptions) {
		o.topicId = topicId
	}
}
func GetTopicsWithOrder(order int) GetTopicsOption {
	return func(o *GetTopicsOptions) {
		o.order = order
	}
}
func GetTopicsWithOffset(offset uint) GetTopicsOption {
	return func(o *GetTopicsOptions) {
		o.offset = offset
	}
}
func GetTopicsWithCount(count uint) GetTopicsOption {
	return func(o *GetTopicsOptions) {
		o.count = count
	}
}
func GetTopicsWithExtended(extended bool) GetTopicsOption {
	return func(o *GetTopicsOptions) {
		o.extended = extended
	}
}
func GetTopicsWithPreview(preview int) GetTopicsOption {
	return func(o *GetTopicsOptions) {
		o.preview = preview
	}
}
func GetTopicsWithPreviewLength(previewLength uint) GetTopicsOption {
	return func(o *GetTopicsOptions) {
		o.previewLength = previewLength
	}
}

// GetTopics Возвращает список тем в обсуждениях указанной группы.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
// •сервисный ключ доступа
func (am *BoardMethods) GetTopics(ctx context.Context, groupId uint, opts ...GetTopicsOption) (types.VkResponse, error) {

	options := &GetTopicsOptions{
		previewLength: 90,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("preview_length", strconv.Itoa(int(options.previewLength)))
	if options.topicId != 0 {
		params.Set("topic_ids", strconv.Itoa(options.topicId))
	}
	if options.order != 0 {
		params.Set("order", strconv.Itoa(options.order))
	}
	if options.offset != 0 {
		params.Set("offset", strconv.Itoa(int(options.offset)))
	}
	if options.count != 0 {
		params.Set("count", strconv.Itoa(int(options.count)))
	}
	if options.extended {
		params.Set("extended", "1")
	}
	if options.preview != 0 {
		params.Set("preview", strconv.Itoa(int(options.preview)))
	}
	if options.previewLength != 90 {
		params.Set("preview_length", strconv.Itoa(int(options.previewLength)))
	}

	VkRequest := types.VkRequest{
		Method: "board.getTopics",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
