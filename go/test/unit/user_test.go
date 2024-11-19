package unit_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/miyabiii1210/ulala/go/library/api"
	"github.com/miyabiii1210/ulala/go/model"
)

func CreateUser(ctx context.Context, path string, internal *model.User) (*model.CreateUserResponse, error) {
	req := api.NewRequest(http.MethodPost, path, internal)
	res, err := api.SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(model.CreateUserResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func TestCreateUser(t *testing.T) {
	type args struct {
		ctx      context.Context
		path     string
		internal model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CreateUser Test",
			args: args{
				ctx:  context.TODO(),
				path: "/signup",
				internal: model.User{
					Name:  "hoge1",
					Email: "hoge1@example.com",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CreateUser(tt.args.ctx, tt.args.path, &tt.args.internal)
			if err != nil {
				t.Errorf("[%s] err: %v\n", tt.name, err.Error())
				return
			}
			t.Log("result:", res)
			t.Logf("%s End\n", tt.name)
		})
	}
}
