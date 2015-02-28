package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
	"gopkg.in/qml.v1"
)

func main() {

	err := qml.Run(run)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	dbmap := initDb()
	defer dbmap.Db.Close()

	err = dbmap.TruncateTables()
	CheckErr(err, "TruncateTables failed")

}

func run() error {
	engine := qml.NewEngine()

	component, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)

	engine.On("quit", func() { os.Exit(0) })

	person := Person{}
	context := engine.Context()
	context.SetVar("person", &person)

	person.Root = window.Root()

	window.Show()
	window.Wait()
	return nil
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "contactcrmdb.bin")
	CheckErr(err, "sql.Open failed!")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(Person{}, "person").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed!")

	return dbmap
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
