package tasks

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/entity"
	"github.com/hayato24s/todo-echo-gorm/usecase"

	"github.com/labstack/echo/v4"
)

type IndexRes struct {
	Tasks []TaskRes `json:"tasks"`
	Total uint64    `json:"total" example:"10"`
}

func ToIndexRes(tasks []entity.Task, total uint64) IndexRes {
	res := IndexRes{
		Total: total,
	}
	res.Tasks = make([]TaskRes, len(tasks))
	for i, t := range tasks {
		res.Tasks[i] = ToTaskRes(&t)
	}
	return res
}

// Index
//
//	@Tags		tasks
//	@Produce	json
//	@Success	200	{object}	IndexRes
//	@Failure	401	{object}	common.ErrorRes
//	@Failure	500	{object}	common.ErrorRes
//	@Router		/tasks [get]
func (h *Handler) Index(c echo.Context) error {
	tasks, total, err := h.uc.GetTasks(c.Request().Context(), usecase.GetTasksIn{Offset: 0, Limit: 20})
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, ToIndexRes(tasks, total))
	return nil
}
