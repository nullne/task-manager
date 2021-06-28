### Task Manager

实现一个简单的任务管理器，功能包括：

1. 创建，销毁任务管理器（Manager）
2. 给任务管理器添加一个任务（Task），这个任务会被尽快执行。添加任务本身应该快速返回。每一个任务是一个 `TaskFn`，执行没有任何参数，仅返回 `error` ，任务可配置超时时间跟最大重试次数 N，该操作确保多线程安全
3. 任务管理器可以设定最大并发量，当并行任务超过最大并发量时，按照任务添加的时间排队执行
4. 添加进去的任务可以通过返回值，查看任务执行状态，或者阻塞等待任务执行完成
5. 当某一个任务执行出错时，通过判断其报错类型，如果是可重试的错误（`TaskFn` 返回的error 类型如果实现了 `temporary` 接口，当 `Temporary()` 返回为 `true` 即表明该任务可以重试， 可参考 [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)），需要给当前任务进行重试，最多重试 N 次

要求： 

1. 参考已经定义好的接口，可按照实际情况酌情进行更改
2. 充分的 unit test 和 bench test， 确保没有内存泄露，或者 Race condition 发生
3. 仅使用标准库