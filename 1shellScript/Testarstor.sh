# 本地单独跑arstor e2e
#!/bin/bash
#- PORTALS: docker exec -it mxsp curl localhost/api/v3/ping
#- TARGETPORTAL: docker exec -it mxsp sh -c "mxSCST showvip"
export KUBECONFIG=/root/.kube/config
export ARSTORPLUGIN_NAMESPACE=default
export PORTALS='[""179.19.2.99:80"]'
export TARGETPORTAL="11.11.11.9:3260"
export IQN="iqn.2010-01.com.example:k8starget"
export GOPROXY=https://goproxy.cn,direct
go test -c -o e2e.test ./tests/

  $ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
`go env GOPATH`/bin/ginkgo e2e.test
