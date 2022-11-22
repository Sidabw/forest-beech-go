package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Test1HelloWorld(context echo.Context) error {
	return context.String(http.StatusCreated, "echo framework success.")
}

func Test2GetParams(context echo.Context) error {
	param := context.Param("id")
	fmt.Println("path-variable:Id ", param)

	queryParamA := context.QueryParam("a")
	fmt.Println("params a  from query string, value: ", queryParamA)

	return context.String(http.StatusCreated, "echo framework success in Test2GetParams")
}
