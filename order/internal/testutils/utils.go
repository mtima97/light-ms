package testutils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func GetRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	return gin.Default()
}

func ToReader(p map[string]any) *bytes.Reader {
	b, _ := json.Marshal(p)

	return bytes.NewReader(b)
}

func RespBody(r *http.Response) string {
	b, _ := io.ReadAll(r.Body)

	return string(b)
}
