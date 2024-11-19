package auth

import (
	"context"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/miyabiii1210/ulala/go/config"
	"google.golang.org/api/option"
)

type IFirebaseClient interface {
	CreateFirebaseUser(ctx context.Context) (*auth.UserRecord, error)
	GetFirebaseUser(ctx context.Context, uid string) (*auth.UserRecord, error)
	CreateFirebaseCustomToken(ctx context.Context, uid string) (string, error)
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
	SessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error)
	VerifySessionCookie(ctx context.Context, idToken string) (*auth.Token, error)
	RevokeRefreshTokens(ctx context.Context, uid string) error
}

type FirebaseClient struct {
	app  *firebase.App
	auth *auth.Client
}

func NewFirebaseClient() (*FirebaseClient, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.EnvConfig.GoogleApplicationCredentials)
	if !(config.IsLocal() || config.IsDevelopment()) {
		opt = option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	}

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseClient{
		app:  app,
		auth: authClient,
	}, nil
}

func (app *FirebaseClient) CreateFirebaseUser(ctx context.Context, params *auth.UserToCreate) (*auth.UserRecord, error) {
	return app.auth.CreateUser(ctx, params)
}

func (app *FirebaseClient) GetFirebaseUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	return app.auth.GetUser(ctx, uid)
}

func (app *FirebaseClient) CreateFirebaseCustomToken(ctx context.Context, uid string) (string, error) {
	return app.auth.CustomToken(ctx, uid)
}

func (app *FirebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return app.auth.VerifyIDToken(ctx, idToken)
}

func (app *FirebaseClient) SessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error) {
	return app.auth.SessionCookie(ctx, idToken, expiresIn)
}

func (app *FirebaseClient) VerifySessionCookie(ctx context.Context, idToken string) (*auth.Token, error) {
	return app.auth.VerifySessionCookie(ctx, idToken)
}

func (app *FirebaseClient) RevokeRefreshTokens(ctx context.Context, uid string) error {
	return app.auth.RevokeRefreshTokens(ctx, uid)
}
