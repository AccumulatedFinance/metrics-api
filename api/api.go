package api

import (
	"math"
	"net/http"
	"strconv"

	"github.com/AccumulatedFinance/metrics-api/store"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const API_VERSION = "1.0"

type API struct {
	HTTP     *echo.Echo
	Validate *validator.Validate
}

type ErrorResponse struct {
	Result bool   `json:"result"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

// StartAPI configures and starts REST API server
func StartAPI(port int) error {

	api := &API{}

	api.HTTP = echo.New()
	api.HTTP.HideBanner = true

	// init validator v10
	api.Validate = validator.New()

	// remove trailing slash middleware
	// https://echo.labstack.com/middleware/trailing-slash/
	api.HTTP.Pre(middleware.RemoveTrailingSlash())

	// recover middleware
	// https://echo.labstack.com/middleware/recover/
	api.HTTP.Use(middleware.Recover())

	// logger middleware
	// https://echo.labstack.com/middleware/logger/
	api.HTTP.Use(middleware.Logger())

	// v1 public metrics API
	api.HTTP.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, API_VERSION)
	})
	publicAPI := api.HTTP.Group("/v1")

	publicAPI.GET("/supply", api.getStACMESupply)

	api.HTTP.Logger.Fatal(api.HTTP.Start(":" + strconv.Itoa(port)))

	return nil

}

func (api *API) getStACMESupply(c echo.Context) error {

	res := int(math.Round(float64(store.StACME.TotalSupply) * math.Pow10(-1*int(store.StACME.Decimals))))

	return c.String(http.StatusOK, strconv.Itoa(res))

}
