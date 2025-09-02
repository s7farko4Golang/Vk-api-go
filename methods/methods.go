package methods

import (
	"Vk-api-go/account"
	"Vk-api-go/client"
	"Vk-api-go/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// APIMethods предоставляет базовый функционал для вызовов API
type APIMethods struct {
	client  *client.Client
	account *account.VkAccount
}

func NewAPIMethods(client *client.Client, acc *account.VkAccount) *APIMethods {
	return &APIMethods{client: client, account: acc}
}

func (m *APIMethods) Call(ctx context.Context, request types.VkRequest) (types.VkResponse, error) {
	VkResponse := types.VkResponse{}

	baseURL := fmt.Sprintf("%s%s", m.client.BaseUrl, request.Method)

	params := request.Params
	fmt.Print(params)
	// Добавляем обязательные параметры
	params.Set("access_token", m.account.AccessToken)
	params.Set("v", "5.199") // версия API
	params.Set("lang", m.client.Language)

	// Создаем полный URL с параметрами
	fullURL := baseURL + "?" + params.Encode()

	fmt.Printf("URL %s: %s\n", request.Method, fullURL)

	// Создаем HTTP запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return types.VkResponse{}, err
	}

	response, err := m.client.HttpClient.Do(req)
	if err != nil {
		return types.VkResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return types.VkResponse{}, err
	}

	err = json.Unmarshal(body, &VkResponse)
	if err != nil {
		return types.VkResponse{}, err
	}

	if VkResponse.Error != nil {
		return types.VkResponse{}, fmt.Errorf("VK API error: %v", VkResponse.Error)
	}

	return VkResponse, nil
}
