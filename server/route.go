package server

import (
	"github.com/gin-gonic/gin"
	"go-holiday/server/holiday"

	"time"
)

func setRoutes(engine *gin.Engine) {
	group := engine.Group("/holidays")
	{
		group.GET("/is-holiday", func(context *gin.Context) {
			date := context.Query("date")
			if date == "" {
				date = time.Now().Format("2006-01-02")
			}
			isHoliday, err := holiday.IsHoliday(date)
			if err != nil {
				sendErr(context, err)
				return
			}
			sendOk(context, gin.H{
				"isHoliday": isHoliday,
			})
		})
	}
}
