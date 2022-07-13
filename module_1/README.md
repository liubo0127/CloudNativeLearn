**go语言特性**

+ Go 语言由来和特性
+ Go 语言环境设置
    - [Hello world](./hello_world)
    - [testing](./pri_add)
    - [go vet](./go_vet)
+ Go 语言控制结构，常用数据结构
  - [if](./if_condition)
  ```go
  // 示例一
  if condition1 {
    // do something1
  } else if condition2 {
    // do something2
  } else {
    // else
  }
  
  // 示例二，带赋值
  if v:=x; v>0{
    return v
  } else {
    return 0
  }
  ```
  - [switch](./switch_case)
  ```go
  switch val{
    case val1: //什么都没做, pass
    case val2:
        fallthrough // 会继续执行下一个 case，不管条件是否符合都会执行
    case val3:
        f() // 执行 f 函数
    default:
        ... // 默认行为，当前面都没匹配到时执行，或者匹配到了，但是有 fallthrough 继续往下执行到这里
  }
  ```
  + [for](./for_loop)
    - 计数器
    
    ```go
    // 示例一 初始化语句；条件语句；修饰语句
    for i:=0;i<10;i++{
    }
    
    // 示例二，只加条件判断，第一个（初始化语句）和第三个（后置语句）都是可以省略的
    for ;sum<1000; {
       sum += sum
    ```
    
    - 无限循环
    
    ```go
    // 示例三，无限循环
    for {
    	if condition1 {
            break
        }
    }
    ```
    
    - 遍历
    
    > 如果使用 for 遍历指针数组，value 取出的指针地址为原指针地址的拷贝
    
    ```go
    // 示例一，遍历字符串
    for index, char := range myString { // index 是字符的位置，char 是字符的值
        fmt.Printf("index: %d, value: %s\n", index, char)
    }
    
    // 示例二，遍历 map
    for key, value := range myMap {
        fmt.Printf("key: %v, value: %v\n", key, value)
    }
    
    // 示例三，遍历数组/切片
    for index, value := range myArray {
        fmt.Printf("index: %d, value: %v", index, value)
    }
    ```
  
+ Go 语言函数
+ Go 语言常用语法，多线程