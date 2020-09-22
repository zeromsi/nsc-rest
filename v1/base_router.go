package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

func Routes(e *echo.Echo) {
	e.GET("/", index)

	// Health Page
	e.GET("/health", health)

	v1AccountsMonitor := e.Group("/api/v1/accounts")
	AccountRouter(v1AccountsMonitor)
	//v1OperatorsMonitor := e.Group("/api/v1/operators")
	//OperatorRouter(v1OperatorsMonitor)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "This is KloverCloud nats service")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I am live!")
}
