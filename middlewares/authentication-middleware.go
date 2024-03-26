package middlewares

import (
	"basictrade-gading/utils/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Authentication() gin.HandlerFunc{

	return func(ctx *gin.Context) {

		verifyToken, err := helpers.VerifyToken(ctx)
		if err != nil {
			log.Error().Msg(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthenticated",
				"messsage": err.Error(),
			})
			return
		}

		ctx.Set("adminData", verifyToken)
		ctx.Next()

	}

}