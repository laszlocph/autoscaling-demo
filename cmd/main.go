package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "autoscaling_demo_requests_total",
		Help: "The total number of api requests",
	})

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {
	http.HandleFunc("/api", api)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8888", nil)
}

func api(w http.ResponseWriter, req *http.Request) {
	requestCounter.Inc()

	min := 100
	max := 200
	time.Sleep(time.Millisecond * time.Duration((r.Intn(max-min+1) + min)))

	w.Write([]byte("{\"response\": \"ok\"}"))
}
