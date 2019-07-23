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

cp config-example.json ./bin/config.json
cd ./bin
for filename in ./*; do
    echo $filename
    if [ $filename = "./config.json" ];then
        continue
    fi
    filename2=$(echo $filename | sed "s/.exe//")
    tar -czvf  ${filename2}.tar.gz  ${filename} config.json
    rm ${filename}
done