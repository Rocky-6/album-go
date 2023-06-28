package service

import "database/sql"

type AlbumService struct {
	db *sql.DB
}

func NewAlbumService(db *sql.DB) *AlbumService {
	return &AlbumService{
		db: db,
	}
}
