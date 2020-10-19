package models

import (
	"bee06/db_mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type  User struct {
	Id int				`form:"id"`
	Phone string		`form:"phone"`
	Password string		`form:"password"`
}
func (u User) SaveUser() (int64,error){
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordBytes)

	row, err := db_mysql.DB.Exec("insert into user (phone, password)" + "values(?,?)",u.Phone,u.Password)
	if err !=nil{
		fmt.Println(err.Error())
		return -1,err
	}
	id,err :=row.RowsAffected()
	if err !=nil{
		fmt.Println(err.Error())
		return -1,err
	}
	return id,nil
}
func (u User)QueryUser()(*User,error){
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordBytes)
	row := db_mysql.DB.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone,u.Password)
	err :=row.Scan(&u.Phone)
	if err !=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return &u,nil
}
