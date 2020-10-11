package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init()  {

	//1.读取conf配置信息
	config := beego.AppConfig

	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//2.组织链接数据库的字符串
	connUrl := dbUser + ":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	//fmt.Println(connUrl)
	//	//fmt.Println(dbDriver)
	//3.链接数据库
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil{
		fmt.Println(err.Error())
		panic("数据库连接错误，请检查配置")
	}
	//4.为全局变量赋值
	DB = db
}
//func InserUser(user models.User)(int64,error){
//	//将用户密码进行hash脱敏，使用md5计算密码hash值，并储存hash值
//	hashMd5 := md5.New()
//	hashMd5.Write([]byte(user.Password))
//	bytes := hashMd5.Sum([]byte("nil"))
//	user.Password = hex.EncodeToString(bytes)
//	fmt.Println("将要保存的用户名:",user.Phone,"密码:",user.Password)
//	result,err :=DB.Exec("insert into user (Phone,password)values (?,?)",user.Phone,user.Password)
//	if err != nil {
//		return -1,err
//	}
//	id, err := result.RowsAffected()
//	if err != nil{
//		return  -1,err
//	}
//
//	return id,nil
//}
//func QueryUser(){
//	DB.QueryRow("select*from")
//}