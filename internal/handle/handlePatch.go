// curl -X PATCH -H "Content-Type: application/json" --data '{"id": 1, "regNum": "Новый регистрационный номер", "owner": {"name": "Новое имя владельца"}}' "http://localhost:8080/cars/1"

package handle

//import (
//	"database/sql"
//	"encoding/json"
//	"log"
//	"net/http"
//	"strconv"
//
//	"github.com/gorilla/mux"
//)
//
//func PatchCars(w http.ResponseWriter, r *http.Request, db *sql.DB) {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//
//	if err != nil {
//		http.Error(w, "Invalid ID", http.StatusBadRequest)
//		return
//	}
//
//	var car Car
//	err = json.NewDecoder(r.Body).Decode(&car)
//	if err != nil {
//		log.Printf("Error decoding request body: %v", err)
//	}
//
//}
