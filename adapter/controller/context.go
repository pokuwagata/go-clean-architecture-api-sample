package controller

type Context interface {
	Param(string) string
	Bind(interface{}) // gin, echoのI/Fではerror型を返り値に持つ
	Status(int)
	JSON(int, interface{})
}