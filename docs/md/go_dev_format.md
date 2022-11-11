
# 一、命名规范

## 1. 标识符

按照Go规定的标志符命名规则，只包含字母、数字和下划线，并且均以字母开头。代码中的命名严禁使用拼音与英文混合的方式，更不允许直接使用中文的方式。正确的英文拼写和语法可以让阅读者易于理解，避免歧义。

说明：拼音同音不同义，会产生理解上的歧义；语言本身的基础包、依赖包均采用英文编写，引入中文变量和方法看起来不协调。

正例：`var standard` [标准] 、`getScoreByName`[评分] 、`var flag int = 1`

反例：`var biaozhun` [标准] 、`getPingfenByName()` [评分] 、`var int 标志 = 3`

## 2. 包

包名采用小写的一个单词，尽量不要和标准库冲突。如果包含多个单词，直接连写，中间无需使用下划线连接， 不使用驼峰形式。多个单词的情况下，通常考虑对包按照每个单词进行分层。

正例：`admin`、`stdmanager`、`encoding/base64`

反例：`Admin`、`std_manager`、`encodingBase64`、`encoding_base64`

## 3. 文件夹

文件夹的名称要与所包含的代码文件中的包名保持一致，但`package main`除外。

正例：文件夹 `controllers` 下包含了全部`package controllers`的go文件

## 4. 文件

go文件的命名采用小写单词，尽量见名思义，看见文件名就可以知道这个文件下的大概内容。测试文件必须以`_test.go`结尾。

正例：`struct Role`所在文件为`role.go`，对应的测试文件为`role_test.go`

## 5. 变量

全局变量、成员变量的命名必须遵从驼峰形式，如果包外可见使用 `UpperCamelCase` 风格，包外不可见采用 `lowerCamelCase`。

```go
  正例：
    type UserController struct
    func isValidNumber(s string)

  反例：
    type Usercontroller struct
    func isValidnumber(s string)
```

局部变量、函数参数的命名全部采用 `lowerCamelCase`。我们尽量让局部变量和函数参数的命名意义明确。

```go
    正例：函数参数 
    func Open(driverName, dataSourceName string) (*DB, error)
```

变量的长短一般与作用域的大小相关，作用域范围很小的局部变量可以尽量简短，比如循环语句中:

```go
    for i := 0; i < 10; i++ 
    for i, v := range slice
```

## 6. 常量

常量命名全部大写，单词间用下划线隔开，力求语义表达完整清楚，不要嫌名字长。

```go
   正例：MAX_STOCK_COUNT
   反例：MAX_COUNT
```

Golang中没有专门的枚举类型（enum），通常使用一组常量来表示，为了更好的区分不同的枚举类型值，应该使用完整的前缀加以区分：

```go
   type PullRequestStatus int
   const  (
     PULL_REQUEST_STATUS_CONFLICT PullRequestStatus = iota
     PULL_REQUEST_STATUS_CHECKING
     PULL_REQUEST_STATUS_MERGEABLE
   )
```

## 7. 接口名

单个函数的接口名以"er"作为后缀，其函数去掉"er"如Reader,Writer。

```go
type Reader interface {
 Read(p []byte) (n int, err error)
}
```

两个函数的接口名综合两个函数的名字，比如：

```go
type WriteFlusher interface {
 Write([]byte) (int, error)
 Flush() error
}
```

三个函数及以上的接口名类似于结构体名，比如：

```go
type Car interface {
 Start([]byte) 
 Stop() error
 Recover()
}
```

# 二、代码格式

## 1. IDE

- IDE 的 text file encoding 设置为 UTF-8
- IDE 中文件的换行符使用 Unix 格式，不要使用 Windows 格式

## 2. gofmt

- gofmt能够自动格式化代码，所有格式有关问题，均以 gofmt 格式化的结果为准
- 代码提交前，必须执行gofmt进行格式化
- IDE一般都会提供gofmt功能

## 3. 缩进

代码对齐应该使用tab对齐，不推荐使用空格对齐。

## 4. 行宽

单行字符数不超过 100 个（主要参考IDE编辑器的宽度可见范围），超出则需要换行。换行时遵循如下原则：

- 第二行相对第一行缩进 1个tab，从第三行开始，不再继续缩进。
- 运算符与下文一起换行。
- 方法调用的点符号与下文一起换行。
- 方法调用时，多个参数，需要换行时，在逗号后进行。
- 在括号前不要换行。

```go
  正例：
 condition := orm.NewCondition()
 // 超过 100 个字符的情况下，换行缩进 1 个tab，点号和方法名称一起换行
 condition.And("implement_date__gte", "a").And("b")...
  .And("c")...
  .And("d")...
  .And("e")
```

## 5. 大括号

大括号的使用约定。

1. 左大括号前不换行。
2. 左大括号后换行。
3. 右大括号前换行。
4. 右大括号后还有 else 等代码则不换行；表示终止的右大括号后必须换行。

```go
正例：
if a == b {
} else {
}
```

## 6. 空格

左小括号和字符之间不出现空格；同样，右小括号和字符之间也不出现空格。

  正例：(a == b)
  反例：(空格 a == b 空格)
任何二目运算符的左右两边都需要加一个空格。一目运算符与变量之间不添加空格。
    正例：v := 3、a + b、a && b、!ok
函数参数在定义和传入时，多个参数逗号后边必须加空格。
    正例：下例中实参的"a"，后边必须要有一个空格。
    func("a", "b", "c")

## 7. 空行

函数体内的执行语句组、变量的定义语句组、不同的业务逻辑之间插入一个空行，但无需插入多个空行进行分隔。相同业务逻辑之间无需插入空行。

## 8. 注释

Go有两种注释方式，块注释 /**/ 和 行注释 // 。 Godoc用来处理Go源文件，抽取有关程序包内容的文档。在顶层声明之前出现，若中间没有换行的注释，会随着声明一起被抽取，作为该项的解释性文本。 每个程序包都应该有一个包注释，位于包声明之前，比如：

```go
/*
Package regexp implements a simple library for regular expressions.
The syntax of the regular expressions accepted is:
    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

如果程序包很简单，则包注释可以非常简短：

```go
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

函数的注释，第一条语句应该为一条概括语句，并且使用被声明的名字作为开头。比如：

```go
// Compile parses a regular expression and returns, if successful, a Regexp
// object that can be used to match against text.
func Compile(str string) (regexp *Regexp, err error) {}
```

变量的注释，可以对声明进行组合，比如：

```go
// Error codes returned by failures to parse an expression.
var (
    ErrInternal = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

# 三、Import规范

import在多行的情况下，使用goimports会自动进行格式化。在一个文件里面引入了一个包，建议采用如下格式：

```go
import (
    "fmt"
)
```

如果文件中引入了三种类型的包：标准库包，程序内部包，第三方包，建议采用如下方式进行组织：

```go
import (
    "encoding/json"            // 标准库包
    "strings"

    "myproject/models"         // 程序内部包
    "myproject/controller"
    
    "git.obc.im/obc/utils"     // 第三方包
    "git.obc.im/dep/beego"
    "git.obc.im/dep/mysql"
)
```

导入包时，尽量使用绝对路径，使用程序内部包无需包含路径：

```go
import "xxxx.com/proj/net"  // 正确
import "../net"             // 错误
```

# 四、错误处理

## 1. error

error作为函数的返回值，必须尽快对error进行处理，绝对不允许忽略。

通常采用独立的错误流进行处理，不要采用下面这种方式：

```go
   if err != nil {
     // error handling
   } else {
     // normal code
   }
   
   而是采用这种方式：
   
   if err != nil {
    // error handling
    return // or continue, etc.
   }
   // normal code
   
  如果返回值需要初始化，则采用如下方式：
   
   x , err := f()
   if err != nil {
     // error handling
     return // or continue, etc.
   }
   // use x
   
```

## 2. recover

recover用于捕获runtime的异常，禁止滥用recover，在开发测试阶段尽量不要用recover，这样可以充分暴露错误。recover一般放在你认为会有不可预期的异常的地方。比如：

```go
   func safelyDo(work *Work) {
     defer func() {
       if err := recover(); err != nil {
          log.Println("work failed:", err)
       }
     }()
     // do函数可能会有不可预期的异常
     do(work)
   }
```

## 3. panic

用来创建一个 RuntimeException 并结束当前程序。该函数接受一个任意类型的参数，并在程序挂掉之前打印该参数内容，通常选择一个字符串作为参数。比如：

```go
   func init() {
     if user == "" {
       panic("no value for $USER")
     }
   }
  
```

应该在逻辑处理中禁用panic。在main包中只有当实在不可运行的情况采用panic，例如文件无法打开，数据库无法连接导致程序无法正常运行，但是对于其他的package对外的接口不能有panic，只能在包内采用，并使用recover进行捕捉处理。建议在main包中使用log.Fatal来记录错误，这样就可以由log来结束程序。

## 4. defer

该函数会在return前执行，对于一些资源的回收使用defer非常方便，但禁止滥用defer，defer是需要消耗性能的，所以频繁调用的函数尽量不要使用defer。

```go
  // Contents returns the file's contents as a string.
  func Contents(filename string) (string, error) {
      f, err := os.Open(filename)
      if err != nil {
          return "", err
      }
      defer f.Close()  // f.Close will run when we're finished.

      var result []byte
      buf := make([]byte, 100)
      for {
          n, err := f.Read(buf[0:])
          result = append(result, buf[0:n]...)
          if err != nil {
              if err == io.EOF {
                  break
              }
              return "", err  // f will be closed if we return here.
          }
      }
      return string(result), nil // f will be closed if we return here.
  }
  
```

# 五、安全处理

代码中执行的SQL语句参数，应该严格使用参数绑定，防止 SQL 注入，禁止字符串拼接 SQL访问数据库。使用参数绑定的方式，带来的另一个好处是无需处理繁琐的字符转义。

```go
正例：
    sqlQueryUnit := "select unit_id from unit_draft where unit_name = ? limit 1" 
    stmtQueryUnit, err := db.Prepare(sqlQueryUnit)
    if err != nil {
      return err
    }
    defer stmtQueryUnit.Close()
    
正例：
    count, err := orm.NewOrm().Raw("SELECT id,name_cn,name_en FROM std_info WHERE std_no = ?", stdNo).QueryRows(&stdNos)
    if err != nil {
      return err
    }
```

# 六、包依赖管理

\1. go1.11.1版本以上可以使用module机制做包依赖管理，主要优点有：

- 项目路径可以脱离GOPATH，不需要将项目必须放在GOPATH/src下。
- 项目依赖的第三方包，不再需要放入GOPATH/src下，放入项目的vendor目录下，与项目代码接受同样的源码管理。
- 不需要使用get预先安装依赖，module在 build、run、test时会检测未下载的依赖，并自动下载它们。

\2. go modules 操作过程

- 初始化命令：go mod init ，项目目录下会产生go.mod文件，里面记录了module路径和go的版本信息。
- 获取依赖包命令：go build、run、test，运行后会自动下载和安装依赖包，代码位于GOPATH/pkg目录下。
- 依赖包交由vendor管理：go mod vendor，GOPATH/pkg目录下的第三方依赖包放入项目下的vendor目录下。
