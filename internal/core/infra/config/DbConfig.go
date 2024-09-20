// /internal/core/infra/config/DbConection.go
package config

import (
	"fmt"
	UserEntity "go-api-template/internal/modules/Users/domain/entitys"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Exporta la variable Db
var Db *gorm.DB

func DbConection() error {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error al cargar archivo .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	// Conectar a la base de datos
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// Crear la extensión uuid-ossp
	if err := createUUIDExtension(Db); err != nil {
		return fmt.Errorf("error creating UUID extension: %v", err)
	}

	// Migrar el esquema
	if err := Db.AutoMigrate(&UserEntity.User{}); err != nil {
		return fmt.Errorf("error migrating the schema: %v", err)
	}

	log.Println("Database migrated successfully")
	return nil
}

func createUUIDExtension(db *gorm.DB) error {
	// Ejecuta el comando SQL para crear la extensión
	err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		log.Printf("Error creating UUID extension: %v", err)
		return err
	}
	log.Println("UUID extension created successfully")
	return nil
}
