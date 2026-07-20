package store

import (
	"context"
	"fmt"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/seed"
)

type MemoryCharacterJutsuStore struct{
	cj []sqlc.CharactersJutsu
}

func CreateMemoryCharacterJutsuStore() *MemoryCharacterJutsuStore{
	return &MemoryCharacterJutsuStore{cj:seed.CharacterJutsuData}
}
func findCharacterJutsu(cID,jID int64) (sqlc.CharactersJutsu, bool) {
	for _, c := range seed.CharacterJutsuData{
		if c.JutsuID == jID && c.CharacterID==cID {
			return c, true
		}
	}
	return sqlc.CharactersJutsu{}, false
}
func (cjs *MemoryCharacterJutsuStore) Create(ctx context.Context, p sqlc.CreateCharacterJutsuParams) (sqlc.CharactersJutsu, error){
	 newcj:=sqlc.CharactersJutsu{
		ID: int64(len(seed.CharacterJutsuData) +1),
		CharacterID: p.CharacterID,
		JutsuID: p.JutsuID,
		CreatedAt: time.Now(),
	 }
	 return newcj,nil
}

func (cjs *MemoryCharacterJutsuStore) ListJutsusByCharacter(ctx context.Context,id int64) ([]sqlc.Jutsu, error){
	var result []sqlc.Jutsu
	for _,cj:=range cjs.cj{
		if cj.CharacterID==id{
			for _,j:=range seed.JutsuData{
				if j.ID==cj.JutsuID{
					result = append(result, j)
				}
			}
		}
	}
	return result,nil
}
func (cjs *MemoryCharacterJutsuStore) ListCharactersByJutsu(ctx context.Context,id int64) ([]sqlc.Character, error){
	var result []sqlc.Character
	for _,cj:=range cjs.cj{
		if cj.JutsuID==id{
			for _,j:=range seed.CharacterData{
				if j.ID==cj.CharacterID{
					result = append(result, j)
				}
			}
		}
	}
	return result,nil
}

func (cj *MemoryCharacterJutsuStore) Delete(ctx context.Context, characterId,jutsuId int64) (sqlc.CharactersJutsu, error){
   characterJutsu,found:=findCharacterJutsu(characterId,jutsuId)
   if found{
	return  characterJutsu,nil
   }
   return sqlc.CharactersJutsu{},fmt.Errorf("Character-Jutsu not found")
}