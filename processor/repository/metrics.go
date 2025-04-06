package repository

import (
	"time"
)

type Metrics interface {
	Register()
	RecordTaskDuration(language string, duration time.Duration)
	RecordLanguageUsage(language string)
}
