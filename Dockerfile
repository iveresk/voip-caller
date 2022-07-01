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
COPY exploit ./src/
RUN go build -o /cve-2022-21907
CMD [ "/cve-2022-21907", "-t", ${INPUT_FILE}, "-d", "false", "-s", "<YOUR_CALL_PHRASE>" ]