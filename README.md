# to_start


## run:
    cd ./cmd
## create file .env
    DB_USERNAME=YOUR_username
    DB_PASSWORD=YOUR_password
    DB_DATABASE=YOUR_dbname
    DATABASE_HOST=YOUR_host or 127.0.0.1
    DATABASE_PORT=YOUR_port or 5432
    SECRET_KEY=your-256-bit-secret
## load libs
    go mod tidy

## then run:
    go build .\main.go
    go run .\main.go 

## authorization order
1. sign up


        curl --location 'http://127.0.0.1:1323/api/public/signup' \
        --form 'name="Kio"' \
        --form 'email="go55@go.ru"' \
        --form 'age="33"' \
        --form 'password="123456"'
2. get jwt token


        curl --location 'http://127.0.0.1:1323/api/public/login' \
        --form 'email="go55@go.ru"' \
        --form 'password="123456"'

3. work with api

    
        curl --location 'http://127.0.0.1:1323/api/protected/users' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Ii
        QyYSQxNCRHVzU0ZjJHaGNtd2ozRmNoUi40N2NlUzJLRENSOGNoZjEwS2NBVEdOQ3dzTmdCRWlFRzVMeSIs
        ImV4cCI6MTY5NjI1MjQ1NSwiaXNzIjoiQXV0aFNlcnZpY2UifQ.USpu4MywQ6Qo8444h2STXoUuIK8Vu9wv-OnhmfEKKAc'


.

        curl --location --request DELETE 'http://127.0.0.1:1323/api/protected/users/11' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
        eyJFbWFpbCI6IiQyYSQxNCRHVzU0ZjJHaGNtd2ozRmNoUi40N2NlUzJLRENSOGNoZjEwS2NBVEdOQ3dzTmdCRW
        lFRzVMeSIsImV4cCI6MTY5NjI1MjQ1NSwiaXNzIjoiQXV0aFNlcnZpY2UifQ.USpu4MywQ6Qo8444h2STXoUuIK
        8Vu9wv-OnhmfEKKAc'