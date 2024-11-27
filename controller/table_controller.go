package controller

import (
	"api-blog-go/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableController struct {
	tableUsecase usecase.TableUsecase
}

func NewTableController(usecase usecase.TableUsecase) TableController {
	return TableController{
		tableUsecase: usecase,
	}
}

func (t *TableController) GetTables(ctx *gin.Context) {
	tables, err := t.tableUsecase.GetTables()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, tables)
}
