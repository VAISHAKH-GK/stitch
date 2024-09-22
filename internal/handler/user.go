package handler

import (
	"github.com/FulgurCode/stitch/utils"
	"github.com/FulgurCode/stitch/view/layout"
	"github.com/FulgurCode/stitch/view/user"
	"github.com/labstack/echo/v4"
)

// Home page handler
func Home(c echo.Context) error {
	var component = user.Home()

	return utils.Render(c, component)
}

// Product page handler
func Products(c echo.Context) error {
	var component = user.Products()

	return utils.Render(c, component)
}

// Item page handler
func Item(c echo.Context) error {
	var component = user.Item()

	return utils.Render(c, component)
}

// Item page handler
func Order(c echo.Context) error {
	var component = user.Order()

	return utils.Render(c, component)
}

// Not Found handler
func NotFound(c echo.Context) error {
	var component = layout.NotFound()

	return utils.Render(c, component)
}
