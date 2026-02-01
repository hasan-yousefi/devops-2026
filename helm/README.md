# The Most Important Helm command i've already use it
# Prometheus Operator
```
# We enable remote write to other prometheus and also enable podmonitor selector & service monitor selector
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm upgrade prometheus prometheus-community/kube-prometheus-stack -f prometheus.yaml -n monitoring
```
# Tetragon
```
helm repo add cilium https://helm.cilium.io
helm install tetragon cilium/tetragon -n tetragon --create-namespace
```
# Chaos-mesh
```
helm repo add chaos-mesh https://charts.chaos-mesh.org
helm install chaos-mesh chaos-mesh/chaos-mesh -n=chaos-mesh --set chaosDaemon.runtime=containerd --set chaosDaemon.socketPath=/run/containerd/containerd.sock --version 2.8.0 --create-namespace
```
# K8gb
```
kubectl create ns k8gb
helm repo add k8gb https://www.k8gb.io
helm install -n k8gb k8gb k8gb/k8gb --version 0.17.0
```
# Gateway-api
```
kubectl apply --server-side -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.4.1/experimental-install.yaml
```
