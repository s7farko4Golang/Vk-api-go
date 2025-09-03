package Ads

import "Vk-api-go/methods"

type AddMethods struct {
	methods *methods.APIMethods
}

func NewAccountMethods(m *methods.APIMethods) *AddMethods {
	return &AddMethods{methods: m}
}
