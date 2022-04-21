package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectToSQL() *sql.DB {
	//	cfg := mysql.Config{
	//		User:   "anurag",
	//		Passwd: "anur@12",
	//		Net:    "tcp",
	//		Addr:   "127.0.0.1:3306",
	//		DBName: "CarDealership",
	//	}

	// get a database handle
	db, err := sql.Open("mysql", "anurag:anur@12@tcp(0.0.0.0:3307)/carDealership")
	if err != nil {
		log.Println(err)
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
	}

	log.Println("Connected!")

	return db
}
