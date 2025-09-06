package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type AddSnippetOptions struct {
	vkRef string //Необязательный параметр. Область применения сниппета. Массив, который содержит одну из следующих строк или обе сразу:
	//• snippet_post — сниппет используется в записи на стене пользователя или в записи сообщества.
	//• snippet_im — сниппет используется в личном сообщении.
	groupId string //Необязательный параметр. Список идентификаторов сообществ, в которых может быть отображён сниппет.
	hash    string //Необязательный параметр. Список хешей для запуска мини-приложений или игр, для которых будет отрисован сниппет.
	// Хеш — подстрока в ссылке после символа #. Сниппет будет отображён, если ссылка содержит хеш, который указан в массиве.
	//Важно. Значения чувствительны к регистру. В массиве можно указать значения с маской — символом *.
	//Маска означает любое количество символов или пустую подстроку. Например, значение join* соответствует следующим ссылкам:
	//https://vk.com/app123#join_invite https://vk.com/app123#join_msg https://vk.com/app123#join
	//Элемент массива без * соответствует ссылкам, в которых хеш точно совпадает со значением элемента.
	//Например, join соответствует ссылке https://vk.com/app123#join, но не https://vk.com/app123#join_invite.
	//Максимальная длина — 100 символов.
	snippetId     int    //Необязательный параметр. Идентификатор сниппета.
	title         string //Необязательный параметр. В настоящее время не используется. Зарезервирован для применения в будущем.
	description   string //Необязательный параметр. Описание сниппета. Максимальная длина: 80 символов.
	imageUrl      string //Необязательный параметр. URL изображения сниппета. Размер изображения: 1120×630 px.
	smallImageUrl string //Необязательный параметр. URL изображения сниппета. Размер изображения: 150×150 px.
	button        string //Необязательный параметр. Текст кнопки сниппета. Возможные значения:
	//• open — «Открыть».
	//• buy — «Купить».
	//• buy_ticket — «Купить билет».
	//• enroll — «Записаться».
	//• contact — «Связаться».
	//• fill — «Заполнить».
	//• go — «Перейти».
	//• play — «Играть».
	//• help — «Помочь».
	//• create — «Создать».
}

type AddSnippetOption func(*AddSnippetOptions)

func AddSnippetWithVkRef(vkRef string) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.vkRef = vkRef
	}
}

func AddSnippetWithGroupId(groupId string) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.groupId = groupId
	}
}

func AddSnippetWithHash(hash string) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.hash = hash
	}
}

func AddSnippetWithSnippetId(snippetId int) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.snippetId = snippetId
	}
}

func AddSnippetWithTitle(title string) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.title = title
	}
}

func AddSnippetWithDescription(description string) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.description = description
	}
}

func AddSnippetWithImageUrl(imageUrl string) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.imageUrl = imageUrl
	}
}

func AddSnippetWithSmallImageUrl(smallImageUrl string) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.smallImageUrl = smallImageUrl
	}
}

func AddSnippetWithButton(button string) AddSnippetOption {
	return func(o *AddSnippetOptions) {
		o.button = button
	}
}

// AddSnippet Метод добавляет новый сниппет в коллекцию сниппетов мини-приложения или игры.
// Сниппеты можно создавать для мини-приложений и игр, опубликованных в каталоге. Для каждого мини-приложения или игры можно создать до 10 сниппетов.
// Время жизни сниппета не ограничено. Он хранится до тех пор, пока вы его не удалите.
// Для вызова метода можно использовать:
// •сервисный ключ доступа
func (am *AppMethods) AddSnippet(ctx context.Context, opts ...AddSnippetOption) (types.VkResponse, error) {

	options := &AddSnippetOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.vkRef != "" {
		params.Set("vk_ref", options.vkRef)
	}
	if options.groupId != "" {
		params.Set("group_id", options.groupId)
	}
	if options.hash != "" {
		params.Set("hash", options.hash)
	}
	if options.snippetId != 0 {
		params.Set("snippet_id", strconv.Itoa(options.snippetId))
	}
	if options.title != "" {
		params.Set("title", options.title)
	}
	if options.description != "" {
		params.Set("description", options.description)
	}
	if options.imageUrl != "" {
		params.Set("image_url", options.imageUrl)
	}
	if options.smallImageUrl != "" {
		params.Set("small_image_url", options.smallImageUrl)
	}
	if options.button != "" {
		params.Set("button", options.button)
	}

	VkRequest := types.VkRequest{
		Method: "apps.addSnippet",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
