package routes

import(
	"github.com/gin-gonic/gin"
	
	controller "golang-restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
	incomingRoutes.DELETE("/allorderitems", controller.DeleteAllOrderItems())

}