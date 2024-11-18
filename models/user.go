package models

import (
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id        int
	FirstName string `orm:"null"`
	LastName  string `orm:"null"`
	Email     string `orm:"null;unique"`
	Password  string `orm:"null"`
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

func GetUserById(id int) *Users {
	o := orm.NewOrm()
	user := Users{Id: id}
	o.Read(&user)
	return &user
}

func InsertOneUser(user Users) *Users {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))

	// get prepared statement
	i, _ := qs.PrepareInsert()

	var u Users

	// hash password
	user.Password, _ = hashPassword(user.Password)

	// Insert
	id, err := i.Insert(&user)
	if err == nil {
		// successfully inserted
		u = Users{Id: int(id)}
		err := o.Read(&u)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &u
}
func UpdateUser(user Users) *Users {
	o := orm.NewOrm()
	u := Users{Id: user.Id}
	var updatedUser Users

	// get existing user
	if o.Read(&u) == nil {

		// updated user
		// hash new password
		user.Password, _ = hashPassword(user.Password)

		u = user
		_, err := o.Update(&u)

		// read updated user
		if err == nil {
			// update successful
			updatedUser = Users{Id: user.Id}
			o.Read(&updatedUser)
		}
	}

	return &updatedUser
}

// DeleteUser deletes a user
func DeleteUser(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Users{Id: id})
	if err == nil {
		return true
	}
	return false
}
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
