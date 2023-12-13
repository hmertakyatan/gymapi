package configs

import (
	"context"
	"fmt"
	"log"

	"github.com/hmertakyatan/gymapi/ent"
	_ "github.com/lib/pq"
)

func ConnectionToDB(config *Config) *ent.Client {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUsername, config.DBName, config.DBPassword)
	fmt.Println("DSN:", dsn)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
