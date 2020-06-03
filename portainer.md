0.拉取portainer:
---docker pull portainer/portainer
1.启动portainer:
---docker run -d --name portainerUI -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock portainer/portainer
2.portainer添加endpoints需要:
---將endpoints的節點的/usr/lib/systemd/system/docker.service配置文件中添加ExecStart選項的值: -H tcp://0.0.0.0:2375
3.portainer給endpoints節點拉取鏡像需要在endpoints節點的/etc/docker/daemon.json中配置{ "insecure-registries":["192.168.0.52:80"] }
