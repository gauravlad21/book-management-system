# Book Management System

git clone https://github.com/gauravlad21/book-management-system.git
cd book-management-system
docker build -f build/Dockerfile.yaml -t tag .
cd build
docker-compose up

use swagger UI or postman collection to test
swagger endpoint: {{BaseUrl}}/swagger/index.html
live on: http://13.48.212.214:5002/swagger/index.html#/Books/get_events