# k8s

## Install docker

```bash
sudo pacman -S docker
sudo systemctl enable --now docker.service
```

验证是否安装成功

```bash 
docker info
docker run -it --rm archlinux bash -c "echo hello world"
```

登录

```bash
docker login -u <username>
```

设置镜像

```bash
vim /etc/docker/daemon.json
```

```json
{
  "registry-mirrors" : [
    "http://docker.mirrors.ustc.edu.cn"
  ],
  "insecure-registries" : [
    "docker.mirrors.ustc.edu.cn"
  ],
  "debug" : true,
  "experimental" : true
}
```

```bash
systemctl daemon-reload
systemctl restart docker
```

#### 在 Dockfile 中修改 alpine 源 

修改项目中的`Dockerfile`，在`Dockerfile`中的所有 `FROM ...alpine...` 语句后面添加一句：

```bash
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
```

其他修改 alpine 源的方法请参阅 [Docker和alpine镜像内国内源配置](https://blog.csdn.net/shuizhongmose/article/details/108992380)

## Install kubernates

```bash
sudo pacman -S etcd kubernetes-control-plane kubernetes-node
sudo pacman -S kubeadm kubelet
sudo pacman -S kubectl
```

## Install minikube

```bash
sudo pacman -S minikube
```

#### 配置代理

```bash
export HTTP_PROXY=http://<proxy hostname:port>
export HTTPS_PROXY=https://<proxy hostname:port>
export NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,192.168.99.0/24,192.168.39.0/24
```

#### 启动集群

```bash
minikube start
minikube dashboard
```

#### 使用 echo-server 创建 Deployment

```bash
kubectl create deployment hello-minikube --image=registry.aliyuncs.com/google_containers/echoserver:1.10

kubectl expose deployment hello-minikube --type=NodePort --port=8080
kubectl get pod
minikube service hello-minikube --url
```

#### 服务扩容

```bash
kubectl scale deployment hello-minikube --replicas=3
kubectl get pods
```

#### 关闭 minikube

```bash
minikube stop
```

