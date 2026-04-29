package controller

import (
	"context"
	"errors"
	"lentera/internal/model"
	"lentera/internal/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Db repository.PgRepo
}

func (ct *Controller) CheckIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var req model.AttendaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := ct.Db.CheckIn(ctx, req)
	if err != nil {
		if errors.Is(err, repository.ErrEmployeeNotFound) || errors.Is(err, repository.ErrEmployeeCheckInAlready) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, model.AttendaceResponse{
		EmployeeId:  req.EmployeeId,
		AttendaceId: id,
	})

}

func (ct *Controller) CheckOut(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var req model.AttendaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := ct.Db.CheckOut(ctx, req)
	if err != nil {
		if errors.Is(err, repository.ErrEmployeeNotFound) || errors.Is(err, repository.ErrEmployeeNotCheckIn) || errors.Is(err, repository.ErrEmployeeCheckOutAlready) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, model.AttendaceResponse{
		EmployeeId:  req.EmployeeId,
		AttendaceId: id,
	})
}

func (ct *Controller) History(c *gin.Context) {
	id := c.Query("employee_id")

	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing required param"})
		return
	}

	empId, _ := strconv.Atoi(id)

	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "5")

	intPage, _ := strconv.Atoi(page)
	intSize, _ := strconv.Atoi(size)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := ct.Db.GetHistory(ctx, empId, intPage, intSize)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, res)

}
