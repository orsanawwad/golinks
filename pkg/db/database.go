package db

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GoLinkDB struct {
	conn *gorm.DB
}

type Link struct {
	gorm.Model
	Owner string
	Key   string
	URL   string
}

func New() *GoLinkDB {
	db, err := gorm.Open(sqlite.Open("golinks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Link{})
	return &GoLinkDB{
		conn: db,
	}
}
func (gldb *GoLinkDB) GetLink(key string) string {
	var link Link
	result := gldb.conn.Limit(1).Find(&link, "key = ?", key)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ""
	}
	return link.URL
}

func (gldb *GoLinkDB) CreateLink(link *Link) {
	gldb.conn.Create(&link)
}
