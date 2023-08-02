package services

import (
	"log"
	"time"

	"github.com/vikashkumar2020/quigo-backend/app/models"
	"gorm.io/gorm"
)

func GetUnusedPrivateKey(db *gorm.DB, email string) (Address string, PrivateKey string) {
    
	var walletProvider models.Wallet;

    if err := db.Where("email IS NULL").First(&walletProvider).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
			log.Println("No unused key found")
            return "", "" // No unused key found
        }
		log.Println("Error occurred")
        return "", "" // Error occurred
    }
	
	// Update the email field
	walletProvider.Email = email
	walletProvider.UpdatedAt = time.Now()
	db.Save(&walletProvider)
	
    return walletProvider.Address, walletProvider.PrivateKey
}