package main

import (
	"time"

	_ "github.com/mattn/go-sqlite3"

	"gopkg.in/qml.v1"
)

type Person struct {
	Root           qml.Object
	Id             int64 `db:"person_id"`
	Created        int64
	FamilyName     string
	GivenName      string
	AdditionalName string
	NickName       string
	Title          string
	Role           string
}

func (p *Person) SavePerson(text qml.Object) {
	p.Created = time.Now().UnixNano()
	p.FamilyName = text.ObjectByName("familyName").String("text")
	p.GivenName = text.ObjectByName("givenName").String("text")
	p.AdditionalName = text.ObjectByName("additionalName").String("text")
	p.NickName = text.ObjectByName("nickName").String("text")
	p.Title = text.ObjectByName("title").String("text")
	p.Role = text.ObjectByName("role").String("text")

	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.Insert(&p)
	CheckErr(err, "DB Insert Failed")
	//fmt.Printf("%v: ", p)
}

func (p *Person) GetPerson(id int64) {

}

func (p *Person) GetPeople() []Person {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var people []Person
	_, err := dbmap.Select(&people, "select * from people order by people_id")
	CheckErr(err, "DB Select All Failed")
	return people
}
