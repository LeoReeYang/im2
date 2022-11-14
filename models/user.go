package models

import (
	"errors"
	"html"
	"strings"

	"github.com/LeoReeYang/im2/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;" json:"name"`
	Password string `gorm:"size:255;not null;" json:"password"`
}
type Friends struct {
	gorm.Model
	UserId   uint `json:"uid"`
	FriendId uint `json:"fid"`
}

type Blocks struct {
	gorm.Model
	UserId  uint `json:"uid"`
	BlockId uint `json:"blockid"`
}

func GetUserByID(uid uint) (User, error) {
	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}

	u.PrepareGive()

	return u, nil
}
func (u *User) PrepareGive() {
	u.Password = ""
}

func (u *User) SaveUser() (*User, error) {
	// user Model().xxx to use hooks
	err := DB.Model(&User{}).Create(&u).Error
	// err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	// log.Println("BeforeSave working!")

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	u := User{}

	err = DB.Model(User{}).Where("name = ?", username).First(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
