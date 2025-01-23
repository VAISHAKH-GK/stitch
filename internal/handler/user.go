package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/FulgurCode/stitch/models"
	"github.com/FulgurCode/stitch/pkg/mysql"
	"github.com/FulgurCode/stitch/utils"
	"github.com/FulgurCode/stitch/view/layout"
	"github.com/FulgurCode/stitch/view/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/razorpay/razorpay-go"
)

// Home page handler
func Home(c echo.Context) error {
	var products, err = mysql.GetProducts()
	if err != nil {
		fmt.Println(err)
	}

	var settings = mysql.GetSettings()

	var component = user.Home(products, settings)

	return utils.Render(c, component)
}

// Product page handler
func Products(c echo.Context) error {
	var query = c.QueryParam("search")
	var products = mysql.SearchProduct(query)

	var component = user.Products(products)
	return utils.Render(c, component)
}

// Item page handler
func Item(c echo.Context) error {
	var productId = c.Param("productId")
	var product, err = mysql.GetProductById(productId)
	if err != nil {
		fmt.Println(err)
	}

	stock, err := mysql.GetStock(productId)
	if err != nil {
		fmt.Println(err)
	}

	var component = user.Item(product, stock)
	return utils.Render(c, component)
}

// Order page handler
func OrderGet(c echo.Context) error {
	var productId = c.Param("productId")
	var size = c.QueryParam("size")

	var product, err = mysql.GetProductById(productId)
	if err != nil {
		fmt.Println(err)
	}

	var component = user.Order(product, size)
	return utils.Render(c, component)
}

func OrderPost(c echo.Context) error {
	var productId = c.Param("productId")
	var order models.Order
	var err = c.Bind(&order)
	if err != nil {
		fmt.Println(err)
	}

	product, err := mysql.GetProductById(productId)
	if err != nil {
		fmt.Println(err)

		if c.Request().Header.Get("HX-Request") == "true" {
			c.Response().Header().Set("HX-Location", "/products")
			return c.NoContent(http.StatusSeeOther)
		}
		c.Redirect(http.StatusSeeOther, "/products")
	}

	order.ProductId = productId
	if order.Payment == "online" {
		order.Status = "pending"
	} else {
		order.Status = "ordered"
	}
	if order.Quantity == 0 {
		order.Quantity = 1
	}
	order.Total = product.Price * order.Quantity

	var orderId = uuid.New()
	order.Id = string(orderId.String())
	err = mysql.MakeOrder(order, orderId)
	if err != nil {
		fmt.Println(err)
	}

	if order.Payment == "online" {
		var key = os.Getenv("RAZORPAY_KEY")
		var secret = os.Getenv("RAZORPAY_SECRET")
		var client = razorpay.NewClient(key, secret)
		var ord = map[string]interface{}{
			"amount":   order.Total * 100,
			"currency": "INR",
		}
		body, err := client.Order.Create(ord, nil)
		if err != nil {
			fmt.Println(err)
		}

		var component = user.OnlinePayment(order, body, key)
		return utils.Render(c, component)
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", fmt.Sprintf("/item/%s", productId))
		return c.NoContent(http.StatusSeeOther)
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/item/%s", productId))
}

// Search page handler
func Search(c echo.Context) error {
	var query = c.QueryParam("search")
	var products = mysql.SearchProduct(query)

	var component = user.Products(products)
	return utils.Render(c, component)
}

// Cart page handler
func Cart(c echo.Context) error {
	var cart = utils.GetSessionAll(c, "cart")
	var products []models.Product
	for id, size := range cart {
		var product, err = mysql.GetProductById(id.(string))
		product.Size = size.(string)
		if err != nil {
			fmt.Println(err)
			var component = user.Cart(products)
			return utils.Render(c, component)
		}
		products = append(products, product)
	}

	var component = user.Cart(products)
	return utils.Render(c, component)
}

// About page handler
func About(c echo.Context) error {
	var component = user.About()

	return utils.Render(c, component)
}

// Comming Soon handler
func CommingSoon(c echo.Context) error {
	var component = layout.CommmingSoon()

	return utils.Render(c, component)
}

// Not Found handler
func NotFound(c echo.Context) error {
	var component = layout.NotFound()

	return utils.Render(c, component)
}

func AddToCart(c echo.Context) error {
	var id = c.Param("productId")
	var size = c.QueryParam("size")
	utils.CreateSession(c, "cart")

	utils.AddSessionValue(c, "cart", id, size)

	return c.NoContent(200)
}

func DeleteFromCart(c echo.Context) error {
	var id = c.Param("productId")
	utils.DeleteSessionValue(c, "cart", id)
	return c.NoContent(200)
}

func CartOrderGet(c echo.Context) error {
	var cart = utils.GetSessionAll(c, "cart")
	var products []models.Product
	for id, size := range cart {
		var product, err = mysql.GetProductById(id.(string))
		product.Size = size.(string)
		if err != nil {
			fmt.Println(err)
			var component = user.Cart(products)
			return utils.Render(c, component)
		}
		products = append(products, product)
	}

	var component = user.CartOrder(products)
	return utils.Render(c, component)
}

func CartOrderPost(c echo.Context) error {
	var body models.Order
	var err = c.Bind(&body)
	if err != nil {
		fmt.Println(err)
	}

	var cart = utils.GetSessionAll(c, "cart")
	var orderId = uuid.New()
	body.Id = orderId.String()
	var total int = 0
	for id, size := range cart {
		var product, _ = mysql.GetProductById(id.(string))
		var order models.Order = body

		order.ProductId = product.Id
		if body.Payment == "online" {
			order.Status = "pending"
		} else {
			order.Status = "ordered"
		}
		order.Quantity = 1
		order.Total = product.Price
		order.Size = size.(string)

		err = mysql.MakeOrder(order, orderId)
		if err != nil {
			fmt.Println(err)
		}
		total += order.Total
	}
	if body.Payment == "online" {
		var key = os.Getenv("RAZORPAY_KEY")
		var secret = os.Getenv("RAZORPAY_SECRET")
		var client = razorpay.NewClient(key, secret)
		var ord = map[string]interface{}{
			"amount":   total * 100,
			"currency": "INR",
		}
		order, err := client.Order.Create(ord, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("submit cart")

		var component = user.OnlinePaymentCart(body, order, key)
		return utils.Render(c, component)
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", "/cart")
		return c.NoContent(http.StatusSeeOther)
	}

	return c.Redirect(http.StatusSeeOther, "/cart")
}

func VerifyOnlinePayment(c echo.Context) error {
	var orderId = c.Param("orderId")
	var data models.VerifyOrder
	var err = c.Bind(&data)
	if err != nil {
		fmt.Println(err)
	}

	if err := utils.RazorPaymentVerification(data); err == nil {
		mysql.OrderStatusPayment(orderId, "ordered")
		return c.JSON(200, "Ordered Successfully And Payment Verified")
	}

	fmt.Println(err)
	return c.JSON(406, "Payment Verification Failed")
}

func DeletePendingOrder(c echo.Context) error {
	var orderId = c.Param("orderId")
	var err = mysql.DeletePendingOrder(orderId)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "Failed")
	}

	return c.JSON(200, "Order Payment Failed / Cancelled")
}
