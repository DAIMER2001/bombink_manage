package database

import "gorm.io/gorm"

type BombinkDB struct {
	conn *gorm.DB
}

func NewBombinkDB(conn *gorm.DB) *BombinkDB {
	return &BombinkDB{
		conn: conn,
	}
}
