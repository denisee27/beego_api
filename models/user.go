package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id        string     `orm:"pk;size(36)" valid:"uuid" json:"id,omitempty"`
	Name      string     `orm:"size(128)" validate:"required" json:"name,omitempty"`
	Email     string     `orm:"size(64);unique" validate:"required,email" json:"email,omitempty"`
	Password  *string    `orm:"size(64),min=6" json:"password,omitempty"`
	CreatedAt *time.Time `orm:"auto_now_add;type(datetime)" json:"created_at,omitempty"`
	UpdatedAt *time.Time `orm:"auto_now;type(datetime)" json:"updated_at,omitempty"`
}

var validate = validator.New()

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

func CreateUser(user Users) error {
	if err := validate.Struct(user); err != nil {
		return fmt.Errorf("%w", err)
	}
	if user.Password == nil {
		return fmt.Errorf("Password can't empty")
	}
	o := orm.NewOrm()
	user.Id = uuid.NewString()
	user.Password, _ = hashPassword(*user.Password)
	sql := `
		INSERT INTO users (id, name, email, password, created_at, updated_at)
		VALUES (?, ?, ?, ?, NOW(), NOW())
		RETURNING id
	`
	var insertedId string
	err := o.Raw(sql, user.Id, user.Name, user.Email, *user.Password).QueryRow(&insertedId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user Users) error {
	if err := validate.Struct(user); err != nil {
		return fmt.Errorf("wrong: %w", err)
	}
	o := orm.NewOrm()
	u := Users{Id: user.Id}
	if err := o.Read(&u); err != nil {
		return fmt.Errorf("user not found: %w", err)
	}
	if user.Email != u.Email {
		existing := Users{Email: user.Email}
		if err := o.Read(&existing, "Email"); err == nil {
			return fmt.Errorf("email already in use")
		}
	}
	if user.Password != nil {
		hashedPassword, err := hashPassword(*user.Password)
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

func hashPassword(password string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	hashed := string(bytes)
	return &hashed, nil

}
func IsEmailExists(email string) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(Users)).Filter("email", email).Exist()
}
