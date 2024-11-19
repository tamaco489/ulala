package unit_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/miyabiii1210/ulala/go/library/api"
	"github.com/miyabiii1210/ulala/go/model"
)

func getMovie(ctx context.Context, path string) (*model.GetMovieResponse, error) {
	req := api.NewRequest(http.MethodGet, path, nil)
	res, err := api.SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(model.GetMovieResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func TestGetMovie(t *testing.T) {
	type args struct {
		ctx     context.Context
		service string
		movieID uint32
	}
	test := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetMovie Test",
			args: args{
				ctx:     context.TODO(),
				service: "/movies/",
				movieID: 10000001,
			},
			wantErr: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.args.service + strconv.FormatUint(uint64(tt.args.movieID), 10) // /movies/10000001
			res, err := getMovie(tt.args.ctx, path)
			if err != nil {
				t.Errorf("[%s] err: %v\n", tt.name, err.Error())
				return
			}
			t.Log("result:", res)
		})
	}
}
