package MiddleWare

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"net/http"
)

const (
	csrfHeader = "CSRF-Token" //header字段名
	csrfField  = "csrf.Token" //http form中字段名
	csrfCookie = "csrf.Token" //cookie中字段名
)

//
// CSRF
// @Description: csrf验证中间件
// @return gin.HandlerFunc
//
func CSRF() gin.HandlerFunc {
	CSRFMd := csrf.Protect(
		[]byte("e10adc3949ba59abbe56e057f20f883e"),
		csrf.Secure(false),
		csrf.HttpOnly(true),
		csrf.RequestHeader(csrfHeader), //header字段名
		csrf.FieldName(csrfField),      //http form中字段名
		csrf.CookieName(csrfCookie),    //cookie中字段名
		csrf.ErrorHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusForbidden)
			_, _ = writer.Write([]byte(`{"message":"Forbidden - CSRF token invalid"}`))
		})),
	)

	return Wrap(CSRFMd)

	/*return func(c *gin.Context) {
		fmt.Println("csrf验证。。。")
	}*/

}

//
// CSRFToken
// @Description: csrf验证中间件
// @return gin.HandlerFunc
//
func CSRFToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(csrfHeader, csrf.Token(c.Request))
	}
}

////参考https://github.com/gwatts/gin-adapter
//默认http转换gin http
type connectHandler struct{}
type middlewareCtx struct {
	ctx         *gin.Context
	childCalled bool
}

//
// New
// @Description: 重写gin 中间件return http return
// @return http.Handler
// @return func(h http.Handler) gin.HandlerFunc
//
func New() (http.Handler, func(h http.Handler) gin.HandlerFunc) {
	nextHandler := new(connectHandler)
	makeGinHandler := func(h http.Handler) gin.HandlerFunc {
		return func(c *gin.Context) {
			state := &middlewareCtx{ctx: c}
			ctx := context.WithValue(c.Request.Context(), nextHandler, state)
			h.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
			if !state.childCalled {
				c.Abort()
			}
		}
	}
	return nextHandler, makeGinHandler
}

//
// ServeHTTP
// @Description: 重写gin http
// @receiver h
// @param w
// @param r
//
func (h *connectHandler) ServeHTTP(_ http.ResponseWriter, r *http.Request) {
	state := r.Context().Value(h).(*middlewareCtx)
	defer func(r *http.Request) { state.ctx.Request = r }(state.ctx.Request)
	state.ctx.Request = r
	state.childCalled = true
	state.ctx.Next()
}

//
// Wrap
// @Description:http中间件转gin中间件http
// @param f
// @return gin.HandlerFunc
//
func Wrap(f func(h http.Handler) http.Handler) gin.HandlerFunc {
	next, adapter := New()
	return adapter(f(next))
}
