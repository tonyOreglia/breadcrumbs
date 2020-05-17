## BreadCrumbs Application 
This server is intended to support an application that [Josh Kemp](https://www.linkedin.com/in/josh-kemp-440a3b83/) (aka "kempatron" aka "kempy wants!" aka "kempy") and I envisioned called BreadCrumbs. BreadCrumbs is an implementation of next generation social networking which revolves around dynamic geography based interaction.

BreadCrumbs allow users to create and find location based data. This data could be anything; from a message, image, or video to a virtual restaurant rating, augmented reality footage, or advertisement. The user need not be connected to any other user in the traditional sense of social networking. No friendships or network needed. The only connection needed is location; irrespective of time. To receive data, the user only needs to be in the same place the data was dropped. Thus the user stumbles upon past messages as they explore the real world. 

# BreadCrumbs Server Tech Stack
A RESTful HTTP server written in Golang. This API allows users to generate and retrieve location based messages. A PostgreSQL database is run in a Docker container. [PostGIS](https://postgis.net/) is used for storage and efficient retrieval of spatial data. 

* [Go Programming Language](https://golang.org/): Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Docker](https://www.docker.com/): Docker uses OS-level virtualization to deliver software in packages called containers.
* [PostgreSQL](https://www.postgresql.org/): PostgreSQL is a powerful, open source object-relational database system.
* [PostGIS](https://postgis.net/): A spatial database extender for PostgreSQL
* [Flyway](https://flywaydb.org/): Flyway is an open-source database migration tool.
* [modd](https://github.com/cortesi/modd): A flexible developer tool that runs processes and responds to filesystem changes


## Dependencies 
 - [Docker](https://www.docker.com/)

### Getting Started Developing Locally

#### Install Dependencies

* Install [Docker Desktop](https://www.docker.com/products/docker-desktop)

#### Run Development Mode

```
$ make run-dev
```

The server will be available on localhost port 80.

#### Run with Production Build
```
$ make run-prod
```


#### Stop the Service
```
$ make stop
```

#### Generate Test Data
This repo contains a data set from [kaggle.com](https://www.kaggle.com/datasets). The data set contains the number of covid-19 deaths recorded in cities around the world as of May 8, 2020. This data set can be stored in the Breadcrumbs server for testing purposes. 

Requires
* python3
* pipenv

```
$ cd example-data-set
$ pipenv shell
$ python store-data-as-breadcrumbs.py
```

# API 

If you use Postman, import the collection from `app/documentation/breadcrumbs.postman_collection.json`

## POST Create Note
http://localhost:80/note
### HEADERS
`Content-Typeapplication/json`

### BODY
```
{
    "message": String,
    "longitude": Number,
    "latitude": Number
}
```

### Example Request - Create Note
```
curl --location --request POST 'http://localhost:80/note' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "Hello World!!",
    "longitude": 9.156324,
    "latitude": 38.7153423
}'
```

## POST Get Notes
http://localhost:80/getNotes

### HEADERS 
`Content-Typeapplication/json`

### BODY 
```
{
    "latitude": Number,
    "longitude": Number,
    "radius_in_meters": Integer
}
```

### Example Request - Get Notes
```
curl --location --request POST 'http://localhost:80/getNotes' \
--header 'Content-Type: application/json' \
--data-raw '{
    "latitude": 38.7153,
    "longitude": 9.1566,
    "radius_in_meters": 1000
}'
```
