// curl -X GET "http://localhost:8080/cars?page=1"
// curl -X GET "http://localhost:8080/cars/mark/Lada?page=0"

package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Car struct {
	ID     int    `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  People `json:"owner"`
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

// GetCars godoc
// @Summary Get a list of cars
// @Description Get a paginated list of cars with optional filtering by field and value
// @Tags cars
// @Param page query int false "Page number"
// @Param filterField path string false "Filter field name"
// @Param valueField path string false "Filter value"
// @Produce json
// @Success 200 {array} Car
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /cars [get]
func GetCars(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	vars := mux.Vars(r)
	field := vars["filterField"]
	value := vars["valueField"]

	cars, total, err := getPage(db, page, field, value)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Print(err)
	}

	renderPage(w, cars, total, page)
}

func getPage(db *sql.DB, page int, field string, value string) ([]Car, int, error) {
	rows, count, err := parseField(db, page, field, value)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cars []Car
	for rows.Next() {
		var car Car
		err := rows.Scan(&car.ID, &car.RegNum, &car.Mark,
			&car.Model, &car.Year, &car.Owner.Name,
			&car.Owner.Surname, &car.Owner.Patronymic)
		if err != nil {
			return nil, 0, err
		}
		cars = append(cars, car)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return cars, count, nil
}

func parseField(db *sql.DB, page int, field string, value string) (*sql.Rows, int, error) {
	var count int
	var rows *sql.Rows
	var err, err2 error
	const LIMIT = 10
	offset := page * LIMIT

	switch field {
	case "":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id 
			LIMIT $1 OFFSET $2`,
			LIMIT, offset)
		err2 = db.QueryRow("SELECT count(*) FROM cars").Scan(&count)
	case "id":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id 
			WHERE c.id = $1 LIMIT $2 OFFSET $3`,
			value, LIMIT, offset)
		err2 = db.QueryRow(`SELECT count(*) FROM cars c
		    WHERE c.id = $1`, value).Scan(&count)
	case "regNum":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id 
			WHERE c.regNum = $1 LIMIT $2 OFFSET $3`,
			value, LIMIT, offset)
		err2 = db.QueryRow(`SELECT count(*) FROM cars c 
			WHERE c.regNum = $1`, value).Scan(&count)
	case "mark":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id 
			WHERE c.mark = $1 LIMIT $2 OFFSET $3`,
			value, LIMIT, offset)
		err2 = db.QueryRow(`SELECT count(*) FROM cars c 
		    WHERE c.mark = $1`, value).Scan(&count)
	case "model":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id
			WHERE c.model = $1 LIMIT $2 OFFSET $3`,
			value, LIMIT, offset)
		err2 = db.QueryRow(`SELECT count(*) FROM cars c 
			WHERE c.model = $1`, value).Scan(&count)
	case "year":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id
			WHERE c.year = $1 LIMIT $2 OFFSET $3`,
			value, LIMIT, offset)
		err2 = db.QueryRow(`SELECT count(*) FROM cars c 
			WHERE c.year = $1`, value).Scan(&count)
	case "name":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id
			WHERE p.name = $1 LIMIT $2 OFFSET $3`,
			value, LIMIT, offset)
		err2 = db.QueryRow(`SELECT count(*) FROM cars c 
			JOIN people p ON c.id = p.car_id WHERE p.name = $1`, value).Scan(&count)
	case "surname":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id
			WHERE p.surname = $1 LIMIT $2 OFFSET $3`,
			value, LIMIT, offset)
		err2 = db.QueryRow(`SELECT count(*) FROM cars c 
			JOIN people p ON c.id = p.car_id WHERE p.surname = $1`, value).Scan(&count)
	case "patronymic":
		rows, err = db.Query(
			`SELECT c.id, c.regNum, c.mark, c.model, c.year, p.name, p.surname, p.patronymic 
			FROM cars c 
			JOIN people p ON c.id = p.car_id
			WHERE p.patronymic = $1 LIMIT $2 OFFSET $3`,
			value, LIMIT, offset)
		err2 = db.QueryRow(`SELECT count(*) FROM cars c 
			JOIN people p ON c.id = p.car_id WHERE p.patronymic = $1`, value).Scan(&count)
	default:
		return nil, 0, fmt.Errorf(fmt.Sprintf("Invalid field for filtering: %s", field))
	}

	if err != nil {
		return nil, 0, err
	}

	if err2 != nil {
		return nil, 0, err
	}

	return rows, count, nil
}

func renderPage(w http.ResponseWriter, cars []Car, total, currentPage int) {
	totalPages := (total + 9) / 10

	if currentPage < 0 || currentPage >= totalPages {
		http.Error(w, fmt.Sprintf("Invalid page: %v", currentPage), http.StatusNotFound)
		return
	}

	pageData := struct {
		Name     string
		Total    int
		Cars     []Car
		PrevPage int
		NextPage int
		LastPage int
	}{
		Name:     "cars",
		Total:    total,
		Cars:     cars,
		PrevPage: max(currentPage-1, 0),
		NextPage: min(currentPage+1, totalPages-1),
		LastPage: totalPages - 1,
	}

	data, err := json.MarshalIndent(pageData, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeHeader(w, data)
}

func writeHeader(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
