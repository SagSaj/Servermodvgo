package timeoutdrop

import (
	"log"
	"main/modules/config"
	"main/modules/subdmongo"
	"time"
)

var t time.Time

func Initialize() {
	t = time.Now()
	t = time.Date(t.Year(), t.Month(), t.Day(), int(0), int(0), int(0), int(0), time.UTC)
	go TimeCount()
}
func TimeCount() {
	for {
		//log.Println("Run timeoutdrop")
		//log.Println(time.Now().UnixNano())
		//log.Println(t.AddDate(0, 0, config.Conf.Days_period).UnixNano())
		if time.Now().UnixNano() > t.AddDate(0, 0, config.Conf.Days_period).UnixNano() {
			//DropNew
			if t.Hour() == config.Conf.Time_reload {
				log.Println("Clone and update")
				subdmongo.Clone()
				subdmongo.UpdateAllForDefault()
				t = time.Date(t.Year(), t.Month(), t.Day(), int(0), int(0), int(0), int(0), time.UTC)
			}
		}

		time.Sleep(time.Minute * 2)
		for time.Now().Minute() < 2 {
			time.Sleep(time.Minute * 1)
		}

	}
}
