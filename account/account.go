package account

type VkAccount struct {
	AccessToken string
	UserID      string
}

func NewVkAccount(accessToken string, userID string) *VkAccount {
	return &VkAccount{
		AccessToken: accessToken,
		UserID:      userID,
	}
}
