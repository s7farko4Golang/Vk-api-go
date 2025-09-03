package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// Unban Удаляет пользователя или группу из черного списка.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) Unban(ctx context.Context, ownerID int) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("owner_id", strconv.Itoa(ownerID))
	VkRequest := types.VkRequest{
		Method: "account.unban",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
