package services

import (
	"OzonTech/internal/models"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
)

func Valid(urlstring string) bool {
	rez, err := url.ParseRequestURI(urlstring)
	fmt.Println("Valid rez are: " + rez.String())
	return err == nil
}

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

func HashLink(origLink string) string {
	h := sha256.New()
	h.Write([]byte(origLink))
	hashedLink := h.Sum(nil)
	encodedLink := base64.URLEncoding.EncodeToString(hashedLink)
	truncatedLink := encodedLink[:10]
	var shortenedLink []byte
	for _, b := range truncatedLink {
		i := int(b) % len(alphabet)
		shortenedLink = append(shortenedLink, alphabet[i])
	}
	fmt.Println("HashLink rez are: " + string(shortenedLink))
	return string(shortenedLink)
}

func AddToMap(link string, short string) error {
	_, ok := models.Inmemory[link]
	if ok {
		return errors.New("url is already exist")
	} else {
		models.Inmemory[short] = link
		fmt.Println(models.Inmemory[short])
		return nil
	}
}

func DbConnection() *sql.DB {
	connStr := "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, errsql := sql.Open("postgres", connStr)
	if errsql != nil {
		panic(errsql)
	}
	fmt.Println("Connected to database")
	return db
}
