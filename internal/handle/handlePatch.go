// curl -X PATCH -H "Content-Type: application/json" --data '{"regNum": "Новый регистрационный номер", "owner": {"name": "Новое имя владельца"}}' "http://localhost:8080/cars/1"

package handle

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PatchCars godoc
// @Summary Update a car by ID
// @Description Update a car in the database by its ID
// @Tags cars
// @Accept json
// @Param id path int true "Car ID"
// @Param car body Car true "Car object"
// @Produce json
// @Success 200 {string} string "Car updated successfully"
// @Failure 400 {string} string "Invalid ID or request body"
// @Failure 500 {string} string "Internal Server Error"
// @Router /cars/{id} [patch]
func PatchCars(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateData map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateQueryCars := "UPDATE cars SET"
	params := []interface{}{}
	index := 2

	params = append(params, id)
	for field, value := range updateData {
		if val, ok := value.(map[string]interface{}); ok && field == "owner" {
			updateQueryPeople := "UPDATE people SET"
			paramsPeople := []interface{}{}
			indexPeople := 2

			paramsPeople = append(paramsPeople, id)
			for fieldPeople, valuePeople := range val {
				updateQueryPeople += " " + fieldPeople + " = $" + strconv.Itoa(indexPeople) + ","
				paramsPeople = append(paramsPeople, valuePeople)
				indexPeople++
			}
			updateQueryPeople = updateQueryPeople[:len(updateQueryPeople)-1]
			updateQueryPeople += " WHERE car_id = $1"

			err = postgresQuery(db, updateQueryPeople, paramsPeople)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("Error updating base: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			updateQueryCars += " " + field + " = $" + strconv.Itoa(index) + ","
			params = append(params, value)
			index++
		}
	}

	updateQueryCars = updateQueryCars[:len(updateQueryCars)-1]
	updateQueryCars += " WHERE id = $1"

	err = postgresQuery(db, updateQueryCars, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error updating base: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Car updated successfully\n"))
}

func postgresQuery(db *sql.DB, updateQuery string, params []interface{}) error {
	_, err := db.Exec(updateQuery, params...)
	if err != nil {
		return err
	}

	return nil
}
