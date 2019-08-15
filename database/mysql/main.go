package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Passanger struct {
	Uid        uint
	Pwd        []byte
	NickName   string
	SignUpTime time.Time
}

func (p *Passanger) String() string {
	s := fmt.Sprintf("uid:%d %s %x %v", p.Uid, p.NickName, p.Pwd, p.SignUpTime.Format("Jan 2 15:04:05 2006"))
	return s
}

func main() {
	uri := readUri()

	db, err := sql.Open("mysql", uri)
	defer db.Close()
	if err != nil {
		log.Printf("Fail to open connection, %v", err)
		return
	}
	query(db)
	insert(db)
	fmt.Println("\nAfter insert")
	query(db)

	update(db)
	fmt.Println("\nAfter update")
	query(db)

	delete(db)
	fmt.Println("\nAfter delete")
	query(db)
}
func query(db *sql.DB) {
	rows, err := db.Query("select uid,pwd_hash,nickname,signup_time from FT_passanger")
	defer rows.Close()

	if err != nil {
		log.Printf("Fail to query, %v", err)
		return
	}
	var p Passanger
	for rows.Next() {

		rows.Scan(&p.Uid, &p.Pwd, &p.NickName, &p.SignUpTime)
		fmt.Println(p.String())

	}

}

func insert(db *sql.DB) {

	p := &Passanger{
		Uid:        2,
		Pwd:        []byte("AGKEFAEOAFDWEF"),
		NickName:   "mike",
		SignUpTime: time.Time{},
	}
	_, err := db.Exec("insert into FT_passanger (uid,pwd_hash,nickname) values(?,?,?)", &p.Uid, &p.Pwd, &p.NickName)

	if err != nil {
		log.Printf("Fail to insert, %v", err)
		return
	}

}
func update(db *sql.DB) {

	p := &Passanger{
		Uid:        2,
		Pwd:        []byte("AGKEFAEOAFDWEF"),
		NickName:   "jack",
		SignUpTime: time.Time{},
	}
	_, err := db.Exec("update FT_passanger set nickname = ? where uid = ?", &p.NickName, &p.Uid)

	if err != nil {
		log.Printf("Fail to update, %v", err)
		return
	}

}
func delete(db *sql.DB) {

	p := &Passanger{
		Uid:        2,
		Pwd:        []byte("AGKEFAEOAFDWEF"),
		NickName:   "jack",
		SignUpTime: time.Time{},
	}
	_, err := db.Exec("delete from FT_passanger where uid = ?", &p.Uid)

	if err != nil {
		log.Printf("Fail to delete, %v", err)
		return
	}

}

// readUri like user:pwd@tcp(127.0.0.1:6603)/database?parseTime=true&loc=Asia%2FShanghai
func readUri() string {

	f, err := os.OpenFile("mysql_uri.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to read config file, %v", err)
		return ""
	}
	buf := [256]byte{}
	n, err := f.Read(buf[:])
	return strings.Trim(string(buf[:n]), "\n")

}
