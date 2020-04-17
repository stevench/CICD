#### 系统 ubuntu18.04 

#### docker-ce 1903  安装过程略

+ 配置 docker

```
$ cat /etc/docker/daemon.json 
{
        "exec-opts": ["native.cgroupdriver=systemd"],
        "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn/"],
        "log-driver": "json-file",
        "log-opts": {
                "max-size": "100m"
        }
}

systemctl daemon-reload
systemctl restart docker

```

+ 关闭 swap  

> swapoff -a
> 注释 /etc/fstab swap 的行
> 重启


+ 阿里云 k8s 镜像地址 https://opsx.alibaba.com/mirror

``` 
apt-get update && apt-get install -y apt-transport-https
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add - 

# 添加镜像源
cat /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main

# 安装
apt-get update
apt-get install -y kubelet kubeadm kubectl
```

+ 开始部署(仅master节点)
```
kubeadm init \
    --apiserver-advertise-address=192.168.6.142 \    
    --image-repository registry.aliyuncs.com/google_containers \
    --pod-network-cidr=10.244.0.0/16 \
    --kubernetes-version v1.16.3 \
    --ignore-preflight-errors=NumCPU 
```  

> 说明 ：
> --apiserver-advertise-address=192.168.6.142   本地地址
> --image-repository registry.aliyuncs.com/google_containers 使用阿里云镜像，不然很慢
> --pod-network-cidr=10.244.0.0/16 指定 pob 网络， 下面安装网络时需要
> --kubernetes-version v1.16.3  指定版本，可选，不加就默认最新版，
> --ignore-preflight-errors=NumCPU  无视 cpu 检查，cpu 至少要求 2 个，我虚拟机只配了一个， 不加报错，加了警告

[问题解决]
kubeadm join 超时报错 error execution phase kubelet-start: error uploading crisocket: timed out waiting for the condition
解决:
swapoff -a
kubeadm reset
systemctl daemon-reload
systemctl restart kubelet
iptables -F && iptables -t nat -F && iptables -t mangle -F && iptables -X  


+ 安装完成提示， 信息要记下，
```
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.6.142:6443 --token 8cikod.edxfpc3auefetiag \
    --discovery-token-ca-cert-hash sha256:f9fe7126906bbb26ed700f749389289784cb83ea78cf0c4b0bcd7bcae0c04c85 


```




+ 安装网络，不安装网络什么都运行不起来， 我这里选  calico ，还有很多可以选的
参考 ： https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/

```
  export KUBECONFIG=/etc/kubernetes/admin.conf
  kubectl apply -f https://docs.projectcalico.org/v3.8/manifests/calico.yaml
   

```

+ 允许master 运行 pob ， 

```
  kubectl taint nodes --all node-role.kubernetes.io/master-
```


+  非root要想使用。。。。

```
    sudo cp /etc/kubernetes/admin.conf $HOME/
    sudo chown $(id -u):$(id -g) $HOME/admin.conf
    export KUBECONFIG=$HOME/admin.conf
```

+  另一个节点加入集群：
> kubeadm join 192.168.6.142:6443 --token 8cikod.edxfpc3auefetiag \
    --discovery-token-ca-cert-hash sha256:f9fe7126906bbb26ed700f749389289784cb83ea78cf0c4b0bcd7bcae0c04c85 


+ dashboard 
> 1.下载阿里云的镜像：docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:v1.10.0
> 2.重新打tag: docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:v1.10.0 k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.0
> 3.通过yml文件拉起pod: kubectl apply -f http://mirror.faasx.com/kubernetes/dashboard/master/src/deploy/recommended/kubernetes-dashboard.yaml

+ 获取token
> kubectl -n kube-system describe $(kubectl -n kube-system get secret -n kube-system -o name | grep namespace) | grep token




+ kubernetes dashboard安装
下载 Dashboard yaml 文件
wget http://pencil-file.oss-cn-hangzhou.aliyuncs.com/blog/kubernetes-dashboard.yaml

打开下载的文件添加一项：type: NodePort，暴露出去 Dashboard 端口，方便外部访问
