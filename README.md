
Use the package
```
go get -u github.com/dieterreuter/osarch
```

Install the CLI program
```
go install github.com/dieterreuter/osarch/cmd/osarch
```

Cross compiling on OSX:
```
git clone https://github.com/dieterreuter/osarch
cd osarch/cmd/osarch
GOOS=linux GOARCH=arm go build .
```

Usage on OSX:
```
./osarch
Arch()= amd64
OsArch()= darwin/amd64
```

Usage on a Raspberry Pi 1:
```
./osarch
Arch()= arm/armv6l
OsArch()= linux/arm/armv6l
```


### Results on different devices:

Device                  |OS             |Docker OS/Arch|Arch()   |OsArch()
--- | --- | --- | --- | ---
MacBook Pro 15" Mid 1014|OS X 10.10.4   |darwin/amd64|amd64      |darwin/amd64
boot2docker@OSX         |Debian Wheezy  |linux/amd64 |amd64      |linux/amd64
boot2docker@OSX         |Debian Jessie  |linux/amd64 |amd64      |linux/amd64
boot2docker@OSX         |Ubuntu 14.04   |linux/amd64 |amd64      |linux/amd64
Raspberry Pi 1 model A  |Raspbian Wheezy|linux/arm   |arm/armv6l |linux/arm/armv6l
Raspberry Pi 1 model A+ |Raspbian Jessie|linux/arm   |arm/armv6l |linux/arm/armv6l
Raspberry Pi 1 model B  |Raspbian Jessie|linux/arm   |arm/armv6l |linux/arm/armv6l
Raspberry Pi 1 model B+ |Raspbian Jessie|linux/arm   |arm/armv6l |linux/arm/armv6l
Raspberry Pi 2 model B  |Raspbian Jessie|linux/arm   |arm/armv7l |linux/arm/armv7l
ODROID-C1               |Ubuntu 14.04   |linux/arm   |arm/armv7l |linux/arm/armv7l
ODROID-XU3              |Ubuntu 14.04   |linux/arm   |arm/armv7l |linux/arm/armv7l
ODROID-XU3 lite         |Ubuntu 14.04   |linux/arm   |arm/armv7l |linux/arm/armv7l
ODROID-XU4              |Ubuntu 14.04   |linux/arm   |arm/armv7l |linux/arm/armv7l
NVIDIA SHIELD Android TV|Android 5.0    |linux/arm64 |arm/aarch64|linux/arm/aarch64
ASUS ZenFone 2 (ZE551ML)|Android 5.0    |linux/amd64 |amd64      |linux/amd64
