package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type GetChangePasswordOptions struct {
	restoreSid string //Идентификатор сессии, полученный при восстановлении доступа используя метод auth.restore.
	// (В случае если пароль меняется сразу после восстановления доступа).
	changePasswordHash string //Хэш, полученный при успешной OAuth авторизации по коду полученному по СМС
	// (В случае если пароль меняется сразу после восстановления доступа).
	oldPassword string //Текущий пароль пользователя.
	newPassword string //Новый пароль, который будет установлен в качестве текущего.
}
type GetChangePasswordOption func(*GetChangePasswordOptions)

func ChangePasswordWithRestoreSid(restoreSid string) GetChangePasswordOption {
	return func(o *GetChangePasswordOptions) {
		o.restoreSid = restoreSid
	}
}

func ChangePasswordWithChangePasswordHash(changePasswordHash string) GetChangePasswordOption {
	return func(o *GetChangePasswordOptions) {
		o.changePasswordHash = changePasswordHash
	}
}

func ChangePasswordWithOldPassword(oldPassword string) GetChangePasswordOption {
	return func(o *GetChangePasswordOptions) {
		o.oldPassword = oldPassword
	}
}

// ChangePassword Позволяет сменить пароль пользователя после успешного восстановления доступа к аккаунту через СМС,
// используя метод auth.restore.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) ChangePassword(ctx context.Context, newPassword string, opts ...GetChangePasswordOption) (types.VkResponse, error) {

	options := &GetChangePasswordOptions{
		newPassword: newPassword,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("new_password", options.newPassword)

	if options.restoreSid != "" {
		params.Set("restore_sid", options.restoreSid)
	}

	if options.changePasswordHash != "" {
		params.Set("change_password_hash", options.changePasswordHash)
	}

	if options.oldPassword != "" {
		params.Set("old_password", options.oldPassword)
	}

	VkRequest := types.VkRequest{
		Method: "account.changePassword",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
