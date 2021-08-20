func Describe(text string, body func()) // 第二个参数是func()类型的，即没有入参也没有返回值，特别适合匿名函数类型func(){ do something}
You can name the TestCart function however you want, as long as it begins with Test, since it’s just a regular Go test function.
The call to RegisterFailHandler registers a handler, the Fail function from the ginkgo package, with Gomega. This creates the coupling between Ginkgo and Gomega


Ginkgo与Go现有的测试基础设施挂钩。这允许您使用 `go test` 运行Ginkgo套件。
这也意味着Ginkgo测试可以与传统的Go测试一起使用。 go test和ginkgo都可以在你的套件中运行所有测试。

Administrator@mli MINGW64 ~/go/src/goproj/test/ginkgo (master)
$ go test -c -o add.test

Administrator@mli MINGW64 ~/go/src/goproj/test/ginkgo (master)
$ ./add.test
Running Suite: Shopping Cart Suite
==================================
Random Seed: 1629441056
Will run 3 of 3 specs

+++
Ran 3 of 3 Specs in 0.011 seconds
SUCCESS! -- 3 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS


------
Administrator@mli MINGW64 ~/go/src/goproj/test/ginkgo (master)
$ ginkgo.exe
Running Suite: Shopping Cart Suite
==================================
Random Seed: 1629442037
Will run 3 of 3 specs

+++
Ran 3 of 3 Specs in 0.014 seconds
SUCCESS! -- 3 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 48.6574082s
Test Suite Passed


