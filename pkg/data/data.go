package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Db ...sql.Db
var Db *sql.DB

func init() {
	var err error
	// TODO : Please set your DB information.
	Db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=chatapp sslmode=disable")

	err = Db.Ping()
	if err != nil {
		println("DB connection failure !")
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt ...hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
