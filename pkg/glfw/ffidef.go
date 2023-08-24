package glfw

import "github.com/jinzhongmin/goffi/pkg/c"

var (
	_typs_F64               = []c.Type{c.F64}
	_typs_I32               = []c.Type{c.I32}
	_typs_I32I32            = []c.Type{c.I32, c.I32}
	_typs_I32I32PPP         = []c.Type{c.I32, c.I32, c.Pointer, c.Pointer, c.Pointer}
	_typs_I32P              = []c.Type{c.I32, c.Pointer}
	_typs_P                 = []c.Type{c.Pointer}
	_typs_PF32              = []c.Type{c.Pointer, c.F32}
	_typs_PF32F32           = []c.Type{c.Pointer, c.F32, c.F32}
	_typs_PI32              = []c.Type{c.Pointer, c.I32}
	_typs_PI32I32           = []c.Type{c.Pointer, c.I32, c.I32}
	_typs_PI32I32I32I32     = []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32}
	_typs_PP                = []c.Type{c.Pointer, c.Pointer}
	_typs_PPI32I32I32I32I32 = []c.Type{c.Pointer, c.Pointer, c.I32, c.I32, c.I32, c.I32, c.I32}
	_typs_PPP               = []c.Type{c.Pointer, c.Pointer, c.Pointer}
	_typs_PPPPP             = []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.Pointer}
)

var (
	_func_glfwSetWindowShouldClose_          = &c.FuncPrototype{Name: "glfwSetWindowShouldClose", OutType: c.Void, InTypes: _typs_PI32}
	_func_glfwDestroyWindow_                 = &c.FuncPrototype{Name: "glfwDestroyWindow", OutType: c.Void, InTypes: _typs_P}
	_func_glfwFocusWindow_                   = &c.FuncPrototype{Name: "glfwFocusWindow", OutType: c.Void, InTypes: _typs_P}
	_func_glfwGetError_                      = &c.FuncPrototype{Name: "glfwGetError", OutType: c.I32, InTypes: _typs_P}
	_func_glfwGetFramebufferSize_            = &c.FuncPrototype{Name: "glfwGetFramebufferSize", OutType: c.Void, InTypes: _typs_PPP}
	_func_glfwGetMonitorContentScale_        = &c.FuncPrototype{Name: "glfwGetMonitorContentScale", OutType: c.Void, InTypes: _typs_PPP}
	_func_glfwGetMonitorPhysicalSize_        = &c.FuncPrototype{Name: "glfwGetMonitorPhysicalSize", OutType: c.Void, InTypes: _typs_PPP}
	_func_glfwGetMonitorPos_                 = &c.FuncPrototype{Name: "glfwGetMonitorPos", OutType: c.Void, InTypes: _typs_PPP}
	_func_glfwGetMonitorWorkarea_            = &c.FuncPrototype{Name: "glfwGetMonitorWorkarea", OutType: c.Void, InTypes: _typs_PPPPP}
	_func_glfwGetVersion_                    = &c.FuncPrototype{Name: "glfwGetVersion", OutType: c.Void, InTypes: _typs_PPP}
	_func_glfwGetWindowContentScale_         = &c.FuncPrototype{Name: "glfwGetWindowContentScale", OutType: c.Void, InTypes: _typs_PPP}
	_func_glfwGetWindowFrameSize_            = &c.FuncPrototype{Name: "glfwGetWindowFrameSize", OutType: c.Void, InTypes: _typs_PPPPP}
	_func_glfwGetWindowPos_                  = &c.FuncPrototype{Name: "glfwGetWindowPos", OutType: c.Void, InTypes: _typs_PPP}
	_func_glfwGetWindowSize_                 = &c.FuncPrototype{Name: "glfwGetWindowSize", OutType: c.Void, InTypes: _typs_PPP}
	_func_glfwHideWindow_                    = &c.FuncPrototype{Name: "glfwHideWindow", OutType: c.Void, InTypes: _typs_P}
	_func_glfwIconifyWindow_                 = &c.FuncPrototype{Name: "glfwIconifyWindow", OutType: c.Void, InTypes: _typs_P}
	_func_glfwInitHint_                      = &c.FuncPrototype{Name: "glfwInitHint", OutType: c.Void, InTypes: _typs_I32I32}
	_func_glfwMakeContextCurrent_            = &c.FuncPrototype{Name: "glfwMakeContextCurrent", OutType: c.Void, InTypes: _typs_P}
	_func_glfwMaximizeWindow_                = &c.FuncPrototype{Name: "glfwMaximizeWindow", OutType: c.Void, InTypes: _typs_P}
	_func_glfwPollEvents_                    = &c.FuncPrototype{Name: "glfwPollEvents", OutType: c.Void, InTypes: nil}
	_func_glfwRequestWindowAttention_        = &c.FuncPrototype{Name: "glfwRequestWindowAttention", OutType: c.Void, InTypes: _typs_P}
	_func_glfwRestoreWindow_                 = &c.FuncPrototype{Name: "glfwRestoreWindow", OutType: c.Void, InTypes: _typs_P}
	_func_glfwSetErrorCallback_              = &c.FuncPrototype{Name: "glfwSetErrorCallback", OutType: c.Pointer, InTypes: _typs_P}
	_func_glfwSetFramebufferSizeCallback_    = &c.FuncPrototype{Name: "glfwSetFramebufferSizeCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetGamma_                      = &c.FuncPrototype{Name: "glfwSetGamma", OutType: c.Void, InTypes: _typs_PF32}
	_func_glfwSetGammaRamp_                  = &c.FuncPrototype{Name: "glfwSetGammaRamp", OutType: c.Void, InTypes: _typs_PP}
	_func_glfwSetWindowAspectRatio_          = &c.FuncPrototype{Name: "glfwSetWindowAspectRatio", OutType: c.Void, InTypes: _typs_PI32I32}
	_func_glfwSetWindowAttrib_               = &c.FuncPrototype{Name: "glfwSetWindowAttrib", OutType: c.Void, InTypes: _typs_PI32}
	_func_glfwSetWindowCloseCallback_        = &c.FuncPrototype{Name: "glfwSetWindowCloseCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetWindowContentScaleCallback_ = &c.FuncPrototype{Name: "glfwSetWindowContentScaleCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetWindowFocusCallback_        = &c.FuncPrototype{Name: "glfwSetWindowFocusCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetWindowIconifyCallback_      = &c.FuncPrototype{Name: "glfwSetWindowIconifyCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetWindowMaximizeCallback_     = &c.FuncPrototype{Name: "glfwSetWindowMaximizeCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetWindowMonitor_              = &c.FuncPrototype{Name: "glfwSetWindowMonitor", OutType: c.Void, InTypes: _typs_PPI32I32I32I32I32}
	_func_glfwSetWindowOpacity_              = &c.FuncPrototype{Name: "glfwSetWindowOpacity", OutType: c.Void, InTypes: _typs_PF32}
	_func_glfwSetWindowPos_                  = &c.FuncPrototype{Name: "glfwSetWindowPos", OutType: c.Void, InTypes: _typs_PI32I32}
	_func_glfwSetWindowPosCallback_          = &c.FuncPrototype{Name: "glfwSetWindowPosCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetWindowRefreshCallback_      = &c.FuncPrototype{Name: "glfwSetWindowRefreshCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetWindowSize_                 = &c.FuncPrototype{Name: "glfwSetWindowSize", OutType: c.Void, InTypes: _typs_PI32I32}
	_func_glfwSetWindowSizeCallback_         = &c.FuncPrototype{Name: "glfwSetWindowSizeCallback", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwSetWindowSizeLimits_           = &c.FuncPrototype{Name: "glfwSetWindowSizeLimits", OutType: c.Void, InTypes: _typs_PI32I32I32I32}
	_func_glfwSetWindowTitle_                = &c.FuncPrototype{Name: "glfwSetWindowTitle", OutType: c.Void, InTypes: _typs_PP}
	_func_glfwSetWindowUserPointer_          = &c.FuncPrototype{Name: "glfwSetWindowUserPointer", OutType: c.Void, InTypes: _typs_PP}
	_func_glfwShowWindow_                    = &c.FuncPrototype{Name: "glfwShowWindow", OutType: c.Void, InTypes: _typs_P}
	_func_glfwSwapBuffers_                   = &c.FuncPrototype{Name: "glfwSwapBuffers", OutType: c.Void, InTypes: _typs_P}
	_func_glfwSwapInterval_                  = &c.FuncPrototype{Name: "glfwSwapInterval", OutType: c.Void, InTypes: _typs_I32}
	_func_glfwTerminate_                     = &c.FuncPrototype{Name: "glfwTerminate", OutType: c.Void, InTypes: nil}
	_func_glfwWaitEvents_                    = &c.FuncPrototype{Name: "glfwWaitEvents", OutType: c.Void, InTypes: nil}
	_func_glfwWaitEventsTimeout_             = &c.FuncPrototype{Name: "glfwWaitEventsTimeout", OutType: c.Void, InTypes: _typs_F64}
	_func_glfwWindowHint_                    = &c.FuncPrototype{Name: "glfwWindowHint", OutType: c.Void, InTypes: _typs_I32I32}
	_func_glfwWindowHintString_              = &c.FuncPrototype{Name: "glfwWindowHintString", OutType: c.Void, InTypes: _typs_I32P}
	_func_glfwCreateWindow_                  = &c.FuncPrototype{Name: "glfwCreateWindow", OutType: c.Pointer, InTypes: _typs_I32I32PPP}
	_func_glfwGetCurrentContext_             = &c.FuncPrototype{Name: "glfwGetCurrentContext", OutType: c.Pointer, InTypes: nil}
	_func_glfwGetMonitors_                   = &c.FuncPrototype{Name: "glfwGetMonitors", OutType: c.Pointer, InTypes: _typs_P}
	_func_glfwGetPrimaryMonitor_             = &c.FuncPrototype{Name: "glfwGetPrimaryMonitor", OutType: c.Pointer, InTypes: nil}
	_func_glfwGetVideoModes_                 = &c.FuncPrototype{Name: "glfwGetVideoModes", OutType: c.Pointer, InTypes: _typs_PP}
	_func_glfwGetWindowMonitor_              = &c.FuncPrototype{Name: "glfwGetWindowMonitor", OutType: c.Pointer, InTypes: _typs_P}
	_func_glfwInit_                          = &c.FuncPrototype{Name: "glfwInit", OutType: c.I32, InTypes: nil}
	_func_glfwWindowShouldClose_             = &c.FuncPrototype{Name: "glfwWindowShouldClose", OutType: c.I32, InTypes: _typs_P}
	_func_glfwExtensionSupported_            = &c.FuncPrototype{Name: "glfwExtensionSupported", OutType: c.I32, InTypes: _typs_P}
	_func_glfwGetMonitorName_                = &c.FuncPrototype{Name: "glfwGetMonitorName", OutType: c.Pointer, InTypes: _typs_P}
	_func_glfwGetVersionString_              = &c.FuncPrototype{Name: "glfwGetVersionString", OutType: c.Pointer, InTypes: nil}
	_func_glfwGetWindowAttrib_               = &c.FuncPrototype{Name: "glfwGetWindowAttrib", OutType: c.I32, InTypes: _typs_PI32}
	_func_glfwGetWindowOpacity_              = &c.FuncPrototype{Name: "glfwGetWindowOpacity", OutType: c.F32, InTypes: _typs_P}
	_func_glfwGetWindowUserPointer_          = &c.FuncPrototype{Name: "glfwGetWindowUserPointer", OutType: c.Pointer, InTypes: _typs_P}
	_func_glfwGetGammaRamp_                  = &c.FuncPrototype{Name: "glfwGetGammaRamp", OutType: c.Pointer, InTypes: _typs_P}
	_func_glfwGetVideoMode_                  = &c.FuncPrototype{Name: "glfwGetVideoMode", OutType: c.Pointer, InTypes: _typs_P}
	_func_glfwDefaultWindowHints_            = &c.FuncPrototype{Name: "glfwDefaultWindowHints", OutType: c.Void, InTypes: nil}
	_func_glfwPostEmptyEvent_                = &c.FuncPrototype{Name: "glfwPostEmptyEvent", OutType: c.Void, InTypes: nil}
)
