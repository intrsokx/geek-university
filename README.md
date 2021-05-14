# geek-university
This is about the homework of the Geek University golang training camp


```
echo 'FROM mysql:5.7\nRUN mkdir data && echo "自定义MySQL镜像" > /data/info.txt' > Dockerfile
docker build -t mysql_init:v1 .

docker run -itd --name mysql_init -p 8888:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql_init:v1

docker exec -it container_id /bin/bash


//可以在dockerfile中做一些初始化MySQL的操作:    失败
mysql -h127.0.0.1 -P3306 -uroot -p123456 -e "create database test"
docker run -itd --name mysql_5.7 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
```
