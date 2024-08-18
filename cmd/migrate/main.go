package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/DBrange/didis-comp-bk/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Usar la conexión existente
	client := database.MongoClient

	// DB Name
	dbName := database.Database
	// Configurar el driver de MongoDB para las migraciones
	driver, err := mongodb.WithInstance(client, &mongodb.Config{
		DatabaseName: dbName,
	})
	if err != nil {
		log.Fatal("Error creating MongoDB driver:", err)
	}

	// Crear una nueva instancia de migración
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // Ruta a tus archivos de migración
		dbName,
		driver,
	)
	if err != nil {
		log.Fatal("Error creating migrate instance:", err)
	}

	// Obtener el comando (up o down)
	cmd := os.Args[len(os.Args)-1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Error applying migrations:", err)
		}
		log.Println("Migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Error reverting migrations:", err)
		}
		log.Println("Migrations reverted successfully")
	default:
		log.Fatal("Invalid command. Use 'up' or 'down'")
	}
}
