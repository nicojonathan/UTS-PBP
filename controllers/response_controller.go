package controllers

import (
	"encoding/json"
	"net/http"
	m "uts/models"
)

func sendRoomSuccessResponse (w http.ResponseWriter, message string, data []m.Room) {
	var response m.RoomsResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendRoomDetailSuccessResponse (w http.ResponseWriter, message string, data []m.RoomDetail) {
	var response m.RoomDetailResponse
	response.Status = 200
	response.Message = message
	response.Data.Room = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// func sendTransactionDetailSuccessResponse (w http.ResponseWriter, message string, data []m.TransactionDetail) {
// 	var response m.TransactionDetailResponse
// 	response.Status = 200
// 	response.Message = message
// 	response.Data.Transaction = data
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func sendSuccessResponse (w http.ResponseWriter, message string) {
	var response m.GeneralResponse
	response.Status = 200
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse (w http.ResponseWriter, status int, message string) {
	var response m.GeneralResponse
	response.Status = status 
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

