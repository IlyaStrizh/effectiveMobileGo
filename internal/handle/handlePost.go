// curl -X POST -H "Content-Type: application/json" --data '{"regNums": ["X123XX150", "Y456YY789"]}' "http://localhost:8080/cars"

package handle

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

// PostCars godoc
// @Summary Create new cars
// @Description Create new cars in the database based on the provided registration numbers
// @Tags cars
// @Accept json
// @Param car body RegNums true "Registration numbers"
// @Produce json
// @Success 201 {string} string "Cars created successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal Server Error"
// @Router /cars [post]
func PostCars(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	type RequestBody struct {
		RegNums []string `json:"regNums"`
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, regNum := range requestBody.RegNums {
		car := Car{
			RegNum: regNum,
		}

		apiUrl := os.Getenv("API_URL") + "?regNum=" + url.QueryEscape(regNum)
		resp, err := http.Get(apiUrl)
		if err != nil {
			log.Printf("Error making external API request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			err = json.NewDecoder(resp.Body).Decode(&car)
			if err != nil {
				log.Printf("Error decoding external API response: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		tx, err := db.Begin()
		if err != nil {
			log.Printf("Error starting database transaction: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		insertCarQuery := "INSERT INTO cars (regNum, mark, model, year) VALUES ($1, $2, $3, $4) RETURNING id"
		var carID int
		err = tx.QueryRow(insertCarQuery, car.RegNum, car.Mark, car.Model, car.Year).Scan(&carID)
		if err != nil {
			tx.Rollback()
			log.Printf("Error inserting car into database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		insertPeopleQuery := "INSERT INTO people (name, surname, patronymic, car_id) VALUES ($1, $2, $3, $4)"
		_, err = tx.Exec(insertPeopleQuery, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic, carID)
		if err != nil {
			tx.Rollback()
			log.Printf("Error inserting people into database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Printf("Error committing database transaction: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cars created successfully\n"))
}
