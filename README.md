##Build Image

```bash
docker build -t mysql_gorm_test .
```

##Run Container
```bash
docker run --name mysql_gorm_test -e MYSQL_ROOT_PASSWORD=password -d -p 3308:3306  mysql_gorm_test
```

***
需要等container初始完才可正常連線

##Run Test
```bash
go run main.go
```
