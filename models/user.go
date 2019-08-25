package models

import (
	"cartoon-gin/common"
	"time"
)

type User struct {
	Id uint `json:"id"`
	Phone string `json:"phone"`
	NickName string `json:"nick_name"`
	CreatedAt common.MyTime `json:"created_at"`
	UpdatedAt common.MyTime `json:"updated_at"`
}

func AddUser(phone,nickname string) (id int64,err error) {
	db,_ := OpenDB()
	now := time.Now().Format("2006-01-02 15:04:05")
	rs,err := db.Exec("insert into users(phone,nick_name,created_at,updated_at) values (?,?,?,?)",phone,nickname,now,now)
	common.CheckError(err)
	return rs.LastInsertId()
}

func GetUsersByPhone(phone string) (users []User,err error) {
	db,_ := OpenDB()
	rows,err := db.Query("select * from users where phone=?",phone)
	common.CheckError(err)

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Phone, &user.NickName,&user.CreatedAt,&user.UpdatedAt)
		common.CheckError(err)
		//if user.DeletedAt.Valid {
		//	user.DeletedAt.Time.Format("2006-01-02 15:04:05")
		//	//= time.Time.Format(user.DeletedAt.Time,"2006-01-02 15:04:05")
		//}
		users = append(users, user)
	}
	return users,err
}
