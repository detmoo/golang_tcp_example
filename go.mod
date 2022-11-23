module bowdata.test.go_tcp_echo

go 1.17

replace bowdata.test.go_module_template/pkg => ../pkg

replace bowdata.test.go_module_template/cmd => ../cmd

require (
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.6.1
)

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.2.0 // indirect
)
