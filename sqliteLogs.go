package sqlitelogs

import (
	"time"
)

// SqliteLog - shape of a log in sqlite
type SqliteLog struct {
	id     int
	epoch  float64
	action string
	result string
}

// SqliteLogger - abstraction for access to sqlite logs
type SqliteLogger interface {
	SaveLog(string, string) error
	GetAllLogs() []SqliteLog
	GetLogsBetween(time.Time, time.Time) []SqliteLog
	Alive() bool
}
