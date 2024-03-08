package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	m "uts/models"
	// "github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)


func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	id := r.Form.Get("id")
	room_name := r.Form.Get("room_name")

	query := "SELECT * FROM rooms WHERE 1"

	if id != "" {
		query += fmt.Sprintf(" AND id='%s'", id)
	}

	if room_name != "" {
		query += fmt.Sprintf(" AND room_name=%s", room_name)
	}


	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {		
		sendErrorResponse(w, 500, "Internal Server Error! Database Query Failed!")
		return
	}

	var room m.Room
	var rooms []m.Room
	var dataFound bool

	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.RoomName, &sql.NullInt64{}); err != nil {
			sendErrorResponse(w, 500, "Fail to scan row")
			return
		}else {
			rooms = append(rooms, room)
			dataFound = true
		}
	}

	if (!dataFound) {
		sendErrorResponse(w, 404, "Data not found!")
		return
	}

	sendRoomSuccessResponse(w, "Get Room Successful", rooms)
}

func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `
	SELECT r.id, r.room_name, p.id, a.id, a.username FROM participants p JOIN rooms r ON p.id_room = r.id JOIN accounts a ON p.id_account = a.id;`


	RoomDetailRow, err := db.Query(query)
	if err != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Database query fail!")
		return
	}

	var roomDetail m.RoomDetail
	var roomDetails []m.RoomDetail
	for RoomDetailRow.Next(){
		if err := RoomDetailRow.Scan(
		&roomDetail.Id, &roomDetail.RoomName, &roomDetail.Participant.Id, &roomDetail.Participant.IdAccount, &roomDetail.Participant.UserName); err != nil {
			sendErrorResponse(w, 500, "Internal server Error! Fail to scan row!")
			return
		} else {
			roomDetails = append(roomDetails, roomDetail)
		}
	}

	sendRoomDetailSuccessResponse(w, "Get Room Detail Successful", roomDetails)
}



	// ParticipantDetailRow, err := db.Query(query)
	// if err != nil {
	// 	sendErrorResponse(w, 500, "Internal Server Error! Database query fail!")
	// 	return
	// }

	// var participant m.Participant
	// var participants []m.Participant
	// for ParticipantDetailRow.Next(){
	// 	if err := ParticipantDetailRow.Scan(
	// 	&sql.NullInt64{}, &sql.NullInt64{}, &participant.Id, &participant.IdAccount, &participant.UserName); err != nil {
	// 		sendErrorResponse(w, 500, "Internal server Error! Fail to scan row!")
	// 		return
	// 	} else {
	// 		participants = append(participants, participant)
	// 	}
	// }

	// query2 := `
	// SELECT r.id, r.room_name FROM participants p JOIN rooms r ON p.id_room = r.id JOIN accounts a ON p.id_account = a.id;`