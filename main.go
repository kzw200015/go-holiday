package main

import (
	"github.com/gin-gonic/gin"
	"go-holiday/holiday"
	"go-holiday/utils/http"
	"log/slog"
	"time"
)

func main() {
	initConfig()
	err := holiday.LoadHolidaysFromRemote()
	if err != nil {
		panic(err)
	}
	err = startSchedule()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	group := r.Group("/holidays")
	{
		group.GET("/is-holiday", func(context *gin.Context) {
			date := context.Query("date")
			if date == "" {
				date = time.Now().Format("2006-01-02")
			}
			isHoliday, err := holiday.IsHoliday(date)
			if err != nil {
				http.Error(context, err)
				return
			}
			http.Ok(context, gin.H{
				"isHoliday": isHoliday,
			})
		})
	}
	err = r.Run()
	if err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
