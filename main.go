package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Flight struct {
	ID             string `json"`
	Brand_id       string `json:"E"`
	Flight_id      string `json:"F"`
	Departure_time string `json:"G"`
	Arrival_time   string `json:"H"`
	Price          string `json:"J"`
}

//type Brand struct {
//	Brand_id   string
//	Brand_code string
//	Brand_des  string
//}

func insertData(db *sql.DB) {
	date := "2023-07-01 "

	var flights []Flight
	err := json.Unmarshal([]byte(data), &flights)
	if err != nil {
		fmt.Println(err)
	}

	for _, flight := range flights {
		_, err = db.Exec("INSERT INTO flight(brand_id, flight_id, departure_time, arrival_time, price) VALUES(?, ?, ?, ?, ?)", flight.Brand_id, flight.Flight_id, date+flight.Departure_time+":00", date+flight.Arrival_time+":00", flight.Price)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	// ket noi voi database
	db, err := sql.Open("mysql", "root:vuvietduy1234@tcp(localhost:3306)/data")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected")
	}

	//luu data lên mysql
	//insertData(db)

	//api tim kiem
	app := fiber.New()

	app.Get("/api/getFlight", func(c *fiber.Ctx) error {
		date := c.Query("date")
		sortby := c.Query("sortby")
		filterbyairline := c.Query("filterbyairline")

		query := "SELECT * FROM flight WHERE DATE(departure_time) = ? "

		if filterbyairline != "" {
			query += "AND brand_id = " + filterbyairline
		}

		switch sortby {
		case "price":
			query += " ORDER BY price;"
		case "departure_time":
			query += " ORDER BY departure_time ASC;"
		case "brand_id":
			query += " ORDER BY brand_id ASC;"
		default:
		}

		rows, err := db.Query(query, date)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		var flights []Flight
		for rows.Next() {
			var flight Flight
			err = rows.Scan(&flight.ID, &flight.Flight_id, &flight.Brand_id, &flight.Departure_time, &flight.Arrival_time, &flight.Price)
			if err != nil {
				fmt.Println(err)
			}
			flights = append(flights, flight)
		}
		return c.JSON(flights)
	})

	//listen port 3000
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
