package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metricsPath     = getEnv("METRICS_PATH", "/metrics") // Путь для получения метрик
	listenAddress   = getEnv("LISTEN_ADDRESS", ":9191")  // Порт для получения метрик
	rancherURL      = os.Getenv("CATTLE_URL")            // URL rancher сервера. Пример: https://rancher.example.com
	accessKey       = os.Getenv("CATTLE_ACCESS_KEY")     // Access Key для Rancher API
	secretKey       = os.Getenv("CATTLE_SECRET_KEY")     // Secret Key для Rancher API
	resourceLimit   = getEnv("API_LIMIT", "100")         // Лимит ресурсов Rancher API (по-умолчанию: 100)
	namespaceFilter = getEnv("NAMESPACES", "all")        // Фильтр по namespace (по-умолчанию все. Пример: "kube-system,default,test")
)

// getEnv - получаем переменную оркужения, если её нет - возвращаем значение по-умолчанию
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	if rancherURL == "" {
		log.Fatal("CATTLE_URL must be set and non-empty")
	}
	if accessKey == "" {
		log.Fatal("CATTLE_ACCESS_KEY must be set and non-empty")
	}
	if secretKey == "" {
		log.Fatal("CATTLE_SECRET_KEY must be set and non-empty")
	}

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)

	exp := newExporter()
	reg.MustRegister(exp)

	http.Handle(metricsPath, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
