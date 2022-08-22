
install:
	go build -o  ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/aws@latest/steampipe-plugin-aws.plugin  *.go


local:
	go build -o  ~/.steampipe/plugins/local/awstest/awstest.plugin *.go

