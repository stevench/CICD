0.拉取portainer:
---docker pull portainer/portainer
1.启动portainer:
---docker run -d --name portainerUI -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock portainer/portainer
2.portainer添加endpoints需要:
---將endpoints的節點的/usr/lib/systemd/system/docker.service配置文件中添加ExecStart選項的值: -H tcp://0.0.0.0:2375
3.portainer給endpoints節點拉取鏡像需要在endpoints節點的/etc/docker/daemon.json中配置{ "insecure-registries":["192.168.0.52:80"] }
4.部署节点和portainer进行tsl证书配置:
---4.1使用ca.sh在部署节点生成证书;
---4.2配置docker的2375端口tsl证书；
    vim  /lib/systemd/system/docker.service
    ExecStart=/usr/bin/dockerd -H unix:///var/run/docker.sock -D -H tcp://0.0.0.0:2375 --tlsverify --tlscacert=/root/.docker/ca.pem -- tlscert=/root/.docker/server-cert.pem --tlskey=/root/.docker/server-key.pem
---4.3将cert.pem和key.pem配置到portainer的endpoint中；
---4.4测试：
    curl -k https://docker服务器IP:2375/info --cert ./cert.pem --key ./key.pem
