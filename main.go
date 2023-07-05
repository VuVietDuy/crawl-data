package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

type Flight struct {
	ID             string `json:"ID"`
	Brand_code     string `json:"E"`
	Flight_id      string `json:"F"`
	Departure_time string `json:"G"`
	Arrival_time   string `json:"H"`
	Price          string `json:"J"`
}

type Data struct {
	Brands  string   `json:"A"`
	B       string   `json:"B"`
	C       string   `json:"C"`
	Flights []Flight `json:"D"`
}

type D struct {
	D string `json:"d"`
}

type Brand struct {
	Brand_id   string
	Brand_code string
	Brand_des  string
}

func insertData(db *sql.DB, flights []Flight) {
	date := "2023-07-03 "

	for _, flight := range flights {
		fmt.Println(flight)
		var brand_id_int int
		brand_id, err := db.Query("SELECT brand_id FROM Brand WHERE brand_code = ?", flight.Brand_code)
		brand_id.Next()
		err = brand_id.Scan(&brand_id_int)
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Exec("INSERT INTO Flight(brand_id, flight_id, departure_timw, arrival_time, price) VALUES(?, ?, ?, ?, ?)", brand_id_int, flight.Flight_id, date+flight.Departure_time+":00", date+flight.Arrival_time+":00", flight.Price)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	// ket noi voi database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/data")
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

	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"input":             "SGN-HAN-1-0-0-05Jul2023-",
			"sortby":            "abay-suggest",
			"filterbyairline":   "",
			"defaultSort":       "1",
			"IsRewriteFare":     "0",
			"display":           "base",
			"edition":           "2",
			"cookieAll":         `{"CID":"164130","CD":"26Jun2023","LVPDI":"1510186","_gid":"GA1.2.882121964.1688344572","TCB":"1","_gcl_au":"1.1.1839823084.1688366182","SumAll":"14","LVPD":"1510186,0,0,1,0,2,14,4","RetDate":"","Client-Browser":"1","DepDate":"05-07-2023","VHP":"05/07/2023 07:00:09","VHPTD":"05JUL2023.1","AT":"4","StartPoint":"SGN","EndPoint":"HAN","Adult":"1","Child":"0","Infant":"0","DV":"05Jul2023","CDV":"6","tsv2":"2023-07-04T23:59:47.619Z"},"waytype":"OutBound","isAllowSearchReal":0}`,
			"waytype":           "OutBound",
			"isAllowSearchReal": "0"}).
		SetHeaders(map[string]string{
			"Accept":             "application/json, text/javascript, */*; q=0.01",
			"Accept-Language":    "en-US,en;q=0.9,vi;q=0.8",
			"Connection":         "keep-alive",
			"Content-Length":     "0",
			"Content-Type":       "application/json",
			"Origin":             "https://www.abay.vn",
			"Referer":            "https://www.abay.vn/",
			"Sec-Fetch-Dest":     "empty",
			"Sec-Fetch-Mode":     "cors",
			"Sec-Fetch-Site":     "same-site",
			"User-Agent":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
			"sec-ch-ua":          `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`,
			"sec-ch-ua-mobile":   "?0",
			"sec-ch-ua-platform": "Linux",
		}).
		Post("https://chuyenbay2.abay.vn/_WEB/ResultDom2/ResultDomAjax.aspx/GetFlights")

	response := string(resp.Body())
	//fmt.Println(response)

	var d D
	err = json.Unmarshal([]byte(response), &d)
	if err != nil {
		fmt.Println(err)
	}

	data := d.D
	data = strings.Replace(data, "\\", "", -1)
	var da Data
	err = json.Unmarshal([]byte(data), &da)
	if err != nil {
		fmt.Println(err)
	}

	//luu data lên mysql
	//insertData(db, da.Flights)

	//api tim kiem
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(da)
	})

	app.Get("/api/getFlight", func(c *fiber.Ctx) error {
		date := c.Query("date")
		sortby := c.Query("sortby")
		filterbyairline := c.Query("filterbyairline")

		query := "SELECT f.id, b.brand_code, f.flight_id, f.departure_timw, f.arrival_time, f.price FROM Flight f INNER JOIN Brand b ON f.brand_id = b.brand_id WHERE DATE(departure_timw) = ? "

		if filterbyairline != "" {
			query = query + " AND b.brand_code = '" + filterbyairline + "' "
		}

		switch sortby {
		case "price":
			query += " ORDER BY price "
		case "departure_time":
			query += " ORDER BY departure_timw ASC "
		case "brand_id":
			query += " ORDER BY brand_id ASC "
		default:
		}

		query += ""

		rows, err := db.Query(query, date)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		var flights []Flight
		for rows.Next() {
			var flight Flight
			err = rows.Scan(&flight.ID, &flight.Flight_id, &flight.Brand_code, &flight.Departure_time, &flight.Arrival_time, &flight.Price)
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
