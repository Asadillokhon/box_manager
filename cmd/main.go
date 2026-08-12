package main

import (
	"box-manager/cmd/api"
	"box-manager/cmd/store/pgstore"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	// migrate create -ext sql -dir cmd/migrations -seq title_table

	ctx := context.Background()
	dsn := os.Getenv("CONN_STRING")
	pgStore, err := pgstore.NewPostgresStore(ctx, dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer pgStore.Close()

	httpHandlers := api.NewHTTPHandler(pgStore)
	httpServer := api.NewHTTPServer(httpHandlers)

	fmt.Println("Сервер запущен на :9091")
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
	// memory := store.NewMemoryStore()
	// httpHandlers := api.NewHTTPHandler(memory)
	// httpServer := api.NewHTTPServer(httpHandlers)

	// if err := httpServer.StartServer(); err != nil {
	// 	fmt.Println("failed to start HTTP server:", err)
	// }

}
