package framework

import (
	"api/adapter/viewmodel"
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{w, r}
}

// TODO: 実装
func (c *Context) Param(string) string {
	return ""
}

func (c *Context) Bind(obj interface{}) {
	len := c.Request.ContentLength
	body := make([]byte, len)
	c.Request.Body.Read(body)
	err := json.Unmarshal(body, &obj)
	if err != nil {
		c.JSON(400, viewmodel.Error{Code: 400, Msg: "invalid parameter"})
	}
}

func (c *Context) Status(status int) {
	c.Writer.WriteHeader(status)
}

func (c *Context) JSON(status int, obj interface{}) {
	c.Writer.WriteHeader(status)
	c.Writer.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&obj, "", "\t\t")
	if err != nil {
		c.JSON(500, viewmodel.Error{Code: 500, Msg: "interanl server error"})
	}
	c.Writer.Write(output)
}
