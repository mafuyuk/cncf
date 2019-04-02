# k8s-training

## minikube install
```bash
$ brew cask install virtualbox
$ brew cask install minikube
```

## kubectl install
```bash
$ brew install kubernetes-cli
```

## create cluster for minikube
```bash
$ minikube start --kubernetes-version=v1.11.3
$ kubectl get node
```

## stop cluster for minikube
```bash
$ minikube stop
```


## delete cluster for minikube
```bash
$ minikube delete
```

## addon
```bash
$ minikube addons list
$ minikube addons open dashboard
$ minikube addons enable dashboard
$ minikube addons disable dashboard
```
