# build stage
FROM golang:1.13 as builder

ARG environment
RUN echo "Running $environment build"

# use go modules for package management
ENV GO111MODULE=on
# set the working directory
WORKDIR /application
# If these haven't changed then they'll be cached
# important as downloading these files takes time
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./app ./app
# move the golang server files to the working directory
# COPY . /app
# build the app into an executable
# result is "breadcrumbs" executable in the current folder

# this tool waits for other docker containers to come ready	
# depends on the environment variable WAIT_HOSTS -- see docker-compose file
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

RUN if [ "$environment" = "production" ]; \
 then \
    echo "building executable for production" && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./app/cmd/breadcrumbs; \
 else \
    echo "preparing development environment"  && \
    go get github.com/githubnemo/CompileDaemon; \
 fi

EXPOSE 80

## example taken from https://levelup.gitconnected.com/docker-for-go-development-a27141f36ba9
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./app/cmd/breadcrumbs" -command="./breadcrumbs"

# final stage -- only get's here on production build
FROM bash as prod
RUN echo "building production image"
# use builder to reduce final image size
# don't need the installed packages just the final build executable
COPY --from=builder /application/breadcrumbs /
COPY --from=builder /wait /
# expose the port this server runs on
EXPOSE 80
CMD ["/breadcrumbs"]