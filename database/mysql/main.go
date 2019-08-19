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

type Error string

func (e Error) Error() string {
	return string(e)
}

type Passanger struct {
	Uid        uint
	Gid        uint
	Pwd        []byte
	NickName   string
	SignUpTime time.Time
}

type Group struct {
	Gid        uint
	Name       string
	Note       string
	CreateTime time.Time
}

func (g *Group) String() string {
	s := fmt.Sprintf("{#%d %s, %s}", g.Gid, g.Name, g.Note)
	return s
}

func (p *Passanger) String() string {
	s := fmt.Sprintf("[$%d %s]", p.Uid, p.NickName)
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

	p := GetPassenger(db, 1)
	if p != nil {
		fmt.Println(p)
	}
	p.Uid = 0
	p.NickName = "Jack"
	p.Save(db)
	if p != nil {
		fmt.Println("New User:", p)
	}

	g := GetGroup(db, 1)
	if g != nil {
		fmt.Println(g)
	} else {
		fmt.Println("Group not found")
	}
	g.AddMember(db, p.Uid)
	members := g.GetMembers(db)
	for i, m := range members {
		fmt.Printf("Member %d: %v\n", i, m)

	}
	fmt.Printf("Try to delete %v\n", p)
	p.Delete(db)
	members = g.GetMembers(db)
	for i, m := range members {
		fmt.Printf("Member %d: %v\n", i, m)

	}
	GetAllUserAndGroup(db)

}

func GetAllUserAndGroup(db *sql.DB) {
	fmt.Println("UID", "name", "GID")
	rows, err := db.Query("SELECT u.uid, u.nickname, g.gid FROM FT_passanger AS u LEFT JOIN FT_user_ext_group AS g ON u.uid = g.uid")
	if err != nil {
		log.Printf("Fail to get passanger, %v", err)
		return
	}
	for rows.Next() {
		var uid uint
		var name string
		var gid uint
		rows.Scan(&uid, &name, &gid)
		fmt.Println(uid, name, gid)

	}

}
func GetPassenger(db *sql.DB, uid uint) *Passanger {

	rows, err := db.Query("select uid,gid,pwd_hash,nickname,signup_time from FT_passanger where uid = ?", uid)
	defer rows.Close()

	if err != nil {
		log.Printf("Fail to get passanger, %v", err)
		return nil
	}
	if rows.Next() {
		p := &Passanger{}
		rows.Scan(&p.Uid, &p.Gid, &p.Pwd, &p.NickName, &p.SignUpTime)
		return p
	} else {
		return nil
	}

}

func GetGroup(db *sql.DB, gid uint) *Group {

	rows, err := db.Query("select `gid`,`name`,`note` from  FT_group where gid = ?", gid)

	if err != nil {
		log.Printf("Fail to get passanger, %v", err)
		return nil
	}
	defer rows.Close()
	if rows.Next() {
		g := &Group{}
		rows.Scan(&g.Gid, &g.Name, &g.Note)
		return g
	} else {
		return nil
	}

}

// GetGroups 获取用户所在的组,
func (u *Passanger) GetGroups(db *sql.DB) (groups []*Group) {

	if u == nil || u.Uid == 0 {

		return

	}
	rows, err := db.Query("select `gid` from  FT_user_ext_group where uid = ?", u.Uid)

	if err != nil {
		log.Printf("Fail to delete, %v", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var gid uint
		rows.Scan(&gid)
		g := GetGroup(db, gid)
		if g != nil {
			groups = append(groups, g)

		}
	}
	return groups
}
func (u *Passanger) Delete(db *sql.DB) error {

	if u == nil || u.Uid == 0 {

		return Error("Nil user")

	}
	groups := u.GetGroups(db)
	for _, g := range groups {
		err := g.DelMember(db, u.Uid)
		if err != nil {
			log.Printf("delete from group fail %v", err)
		}
	}

	_, err := db.Exec("delete from FT_passanger where uid = ? ", &u.Uid)

	if err != nil {
		log.Printf("Fail to delete, %v", err)
		return Error("delete user failed")
	}

	return nil
}
func (u *Passanger) Save(db *sql.DB) error {

	if u.Uid != 0 {

		rows, err := db.Query("select uid from FT_passanger where uid = ?", u.Uid)

		if err != nil {
			log.Printf("Fail to get passanger, %v", err)
			return Error(fmt.Sprintf("Fail to get passanger, %v", err))
		}
		defer rows.Close()
		if rows.Next() {
			_, err := db.Exec("update FT_passanger set pwd_hash=?,gid=?, nickname = ? where uid = ?", &u.Gid, &u.Pwd, &u.NickName, &u.Uid)

			if err != nil {
				log.Printf("Fail to update, %v", err)
				return Error("Fail to update")
			}
			return nil

		}
	}

	r, err := db.Exec("insert into FT_passanger (gid,pwd_hash,nickname) values(?,?,?)", &u.Gid, &u.Pwd, &u.NickName)
	newid, _ := r.LastInsertId()
	u.Uid = uint(newid)

	if err != nil {
		log.Printf("Fail to insert, %v", err)
		return Error("Save user failed")
	}

	return nil
}
func (u *Passanger) AddToGroup(db *sql.DB, gid uint) error {
	rows, err := db.Query("select name from FT_group where gid = ?", gid)
	defer rows.Close()

	if err != nil {
		log.Printf("Fail to query, %v", err)
		return err
	}
	var gName string
	if rows.Next() {
		err := rows.Scan(&gName)
		if err != nil {
			return err
		}
	} else {
		return Error("Group not found")
	}
	_, err = db.Exec("insert into FT_user_ext_group values (?,?)", u.Uid, gid)
	if err != nil {
		log.Printf("Fail to insert , %v", err)
		return Error("Fail to insert ")
	}

	return nil
}
func (g *Group) DelMember(db *sql.DB, uid uint) error {
	_, err := db.Exec("delete from FT_user_ext_group where uid = ? and gid = ?", uid, g.Gid)
	if err != nil {
		log.Printf("Fail to del member, %v", err)
		return Error("Fail to del member")
	}
	return nil
}

func (g *Group) AddMember(db *sql.DB, uid uint) error {
	_, err := db.Exec("insert into FT_user_ext_group values (?,?)", uid, g.Gid)
	if err != nil {
		log.Printf("Fail to add member, %v", err)
		return Error("Fail to add member")
	}

	return nil
}

func (g *Group) GetMembers(db *sql.DB) (members []*Passanger) {

	uidset, err := db.Query("select uid from FT_user_ext_group where gid = ?", g.Gid)
	defer uidset.Close()

	if err != nil {
		log.Printf("Fail to find group, %v", err)
		return nil
	}
	var uid uint
	for uidset.Next() {
		err := uidset.Scan(&uid)
		if err != nil {
			return nil
		}
		p := GetPassenger(db, uid)
		if p == nil {
			log.Printf("Fail to get user %d\n", uid)
		}
		members = append(members, p)

	}

	return members
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
