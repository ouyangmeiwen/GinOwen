# GinOwen

go mod init GoTools //初始化项目
go mod tidy  //更新依赖: 确保你的 go.mod 文件中列出了所有项目所需的依赖项，并确保它们是正确的版本。可以通过运行 go mod tidy 来清理和更新依赖项。
go mod vendor // 如果你的项目需要使用 vendor 目录作为依赖项的本地副本，可以使用以下命令重新生成
如果代码无法弹出正常提示，则检查文件是否为.go结尾

重新生成swagger
go install github.com/swaggo/swag/cmd/swag@latest

swag init
如果为生成的对应接口的swaggerUI
1.格式错误
2.清除浏览器缓存



go 打包
    window下用set 如果在linux 则用export

    windwos:
        SET GOOS=windows
        SET GOARCH=amd64
        go build -o ginowen.exe    然后依赖文件过去比如config.yaml
