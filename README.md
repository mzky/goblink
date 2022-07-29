- 这是miniblink的go封装，没用cgo，直接调用win32API，demo文件夹是测试启动程序
- 嵌入upx压缩后的dll和devtools（均为20220405版，通过构建时添加参数`-tags=debug`嵌入）
- vip版支持web页面正常下载文件，标准版不支持 （https://github.com/weolar/miniblink49/issues/430）
- 32位版很小，编译并压缩后exe小于10M，兼容64位
- 此库不建议正式用，玩玩可以，正式用推荐用其它库，比如
```
https://github.com/mzky/blink
https://github.com/suiyunonghen/GVCL/tree/master/Components/DxControls/gminiblink
https://gitee.com/aochulai/GoMiniblink
```

```go
func main() {
	ico, _ := res.ReadFile("res/gear22.ico")
	win32.LoadIconFromBytes(ico)
	err := goblink.NewWebBrowser("https://www.baidu.com", windowTitle, devtoolsPath, false, true, true, 1280, 800)
	if err != nil {
		log.Println(*url, err.Error())
	}
}
```


```bash
gox -osarch="windows/amd64" -ldflags "-w -s -H=windowsgui"
gox -osarch="windows/386" -ldflags "-w -s -H=windowsgui" -tags="debug"
go build -ldflags "-w -s -H=windowsgui" -tags="debug"
```

