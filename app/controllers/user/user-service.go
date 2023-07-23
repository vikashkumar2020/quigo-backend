package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/models"
)

func GetMe() gin.HandlerFunc{
	return func(ctx *gin.Context){
	
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Phone: currentUser.Phone,
		Role: currentUser.Role,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
}
