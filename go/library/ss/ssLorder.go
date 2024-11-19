package ss

import (
	"context"
	"errors"
	"strings"

	"github.com/miyabiii1210/ulala/go/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var credentials string

func init() {
	if credentials = config.EnvConfig.GoogleApplicationCredentials; credentials == "" {
		panic("GoogleApplicationCredentials is not set")
	}
}

func Escape(s string) string {
	x := strings.Replace(s, ",", "_+_", -1)
	x = strings.Replace(x, "\n", "\\n", -1)

	return x
}

func GetCSVFormSpreadSheet(ctx context.Context, spreadsheetID, sheet string) (csv string, err error) {
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentials))
	if err != nil {
		return "", err
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, sheet).Do()
	if err != nil {
		return "", err
	}

	if len(resp.Values) == 0 {
		return "", errors.New("No data found")
	}

	rows := make([]string, len(resp.Values))
	for i, row := range resp.Values {
		colomns := make([]string, len(row))
		for j, c := range row {
			s := Escape(c.(string))
			colomns[j] = s
		}
		rows[i] = strings.Join(colomns, ",")
	}

	return strings.Join(rows, "\n"), err
}
