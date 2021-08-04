
# 模型：
       ->          ->
client    server1    server2
       <-         <-
# why?
比如第二个req&res 耗时1mins，而client超过2s就中断请求，此时server1和server2还在处理请求，造成资源浪费

期望当client中断时（不想要这次请求了），server1&2立刻知道，停止相应的请求，

# How ?
需要有个上下文（context），当client停止请求时，可以传递给server1，server1传给server2

新建一个空的context
ctx := context.Background() 根context， 该context不能被取消，用在main或者顶级请求中来派生其他context
txt := context.TODO() // 空context, 只能用于高等级或者当您不确定使用什么context类型