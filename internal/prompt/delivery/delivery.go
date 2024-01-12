package delivery

import (
	"promptgo/internal/prompt/dto/request"
	"promptgo/internal/prompt/usecase"
	"promptgo/util/constanta"
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
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = prompt.promptUsecase.Insert(input)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(201, helper.SuccessResponse("success insert data"))
}

func (prompt *promptDelivery) GetAllInstructions(c *gin.Context) {

	data, err := prompt.promptUsecase.SelectAll()
	if err != nil {
		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	if len(data) == 0 {
		c.JSON(200, helper.SuccessResponse(constanta.SUCCESS_NUll))
		return
	}

	c.JSON(200, helper.SuccessWithDataResponse("success get all data", data))
}

func (prompt *promptDelivery) GetInstructionById(c *gin.Context) {

	id := c.Param("id")

	data, err := prompt.promptUsecase.SelectByID(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_NOT_FOUND) {
			c.AbortWithStatusJSON(404, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessWithDataResponse("success get data", data))
}

func (prompt *promptDelivery) UpdatePrompt(c *gin.Context) {

	input := request.RequestPrompt{}
	id := c.Param("id")

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = prompt.promptUsecase.Update(id, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_NOT_FOUND) {
			c.AbortWithStatusJSON(404, helper.ErrorResponse(err.Error()))
			return
		}

		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessResponse("success update data"))
}

func (prompt *promptDelivery) DeletePrompt(c *gin.Context) {

	id := c.Param("id")

	err := prompt.promptUsecase.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_NOT_FOUND) {
			c.AbortWithStatusJSON(404, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessResponse("success delete data"))
}
