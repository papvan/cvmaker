package repeatable

import (
	"strings"
	"time"
)

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--

			continue
		}

		return nil
	}

	return
}

func FormatSQLToLine(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
