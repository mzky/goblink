- 这是miniblink的go封装，没用cgo，直接调用win32API，demo文件夹是测试启动程序
- 嵌入upx压缩后的dll和devtools（均为20220405版，通过构建时添加参数`-tags=debug`嵌入）
- vip版支持web页面正常下载文件，标准版不支持 （https://github.com/weolar/miniblink49/issues/430）



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