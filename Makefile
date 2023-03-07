
install:
	go build -o  ~/.steampipe/plugins/zigbang.com/aws-zb/steampipe-plugin-aws-zb.plugin  *.go


local:
	go build -o  ~/.steampipe/plugins/local/aws-zb-local/aws-zb-local.plugin *.go

