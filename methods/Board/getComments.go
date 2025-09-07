package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetCommentsOptions struct {
	groupId        uint //Идентификатор сообщества, информацию об обсуждениях которого нужно получить.
	topicId        uint //Идентификатор обсуждения.
	needLikes      bool //1 — будет возвращено дополнительное поле likes. По умолчанию поле likes не возвращается.
	startCommentId uint //Идентификатор комментария, начиная с которого нужно вернуть список (подробности см. ниже).
	offset         int  //Смещение, необходимое для выборки определенного подмножества сообщений.
	count          uint //Количество сообщений, которое необходимо получить (но не более 100). По умолчанию — 20.
	extended       bool //Если указать в качестве этого параметра 1, то будет возвращена информация о пользователях,
	// являющихся авторами сообщений. По умолчанию 0.
	sort string //Порядок сортировки комментариев:
	//• asc — хронологический;
	//• desc — антихронологический.
}

//Если был передан параметр start_comment_id, будет найдена позиция комментария в списке (или ближайший к нему более ранний).
//Начиная с этой позиции будет возвращено count комментариев. Смещение offset в этом случае будет отсчитываться от этой позиции
//(оно может быть отрицательным).

type GetCommentsOption func(*GetCommentsOptions)

func GetCommentsWithNeedLikes(needLikes bool) GetCommentsOption {
	return func(o *GetCommentsOptions) {
		o.needLikes = needLikes
	}
}
func GetCommentsWithStartCommentId(startCommentId uint) GetCommentsOption {
	return func(o *GetCommentsOptions) {
		o.startCommentId = startCommentId
	}
}
func GetCommentsWithOffset(offset int) GetCommentsOption {
	return func(o *GetCommentsOptions) {
		o.offset = offset
	}
}
func GetCommentsWithCount(count uint) GetCommentsOption {
	return func(o *GetCommentsOptions) {
		o.count = count
	}
}
func GetCommentsWithExtended(extended bool) GetCommentsOption {
	return func(o *GetCommentsOptions) {
		o.extended = extended
	}
}
func GetCommentsWithSort(sort string) GetCommentsOption {
	return func(o *GetCommentsOptions) {
		o.sort = sort
	}
}

// GetComments Возвращает список сообщений в указанной теме.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
// •сервисный ключ доступа
func (am *BoardMethods) GetComments(ctx context.Context, groupId, topicId uint, opts ...GetCommentsOption) (types.VkResponse, error) {

	options := &GetCommentsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))
	// Add parameters to values based on options
	if options.needLikes {
		params.Set("need_likes", "1")
	}
	if options.startCommentId > 0 {
		params.Set("start_comment_id", strconv.Itoa(int(options.startCommentId)))
	}
	if options.offset > 0 {
		params.Set("offset", strconv.Itoa(options.offset))
	}
	if options.count > 0 {
		params.Set("count", strconv.Itoa(int(options.count)))
	}
	if options.extended {
		params.Set("extended", "1")
	}
	if options.sort != "" {
		params.Set("sort", options.sort)
	}

	VkRequest := types.VkRequest{
		Method: "board.getComments",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
