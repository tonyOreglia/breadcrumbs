## BreadCrumbs Application 

"It's like exploring history for a place."
- Jeremy Rovelli, product advocate 1

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

#### Environment Variables
The postgres server depends on environment variables for username, password, etc. 

Store these in a local .env file. 

For example, the file contents may be: 
```
".env" [noeol] 6L, 121C                                                                          1,1           All
DB_NAME=<db name>
DB_USER=<db user> 
DB_PW=<password>
DB_PORT=5432
DB_HOST=postgis
MAX_DB_CONNECTIONS=2
``` 

#### Run Development Mode

```
$ make run-dev
```

The server will be available on localhost port 8081.

#### Run with Production Build
```
$ make run-prod
```


#### Stop the Service
```
$ make stop
```

#### Connect to postgres server with psql command line tool
```
$ psql -h 192.168.1.203  postgis -d breadcrumbs -U breadcrumbs_user -W
```

Get the password from local .env file

#### Migrating Data to new DB 
Note that this is data only; the tables are setup automatically by the flyway migration script

##### Dump data
```
$ sudo docker exec postgis pg_dump --data-only -d breadcrumbs -U breadcrumbs_user > breadcrumbs_dump.sql
```

##### Move the date to new place 
```
$ scp  tony@192.168.1.203:/home/tony/dev/breadcrumbs/breadcrumbs_dump.sql ~/dev/breadcrumbs/
```

##### Load the data
```
$ psql -h localhost -d breadcrumbs -U breadcrumbs_user < breadcrumbs_dump.sql
```


#### Generate Test Data
See `example-data-set/README.md`

# API 

If you use Postman, import the collection from `app/documentation/breadcrumbs.postman_collection.json`

## POST Create Note
http://localhost:8081/note
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
curl --location --request POST 'http://localhost:8081/note' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "Hello World!!",
    "longitude": 9.156324,
    "latitude": 38.7153423
}'
```

## POST Get Notes
http://localhost:8081/getNotes

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
curl --location --request POST 'http://localhost:8081/getNotes' \
--header 'Content-Type: application/json' \
--data-raw '{
    "latitude": 38.7153,
    "longitude": 9.1566,
    "radius_in_meters": 1000
}'
```

# GUI Demonstration 
## Options 
- [react google maps](https://github.com/tomchentw/react-google-maps)
- [mapBox](https://www.mapbox.com/pricing/)

I'll try google maps mainly because it is less configurable i.e. works more easily "out of the box". I don't want to do anything fancy, just post and display points so I don't want the extra flexibility of mapBox. Both free plans are more than sufficient for my needs. 

## Resources 
- overview of react google maps and react with map box: https://www.telerik.com/blogs/maps-in-react


## What features would the demo include? 
- could enable anyone to drop a pin with a note as a tonycodes.com demo universal user -- and likewise people may see all notes dropped


## What will this requires?
- Self hosting
  - hosting breadcrumbs as a backend on the server
  - hosting the data generated on my small ssd
  - allowing breadcrumbs requests / responses through the router -- can put it behind nginx similarly to the chess engine

- host on cloud provider 
  - might have to pay? 
  - less fun 


## Steps 
1. Generate a google api key https://console.cloud.google.com/google/maps-apis/credentials
2. Enable your API Key to access maps https://developers.google.com/maps/gmp-get-started#enable-api-sdk
You are looking for "Maps Javascript API" 



# To Do
- [ ] attach a user to the breadcrumbs
- [ ] implement rate limiting or some kind of authentication mechanism ... 
- [ ] add bc's to deployment pipeline on codeship
