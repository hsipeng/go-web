# go-web
> go web 框架 template


## 本地启动
1. 安装依赖

```bash
make install
```

2. 启动

```bash
make start
```

# docker-compose 部署启动

```bash
docker-compose up
```
如果更改了配置等需要重新编译，使用 `docker-compose build` or `docker-compose up --build`

提示: 

解决 Mysql 启动时会报链接错误
```bash
1. 进入 mysql ，执行命令
	# docker exec -it go-web_mysql57_1 bash
	# mysql_secure_installation
2.  新建一个admin账户，赋予权限。
    # docker exec -it go-web_mysql57_1 bash
    # mysql -u root -p
    # CREATE USER 'admin'@'%' IDENTIFIED BY '123456';
    # grant all on *.* to 'admin'@'%';

参考 https://github.com/docker-library/mysql/issues/275

```