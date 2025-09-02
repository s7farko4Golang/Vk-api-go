package account_methods

import (
	"Vk-api-go/methods"
)

type AccountMethods struct {
	methods *methods.APIMethods
}

func NewAccountMethods(m *methods.APIMethods) *AccountMethods {
	return &AccountMethods{methods: m}
}
