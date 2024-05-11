package usecases

import (
	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/internal/cloudinary"
)

type media struct {}

func NewMediaUpload() entities.CloudinaryRepository {
    return &media{}
}

func (*media) FileUpload(file dto.File, config entities.CloudinaryEnvSetting) (string, error) {
	uploadUrl, err := cloudinary.ImageUploadHelper(file.File, config)

	if err != nil {
		return "", err
	}

	return uploadUrl, nil
}

func (*media) RemoteUpload(url dto.Url, config entities.CloudinaryEnvSetting) (string, error) {
	uploadUrl, errUrl := cloudinary.ImageUploadHelper(url.Url, config)
	if errUrl != nil {
		return "", errUrl
	}
	return uploadUrl, nil
}