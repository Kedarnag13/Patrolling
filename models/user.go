package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id           int
	FirstName    string `json:"firstname" valid:"Required;Alpha"`
	LastName     string `json:"lastname" valid:"Required;Alpha"`
	Email        string `json:"email" valid:"Required;Email" orm:"unique"`
	MobileNumber string `json:"mobile_number" valid:"Required" orm:"unique"`
	Designation  string `json:"designation" valid:"Alpha"`
}

type Sessions struct {
	Id          int
	User        *Users   `json:"user_id" orm:"rel(fk)"`
	DeviseToken *Devises `json:"devise_token" orm:"rel(fk)"`
}

type Devises struct {
	Id    int
	Token string `orm:"pk"`
}

// var (
// 	UserList map[string]*User
// )

// func init() {
// 	UserList = make(map[string]*User)
// 	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
// 	UserList["user_11111"] = &u
// }

// type User struct {
// 	Id       string
// 	Username string
// 	Password string
// 	Profile  Profile
// }

// type Profile struct {
// 	Gender  string
// 	Age     int
// 	Address string
// 	Email   string
// }

func CreateUser(u Users) *Users {
	o := orm.NewOrm()
	new_user := &Users{}
	if o.QueryTable("users").Filter("MobileNumber", u.MobileNumber).Exist() == false {
		new_user.FirstName = u.FirstName
		new_user.LastName = u.LastName
		new_user.Email = u.Email
		new_user.MobileNumber = u.MobileNumber
		new_user.Designation = u.Designation
		o.Insert(new_user)
	} else {
		fmt.Println("User already exists with the Mobile number!")
	}
	return new_user
}

func CreateSession(s Sessions) *Sessions {
	o := orm.NewOrm()
	new_session := &Sessions{}
	if o.QueryTable("sessions").Filter("User", s.User).Exist() == false {
		new_session.User = s.User
		new_session.DeviseToken = s.DeviseToken
	} else {
		fmt.Println("Session does not exist!")
	}
	return new_session
}

// func AddUser(u User) string {
// 	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
// 	UserList[u.Id] = &u
// 	return u.Id
// }

// func GetUser(uid string) (u *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		return u, nil
// 	}
// 	return nil, errors.New("User not exists")
// }

// func GetAllUsers() map[string]*User {
// 	return UserList
// }

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		if uu.Username != "" {
// 			u.Username = uu.Username
// 		}
// 		if uu.Password != "" {
// 			u.Password = uu.Password
// 		}
// 		if uu.Profile.Age != 0 {
// 			u.Profile.Age = uu.Profile.Age
// 		}
// 		if uu.Profile.Address != "" {
// 			u.Profile.Address = uu.Profile.Address
// 		}
// 		if uu.Profile.Gender != "" {
// 			u.Profile.Gender = uu.Profile.Gender
// 		}
// 		if uu.Profile.Email != "" {
// 			u.Profile.Email = uu.Profile.Email
// 		}
// 		return u, nil
// 	}
// 	return nil, errors.New("User Not Exist")
// }

// func Login(username, password string) bool {
// 	for _, u := range UserList {
// 		if u.Username == username && u.Password == password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func DeleteUser(uid string) {
// 	delete(UserList, uid)
// }
