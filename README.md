# bookstore-users-api

## Crear Go Mod
- go mod init github.com/burnout09/bookstore-users-api

## Add Dependencies
- go get github.com/VividCortex/mysqlerr
- go get github.com/gin-gonic/gin
- go get github.com/go-sql-driver/mysql
- go get -u go.uber.org/zap

## Creacion de MySQL Schema
- CREATE SCHEMA `users_db` DEFAULT CHAR SET utf8 COLLATE utf8_spanish2_ci;

## Variables de Entorno
- mysql_users_username=root;mysql_users_password=burnout;mysql_users_host=127.0.0.1:3306;mysql_users_schema=users_db

## Creacion de Table

````
create table users
(
id           bigint auto_increment
primary key,
first_name   varchar(45) null,
last_name    varchar(45) null,
email        varchar(45) not null,
date_created varchar(45) null,
constraint users_email_uindex
unique (email)
);
````

## Alter Table

````
alter table users
	add status varchar(45) not null;

alter table users
	add password varchar(32) not null;
````

## Run App
-  mysql_users_username=root mysql_users_password=burnout mysql_users_host=127.0.0.1:3306 mysql_users_schema=users_db go run main.go