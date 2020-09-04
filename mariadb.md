---master配置文件添加：
server-id=1 #配置服务器的服务号
log-bin=master #指定数据库操作的日志文件
replicate-do-db=pos #需要同步的数据库，多个就写多行
max_binlog_size=500M #设置日志文件最大值
binlog_cache_size= 128k #设置日志文件缓存大小
bind-address = 0.0.0.0  #允许远程访问
---slave配置文件添加:
server-id=2 #配置服务器的服务号
log-bin=master #指定数据库操作的日志文件
replicate-do-db=pos #需要同步的数据库，多个就写多行
max_binlog_size=500M #设置日志文件最大值
binlog_cache_size= 128k #设置日志文件缓存大小
bind-address = 0.0.0.0  #允许远程访问
read-only=1 #设置从服务器为只读模式
---创建同步账户
---master主机上：
grantreplication slave on *.* to‘slave‘@'192.168.159.%' identifiedby 'Admin@123';
flush privileges;
---slave主机上：
grantreplication slave on *.* to‘master’@'192.168.159.%' identifiedby 'Admin@123';
flush privileges;
---配置主从服务器
---master主机上：
changemaster to      master_host='192.168.159.51',master_user='master',master_password='Admin@123',master_log_file='user.000004',master_log_pos=0;
---slave主机上:
changemaster to master_host='192.168.159.50',master_user=’slave',master_password='Admin@123',master_log_file='user.000004',master_log_pos=0;
---启动主从服务器
reset master;
reset slave;
slave start;
---重启主从节点mariadb:
systemctl restart mariadb
