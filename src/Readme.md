### go enviroment build
#### 1. 下载golang环境
windows 下载go1.13.1.windows-amd64.msi。
下载后安装一路默认即可，安装完成后命令行输入`go version`，检查安装成功。
```cmd
c:\>go version
go version go1.13.1 windows/amd64
```

Linux 下载go1.13.1.linux-amd64.tar.gz
```bash
tar -C /usr/local -xzf go1.13.1.linux-amd64.tar.gz
vim /etc/profile
# 添加以下环境变量
# export PATH=$PATH:/usr/local/go/bin
```

#### 2. 构建vscode-go环境
##### 2.1 vscode terminal设置为git bath
编辑 `settings.json`，便捷`terminal.integrated.shell.windows`项
```json
{
    "terminal.integrated.shell.windows": "your_git_path\\Git\\bin\\bash.exe"
}
```
设置GOPATH下载vs code插件，GOPATH不能和GOROOT（GO安装目录）一样
有些插件无法下载，需要下载到gopath下面，手动安装,以`sync\errgroup`为例
```bash
$ go install golang.org/x/tools/gopls
D:\Go\src\golang.org\x\tools\internal\lsp\source\analysis.go:20:2: cannot find package "golang.org/x/sync/errgroup" in any of:
        D:\Go\src\vendor\golang.org\x\sync\errgroup (vendor tree)
        D:\Go\src\golang.org\x\sync\errgroup (from $GOROOT)
        D:\gopath\src\golang.org\x\sync\errgroup (from $GOPATH)
D:\Go\src\golang.org\x\tools\internal\lsp\protocol\span.go:13:2: cannot find package "golang.org/x/xerrors" in any of:
        D:\Go\src\vendor\golang.org\x\xerrors (vendor tree)
        D:\Go\src\golang.org\x\xerrors (from $GOROOT)
        D:\gopath\src\golang.org\x\xerrors (from $GOPATH)
# 下载代码
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x


```


##### 2.2 编辑，运行go
使用vscode新建文件`hello.go`
```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```
在terminal中运行
```bash
go run hello.go
hello, world
```

#### 3. 问题
* 指针和接口方法定义关系（实现接口时，类型参数和指针类型参数）？
* struct 怎么取指针？

#### 相关链接
* [golang](https://golang.google.cn/dl/)
* [vscode](https://code.visualstudio.com/docs?start=true)
* [vscode-plugin-go](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)