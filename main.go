package main

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/biplob-codes/mockly/internal/config"
	"github.com/biplob-codes/mockly/internal/database"
	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/handlers"
	"github.com/biplob-codes/mockly/internal/store"
)

func main() {
	mux := http.NewServeMux()
	var villageStore store.VillageStore
	cfg:=config.Load()
	if cfg.Env=="local"{
    db, err := database.Connect(cfg.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Database connection successfull")
	defer db.Close()
	queries:=sqlc.New(db)
        villageStore=store.CreateDBVillageStore(queries)
	}else{
		villageStore=store.CreateMemoryVillageStore()
	}
	
 
	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /health", handlers.Health)
   
	villageHandler:=handlers.NewVillageHandler(villageStore)

    mux.HandleFunc("GET /villages", villageHandler.ListVillages)
	mux.HandleFunc("GET /villages/{id}", villageHandler.GetVillage)
	mux.HandleFunc("POST /villages", villageHandler.CreateVillage)
	mux.HandleFunc("DELETE /villages/{id}", villageHandler.DeleteVillage)
	 
 

	addr := ":" + cfg.Port
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}


	fmt.Println("Listening on port",cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
