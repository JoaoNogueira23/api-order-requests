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

	if tables == nil {
		response := model.Response{
			Message: "No tables found!",
		}
		ctx.JSON(http.StatusNoContent, response)
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

func (t *TableController) GetTableById(ctx *gin.Context) {
	id := ctx.Param("tableId")

	if id == "" {
		reponse := model.Response{
			Message: "Id cannot null!",
		}

		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	table, err := t.tableUsecase.GetTableById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if table == nil {
		response := model.Response{
			Message: "Table not found!",
		}

		ctx.JSON(http.StatusNotFound, response)
	}

	ctx.JSON(http.StatusOK, table)

}
