package timeoutdrop

import "time"
import "main/modules/config"
import "main/modules/subdmongo"

var t time.Time

func Initialize() {
	t = time.Now()
	t = time.Date(t.Year(), t.Month(), t.Day(), int(0), int(0), int(0), int(0), nil)
	go TimeCount()
}
func TimeCount() {
	for {
		if time.Now().UnixNano() > t.AddDate(0, 0, config.Conf.Days_period).UnixNano() {
			//DropNew
			if t.Hour() == config.Conf.Time_reload {
				subdmongo.Clone()
				subdmongo.UpdateAllForDefault()
				t = time.Date(t.Year(), t.Month(), t.Day(), int(0), int(0), int(0), int(0), nil)
			}
		}
		time.Sleep(time.Minute * 2)
		for time.Now().Minute() < 2 {
			time.Sleep(time.Minute * 1)
		}

	}
}
