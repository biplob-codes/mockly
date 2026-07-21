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
	var characterStore store.CharacterStore
	var jutsuStore store.JutsuStore
	var characterJutsuStore store.CharacterJutsuStore
	var teamStore store.TeamStore
	var missionStore store.MissionStore
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
        characterStore=store.CreateDBCharacterStore(queries)
		jutsuStore=store.CreateDBJutsuStore(queries)
		characterJutsuStore=store.CreateDBCharacterJutsuStore(queries)
		teamStore=store.CreateDBTeamStore(queries)
		missionStore=store.CreateDBMissionStore(queries)
	}else{
		villageStore=store.CreateMemoryVillageStore()
		characterStore=store.CreateMemoryCharacterStore()
		jutsuStore=store.CreateMemoryJutsuStore()
		characterJutsuStore=store.CreateMemoryCharacterJutsuStore()
		teamStore=store.CreateMemoryTeamStore()
		missionStore=store.CreateMemoryMissionStore()
	}
	
 
	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /health", handlers.Health)
   
	villageHandler:=handlers.NewVillageHandler(villageStore)
	characterHandler:=handlers.NewCharacterHandler(characterStore)
	jutsuHandler:=handlers.NewJutsuHandler(jutsuStore)
	characterJutsuHandler:=handlers.NewCharacterJutsuHandler(characterJutsuStore)
    teamHandler:=handlers.NewTeamHandler(teamStore)
	missionHandler:=handlers.NewMissionHandler(missionStore)

    mux.HandleFunc("GET /villages", villageHandler.ListVillages)
	mux.HandleFunc("GET /villages/{id}", villageHandler.GetVillage)
	mux.HandleFunc("POST /villages", villageHandler.CreateVillage)
	mux.HandleFunc("DELETE /villages/{id}", villageHandler.DeleteVillage)

    mux.HandleFunc("GET /characters", characterHandler.ListCharacters)
	mux.HandleFunc("GET /characters/{id}", characterHandler.GetCharacter)
	mux.HandleFunc("GET /characters/{id}/jutsus", characterJutsuHandler.ListJutsusByCharacter)
	mux.HandleFunc("POST /characters", characterHandler.CreateCharacter)
	mux.HandleFunc("DELETE /characters/{id}", characterHandler.DeleteCharacter)
    mux.HandleFunc("POST /characters/{id}/jutsus", characterJutsuHandler.CreateCharacterJutsu)
    mux.HandleFunc("DELETE /characters/{characterId}/jutsus/{jutsuId}",characterJutsuHandler.DeleteCharacterJutsu)

    mux.HandleFunc("GET /jutsus", jutsuHandler.ListJutsus)
	mux.HandleFunc("GET /jutsus/{id}", jutsuHandler.GetJutsu)
	mux.HandleFunc("GET /jutsus/{id}/characters", characterJutsuHandler.ListCharactersByJutsu)
	mux.HandleFunc("POST /jutsus", jutsuHandler.CreateJutsu)
	mux.HandleFunc("DELETE /jutsus/{id}", jutsuHandler.DeleteJutsu)

   mux.HandleFunc("GET /teams", teamHandler.ListTeams)
	mux.HandleFunc("POST /teams", teamHandler.CreateTeam)
mux.HandleFunc("DELETE /teams/{id}", teamHandler.DeleteTeam)
mux.HandleFunc("GET /teams/{id}", teamHandler.GetTeam)
mux.HandleFunc("GET /teams/{id}/members", teamHandler.GetTeamDetails)

mux.HandleFunc("POST /missions", missionHandler.CreateMission)
mux.HandleFunc("GET /missions/{id}", missionHandler.GetMission)
mux.HandleFunc("GET /missions", missionHandler.ListMissions)
mux.HandleFunc("GET /teams/{teamId}/missions", missionHandler.GetMissionsByTeam)
mux.HandleFunc("DELETE /missions/{id}", missionHandler.DeleteMission)

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
