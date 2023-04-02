package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
	"student-app/internal/model"
	"student-app/internal/request"
	"student-app/internal/store"
)

type Student struct {
	Store  store.Students
	Logger *zap.Logger
}

func (s Student) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	ss, err := s.Store.GetAll(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, ss)
}

func (s Student) Get(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	st, err := s.Store.Get(ctx, id)
	if err != nil {
		var errNotFound store.StudentNotFoundError

		if ok := errors.As(err, &errNotFound); ok {
			s.Logger.Error("student not found", zap.Error(err))

			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, st)
}

func (s Student) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.Student
	if err := c.Bind(&req); err != nil {
		body, _ := io.ReadAll(c.Request().Body)
		s.Logger.Error("cannot bind request to student", zap.Error(err), zap.Any("body", body))

		return echo.ErrBadRequest
	}
	if err := req.Validate(); err != nil {
		s.Logger.Error("request validation failed", zap.Error(err), zap.Any("request", req))

		return echo.ErrBadRequest
	}

	id, _ := strconv.ParseUint(req.ID, 10, 64)

	stu := model.Student{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Average:   0,
	}

	if err := s.Store.Save(ctx, stu); err != nil {
		var errDuplicateErr store.StudentDuplicateError
		if ok := errors.As(err, &errDuplicateErr); ok {
			s.Logger.Error("duplicate student", zap.Error(err), zap.Uint64("id", id))

			return echo.ErrBadRequest
		}

		return echo.ErrInternalServerError
	}

	s.Logger.Info("student creation success")

	return c.JSON(http.StatusCreated, stu)
}

func (s Student) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	if err := s.Store.Delete(ctx, id); err != nil {
		var errNotFound store.StudentNotFoundError

		if ok := errors.As(err, &errNotFound); ok {
			s.Logger.Error("student not found", zap.Error(err))

			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	s.Logger.Info("student deletion success")

	type responseMsg struct {
		Message string `json:"message"`
	}

	m := responseMsg{
		Message: fmt.Sprintf("student with id %d successfuly deleted", id),
	}

	return c.JSON(http.StatusOK, &m)
}

func (s Student) Register(g *echo.Group) {
	g.GET("", s.GetAll)
	g.GET("/:id", s.Get)
	g.POST("", s.Create)
	g.DELETE("/:id", s.Delete)
}
