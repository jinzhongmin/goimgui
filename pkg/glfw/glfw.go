package glfw

import (
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/goimgui/pkg/gl"
	"github.com/jinzhongmin/usf"
)

var glfwLib *c.Lib

// init library by shared library
func InitLib(path string, mod c.LibMode) {
	if glfwLib != nil {
		return
	}
	var err error
	glfwLib, err = c.NewLib(path, mod)
	if err != nil {
		panic(err)
	}
	Init()
}

func InitGL() {
	lib := c.NewLibFrom(glfwLib)
	gl.InitProc(lib)
}

// int glfwInit(void)
// Initializes the GLFW library.
func Init() int32 {
	return glfwLib.Call(_func_glfwInit_, nil).I32Free()
}

const _fn_glfwGetProcAddress_ = "glfwGetProcAddress"

func SymbolGetProcAddress() unsafe.Pointer {
	return glfwLib.Symbol(_fn_glfwGetProcAddress_)
}

// var _GLProcAddr map[string]unsafe.Pointer
// func init() {
//     _GLProcAddr = make(map[string]unsafe.Pointer)
// }
// GLFWAPI GLFWglproc glfwGetProcAddress(const char* procname);
// func GetProcAddress(procname string) unsafe.Pointer {
//     p, ok := _GLProcAddr[procname]
//     if ok {
//         return p
//     }

//     n := c.CStr(procname)
//     defer usf.Free(n)

//     ptr := glfwLib.Call(_func_glfwGetProcAddress_, []interface{}{&n})
//     defer ptr.Free()
//     _GLProcAddr[procname] = ptr.Ptr()

//     return ptr.Ptr()
// }

// void glfwTerminate(void)
// Terminates the GLFW library.
func Terminate() {
	glfwLib.Call(_func_glfwTerminate_, nil)
}

// void glfwInitHint(int hint, int value)
// Sets the specified init hint to the desired value.
func InitHint(hit int32, value int32) {
	glfwLib.Call(_func_glfwInitHint_, []interface{}{&hit, &value})
}

// void glfwGetVersion(int *major, int *minor, int *rev)
// Retrieves the version of the GLFW library.
func GetVersion() (major, minor, rev int32) {
	_major, _minor, _rev := new(int32), new(int32), new(int32)
	glfwLib.Call(_func_glfwGetVersion_, []interface{}{&_major, &_minor, _rev})
	return *_major, *_minor, *_rev
}

// const char * glfwGetVersionString(void)
// Returns a string describing the compile-time configuration.
func GetVersionString() string {
	return glfwLib.Call(_func_glfwGetVersionString_, nil).StrFree()
}

// int glfwGetError(const char **description)
func GetError() string {
	ret := usf.Malloc(8)
	usf.Memset(ret, 0, 8)

	glfwLib.Call(_func_glfwGetError_, []interface{}{&ret}).Free()

	v := (*c.Value)(ret)
	defer v.Free()
	return v.Str()
}

// Returns and clears the last error for the calling thread.

var _ErrorCallback *c.Callback

// GLFWerrorfun glfwSetErrorCallback(GLFWerrorfun callback)
// Sets the error callback
func SetErrorCallback(fn func(code int32, des string)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if _ErrorCallback == nil {
			_ErrorCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_I32P)
			_ErrorCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				if _ErrorCallback == nil {
					return
				}
				code := args[0].I32()
				des := args[1].Str()
				fn := cb.CallbackFunc.(func(code int32, des string))
				if fn != nil {
					fn(code, des)
				}
			}
		}
		_ErrorCallback.CallbackFunc = fn
		cfn_ptr = _ErrorCallback.FuncPtr()
	} else {
		if _ErrorCallback != nil {
			_ErrorCallback.CallbackCvt = nil
			_ErrorCallback.CallbackFunc = nil
			_ErrorCallback.Free()
			_ErrorCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetErrorCallback_, []interface{}{&cfn_ptr}).Free()
}

type Monitor struct{ ptr unsafe.Pointer }

// GLFWmonitor ** glfwGetMonitors(int *count)
// Returns the currently connected monitors.
func GetMonitors() []*Monitor {
	count := new(int32)
	p := glfwLib.Call(_func_glfwGetMonitors_, []interface{}{&count})
	ptrs := *(*[]unsafe.Pointer)(usf.Slice(p.PtrFree(), uint64(*count)))

	ms := make([]*Monitor, 0)
	for _, p := range ptrs {
		ms = append(ms, &Monitor{p})
	}

	return ms
}

// GLFWmonitor * glfwGetPrimaryMonitor(void)
// Returns the primary monitor.
func GetPrimaryMonitor() *Monitor {
	p := glfwLib.Call(_func_glfwGetPrimaryMonitor_, nil)
	return &Monitor{p.PtrFree()}
}

func (m *Monitor) CPtr() unsafe.Pointer { return m.ptr }

// void glfwGetMonitorPos (GLFWmonitor *monitor, int *xpos, int *ypos)
// Returns the position of the monitor's viewport on the virtual screen.
func (monitor *Monitor) GetMonitorPos() (xpos, ypos int32) {
	_xpos, _ypos := new(int32), new(int32)
	glfwLib.Call(_func_glfwGetMonitorPos_, []interface{}{&monitor.ptr, &_xpos, &_ypos})

	return *_xpos, *_ypos
}

// void glfwGetMonitorWorkarea(GLFWmonitor *monitor, int *xpos, int *ypos, int *width, int *height)
// Retrieves the work area of the monitor.
func (monitor *Monitor) GetMonitorWorkarea() (xpos, ypos, width, height int32) {
	_xpos, _ypos, _width, _height := new(int32), new(int32), new(int32), new(int32)
	glfwLib.Call(_func_glfwGetMonitorWorkarea_, []interface{}{&monitor.ptr, &_xpos, &_ypos, &_width, &_height})

	return *_xpos, *_ypos, *_width, *_height
}

// void glfwGetMonitorPhysicalSize(GLFWmonitor *monitor, int *widthMM, int *heightMM)
// Returns the physical size of the monitor.
func (monitor *Monitor) GetMonitorPhysicalSize() (widthMM, heightMM int32) {
	_widthMM, _heightMM := new(int32), new(int32)
	glfwLib.Call(_func_glfwGetMonitorPhysicalSize_, []interface{}{&monitor.ptr, &_widthMM, &_heightMM})

	return *_widthMM, *_heightMM
}

// void glfwGetMonitorContentScale(GLFWmonitor *monitor, float *xscale, float *yscale)
// Retrieves the content scale for the specified monitor. More...
func (monitor *Monitor) GetMonitorContentScale() (xscale, yscale float32) {
	_xscale, _yscale := new(float32), new(float32)
	glfwLib.Call(_func_glfwGetMonitorContentScale_, []interface{}{&monitor.ptr, &_xscale, &_yscale})

	return *_xscale, *_yscale
}

// const char * glfwGetMonitorName(GLFWmonitor *monitor)
// Returns the name of the specified monitor.
func (monitor *Monitor) GetMonitorName() string {
	return glfwLib.Call(_func_glfwGetMonitorName_, []interface{}{&monitor.ptr}).StrFree()

}

// void     glfwSetMonitorUserPointer (GLFWmonitor *monitor, void *pointer)
// Sets the user pointer of the specified monitor. More...

// void *     glfwGetMonitorUserPointer (GLFWmonitor *monitor)
// Returns the user pointer of the specified monitor. More...

// GLFWmonitorfun     glfwSetMonitorCallback (GLFWmonitorfun callback)
// Sets the monitor configuration callback. More...

type Vidmode struct {
	Width       int32
	Height      int32
	RedBits     int32
	GreenBits   int32
	BlueBits    int32
	RefreshRate int32
}

// const GLFWvidmode * glfwGetVideoModes(GLFWmonitor *monitor, int *count)
// Returns the available video modes for the specified monitor.
func (monitor *Monitor) GetVideoModes() []Vidmode {
	count := new(int32)
	p := glfwLib.Call(_func_glfwGetVideoModes_, []interface{}{&monitor.ptr, &count})
	vidmods := *(*[]Vidmode)(usf.Slice(p.PtrFree(), uint64(*count)))
	return vidmods
}

// const GLFWvidmode * glfwGetVideoMode(GLFWmonitor *monitor)
// Returns the current mode of the specified monitor.
func (monitor *Monitor) GetVideoMode() Vidmode {
	v := glfwLib.Call(_func_glfwGetVideoMode_, []interface{}{&monitor.ptr})
	return *(*Vidmode)(v.PtrFree())
}

type Gammaramp struct {
	Red   unsafe.Pointer //*c.uchar
	Green unsafe.Pointer //*c.uchar
	Blue  unsafe.Pointer //*c.uchar
	Size  uint32
}

// void glfwSetGamma (GLFWmonitor *monitor, float gamma)
// Generates a gamma ramp and sets it for the specified monitor.
func (monitor *Monitor) SetGamma(gamma float32) {
	glfwLib.Call(_func_glfwSetGamma_, []interface{}{&monitor.ptr, &gamma})
}

// const GLFWgammaramp * glfwGetGammaRamp(GLFWmonitor *monitor)
// Returns the current gamma ramp for the specified monitor.
func (monitor *Monitor) GetGammaRamp() Gammaramp {
	v := glfwLib.Call(_func_glfwGetGammaRamp_, []interface{}{&monitor.ptr})
	return *(*Gammaramp)(v.PtrFree())
}

// void glfwSetGammaRamp(GLFWmonitor *monitor, const GLFWgammaramp *ramp)
// Sets the current gamma ramp for the specified monitor. More...
func (monitor *Monitor) SetGammaRamp(ramp *Gammaramp) {
	glfwLib.Call(_func_glfwSetGammaRamp_, []interface{}{&monitor.ptr, &ramp})
}

type Window struct {
	ptr                        unsafe.Pointer
	windowPosCallback          *c.Callback
	windowSizeCallback         *c.Callback
	framebufferSizeCallback    *c.Callback
	windowContentScaleCallback *c.Callback

	windowCloseCallback   *c.Callback
	windowRefreshCallback *c.Callback

	windowFocusCallback    *c.Callback
	windowIconifyCallback  *c.Callback
	windowMaximizeCallback *c.Callback
}

var _WINDOWS_LIST map[unsafe.Pointer]*Window

func _WINDOWS(p unsafe.Pointer) *Window {
	v, ok := _WINDOWS_LIST[p]
	if ok {
		return v
	}
	_WINDOWS_LIST[p] = &Window{ptr: p}
	return _WINDOWS_LIST[p]
}

func init() {
	_WINDOWS_LIST = make(map[unsafe.Pointer]*Window)
}

func NewWindowFromNativeCPointer(p unsafe.Pointer) *Window { return _WINDOWS(p) }

// void glfwDefaultWindowHints(void)
// Resets all window hints to their default values. More...
func DefaultWindowHints() { glfwLib.Call(_func_glfwDefaultWindowHints_, nil) }

type HintCode int32

const (
	GLFW_CONTEXT_VERSION_MAJOR HintCode = 0x00022002
	GLFW_CONTEXT_VERSION_MINOR HintCode = 0x00022003
	GLFW_OPENGL_PROFILE        HintCode = 0x00022008

	GLFW_OPENGL_ANY_PROFILE    int32 = 0
	GLFW_OPENGL_CORE_PROFILE   int32 = 0x00032001
	GLFW_OPENGL_COMPAT_PROFILE int32 = 0x00032002
)

// void glfwWindowHint (int hint, int value)
// Sets the specified window hint to the desired value.
func WindowHint(hint HintCode, value int32) {
	glfwLib.Call(_func_glfwWindowHint_, []interface{}{&hint, &value})
}

// void glfwWindowHintString(int hint, const char *value)
// Sets the specified window hint to the desired value. More...
func WindowHintString(hint HintCode, value string) {
	_value := c.CStr(value)
	glfwLib.Call(_func_glfwWindowHintString_, []interface{}{&hint, &_value})
}

// GLFWwindow * glfwCreateWindow(int width, int height, const char *title, GLFWmonitor *monitor, GLFWwindow *share)
// Creates a window and its associated context.
func CreateWindow(width, height int32, title string, monitor *Monitor, share *Window) *Window {
	t := c.CStr(title)
	defer usf.Free(t)

	var m, s unsafe.Pointer = nil, nil

	if monitor != nil {
		m = monitor.ptr
	}
	if share != nil {
		s = share.ptr
	}

	return _WINDOWS(glfwLib.Call(_func_glfwCreateWindow_, []interface{}{&width, &height, &t, &m, &s}).PtrFree())
}
func (win *Window) CPtr() unsafe.Pointer { return win.ptr }

// void glfwDestroyWindow (GLFWwindow *window)
// Destroys the specified window and its context.
func (win *Window) DestroyWindow() {
	glfwLib.Call(_func_glfwDestroyWindow_, []interface{}{&win.ptr})
	// type _wincallback_free interface {
	//     Free()
	// }
	// frees := []_wincallback_free{
	//     win.windowPosCallback,
	//     win.windowSizeCallback,
	//     win.framebufferSizeCallback,
	//     win.windowContentScaleCallback,
	//     win.windowCloseCallback,
	//     win.windowRefreshCallback,
	//     win.windowFocusCallback,
	//     win.windowIconifyCallback,
	//     win.windowMaximizeCallback,
	// }
	// for _, f := range frees {
	//     if f != nil {
	//         f.Free()
	//     }
	// }
}

// int glfwWindowShouldClose(GLFWwindow *window)
// Checks the close flag of the specified window.
func (win *Window) WindowShouldClose() bool {
	return glfwLib.Call(_func_glfwWindowShouldClose_, []interface{}{&win.ptr}).BoolFree()

}

// void glfwSetWindowShouldClose(GLFWwindow *window, int value)
// Sets the close flag of the specified window.
func (win *Window) SetWindowShouldClose(value bool) {
	v := c.CBool(value)
	glfwLib.Call(_func_glfwSetWindowShouldClose_, []interface{}{&win.ptr, &v})
}

// void glfwSetWindowTitle(GLFWwindow *window, const char *title)
// Sets the title of the specified window.
func (win *Window) SetWindowTitle(title string) {
	t := c.CStr(title)
	defer usf.Free(t)
	glfwLib.Call(_func_glfwSetWindowTitle_, []interface{}{&win.ptr, &t})
}

// void glfwSetWindowIcon(GLFWwindow *window, int count, const GLFWimage *images)
// Sets the icon for the specified window.

// void glfwGetWindowPos(GLFWwindow *window, int *xpos, int *ypos)
// Retrieves the position of the content area of the specified window.
func (win *Window) GetWindowPos() (xpos, ypos int32) {
	_xpos, _ypos := new(int32), new(int32)
	glfwLib.Call(_func_glfwGetWindowPos_, []interface{}{&win.ptr, &_xpos, &_ypos})
	return *_xpos, *_ypos
}

// void glfwSetWindowPos(GLFWwindow *window, int xpos, int ypos)
// Sets the position of the content area of the specified window.
func (win *Window) SetWindowPos(xpos, ypos int32) {
	glfwLib.Call(_func_glfwSetWindowPos_, []interface{}{&win.ptr, &xpos, &xpos})
}

// void glfwGetWindowSize(GLFWwindow *window, int *width, int *height)
// Retrieves the size of the content area of the specified window.
func (win *Window) GetWindowSize() (width, height int32) {
	_width, _height := new(int32), new(int32)
	glfwLib.Call(_func_glfwGetWindowSize_, []interface{}{&win.ptr, &_width, &_height})
	return *_width, *_height
}

// void glfwSetWindowSizeLimits(GLFWwindow *window, int minwidth, int minheight, int maxwidth, int maxheight)
// Sets the size limits of the specified window.
func (win *Window) SetWindowSizeLimits(minwidth, minheight, maxwidth, maxheight int32) {
	glfwLib.Call(_func_glfwSetWindowSizeLimits_, []interface{}{&win.ptr, &minwidth, &minheight, &maxwidth, &maxheight})

}

// void glfwSetWindowAspectRatio(GLFWwindow *window, int numer, int denom)
// Sets the aspect ratio of the specified window.
func (win *Window) SetWindowAspectRatio(numer, denom int32) {
	glfwLib.Call(_func_glfwSetWindowAspectRatio_, []interface{}{&win.ptr, &numer, &denom})
}

// void glfwSetWindowSize(GLFWwindow *window, int width, int height)
// Sets the size of the content area of the specified window.
func (win *Window) SetWindowSize(width, height int32) {
	glfwLib.Call(_func_glfwSetWindowSize_, []interface{}{&win.ptr, &width, &height})
}

// void glfwGetFramebufferSize(GLFWwindow *window, int *width, int *height)
// Retrieves the size of the framebuffer of the specified window.
func (win *Window) GetFramebufferSize() (width, height int32) {
	_width, _height := &width, &height
	glfwLib.Call(_func_glfwGetFramebufferSize_, []interface{}{&win.ptr, &_width, &_height})

	return
}

// void glfwGetFramebufferSize(GLFWwindow *window, int *width, int *height)
// Retrieves the size of the framebuffer of the specified window.
func (win *Window) GetFramebufferSizeRef(width, height *int32) {
	glfwLib.Call(_func_glfwGetFramebufferSize_, []interface{}{&win.ptr, &width, &height})

}

// void glfwGetWindowFrameSize(GLFWwindow *window, int *left, int *top, int *right, int *bottom)
// Retrieves the size of the frame of the window.
func (win *Window) GetWindowFrameSize() (left, top, right, bottom int32) {
	_left, _top, _right, _bottom := &left, &top, &right, &bottom
	glfwLib.Call(_func_glfwGetWindowFrameSize_, []interface{}{&win.ptr, &_left, &_top, &_right, &_bottom})

	return
}

// void glfwGetWindowContentScale(GLFWwindow *window, float *xscale, float *yscale)
// Retrieves the content scale for the specified window.
func (win *Window) GetWindowContentScale() (xscale, yscale float32) {
	_xscale, _yscale := &xscale, &yscale
	glfwLib.Call(_func_glfwGetWindowContentScale_, []interface{}{&win.ptr, &_xscale, &_yscale})
	return
}

// float glfwGetWindowOpacity(GLFWwindow *window)
// Returns the opacity of the whole window.
func (win *Window) GetWindowOpacity() float32 {
	return glfwLib.Call(_func_glfwGetWindowOpacity_, []interface{}{&win.ptr}).F32Free()
}

// void glfwSetWindowOpacity(GLFWwindow *window, float opacity)
// Sets the opacity of the whole window.
func (win *Window) SetWindowOpacity(Opacity float32) {
	glfwLib.Call(_func_glfwSetWindowOpacity_, []interface{}{&win.ptr, &Opacity})
}

// void glfwIconifyWindow(GLFWwindow *window)
// Iconifies the specified window.
func (win *Window) IconifyWindow() {
	glfwLib.Call(_func_glfwIconifyWindow_, []interface{}{&win.ptr})
}

// void glfwRestoreWindow(GLFWwindow *window)
// Restores the specified window.
func (win *Window) RestoreWindow() {
	glfwLib.Call(_func_glfwRestoreWindow_, []interface{}{&win.ptr})
}

// void glfwMaximizeWindow(GLFWwindow *window)
// Maximizes the specified window.
func (win *Window) MaximizeWindow() {
	glfwLib.Call(_func_glfwMaximizeWindow_, []interface{}{&win.ptr})
}

// void glfwShowWindow(GLFWwindow *window)
// Makes the specified window visible.
func (win *Window) ShowWindow() {
	glfwLib.Call(_func_glfwShowWindow_, []interface{}{&win.ptr})
}

// void glfwHideWindow(GLFWwindow *window)
// Hides the specified window.
func (win *Window) HideWindow() {
	glfwLib.Call(_func_glfwHideWindow_, []interface{}{&win.ptr})
}

// void glfwFocusWindow(GLFWwindow *window)
// Brings the specified window to front and sets input focus.
func (win *Window) FocusWindow() {
	glfwLib.Call(_func_glfwFocusWindow_, []interface{}{&win.ptr})
}

// void glfwRequestWindowAttention(GLFWwindow *window)
// Requests user attention to the specified window. More...
func (win *Window) RequestWindowAttention() {
	glfwLib.Call(_func_glfwRequestWindowAttention_, []interface{}{&win.ptr})
}

// GLFWmonitor * glfwGetWindowMonitor(GLFWwindow *window)
// Returns the monitor that the window uses for full screen mode.
func (win *Window) GetWindowMonitor() *Monitor {
	p := glfwLib.Call(_func_glfwGetWindowMonitor_, []interface{}{&win.ptr})
	return &Monitor{p.PtrFree()}
}

// void glfwSetWindowMonitor(GLFWwindow *window, GLFWmonitor *monitor,
// int xpos, int ypos, int width, int height, int refreshRate)
// Sets the mode, monitor, video mode and placement of a window.
func (win *Window) SetWindowMonitor(monitor *Monitor, xpos, ypos int32, width, height int32, refreshRate int32) {
	glfwLib.Call(_func_glfwSetWindowMonitor_, []interface{}{&win.ptr, &monitor.ptr, &xpos, &ypos, &width, &height, &refreshRate})

}

type WindowAttribs int32

const (
	GLFW_FOCUSED                 WindowAttribs = 0x00020001
	GLFW_ICONIFIED               WindowAttribs = 0x00020002
	GLFW_MAXIMIZED               WindowAttribs = 0x00020008
	GLFW_HOVERED                 WindowAttribs = 0x0002000B
	GLFW_VISIBLE                 WindowAttribs = 0x00020004
	GLFW_RESIZABLE               WindowAttribs = 0x00020003
	GLFW_DECORATED               WindowAttribs = 0x00020005
	GLFW_AUTO_ICONIFY            WindowAttribs = 0x00020006
	GLFW_FLOATING                WindowAttribs = 0x00020007
	GLFW_TRANSPARENT_FRAMEBUFFER WindowAttribs = 0x0002000A
	GLFW_FOCUS_ON_SHOW           WindowAttribs = 0x0002000C
)

// int glfwGetWindowAttrib(GLFWwindow *window, int attrib)
// Returns an attribute of the specified window.
func (win *Window) GetWindowAttrib(attrib WindowAttribs) bool {
	return glfwLib.Call(_func_glfwGetWindowAttrib_, []interface{}{&win.ptr, &attrib}).BoolFree()
}

// void glfwSetWindowAttrib (GLFWwindow *window, int attrib, int value)
// Sets an attribute of the specified window.
func (win *Window) SetWindowAttrib(attrib WindowAttribs, value int32) {
	glfwLib.Call(_func_glfwSetWindowAttrib_, []interface{}{&win.ptr, &attrib, &value})
}

// void glfwSetWindowUserPointer(GLFWwindow *window, void *pointer)
// Sets the user pointer of the specified window.
func (win *Window) SetWindowUserPointer(ptr unsafe.Pointer) {
	glfwLib.Call(_func_glfwSetWindowUserPointer_, []interface{}{&win.ptr, &ptr})
}

// void glfwGetWindowUserPointer(GLFWwindow *window)
// Returns the user pointer of the specified window.
func (win *Window) GetWindowUserPointer() unsafe.Pointer {
	return glfwLib.Call(_func_glfwGetWindowUserPointer_, []interface{}{&win.ptr}).PtrFree()
}

// GLFWwindowposfun glfwSetWindowPosCallback(GLFWwindow *window, GLFWwindowposfun callback)
// Sets the position callback for the specified window.
func (win *Window) SetWindowPosCallback(fn func(win *Window, xpos, ypos int32)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.windowPosCallback == nil { //not define
			win.windowPosCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_PI32I32)
			win.windowPosCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				_xpos := args[1].I32()
				_ypos := args[2].I32()
				fn := cb.CallbackFunc.(func(win *Window, xpos, ypos int32))
				if fn != nil {
					fn(win, _xpos, _ypos)
				}
			}
		}

		win.windowPosCallback.CallbackFunc = fn
		cfn_ptr = win.windowPosCallback.FuncPtr()
	} else { //remove callback
		if win.windowPosCallback != nil {
			win.windowPosCallback.CallbackCvt = nil
			win.windowPosCallback.CallbackFunc = nil
			win.windowPosCallback.Free()
			win.windowPosCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetWindowPosCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// GLFWwindowsizefun glfwSetWindowSizeCallback(GLFWwindow *window, GLFWwindowsizefun callback)
// Sets the size callback for the specified window.
func (win *Window) SetWindowSizeCallback(fn func(win *Window, width, height int32)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.windowSizeCallback == nil { //not define
			win.windowSizeCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_PI32I32)
			win.windowSizeCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				_width := args[1].I32()
				_height := args[2].I32()
				fn := cb.CallbackFunc.(func(win *Window, width, height int32))
				if fn != nil {
					fn(win, _width, _height)
				}
			}
		}
		win.windowSizeCallback.CallbackFunc = fn
		cfn_ptr = win.windowSizeCallback.FuncPtr()
	} else { //remove callback
		if win.windowSizeCallback != nil {
			win.windowSizeCallback.CallbackCvt = nil
			win.windowSizeCallback.CallbackFunc = nil
			win.windowSizeCallback.Free()
			win.windowSizeCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetWindowSizeCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// GLFWwindowclosefun glfwSetWindowCloseCallback(GLFWwindow *window, GLFWwindowclosefun callback)
// Sets the close callback for the specified window...
func (win *Window) SetWindowCloseCallback(fn func(win *Window)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.windowCloseCallback == nil { //not define
			win.windowCloseCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_P)
			win.windowCloseCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				fn := cb.CallbackFunc.(func(win *Window))
				if fn != nil {
					fn(win)
				}
			}
		}
		win.windowCloseCallback.CallbackFunc = fn
		cfn_ptr = win.windowCloseCallback.FuncPtr()
	} else { //remove callback
		if win.windowCloseCallback != nil {
			win.windowCloseCallback.CallbackCvt = nil
			win.windowCloseCallback.CallbackFunc = nil
			win.windowCloseCallback.Free()
			win.windowCloseCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetWindowCloseCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// GLFWwindowrefreshfun glfwSetWindowRefreshCallback(GLFWwindow *window, GLFWwindowrefreshfun callback)
// Sets the refresh callback for the specified window.
func (win *Window) SetWindowRefreshCallback(fn func(win *Window)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.windowRefreshCallback == nil { //not define
			win.windowRefreshCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_P)
			win.windowRefreshCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				fn := cb.CallbackFunc.(func(win *Window))
				if fn != nil {
					fn(win)
				}
			}
		}
		win.windowRefreshCallback.CallbackFunc = fn
		cfn_ptr = win.windowRefreshCallback.FuncPtr()
	} else { //remove callback
		if win.windowRefreshCallback != nil {
			win.windowRefreshCallback.CallbackCvt = nil
			win.windowRefreshCallback.CallbackFunc = nil
			win.windowRefreshCallback.Free()
			win.windowRefreshCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetWindowRefreshCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// GLFWwindowfocusfun glfwSetWindowFocusCallback(GLFWwindow *window, GLFWwindowfocusfun callback)
// Sets the focus callback for the specified window. More...
func (win *Window) SetWindowFocusCallback(fn func(win *Window, focused bool)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.windowFocusCallback == nil { //not define
			win.windowFocusCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_PI32)
			win.windowFocusCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				focused := args[1].Bool()
				fn := cb.CallbackFunc.(func(win *Window, focused bool))
				if fn != nil {
					fn(win, focused)
				}
			}
		}
		win.windowFocusCallback.CallbackFunc = fn
		cfn_ptr = win.windowFocusCallback.FuncPtr()
	} else { //remove callback
		if win.windowFocusCallback != nil {
			win.windowFocusCallback.CallbackCvt = nil
			win.windowFocusCallback.CallbackFunc = nil
			win.windowFocusCallback.Free()
			win.windowFocusCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetWindowFocusCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// GLFWwindowiconifyfun glfwSetWindowIconifyCallback(GLFWwindow *window, GLFWwindowiconifyfun callback)
// Sets the iconify callback for the specified window.
func (win *Window) SetWindowIconifyCallback(fn func(win *Window, iconified bool)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.windowIconifyCallback == nil { //not define
			win.windowIconifyCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_PI32)
			win.windowIconifyCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				iconified := args[1].Bool()
				fn := cb.CallbackFunc.(func(win *Window, iconified bool))
				if fn != nil {
					fn(win, iconified)
				}
			}
		}
		win.windowIconifyCallback.CallbackFunc = fn
		cfn_ptr = win.windowIconifyCallback.FuncPtr()
	} else { //remove callback
		if win.windowIconifyCallback != nil {
			win.windowIconifyCallback.CallbackCvt = nil
			win.windowIconifyCallback.CallbackFunc = nil
			win.windowIconifyCallback.Free()
			win.windowIconifyCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetWindowIconifyCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// GLFWwindowmaximizefun glfwSetWindowMaximizeCallback(GLFWwindow *window, GLFWwindowmaximizefun callback)
// Sets the maximize callback for the specified window.
func (win *Window) SetWindowMaximizeCallback(fn func(win *Window, maximized bool)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.windowMaximizeCallback == nil { //not define
			win.windowMaximizeCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_PI32)
			win.windowMaximizeCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				maximized := args[1].Bool()
				fn := cb.CallbackFunc.(func(win *Window, maximized bool))
				if fn != nil {
					fn(win, maximized)
				}
			}
		}
		win.windowMaximizeCallback.CallbackFunc = fn
		cfn_ptr = win.windowMaximizeCallback.FuncPtr()
	} else { //remove callback
		if win.windowMaximizeCallback != nil {
			win.windowMaximizeCallback.CallbackCvt = nil
			win.windowMaximizeCallback.CallbackFunc = nil
			win.windowMaximizeCallback.Free()
			win.windowMaximizeCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetWindowMaximizeCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// GLFWframebuffersizefun glfwSetFramebufferSizeCallback(GLFWwindow *window, GLFWframebuffersizefun callback)
// Sets the framebuffer resize callback for the specified window.
func (win *Window) SetFramebufferSizeCallback(fn func(win *Window, width, height int32)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.framebufferSizeCallback == nil { //not define
			win.framebufferSizeCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_PI32I32)
			win.framebufferSizeCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				_width := args[1].I32()
				_height := args[2].I32()
				fn := cb.CallbackFunc.(func(win *Window, width, height int32))
				if fn != nil {
					fn(win, _width, _height)
				}
			}
		}
		win.framebufferSizeCallback.CallbackFunc = fn
		cfn_ptr = win.framebufferSizeCallback.FuncPtr()
	} else { //remove callback
		if win.framebufferSizeCallback != nil {
			win.framebufferSizeCallback.CallbackCvt = nil
			win.framebufferSizeCallback.CallbackFunc = nil
			win.framebufferSizeCallback.Free()
			win.framebufferSizeCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetFramebufferSizeCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// GLFWwindowcontentscalefun glfwSetWindowContentScaleCallback(GLFWwindow *window, GLFWwindowcontentscalefun callback)
// Sets the window content scale callback for the specified window.
func (win *Window) SetWindowContentScaleCallback(fn func(win *Window, xscale, yscale float32)) {
	var cfn_ptr unsafe.Pointer = nil
	if fn != nil { //add callback
		if win.windowContentScaleCallback == nil { //not define
			win.windowContentScaleCallback = c.NewCallback(c.AbiDefault, c.Void, _typs_PF32F32)
			win.windowContentScaleCallback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				_xscale := args[1].F32()
				_yscale := args[2].F32()
				fn := cb.CallbackFunc.(func(win *Window, xscale, yscale float32))
				if fn != nil {
					fn(win, _xscale, _yscale)
				}
			}
		}
		win.windowContentScaleCallback.CallbackFunc = fn
		cfn_ptr = win.windowContentScaleCallback.FuncPtr()
	} else { //remove callback
		if win.windowContentScaleCallback != nil {
			win.windowContentScaleCallback.CallbackCvt = nil
			win.windowContentScaleCallback.CallbackFunc = nil
			win.windowContentScaleCallback.Free()
			win.windowContentScaleCallback = nil
		}
	}
	glfwLib.Call(_func_glfwSetWindowContentScaleCallback_, []interface{}{&win.ptr, &cfn_ptr}).Free()

}

// void glfwPollEvents(void)
// Processes all pending events.
func PollEvents() {
	glfwLib.Call(_func_glfwPollEvents_, nil)
}

// void glfwWaitEvents (void)
// Waits until events are queued and processes them.
func WaitEvents() {
	glfwLib.Call(_func_glfwWaitEvents_, nil)
}

// void  glfwWaitEventsTimeout(double timeout)
// Waits with timeout until events are queued and processes them.
func WaitEventsTimeout(timeout float64) {
	glfwLib.Call(_func_glfwWaitEventsTimeout_, []interface{}{&timeout})
}

// void glfwPostEmptyEvent(void)
// Posts an empty event to the event queue.
func PostEmptyEvent() { glfwLib.Call(_func_glfwPostEmptyEvent_, nil) }

// void glfwSwapBuffers(GLFWwindow *window)
// Swaps the front and back buffers of the specified window.
func (win *Window) SwapBuffers() {
	glfwLib.Call(_func_glfwSwapBuffers_, []interface{}{&win.ptr})
}

// void glfwMakeContextCurrent(GLFWwindow *window)
// Makes the context of the specified window current for the calling thread.
func (win *Window) MakeContextCurrent() {
	glfwLib.Call(_func_glfwMakeContextCurrent_, []interface{}{&win.ptr})
}

// GLFWwindow * glfwGetCurrentContext(void)
// Returns the window whose context is current on the calling thread.
func GetCurrentContext() *Window {
	p := glfwLib.Call(_func_glfwGetCurrentContext_, nil)
	return _WINDOWS(p.PtrFree())
}

// void glfwSwapInterval(int interval)
// Sets the swap interval for the current context.
func SwapInterval(interval int32) {
	glfwLib.Call(_func_glfwSwapInterval_, []interface{}{&interval})
}

// int glfwExtensionSupported(const char *extension)
// Returns whether the specified extension is available.
func ExtensionSupported(extension string) bool {
	exs := c.CStr(extension)
	defer usf.Free(exs)
	return glfwLib.Call(_func_glfwExtensionSupported_, []interface{}{&exs}).BoolFree()
}
