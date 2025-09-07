package Board

import "Vk-api-go/methods"

type BoardMethods struct {
	methods *methods.APIMethods
}

func NewBoardMethods(m *methods.APIMethods) *BoardMethods {
	return &BoardMethods{methods: m}
}
