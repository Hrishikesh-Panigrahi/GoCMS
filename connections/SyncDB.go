package connections

import (
	"fmt"
	"log"

	models "github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/joho/godotenv"
)

func SyncDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Role{}, &models.Image{})
	if err != nil {
		log.Fatal("Error migrating the database")
	}
	fmt.Println("Database Migrated")
}

// func seedData() {
// 	roles := []models.Role{{Name: "Admin", Description: "Admin Role"}, {Name: "Client", Description: "Client Role"}}
// 	user := []models.User{{Name: "blah", Email: "blah", Password: "blah", RoleID: 1}}
// 	DB.Save(&roles)
// 	DB.Save(&user)
// }
