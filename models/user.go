package models

type User struct {
	Id       int    `xorm:"not null INT(11)"`
	Username string `xorm:"not null CHAR(30)"`
	Email    string `xorm:"not null CHAR(50)"`
	Password string `xorm:"not null CHAR(32)"`
}
