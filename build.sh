rm -f ./bin/*
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o ./bin/wifi_login_darwin_amd64
env CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build  -o ./bin/wifi_login_freebsd_386
env CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build  -o ./bin/wifi_login_freebsd_amd64
env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build  -o ./bin/wifi_login_linux_386
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ./bin/wifi_login_linux_amd64
env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build  -o ./bin/wifi_login_linux_arm
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build  -o ./bin/wifi_login_linux_arm64
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./bin/wifi_login_windows_amd64.exe
env CGO_ENABLED=0 GOOS=windows GOARCH=386 go build  -o ./bin/wifi_login_windows_386.exe
env CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build  -o ./bin/wifi_login_linux_mips64
env CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build  -o ./bin/wifi_login_linux_mips64le
env CGO_ENABLED=0 GOOS=linux GOARCH=mips GOMIPS=softfloat go build  -o ./bin/wifi_login_linux_mips
env CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build  -o ./bin/wifi_login_linux_mipsle

for filename in ./bin/*; do
    echo $filename
    tar -czvf  ${filename}.tar.gz  ${filename} config-example.json
    rm ${filename}
done