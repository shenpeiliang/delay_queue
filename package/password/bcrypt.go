package password

//noinspection ALL
import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	Cost int
}

func New() *Bcrypt {
	return &Bcrypt{}
}

//加密
func (bc Bcrypt) GeneratePassword(userPassword string) ([]byte, error) {
	if bc.Cost == 0 {
		bc.Cost = bcrypt.DefaultCost
	}
	return bcrypt.GenerateFromPassword([]byte(userPassword), bc.Cost)
}

//密码比对
func (bc Bcrypt) ValidatePassword(hashed, userPassword string) (ok bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码错误！")
	}
	return true, nil

}
