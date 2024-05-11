package repositories

import (
	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/internal/cloudinary"
)


type media struct {}

type CloudinaryRepository struct {
	Media *media
}

func NewCloudinaryRepository(media *media) entities.CloudinaryRepository {
	return &CloudinaryRepository{Media: media}
}


func (cr *CloudinaryRepository) FileUpload(file dto.File, config entities.CloudinaryEnvSetting) (string, error) {
	uploadUrl, err := cloudinary.ImageUploadHelper(file.File, config)

	if err != nil {
		return "", err
	}

	return uploadUrl, nil
}

func (cr *CloudinaryRepository) RemoteUpload(url dto.Url, config entities.CloudinaryEnvSetting) (string, error) {
	uploadUrl, errUrl := cloudinary.ImageUploadHelper(url.Url, config)
	if errUrl != nil {
		return "", errUrl
	}
	return uploadUrl, nil
}
