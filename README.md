# bookstore-users-api

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