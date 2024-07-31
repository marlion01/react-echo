package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Signup(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}
type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}
func (uc *userController) Signup(c echo.Context) error {
	user := model.User{}                  //model.Userはgo-rest-api/model/user.goで定義されている
	if err := c.Bind(&user); nil != err { //リクエストのbodyをuserにバインド
		return c.JSON(http.StatusBadRequest, err.Error()) //エラーがあれば400を返す
	}
	resUser, err := uc.uu.SignUp(user) //usecaseのSignUpメソッドを呼び出す
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) //エラーがあれば500を返す
	}
	return c.JSON(http.StatusCreated, resUser)
}
func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); nil != err {
		return c.JSON(http.StatusBadRequest, err.Error()) //エラーがあれば400を返す
	}
	token, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) //エラーがあれば500を返す
	}
	cookie := new(http.Cookie)                      //cookieの生成
	cookie.Name = "token"                           //cookieの名前
	cookie.Value = token                            //JWTのtokenをcookieに格納
	cookie.Expires = time.Now().Add(24 * time.Hour) //cookieの有効期限
	cookie.Path = "/"                               //cookieの有効なパス
	cookie.Domain = os.Getenv("API_DOMAIN")         //cookieno有効なドメイン
	cookie.HttpOnly = true                          //cookieのhttpOnly属性
	cookie.SameSite = http.SameSiteNoneMode         //cookieのSameSite属性
	c.SetCookie(cookie)                             //cookieをセット
	return c.NoContent(http.StatusOK)
}
func (uc *userController) Logout(c echo.Context) error {
	cookie := new(http.Cookie)              //cookieの生成
	cookie.Name = "token"                   //cookieの名前
	cookie.Value = ""                       //cookieの値を空にする
	cookie.Expires = time.Now()             //cookieの有効期限を過去に設定して削除
	cookie.Path = "/"                       //cookieの有効なパス
	cookie.Domain = os.Getenv("API_DOMAIN") //cookieno有効なドメイン
	cookie.HttpOnly = true                  //cookieのhttpOnly属性
	cookie.SameSite = http.SameSiteNoneMode //cookieのSameSite属性
	c.SetCookie(cookie)                     //cookieをセット
	return c.NoContent(http.StatusOK)
}
