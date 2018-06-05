package middleware

import (
	mon "git.finogeeks.club/monitor/go-client/monitor"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var monitor mon.Monitor

func init() {
	monitor = mon.GetInstance()
}

func FinoMonitor() gin.HandlerFunc {
	var summary mon.LabeledSummary = monitor.NewLabeledSummary(
		"http_request_duration_seconds",
		[]string{"method", "path", "code"},
		map[float64]float64{0.5: 0.01, 0.9: 0.01, 0.99: 0.001, 0.999: 0.0001},
	)

	var histogram mon.LabeledHistogram = monitor.NewLabeledHistogram(
		"http_request_buckets_seconds",
		[]string{"method", "path", "code"},
		[]float64{0.05, 0.1, 0.5, 1, 2, 5},
	)

	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		//end := time.Now()
		//latency := end.Sub(start)
		duration := float64(time.Since(start)) / float64(time.Second)

		//fmt.Printf("duration |%.3f|%.3g|\n", duration, duration)
		//fmt.Printf("latency %#v    %13v \n", latency, latency)

		method := c.Request.Method
		path := c.Request.URL.Path
		code := strconv.Itoa(c.Writer.Status())

		summary.WithLabelValues(method, path, code).Observe(duration)
		histogram.WithLabelValues(method, path, code).Observe(duration)
	}
}
