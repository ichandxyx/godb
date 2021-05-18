package api

import (
	"encoding/json"
	"net/http"
)

func (a *API) handleGetusers(w http.ResponseWriter, r *http.Request) {
	users, err := a.store.GetUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"msg":    "someting went wrong",
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"users":  users,
	})

}

func (a *API) handleAddUsers(w http.ResponseWriter, r *http.Request){
	var pr struct{
	Name string `json:"name"`
	Age int `json:"age"`
}
	json.NewDecoder(r.Body).Decode(&pr); 
	

	users,err :=a.store.AddUsers(r.Context(),pr.Name,pr.Age)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"msg":    "someting went wrong",
		})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"users":   users,

	})
}
// type name string
// var na string ="chandan"
// func (x name) sayhello() string
// {
// 	return "hello"+x
// }
