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
	sheetName = "movie_types"
}

type master struct {
	Exec        bool   `csv:"exec"`
	TypeID      uint32 `csv:"type_id"`
	TypeName    string `csv:"type_name"`
	Title       string `csv:"title"`
	Description string `csv:"description"`
}

func getMaster(ctx context.Context) ([]*master, error) {
	csv, err := ss.GetCSVFormSpreadSheet(ctx, spreadsheetID, sheetName)
	if err != nil {
		return nil, fmt.Errorf("GetCSVFormSpreadSheet Error: %w", err)
	}

	masters := []*master{}
	if err = gocsv.Unmarshal(bytes.NewBufferString(csv), &masters); err != nil {
		return nil, fmt.Errorf("Unmarshal CSV Error: %w", err)
	}

	return masters, nil
}

func handler(ctx context.Context) error {
	masters, err := getMaster(ctx)
	if err != nil {
		return err
	}

	for _, m := range masters {
		if !m.Exec {
			continue
		}

		data := model.MovieType{
			TypeID:      m.TypeID,
			TypeName:    m.TypeName,
			Title:       m.Title,
			Description: m.Description,
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
	if err := handler(context.TODO()); err != nil {
		panic(err)
	}

	fmt.Println("movie types seed load success")
	return
}
