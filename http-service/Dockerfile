FROM golang:1.17.0 AS build_stage

# setup environment
ARG work_dir=/app
RUN mkdir -p ${work_dir}/src
ADD . ${work_dir}/src
WORKDIR ${work_dir}/src

# In build process, makefile will copy your ${HOME}/.ssh into project folder as ssh_keys.
COPY ssh_keys /root/.ssh
RUN go install github.com/google/wire/cmd/wire@v0.5.0 && CGO_ENABLED=0 ${GOPATH}/bin/wire gen ./internal/injector
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cmd/main cmd/main.go

# In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:3.13.2
ARG work_dir=/app
WORKDIR ${work_dir}
RUN mkdir -p ${work_dir} && mkdir -p ${work_dir}/configs

# copy binary into image
COPY --from=build_stage ${work_dir}/src/cmd/main ${work_dir}/main
COPY --from=build_stage ${work_dir}/src/cert ${work_dir}/cert

# entrypoint
COPY entrypoint.sh ${work_dir}/entrypoint.sh
ENV WORK_DIR ${work_dir}
ENV FLAGS "--file ${WORK_DIR}/configs/config.yaml"
ENTRYPOINT ["./entrypoint.sh"]
CMD ""
