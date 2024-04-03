package helpers

import (
	"bytes"
	"context"
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

func RemoveExtension(filename string) string {
	return path.Base(filename[:len(filename)-len(path.Ext(filename))])
}
