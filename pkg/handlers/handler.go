package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Группа для аутентификации
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)  
		auth.POST("/sign-in", h.signIn)
	}

	// Группа для API
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)             
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group("/items")
			{
				items.POST("/:id", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:items_id", h.getItemById)
				items.PUT("/:items_id", h.updateItem)
				items.DELETE("/:items_id", h.deleteItem)
			}
		}
	}

	return router
}
