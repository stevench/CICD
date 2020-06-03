0.portainer添加endpoints需要:
---將endpoints的節點的/usr/lib/systemd/system/docker.service配置文件中添加ExecStart選項的值: -H tcp://0.0.0.0:2375
1.portainer給endpoints節點拉取鏡像需要在endpoints節點的/etc/docker/daemon.json中配置{ "insecure-registries":["192.168.0.52:80"] }
