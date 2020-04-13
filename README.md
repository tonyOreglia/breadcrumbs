# BreadCrumbs Server 
A RESTful HTTP server which allows users to generate and receive location based notes.

The server will validate and attempt to fix mobile numbers before storing. 
Upon storing a batch of numbers, the service will return some basic information about the numbers stored.

## The App 
This server is intended to support an application that Josh Kempy and I envisioned. The app would be an implementation of next generation social networking – dynamic geo based interaction.

The idea would allow users to generate and receive location based data. The user need not be connected to another user to recieve data -- only to be in the place the data was dropped. This data could be as simple as a note, image, or video. As a first step, I am planning to build an API to support such an app.

## Dependencies 
 - [Docker](https://www.docker.com/)
 - [Flyway](https://flywaydb.org/)
 - [Golang >1.11](https://github.com/golang/go/wiki/Modules)
 - [modd](https://github.com/cortesi/modd) 
 - [psql]

### Developing on Mac

#### Install Dependencies
  1. Install [Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
  1. Install modd 
  `$ brew install modd`
  1. Install [Go](https://golang.org/doc/install)

#### Initialize DB and Tables

Start Postgres, initialize DB and create tables
```
$ docker-compose up -d
```


Connect to the Postgres DB using psql
```
$ docker exec -it postgis /bin/bash -c "PGPASSWORD=<password> psql -d <db_name> -U <username> -h localhost"
```



