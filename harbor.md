0.下載安裝包：
wget https://storage.googleapis.com/harbor-releases/release-1.7.0/harbor-online-installer-v1.7.1.tgz
1.解壓並修改配置：
tar -xzf harbor-online-installer-v1.7.1.tgz && cd harbor && vim harbor.cfg
############################
#配置访问的地址
hostname = 198.127.0.1
#使用http方式访问管理界面
ui_url_protocol = http
#配置admin的密码，默认是Harbor12345
harbor_admin_password = 12345
##jenkin docker login error 需要安裝##
sudo apt install gnupg2 pass
