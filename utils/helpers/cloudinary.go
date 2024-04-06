package helpers

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func UploadFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cloudName := viper.GetString("CLOUDINARY_CLOUD_NAME")
	cloudApiKey := viper.GetString("CLOUDINARY_API_KEY")
	cloudApiSecret := viper.GetString("CLOUDINARY_API_SECRET")
	cloudFolderLocation := viper.GetString("CLOUDINARY_UPLOAD_FOLDER")

	cloud, err := cloudinary.NewFromParams(cloudName, cloudApiKey, cloudApiSecret)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	fileReader, err := convertFile(fileHeader)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	uploadParam, err := cloud.Upload.Upload(ctx, fileReader, uploader.UploadParams{
		PublicID: fileName,
		Folder:   cloudFolderLocation,
	})

	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	return uploadParam.SecureURL, nil

}

func convertFile(fileHeader *multipart.FileHeader) (*bytes.Reader, error) {

	file, err := fileHeader.Open()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}

	fileReader := bytes.NewReader(buffer.Bytes())
	return fileReader, nil

}

func GenerateFileNameImage() (string, error) {

	timeNow := time.Now()
	timeFormat := timeNow.Format("02012006")
	byteCode := make([]byte, 5)

	_, err := rand.Read(byteCode)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	byteString := hex.EncodeToString(byteCode)

	codeImage := fmt.Sprintf("PRD%s%s", timeFormat, byteString)

	return codeImage, nil
}
