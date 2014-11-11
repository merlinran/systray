Systray (Trayicon/Menu Extras)
=======


## Cross platform systray for golang

Instead of gui program, "go-program + systray + web-console" might be a interesting choise.


# Platform

Mac: avalid  
Win: avalid  
Linux: coming soon  


## Run example

Mac:
```sh
cd $GOPATH
go get github.com/merlinran/systray
cd $GOPATH/src/github.com/merlinran/systray
xcodebuild -project $GOPATH/src/github.com/merlinran/systray/darwin/systray/systray.xcodeproj CONFIGURATION_BUILD_DIR="$GOPATH/src/github.com/merlinran/systray/example/systray"
go run example/main.go example/icons/mac/ example/systray/
```

## How it works

Win:  
    [your code in go] -> [systray: win32 api call in go]

Mac:  
    [your code in go] -> [systray.Server in go] -(tcp)-> [systray.Client in objc]

Linux:  
    [your code in go] -> [systray.Server in go] -(tcp)-> [systray.Client in c]


