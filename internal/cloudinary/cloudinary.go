package cloudinary

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"

	"redoocehub/domains/entities"
)

func ImageUploadHelper(input interface{}, config entities.CloudinaryEnvSetting) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.CloudName, config.ApiKey, config.ApiSecret)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.UploadFolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
