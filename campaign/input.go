package campaign

import "auth2/user"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	BackerCount      int    `json:"backer_count"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	User             user.User
}

type CreateCampaignImageInput struct {
	CampaignID int `form:"campaign_id" binding:"required"`
	IsPrimary  int `form:"is_primary" `
	User       user.User
}
