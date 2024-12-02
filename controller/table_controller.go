package controller

import (
	"api-blog-go/model"
	"api-blog-go/usecase"
	"fmt"
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

func (t *TableController) CreateTable(ctx *gin.Context) {
	var table model.Table
	err := ctx.BindJSON(&table)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	insertedTable, err := t.tableUsecase.CreateTable(table)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTable)

}
