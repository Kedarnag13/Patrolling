package models

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io"
)

type Users struct {
	Id                   int
	FirstName            string `json:"firstname" valid:"Required;Alpha"`
	LastName             string `json:"lastname" valid:"Required;Alpha"`
	Email                string `json:"email" valid:"Required;Email" orm:"unique"`
	MobileNumber         string `json:"mobile_number" valid:"Required" orm:"unique"`
	Designation          string `json:"designation" valid:"Alpha"`
	Password             string `json:"password" valid:"Required"; MinSize(7); MaxSize(15)`
	PasswordConfirmation string `json:"password_confirmation" valid:"Required"; MinSize(7); MaxSize(15)`
}

type Sessions struct {
	Id          int
	User        *Users   `json:"user" orm:"rel(fk)"`
	DeviseToken *Devises `json:"devise_token" orm:"rel(fk)"`
}

type Devises struct {
	Id    int
	Token string `json:"token" orm:"pk"`
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
	o.Using("default")
	new_user := &Users{}
	new_user.FirstName = u.FirstName
	new_user.LastName = u.LastName
	new_user.Email = u.Email
	new_user.MobileNumber = u.MobileNumber
	new_user.Designation = u.Designation
	key := []byte("traveling is fun")
	Password := []byte(u.Password)
	PasswordConfirmation := []byte(u.PasswordConfirmation)
	new_user.Password = Encrypt(key, Password)
	new_user.PasswordConfirmation = Encrypt(key, PasswordConfirmation)
	if o.QueryTable("users").Filter("MobileNumber", u.MobileNumber).Exist() == false {
		o.Insert(new_user)
	} else {
		fmt.Println("User already exists with the Mobile number!")
	}
	return new_user
}

func CreateSession(s Sessions) *Sessions {
	o := orm.NewOrm()
	o.Using("default")
	new_session := &Sessions{}
	if o.QueryTable("sessions").Filter("devise_token_id", s.DeviseToken).Exist() == false {
		key := []byte("traveling is fun")

		user_id := o.Raw("SELECT id FROM users WHERE mobile_number = ?", s.User.MobileNumber)
		db_password := o.Raw("SELECT password FROM users WHERE mobile_number = ?", s.User.MobileNumber)
		decrypt_password := Decrypt(key, db_password)
		new_session.User.MobileNumber = s.User.MobileNumber
		new_session.User.Password = decrypt_password
		if new_session.User.MobileNumber == s.User.MobileNumber && new_session.User.Password == decrypt_password {
			new_session.User = s.User
			new_session.User = user_id
			new_session.DeviseToken = s.DeviseToken
			o.Insert(new_session)
		} else {
			fmt.Println("Invalid Mobile Number or Password")
		}
		// var users []*Users
		// o.QueryTable("users")
		// fmt.Println(o.Read(&users))
	} else {
		fmt.Println("Session already Exists!")
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

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(key, text []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(crand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return encodeBase64(ciphertext)
}

func Decrypt(key []byte, b64 string) string {
	text := decodeBase64(b64)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(text)
}
