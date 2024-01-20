package delivery

import (
	"promptgo/internal/chat/usecase"
	"promptgo/internal/prompt/dto/request"
	"promptgo/util/helper"

	"github.com/gin-gonic/gin"
)


type chatDelivery struct {
	chatUsecase usecase.UsecaseInterface
}

func NewChatDelivery(chatUsecase usecase.UsecaseInterface) *chatDelivery {
	return &chatDelivery{
		chatUsecase: chatUsecase,
	}
}


func (chat *chatDelivery) GetCompletion(c *gin.Context) {
	
	input := request.RequestQuestion{}

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	data, err := chat.chatUsecase.Completion(input.Question)
	if err != nil {
		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessWithDataResponse("success get all data", data))
}
