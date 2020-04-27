FROM golang:1.13
# use go modules for package management
ENV GO111MODULE=on
# set the working directory
WORKDIR /app
# If these haven't changed then they'll be cached
# important as downloading these files takes time
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
# move the golang server files to the working directory
COPY . /app
# build the app into an executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/breadcrumbs
# this tool waits for other docker containers to come ready
# depends on the environment variable WAIT_HOSTS -- see docker-compose file
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait
# expose the port this server runs on
EXPOSE 80
# wait for needed containers and then run the server
CMD /wait && go run ./cmd/breadcrumbs