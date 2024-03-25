package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mojocn/base64Captcha"
)

type CaptchaResult struct {
	CaptchaId string `json:"captchaId"`
	Image     string `json:"image"`
}

// GenerateCaptchaHandler is a handler to create captcha
func GenerateCaptchaHandler(c echo.Context) error {
	// create a new captcha object by calling NewDriverDigit function
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	// get the id and the base64 string of captcha image
	id, b64s, answer, err := cp.Generate()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(answer)

	// return the captchaId and the base64 image string
	return c.JSON(http.StatusOK, &CaptchaResult{
		CaptchaId: id,
		Image:     b64s,
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/captcha", GenerateCaptchaHandler)

	e.Start(":8080")
}
