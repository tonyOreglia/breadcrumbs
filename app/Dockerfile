FROM golang:1.13
# use go modules for package management
ENV GO111MODULE=on
# set the working directory
WORKDIR /app
# move the golang server files to the working directory
COPY . /app
# build the app into an executable
RUN go build ./cmd/breadcrumbs
# this tool waits for other docker containers to come ready
# depends on the environment variable WAIT_HOSTS -- see docker-compose file
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait
# expose the port this server runs on
EXPOSE 80
# wait for needed containers and then run the server
CMD /wait && go run ./cmd/breadcrumbs