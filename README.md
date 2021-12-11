# go-prom
Go server exposing a prometheus endpoint

# Pre-requisites
- Running kubernetes cluster with prometheus agent installed as outlined in Prometheus installation section
- Deploy the go-prom application to the cluster. Metrics are exposed at `:2112/metrics`
- Edit the configmap called `prometheus-server` to add a job which would scrape the metrics from this. The job config is specified below -
    ```yaml
    - job_name: 'go-prom'
        static_configs:
        - targets: ['go-prom:2112]
    ```
- This would create a new target that can be viewed in the prometheus dashboard

# Prometheus installation steps
- `helm repo add prometheus-community https://prometheus-community.github.io/helm-charts`
- `helm repo update`
- `helm install prometheus prometheus-community/prometheus`

# Local setup
- Launch a `kind` cluster - 
    ```
    kind create cluster
    ```
- Build the docker image locally and load it into the cluster. These steps would make the image available locally within the cluster
    ```
    docker build --tag go-prom:1.0.0 .
    kind load docker-image go-prom:1.0.0
    ```
- Create the service and deployment for `go-prom` within the cluster
    ```
    kubectl apply -f go-prom-app.yaml
    ```
- Port forward `prometheus-server` port 80 to view the prometheus dashboard locally
    ```
    kubectl port-forward service/prometheus-server 80
    ```
    