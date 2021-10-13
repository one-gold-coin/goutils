package goutils

import (
	"time"
)

const (
	LayoutDefault = "2006-01-02 15:04:05" //常规类型
	LayoutYmd     = "2006-01-02"
	LayoutHis     = "15:04:05"

	//location
	LocationAsiaSh         = "Asia/Shanghai"
	LocationAmericaNewYork = "America/New_York"
)

func DefaultLocation() *time.Location {
	loc, _ := time.LoadLocation(LocationAsiaSh)
	return loc
}

func Now(name ...string) int64 {
	if name != nil {
		loc, _ := time.LoadLocation(name[0])
		return time.Now().In(loc).Unix()
	}
	return time.Now().In(DefaultLocation()).Unix()
}

// layout
func DayStartUnix(layout ...string) int64 {
	layoutS := LayoutYmd
	if layout != nil {
		layoutS = layout[0]
	}
	timeStr := time.Now().Format(layoutS)
	t, _ := time.ParseInLocation(layoutS, timeStr, DefaultLocation())
	unix := t.Unix()
	if unix > 0 {
		return unix
	}
	return 0
}

func DayEndUnix(layout ...string) int64 {
	unix := DayStartUnix(layout...)
	if unix > 0 {
		return unix + 86399
	}
	return 0
}

type timeFormat struct {
	layout   string
	unixTime int64
	timeStr  string
	loc      *time.Location
	err      error
}

func (m *timeFormat) GetErr() error {
	return m.err
}

func (m *timeFormat) SetLoc(name string) *timeFormat {
	loc, err := time.LoadLocation(name)
	m.loc = loc
	m.err = err
	return m
}

func (m *timeFormat) GetLoc() *time.Location {
	if m.loc == nil {
		m.loc = DefaultLocation()
	}
	return m.loc
}

func StringTime(timeStr string) *timeFormat {
	return &timeFormat{
		layout:  LayoutDefault,
		timeStr: timeStr,
	}
}

func Int64Time(unixTime int64) *timeFormat {
	return &timeFormat{
		layout:   LayoutDefault,
		unixTime: unixTime,
	}
}

func (m *timeFormat) String(layout ...string) string {
	if m.unixTime == 0 {
		if m.timeStr != "" {
			return m.timeStr
		}
		return ""
	}
	if layout != nil {
		m.layout = layout[0]
	}
	return time.Unix(m.unixTime, 0).Format(m.layout)
}

func (m *timeFormat) Int64(layout ...string) int64 {
	if m.timeStr == "" {
		if m.unixTime > 0 {
			return m.unixTime
		}
		return 0
	}
	if layout != nil {
		m.layout = layout[0]
	}
	stamp, _ := time.ParseInLocation(m.layout, m.timeStr, m.GetLoc())
	return stamp.Unix()
}
