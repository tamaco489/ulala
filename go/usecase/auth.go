package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/miyabiii1210/ulala/go/middlewear/auth"
	"github.com/miyabiii1210/ulala/go/model"
	"github.com/miyabiii1210/ulala/go/repository"
	"github.com/miyabiii1210/ulala/go/validator"
)

type IAuthUsecase interface {
	SignUp(c echo.Context, req model.SignUpRequest) (model.SignUpResponse, error)
	SignIn(c echo.Context, req model.SignInRequest) (model.SignInResponse, error)
	SignOut(c echo.Context, req model.SignOutRequest) error
}

type authUsecase struct {
	ar repository.IAuthRepository
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewAuthUsecase(ar repository.IAuthRepository, ur repository.IUserRepository, uv validator.IUserValidator) IAuthUsecase {
	return &authUsecase{ar, ur, uv}
}

func makeCookie(value string, expires time.Duration) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "_session", // TODO: 環境変数化したい
		Value:    value,
		MaxAge:   int(expires.Seconds()),
		Domain:   "localhost", // TODO: 環境変数化したい
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Secure:   true,
	}

	return cookie
}

func (au *authUsecase) SignIn(c echo.Context, req model.SignInRequest) (model.SignInResponse, error) {
	requestEmail, requestToken := req.Email, req.FirebaseToken
	if requestEmail == "" || requestToken == "" {
		return model.SignInResponse{}, fmt.Errorf("invalid request. email or token is nil")
	}

	app, err := auth.NewFirebaseClient()
	if err != nil {
		return model.SignInResponse{}, err
	}

	ctx := c.Request().Context()
	decoded, err := app.VerifyIDToken(ctx, requestToken)
	if err != nil {
		return model.SignInResponse{}, err
	}

	firebaseUID := decoded.UID
	if firebaseUID == "" {
		return model.SignInResponse{}, fmt.Errorf("invalid firebase user_id")
	}

	fu, err := app.GetFirebaseUser(ctx, firebaseUID)
	if err != nil {
		return model.SignInResponse{}, err
	}

	if fu.Email != requestEmail {
		return model.SignInResponse{}, fmt.Errorf("Invalid email address: %v", fu.Email)
	}

	u := model.User{}
	if err := au.ar.SignIn(firebaseUID, &u); err != nil {
		return model.SignInResponse{}, err
	}

	uid := u.UID
	if uid == 0 {
		return model.SignInResponse{}, fmt.Errorf("uid not exist")
	}

	expiresIn := time.Hour * 24 * 5
	cookie, err := app.SessionCookie(ctx, requestToken, expiresIn)
	if err != nil {
		return model.SignInResponse{}, err
	}

	c.SetCookie(makeCookie(cookie, expiresIn))

	return model.SignInResponse{
		UID: uid,
	}, nil
}

func (au *authUsecase) SignUp(c echo.Context, req model.SignUpRequest) (model.SignUpResponse, error) {
	requestEmail, requestToken := req.Email, req.FirebaseToken
	if requestEmail == "" || requestToken == "" {
		return model.SignUpResponse{}, fmt.Errorf("invalid request. email or token is nil")
	}

	app, err := auth.NewFirebaseClient()
	if err != nil {
		return model.SignUpResponse{}, err
	}

	// リクエストで受け取ったtokenの正常性確認
	ctx := c.Request().Context()
	decoded, err := app.VerifyIDToken(ctx, requestToken)
	if err != nil {
		return model.SignUpResponse{}, err
	}

	firebaseUID := decoded.UID
	if firebaseUID == "" {
		return model.SignUpResponse{}, fmt.Errorf("invalid firebase user_id ")
	}

	// 検証したfirebase_uidを使用して、firebase側で保持するユーザ情報を取得し、signup時に使用したemailアドレスと一致していることを確認
	fu, err := app.GetFirebaseUser(ctx, firebaseUID)
	if err != nil {
		return model.SignUpResponse{}, err
	}

	if fu.Email != requestEmail {
		return model.SignUpResponse{}, fmt.Errorf("Invalid email address: %v", fu.Email)
	}

	uuid := uuid.New().String()
	a := model.UserFirebaseAuthentication{
		FirebaseUID: firebaseUID,
		UUID:        uuid,
	}

	if err := au.ar.SignUp(&a); err != nil {
		return model.SignUpResponse{}, err
	}

	u := model.User{
		Name:     "default",
		Email:    requestEmail,
		AuthUUID: uuid,
	}

	if err := au.uv.CreateUserValidate(u); err != nil {
		return model.SignUpResponse{}, err
	}

	if err := au.ur.CreateUser(&u); err != nil {
		return model.SignUpResponse{}, err
	}

	// firebaseのsessionを生成
	expiresIn := time.Hour * 24 * 5
	cookie, err := app.SessionCookie(ctx, requestToken, expiresIn)
	if err != nil {
		return model.SignUpResponse{}, err
	}

	// 生成したsessionをHTTPレスポンスのSet-Cookieヘッダに追加
	c.SetCookie(makeCookie(cookie, expiresIn))

	return model.SignUpResponse{
		FirebaseUID: firebaseUID,
	}, nil
}

func (au *authUsecase) SignOut(c echo.Context, req model.SignOutRequest) error {
	requestToken := req.FirebaseToken
	if requestToken == "" {
		return fmt.Errorf("invalid request. token is nil")
	}

	app, err := auth.NewFirebaseClient()
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	decoded, err := app.VerifyIDToken(ctx, requestToken)
	if err != nil {
		return err
	}

	firebaseUID := decoded.UID
	if firebaseUID == "" {
		return fmt.Errorf("invalid firebase user_id ")
	}

	// firebaseのセッションを終了
	if err := app.RevokeRefreshTokens(ctx, decoded.UID); err != nil {
		return err
	}

	c.SetCookie(makeCookie("", 0))

	return nil
}
