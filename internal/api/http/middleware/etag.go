package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.body = append(w.body, data...)
	return w.ResponseWriter.Write(data)
}

func GenerateEtag(body []byte) string {
	hash := md5.Sum(body)
	return `W/"` + hex.EncodeToString(hash[:]) + `"`
}

func EtagMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodGet && ctx.Request.Method != http.MethodHead {
			ctx.Next()
			return
		}

		writer := &ResponseWriter{
			ResponseWriter: ctx.Writer,
			body:           []byte{},
		}
		ctx.Writer = writer

		ctx.Next()

		if ctx.Writer.Status() == http.StatusOK {
			ifNoneMatch := ctx.Request.Header.Get("If-None-Match")
			etag := GenerateEtag(writer.body)
			if ifNoneMatch != "" {
				if ifNoneMatch == etag {
					ctx.AbortWithStatus(http.StatusNotModified)
					return
				}
			}

			ctx.Writer.Header().Set("ETag", etag)
		}
	}
}
