// curl -X POST -H "Content-Type: application/json" --data '{"regNums": ["X123XX150", "Y456YY789"]}' "http://localhost:8080/cars"

package handle

import (
	"database/sql"
	"net/http"
)

func PostCars(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}
