package Group

type VkGroup struct {
	AccessToken string
	GroupID     string
}

func (group *VkGroup) NewGroup(accessToken string, groupId string) *VkGroup {
	return &VkGroup{
		AccessToken: accessToken,
		GroupID:     groupId,
	}
}
