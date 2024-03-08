package controllers

import (
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	id_account := r.Form.Get("id_account")
	id_room := r.Form.Get("id_room")

	queryCheckRoom := "SELECT COUNT(p.id) AS 'total_player', g.max_player FROM games g JOIN rooms r ON g.id = r.id_game JOIN participants p ON r.id = p.id_room WHERE r.id = ? GROUP BY r.id;"

	var total_player int
	var max_player int

	resultRoom, err := db.Query(queryCheckRoom, id_room)
	if err != nil {
		sendErrorResponse(w, 500, "Internal Server Error! DB Query to Room Fail!")
		return
	}

	if resultRoom.Next() {
		errScan := resultRoom.Scan(&total_player, &max_player)

		if errScan != nil {
			fmt.Println(errScan)
			sendErrorResponse(w, 500, "Internal Server Error! Fail to Scan resultRoom!")
			return
		}

	}else{
		sendErrorResponse(w, 404, "room not found!")
		return
	}

	if total_player >= max_player {
		sendErrorResponse(w, 400, "Opps room already full!")
		return
	}

	queryInsert := "INSERT INTO participants (`id_room`, `id_account`) VALUES (?, ?);"

	_, errInsert := db.Exec(queryInsert, id_room, id_account)
	if errInsert != nil {
		sendErrorResponse(w, 500, "Internal Server Error! DB Query to Participant Fail")
		return
	}

	sendSuccessResponse(w, "Insert Participant to Room Successfull")
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	id_account := r.Form.Get("id_account")
	id_room := r.Form.Get("id_room")

	query := "DELETE FROM participants WHERE id_account = ? AND id_room = ?"

	_, err := db.Exec(query, id_account, id_room)
	if err != nil {
		sendErrorResponse(w, 500, "DB Query Fail")
		return
	}

	sendSuccessResponse(w, "Leave Room Successfull")
}