package database

import "fmt"

var connection string

func init() {
	connection = "postgres"
	fmt.Println("init dipanggil di database")
} // secara auto dieksekusi saat package ini diimport

func GetDB() string {
	return connection
}
