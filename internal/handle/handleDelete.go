// curl -X DELETE "http://localhost:8080/cars/id"

package handle

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteCars(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ids, err := getExistingIDs(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := ids[id]; ok {
		_, err = db.Query("DELETE FROM cars WHERE id = $1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getExistingIDs(db *sql.DB) (map[int]struct{}, error) {
	rows, err := db.Query("SELECT id FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := make(map[int]struct{})
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids[id] = struct{}{}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}
