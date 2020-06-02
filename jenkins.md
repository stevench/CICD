###ubuntu install jenkins###
sudo apt-get install openjdk-8-jdk
wget -q -O - https://pkg.jenkins.io/debian/jenkins-ci.org.key | sudo apt-key add -
echo deb http://pkg.jenkins.io/debian-stable binary/ | sudo tee /etc/apt/sources.list.d/jenkins.list
sudo apt-get update
sudo apt-get install jenkins
sudo systemctl start jenkins
sudo systemctl status jenkins
sudo ufw allow 8080
sudo ufw status
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
#0.讓jenkins用戶有docker執行權限:
sudo gpasswd -a jenkins docker     #将登陆用户加入到docker用户组中
sudo newgrp docker     #更新用户组
sudo service jenkins restart #重啟jenkins
#1.安裝推薦插件；
#2.安裝中文插件；
#3.重啟顯示中文；
