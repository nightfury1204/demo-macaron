package routers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/nightfury1204/demo-macaron/models"
	"gopkg.in/macaron.v1"
)

func GetAllBooks(ctx *macaron.Context) {
	var books []models.Book
	engine, err := models.GetDBEngine()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	if err := engine.Find(&books); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, books)
}

func GetBook(ctx *macaron.Context) {
	id := ctx.Params(":id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "id must be non empty",
		})
		return
	}

	var b models.Book
	engine, err := models.GetDBEngine()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if has, err := engine.ID(id).NoAutoCondition().Get(&b); has && err == nil {
		ctx.JSON(http.StatusOK, b)
		return
	} else {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, map[string]string{
				"error": "book not found",
			})
		}
		return
	}
}

func CreateBook(ctx *macaron.Context) {
	b := &models.Book{}
	payload, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = json.Unmarshal(payload, b)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	b.ID = rand.Int()

	err = b.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	engine, err := models.GetDBEngine()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if _, err := engine.InsertOne(b); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"msg": "book added successfully",
	})
}

func EditBook(ctx *macaron.Context) {
	var oldB models.Book
	idS := ctx.Params(":id")
	if idS == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "id must be non empty",
		})
		return
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	oldB.ID = id

	engine, err := models.GetDBEngine()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if has, err := engine.Get(&oldB); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	} else if !has {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "book does not exists",
		})
		return
	}

	newB := &models.Book{}
	payload, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = json.Unmarshal(payload, newB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	finalB := models.Merge(oldB, *newB)
	if _, err := engine.Update(finalB); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"msg": "book updated successfully",
	})
}

func DeleteBook(ctx *macaron.Context) {
	idS := ctx.Params(":id")
	if idS == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "id must be non empty",
		})
		return
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	engine, err := models.GetDBEngine()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if _, err := engine.Delete(&models.Book{ID: id}); err == nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "book deleted successfully",
		})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
}
