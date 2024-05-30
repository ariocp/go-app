package v1

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()

	router.POST("sign-up", h.signUp)
	router.POST("sign-in", h.signIn)

	return router
}
