package database

import "github.com/gohouse/gorose"

var Database *gorose.Connection

func GetOrm() *gorose.Session  {
	return Database.NewSession()
}