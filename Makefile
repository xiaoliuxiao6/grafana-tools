go_files = main.go backupDashboards.go

BINS :=

build:
	rm -f bin/grafana-tools
	GO11MODULE=on GO111MODULE=on GOPROXY=https://goproxy.io \
	go build -o bin/grafana-tools $(go_files)
BINS+=bin/grafana-tools

all:
	echo "Compiling for every OS and Platform"
	rm -f bin/grafana-tools-linux-arm bin/grafana-tools-linux-arm64 bin/grafana-tools-linux-386 bin/grafana-tools-linux-amd64
	GOOS=linux GOARCH=arm go build -o bin/grafana-tools-linux-arm $(go_files)
	GOOS=linux GOARCH=arm64 go build -o bin/grafana-tools-linux-arm64 $(go_files)
	GOOS=freebsd GOARCH=386 go build -o bin/grafana-tools-linux-386 $(go_files)
	GOOS=freebsd GOARCH=amd64 go build -o bin/grafana-tools-linux-amd64 $(go_files)
BINS+=bin/grafana-tools-linux-arm
BINS+=bin/grafana-tools-linux-arm64
BINS+=bin/grafana-tools-linux-386
BINS+=bin/grafana-tools-linux-amd64

clean:
	rm -rf ./bin
