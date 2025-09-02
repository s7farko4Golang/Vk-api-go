package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

func (am *AccountMethods) Ban(ctx context.Context, ownerID interface{}) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("owner_id", ownerID.(string))
	VkRequest := types.VkRequest{
		Method: "account.ban",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
