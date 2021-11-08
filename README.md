# goblog
a web blog system in go for practice.

# go program

go语言通过package形成模块化，从package main开始，main.go文件中的func main() 为入口。

从main.go中import一系列的package，完成整个程序。

创建go项目的步骤：
1. 创建src/path/to/your/package go init，创建go.mod，完善模块化信息
2. 创建子包，引入子包
3. 如果需要第三方包，引入第三方包，通过go get -u url来获取对应的包
