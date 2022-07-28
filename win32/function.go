package win32

import (
	"github.com/mzky/goblink/blink"
	"github.com/mzky/win"
	"golang.org/x/sys/windows"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

type Blink struct {
	_dll *windows.LazyDLL

	_wkeInitialize                     *windows.LazyProc
	_wkeUnInitialize                   *windows.LazyProc
	_wkeCreateWebView                  *windows.LazyProc
	_wkeSetHandle                      *windows.LazyProc
	_wkeOnPaintBitUpdated              *windows.LazyProc
	_wkeOnPaintUpdated                 *windows.LazyProc
	_wkeLoadURL                        *windows.LazyProc
	_wkeGetHostHWND                    *windows.LazyProc
	_wkeResize                         *windows.LazyProc
	_wkeNetOnResponse                  *windows.LazyProc
	_wkeOnLoadUrlBegin                 *windows.LazyProc
	_wkeFireMouseEvent                 *windows.LazyProc
	_wkeFireContextMenuEvent           *windows.LazyProc
	_wkeFireWindowsMessage             *windows.LazyProc
	_wkeCreateWebWindow                *windows.LazyProc
	_wkeShowWindow                     *windows.LazyProc
	_wkeFireMouseWheelEvent            *windows.LazyProc
	_wkeFireKeyUpEvent                 *windows.LazyProc
	_wkeFireKeyDownEvent               *windows.LazyProc
	_wkeFireKeyPressEvent              *windows.LazyProc
	_wkeSetFocus                       *windows.LazyProc
	_wkeNetGetRequestMethod            *windows.LazyProc
	_wkeNetSetData                     *windows.LazyProc
	_wkeNetCancelRequest               *windows.LazyProc
	_wkeDestroyWebView                 *windows.LazyProc
	_jsGetWebView                      *windows.LazyProc
	_wkeKillFocus                      *windows.LazyProc
	_wkeOnDidCreateScriptContext       *windows.LazyProc
	_wkeIsMainFrame                    *windows.LazyProc
	_wkeGetString                      *windows.LazyProc
	_wkeNetSetHTTPHeaderField          *windows.LazyProc
	_wkeNetChangeRequestUrl            *windows.LazyProc
	_wkeNetHookRequest                 *windows.LazyProc
	_wkeNetHoldJobToAsynCommit         *windows.LazyProc
	_wkeNetContinueJob                 *windows.LazyProc
	_wkeOnLoadUrlEnd                   *windows.LazyProc
	_wkeOnConsole                      *windows.LazyProc
	_wkeOnLoadUrlFail                  *windows.LazyProc
	_wkeOnJsQuery                      *windows.LazyProc
	_wkeOnDocumentReady2               *windows.LazyProc
	_wkeOnDownload                     *windows.LazyProc
	_wkeOnAlertBox                     *windows.LazyProc
	_wkeOnCreateView                   *windows.LazyProc
	_wkeSetContextMenuEnabled          *windows.LazyProc
	_wkeResponseQuery                  *windows.LazyProc
	_wkeNetGetMIMEType                 *windows.LazyProc
	_wkeNetSetMIMEType                 *windows.LazyProc
	_wkeNetGetRawResponseHead          *windows.LazyProc
	_wkeSetTransparent                 *windows.LazyProc
	_wkeSetViewProxy                   *windows.LazyProc
	_wkeSetNavigationToNewWindowEnable *windows.LazyProc
	_wkeSetUserAgent                   *windows.LazyProc
	_wkeSetDebugConfig                 *windows.LazyProc
	_wkePopupDialogAndDownload         *windows.LazyProc
	_wkeGetLockedViewDC                *windows.LazyProc
	_wkeRunMessageLoop                 *windows.LazyProc
	_wkeWebFrameGetMainFrame           *windows.LazyProc
	_wkeRunJs                          *windows.LazyProc
	_wkeOnLoadingFinish                *windows.LazyProc
	_wkeEnableHighDPISupport           *windows.LazyProc
}

var (
	DllName = "node"
	// TempPath 临时目录,用于存放临时文件如:dll,cookie等
	TempPath = filepath.Join(os.TempDir(), "blink")
)

func (b *Blink) Init() *Blink {
	tmpFile, err := ioutil.TempFile(TempPath, DllName+"*.dll")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}
	dllPath := tmpFile.Name()
	ioutil.WriteFile(dllPath, blink.Node, 755)
	tmpFile.Close() // 必须放前面
	lib := windows.NewLazyDLL(dllPath)

	b._wkeSetViewProxy = lib.NewProc("mbSetViewProxy")
	b._wkeSetTransparent = lib.NewProc("mbSetTransparent")
	b._wkeOnDocumentReady2 = lib.NewProc("mbOnDocumentReady")
	b._wkeNetGetRawResponseHead = lib.NewProc("mbNetGetRawResponseHeadInBlinkThread")
	b._wkeNetSetMIMEType = lib.NewProc("mbNetSetMIMEType")
	b._wkeNetGetMIMEType = lib.NewProc("mbNetGetMIMEType")
	b._wkeOnLoadUrlFail = lib.NewProc("mbOnLoadUrlFail")
	b._wkeOnLoadUrlEnd = lib.NewProc("mbOnLoadUrlEnd")
	b._wkeNetContinueJob = lib.NewProc("mbNetContinueJob")
	b._wkeNetHoldJobToAsynCommit = lib.NewProc("mbNetHoldJobToAsynCommit")
	b._wkeNetHookRequest = lib.NewProc("mbNetHookRequest")
	b._wkeNetChangeRequestUrl = lib.NewProc("mbNetChangeRequestUrl")
	b._wkeNetSetHTTPHeaderField = lib.NewProc("mbNetSetHTTPHeaderField")
	b._wkeGetString = lib.NewProc("mbGetString")
	b._wkeOnConsole = lib.NewProc("mbOnConsole")
	b._wkeIsMainFrame = lib.NewProc("mbIsMainFrame")
	b._wkeOnDidCreateScriptContext = lib.NewProc("mbOnDidCreateScriptContext")
	b._wkeKillFocus = lib.NewProc("mbKillFocus")
	b._wkeNetCancelRequest = lib.NewProc("mbNetCancelRequest")
	b._wkeNetSetData = lib.NewProc("mbNetSetData")
	b._wkeNetGetRequestMethod = lib.NewProc("mbNetGetRequestMethod")
	b._wkeFireKeyPressEvent = lib.NewProc("mbFireKeyPressEvent")
	b._wkeFireKeyUpEvent = lib.NewProc("mbFireKeyUpEvent")
	b._wkeFireKeyDownEvent = lib.NewProc("mbFireKeyDownEvent")
	b._wkeFireMouseWheelEvent = lib.NewProc("mbFireMouseWheelEvent")
	b._wkeFireContextMenuEvent = lib.NewProc("mbFireContextMenuEvent")
	b._wkeFireWindowsMessage = lib.NewProc("mbFireWindowsMessage")
	b._wkeCreateWebWindow = lib.NewProc("mbCreateWebWindow")
	b._wkeShowWindow = lib.NewProc("mbShowWindow")
	b._wkeFireMouseEvent = lib.NewProc("mbFireMouseEvent")
	b._wkeOnLoadUrlBegin = lib.NewProc("mbOnLoadUrlBegin")
	b._wkeNetOnResponse = lib.NewProc("mbNetOnResponse")
	b._wkeLoadURL = lib.NewProc("mbLoadURL")
	b._wkeGetHostHWND = lib.NewProc("mbGetHostHWND")
	b._wkeResize = lib.NewProc("mbResize")
	b._wkeOnPaintBitUpdated = lib.NewProc("mbOnPaintBitUpdated")
	b._wkeOnPaintUpdated = lib.NewProc("mbOnPaintUpdated")
	b._wkeSetHandle = lib.NewProc("mbSetHandle")
	b._wkeCreateWebView = lib.NewProc("mbCreateWebView")
	b._wkeInitialize = lib.NewProc("mbInit")
	b._wkeUnInitialize = lib.NewProc("mbUninit")
	b._wkeSetFocus = lib.NewProc("mbSetFocus")
	b._wkeDestroyWebView = lib.NewProc("mbDestroyWebView")
	b._jsGetWebView = lib.NewProc("jsGetWebView")
	b._wkeOnDownload = lib.NewProc("mbOnDownloadInBlinkThread")
	b._wkeOnAlertBox = lib.NewProc("mbOnAlertBox")
	b._wkeOnCreateView = lib.NewProc("mbOnCreateView")
	b._wkeSetContextMenuEnabled = lib.NewProc("mbSetContextMenuEnabled")
	b._wkeSetNavigationToNewWindowEnable = lib.NewProc("mbSetNavigationToNewWindowEnable")
	b._wkeSetUserAgent = lib.NewProc("mbSetUserAgent")
	b._wkePopupDialogAndDownload = lib.NewProc("mbPopupDialogAndDownload")
	b._wkeSetDebugConfig = lib.NewProc("mbSetDebugConfig")
	b._wkeOnJsQuery = lib.NewProc("mbOnJsQuery")
	b._wkeResponseQuery = lib.NewProc("mbResponseQuery")
	b._wkeGetLockedViewDC = lib.NewProc("mbGetLockedViewDC")
	b._wkeRunMessageLoop = lib.NewProc("mbRunMessageLoop")
	b._wkeWebFrameGetMainFrame = lib.NewProc("mbWebFrameGetMainFrame")
	b._wkeRunJs = lib.NewProc("mbRunJs")
	b._wkeOnLoadingFinish = lib.NewProc("mbOnLoadingFinish")
	b._wkeEnableHighDPISupport = lib.NewProc("mbEnableHighDPISupport")
	var set MbSettings
	set.Mask = MB_ENABLE_NODEJS
	r, _, err := b._wkeInitialize.Call(uintptr(unsafe.Pointer(&set)))
	if r != 0 {
		return b
	}
	return b
}

func GetBound(h win.HWND) win.RECT {
	rect := win.RECT{}
	win.GetWindowRect(h, &rect)
	bn := win.RECT{
		Left: rect.Left,
		Top:  rect.Top,
	}
	win.GetClientRect(h, &rect)

	bn.Right = rect.Width() + bn.Left
	bn.Bottom = rect.Height() + bn.Top
	return bn
}

func (b *Blink) WkeUnInit() {
	b._wkeUnInitialize.Call()
}

func (b *Blink) WkeOnDownload(wke WkeHandle, callback WkeOnDownloadCallback, param uintptr) {
	b._wkeOnDownload.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) WkeOnAlertBox(wke WkeHandle, callback WkeOnAlertBoxCallback, param uintptr) {
	b._wkeOnAlertBox.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeOnCreateView(wke WkeHandle, callback WkeOnCreateViewCallback, param uintptr) {
	b._wkeOnCreateView.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeSetContextMenuEnabled(wke WkeHandle, show bool) {
	b._wkeSetContextMenuEnabled.Call(uintptr(wke), uintptr(toBool(show)))
}

func (b *Blink) WkeSetNavigationToNewWindowEnable(wke WkeHandle, bl bool) {
	b._wkeSetNavigationToNewWindowEnable.Call(uintptr(wke), uintptr(toBool(bl)))
}

func (b *Blink) WkeSetUserAgent(wke WkeHandle, userAgent string) {
	p := StrToCharPtr(userAgent)
	b._wkeSetUserAgent.Call(uintptr(wke), p)
}

func (b *Blink) wkeSetViewProxy(wke WkeHandle, proxy ProxyInfo) {
	px := WkeProxy{
		Type: int32(proxy.Type),
		Port: uint16(proxy.Port),
	}
	for i, c := range proxy.HostName {
		px.HostName[i] = byte(c)
	}
	if proxy.UserName != "" {
		for i, c := range proxy.UserName {
			px.UserName[i] = byte(c)
		}
	}
	if proxy.Password != "" {
		for i, c := range proxy.Password {
			px.Password[i] = byte(c)
		}
	}
	b._wkeSetViewProxy.Call(uintptr(wke), uintptr(unsafe.Pointer(&px)))
}

func (b *Blink) WkeSetTransparent(wke WkeHandle, enable bool) {
	b._wkeSetTransparent.Call(uintptr(wke), uintptr(toBool(enable)))
}

func (b *Blink) WkeOnDocumentReady(wke WkeHandle, callback WkeDocumentReady2Callback, param uintptr) {
	b._wkeOnDocumentReady2.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeNetGetRawResponseHead(job WkeNetJob) map[string]string {
	r, _, _ := b._wkeNetGetRawResponseHead.Call(uintptr(job))
	var list []string
	sList := *((*WkeSlist)(unsafe.Pointer(r)))
	for sList.Str != 0 {
		list = append(list, PtrToUtf8(sList.Str))
		if sList.Next == 0 {
			break
		} else {
			sList = *((*WkeSlist)(unsafe.Pointer(sList.Next)))
		}
	}
	hMap := make(map[string]string)
	for i := 0; i < len(list); i += 2 {
		hMap[list[i]] = list[i+1]
	}
	return hMap
}

func (b *Blink) wkeNetSetMIMEType(job WkeNetJob, mime string) {
	p := StrToCharPtr(mime)
	b._wkeNetSetMIMEType.Call(uintptr(job), p)
}

func (b *Blink) wkeNetGetMIMEType(job WkeNetJob) string {
	r, _, _ := b._wkeNetGetMIMEType.Call(uintptr(job))
	return PtrToUtf8(r)
}

func (b *Blink) wkeOnLoadUrlFail(wke WkeHandle, callback WkeLoadUrlFailCallback, param uintptr) {
	b._wkeOnLoadUrlFail.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeOnLoadUrlEnd(wke WkeHandle, callback WkeLoadUrlEndCallback, param uintptr) {
	b._wkeOnLoadUrlEnd.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeNetContinueJob(job WkeNetJob) {
	b._wkeNetContinueJob.Call(uintptr(job))
}

func (b *Blink) wkeNetHoldJobToAsynCommit(job WkeNetJob) {
	b._wkeNetHoldJobToAsynCommit.Call(uintptr(job))
}

func (b *Blink) wkeNetHookRequest(job WkeNetJob) {
	b._wkeNetHookRequest.Call(uintptr(job))
}

func (b *Blink) wkeNetChangeRequestUrl(job WkeNetJob, url string) {
	p := StrToCharPtr(url)
	b._wkeNetChangeRequestUrl.Call(uintptr(job), p)
}

func (b *Blink) wkeNetSetHTTPHeaderField(job WkeNetJob, name, value string) {
	np := StrToCharPtr(name)
	vp := StrToCharPtr(value)
	b._wkeNetSetHTTPHeaderField.Call(uintptr(job), np, vp)
}

func (b *Blink) wkeGetString(str WkeString) string {
	r, _, _ := b._wkeGetString.Call(uintptr(str))
	return PtrToUtf8(r)
}

func (b *Blink) wkeOnConsole(wke WkeHandle, callback WkeConsoleCallback, param uintptr) {
	b._wkeOnConsole.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeIsMainFrame(wke WkeHandle, frame WkeFrame) bool {
	r, _, _ := b._wkeIsMainFrame.Call(uintptr(wke), uintptr(frame))
	return r != 0
}

func (b *Blink) wkeOnDidCreateScriptContext(wke WkeHandle, callback WkeDidCreateScriptContextCallback, param uintptr) {
	b._wkeOnDidCreateScriptContext.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeKillFocus(wke WkeHandle) {
	b._wkeKillFocus.Call(uintptr(wke))
}

func (b *Blink) jsGetWebView(es JsExecState) WkeHandle {
	r, _, _ := b._jsGetWebView.Call(uintptr(es))
	return WkeHandle(r)
}

func (b *Blink) WkeDestroyWebView(wke WkeHandle) {
	b._wkeDestroyWebView.Call(uintptr(wke))
}

func (b *Blink) wkeNetCancelRequest(job WkeNetJob) {
	b._wkeNetCancelRequest.Call(uintptr(job))
}

func (b *Blink) wkeNetOnResponse(wke WkeHandle, callback WkeNetResponseCallback, param uintptr) {
	b._wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeOnLoadUrlBegin(wke WkeHandle, callback WkeLoadUrlBeginCallback, param uintptr) {
	b._wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (b *Blink) wkeNetGetRequestMethod(job WkeNetJob) WkeRequestType {
	r, _, _ := b._wkeNetGetRequestMethod.Call(uintptr(job))
	return WkeRequestType(r)
}

func (b *Blink) wkeNetSetData(job WkeNetJob, buf []byte) {
	if len(buf) == 0 {
		buf = []byte{0}
	}
	b._wkeNetSetData.Call(uintptr(job), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func (b *Blink) wkeSetFocus(wke WkeHandle) {
	b._wkeSetFocus.Call(uintptr(wke))
}

func (b *Blink) wkeFireKeyPressEvent(wke WkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := b._wkeFireKeyPressEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (b *Blink) wkeFireKeyDownEvent(wke WkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := b._wkeFireKeyDownEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (b *Blink) wkeFireKeyUpEvent(wke WkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := b._wkeFireKeyUpEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (b *Blink) wkeFireMouseWheelEvent(wke WkeHandle, x, y, delta, flags int32) bool {
	r, _, _ := b._wkeFireMouseWheelEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(delta),
		uintptr(flags))
	return byte(r) != 0
}
func (b *Blink) wkeFireContextMenuEvent(wke WkeHandle, x, y, flags int32) bool {
	r, _, _ := b._wkeFireContextMenuEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}
func (b *Blink) wkeFireWindowsMessage(wke WkeHandle, hWnd win.HWND, message, wParam, lParam int32) bool {
	r, _, _ := b._wkeFireWindowsMessage.Call(
		uintptr(wke),
		uintptr(hWnd),
		uintptr(message),
		uintptr(wParam),
		uintptr(lParam),
		uintptr(0))
	return byte(r) != 0
}

func (b *Blink) wkeCreateWebWindow(wt WindowType, parent win.HWND, x, y, width, height int32) WkeHandle {
	r, _, _ := b._wkeCreateWebWindow.Call(
		uintptr(wt),
		uintptr(parent),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return WkeHandle(r)
}

func (b *Blink) wkeShowWindow(wke WkeHandle, show bool) {
	b._wkeShowWindow.Call(uintptr(wke), uintptr(toBool(show)))
}

func (b *Blink) wkeFireMouseEvent(wke WkeHandle, message, x, y, flags int32) bool {
	r, _, _ := b._wkeFireMouseEvent.Call(
		uintptr(wke),
		uintptr(message),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}

func (b *Blink) wkeResize(wke WkeHandle, w, h uint32) {
	b._wkeResize.Call(uintptr(wke), uintptr(w), uintptr(h))
}

func (b *Blink) wkeLoadURL(wke WkeHandle, url string) {
	ptr := StrToCharPtr(url)
	b._wkeLoadURL.Call(uintptr(wke), ptr)
}

/*
设置一些实验性选项。debugString可用参数有：
*/
func (b *Blink) WkeSetDebugConfig(wke WkeHandle, debug DebugType, param string) {
	dp := StrToCharPtr(string(debug))
	pp := StrToCharPtr(param)
	b._wkeSetDebugConfig.Call(uintptr(wke), dp, pp)
}

func (b *Blink) wkeOnPaintBitUpdated(wke WkeHandle, callback WkePaintBitUpdatedCallback, param uintptr) {
	b._wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (b *Blink) WkeOnPaintUpdated(wke WkeHandle, callback WkePaintUpdatedCallback, param uintptr) {
	b._wkeOnPaintUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (b *Blink) wkeOnLoadingFinish(wke WkeHandle, callback WkeLoadingFinishCallback, param uintptr) {
	b._wkeOnLoadingFinish.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (b *Blink) wkeEnableHighDPISupport() {
	b._wkeEnableHighDPISupport.Call()
}

func (b *Blink) wkeRunJs(handle WkeHandle, frame WkeFrame, script uintptr, isInClosure bool, param, unUse uintptr) {
	b._wkeRunJs.Call(uintptr(handle), uintptr(frame), script, uintptr(toBool(isInClosure)), 0, param, unUse)
}

func (b *Blink) WkeSetHandle(wke WkeHandle, handle uintptr) {
	b._wkeSetHandle.Call(uintptr(wke), handle)
}
func (b *Blink) wkeOnShowDevtoolsCallback(wke uintptr, param uintptr) uintptr {
	return 0
}
func (b *Blink) WkeCreateWebView() WkeHandle {
	r, _, _ := b._wkeCreateWebView.Call()
	return WkeHandle(r)
}
func (b *Blink) wkeGetHostHWND() win.HWND {
	r, _, _ := b._wkeGetHostHWND.Call()
	return win.HWND(r)
}

func (b *Blink) wkeGetLockedViewDC(handle WkeHandle) win.HDC {
	r, _, _ := b._wkeGetLockedViewDC.Call(uintptr(handle))
	return win.HDC(r)
}
func (b *Blink) wkeRunMessageLoop() {
	b._wkeRunMessageLoop.Call()
}

func (b *Blink) wkeWebFrameGetMainFrame(handle WkeHandle) WkeFrame {
	r, _, _ := b._wkeWebFrameGetMainFrame.Call(uintptr(handle))
	return WkeFrame(r)
}

func StrToCharPtr(str string) uintptr {
	buf := []byte(str)
	rs := make([]byte, len(str)+1)
	for i, v := range buf {
		rs[i] = v
	}
	return uintptr(unsafe.Pointer(&rs[0]))
}

func StrPtr(s string) (uintptr, error) {
	fromString, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return 0, err
	}
	return uintptr(unsafe.Pointer(fromString)), nil
}

func PtrToUtf8(ptr uintptr) string {
	var seq []byte
	for ptr > 0 {
		b := *((*byte)(unsafe.Pointer(ptr)))
		if b != 0 {
			seq = append(seq, b)
			ptr++
		} else {
			break
		}
	}
	return string(seq)
}

func toBool(b bool) byte {
	if b {
		return 1
	} else {
		return 0
	}
}
