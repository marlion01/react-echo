package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"

	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface { //interface型は抽象クラスのようなもの
	SignUp(user model.User) (model.UserResponse, error) //UserResponseは返す
	Login(user model.User) (string, error)              //striingはJWTのtokenを返す
}
type useUserUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &useUserUsecase{ur}
}
func (uu *useUserUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err = uu.ur.CreateUser(&newUser); nil != err {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{ID: newUser.ID, Email: newUser.Email}
	return resUser, nil

}
func (uu *useUserUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}                                            //Emailで検索して取得したユーザー情報を格納するための空のUser構造体
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); nil != err { //Emailで検索して取得したユーザー情報を格納する
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); nil != err { //パスワードの比較
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //JWTの生成
		"passward": storedUser.ID,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret")) //JWTの文字列化
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
