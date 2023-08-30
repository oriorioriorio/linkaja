# linkaja

Ini adalah aplikasi sederhana untuk melihat saldo dan mengirim uang.

TECH STACK
- GO
- Mysql
- Docker
 
 
langkah untuk menjalankan aplikasi:
- import .env dari .env.example (abaikan jika sudah ada)

- Manual Setting
1. sediakan mysql, go
2. set config untuk connection db
3. run query pada mysql/migrations/migration.sql
4. jalankan aplikasi dengan (make run)

- Docker-Compose
1. sediakan docker
2. make run-docker

API Collection
- import postman collection