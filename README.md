# Chat
1.) To create MySQL Database on Docker: Link-Create Docker Container for sql database Working: https://www.youtube.com/watch?v=X8W5Xq9e2Os 

a.) Download Docker Image docker pull mysql/mysql-server:latest

b.) Run docker container
    docker run -p 3307:3306 --name chatAppLanguageTranslation -e MYSQL_ROOT_PASSWORD=password -d mysql

c.) Execute container
    docker exec -it chatAppLanguageTranslation /bin/bash

d.) Login to mysql
    mysql -uroot -p -A
    enter password:password

e.) To create a table in sql database in container
    1.) CREATE DATABASE IF NOT EXISTS userAccountDB;
    2.) USE userAccountDB;
    3.) Press enter
    4.) CREATE TABLE Users (UserName VARCHAR(30) NOT NULL PRIMARY KEY, Password VARCHAR(256), FirstName VARCHAR(30), LastName VARCHAR(30), Language VARCHAR(30));
    5.) Press enter

f.) To check content of table:
    USE userAccountDB; SELECT * FROM Users;

g.) To see database:
    SHOW DATABASES;

h.) To see what is in userAccountDB User table:
    SHOW COLUMNS FROM userAccountDB.Users;

====================================================================================================================================================================
Face Recognition:
1.) Run Machine Box Docker container:
    docker run -d -p 8080:8080 -e "MB_KEY=$MB_KEY" machinebox/facebox

2.) Install OpenCV
    a.) brew install opencv3
    b.) go get gocv.io/x/gocv



====================================================================================================================================================================
TO RUN PROGRAM:

a.) Open folder C:\Users\user1\Desktop\GitHub\Chat\realtime chat\src
    - go run main.go
    
b.) Open folder C:\Users\user1\Desktop\GitHub\Chat\server
    - go run .
