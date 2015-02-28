package main

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"gopkg.in/qml.v1"
)

type Person struct {
	Root           qml.Object `db:"-"`
	Id             int64      `db:"person_id"`
	Created        int64
	FamilyName     string
	GivenName      string
	AdditionalName string
	NickName       string
	Title          string
	Role           string
}

func (p *Person) SavePerson(text qml.Object) {
	//log.Println("Trying to save!")

	// Fetch data from qml form
	p.Created = time.Now().UnixNano()
	p.FamilyName = text.ObjectByName("familyName").String("text")
	p.GivenName = text.ObjectByName("givenName").String("text")
	p.AdditionalName = text.ObjectByName("additionalName").String("text")
	p.NickName = text.ObjectByName("nickName").String("text")
	p.Title = text.ObjectByName("title").String("text")
	p.Role = text.ObjectByName("role").String("text")
	log.Println(p.Created, p.FamilyName, p.GivenName, p.AdditionalName, p.NickName, p.Title, p.Role)

	// Initialize DB Connection.
	dbmap := initDb()
	defer dbmap.Db.Close()

	// Insert Data
	err := dbmap.Insert(p)
	CheckErr(err, "DB Insert Failed: ")

	// Count All Rows
	count, err := dbmap.SelectInt("select count(*) from person")
	CheckErr(err, "select count(*) failed")
	log.Println("Rows after inserting:", count)

	// Fetch all rows
	var person []Person
	_, err = dbmap.Select(&person, "select * from person")
	CheckErr(err, "Select failed")
	log.Println("All rows:")
	for x, p := range person {
		log.Printf("    %d: %v\n", x, p)
	}

	log.Println("Everything must have worked!")

	//fmt.Printf("%v: ", p)
}

func (p *Person) GetPerson(id int64) {

}

func (p *Person) GetPeople() []Person {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var person []Person
	_, err := dbmap.Select(&person, "select * from person order by people_id")
	CheckErr(err, "DB Select All Failed")
	return person
}
