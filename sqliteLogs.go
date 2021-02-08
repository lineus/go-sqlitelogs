package sqlitelogs

import (
	"database/sql"
	"time"
)

// SqliteLog - shape of a log in sqlite
type SqliteLog struct {
	ID     int
	Epoch  float64
	Action string
	Result string
}

// SqliteLogger - abstraction for access to sqlite logs
type SqliteLogger interface {
	SaveLog(string, string) (sql.Result, error)
	GetAllLogs() []SqliteLog
	GetLogsBetween(time.Time, time.Time) []SqliteLog
	Alive() bool
}
