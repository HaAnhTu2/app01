package controllers

import (
	"app02/log"
	"app02/models"
	req "app02/models/req"
	"app02/reponsitory"
	"app02/security"
	"html/template"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Loginout struct {
	UserRepo reponsitory.UserRepo
}

func (u *Loginout) Formregister(c echo.Context) error {
	t, err := template.ParseFiles("view/form.html")
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return t.Execute(c.Response().Writer, err)
}

func (u *Loginout) Register(c echo.Context) error {
	//lấy dữ liệu từ form
	//kiểm tra và tạo tài khoản
	req := req.ReqUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	hash := security.HashAndSalt([]byte(req.Password))
	Id, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, models.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := models.User{
		Id:       Id.String(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hash,
	}
	user, err = u.UserRepo.CreateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, models.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Id = ""
	user.Password = ""
	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Tài khoản đã được tạo thành công",
		Data:       user,
	})
}

func (u *Loginout) FormLogin(c echo.Context) error {
	t, err := template.ParseFiles("view/login.html")
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return t.Execute(c.Response().Writer, err)
}

func (u *Loginout) Login(c echo.Context) error {
	req := req.ReqIn{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	// check pass
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Đăng nhập thất bại",
			Data:       nil,
		})
	}
	user.Password = ""
	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Tài khoản đăng nhập thành công",
		Data:       user,
	})
}
