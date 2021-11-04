module 1stgoproj

go 1.15

require (
	github.com/chilts/sid v0.0.0-20190607042430-660e94789ec9
	github.com/go-openapi/spec v0.19.3 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.5.0
	github.com/google/uuid v1.1.2
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inhies/go-bytesize v0.0.0-20201103132853-d0aed0d254f8
	github.com/kjk/betterguid v0.0.0-20170621091430-c442874ba63a
	github.com/oklog/ulid v1.3.1
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.10.1
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/rs/xid v1.3.0
	github.com/satori/go.uuid v1.2.0
	github.com/segmentio/ksuid v1.0.3
	github.com/sony/sonyflake v1.0.0
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.22.1 // indirect
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.18.15
	k8s.io/klog/v2 v2.9.0
	k8s.io/kubernetes/pkg/scheduler/framework v1.20.0
)

replace k8s.io/kubernetes => k8s.io/kubernetes v1.21.4
