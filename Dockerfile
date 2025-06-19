   # 以官方的Go基础镜像为基础，这里使用Go 1.19，你可以根据实际情况选择版本
FROM docker.wanpeng.top/library/golang:latest


   # 设置工作目录，后续的操作将在这个目录下进行
WORKDIR /app

   # 将当前目录下的所有文件（包括Go源文件、模块文件等）复制到镜像的 /app 目录下
COPY . /app

   # 构建Go应用程序，假设主程序是main.go，这将生成一个可执行文件
ENV GOARCH=amd64
ENV GOOS=linux
RUN go build -o mobile-storage-test-service

   # 暴露应用程序可能需要的端口，比如如果你的Go应用是一个Web服务，运行在8080端口，就暴露8080端口
EXPOSE 8080

   # 定义容器启动时要执行的命令，这里是启动刚才构建的Go应用程序
CMD ["/app/mobile-storage-test-service"]

# docker buildx build --platform linux/amd64 -t mobile-storage-test-service:1.0.0 . --load