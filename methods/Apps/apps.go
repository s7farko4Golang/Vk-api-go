package Apps

import "Vk-api-go/methods"

type AppMethods struct {
	methods *methods.APIMethods
}

func NewAppMethods(m *methods.APIMethods) *AppMethods {
	return &AppMethods{methods: m}
}
