package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/models"
	pgdatabase "github.com/vikashkumar2020/quigo-backend/infra/postgres/database"
)

func GetMe() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		currentUser := ctx.MustGet("currentUser").(models.User)

		userResponse := &models.UserResponse{
			ID:        currentUser.ID,
			Name:      currentUser.Name,
			Email:     currentUser.Email,
			Phone:     currentUser.Phone,
			Role:      currentUser.Role,
			CreatedAt: currentUser.CreatedAt,
			UpdatedAt: currentUser.UpdatedAt,
		}

		db := pgdatabase.GetDBInstance().GetDB()
		var balance models.Wallet
		result := db.Where("email = ?", currentUser.Email).First(&balance)

		if result.Error != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
			return
		}
		
		userResponse.Balance = strconv.FormatInt(balance.Balance, 10)
	
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": userResponse})
	}
}
