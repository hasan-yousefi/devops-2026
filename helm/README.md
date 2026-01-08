# The Most Important Helm command i've already use it
# Prometheus Operator
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm upgrade prometheus prometheus-community/kube-prometheus-stack -f prometheus.yaml -n monitoring
