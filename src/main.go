package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_errors_total",
			Help: "Total number of application errors",
		},
		[]string{"error_type"},
	)
)

func init() {
	prometheus.MustRegister(errorCounter)
}

func logError(message string) {
	log.Printf("ERROR: %s | Timestamp: %s | Pod: %s | Container: %s",
		message, time.Now().Format(time.RFC3339), os.Getenv("POD_NAME"), os.Getenv("CONTAINER_NAME"))
	errorCounter.WithLabelValues("runtime_error").Inc()
}

func handler(w http.ResponseWriter, r *http.Request) {
	logError("Simulated application error occurred")
	fmt.Fprintln(w, "An error occurred!")
}

func main() {
	// Testing analyzeError function
	testLog := "ERROR: database connection failed"
	log.Println("Inference:", analyzeError(testLog))

	// Starting HTTP Server
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)

	log.Println("Starting Go SRE Agent on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
