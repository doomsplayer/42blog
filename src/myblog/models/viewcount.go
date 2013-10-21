package models

import (
	"strconv"
	"time"
)

var ViewCountCltn viewCountCltn

type viewCountCltn struct{}

func (viewCountCltn) GetTodayView() int {
	i, _ := strconv.Atoi(string(Meta.Get(`view_today`).([]uint8)))
	return i
}

func (viewCountCltn) GetAllView() int {
	i, _ := strconv.Atoi(string(Meta.Get(`view_all`).([]uint8)))
	return i
}

func (viewCountCltn) LastViewDate() string {
	t := Meta.Get(`today`)
	if t == nil {
		Meta.Put(`today`, time.Now().Format(`2006-01-02`), 0)
		return time.Now().Format(`2006-01-02`)
	}

	return string(t.([]uint8))
}

func (viewCountCltn) IncrView() {
	if ViewCountCltn.LastViewDate() != time.Now().Format(`2006-01-02`) {
		Meta.Put(`today`, time.Now().Format(`2006-01-02`), 0)
		Meta.Put(`view_today`, 0, 0)
	}

	Meta.Incr(`view_today`)
	Meta.Incr(`view_all`)
}
