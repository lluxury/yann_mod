
// 需要go 1.11


//   大纲
// 创建一个 Module
// 版本化
// 发布版本
// 使用module
// 版本修复漏洞
// 更新 modules
// 主要版本号
// 更新到新的主要版本
// 整理一下
// Vendor 机制



// 创建一个叫 yann_mod 的包
// 目录要在 $GOPATH 之外，默认$GOPATH 里面是禁用 modules 支持


mkdir yann_mod
cd yann_mod

go mod init github.com/lluxury/yann_mod
ls go.mod

vi yann_mod.go

// 手动建库?

git init<br>
git add *<br>
git commit -am "First commit"<br>
git remote add origin https://github.com/lluxury/yann_mod.git<br>
git push -u origin master<br>

go get github.com/lluxury/yann_mod<br>

// Go 根据代码仓库的标签来确定版本
// Go 会获取代码仓库里面设置好标签的最新的版本


// 制作发布版本
git tag v1.0.0<br>
git push --tags<br>


// 创建一个叫 v1 的分支
git checkout -b v1<br>
git push -u origin v1<br>


// git push -u origin master 上面命令将本地的master分支推送到origin主机，同时指定origin为默认主机<br>
// 后面就可以不加任何参数使用git push了

//  不带任何参数的git push，默认只推送当前分支，这叫做simple方式
//  matching方式，会推送所有有对应的远程分支的本地分支
//  Git 2.0版本之前，默认采用matching方法，现在改为默认采用simple方式


// 不能使用 main.go,否则
// build mod: cannot find module for path xxx

go build
// go: finding github.com/lluxury/yann_mod v1.0.0
// go: downloading github.com/lluxury/yann_mod v1.0.0

// module mod
// require github.com/lluxury/yann_mod v1.0.0

// github.com/lluxury/yann_mod v1.0.0 h1:9EdH0EArQ/rkpss9Tj8gUnwx3w5p0jkzJrd5tRAhxnA= <br>
// github.com/lluxury/yann_mod v1.0.0/go.mod h1:UVhi5McON9ZLc5kl5iN2bTXlL6ylcxE9VInV71RrlO8=

// 用过的包都会记录进去

./mod


// 修复bug
vi yann_mod.go

git commit -m "Emphasize our friendliness" yann_mod.go
git tag v1.0.1
git push --tags origin v1

// 在 v1 分支做这些改动，因为这个 bug 只在 v1 版本中存在
// 需要把这个改动应用到多个版本，需要在 master 分支做这些改动，然后再向后移植



// 默认情况下，Go 不会自己更新模块

    go get -u 
    go get -u=patch
    go get package@version 
// 大版本,小版本,指定版本
 // v1.2.3 中的 1 定义为主要版本号，2 为次要版本号，3 为修订版本号


go get -u=patch
go build
./mod

// 主要版本可以打破向后兼容性


vi yann_mod.go 
vi go.mod 
    module github.com/lluxury/yann_mod/v2

git commit yann_mod.go -m "Change Hi to allow multilang"
git checkout -b v2 
echo "module github.com/lluxury/yann_mod/v2" > go.mod
git commit go.mod -m "Bump version to v2"
git tag v2.0.0
git push --tags origin v2 


// go get -u 不会把项目使用的包的版本升级到 2.0.0
// 需要使用要重新引入包  "github.com/lluxury/yann_mod/v2"


// 移除掉不使用的依赖
go mod tidy


go run fmt // 也可以更新包

// 兼容现有的方式,需要主动编译
go mod vendor
go build -mod vendor
