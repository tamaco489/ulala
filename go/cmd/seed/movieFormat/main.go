package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/miyabiii1210/ulala/go/datastore"
	"github.com/miyabiii1210/ulala/go/library/ss"
	"github.com/miyabiii1210/ulala/go/model"
)

var spreadsheetID, sheetName string

func init() {
	if spreadsheetID = os.Getenv("GOOGLE_SPREADSHEET_ID"); spreadsheetID == "" {
		panic("GOOGLE_SPREADSHEET_ID is not set")
	}
	sheetName = "movie_formats"
}

type master struct {
	Exec        bool   `csv:"exec"`
	FormatID    uint32 `csv:"format_id"`
	MovieFormat string `csv:"movie_format"`
}

func getMaster(ctx context.Context) ([]*master, error) {
	csv, err := ss.GetCSVFormSpreadSheet(ctx, spreadsheetID, sheetName)
	if err != nil {
		return nil, fmt.Errorf("GetCSVFormSpreadSheet Error: %w", err)
	}

	master := []*master{}
	if err = gocsv.Unmarshal(bytes.NewBufferString(csv), &master); err != nil {
		return nil, fmt.Errorf("Unmarshal CSV Error: %w", err)
	}

	return master, nil
}

func handler(ctx context.Context) error {
	master, err := getMaster(ctx)
	if err != nil {
		return err
	}

	for _, m := range master {
		if !m.Exec {
			continue
		}

		data := model.MovieFormat{
			FormatID:    m.FormatID,
			MovieFormat: m.MovieFormat,
		}

		db := datastore.NewDBConnection()
		if db.Error != nil {
			return fmt.Errorf("NewDBConnection Error: %w", db.Error)
		}

		if err := db.Create(&data).Error; err != nil {
			return fmt.Errorf("Create Error: %w", err)
		}
	}

	return nil
}

func main() {
	if err := handler(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println("movie formats seed load success")
}
