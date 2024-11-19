package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

const (
	MOVIE_TYPE_RANGE = 8
	BASE_PATH        = "../../public/proto/img/original"
	THUMBNAIL_PATH   = "../../public/proto/img/thumbnail"
	BASE_FILE_NAME   = "0000000.png"
	LIMIT_EDGE       = 256
)

type imageData struct {
	typeID       int
	path         string
	decodeImage  image.Image
	imgRectangle image.Rectangle
}

func getImagePathList() ([]imageData, error) {
	localImageDataList := []imageData{}
	for i := 1; i <= MOVIE_TYPE_RANGE; i++ {
		localImageDataList = append(localImageDataList, imageData{
			typeID: i,
		})
	}

	for i, localImageData := range localImageDataList {
		path := fmt.Sprintf("%s/type%d/%d%s",
			BASE_PATH, localImageData.typeID, localImageData.typeID, BASE_FILE_NAME,
		)

		matches, err := filepath.Glob(path)
		if err != nil {
			return nil, err
		}

		if len(matches) != 1 {
			return nil, fmt.Errorf("%d: invalid file: %s", i+1, path)
		}

		localImageData.path = matches[0]
		localImageDataList[i] = localImageData
	}

	return localImageDataList, nil
}

func getImageSize(imageDataList []imageData) ([]imageData, error) {
	newImageDataList := []imageData{}
	for _, img := range imageDataList {
		fileData, err := os.Open(img.path)
		if err != nil {
			return nil, err
		}
		defer fileData.Close()

		decodeImage, _, err := image.Decode(fileData)
		if err != nil {
			return nil, err
		}

		tmpImgRectangle := decodeImage.Bounds()
		newImageDataList = append(newImageDataList, imageData{
			typeID:      img.typeID,
			path:        img.path,
			decodeImage: decodeImage,
			imgRectangle: image.Rectangle{
				Min: image.Point{0, 0},
				Max: image.Point{tmpImgRectangle.Dx(), tmpImgRectangle.Dy()},
			},
		})
	}

	return newImageDataList, nil
}

func resizeImage(imageDataList []imageData) error {
	for _, img := range imageDataList {
		newImgData := &image.RGBA{}
		width := img.imgRectangle.Bounds().Size().X
		height := img.imgRectangle.Bounds().Size().Y
		if height >= width {
			// 縦長の画像
			f := float64(LIMIT_EDGE * height)
			h := math.Round((f / float64(width)))
			newImgData = image.NewRGBA(image.Rect(0, 0, LIMIT_EDGE, int(h)))
		} else {
			// 横長の画像
			f := float64(width * LIMIT_EDGE)
			w := math.Round((f / float64(height)))
			newImgData = image.NewRGBA(image.Rect(0, 0, int(w), LIMIT_EDGE))
		}

		draw.CatmullRom.Scale(
			newImgData, newImgData.Bounds(), img.decodeImage, img.imgRectangle, draw.Over, nil,
		)
		newImage, err := os.Create(fmt.Sprintf("%s/type%d/%d%s",
			THUMBNAIL_PATH, img.typeID, img.typeID, BASE_FILE_NAME,
		))
		if err != nil {
			return err
		}
		defer newImage.Close()

		if err := png.Encode(newImage, newImgData); err != nil {
			return err
		}
	}

	return nil
}

func handler() error {
	imageDataList, err := getImagePathList()
	if err != nil {
		return err
	}

	newImageDataList, err := getImageSize(imageDataList)
	if err != nil {
		return err
	}

	// for i, img := range newImageDataList {
	// 	fmt.Printf("[index:%d] typeID: %d, path: %s width: %d, height: %d\n", i+1, img.typeID, img.path, img.imgRectangle.Size().X, img.imgRectangle.Size().Y)
	// }

	if err := resizeImage(newImageDataList); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := handler(); err != nil {
		panic(err)
	}
}
