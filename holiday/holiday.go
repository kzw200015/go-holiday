package holiday

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"log/slog"
	"strconv"
	"strings"
	"sync"
	"time"
)

var client = resty.New()

var holidays = map[string]*Holiday{}
var mu sync.RWMutex

func IsHoliday(date string) (bool, error) {
	mu.RLock()
	h := holidays[date]
	mu.RUnlock()
	if h != nil {
		return h.IsOffDay, nil
	} else {
		parsed, err := time.Parse("2006-01-02", date)
		if err != nil {
			return false, err
		}

		return parsed.Weekday() == time.Saturday || parsed.Weekday() == time.Sunday, nil
	}
}

func LoadHolidaysFromRemote() error {
	response, err := client.R().Get(strings.ReplaceAll(viper.GetString("holiday.remote_url"), "{year}", strconv.Itoa(time.Now().Year())))
	if err != nil {
		slog.Error("failed to load holidays from remote", "error", err)
		return err
	}
	mu.Lock()
	for _, ele := range gjson.GetBytes(response.Body(), "days").Array() {
		name := ele.Get("name").String()
		date := ele.Get("date").String()
		isOffDay := ele.Get("date").Bool()
		holidays[date] = &Holiday{
			Name:     name,
			Date:     date,
			IsOffDay: isOffDay,
		}
	}
	mu.Unlock()

	slog.Info("loaded holidays from remote")
	return nil
}

type Holiday struct {
	Name     string
	Date     string
	IsOffDay bool
}
