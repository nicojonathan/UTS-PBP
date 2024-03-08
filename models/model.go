package models

//import "time"

type Room struct {
	Id       int    `json:"id,omitempty"`
	RoomName     string `json:"room_name,omitempty"`
	IdGame    int `json:"id_game,omitempty"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data,omitempty"`
}

type Account struct {
	Id       int    `json:"id,omitempty"`
	UserName     string `json:"username,omitempty"`
}

// type Participant struct {
// 	Id       int    `json:"id,omitempty"`
// 	IdAccount int `json:"id_account,omitempty"`
// 	Account Account `json:"user,omitempty"`
// }

type Participant struct {
	Id       int    `json:"id,omitempty"`
	IdAccount int `json:"id_account,omitempty"`
	UserName string `json:"username,omitempty"`
}

type RoomDetail struct {
	Id       int    `json:"id,omitempty"`
	RoomName     string `json:"room_name,omitempty"`
	Participant Participant `json:"pariticipant"`
}

type RoomsDetail struct {
	Room []RoomDetail `json:"room"`
}
	
type RoomDetailResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    RoomsDetail `json:"data"`
}


// type TransactionsDetail struct {
// 	Transaction []TransactionDetail `json:"transactions"`
// }
	


// type Playlist struct {
// 	PlaylistId          int    `json:"id,omitempty" gorm:"primaryKey"`
// 	PlaylistName        string `json:"name,omitempty"`
// 	PlayListDateCreated time.Time`json:"date_created,omitempty"`
// 	PlayListState bool `json:"playlistState,omitempty"`
// 	UserId int `json:"userId,omitempty"`
// }

// type Song struct {
// 	SongId          int    `json:"id,omitempty" gorm:"primaryKey"`
// 	SongTitle        string `json:"title,omitempty"`
// 	SongDuration float64 `json:"duration,omitempty"`
// 	SongSinger       string `json:"singer,omitempty"`
// }

// type PopularSong struct {
// 	Song Song `json:"Song"`
// 	TimePlayed int `json:"time played"`
// }


// type PopularSongsResponse struct {
// 	Status  int                `json:"status"`
// 	Message string             `json:"message"`
// 	Data    []PopularSong `json:"data"`
// }

// type RecommendedSongResponse struct {
// 	Status  int                `json:"status"`
// 	Message string             `json:"message"`
// 	Data    Song `json:"data"`
// }

// type DetailPlaylistSong struct {
// 	PlayList Playlist `json:"playlist"`
// 	Song Song `json:"song"`
// 	TimePlayed int `json:"timeplayed"`
// }

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}


