# Spy Cat Agency Build & Run

This project includes Go backend service for the Spy Cat Agency, designed to manage and track spy cats.
It uses a PostgreSQL database for persistent data storage and exposes a Swagger API Documentation for easy testing.

*All of the commands should be performed from the root of the project.*

## Step 1: Create Docker Network:

```
docker network create spycats-net
```

##  Step 2: Build and Start PostgreSQL Container

```
docker-compose -f docker/docker_compose_files/postgres-docker-compose.yml up --build -d
```

##  Step 3: Build and Start Go Application Container

```
docker-compose -f docker/docker_compose_files/application-docker-compose.yml up --build -d
```

**Now you can test the application! The Swagger API specification is available at:**  

*http://localhost:8080/swagger/index.html*

## Stop & CleanUp

```
docker-compose -f docker/docker_compose_files/postgres-docker-compose.yml down -v
```

```
docker-compose -f docker/docker_compose_files/application-docker-compose.yml down
```

```
docker network rm spycats-net
```