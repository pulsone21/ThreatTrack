# How To Run it

- download docker desktop or make sure you have docker and docker compose installed
- create .env file which have the following parameter

```conf
FRONTEND_ADDR=http://Frontend
FRONTEND_PORT=5051
BACKEND_PORT=5051
BACKEND_ADDR=http://Backend
MYSQL_ADDR=DB
MYSQL_PORT=3306
MYSQLROOTPW=root
MYSQL_USER=contentUsr
MYSQL_PW=root
MYSQL_DBNAME=threattrack
```

- run the docker-compose.yml
- the prod_docker-compose.yml dosn't contains pre generated data
