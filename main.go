package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func getEnvDefault(name string, defaultVal string) string {
	envValue, ok := os.LookupEnv(name)
	if ok {
		return envValue
	}
	return defaultVal
}

func main() {

	listendAddr := getEnvDefault("HTTP_LISTENADDR", ":9112")

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	tfeRuns := newTfeCollector()
	prometheus.MustRegister(tfeRuns)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	log.Info("Now listening on ", listendAddr)
	log.Fatal(http.ListenAndServe(listendAddr, nil))

}
