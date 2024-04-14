package services

import (
	"OzonTech/internal/models"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/url"
)

func Valid(urlstring string) bool {
	_, err := url.ParseRequestURI(urlstring)
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
	return string(shortenedLink)
}

func AddToMap(link string, short string) error {
	_, ok := models.Inmemory[link]
	if ok {
		return errors.New("url is exist")
	} else {
		models.Inmemory[link] = short
		return nil
	}
}
