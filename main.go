package main

import (
	"database/sql"
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

	//client := resty.New()
	//
	//resp, error := client.R().Get("https://www.agoda.com/api/personalization/PersonalizeRecommendedProperties/v1?recommendationType=2&hotelId=0&finalPriceView=1&hasSearchCriteria=true&cityId=13170&lengthOfStay=1&checkIn=2023-07-13T00%3A00%3A00&adults=1&children=0&rooms=1&_ts=1688365681981")
	//if error != nil {
	//	fmt.Println(error)
	//}
	//
	//if resp.StatusCode() != 200 {
	//	fmt.Println("error")
	//}
	//
	//body := resp.Body()
	//
	//fmt.Println(string(body))

	//luu data lên mysql
	//date := "2023-07-01 "
	//
	//var flights []Flight
	//err = json.Unmarshal([]byte(data), &flights)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for _, flight := range flights {
	//	_, err = db.Exec("INSERT INTO flight(brand_id, flight_id, departure_time, arrival_time, price) VALUES(?, ?, ?, ?, ?)", flight.Brand_id, flight.Flight_id, date+flight.Departure_time+":00", date+flight.Arrival_time+":00", flight.Price)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	app := fiber.New()

	//app.Get("/api/getFlight", func(c *fiber.Ctx) error {
	//	date := c.Query("date")
	//
	//	rows, err := db.Query("SELECT * FROM flight WHERE departure_time LIKE ?", date)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	defer rows.Close()
	//
	//	var flights []Flight
	//	for rows.Next() {
	//		var flight Flight
	//		err = rows.Scan(&flight.ID, &flight.Flight_id, &flight.Brand_id, &flight.Departure_time, &flight.Arrival_time, &flight.Price)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		flights = append(flights, flight)
	//	}
	//	return c.JSON(flights)
	//})

	app.Get("/api/getFlight", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM flight")
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

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
