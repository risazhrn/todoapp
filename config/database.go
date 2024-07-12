package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "log"
    "os"
    "todoapp/models"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    dsn := "host=" + os.Getenv("DB_HOST") + 
           " user=" + os.Getenv("DB_USER") + 
           " dbname=" + os.Getenv("DB_NAME") + 
           " sslmode=" + os.Getenv("DB_SSLMODE") + 
           " password=" + os.Getenv("DB_PASSWORD")
    DB, err = gorm.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Database connected successfully")
    DB.AutoMigrate(&models.ToDo{})
}
