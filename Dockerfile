FROM golang:alpine AS builder

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

# 将代码复制到容器中
COPY . .

RUN chmod 755 wait-for-it.sh

CMD ["air", "-c", ".air.toml"]

# ENTRYPOINT ["air"]

# 需要运行的命令（注释掉下面这一行）
# ENTRYPOINT ["/ginchat", "config/config.ini"]
