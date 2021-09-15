# 本地单独跑arstor e2e
#!/bin/bash
export KUBECONFIG=/root/.kube/config
export ARSTORPLUGIN_NAMESPACE=default
export PORTALS='["10.0.211.10:80"]'
export TARGETPORTAL="10.10.10.5:3260"
export IQN="iqn.2010-01.com.huayun:k8starget"
export GOPROXY=https://goproxy.cn,direct
go test -c -o e2e.test ./tests/
`go env GOPATH`/bin/ginkgo e2e.test
