package db

import (
	"OzonTech/internal/services"
)

type Psql struct {
	Id        int
	Name      string
	Shortname string
}

var rezult = make([]Psql, 0)

func AddToDatabase(name string, shortname string) error {
	d := services.DbConnection()
	defer d.Close()
	rez, err := d.Query("insert into ")
	if err != nil {
		return err
	}
	defer rez.Close()
	for rez.Next() {
		var p Psql
		err := rez.Scan(&p.Id, &p.Name, &p.Shortname)
		if err != nil {
			return err
		}
		rezult = append(rezult, p)
	}
	return nil
}

func GetFromDatabase(name string) string {
	d := services.DbConnection()
	defer d.Close()
	rez, err := d.Query("Select * from urls")
	if err != nil {

	}
	defer rez.Close()
	for rez.Next() {
		var p Psql
		err := rez.Scan(&p.Id, &p.Name, &p.Shortname)
		if err != nil {

		}
		rezult = append(rezult, p)
	}
	return ""
}
