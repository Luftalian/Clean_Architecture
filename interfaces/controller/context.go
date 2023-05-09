package controller

type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(key string) string
}
