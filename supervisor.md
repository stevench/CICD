0.安装supervisor:
sudo apt-get install supervisor
1.添加到系统用户组:
sudo addgroup --system supervisor
2.重新加载:
sudo supervisorctl reload
3.检查状态:
sudo supervisorctl status
