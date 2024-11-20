package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id        string     `orm:"pk;size(36)" valid:"uuid" json:"id,omitempty"`
	Name      string     `orm:"size(128)" validate:"required" json:"name,omitempty"`
	Email     string     `orm:"size(64);unique" validate:"required,email" json:"email,omitempty"`
	Password  string     `orm:"size(64)" validate:"required,min=6" json:"password,omitempty"`
	CreatedAt *time.Time `orm:"auto_now_add;type(datetime)" json:"created_at,omitempty"`
	UpdatedAt *time.Time `orm:"auto_now;type(datetime)" json:"updated_at,omitempty"`
}

func init() {
	orm.RegisterModel(new(Users))
}
func GetAllUsers() []*Users {
	o := orm.NewOrm()
	var users []*Users
	o.QueryTable(new(Users)).All(&users)
	return users
}

func GetUserById(ids string) *Users {
	o := orm.NewOrm()
	user := Users{Id: ids}
	if err := o.Read(&user); err == orm.ErrNoRows {
		return &Users{}
	}
	return &user
}

func CreateUser(user Users) (*Users, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
	i, _ := qs.PrepareInsert()
	var u Users
	uid := uuid.New()
	user.Id = uid.String()
	user.Password, _ = hashPassword(user.Password)
	id, err := i.Insert(&user)
	if err == nil {
		u = Users{Id: string(id)}
		err := o.Read(&u)
		if err == orm.ErrNoRows {
			return nil, err
		}
	} else {
		return nil, err
	}
	return &u, nil
}

func UpdateUser(user Users) error {
	o := orm.NewOrm()
	u := Users{Id: user.Id}
	if err := o.Read(&u); err != nil {
		return fmt.Errorf("user not found: %w", err)
	}
	if user.Password != "" {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = hashedPassword
	}
	if _, err := o.Update(&user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func DeleteUser(ids string) error {
	o := orm.NewOrm()
	getId := GetUserById(ids)
	if getId == nil {
		return errors.New("user not found")
	}
	if _, err := o.Delete(&Users{Id: ids}); err == nil {
		return nil
	} else {
		return err
	}
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func IsEmailExists(email string) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(Users)).Filter("email", email).Exist()
}
