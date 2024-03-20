package middlewares

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {

	return func (ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin","*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type,Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Method","POST, GET, DELETE, PUT")
		
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		
		ctx.Next()	
	}
}