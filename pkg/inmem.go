package pkg

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type Inmem struct {
	mtx sync.Mutex
	store map[string]Book
}

func NewInmemStorage() *Inmem {
	return &Inmem{
		store: map[string]Book{},
	}
}


func (s *Inmem) GetAllBooks(ctx *macaron.Context) {
	ctx.JSON(200, s.store)
}

func (s *Inmem) GetBook(ctx *macaron.Context) {
	id := ctx.Params(":id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "id must be non empty",
		})
		return
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	if b, ok := s.store[id]; ok {
		ctx.JSON(http.StatusOK, b)
		return
	} else {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "book not found",
		})
		return
	}
}

func (s *Inmem) CreateBook(ctx *macaron.Context) {
	b := &Book{}
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

	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.store[strconv.Itoa(b.ID)] = *b
	ctx.JSON(http.StatusOK, map[string]string{
		"msg": "book added successfully",
	})
}

func (s *Inmem) EditBook(ctx *macaron.Context) {
	var oldB Book
	id := ctx.Params(":id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "id must be non empty",
		})
		return
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	if b, ok := s.store[id]; ok {
		oldB = b
	} else {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "book not found",
		})
		return
	}

	newB := &Book{}
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

	s.store[id] = Merge(oldB, *newB)
	ctx.JSON(http.StatusOK, map[string]string{
		"msg": "book updated successfully",
	})
}

func (s *Inmem) DeleteBook(ctx *macaron.Context) {
	id := ctx.Params(":id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "id must be non empty",
		})
		return
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.store[id]; ok {
		delete(s.store, id)
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "book deleted successfully",
		})
		return
	} else {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "book not found",
		})
		return
	}
}