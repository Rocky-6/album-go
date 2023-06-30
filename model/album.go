package model

type (
	Album struct {
		ID     int64   `json:"id"`
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float32 `json:"price"`
	}

	AddAlbumRequest struct {
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float32 `json:"price"`
	}

	AddAlbumResponse struct {
		Album Album `json:"album"`
	}

	GetAlbumByIDRequest struct {
		ID int64 `json:"id"`
	}

	GetAlbumByIDResponse struct {
		Album Album `json:"album"`
	}

	GetAlbumsByTitelRequest struct {
		Title string `json:"title"`
	}

	GetAlbumsByTitleResponse struct {
		Album []Album `json:"albums"`
	}

	GetAlbumsByArtistRequest struct {
		Artist string `json:"artist"`
	}

	GetAlbumsByArtistResponse struct {
		Album []Album `json:"albums"`
	}
)
