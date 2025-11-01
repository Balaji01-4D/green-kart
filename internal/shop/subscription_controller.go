package shop

import (
	"net/http"
	"shop-near-u/internal/models"
	"shop-near-u/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) SubscribeFarmer(c *gin.Context) {
	// Get shop ID from URL parameter
	shopIDParam := c.Param("id")
	shopID, err := strconv.ParseUint(shopIDParam, 10, 32)
	if err != nil {
		utils.ErrorResponseSimple(c, 400, "invalid shop ID")
		return
	}

	user, exists := c.Get("user")
	if !exists {
		utils.ErrorResponseSimple(c, http.StatusUnauthorized, "unauthorized")
		c.Abort()
		return
	}

	u := user.(models.User)

	subscriberCount, err := ctrl.shopService.SubscribeFarmer(uint(u.ID), uint(shopID))
	if err != nil {
		if err.Error() == "already subscribed" {
			utils.ErrorResponseSimple(c, 400, "User is already subscribed to this shop")
			return
		}
		utils.ErrorResponseSimple(c, 500, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Subscribed successfully", SubscribeShopDTOResponse{
		Message:         "Subscribed successfully",
		SubscriberCount: subscriberCount,
	})
}

func (ctrl *Controller) UnsubscribeFarmer(c *gin.Context) {
	// Get shop ID from URL parameter
	shopIDParam := c.Param("id")
	shopID, err := strconv.ParseUint(shopIDParam, 10, 32)
	if err != nil {
		utils.ErrorResponseSimple(c, 400, "invalid shop ID")
		return
	}

	user, exists := c.Get("user")
	if !exists {
		utils.ErrorResponseSimple(c, http.StatusUnauthorized, "unauthorized")
		c.Abort()
		return
	}

	u := user.(models.User)

	subscriberCount, err := ctrl.shopService.UnsubscribeFarmer(uint(u.ID), uint(shopID))
	if err != nil {
		if err.Error() == "not subscribed" {
			utils.ErrorResponseSimple(c, 400, "User is not subscribed to this shop")
			return
		}
		utils.ErrorResponseSimple(c, 500, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Unsubscribed successfully", UnsubscribeShopDTOResponse{
		Message:         "Unsubscribed successfully",
		SubscriberCount: subscriberCount,
	})
}

func (ctrl *Controller) GetShopDetails(c *gin.Context) {
	shopIDParam := c.Param("id")
	shopID, err := strconv.ParseUint(shopIDParam, 10, 32)
	if err != nil {
		utils.ErrorResponseSimple(c, 400, "invalid shop ID")
		return
	}

	user, exists := c.Get("user")
	if !exists {
		utils.ErrorResponseSimple(c, http.StatusUnauthorized, "unauthorized")
		c.Abort()
		return
	}

	u := user.(models.User)

	shop, isSubscribed, err := ctrl.shopService.GetFarmerDetails(uint(shopID), u.ID)
	if err != nil {
		if err.Error() == "record not found" {
			utils.ErrorResponseSimple(c, 404, "farmer not found")
			return
		}
		utils.ErrorResponseSimple(c, 500, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Farmer details retrieved successfully", GetShopDetailsDTOResponse{
		ID:              shop.ID,
		Name:            shop.Name,
		SubscriberCount: shop.SubscriberCount,
		IsSubscribed:    isSubscribed,
	})
}

func (ctrl *Controller) GetUserSubscriptions(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.ErrorResponseSimple(c, http.StatusUnauthorized, "unauthorized")
		c.Abort()
		return
	}

	u := user.(models.User)

	shops, err := ctrl.shopService.GetUserSubscribedShops(uint(u.ID))
	if err != nil {
		utils.ErrorResponseSimple(c, 500, err.Error())
		return
	}

	// Convert to response DTOs
	var response []SubscribedShopDTOResponse
	for _, shop := range shops {
		response = append(response, SubscribedShopDTOResponse{
			ID:              shop.ID,
			Name:            shop.Name,
			Address:         shop.Address,
			Latitude:        shop.Latitude,
			Longitude:       shop.Longitude,
			SubscriberCount: shop.SubscriberCount,
		})
	}

	utils.SuccessResponse(c, http.StatusOK, "User subscriptions retrieved successfully", response)
}
