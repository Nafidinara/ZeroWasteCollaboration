package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"redoocehub/bootstrap"
	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/domains/infra"
	"redoocehub/internal/constant"
	"redoocehub/internal/validation"
	"redoocehub/usecases"
)

type UserController struct {
	UserUsecase    entities.UserUsecase
	ChatbotUsecase entities.ChatbotUsecase
	Env            *bootstrap.Env
}

func (uc *UserController) Register(c echo.Context) error {
	var request dto.RegisterRequest

	formHeader, errFile := c.FormFile("profile_image")

	if errFile != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrFailedGetFile, errFile.Error())
	}

	formFile, errFile := formHeader.Open()

	if errFile != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrFailedOpenFile, errFile.Error())
	}

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	if _, err := uc.UserUsecase.GetUserByEmail(c.Request().Context(), request.Email); err == nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.Conflict, constant.ErrEmailAlreadyExist, err)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGeneratePassword, err.Error())
	}

	uploadUrl, err := usecases.NewMediaUpload().FileUpload(dto.File{File: formFile}, entities.CloudinaryEnvSetting{
		CloudName:    uc.Env.CLOUDINARY_CLOUD_NAME,
		ApiKey:       uc.Env.CLOUDINARY_API_KEY,
		ApiSecret:    uc.Env.CLOUDINARY_API_SECRET,
		UploadFolder: uc.Env.CLOUDINARY_UPLOAD_FOLDER,
	})

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrFailedUploadFile, err.Error())
	}

	request.Password = string(encryptedPassword)
	request.ProfileImage = uploadUrl

	user, err := uc.UserUsecase.Create(c.Request().Context(), &request)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrRegisterUser, err.Error())
	}

	accessToken, err := uc.UserUsecase.CreateAccessToken(user, uc.Env.ACCESS_TOKEN_SECRET, uc.Env.ACCESS_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrAccessToken, err.Error())
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(user, uc.Env.REFRESH_TOKEN_SECRET, uc.Env.REFRESH_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrRefreshToken, err.Error())
	}

	user.RefreshToken = refreshToken

	response := entities.ToRegisterResponseUser(user, accessToken)

	return infra.NewSuccessResponse(c, http.StatusCreated, constant.SuccessCreated, constant.SuccessRegisterUser, response)
}

func (uc *UserController) Login(c echo.Context) error {
	var request dto.LoginRequest

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	user, err := uc.UserUsecase.GetUserByEmail(c.Request().Context(), request.Email)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusUnauthorized, constant.ErrUnauthorized, constant.ErrNotFoundUser, err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return infra.NewErrorResponse(c, http.StatusUnauthorized, constant.ErrUnauthorized, constant.ErrInvalidCredentials, err)
	}

	accessToken, err := uc.UserUsecase.CreateAccessToken(&user, uc.Env.ACCESS_TOKEN_SECRET, uc.Env.ACCESS_TOKEN_EXPIRY_HOUR)
	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrAccessToken, err.Error())
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(&user, uc.Env.REFRESH_TOKEN_SECRET, uc.Env.REFRESH_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrRefreshToken, err.Error())
	}

	loginResponse := entities.ToLoginResponseUser(entities.EntityToDtoUser(&user), accessToken, refreshToken)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessLoginUser, loginResponse)
}

func (uc *UserController) Profile(c echo.Context) error {
	profile, err := uc.UserUsecase.GetProfileByID(c.Request().Context(), c.Get("x-user-id").(string))

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGetProfile, err.Error())
	}

	response := entities.ToProfileResponseUser(profile)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Profile retrieved successfully",
		Data:       response,
	})
}

func (uc *UserController) RefreshToken(c echo.Context) error {
	var request dto.RefreshTokenRequest

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	user, err := uc.UserUsecase.GetUserByID(c.Request().Context(), request.RefreshToken)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusUnauthorized, constant.ErrUnauthorized, constant.ErrNotFoundUser, err)
	}

	accessToken, err := uc.UserUsecase.CreateAccessToken(&user, uc.Env.ACCESS_TOKEN_SECRET, uc.Env.ACCESS_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrAccessToken, err.Error())
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(&user, uc.Env.REFRESH_TOKEN_SECRET, uc.Env.REFRESH_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrRefreshToken, err.Error())
	}

	response := entities.ToRefreshTokenResponseUser(accessToken, refreshToken)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessRefreshToken, response)
}

func (uc *UserController) Update(c echo.Context) error {
	var request dto.UpdateUserRequest

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	updatedUser, err := uc.UserUsecase.Update(uuid.MustParse(c.Get("x-user-id").(string)), &request)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrUpdateUser, err.Error())
	}

	response := entities.EntityToDtoUser(updatedUser)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessUpdateUser, response)
}

func (uc *UserController) Dashboard(c echo.Context) error {
	dashboardData, err := uc.UserUsecase.GetDashboardData(c.Get("x-user-id").(string))

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGetDashboardData, err.Error())
	}

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessGetDashboardData, dashboardData)
}

func (uc *UserController) SendMessageChatbot(c echo.Context) error {

	var request dto.ChatbotRequest

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	chatbotConfig := entities.ChatbotConfig{
		APIKey: uc.Env.OPENAI_API_KEY,
	}

	reply, err := uc.ChatbotUsecase.SendMessage(chatbotConfig, request)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrSendMessageChatbot, err.Error())
	}

	response := dto.ChatbotResponse{
		Message:  request.Message,
		Response: reply,
	}

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessSendMessageChatbot, response)
}
