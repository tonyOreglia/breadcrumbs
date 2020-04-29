## BreadCrumbs Application 
This server is intended to support an application that Josh Kempy and I envisioned called BreadCrumbs. BreadCrumbs is an example of next generation social networking which centers around dynamic geography based interaction.

BreadCrumbs allow users to create and find location based data. This data could be anything; from a message, image, or video to a virtual restaurant rating, augmented reality footage, or advertisement. The user need not be connected to any other user in the traditional sense of social networking. No friendships or network needed. The only connection needed is location; irrespective of time. To receive data, the user only needs to be in the same place the data was dropped. Thus the user stumbles upon past messages as they explore the real world. 

# BreadCrumbs Server Tech Stack
A RESTful HTTP server builtt in Golang which allows users to generate and receive location based messages. A PostgreSQL database is run in a Docker container. [PostGIS](https://postgis.net/) is used for storage and efficient retrieval of spatial data. 

* [Go Programming Language](https://golang.org/): Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Docker](https://www.docker.com/): Docker uses OS-level virtualization to deliver software in packages called containers.
* [PostgreSQL](https://www.postgresql.org/): PostgreSQL is a powerful, open source object-relational database system.
* [PostGIS](https://postgis.net/): A spatial database extender for PostgreSQL
* [Flyway](https://flywaydb.org/): Flyway is an open-source database migration tool.
* [modd](https://github.com/cortesi/modd): A flexible developer tool that runs processes and responds to filesystem changes


## Dependencies 
 - [Docker](https://www.docker.com/)
 - [Golang >1.11](https://github.com/golang/go/wiki/Modules)

## Dev Dependencies
 - [modd](https://github.com/cortesi/modd) 

### Developing on Mac

#### Install Dependencies
  1. Install [Docker for Mac](https://docs.docker.com/docker-for-mac/install/)

#### Run Locally in Development Mode

```
$ make run-dev
```

#### Run with Production Build

```
$ make rur-prod
```


#### Stop the Servce

```
$ make stop
```

