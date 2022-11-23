package pojo

type User struct {
	// 这些tag是echo框架给支持的
	Id   int    `json:"id" form:"id" query:"id"`
	Name string `json:"name"`
}
