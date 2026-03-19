package health

import (
	"runtime"
	"time"
)

// Часовой пояс для полей времени в API: UTC+3 (Москва).
var moscow = func() *time.Location {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return time.FixedZone("MSK", 3*3600)
	}
	return loc
}()

// Response — ответ API /api/health.
type Response struct {
	Status       string    `json:"status"`
	Timestamp    time.Time `json:"timestamp"`
	Uptime       string    `json:"uptime"`
	RequestCount int64     `json:"request_count"`
	GoVersion    string    `json:"go_version"`
}

// BuildResponse создаёт ответ для health endpoint.
func BuildResponse(uptime time.Duration, requestCount int64) Response {
	return Response{
		Status:       "ok",
		Timestamp:    time.Now().In(moscow),
		Uptime:       FormatUptime(uptime),
		RequestCount: requestCount,
		GoVersion:    runtime.Version(),
	}
}

// FormatUptime форматирует длительность в строку uptime.
func FormatUptime(d time.Duration) string {
	return d.Round(time.Second).String()
}
