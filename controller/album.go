package controller

import (
	"album/model"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (con *Controller) AddAlbum(ctx *gin.Context) {
	var album model.Album
	queryAddAlbum := `INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)`
	checkExistence := `SELECT EXISTS (SELECT * FROM album WHERE title = $1 AND artist = $2)`

	if err := ctx.BindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var exist bool
	err := con.db.QueryRow(checkExistence, album.Title, album.Artist).Scan(&exist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "album already exists"})
		return
	}

	_, err = con.db.Exec(queryAddAlbum, album.Title, album.Artist, album.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, album)
}

func (con *Controller) GetAlbums(ctx *gin.Context) {
	var albums []model.Album
	queryGetAlbums := "SELECT * FROM album"

	rows, err := con.db.Query(queryGetAlbums)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var album model.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}
		albums = append(albums, album)
	}
	ctx.JSON(http.StatusOK, albums)
}
