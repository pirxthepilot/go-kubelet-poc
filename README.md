# Go kubelet client proof-of-concept

This program is meant to be run inside a Kubernetes node, and queries the PodList of the kubelet locally (default endpoint is `http://127.0.0.1:10255/pods`).

# Usage

```
$ ./go-kubelet-poc -help
Usage of ./go-kubelet-poc:
  -url string
    	Complete URL to pods endpoint (default "http://127.0.0.1:10255/pods")
```

# Sample Output

```
$ ./go-kubelet-poc | jq .
[
  {
    "hostip": "",
    "name": "etcd-minikube",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "",
    "status": "Pending"
  },
  {
    "hostip": "",
    "name": "kube-apiserver-minikube",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "",
    "status": "Pending"
  },
  {
    "hostip": "",
    "name": "kube-controller-manager-minikube",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "",
    "status": "Pending"
  },
  {
    "hostip": "",
    "name": "kube-scheduler-minikube",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "",
    "status": "Pending"
  },
  {
    "hostip": "10.0.2.15",
    "name": "kubernetes-dashboard-5498ccf677-98qq9",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "172.17.0.2",
    "status": "Running"
  },
  {
    "hostip": "",
    "name": "kube-addon-manager-minikube",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "",
    "status": "Pending"
  },
  {
    "hostip": "10.0.2.15",
    "name": "kube-proxy-qm5rw",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "10.0.2.15",
    "status": "Running"
  },
  {
    "hostip": "10.0.2.15",
    "name": "kube-dns-86f4d74b45-dslpq",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "172.17.0.3",
    "status": "Running"
  },
  {
    "hostip": "10.0.2.15",
    "name": "storage-provisioner",
    "namespace": "kube-system",
    "node": "minikube",
    "podip": "10.0.2.15",
    "status": "Running"
  }
]
```
