# prometheus-simple-rancher-exporter

Экспортер позволяет получать метрики о статусе кластеров (ноды, поды, контейнеры) из Rancher API в формате Prometheus

## Описание

Приложение может быть запущено с помощью docker образа "dru90i/prometheus-simple-rancher-exporter".

Приложению требуется как минимум URL-адрес Rancher API, RANCHER_ACCESS_KEY и RANCHER_SECRET_KEY

**Обязательно**
* `CATTLE_URL`          // URL rancher сервера. Пример: https://rancher.example.com
* `CATTLE_ACCESS_KEY`   // Access Key для Rancher API
* `CATTLE_SECRET_KEY`   // Secret Key для Rancher API

**Опционально**
* `METRICS_PATH`        // Путь для получения метрик (по-умолчанию: /metrics)
* `LISTEN_ADDRESS`      // Порт для получения метрик (по-умолчанию: 9191)
* `NAMESPACES`          // Фильтр по namespace (по-умолчанию все. Пример: "kube-system,default,test")
* `API_LIMIT`           // Лимит ресурсов Rancher API (по-умолчанию: 100)

## Установка и запуск

Запуск с помощью образа из docker hub
```
docker run -d -e CATTLE_ACCESS_KEY="XXXXXXXX" -e CATTLE_SECRET_KEY="XXXXXXX" -e CATTLE_URL="XXXXXXX" -p 9191:9191 dru90i/prometheus-simple-rancher-exporter
```

Сборка docker образа:
```
docker build -t <image-name> .
docker run -d -e CATTLE_ACCESS_KEY="XXXXXXXX" -e CATTLE_SECRET_KEY="XXXXXXX" -e CATTLE_URL="XXXXXXX" -p 9191:9191 <image-name>
```

## Метрики

Метрики будут доступны через порт 9191 по-умолчанию, или вы можете передать переменную окружения ```LISTEN_ADDRESS``` для переопределения.
Пример метрик, которые вы должны увидеть, можно найти на METRICS.md.

**Интерпретация значений метрик:**

rancher_cluster - возвращает id статуса (порядковый номер в списке начиная с 0):
{"active", "cordoned", "degraded", "disconnected", "drained", "draining", "healthy", "initializing", "locked", "purged", "purging", "reconnecting", "reinitializing", "removed", "running", "unavailable", "unhealthy", "upgraded", "upgrading"}

rancher_cluster_node - возвращает id статуса (порядковый номер в списке начиная с 0):
{"active", "cordoned", "drained", "draining", "provisioning", "registering", "unavailable"}

rancher_cluster_pod - возвращает id статуса (порядковый номер в списке начиная с 0):
{"Pending", "Running", "Succeeded", "Failed", "Unknown", "Unavailable"}

rancher_cluster_pod_container - возвращает id статуса (порядковый номер в списке начиная с 0):
{"waiting", "running", "terminated"}
