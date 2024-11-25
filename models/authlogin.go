package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthLogin(request *LoginRequest) (response map[string]interface{}, err error) {
	email := request.Email
	password := request.Password

	if len(email) <= 0 || len(password) <= 0 {
		return nil, errors.New("email or password is empty")
	}
	o := orm.NewOrm()
	user := &Users{Email: email}
	find := o.Read(user, "email")
	if find != nil {
		return nil, errors.New("email is not existing")
	}
	pass := CheckPasswordHash(password, *user.Password)
	if !pass {
		return nil, errors.New("password is wrong")
	}
	result, err := GenerateToken(user, user.Id, 48)
	if err != nil {
		return nil, errors.New("Error Generate Token")
	} else {
		return result, nil
	}

}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
