package models

import "strconv"

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;comment:用户手机号"`
	Password string `json:"-" gorm:"not null;comment:用户密码"` //密码不返回给前端
	Timestamps
	SoftDeletes
}

func (User) TableName() string {
	return "user"
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
