package route

import (
	"promptgo/internal/chat/delivery"
	"promptgo/internal/prompt/repository"
	"promptgo/internal/chat/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteChat(c *gin.RouterGroup, db *gorm.DB) {
	promptRepository := repository.NewPromptRepository(db)
	chatUsecase := usecase.NewChatUsecase(promptRepository)
	chatDelivery := delivery.NewChatDelivery(chatUsecase)

	chat := c.Group("chat")
	{
		chat.POST("/completion", chatDelivery.GetCompletion)
	}
}