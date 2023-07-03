package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Flight struct {
	//ID             string `json"`
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

	// ket noi voi databse
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

	//app := fiber.New()

	// lưu data lên mysql
	date := "2023-07-03 "
	data := `[
	 {
	   "E": "VN",
	   "F": "VN271",
	   "G": "20:30",
	   "H": "22:55",
	   "I": true,
	   "J": "909,000",
	   "K": false,
	   "L": "1,029,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-20:30|VN271-20:30-22:55,10kg-23kg-1029000,E-787-EPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 20:30",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-20:30%7CVN271-20:30-22:55,10kg-23kg-1029000,E-787-EPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN267",
	   "G": "22:00",
	   "H": "00:15",
	   "I": true,
	   "J": "909,000",
	   "K": false,
	   "L": "1,029,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-22:00|VN267-22:00-00:15,10kg-23kg-1029000,E-787-EPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 22:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-22:00%7CVN267-22:00-00:15,10kg-23kg-1029000,E-787-EPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN7219",
	   "G": "22:30",
	   "H": "00:45",
	   "I": true,
	   "J": "909,000",
	   "K": false,
	   "L": "1,029,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-22:30|VN7219-22:30-00:45,10kg-23kg-1029000,E-321-EPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 22:30",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-22:30%7CVN7219-22:30-00:45,10kg-23kg-1029000,E-321-EPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "BL",
	   "F": "VN6025",
	   "G": "06:20",
	   "H": "08:50",
	   "I": true,
	   "J": "949,000",
	   "K": false,
	   "L": "1,069,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-06:20|VN6025-06:20-08:50,7kg-23kg-1069000,N-320-NPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120",
	   "O": "Chọn Pacific 06:20",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-06:20%7CVN6025-06:20-08:50,7kg-23kg-1069000,N-320-NPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": false,
	   "S": false
	 },
	 {
	   "E": "BL",
	   "F": "VN6039",
	   "G": "18:45",
	   "H": "21:00",
	   "I": true,
	   "J": "949,000",
	   "K": false,
	   "L": "1,069,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-18:45|VN6039-18:45-21:00,7kg-23kg-1069000,N-320-NPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120",
	   "O": "Chọn Pacific 18:45",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-18:45%7CVN6039-18:45-21:00,7kg-23kg-1069000,N-320-NPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": false,
	   "S": false
	 },
	 {
	   "E": "BL",
	   "F": "VN6021",
	   "G": "21:10",
	   "H": "23:30",
	   "I": true,
	   "J": "949,000",
	   "K": false,
	   "L": "1,069,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-21:10|VN6021-21:10-23:30,7kg-23kg-1069000,N-320-NPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120",
	   "O": "Chọn Pacific 21:10",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-21:10%7CVN6021-21:10-23:30,7kg-23kg-1069000,N-320-NPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": false,
	   "S": false
	 },
	 {
	   "E": "BL",
	   "F": "VN6045",
	   "G": "21:45",
	   "H": "23:55",
	   "I": true,
	   "J": "949,000",
	   "K": false,
	   "L": "1,069,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-21:45|VN6045-21:45-23:55,7kg-23kg-1069000,N-320-NPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120",
	   "O": "Chọn Pacific 21:45",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-21:45%7CVN6045-21:45-23:55,7kg-23kg-1069000,N-320-NPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": false,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN223",
	   "G": "23:00",
	   "H": "01:25",
	   "I": true,
	   "J": "959,000",
	   "K": false,
	   "L": "1,029,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-23:00|VN223-23:00-01:25,10kg-23kg-1029000,E-321-EPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD70",
	   "O": "Chọn VNA 23:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-23:00%7CVN223-23:00-01:25,10kg-23kg-1029000,E-321-EPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD70&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN243",
	   "G": "06:00",
	   "H": "08:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-06:00|VN243-06:00-08:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 06:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-06:00%7CVN243-06:00-08:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN245",
	   "G": "08:00",
	   "H": "10:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-08:00|VN245-08:00-10:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 08:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-08:00%7CVN245-08:00-10:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN209",
	   "G": "09:00",
	   "H": "11:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-09:00|VN209-09:00-11:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 09:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-09:00%7CVN209-09:00-11:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN247",
	   "G": "10:00",
	   "H": "12:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-10:00|VN247-10:00-12:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 10:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-10:00%7CVN247-10:00-12:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN211",
	   "G": "11:00",
	   "H": "13:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-11:00|VN211-11:00-13:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 11:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-11:00%7CVN211-11:00-13:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN251",
	   "G": "12:00",
	   "H": "14:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-12:00|VN251-12:00-14:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 12:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-12:00%7CVN251-12:00-14:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN213",
	   "G": "13:00",
	   "H": "15:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-13:00|VN213-13:00-15:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 13:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-13:00%7CVN213-13:00-15:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN253",
	   "G": "14:00",
	   "H": "16:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-14:00|VN253-14:00-16:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 14:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-14:00%7CVN253-14:00-16:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN215",
	   "G": "15:00",
	   "H": "17:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-15:00|VN215-15:00-17:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 15:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-15:00%7CVN215-15:00-17:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN261",
	   "G": "16:55",
	   "H": "19:10",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-16:55|VN261-16:55-19:10,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 16:55",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-16:55%7CVN261-16:55-19:10,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN217",
	   "G": "17:00",
	   "H": "19:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-17:00|VN217-17:00-19:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 17:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-17:00%7CVN217-17:00-19:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN257",
	   "G": "17:10",
	   "H": "19:25",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-17:10|VN257-17:10-19:25,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 17:10",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-17:10%7CVN257-17:10-19:25,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN259",
	   "G": "18:00",
	   "H": "20:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-18:00|VN259-18:00-20:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 18:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-18:00%7CVN259-18:00-20:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN265",
	   "G": "18:30",
	   "H": "20:45",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-18:30|VN265-18:30-20:45,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 18:30",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-18:30%7CVN265-18:30-20:45,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN219",
	   "G": "19:00",
	   "H": "21:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-19:00|VN219-19:00-21:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 19:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-19:00%7CVN219-19:00-21:15,10kg-23kg-1219000,T-787-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN221",
	   "G": "21:00",
	   "H": "23:15",
	   "I": true,
	   "J": "1,099,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-21:00|VN221-21:00-23:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 21:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-21:00%7CVN221-21:00-23:15,10kg-23kg-1219000,T-359-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN205",
	   "G": "05:00",
	   "H": "07:15",
	   "I": true,
	   "J": "1,149,000",
	   "K": false,
	   "L": "1,219,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-05:00|VN205-05:00-07:15,10kg-23kg-1219000,T-321-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD70",
	   "O": "Chọn VNA 05:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-05:00%7CVN205-05:00-07:15,10kg-23kg-1219000,T-321-TPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD70&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "BL",
	   "F": "VN6009",
	   "G": "09:25",
	   "H": "11:45",
	   "I": true,
	   "J": "1,179,000",
	   "K": false,
	   "L": "1,299,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-09:25|VN6009-09:25-11:45,7kg-23kg-1299000,Q-320-QPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120",
	   "O": "Chọn Pacific 09:25",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-09:25%7CVN6009-09:25-11:45,7kg-23kg-1299000,Q-320-QPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": false,
	   "S": false
	 },
	 {
	   "E": "BL",
	   "F": "VN6015",
	   "G": "15:25",
	   "H": "17:40",
	   "I": true,
	   "J": "1,179,000",
	   "K": false,
	   "L": "1,299,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-15:25|VN6015-15:25-17:40,7kg-23kg-1299000,Q-320-QPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120",
	   "O": "Chọn Pacific 15:25",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-15:25%7CVN6015-15:25-17:40,7kg-23kg-1299000,Q-320-QPXVNF9-OperateAirlineBL-SeatRemain9-1549000-L-9-LPXVNF9-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": false,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN207",
	   "G": "07:00",
	   "H": "09:15",
	   "I": true,
	   "J": "1,199,000",
	   "K": false,
	   "L": "1,319,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-07:00|VN207-07:00-09:15,10kg-23kg-1319000,R-359-RPXVNFP-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 07:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-07:00%7CVN207-07:00-09:15,10kg-23kg-1319000,R-359-RPXVNFP-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN255",
	   "G": "16:00",
	   "H": "18:15",
	   "I": true,
	   "J": "1,199,000",
	   "K": false,
	   "L": "1,319,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-16:00|VN255-16:00-18:15,10kg-23kg-1319000,R-787-RPXVNFP-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 16:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-16:00%7CVN255-16:00-18:15,10kg-23kg-1319000,R-787-RPXVNFP-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 },
	 {
	   "E": "VN",
	   "F": "VN263",
	   "G": "20:00",
	   "H": "22:15",
	   "I": true,
	   "J": "1,649,000",
	   "K": false,
	   "L": "1,769,000",
	   "C": "OutBound",
	   "N": "HAN-SGN-05Jul2023-20:00|VN263-20:00-22:15,10kg-23kg-1769000,Q-787-QPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120",
	   "O": "Chọn VNA 20:00",
	   "P": "",
	   "Q": "",
	   "T": "/_Web/ResultDom/usrDesktop/FlightDetail.aspx?fi=HAN-SGN-05Jul2023-20:00%7CVN263-20:00-22:15,10kg-23kg-1769000,Q-787-QPXVNF-OperateAirlineVN-SeatRemain9-1959000-L-9-LPXVNF-MD120&input=HAN-SGN-1-0-0-05Jul2023-",
	   "R": true,
	   "S": false
	 }
	]`

	var flights []Flight
	err = json.Unmarshal([]byte(data), &flights)
	if err != nil {
		fmt.Println(err)
	}

	for _, flight := range flights {
		_, err = db.Exec("INSERT INTO flight(brand_id, flight_id, departure_time, arrival_time, price) VALUES(?, ?, ?, ?, ?)", flight.Brand_id, flight.Flight_id, date+flight.Departure_time+":00", date+flight.Arrival_time+":00", flight.Price)
		if err != nil {
			log.Fatal(err)
		}

	}

	//app.Get("/api", func(c *fiber.Ctx) error {
	//
	//	//body := resp.Body()
	//
	//	// Xử lý phản hồi từ API
	//
	//	// Trả về kết quả từ API
	//	return c.SendString(data)
	//})
	//
	//err = app.Listen(":3000")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
