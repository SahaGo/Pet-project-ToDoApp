package handler // имплементируем наши хендлеры

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine { // инициализирует все наши эндпоинты
	router := gin.New()

	auth := router.Group("/auth") // объявим методы, сгруппировав их по маршрутам
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById) // используя двоеточие в марруте эндпоинта, мы указываем, что тут может быть любое значение, которому мы можем обратиться, при помощи по имени параметра id, это фишка библиотеки гин(джин)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
