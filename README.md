# Go Database Connection
## How to Install Project
### เปิด Docker
สำหรับ MAC OS ให้กดเปิด Application Docker ก่อน

### Run with docker image
```bash
docker pull mysql:5.7
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=sckshuhari -d mysql:5.7
```

### คำสั่งสำหรับดู Image ใน docker
```
docker image ls
```
หรือ
```
docker images
```

### คำสั่งสำหรับดูว่า Container กำลังรันอยู่หรือเปล่าใน Docker
```
docker ps
```

#### คำสั่งสำหรับดู Container ที่มีอยู่ทั้งหมดใน Docker
```
docker ps -a
```

### คำสั่งสำหรับหยุดการทำงานของ Container ใน Docker
```
docker stop <ชื่อของ container หรือ container id>
```

### คำสั่งสำหรับลบ container ที่หยุดการทำงานแล้ว
```
docker rm <ชื่อของ container หรือ container id>
```

### Install MySql Workbench
```bash
https://dev.mysql.com/downloads/workbench/?utm_source=tuicool
```
