package main


import "testing"
// 基准测试
//run: #go test -v -run="none" -bench=.
// -run 只跑benchmark，不跑正常的testcase(t *testing.T结尾的用例)
// -bench=. 跑所有的benchmark
//
//leemingeer@mbp concurrent % go test -v -run="none" -bench=.
//goos: darwin
//goarch: amd64
//pkg: 1stgoproj/aLearnFromBli/7testing/concurrent
//BenchmarkConcurrentAtomicAdd
//BenchmarkConcurrentAtomicAdd-8              5232            232761 ns/op Atomic 232761ns,更快
//BenchmarkConcurrentMutexAdd
//BenchmarkConcurrentMutexAdd-8               3338            353589 ns/op
//PASS
//ok      1stgoproj/aLearnFromBli/7testing/concurrent     4.178s

func BenchmarkConcurrentAtomicAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentAtomicAdd()
	}
}

func BenchmarkConcurrentMutexAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentMutexAdd()
	}
}