package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	UserName = "coin"
	PassWord = "LINUX2020"
	NetWork  = "tcp"
	Server   = "rocc.top"
	Port     = 3306
	Database = "coin"
)

type Price struct {
	Time  time.Time `json:"time"`
	Price float64   `json:"price"`
}

func CoinMySQLData(coinPrice map[string]float64) {

	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		UserName, PassWord, NetWork, Server, Port, Database)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println("connection to mysql failed:", err)
		return
	}

	defer DB.Close()
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	InsertData(DB, coinPrice)

}

func InsertData(DB *sql.DB, coinPrice map[string]float64) {
	//CST := time.FixedZone("CST", 8 * 3600)

	result, err := DB.Exec("insert INTO price(time,eth,btc,doge) values(?,?,?,?)", time.Now(), coinPrice["ethusdt"], coinPrice["btcusdt"], coinPrice["dogeusdt"])

	if err != nil {
		log.Printf("Insert data failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId() // 获取插入数据的自增 ID
	if err != nil {
		log.Printf("Get insert id failed,err:%v", err)
		return
	}
	log.Println("Insert data id:", lastInsertID)

	rowsaffected, err := result.RowsAffected() // 通过 RowsAffected 获取受影响的行数
	if err != nil {
		log.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	log.Println("Affected rows:", rowsaffected)
}
