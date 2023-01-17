package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ParseJwt(str string) {

}

func Create() {
	jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
}

const SECRET = "taoshihan"

type UserClaims struct {
	Id         uint      `json:"id"`
	Pid        uint      `json:"pid"`
	Username   string    `json:"username"`
	RoleId     uint      `json:"role_id"`
	CreateTime time.Time `json:"create_time"`
	jwt.RegisteredClaims
}

func MakeCliamsToken(obj UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, obj)
	tokenString, err := token.SignedString([]byte(SECRET))
	return tokenString, err
}

func TestJwt() {
	tokenCliams := UserClaims{
		Id:         1,
		Username:   "kefu2",
		RoleId:     2,
		Pid:        1,
		CreateTime: time.Now(),
		RegisteredClaims: jwt.RegisteredClaims{
			ID: "1",
		},
	}
	token, _ := MakeCliamsToken(tokenCliams)
	fmt.Println("#v", token)
	fmt.Printf("%v", ok)
}
