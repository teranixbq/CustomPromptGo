package route

import (
	"promptgo/internal/prompt/delivery"
	"promptgo/internal/prompt/repository"
	"promptgo/internal/prompt/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutePrompt(c *gin.RouterGroup, db *gorm.DB) {
	promptRepository := repository.NewPromptRepository(db)
	promptService := usecase.NewPromptUsecase(promptRepository)
	promptDelivery := delivery.NewPromptDelivery(promptService)

	prompt := c.Group("prompt")
	{
		prompt.POST("", promptDelivery.PostInstruction)
		prompt.GET("",promptDelivery.GetAllInstructions)
		prompt.GET("/:id",promptDelivery.GetInstructionById)
		prompt.PUT("/:id",promptDelivery.UpdatePrompt)
		prompt.DELETE("/:id",promptDelivery.DeletePrompt)
	}
}