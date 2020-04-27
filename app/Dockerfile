# build stage
FROM golang:1.13 as builder
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
RUN echo 'here'
# build the app into an executable
# result is "breadcrumbs" executable in the current folder
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/breadcrumbs

# final stage
FROM scratch
# use builder to reduce final image size
# don't need the installed packages just the final build executable
COPY --from=builder /app/breadcrumbs /app/
# COPY --from=builder wait ./
# expose the port this server runs on
EXPOSE 80
# run this container as an executable
ENTRYPOINT ["/app/breadcrumbs"]