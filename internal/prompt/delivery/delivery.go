package delivery

import (
	"net/http"
	"promptgo/internal/prompt/dto/request"
	"promptgo/internal/prompt/usecase"
	"promptgo/util/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type promptDelivery struct {
	promptUsecase usecase.UsecaseInterface
}

func NewPromptDelivery(promptUsecase usecase.UsecaseInterface) *promptDelivery {
	return &promptDelivery{
		promptUsecase: promptUsecase,
	}
}

func (prompt *promptDelivery) PostInstruction(c *gin.Context) {
	input := request.RequestPrompt{}

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		return
	}

	err = prompt.promptUsecase.Insert(input)
	if err != nil {
		if strings.Contains(err.Error(), "") {
			c.AbortWithStatusJSON(400,helper.ErrorResponse(err.Error()))
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func (prompt *promptDelivery) GetAllInstructions(c *gin.Context) {
	question := c.Query("question")

	data, err := prompt.promptUsecase.SelectAll(question)
	if err != nil {
		if strings.Contains(err.Error(), "") {
			c.AbortWithStatusJSON(400,helper.ErrorResponse(err.Error()))
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
	}

	c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all data", data))
}