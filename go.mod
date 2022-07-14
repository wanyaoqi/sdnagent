module yunion.io/x/sdnagent

go 1.18

require (
	github.com/coreos/go-iptables v0.6.0
	github.com/digitalocean/go-openvswitch v0.0.0-20190515160856-1141932ed5cf
	github.com/fsnotify/fsnotify v1.4.9
	github.com/golang/protobuf v1.5.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.3.2
	github.com/vishvananda/netlink v1.0.0
	github.com/vishvananda/netns v0.0.0-20211101163701-50045581ed74
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
	yunion.io/x/jsonutils v1.0.0
	yunion.io/x/log v1.0.0
	yunion.io/x/onecloud v0.0.0-20220714042151-0ec0a182c004
	yunion.io/x/pkg v1.0.1-0.20220630095420-9925accd7c5e
)

replace (
	github.com/digitalocean/go-openvswitch => github.com/yousong/go-openvswitch v0.0.0-20200422025222-6b2d502be872
	github.com/go-logr/logr => github.com/go-logr/logr v0.4.0
	github.com/jaypipes/ghw => github.com/zexi/ghw v0.9.1
)
