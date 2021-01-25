package api

import (
	"database/sql"
	"learn-go/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var task model.Task

		c.Bind(&task)

		id, err := model.PutTask(db, task.Name, task.Status)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}

	}
}

func EditTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var task model.Task
		c.Bind(&task)

		_, err := model.EditTask(db, task.ID, task.Name, task.Status)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": task,
			})
		} else {
			return err
		}
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := model.DeleteTask(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}

	}
}
