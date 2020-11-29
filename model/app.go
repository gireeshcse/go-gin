package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App :  App model
type App struct {
	gorm.Model
	Title       string
	Description string
}

// InitiateApp : Initiates app database
func InitiateApp() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&App{})

	// Create
	db.Create(&App{Title: "App", Description: "Some Info About App"})

	var app App
	db.First(&app, 1)                  // find App with integer primary key
	db.First(&app, "title = ?", "App") // find App with title App

	// Update - update app's description to Updated app info
	db.Model(&app).Update("Description", "Updated app info")
	// Update - update multiple fields
	db.Model(&app).Updates(App{Description: "New Description", Title: "New Title"}) // non-zero fields
	db.Model(&app).Updates(map[string]interface{}{"Description": "Description some", "Title": "F42"})

	// Delete - delete product
	db.Delete(&app, 1)

}
