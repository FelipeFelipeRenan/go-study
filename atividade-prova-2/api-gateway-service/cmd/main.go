package main

import (
    "github.com/gin-gonic/gin"
    "foods/internal/service"
    "foods/internal/handlers"
    "foods/internal/transport"
)

func main() {
    // Inicialize os serviços de comida e pedidos
    foodService := service.NewFoodService()
    orderService := service.NewOrderService()

    // Crie os endpoints para cada serviço
    foodEndpoints := endpoint.MakeFoodEndpoints(foodService)
    orderEndpoints := endpoint.MakeOrderEndpoints(orderService)

    // Crie um roteador HTTP com Gin
    router := gin.Default()

    // Adicione os handlers HTTP para cada endpoint
    router.GET("/foods", transport.MakeGetAllFoodsHandler(foodEndpoints.GetAllFoods))
    // Outras rotas...

    // Inicie o servidor do API Gateway
    router.Run(":8080")
}
