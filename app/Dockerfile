# build stage
FROM golang:1.13 as builder

ARG environment
RUN echo "Running $environment build"

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
# result is "breadcrumbs" executable in the current folder

# if dev build then
RUN if [ "$environment" = "production" ]; \
 then echo "building executable for production" && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/breadcrumbs; \
 else echo "preparing development environment"  && go get github.com/githubnemo/CompileDaemon; \
 fi

# final stage -- only get's here on production build
FROM scratch as prod
# use builder to reduce final image size
# don't need the installed packages just the final build executable
COPY --from=builder /app/breadcrumbs /app/
# COPY --from=builder wait ./
# expose the port this server runs on
EXPOSE 80
# run this container as an executable
ENTRYPOINT ["/app/breadcrumbs"]