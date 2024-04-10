package router

import (
	"app02/controllers"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo     *echo.Echo
	Loginout controllers.Loginout
}

func (api *API) SetUpRouter() {
	api.Echo.GET("/api/form", api.Loginout.Formregister)
	api.Echo.POST("/api/register", api.Loginout.Register)
	api.Echo.GET("/api/formlogin", api.Loginout.FormLogin)
	api.Echo.POST("/api/login", api.Loginout.Login)

}
