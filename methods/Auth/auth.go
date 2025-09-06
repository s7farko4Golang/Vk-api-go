package Auth

import "Vk-api-go/methods"

type AuthMethods struct {
	methods *methods.APIMethods
}

func NewAuthMethods(m *methods.APIMethods) *AuthMethods {
	return &AuthMethods{methods: m}
}
