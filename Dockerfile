FROM golang:1.12-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Sami Belaid & Ons Tliba"

# Set the Current Working Directory inside the container
WORKDIR /app

# Fetch dependencies on separate layer as they are less likely to
# change on every build and will therefore be cached for speeding
# up the next build
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

# copy source from the host to the working directory inside
# the container
COPY ./app .

# Build the Go app
RUN go build -o main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Run the executable
CMD ["./main"]
