package users

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"

	"github.com/PonyWilliam/sqld"
)

type Person struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
}

var Db *sqld.MySQL_D
var User []Person
var slat string

func init() {
	slat = "54088"
	//连接到数据库
	Db = sqld.SQL_init()
	Db.SetType("mysql")
	err, _ := Db.Connect("sh-cynosdbmysql-grp-2o1mkprk.sql.tencentcdb.com", 21270, "root", "Xiaowei123", "iotusers")
	if err != nil {
		log.Fatal(err)
	}
}

func Encode(pass string) string {
	//传入一个string，返回一个加密后的string
	temp := slat + pass
	temp += (temp + slat)
	h := md5.New()
	h.Write([]byte(temp))
	temp = hex.EncodeToString(h.Sum(nil))
	temp += slat + "tjw666"
	return hex.EncodeToString(h.Sum(nil))
}

func FindUser(username string) bool {
	User = nil
	_ = Db.Select(&User, "select * from person where username = ?", username)
	if User != nil {
		return true
	}
	return false
}

func Registry(username string, password string, phoneNum string) error {
	if FindUser(username) {
		return errors.New("已有同名账号，请换一个用户名注册")
	}
	//注册使用
	key := []string{"username", "password", "phone"}
	val := []string{username, Encode(password), phoneNum}
	err, _ := Db.Insert("person", key, val)
	return err
}
func Login(username string, password string) (bool, []Person) {
	if !FindUser(username) {
		return false, nil
	}
	//有，判断下账号密码
	if User[0].Password == Encode(password) {
		//密码正确
		return true, User
	} else {
		return false, nil
	}
}
