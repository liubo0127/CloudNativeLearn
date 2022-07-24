**编写 GO 程序**

+ 线程加锁
  + Sync 包(锁)
    + [sync.Mutex 互斥锁](sync_mutex_lock/main.go)
      + Lock() 加锁
      + Unlock() 解锁
    + [sync.RWMutex 读写分离锁](sync_mutex_lock/main.go)
      > 不限制并发读，只限制并发写和并发读写
    + [sync.WaitGroup](sync_waitgroup/main.go)
      > 等待一组 goroutine 返回
    + sync.Once
      > 保证某段代码只执行一次
    + [sync.Cond](sync_cond/main.go)
      > 让一组 goroutine 在满足特定条件时被唤醒
+ 线程调度
+ 内存管理
  > 是在 TCMalloc 原型上做的修改
  + TCMalloc
    + PageHeap
    + CentralCache
    + ThreadCache
    + 概念
      + page 内存页
      + span 内存块
      + size class 空间规格
      + object 对象
        + 小对象: 0 ~ 256KB
        + 中对象: 256KB ~ 1MB
        + 大对象: > 1MB
  + Go 语言内存分配
    + mheap
      + 不适用链表了，维护了两个排序二叉树，为了搜索更快
        + free
        + scav
    + mcentral
      + mcentral 每个 span class 维护两个链表
      > 先去 non empty 申请内存，不够的话再去 empty 申请，还不够的话再去 mheap 申请内存
        + non empty
        + empty
    + mcache
      + 一个 size class 对应两个 span class，主要用于 GC
        + 一个 span class 存储指针，GC 需要扫链表
        + 一个 span class 存储直接引用，不需要回收
  + Go 内存回收
    > 标记-清除
    > 从根变量开始遍历所有引用对象，引用的对象标记为 "被引用"，没有被标记的进行回收
    + 原理
      + mspan
        + allocBits
          > 记录了每块内存分配的情况
        + gcmarkBits
          > 记录了每块内存的引用情况，标记阶段对每块内存进行标记，有对象引用的内存标记为 1，没有的标记为 0
          > 标记为 0 的就可以在申请内存时分配出去
    + GC 工作流程
      + Mark
        + Mark Prepare
          > 初始化 GC 任务，包括开启写屏障（write barrier）和辅助 GC（mutator assist），统计 root 对象的任务数量等。需要 STW
        + Mark Drains
          > 扫描所有 root 对象，包括全局指针和 goroutine（G）栈上的指针（扫描对应 G 栈时需停止该 G），将其加入标记队列（灰色队列），并循环处理灰色队列的对象，直到灰色队列为空。后台并行执行
      + Mark Termination
        > 完成标记工作，重新扫描（re-scan）全局指针和栈。因为 Mark 和用户程序是并行的，所以在 Mark 过程中可能会有新的对象分配和指针赋值，这个时候就需要通过写屏障记录下来，re-scan 再检查一下。也会 STW
      + Sweep
        > 按照标记结果回收所有的白色对象，后台并行执行
      + Sweep Termination
        > 对未清扫的 span 进行清扫，只有上一轮的 GC 清扫工作完成才可以开始新一轮 GC
    + 三色标记
      + GC 开始时，认为所有 object 都是白色，即垃圾
      + 从 root 区开始遍历，被触达的 object 置成灰色
      + 遍历所有灰色 object，将他们内部的引用变量置成灰色，自身置成黑色
      + 循环第三步，直到没有灰色 object 了，只剩下了黑白两种，白色的都是垃圾
      + 对于黑色 object，如果在标记期间发生了写操作，写屏障会在真正赋值之前将新对象标记为灰色
      + 标记过程中，mallocgc 新分配的 object 会先被标记成黑色再返回
    + 垃圾回收触发机制
      + 内存分配量达到阈值触发 GC
        > 每次内存分配时都会检查当前内存分配量是否已达到阈值，如果达到阈值则立即启动 GC
        + 阈值：上次 GC 内存分配量 * 内存增长率
        + 内存增长率由环境变量 GOGC 控制，默认为 100，即每当内存扩大一倍时启动 GC
      + 定期触发 GC
        > 默认情况下，最长 2 分钟触发一次 GC，这个间隔在 `src/runtime/proc.go:forcegcperiod` 变量中被声明
      + 手动触发
        > 程序代码中可以使用 runtime.GC() 来触发 GC。主要用于 GC 性能测试和统计。
+ 包引用与依赖管理
  + vendor
    > 1.6 版本之后支持 vendor
    > 在项目根目录创建 vendor 目录，并将依赖拷贝至该目录。需要手动操作
    > Go 语言项目会自动将 vendor 目录作为自身的项目依赖路径
  + 依赖管理工具
    + 社区
      + Godeps
      + Glide
    + 官方
      + Gopkg（旧工具）
      + gomod（新工具）
        > `go mod vendor` 将依赖放到 vendor 目录
  + 环境变量
    + GOPROXY
      > 为拉取 Go 依赖设置代理
    + GOPRIVATE
      > 声明私有仓库
  + Makefile
+ [编写 HTTP Server](./httpserver/main.go)
+ Go 语言调试
  + Debug
    + 工具
      + gdb
        + Gccgo 原生支持 gdb，因此可以用 gdb 调试 Go 语言代码，但 dlv 对 Go 语言 debug 的支持更好
        + gdb 对 Go 语言的栈管理，多线程支持等方面做的不够好，调试代码可能有错乱现象
      + dlv
        + Go 语言的专有 debugger
    + 添加日志
      + fmt
      + 日志框架
        > 可配置 appender，输出到标准输出还是文件
        > 日志分级
        > 时间戳，代码行等
        + glog(官方)
          ```go
          flag.Parse() // 解析参数，可以使用 --help 查看 log 相关参数
          defer glog.Flush()
          glog.V(2).Info("")
          glog.Info("This is info message")
          glog.Infof("This is info message: %v", 12345)
          glog.InfoDepth(1, "This is info message", 12345)
          ```
        + 其他第三方日志框架
  + 性能分析
    + `runtime.pprof`
      > 使用 `go tool pprof $filename` 来查看生成的 pprof 二进制文件
      + [CPU profile](./cpuprofile/main.go)
      + Memory profile
      + Block Profiling
      + Goroutine Profiling
    + [net/http/pprof 示例](./httpserver/main.go)