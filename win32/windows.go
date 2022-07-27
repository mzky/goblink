package win32

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/mzky/win"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"
	"unsafe"
)

const (
	className      = "thublink_class"
	windowName     = "thublink_window"
	classViewName  = "thublink_view_class"
	windowViewName = "thublink_view_window"
)

var (
	classNamePtr      *uint16
	windowNamePtr     *uint16
	classViewNamePtr  *uint16
	windowViewNamePtr *uint16
	hInst             win.HINSTANCE
	procMap           map[win.HWND]uintptr
	WebView           *Blink
	iconHandle        win.HANDLE
	urls              []string
)

type SaveCallback func(url, path string)
type FinishCallback func(url string, success bool)

func init() {
	WebView = new(Blink).Init()
	var err error
	classNamePtr, err = syscall.UTF16PtrFromString(className)
	if err != nil {
		fmt.Println(err)
		return
	}
	windowNamePtr, err = syscall.UTF16PtrFromString(windowName)
	if err != nil {
		fmt.Println(err)
		return
	}
	classViewNamePtr, err = syscall.UTF16PtrFromString(classViewName)
	if err != nil {
		fmt.Println(err)
		return
	}
	windowViewNamePtr, err = syscall.UTF16PtrFromString(windowViewName)
	if err != nil {
		fmt.Println(err)
		return
	}
	hInst = win.GetModuleHandle(nil)
	wndClass := win.WNDCLASSEX{
		Style:         win.CS_HREDRAW | win.CS_VREDRAW,
		LpfnWndProc:   syscall.NewCallbackCDecl(classMsgProc),
		HInstance:     hInst,
		LpszClassName: classNamePtr,
		HCursor:       win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW)),
		HbrBackground: win.GetSysColorBrush(win.COLOR_WINDOW + 1),
	}
	wndClass.CbSize = uint32(unsafe.Sizeof(wndClass))
	win.RegisterClassEx(&wndClass)
	wndClass = win.WNDCLASSEX{
		Style:         win.CS_DBLCLKS,
		LpfnWndProc:   syscall.NewCallbackCDecl(classMsgProc),
		HInstance:     hInst,
		LpszClassName: classViewNamePtr,
		HbrBackground: win.GetSysColorBrush(win.COLOR_WINDOW + 1),
		HCursor:       win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW)),
	}
	wndClass.CbSize = uint32(unsafe.Sizeof(wndClass))
	win.RegisterClassEx(&wndClass)
	procMap = make(map[win.HWND]uintptr)
}
func classMsgProc(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if v, e := procMap[hWnd]; e {
		return win.CallWindowProc(v, hWnd, msg, wParam, lParam)
	}
	return win.DefWindowProc(hWnd, msg, wParam, lParam)
}
func newWindow(exStyle, style uint32, parent win.HWND, width, height int32, proc func(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr) win.HWND {
	return newClassWindow(exStyle, style, parent, width, height, classNamePtr, windowNamePtr, proc)
}
func newClassWindow(exStyle, style uint32, parent win.HWND, width, height int32, className, windowName *uint16,
	proc func(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr) win.HWND {
	var x, y int32
	if parent == 0 && style&win.WS_MAXIMIZE == 0 { // 居中
		sw := win.GetSystemMetrics(win.SM_CXFULLSCREEN)
		sh := win.GetSystemMetrics(win.SM_CYFULLSCREEN)
		x = (sw - width) / 2
		y = (sh - height) / 2
	}
	wnd := win.CreateWindowEx(exStyle, className, windowName, style, x, y, width, height,
		parent, 0, hInst, unsafe.Pointer(nil))
	if wnd != 0 {
		procMap[wnd] = syscall.NewCallbackCDecl(proc)
	}
	return wnd
}

type FormProfile struct {
	Title        string
	UserAgent    string
	Width        int
	Height       int
	Max          bool
	Mb           bool
	Ib           bool
	Index        string
	DevtoolsPath string
	Subs         map[string]FormProfile
	Main         bool
}

func StartBlinkMain(url, title, devtoolsPath string, max, mb, ib bool, width, height int) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	main := FormProfile{Title: title, Index: url, DevtoolsPath: devtoolsPath, Max: max, Mb: mb, Ib: ib, Width: width, Height: height}
	main.NewBlinkWindow()
	// 3. 主消息循环
	msg := (*win.MSG)(unsafe.Pointer(win.GlobalAlloc(0, unsafe.Sizeof(win.MSG{}))))
	defer win.GlobalFree(win.HGLOBAL(unsafe.Pointer(msg)))
	for win.GetMessage(msg, 0, 0, 0) > 0 {
		// fmt.Println(msg.Message, msg.HWnd, msg.LParam, msg.WParam)
		if msg.Message == win.WM_QUIT {
			WebView.wkeUnInit()
			break
		}
		win.TranslateMessage(msg)
		win.DispatchMessage(msg)
	}
	return nil
}

func (fp FormProfile) NewBlinkWindow() {
	w := window{profile: fp}
	w.init()
	v := BlinkView{}
	var r win.RECT
	win.GetClientRect(w.hWnd, &r)
	v.init(fp.UserAgent, fp.DevtoolsPath)
	v.SetOnNewWindow(w.OnCreateView)
	v.setDownloadCallback(w.WkeOnDownloadCallback)
	w.child = newClassWindow(0, win.WS_CHILD|win.WS_VISIBLE|win.WS_CLIPSIBLINGS|win.WS_CLIPCHILDREN, w.hWnd, r.Width(), r.Height(), classViewNamePtr, windowViewNamePtr, v.OnWndProc)
	v.setHWnd(w.child)
	v.resize(r.Width(), r.Height(), true)
	v.LoadUrl(fp.Index)
	WebView.wkeOnLoadUrlBegin(v.handle, v.WkeLoadUrlBeginCallback, 0)
	urls = append(urls, fp.Index)
	w.view = &v
}

func loadIcon(ico string) {
	hInst := win.GetModuleHandle(nil)
	fromString, err := syscall.UTF16PtrFromString(ico)
	if err != nil {
		fmt.Println(err)
		return
	}
	iconHandle = win.LoadImage(hInst, fromString, win.IMAGE_ICON, 0, 0, win.LR_LOADFROMFILE)
}

// LoadIconFromBytes 先把ico二进制数据存到本地(common.TempPath),再使用winapi的LoadImage加载图标
func LoadIconFromBytes(iconData []byte) {
	//计算数据的hash
	bh := md5.Sum(iconData)
	dataHash := hex.EncodeToString(bh[:])

	//缓存中没有,则释放到本地目录
	iconFilePath := filepath.Join(TempPath, "icon_"+dataHash)
	if _, err := os.Stat(iconFilePath); os.IsNotExist(err) {
		if err := ioutil.WriteFile(iconFilePath, iconData, 0644); err != nil {
			log.Println("无法创建临时icon文件: " + err.Error())
		}
	}

	loadIcon(iconFilePath)
}

type window struct {
	hWnd    win.HWND
	child   win.HWND
	profile FormProfile
	view    *BlinkView
	down    map[string]*downInfo
	bind    map[string]*wkeDownloadBind
	mux     sync.Mutex
}

func (w *window) init() {
	w.down = make(map[string]*downInfo)
	w.bind = make(map[string]*wkeDownloadBind)
	w.hWnd = newWindow(0, w.style(), 0, int32(w.profile.Width), int32(w.profile.Height), w.windowMsgProc)
	if w.hWnd == 0 {
		return
	}
	if iconHandle != 0 {
		win.SendMessage(w.hWnd, win.WM_SETICON, 1, uintptr(iconHandle))
	}
	win.SetWindowText(w.hWnd, w.profile.Title)
	win.ShowWindow(w.hWnd, win.SW_SHOW)
}

func (w *window) style() uint32 {
	var style uint32 = win.WS_OVERLAPPEDWINDOW | win.WS_VISIBLE | win.WS_CLIPSIBLINGS | win.WS_CLIPCHILDREN
	if !w.profile.Ib {
		style ^= win.WS_MINIMIZEBOX
	}
	if !w.profile.Mb {
		style ^= win.WS_MAXIMIZEBOX
	}
	if w.profile.Max {
		style |= win.WS_MAXIMIZE
	}
	return style
}

func (w *window) roundRect() { // 有效果，但是很丑，还有bug
	/*
		您可以创建一个没有任何框架的窗口，使用WS_EX_LAYERED获取透明度，然后在WM_PAINT中“正常”绘制包含自定义框架的窗口，或者组成一个离屏位图，并使用
		UpdateLayeredWindow
		（后一种方法更有效） 。 当然，您必须将绘制的内容调整为窗口的当前大小。通常，您可以从不同的元素组成 - 例如使用四个“角落”位图（或椭圆函数）绘制角点，然后绘制边框等。 此外，您可以处理
		WM_NCHITTEST
		将“标题”/“边框”/“角落”功能（即移动和调整窗口大小）分配到窗口的任意区域。
	*/
	var r win.RECT
	win.GetWindowRect(w.hWnd, &r)
	rgn := win.CreateRoundRectRgn(r.Left, r.Top, r.Right, r.Bottom, 20, 20)
	win.SetWindowRgn(w.hWnd, rgn, true)
}

func (w *window) windowMsgProc(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case win.WM_SIZE:
		if w.child > 0 && w.view != nil {
			var r win.RECT
			win.GetClientRect(hWnd, &r)
			win.MoveWindow(w.child, r.Left, r.Top, r.Width(), r.Height(), true)
			w.view.resize(r.Width(), r.Height(), false)
		}
	case win.WM_CLOSE:
		if w.view == nil {
			break
		}
		w.view.close()
		if w.profile.Main {
			win.PostQuitMessage(0)
		}
	}
	return win.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (w *window) WkeOnDownloadCallback(wke WkeHandle, param uintptr, length uint32, url, mime, disposition uintptr, job WkeNetJob, dataBind uintptr) wkeDownloadOpt {
	info := downInfo{}
	urlStr := PtrToUtf8(url)
	info.url = StrToCharPtr(urlStr)
	info.recvSize = 0
	info.allSize = length
	bind := wkeDownloadBind{param: uintptr(unsafe.Pointer(&info))}
	w.mux.Lock()
	defer w.mux.Unlock()
	w.down[urlStr] = &info
	w.bind[urlStr] = &bind
	return w.view.wkePopupDialogAndDownload(param, length, url, mime, disposition, job, dataBind, &bind)
}
func (w *window) OnCreateView(wke WkeHandle, param uintptr, naviType wkeNavigationType, url, feature uintptr) uintptr {
	a := PtrToUtf8(url)
	if Debug() {
		fmt.Println("OnCreateView", a)
	}
	urls = append(urls, a)
	if v, e := w.profile.Subs[a]; e {
		v.NewBlinkWindow()
	} else {
		o := operateUri(a)
		if o == 1 {
			return 0
		}
		n := w.profile
		n.Index = a
		n.Main = false
		n.NewBlinkWindow()
	}
	return 0
}

func Debug() bool {
	return true
}
