# How To Run it 
- download docker desktop or make sure you have docker and docker compose installed
- create .env file which have the following parameter
``` conf
BACKEND_ADRESS=http://Backend
BACKEND_PORT=5050
FRONTEND_ADRESS=http://Frontend
FRONTEND_PORT=5051
DB_ADRESS=DB
DB_PORT=3306
MYSQLROOTPW=root
MYSQLUSER=contentUsr
MYSQLPW=root
```
- run the docker-compose.yml
- the prod_docker-compose.yml dosn't contains pre generated data