package handler

import (
	"auth2/campaign"
	"auth2/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// tangkap parameter di handler
// handler ke service
// service yg menentukan repository mana yg di call
// repository findall, findbyid
// db
type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", campaign.FormatCampaigns(campaigns))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}
