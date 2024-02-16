FROM golang:1.17 AS build_stage

# setup environment
ARG work_dir=/app
RUN mkdir -p ${work_dir}/src
WORKDIR ${work_dir}/src

# build code
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"
ENV GO111MODULE=on
ENV GOPATH=${work_dir}
ENV GOPRIVATE=github.com/17media

RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts
COPY go.mod go.mod
COPY go.sum go.sum
RUN --mount=type=ssh,id=main go mod download
COPY . ${work_dir}/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/main cmd/main.go

# In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:3.16.2
ARG work_dir=/app
WORKDIR ${work_dir}
RUN mkdir -p ${work_dir} && mkdir -p ${work_dir}/configs

# copy binary into image
COPY --from=build_stage ${work_dir}/src/cmd/main ${work_dir}/main
COPY --from=build_stage ${work_dir}/src/cert ${work_dir}/cert
COPY --from=build_stage ${work_dir}/src/configs/config.yaml ${work_dir}/configs/config.yaml

# entrypoint
COPY entrypoint.sh ${work_dir}/entrypoint.sh
ENV WORK_DIR ${work_dir}
ENV CONFIG_FILE "${WORK_DIR}/configs/config.yaml"
ENTRYPOINT ["./entrypoint.sh"]
CMD ""
