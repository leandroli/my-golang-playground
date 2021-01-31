# my go learning note

### 关于range
比如 i, val := range arr
range会给i和val分别分配一个且仅此一个空间，每次循环都会换空间里的内容，但是空间还是那个空间，
所以如果在循环里往存函数的slice里加新函数并且新函数中用了i或者val，记得在循环里复制一下再在新函数里使用，
否则会造成所有加入的新函数被调用时使用的i和val实际上都是同一个变量。Page141 Chapter5.6 gopl
同时记得：更改i，val不会改变arr的值！就像传值引用一样。

### 关于用model怎么import以及import远程包

用model import:
import "modelName/projectName/pkgName"
model name 和 project name 都可以在go.mod文件中找到
import远程包:
import "github/LIZHUO99/pathAndPkgName"

### goroutine和scheduler
goroutine实际上相当于 user level thread 一个系统线程对应多个 goroutine。
schedule实现：https://www.bilibili.com/video/BV1kz411e7dL?t=36 gophercon 2018 Kavya Joshi 