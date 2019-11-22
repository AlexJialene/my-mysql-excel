package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"strings"
	"time"
)

const (
	/*USERNAME = "lvji"
	PASSWORD = "L@#$j"*/
	USERNAME = "root"
	PASSWORD = "950426"
	NETWORK  = "tcp"
	//SERVER   = "192.168.32.88"
	SERVER   = "47.98.45.180"
	PORT     = 3306
	DATABASE = "nb-smork"
)

var mysql *gorm.DB

type Table struct {
	TableName    string
	TableComment string
	Columns      []Column
}
type Column struct {
	ColumnName    string
	ColumnComment string
}

var ch = make(chan int)

func main() {
	db, _ := openDb()
	start(db)
	i := 0
	for range ch {
		i += 1
		if i == 17 {
			fmt.Println("Generate Excel Done")
			break
		}
	}

}

func start(db *gorm.DB) {
	var tableNames []Table
	db.Table("information_schema.tables").Select("table_name , table_comment").Where("table_schema=?", DATABASE).Find(&tableNames)

	for _, v := range tableNames {
		var c []Column
		db.Table("information_schema.columns").Select("column_name , column_comment").Where("table_name=? and table_schema=?", v.TableName, DATABASE).Find(&c)
		v.Columns = c
		go GenExcel(v)
	}
}

func GenExcel(table Table) {
	fmt.Println("gen excel ...")
	file := excelize.NewFile()
	for i, v := range table.Columns {
		axis := AxisName(i + 1)
		file.SetCellValue("Sheet1", axis+strconv.Itoa(1), v.ColumnComment)
		file.SetCellValue("Sheet1", axis+strconv.Itoa(2), v.ColumnName)

	}
	as := file.SaveAs("./" + table.TableComment + "-" + table.TableName + ".xlsx")
	if as != nil {
		fmt.Println("gen Excel error")
	} else {
		ch <- 1
	}

}

//生成excel列名
func AxisName(i int) string {
	var str string
	for {
		if i <= 0 {
			break
		}
		i--
		i2 := i % 26
		str += string(i2 + 97)
		i = (i - i2) / 26

	}

	var result string
	for i := len(str); i > 0; i-- {
		result += string(str[i-1])

	}
	return strings.ToUpper(result)
}

func openDb() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("[ERROR] | connect mysql error: ", err)
		return
	}
	db.DB().SetConnMaxLifetime(100 * time.Second)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(16)
	db.Debug()
	db.LogMode(true)

	fmt.Println("[INFO] | connect mysql success")
	return
}
