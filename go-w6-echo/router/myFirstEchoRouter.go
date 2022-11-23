package router

import (
	"fmt"
	"github.com/sidabw/go-w6-echo/pojo"
	"io"
	"net/http"
	"os"

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

func Test3PostForm(context echo.Context) error {
	aValue := context.FormValue("a")
	fmt.Println("aValue from form k-v pair: ", aValue)

	return context.String(http.StatusCreated, "echo framework success in Test3PostForm")
}

func Test4PostJson(context echo.Context) error {
	user := new(pojo.User)

	if err := context.Bind(user); err != nil {
		return err
	}

	fmt.Printf("json body map: id: %v, name: %v \n", user.Id, user.Name)

	//以json格式返回
	res := pojo.User{Id: 1, Name: "dd"}
	return context.JSON(http.StatusOK, res)

	//返回html的可以context.HTML
}

func Test5Download(context echo.Context) error {
	//其他Blob字节数组形式未列举
	return context.Attachment("README.md", "r.md")
}

func Test6Upload(context echo.Context) error {
	file, err := context.FormFile("file")
	if err != nil {
		return err
	}
	//先打开文件源
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	nasusFile, err := os.OpenFile("/var/www/dream/webapp/logs/nasus3.log",
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	_, err2 := io.Copy(nasusFile, src)
	if err2 != nil {
		return err2
	}

	return context.String(200, "success on copy")

}
