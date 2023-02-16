package auth

import (
	"errors"
	"redisadmin/internal/configs"
	"redisadmin/internal/consts"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 密钥
var jwtKey []byte

type JwtClaims struct {
	UId  uint
	RId  uint
	Name string
	jwt.StandardClaims
}

func init() {
	salt := configs.GetEnvVal(consts.ENV_CONF_JWT_KEY)
	jwtKey = []byte(salt)
}

func BuildJwtToken(uId uint, rId uint, name string) (string, error) {
	expireTime := time.Now().Add(2 * time.Hour)
	claims := &JwtClaims{
		UId:  uId,
		RId:  rId,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			// Issuer:    "127.0.0.1",  // 签名颁发者
			// Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//解析token
func CheckJwtToken(tokenString string) (*JwtClaims, error) {
	//vcalidate token formate
	if tokenString == "" {
		return nil, errors.New("token为空")
	}

	token, claims, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		return nil, errors.New("token解析失败或已过期")
	}

	return claims, nil
}

func parseToken(tokenString string) (*jwt.Token, *JwtClaims, error) {
	Claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, Claims, err
}

//TestBuildJwtToken 测试 BuildJwtToken 函数
func TestBuildJwtToken(t *testing.T) {
	//声明用户信息
	uId := uint(123)
	rId := uint(456)
	name := "test user"
	//调用 BuildJwtToken 函数，获取 token
	tokenString, err := BuildJwtToken(uId, rId, name)
	if err != nil {
		t.Error("generate token failed, err:", err)
	}

	//检查 token 是否正确
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		t.Error("parse token failed, err:", err)
	}

	//检查用户信息是否正确
	if claims.Name != name || claims.UId != uId || claims.RId != rId {
		t.Error("invalid token!")
	}
}
