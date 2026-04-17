# Stage 1
# Here we are pulling the base image for golang with enough ability to run a go project
FROM golang:alpine3.22 AS builder

# Here we are creating a directory inside the image where we will store our go project files
WORKDIR /ascii-art-web-stylize

# Here we are copying the go.mod file in our go project to the directory we created in the image
COPY go.mod ./

# Here we run the command below to download all the neccessary dependencies needed to run our docker project
RUN go mod download

# Here we are copying all the other files and directories in our project to the directory we created in the image
COPY . .

# Here we are exposing the port our project is working with for accepting requests from the web
EXPOSE 8081

# Here we are building our go project into a binary file
RUN CGO_ENABLED=0 GOOS=linux go build -o /ascii-art ./cmd/

# Stage 2
FROM scratch

WORKDIR /root/

COPY --from=builder /ascii-art .

COPY --from=builder /ascii-art-web-stylize/fonts /root/fonts

COPY --from=builder /ascii-art-web-stylize/web /root/web

# This is our entry point from which we can run an instance of this image (container) in docker
CMD [ "./ascii-art" ]
