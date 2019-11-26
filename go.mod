module github.com/sensu/sensu-go

go 1.13

require (
	github.com/AlecAivazis/survey v1.4.1
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/NYTimes/gziphandler v0.0.0-20180227021810-5032c8878b9d
	github.com/atlassian/gostatsd v0.0.0-20191116233337-abc015413f83
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/dave/jennifer v0.0.0-20171207062344-d8bdbdbee4e1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/docker v0.0.0-20180409082103-cbde00b44273
	github.com/dsnet/compress v0.0.0-20170928175515-f41072d47fff // indirect
	github.com/echlebek/crock v1.0.1
	github.com/echlebek/timeproxy v1.0.0
	github.com/emicklei/proto v1.1.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-resty/resty v0.0.0-20170925192930-9ac9c42358f7
	github.com/gogo/protobuf v1.3.1
	github.com/golang/groupcache v0.0.0-20191002201903-404acd9df4cc // indirect
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.0
	github.com/gorilla/websocket v1.4.1
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/graph-gophers/dataloader v0.0.0-20180104184831-78139374585c
	github.com/graphql-go/graphql v0.7.9-0.20190403165646-199d20bbfed7
	github.com/grpc-ecosystem/grpc-gateway v1.11.3 // indirect
	github.com/hashicorp/go-version v1.2.0
	github.com/json-iterator/go v1.1.7
	github.com/mattn/go-runewidth v0.0.2 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
	github.com/mholt/archiver v0.0.0-20180816053333-85d3d0b511ea
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/nwaples/rardecode v0.0.0-20170112110516-f22b7ef81a0a // indirect
	github.com/olekukonko/tablewriter v0.0.0-20180506121414-d4647c9c7a84
	github.com/pierrec/lz4 v0.0.0-20171218195038-2fcda4cb7018 // indirect
	github.com/pierrec/xxHash v0.1.1 // indirect
	github.com/prometheus/client_golang v1.2.0
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/robfig/cron/v3 v3.0.0
	github.com/sensu/lasr v1.2.1
	github.com/shirou/gopsutil v2.19.10+incompatible
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0
	github.com/ulikunitz/xz v0.5.4 // indirect
	github.com/willf/pad v0.0.0-20160331131008-b3d780601022
	go.etcd.io/bbolt v1.3.4-0.20191122203157-7f8bb47fcaf8
	go.uber.org/multierr v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20191014212845-da9a3fd4c582
	golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03 // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/AlecAivazis/survey.v1 v1.4.0 // indirect
	gopkg.in/h2non/filetype.v1 v1.0.3
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.2.4
	gotest.tools v2.2.0+incompatible // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace github.com/spf13/afero => github.com/sensu/afero v1.1.3-0.20191126053333-07cad5c1a1bd

replace github.com/fsnotify/fsnotify => github.com/sensu/fsnotify v1.4.8-0.20191126053121-0adce4777482
