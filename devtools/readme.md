```go
r := gin.Default()
fs, err := devtools.FS() 
if err != nil {
panic(err)
}
r.StaticFS("/devtools", http.FS(fs))
```