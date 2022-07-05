# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
from golang:1.17-alpine

# Adding Variables
ARG INPUT_FILE="targets.txt"
ENV INPUT_FILE=${INPUT_FILE}

# Download necessary Go modules
WORKDIR app/
COPY go.mod ./
RUN go mod download

# COPY necessary files
ADD . ./
COPY *.go ./src/
COPY *.txt ./src/
RUN go build -o /voip-caller
CMD [ "/voip-caller", "-t", ${INPUT_FILE}, "-s", "<YOUR-PHRASE>", "-d", "false" ]
