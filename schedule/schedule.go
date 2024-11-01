package schedule

import (
	"github.com/robfig/cron/v3"
	"go-holiday/server/holiday"
	"log/slog"
)

var c = cron.New()

func StartSchedule() error {
	_, err := c.AddFunc("0 23 * * *", func() {
		err := holiday.LoadHolidaysFromRemote()
		if err != nil {
			slog.Error("failed to load holidays", "error", err)
		}
	})
	if err != nil {
		return err
	}

	return nil
}
