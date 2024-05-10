package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type API struct{}

func (api *API) HandleHome(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitParam)
	json.NewEncoder(w).Encode(getUsers("./data.json")[:limit])
}

type User struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func getUsers(path string) []User {
	dataJson, _ := os.ReadFile(path)
	var users []User
	json.Unmarshal(dataJson, &users)
	return users
}
