FROM golang:1.19 as builder
COPY . /src
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct
RUN cd /src && \
 go build -o app .

FROM ubuntu
COPY --from=builder /src/app /src/app
COPY ./configs /src/configs
WORKDIR /src
CMD ["/src/app"]

EXPOSE 8080

# docker build -t zlyuan/http-print .
