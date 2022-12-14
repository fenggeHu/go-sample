module go-sample

go 1.19

require (
	github.com/confluentinc/confluent-kafka-go v1.8.2
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/assert/v2 v2.0.1
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/henrylee2cn/pholcus v1.3.4
	github.com/henrylee2cn/pholcus_lib v0.0.0-20180312025545-3ffb61273b75
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/robfig/cron/v3 v3.0.1
	github.com/segmentio/kafka-go v0.4.30
	gorm.io/driver/sqlite v1.4.3
	gorm.io/gorm v1.24.2
)

// 引入本地项目
replace github.com/henrylee2cn/pholcus_lib => ../pholcus_lib

require (
	github.com/DataDog/zstd v1.3.6-0.20190409195224-796139022798 // indirect
	github.com/Shopify/sarama v1.23.1 // indirect
	github.com/andybalholm/cascadia v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/hashicorp/go-uuid v1.0.1 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20190807075143-e8afa09140e9 // indirect
	github.com/henrylee2cn/teleport v1.0.0 // indirect
	github.com/jcmturner/gofork v0.0.0-20190328161633-dc7c13fece03 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/kr/beanstalk v0.0.0-20180818045031-cae1762e4858 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lxn/walk v0.0.0-20190619151032-86d8802c197a // indirect
	github.com/lxn/win v0.0.0-20190716185335-d1d36f0e4f48 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/pierrec/lz4 v0.0.0-20190327172049-315a67e90e41 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d // indirect
	github.com/tidwall/gjson v1.14.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.2.3 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
