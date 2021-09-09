#!/usr/bin/env bash
curl -SLO https://golang.google.cn/dl/go1.16.5.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
export PATH=/usr/local/go/bin:$PATH
go versin