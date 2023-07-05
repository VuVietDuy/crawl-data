package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
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

	//url := "https://chuyenbay2.abay.vn/_WEB/ResultDom2/ResultDomAjax.aspx/GetFlights"
	//data := `{"input":"SGN-HAN-1-0-0-05Jul2023-","sortby":"abay-suggest","filterbyairline":"","defaultSort":1,"IsRewriteFare":0,"display":"base","edition":2,"cookieAll":{"CID":"164130","CD":"26Jun2023","LVPDI":"1510186","_gid":"GA1.2.882121964.1688344572","TCB":"1","_gcl_au":"1.1.1839823084.1688366182","SumAll":"14","LVPD":"1510186,0,0,1,0,2,14,4","RetDate":"","Client-Browser":"1","DepDate":"05-07-2023","VHP":"05/07/2023 07:00:09","VHPTD":"05JUL2023.1","AT":"4","StartPoint":"SGN","EndPoint":"HAN","Adult":"1","Child":"0","Infant":"0","DV":"05Jul2023","CDV":"6","tsv2":"2023-07-04T23:59:47.619Z"},"waytype":"OutBound","isAllowSearchReal":0}`
	//
	//// Tạo một payload từ dữ liệu JSON
	//payload := strings.NewReader(data)
	//
	//// Gửi yêu cầu POST
	//resp, err := http.Post(url, "application/json", payload)
	//if err != nil {
	//	fmt.Println("Lỗi trong quá trình gửi yêu cầu POST:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//// Đọc nội dung phản hồi
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Lỗi trong quá trình đọc phản hồi:", err)
	//	return
	//}
	//
	//// Hiển thị nội dung phản hồi
	//fmt.Println(string(body))
	//fmt.Println("Phản hồi từ server:", reflect.TypeOf(body))

	url := "https://chuyenbay2.abay.vn/_WEB/ResultDom2/ResultDomAjax.aspx/GetFlights"
	payload := []byte(`{"input":"SGN-HAN-1-0-0-05Jul2023-","sortby":"abay-suggest","filterbyairline":"","defaultSort":1,"IsRewriteFare":0,"display":"base","edition":2,"cookieAll":"{\"CID\":\"164130\",\"CD\":\"26Jun2023\",\"LVPDI\":\"1510186\",\"_gid\":\"GA1.2.882121964.1688344572\",\"TCB\":\"1\",\"_gcl_au\":\"1.1.1839823084.1688366182\",\"SumAll\":\"14\",\"LVPD\":\"1510186,0,0,1,0,2,14,4\",\"RetDate\":\"\",\"Client-Browser\":\"1\",\"DepDate\":\"05-07-2023\",\"VHP\":\"05/07/2023 07:00:09\",\"VHPTD\":\"05JUL2023.1\",\"AT\":\"4\",\"StartPoint\":\"SGN\",\"EndPoint\":\"HAN\",\"Adult\":\"1\",\"Child\":\"0\",\"Infant\":\"0\",\"DV\":\"05Jul2023\",\"CDV\":\"6\",\"tsv2\":\"2023-07-04T23:59:47.619Z\"}","waytype":"OutBound","isAllowSearchReal":0}`)

	// Tạo yêu cầu POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Lỗi trong quá trình tạo yêu cầu POST:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Gửi yêu cầu POST
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Lỗi trong quá trình gửi yêu cầu POST:", err)
		return
	}
	defer resp.Body.Close()

	// Đọc nội dung phản hồi
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Lỗi trong quá trình đọc phản hồi:", err)
		return
	}

	// Phân tích phản hồi thành đối tượng FlightData
	fmt.Println(string(body))
	var flightData Flight
	err = json.Unmarshal(body, &flightData)
	if err != nil {
		fmt.Println("Lỗi trong quá trình phân tích phản hồi:", err)
		return
	}

	// Hiển thị dữ liệu phản hồi
	fmt.Println("Dữ liệu phản hồi:", flightData)

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
			query = query + " AND brand_id = '" + filterbyairline + "' "
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
