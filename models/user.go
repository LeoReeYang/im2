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
	Name     string   `gorm:"size:255;not null;" json:"name"`
	Password string   `gorm:"size:255;not null;" json:"password"`
	Friends  []Friend `gorm:"foreignKey:UserID" json:"friends"`
	Blocks   []Block  `gorm:"foreignKey:UserID" json:"blocks"`
}
type Friend struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	FriendID uint `json:"friend_id"`
}
type Block struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	BlockID uint `json:"block_id"`
}

func Get(uid uint) (*User, error) {
	u := User{}

	if err := DB.First(&u, uid).Error; err != nil {
		return &u, errors.New("user not found")
	}

	u.PrepareGive()

	return &u, nil
}

func All() (users []User, err error) {
	if err = DB.Find(&users).Error; err != nil {
		return
	}
	return
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func (u *User) Save() (*User, error) {
	// user Model().xxx to use hooks this is correct way.
	// err := DB.Create(&u).Error   this is wrong to add user when using hooks.
	err := DB.Model(&User{}).Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
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

func UserVerify(username string, password string) (string, error) {
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

	token, err := token.GenerateToken(u.ID, username)

	if err != nil {
		return "", err
	}

	return token, nil
}
