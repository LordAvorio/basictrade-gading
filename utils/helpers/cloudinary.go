package helpers

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/rs/zerolog/log"
)

func UploadFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudApiKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudApiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudFolderLocation := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")

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

func DeleteFile(pathFile string) error {

	pathSplit := strings.Split(pathFile, "/")
	fileName := pathSplit[len(pathSplit)-1]

	fileNameSplit := strings.Split(fileName, ".")
	fileNameWithoutExt := fileNameSplit[0]

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudApiKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudApiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudFolderLocation := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")

	publicId := fmt.Sprintf("%s/%s", cloudFolderLocation, fileNameWithoutExt)

	cloud, err := cloudinary.NewFromParams(cloudName, cloudApiKey, cloudApiSecret)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	response, errDelete := cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicId,
	})

	if errDelete != nil {
		log.Error().Msg(errDelete.Error())
		return errDelete
	}

	if response.Error.Message != "" {
		log.Error().Msg(response.Error.Message)
		return errors.New(response.Error.Message)
	}

	if response.Result == "not found" {
		log.Error().Msg("Cannot found data image on cloudinary")
		return errors.New("cannot found data image on cloudinary")
	}

	return nil

}

