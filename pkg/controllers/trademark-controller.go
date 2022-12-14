package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-demo/pkg/models"
	"go-demo/pkg/utils"
	"net/http"
	"strconv"
)

var NewRecord models.Root

func GetRecordById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recordId := vars["recordId"]
	ID, err := strconv.ParseInt(recordId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	recordDetails, _ := models.GetRecordById(ID)
	res, _ := json.Marshal(recordDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	CreateRecord := &models.Root{}
	utils.ParseBody(r, CreateRecord)
	b := CreateRecord.CreateRecord
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
