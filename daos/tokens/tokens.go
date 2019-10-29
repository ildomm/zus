package urls

import (
	"github.com/go-openapi/strfmt"
	db "github.com/ildomm/zus/database"
	"github.com/ildomm/zus/models"
	"github.com/lib/pq"
	"log"
	"time"
)
	
var tableName = "tokens"
var tableFields = "token, hash, created_at"

func All() []*models.TokenInfo {
	var entry []*models.TokenInfo

	session := db.Postgres()
	_, err := session.Select(tableFields).From(tableName).Load(&entry)
	if err != nil {
		return nil
	}

	return entry
}

func ByHash(hash string) *models.TokenInfo {
	var entry *models.TokenInfo

	session := db.Postgres()
	session.Select(tableFields).From(tableName).
		Where("hash = ?", hash).
		LoadOne(&entry)
	return entry
}

func TokenExists(token string) bool {
	var total int

	session := db.Postgres()
	err :=
		session.Select("COUNT(*) as total").
			From(tableName).
			Where("token = ?", token).
			LoadOne(&total)

	if err != nil {
		log.Printf("Error select from table %s : %s", tableName, err.Error())
	}
	if total > 0 {
		return true
	} else {
		return false
	}
}

func Insert(entry models.TokenInfo) (*models.TokenInfo, *pq.Error) {

	entry.Hash = entry.GenerateHash()

	session := db.Postgres()
	stmt := session.InsertInto(tableName).
		Pair("token", entry.Token).
		Pair("hash", entry.Hash).
		Pair("created_at", time.Now().Format(db.TIME_FORMAT)).
		Returning("id")

	var entryID strfmt.UUID
	err := stmt.Load(&entryID)

	if err != nil {
		errDB := err.(*pq.Error)
		return nil, errDB
	} else {
		entry.ID = entryID
		return &entry, nil
	}

	return nil, nil
}