package entities

import "redoocehub/domains/dto"

type CloudinaryEnvSetting struct {
	CloudName    string
	ApiKey       string
	ApiSecret    string
	UploadFolder string
}

type CloudinaryRepository interface {
	FileUpload(file dto.File, config CloudinaryEnvSetting) (string, error)
    RemoteUpload(url dto.Url, config CloudinaryEnvSetting) (string, error)
}

type CloudinaryUsecase interface {
	NewMediaUpload() CloudinaryRepository
	FileUpload(file dto.File, config CloudinaryEnvSetting) (string, error)
    RemoteUpload(url dto.Url, config CloudinaryEnvSetting) (string, error)
}