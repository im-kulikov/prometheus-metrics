# Prometheus + Grafana metrics

## How to use

```go
// ...

// PROMETHEUS=127.0.0.1:18081

prometheus.Serve(nil)

// ...
```

See examples for setup [Prometheus and Grafana](./examples)

> INFO: replace `PROMETHEUS_METRICS_BIN` with your `PROMETHEUS` env