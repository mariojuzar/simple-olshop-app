package service

import (
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/mariotj/olshop-app/api/entity/model"
	"github.com/mariotj/olshop-app/api/libraries/util"
	"io"
	"log"
	"os"
)

type DatabaseService struct {
	DB 				*gorm.DB
	IsInitialized 	bool
}

var databaseService DatabaseService

func Initialize() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DbUser, DbPassword, DbPort, DbHost, DbName := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		log.Fatal("Failed to init db: ", err)
	}
	db.LogMode(true)

	databaseService = DatabaseService{DB:db, IsInitialized:true}

	loadFromCsv(databaseService.DB)
}

/*
Load data from csv files into database
 */
func loadFromCsv(db *gorm.DB)  {
	loadCustomers(db)
	loadEmployees(db)
	loadShippingMethods(db)
	loadProducts(db)
	loadOrders(db)
	loadOrderDetails(db)
}

/*
Load customer data
 */
func loadCustomers(db *gorm.DB)  {
	csvFile, err := os.Open("file/Customers.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvFile)
	r.Comma = ';'

	if _, err := r.Read(); err != nil { //read header
		log.Fatal(err)
	}

	for {
		record, er := r.Read()
		if er == io.EOF {
			break
		}

		customer := model.Customer{
			ID:                  util.StrToUint(record[0]),
			CompanyName:         record[1],
			FirstName:           record[2],
			LastName:            record[3],
			BillingAddress:      record[4],
			City:                record[5],
			StateOrProvince:     record[6],
			ZipCode:             record[7],
			Email:               record[8],
			CompanyWebsite:      record[9],
			PhoneNumber:         record[10],
			FaxNumber:           record[11],
			ShipAddress:         record[12],
			ShipCity:            record[13],
			ShipStateOrProvince: record[14],
			ShipZipCode:         record[15],
			ShipPhoneNumber:     record[16],
		}

		db.Create(&customer)
	}
}

/*
Load employee data
 */
func loadEmployees(db *gorm.DB)  {
	csvFile, err := os.Open("file/Employees.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvFile)
	r.Comma = ';'

	if _, err := r.Read(); err != nil { //read header
		log.Fatal(err)
	}

	for {
		record, er := r.Read()
		if er == io.EOF {
			break
		}

		employee := model.Employee{
			ID:        util.StrToUint(record[0]),
			FirstName: record[1],
			LastName:  record[2],
			Title:     record[3],
			WorkPhone: record[4],
		}

		db.Create(&employee)
	}
}

/*
Load shipping methods data
 */
func loadShippingMethods(db *gorm.DB)  {
	csvFile, err := os.Open("file/ShippingMethods.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvFile)
	r.Comma = ';'

	if _, err := r.Read(); err != nil { //read header
		log.Fatal(err)
	}

	for {
		record, er := r.Read()
		if er == io.EOF {
			break
		}

		method := model.ShippingMethod{
			ID:             util.StrToUint(record[0]),
			ShippingMethod: record[1],
		}

		db.Create(&method)
	}
}

/*
Load products data
 */
func loadProducts(db *gorm.DB)  {
	csvFile, err := os.Open("file/Products.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvFile)
	r.Comma = ';'

	if _, err := r.Read(); err != nil { //read header
		log.Fatal(err)
	}

	for {
		record, er := r.Read()
		if er == io.EOF {
			break
		}

		product := model.Product{
			ID:          util.StrToUint(record[0]),
			ProductName: record[1],
			UnitPrice:   util.StrToFloat(record[2]),
			InStock:     util.StrToBool(record[3]),
		}

		db.Create(&product)
	}
}

/*
Load order data
 */
func loadOrders(db *gorm.DB) {
	csvFile, err := os.Open("file/Orders.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvFile)
	r.Comma = ';'

	if _, err := r.Read(); err != nil { //read header
		log.Fatal(err)
	}

	for {
		record, er := r.Read()
		if er == io.EOF {
			break
		}

		order := model.Order{
			ID:                  util.StrToUint(record[0]),
			CustomerId:          util.StrToUint(record[1]),
			EmployeeId:          util.StrToUint(record[2]),
			OrderDate:           util.StrToDate(record[3]),
			PurchaseOrderNumber: record[4],
			ShippingMethodId:    util.StrToUint(record[6]),
			FreightCharge:       util.StrToFloat(record[7]),
			Taxes:               util.StrToFloat(record[8]),
			PaymentReceived:     util.StrToBool(record[9]),
			Comment:             record[10],
		}

		if record[5] != "" {
			order.ShipDate = util.StrToNullTime(record[5], true)
		} else {
			order.ShipDate = util.StrToNullTime("00/00/0000", false)
		}

		db.Create(&order)
	}
}

/*
Load order details data
 */
func loadOrderDetails(db *gorm.DB)  {
	csvFile, err := os.Open("file/OrderDetails.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvFile)
	r.Comma = ';'

	if _, err := r.Read(); err != nil { //read header
		log.Fatal(err)
	}

	for {
		record, er := r.Read()
		if er == io.EOF {
			break
		}

		detail := model.OrderDetail{
			ID:        util.StrToUint(record[0]),
			OrderId:   util.StrToUint(record[1]),
			ProductId: util.StrToUint(record[2]),
			Quantity:  util.StrToInt(record[3]),
			UnitPrice: util.StrToFloat(record[4]),
			Discount:  util.StrToFloat(record[5]),
		}

		db.Create(&detail)
	}
}