package inputsources

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	mysql string = "mysql"
	pg    string = "postgres"
)

//- postgres
var lock = &sync.RWMutex{}
var postgresInstance *PostgresStruct

type PostgresStruct struct {
	db *sql.DB
}

func PostgresInstance() *sql.DB {
	if postgresInstance == nil {
		//- lock the thread and check
		lock.Lock()
		defer lock.Unlock()
		if postgresInstance == nil {
			fmt.Println("Create new POSTGRES instance now")
			//- dbType will be aroud mysql or postgres
			dsnBuilderString := DSNBuilder("pg")
			pi, err := sql.Open(`postgres`, dsnBuilderString)
			if err != nil {
				panic("Postgres db conenction is not successful")
			}
			postgresInstance = &PostgresStruct{
				db: pi,
			}
		} else {
			//- instance can reuse
			fmt.Println("POSTGRES Instance reuse")
		}

	} else {
		//- instance can reuse
		fmt.Println("POSTGRES Instance reuse")
	}
	return postgresInstance.db
}

//- mysql
var lock2 = &sync.RWMutex{}
var mysqlInstance *MysqlStruct

type MysqlStruct struct {
	db *sql.DB
}

func MysqlInstance() *sql.DB {
	fmt.Println("mysql instance == nil 1", mysqlInstance == nil)
	if mysqlInstance == nil {
		//- lock the thread and check
		lock2.Lock()
		defer lock2.Unlock()
		fmt.Println("mysql instance == nil 2", mysqlInstance == nil)
		if mysqlInstance == nil {
			fmt.Println("Create new MYSQL instance now")
			//- dbType will be aroud mysql or postgres
			dsnBuilderString := DSNBuilder("mysql")
			fmt.Println("dsn builder", dsnBuilderString)
			mi, err := sql.Open(`mysql`, dsnBuilderString)
			if err != nil {
				panic("Postgres db conenction is not successful")
			}
			mysqlInstance = &MysqlStruct{
				db: mi,
			}
		} else {
			//- instance can reuse
			fmt.Println("MYSQL Instance reuse")
		}

	} else {
		//- instance can reuse
		fmt.Println("MYSQL Instance reuse")
	}
	return mysqlInstance.db
}
