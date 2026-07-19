package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/biplob-codes/mockly/internal/database"
	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/handlers"
	"github.com/biplob-codes/mockly/internal/store"
)

func main() {
	mux := http.NewServeMux()
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
    

	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /health", handlers.Health)
    app_env:="binary"

	var villageStore store.VillageStore
	if app_env=="binary"{
		queries:=sqlc.New(db)
        villageStore=store.CreateDBVillageStore(queries)
	}else{

		villageStore=store.CreateMemoryVillageStore()
	}
	villageHandler:=handlers.NewVillageHandler(villageStore)

    mux.HandleFunc("GET /villages", villageHandler.ListVillages)
	mux.HandleFunc("GET /villages/{id}", villageHandler.GetVillage)
	mux.HandleFunc("POST /villages", villageHandler.CreateVillage)
	mux.HandleFunc("DELETE /villages/{id}", villageHandler.DeleteVillage)
	 
 

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}


	fmt.Println("Database connection successfull")
	fmt.Println("Listening on port", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
