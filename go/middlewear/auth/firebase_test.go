package auth_test

import (
	"context"
	"testing"
	"time"

	firebase_auth "firebase.google.com/go/auth"
	"github.com/miyabiii1210/ulala/go/middlewear/auth"
)

func TestNewFirebaseClient(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "NewFirebaseClient Test",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := auth.NewFirebaseClient()
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}

			t.Logf("firebase auth client: %v", app)
			t.Logf("%s End\n", tt.name)
		})
	}
}

func TestCreateFirebaseUser(t *testing.T) {
	type args struct {
		ctx          context.Context
		email        string
		emailVarfied bool
		phoneNumber  string
		password     string
		displayName  string
		photoURL     string
		disabled     bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CreateFirebaseUser Test",
			args: args{
				ctx:          context.TODO(),
				email:        "user1@example.com",
				emailVarfied: false,
				phoneNumber:  "+15555550200",
				password:     "Gorilla0#",
				displayName:  "Gorilla789",
				photoURL:     "http://www.example.com/12345678/photo.png",
				disabled:     false,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := auth.NewFirebaseClient()
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}

			params := (&firebase_auth.UserToCreate{}).
				Email(tt.args.email).
				EmailVerified(tt.args.emailVarfied).
				PhoneNumber(tt.args.phoneNumber).
				Password(tt.args.password).
				DisplayName(tt.args.displayName).
				Disabled(tt.args.disabled)

			u, err := app.CreateFirebaseUser(tt.args.ctx, params)
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}

			t.Logf("create user successfully: %v", u)
			t.Logf("%s End\n", tt.name)
		})
	}
}

func TestGetFirebaseUser(t *testing.T) {
	type args struct {
		ctx context.Context
		uid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetFirebaseUser Test",
			args: args{
				ctx: context.TODO(),
				uid: "hoge",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := auth.NewFirebaseClient()
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}
			u, err := app.GetFirebaseUser(tt.args.ctx, tt.args.uid)
			if err != nil {
				t.Error(err)
				return
			}

			t.Logf("UserInfo     : %s", u.UserInfo)      // &{Gorilla789 user1@example.com +15555550200  firebase U5UVhXunTYgkagFni7Z6HK0mvxh1}
			t.Logf("ProviderID   : %s", u.ProviderID)    // firebase
			t.Logf("UiD          : %s", u.UID)           // U5UVhXunTYgkagFni7Z6HK0mvxh1
			t.Logf("Email        : %s", u.Email)         // user1@example.com
			t.Logf("EmailVerified: %t", u.EmailVerified) // false
			t.Logf("PhoneNumber  : %s", u.PhoneNumber)   // PhoneNumber: +15555550200
			t.Logf("Disabled     : %t", u.Disabled)      // false
			t.Logf("TenantID     : %s", u.TenantID)      // empty
			t.Logf("CustomClaims : %s", u.CustomClaims)  // map[]
			t.Logf("%s End\n", tt.name)
		})
	}
}

func TestCreateFirebaseCustomToken(t *testing.T) {
	type args struct {
		ctx context.Context
		uid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CreateFirebaseCustomToken Test",
			args: args{
				ctx: context.TODO(),
				uid: "hoge",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := auth.NewFirebaseClient()
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}

			token, err := app.CreateFirebaseCustomToken(tt.args.ctx, tt.args.uid)
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}

			t.Logf("firebase custom token successfully: %s", token) // eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJmaXJlYmFzZS1hZG1pbnNkay1mOTI1c0BmaXItZGVtby03MDcyYy5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHl0b29sa2l0Lmdvb2dsZWFwaXMuY29tL2dvb2dsZS5pZGVudGl0eS5pZGVudGl0eXRvb2xraXQudjEuSWRlbnRpdHlUb29sa2l0IiwiZXhwIjoxNjg4MzA1ODk5LCJpYXQiOjE2ODgzMDIyOTksInN1YiI6ImZpcmViYXNlLWFkbWluc2RrLWY5MjVzQGZpci1kZW1vLTcwNzJjLmlhbS5nc2VydmljZWFjY291bnQuY29tIiwidWlkIjoiT1FzTHlLWGo3eFpUSXBoUFNNU0RKYXlKaGVGMiJ9.AWOktQ_aaa8tHulqU-9f7hD22ad4qSk3PGgfyGiVN0pPLXT5NnLetmHkqRZ2y_1U7VFM3pE8-qNdqddRvWSGSmRJmdYOJS0yQsX-2UoAgsLlWM_vz7PHAQGJNSobBfV2cyBlLcj4duMdEfYyW4QLqTjKfdoBoNfuBk3M9dpLab60lr-n1wxt2z8oTAxYKpCrJ2pZXriGY5dlQlvQ-bimZRFNzkJLuvc9u8ttmcueehsvsRZfHnH9J0H31yoj5kufHCd3msNs134xWDuN1XlnY8n-pYm2tXL4PxuEGkYi5ugORuH3HW5NvoO63Wgj7puRsSJilHQIku8CVMaGT_IZog
			t.Logf("%s End\n", tt.name)
		})
	}
}

func TestVerifyIDToken(t *testing.T) {
	type args struct {
		ctx   context.Context
		token string // It must be obtained by the client.
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "VerifyIDToken Test",
			args: args{
				ctx:   context.TODO(),
				token: "hoge",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := auth.NewFirebaseClient()
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}

			aToken, err := app.VerifyIDToken(tt.args.ctx, tt.args.token)
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}

			t.Logf("Audience      : %v", aToken.Audience)                // <firebase_project-id>
			t.Logf("AuthTime      : %v", aToken.AuthTime)                // 1688307303
			t.Logf("Claims        : %v", aToken.Claims)                  // map[auth_time:1.688307303e+09 email:hoge_test_001@gmail.com email_verified:true firebase:map[identities:map[email:[hoge_test_001@gmail.com]] sign_in_provider:password] user_id:mg6uSbN7LlaaE9K66j8Jes5L8EL2]
			t.Logf("Expires       : %v", aToken.Expires)                 // 1688310903
			t.Logf("Firebase      : %v", aToken.Firebase)                // {password  map[email:[hoge_test_001@gmail.com]]}
			t.Logf("IssuedAt      : %v", aToken.IssuedAt)                // 1688307303
			t.Logf("Issuer        : %v", aToken.Issuer)                  // https://securetoken.google.com/<firebase_project-id>
			t.Logf("Subject       : %v", aToken.Subject)                 // mg6uSbN7LlaaE9K66j8Jes5L8EL2
			t.Logf("UID           : %v", aToken.UID)                     // mg6uSbN7LlaaE9K66j8Jes5L8EL2
			t.Logf("Identities    : %v", aToken.Firebase.Identities)     // map[email:[hoge_test_001@gmail.com]]
			t.Logf("SignInProvider: %v", aToken.Firebase.SignInProvider) // password
			t.Logf("Tenant        : %v", aToken.Firebase.Tenant)         // none
			t.Logf("%s End\n", tt.name)
		})
	}
}

func TestSessionCookie(t *testing.T) {
	type args struct {
		ctx       context.Context
		token     string // It must be obtained by the client.
		expiresIn time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SessionCookie Test",
			args: args{
				ctx:       context.TODO(),
				token:     "eyJhbGciOiJSUzI1NiIsImtpZCI6ImY5N2U3ZWVlY2YwMWM4MDhiZjRhYjkzOTczNDBiZmIyOTgyZTg0NzUiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vZmlyLWRlbW8tNzA3MmMiLCJhdWQiOiJmaXItZGVtby03MDcyYyIsImF1dGhfdGltZSI6MTY4ODU0NzY4OCwidXNlcl9pZCI6IkZ5RU9ReEtLMm1TajNkVXFTOVdBTEZaR1AwcjEiLCJzdWIiOiJGeUVPUXhLSzJtU2ozZFVxUzlXQUxGWkdQMHIxIiwiaWF0IjoxNjg4NTQ3Njg4LCJleHAiOjE2ODg1NTEyODgsImVtYWlsIjoibWFzYXRvLm1laWppMEBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsibWFzYXRvLm1laWppMEBnbWFpbC5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9fQ.x1KR9URtdmvmjH7Utxx2eF3ngmCuHdC-AI1nxJoVU9llMs0RYyawEdZQAqJxaPJBwH2I56ghLnCU_F340k6RZQu9JkCILwPhtS642zv7JLiKQ1-xmmPoguea5VH4dC0DmkRnAmenCiye9ju_tnqxZv4l_8jpJzIuBzbJ4uSEml83PVBDkWoCB7nZlYfECTQQCsaXOlwXwdiulgTurlFYlLNgh9FHtRMLh3EwoZfYy9uxxnsuSKb4o4BN3DUUxgnx-Vu4YCLub2jsnJX1tYWudai7w-Ppx20r0v4aOPcc-8kOofvcfSKc0veQS-hcl8WIpM98Mcb3sbDOXNdpdq-Oqw",
				expiresIn: time.Hour * 24 * 7,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := auth.NewFirebaseClient()
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}

			cookie, err := app.SessionCookie(tt.args.ctx, tt.args.token, tt.args.expiresIn)
			if err != nil {
				t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
				return
			}
			t.Logf("firebase session cookie: %s", cookie)

			// validToken, err := app.VerifySessionCookie(tt.args.ctx, tt.args.token)
			// if err != nil {
			// 	t.Errorf("unexpected error: %v, wantErr: %v", err, tt.wantErr)
			// 	return
			// }
			// t.Logf("firebase session cookie: %v", validToken)

			t.Logf("%s End\n", tt.name)
		})
	}
}
