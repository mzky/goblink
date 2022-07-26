package win32

import (
	"fmt"
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

func (t *Blink) Init() *Blink {
	tmpFile, err := ioutil.TempFile(os.TempDir(), DllName+"*.dll")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}
	ioutil.WriteFile(tmpFile.Name(), blink.Node, 755)
	fp := tmpFile.Name()
	tmpFile.Close()
	lib := windows.NewLazyDLL(fp)

	t._wkeSetViewProxy = lib.NewProc("mbSetViewProxy")
	t._wkeSetTransparent = lib.NewProc("mbSetTransparent")
	t._wkeOnDocumentReady2 = lib.NewProc("mbOnDocumentReady")
	t._wkeNetGetRawResponseHead = lib.NewProc("mbNetGetRawResponseHeadInBlinkThread")
	t._wkeNetSetMIMEType = lib.NewProc("mbNetSetMIMEType")
	t._wkeNetGetMIMEType = lib.NewProc("mbNetGetMIMEType")
	t._wkeOnLoadUrlFail = lib.NewProc("mbOnLoadUrlFail")
	t._wkeOnLoadUrlEnd = lib.NewProc("mbOnLoadUrlEnd")
	t._wkeNetContinueJob = lib.NewProc("mbNetContinueJob")
	t._wkeNetHoldJobToAsynCommit = lib.NewProc("mbNetHoldJobToAsynCommit")
	t._wkeNetHookRequest = lib.NewProc("mbNetHookRequest")
	t._wkeNetChangeRequestUrl = lib.NewProc("mbNetChangeRequestUrl")
	t._wkeNetSetHTTPHeaderField = lib.NewProc("mbNetSetHTTPHeaderField")
	t._wkeGetString = lib.NewProc("mbGetString")
	t._wkeOnConsole = lib.NewProc("mbOnConsole")
	t._wkeIsMainFrame = lib.NewProc("mbIsMainFrame")
	t._wkeOnDidCreateScriptContext = lib.NewProc("mbOnDidCreateScriptContext")
	t._wkeKillFocus = lib.NewProc("mbKillFocus")
	t._wkeNetCancelRequest = lib.NewProc("mbNetCancelRequest")
	t._wkeNetSetData = lib.NewProc("mbNetSetData")
	t._wkeNetGetRequestMethod = lib.NewProc("mbNetGetRequestMethod")
	t._wkeFireKeyPressEvent = lib.NewProc("mbFireKeyPressEvent")
	t._wkeFireKeyUpEvent = lib.NewProc("mbFireKeyUpEvent")
	t._wkeFireKeyDownEvent = lib.NewProc("mbFireKeyDownEvent")
	t._wkeFireMouseWheelEvent = lib.NewProc("mbFireMouseWheelEvent")
	t._wkeFireContextMenuEvent = lib.NewProc("mbFireContextMenuEvent")
	t._wkeFireWindowsMessage = lib.NewProc("mbFireWindowsMessage")
	t._wkeCreateWebWindow = lib.NewProc("mbCreateWebWindow")
	t._wkeShowWindow = lib.NewProc("mbShowWindow")
	t._wkeFireMouseEvent = lib.NewProc("mbFireMouseEvent")
	t._wkeOnLoadUrlBegin = lib.NewProc("mbOnLoadUrlBegin")
	t._wkeNetOnResponse = lib.NewProc("mbNetOnResponse")
	t._wkeLoadURL = lib.NewProc("mbLoadURL")
	t._wkeGetHostHWND = lib.NewProc("mbGetHostHWND")
	t._wkeResize = lib.NewProc("mbResize")
	t._wkeOnPaintBitUpdated = lib.NewProc("mbOnPaintBitUpdated")
	t._wkeOnPaintUpdated = lib.NewProc("mbOnPaintUpdated")
	t._wkeSetHandle = lib.NewProc("mbSetHandle")
	t._wkeCreateWebView = lib.NewProc("mbCreateWebView")
	t._wkeInitialize = lib.NewProc("mbInit")
	t._wkeUnInitialize = lib.NewProc("mbUninit")
	t._wkeSetFocus = lib.NewProc("mbSetFocus")
	t._wkeDestroyWebView = lib.NewProc("mbDestroyWebView")
	t._jsGetWebView = lib.NewProc("jsGetWebView")
	t._wkeOnDownload = lib.NewProc("mbOnDownloadInBlinkThread")
	t._wkeOnAlertBox = lib.NewProc("mbOnAlertBox")
	t._wkeOnCreateView = lib.NewProc("mbOnCreateView")
	t._wkeSetContextMenuEnabled = lib.NewProc("mbSetContextMenuEnabled")
	t._wkeSetNavigationToNewWindowEnable = lib.NewProc("mbSetNavigationToNewWindowEnable")
	t._wkeSetUserAgent = lib.NewProc("mbSetUserAgent")
	t._wkePopupDialogAndDownload = lib.NewProc("mbPopupDialogAndDownload")
	t._wkeSetDebugConfig = lib.NewProc("mbSetDebugConfig")
	t._wkeOnJsQuery = lib.NewProc("mbOnJsQuery")
	t._wkeResponseQuery = lib.NewProc("mbResponseQuery")
	t._wkeGetLockedViewDC = lib.NewProc("mbGetLockedViewDC")
	t._wkeRunMessageLoop = lib.NewProc("mbRunMessageLoop")
	t._wkeWebFrameGetMainFrame = lib.NewProc("mbWebFrameGetMainFrame")
	t._wkeRunJs = lib.NewProc("mbRunJs")
	t._wkeOnLoadingFinish = lib.NewProc("mbOnLoadingFinish")
	t._wkeEnableHighDPISupport = lib.NewProc("mbEnableHighDPISupport")
	var set MbSettings
	set.Mask = MB_ENABLE_NODEJS
	r, _, err := t._wkeInitialize.Call(uintptr(unsafe.Pointer(&set)))
	if r != 0 {
		fmt.Println(err.Error())
	}
	return t
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

func (t *Blink) wkeUnInit() {
	t._wkeUnInitialize.Call()
}
func (t *Blink) WkeOnDownload(wke WkeHandle, callback WkeOnDownloadCallback, param uintptr) {
	t._wkeOnDownload.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Blink) WkeOnAlertBox(wke WkeHandle, callback WkeOnAlertBoxCallback, param uintptr) {
	t._wkeOnAlertBox.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Blink) wkeOnCreateView(wke WkeHandle, callback WkeOnCreateViewCallback, param uintptr) {
	t._wkeOnCreateView.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Blink) wkeSetContextMenuEnabled(wke WkeHandle, show bool) {
	t._wkeSetContextMenuEnabled.Call(uintptr(wke), uintptr(toBool(show)))
}
func (t *Blink) WkeSetNavigationToNewWindowEnable(wke WkeHandle, b bool) {
	t._wkeSetNavigationToNewWindowEnable.Call(uintptr(wke), uintptr(toBool(b)))
}
func (t *Blink) WkeSetUserAgent(wke WkeHandle, ua string) {
	p := StrToCharPtr(ua)
	t._wkeSetUserAgent.Call(uintptr(wke), p)
}

func (t *Blink) wkeSetViewProxy(wke WkeHandle, proxy ProxyInfo) {
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
	t._wkeSetViewProxy.Call(uintptr(wke), uintptr(unsafe.Pointer(&px)))
}

func (t *Blink) WkeSetTransparent(wke WkeHandle, enable bool) {
	t._wkeSetTransparent.Call(uintptr(wke), uintptr(toBool(enable)))
}

func (t *Blink) WkeOnDocumentReady(wke WkeHandle, callback WkeDocumentReady2Callback, param uintptr) {
	t._wkeOnDocumentReady2.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Blink) wkeNetGetRawResponseHead(job WkeNetJob) map[string]string {
	r, _, _ := t._wkeNetGetRawResponseHead.Call(uintptr(job))
	var list []string
	slist := *((*WkeSlist)(unsafe.Pointer(r)))
	for slist.Str != 0 {
		list = append(list, PtrToUtf8(slist.Str))
		if slist.Next == 0 {
			break
		} else {
			slist = *((*WkeSlist)(unsafe.Pointer(slist.Next)))
		}
	}
	hMap := make(map[string]string)
	for i := 0; i < len(list); i += 2 {
		hMap[list[i]] = list[i+1]
	}
	return hMap
}

func (t *Blink) wkeNetSetMIMEType(job WkeNetJob, mime string) {
	p := StrToCharPtr(mime)
	t._wkeNetSetMIMEType.Call(uintptr(job), p)
}

func (t *Blink) wkeNetGetMIMEType(job WkeNetJob) string {
	r, _, _ := t._wkeNetGetMIMEType.Call(uintptr(job))
	return PtrToUtf8(r)
}

func (t *Blink) wkeOnLoadUrlFail(wke WkeHandle, callback WkeLoadUrlFailCallback, param uintptr) {
	t._wkeOnLoadUrlFail.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Blink) wkeOnLoadUrlEnd(wke WkeHandle, callback WkeLoadUrlEndCallback, param uintptr) {
	t._wkeOnLoadUrlEnd.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Blink) wkeNetContinueJob(job WkeNetJob) {
	t._wkeNetContinueJob.Call(uintptr(job))
}

func (t *Blink) wkeNetHoldJobToAsynCommit(job WkeNetJob) {
	t._wkeNetHoldJobToAsynCommit.Call(uintptr(job))
}

func (t *Blink) wkeNetHookRequest(job WkeNetJob) {
	t._wkeNetHookRequest.Call(uintptr(job))
}

func (t *Blink) wkeNetChangeRequestUrl(job WkeNetJob, url string) {
	p := StrToCharPtr(url)
	t._wkeNetChangeRequestUrl.Call(uintptr(job), p)
}

func (t *Blink) wkeNetSetHTTPHeaderField(job WkeNetJob, name, value string) {
	np := StrToCharPtr(name)
	vp := StrToCharPtr(value)
	t._wkeNetSetHTTPHeaderField.Call(uintptr(job), np, vp)
}

func (t *Blink) wkeGetString(str WkeString) string {
	r, _, _ := t._wkeGetString.Call(uintptr(str))
	return PtrToUtf8(r)
}

func (t *Blink) wkeOnConsole(wke WkeHandle, callback WkeConsoleCallback, param uintptr) {
	t._wkeOnConsole.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Blink) wkeIsMainFrame(wke WkeHandle, frame WkeFrame) bool {
	r, _, _ := t._wkeIsMainFrame.Call(uintptr(wke), uintptr(frame))
	return r != 0
}

func (t *Blink) wkeOnDidCreateScriptContext(wke WkeHandle, callback WkeDidCreateScriptContextCallback, param uintptr) {
	t._wkeOnDidCreateScriptContext.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Blink) wkeKillFocus(wke WkeHandle) {
	t._wkeKillFocus.Call(uintptr(wke))
}

func (t *Blink) jsGetWebView(es JsExecState) WkeHandle {
	r, _, _ := t._jsGetWebView.Call(uintptr(es))
	return WkeHandle(r)
}

func (t *Blink) WkeDestroyWebView(wke WkeHandle) {
	t._wkeDestroyWebView.Call(uintptr(wke))
}

func (t *Blink) wkeNetCancelRequest(job WkeNetJob) {
	t._wkeNetCancelRequest.Call(uintptr(job))
}

func (t *Blink) wkeNetOnResponse(wke WkeHandle, callback WkeNetResponseCallback, param uintptr) {
	t._wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Blink) wkeOnLoadUrlBegin(wke WkeHandle, callback WkeLoadUrlBeginCallback, param uintptr) {
	t._wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Blink) wkeNetGetRequestMethod(job WkeNetJob) WkeRequestType {
	r, _, _ := t._wkeNetGetRequestMethod.Call(uintptr(job))
	return WkeRequestType(r)
}

func (t *Blink) wkeNetSetData(job WkeNetJob, buf []byte) {
	if len(buf) == 0 {
		buf = []byte{0}
	}
	t._wkeNetSetData.Call(uintptr(job), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func (t *Blink) wkeSetFocus(wke WkeHandle) {
	t._wkeSetFocus.Call(uintptr(wke))
}

func (t *Blink) wkeFireKeyPressEvent(wke WkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := t._wkeFireKeyPressEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (t *Blink) wkeFireKeyDownEvent(wke WkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := t._wkeFireKeyDownEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (t *Blink) wkeFireKeyUpEvent(wke WkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := t._wkeFireKeyUpEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (t *Blink) wkeFireMouseWheelEvent(wke WkeHandle, x, y, delta, flags int32) bool {
	r, _, _ := t._wkeFireMouseWheelEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(delta),
		uintptr(flags))
	return byte(r) != 0
}
func (t *Blink) wkeFireContextMenuEvent(wke WkeHandle, x, y, flags int32) bool {
	r, _, _ := t._wkeFireContextMenuEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}
func (t *Blink) wkeFireWindowsMessage(wke WkeHandle, hWnd win.HWND, message, wParam, lParam int32) bool {
	r, _, _ := t._wkeFireWindowsMessage.Call(
		uintptr(wke),
		uintptr(hWnd),
		uintptr(message),
		uintptr(wParam),
		uintptr(lParam),
		uintptr(0))
	return byte(r) != 0
}

func (t *Blink) wkeCreateWebWindow(wt WindowType, parent win.HWND, x, y, width, height int32) WkeHandle {
	r, _, _ := t._wkeCreateWebWindow.Call(
		uintptr(wt),
		uintptr(parent),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return WkeHandle(r)
}

func (t *Blink) wkeShowWindow(wke WkeHandle, show bool) {
	t._wkeShowWindow.Call(uintptr(wke), uintptr(toBool(show)))
}

func (t *Blink) wkeFireMouseEvent(wke WkeHandle, message, x, y, flags int32) bool {
	r, _, _ := t._wkeFireMouseEvent.Call(
		uintptr(wke),
		uintptr(message),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}

func (t *Blink) wkeResize(wke WkeHandle, w, h uint32) {
	t._wkeResize.Call(uintptr(wke), uintptr(w), uintptr(h))
}

func (t *Blink) wkeLoadURL(wke WkeHandle, url string) {
	ptr := StrToCharPtr(url)
	t._wkeLoadURL.Call(uintptr(wke), ptr)
}

/*
设置一些实验性选项。debugString可用参数有：
*/
func (t *Blink) WkeSetDebugConfig(wke WkeHandle, debug DebugType, param string) {
	dp := StrToCharPtr(string(debug))
	pp := StrToCharPtr(param)
	t._wkeSetDebugConfig.Call(uintptr(wke), dp, pp)
}

func (t *Blink) wkeOnPaintBitUpdated(wke WkeHandle, callback WkePaintBitUpdatedCallback, param uintptr) {
	t._wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Blink) WkeOnPaintUpdated(wke WkeHandle, callback WkePaintUpdatedCallback, param uintptr) {
	t._wkeOnPaintUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Blink) wkeOnLoadingFinish(wke WkeHandle, callback WkeLoadingFinishCallback, param uintptr) {
	t._wkeOnLoadingFinish.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Blink) wkeEnableHighDPISupport() {
	t._wkeEnableHighDPISupport.Call()
}

func (t *Blink) wkeRunJs(handle WkeHandle, frame WkeFrame, script uintptr, isInClosure bool, param, unUse uintptr) {
	t._wkeRunJs.Call(uintptr(handle), uintptr(frame), script, uintptr(toBool(isInClosure)), 0, param, unUse)
}

func (t *Blink) WkeSetHandle(wke WkeHandle, handle uintptr) {
	t._wkeSetHandle.Call(uintptr(wke), handle)
}
func (t *Blink) wkeOnShowDevtoolsCallback(wke uintptr, param uintptr) uintptr {
	return 0
}
func (t *Blink) WkeCreateWebView() WkeHandle {
	r, _, _ := t._wkeCreateWebView.Call()
	return WkeHandle(r)
}
func (t *Blink) wkeGetHostHWND() win.HWND {
	r, _, _ := t._wkeGetHostHWND.Call()
	return win.HWND(r)
}

func (t *Blink) wkeGetLockedViewDC(handle WkeHandle) win.HDC {
	r, _, _ := t._wkeGetLockedViewDC.Call(uintptr(handle))
	return win.HDC(r)
}
func (t *Blink) wkeRunMessageLoop() {
	t._wkeRunMessageLoop.Call()
}
func (t *Blink) wkeWebFrameGetMainFrame(handle WkeHandle) WkeFrame {
	r, _, _ := t._wkeWebFrameGetMainFrame.Call(uintptr(handle))
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
