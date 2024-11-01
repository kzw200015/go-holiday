package server

import "go-holiday/server/holiday"

func bootstrap() error {
	err := holiday.LoadHolidaysFromRemote()
	if err != nil {
		return err
	}

	return nil
}
