package index

import (
	"emoji/pkg/config"
	"emoji/pkg/unity"
	"os"
	"time"
)

type Task struct {
	Path string
	Tick *time.Ticker
}

func NewTask(path string)*Task  {
	return &Task{
		Path:path,
		Tick:time.NewTicker(time.Minute * 60),
	}
}

func (this *Task)DeleteExpireAssFile()  {
	go func() {
		for{
			dir := this.Path + time.Now().Add(-24 * time.Hour).Format(config.DateFormat)
			select {
			case  <- this.Tick.C:
				if unity.DirExistValidate(dir){
					err := os.RemoveAll(dir)
					unity.ErrorCheck(err)
				}
			}
		}
	}()
}
