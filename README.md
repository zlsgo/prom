# prom

Prometheus metrics exporter


**prometheus.yml**

```yml
scrape_configs:
  - job_name: 'api'
    scrape_interval: 20s
    metrics_path: /prom_metrics
    static_configs:
      - targets: ['localhost:3788']
```

**Import via panel json**

export `dashboard.json`