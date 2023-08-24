package gi

import (
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/usf"
)

var giLib *c.Lib

// init library by shared library
func InitLib(path string, mod c.LibMode) {
	var err error
	giLib, err = c.NewLib(path, mod)
	if err != nil {
		panic(err)
	}
}

// CIMGUI_API ImVec2* ImVec2_ImVec2_Nil(void);
func NewImVec2() *ImVec2 {
	p := giLib.Call(_func_ImVec2_ImVec2_Nil_, nil)
	defer p.Free()
	return (*ImVec2)(p.Ptr())
}

// CIMGUI_API void ImVec2_destroy(ImVec2* self);
func (vec2 *ImVec2) Destroy() {
	giLib.Call(_func_ImVec2_destroy_, []interface{}{&vec2})
}

// CIMGUI_API ImVec2* ImVec2_ImVec2_Float(float _x,float _y);
func NewImVec2Float(x, y float32) *ImVec2 {
	p := giLib.Call(_func_ImVec2_ImVec2_Float_, []interface{}{&x, &y})
	return (*ImVec2)(p.PtrFree())
}

// CIMGUI_API ImVec4* ImVec4_ImVec4_Nil(void);
func NewImVec4() *ImVec4 {
	p := giLib.Call(_func_ImVec4_ImVec4_Nil_, nil)
	return (*ImVec4)(p.PtrFree())
}

// CIMGUI_API void ImVec4_destroy(ImVec4* self);
func (vec4 *ImVec4) Destroy() {
	giLib.Call(_func_ImVec4_destroy_, []interface{}{&vec4})
}

// CIMGUI_API ImVec4* ImVec4_ImVec4_Float(float _x,float _y,float _z,float _w);
func NewImVec4Float(x, y, z, w float32) *ImVec4 {
	p := giLib.Call(_func_ImVec4_ImVec4_Float_, []interface{}{&x, &y, &z, &w})
	return (*ImVec4)(p.PtrFree())
}

// Context creation and access
// - Each context create its own ImFontAtlas by default. You may instance one yourself and pass it to CreateContext() to share a font atlas between contexts.
// - DLL users: heaps and globals are not shared across DLL boundaries! You will need to call SetCurrentContext() + SetAllocatorFunctions()
//   for each static/DLL boundary you are calling from. Read "Context and Memory Allocators" section of imgui.cpp for details.

// CIMGUI_API ImGuiContext* igCreateContext(ImFontAtlas* shared_font_atlas);
// IMGUI_API ImGuiContext* CreateContext(ImFontAtlas* shared_font_atlas = NULL);
func CreateContext(font *ImFontAtlas) *ImGuiContext {
	return (*ImGuiContext)(giLib.Call(_func_igCreateContext_, []interface{}{&font}).PtrFree())

}

// CIMGUI_API ImGuiContext* igCreateContext(ImFontAtlas* shared_font_atlas);
// IMGUI_API ImGuiContext* CreateContext(ImFontAtlas* shared_font_atlas = NULL);
func CreateContextDefault() *ImGuiContext {
	return (*ImGuiContext)(giLib.Call(_func_igCreateContext_, []interface{}{_ptr_zero}).PtrFree())

}

// CIMGUI_API void igDestroyContext(ImGuiContext* ctx);
// IMGUI_API void DestroyContext(ImGuiContext* ctx = NULL);
// NULL = destroy current context
func DestroyContext(ctx *ImGuiContext) {
	giLib.Call(_func_igDestroyContext_, []interface{}{&ctx})
}

// CIMGUI_API void igDestroyContext(ImGuiContext* ctx);
// IMGUI_API void DestroyContext(ImGuiContext* ctx = NULL);
// NULL = destroy current context
func DestroyContextDefault() {
	giLib.Call(_func_igDestroyContext_, []interface{}{_ptr_zero})
}

// CIMGUI_API ImGuiContext* igGetCurrentContext(void);
// IMGUI_API ImGuiContext* GetCurrentContext();
func GetCurrentContext() *ImGuiContext {
	p := giLib.Call(_func_igGetCurrentContext_, nil)
	return (*ImGuiContext)(p.PtrFree())
}

// CIMGUI_API void igSetCurrentContext(ImGuiContext* ctx);
// IMGUI_API void SetCurrentContext(ImGuiContext* ctx);
func SetCurrentContext(ctx *ImGuiContext) {
	giLib.Call(_func_igSetCurrentContext_, []interface{}{&ctx})
}

// CIMGUI_API ImGuiIO* igGetIO(void);
// IMGUI_API ImGuiIO& GetIO();
// access the IO structure (mouse/keyboard/gamepad inputs, time, various configuration options/flags)
func GetIO() *ImGuiIO {
	return (*ImGuiIO)(giLib.Call(_func_igGetIO_, nil).PtrFree())
}

// CIMGUI_API ImGuiStyle* igGetStyle(void);
// IMGUI_API ImGuiStyle& GetStyle();
// access the Style structure (colors, sizes). Always use PushStyleCol(), PushStyleVar() to modify style mid-frame!
func GetStyle() *ImGuiStyle {
	i := giLib.Call(_func_igGetStyle_, nil)
	return (*ImGuiStyle)(i.PtrFree())
}

// CIMGUI_API void igNewFrame(void);
// IMGUI_API void NewFrame();
// start a new Dear ImGui frame, you can submit any command from this point until Render()/EndFrame().
func NewFrame() { giLib.Call(_func_igNewFrame_, nil) }

// CIMGUI_API void igEndFrame(void);
// IMGUI_API void EndFrame();
// ends the Dear ImGui frame. automatically called by Render(). If you don't need to render data (skipping rendering) you may call EndFrame() without Render()... but you'll have wasted CPU already! If you don't need to render, better to not create any windows and not call NewFrame() at all!
func EndFrame() { giLib.Call(_func_igEndFrame_, nil) }

// CIMGUI_API void igRender(void);
// IMGUI_API void Render();
// ends the Dear ImGui frame, finalize the draw data. You can then get call GetDrawData().
func Render() { giLib.Call(_func_igRender_, nil) }

// CIMGUI_API ImDrawData* igGetDrawData(void);
// IMGUI_API ImDrawData* GetDrawData();
// valid after Render() and until the next call to NewFrame(). this is what you have to render.
func GetDrawData() *ImDrawData {
	return (*ImDrawData)(giLib.Call(_func_igGetDrawData_, nil).PtrFree())
}

func cboolPtr(b *bool) *int32 {
	if b == nil {
		return nil
	}
	i := c.CBool(*b)
	return &i
}
func cboolGet(goptr *bool, cptr *int32) {
	if cptr == nil {
		return
	}
	(*goptr) = c.GoBool(*cptr)
}

// CIMGUI_API void igShowDemoWindow(bool* p_open);
// IMGUI_API void ShowDemoWindow(bool* p_open = NULL);
// create Demo window. demonstrate most ImGui features. call this to learn about the library! try to make it always available in your application!
func ShowDemoWindow(open *bool) {
	op := cboolPtr(open)
	giLib.Call(_func_igShowDemoWindow_, []interface{}{&op})
	cboolGet(open, op)
}

// CIMGUI_API void igShowDemoWindow(bool* p_open);
// IMGUI_API void ShowDemoWindow(bool* p_open = NULL);
// create Demo window. demonstrate most ImGui features. call this to learn about the library! try to make it always available in your application!
func ShowDemoWindowDefault() {
	giLib.Call(_func_igShowDemoWindow_, []interface{}{_ptr_zero})
}

// CIMGUI_API void igShowMetricsWindow(bool* p_open);
// IMGUI_API void ShowMetricsWindow(bool* p_open = NULL);
// create Metrics/Debugger window. display Dear ImGui internals: windows, draw commands, various internal state, etc.
func ShowMetricsWindow(open *bool) {
	op := cboolPtr(open)
	giLib.Call(_func_igShowMetricsWindow_, []interface{}{&op})
	cboolGet(open, op)
}

// CIMGUI_API void igShowMetricsWindow(bool* p_open);
// IMGUI_API void ShowMetricsWindow(bool* p_open = NULL);
// create Metrics/Debugger window. display Dear ImGui internals: windows, draw commands, various internal state, etc.
func ShowMetricsWindowDefault() {
	giLib.Call(_func_igShowMetricsWindow_, []interface{}{_ptr_zero})
}

// CIMGUI_API void igShowDebugLogWindow(bool* p_open);
// IMGUI_API void ShowDebugLogWindow(bool* p_open = NULL);
// create Debug Log window. display a simplified log of important dear imgui events.
func ShowDebugLogWindow(open *bool) {
	op := cboolPtr(open)
	giLib.Call(_func_igShowDebugLogWindow_, []interface{}{&op})
	cboolGet(open, op)
}

// CIMGUI_API void igShowDebugLogWindow(bool* p_open);
// IMGUI_API void ShowDebugLogWindow(bool* p_open = NULL);
// create Debug Log window. display a simplified log of important dear imgui events.
func ShowDebugLogWindowDefault() {
	giLib.Call(_func_igShowDebugLogWindow_, []interface{}{_ptr_zero})
}

// CIMGUI_API void igShowStackToolWindow(bool* p_open);
// IMGUI_API void ShowStackToolWindow(bool* p_open = NULL);
// create Stack Tool window. hover items with mouse to query information about the source of their unique ID.
func ShowStackToolWindow(open *bool) {
	op := cboolPtr(open)
	giLib.Call(_func_igShowStackToolWindow_, []interface{}{&op})
	cboolGet(open, op)
}

// CIMGUI_API void igShowAboutWindow(bool* p_open);
// IMGUI_API void ShowAboutWindow(bool* p_open = NULL);
// create About window. display Dear ImGui version, credits and build/system information.
func ShowAboutWindow(open *bool) {
	op := cboolPtr(open)
	giLib.Call(_func_igShowAboutWindow_, []interface{}{&op})
	cboolGet(open, op)
}

// CIMGUI_API void igShowAboutWindow(bool* p_open);
// IMGUI_API void ShowAboutWindow(bool* p_open = NULL);
// create About window. display Dear ImGui version, credits and build/system information.
func ShowAboutWindowDefault() {
	giLib.Call(_func_igShowAboutWindow_, []interface{}{_ptr_zero})
}

// CIMGUI_API void igShowStyleEditor(ImGuiStyle* ref);
// IMGUI_API void ShowStyleEditor(ImGuiStyle* ref = NULL);
// add style editor block (not a window). you can pass in a reference ImGuiStyle structure to compare to, revert to and save to (else it uses the default style)
func ShowStyleEditor(ref *ImGuiStyle) {
	giLib.Call(_func_igShowStyleEditor_, []interface{}{&ref})
}

// CIMGUI_API void igShowStyleEditor(ImGuiStyle* ref);
// IMGUI_API void ShowStyleEditor(ImGuiStyle* ref = NULL);
// add style editor block (not a window). you can pass in a reference ImGuiStyle structure to compare to, revert to and save to (else it uses the default style)
func ShowStyleEditorDefault() {
	giLib.Call(_func_igShowStyleEditor_, []interface{}{_ptr_zero})
}

// CIMGUI_API bool igShowStyleSelector(const char* label);
// IMGUI_API bool ShowStyleSelector(const char* label);
// add style selector block (not a window), essentially a combo listing the default styles.
func ShowStyleSelector(label string) bool {
	l := c.CStr(label)
	defer usf.Free(l)
	return giLib.Call(_func_igShowStyleSelector_, []interface{}{&l}).BoolFree()
}

// CIMGUI_API void igShowFontSelector(const char* label);
// IMGUI_API void ShowFontSelector(const char* label);
// add font selector block (not a window), essentially a combo listing the loaded fonts.
func ShowFontSelector(label string) {
	l := c.CStr(label)
	defer usf.Free(l)
	giLib.Call(_func_igShowFontSelector_, []interface{}{&l})
}

// CIMGUI_API void igShowUserGuide(void);
// IMGUI_API void ShowUserGuide();
// add basic help/info block (not a window): how to manipulate ImGui as an end-user (mouse/keyboard controls).
func ShowUserGuide() {
	giLib.Call(_func_igShowUserGuide_, nil)
}

// CIMGUI_API const char* igGetVersion(void);
// IMGUI_API const char* GetVersion();
// get the compiled version string e.g. "1.80 WIP" (essentially the value for IMGUI_VERSION from the compiled version of imgui.cpp)
func GetVersion() string {
	return giLib.Call(_func_igGetVersion_, nil).StrFree()
}

// CIMGUI_API void igStyleColorsDark(ImGuiStyle* dst);
// IMGUI_API void StyleColorsDark(ImGuiStyle* dst = NULL);
// new, recommended style (default)
func StyleColorsDark(dst *ImGuiStyle) {
	giLib.Call(_func_igStyleColorsDark_, []interface{}{&dst})
}

// CIMGUI_API void igStyleColorsLight(ImGuiStyle* dst);
// IMGUI_API void StyleColorsLight(ImGuiStyle* dst = NULL);
// best used with borders and a custom, thicker font
func StyleColorsLight(dst *ImGuiStyle) {
	giLib.Call(_func_igStyleColorsLight_, []interface{}{&dst})
}

// CIMGUI_API void igStyleColorsClassic(ImGuiStyle* dst);
// IMGUI_API void StyleColorsClassic(ImGuiStyle* dst = NULL);
// classic imgui style
func StyleColorsClassic(dst *ImGuiStyle) {
	giLib.Call(_func_igStyleColorsClassic_, []interface{}{&dst})
}

// Windows
// - Begin() = push window to the stack and start appending to it. End() = pop window from the stack.
// - Passing 'bool* p_open != NULL' shows a window-closing widget in the upper-right corner of the window,
//   which clicking will set the boolean to false when clicked.
// - You may append multiple times to the same window during the same frame by calling Begin()/End() pairs multiple times.
//   Some information such as 'flags' or 'p_open' will only be considered by the first call to Begin().
// - Begin() return false to indicate the window is collapsed or fully clipped, so you may early out and omit submitting
//   anything to the window. Always call a matching End() for each Begin() call, regardless of its return value!
//   [Important: due to legacy reason, this is inconsistent with most other functions such as BeginMenu/EndMenu,
//    BeginPopup/EndPopup, etc. where the EndXXX call should only be called if the corresponding BeginXXX function
//    returned true. Begin and BeginChild are the only odd ones out. Will be fixed in a future update.]
// - Note that the bottom of window stack always contains a window called "Debug".

// CIMGUI_API bool igBegin(const char* name,bool* p_open,ImGuiWindowFlags flags);
// IMGUI_API bool Begin(const char* name, bool* p_open = NULL, ImGuiWindowFlags flags = 0);
func Begin(name string, open *bool, flags ImGuiWindowFlags) bool {
	n := c.CStr(name)
	defer usf.Free(n)
	op := cboolPtr(open)
	p := giLib.Call(_func_igBegin_, []interface{}{&n, &op, &flags}).BoolFree()
	cboolGet(open, op)
	return p
}

// CIMGUI_API bool igBegin(const char* name,bool* p_open,ImGuiWindowFlags flags);
// IMGUI_API bool Begin(const char* name, bool* p_open = NULL, ImGuiWindowFlags flags = 0);
func BeginDefault(name string) bool {
	n := c.CStr(name)
	defer usf.Free(n)
	return giLib.Call(_func_igBegin_, []interface{}{&n, _ptr_zero, _u32_zero}).BoolFree()
}

// CIMGUI_API void igEnd(void);
// IMGUI_API void End();
func End() {
	giLib.Call(_func_igEnd_, nil)
}

// Child Windows
// - Use child windows to begin into a self-contained independent scrolling/clipping regions within a host window. Child windows can embed their own child.
// - For each independent axis of 'size': ==0.0f: use remaining host window size / >0.0f: fixed size / <0.0f: use remaining window size minus abs(size) / Each axis can use a different mode, e.g. ImVec2(0,400).
// - BeginChild() returns false to indicate the window is collapsed or fully clipped, so you may early out and omit submitting anything to the window.
//   Always call a matching EndChild() for each BeginChild() call, regardless of its return value.
//   [Important: due to legacy reason, this is inconsistent with most other functions such as BeginMenu/EndMenu,
//    BeginPopup/EndPopup, etc. where the EndXXX call should only be called if the corresponding BeginXXX function
//    returned true. Begin and BeginChild are the only odd ones out. Will be fixed in a future update.]

// CIMGUI_API bool igBeginChild_Str(const char* str_id,const ImVec2 size,bool border,ImGuiWindowFlags flags);
// IMGUI_API bool BeginChild(const char* str_id, const ImVec2& size = ImVec2(0, 0), bool border = false, ImGuiWindowFlags flags = 0);
func BeginChildByStrVec2(str_id string, size ImVec2, border bool, flags ImGuiWindowFlags) bool {
	s, b := c.CStr(str_id), c.CBool(border)
	defer usf.Free(s)

	return giLib.Call(_func_igBeginChild_Str_, []interface{}{&s, &size, &b, &flags}).BoolFree()

}

// CIMGUI_API bool igBeginChild_Str(const char* str_id,const ImVec2 size,bool border,ImGuiWindowFlags flags);
// IMGUI_API bool BeginChild(const char* str_id, const ImVec2& size = ImVec2(0, 0), bool border = false, ImGuiWindowFlags flags = 0);
func BeginChildByStr(str_id string, sizeW, sizeH float32, border bool, flags ImGuiWindowFlags) bool {
	s, b, size := c.CStr(str_id), c.CBool(border), NewImVec2Float(sizeW, sizeH)
	defer usf.Free(s)
	defer size.Destroy()

	return giLib.Call(_func_igBeginChild_Str_, []interface{}{&s, &size, &b, &flags}).BoolFree()

}

// CIMGUI_API bool igBeginChild_Str(const char* str_id,const ImVec2 size,bool border,ImGuiWindowFlags flags);
// IMGUI_API bool BeginChild(const char* str_id, const ImVec2& size = ImVec2(0, 0), bool border = false, ImGuiWindowFlags flags = 0);
func BeginChildByStrDefault(str_id string) bool {
	s := c.CStr(str_id)
	defer usf.Free(s)

	return giLib.Call(_func_igBeginChild_Str_, []interface{}{&s, _vec2_zero, _bool_false, _u32_zero}).BoolFree()

}

// CIMGUI_API bool igBeginChild_ID(ImGuiID id,const ImVec2 size,bool border,ImGuiWindowFlags flags);
// IMGUI_API bool BeginChild(const char* str_id, const ImVec2& size = ImVec2(0, 0), bool border = false, ImGuiWindowFlags flags = 0);
func BeginChildVec2(id ImGuiID, size ImVec2, border bool, flags ImGuiWindowFlags) bool {
	b := c.CBool(border)
	return giLib.Call(_func_igBeginChild_ID_, []interface{}{&id, &size, &b, &flags}).BoolFree()
}

// CIMGUI_API bool igBeginChild_ID(ImGuiID id,const ImVec2 size,bool border,ImGuiWindowFlags flags);
// IMGUI_API bool BeginChild(ImGuiID id, const ImVec2& size = ImVec2(0, 0), bool border = false, ImGuiWindowFlags flags = 0);
func BeginChild(id ImGuiID, sizeW, sizeH float32, border bool, flags ImGuiWindowFlags) bool {
	b, size := c.CBool(border), NewImVec2Float(sizeW, sizeH)
	defer size.Destroy()
	return giLib.Call(_func_igBeginChild_ID_, []interface{}{&id, &size, &b, &flags}).BoolFree()
}

// CIMGUI_API bool igBeginChild_ID(ImGuiID id,const ImVec2 size,bool border,ImGuiWindowFlags flags);
// IMGUI_API bool BeginChild(ImGuiID id, const ImVec2& size = ImVec2(0, 0), bool border = false, ImGuiWindowFlags flags = 0);
func BeginChildDefault(id ImGuiID) bool {
	return giLib.Call(_func_igBeginChild_ID_, []interface{}{&id, _vec2_zero, _bool_false, _i32_zero}).BoolFree()

}

// CIMGUI_API void igEndChild(void);
// IMGUI_API void EndChild();
func EndChild() {
	giLib.Call(_func_igEndChild_, nil)
}

// Windows Utilities
// - 'current window' = the window we are appending into while inside a Begin()/End() block. 'next window' = next window we will Begin() into.

// CIMGUI_API bool igIsWindowAppearing(void);
// IMGUI_API bool IsWindowAppearing();
func IsWindowAppearing() bool {
	return giLib.Call(_func_igIsWindowAppearing_, nil).BoolFree()
}

// CIMGUI_API bool igIsWindowCollapsed(void);
// IMGUI_API bool IsWindowCollapsed();
func IsWindowCollapsed() bool {
	return giLib.Call(_func_igIsWindowCollapsed_, nil).BoolFree()
}

// CIMGUI_API bool igIsWindowFocused(ImGuiFocusedFlags flags);
// IMGUI_API bool IsWindowFocused(ImGuiFocusedFlags flags=0);
// is current window focused? or its root/child, depending on flags. see flags for options.
func IsWindowFocused(flags ImGuiFocusedFlags) bool {
	return giLib.Call(_func_igIsWindowFocused_, []interface{}{&flags}).BoolFree()
}

// CIMGUI_API bool igIsWindowHovered(ImGuiHoveredFlags flags);
// IMGUI_API bool IsWindowHovered(ImGuiHoveredFlags flags=0);
// is current window hovered (and typically: not blocked by a popup/modal)? see flags for options. NB: If you are trying to check whether your mouse should be dispatched to imgui or to your app, you should use the 'io.WantCaptureMouse' boolean for that! Please read the FAQ!
func IsWindowHovered(flags ImGuiHoveredFlags) bool {
	return giLib.Call(_func_igIsWindowHovered_, []interface{}{&flags}).BoolFree()
}

// CIMGUI_API ImDrawList* igGetWindowDrawList(void);
// IMGUI_API ImDrawList* GetWindowDrawList();
// get draw list associated to the current window, to append your own drawing primitives
func GetWindowDrawList() *ImDrawList {
	return (*ImDrawList)(giLib.Call(_func_igGetWindowDrawList_, nil).PtrFree())
}

// CIMGUI_API float igGetWindowDpiScale(void);
// IMGUI_API float GetWindowDpiScale();
// get DPI scale currently associated to the current window's viewport.
func GetWindowDpiScale() float32 {
	return giLib.Call(_func_igGetWindowDpiScale_, nil).F32Free()
}

// CIMGUI_API void igGetWindowPos(ImVec2* pOut);
// IMGUI_API ImVec2 GetWindowPos();
// get current window position in screen space (useful if you want to do your own drawing via the DrawList API)
func GetWindowPosVec2() *ImVec2 {
	vec2 := NewImVec2()
	giLib.Call(_func_igGetWindowPos_, []interface{}{&vec2})
	return vec2
}

// CIMGUI_API void igGetWindowPos(ImVec2* pOut);
// IMGUI_API ImVec2 GetWindowPos();
// get current window position in screen space (useful if you want to do your own drawing via the DrawList API)
func GetWindowPos() (x float32, y float32) {
	vec2 := NewImVec2()
	giLib.Call(_func_igGetWindowPos_, []interface{}{&vec2})
	x, y = vec2.X, vec2.Y
	vec2.Destroy()
	return
}

// CIMGUI_API void igGetWindowSize(ImVec2* pOut);
// IMGUI_API ImVec2 GetWindowSize();
// get current window size
func GetWindowSizeVec2() *ImVec2 {
	vec2 := NewImVec2()
	giLib.Call(_func_igGetWindowSize_, []interface{}{&vec2})
	return vec2
}

// CIMGUI_API void igGetWindowSize(ImVec2* pOut);
// IMGUI_API ImVec2 GetWindowSize();
// get current window size
func GetWindowSize() (w float32, h float32) {
	vec2 := NewImVec2()
	giLib.Call(_func_igGetWindowSize_, []interface{}{&vec2})
	w, h = vec2.X, vec2.Y
	vec2.Destroy()
	return
}

// CIMGUI_API float igGetWindowWidth(void);
// IMGUI_API float  GetWindowWidth();
// get current window width (shortcut for GetWindowSize().x)
func GetWindowWidth() float32 {
	return giLib.Call(_func_igGetWindowWidth_, nil).F32Free()
}

// CIMGUI_API float igGetWindowHeight(void);
// IMGUI_API float GetWindowHeight();
// get current window height (shortcut for GetWindowSize().y)
func GetWindowHeight() float32 {
	return giLib.Call(_func_igGetWindowHeight_, nil).F32Free()
}

// CIMGUI_API ImGuiViewport* igGetWindowViewport(void);
// IMGUI_API ImGuiViewport*GetWindowViewport();
// get viewport currently associated to the current window.
func GetWindowViewport() *ImGuiViewport {
	return (*ImGuiViewport)(giLib.Call(_func_igGetWindowViewport_, nil).PtrFree())
}

// Window manipulation
// - Prefer using SetNextXXX functions (before Begin) rather that SetXXX functions (after Begin).

// CIMGUI_API void igSetNextWindowPos(const ImVec2 pos,ImGuiCond cond,const ImVec2 pivot);
// IMGUI_API void SetNextWindowPos(const ImVec2& pos, ImGuiCond cond = 0, const ImVec2& pivot = ImVec2(0, 0));
// set next window position. call before Begin(). use pivot=(0.5f,0.5f) to center on given point, etc.
func SetNextWindowPosVec2(pos ImVec2, cond ImGuiCond, pivot ImVec2) {
	giLib.Call(_func_igSetNextWindowPos_, []interface{}{&pos, &cond, &pivot})
}

// CIMGUI_API void igSetNextWindowPos(const ImVec2 pos,ImGuiCond cond,const ImVec2 pivot);
// IMGUI_API void SetNextWindowPos(const ImVec2& pos, ImGuiCond cond = 0, const ImVec2& pivot = ImVec2(0, 0));
// set next window position. call before Begin(). use pivot=(0.5f,0.5f) to center on given point, etc.
func SetNextWindowPos(posX float32, posY float32, cond ImGuiCond, pivotX float32, pivotY float32) {
	pos, pivot := NewImVec2Float(posX, posY), NewImVec2Float(pivotX, pivotY)
	defer pos.Destroy()
	defer pivot.Destroy()
	giLib.Call(_func_igSetNextWindowPos_, []interface{}{pos, &cond, pivot})
}

// CIMGUI_API void igSetNextWindowPos(const ImVec2 pos,ImGuiCond cond,const ImVec2 pivot);
// IMGUI_API void SetNextWindowPos(const ImVec2& pos, ImGuiCond cond = 0, const ImVec2& pivot = ImVec2(0, 0));
// set next window position. call before Begin(). use pivot=(0.5f,0.5f) to center on given point, etc.
func SetNextWindowPosDefault(posX float32, posY float32) {
	pos := NewImVec2Float(posX, posY)
	defer pos.Destroy()
	giLib.Call(_func_igSetNextWindowPos_, []interface{}{pos, _u32_zero, _vec2_zero})
}

// CIMGUI_API void igSetNextWindowSize(const ImVec2 size,ImGuiCond cond);
// IMGUI_API void SetNextWindowSize(const ImVec2& size, ImGuiCond cond = 0);
// set next window size. set axis to 0.0f to force an auto-fit on this axis. call before Begin()
func SetNextWindowSizeVec2(size ImVec2, cond ImGuiCond) {
	giLib.Call(_func_igSetNextWindowSize_, []interface{}{&size, &cond})
}

// CIMGUI_API void igSetNextWindowSize(const ImVec2 size,ImGuiCond cond);
// IMGUI_API void SetNextWindowSize(const ImVec2& size, ImGuiCond cond = 0);
// set next window size. set axis to 0.0f to force an auto-fit on this axis. call before Begin()
func SetNextWindowSize(sizeW float32, sizeH float32, cond ImGuiCond) {
	size := NewImVec2Float(sizeW, sizeH)
	defer size.Destroy()
	giLib.Call(_func_igSetNextWindowSize_, []interface{}{size, &cond})
}

// CIMGUI_API void igSetNextWindowSizeConstraints(const ImVec2 size_min,const ImVec2 size_max,ImGuiSizeCallback custom_callback,void* custom_callback_data);
// IMGUI_API void SetNextWindowSizeConstraints(const ImVec2& size_min, const ImVec2& size_max, ImGuiSizeCallback custom_callback = NULL, void* custom_callback_data = NULL);
// set next window size limits. use -1,-1 on either X/Y axis to preserve the current size. Sizes will be rounded down. Use callback to apply non-trivial programmatic constraints.
func SetNextWindowSizeConstraintsVec2(size_min ImVec2, size_max ImVec2, callback *c.Callback) {
	fn, data := unsafe.Pointer(nil), unsafe.Pointer(nil)
	if callback != nil {
		if callback.CallbackCvt == nil {
			callback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				p := (*ImGuiSizeCallbackData)(args[0].Ptr())
				fn := cb.CallbackFunc.(func(*ImGuiSizeCallbackData))
				if fn != nil {
					fn(p)
				}
			}
		}
		fn = callback.Cfunc
	}
	giLib.Call(_func_igSetNextWindowSizeConstraints_, []interface{}{&size_min, &size_max, &fn, &data})
}

// CIMGUI_API void igSetNextWindowSizeConstraints(const ImVec2 size_min,const ImVec2 size_max,ImGuiSizeCallback custom_callback,void* custom_callback_data);
// IMGUI_API void SetNextWindowSizeConstraints(const ImVec2& size_min, const ImVec2& size_max, ImGuiSizeCallback custom_callback = NULL, void* custom_callback_data = NULL);
// set next window size limits. use -1,-1 on either X/Y axis to preserve the current size. Sizes will be rounded down. Use callback to apply non-trivial programmatic constraints.
func SetNextWindowSizeConstraints(minX, minY float32, maxX, maxY float32, callback *c.Callback) {
	min, max := NewImVec2Float(minX, minY), NewImVec2Float(maxX, maxY)
	defer min.Destroy()
	defer max.Destroy()

	fn, data := unsafe.Pointer(nil), unsafe.Pointer(nil)
	if callback != nil {
		if callback.CallbackCvt == nil {
			callback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				p := (*ImGuiSizeCallbackData)(args[0].Ptr())
				fn := cb.CallbackFunc.(func(*ImGuiSizeCallbackData))
				if fn != nil {
					fn(p)
				}
			}
		}
		fn = callback.Cfunc
	}
	giLib.Call(_func_igSetNextWindowSizeConstraints_, []interface{}{min, max, &fn, &data})
}

// CIMGUI_API void igSetNextWindowContentSize(const ImVec2 size);
// IMGUI_API void SetNextWindowContentSize(const ImVec2& size);
// set next window content size (~ scrollable client area, which enforce the range of scrollbars). Not including window decorations (title bar, menu bar, etc.) nor WindowPadding. set an axis to 0.0f to leave it automatic. call before Begin()
func SetNextWindowContentSizeVec2(size ImVec2) {
	giLib.Call(_func_igSetNextWindowContentSize_, []interface{}{&size})
}

// CIMGUI_API void igSetNextWindowContentSize(const ImVec2 size);
// IMGUI_API void SetNextWindowContentSize(const ImVec2& size);
// set next window content size (~ scrollable client area, which enforce the range of scrollbars). Not including window decorations (title bar, menu bar, etc.) nor WindowPadding. set an axis to 0.0f to leave it automatic. call before Begin()
func SetNextWindowContentSize(w, h float32) {
	size := NewImVec2Float(w, h)
	defer size.Destroy()
	giLib.Call(_func_igSetNextWindowContentSize_, []interface{}{size})
}

// CIMGUI_API void igSetNextWindowCollapsed(bool collapsed,ImGuiCond cond);
// IMGUI_API void SetNextWindowCollapsed(bool collapsed, ImGuiCond cond = 0);
// set next window collapsed state. call before Begin()
func SetNextWindowCollapsed(collapsed bool, cond ImGuiCond) {
	col := c.CBool(collapsed)
	giLib.Call(_func_igSetNextWindowCollapsed_, []interface{}{&col, &cond})
}

// CIMGUI_API void igSetNextWindowFocus(void);
// IMGUI_API void SetNextWindowFocus();
// set next window to be focused / top-most. call before Begin()
func SetNextWindowFocus() {
	giLib.Call(_func_igSetNextWindowFocus_, nil)
}

// CIMGUI_API void igSetNextWindowScroll(const ImVec2 scroll);
// IMGUI_API void SetNextWindowScroll(const ImVec2& scroll);
// set next window scrolling value (use < 0.0f to not affect a given axis).
func SetNextWindowScrollVec2(scroll ImVec2) {
	giLib.Call(_func_igSetNextWindowScroll_, []interface{}{&scroll})
}

// CIMGUI_API void igSetNextWindowScroll(const ImVec2 scroll);
// IMGUI_API void SetNextWindowScroll(const ImVec2& scroll);
// set next window scrolling value (use < 0.0f to not affect a given axis).
func SetNextWindowScroll(v, h float32) {
	size := NewImVec2Float(v, h)
	defer size.Destroy()
	giLib.Call(_func_igSetNextWindowScroll_, []interface{}{size})
}

// CIMGUI_API void igSetNextWindowBgAlpha(float alpha);
// IMGUI_API void SetNextWindowBgAlpha(float alpha);
// set next window background color alpha. helper to easily override the Alpha component of ImGuiCol_WindowBg/ChildBg/PopupBg. you may also use ImGuiWindowFlags_NoBackground.
func SetNextWindowBgAlpha(alpha float32) {
	giLib.Call(_func_igSetNextWindowBgAlpha_, []interface{}{&alpha})
}

// CIMGUI_API void igSetNextWindowViewport(ImGuiID viewport_id);
// IMGUI_API void SetNextWindowViewport(ImGuiID viewport_id);
// set next window viewport
func SetNextWindowViewport(viewport_id ImGuiID) {
	giLib.Call(_func_igSetNextWindowViewport_, []interface{}{&viewport_id})
}

// CIMGUI_API void igSetWindowPos_Vec2(const ImVec2 pos,ImGuiCond cond);
// IMGUI_API void SetWindowPos(const ImVec2& pos, ImGuiCond cond = 0);
// (not recommended) set current window position - call within Begin()/End(). prefer using SetNextWindowPos(), as this may incur tearing and side-effects.
func SetWindowPosVec2(pos ImVec2, cond ImGuiCond) {
	giLib.Call(_func_igSetWindowPos_Vec2_, []interface{}{&pos, &cond})
}

// CIMGUI_API void igSetWindowPos_Vec2(const ImVec2 pos,ImGuiCond cond);
// IMGUI_API void SetWindowPos(const ImVec2& pos, ImGuiCond cond = 0);
// (not recommended) set current window position - call within Begin()/End(). prefer using SetNextWindowPos(), as this may incur tearing and side-effects.
func SetWindowPos(x, y float32, cond ImGuiCond) {
	size := NewImVec2Float(x, y)
	defer size.Destroy()
	giLib.Call(_func_igSetWindowPos_Vec2_, []interface{}{size, &cond})
}

// CIMGUI_API void igSetWindowSize_Vec2(const ImVec2 size,ImGuiCond cond);
// IMGUI_API void SetWindowSize(const ImVec2& size, ImGuiCond cond = 0);
// (not recommended) set current window size - call within Begin()/End(). set to ImVec2(0, 0) to force an auto-fit. prefer using SetNextWindowSize(), as this may incur tearing and minor side-effects.
func SetWindowSizeVec2(size ImVec2, cond ImGuiCond) {
	giLib.Call(_func_igSetWindowSize_Vec2_, []interface{}{&size, &cond})
}

// CIMGUI_API void igSetWindowSize_Vec2(const ImVec2 size,ImGuiCond cond);
// IMGUI_API void SetWindowSize(const ImVec2& size, ImGuiCond cond = 0);
// (not recommended) set current window size - call within Begin()/End(). set to ImVec2(0, 0) to force an auto-fit. prefer using SetNextWindowSize(), as this may incur tearing and minor side-effects.
func SetWindowSize(x float32, y float32, cond ImGuiCond) {
	size := NewImVec2Float(x, y)
	defer size.Destroy()
	giLib.Call(_func_igSetWindowSize_Vec2_, []interface{}{size, &cond})
}

// CIMGUI_API void igSetWindowCollapsed_Bool(bool collapsed,ImGuiCond cond);
// IMGUI_API void  SetWindowCollapsed(bool collapsed, ImGuiCond cond = 0);
// (not recommended) set current window collapsed state. prefer using SetNextWindowCollapsed().
func SetWindowCollapsed(collapsed bool, cond ImGuiCond) {
	col := c.CBool(collapsed)
	giLib.Call(_func_igSetWindowCollapsed_Bool_, []interface{}{&col, &cond})
}

// CIMGUI_API void igSetWindowFocus_Nil(void);
// IMGUI_API void SetWindowFocus();
// (not recommended) set current window to be focused / top-most. prefer using SetNextWindowFocus().
func SetWindowFocus() {
	giLib.Call(_func_igSetWindowFocus_Nil_, nil)
}

// CIMGUI_API void igSetWindowFontScale(float scale);
// IMGUI_API void SetWindowFontScale(float scale);
// [OBSOLETE] set font scale. Adjust IO.FontGlobalScale if you want to scale all windows. This is an old API! For correct scaling, prefer to reload font + rebuild ImFontAtlas + call style.ScaleAllSizes().
func SetWindowFontScale(scale float32) {
	giLib.Call(_func_igSetWindowFontScale_, []interface{}{&scale})
}

// CIMGUI_API void igSetWindowPos_Str(const char* name,const ImVec2 pos,ImGuiCond cond);
// IMGUI_API void SetWindowPos(const char* name, const ImVec2& pos, ImGuiCond cond = 0);
// set named window position.
func SetWindowPosByNameVec2(name string, pos ImVec2, cond ImGuiCond) {
	n := c.CStr(name)
	defer usf.Free(n)
	giLib.Call(_func_igSetWindowPos_Str_, []interface{}{&n, pos, &cond})
}

// CIMGUI_API void igSetWindowPos_Str(const char* name,const ImVec2 pos,ImGuiCond cond);
// IMGUI_API void SetWindowPos(const char* name, const ImVec2& pos, ImGuiCond cond = 0);
// set named window position.
func SetWindowPosByName(name string, posX, posY float32, cond ImGuiCond) {
	n, pos := c.CStr(name), NewImVec2Float(posX, posY)
	defer usf.Free(n)
	defer pos.Destroy()
	giLib.Call(_func_igSetWindowPos_Str_, []interface{}{&n, pos, &cond})
}

// CIMGUI_API void igSetWindowSize_Str(const char* name,const ImVec2 size,ImGuiCond cond);
// IMGUI_API void SetWindowSize(const char* name, const ImVec2& size, ImGuiCond cond = 0);
// set named window size. set axis to 0.0f to force an auto-fit on this axis.
func SetWindowSizeByNameVec2(name string, size ImVec2, cond ImGuiCond) {
	n := c.CStr(name)
	defer usf.Free(n)
	giLib.Call(_func_igSetWindowSize_Str_, []interface{}{&n, &size, &cond})
}

// CIMGUI_API void igSetWindowSize_Str(const char* name,const ImVec2 size,ImGuiCond cond);
// IMGUI_API void SetWindowSize(const char* name, const ImVec2& size, ImGuiCond cond = 0);
// set named window size. set axis to 0.0f to force an auto-fit on this axis.
func SetWindowSizeByName(name string, sizeW, sizeH float32, cond ImGuiCond) {
	n, size := c.CStr(name), NewImVec2Float(sizeW, sizeH)
	defer usf.Free(n)
	defer size.Destroy()
	giLib.Call(_func_igSetWindowSize_Str_, []interface{}{&n, size, &cond})
}

// CIMGUI_API void igSetWindowCollapsed_Str(const char* name,bool collapsed,ImGuiCond cond);
// IMGUI_API void SetWindowCollapsed(const char* name, bool collapsed, ImGuiCond cond = 0);
// set named window collapsed state
func SetWindowCollapsedByName(name string, collapsed bool, cond ImGuiCond) {
	n, b := c.CStr(name), c.CBool(collapsed)
	defer usf.Free(n)
	giLib.Call(_func_igSetWindowCollapsed_Str_, []interface{}{&n, &b, &cond})
}

// CIMGUI_API void igSetWindowFocus_Str(const char* name);
// IMGUI_API void SetWindowFocus(const char* name);
// set named window to be focused / top-most. use NULL to remove focus.
func SetWindowFocusByName(name string) {
	n := c.CStr(name)
	defer usf.Free(n)
	giLib.Call(_func_igSetWindowFocus_Str_, []interface{}{&n})
}

// Content region
// - Retrieve available space from a given point. GetContentRegionAvail() is frequently useful.
// - Those functions are bound to be redesigned (they are confusing, incomplete and the Min/Max return values are in local window coordinates which increases confusion)

// CIMGUI_API void igGetContentRegionAvail(ImVec2* pOut);
// IMGUI_API ImVec2 GetContentRegionAvail();
// == GetContentRegionMax() - GetCursorPos()
func GetContentRegionAvailVec2() *ImVec2 {
	p := NewImVec2()
	giLib.Call(_func_igGetContentRegionAvail_, []interface{}{&p})
	return p
}

// CIMGUI_API void igGetContentRegionAvail(ImVec2* pOut);
// IMGUI_API ImVec2 GetContentRegionAvail();
// == GetContentRegionMax() - GetCursorPos()
func GetContentRegionAvail() (float32, float32) {
	p := NewImVec2()
	defer p.Destroy()
	giLib.Call(_func_igGetContentRegionAvail_, []interface{}{&p})
	return p.X, p.Y
}

// CIMGUI_API void igGetContentRegionMax(ImVec2* pOut);
// IMGUI_API ImVec2 GetContentRegionMax();
// current content boundaries (typically window boundaries including scrolling, or current column boundaries), in windows coordinates
func GetContentRegionMaxVec2() *ImVec2 {
	p := NewImVec2()
	giLib.Call(_func_igGetContentRegionMax_, []interface{}{&p})
	return p
}

// CIMGUI_API void igGetContentRegionMax(ImVec2* pOut);
// IMGUI_API ImVec2 GetContentRegionMax();
// current content boundaries (typically window boundaries including scrolling, or current column boundaries), in windows coordinates
func GetContentRegionMax() (float32, float32) {
	p := NewImVec2()
	defer p.Destroy()
	giLib.Call(_func_igGetContentRegionMax_, []interface{}{&p})
	return p.X, p.Y
}

// CIMGUI_API void igGetWindowContentRegionMin(ImVec2* pOut);
// IMGUI_API ImVec2 GetWindowContentRegionMin();
// content boundaries min for the full window (roughly (0,0)-Scroll), in window coordinates
func GetWindowContentRegionMinVec2() *ImVec2 {
	p := NewImVec2()
	giLib.Call(_func_igGetWindowContentRegionMin_, []interface{}{&p})
	return p
}

// CIMGUI_API void igGetWindowContentRegionMin(ImVec2* pOut);
// IMGUI_API ImVec2 GetWindowContentRegionMin();
// content boundaries min for the full window (roughly (0,0)-Scroll), in window coordinates
func GetWindowContentRegionMin() (float32, float32) {
	p := NewImVec2()
	defer p.Destroy()
	giLib.Call(_func_igGetWindowContentRegionMin_, []interface{}{&p})
	return p.X, p.Y
}

// CIMGUI_API void igGetWindowContentRegionMax(ImVec2* pOut);
// IMGUI_API ImVec2 GetWindowContentRegionMax();
// content boundaries max for the full window (roughly (0,0)+Size-Scroll) where Size can be overridden with SetNextWindowContentSize(), in window coordinates
func GetWindowContentRegionMaxVec2() *ImVec2 {
	p := NewImVec2()
	giLib.Call(_func_igGetWindowContentRegionMax_, []interface{}{&p})
	return p
}

// CIMGUI_API void igGetWindowContentRegionMax(ImVec2* pOut);
// IMGUI_API ImVec2 GetWindowContentRegionMax();
// content boundaries max for the full window (roughly (0,0)+Size-Scroll) where Size can be overridden with SetNextWindowContentSize(), in window coordinates
func GetWindowContentRegionMax() (float32, float32) {
	p := NewImVec2()
	defer p.Destroy()
	giLib.Call(_func_igGetWindowContentRegionMax_, []interface{}{&p})
	return p.X, p.Y
}

// Windows Scrolling
// - Any change of Scroll will be applied at the beginning of next frame in the first call to Begin().
// - You may instead use SetNextWindowScroll() prior to calling Begin() to avoid this delay, as an alternative to using SetScrollX()/SetScrollY().

// CIMGUI_API float igGetScrollX(void);
// IMGUI_API float GetScrollX();
// get scrolling amount [0 .. GetScrollMaxX()]
func GetScrollX() float32 {
	return giLib.Call(_func_igGetScrollX_, nil).F32Free()
}

// CIMGUI_API float igGetScrollY(void);
// IMGUI_API float GetScrollY();
// get scrolling amount [0 .. GetScrollMaxY()]
func GetScrollY() float32 {
	return giLib.Call(_func_igGetScrollY_, nil).F32Free()
}

// CIMGUI_API void igSetScrollX_Float(float scroll_x);
// IMGUI_API void SetScrollX(float scroll_x);
// set scrolling amount [0 .. GetScrollMaxX()]
func SetScrollX(x float32) {
	giLib.Call(_func_igSetScrollX_Float_, []interface{}{&x})
}

// CIMGUI_API void igSetScrollY_Float(float scroll_y);
// IMGUI_API void SetScrollY(float scroll_y);
// set scrolling amount [0 .. GetScrollMaxY()]
func SetScrollY(y float32) {
	giLib.Call(_func_igSetScrollY_Float_, []interface{}{&y})
}

// CIMGUI_API float igGetScrollMaxX(void);
func GetScrollMaxX() float32 {
	return giLib.Call(_func_igGetScrollMaxX_, nil).F32Free()
}

// CIMGUI_API float igGetScrollMaxY(void);
func GetScrollMaxY() float32 {
	return giLib.Call(_func_igGetScrollMaxY_, nil).F32Free()
}

// CIMGUI_API void igSetScrollHereX(float center_x_ratio);
func SetScrollHereX(center_x_ratio float32) {
	giLib.Call(_func_igSetScrollHereX_, []interface{}{&center_x_ratio})
}

// CIMGUI_API void igSetScrollHereY(float center_y_ratio);
func SetScrollHereY(center_y_ratio float32) {
	giLib.Call(_func_igSetScrollHereY_, []interface{}{&center_y_ratio})
}

// CIMGUI_API void igSetScrollFromPosX_Float(float local_x,float center_x_ratio);
func SetScrollFromPosX(local_x, center_x_ratio float32) {
	giLib.Call(_func_igSetScrollFromPosX_Float_, []interface{}{&local_x, &center_x_ratio})
}

// CIMGUI_API void igSetScrollFromPosY_Float(float local_y,float center_y_ratio);
func SetScrollFromPosY(local_y, center_y_ratio float32) {
	giLib.Call(_func_igSetScrollFromPosY_Float_, []interface{}{&local_y, &center_y_ratio})
}

// CIMGUI_API void igPushFont(ImFont* font);
func PushFont(font *ImFont) {
	giLib.Call(_func_igPushFont_, []interface{}{&font})
}

// CIMGUI_API void igPopFont(void);
func PopFont() {
	giLib.Call(_func_igPopFont_, nil)
}

// CIMGUI_API void igPushStyleColor_U32(ImGuiCol idx,ImU32 col);
func PushStyleColorU32(idx ImGuiCol, col ImU32) {
	giLib.Call(_func_igPushStyleColor_U32_, []interface{}{&idx, &col})
}

// CIMGUI_API void igPushStyleColor_Vec4(ImGuiCol idx,const ImVec4 col);
func PushStyleColorVec4(idx ImGuiCol, col ImVec4) {
	giLib.Call(_func_igPushStyleColor_Vec4_, []interface{}{&idx, &col})
}

// CIMGUI_API void igPushStyleColor_Vec4(ImGuiCol idx,const ImVec4 col);
func PushStyleColor(idx ImGuiCol, r, g, b, a float32) {
	col := NewImVec4Float(r, g, b, a)
	defer col.Destroy()
	giLib.Call(_func_igPushStyleColor_Vec4_, []interface{}{&idx, col})
}

// CIMGUI_API void igPopStyleColor(int count);
func PopStyleColor(count int32) {
	giLib.Call(_func_igPopStyleColor_, []interface{}{&count})
}

// CIMGUI_API void igPushStyleVar_Float(ImGuiStyleVar idx,float val);
func PushStyleVar(idx ImGuiStyleVar, val float32) {
	giLib.Call(_func_igPushStyleVar_Float_, []interface{}{&idx, &val})
}

// CIMGUI_API void igPushStyleVar_Vec2(ImGuiStyleVar idx,const ImVec2 val);
func PushStyleVarVec2(idx ImGuiStyleVar, val ImVec2) {
	giLib.Call(_func_igPushStyleVar_Vec2_, []interface{}{&idx, &val})
}

// CIMGUI_API void igPopStyleVar(int count);
func PopStyleVar(count int32) {
	giLib.Call(_func_igPopStyleVar_, []interface{}{&count})
}

// CIMGUI_API void igPushTabStop(bool tab_stop);
func PushTabStop(tab_stop bool) {
	t := c.CBool(tab_stop)
	giLib.Call(_func_igPushTabStop_, []interface{}{&t})
}

// CIMGUI_API void igPopTabStop(void);
func PopTabStop() { giLib.Call(_func_igPopTabStop_, nil) }

// CIMGUI_API void igPushButtonRepeat(bool repeat);
func PushButtonRepeat(repeat bool) {
	r := c.CBool(repeat)
	giLib.Call(_func_igPushButtonRepeat_, []interface{}{&r})
}

// CIMGUI_API void igPopButtonRepeat(void);
func PopButtonRepeat() { giLib.Call(_func_igPopButtonRepeat_, nil) }

// CIMGUI_API void igPushItemWidth(float item_width);
func PushItemWidth(item_width float32) {
	giLib.Call(_func_igPushItemWidth_, []interface{}{&item_width})
}

// CIMGUI_API void igPopItemWidth(void);
func PopItemWidth() { giLib.Call(_func_igPopItemWidth_, nil) }

// CIMGUI_API void igSetNextItemWidth(float item_width);
func SetNextItemWidth(item_width float32) {
	giLib.Call(_func_igSetNextItemWidth_, []interface{}{&item_width})
}

// CIMGUI_API float igCalcItemWidth(void);
func CalcItemWidth() float32 {
	return giLib.Call(_func_igCalcItemWidth_, nil).F32Free()
}

// CIMGUI_API void igPushTextWrapPos(float wrap_local_pos_x);
func PushTextWrapPos(wrap_local_pos_x float32) {
	giLib.Call(_func_igPushTextWrapPos_, []interface{}{&wrap_local_pos_x})
}

// CIMGUI_API void igPopTextWrapPos(void);
func PopTextWrapPos() { giLib.Call(_func_igPopTextWrapPos_, nil) }

// CIMGUI_API ImFont* igGetFont(void);
func GetFont() *ImFont {
	p := giLib.Call(_func_igGetFont_, nil)
	return (*ImFont)(p.PtrFree())
}

// CIMGUI_API float igGetFontSize(void);
func GetFontSize() float32 {
	return giLib.Call(_func_GetFontSize_, nil).F32Free()
}

// CIMGUI_API void igGetFontTexUvWhitePixel(ImVec2* pOut);
func GetFontTexUvWhitePixel() *ImVec2 {
	p := NewImVec2()
	giLib.Call(_func_igGetFontTexUvWhitePixel_, []interface{}{&p})
	return p
}

// CIMGUI_API ImU32 igGetColorU32_Col(ImGuiCol idx,float alpha_mul);
func GetColorU32Col(idx ImGuiCol, alpha_mul float32) uint32 {
	return giLib.Call(_func_igGetColorU32_Col_, []interface{}{&idx, &alpha_mul}).U32Free()
}

// CIMGUI_API ImU32 igGetColorU32_Vec4(const ImVec4 col);
func GetColorU32Vec4(col ImVec4) uint32 {
	return giLib.Call(_func_igGetColorU32_Vec4_, []interface{}{&col}).U32Free()
}

// CIMGUI_API ImU32 igGetColorU32_U32(ImU32 col);
func GetColorU32(col uint32) uint32 {
	return giLib.Call(_func_igGetColorU32_U32_, []interface{}{&col}).U32Free()
}

// CIMGUI_API const ImVec4* igGetStyleColorVec4(ImGuiCol idx);
func GetStyleColorVec4(idx ImGuiCol) *ImVec4 {
	p := giLib.Call(_func_igGetStyleColorVec4_, []interface{}{&idx})
	return (*ImVec4)(p.PtrFree())
}

// CIMGUI_API void igSeparator(void);
func Separator() { giLib.Call(_func_igSeparator_, nil) }

// CIMGUI_API void igSameLine(float offset_from_start_x,float spacing);
func SameLine(offset_from_start_x, spacing float32) {
	giLib.Call(_func_igSameLine_, []interface{}{&offset_from_start_x, &spacing})
}

// CIMGUI_API void igNewLine(void);
func NewLine() { giLib.Call(_func_igNewLine_, nil) }

// CIMGUI_API void igSpacing(void);
func Spacing() { giLib.Call(_func_igSpacing_, nil) }

// CIMGUI_API void igDummy(const ImVec2 size);
func DummyVec2(size ImVec2) {
	giLib.Call(_func_igDummy_, []interface{}{&size})
}

// CIMGUI_API void igDummy(const ImVec2 size);
func Dummy(w, h float32) {
	size := NewImVec2Float(w, h)
	defer size.Destroy()
	giLib.Call(_func_igDummy_, []interface{}{size})
}

// CIMGUI_API void igIndent(float indent_w);
func Indent(indent_w float32) {
	giLib.Call(_func_igIndent_, []interface{}{&indent_w})
}

// CIMGUI_API void igUnindent(float indent_w);
func Unindent(indent_w float32) {
	giLib.Call(_func_igUnindent_, []interface{}{&indent_w})
}

// CIMGUI_API void igBeginGroup(void);
func BeginGroup() { giLib.Call(_func_igBeginGroup_, nil) }

// CIMGUI_API void igEndGroup(void);
func EndGroup() { giLib.Call(_func_igEndGroup_, nil) }

// CIMGUI_API void igGetCursorPos(ImVec2* pOut);
func GetCursorPosVec2() *ImVec2 {
	pos := NewImVec2()
	giLib.Call(_func_igGetCursorPos_, []interface{}{&pos})
	return pos
}

// CIMGUI_API void igGetCursorPos(ImVec2* pOut);
func GetCursorPos() (x, y float32) {
	pos := NewImVec2()
	giLib.Call(_func_igGetCursorPos_, []interface{}{&pos})
	defer pos.Destroy()
	return pos.X, pos.Y
}

// CIMGUI_API float igGetCursorPosX(void);
func GetCursorPosX() float32 {
	return giLib.Call(_func_igGetCursorPosX_, nil).F32Free()
}

// CIMGUI_API float igGetCursorPosY(void);
func GetCursorPosY() float32 {
	return giLib.Call(_func_igGetCursorPosY_, nil).F32Free()
}

// CIMGUI_API void igSetCursorPos(const ImVec2 local_pos);
func SetCursorPosVec2(local_pos ImVec2) {
	giLib.Call(_func_igSetCursorPos_, []interface{}{&local_pos})
}

// CIMGUI_API void igSetCursorPos(const ImVec2 local_pos);
func SetCursorPos(x, y float32) {
	local_pos := NewImVec2Float(x, y)
	defer local_pos.Destroy()
	giLib.Call(_func_igSetCursorPos_, []interface{}{local_pos})
}

// CIMGUI_API void igSetCursorPosX(float local_x);
func SetCursorPosX(local_x float32) {
	giLib.Call(_func_igSetCursorPosX_, []interface{}{&local_x})
}

// CIMGUI_API void igSetCursorPosY(float local_y);
func SetCursorPosY(local_y float32) {
	giLib.Call(_func_igSetCursorPosY_, []interface{}{&local_y})
}

// CIMGUI_API void igGetCursorStartPos(ImVec2* pOut);
func GetCursorStartPosVec2() *ImVec2 {
	p := NewImVec2()
	giLib.Call(_func_igGetCursorStartPos_, []interface{}{&p})
	return p
}

// CIMGUI_API void igGetCursorStart(ImVec2* pOut);
func GetCursorStartPos() (float32, float32) {
	p := NewImVec2()
	defer p.Destroy()
	giLib.Call(_func_igGetCursorStartPos_, []interface{}{&p})
	return p.X, p.Y
}

// CIMGUI_API void igGetCursorScreenPos(ImVec2* pOut);
func GetCursorScreenPosVec2() *ImVec2 {
	p := NewImVec2()
	giLib.Call(_func_igGetCursorScreenPos_, []interface{}{&p})
	return p
}

// CIMGUI_API void igGetCursorScreenPos(ImVec2* pOut);
func GetCursorScreenPos() (float32, float32) {
	p := NewImVec2()
	defer p.Destroy()
	giLib.Call(_func_igGetCursorScreenPos_, []interface{}{&p})
	return p.X, p.Y
}

// CIMGUI_API void igSetCursorScreenPos(const ImVec2 pos);
func SetCursorScreenPosVec2(pos ImVec2) {
	giLib.Call(_func_igSetCursorScreenPos_, []interface{}{&pos})
}

// CIMGUI_API void igSetCursorScreenPos(const ImVec2 pos);
func SetCursorScreenPos(x, y float32) {
	pos := NewImVec2Float(x, y)
	defer pos.Destroy()
	giLib.Call(_func_igSetCursorScreenPos_, []interface{}{pos})
}

// CIMGUI_API void igAlignTextToFramePadding(void);
func AlignTextToFramePadding() {
	giLib.Call(_func_igAlignTextToFramePadding_, nil)
}

// CIMGUI_API float igGetTextLineHeight(void);
func GetTextLineHeight() float32 {
	return giLib.Call(_func_igGetTextLineHeight_, nil).F32Free()
}

// CIMGUI_API float igGetTextLineHeightWithSpacing(void);
func GetTextLineHeightWithSpacing() float32 {
	return giLib.Call(_func_igGetTextLineHeightWithSpacing_, nil).F32Free()
}

// CIMGUI_API float igGetFrameHeight(void);
func GetFrameHeight() float32 {
	return giLib.Call(_func_igGetFrameHeight_, nil).F32Free()
}

// CIMGUI_API float igGetFrameHeightWithSpacing(void);
func GetFrameHeightWithSpacing() float32 {
	return giLib.Call(_func_igGetFrameHeightWithSpacing_, nil).F32Free()
}

// CIMGUI_API void igPushID_Str(const char* str_id);
func PushIDStr(str_id string) {
	id := c.CStr(str_id)
	defer usf.Free(id)
	giLib.Call(_func_igPushID_Str_, []interface{}{&id})
}

// CIMGUI_API void igPushID_StrStr(const char* str_id_begin,const char* str_id_end);
func PushIDStrStr(str_id_begin, str_id_end string) {
	begin, end := c.CStr(str_id_begin), c.CStr(str_id_end)
	defer usf.Free(begin)
	defer usf.Free(end)

	giLib.Call(_func_igPushID_StrStr_, []interface{}{&begin, &end})
}

// CIMGUI_API void igPushID_Ptr(const void* ptr_id);
func PushIDPtr(ptr_id unsafe.Pointer) {
	giLib.Call(_func_igPushID_Ptr_, []interface{}{&ptr_id})
}

// CIMGUI_API void igPushID_Int(int int_id);
func PushIDInt(int_id int32) {
	giLib.Call(_func_igPushID_Int_, []interface{}{&int_id})
}

// CIMGUI_API void igPopID(void);
func PopID() { giLib.Call(_func_igPopID_, nil) }

// CIMGUI_API ImGuiID igGetID_Str(const char* str_id);
func GetIDtr(str_id string) ImGuiID {
	id := c.CStr(str_id)
	defer usf.Free(id)

	p := giLib.Call(_func_igGetID_Str_, []interface{}{&id})
	return ImGuiID(p.U32Free())
}

// CIMGUI_API ImGuiID igGetID_StrStr(const char* str_id_begin,const char* str_id_end);
func GetIDStrStr(str_id_begin, str_id_end string) ImGuiID {
	begin, end := c.CStr(str_id_begin), c.CStr(str_id_end)
	defer usf.Free(begin)
	defer usf.Free(end)

	p := giLib.Call(_func_igGetID_StrStr_, []interface{}{&begin, &end})
	return ImGuiID(p.U32Free())
}

// CIMGUI_API ImGuiID igGetID_Ptr(const void* ptr_id);
func GetID_Ptr(ptr_id unsafe.Pointer) ImGuiID {
	p := giLib.Call(_func_igGetID_Ptr_, []interface{}{&ptr_id})
	return ImGuiID(p.U32Free())
}

// CIMGUI_API void igTextUnformatted(const char* text,const char* text_end);
func TextUnformatted(text, text_end string) {
	txt, txt_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(txt)
	defer usf.Free(txt_end)

	giLib.Call(_func_igTextUnformatted_, []interface{}{&txt, &txt_end})
}

// CIMGUI_API void igText(const char* fmt,...);
func Text(text string) {
	txt := c.CStr(text)
	defer usf.Free(txt)
	giLib.Call(_func_igText_, []interface{}{&txt})
}

// CIMGUI_API void igTextV(const char* fmt,va_list args);

// CIMGUI_API void igTextColored(const ImVec4 col,const char* fmt,...);
func TextColoredVec4(col ImVec4, text string) {
	txt := c.CStr(text)
	defer usf.Free(txt)
	giLib.Call(_func_igTextColored_, []interface{}{&col, &txt})
}

// CIMGUI_API void igTextColored(const ImVec4 col,const char* fmt,...);
func TextColored(r, g, b, a float32, text string) {
	txt, col := c.CStr(text), NewImVec4Float(r, g, b, a)
	defer usf.Free(txt)
	defer col.Destroy()
	giLib.Call(_func_igTextColored_, []interface{}{col, &txt})
}

// CIMGUI_API void igTextColoredV(const ImVec4 col,const char* fmt,va_list args);

// CIMGUI_API void igTextDisabled(const char* fmt,...);
func TextDisabled(text string) {
	txt := c.CStr(text)
	defer usf.Free(txt)
	giLib.Call(_func_igTextDisabled_, []interface{}{&txt})
}

// CIMGUI_API void igTextDisabledV(const char* fmt,va_list args);

// CIMGUI_API void igTextWrapped(const char* fmt,...);
func TextWrapped(text string) {
	txt := c.CStr(text)
	defer usf.Free(txt)
	giLib.Call(_func_igTextWrapped_, []interface{}{&txt})
}

// CIMGUI_API void igTextWrappedV(const char* fmt,va_list args);

// CIMGUI_API void igLabelText(const char* label,const char* fmt,...);
func LabelText(label, text string) {
	lb, txt := c.CStr(label), c.CStr(text)
	defer usf.Free(txt)
	defer usf.Free(lb)

	giLib.Call(_func_igLabelText_, []interface{}{&lb, &txt})
}

// CIMGUI_API void igLabelTextV(const char* label,const char* fmt,va_list args);

// CIMGUI_API void igBulletText(const char* fmt,...);
func BulletText(text string) {
	txt := c.CStr(text)
	defer usf.Free(txt)
	giLib.Call(_func_igBulletText_, []interface{}{&txt})
}

// CIMGUI_API void igBulletTextV(const char* fmt,va_list args);

// CIMGUI_API void igSeparatorText(const char* label);
func SeparatorText(text string) {
	txt := c.CStr(text)
	defer usf.Free(txt)
	giLib.Call(_func_igSeparatorText_, []interface{}{&txt})
}

// CIMGUI_API bool igButton(const char* label,const ImVec2 size);
func ButtonVec2(label string, size ImVec2) (check bool) {
	lb := c.CStr(label)
	defer usf.Free(lb)
	return giLib.Call(_func_igButton_, []interface{}{&lb, &size}).BoolFree()
}

// CIMGUI_API bool igButton(const char* label,const ImVec2 size);
func Button(label string, w, h float32) (check bool) {
	lb, size := c.CStr(label), NewImVec2Float(w, h)
	defer usf.Free(lb)
	defer size.Destroy()
	return giLib.Call(_func_igButton_, []interface{}{&lb, size}).BoolFree()
}

// CIMGUI_API bool igSmallButton(const char* label);
func SmallButton(label string) (click bool) {
	lb := c.CStr(label)
	defer usf.Free(lb)
	return giLib.Call(_func_igSmallButton_, []interface{}{&lb}).BoolFree()
}

// CIMGUI_API bool igInvisibleButton(const char* str_id,const ImVec2 size,ImGuiButtonFlags flags);
func InvisibleButtonVec2(str_id string, size ImVec2, flag ImGuiButtonFlags) (click bool) {
	id := c.CStr(str_id)
	defer usf.Free(id)

	return giLib.Call(_func_igInvisibleButton_, []interface{}{&id, &size, &flag}).BoolFree()
}

// CIMGUI_API bool igInvisibleButton(const char* str_id,const ImVec2 size,ImGuiButtonFlags flags);
func InvisibleButton(str_id string, w, h float32, flag ImGuiButtonFlags) (click bool) {
	id, size := c.CStr(str_id), NewImVec2Float(w, h)
	defer usf.Free(id)
	defer size.Destroy()

	return giLib.Call(_func_igInvisibleButton_, []interface{}{&id, size, &flag}).BoolFree()
}

// CIMGUI_API bool igArrowButton(const char* str_id,ImGuiDir dir);
func ArrowButton(str_id string, dir ImGuiDir) (click bool) {
	id := c.CStr(str_id)
	defer usf.Free(id)

	return giLib.Call(_func_igArrowButton_, []interface{}{&id, &dir}).BoolFree()
}

// CIMGUI_API bool igCheckbox(const char* label,bool* v);
func Checkbox(label string, checked *bool) (check bool) {
	lb, v := c.CStr(label), new(int32)
	defer usf.Free(lb)
	*v = c.CBool(*checked)

	p := giLib.Call(_func_igCheckbox_, []interface{}{&lb, &v})
	(*checked) = c.GoBool(*v)
	return p.BoolFree()
}

// CIMGUI_API bool igCheckboxFlags_IntPtr(const char* label,int* flags,int flags_value);
func CheckboxFlagsIntPtr(label string, flags *int32, mask int32) (check bool) {
	lb := c.CStr(label)
	defer usf.Free(lb)

	return giLib.Call(_func_igCheckboxFlags_IntPtr_, []interface{}{&lb, &flags, &mask}).BoolFree()
}

// CIMGUI_API bool igCheckboxFlags_UintPtr(const char* label,unsigned int* flags,unsigned int flags_value);
func CheckboxFlagsUintPtr(label string, flags *uint32, mask uint32) (check bool) {
	lb := c.CStr(label)
	defer usf.Free(lb)

	return giLib.Call(_func_igCheckboxFlags_UintPtr_, []interface{}{&lb, &flags, &mask}).BoolFree()
}

// CIMGUI_API bool igRadioButton_Bool(const char* label,bool active);
func RadioButton(label string, active bool) (check bool) {
	lb, act := c.CStr(label), c.CBool(active)
	defer usf.Free(lb)

	return giLib.Call(_func_igRadioButton_Bool_, []interface{}{&lb, &act}).BoolFree()
}

// CIMGUI_API bool igRadioButton_IntPtr(const char* label,int* v,int v_button);
func RadioButtonIntPtr(label string, flags *int32, mask int32) (check bool) {
	lb := c.CStr(label)
	defer usf.Free(lb)

	return giLib.Call(_func_igRadioButton_IntPtr_, []interface{}{&lb, &flags, &mask}).BoolFree()
}

// CIMGUI_API void igProgressBar(float fraction,const ImVec2 size_arg,const char* overlay);
func ProgressBarVec2(fraction float32, size ImVec2, overlay string) {
	o := c.CStr(overlay)
	defer usf.Free(o)
	giLib.Call(_func_igProgressBar_, []interface{}{&fraction, &size, &o})
}

// CIMGUI_API void igProgressBar(float fraction,const ImVec2 size_arg,const char* overlay);
func ProgressBar(fraction float32, w, h float32, overlay string) {
	o, size := c.CStr(overlay), NewImVec2Float(w, h)
	defer usf.Free(o)
	defer size.Destroy()
	giLib.Call(_func_igProgressBar_, []interface{}{&fraction, &size, &o})
}

// CIMGUI_API void igBullet(void);
func Bullet() { giLib.Call(_func_igBullet_, nil) }

// CIMGUI_API void igImage(ImTextureID user_texture_id,const ImVec2 size,
// const ImVec2 uv0,const ImVec2 uv1,const ImVec4 tint_col,const ImVec4 border_col);
func Image(user_texture_id ImTextureID, size, uv0, uv1 ImVec2, tint_col, border_col ImVec4) {
	giLib.Call(_func_igImage_, []interface{}{&user_texture_id, &size, &uv0, &uv1, &tint_col, &border_col})

}

// CIMGUI_API bool igImageButton(const char* str_id,ImTextureID user_texture_id,const ImVec2 size,const ImVec2 uv0,const ImVec2 uv1,const ImVec4 bg_col,const ImVec4 tint_col);
func ImageButton(label string, user_texture_id ImTextureID, size, uv0, uv1 ImVec2, tint_col, border_col ImVec4) (click bool) {
	l := c.CStr(label)
	defer usf.Free(l)

	return giLib.Call(_func_igImageButton_, []interface{}{&l, &user_texture_id, &size, &uv0, &uv1, &tint_col, &border_col}).BoolFree()

}

// CIMGUI_API bool igBeginCombo(const char* label,const char* preview_value,ImGuiComboFlags flags);
func BeginCombo(label, preview string, flag ImGuiComboFlags) (click bool) {
	l, p := c.CStr(label), c.CStr(preview)
	defer usf.Free(l)
	defer usf.Free(p)

	return giLib.Call(_func_igBeginCombo_, []interface{}{&l, &p, &flag}).BoolFree()
}

// CIMGUI_API void igEndCombo(void);
func EndCombo() { giLib.Call(_func_igEndCombo_, nil) }

// CIMGUI_API bool igCombo_Str_arr(const char* label,int* current_item,const char* const items[],int items_count,int popup_max_height_in_items);
func ComboStrArr(label string, currItem *int32, items []string, popup_max_height_in_items int32) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)

	_itemsLen := uint64(uint64(len(items)))
	_items := usf.MallocN(_itemsLen, 8)
	_items_count := int32(_itemsLen)
	defer usf.Free(_items)
	for i := uint64(0); i < _itemsLen; i++ {
		_item := c.CStr(items[i])
		defer usf.Free(_item)
		usf.PushAt(_items, i, _item)
	}

	return giLib.Call(_func_igCombo_Str_arr_, []interface{}{&_label, &currItem, &_items, &_items_count, &popup_max_height_in_items}).BoolFree()

}

// CIMGUI_API bool igCombo_Str(const char* label,int* current_item,const char* items_separated_by_zeros,int popup_max_height_in_items);
func ComboStr(label string, currItem *int32, items_separated_by_zeros string, popup_max_height_in_items int32) (click bool) {
	_label, _items_separated_by_zeros := c.CStr(label), c.CStr(items_separated_by_zeros)
	defer usf.Free(_label)
	defer usf.Free(_items_separated_by_zeros)

	return giLib.Call(_func_igCombo_Str_, []interface{}{&_label, &currItem, &_items_separated_by_zeros, &popup_max_height_in_items}).BoolFree()

}

// CIMGUI_API bool igCombo_FnBoolPtr(const char* label,int* current_item,bool(*items_getter)(void* data,int idx,const char** out_text),void* data,int items_count,int popup_max_height_in_items);

// CIMGUI_API bool igDragFloat(const char* label,float* v,float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func DragFloat(label string, v *float32, v_speed, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igDragFloat_, []interface{}{&_label, &v, &v_speed, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragFloat2(const char* label,float v[2],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func DragFloat2(label string, v []float32, v_speed, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igDragFloat2_, []interface{}{&_label, &p, &v_speed, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragFloat3(const char* label,float v[3],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func DragFloat3(label string, v []float32, v_speed, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igDragFloat3_, []interface{}{&_label, &p, &v_speed, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragFloat4(const char* label,float v[4],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func DragFloat4(label string, v []float32, v_speed, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igDragFloat4_, []interface{}{&_label, &p, &v_speed, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragFloatRange2(const char* label,float* v_current_min,float* v_current_max,float v_speed,float v_min,float v_max,const char* format,const char* format_max,ImGuiSliderFlags flags);
func DragFloatRange2(label string, currMin, currMax *float32, v_speed, v_min, v_max float32, format, format_max string, flag ImGuiSliderFlags) bool {
	_label, _fmt, _fmtMax := c.CStr(label), c.CStr(format), c.CStr(format_max)
	defer usf.Free(_label)
	defer usf.Free(_fmt)
	defer usf.Free(_fmtMax)

	return giLib.Call(_func_igDragFloatRange2_, []interface{}{&_label, &currMin, &currMax, &v_speed, &v_min, &v_max, &_fmt, &_fmtMax, &flag}).BoolFree()

}

// CIMGUI_API bool igDragInt(const char* label,int* v,float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func DragInt(label string, v *int32, v_speed float32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igDragInt_, []interface{}{&_label, &v, &v_speed, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragInt2(const char* label,int v[2],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func DragInt2(label string, v []int32, v_speed float32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igDragInt2_, []interface{}{&_label, &p, &v_speed, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragInt3(const char* label,int v[3],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func DragInt3(label string, v []int32, v_speed float32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igDragInt3_, []interface{}{&_label, &p, &v_speed, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragInt4(const char* label,int v[4],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func DragInt4(label string, v []int32, v_speed float32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igDragInt4_, []interface{}{&_label, &p, &v_speed, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragIntRange2(const char* label,int* v_current_min,int* v_current_max,float v_speed,int v_min,int v_max,const char* format,const char* format_max,ImGuiSliderFlags flags);
func DragIntRange2(label string, currMin, currMax *int32, v_speed float32, v_min, v_max int, format, format_max string, flag ImGuiSliderFlags) bool {
	_label, _fmt, _fmtMax := c.CStr(label), c.CStr(format), c.CStr(format_max)
	defer usf.Free(_label)
	defer usf.Free(_fmt)
	defer usf.Free(_fmtMax)

	return giLib.Call(_func_igDragIntRange2_, []interface{}{&_label, &currMin, &currMax, &v_speed, &v_min, &v_max, &_fmt, &_fmtMax, &flag}).BoolFree()

}

// CIMGUI_API bool igDragScalar(const char* label,ImGuiDataType data_type,void* p_data,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags);
func DragScalar(label string, dataType ImGuiDataType, dataAddr interface{},
	v_speed float32, dataMinAddr, dataMaxAddr interface{}, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p_data := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataAddr), 2)))[1]
	p_min := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMinAddr), 2)))[1]
	p_max := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMaxAddr), 2)))[1]
	return giLib.Call(_func_igDragScalar_, []interface{}{&_label, &dataType, &p_data, &v_speed, &p_min, &p_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igDragScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags);
func DragScalarN(label string, dataType ImGuiDataType, arryFirstAddr interface{}, components int32,
	v_speed float32, dataMinAddr, dataMaxAddr interface{}, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p_data := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&arryFirstAddr), 2)))[1]
	p_min := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMinAddr), 2)))[1]
	p_max := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMaxAddr), 2)))[1]
	return giLib.Call(_func_igDragScalarN_, []interface{}{&_label, &dataType, &p_data, &components, &v_speed, &p_min, &p_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderFloat(const char* label,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func SliderFloat(label string, v *float32, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igSliderFloat_, []interface{}{&_label, &v, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderFloat2(const char* label,float v[2],float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func SliderFloat2(label string, v []float32, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igSliderFloat2_, []interface{}{&_label, &p, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderFloat3(const char* label,float v[3],float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func SliderFloat3(label string, v []float32, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igSliderFloat3_, []interface{}{&_label, &p, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderFloat4(const char* label,float v[4],float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func SliderFloat4(label string, v []float32, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igSliderFloat4_, []interface{}{&_label, &p, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderAngle(const char* label,float* v_rad,float v_degrees_min,float v_degrees_max,const char* format,ImGuiSliderFlags flags);
func SliderAngle(label string, v *float32, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igSliderAngle_, []interface{}{&_label, &v, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderInt(const char* label,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func SliderInt(label string, v *int32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igSliderInt_, []interface{}{&_label, &v, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderInt2(const char* label,int v[2],int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func SliderInt2(label string, v []int32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igSliderInt2_, []interface{}{&_label, &p, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderInt3(const char* label,int v[3],int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func SliderInt3(label string, v []int32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igSliderInt3_, []interface{}{&_label, &p, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderInt4(const char* label,int v[4],int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func SliderInt4(label string, v []int32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p := &v[0]
	return giLib.Call(_func_igSliderInt4_, []interface{}{&_label, &p, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags);
func SliderScalar(label string, dataType ImGuiDataType, dataAddr interface{},
	dataMinAddr, dataMaxAddr interface{}, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)
	p_data := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataAddr), 2)))[1]
	p_min := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMinAddr), 2)))[1]
	p_max := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMaxAddr), 2)))[1]
	return giLib.Call(_func_igSliderScalar_, []interface{}{&_label, &dataType, &p_data, &p_min, &p_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igSliderScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags);
func SliderScalarN(label string, dataType ImGuiDataType, arryFirstAddr interface{}, components int32,
	dataMinAddr, dataMaxAddr interface{}, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p_data := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&arryFirstAddr), 2)))[1]
	p_min := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMinAddr), 2)))[1]
	p_max := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMaxAddr), 2)))[1]
	return giLib.Call(_func_igSliderScalarN_, []interface{}{&_label, &dataType, &p_data, &components, &p_min, &p_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igVSliderFloat(const char* label,const ImVec2 size,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func VSliderFloatVec2(label string, size ImVec2, v *float32, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igVSliderFloat_, []interface{}{&_label, &size, &v, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igVSliderFloat(const char* label,const ImVec2 size,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags);
func VSliderFloat(label string, w, h float32, val *float32, v_min, v_max float32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt, size := c.CStr(label), c.CStr(format), NewImVec2Float(w, h)
	defer usf.Free(_label)
	defer usf.Free(_fmt)
	defer size.Destroy()

	return giLib.Call(_func_igVSliderFloat_, []interface{}{&_label, size, &val, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igVSliderInt(const char* label,const ImVec2 size,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func VSliderIntVec2(label string, size ImVec2, v *int32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igVSliderInt_, []interface{}{&_label, &size, &v, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igVSliderInt(const char* label,const ImVec2 size,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags);
func VSliderInt(label string, w, h float32, v *int32, v_min, v_max int32, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt, size := c.CStr(label), c.CStr(format), NewImVec2Float(w, h)
	defer usf.Free(_label)
	defer usf.Free(_fmt)
	defer size.Destroy()

	return giLib.Call(_func_igVSliderInt_, []interface{}{&_label, size, &v, &v_min, &v_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igVSliderScalar(const char* label,const ImVec2 size,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags);
func VSliderScalarVec2(label string, size ImVec2, dataType ImGuiDataType, dataAddr interface{},
	dataMinAddr, dataMaxAddr interface{}, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)
	p_data := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataAddr), 2)))[1]
	p_min := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMinAddr), 2)))[1]
	p_max := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMaxAddr), 2)))[1]
	return giLib.Call(_func_igVSliderScalar_, []interface{}{&_label, &size, &dataType, &p_data, &p_min, &p_max, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igVSliderScalar(const char* label,const ImVec2 size,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags);
func VSliderScalar(label string, w, h float32, dataType ImGuiDataType, dataAddr interface{},
	dataMinAddr, dataMaxAddr interface{}, format string, flag ImGuiSliderFlags) bool {
	_label, _fmt, size := c.CStr(label), c.CStr(format), NewImVec2Float(w, h)
	defer usf.Free(_label)
	defer usf.Free(_fmt)
	defer size.Destroy()
	p_data := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataAddr), 2)))[1]
	p_min := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMinAddr), 2)))[1]
	p_max := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataMaxAddr), 2)))[1]
	return giLib.Call(_func_igVSliderScalar_, []interface{}{&_label, size, &dataType, &p_data, &p_min, &p_max, &_fmt, &flag}).BoolFree()

}

func NewInputTextCallback(fn func(data *ImGuiInputTextCallbackData) int32) *c.Callback {
	cb := c.NewCallback(c.AbiDefault, c.I32, []c.Type{c.Pointer})
	cb.CallbackFunc = fn
	return cb
}

// CIMGUI_API bool igInputText(const char* label,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data);
//
// typedef int (*ImGuiInputTextCallback)(ImGuiInputTextCallbackData* data);
//
// callback : func NewInputTextCallback(fn func(data *ImGuiInputTextCallbackData) int32)
//
// IMGUI_API bool InputText(const char* label, char* buf, size_t buf_size, ImGuiInputTextFlags flags = 0, ImGuiInputTextCallback callback = NULL, void* user_data = NULL);
func InputText(label string, val *string, valMaxLen uint64, flag ImGuiInputTextFlags, callback *c.Callback) (act bool, currVal string) {
	_label, src := c.CStr(label), c.CStr(*val)
	defer usf.Free(_label)
	defer usf.Free(src)

	_buf := usf.MallocN(valMaxLen, 1)
	usf.Memset(_buf, 0, valMaxLen)
	usf.Memcpy(_buf, src, uint64(len(*val)))
	defer usf.Free(_buf)

	cb := unsafe.Pointer(nil)
	if callback != nil {
		if callback.CallbackCvt == nil {
			callback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				data := (*ImGuiInputTextCallbackData)(args[0].Ptr())
				fn := cb.CallbackFunc.(func(*ImGuiInputTextCallbackData) int32)
				if fn != nil {
					i := fn(data)
					ret.SetI32(i)
				}
			}
		}
		cb = callback.Cfunc
	}

	act = giLib.Call(_func_igInputText_, []interface{}{&_label, &_buf, &valMaxLen, &flag, &cb, _ptr_zero}).BoolFree()

	currVal = c.GoStr(_buf)
	*val = currVal
	return
}

// CIMGUI_API bool igInputText(const char* label,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data);
// IMGUI_API bool InputText(const char* label, char* buf, size_t buf_size, ImGuiInputTextFlags flags = 0, ImGuiInputTextCallback callback = NULL, void* user_data = NULL);
func InputTextDefault(label string, val string, valMaxLen uint64) (act bool, currVal string) {
	_label, src := c.CStr(label), c.CStr(val)
	defer usf.Free(_label)
	defer usf.Free(src)

	_buf := usf.MallocN(valMaxLen, 1)
	usf.Memset(_buf, 0, valMaxLen)
	usf.Memcpy(_buf, src, uint64(len(val)))
	defer usf.Free(_buf)

	act = giLib.Call(_func_igInputText_, []interface{}{&_label, &_buf, &valMaxLen, _u32_zero, _ptr_zero, _ptr_zero}).BoolFree()

	currVal = c.GoStr(_buf)
	return
}

// CIMGUI_API bool igInputTextMultiline(const char* label,char* buf,size_t buf_size,const ImVec2 size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data);
func InputTextMultilineVec2(label string, val string, valMaxLen uint64, size ImVec2, flag ImGuiInputTextFlags, callback *c.Callback) (act bool, currVal string) {
	_label, src := c.CStr(label), c.CStr(val)
	defer usf.Free(_label)
	defer usf.Free(src)

	_buf := usf.MallocN(valMaxLen, 1)
	usf.Memset(_buf, 0, valMaxLen)
	usf.Memcpy(_buf, src, uint64(len(val)))
	defer usf.Free(_buf)

	cb, data := unsafe.Pointer(nil), usf.MallocN(1, 8)
	if callback != nil {
		if callback.CallbackCvt == nil {
			callback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				data := (*ImGuiInputTextCallbackData)(args[0].Ptr())
				fn := cb.CallbackFunc.(func(data *ImGuiInputTextCallbackData) int32)
				if fn != nil {
					i := fn(data)
					ret.SetI32(i)
				}
			}
		}
		cb = callback.Cfunc
	}
	usf.Memset(data, 0, 8)
	defer usf.Free(data)
	act = giLib.Call(_func_igInputTextMultiline_, []interface{}{&_label, &_buf, &valMaxLen, &size, &flag, &cb, &data}).BoolFree()

	currVal = c.GoStr(_buf)
	return
}

// CIMGUI_API bool igInputTextMultiline(const char* label,char* buf,size_t buf_size,const ImVec2 size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data);
//
// IMGUI_API bool InputTextMultiline(const char* label, char* buf, size_t buf_size, const ImVec2& size = ImVec2(0, 0), ImGuiInputTextFlags flags = 0, ImGuiInputTextCallback callback = NULL, void* user_data = NULL);
func InputTextMultiline(label string, val *string, valMaxLen uint64, w, h float32, flag ImGuiInputTextFlags, callback *c.Callback) (act bool) {
	_label, src, size := c.CStr(label), c.CStr(*val), NewImVec2Float(w, h)

	_buf := usf.MallocN(valMaxLen, 1)
	usf.Memset(_buf, 0, valMaxLen)
	usf.Memcpy(_buf, src, uint64(len(*val)))
	usf.Free(src)

	cb := unsafe.Pointer(nil)
	if callback != nil {
		if callback.CallbackCvt == nil {
			callback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				data := (*ImGuiInputTextCallbackData)(args[0].Ptr())
				fn := cb.CallbackFunc.(func(data *ImGuiInputTextCallbackData) int32)
				if fn != nil {
					i := fn(data)
					ret.SetI32(i)
				}
			}
		}
		cb = callback.Cfunc
	}

	act = giLib.Call(_func_igInputTextMultiline_, []interface{}{&_label, &_buf, &valMaxLen, size, &flag, &cb, _ptr_zero}).BoolFree()
	usf.Free(_label)

	(*val) = c.GoStr(_buf)
	usf.Free(_buf)
	return
}

// CIMGUI_API bool igInputTextWithHint(const char* label,const char* hint,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data);
func InputTextWithHint(label string, hint string, val string, valMaxLen uint64, flag ImGuiInputTextFlags, callback *c.Callback) (act bool, currVal string) {
	_label, _hint, src := c.CStr(label), c.CStr(hint), c.CStr(val)

	_buf := usf.MallocN(valMaxLen, 1)
	usf.Memset(_buf, 0, valMaxLen)
	usf.Memcpy(_buf, src, uint64(len(val)))
	usf.Free(src)

	cb := unsafe.Pointer(nil)
	if callback != nil {
		if callback.CallbackCvt != nil {
			callback.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
				data := (*ImGuiInputTextCallbackData)(args[0].Ptr())
				fn := cb.CallbackFunc.(func(data *ImGuiInputTextCallbackData) int32)
				if fn != nil {
					i := fn(data)
					ret.SetI32(i)
				}
			}
		}
		cb = callback.Cfunc
	}

	act = giLib.Call(_func_igInputTextWithHint_, []interface{}{&_label, &_hint, &_buf, &valMaxLen, &flag, &cb, _ptr_zero}).BoolFree()
	usf.Free(_label)
	usf.Free(_hint)

	currVal = c.GoStr(_buf)
	usf.Free(_buf)
	return
}

// CIMGUI_API bool igInputFloat(const char* label,float* v,float step,float step_fast,const char* format,ImGuiInputTextFlags flags);
func InputFloat(label string, v *float32, step, step_fast float32, format string, flag ImGuiInputTextFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igInputFloat_, []interface{}{&_label, &v, &step, &step_fast, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igInputFloat2(const char* label,float v[2],const char* format,ImGuiInputTextFlags flags);
func InputFloat2(label string, v []float32, format string, flag ImGuiInputTextFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	_v := &v[0]
	return giLib.Call(_func_igInputFloat2_, []interface{}{&_label, &_v, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igInputFloat3(const char* label,float v[3],const char* format,ImGuiInputTextFlags flags);
func InputFloat3(label string, v []float32, format string, flag ImGuiInputTextFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	_v := &v[0]
	return giLib.Call(_func_igInputFloat3_, []interface{}{&_label, &_v, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igInputFloat4(const char* label,float v[4],const char* format,ImGuiInputTextFlags flags);
func InputFloat4(label string, v []float32, format string, flag ImGuiInputTextFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	_v := &v[0]
	return giLib.Call(_func_igInputFloat4_, []interface{}{&_label, &_v, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igInputInt(const char* label,int* v,int step,int step_fast,ImGuiInputTextFlags flags);
func InputInt(label string, v *int32, step, step_fast int32, flag ImGuiInputTextFlags) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)

	return giLib.Call(_func_igInputInt_, []interface{}{&_label, &v, &step, &step_fast, &flag}).BoolFree()

}

// CIMGUI_API bool igInputInt2(const char* label,int v[2],ImGuiInputTextFlags flags);
func InputInt2(label string, v []int32, flag ImGuiInputTextFlags) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)
	return giLib.Call(_func_igInputInt2_, []interface{}{&_label, &v, &flag}).BoolFree()
}

// CIMGUI_API bool igInputInt3(const char* label,int v[3],ImGuiInputTextFlags flags);
func InputInt3(label string, v []int32, flag ImGuiInputTextFlags) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)
	return giLib.Call(_func_igInputInt3_, []interface{}{&_label, &v, &flag}).BoolFree()
}

// CIMGUI_API bool igInputInt4(const char* label,int v[4],ImGuiInputTextFlags flags);
func InputInt4(label string, v []int32, flag ImGuiInputTextFlags) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)
	return giLib.Call(_func_igInputInt4_, []interface{}{&_label, &v, &flag}).BoolFree()
}

// CIMGUI_API bool igInputDouble(const char* label,double* v,double step,double step_fast,const char* format,ImGuiInputTextFlags flags);
func InputDouble(label string, v *float64, step, step_fast float64, format string, flag ImGuiInputTextFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	return giLib.Call(_func_igInputDouble_, []interface{}{&_label, &v, &step, &step_fast, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igInputScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags);
func InputScalar(label string, dataType ImGuiDataType, dataAddr interface{},
	stepAddr, stepFastAddr interface{}, format string, flag ImGuiInputTextFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p_data := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataAddr), 2)))[1]
	p_step := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&stepAddr), 2)))[1]
	p_step_fast := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&stepFastAddr), 2)))[1]
	return giLib.Call(_func_igInputScalar_, []interface{}{&_label, &dataType, &p_data, &p_step, &p_step_fast, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igInputScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags);
func InputScalarN(label string, dataType ImGuiDataType, arryFirstAddr interface{}, components int32,
	stepAddr, stepFastAddr interface{}, format string, flag ImGuiInputTextFlags) bool {
	_label, _fmt := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_fmt)

	p_data := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&arryFirstAddr), 2)))[1]
	p_step := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&stepAddr), 2)))[1]
	p_step_fast := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&stepFastAddr), 2)))[1]
	return giLib.Call(_func_igInputScalarN_, []interface{}{&_label, &dataType, &p_data, &components, &p_step, &p_step_fast, &_fmt, &flag}).BoolFree()

}

// CIMGUI_API bool igColorEdit3(const char* label,float col[3],ImGuiColorEditFlags flags);
func ColorEdit3(label string, col []float32, flag ImGuiColorEditFlags) bool {
	_label, _col := c.CStr(label), &col[0]
	defer usf.Free(_label)
	return giLib.Call(_func_igColorEdit3_, []interface{}{&_label, &_col, &flag}).BoolFree()
}

// CIMGUI_API bool igColorEdit4(const char* label,float col[4],ImGuiColorEditFlags flags);
func ColorEdit4(label string, col []float32, flag ImGuiColorEditFlags) bool {
	_label, _col := c.CStr(label), &col[0]
	defer usf.Free(_label)
	return giLib.Call(_func_igColorEdit4_, []interface{}{&_label, &_col, &flag}).BoolFree()
}

// CIMGUI_API bool igColorPicker3(const char* label,float col[3],ImGuiColorEditFlags flags);
func ColorPicker3(label string, col []float32, flag ImGuiColorEditFlags) bool {
	_label, _col := c.CStr(label), &col[0]
	defer usf.Free(_label)
	return giLib.Call(_func_igColorPicker3_, []interface{}{&_label, &_col, &flag}).BoolFree()
}

// CIMGUI_API bool igColorPicker4(const char* label,float col[4],ImGuiColorEditFlags flags,const float* ref_col);
func ColorPicker4(label string, col []float32, flag ImGuiColorEditFlags, ref_col *float32) bool {
	_label, _col := c.CStr(label), &col[0]
	defer usf.Free(_label)
	return giLib.Call(_func_igColorPicker4_, []interface{}{&_label, &_col, &flag, &ref_col}).BoolFree()
}

// CIMGUI_API bool igColorButton(const char* desc_id,const ImVec4 col,ImGuiColorEditFlags flags,const ImVec2 size);
func ColorButtonVec(desc_id string, col ImVec4, flag ImGuiColorEditFlags, size ImVec2) bool {
	_des_id := c.CStr(desc_id)
	defer usf.Free(_des_id)
	return giLib.Call(_func_igColorButton_, []interface{}{&_des_id, &col, &flag, &size}).BoolFree()
}

// CIMGUI_API bool igColorButton(const char* desc_id,const ImVec4 col,ImGuiColorEditFlags flags,const ImVec2 size);
func ColorButton(desc_id string, col1, col2, col3, col4 float32, flag ImGuiColorEditFlags, w, h float32) bool {
	_des_id, col, size := c.CStr(desc_id), NewImVec4Float(col1, col2, col3, col4), NewImVec2Float(w, h)
	defer usf.Free(_des_id)
	defer col.Destroy()
	defer size.Destroy()
	return giLib.Call(_func_igColorButton_, []interface{}{&_des_id, col, &flag, size}).BoolFree()
}

// CIMGUI_API void igSetColorEditOptions(ImGuiColorEditFlags flags);
func SetColorEditOptions(flag ImGuiColorEditFlags) {
	giLib.Call(_func_igSetColorEditOptions_, []interface{}{&flag})
}

// CIMGUI_API bool igTreeNode_Str(const char* label);
func TreeNodeStr(label string) bool {
	l := c.CStr(label)
	defer usf.Free(l)
	return giLib.Call(_func_igTreeNode_Str_, []interface{}{&l}).BoolFree()
}

// CIMGUI_API bool igTreeNode_StrStr(const char* str_id,const char* fmt,...);
func TreeNodeStrStr(ptr_id string, str string) bool {
	l, s := c.CStr(ptr_id), c.CStr(str)
	defer usf.Free(l)
	defer usf.Free(s)
	return giLib.Call(_func_igTreeNode_StrStr_, []interface{}{&l, &s}).BoolFree()
}

// CIMGUI_API bool igTreeNode_Ptr(const void* ptr_id,const char* fmt,...);
func TreeNodePtr(ptr_id unsafe.Pointer, str string) bool {
	s := c.CStr(str)
	defer usf.Free(s)
	return giLib.Call(_func_igTreeNode_Ptr_, []interface{}{&ptr_id, &s}).BoolFree()
}

// CIMGUI_API bool igTreeNodeV_Str(const char* str_id,const char* fmt,va_list args);
// CIMGUI_API bool igTreeNodeV_Ptr(const void* ptr_id,const char* fmt,va_list args);

// CIMGUI_API bool igTreeNodeEx_Str(const char* label,ImGuiTreeNodeFlags flags);
func TreeNodeExStr(label string, flags ImGuiTreeNodeFlags) bool {
	l := c.CStr(label)
	defer usf.Free(l)
	return giLib.Call(_func_igTreeNodeEx_Str_, []interface{}{&l, &flags}).BoolFree()
}

// CIMGUI_API bool igTreeNodeEx_StrStr(const char* str_id,ImGuiTreeNodeFlags flags,const char* fmt,...);
func TreeNodeExStrStr(id string, flags ImGuiTreeNodeFlags, str string) bool {
	l, s := c.CStr(id), c.CStr(str)
	defer usf.Free(l)
	defer usf.Free(s)
	return giLib.Call(_func_igTreeNodeEx_StrStr_, []interface{}{&l, &flags, &s}).BoolFree()
}

// CIMGUI_API bool igTreeNodeEx_Ptr(const void* ptr_id,ImGuiTreeNodeFlags flags,const char* fmt,...);
func TreeNodeExPtr(ptr_id unsafe.Pointer, flags ImGuiTreeNodeFlags, str string) bool {
	s := c.CStr(str)
	defer usf.Free(s)
	return giLib.Call(_func_igTreeNodeEx_Ptr_, []interface{}{&ptr_id, &flags, &s}).BoolFree()
}

// CIMGUI_API bool igTreeNodeExV_Str(const char* str_id,ImGuiTreeNodeFlags flags,const char* fmt,va_list args);
// CIMGUI_API bool igTreeNodeExV_Ptr(const void* ptr_id,ImGuiTreeNodeFlags flags,const char* fmt,va_list args);

// CIMGUI_API void igTreePush_Str(const char* str_id);
func TreePushStr(str_id string) {
	id := c.CStr(str_id)
	defer usf.Free(id)
	giLib.Call(_func_igTreePush_Str_, []interface{}{&id})
}

// CIMGUI_API void igTreePush_Ptr(const void* ptr_id);
func TreePushPtr(ptr_id unsafe.Pointer) {
	giLib.Call(_func_igTreePush_Ptr_, []interface{}{&ptr_id})
}

// CIMGUI_API void igTreePop(void);
func TreePop() { giLib.Call(_func_igTreePop_, nil) }

// CIMGUI_API float igGetTreeNodeToLabelSpacing(void);
func GetTreeNodeToLabelSpacing() float32 {
	return giLib.Call(_func_igGetTreeNodeToLabelSpacing_, nil).F32Free()
}

// CIMGUI_API bool igCollapsingHeader_TreeNodeFlags(const char* label,ImGuiTreeNodeFlags flags);
func CollapsingHeaderTreeNodeFlags(label string, flag ImGuiTreeNodeFlags) bool {
	l := c.CStr(label)
	defer usf.Free(l)
	return giLib.Call(_func_igCollapsingHeader_TreeNodeFlags_, []interface{}{&l, &flag}).BoolFree()

}

// CIMGUI_API bool igCollapsingHeader_BoolPtr(const char* label,bool* p_visible,ImGuiTreeNodeFlags flags);
func CollapsingHeaderBoolPtr(label string, visible *bool, flag ImGuiTreeNodeFlags) bool {
	l, v := c.CStr(label), c.CBool(*visible)
	defer usf.Free(l)
	_v := &v
	b := giLib.Call(_func_igCollapsingHeader_BoolPtr_, []interface{}{&l, &_v, &flag}).BoolFree()

	(*visible) = c.GoBool(v)
	return b
}

// CIMGUI_API void igSetNextItemOpen(bool is_open,ImGuiCond cond);
func SetNextItemOpen(is_open bool, cond ImGuiCond) {
	o := c.CBool(is_open)
	giLib.Call(_func_igSetNextItemOpen_, []interface{}{&o, &cond})
}

// CIMGUI_API bool igSelectable_Bool(const char* label,bool selected,ImGuiSelectableFlags flags,const ImVec2 size);
func SelectableVec2(label string, selected bool, flags ImGuiSelectableFlags, size ImVec2) bool {
	lb, selt := c.CStr(label), c.CBool(selected)
	defer usf.Free(lb)
	return giLib.Call(_func_igSelectable_Bool_, []interface{}{&lb, &selt, &flags, &size}).BoolFree()

}

// CIMGUI_API bool igSelectable_Bool(const char* label,bool selected,ImGuiSelectableFlags flags,const ImVec2 size);
func Selectable(label string, selected bool, flags ImGuiSelectableFlags, w, h float32) bool {
	lb, selt, size := c.CStr(label), c.CBool(selected), NewImVec2Float(w, h)
	defer usf.Free(lb)
	defer size.Destroy()
	return giLib.Call(_func_igSelectable_Bool_, []interface{}{&lb, &selt, &flags, size}).BoolFree()

}

// CIMGUI_API bool igSelectable_BoolPtr(const char* label,bool* p_selected,ImGuiSelectableFlags flags,const ImVec2 size);
func SelectablelPtrVec2(label string, selected *bool, flags ImGuiSelectableFlags, size ImVec2) bool {
	lb, selt := c.CStr(label), c.CBool(*selected)
	defer usf.Free(lb)
	_selt := &selt
	p := giLib.Call(_func_igSelectable_BoolPtr_, []interface{}{&lb, &_selt, &flags, &size}).BoolFree()

	(*selected) = c.GoBool((*_selt))
	return p
}

// CIMGUI_API bool igSelectable_BoolPtr(const char* label,bool* p_selected,ImGuiSelectableFlags flags,const ImVec2 size);
func SelectablelPtr(label string, selected *bool, flags ImGuiSelectableFlags, w, h float32) bool {
	lb, selt, size := c.CStr(label), c.CBool(*selected), NewImVec2Float(w, h)
	defer usf.Free(lb)
	defer size.Destroy()
	_selt := &selt
	p := giLib.Call(_func_igSelectable_BoolPtr_, []interface{}{&lb, &_selt, &flags, size}).BoolFree()

	(*selected) = c.GoBool((*_selt))
	return p
}

// CIMGUI_API bool igBeginListBox(const char* label,const ImVec2 size);
func BeginListBoxVec2(label string, size ImVec2) bool {
	lb := c.CStr(label)
	defer usf.Free(lb)
	return giLib.Call(_func_igBeginListBox_, []interface{}{&lb, &size}).BoolFree()

}

// CIMGUI_API bool igBeginListBox(const char* label,const ImVec2 size);
func BeginListBox(label string, w, h float32) bool {
	lb, size := c.CStr(label), NewImVec2Float(w, h)
	defer usf.Free(lb)
	return giLib.Call(_func_igBeginListBox_, []interface{}{&lb, size}).BoolFree()

}

// CIMGUI_API void igEndListBox(void);
func EndListBox() { giLib.Call(_func_igEndListBox_, nil) }

// CIMGUI_API bool igListBox_Str_arr(const char* label,int* current_item,const char* const items[],int items_count,int height_in_items);
func ListBoxStrs(label string, item_curr *int32, items []string, height_int_items int32) bool {
	_label, _count := c.CStr(label), int32(len(items))
	defer usf.Free(_label)
	__items := c.NewCStrs(items)
	defer __items.Free()
	_items := __items.Ptr()

	return giLib.Call(_func_igListBox_Str_arr_, []interface{}{&_label, &item_curr, &_items, &_count, &height_int_items}).BoolFree()

}

// func NewListBoxCallback(fn func(data *ImGuiInputTextCallbackData) int32) *c.Fn {
//  return c.NewFn(c.AbiDefault, c.I32, []c.Type{c.Pointer},
//      func(args []c.Value, ret c.Value) {
//          d := (*ImGuiInputTextCallbackData)(args[0].Ptr())
//          ret.SetI32(fn(d))
//      })
// }
// // CIMGUI_API bool igListBox_FnBoolPtr(const char* label,int* current_item,bool(*items_getter)(void* data,int idx,const char** out_text),void* data,int items_count,int height_in_items);
// func ListBoxFnBoolPtr(label string, item_curr *int32, items []string, height_int_items int32) bool {
//  _label, _count := c.CStr(label), int32(len(items))
//  defer usf.Free(_label)
//  __items := c.NewCStrs(items)
//  defer __items.Free()
//  _items := __items.Ptr()

//  return giLib.Call(_func_igListBox_FnBoolPtr_, []interface{}{&_label, &item_curr, &_items, &_count, &height_int_items}).BoolFree()

// }

// CIMGUI_API void igPlotLines_FloatPtr(const char* label,const float* values,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size,int stride);
// IMGUI_API void PlotLines(const char* label, const float* values, int values_count, int values_offset = 0, const char* overlay_text = NULL, float scale_min = FLT_MAX, float scale_max = FLT_MAX, ImVec2 graph_size = ImVec2(0, 0), int stride = sizeof(float));
func PlotLinesVec(label string, values []float32, offset int32,
	overlay_text string, scale_min, scale_max float32, graph_size ImVec2, stride int32) {
	_label, _overlay := c.CStr(label), c.CStr(overlay_text)
	defer usf.Free(_label)
	defer usf.Free(_overlay)

	_value := &values[0]
	_vCount := int32(len(values))
	_vOffset := int32(offset)
	giLib.Call(_func_igPlotLines_FloatPtr_, []interface{}{

		&_label, &_value, &_vCount, &_vOffset,
		&_overlay, &scale_min, &scale_max, &graph_size, &stride})
}

func PlotLines(label string, values []float32, offset int32,
	overlay_text string, scale_min, scale_max float32, graphW, graphH float32, stride int32) {
	_label, _overlay, graph_size := c.CStr(label), c.CStr(overlay_text), NewImVec2Float(graphW, graphH)
	defer usf.Free(_label)
	defer usf.Free(_overlay)
	defer graph_size.Destroy()

	_value := &values[0]
	_vCount := int32(len(values))
	_vOffset := int32(offset)
	giLib.Call(_func_igPlotLines_FloatPtr_, []interface{}{

		&_label, &_value, &_vCount, &_vOffset,
		&_overlay, &scale_min, &scale_max, graph_size, &stride})
}

func NewPlotLinesCallback(fn func(data unsafe.Pointer, idx int32) float32) *c.Callback {
	cb := c.NewCallback(c.AbiDefault, c.F32, []c.Type{c.Pointer, c.I32})
	cb.CallbackFunc = fn
	return cb
}

// CIMGUI_API void igPlotLines_FnFloatPtr(const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size);
func PlotLinesFn(label string, fn *c.Callback, dataFristAddr interface{}, count, offset int32,
	overlay_text string, scale_min, scale_max float32, graph_size ImVec2, stride int32) {
	_label, _overlay := c.CStr(label), c.CStr(overlay_text)
	defer usf.Free(_label)
	defer usf.Free(_overlay)

	if fn.CallbackCvt == nil {
		fn.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
			data := args[0].Ptr()
			idx := args[1].I32()
			_fn := cb.CallbackFunc.(func(data unsafe.Pointer, idx int32) float32)
			if _fn != nil {
				ret.SetF32(_fn(data, idx))
			}
		}
	}
	_vFn := fn.Cfunc
	_value := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataFristAddr), 2)))[1]
	_vOffset := int32(offset)
	giLib.Call(_func_igPlotLines_FnFloatPtr_, []interface{}{

		&_label, &_vFn, &_value, &count, &_vOffset,
		&_overlay, &scale_min, &scale_max, &graph_size, &stride})
}

// CIMGUI_API void igPlotHistogram_FloatPtr(const char* label,const float* values,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size,int stride);
func PlotHistogramVec2(label string, values []float32, offset int32,
	overlay_text string, scale_min, scale_max float32, graph_size ImVec2, stride int32) {
	_label, _overlay := c.CStr(label), c.CStr(overlay_text)
	defer usf.Free(_label)
	defer usf.Free(_overlay)

	_value := &values[0]
	_vCount := int32(len(values))
	_vOffset := int32(offset)
	giLib.Call(_func_igPlotHistogram_FloatPtr_, []interface{}{

		&_label, &_value, &_vCount, &_vOffset,
		&_overlay, &scale_min, &scale_max, &graph_size, &stride})
}

// CIMGUI_API void igPlotHistogram_FloatPtr(const char* label,const float* values,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size,int stride);
func PlotHistogram(label string, values []float32, offset int32,
	overlay_text string, scale_min, scale_max float32, graphW, graphH float32, stride int32) {
	_label, _overlay, graph_size := c.CStr(label), c.CStr(overlay_text), NewImVec2Float(graphW, graphH)
	defer usf.Free(_label)
	defer usf.Free(_overlay)
	defer graph_size.Destroy()

	_value := &values[0]
	_vCount := int32(len(values))
	_vOffset := int32(offset)
	giLib.Call(_func_igPlotHistogram_FloatPtr_, []interface{}{

		&_label, &_value, &_vCount, &_vOffset,
		&_overlay, &scale_min, &scale_max, graph_size, &stride})
}

// CIMGUI_API void igPlotHistogram_FnFloatPtr(const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size);
func PlotHistogramFn(label string, fn *c.Callback, dataFristAddr interface{}, count, offset int32,
	overlay_text string, scale_min, scale_max float32, graph_size ImVec2, stride int32) {
	_label, _overlay := c.CStr(label), c.CStr(overlay_text)
	defer usf.Free(_label)
	defer usf.Free(_overlay)

	if fn.CallbackCvt == nil {
		fn.CallbackCvt = func(cb *c.Callback, args []*c.Value, ret *c.Value) {
			data := args[0].Ptr()
			idx := args[1].I32()
			_fn := cb.CallbackFunc.(func(data unsafe.Pointer, idx int32) float32)
			if _fn != nil {
				ret.SetF32(_fn(data, idx))
			}
		}
	}
	_vFn := fn.Cfunc
	_value := (*(*[]unsafe.Pointer)(usf.Slice(unsafe.Pointer(&dataFristAddr), 2)))[1]
	_vOffset := int32(offset)
	giLib.Call(_func_igPlotHistogram_FnFloatPtr_, []interface{}{

		&_label, &_vFn, &_value, &count, &_vOffset,
		&_overlay, &scale_min, &scale_max, &graph_size, &stride})
}

// CIMGUI_API void igValue_Bool(const char* prefix,bool b);
func ValueBool(prefix string, b bool) {
	p, _b := c.CStr(prefix), c.CBool(b)
	defer usf.Free(p)
	giLib.Call(_func_igValue_Bool_, []interface{}{&p, &_b})
}

// CIMGUI_API void igValue_Int(const char* prefix,int v);
func ValueInt(prefix string, v int32) {
	p := c.CStr(prefix)
	defer usf.Free(p)
	giLib.Call(_func_igValue_Int_, []interface{}{&p, &v})
}

// CIMGUI_API void igValue_Uint(const char* prefix,unsigned int v);
func ValueUint(prefix string, v uint32) {
	p := c.CStr(prefix)
	defer usf.Free(p)
	giLib.Call(_func_igValue_Uint_, []interface{}{&p, &v})
}

// CIMGUI_API void igValue_Float(const char* prefix,float v,const char* float_format);
func ValueFloat(prefix string, v float32, format string) {
	p, f := c.CStr(prefix), c.CStr(format)
	defer usf.Free(p)
	defer usf.Free(f)
	giLib.Call(_func_igValue_Float_, []interface{}{&p, &v, &f})
}

// CIMGUI_API bool igBeginMenuBar(void);
func BeginMenuBar() bool {
	return giLib.Call(_func_igBeginMenuBar_, nil).BoolFree()
}

// CIMGUI_API void igEndMenuBar(void);
func EndMenuBar() {
	giLib.Call(_func_igEndMenuBar_, nil)
}

// CIMGUI_API bool igBeginMainMenuBar(void);
func BeginMainMenuBar() bool {
	return giLib.Call(_func_igBeginMainMenuBar_, nil).BoolFree()
}

// CIMGUI_API void igEndMainMenuBar(void);
func EndMainMenuBar() {
	giLib.Call(_func_igEndMainMenuBar_, nil)
}

// CIMGUI_API bool igBeginMenu(const char* label,bool enabled);
func BeginMenu(label string, enabled bool) bool {
	l, e := c.CStr(label), c.CBool(enabled)
	defer usf.Free(l)
	return giLib.Call(_func_igBeginMenu_, []interface{}{&l, &e}).BoolFree()
}

// CIMGUI_API void igEndMenu(void);
func EndMenu() {
	giLib.Call(_func_igEndMenu_, nil)
}

// CIMGUI_API bool igMenuItem_Bool(const char* label,const char* shortcut,bool selected,bool enabled);
func MenuItem(label string, shortcut string, selected, enabled bool) bool {
	_label, _shortcut, _selected, _enabled := c.CStr(label), c.CStr(shortcut), c.CBool(selected), c.CBool(enabled)
	defer usf.Free(_label)
	defer usf.Free(_shortcut)
	return giLib.Call(_func_igMenuItem_Bool_, []interface{}{&_label, &_shortcut, &_selected, &_enabled}).BoolFree()

}

// CIMGUI_API bool igMenuItem_BoolPtr(const char* label,const char* shortcut,bool* p_selected,bool enabled);
func MenuItemPtr(label string, shortcut string, selected *bool, enabled bool) bool {
	_label, _shortcut, _selected, _enabled := c.CStr(label), c.CStr(shortcut), c.CBool(*selected), c.CBool(enabled)
	defer usf.Free(_label)
	defer usf.Free(_shortcut)
	_selt := &_selected
	p := giLib.Call(_func_igMenuItem_BoolPtr_, []interface{}{&_label, &_shortcut, &_selt, &_enabled}).BoolFree()

	(*selected) = c.GoBool(*_selt)
	return p
}

// CIMGUI_API bool igBeginTooltip(void);
func BeginTooltip() bool {
	return giLib.Call(_func_igBeginTooltip_, nil).BoolFree()
}

// CIMGUI_API void igEndTooltip(void);
func EndTooltip() {
	giLib.Call(_func_igEndTooltip_, nil)
}

// CIMGUI_API void igSetTooltip(const char* fmt,...);
func SetTooltip(format string) {
	fm := c.CStr(format)
	defer usf.Free(fm)
	giLib.Call(_func_igSetTooltip_, []interface{}{&fm})
}

// CIMGUI_API void igSetTooltipV(const char* fmt,va_list args);

// CIMGUI_API bool igBeginItemTooltip(void);
func BeginItemTooltip() bool {
	return giLib.Call(_func_igBeginItemTooltip_, nil).BoolFree()
}

// CIMGUI_API void igSetItemTooltip(const char* fmt,...);
func SetItemTooltip(format string) {
	fm := c.CStr(format)
	defer usf.Free(fm)
	giLib.Call(_func_igSetItemTooltip_, []interface{}{&fm})
}

// CIMGUI_API void igSetItemTooltipV(const char* fmt,va_list args);

// CIMGUI_API bool igBeginPopup(const char* str_id,ImGuiWindowFlags flags);
func BeginPopup(str_id string, flags ImGuiWindowFlags) bool {
	id := c.CStr(str_id)
	defer usf.Free(id)
	return giLib.Call(_func_igBeginPopup_, []interface{}{&id, &flags}).BoolFree()
}

// CIMGUI_API bool igBeginPopupModal(const char* name,bool* p_open,ImGuiWindowFlags flags);
func BeginPopupModal(str_id string, open *bool, flags ImGuiWindowFlags) bool {
	id, _open := c.CStr(str_id), c.CBool(*open)
	defer usf.Free(id)
	p_open := &_open
	p := giLib.Call(_func_igBeginPopupModal_, []interface{}{&id, &p_open, &flags}).BoolFree()

	(*open) = c.GoBool(*p_open)
	return p
}

// CIMGUI_API void igEndPopup(void);
func EndPopup() {
	giLib.Call(_func_igEndPopup_, nil)
}

// CIMGUI_API void igOpenPopup_Str(const char* str_id,ImGuiPopupFlags popup_flags);
func OpenPopupStr(str_id string, flags ImGuiPopupFlags) {
	id := c.CStr(str_id)
	defer usf.Free(id)
	giLib.Call(_func_igOpenPopup_Str_, []interface{}{&id, &flags})
}

// CIMGUI_API void igOpenPopup_ID(ImGuiID id,ImGuiPopupFlags popup_flags);
func OpenPopupID(id ImGuiID, flags ImGuiPopupFlags) {
	giLib.Call(_func_igOpenPopup_ID_, []interface{}{&id, &flags})
}

// CIMGUI_API void igOpenPopupOnItemClick(const char* str_id,ImGuiPopupFlags popup_flags);
func OpenPopupOnItemClick(str_id string, flags ImGuiPopupFlags) {
	id := c.CStr(str_id)
	defer usf.Free(id)
	giLib.Call(_func_igOpenPopupOnItemClick_, []interface{}{&id, &flags})
}

// CIMGUI_API void igCloseCurrentPopup(void);
func CloseCurrentPopup() {
	giLib.Call(_func_igCloseCurrentPopup_, nil)
}

// CIMGUI_API bool igBeginPopupContextItem(const char* str_id,ImGuiPopupFlags popup_flags);
func BeginPopupContextItem(str_id string, flags ImGuiPopupFlags) bool {
	id := c.CStr(str_id)
	defer usf.Free(id)
	return giLib.Call(_func_igBeginPopupContextItem_, []interface{}{&id, &flags}).BoolFree()

}

// CIMGUI_API bool igBeginPopupContextWindow(const char* str_id,ImGuiPopupFlags popup_flags);
func BeginPopupContextWindow(str_id string, flags ImGuiPopupFlags) bool {
	id := c.CStr(str_id)
	defer usf.Free(id)
	return giLib.Call(_func_igBeginPopupContextWindow_, []interface{}{&id, &flags}).BoolFree()

}

// CIMGUI_API bool igBeginPopupContextVoid(const char* str_id,ImGuiPopupFlags popup_flags);
func BeginPopupContextVoid(str_id string, flags ImGuiPopupFlags) bool {
	id := c.CStr(str_id)
	defer usf.Free(id)
	return giLib.Call(_func_igBeginPopupContextVoid_, []interface{}{&id, &flags}).BoolFree()

}

// CIMGUI_API bool igIsPopupOpen_Str(const char* str_id,ImGuiPopupFlags flags);
func IsPopupOpen(str_id string, flags ImGuiPopupFlags) bool {
	id := c.CStr(str_id)
	defer usf.Free(id)
	return giLib.Call(_func_igIsPopupOpen_Str_, []interface{}{&id, &flags}).BoolFree()

}

// CIMGUI_API bool igBeginTable(const char* str_id,int column,ImGuiTableFlags flags,const ImVec2 outer_size,float inner_width);
func BeginTableVec2(str_id string, col int32, flags ImGuiTableFlags, outer_size ImVec2, inner_width float32) bool {
	id := c.CStr(str_id)
	defer usf.Free(id)
	return giLib.Call(_func_igBeginTable_, []interface{}{&id, &col, &flags, &outer_size, &inner_width}).BoolFree()

}

// CIMGUI_API bool igBeginTable(const char* str_id,int column,ImGuiTableFlags flags,const ImVec2 outer_size,float inner_width);
func BeginTable(str_id string, col int32, flags ImGuiTableFlags, w, h float32, inner_width float32) bool {
	id, outer_size := c.CStr(str_id), NewImVec2Float(w, h)
	defer usf.Free(id)
	defer outer_size.Destroy()
	return giLib.Call(_func_igBeginTable_, []interface{}{&id, &col, &flags, outer_size, &inner_width}).BoolFree()

}

// BeginTable(const char* str_id, int column, ImGuiTableFlags flags = 0, const ImVec2& outer_size = ImVec2(0.0f, 0.0f), float inner_width = 0.0f);
func BeginTableDefault(str_id string, col int32) bool {
	id := c.CStr(str_id)
	defer usf.Free(id)
	return giLib.Call(_func_igBeginTable_, []interface{}{&id, &col, _u32_zero, _vec2_zero, _f32_zero}).BoolFree()

}

// CIMGUI_API void igEndTable(void);
func EndTable() { giLib.Call(_func_igEndTable_, nil) }

// CIMGUI_API void igTableNextRow(ImGuiTableRowFlags row_flags,float min_row_height);
func TableNextRow(row_flags ImGuiTableRowFlags, min_row_height float32) {
	giLib.Call(_func_igTableNextRow_, []interface{}{&row_flags, &min_row_height})
}

// IMGUI_API void TableNextRow(ImGuiTableRowFlags row_flags = 0, float min_row_height = 0.0f); // append into the first cell of a new row.
func TableNextRowDefault() {
	giLib.Call(_func_igTableNextRow_, []interface{}{_u32_zero, _f32_zero})
}

// CIMGUI_API bool igTableNextColumn(void);
func TableNextColumn() bool {
	return giLib.Call(_func_igTableNextColumn_, nil).BoolFree()
}

// CIMGUI_API bool igTableSetColumnIndex(int column_n);
func TableSetColumnIndex(column_n int32) bool {
	return giLib.Call(_func_igTableSetColumnIndex_, []interface{}{&column_n}).BoolFree()
}

// CIMGUI_API void igTableSetupColumn(const char* label,ImGuiTableColumnFlags flags,float init_width_or_weight,ImGuiID user_id);
func TableSetupColumn(label string, flags ImGuiTableColumnFlags, init_width_or_weight float32, user_id ImGuiID) {
	l := c.CStr(label)
	defer usf.Free(l)
	giLib.Call(_func_igTableSetupColumn_, []interface{}{&l, &flags, &init_width_or_weight, &user_id})

}

// IMGUI_API void          TableSetupColumn(const char* label, ImGuiTableColumnFlags flags = 0, float init_width_or_weight = 0.0f, ImGuiID user_id = 0);
func TableSetupColumnDefaule(label string) {
	l, flags, init_width_or_weight, user_id := c.CStr(label), ImGuiTableColumnFlags(0), float32(0), ImGuiID(0)
	defer usf.Free(l)
	giLib.Call(_func_igTableSetupColumn_, []interface{}{&l, &flags, &init_width_or_weight, &user_id})

}

// CIMGUI_API void igTableSetupScrollFreeze(int cols,int rows);
func TableSetupScrollFreeze(cols, rows int32) {
	giLib.Call(_func_igTableSetupScrollFreeze_, []interface{}{&cols, &rows})
}

// CIMGUI_API void igTableHeadersRow(void);
func TableHeadersRow() { giLib.Call(_func_igTableHeadersRow_, nil) }

// CIMGUI_API void igTableHeader(const char* label);
func TableHeader(label string) {
	l := c.CStr(label)
	defer usf.Free(l)
	giLib.Call(_func_igTableHeader_, []interface{}{&l})
}

// CIMGUI_API ImGuiTableSortSpecs* igTableGetSortSpecs(void);
func TableGetSortSpecs() *ImGuiTableSortSpecs {
	return (*ImGuiTableSortSpecs)(giLib.Call(_func_igTableGetSortSpecs_, nil).PtrFree())
}

// CIMGUI_API int igTableGetColumnCount(void);
func TableGetColumnCount() int32 {
	return giLib.Call(_func_igTableGetColumnCount_, nil).I32Free()
}

// CIMGUI_API int igTableGetColumnIndex(void);
func TableGetColumnIndex() int32 {
	return giLib.Call(_func_igTableGetColumnIndex_, nil).I32Free()
}

// CIMGUI_API int igTableGetRowIndex(void);
func TableGetRowIndex() int32 {
	return giLib.Call(_func_igTableGetRowIndex_, nil).I32Free()
}

// CIMGUI_API const char* igTableGetColumnName_Int(int column_n);
func TableGetColumnName(column_n int32) string {
	return giLib.Call(_func_igTableGetColumnName_Int_, []interface{}{&column_n}).StrFree()
}

// CIMGUI_API ImGuiTableColumnFlags igTableGetColumnFlags(int column_n);
func TableGetColumnFlags(column_n int32) ImGuiTableColumnFlags {
	return ImGuiTableColumnFlags(giLib.Call(_func_igTableGetColumnFlags_, []interface{}{&column_n}).U32Free())
}

// CIMGUI_API void igTableSetColumnEnabled(int column_n,bool v);
func TableSetColumnEnabled(column_n int32, enable bool) {
	v := c.CBool(enable)
	giLib.Call(_func_igTableSetColumnEnabled_, []interface{}{&column_n, &v})
}

// target: ABGR:0x[xx|xx|xx|xx]
// CIMGUI_API void igTableSetBgColor(ImGuiTableBgTarget target,ImU32 color,int column_n);
func TableSetBgColor(target ImGuiTableBgTarget, color uint32, column_n int32) {
	giLib.Call(_func_igTableSetBgColor_, []interface{}{&target, &color, &column_n})
}

// CIMGUI_API void igColumns(int count,const char* id,bool border);
func Columns(count int32, id string, border bool) {
	i, b := c.CStr(id), c.CBool(border)
	defer usf.Free(i)
	giLib.Call(_func_igColumns_, []interface{}{&count, &i, &b})
}

// IMGUI_API void Columns(int count = 1, const char* id = NULL, bool border = true);
func ColumnsDefault() {
	giLib.Call(_func_igColumns_, []interface{}{_i32_one, _ptr_zero, _bool_true})

}

// CIMGUI_API void igNextColumn(void);
func NextColumn() { giLib.Call(_func_igNextColumn_, nil) }

// CIMGUI_API int igGetColumnIndex(void);
func GetColumnIndex() int32 { return giLib.Call(_func_igGetColumnIndex_, nil).I32Free() }

// CIMGUI_API float igGetColumnWidth(int column_index);
func GetColumnWidth(column_index int32) float32 {
	return giLib.Call(_func_igGetColumnWidth_, []interface{}{&column_index}).F32Free()
}

// CIMGUI_API void igSetColumnWidth(int column_index,float width);
func SetColumnWidth(column_index int32, width float32) {
	giLib.Call(_func_igSetColumnWidth_, []interface{}{&column_index, &width})
}

// CIMGUI_API float igGetColumnOffset(int column_index);
func GetColumnOffset(column_index int32) float32 {
	return giLib.Call(_func_igGetColumnOffset_, []interface{}{&column_index}).F32Free()
}

// CIMGUI_API void igSetColumnOffset(int column_index,float offset_x);
func SetColumnOffset(column_index int32, offset_x float32) {
	giLib.Call(_func_igSetColumnOffset_, []interface{}{&column_index, &offset_x})
}

// CIMGUI_API int igGetColumnsCount(void);
func GetColumnsCount() int32 {
	return giLib.Call(_func_igGetColumnsCount_, nil).I32Free()
}

// CIMGUI_API bool igBeginTabBar(const char* str_id,ImGuiTabBarFlags flags);
func BeginTabBar(str_id string, flags ImGuiTabBarFlags) bool {
	i := c.CStr(str_id)
	defer usf.Free(i)
	return giLib.Call(_func_igBeginTabBar_, []interface{}{&i, &flags}).BoolFree()
}

// IMGUI_API bool BeginTabBar(const char* str_id, ImGuiTabBarFlags flags = 0);
// create and append into a TabBar
func BeginTabBarDefault(str_id string) bool {
	i := c.CStr(str_id)
	defer usf.Free(i)
	return giLib.Call(_func_igBeginTabBar_, []interface{}{&i, _i32_zero}).BoolFree()
}

// CIMGUI_API void igEndTabBar(void);
func EndTabBar() { giLib.Call(_func_igEndTabBar_, nil) }

// CIMGUI_API bool igBeginTabItem(const char* label,bool* p_open,ImGuiTabItemFlags flags);
func BeginTabItem(label string, open *bool, flags ImGuiTabItemFlags) bool {
	l, o := c.CStr(label), c.CBool(*open)
	defer usf.Free(l)

	_o := &o
	p := giLib.Call(_func_igBeginTabItem_, []interface{}{&l, &_o, &flags}).BoolFree()
	(*open) = c.GoBool(*_o)
	return p
}

// IMGUI_API bool BeginTabItem(const char* label, bool* p_open = NULL, ImGuiTabItemFlags flags = 0);
// create a Tab. Returns true if the Tab is selected.
func BeginTabItemDefault(label string) bool {
	l := c.CStr(label)
	defer usf.Free(l)
	return giLib.Call(_func_igBeginTabItem_, []interface{}{&l, _ptr_zero, _u32_zero}).BoolFree()

}

// CIMGUI_API void igEndTabItem(void);
func EndTabItem() { giLib.Call(_func_igEndTabItem_, nil) }

// CIMGUI_API bool igTabItemButton(const char* label,ImGuiTabItemFlags flags);
func TabItemButton(label string, flags ImGuiTabItemFlags) bool {
	l := c.CStr(label)
	defer usf.Free(l)
	return giLib.Call(_func_igTabItemButton_, []interface{}{&l, &flags}).BoolFree()
}

// CIMGUI_API void igSetTabItemClosed(const char* tab_or_docked_window_label);
func SetTabItemClosed(tab_or_docked_window_label string) {
	l := c.CStr(tab_or_docked_window_label)
	defer usf.Free(l)
	giLib.Call(_func_igSetTabItemClosed_, []interface{}{&l})
}

// CIMGUI_API ImGuiID igDockSpace(ImGuiID id,const ImVec2 size,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class);
func DockSpaceVec2(id ImGuiID, size ImVec2, flags ImGuiDockNodeFlags, window_class *ImGuiWindowClass) ImGuiID {
	return ImGuiID(giLib.Call(_func_igDockSpace_, []interface{}{&id, &size, &flags, &window_class}).U32Free())

}

// CIMGUI_API ImGuiID igDockSpace(ImGuiID id,const ImVec2 size,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class);
func DockSpace(id ImGuiID, w, h float32, flags ImGuiDockNodeFlags, window_class *ImGuiWindowClass) ImGuiID {
	size := NewImVec2Float(w, h)
	defer size.Destroy()
	return ImGuiID(giLib.Call(_func_igDockSpace_, []interface{}{&id, size, &flags, &window_class}).U32Free())

}

// IMGUI_API ImGuiID DockSpace(ImGuiID id, const ImVec2& size = ImVec2(0, 0), ImGuiDockNodeFlags flags = 0, const ImGuiWindowClass* window_class = NULL);
func DockSpaceDefault(id ImGuiID) ImGuiID {
	return ImGuiID(giLib.Call(_func_igDockSpace_, []interface{}{&id, _vec2_zero, _u32_zero, _ptr_zero}).U32Free())

}

// CIMGUI_API ImGuiID igDockSpaceOverViewport(const ImGuiViewport* viewport,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class);
func DockSpaceOverViewport(viewport *ImGuiViewport, flags ImGuiDockNodeFlags, window_class *ImGuiWindowClass) ImGuiID {
	return ImGuiID(giLib.Call(_func_igDockSpaceOverViewport_, []interface{}{&viewport, &flags, &window_class}).U32Free())

}

// IMGUI_API ImGuiID DockSpaceOverViewport(const ImGuiViewport* viewport = NULL, ImGuiDockNodeFlags flags = 0, const ImGuiWindowClass* window_class = NULL);
func DockSpaceOverViewportDefault() ImGuiID {
	return ImGuiID(giLib.Call(_func_igDockSpaceOverViewport_, []interface{}{_ptr_zero, _u32_zero, _ptr_zero}).U32Free())

}

// CIMGUI_API void igSetNextWindowDockID(ImGuiID dock_id,ImGuiCond cond);
func SetNextWindowDockID(dock_id ImGuiID, cond ImGuiCond) {
	giLib.Call(_func_igSetNextWindowDockID_, []interface{}{&dock_id, &cond})
}

// CIMGUI_API void igSetNextWindowClass(const ImGuiWindowClass* window_class);
func SetNextWindowClass(window_class *ImGuiWindowClass) {
	giLib.Call(_func_igSetNextWindowClass_, []interface{}{&window_class})
}

// CIMGUI_API ImGuiID igGetWindowDockID(void);
func GetWindowDockID() ImGuiID {
	return ImGuiID(giLib.Call(_func_igSetNextWindowClass_, nil).U32Free())
}

// CIMGUI_API bool igIsWindowDocked(void);
func IsWindowDocked() bool {
	return giLib.Call(_func_igIsWindowDocked_, nil).BoolFree()
}

// CIMGUI_API void igLogToTTY(int auto_open_depth);
func LogToTTY(auto_open_depth int32) {
	giLib.Call(_func_igLogToTTY_, []interface{}{&auto_open_depth})
}

// IMGUI_API void LogToTTY(int auto_open_depth = -1);                                 // start logging to tty (stdout)
func LogToTTYDefault() {
	giLib.Call(_func_igLogToTTY_, []interface{}{_i32_minus_one})
}

// CIMGUI_API void igLogToFile(int auto_open_depth,const char* filename);
func LogToFile(auto_open_depth int32, filename string) {
	f := c.CStr(filename)
	defer usf.Free(f)
	giLib.Call(_func_igLogToFile_, []interface{}{&auto_open_depth, &f})
}

// IMGUI_API void          LogToFile(int auto_open_depth = -1, const char* filename = NULL);   // start logging to file
func LogToFileDefault() {
	giLib.Call(_func_igLogToFile_, []interface{}{_i32_minus_one, _ptr_zero})
}

// CIMGUI_API void igLogToClipboard(int auto_open_depth);
func LogToClipboard(auto_open_depth int32) {
	giLib.Call(_func_igLogToClipboard_, []interface{}{&auto_open_depth})
}

// IMGUI_API void          LogToClipboard(int auto_open_depth = -1);                           // start logging to OS clipboard
func LogToClipboardDefault() {
	giLib.Call(_func_igLogToClipboard_, []interface{}{_i32_minus_one})
}

// CIMGUI_API void igLogFinish(void);
func LogFinish() { giLib.Call(_func_igLogFinish_, nil) }

// CIMGUI_API void igLogButtons(void);
func LogButtons() { giLib.Call(_func_igLogButtons_, nil) }

// CIMGUI_API void igLogTextV(const char* fmt,va_list args);
func LogTextV(txt string) {
	t := c.CStr(txt)
	defer usf.Free(t)
	giLib.Call(_func_igLogTextV_, []interface{}{&t})
}

// CIMGUI_API bool igBeginDragDropSource(ImGuiDragDropFlags flags);
func BeginDragDropSource(flags ImGuiDragDropFlags) bool {
	return giLib.Call(_func_igBeginDragDropSource_, []interface{}{&flags}).BoolFree()
}

// IMGUI_API bool          BeginDragDropSource(ImGuiDragDropFlags flags = 0);                                      // call after submitting an item which may be dragged. when this return true, you can call SetDragDropPayload() + EndDragDropSource()
func BeginDragDropSourceDefault() bool {
	return giLib.Call(_func_igBeginDragDropSource_, []interface{}{_u32_zero}).BoolFree()

}

// CIMGUI_API bool igSetDragDropPayload(const char* type,const void* data,size_t sz,ImGuiCond cond);
func SetDragDropPayload(typ string, data unsafe.Pointer, size uint64, cond ImGuiCond) bool {
	t := c.CStr(typ)
	defer usf.Free(t)
	return giLib.Call(_func_igSetDragDropPayload_, []interface{}{&t, &data, &size, &cond}).BoolFree()
}

// CIMGUI_API void igEndDragDropSource(void);
func EndDragDropSource() { giLib.Call(_func_igEndDragDropSource_, nil) }

// CIMGUI_API bool igBeginDragDropTarget(void);
func BeginDragDropTarget() bool {
	return giLib.Call(_func_igBeginDragDropTarget_, nil).BoolFree()
}

// CIMGUI_API const ImGuiPayload* igAcceptDragDropPayload(const char* type,ImGuiDragDropFlags flags);
func AcceptDragDropPayload(typ string, flags ImGuiDragDropFlags) *ImGuiPayload {
	t := c.CStr(typ)
	defer usf.Free(t)
	return (*ImGuiPayload)(giLib.Call(_func_igAcceptDragDropPayload_, []interface{}{&t, &flags}).PtrFree())

}

// CIMGUI_API void igEndDragDropTarget(void);
func EndDragDropTarget() { giLib.Call(_func_igEndDragDropTarget_, nil) }

// CIMGUI_API const ImGuiPayload* igGetDragDropPayload(void);
func GetDragDropPayload() *ImGuiPayload {
	return (*ImGuiPayload)(giLib.Call(_func_igGetDragDropPayload_, nil).PtrFree())
}

// CIMGUI_API void igBeginDisabled(bool disabled);
func BeginDisabled(disabled bool) {
	b := c.CBool(disabled)
	giLib.Call(_func_igBeginDisabled_, []interface{}{&b})
}

// CIMGUI_API void igEndDisabled(void);
func EndDisabled() { giLib.Call(_func_igEndDisabled_, nil) }

// CIMGUI_API void igPushClipRect(const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect);
func PushClipRect(clip_rect_min, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	b := c.CBool(intersect_with_current_clip_rect)
	giLib.Call(_func_igPushClipRect_, []interface{}{&clip_rect_min, &clip_rect_min, &b})
}

// CIMGUI_API void igPopClipRect(void);
func PopClipRect() { giLib.Call(_func_igPopClipRect_, nil) }

// CIMGUI_API void igSetItemDefaultFocus(void);
// IMGUI_API void SetItemDefaultFocus(); // make last item the default focused item of a window.
func SetItemDefaultFocus() { giLib.Call(_func_igSetItemDefaultFocus_, nil) }

// CIMGUI_API void igSetKeyboardFocusHere(int offset);
// IMGUI_API void SetKeyboardFocusHere(int offset = 0); // focus keyboard on the next widget. Use positive 'offset' to access sub components of a multiple component widget. Use -1 to access previous widget.
func SetKeyboardFocusHere(offset int32) {
	giLib.Call(_func_igSetKeyboardFocusHere_, []interface{}{&offset})
}

// CIMGUI_API void igSetNextItemAllowOverlap(void);
// IMGUI_API void SetNextItemAllowOverlap(); // allow next item to be overlapped by a subsequent item. Useful with invisible buttons, selectable, treenode covering an area where subsequent items may need to be added. Note that both Selectable() and TreeNode() have dedicated flags doing this.
func SetNextItemAllowOverlap() { giLib.Call(_func_igSetNextItemAllowOverlap_, nil) }

// CIMGUI_API bool igIsItemHovered(ImGuiHoveredFlags flags);
// IMGUI_API bool IsItemHovered(ImGuiHoveredFlags flags = 0);
// is the last item hovered? (and usable, aka not blocked by a popup, etc.). See ImGuiHoveredFlags for more options.
func IsItemHovered(flags ImGuiHoveredFlags) bool {
	return giLib.Call(_func_igIsItemHovered_, []interface{}{&flags}).BoolFree()
}

// CIMGUI_API bool igIsItemActive(void);
//
// IMGUI_API bool IsItemActive();
//
// is the last item active? (e.g. button being held, text field being edited. This will continuously return true while holding mouse button on an item. Items that don't interact will always return false)
func IsItemActive() bool {
	return giLib.Call(_func_igIsItemActive_, nil).BoolFree()
}

// CIMGUI_API bool igIsItemFocused(void);
// IMGUI_API bool IsItemFocused();
// is the last item focused for keyboard/gamepad navigation?
func IsItemFocused() bool {
	return giLib.Call(_func_igIsItemFocused_, nil).BoolFree()
}

// CIMGUI_API bool igIsItemClicked(ImGuiMouseButton mouse_button);
// IMGUI_API bool IsItemClicked(ImGuiMouseButton mouse_button = 0);
// is the last item hovered and mouse clicked on? (**)  == IsMouseClicked(mouse_button) && IsItemHovered()Important. (**) this is NOT equivalent to the behavior of e.g. Button(). Read comments in function definition.
func IsItemClicked(mouse_button ImGuiMouseButton) bool {
	return giLib.Call(_func_igIsItemClicked_, []interface{}{&mouse_button}).BoolFree()
}

// CIMGUI_API bool igIsItemVisible(void);
// IMGUI_API bool IsItemVisible();
// is the last item visible? (items may be out of sight because of clipping/scrolling)
func IsItemVisible() bool {
	return giLib.Call(_func_igIsItemVisible_, nil).BoolFree()
}

// CIMGUI_API bool igIsItemEdited(void);
// IMGUI_API bool IsItemEdited();
// did the last item modify its underlying value this frame? or was pressed? This is generally the same as the "bool" return value of many widgets.
func IsItemEdited() bool {
	return giLib.Call(_func_igIsItemEdited_, nil).BoolFree()
}

// CIMGUI_API bool igIsItemActivated(void);
// IMGUI_API bool IsItemActivated();
// was the last item just made active (item was previously inactive).
func IsItemActivated() bool {
	return giLib.Call(_func_igIsItemActivated_, nil).BoolFree()
}

// CIMGUI_API bool igIsItemDeactivated(void);
// IMGUI_API bool IsItemDeactivated();
// was the last item just made inactive (item was previously active). Useful for Undo/Redo patterns with widgets that require continuous editing.
func IsItemDeactivated() bool {
	return giLib.Call(_func_igIsItemDeactivated_, nil).BoolFree()
}

// CIMGUI_API bool igIsItemDeactivatedAfterEdit(void);
// IMGUI_API bool IsItemDeactivatedAfterEdit();
// was the last item just made inactive and made a value change when it was active? (e.g. Slider/Drag moved). Useful for Undo/Redo patterns with widgets that require continuous editing. Note that you may get false positives (some widgets such as Combo()/ListBox()/Selectable() will return true even when clicking an already selected item).
func IsItemDeactivatedAfterEdit() bool {
	return giLib.Call(_func_igIsItemDeactivatedAfterEdit_, nil).BoolFree()
}

// CIMGUI_API bool igIsItemToggledOpen(void);
// IMGUI_API bool IsItemToggledOpen();
// was the last item open state toggled? set by TreeNode().
func IsItemToggledOpen() bool {
	return giLib.Call(_func_igIsItemToggledOpen_, nil).BoolFree()
}

// CIMGUI_API bool igIsAnyItemHovered(void);
// IMGUI_API bool IsAnyItemHovered();
// is any item hovered?
func IsAnyItemHovered() bool {
	return giLib.Call(_func_igIsAnyItemHovered_, nil).BoolFree()
}

// CIMGUI_API bool igIsAnyItemActive(void);
// IMGUI_API bool          IsAnyItemActive();
// is any item active?
func IsAnyItemActive() bool {
	return giLib.Call(_func_igIsAnyItemActive_, nil).BoolFree()
}

// CIMGUI_API bool igIsAnyItemFocused(void);
// IMGUI_API bool IsAnyItemFocused();
// is any item focused?
func IsAnyItemFocused() bool {
	return giLib.Call(_func_igIsAnyItemFocused_, nil).BoolFree()
}

// CIMGUI_API ImGuiID igGetItemID(void);
// IMGUI_API ImGuiID       GetItemID();
// get ID of last item (~~ often same ImGui::GetID(label) beforehand)
func GetItemID() ImGuiID {
	return ImGuiID(giLib.Call(_func_igGetItemID_, nil).U32Free())
}

// CIMGUI_API void igGetItemRectMin(ImVec2* pOut);
// IMGUI_API ImVec2        GetItemRectMin();
// get upper-left bounding rectangle of the last item (screen space)
func GetItemRectMin() (w, h float32) {
	pOut := NewImVec2Float(0, 0)
	defer pOut.Destroy()
	giLib.Call(_func_igGetItemRectMin_, []interface{}{&pOut})
	w, h = pOut.X, pOut.Y
	return
}

// CIMGUI_API void igGetItemRectMin(ImVec2* pOut);
// IMGUI_API ImVec2        GetItemRectMin();
// get upper-left bounding rectangle of the last item (screen space)
// return ImVec2 need Destory
func GetItemRectMinVec2() *ImVec2 {
	pOut := NewImVec2Float(0, 0)
	giLib.Call(_func_igGetItemRectMin_, []interface{}{&pOut})
	return pOut
}

// CIMGUI_API void igGetItemRectMax(ImVec2* pOut);
// IMGUI_API ImVec2 GetItemRectMax();
// get lower-right bounding rectangle of the last item (screen space)
func GetItemRectMax() (w, h float32) {
	pOut := NewImVec2Float(0, 0)
	defer pOut.Destroy()
	giLib.Call(_func_igGetItemRectMax_, []interface{}{&pOut})
	w, h = pOut.X, pOut.Y
	return
}

// CIMGUI_API void igGetItemRectMax(ImVec2* pOut);
// IMGUI_API ImVec2 GetItemRectMax();
// get lower-right bounding rectangle of the last item (screen space)
// return ImVec2 need Destory
func GetItemRectMaxVec2() *ImVec2 {
	pOut := NewImVec2Float(0, 0)
	giLib.Call(_func_igGetItemRectMax_, []interface{}{&pOut})
	return pOut
}

// CIMGUI_API void igGetItemRectSize(ImVec2* pOut);
// IMGUI_API ImVec2 GetItemRectSize();
// get size of last item
func GetItemRectSize() (w, h float32) {
	pOut := NewImVec2Float(0, 0)
	defer pOut.Destroy()
	giLib.Call(_func_igGetItemRectSize_, []interface{}{&pOut})
	w, h = pOut.X, pOut.Y
	return
}

// CIMGUI_API void igGetItemRectSize(ImVec2* pOut);
// IMGUI_API ImVec2 GetItemRectSize();
// get size of last item
// return ImVec2 need Destory
func GetItemRectSizeVec2() *ImVec2 {
	pOut := NewImVec2Float(0, 0)
	giLib.Call(_func_igGetItemRectSize_, []interface{}{&pOut})
	return pOut
}

// CIMGUI_API ImGuiViewport* igGetMainViewport(void);
// Viewports
// - Currently represents the Platform Window created by the application which is hosting our Dear ImGui windows.
// - In 'docking' branch with multi-viewport enabled, we extend this concept to have multiple active viewports.
// - In the future we will extend this concept further to also represent Platform Monitor and support a "no main platform window" operation mode.

// IMGUI_API ImGuiViewport* GetMainViewport();
// return primary/default viewport. This can never be NULL.
func GetMainViewport() *ImGuiViewport {
	return (*ImGuiViewport)(giLib.Call(_func_igGetMainViewport_, nil).PtrFree())
}

// CIMGUI_API ImDrawList* igGetBackgroundDrawList_Nil(void);
// IMGUI_API ImDrawList*   GetBackgroundDrawList();
// get background draw list for the viewport associated to the current window. this draw list will be the first rendering one. Useful to quickly draw shapes/text behind dear imgui contents.
func GetBackgroundDrawList() *ImDrawList {
	return (*ImDrawList)(giLib.Call(_func_igGetBackgroundDrawList_Nil_, nil).PtrFree())
}

// CIMGUI_API ImDrawList* igGetForegroundDrawList_Nil(void);
// IMGUI_API ImDrawList* GetForegroundDrawList();
// get foreground draw list for the viewport associated to the current window. this draw list will be the last rendered one. Useful to quickly draw shapes/text over dear imgui contents.
func GetForegroundDrawList() *ImDrawList {
	return (*ImDrawList)(giLib.Call(_func_igGetForegroundDrawList_Nil_, nil).PtrFree())
}

// CIMGUI_API ImDrawList* igGetBackgroundDrawList_ViewportPtr(ImGuiViewport* viewport);
// IMGUI_API ImDrawList* GetBackgroundDrawList(ImGuiViewport* viewport);
// get background draw list for the given viewport. this draw list will be the first rendering one. Useful to quickly draw shapes/text behind dear imgui contents.
func GetBackgroundDrawListByViewport(viewport *ImGuiViewport) *ImDrawList {
	return (*ImDrawList)(giLib.Call(_func_igGetBackgroundDrawList_ViewportPtr_, []interface{}{&viewport}).PtrFree())

}

// CIMGUI_API ImDrawList* igGetForegroundDrawList_ViewportPtr(ImGuiViewport* viewport);
// IMGUI_API ImDrawList* GetForegroundDrawList(ImGuiViewport* viewport);
// get foreground draw list for the given viewport. this draw list will be the last rendered one. Useful to quickly draw shapes/text over dear imgui contents.
func GetForegroundDrawListByViewport(viewport *ImGuiViewport) *ImDrawList {
	return (*ImDrawList)(giLib.Call(_func_igGetForegroundDrawList_ViewportPtr_, []interface{}{&viewport}).PtrFree())

}

// CIMGUI_API bool igIsRectVisible_Nil(const ImVec2 size);
// IMGUI_API bool IsRectVisible(const ImVec2& size);
// test if rectangle (of given size, starting from cursor position) is visible / not clipped.
func IsRectVisibleVec2(size ImVec2) bool {
	return giLib.Call(_func_igIsRectVisible_Nil_, []interface{}{&size}).BoolFree()
}

// CIMGUI_API bool igIsRectVisible_Nil(const ImVec2 size);
// IMGUI_API bool IsRectVisible(const ImVec2& size);
// test if rectangle (of given size, starting from cursor position) is visible / not clipped.
func IsRectVisible(w, h float32) bool {
	size := NewImVec2Float(w, h)
	defer size.Destroy()
	return giLib.Call(_func_igIsRectVisible_Nil_, []interface{}{size}).BoolFree()
}

// CIMGUI_API bool igIsRectVisible_Vec2(const ImVec2 rect_min,const ImVec2 rect_max);
// IMGUI_API bool IsRectVisible(const ImVec2& rect_min, const ImVec2& rect_max);
// test if rectangle (in screen space) is visible / not clipped. to perform coarse clipping on user's side.
func IsRectVisibleByMinMaxVec2(rect_min, rect_max ImVec2) bool {
	return giLib.Call(_func_igIsRectVisible_Vec2_, []interface{}{&rect_min, &rect_max}).BoolFree()

}

// CIMGUI_API bool igIsRectVisible_Vec2(const ImVec2 rect_min,const ImVec2 rect_max);
// IMGUI_API bool IsRectVisible(const ImVec2& rect_min, const ImVec2& rect_max);
// test if rectangle (in screen space) is visible / not clipped. to perform coarse clipping on user's side.
func IsRectVisibleByMinMax(minX, minY, maxX, maxY float32) bool {
	rect_min, rect_max := NewImVec2Float(minX, minY), NewImVec2Float(maxX, maxY)
	defer rect_min.Destroy()
	defer rect_max.Destroy()
	return giLib.Call(_func_igIsRectVisible_Vec2_, []interface{}{rect_min, rect_max}).BoolFree()

}

// CIMGUI_API double igGetTime(void);
// IMGUI_API double GetTime();
// get global imgui time. incremented by io.DeltaTime every frame.
func GetTime() float64 {
	return giLib.Call(_func_igGetTime_, nil).F64Free()
}

// CIMGUI_API int igGetFrameCount(void);
// IMGUI_API int GetFrameCount();
// get global imgui frame count. incremented by 1 every frame.
func GetFrameCount() int32 {
	return giLib.Call(_func_igGetFrameCount_, nil).I32Free()
}

// CIMGUI_API ImDrawListSharedData* igGetDrawListSharedData(void);
// IMGUI_API ImDrawListSharedData* GetDrawListSharedData();
// you may use this when creating your own ImDrawList instances.
func GetDrawListSharedData() *ImDrawListSharedData {
	return (*ImDrawListSharedData)(giLib.Call(_func_igGetDrawListSharedData_, nil).PtrFree())
}

// CIMGUI_API const char* igGetStyleColorName(ImGuiCol idx);
// IMGUI_API const char*   GetStyleColorName(ImGuiCol idx);
// get a string corresponding to the enum value (for display, saving, etc.).
func GetStyleColorName(idx ImGuiCol) string {
	return giLib.Call(_func_igGetStyleColorName_, []interface{}{&idx}).StrFree()
}

// CIMGUI_API void igSetStateStorage(ImGuiStorage* storage);
// IMGUI_API void SetStateStorage(ImGuiStorage* storage);
// replace current window storage with our own (if you want to manipulate it yourself, typically clear subsection of it)
func SetStateStorage(storage *ImGuiStorage) {
	giLib.Call(_func_igSetStateStorage_, []interface{}{&storage})
}

// CIMGUI_API ImGuiStorage* igGetStateStorage(void);
// IMGUI_API ImGuiStorage* GetStateStorage();
func GetStateStorage() *ImGuiStorage {
	return (*ImGuiStorage)(giLib.Call(_func_igGetStateStorage_, nil).PtrFree())
}

// CIMGUI_API bool igBeginChildFrame(ImGuiID id,const ImVec2 size,ImGuiWindowFlags flags);
// IMGUI_API bool BeginChildFrame(ImGuiID id, const ImVec2& size, ImGuiWindowFlags flags = 0);
// helper to create a child window / scrolling region that looks like a normal widget frame
func BeginChildFrameVec2(id ImGuiID, size ImVec2, flags ImGuiWindowFlags) bool {
	return giLib.Call(_func_igBeginChildFrame_, []interface{}{&id, &size, &flags}).BoolFree()

}

// CIMGUI_API bool igBeginChildFrame(ImGuiID id,const ImVec2 size,ImGuiWindowFlags flags);
// IMGUI_API bool BeginChildFrame(ImGuiID id, const ImVec2& size, ImGuiWindowFlags flags = 0);
// helper to create a child window / scrolling region that looks like a normal widget frame
func BeginChildFrame(id ImGuiID, w, h float32, flags ImGuiWindowFlags) bool {
	size := NewImVec2Float(w, h)
	defer size.Destroy()
	return giLib.Call(_func_igBeginChildFrame_, []interface{}{&id, size, &flags}).BoolFree()

}

// CIMGUI_API void igEndChildFrame(void);
// IMGUI_API void EndChildFrame();
// always call EndChildFrame() regardless of BeginChildFrame() return values (which indicates a collapsed/clipped window)
func EndChildFrame() { giLib.Call(_func_igEndChildFrame_, nil) }

// Text Utilities
// CIMGUI_API void igCalcTextSize(ImVec2* pOut,const char* text,const char* text_end,bool hide_text_after_double_hash,float wrap_width);
// IMGUI_API ImVec2 CalcTextSize(const char* text, const char* text_end = NULL, bool hide_text_after_double_hash = false, float wrap_width = -1.0f);
func CalcTextSize(text, text_end string, hide_text_after_double_hash bool, wrap_width float32) (w, h float32) {
	pOut, txt, txt_end, b := NewImVec2Float(0, 0), c.CStr(text), c.CStr(text_end), c.CBool(hide_text_after_double_hash)
	defer pOut.Destroy()
	defer usf.Free(txt)
	defer usf.Free(txt_end)
	giLib.Call(_func_igCalcTextSize_, []interface{}{&pOut, &txt, &txt_end, &b, &wrap_width})
	w, h = pOut.X, pOut.Y
	return
}

// Text Utilities
// CIMGUI_API void igCalcTextSize(ImVec2* pOut,const char* text,const char* text_end,bool hide_text_after_double_hash,float wrap_width);
// IMGUI_API ImVec2 CalcTextSize(const char* text, const char* text_end = NULL, bool hide_text_after_double_hash = false, float wrap_width = -1.0f);
func CalcTextSizeDefault(text string) (w, h float32) {
	pOut, txt := NewImVec2Float(0, 0), c.CStr(text)
	defer pOut.Destroy()
	defer usf.Free(txt)
	giLib.Call(_func_igCalcTextSize_, []interface{}{&pOut, &txt, _ptr_zero, _bool_false, _f32_minus_one})
	w, h = pOut.X, pOut.Y
	return
}

// Text Utilities
// CIMGUI_API void igCalcTextSize(ImVec2* pOut,const char* text,const char* text_end,bool hide_text_after_double_hash,float wrap_width);
// IMGUI_API ImVec2 CalcTextSize(const char* text, const char* text_end = NULL, bool hide_text_after_double_hash = false, float wrap_width = -1.0f);
func CalcTextSizeVec2(text, text_end string, hide_text_after_double_hash bool, wrap_width float32) *ImVec2 {
	pOut, txt, txt_end, b := NewImVec2Float(0, 0), c.CStr(text), c.CStr(text_end), c.CBool(hide_text_after_double_hash)
	defer usf.Free(txt)
	defer usf.Free(txt_end)
	giLib.Call(_func_igCalcTextSize_, []interface{}{&pOut, &txt, &txt_end, &b, &wrap_width})
	return pOut
}

// CIMGUI_API void igColorConvertU32ToFloat4(ImVec4* pOut,ImU32 in);
// IMGUI_API ImVec4 ColorConvertU32ToFloat4(ImU32 in);
func ColorConvertU32ToFloat4(in ImU32) (x, y, z, w float32) {
	pOut := NewImVec4Float(0, 0, 0, 0)
	defer pOut.Destroy()
	giLib.Call(_func_igColorConvertU32ToFloat4_, []interface{}{&pOut, &in})
	x, y, z, w = pOut.X, pOut.Y, pOut.Z, pOut.W
	return
}

// CIMGUI_API void igColorConvertU32ToFloat4(ImVec4* pOut,ImU32 in);
// IMGUI_API ImVec4 ColorConvertU32ToFloat4(ImU32 in);
func ColorConvertU32ToImVec4(in uint32) *ImVec4 {
	pOut := NewImVec4Float(0, 0, 0, 0)
	giLib.Call(_func_igColorConvertU32ToFloat4_, []interface{}{&pOut, &in})
	return pOut
}

// CIMGUI_API ImU32 igColorConvertFloat4ToU32(const ImVec4 in);
// IMGUI_API ImU32 ColorConvertFloat4ToU32(const ImVec4& in);
func ColorConvertImVec4ToU32(in ImVec4) uint32 {
	return giLib.Call(_func_igColorConvertFloat4ToU32_, []interface{}{&in}).U32Free()

}

// CIMGUI_API ImU32 igColorConvertFloat4ToU32(const ImVec4 in);
// IMGUI_API ImU32 ColorConvertFloat4ToU32(const ImVec4& in);
func ColorConvertFloat4ToU32(x, y, z, w float32) uint32 {
	in := NewImVec4Float(x, y, z, w)
	return giLib.Call(_func_igColorConvertFloat4ToU32_, []interface{}{in}).U32Free()

}

// CIMGUI_API void igColorConvertRGBtoHSV(float r,float g,float b,float* out_h,float* out_s,float* out_v);
// IMGUI_API void ColorConvertRGBtoHSV(float r, float g, float b, float& out_h, float& out_s, float& out_v);
func ColorConvertRGBtoHSV(r, g, b float32) (h, s, v float32) {
	_h, _s, _v := &h, &s, &v
	giLib.Call(_func_igColorConvertRGBtoHSV_, []interface{}{&r, &g, &b, &_h, &_s, &_v})

	return
}

// CIMGUI_API void igColorConvertHSVtoRGB(float h,float s,float v,float* out_r,float* out_g,float* out_b);
// IMGUI_API void ColorConvertHSVtoRGB(float h, float s, float v, float& out_r, float& out_g, float& out_b);
func ColorConvertHSVtoRGB(h, s, v float32) (r, g, b float32) {
	_r, _g, _b := &r, &g, &b
	giLib.Call(_func_igColorConvertHSVtoRGB_, []interface{}{&h, &s, &v, &_r, &_g, &_b})

	return
}

// Inputs Utilities: Keyboard/Mouse/Gamepad
// - the ImGuiKey enum contains all possible keyboard, mouse and gamepad inputs (e.g. ImGuiKey_A, ImGuiKey_MouseLeft, ImGuiKey_GamepadDpadUp...).
// - before v1.87, we used ImGuiKey to carry native/user indices as defined by each backends. About use of those legacy ImGuiKey values:
//  - without IMGUI_DISABLE_OBSOLETE_KEYIO (legacy support): you can still use your legacy native/user indices (< 512) according to how your backend/engine stored them in io.KeysDown[], but need to cast them to ImGuiKey.
//  - with    IMGUI_DISABLE_OBSOLETE_KEYIO (this is the way forward): any use of ImGuiKey will assert with key < 512. GetKeyIndex() is pass-through and therefore deprecated (gone if IMGUI_DISABLE_OBSOLETE_KEYIO is defined).

// CIMGUI_API bool igIsKeyDown_Nil(ImGuiKey key);
// IMGUI_API bool IsKeyDown(ImGuiKey key);
// is key being held.
func IsKeyDown(key ImGuiKey) bool {
	return giLib.Call(_func_igIsKeyDown_Nil_, []interface{}{&key}).BoolFree()
}

// CIMGUI_API bool igIsKeyPressed_Bool(ImGuiKey key,bool repeat);
// IMGUI_API bool IsKeyPressed(ImGuiKey key, bool repeat = true);
// was key pressed (went from !Down to Down)? if repeat=true, uses io.KeyRepeatDelay / KeyRepeatRate
func IsKeyPressed(key ImGuiKey, repeat bool) bool {
	r := c.CBool(repeat)
	return giLib.Call(_func_igIsKeyPressed_Bool_, []interface{}{&key, &r}).BoolFree()

}

// CIMGUI_API bool igIsKeyReleased_Nil(ImGuiKey key);
// IMGUI_API bool IsKeyReleased(ImGuiKey key);
// was key released (went from Down to !Down)?
func IsKeyReleased(key ImGuiKey) bool {
	return giLib.Call(_func_igIsKeyReleased_Nil_, []interface{}{&key}).BoolFree()
}

// CIMGUI_API int igGetKeyPressedAmount(ImGuiKey key,float repeat_delay,float rate);
// IMGUI_API int GetKeyPressedAmount(ImGuiKey key, float repeat_delay, float rate);
// uses provided repeat rate/delay. return a count, most often 0 or 1 but might be >1 if RepeatRate is small enough that DeltaTime > RepeatRate
func GetKeyPressedAmount(key ImGuiKey, repeat_delay, rate float32) int32 {
	return giLib.Call(_func_igGetKeyPressedAmount_, []interface{}{&key, &repeat_delay, &rate}).I32Free()

}

// CIMGUI_API const char* igGetKeyName(ImGuiKey key);
// IMGUI_API const char* GetKeyName(ImGuiKey key);
// [DEBUG] returns English name of the key. Those names a provided for debugging purpose and are not meant to be saved persistently not compared.
func GetKeyName(key ImGuiKey) string {
	return giLib.Call(_func_igGetKeyName_, []interface{}{&key}).StrFree()
}

// CIMGUI_API void igSetNextFrameWantCaptureKeyboard(bool want_capture_keyboard);
// IMGUI_API void SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard);
// Override io.WantCaptureKeyboard flag next frame (said flag is left for your application to handle, typically when true it instructs your app to ignore inputs). e.g. force capture keyboard when your widget is being hovered. This is equivalent to setting "io.WantCaptureKeyboard = want_capture_keyboard"; after the next NewFrame() call.
func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
	k := c.CBool(want_capture_keyboard)
	giLib.Call(_func_igSetNextFrameWantCaptureKeyboard_, []interface{}{&k})

}

// Inputs Utilities: Mouse specific
// - To refer to a mouse button, you may use named enums in your code e.g. ImGuiMouseButton_Left, ImGuiMouseButton_Right.
// - You can also use regular integer: it is forever guaranteed that 0=Left, 1=Right, 2=Middle.
// - Dragging operations are only reported after mouse has moved a certain distance away from the initial clicking position (see 'lock_threshold' and 'io.MouseDraggingThreshold')

// CIMGUI_API bool igIsMouseDown_Nil(ImGuiMouseButton button);
// IMGUI_API bool IsMouseDown(ImGuiMouseButton button);
// is mouse button held?
func IsMouseDown(button ImGuiMouseButton) bool {
	return giLib.Call(_func_igIsMouseDown_Nil_, []interface{}{&button}).BoolFree()
}

// CIMGUI_API bool igIsMouseClicked_Bool(ImGuiMouseButton button,bool repeat);
// IMGUI_API bool IsMouseClicked(ImGuiMouseButton button, bool repeat = false);
// did mouse button clicked? (went from !Down to Down). Same as GetMouseClickedCount() == 1.
func IsMouseClicked(button ImGuiMouseButton, repeat bool) bool {
	r := c.CBool(repeat)
	return giLib.Call(_func_igIsMouseClicked_Bool_, []interface{}{&button, &r}).BoolFree()

}

// CIMGUI_API bool igIsMouseReleased_Nil(ImGuiMouseButton button);
// IMGUI_API bool IsMouseReleased(ImGuiMouseButton button);
// did mouse button released? (went from Down to !Down)
func IsMouseReleased(button ImGuiMouseButton) bool {
	return giLib.Call(_func_igIsMouseReleased_Nil_, []interface{}{&button}).BoolFree()
}

// CIMGUI_API bool igIsMouseDoubleClicked(ImGuiMouseButton button);
// IMGUI_API bool IsMouseDoubleClicked(ImGuiMouseButton button);
// did mouse button double-clicked? Same as GetMouseClickedCount() == 2. (note that a double-click will also report IsMouseClicked() == true)
func IsMouseDoubleClicked(button ImGuiMouseButton) bool {
	return giLib.Call(_func_igIsMouseDoubleClicked_, []interface{}{&button}).BoolFree()
}

// CIMGUI_API int igGetMouseClickedCount(ImGuiMouseButton button);
// IMGUI_API int GetMouseClickedCount(ImGuiMouseButton button);
// return the number of successive mouse-clicks at the time where a click happen (otherwise 0).
func GetMouseClickedCount(button ImGuiMouseButton) int32 {
	return giLib.Call(_func_igGetMouseClickedCount_, []interface{}{&button}).I32Free()
}

// CIMGUI_API bool igIsMouseHoveringRect(const ImVec2 r_min,const ImVec2 r_max,bool clip);
// IMGUI_API bool IsMouseHoveringRect(const ImVec2& r_min, const ImVec2& r_max, bool clip = true);
// is mouse hovering given bounding rect (in screen space). clipped by current clipping settings, but disregarding of other consideration of focus/window ordering/popup-block.
func IsMouseHoveringRectVec2(r_min, r_max ImVec2, clip bool) bool {
	_clip := c.CBool(clip)
	return giLib.Call(_func_igIsMouseHoveringRect_, []interface{}{&r_min, &r_max, &_clip}).BoolFree()

}

// CIMGUI_API bool igIsMouseHoveringRect(const ImVec2 r_min,const ImVec2 r_max,bool clip);
// IMGUI_API bool IsMouseHoveringRect(const ImVec2& r_min, const ImVec2& r_max, bool clip = true);
// is mouse hovering given bounding rect (in screen space). clipped by current clipping settings, but disregarding of other consideration of focus/window ordering/popup-block.
func IsMouseHoveringRect(minX, minY, maxX, maxY float32, clip bool) bool {
	r_min, r_max, _clip := NewImVec2Float(minX, minY), NewImVec2Float(maxX, maxY), c.CBool(clip)
	defer r_min.Destroy()
	defer r_max.Destroy()
	return giLib.Call(_func_igIsMouseHoveringRect_, []interface{}{r_min, r_max, &_clip}).BoolFree()

}

// CIMGUI_API bool igIsMousePosValid(const ImVec2* mouse_pos);
// IMGUI_API bool IsMousePosValid(const ImVec2* mouse_pos = NULL);
// by convention we use (-FLT_MAX,-FLT_MAX) to denote that there is no mouse available
func IsMousePosValidVec2(mouse_pos *ImVec2) bool {
	return giLib.Call(_func_igIsMousePosValid_, []interface{}{&mouse_pos}).BoolFree()

}

// CIMGUI_API bool igIsMousePosValid(const ImVec2* mouse_pos);
// IMGUI_API bool IsMousePosValid(const ImVec2* mouse_pos = NULL);
// by convention we use (-FLT_MAX,-FLT_MAX) to denote that there is no mouse available
func IsMousePosValid(x, y float32) bool {
	mouse_pos := NewImVec2Float(x, y)
	defer mouse_pos.Destroy()
	return giLib.Call(_func_igIsMousePosValid_, []interface{}{&mouse_pos}).BoolFree()

}

// CIMGUI_API bool igIsAnyMouseDown(void);
// IMGUI_API bool IsAnyMouseDown();
// [WILL OBSOLETE] is any mouse button held? This was designed for backends, but prefer having backend maintain a mask of held mouse buttons, because upcoming input queue system will make this invalid.
func IsAnyMouseDown() bool {
	return giLib.Call(_func_igIsAnyMouseDown_, nil).BoolFree()
}

// CIMGUI_API void igGetMousePos(ImVec2* pOut);
// IMGUI_API ImVec2 GetMousePos();
// shortcut to ImGui::GetIO().MousePos provided by user, to be consistent with other calls
func GetMousePos() (x, y float32) {
	pOut := NewImVec2Float(0, 0)
	defer pOut.Destroy()
	giLib.Call(_func_igGetMousePos_, []interface{}{&pOut})
	return pOut.X, pOut.Y
}

// CIMGUI_API void igGetMousePos(ImVec2* pOut);
// IMGUI_API ImVec2 GetMousePos();
// shortcut to ImGui::GetIO().MousePos provided by user, to be consistent with other calls
func GetMousePosVec2() *ImVec2 {
	pOut := NewImVec2Float(0, 0)
	giLib.Call(_func_igGetMousePos_, []interface{}{&pOut})
	return pOut
}

// CIMGUI_API void igGetMousePosOnOpeningCurrentPopup(ImVec2* pOut);
// IMGUI_API ImVec2 GetMousePosOnOpeningCurrentPopup();
// retrieve mouse position at the time of opening popup we have BeginPopup() into (helper to avoid user backing that value themselves)
func GetMousePosOnOpeningCurrentPopupVec2() *ImVec2 {
	pOut := NewImVec2Float(0, 0)
	giLib.Call(_func_igGetMousePosOnOpeningCurrentPopup_, []interface{}{&pOut})
	return pOut
}

// CIMGUI_API void igGetMousePosOnOpeningCurrentPopup(ImVec2* pOut);
// IMGUI_API ImVec2 GetMousePosOnOpeningCurrentPopup();
// retrieve mouse position at the time of opening popup we have BeginPopup() into (helper to avoid user backing that value themselves)
func GetMousePosOnOpeningCurrentPopup() (x, y float32) {
	pOut := NewImVec2Float(0, 0)
	defer pOut.Destroy()
	giLib.Call(_func_igGetMousePosOnOpeningCurrentPopup_, []interface{}{&pOut})
	return pOut.X, pOut.Y
}

// CIMGUI_API bool igIsMouseDragging(ImGuiMouseButton button,float lock_threshold);
// IMGUI_API bool IsMouseDragging(ImGuiMouseButton button, float lock_threshold = -1.0f);
// is mouse dragging? (if lock_threshold < -1.0f, uses io.MouseDraggingThreshold)
func IsMouseDragging(button ImGuiMouseButton, lock_threshold float32) bool {
	return giLib.Call(_func_igIsMouseDragging_, []interface{}{&button, &lock_threshold}).BoolFree()

}

// CIMGUI_API void igGetMouseDragDelta(ImVec2* pOut,ImGuiMouseButton button,float lock_threshold);
// IMGUI_API ImVec2 GetMouseDragDelta(ImGuiMouseButton button = 0, float lock_threshold = -1.0f);
// return the delta from the initial clicking position while the mouse button is pressed or was just released. This is locked and return 0.0f until the mouse moves past a distance threshold at least once (if lock_threshold < -1.0f, uses io.MouseDraggingThreshold)
func GetMouseDragDelta(button ImGuiMouseButton, lock_threshold float32) (dx, dy float32) {
	pOut := NewImVec2Float(0, 0)
	defer pOut.Destroy()
	giLib.Call(_func_igGetMouseDragDelta_, []interface{}{&pOut, &button, &lock_threshold})

	dx, dy = pOut.X, pOut.Y
	return
}

// CIMGUI_API void igGetMouseDragDelta(ImVec2* pOut,ImGuiMouseButton button,float lock_threshold);
// IMGUI_API ImVec2 GetMouseDragDelta(ImGuiMouseButton button = 0, float lock_threshold = -1.0f);
// return the delta from the initial clicking position while the mouse button is pressed or was just released. This is locked and return 0.0f until the mouse moves past a distance threshold at least once (if lock_threshold < -1.0f, uses io.MouseDraggingThreshold)
func GetMouseDragDeltaVec2(button ImGuiMouseButton, lock_threshold float32) *ImVec2 {
	pOut := NewImVec2Float(0, 0)
	giLib.Call(_func_igGetMouseDragDelta_, []interface{}{&pOut, &button, &lock_threshold})
	return pOut
}

// CIMGUI_API void igResetMouseDragDelta(ImGuiMouseButton button);
// IMGUI_API void ResetMouseDragDelta(ImGuiMouseButton button = 0);
func ResetMouseDragDelta(button ImGuiMouseButton) {
	giLib.Call(_func_igResetMouseDragDelta_, []interface{}{&button})
}

// CIMGUI_API ImGuiMouseCursor igGetMouseCursor(void);
// IMGUI_API ImGuiMouseCursor GetMouseCursor();
// get desired mouse cursor shape. Important: reset in ImGui::NewFrame(), this is updated during the frame. valid before Render(). If you use software rendering by setting io.MouseDrawCursor ImGui will render those for you
func GetMouseCursor() ImGuiMouseCursor {
	return ImGuiMouseCursor(giLib.Call(_func_igGetMouseCursor_, nil).I32Free())
}

// CIMGUI_API void igSetMouseCursor(ImGuiMouseCursor cursor_type);
// IMGUI_API void SetMouseCursor(ImGuiMouseCursor cursor_type);
// set desired mouse cursor shape
func SetMouseCursor(cursor_type ImGuiMouseCursor) {
	giLib.Call(_func_igSetMouseCursor_, []interface{}{&cursor_type})
}

// CIMGUI_API void igSetNextFrameWantCaptureMouse(bool want_capture_mouse);
// IMGUI_API void SetNextFrameWantCaptureMouse(bool want_capture_mouse);
// Override io.WantCaptureMouse flag next frame (said flag is left for your application to handle, typical when true it instucts your app to ignore inputs). This is equivalent to setting "io.WantCaptureMouse = want_capture_mouse;" after the next NewFrame() call.
func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
	b := c.CBool(want_capture_mouse)
	giLib.Call(_func_igSetNextFrameWantCaptureMouse_, []interface{}{&b})

}

// Clipboard Utilities
// - Also see the LogToClipboard() function to capture GUI into clipboard, or easily output text data to the clipboard.

// CIMGUI_API const char* igGetClipboardText(void);
// IMGUI_API const char* GetClipboardText();
func GetClipboardText() string {
	return giLib.Call(_func_igGetClipboardText_, nil).StrFree()
}

// CIMGUI_API void igSetClipboardText(const char* text);
// IMGUI_API void SetClipboardText(const char* text);
func SetClipboardText(text string) {
	txt := c.CStr(text)
	defer usf.Free(txt)
	giLib.Call(_func_igSetClipboardText_, []interface{}{&txt})
}

// Settings/.Ini Utilities
// - The disk functions are automatically called if io.IniFilename != NULL (default is "imgui.ini").
// - Set io.IniFilename to NULL to load/save manually. Read io.WantSaveIniSettings description about handling .ini saving manually.
// - Important: default value "imgui.ini" is relative to current working dir! Most apps will want to lock this to an absolute path (e.g. same path as executables).

// CIMGUI_API void igLoadIniSettingsFromDisk(const char* ini_filename);
// IMGUI_API void LoadIniSettingsFromDisk(const char* ini_filename);
// call after CreateContext() and before the first call to NewFrame(). NewFrame() automatically calls LoadIniSettingsFromDisk(io.IniFilename).
func LoadIniSettingsFromDisk(ini_filename string) {
	filename := c.CStr(ini_filename)
	defer usf.Free(filename)
	giLib.Call(_func_igLoadIniSettingsFromDisk_, []interface{}{&filename})
}

// CIMGUI_API void igLoadIniSettingsFromMemory(const char* ini_data,size_t ini_size);
// IMGUI_API void LoadIniSettingsFromMemory(const char* ini_data, size_t ini_size=0);
// call after CreateContext() and before the first call to NewFrame() to provide .ini data from your own data source.
func LoadIniSettingsFromMemory(ini_data []byte) {
	data := &ini_data[0]
	size := uint64(len(ini_data))
	giLib.Call(_func_igLoadIniSettingsFromMemory_, []interface{}{&data, &size})
}

// CIMGUI_API void igSaveIniSettingsToDisk(const char* ini_filename);
// IMGUI_API void SaveIniSettingsToDisk(const char* ini_filename);
// this is automatically called (if io.IniFilename is not empty) a few seconds after any modification that should be reflected in the .ini file (and also by DestroyContext).
func SaveIniSettingsToDisk(ini_filename string) {
	filename := c.CStr(ini_filename)
	defer usf.Free(filename)
	giLib.Call(_func_igSaveIniSettingsToDisk_, []interface{}{&filename})
}

// CIMGUI_API const char* igSaveIniSettingsToMemory(size_t* out_ini_size);
// IMGUI_API const char* SaveIniSettingsToMemory(size_t* out_ini_size = NULL);
// return a zero-terminated string with the .ini data which you can save by your own mean. call when io.WantSaveIniSettings is set, then save data by your own mean and clear io.WantSaveIniSettings.
func SaveIniSettingsToMemory() string {
	out_ini_size := new(uint64)
	p := giLib.Call(_func_igSaveIniSettingsToMemory_, []interface{}{&out_ini_size}).PtrFree()

	s := *(*[]byte)(usf.Slice(p, *out_ini_size))
	return string(s)
}

// Debug Utilities

// CIMGUI_API void igDebugTextEncoding(const char* text);
// IMGUI_API void DebugTextEncoding(const char* text);
func DebugTextEncoding(text string) {
	txt := c.CStr(text)
	defer usf.Free(txt)
	giLib.Call(_func_igDebugTextEncoding_, []interface{}{&txt})
}

// CIMGUI_API bool igDebugCheckVersionAndDataLayout(const char* version_str,size_t sz_io,size_t sz_style,size_t sz_vec2,size_t sz_vec4,size_t sz_drawvert,size_t sz_drawidx);
// IMGUI_API bool DebugCheckVersionAndDataLayout(const char* version_str, size_t sz_io, size_t sz_style, size_t sz_vec2, size_t sz_vec4, size_t sz_drawvert, size_t sz_drawidx);
// This is called by IMGUI_CHECKVERSION() macro.
func DebugCheckVersionAndDataLayout(version_str string, sz_io, sz_style, sz_vec2, sz_vec4, sz_drawvert, sz_drawidx uint64) bool {
	ver_str := c.CStr(version_str)
	defer usf.Free(ver_str)
	return giLib.Call(_func_igDebugCheckVersionAndDataLayout_,

		[]interface{}{&ver_str, &sz_io, &sz_style, &sz_vec2, &sz_vec4, &sz_drawvert, &sz_drawidx}).BoolFree()
}

// // Memory Allocators
// // - Those functions are not reliant on the current context.
// // - DLL users: heaps and globals are not shared across DLL boundaries! You will need to call SetCurrentContext() + SetAllocatorFunctions()
// //   for each static/DLL boundary you are calling from. Read "Context and Memory Allocators" section of imgui.cpp for more details.
// IMGUI_API void          SetAllocatorFunctions(ImGuiMemAllocFunc alloc_func, ImGuiMemFreeFunc free_func, void* user_data = NULL);
// IMGUI_API void          GetAllocatorFunctions(ImGuiMemAllocFunc* p_alloc_func, ImGuiMemFreeFunc* p_free_func, void** p_user_data);
// IMGUI_API void*         MemAlloc(size_t size);
// IMGUI_API void          MemFree(void* ptr);
// CIMGUI_API void igSetAllocatorFunctions(ImGuiMemAllocFunc alloc_func,ImGuiMemFreeFunc free_func,void* user_data);
// CIMGUI_API void igGetAllocatorFunctions(ImGuiMemAllocFunc* p_alloc_func,ImGuiMemFreeFunc* p_free_func,void** p_user_data);
// CIMGUI_API void* igMemAlloc(size_t size);
// CIMGUI_API void igMemFree(void* ptr);

// (Optional) Platform/OS interface for multi-viewport support
// Read comments around the ImGuiPlatformIO structure for more details.
// Note: You may use GetWindowViewport() to get the current viewport of the current window.

// CIMGUI_API ImGuiPlatformIO* igGetPlatformIO(void);
// IMGUI_API ImGuiPlatformIO&  GetPlatformIO();
// platform/renderer functions, for backend to setup + viewports list.
func GetPlatformIO() *ImGuiPlatformIO {
	return (*ImGuiPlatformIO)(giLib.Call(_func_igGetPlatformIO_, nil).PtrFree())
}

// CIMGUI_API void igUpdatePlatformWindows(void);
// IMGUI_API void UpdatePlatformWindows();
// call in main loop. will call CreateWindow/ResizeWindow/etc. platform functions for each secondary viewport, and DestroyWindow for each inactive viewport.
func UpdatePlatformWindows() { giLib.Call(_func_igUpdatePlatformWindows_, nil) }

// CIMGUI_API void igRenderPlatformWindowsDefault(void* platform_render_arg,void* renderer_render_arg);
// IMGUI_API void RenderPlatformWindowsDefault(void* platform_render_arg = NULL, void* renderer_render_arg = NULL);
// call in main loop. will call RenderWindow/SwapBuffers platform functions for each secondary viewport which doesn't have the ImGuiViewportFlags_Minimized flag set. May be reimplemented by user for custom rendering needs.
func RenderPlatformWindowsDefault(platform_render_arg, renderer_render_arg unsafe.Pointer) {
	giLib.Call(_func_igRenderPlatformWindowsDefault_, []interface{}{&platform_render_arg, &renderer_render_arg})

}

// CIMGUI_API void igDestroyPlatformWindows(void);
// IMGUI_API void DestroyPlatformWindows();
// call DestroyWindow platform functions for all viewports. call from backend Shutdown() if you need to close platform windows before imgui shutdown. otherwise will be called by DestroyContext().
func DestroyPlatformWindows() { giLib.Call(_func_igDestroyPlatformWindows_, nil) }

// CIMGUI_API ImGuiViewport* igFindViewportByID(ImGuiID id);
// IMGUI_API ImGuiViewport* FindViewportByID(ImGuiID id);
// this is a helper for backends.
func FindViewportByID(id ImGuiID) *ImGuiViewport {
	return (*ImGuiViewport)(giLib.Call(_func_igFindViewportByID_, []interface{}{&id}).PtrFree())

}

// CIMGUI_API ImGuiViewport* igFindViewportByPlatformHandle(void* platform_handle);
// IMGUI_API ImGuiViewport* FindViewportByPlatformHandle(void* platform_handle);
// this is a helper for backends. the type platform_handle is decided by the backend (e.g. HWND, MyWindow*, GLFWwindow* etc.)
func FindViewportByPlatformHandle(platform_handle unsafe.Pointer) *ImGuiViewport {
	return (*ImGuiViewport)(giLib.Call(_func_igFindViewportByPlatformHandle_, []interface{}{&platform_handle}).PtrFree())

}

//
//ImGuiStyle
//

// CIMGUI_API ImGuiStyle* ImGuiStyle_ImGuiStyle(void);
func NewImGuiStyle() *ImGuiStyle {
	return (*ImGuiStyle)(giLib.Call(_func_ImGuiStyle_ImGuiStyle_, nil).PtrFree())
}

// CIMGUI_API void ImGuiStyle_destroy(ImGuiStyle* self);
func (style *ImGuiStyle) Destroy() {
	giLib.Call(_func_ImGuiStyle_destroy_, []interface{}{&style})
}

// CIMGUI_API void ImGuiStyle_ScaleAllSizes(ImGuiStyle* self,float scale_factor);
// IMGUI_API void ScaleAllSizes(float scale_factor);
func (style *ImGuiStyle) ScaleAllSizes(scale_factor float32) {
	giLib.Call(_func_ImGuiStyle_ScaleAllSizes_, []interface{}{&style, &scale_factor})
}

//
//ImGuiIO
//

// CIMGUI_API void ImGuiIO_AddKeyEvent(ImGuiIO* self,ImGuiKey key,bool down);
// IMGUI_API void  AddKeyEvent(ImGuiKey key, bool down);
// Queue a new key down/up event. Key should be "translated" (as in, generally ImGuiKey_A matches the key end-user would use to emit an 'A' character)
func (io *ImGuiIO) AddKeyEvent(key ImGuiKey, down bool) {
	d := c.CBool(down)
	giLib.Call(_func_ImGuiIO_AddKeyEvent_, []interface{}{&io, &key, &d})
}

// CIMGUI_API void ImGuiIO_AddKeyAnalogEvent(ImGuiIO* self,ImGuiKey key,bool down,float v);
// IMGUI_API void  AddKeyAnalogEvent(ImGuiKey key, bool down, float v);
// Queue a new key down/up event for analog values (e.g. ImGuiKey_Gamepad_ values). Dead-zones should be handled by the backend.
func (io *ImGuiIO) AddKeyAnalogEvent(key ImGuiKey, down bool, v float32) {
	d := c.CBool(down)
	giLib.Call(_func_ImGuiIO_AddKeyAnalogEvent_, []interface{}{&io, &key, &d, &v})
}

// CIMGUI_API void ImGuiIO_AddMousePosEvent(ImGuiIO* self,float x,float y);
// IMGUI_API void  AddMousePosEvent(float x, float y);
// Queue a mouse position update. Use -FLT_MAX,-FLT_MAX to signify no mouse (e.g. app not focused and not hovered)
func (io *ImGuiIO) AddMousePosEvent(x, y float32) {
	giLib.Call(_func_ImGuiIO_AddMousePosEvent_, []interface{}{&io, &x, &y})
}

// CIMGUI_API void ImGuiIO_AddMouseButtonEvent(ImGuiIO* self,int button,bool down);
// IMGUI_API void  AddMouseButtonEvent(int button, bool down);
// Queue a mouse button change
func (io *ImGuiIO) AddMouseButtonEvent(button int32, down bool) {
	d := c.CBool(down)
	giLib.Call(_func_ImGuiIO_AddMouseButtonEvent_, []interface{}{&io, &button, &d})
}

// CIMGUI_API void ImGuiIO_AddMouseWheelEvent(ImGuiIO* self,float wheel_x,float wheel_y);
// IMGUI_API void  AddMouseWheelEvent(float wheel_x, float wheel_y);
// Queue a mouse wheel update. wheel_y<0: scroll down, wheel_y>0: scroll up, wheel_x<0: scroll right, wheel_x>0: scroll left.
func (io *ImGuiIO) AddMouseWheelEvent(wheel_x, wheel_y float32) {
	giLib.Call(_func_ImGuiIO_AddMouseWheelEvent_, []interface{}{&io, &wheel_x, &wheel_y})
}

// CIMGUI_API void ImGuiIO_AddMouseSourceEvent(ImGuiIO* self,ImGuiMouseSource source);
// IMGUI_API void  AddMouseSourceEvent(ImGuiMouseSource source);
// Queue a mouse source change (Mouse/TouchScreen/Pen)
func (io *ImGuiIO) AddMouseSourceEvent(source ImGuiMouseSource) {
	giLib.Call(_func_ImGuiIO_AddMouseSourceEvent_, []interface{}{&io, &source})
}

// CIMGUI_API void ImGuiIO_AddMouseViewportEvent(ImGuiIO* self,ImGuiID id);
// IMGUI_API void  AddMouseViewportEvent(ImGuiID id);
// Queue a mouse hovered viewport. Requires backend to set ImGuiBackendFlags_HasMouseHoveredViewport to call this (for multi-viewport support).
func (io *ImGuiIO) AddMouseViewportEvent(id ImGuiID) {
	giLib.Call(_func_ImGuiIO_AddMouseViewportEvent_, []interface{}{&io, &id})
}

// CIMGUI_API void ImGuiIO_AddFocusEvent(ImGuiIO* self,bool focused);
// IMGUI_API void  AddFocusEvent(bool focused);
// Queue a gain/loss of focus for the application (generally based on OS/platform focus of your window)
func (io *ImGuiIO) AddFocusEvent(focused bool) {
	f := c.CBool(focused)
	giLib.Call(_func_ImGuiIO_AddFocusEvent_, []interface{}{&io, &f})
}

// CIMGUI_API void ImGuiIO_AddInputCharacter(ImGuiIO* self,unsigned int c);
// IMGUI_API void  AddInputCharacter(unsigned int c);
// Queue a new character input
func (io *ImGuiIO) AddInputCharacter(ch uint32) {
	giLib.Call(_func_ImGuiIO_AddInputCharacter_, []interface{}{&io, &ch})
}

// CIMGUI_API void ImGuiIO_AddInputCharacterUTF16(ImGuiIO* self,ImWchar16 c);
// IMGUI_API void  AddInputCharacterUTF16(ImWchar16 c);
// Queue a new character input from a UTF-16 character, it can be a surrogate
func (io *ImGuiIO) AddInputCharacterUTF16(ch ImWchar16) {
	giLib.Call(_func_ImGuiIO_AddInputCharacterUTF16_, []interface{}{&io, &ch})
}

// CIMGUI_API void ImGuiIO_AddInputCharactersUTF8(ImGuiIO* self,const char* str);
// IMGUI_API void  AddInputCharactersUTF8(const char* str);
// Queue a new characters input from a UTF-8 string
func (io *ImGuiIO) AddInputCharactersUTF8(str string) {
	s := c.CStr(str)
	defer usf.Free(s)
	giLib.Call(_func_ImGuiIO_AddInputCharactersUTF8_, []interface{}{&io, &s})
}

// CIMGUI_API void ImGuiIO_SetKeyEventNativeData(ImGuiIO* self,ImGuiKey key,int native_keycode,int native_scancode,int native_legacy_index);
// IMGUI_API void  SetKeyEventNativeData(ImGuiKey key, int native_keycode, int native_scancode, int native_legacy_index = -1);
// [Optional] Specify index for legacy <1.87 IsKeyXXX() functions with native indices + specify native keycode, scancode.
func (io *ImGuiIO) SetKeyEventNativeData(key ImGuiKey, native_keycode, native_scancode, native_legacy_index int32) {
	giLib.Call(_func_ImGuiIO_SetKeyEventNativeData_, []interface{}{&io, &key, &native_keycode, &native_scancode, &native_legacy_index})

}

// CIMGUI_API void ImGuiIO_SetAppAcceptingEvents(ImGuiIO* self,bool accepting_events);
// IMGUI_API void  SetAppAcceptingEvents(bool accepting_events);
// Set master flag for accepting key/mouse/text events (default to true). Useful if you have native dialog boxes that are interrupting your application loop/refresh, and you want to disable events being queued while your app is frozen.
func (io *ImGuiIO) SetAppAcceptingEvents(accepting_events bool) {
	b := c.CBool(accepting_events)
	giLib.Call(_func_ImGuiIO_SetAppAcceptingEvents_, []interface{}{&io, &b})
}

// CIMGUI_API void ImGuiIO_ClearInputCharacters(ImGuiIO* self);
// IMGUI_API void  ClearInputCharacters();
// [Internal] Clear the text input buffer manually
func (io *ImGuiIO) ClearInputCharacters() {
	giLib.Call(_func_ImGuiIO_ClearInputCharacters_, []interface{}{&io})
}

// CIMGUI_API void ImGuiIO_ClearInputKeys(ImGuiIO* self);
// IMGUI_API void  ClearInputKeys();
// [Internal] Release all keys
func (io *ImGuiIO) ClearInputKeys() {
	giLib.Call(_func_ImGuiIO_ClearInputKeys_, []interface{}{&io})
}

// CIMGUI_API ImGuiIO* ImGuiIO_ImGuiIO(void);
func NewImGuiIO() *ImGuiIO {
	return (*ImGuiIO)(giLib.Call(_func_ImGuiIO_ImGuiIO_, nil).PtrFree())
}

// CIMGUI_API void ImGuiIO_destroy(ImGuiIO* self);
func (io *ImGuiIO) Destroy() {
	giLib.Call(_func_ImGuiIO_destroy_, []interface{}{&io})
}

// CIMGUI_API ImGuiInputTextCallbackData* ImGuiInputTextCallbackData_ImGuiInputTextCallbackData(void);
func NewImGuiInputTextCallbackData() *ImGuiInputTextCallbackData {
	return (*ImGuiInputTextCallbackData)(giLib.Call(_func_ImGuiInputTextCallbackData_ImGuiInputTextCallbackData_, nil).PtrFree())

}

// CIMGUI_API void ImGuiInputTextCallbackData_destroy(ImGuiInputTextCallbackData* self);
func (cbd *ImGuiInputTextCallbackData) Destroy() {
	giLib.Call(_func_ImGuiInputTextCallbackData_destroy_, []interface{}{&cbd})
}

// CIMGUI_API void ImGuiInputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData* self,int pos,int bytes_count);
// IMGUI_API void DeleteChars(int pos, int bytes_count);
func (cbd *ImGuiInputTextCallbackData) DeleteChars(pos, bytes_count int32) {
	giLib.Call(_func_ImGuiInputTextCallbackData_DeleteChars_, []interface{}{&cbd, &pos, &bytes_count})

}

// CIMGUI_API void ImGuiInputTextCallbackData_InsertChars(ImGuiInputTextCallbackData* self,int pos,const char* text,const char* text_end);
// IMGUI_API void InsertChars(int pos, const char* text, const char* text_end = NULL);
func (cbd *ImGuiInputTextCallbackData) InsertChars(pos int32, text, text_end string) {
	txt, txt_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(txt)
	defer usf.Free(txt_end)
	giLib.Call(_func_ImGuiInputTextCallbackData_InsertChars_, []interface{}{&cbd, &pos, &txt, &txt_end})
}

// CIMGUI_API void ImGuiInputTextCallbackData_SelectAll(ImGuiInputTextCallbackData* self);
//
// void SelectAll(){ SelectionStart = 0; SelectionEnd = BufTextLen; }
func (cbd *ImGuiInputTextCallbackData) SelectAll() {
	giLib.Call(_func_ImGuiInputTextCallbackData_SelectAll_, []interface{}{&cbd})
}

// CIMGUI_API void ImGuiInputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData* self);
//
// void ClearSelection() { SelectionStart = SelectionEnd = BufTextLen; }
func (cbd *ImGuiInputTextCallbackData) ClearSelection() {
	giLib.Call(_func_ImGuiInputTextCallbackData_ClearSelection_, []interface{}{&cbd})
}

// CIMGUI_API bool ImGuiInputTextCallbackData_HasSelection(ImGuiInputTextCallbackData* self);
//
// bool HasSelection() const{ return SelectionStart != SelectionEnd; }
func (cbd *ImGuiInputTextCallbackData) HasSelection() bool {
	return giLib.Call(_func_ImGuiInputTextCallbackData_HasSelection_, []interface{}{&cbd}).BoolFree()
}

// CIMGUI_API ImGuiWindowClass* ImGuiWindowClass_ImGuiWindowClass(void);
func NewImGuiWindowClass() *ImGuiWindowClass {
	return (*ImGuiWindowClass)(giLib.Call(_func_ImGuiWindowClass_ImGuiWindowClass_, nil).PtrFree())
}

// CIMGUI_API void ImGuiWindowClass_destroy(ImGuiWindowClass* self);
func (wc *ImGuiWindowClass) Destroy() {
	giLib.Call(_func_ImGuiWindowClass_destroy_, []interface{}{&wc})
}

// CIMGUI_API ImGuiPayload* ImGuiPayload_ImGuiPayload(void);
func NewImGuiPayload() *ImGuiPayload {
	return (*ImGuiPayload)(giLib.Call(_func_ImGuiPayload_ImGuiPayload_, nil).PtrFree())
}

// CIMGUI_API void ImGuiPayload_destroy(ImGuiPayload* self);
func (paylod *ImGuiPayload) Destroy() {
	giLib.Call(_func_ImGuiPayload_destroy_, []interface{}{&paylod})
}

// CIMGUI_API void ImGuiPayload_Clear(ImGuiPayload* self);
//
// void Clear() { SourceId = SourceParentId = 0; Data = NULL; DataSize = 0; memset(DataType, 0, sizeof(DataType)); DataFrameCount = -1; Preview = Delivery = false; }
func (paylod *ImGuiPayload) Clear() {
	giLib.Call(_func_ImGuiPayload_Clear_, []interface{}{&paylod})
}

// CIMGUI_API bool ImGuiPayload_IsDataType(ImGuiPayload* self,const char* type);
//
// bool IsDataType(const char* type) const { return DataFrameCount != -1 && strcmp(type, DataType) == 0; }
func (paylod *ImGuiPayload) IsDataType(typ string) bool {
	_type := c.CStr(typ)
	defer usf.Free(_type)

	return giLib.Call(_func_ImGuiPayload_IsDataType_, []interface{}{&paylod, &_type}).BoolFree()
}

// CIMGUI_API bool ImGuiPayload_IsPreview(ImGuiPayload* self);
//
// bool IsPreview() const { return Preview; }
func (payload *ImGuiPayload) IsPreview() bool {
	return giLib.Call(_func_ImGuiPayload_IsPreview_, []interface{}{&payload}).BoolFree()
}

// CIMGUI_API bool ImGuiPayload_IsDelivery(ImGuiPayload* self);
//
// bool IsDelivery() const { return Delivery; }
func (payload *ImGuiPayload) IsDelivery() bool {
	return giLib.Call(_func_ImGuiPayload_IsDelivery_, []interface{}{&payload}).BoolFree()
}

// CIMGUI_API ImGuiTableColumnSortSpecs* ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs(void);
//
// Sorting specification for one column of a table (sizeof == 12 bytes)
func NewImGuiTableColumnSortSpecs() *ImGuiTableColumnSortSpecs {
	return (*ImGuiTableColumnSortSpecs)(giLib.Call(_func_ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs_, nil).PtrFree())
}

// CIMGUI_API void ImGuiTableColumnSortSpecs_destroy(ImGuiTableColumnSortSpecs* self);
func (ss *ImGuiTableColumnSortSpecs) Destroy() {
	giLib.Call(_func_ImGuiTableColumnSortSpecs_destroy_, []interface{}{&ss})
}

// CIMGUI_API ImGuiTableSortSpecs* ImGuiTableSortSpecs_ImGuiTableSortSpecs(void);
//
// Sorting specifications for a table (often handling sort specs for a single column, occasionally more)
// Obtained by calling TableGetSortSpecs().
// When 'SpecsDirty == true' you can sort your data. It will be true with sorting specs have changed since last call, or the first time.
// Make sure to set 'SpecsDirty = false' after sorting, else you may wastefully sort your data every frame!
func NewImGuiTableSortSpecs() *ImGuiTableSortSpecs {
	return (*ImGuiTableSortSpecs)(giLib.Call(_func_ImGuiTableSortSpecs_ImGuiTableSortSpecs_, nil).PtrFree())
}

// CIMGUI_API void ImGuiTableSortSpecs_destroy(ImGuiTableSortSpecs* self);
func (ss *ImGuiTableSortSpecs) Destroy() {
	giLib.Call(_func_ImGuiTableSortSpecs_destroy_, []interface{}{&ss})
}

// CIMGUI_API ImGuiOnceUponAFrame* ImGuiOnceUponAFrame_ImGuiOnceUponAFrame(void);
func NewImGuiOnceUponAFrame() *ImGuiOnceUponAFrame {
	return (*ImGuiOnceUponAFrame)(giLib.Call(_func_ImGuiOnceUponAFrame_ImGuiOnceUponAFrame_, nil).PtrFree())
}

// CIMGUI_API void ImGuiOnceUponAFrame_destroy(ImGuiOnceUponAFrame* self);
func (fr *ImGuiOnceUponAFrame) Destroy() {
	giLib.Call(_func_ImGuiOnceUponAFrame_destroy_, []interface{}{&fr})
}

// CIMGUI_API ImGuiTextFilter* ImGuiTextFilter_ImGuiTextFilter(const char* default_filter);
//
// IMGUI_API ImGuiTextFilter(const char* default_filter = "");
func NewImGuiTextFilter(default_filter string) *ImGuiTextFilter {
	_default_filter := c.CStr(default_filter)
	defer usf.Free(_default_filter)

	return (*ImGuiTextFilter)(giLib.Call(_func_ImGuiTextFilter_ImGuiTextFilter_, []interface{}{&_default_filter}).PtrFree())
}

// CIMGUI_API void ImGuiTextFilter_destroy(ImGuiTextFilter* self);
func (filter *ImGuiTextFilter) Destroy() {
	giLib.Call(_func_ImGuiTextFilter_destroy_, []interface{}{&filter})
}

// CIMGUI_API bool ImGuiTextFilter_Draw(ImGuiTextFilter* self,const char* label,float width);
//
// IMGUI_API bool Draw(const char* label = "Filter (inc,-exc)", float width = 0.0f);
//
// Helper calling InputText+Build
func (filter *ImGuiTextFilter) Draw(label string, width float32) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)

	return giLib.Call(_func_ImGuiTextFilter_Draw_, []interface{}{&filter, &_label, &width}).BoolFree()
}

// CIMGUI_API bool ImGuiTextFilter_PassFilter(ImGuiTextFilter* self,const char* text,const char* text_end);
//
// IMGUI_API bool PassFilter(const char* text, const char* text_end = NULL) const;
func (filter *ImGuiTextFilter) PassFilter(text string, text_end string) bool {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)

	return giLib.Call(_func_ImGuiTextFilter_PassFilter_, []interface{}{&filter, &_text, &_text_end}).BoolFree()
}

// CIMGUI_API void ImGuiTextFilter_Build(ImGuiTextFilter* self);
//
// IMGUI_API void      Build();
func (filter *ImGuiTextFilter) Build() {
	giLib.Call(_func_ImGuiTextFilter_Build_, []interface{}{&filter})
}

// CIMGUI_API void ImGuiTextFilter_Clear(ImGuiTextFilter* self);
//
// void  Clear() { InputBuf[0] = 0; Build(); }
func (filter *ImGuiTextFilter) Clear() {
	giLib.Call(_func_ImGuiTextFilter_Clear_, []interface{}{&filter})
}

// CIMGUI_API bool ImGuiTextFilter_IsActive(ImGuiTextFilter* self);
//
// bool IsActive() const { return !Filters.empty(); }
func (filter *ImGuiTextFilter) IsActive() bool {
	return giLib.Call(_func_ImGuiTextFilter_IsActive_, []interface{}{&filter}).BoolFree()
}

// CIMGUI_API ImGuiTextRange* ImGuiTextRange_ImGuiTextRange_Nil(void);
func NewImGuiTextRange() *ImGuiTextRange {
	return (*ImGuiTextRange)(giLib.Call(_func_ImGuiTextRange_ImGuiTextRange_Nil_, nil).PtrFree())
}

// CIMGUI_API void ImGuiTextRange_destroy(ImGuiTextRange* self);
func (tr *ImGuiTextRange) Destroy() {
	giLib.Call(_func_ImGuiTextRange_destroy_, []interface{}{&tr})
}

// CIMGUI_API ImGuiTextRange* ImGuiTextRange_ImGuiTextRange_Str(const char* _b,const char* _e);
//
// ImGuiTextRange(const char* _b, const char* _e)  { b = _b; e = _e; }
func NewImGuiTextRangeWithStr(b string, e string) *ImGuiTextRange {
	_b, _e := c.CStr(b), c.CStr(e)
	defer usf.Free(_b)
	defer usf.Free(_e)

	return (*ImGuiTextRange)(giLib.Call(_func_ImGuiTextRange_ImGuiTextRange_Str_, []interface{}{&_b, &_e}).PtrFree())
}

// CIMGUI_API bool ImGuiTextRange_empty(ImGuiTextRange* self);
//
// bool empty() const { return b == e; }
func (tr *ImGuiTextRange) Empty() bool {
	return giLib.Call(_func_ImGuiTextRange_empty_, []interface{}{&tr}).BoolFree()
}

// CIMGUI_API void ImGuiTextRange_split(ImGuiTextRange* self,char separator,ImVector_ImGuiTextRange* out);
//
// IMGUI_API void  split(char separator, ImVector<ImGuiTextRange>* out) const;
func (tr *ImGuiTextRange) ImGuiTextRange_split(separator byte, out *ImVector_ImGuiTextRange) {
	giLib.Call(_func_ImGuiTextRange_split_, []interface{}{&tr, &separator, &out})
}

// CIMGUI_API ImGuiTextBuffer* ImGuiTextBuffer_ImGuiTextBuffer(void);
func NewImGuiTextBuffer() *ImGuiTextBuffer {
	return (*ImGuiTextBuffer)(giLib.Call(_func_ImGuiTextBuffer_ImGuiTextBuffer_, nil).PtrFree())
}

// CIMGUI_API void ImGuiTextBuffer_destroy(ImGuiTextBuffer* self);
func (tb *ImGuiTextBuffer) Destroy() {
	giLib.Call(_func_ImGuiTextBuffer_destroy_, []interface{}{&tb})
}

// CIMGUI_API const char* ImGuiTextBuffer_begin(ImGuiTextBuffer* self);
//
// const char* begin() const { return Buf.Data ? &Buf.front() : EmptyString; }
func (tb *ImGuiTextBuffer) Begin() string {
	return giLib.Call(_func_ImGuiTextBuffer_begin_, []interface{}{&tb}).StrFree()
}

var _func_ImGuiTextBuffer_end_ = &c.FuncPrototype{Name: "ImGuiTextBuffer_end", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API const char* ImGuiTextBuffer_end(ImGuiTextBuffer* self);
//
// const char*  end() const { return Buf.Data ? &Buf.back() : EmptyString; }
// Buf is zero-terminated, so end() will point on the zero-terminator
func (tb *ImGuiTextBuffer) End() string {
	return giLib.Call(_func_ImGuiTextBuffer_end_, []interface{}{&tb}).StrFree()
}

var _func_ImGuiTextBuffer_size_ = &c.FuncPrototype{Name: "ImGuiTextBuffer_size", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int ImGuiTextBuffer_size(ImGuiTextBuffer* self);
//
// int size() const { return Buf.Size ? Buf.Size - 1 : 0; }
func (tb *ImGuiTextBuffer) Size() int32 {
	return giLib.Call(_func_ImGuiTextBuffer_size_, []interface{}{&tb}).I32Free()
}

var _func_ImGuiTextBuffer_empty_ = &c.FuncPrototype{Name: "ImGuiTextBuffer_empty", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiTextBuffer_empty(ImGuiTextBuffer* self);
func (tb *ImGuiTextBuffer) Empty() bool {
	return giLib.Call(_func_ImGuiTextBuffer_empty_, []interface{}{&tb}).BoolFree()
}

var _func_ImGuiTextBuffer_clear_ = &c.FuncPrototype{Name: "ImGuiTextBuffer_clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTextBuffer_clear(ImGuiTextBuffer* self);
//
// void clear() { Buf.clear(); }
func (tb *ImGuiTextBuffer) Clear() {
	giLib.Call(_func_ImGuiTextBuffer_clear_, []interface{}{&tb})
}

var _func_ImGuiTextBuffer_reserve_ = &c.FuncPrototype{Name: "ImGuiTextBuffer_reserve", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void ImGuiTextBuffer_reserve(ImGuiTextBuffer* self,int capacity);
//
// void reserve(int capacity) { Buf.reserve(capacity); }
func (tb *ImGuiTextBuffer) Reserve(capacity int32) {
	giLib.Call(_func_ImGuiTextBuffer_reserve_, []interface{}{&tb, &capacity})
}

var _func_ImGuiTextBuffer_c_str_ = &c.FuncPrototype{Name: "ImGuiTextBuffer_c_str", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API const char* ImGuiTextBuffer_c_str(ImGuiTextBuffer* self);
//
// const char* c_str() const { return Buf.Data ? Buf.Data : EmptyString; }
func (tb *ImGuiTextBuffer) Str() string {
	return giLib.Call(_func_ImGuiTextBuffer_c_str_, []interface{}{&tb}).StrFree()
}

var _func_ImGuiTextBuffer_append_ = &c.FuncPrototype{Name: "ImGuiTextBuffer_append", OutType: c.Void, InTypes: _typs_PPP}

// CIMGUI_API void ImGuiTextBuffer_append(ImGuiTextBuffer* self,const char* str,const char* str_end);
func (tb *ImGuiTextBuffer) Append(str string, str_end string) {
	_str, _str_end := c.CStr(str), c.CStr(str_end)
	defer usf.Free(_str)
	defer usf.Free(_str_end)
	giLib.Call(_func_ImGuiTextBuffer_append_, []interface{}{&tb, &_str, &_str_end})
}

// CIMGUI_API void ImGuiTextBuffer_appendfv(ImGuiTextBuffer* self,const char* fmt,va_list args);

var _func_ImGuiStoragePair_ImGuiStoragePair_Int_ = &c.FuncPrototype{Name: "ImGuiStoragePair_ImGuiStoragePair_Int", OutType: c.Pointer, InTypes: _typs_U32I32}

// CIMGUI_API ImGuiStoragePair* ImGuiStoragePair_ImGuiStoragePair_Int(ImGuiID _key,int _val_i);
func NewImGuiStoragePairInt(_key ImGuiID, _val_i int32) *ImGuiStoragePair {
	return (*ImGuiStoragePair)(giLib.Call(_func_ImGuiStoragePair_ImGuiStoragePair_Int_, []interface{}{&_key, &_val_i}).PtrFree())
}

var _func_ImGuiStoragePair_destroy_ = &c.FuncPrototype{Name: "ImGuiStoragePair_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiStoragePair_destroy(ImGuiStoragePair* self);
func (sp *ImGuiStoragePair) Destroy() {
	giLib.Call(_func_ImGuiStoragePair_destroy_, []interface{}{&sp})
}

var _func_ImGuiStoragePair_ImGuiStoragePair_Float_ = &c.FuncPrototype{Name: "ImGuiStoragePair_ImGuiStoragePair_Float", OutType: c.Pointer, InTypes: _typs_U32F32}

// CIMGUI_API ImGuiStoragePair* ImGuiStoragePair_ImGuiStoragePair_Float(ImGuiID _key,float _val_f);
func NewImGuiStoragePairFloat(_key ImGuiID, _val_f float32) *ImGuiStoragePair {
	return (*ImGuiStoragePair)(giLib.Call(_func_ImGuiStoragePair_ImGuiStoragePair_Float_, []interface{}{&_key, &_val_f}).PtrFree())
}

var _func_ImGuiStoragePair_ImGuiStoragePair_Ptr_ = &c.FuncPrototype{Name: "ImGuiStoragePair_ImGuiStoragePair_Ptr", OutType: c.Pointer, InTypes: _typs_U32P}

// CIMGUI_API ImGuiStoragePair* ImGuiStoragePair_ImGuiStoragePair_Ptr(ImGuiID _key,void* _val_p);
func NewImGuiStoragePairPtr(_key ImGuiID, _val_p unsafe.Pointer) *ImGuiStoragePair {
	return (*ImGuiStoragePair)(giLib.Call(_func_ImGuiStoragePair_ImGuiStoragePair_Ptr_, []interface{}{&_key, &_val_p}).PtrFree())
}

var _func_ImGuiStorage_Clear_ = &c.FuncPrototype{Name: "ImGuiStorage_Clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiStorage_Clear(ImGuiStorage* self);
func (store *ImGuiStorage) Clear() {
	giLib.Call(_func_ImGuiStorage_Clear_, []interface{}{&store})
}

var _func_ImGuiStorage_GetInt_ = &c.FuncPrototype{Name: "ImGuiStorage_GetInt", OutType: c.I32, InTypes: _typs_PU32I32}

// CIMGUI_API int ImGuiStorage_GetInt(ImGuiStorage* self,ImGuiID key,int default_val);
//
// IMGUI_API int  GetInt(ImGuiID key, int default_val = 0) const;
func (store *ImGuiStorage) GetInt(key ImGuiID, default_val int32) int32 {
	return giLib.Call(_func_ImGuiStorage_GetInt_, []interface{}{&store, &key, &default_val}).I32Free()
}

var _func_ImGuiStorage_SetInt_ = &c.FuncPrototype{Name: "ImGuiStorage_SetInt", OutType: c.Void, InTypes: _typs_PU32I32}

// CIMGUI_API void ImGuiStorage_SetInt(ImGuiStorage* self,ImGuiID key,int val);
func (store *ImGuiStorage) SetInt(key ImGuiID, val int32) {
	giLib.Call(_func_ImGuiStorage_SetInt_, []interface{}{&store, &key, &val})
}

var _func_ImGuiStorage_GetBool_ = &c.FuncPrototype{Name: "ImGuiStorage_GetBool", OutType: c.I32, InTypes: _typs_PU32I32}

// CIMGUI_API bool ImGuiStorage_GetBool(ImGuiStorage* self,ImGuiID key,bool default_val);
//
// IMGUI_API bool GetBool(ImGuiID key, bool default_val = false) const;
func (store *ImGuiStorage) GetBool(key ImGuiID, default_val bool) bool {
	_default_val := c.CBool(default_val)
	return giLib.Call(_func_ImGuiStorage_GetBool_, []interface{}{&store, &key, &_default_val}).BoolFree()
}

var _func_ImGuiStorage_SetBool_ = &c.FuncPrototype{Name: "ImGuiStorage_SetBool", OutType: c.Void, InTypes: _typs_PU32I32}

// CIMGUI_API void ImGuiStorage_SetBool(ImGuiStorage* self,ImGuiID key,bool val);
func (store *ImGuiStorage) SetBool(key ImGuiID, val bool) {
	_val := c.CBool(val)
	giLib.Call(_func_ImGuiStorage_SetBool_, []interface{}{&store, &key, &_val})
}

var _func_ImGuiStorage_GetFloat_ = &c.FuncPrototype{Name: "ImGuiStorage_GetFloat", OutType: c.F32, InTypes: _typs_PU32F32}

// CIMGUI_API float ImGuiStorage_GetFloat(ImGuiStorage* self,ImGuiID key,float default_val);
//
// IMGUI_API float GetFloat(ImGuiID key, float default_val = 0.0f) const;
func (store *ImGuiStorage) GetFloat(key ImGuiID, default_val float32) float32 {
	return giLib.Call(_func_ImGuiStorage_GetFloat_, []interface{}{&store, &key, &default_val}).F32Free()
}

var _func_ImGuiStorage_SetFloat_ = &c.FuncPrototype{Name: "ImGuiStorage_SetFloat", OutType: c.Void, InTypes: _typs_PU32F32}

// CIMGUI_API void ImGuiStorage_SetFloat(ImGuiStorage* self,ImGuiID key,float val);
func (store *ImGuiStorage) SetFloat(key ImGuiID, val float32) {
	giLib.Call(_func_ImGuiStorage_SetFloat_, []interface{}{&store, &key, &val})
}

var _func_ImGuiStorage_GetVoidPtr_ = &c.FuncPrototype{Name: "ImGuiStorage_GetVoidPtr", OutType: c.Pointer, InTypes: _typs_PU32}

// CIMGUI_API void* ImGuiStorage_GetVoidPtr(ImGuiStorage* self,ImGuiID key);
func (store *ImGuiStorage) GetPtr(key ImGuiID) unsafe.Pointer {
	return giLib.Call(_func_ImGuiStorage_GetVoidPtr_, []interface{}{&store, &key}).PtrFree()
}

var _func_ImGuiStorage_SetVoidPtr_ = &c.FuncPrototype{Name: "ImGuiStorage_SetVoidPtr", OutType: c.Void, InTypes: _typs_PU32P}

// CIMGUI_API void ImGuiStorage_SetVoidPtr(ImGuiStorage* self,ImGuiID key,void* val);
func (store *ImGuiStorage) SetPtr(key ImGuiID, val unsafe.Pointer) {
	giLib.Call(_func_ImGuiStorage_SetVoidPtr_, []interface{}{&store, &key, &val})
}

var _func_ImGuiStorage_GetIntRef_ = &c.FuncPrototype{Name: "ImGuiStorage_GetIntRef", OutType: c.Pointer, InTypes: _typs_PU32I32}

// CIMGUI_API int* ImGuiStorage_GetIntRef(ImGuiStorage* self,ImGuiID key,int default_val);
func (store *ImGuiStorage) GetIntRef(key ImGuiID, default_val int32) *int32 {
	return (*int32)(giLib.Call(_func_ImGuiStorage_GetIntRef_, []interface{}{&store, &key, &default_val}).PtrFree())
}

// CIMGUI_API bool* ImGuiStorage_GetBoolRef(ImGuiStorage* self,ImGuiID key,bool default_val);

var _func_ImGuiStorage_GetFloatRef_ = &c.FuncPrototype{Name: "ImGuiStorage_GetFloatRef", OutType: c.Pointer, InTypes: _typs_PU32F32}

// CIMGUI_API float* ImGuiStorage_GetFloatRef(ImGuiStorage* self,ImGuiID key,float default_val);
func (store *ImGuiStorage) GetFloatRef(key ImGuiID, default_val float32) *float32 {
	return (*float32)(giLib.Call(_func_ImGuiStorage_GetFloatRef_, []interface{}{&store, &key, &default_val}).PtrFree())
}

var _func_ImGuiStorage_GetVoidPtrRef_ = &c.FuncPrototype{Name: "ImGuiStorage_GetVoidPtrRef", OutType: c.Pointer, InTypes: _typs_PU32P}

// CIMGUI_API void** ImGuiStorage_GetVoidPtrRef(ImGuiStorage* self,ImGuiID key,void* default_val);
func (store *ImGuiStorage) GetVoidPtrRef(key ImGuiID, default_val unsafe.Pointer) *unsafe.Pointer {
	return (*unsafe.Pointer)(giLib.Call(_func_ImGuiStorage_GetVoidPtrRef_, []interface{}{&store, &key, &default_val}).PtrFree())
}

var _func_ImGuiStorage_SetAllInt_ = &c.FuncPrototype{Name: "ImGuiStorage_SetAllInt", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void ImGuiStorage_SetAllInt(ImGuiStorage* self,int val);
//
// IMGUI_API void SetAllInt(int val);
//
// Use on your own storage if you know only integer are being stored (open/close all tree nodes)
func (store *ImGuiStorage) SetAllInt(val int32) {
	giLib.Call(_func_ImGuiStorage_SetAllInt_, []interface{}{&store, &val})
}

var _func_ImGuiStorage_BuildSortByKey_ = &c.FuncPrototype{Name: "ImGuiStorage_BuildSortByKey", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiStorage_BuildSortByKey(ImGuiStorage* self);
//
// IMGUI_API void BuildSortByKey();
//
// For quicker full rebuild of a storage (instead of an incremental one), you may add all your contents and then sort once.
func (store *ImGuiStorage) BuildSortByKey() {
	giLib.Call(_func_ImGuiStorage_BuildSortByKey_, []interface{}{&store})
}

var _func_ImGuiListClipper_ImGuiListClipper_ = &c.FuncPrototype{Name: "ImGuiListClipper_ImGuiListClipper", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiListClipper* ImGuiListClipper_ImGuiListClipper(void);
func NewImGuiListClipper() *ImGuiListClipper {
	return (*ImGuiListClipper)(giLib.Call(_func_ImGuiListClipper_ImGuiListClipper_, nil).PtrFree())
}

var _func_ImGuiListClipper_destroy_ = &c.FuncPrototype{Name: "ImGuiListClipper_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiListClipper_destroy(ImGuiListClipper* self);
func (lc *ImGuiListClipper) Destroy(self *ImGuiListClipper) {
	giLib.Call(_func_ImGuiListClipper_destroy_, []interface{}{&lc})
}

var _func_ImGuiListClipper_Begin_ = &c.FuncPrototype{Name: "ImGuiListClipper_Begin", OutType: c.Void, InTypes: _typs_PI32F32}

// CIMGUI_API void ImGuiListClipper_Begin(ImGuiListClipper* self,int items_count,float items_height);
//
// IMGUI_API void Begin(int items_count, float items_height = -1.0f);
func (lc *ImGuiListClipper) Begin(items_count int32, items_height float32) {
	giLib.Call(_func_ImGuiListClipper_Begin_, []interface{}{&lc, &items_count, &items_height})
}

var _func_ImGuiListClipper_End_ = &c.FuncPrototype{Name: "ImGuiListClipper_End", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiListClipper_End(ImGuiListClipper* self);
//
// IMGUI_API void End();
//
// Automatically called on the last call of Step() that returns false.
func (lc *ImGuiListClipper) End() {
	giLib.Call(_func_ImGuiListClipper_End_, []interface{}{&lc})
}

var _func_ImGuiListClipper_Step_ = &c.FuncPrototype{Name: "ImGuiListClipper_Step", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiListClipper_Step(ImGuiListClipper* self);
//
// IMGUI_API bool Step();
// Call until it returns false. The DisplayStart/DisplayEnd fields
// will be set and you can process/draw those items.
func (lc *ImGuiListClipper) Step() bool {
	return giLib.Call(_func_ImGuiListClipper_Step_, []interface{}{&lc}).BoolFree()
}

var _func_ImGuiListClipper_IncludeRangeByIndices_ = &c.FuncPrototype{Name: "ImGuiListClipper_IncludeRangeByIndices", OutType: c.Void, InTypes: _typs_PI32I32}

// CIMGUI_API void ImGuiListClipper_IncludeRangeByIndices(ImGuiListClipper* self,int item_begin,int item_end);
//
// IMGUI_API void IncludeRangeByIndices(int item_begin, int item_end);
//
// item_end is exclusive e.g. use (42, 42+1) to make item 42 never clipped.
func (lc *ImGuiListClipper) IncludeRangeByIndices(item_begin int32, item_end int32) {
	giLib.Call(_func_ImGuiListClipper_IncludeRangeByIndices_, []interface{}{&lc, &item_begin, &item_end})
}

var _func_ImColor_ImColor_Nil_ = &c.FuncPrototype{Name: "ImColor_ImColor_Nil", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImColor* ImColor_ImColor_Nil(void);
func NewImColor() *ImColor {
	return (*ImColor)(giLib.Call(_func_ImColor_ImColor_Nil_, nil).PtrFree())
}

var _func_ImColor_destroy_ = &c.FuncPrototype{Name: "ImColor_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImColor_destroy(ImColor* self);
func (clr *ImColor) Destroy() {
	giLib.Call(_func_ImColor_destroy_, []interface{}{&clr})
}

var _func_ImColor_ImColor_Float_ = &c.FuncPrototype{Name: "ImColor_ImColor_Float", OutType: c.Pointer, InTypes: _typs_F32F32F32F32}

// CIMGUI_API ImColor* ImColor_ImColor_Float(float r,float g,float b,float a);
//
// constexpr ImColor(float r, float g, float b, float a = 1.0f)    : Value(r, g, b, a) { }
func NewImColorFloat(r float32, g float32, b float32, a float32) *ImColor {
	return (*ImColor)(giLib.Call(_func_ImColor_ImColor_Float_, []interface{}{&r, &g, &b, &a}).PtrFree())
}

var _func_ImColor_ImColor_Vec4_ = &c.FuncPrototype{Name: "ImColor_ImColor_Vec4", OutType: c.Pointer, InTypes: _typs_Vec4}

// CIMGUI_API ImColor* ImColor_ImColor_Vec4(const ImVec4 col);
func NewImColorVec4(col ImVec4) *ImColor {
	return (*ImColor)(giLib.Call(_func_ImColor_ImColor_Vec4_, []interface{}{&col}).PtrFree())
}

var _func_ImColor_ImColor_Int_ = &c.FuncPrototype{Name: "ImColor_ImColor_Int", OutType: c.Pointer, InTypes: _typs_I32I32I32I32}

// CIMGUI_API ImColor* ImColor_ImColor_Int(int r,int g,int b,int a);
func NewImColorInt(r int32, g int32, b int32, a int32) *ImColor {
	return (*ImColor)(giLib.Call(_func_ImColor_ImColor_Int_, []interface{}{&r, &g, &b, &a}).PtrFree())
}

var _func_ImColor_ImColor_U32_ = &c.FuncPrototype{Name: "ImColor_ImColor_U32", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImColor* ImColor_ImColor_U32(ImU32 rgba);
func NewImColorU32(rgba uint32) *ImColor {
	return (*ImColor)(giLib.Call(_func_ImColor_ImColor_U32_, []interface{}{&rgba}).PtrFree())
}

var _func_ImColor_SetHSV_ = &c.FuncPrototype{Name: "ImColor_SetHSV", OutType: c.Void, InTypes: _typs_PF32F32F32F32}

// CIMGUI_API void ImColor_SetHSV(ImColor* self,float h,float s,float v,float a);
//
// inline void SetHSV(float h, float s, float v, float a = 1.0f){ ImGui::ColorConvertHSVtoRGB(h, s, v, Value.x, Value.y, Value.z); Value.w = a; }
func (clr *ImColor) SetHSV(h float32, s float32, v float32, a float32) {
	giLib.Call(_func_ImColor_SetHSV_, []interface{}{&clr, &h, &s, &v, &a})
}

var _func_ImColor_HSV_ = &c.FuncPrototype{Name: "ImColor_HSV", OutType: c.Void, InTypes: _typs_PF32F32F32F32}

// CIMGUI_API void ImColor_HSV(ImColor* pOut,float h,float s,float v,float a);
//
// static ImColor HSV(float h, float s, float v, float a = 1.0f)   { float r, g, b; ImGui::ColorConvertHSVtoRGB(h, s, v, r, g, b); return ImColor(r, g, b, a); }
func (clr *ImColor) HSV(h float32, s float32, v float32, a float32) {
	giLib.Call(_func_ImColor_HSV_, []interface{}{&clr, &h, &s, &v, &a})
}

var _func_ImDrawCmd_ImDrawCmd_ = &c.FuncPrototype{Name: "ImDrawCmd_ImDrawCmd", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImDrawCmd* ImDrawCmd_ImDrawCmd(void);
func NewImDrawCmd() *ImDrawCmd {
	return (*ImDrawCmd)(giLib.Call(_func_ImDrawCmd_ImDrawCmd_, nil).PtrFree())
}

var _func_ImDrawCmd_destroy_ = &c.FuncPrototype{Name: "ImDrawCmd_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawCmd_destroy(ImDrawCmd* self);
func (dc *ImDrawCmd) Destroy() {
	giLib.Call(_func_ImDrawCmd_destroy_, []interface{}{&dc})
}

var _func_ImDrawCmd_GetTexID_ = &c.FuncPrototype{Name: "ImDrawCmd_GetTexID", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImTextureID ImDrawCmd_GetTexID(ImDrawCmd* self);
func (dc *ImDrawCmd) GetTexID() ImTextureID {
	return (ImTextureID)(giLib.Call(_func_ImDrawCmd_GetTexID_, []interface{}{&dc}).PtrFree())
}

var _func_ImDrawListSplitter_ImDrawListSplitter_ = &c.FuncPrototype{Name: "ImDrawListSplitter_ImDrawListSplitter", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImDrawListSplitter* ImDrawListSplitter_ImDrawListSplitter(void);
func NewImDrawListSplitter() *ImDrawListSplitter {
	return (*ImDrawListSplitter)(giLib.Call(_func_ImDrawListSplitter_ImDrawListSplitter_, nil).PtrFree())
}

var _func_ImDrawListSplitter_destroy_ = &c.FuncPrototype{Name: "ImDrawListSplitter_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawListSplitter_destroy(ImDrawListSplitter* self);
func (dls *ImDrawListSplitter) Destroy() {
	giLib.Call(_func_ImDrawListSplitter_destroy_, []interface{}{&dls})
}

var _func_ImDrawListSplitter_Clear_ = &c.FuncPrototype{Name: "ImDrawListSplitter_Clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawListSplitter_Clear(ImDrawListSplitter* self);
//
// inline void Clear() { _Current = 0; _Count = 1; }
//
// Do not clear Channels[] so our allocations are reused next frame
func (dls *ImDrawListSplitter) Clear() {
	giLib.Call(_func_ImDrawListSplitter_Clear_, []interface{}{&dls})
}

var _func_ImDrawListSplitter_ClearFreeMemory_ = &c.FuncPrototype{Name: "ImDrawListSplitter_ClearFreeMemory", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawListSplitter_ClearFreeMemory(ImDrawListSplitter* self);
func (dls *ImDrawListSplitter) ClearFreeMemory() {
	giLib.Call(_func_ImDrawListSplitter_ClearFreeMemory_, []interface{}{&dls})
}

var _func_ImDrawListSplitter_Split_ = &c.FuncPrototype{Name: "ImDrawListSplitter_Split", OutType: c.Void, InTypes: _typs_PPI32}

// CIMGUI_API void ImDrawListSplitter_Split(ImDrawListSplitter* self,ImDrawList* draw_list,int count);
func (dls *ImDrawListSplitter) Split(draw_list *ImDrawList, count int32) {
	giLib.Call(_func_ImDrawListSplitter_Split_, []interface{}{&dls, &draw_list, &count})
}

var _func_ImDrawListSplitter_Merge_ = &c.FuncPrototype{Name: "ImDrawListSplitter_Merge", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImDrawListSplitter_Merge(ImDrawListSplitter* self,ImDrawList* draw_list);
func (dls *ImDrawListSplitter) Merge(draw_list *ImDrawList) {
	giLib.Call(_func_ImDrawListSplitter_Merge_, []interface{}{&dls, &draw_list})
}

var _func_ImDrawListSplitter_SetCurrentChannel_ = &c.FuncPrototype{Name: "ImDrawListSplitter_SetCurrentChannel", OutType: c.Void, InTypes: _typs_PPI32}

// CIMGUI_API void ImDrawListSplitter_SetCurrentChannel(ImDrawListSplitter* self,ImDrawList* draw_list,int channel_idx);
func (dls *ImDrawListSplitter) SetCurrentChannel(draw_list *ImDrawList, channel_idx int32) {
	giLib.Call(_func_ImDrawListSplitter_SetCurrentChannel_, []interface{}{&dls, &draw_list, &channel_idx})
}

var _func_ImDrawList_ImDrawList_ = &c.FuncPrototype{Name: "ImDrawList_ImDrawList", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImDrawList* ImDrawList_ImDrawList(ImDrawListSharedData* shared_data);
func NewImDrawList(shared_data *ImDrawListSharedData) *ImDrawList {
	return (*ImDrawList)(giLib.Call(_func_ImDrawList_ImDrawList_, []interface{}{&shared_data}).PtrFree())
}

var _func_ImDrawList_destroy_ = &c.FuncPrototype{Name: "ImDrawList_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList_destroy(ImDrawList* self);
func (dl *ImDrawList) Destroy() {
	giLib.Call(_func_ImDrawList_destroy_, []interface{}{&dl})
}

var _func_ImDrawList_PushClipRect_ = &c.FuncPrototype{Name: "ImDrawList_PushClipRect", OutType: c.Void, InTypes: _typs_PVec2Vec2I32}

// CIMGUI_API void ImDrawList_PushClipRect(ImDrawList* self,const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect);
//
// IMGUI_API void  PushClipRect(const ImVec2& clip_rect_min, const ImVec2& clip_rect_max, bool intersect_with_current_clip_rect = false);
//
// Render-level scissoring. This is passed down to your render function but not used for CPU-side coarse clipping. Prefer using higher-level ImGui::PushClipRect() to affect logic (hit-testing and widget culling)
func (dl *ImDrawList) PushClipRectVec2(clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	_intersect_with_current_clip_rect := c.CBool(intersect_with_current_clip_rect)
	giLib.Call(_func_ImDrawList_PushClipRect_, []interface{}{&dl, &clip_rect_min, &clip_rect_max, &_intersect_with_current_clip_rect})
}

// CIMGUI_API void ImDrawList_PushClipRect(ImDrawList* self,const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect);
//
// IMGUI_API void  PushClipRect(const ImVec2& clip_rect_min, const ImVec2& clip_rect_max, bool intersect_with_current_clip_rect = false);
//
// Render-level scissoring. This is passed down to your render function but not used for CPU-side coarse clipping. Prefer using higher-level ImGui::PushClipRect() to affect logic (hit-testing and widget culling)
func (dl *ImDrawList) PushClipRect(minX, minY float32, maxX, maxY float32, intersect_with_current_clip_rect bool) {
	clip_rect_min, clip_rect_max := NewImVec2Float(minX, minY), NewImVec2Float(maxX, maxY)
	_intersect_with_current_clip_rect := c.CBool(intersect_with_current_clip_rect)
	defer clip_rect_min.Destroy()
	defer clip_rect_max.Destroy()
	giLib.Call(_func_ImDrawList_PushClipRect_, []interface{}{&dl, clip_rect_min, clip_rect_max, &_intersect_with_current_clip_rect})
}

var _func_ImDrawList_PushClipRectFullScreen_ = &c.FuncPrototype{Name: "ImDrawList_PushClipRectFullScreen", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList_PushClipRectFullScreen(ImDrawList* self);
func (dl *ImDrawList) PushClipRectFullScreen() {
	giLib.Call(_func_ImDrawList_PushClipRectFullScreen_, []interface{}{&dl})
}

var _func_ImDrawList_PopClipRect_ = &c.FuncPrototype{Name: "ImDrawList_PopClipRect", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList_PopClipRect(ImDrawList* self);
func (dl *ImDrawList) PopClipRect() {
	giLib.Call(_func_ImDrawList_PopClipRect_, []interface{}{&dl})
}

var _func_ImDrawList_PushTextureID_ = &c.FuncPrototype{Name: "ImDrawList_PushTextureID", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImDrawList_PushTextureID(ImDrawList* self,ImTextureID texture_id);
func (dl *ImDrawList) PushTextureID(texture_id ImTextureID) {
	giLib.Call(_func_ImDrawList_PushTextureID_, []interface{}{&dl, &texture_id})
}

var _func_ImDrawList_PopTextureID_ = &c.FuncPrototype{Name: "ImDrawList_PopTextureID", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList_PopTextureID(ImDrawList* self);
func (dl *ImDrawList) PopTextureID() {
	giLib.Call(_func_ImDrawList_PopTextureID_, []interface{}{&dl})
}

var _func_ImDrawList_GetClipRectMin_ = &c.FuncPrototype{Name: "ImDrawList_GetClipRectMin", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImDrawList_GetClipRectMin(ImVec2* pOut,ImDrawList* self);
func (dl *ImDrawList) GetClipRectMinVec2(pOut *ImVec2) {
	giLib.Call(_func_ImDrawList_GetClipRectMin_, []interface{}{&pOut, &dl})
}

var _func_ImDrawList_GetClipRectMax_ = &c.FuncPrototype{Name: "ImDrawList_GetClipRectMax", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImDrawList_GetClipRectMax(ImVec2* pOut,ImDrawList* self);
func (dl *ImDrawList) GetClipRectMax(pOut *ImVec2) {
	giLib.Call(_func_ImDrawList_GetClipRectMax_, []interface{}{&pOut, &dl})
}

var _func_ImDrawList_AddLine_ = &c.FuncPrototype{Name: "ImDrawList_AddLine", OutType: c.Void, InTypes: _typs_PVec2Vec2U32F32}

// CIMGUI_API void ImDrawList_AddLine(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,ImU32 col,float thickness);
//
// IMGUI_API void AddLine(const ImVec2& p1, const ImVec2& p2, ImU32 col, float thickness = 1.0f);
func (dl *ImDrawList) AddLineVec2(p1 ImVec2, p2 ImVec2, col uint32, thickness float32) {
	giLib.Call(_func_ImDrawList_AddLine_, []interface{}{&dl, &p1, &p2, &col, &thickness})
}

// CIMGUI_API void ImDrawList_AddLine(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,ImU32 col,float thickness);
//
// IMGUI_API void AddLine(const ImVec2& p1, const ImVec2& p2, ImU32 col, float thickness = 1.0f);
func (dl *ImDrawList) AddLine(x1, y1 float32, x2, y2 float32, col uint32, thickness float32) {
	p1, p2 := NewImVec2Float(x1, y1), NewImVec2Float(x2, y2)
	defer p1.Destroy()
	defer p2.Destroy()
	giLib.Call(_func_ImDrawList_AddLine_, []interface{}{&dl, p1, p2, &col, &thickness})
}

var _func_ImDrawList_AddRect_ = &c.FuncPrototype{Name: "ImDrawList_AddRect", OutType: c.Void, InTypes: _typs_PVec2Vec2U32F32U32F32}

// CIMGUI_API void ImDrawList_AddRect(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags,float thickness);
//
// IMGUI_API void  AddRect(const ImVec2& p_min, const ImVec2& p_max, ImU32 col, float rounding = 0.0f, ImDrawFlags flags = 0, float thickness = 1.0f);
//
// a: upper-left, b: lower-right (== upper-left + size)
func (dl *ImDrawList) AddRectVec2(p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags, thickness float32) {
	giLib.Call(_func_ImDrawList_AddRect_, []interface{}{&dl, &p_min, &p_max, &col, &rounding, &flags, &thickness})
}

// CIMGUI_API void ImDrawList_AddRect(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags,float thickness);
//
// IMGUI_API void  AddRect(const ImVec2& p_min, const ImVec2& p_max, ImU32 col, float rounding = 0.0f, ImDrawFlags flags = 0, float thickness = 1.0f);
//
// a: upper-left, b: lower-right (== upper-left + size)
func (dl *ImDrawList) AddRect(minX, minY float32, maxX, maxY float32, col uint32, rounding float32, flags ImDrawFlags, thickness float32) {
	p_min, p_max := NewImVec2Float(minX, minY), NewImVec2Float(maxX, maxY)
	defer p_min.Destroy()
	defer p_max.Destroy()
	giLib.Call(_func_ImDrawList_AddRect_, []interface{}{&dl, p_min, p_max, &col, &rounding, &flags, &thickness})
}

var _func_ImDrawList_AddRectFilled_ = &c.FuncPrototype{Name: "ImDrawList_AddRectFilled", OutType: c.Void, InTypes: _typs_PVec2Vec2U32F32U32}

// CIMGUI_API void ImDrawList_AddRectFilled(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags);
//
// IMGUI_API void  AddRectFilled(const ImVec2& p_min, const ImVec2& p_max, ImU32 col, float rounding = 0.0f, ImDrawFlags flags = 0);
//
// a: upper-left, b: lower-right (== upper-left + size)
func (dl *ImDrawList) AddRectFilledVec2(p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	giLib.Call(_func_ImDrawList_AddRectFilled_, []interface{}{&dl, &p_min, &p_max, &col, &rounding, &flags})
}

// CIMGUI_API void ImDrawList_AddRectFilled(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags);
//
// IMGUI_API void  AddRectFilled(const ImVec2& p_min, const ImVec2& p_max, ImU32 col, float rounding = 0.0f, ImDrawFlags flags = 0);
//
// a: upper-left, b: lower-right (== upper-left + size)
func (dl *ImDrawList) AddRectFilled(minX, minY float32, maxX, maxY float32, col uint32, rounding float32, flags ImDrawFlags) {
	p_min, p_max := NewImVec2Float(minX, minY), NewImVec2Float(maxX, maxY)
	defer p_min.Destroy()
	defer p_max.Destroy()
	giLib.Call(_func_ImDrawList_AddRectFilled_, []interface{}{&dl, p_min, p_max, &col, &rounding, &flags})
}

var _func_ImDrawList_AddRectFilledMultiColor_ = &c.FuncPrototype{Name: "ImDrawList_AddRectFilledMultiColor", OutType: c.Void, InTypes: _typs_PVec2Vec2U32U32U32U32}

// CIMGUI_API void ImDrawList_AddRectFilledMultiColor(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col_upr_left,ImU32 col_upr_right,ImU32 col_bot_right,ImU32 col_bot_left);
//
// IMGUI_API void  AddRectFilledMultiColor(const ImVec2& p_min, const ImVec2& p_max, ImU32 col_upr_left, ImU32 col_upr_right, ImU32 col_bot_right, ImU32 col_bot_left);
func (dl *ImDrawList) AddRectFilledMultiColorVec2(p_min ImVec2, p_max ImVec2,
	col_upr_left uint32, col_upr_right uint32, col_bot_right uint32, col_bot_left uint32) {
	giLib.Call(_func_ImDrawList_AddRectFilledMultiColor_, []interface{}{&dl, &p_min, &p_max, &col_upr_left, &col_upr_right, &col_bot_right, &col_bot_left})
}

// CIMGUI_API void ImDrawList_AddRectFilledMultiColor(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col_upr_left,ImU32 col_upr_right,ImU32 col_bot_right,ImU32 col_bot_left);
//
// IMGUI_API void  AddRectFilledMultiColor(const ImVec2& p_min, const ImVec2& p_max, ImU32 col_upr_left, ImU32 col_upr_right, ImU32 col_bot_right, ImU32 col_bot_left);
func (dl *ImDrawList) AddRectFilledMultiColor(minX, minY float32, maxX, maxY float32,
	col_upr_left uint32, col_upr_right uint32, col_bot_right uint32, col_bot_left uint32) {
	p_min, p_max := NewImVec2Float(minX, minY), NewImVec2Float(maxX, maxY)
	defer p_min.Destroy()
	defer p_max.Destroy()
	giLib.Call(_func_ImDrawList_AddRectFilledMultiColor_, []interface{}{&dl, p_min, p_max, &col_upr_left, &col_upr_right, &col_bot_right, &col_bot_left})
}

var _func_ImDrawList_AddQuad_ = &c.FuncPrototype{Name: "ImDrawList_AddQuad", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2U32F32}

// CIMGUI_API void ImDrawList_AddQuad(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness);
//
// IMGUI_API void  AddQuad(const ImVec2& p1, const ImVec2& p2, const ImVec2& p3, const ImVec2& p4, ImU32 col, float thickness = 1.0f);
func (dl *ImDrawList) AddQuadVec2(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32) {
	giLib.Call(_func_ImDrawList_AddQuad_, []interface{}{&dl, &p1, &p2, &p3, &p4, &col, &thickness})
}

var _func_ImDrawList_AddQuadFilled_ = &c.FuncPrototype{Name: "ImDrawList_AddQuadFilled", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2U32}

// CIMGUI_API void ImDrawList_AddQuadFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col);
//
// IMGUI_API void  AddQuadFilled(const ImVec2& p1, const ImVec2& p2, const ImVec2& p3, const ImVec2& p4, ImU32 col);
func (dl *ImDrawList) AddQuadFilled(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_AddQuadFilled_, []interface{}{&dl, &p1, &p2, &p3, &p4, &col})
}

var _func_ImDrawList_AddTriangle_ = &c.FuncPrototype{Name: "ImDrawList_AddTriangle", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2U32F32}

// CIMGUI_API void ImDrawList_AddTriangle(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness);
//
// IMGUI_API void  AddTriangle(const ImVec2& p1, const ImVec2& p2, const ImVec2& p3, ImU32 col, float thickness = 1.0f);
func (dl *ImDrawList) AddTriangle(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32) {
	giLib.Call(_func_ImDrawList_AddTriangle_, []interface{}{&dl, &p1, &p2, &p3, &col, &thickness})
}

var _func_ImDrawList_AddTriangleFilled_ = &c.FuncPrototype{Name: "ImDrawList_AddTriangleFilled", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2U32}

// CIMGUI_API void ImDrawList_AddTriangleFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col);
//
// IMGUI_API void  AddTriangleFilled(const ImVec2& p1, const ImVec2& p2, const ImVec2& p3, ImU32 col);
func (dl *ImDrawList) AddTriangleFilled(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_AddTriangleFilled_, []interface{}{&dl, &p1, &p2, &p3, &col})
}

var _func_ImDrawList_AddCircle_ = &c.FuncPrototype{Name: "ImDrawList_AddCircle", OutType: c.Void, InTypes: _typs_PVec2F32U32I32F32}

// CIMGUI_API void ImDrawList_AddCircle(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness);
//
// IMGUI_API void  AddCircle(const ImVec2& center, float radius, ImU32 col, int num_segments = 0, float thickness = 1.0f);
func (dl *ImDrawList) AddCircle(center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	giLib.Call(_func_ImDrawList_AddCircle_, []interface{}{&dl, &center, &radius, &col, &num_segments, &thickness})
}

var _func_ImDrawList_AddCircleFilled_ = &c.FuncPrototype{Name: "ImDrawList_AddCircleFilled", OutType: c.Void, InTypes: _typs_PVec2F32U32I32}

// CIMGUI_API void ImDrawList_AddCircleFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments);
//
// IMGUI_API void  AddCircleFilled(const ImVec2& center, float radius, ImU32 col, int num_segments = 0);
func (dl *ImDrawList) AddCircleFilled(center ImVec2, radius float32, col uint32, num_segments int32) {
	giLib.Call(_func_ImDrawList_AddCircleFilled_, []interface{}{&dl, &center, &radius, &col, &num_segments})
}

var _func_ImDrawList_AddNgon_ = &c.FuncPrototype{Name: "ImDrawList_AddNgon", OutType: c.Void, InTypes: _typs_PVec2F32U32I32F32}

// CIMGUI_API void ImDrawList_AddNgon(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness);
//
// IMGUI_API void  AddNgon(const ImVec2& center, float radius, ImU32 col, int num_segments, float thickness = 1.0f);
func (dl *ImDrawList) AddNgon(center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	giLib.Call(_func_ImDrawList_AddNgon_, []interface{}{&dl, &center, &radius, &col, &num_segments, &thickness})
}

var _func_ImDrawList_AddNgonFilled_ = &c.FuncPrototype{Name: "ImDrawList_AddNgonFilled", OutType: c.Void, InTypes: _typs_PVec2F32U32I32}

// CIMGUI_API void ImDrawList_AddNgonFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments);
func (dl *ImDrawList) AddNgonFilled(center ImVec2, radius float32, col uint32, num_segments int32) {
	giLib.Call(_func_ImDrawList_AddNgonFilled_, []interface{}{&dl, &center, &radius, &col, &num_segments})
}

var _func_ImDrawList_AddText_Vec2_ = &c.FuncPrototype{Name: "ImDrawList_AddText_Vec2", OutType: c.Void, InTypes: _typs_PVec2U32PP}

// CIMGUI_API void ImDrawList_AddText_Vec2(ImDrawList* self,const ImVec2 pos,ImU32 col,const char* text_begin,const char* text_end);
//
// IMGUI_API void  AddText(const ImVec2& pos, ImU32 col, const char* text_begin, const char* text_end = NULL);
func (dl *ImDrawList) AddTextVec2(pos ImVec2, col uint32, text_begin string, text_end string) {
	_text_begin, _text_end := c.CStr(text_begin), c.CStr(text_end)
	defer usf.Free(_text_begin)
	defer usf.Free(_text_end)
	giLib.Call(_func_ImDrawList_AddText_Vec2_, []interface{}{&dl, &pos, &col, &_text_begin, &_text_end})
}

var _func_ImDrawList_AddText_FontPtr_ = &c.FuncPrototype{Name: "ImDrawList_AddText_FontPtr", OutType: c.Void, InTypes: _typs_PPF32Vec2U32PPF32P}

// CIMGUI_API void ImDrawList_AddText_FontPtr(ImDrawList* self,const ImFont* font,float font_size,const ImVec2 pos,ImU32 col,const char* text_begin,const char* text_end,float wrap_width,const ImVec4* cpu_fine_clip_rect);
//
// IMGUI_API void  AddText(const ImFont* font, float font_size, const ImVec2& pos, ImU32 col, const char* text_begin, const char* text_end = NULL, float wrap_width = 0.0f, const ImVec4* cpu_fine_clip_rect = NULL);
func (dl *ImDrawList) AddTextWithFont(font *ImFont, font_size float32, pos ImVec2, col uint32, text_begin string, text_end string, wrap_width float32, cpu_fine_clip_rect *ImVec4) {
	_text_begin, _text_end := c.CStr(text_begin), c.CStr(text_end)
	defer usf.Free(_text_begin)
	defer usf.Free(_text_end)
	giLib.Call(_func_ImDrawList_AddText_FontPtr_, []interface{}{&dl, &font, &font_size, &pos, &col, &_text_begin, &_text_end, &wrap_width, &cpu_fine_clip_rect})
}

// CIMGUI_API void ImDrawList_AddText_FontPtr(ImDrawList* self,const ImFont* font,float font_size,const ImVec2 pos,ImU32 col,const char* text_begin,const char* text_end,float wrap_width,const ImVec4* cpu_fine_clip_rect);
//
// IMGUI_API void  AddText(const ImFont* font, float font_size, const ImVec2& pos, ImU32 col, const char* text_begin, const char* text_end = NULL, float wrap_width = 0.0f, const ImVec4* cpu_fine_clip_rect = NULL);
func (dl *ImDrawList) AddTextWithFontDefault(font *ImFont, font_size float32, pos ImVec2, col uint32, text_begin string) {
	_text_begin, _text_end, wrap_width, cpu_fine_clip_rect := c.CStr(text_begin), unsafe.Pointer(nil), _f32_zero, unsafe.Pointer(nil)
	defer usf.Free(_text_begin)
	giLib.Call(_func_ImDrawList_AddText_FontPtr_, []interface{}{&dl, &font, &font_size, &pos, &col, &_text_begin, &_text_end, wrap_width, &cpu_fine_clip_rect})
}

var _func_ImDrawList_AddPolyline_ = &c.FuncPrototype{Name: "ImDrawList_AddPolyline", OutType: c.Void, InTypes: _typs_PPI32U32U32F32}

// CIMGUI_API void ImDrawList_AddPolyline(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col,ImDrawFlags flags,float thickness);
//
// IMGUI_API void  AddPolyline(const ImVec2* points, int num_points, ImU32 col, ImDrawFlags flags, float thickness);
func (dl *ImDrawList) AddPolyline(points []ImVec2, col uint32, flags ImDrawFlags, thickness float32) {
	var _points *ImVec2
	var num_points int32
	if points != nil {
		_points, num_points = &points[0], int32(len(points))
	}
	giLib.Call(_func_ImDrawList_AddPolyline_, []interface{}{&dl, &_points, &num_points, &col, &flags, &thickness})
}

var _func_ImDrawList_AddConvexPolyFilled_ = &c.FuncPrototype{Name: "ImDrawList_AddConvexPolyFilled", OutType: c.Void, InTypes: _typs_PPI32U32}

// CIMGUI_API void ImDrawList_AddConvexPolyFilled(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col);
func (dl *ImDrawList) AddConvexPolyFilled(points []ImVec2, col uint32) {
	var _points *ImVec2
	var num_points int32
	if points != nil {
		_points, num_points = &points[0], int32(len(points))
	}
	giLib.Call(_func_ImDrawList_AddConvexPolyFilled_, []interface{}{&dl, &_points, &num_points, &col})
}

var _func_ImDrawList_AddBezierCubic_ = &c.FuncPrototype{Name: "ImDrawList_AddBezierCubic", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2U32F32I32}

// CIMGUI_API void ImDrawList_AddBezierCubic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness,int num_segments);
//
// IMGUI_API void  AddBezierCubic(const ImVec2& p1, const ImVec2& p2, const ImVec2& p3, const ImVec2& p4, ImU32 col, float thickness, int num_segments = 0);
//
// Cubic Bezier (4 control points)
func (dl *ImDrawList) AddBezierCubic(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32, num_segments int32) {
	giLib.Call(_func_ImDrawList_AddBezierCubic_, []interface{}{&dl, &p1, &p2, &p3, &p4, &col, &thickness, &num_segments})
}

var _func_ImDrawList_AddBezierQuadratic_ = &c.FuncPrototype{Name: "ImDrawList_AddBezierQuadratic", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2U32F32I32}

// CIMGUI_API void ImDrawList_AddBezierQuadratic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness,int num_segments);
//
// IMGUI_API void  AddBezierQuadratic(const ImVec2& p1, const ImVec2& p2, const ImVec2& p3, ImU32 col, float thickness, int num_segments = 0);               // Quadratic Bezier (3 control points)
func (dl *ImDrawList) AddBezierQuadratic(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32, num_segments int32) {
	giLib.Call(_func_ImDrawList_AddBezierQuadratic_, []interface{}{&dl, &p1, &p2, &p3, &col, &thickness, &num_segments})
}

var _func_ImDrawList_AddImage_ = &c.FuncPrototype{Name: "ImDrawList_AddImage", OutType: c.Void, InTypes: _typs_PPVec2Vec2Vec2Vec2U32}

// CIMGUI_API void ImDrawList_AddImage(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col);
//
// IMGUI_API void  AddImage(ImTextureID user_texture_id, const ImVec2& p_min, const ImVec2& p_max, const ImVec2& uv_min = ImVec2(0, 0), const ImVec2& uv_max = ImVec2(1, 1), ImU32 col = IM_COL32_WHITE);
func (dl *ImDrawList) AddImage(user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_AddImage_, []interface{}{&dl, &user_texture_id, &p_min, &p_max, &uv_min, &uv_max, &col})
}

// CIMGUI_API void ImDrawList_AddImage(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col);
//
// IMGUI_API void  AddImage(ImTextureID user_texture_id, const ImVec2& p_min, const ImVec2& p_max, const ImVec2& uv_min = ImVec2(0, 0), const ImVec2& uv_max = ImVec2(1, 1), ImU32 col = IM_COL32_WHITE);
func (dl *ImDrawList) AddImageDefault(user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2) {
	uv_min, uv_max, col := _vec2_zero, _vec2_ones, uint32(0xFFFFFFFF)
	giLib.Call(_func_ImDrawList_AddImage_, []interface{}{&dl, &user_texture_id, &p_min, &p_max, uv_min, uv_max, &col})
}

var _func_ImDrawList_AddImageQuad_ = &c.FuncPrototype{Name: "ImDrawList_AddImageQuad", OutType: c.Void, InTypes: _typs_PPVec2Vec2Vec2Vec2Vec2Vec2Vec2Vec2U32}

// CIMGUI_API void ImDrawList_AddImageQuad(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 uv1,const ImVec2 uv2,const ImVec2 uv3,const ImVec2 uv4,ImU32 col);
//
// IMGUI_API void  AddImageQuad(ImTextureID user_texture_id, const ImVec2& p1, const ImVec2& p2, const ImVec2& p3, const ImVec2& p4, const ImVec2& uv1 = ImVec2(0, 0), const ImVec2& uv2 = ImVec2(1, 0), const ImVec2& uv3 = ImVec2(1, 1), const ImVec2& uv4 = ImVec2(0, 1), ImU32 col = IM_COL32_WHITE);
func (dl *ImDrawList) AddImageQuad(user_texture_id ImTextureID, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, uv1 ImVec2, uv2 ImVec2, uv3 ImVec2, uv4 ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_AddImageQuad_, []interface{}{&dl, &user_texture_id, &p1, &p2, &p3, &p4, &uv1, &uv2, &uv3, &uv4, &col})
}

// CIMGUI_API void ImDrawList_AddImageQuad(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 uv1,const ImVec2 uv2,const ImVec2 uv3,const ImVec2 uv4,ImU32 col);
//
// IMGUI_API void  AddImageQuad(ImTextureID user_texture_id, const ImVec2& p1, const ImVec2& p2, const ImVec2& p3, const ImVec2& p4, const ImVec2& uv1 = ImVec2(0, 0), const ImVec2& uv2 = ImVec2(1, 0), const ImVec2& uv3 = ImVec2(1, 1), const ImVec2& uv4 = ImVec2(0, 1), ImU32 col = IM_COL32_WHITE);
func (dl *ImDrawList) AddImageQuadDefault(user_texture_id ImTextureID, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2) {
	uv1, uv2, uv3, uv4, col := _vec2_zero, _vec2_10, _vec2_ones, _vec2_01, uint32(0xFFFFFFFF)
	giLib.Call(_func_ImDrawList_AddImageQuad_, []interface{}{&dl, &user_texture_id, &p1, &p2, &p3, &p4, uv1, uv2, uv3, uv4, &col})
}

var _func_ImDrawList_AddImageRounded_ = &c.FuncPrototype{Name: "ImDrawList_AddImageRounded", OutType: c.Void, InTypes: _typs_PPVec2Vec2Vec2Vec2U32F32U32}

// CIMGUI_API void ImDrawList_AddImageRounded(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col,float rounding,ImDrawFlags flags);
//
// IMGUI_API void  AddImageRounded(ImTextureID user_texture_id, const ImVec2& p_min, const ImVec2& p_max, const ImVec2& uv_min, const ImVec2& uv_max, ImU32 col, float rounding, ImDrawFlags flags = 0);
func (dl *ImDrawList) AddImageRounded(user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	giLib.Call(_func_ImDrawList_AddImageRounded_, []interface{}{&dl, &user_texture_id, &p_min, &p_max, &uv_min, &uv_max, &col, &rounding, &flags})
}

var _func_ImDrawList_PathClear_ = &c.FuncPrototype{Name: "ImDrawList_PathClear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList_PathClear(ImDrawList* self);
//
// inline void  PathClear() { _Path.Size = 0; }
func (dl *ImDrawList) PathClear() {
	giLib.Call(_func_ImDrawList_PathClear_, []interface{}{&dl})
}

var _func_ImDrawList_PathLineTo_ = &c.FuncPrototype{Name: "ImDrawList_PathLineTo", OutType: c.Void, InTypes: _typs_PVec2}

// CIMGUI_API void ImDrawList_PathLineTo(ImDrawList* self,const ImVec2 pos);
//
// inline  void  PathLineTo(const ImVec2& pos) { _Path.push_back(pos); }
func (dl *ImDrawList) PathLineTo(pos ImVec2) {
	giLib.Call(_func_ImDrawList_PathLineTo_, []interface{}{&dl, &pos})
}

var _func_ImDrawList_PathLineToMergeDuplicate_ = &c.FuncPrototype{Name: "ImDrawList_PathLineToMergeDuplicate", OutType: c.Void, InTypes: _typs_PVec2}

// CIMGUI_API void ImDrawList_PathLineToMergeDuplicate(ImDrawList* self,const ImVec2 pos);
func (dl *ImDrawList) PathLineToMergeDuplicate(pos ImVec2) {
	giLib.Call(_func_ImDrawList_PathLineToMergeDuplicate_, []interface{}{&dl, &pos})
}

var _func_ImDrawList_PathFillConvex_ = &c.FuncPrototype{Name: "ImDrawList_PathFillConvex", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void ImDrawList_PathFillConvex(ImDrawList* self,ImU32 col);
func (dl *ImDrawList) PathFillConvex(col uint32) {
	giLib.Call(_func_ImDrawList_PathFillConvex_, []interface{}{&dl, &col})
}

var _func_ImDrawList_PathStroke_ = &c.FuncPrototype{Name: "ImDrawList_PathStroke", OutType: c.Void, InTypes: _typs_PU32U32F32}

// CIMGUI_API void ImDrawList_PathStroke(ImDrawList* self,ImU32 col,ImDrawFlags flags,float thickness);
//
// inline void  PathStroke(ImU32 col, ImDrawFlags flags = 0, float thickness = 1.0f)
func (dl *ImDrawList) PathStroke(col uint32, flags ImDrawFlags, thickness float32) {
	giLib.Call(_func_ImDrawList_PathStroke_, []interface{}{&dl, &col, &flags, &thickness})
}

var _func_ImDrawList_PathArcTo_ = &c.FuncPrototype{Name: "ImDrawList_PathArcTo", OutType: c.Void, InTypes: _typs_PVec2F32F32F32I32}

// CIMGUI_API void ImDrawList_PathArcTo(ImDrawList* self,const ImVec2 center,float radius,float a_min,float a_max,int num_segments);
//
// IMGUI_API void  PathArcTo(const ImVec2& center, float radius, float a_min, float a_max, int num_segments = 0);
func (dl *ImDrawList) PathArcTo(center ImVec2, radius float32, a_min float32, a_max float32, num_segments int32) {
	giLib.Call(_func_ImDrawList_PathArcTo_, []interface{}{&dl, &center, &radius, &a_min, &a_max, &num_segments})
}

var _func_ImDrawList_PathArcToFast_ = &c.FuncPrototype{Name: "ImDrawList_PathArcToFast", OutType: c.Void, InTypes: _typs_PVec2F32I32I32}

// CIMGUI_API void ImDrawList_PathArcToFast(ImDrawList* self,const ImVec2 center,float radius,int a_min_of_12,int a_max_of_12);
func (dl *ImDrawList) PathArcToFast(center ImVec2, radius float32, a_min_of_12 int32, a_max_of_12 int32) {
	giLib.Call(_func_ImDrawList_PathArcToFast_, []interface{}{&dl, &center, &radius, &a_min_of_12, &a_max_of_12})
}

var _func_ImDrawList_PathBezierCubicCurveTo_ = &c.FuncPrototype{Name: "ImDrawList_PathBezierCubicCurveTo", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2I32}

// CIMGUI_API void ImDrawList_PathBezierCubicCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,int num_segments);
func (dl *ImDrawList) PathBezierCubicCurveTo(p2, p3, p4 ImVec2, num_segments int32) {
	giLib.Call(_func_ImDrawList_PathBezierCubicCurveTo_, []interface{}{&dl, &p2, &p3, &p4, &num_segments})
}

var _func_ImDrawList_PathBezierQuadraticCurveTo_ = &c.FuncPrototype{Name: "ImDrawList_PathBezierQuadraticCurveTo", OutType: c.Void, InTypes: _typs_PVec2Vec2I32}

// CIMGUI_API void ImDrawList_PathBezierQuadraticCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,int num_segments);
func (dl *ImDrawList) PathBezierQuadraticCurveTo(p2, p3 ImVec2, num_segments int32) {
	giLib.Call(_func_ImDrawList_PathBezierQuadraticCurveTo_, []interface{}{&dl, &p2, &p3, &num_segments})
}

var _func_ImDrawList_PathRect_ = &c.FuncPrototype{Name: "ImDrawList_PathRect", OutType: c.Void, InTypes: _typs_PVec2Vec2F32U32}

// CIMGUI_API void ImDrawList_PathRect(ImDrawList* self,const ImVec2 rect_min,const ImVec2 rect_max,float rounding,ImDrawFlags flags);
func (dl *ImDrawList) PathRect(rect_min, rect_max ImVec2, rounding float32, flags ImDrawFlags) {
	giLib.Call(_func_ImDrawList_PathRect_, []interface{}{&dl, &rect_min, &rect_max, &rounding, &flags})
}

// CIMGUI_API void ImDrawList_AddCallback(ImDrawList* self,ImDrawCallback callback,void* callback_data);

var _func_ImDrawList_AddDrawCmd_ = &c.FuncPrototype{Name: "ImDrawList_AddDrawCmd", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList_AddDrawCmd(ImDrawList* self);
func (dl *ImDrawList) AddDrawCmd() {
	giLib.Call(_func_ImDrawList_AddDrawCmd_, []interface{}{&dl})
}

var _func_ImDrawList_CloneOutput_ = &c.FuncPrototype{Name: "ImDrawList_CloneOutput", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImDrawList* ImDrawList_CloneOutput(ImDrawList* self);
func (dl *ImDrawList) CloneOutput() *ImDrawList {
	return (*ImDrawList)(giLib.Call(_func_ImDrawList_CloneOutput_, []interface{}{&dl}).PtrFree())
}

var _func_ImDrawList_ChannelsSplit_ = &c.FuncPrototype{Name: "ImDrawList_ChannelsSplit", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void ImDrawList_ChannelsSplit(ImDrawList* self,int count);
func (dl *ImDrawList) ChannelsSplit(count int32) {
	giLib.Call(_func_ImDrawList_ChannelsSplit_, []interface{}{&dl, &count})
}

var _func_ImDrawList_ChannelsMerge_ = &c.FuncPrototype{Name: "ImDrawList_ChannelsMerge", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList_ChannelsMerge(ImDrawList* self);
func (dl *ImDrawList) ChannelsMerge() {
	giLib.Call(_func_ImDrawList_ChannelsMerge_, []interface{}{&dl})
}

var _func_ImDrawList_ChannelsSetCurrent_ = &c.FuncPrototype{Name: "ImDrawList_ChannelsSetCurrent", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void ImDrawList_ChannelsSetCurrent(ImDrawList* self,int n);
func (dl *ImDrawList) ChannelsSetCurrent(n int32) {
	giLib.Call(_func_ImDrawList_ChannelsSetCurrent_, []interface{}{&dl, &n})
}

var _func_ImDrawList_PrimReserve_ = &c.FuncPrototype{Name: "ImDrawList_PrimReserve", OutType: c.Void, InTypes: _typs_PI32I32}

// CIMGUI_API void ImDrawList_PrimReserve(ImDrawList* self,int idx_count,int vtx_count);
func (dl *ImDrawList) PrimReserve(idx_count int32, vtx_count int32) {
	giLib.Call(_func_ImDrawList_PrimReserve_, []interface{}{&dl, &idx_count, &vtx_count})
}

var _func_ImDrawList_PrimUnreserve_ = &c.FuncPrototype{Name: "ImDrawList_PrimUnreserve", OutType: c.Void, InTypes: _typs_PI32I32}

// CIMGUI_API void ImDrawList_PrimUnreserve(ImDrawList* self,int idx_count,int vtx_count);
func (dl *ImDrawList) PrimUnreserve(idx_count int32, vtx_count int32) {
	giLib.Call(_func_ImDrawList_PrimUnreserve_, []interface{}{&dl, &idx_count, &vtx_count})
}

var _func_ImDrawList_PrimRect_ = &c.FuncPrototype{Name: "ImDrawList_PrimRect", OutType: c.Void, InTypes: _typs_PVec2Vec2U32}

// CIMGUI_API void ImDrawList_PrimRect(ImDrawList* self,const ImVec2 a,const ImVec2 b,ImU32 col);
func (dl *ImDrawList) PrimRect(a ImVec2, b ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_PrimRect_, []interface{}{&dl, &a, &b, &col})
}

var _func_ImDrawList_PrimRectUV_ = &c.FuncPrototype{Name: "ImDrawList_PrimRectUV", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2U32}

// CIMGUI_API void ImDrawList_PrimRectUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 uv_a,const ImVec2 uv_b,ImU32 col);
func (dl *ImDrawList) PrimRectUV(a ImVec2, b ImVec2, uv_a ImVec2, uv_b ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_PrimRectUV_, []interface{}{&dl, &a, &b, &uv_a, &uv_b, &col})
}

var _func_ImDrawList_PrimQuadUV_ = &c.FuncPrototype{Name: "ImDrawList_PrimQuadUV", OutType: c.Void, InTypes: _typs_PVec2Vec2Vec2Vec2Vec2Vec2Vec2Vec2U32}

// CIMGUI_API void ImDrawList_PrimQuadUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 d,const ImVec2 uv_a,const ImVec2 uv_b,const ImVec2 uv_c,const ImVec2 uv_d,ImU32 col);
func (dl *ImDrawList) PrimQuadUV(a ImVec2, b ImVec2, c ImVec2, d ImVec2, uv_a ImVec2, uv_b ImVec2, uv_c ImVec2, uv_d ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_PrimQuadUV_, []interface{}{&dl, &a, &b, &c, &d, &uv_a, &uv_b, &uv_c, &uv_d, &col})
}

var _func_ImDrawList_PrimWriteVtx_ = &c.FuncPrototype{Name: "ImDrawList_PrimWriteVtx", OutType: c.Void, InTypes: _typs_PVec2Vec2U32}

// CIMGUI_API void ImDrawList_PrimWriteVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col);
func (dl *ImDrawList) PrimWriteVtx(pos ImVec2, uv ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_PrimWriteVtx_, []interface{}{&dl, &pos, &uv, &col})
}

var _func_ImDrawList_PrimWriteIdx_ = &c.FuncPrototype{Name: "ImDrawList_PrimWriteIdx", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList_PrimWriteIdx(ImDrawList* self,ImDrawIdx idx);
func (dl *ImDrawList) PrimWriteIdx(idx ImDrawIdx) {
	giLib.Call(_func_ImDrawList_PrimWriteIdx_, []interface{}{&dl, &idx})
}

var _func_ImDrawList_PrimVtx_ = &c.FuncPrototype{Name: "ImDrawList_PrimVtx", OutType: c.Void, InTypes: _typs_PVec2Vec2U32}

// CIMGUI_API void ImDrawList_PrimVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col);
func (dl *ImDrawList) PrimVtx(self *ImDrawList, pos ImVec2, uv ImVec2, col uint32) {
	giLib.Call(_func_ImDrawList_PrimVtx_, []interface{}{&dl, &pos, &uv, &col})
}

var _func_ImDrawList__ResetForNewFrame_ = &c.FuncPrototype{Name: "ImDrawList__ResetForNewFrame", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList__ResetForNewFrame(ImDrawList* self);
func (dl *ImDrawList) ResetForNewFrame() {
	giLib.Call(_func_ImDrawList__ResetForNewFrame_, []interface{}{&dl})
}

var _func_ImDrawList__ClearFreeMemory_ = &c.FuncPrototype{Name: "ImDrawList__ClearFreeMemory", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList__ClearFreeMemory(ImDrawList* self);
func (dl *ImDrawList) ClearFreeMemory() {
	giLib.Call(_func_ImDrawList__ClearFreeMemory_, []interface{}{&dl})
}

var _func_ImDrawList__PopUnusedDrawCmd_ = &c.FuncPrototype{Name: "ImDrawList__PopUnusedDrawCmd", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList__PopUnusedDrawCmd(ImDrawList* self);
func (dl *ImDrawList) PopUnusedDrawCmd() {
	giLib.Call(_func_ImDrawList__PopUnusedDrawCmd_, []interface{}{&dl})
}

var _func_ImDrawList__TryMergeDrawCmds_ = &c.FuncPrototype{Name: "ImDrawList__TryMergeDrawCmds", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList__TryMergeDrawCmds(ImDrawList* self);
func (dl *ImDrawList) TryMergeDrawCmds() {
	giLib.Call(_func_ImDrawList__TryMergeDrawCmds_, []interface{}{&dl})
}

var _func_ImDrawList__OnChangedClipRect_ = &c.FuncPrototype{Name: "ImDrawList__OnChangedClipRect", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList__OnChangedClipRect(ImDrawList* self);
func (dl *ImDrawList) OnChangedClipRect(self *ImDrawList) {
	giLib.Call(_func_ImDrawList__OnChangedClipRect_, []interface{}{&dl})
}

var _func_ImDrawList__OnChangedTextureID_ = &c.FuncPrototype{Name: "ImDrawList__OnChangedTextureID", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList__OnChangedTextureID(ImDrawList* self);
func (dl *ImDrawList) OnChangedTextureID() {
	giLib.Call(_func_ImDrawList__OnChangedTextureID_, []interface{}{&dl})
}

var _func_ImDrawList__OnChangedVtxOffset_ = &c.FuncPrototype{Name: "ImDrawList__OnChangedVtxOffset", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawList__OnChangedVtxOffset(ImDrawList* self);
func (dl *ImDrawList) OnChangedVtxOffset() {
	giLib.Call(_func_ImDrawList__OnChangedVtxOffset_, []interface{}{&dl})
}

var _func_ImDrawList__CalcCircleAutoSegmentCount_ = &c.FuncPrototype{Name: "ImDrawList__CalcCircleAutoSegmentCount", OutType: c.I32, InTypes: _typs_PF32}

// CIMGUI_API int ImDrawList__CalcCircleAutoSegmentCount(ImDrawList* self,float radius);
func (dl *ImDrawList) CalcCircleAutoSegmentCount(self *ImDrawList, radius float32) int32 {
	return giLib.Call(_func_ImDrawList__CalcCircleAutoSegmentCount_, []interface{}{&dl, &radius}).I32Free()
}

var _func_ImDrawList__PathArcToFastEx_ = &c.FuncPrototype{Name: "ImDrawList__PathArcToFastEx", OutType: c.Void, InTypes: _typs_PVec2F32I32I32I32}

// CIMGUI_API void ImDrawList__PathArcToFastEx(ImDrawList* self,const ImVec2 center,float radius,int a_min_sample,int a_max_sample,int a_step);
func (dl *ImDrawList) PathArcToFastEx(center ImVec2, radius float32, a_min_sample int32, a_max_sample int32, a_step int32) {
	giLib.Call(_func_ImDrawList__PathArcToFastEx_, []interface{}{&dl, &center, &radius, &a_min_sample, &a_max_sample, &a_step})
}

var _func_ImDrawList__PathArcToN_ = &c.FuncPrototype{Name: "ImDrawList__PathArcToN", OutType: c.Void, InTypes: _typs_PVec2F32F32F32I32}

// CIMGUI_API void ImDrawList__PathArcToN(ImDrawList* self,const ImVec2 center,float radius,float a_min,float a_max,int num_segments);
func (dl *ImDrawList) PathArcToN(center ImVec2, radius float32, a_min float32, a_max float32, num_segments int32) {
	giLib.Call(_func_ImDrawList__PathArcToN_, []interface{}{&dl, &center, &radius, &a_min, &a_max, &num_segments})
}

var _func_ImDrawData_ImDrawData_ = &c.FuncPrototype{Name: "ImDrawData_ImDrawData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImDrawData* ImDrawData_ImDrawData(void);
func NewImDrawData() *ImDrawData {
	return (*ImDrawData)(giLib.Call(_func_ImDrawData_ImDrawData_, nil).PtrFree())
}

var _func_ImDrawData_destroy_ = &c.FuncPrototype{Name: "ImDrawData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawData_destroy(ImDrawData* self);
func (dd *ImDrawData) Destroy() {
	giLib.Call(_func_ImDrawData_destroy_, []interface{}{&dd})
}

var _func_ImDrawData_Clear_ = &c.FuncPrototype{Name: "ImDrawData_Clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawData_Clear(ImDrawData* self);
func (dd *ImDrawData) Clear() {
	giLib.Call(_func_ImDrawData_Clear_, []interface{}{&dd})
}

var _func_ImDrawData_DeIndexAllBuffers_ = &c.FuncPrototype{Name: "ImDrawData_DeIndexAllBuffers", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawData_DeIndexAllBuffers(ImDrawData* self);
func (dd *ImDrawData) DeIndexAllBuffers() {
	giLib.Call(_func_ImDrawData_DeIndexAllBuffers_, []interface{}{&dd})
}

var _func_ImDrawData_ScaleClipRects_ = &c.FuncPrototype{Name: "ImDrawData_ScaleClipRects", OutType: c.Void, InTypes: _typs_PVec2}

// CIMGUI_API void ImDrawData_ScaleClipRects(ImDrawData* self,const ImVec2 fb_scale);
func (dd *ImDrawData) ScaleClipRects(fb_scale ImVec2) {
	giLib.Call(_func_ImDrawData_ScaleClipRects_, []interface{}{&dd, &fb_scale})
}

// CIMGUI_API ImFontConfig* ImFontConfig_ImFontConfig(void);
func NewImFontConfig() *ImFontConfig {
	return (*ImFontConfig)(giLib.Call(_func_ImFontConfig_ImFontConfig_, nil).PtrFree())
}

// CIMGUI_API void ImFontConfig_destroy(ImFontConfig* self);
func (fc *ImFontConfig) Destroy() {
	giLib.Call(_func_ImFontConfig_destroy_, []interface{}{&fc})
}

var _func_ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImFontGlyphRangesBuilder* ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder(void);
func NewImFontGlyphRangesBuilder() *ImFontGlyphRangesBuilder {
	return (*ImFontGlyphRangesBuilder)(giLib.Call(_func_ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder_, nil).PtrFree())
}

var _func_ImFontGlyphRangesBuilder_destroy_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFontGlyphRangesBuilder_destroy(ImFontGlyphRangesBuilder* self);
func (fgrb *ImFontGlyphRangesBuilder) Destroy() {
	giLib.Call(_func_ImFontGlyphRangesBuilder_destroy_, []interface{}{&fgrb})
}

var _func_ImFontGlyphRangesBuilder_Clear_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_Clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder* self);
func (fgrb *ImFontGlyphRangesBuilder) Clear() {
	giLib.Call(_func_ImFontGlyphRangesBuilder_Clear_, []interface{}{&fgrb})
}

var _func_ImFontGlyphRangesBuilder_GetBit_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_GetBit", OutType: c.I32, InTypes: _typs_PU64}

// CIMGUI_API bool ImFontGlyphRangesBuilder_GetBit(ImFontGlyphRangesBuilder* self,size_t n);
func (fgrb *ImFontGlyphRangesBuilder) GetBit(n uint64) bool {
	return giLib.Call(_func_ImFontGlyphRangesBuilder_GetBit_, []interface{}{&fgrb, &n}).BoolFree()
}

var _func_ImFontGlyphRangesBuilder_SetBit_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_SetBit", OutType: c.Void, InTypes: _typs_PU64}

// CIMGUI_API void ImFontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder* self,size_t n);
func (fgrb *ImFontGlyphRangesBuilder) SetBit(n uint64) {
	giLib.Call(_func_ImFontGlyphRangesBuilder_SetBit_, []interface{}{&fgrb, &n})
}

var _func_ImFontGlyphRangesBuilder_AddChar_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_AddChar", OutType: c.Void, InTypes: _typs_PU16}

// CIMGUI_API void ImFontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder* self,ImWchar c);
func (fgrb *ImFontGlyphRangesBuilder) AddChar(c uint16) {
	giLib.Call(_func_ImFontGlyphRangesBuilder_AddChar_, []interface{}{&fgrb, &c})
}

var _func_ImFontGlyphRangesBuilder_AddText_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_AddText", OutType: c.Void, InTypes: _typs_PPP}

// CIMGUI_API void ImFontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder* self,const char* text,const char* text_end);
//
// IMGUI_API void  AddText(const char* text, const char* text_end = NULL);
//
// Add string (each character of the UTF-8 string are added)
func (fgrb *ImFontGlyphRangesBuilder) AddText(text string, text_end string) {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	giLib.Call(_func_ImFontGlyphRangesBuilder_AddText_, []interface{}{&fgrb, &_text, &_text_end})
}

var _func_ImFontGlyphRangesBuilder_AddRanges_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_AddRanges", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImFontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder* self,const ImWchar* ranges);
func (fgrb *ImFontGlyphRangesBuilder) AddRanges(ranges []ImWchar) {
	var r *ImWchar
	if ranges != nil {
		r = &ranges[0]
	}
	giLib.Call(_func_ImFontGlyphRangesBuilder_AddRanges_, []interface{}{&fgrb, &r})
}

var _func_ImFontGlyphRangesBuilder_BuildRanges_ = &c.FuncPrototype{Name: "ImFontGlyphRangesBuilder_BuildRanges", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImFontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder* self,ImVector_ImWchar* out_ranges)
func (fgrb *ImFontGlyphRangesBuilder) BuildRanges(out_ranges *ImVector_ImWchar) {
	giLib.Call(_func_ImFontGlyphRangesBuilder_BuildRanges_, []interface{}{&fgrb, &out_ranges})
}

var _func_ImFontAtlasCustomRect_ImFontAtlasCustomRect_ = &c.FuncPrototype{Name: "ImFontAtlasCustomRect_ImFontAtlasCustomRect", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImFontAtlasCustomRect* ImFontAtlasCustomRect_ImFontAtlasCustomRect(void);
func NewImFontAtlasCustomRect() *ImFontAtlasCustomRect {
	return (*ImFontAtlasCustomRect)(giLib.Call(_func_ImFontAtlasCustomRect_ImFontAtlasCustomRect_, nil).PtrFree())
}

var _func_ImFontAtlasCustomRect_destroy_ = &c.FuncPrototype{Name: "ImFontAtlasCustomRect_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFontAtlasCustomRect_destroy(ImFontAtlasCustomRect* self);
func (facr *ImFontAtlasCustomRect) Destroy() {
	giLib.Call(_func_ImFontAtlasCustomRect_destroy_, []interface{}{&facr})
}

var _func_ImFontAtlasCustomRect_IsPacked_ = &c.FuncPrototype{Name: "ImFontAtlasCustomRect_IsPacked", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImFontAtlasCustomRect_IsPacked(ImFontAtlasCustomRect* self);
func (facr *ImFontAtlasCustomRect) IsPacked() bool {
	return giLib.Call(_func_ImFontAtlasCustomRect_IsPacked_, []interface{}{&facr}).BoolFree()
}

// CIMGUI_API ImFontAtlas* ImFontAtlas_ImFontAtlas(void);
func NewImFontAtlas() *ImFontAtlas {
	return (*ImFontAtlas)(giLib.Call(_func_ImFontAtlas_ImFontAtlas_, nil).PtrFree())
}

// CIMGUI_API void ImFontAtlas_destroy(ImFontAtlas* self);
func (fa *ImFontAtlas) Destroy() {
	giLib.Call(_func_ImFontAtlas_destroy_, []interface{}{&fa})
}

// CIMGUI_API ImFont* ImFontAtlas_AddFont(ImFontAtlas* self,const ImFontConfig* font_cfg);
func (fa *ImFontAtlas) AddFont(font_cfg *ImFontConfig) *ImFont {
	return (*ImFont)(giLib.Call(_func_ImFontAtlas_AddFont_, []interface{}{&fa, &font_cfg}).PtrFree())
}

// CIMGUI_API ImFont* ImFontAtlas_AddFontDefault(ImFontAtlas* self,const ImFontConfig* font_cfg);
func (fa *ImFontAtlas) AddFontDefault(font_cfg *ImFontConfig) *ImFont {
	return (*ImFont)(giLib.Call(_func_ImFontAtlas_AddFontDefault_, []interface{}{&fa, &font_cfg}).Ptr())
}

// CIMGUI_API ImFont* ImFontAtlas_AddFontFromFileTTF(ImFontAtlas* self,const char* filename,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges);
func (fa *ImFontAtlas) AddFontFromFileTTF(filename string, size_pixels float32, font_cfg *ImFontConfig, glyph_ranges *ImWchar) *ImFont {
	_filename := c.CStr(filename)
	defer usf.Free(_filename)
	return (*ImFont)(giLib.Call(_func_ImFontAtlas_AddFontFromFileTTF_,
		[]interface{}{&fa, &_filename, &size_pixels, &font_cfg, &glyph_ranges}).PtrFree())
}

// CIMGUI_API ImFont* ImFontAtlas_AddFontFromMemoryTTF(ImFontAtlas* self,void* font_data,int font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges);
func (fa *ImFontAtlas) AddFontFromMemoryTTF(font_data unsafe.Pointer, font_size int32, size_pixels float32, font_cfg *ImFontConfig, glyph_ranges *ImWchar) *ImFont {
	return (*ImFont)(giLib.Call(_func_ImFontAtlas_AddFontFromMemoryTTF_,
		[]interface{}{&fa, &font_data, &font_size, &size_pixels, &font_cfg, &glyph_ranges}).PtrFree())
}

// CIMGUI_API ImFont* ImFontAtlas_AddFontFromMemoryCompressedTTF(ImFontAtlas* self,const void* compressed_font_data,int compressed_font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges);
func (fa *ImFontAtlas) AddFontFromMemoryCompressedTTF(font_data unsafe.Pointer, font_size int32, size_pixels float32, font_cfg *ImFontConfig, glyph_ranges *ImWchar) *ImFont {
	return (*ImFont)(giLib.Call(_func_ImFontAtlas_AddFontFromMemoryCompressedTTF_,
		[]interface{}{&fa, &font_data, &font_size, &size_pixels, &font_cfg, &glyph_ranges}).PtrFree())
}

// CIMGUI_API ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas* self,const char* compressed_font_data_base85,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges);
func (fa *ImFontAtlas) AddFontFromMemoryCompressedBase85TTF(font_data_base85 unsafe.Pointer, size_pixels float32, font_cfg *ImFontConfig, glyph_ranges *ImWchar) *ImFont {
	return (*ImFont)(giLib.Call(_func_ImFontAtlas_AddFontFromMemoryCompressedBase85TTF_,
		[]interface{}{&fa, &font_data_base85, &size_pixels, &font_cfg, &glyph_ranges}).PtrFree())
}

var _func_ImFontAtlas_ClearInputData_ = &c.FuncPrototype{Name: "ImFontAtlas_ClearInputData", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFontAtlas_ClearInputData(ImFontAtlas* self);
func (fa *ImFontAtlas) ClearInputData() {
	giLib.Call(_func_ImFontAtlas_ClearInputData_, []interface{}{&fa})
}

var _func_ImFontAtlas_ClearTexData_ = &c.FuncPrototype{Name: "ImFontAtlas_ClearTexData", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFontAtlas_ClearTexData(ImFontAtlas* self);
func (fa *ImFontAtlas) ClearTexData() {
	giLib.Call(_func_ImFontAtlas_ClearTexData_, []interface{}{&fa})
}

var _func_ImFontAtlas_ClearFonts_ = &c.FuncPrototype{Name: "ImFontAtlas_ClearFonts", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFontAtlas_ClearFonts(ImFontAtlas* self);
func (fa *ImFontAtlas) ClearFonts() {
	giLib.Call(_func_ImFontAtlas_ClearFonts_, []interface{}{&fa})
}

var _func_ImFontAtlas_Clear_ = &c.FuncPrototype{Name: "ImFontAtlas_Clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFontAtlas_Clear(ImFontAtlas* self);
func (fa *ImFontAtlas) Clear() {
	giLib.Call(_func_ImFontAtlas_Clear_, []interface{}{&fa})
}

var _func_ImFontAtlas_Build_ = &c.FuncPrototype{Name: "ImFontAtlas_Build", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImFontAtlas_Build(ImFontAtlas* self);
func (fa *ImFontAtlas) Build() bool {
	return giLib.Call(_func_ImFontAtlas_Build_, []interface{}{&fa}).BoolFree()
}

// CIMGUI_API void ImFontAtlas_GetTexDataAsAlpha8(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel);
// CIMGUI_API void ImFontAtlas_GetTexDataAsRGBA32(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel);

var _func_ImFontAtlas_IsBuilt_ = &c.FuncPrototype{Name: "ImFontAtlas_IsBuilt", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImFontAtlas_IsBuilt(ImFontAtlas* self);
func (fa *ImFontAtlas) IsBuilt() bool {
	return giLib.Call(_func_ImFontAtlas_IsBuilt_, []interface{}{&fa}).BoolFree()
}

var _func_ImFontAtlas_SetTexID_ = &c.FuncPrototype{Name: "ImFontAtlas_SetTexID", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImFontAtlas_SetTexID(ImFontAtlas* self,ImTextureID id);
func (fa *ImFontAtlas) SetTexID(id ImTextureID) {
	giLib.Call(_func_ImFontAtlas_SetTexID_, []interface{}{&fa, &id})
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesDefault(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesDefault() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesDefault_, []interface{}{&fa}).PtrFree())
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesGreek(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesGreek() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesGreek_, []interface{}{&fa}).PtrFree())
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesKorean(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesKorean() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesKorean_, []interface{}{&fa}).PtrFree())
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesJapanese(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesJapanese() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesJapanese_, []interface{}{&fa}).PtrFree())
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesChineseFull(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesChineseFull() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesChineseFull_, []interface{}{&fa}).PtrFree())
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesChineseSimplifiedCommon() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon_, []interface{}{&fa}).PtrFree())
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesCyrillic(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesCyrillic() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesCyrillic_, []interface{}{&fa}).PtrFree())
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesThai(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesThai() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesThai_, []interface{}{&fa}).PtrFree())
}

// CIMGUI_API const ImWchar* ImFontAtlas_GetGlyphRangesVietnamese(ImFontAtlas* self);
func (fa *ImFontAtlas) GetGlyphRangesVietnamese() *ImWchar {
	return (*ImWchar)(giLib.Call(_func_ImFontAtlas_GetGlyphRangesVietnamese_, []interface{}{&fa}).PtrFree())
}

var _func_ImFontAtlas_AddCustomRectRegular_ = &c.FuncPrototype{Name: "ImFontAtlas_AddCustomRectRegular", OutType: c.I32, InTypes: _typs_PI32I32}

// CIMGUI_API int ImFontAtlas_AddCustomRectRegular(ImFontAtlas* self,int width,int height);
func (fa *ImFontAtlas) AddCustomRectRegular(width int32, height int32) int32 {
	return giLib.Call(_func_ImFontAtlas_AddCustomRectRegular_, []interface{}{&fa, &width, &height}).I32Free()
}

var _func_ImFontAtlas_AddCustomRectFontGlyph_ = &c.FuncPrototype{Name: "ImFontAtlas_AddCustomRectFontGlyph", OutType: c.I32, InTypes: _typs_PPU16I32I32F32Vec2}

// CIMGUI_API int ImFontAtlas_AddCustomRectFontGlyph(ImFontAtlas* self,ImFont* font,ImWchar id,int width,int height,float advance_x,const ImVec2 offset);
func (fa *ImFontAtlas) AddCustomRectFontGlyph(font *ImFont, id uint16, width int32, height int32, advance_x float32, offset ImVec2) int32 {
	return giLib.Call(_func_ImFontAtlas_AddCustomRectFontGlyph_, []interface{}{&fa, &font, &id, &width, &height, &advance_x, &offset}).I32Free()
}

var _func_ImFontAtlas_GetCustomRectByIndex_ = &c.FuncPrototype{Name: "ImFontAtlas_GetCustomRectByIndex", OutType: c.Pointer, InTypes: _typs_PI32}

// CIMGUI_API ImFontAtlasCustomRect* ImFontAtlas_GetCustomRectByIndex(ImFontAtlas* self,int index);
func (fa *ImFontAtlas) GetCustomRectByIndex(index int32) *ImFontAtlasCustomRect {
	return (*ImFontAtlasCustomRect)(giLib.Call(_func_ImFontAtlas_GetCustomRectByIndex_, []interface{}{&fa, &index}).PtrFree())
}

var _func_ImFontAtlas_CalcCustomRectUV_ = &c.FuncPrototype{Name: "ImFontAtlas_CalcCustomRectUV", OutType: c.Void, InTypes: _typs_PPPP}

// CIMGUI_API void ImFontAtlas_CalcCustomRectUV(ImFontAtlas* self,const ImFontAtlasCustomRect* rect,ImVec2* out_uv_min,ImVec2* out_uv_max);
func (fa *ImFontAtlas) CalcCustomRectUV(rect *ImFontAtlasCustomRect, out_uv_min *ImVec2, out_uv_max *ImVec2) {
	giLib.Call(_func_ImFontAtlas_CalcCustomRectUV_, []interface{}{&fa, &rect, &out_uv_min, &out_uv_max})
}

// CIMGUI_API bool ImFontAtlas_GetMouseCursorTexData(ImFontAtlas* self,ImGuiMouseCursor cursor,ImVec2* out_offset,ImVec2* out_size,ImVec2 out_uv_border[2],ImVec2 out_uv_fill[2]);

var _func_ImFont_ImFont_ = &c.FuncPrototype{Name: "ImFont_ImFont", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImFont* ImFont_ImFont(void);
func NewImFont() *ImFont {
	return (*ImFont)(giLib.Call(_func_ImFont_ImFont_, nil).PtrFree())
}

var _func_ImFont_destroy_ = &c.FuncPrototype{Name: "ImFont_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFont_destroy(ImFont* self);
func (ft *ImFont) Destroy() {
	giLib.Call(_func_ImFont_destroy_, []interface{}{&ft})
}

var _func_ImFont_FindGlyph_ = &c.FuncPrototype{Name: "ImFont_FindGlyph", OutType: c.Pointer, InTypes: _typs_PU16}

// CIMGUI_API const ImFontGlyph* ImFont_FindGlyph(ImFont* self,ImWchar c);
func (ft *ImFont) FindGlyph(c uint16) *ImFontGlyph {
	return (*ImFontGlyph)(giLib.Call(_func_ImFont_FindGlyph_, []interface{}{&ft, &c}).PtrFree())
}

var _func_ImFont_FindGlyphNoFallback_ = &c.FuncPrototype{Name: "ImFont_FindGlyphNoFallback", OutType: c.Pointer, InTypes: _typs_PU16}

// CIMGUI_API const ImFontGlyph* ImFont_FindGlyphNoFallback(ImFont* self,ImWchar c);
func (ft *ImFont) FindGlyphNoFallback(c uint16) *ImFontGlyph {
	return (*ImFontGlyph)(giLib.Call(_func_ImFont_FindGlyphNoFallback_, []interface{}{&ft, &c}).PtrFree())
}

var _func_ImFont_GetCharAdvance_ = &c.FuncPrototype{Name: "ImFont_GetCharAdvance", OutType: c.F32, InTypes: _typs_PU16}

// CIMGUI_API float ImFont_GetCharAdvance(ImFont* self,ImWchar c);
func (ft *ImFont) GetCharAdvance(c uint16) float32 {
	return giLib.Call(_func_ImFont_GetCharAdvance_, []interface{}{&ft, &c}).F32Free()
}

var _func_ImFont_IsLoaded_ = &c.FuncPrototype{Name: "ImFont_IsLoaded", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImFont_IsLoaded(ImFont* self);
func (ft *ImFont) IsLoaded() bool {
	return giLib.Call(_func_ImFont_IsLoaded_, []interface{}{&ft}).BoolFree()
}

var _func_ImFont_GetDebugName_ = &c.FuncPrototype{Name: "ImFont_GetDebugName", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API const char* ImFont_GetDebugName(ImFont* self);
func (ft *ImFont) GetDebugName() string {
	return giLib.Call(_func_ImFont_GetDebugName_, []interface{}{&ft}).StrFree()
}

// CIMGUI_API void ImFont_CalcTextSizeA(ImVec2* pOut,ImFont* self,float size,float max_width,float wrap_width,const char* text_begin,const char* text_end,const char** remaining);

var _func_ImFont_CalcWordWrapPositionA_ = &c.FuncPrototype{Name: "ImFont_CalcWordWrapPositionA", OutType: c.Pointer, InTypes: _typs_PF32PPF32}

// CIMGUI_API const char* ImFont_CalcWordWrapPositionA(ImFont* self,float scale,const char* text,const char* text_end,float wrap_width);
func (ft *ImFont) GCalcWordWrapPositionA(scale float32, text string, text_end string, wrap_width float32) string {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	return giLib.Call(_func_ImFont_CalcWordWrapPositionA_, []interface{}{&ft, &scale, &_text, &_text_end, &wrap_width}).StrFree()
}

var _func_ImFont_RenderChar_ = &c.FuncPrototype{Name: "ImFont_RenderChar", OutType: c.Void, InTypes: _typs_PPF32Vec2U32U16}

// CIMGUI_API void ImFont_RenderChar(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,ImWchar c);
func (ft *ImFont) RenderChar(draw_list *ImDrawList, size float32, pos ImVec2, col uint32, c uint16) {
	giLib.Call(_func_ImFont_RenderChar_, []interface{}{&ft, &draw_list, &size, &pos, &col, &c})
}

var _func_ImFont_RenderText_ = &c.FuncPrototype{Name: "ImFont_RenderText", OutType: c.Void, InTypes: _typs_PPF32Vec2U32Vec4PPF32I32}

// CIMGUI_API void ImFont_RenderText(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,const ImVec4 clip_rect,const char* text_begin,const char* text_end,float wrap_width,bool cpu_fine_clip);
func (ft *ImFont) RenderText(draw_list *ImDrawList, size float32, pos ImVec2, col uint32, clip_rect ImVec4, text_begin string, text_end string, wrap_width float32, cpu_fine_clip bool) {
	_text_begin, _text_end, _cpu_fine_clip := c.CStr(text_begin), c.CStr(text_end), c.CBool(cpu_fine_clip)
	defer usf.Free(_text_begin)
	defer usf.Free(_text_end)
	giLib.Call(_func_ImFont_RenderText_, []interface{}{&ft, &draw_list, &size, &pos, &col, &clip_rect, &_text_begin, &_text_end, &wrap_width, &_cpu_fine_clip})
}

var _func_ImFont_BuildLookupTable_ = &c.FuncPrototype{Name: "ImFont_BuildLookupTable", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFont_BuildLookupTable(ImFont* self);
func (ft *ImFont) BuildLookupTable() {
	giLib.Call(_func_ImFont_BuildLookupTable_, []interface{}{&ft})
}

var _func_ImFont_ClearOutputData_ = &c.FuncPrototype{Name: "ImFont_ClearOutputData", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImFont_ClearOutputData(ImFont* self);
func (ft *ImFont) ClearOutputData() {
	giLib.Call(_func_ImFont_ClearOutputData_, []interface{}{&ft})
}

var _func_ImFont_GrowIndex_ = &c.FuncPrototype{Name: "ImFont_GrowIndex", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void ImFont_GrowIndex(ImFont* self,int new_size);
func (ft *ImFont) GrowIndex(new_size int32) {
	giLib.Call(_func_ImFont_GrowIndex_, []interface{}{&ft, &new_size})
}

var _func_ImFont_AddGlyph_ = &c.FuncPrototype{Name: "ImFont_AddGlyph", OutType: c.Void, InTypes: _typs_PPU16F32F32F32F32F32F32F32F32F32}

// CIMGUI_API void ImFont_AddGlyph(ImFont* self,const ImFontConfig* src_cfg,ImWchar c,float x0,float y0,float x1,float y1,float u0,float v0,float u1,float v1,float advance_x);
func (ft *ImFont) AddGlyph(src_cfg *ImFontConfig, c uint16, x0, y0, x1, y1 float32, u0, v0, u1, v1 float32, advance_x float32) {
	giLib.Call(_func_ImFont_AddGlyph_, []interface{}{&ft, &src_cfg, &c, &x0, &y0, &x1, &y1, &u0, &v0, &u1, &v1, &advance_x})
}

var _func_ImFont_AddRemapChar_ = &c.FuncPrototype{Name: "ImFont_AddRemapChar", OutType: c.Void, InTypes: _typs_PU16U16I32}

// CIMGUI_API void ImFont_AddRemapChar(ImFont* self,ImWchar dst,ImWchar src,bool overwrite_dst);
func (ft *ImFont) AddRemapChar(dst uint16, src uint16, overwrite_dst bool) {
	_overwrite_dst := c.CBool(overwrite_dst)
	giLib.Call(_func_ImFont_AddRemapChar_, []interface{}{&ft, &dst, &src, &_overwrite_dst})
}

var _func_ImFont_SetGlyphVisible_ = &c.FuncPrototype{Name: "ImFont_SetGlyphVisible", OutType: c.Void, InTypes: _typs_PU16I32}

// CIMGUI_API void ImFont_SetGlyphVisible(ImFont* self,ImWchar c,bool visible);
func (ft *ImFont) SetGlyphVisible(cc uint16, visible bool) {
	_visible := c.CBool(visible)
	giLib.Call(_func_ImFont_SetGlyphVisible_, []interface{}{&ft, &cc, &_visible})
}

var _func_ImFont_IsGlyphRangeUnused_ = &c.FuncPrototype{Name: "ImFont_IsGlyphRangeUnused", OutType: c.I32, InTypes: _typs_PU32U32}

// CIMGUI_API bool ImFont_IsGlyphRangeUnused(ImFont* self,uint c_begin,uint c_last);
func (ft *ImFont) IsGlyphRangeUnused(c_begin uint32, c_last uint32) bool {
	return giLib.Call(_func_ImFont_IsGlyphRangeUnused_, []interface{}{&ft, &c_begin, &c_last}).BoolFree()
}

var _func_ImGuiViewport_ImGuiViewport_ = &c.FuncPrototype{Name: "ImGuiViewport_ImGuiViewport", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiViewport* ImGuiViewport_ImGuiViewport(void);
func NewImGuiViewport() *ImGuiViewport {
	return (*ImGuiViewport)(giLib.Call(_func_ImGuiViewport_ImGuiViewport_, nil).PtrFree())
}

var _func_ImGuiViewport_destroy_ = &c.FuncPrototype{Name: "ImGuiViewport_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiViewport_destroy(ImGuiViewport* self);
func (vp *ImGuiViewport) Destroy() {
	giLib.Call(_func_ImGuiViewport_destroy_, []interface{}{&vp})
}

var _func_ImGuiViewport_GetCenter_ = &c.FuncPrototype{Name: "ImGuiViewport_GetCenter", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiViewport_GetCenter(ImVec2* pOut,ImGuiViewport* self);
func (vp *ImGuiViewport) GetCenter(pOut *ImVec2) {
	giLib.Call(_func_ImGuiViewport_GetCenter_, []interface{}{&pOut, &vp})
}

var _func_ImGuiViewport_GetWorkCenter_ = &c.FuncPrototype{Name: "ImGuiViewport_GetWorkCenter", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiViewport_GetWorkCenter(ImVec2* pOut,ImGuiViewport* self);
func (vp *ImGuiViewport) GetWorkCenter(pOut *ImVec2) {
	giLib.Call(_func_ImGuiViewport_GetWorkCenter_, []interface{}{&pOut, &vp})
}

var _func_ImGuiPlatformIO_ImGuiPlatformIO_ = &c.FuncPrototype{Name: "ImGuiPlatformIO_ImGuiPlatformIO", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiPlatformIO* ImGuiPlatformIO_ImGuiPlatformIO(void);
func NewImGuiPlatformIO() *ImGuiPlatformIO {
	return (*ImGuiPlatformIO)(giLib.Call(_func_ImGuiPlatformIO_ImGuiPlatformIO_, nil).PtrFree())
}

var _func_ImGuiPlatformIO_destroy_ = &c.FuncPrototype{Name: "ImGuiPlatformIO_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiPlatformIO_destroy(ImGuiPlatformIO* self);
func (pio *ImGuiPlatformIO) Destroy() {
	giLib.Call(_func_ImGuiPlatformIO_destroy_, []interface{}{&pio})
}

var _func_ImGuiPlatformMonitor_ImGuiPlatformMonitor_ = &c.FuncPrototype{Name: "ImGuiPlatformMonitor_ImGuiPlatformMonitor", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiPlatformMonitor* ImGuiPlatformMonitor_ImGuiPlatformMonitor(void);
func NewImGuiPlatformMonitor() *ImGuiPlatformMonitor {
	return (*ImGuiPlatformMonitor)(giLib.Call(_func_ImGuiPlatformMonitor_ImGuiPlatformMonitor_, nil).PtrFree())
}

var _func_ImGuiPlatformMonitor_destroy_ = &c.FuncPrototype{Name: "ImGuiPlatformMonitor_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiPlatformMonitor_destroy(ImGuiPlatformMonitor* self);
func (pm *ImGuiPlatformMonitor) Dsestroy() {
	giLib.Call(_func_ImGuiPlatformMonitor_destroy_, []interface{}{&pm})
}

var _func_ImGuiPlatformImeData_ImGuiPlatformImeData_ = &c.FuncPrototype{Name: "ImGuiPlatformImeData_ImGuiPlatformImeData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiPlatformImeData* ImGuiPlatformImeData_ImGuiPlatformImeData(void);
func NewImGuiPlatformImeData() *ImGuiPlatformImeData {
	return (*ImGuiPlatformImeData)(giLib.Call(_func_ImGuiPlatformImeData_ImGuiPlatformImeData_, nil).PtrFree())
}

var _func_ImGuiPlatformImeData_destroy_ = &c.FuncPrototype{Name: "ImGuiPlatformImeData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiPlatformImeData_destroy(ImGuiPlatformImeData* self);
func (pd *ImGuiPlatformImeData) Destroy() {
	giLib.Call(_func_ImGuiPlatformImeData_destroy_, []interface{}{&pd})
}

var _func_igGetKeyIndex_ = &c.FuncPrototype{Name: "igGetKeyIndex", OutType: c.U32, InTypes: _typs_U32}

// CIMGUI_API ImGuiKey igGetKeyIndex(ImGuiKey key);
func GetKeyIndex(key ImGuiKey) ImGuiKey {
	return (ImGuiKey)(giLib.Call(_func_igGetKeyIndex_, []interface{}{&key}).U32Free())
}

// CIMGUI_API ImGuiID igImHashData(const void* data,size_t data_size,ImGuiID seed);
// CIMGUI_API ImGuiID igImHashStr(const char* data,size_t data_size,ImGuiID seed);
// CIMGUI_API void igImQsort(void* base,size_t count,size_t size_of_element,int(*compare_func)(void const*,void const*));
// CIMGUI_API ImU32 igImAlphaBlendColors(ImU32 col_a,ImU32 col_b);
// CIMGUI_API bool igImIsPowerOfTwo_Int(int v);
// CIMGUI_API bool igImIsPowerOfTwo_U64(ImU64 v);
// CIMGUI_API int igImUpperPowerOfTwo(int v);
// CIMGUI_API int igImStricmp(const char* str1,const char* str2);
// CIMGUI_API int igImStrnicmp(const char* str1,const char* str2,size_t count);
// CIMGUI_API void igImStrncpy(char* dst,const char* src,size_t count);
// CIMGUI_API char* igImStrdup(const char* str);
// CIMGUI_API char* igImStrdupcpy(char* dst,size_t* p_dst_size,const char* str);
// CIMGUI_API const char* igImStrchrRange(const char* str_begin,const char* str_end,char c);
// CIMGUI_API int igImStrlenW(const ImWchar* str);
// CIMGUI_API const char* igImStreolRange(const char* str,const char* str_end);
// CIMGUI_API const ImWchar* igImStrbolW(const ImWchar* buf_mid_line,const ImWchar* buf_begin);
// CIMGUI_API const char* igImStristr(const char* haystack,const char* haystack_end,const char* needle,const char* needle_end);
// CIMGUI_API void igImStrTrimBlanks(char* str);
// CIMGUI_API const char* igImStrSkipBlank(const char* str);
// CIMGUI_API char igImToUpper(char c);
// CIMGUI_API bool igImCharIsBlankA(char c);
// CIMGUI_API bool igImCharIsBlankW(unsigned int c);
// CIMGUI_API int igImFormatString(char* buf,size_t buf_size,const char* fmt,...);
// CIMGUI_API int igImFormatStringV(char* buf,size_t buf_size,const char* fmt,va_list args);
// CIMGUI_API void igImFormatStringToTempBuffer(const char** out_buf,const char** out_buf_end,const char* fmt,...);
// CIMGUI_API void igImFormatStringToTempBufferV(const char** out_buf,const char** out_buf_end,const char* fmt,va_list args);
// CIMGUI_API const char* igImParseFormatFindStart(const char* format);
// CIMGUI_API const char* igImParseFormatFindEnd(const char* format);
// CIMGUI_API const char* igImParseFormatTrimDecorations(const char* format,char* buf,size_t buf_size);
// CIMGUI_API void igImParseFormatSanitizeForPrinting(const char* fmt_in,char* fmt_out,size_t fmt_out_size);
// CIMGUI_API const char* igImParseFormatSanitizeForScanning(const char* fmt_in,char* fmt_out,size_t fmt_out_size);
// CIMGUI_API int igImParseFormatPrecision(const char* format,int default_value);
// CIMGUI_API const char* igImTextCharToUtf8(char out_buf[5],unsigned int c);
// CIMGUI_API int igImTextStrToUtf8(char* out_buf,int out_buf_size,const ImWchar* in_text,const ImWchar* in_text_end);
// CIMGUI_API int igImTextCharFromUtf8(unsigned int* out_char,const char* in_text,const char* in_text_end);
// CIMGUI_API int igImTextStrFromUtf8(ImWchar* out_buf,int out_buf_size,const char* in_text,const char* in_text_end,const char** in_remaining);
// CIMGUI_API int igImTextCountCharsFromUtf8(const char* in_text,const char* in_text_end);
// CIMGUI_API int igImTextCountUtf8BytesFromChar(const char* in_text,const char* in_text_end);
// CIMGUI_API int igImTextCountUtf8BytesFromStr(const ImWchar* in_text,const ImWchar* in_text_end);
// CIMGUI_API ImFileHandle igImFileOpen(const char* filename,const char* mode);
// CIMGUI_API bool igImFileClose(ImFileHandle file);
// CIMGUI_API ImU64 igImFileGetSize(ImFileHandle file);
// CIMGUI_API ImU64 igImFileRead(void* data,ImU64 size,ImU64 count,ImFileHandle file);
// CIMGUI_API ImU64 igImFileWrite(const void* data,ImU64 size,ImU64 count,ImFileHandle file);
// CIMGUI_API void* igImFileLoadToMemory(const char* filename,const char* mode,size_t* out_file_size,int padding_bytes);
// CIMGUI_API float igImPow_Float(float x,float y);
// CIMGUI_API double igImPow_double(double x,double y);
// CIMGUI_API float igImLog_Float(float x);
// CIMGUI_API double igImLog_double(double x);
// CIMGUI_API int igImAbs_Int(int x);
// CIMGUI_API float igImAbs_Float(float x);
// CIMGUI_API double igImAbs_double(double x);
// CIMGUI_API float igImSign_Float(float x);
// CIMGUI_API double igImSign_double(double x);
// CIMGUI_API float igImRsqrt_Float(float x);
// CIMGUI_API double igImRsqrt_double(double x);
// CIMGUI_API void igImMin(ImVec2* pOut,const ImVec2 lhs,const ImVec2 rhs);
// CIMGUI_API void igImMax(ImVec2* pOut,const ImVec2 lhs,const ImVec2 rhs);
// CIMGUI_API void igImClamp(ImVec2* pOut,const ImVec2 v,const ImVec2 mn,ImVec2 mx);
// CIMGUI_API void igImLerp_Vec2Float(ImVec2* pOut,const ImVec2 a,const ImVec2 b,float t);
// CIMGUI_API void igImLerp_Vec2Vec2(ImVec2* pOut,const ImVec2 a,const ImVec2 b,const ImVec2 t);
// CIMGUI_API void igImLerp_Vec4(ImVec4* pOut,const ImVec4 a,const ImVec4 b,float t);
// CIMGUI_API float igImSaturate(float f);
// CIMGUI_API float igImLengthSqr_Vec2(const ImVec2 lhs);
// CIMGUI_API float igImLengthSqr_Vec4(const ImVec4 lhs);
// CIMGUI_API float igImInvLength(const ImVec2 lhs,float fail_value);
// CIMGUI_API float igImFloor_Float(float f);
// CIMGUI_API float igImFloorSigned_Float(float f);
// CIMGUI_API void igImFloor_Vec2(ImVec2* pOut,const ImVec2 v);
// CIMGUI_API void igImFloorSigned_Vec2(ImVec2* pOut,const ImVec2 v);
// CIMGUI_API int igImModPositive(int a,int b);
// CIMGUI_API float igImDot(const ImVec2 a,const ImVec2 b);
// CIMGUI_API void igImRotate(ImVec2* pOut,const ImVec2 v,float cos_a,float sin_a);
// CIMGUI_API float igImLinearSweep(float current,float target,float speed);
// CIMGUI_API void igImMul(ImVec2* pOut,const ImVec2 lhs,const ImVec2 rhs);
// CIMGUI_API bool igImIsFloatAboveGuaranteedIntegerPrecision(float f);
// CIMGUI_API float igImExponentialMovingAverage(float avg,float sample,int n);
// CIMGUI_API void igImBezierCubicCalc(ImVec2* pOut,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,float t);
// CIMGUI_API void igImBezierCubicClosestPoint(ImVec2* pOut,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 p,int num_segments);
// CIMGUI_API void igImBezierCubicClosestPointCasteljau(ImVec2* pOut,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 p,float tess_tol);
// CIMGUI_API void igImBezierQuadraticCalc(ImVec2* pOut,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,float t);
// CIMGUI_API void igImLineClosestPoint(ImVec2* pOut,const ImVec2 a,const ImVec2 b,const ImVec2 p);
// CIMGUI_API bool igImTriangleContainsPoint(const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 p);
// CIMGUI_API void igImTriangleClosestPoint(ImVec2* pOut,const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 p);
// CIMGUI_API void igImTriangleBarycentricCoords(const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 p,float* out_u,float* out_v,float* out_w);
// CIMGUI_API float igImTriangleArea(const ImVec2 a,const ImVec2 b,const ImVec2 c);

var _func_ImVec1_ImVec1_Nil_ = &c.FuncPrototype{Name: "ImVec1_ImVec1_Nil", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImVec1* ImVec1_ImVec1_Nil(void);
func NewImVec1() *ImVec1 {
	return (*ImVec1)(giLib.Call(_func_ImVec1_ImVec1_Nil_, nil).PtrFree())
}

var _func_ImVec1_destroy_ = &c.FuncPrototype{Name: "ImVec1_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImVec1_destroy(ImVec1* self);
func (v1 *ImVec1) Destroy(self *ImVec1) {
	giLib.Call(_func_ImVec1_destroy_, []interface{}{&v1})
}

var _func_ImVec1_ImVec1_Float_ = &c.FuncPrototype{Name: "ImVec1_ImVec1_Float", OutType: c.Pointer, InTypes: _typs_F32}

// CIMGUI_API ImVec1* ImVec1_ImVec1_Float(float _x);
func NewImVec1Float(_x float32) *ImVec1 {
	return (*ImVec1)(giLib.Call(_func_ImVec1_ImVec1_Float_, []interface{}{&_x}).PtrFree())
}

var _func_ImVec2ih_ImVec2ih_Nil_ = &c.FuncPrototype{Name: "ImVec2ih_ImVec2ih_Nil", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImVec2ih* ImVec2ih_ImVec2ih_Nil(void);
func NewImVec2ih() *ImVec2ih {
	return (*ImVec2ih)(giLib.Call(_func_ImVec2ih_ImVec2ih_Nil_, nil).PtrFree())
}

var _func_ImVec2ih_destroy_ = &c.FuncPrototype{Name: "ImVec2ih_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImVec2ih_destroy(ImVec2ih* self);
func (v2 *ImVec2ih) Destroy() {
	giLib.Call(_func_ImVec2ih_destroy_, []interface{}{&v2})
}

var _func_ImVec2ih_ImVec2ih_short_ = &c.FuncPrototype{Name: "ImVec2ih_ImVec2ih_short", OutType: c.Pointer, InTypes: _typs_I16I16}

// CIMGUI_API ImVec2ih* ImVec2ih_ImVec2ih_short(short _x,short _y);
func NewImVec2ihShort(_x int16, _y int16) *ImVec2ih {
	return (*ImVec2ih)(giLib.Call(_func_ImVec2ih_ImVec2ih_short_, []interface{}{&_x, &_y}).PtrFree())
}

var _func_ImVec2ih_ImVec2ih_Vec2_ = &c.FuncPrototype{Name: "ImVec2ih_ImVec2ih_Vec2", OutType: c.Pointer, InTypes: _typs_Vec2}

// CIMGUI_API ImVec2ih* ImVec2ih_ImVec2ih_Vec2(const ImVec2 rhs);
func NewImVec2ihVec2(rhs ImVec2) *ImVec2ih {
	return (*ImVec2ih)(giLib.Call(_func_ImVec2ih_ImVec2ih_Vec2_, []interface{}{&rhs}).PtrFree())
}

var _func_ImRect_ImRect_Nil_ = &c.FuncPrototype{Name: "ImRect_ImRect_Nil", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImRect* ImRect_ImRect_Nil(void);
func NewImRect() *ImRect {
	return (*ImRect)(giLib.Call(_func_ImRect_ImRect_Nil_, nil).PtrFree())
}

var _func_ImRect_destroy_ = &c.FuncPrototype{Name: "ImRect_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImRect_destroy(ImRect* self);
func (rct *ImRect) Destroy() {
	giLib.Call(_func_ImRect_destroy_, []interface{}{&rct})
}

var _func_ImRect_ImRect_Vec2_ = &c.FuncPrototype{Name: "ImRect_ImRect_Vec2", OutType: c.Pointer, InTypes: _typs_Vec2Vec2}

// CIMGUI_API ImRect* ImRect_ImRect_Vec2(const ImVec2 min,const ImVec2 max);
func NewImRectVec2(min ImVec2, max ImVec2) *ImRect {
	return (*ImRect)(giLib.Call(_func_ImRect_ImRect_Vec2_, []interface{}{&min, &max}).PtrFree())
}

var _func_ImRect_ImRect_Vec4_ = &c.FuncPrototype{Name: "ImRect_ImRect_Vec4", OutType: c.Pointer, InTypes: _typs_Vec4}

// CIMGUI_API ImRect* ImRect_ImRect_Vec4(const ImVec4 v);
func NewImRectVec4(v ImVec4) *ImRect {
	return (*ImRect)(giLib.Call(_func_ImRect_ImRect_Vec4_, []interface{}{&v}).PtrFree())
}

var _func_ImRect_ImRect_Float_ = &c.FuncPrototype{Name: "ImRect_ImRect_Float", OutType: c.Pointer, InTypes: _typs_F32F32F32F32}

// CIMGUI_API ImRect* ImRect_ImRect_Float(float x1,float y1,float x2,float y2);
func NewImRectFloat(x1 float32, y1 float32, x2 float32, y2 float32) *ImRect {
	return (*ImRect)(giLib.Call(_func_ImRect_ImRect_Float_, []interface{}{&x1, &y1, &x2, &y2}).PtrFree())
}

var _func_ImRect_GetCenter_ = &c.FuncPrototype{Name: "ImRect_GetCenter", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImRect_GetCenter(ImVec2* pOut,ImRect* self);
func (rct *ImRect) GetCenter(pOut *ImVec2) {
	giLib.Call(_func_ImRect_GetCenter_, []interface{}{&pOut, &rct})
}

var _func_ImRect_GetSize_ = &c.FuncPrototype{Name: "ImRect_GetSize", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImRect_GetSize(ImVec2* pOut,ImRect* self);
func (rct *ImRect) GetSize(pOut *ImVec2) {
	giLib.Call(_func_ImRect_GetSize_, []interface{}{&pOut, &rct})
}

var _func_ImRect_GetWidth_ = &c.FuncPrototype{Name: "ImRect_GetWidth", OutType: c.F32, InTypes: _typs_P}

// CIMGUI_API float ImRect_GetWidth(ImRect* self);
func (rct *ImRect) GetWidth() float32 {
	return giLib.Call(_func_ImRect_GetWidth_, []interface{}{&rct}).F32Free()
}

var _func_ImRect_GetHeight_ = &c.FuncPrototype{Name: "ImRect_GetHeight", OutType: c.F32, InTypes: _typs_P}

// CIMGUI_API float ImRect_GetHeight(ImRect* self);
func (rct *ImRect) GetHeight() float32 {
	return giLib.Call(_func_ImRect_GetHeight_, []interface{}{&rct}).F32Free()
}

var _func_ImRect_GetArea_ = &c.FuncPrototype{Name: "ImRect_GetArea", OutType: c.F32, InTypes: _typs_P}

// CIMGUI_API float ImRect_GetArea(ImRect* self);
func (rct *ImRect) GetArea() float32 {
	return giLib.Call(_func_ImRect_GetArea_, []interface{}{&rct}).F32Free()
}

var _func_ImRect_GetTL_ = &c.FuncPrototype{Name: "ImRect_GetTL", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImRect_GetTL(ImVec2* pOut,ImRect* self);
func (rct *ImRect) GetTL(pOut *ImVec2) {
	giLib.Call(_func_ImRect_GetTL_, []interface{}{&pOut, &rct})
}

var _func_ImRect_GetTR_ = &c.FuncPrototype{Name: "ImRect_GetTR", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImRect_GetTR(ImVec2* pOut,ImRect* self);
func (rct *ImRect) GetTR(pOut *ImVec2) {
	giLib.Call(_func_ImRect_GetTR_, []interface{}{&pOut, &rct})
}

var _func_ImRect_GetBL_ = &c.FuncPrototype{Name: "ImRect_GetBL", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImRect_GetBL(ImVec2* pOut,ImRect* self);
func (rct *ImRect) GetBL(pOut *ImVec2) {
	giLib.Call(_func_ImRect_GetBL_, []interface{}{&pOut, &rct})
}

var _func_ImRect_GetBR_ = &c.FuncPrototype{Name: "ImRect_GetBR", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImRect_GetBR(ImVec2* pOut,ImRect* self);
func (rct *ImRect) GetBR(pOut *ImVec2) {
	giLib.Call(_func_ImRect_GetBR_, []interface{}{&pOut, &rct})
}

var _func_ImRect_Contains_Vec2_ = &c.FuncPrototype{Name: "ImRect_Contains_Vec2", OutType: c.I32, InTypes: _typs_PVec2}

// CIMGUI_API bool ImRect_Contains_Vec2(ImRect* self,const ImVec2 p);
func (rct *ImRect) ContainsVec2(p ImVec2) bool {
	return giLib.Call(_func_ImRect_Contains_Vec2_, []interface{}{&rct, &p}).BoolFree()
}

var _func_ImRect_Contains_Rect_ = &c.FuncPrototype{Name: "ImRect_Contains_Rect", OutType: c.I32, InTypes: _typs_PRect}

// CIMGUI_API bool ImRect_Contains_Rect(ImRect* self,const ImRect r);
func (rct *ImRect) ContainsRect(r ImRect) bool {
	return giLib.Call(_func_ImRect_Contains_Rect_, []interface{}{&rct, &r}).BoolFree()
}

var _func_ImRect_Overlaps_ = &c.FuncPrototype{Name: "ImRect_Overlaps", OutType: c.I32, InTypes: _typs_PRect}

// CIMGUI_API bool ImRect_Overlaps(ImRect* self,const ImRect r);
func (rct *ImRect) Overlaps(r ImRect) bool {
	return giLib.Call(_func_ImRect_Overlaps_, []interface{}{&rct, &r}).BoolFree()
}

var _func_ImRect_Add_Vec2_ = &c.FuncPrototype{Name: "ImRect_Add_Vec2", OutType: c.Void, InTypes: _typs_PVec2}

// CIMGUI_API void ImRect_Add_Vec2(ImRect* self,const ImVec2 p);
func (rct *ImRect) AddVec2(p ImVec2) {
	giLib.Call(_func_ImRect_Add_Vec2_, []interface{}{&rct, &p})
}

var _func_ImRect_Add_Rect_ = &c.FuncPrototype{Name: "ImRect_Add_Rect", OutType: c.Void, InTypes: _typs_PRect}

// CIMGUI_API void ImRect_Add_Rect(ImRect* self,const ImRect r);
func (rct *ImRect) AddRect(r ImRect) {
	giLib.Call(_func_ImRect_Add_Rect_, []interface{}{&rct, &r})
}

var _func_ImRect_Expand_Float_ = &c.FuncPrototype{Name: "ImRect_Expand_Float", OutType: c.Void, InTypes: _typs_PF32}

// CIMGUI_API void ImRect_Expand_Float(ImRect* self,const float amount);
func (rct *ImRect) ExpandFloat(amount float32) {
	giLib.Call(_func_ImRect_Expand_Float_, []interface{}{&rct, &amount})
}

var _func_ImRect_Expand_Vec2_ = &c.FuncPrototype{Name: "ImRect_Expand_Vec2", OutType: c.Void, InTypes: _typs_PVec2}

// CIMGUI_API void ImRect_Expand_Vec2(ImRect* self,const ImVec2 amount);
func (rct *ImRect) ExpandVec2(amount ImVec2) {
	giLib.Call(_func_ImRect_Expand_Vec2_, []interface{}{&rct, &amount})
}

var _func_ImRect_Translate_ = &c.FuncPrototype{Name: "ImRect_Translate", OutType: c.Void, InTypes: _typs_PVec2}

// CIMGUI_API void ImRect_Translate(ImRect* self,const ImVec2 d);
func (rct *ImRect) Translate(d ImVec2) {
	giLib.Call(_func_ImRect_Translate_, []interface{}{&rct, &d})
}

var _func_ImRect_TranslateX_ = &c.FuncPrototype{Name: "ImRect_TranslateX", OutType: c.Void, InTypes: _typs_PF32}

// CIMGUI_API void ImRect_TranslateX(ImRect* self,float dx);
func (rct *ImRect) TranslateX(dx float32) {
	giLib.Call(_func_ImRect_TranslateX_, []interface{}{&rct, &dx})
}

var _func_ImRect_TranslateY_ = &c.FuncPrototype{Name: "ImRect_TranslateY", OutType: c.Void, InTypes: _typs_PF32}

// CIMGUI_API void ImRect_TranslateY(ImRect* self,float dy);
func (rct *ImRect) TranslateY(dy float32) {
	giLib.Call(_func_ImRect_TranslateY_, []interface{}{&rct, &dy})
}

var _func_ImRect_ClipWith_ = &c.FuncPrototype{Name: "ImRect_ClipWith", OutType: c.Void, InTypes: _typs_PRect}

// CIMGUI_API void ImRect_ClipWith(ImRect* self,const ImRect r);
func (rct *ImRect) ClipWith(r ImRect) {
	giLib.Call(_func_ImRect_ClipWith_, []interface{}{&rct, &r})
}

var _func_ImRect_ClipWithFull_ = &c.FuncPrototype{Name: "ImRect_ClipWithFull", OutType: c.Void, InTypes: _typs_PRect}

// CIMGUI_API void ImRect_ClipWithFull(ImRect* self,const ImRect r);
func (rct *ImRect) ClipWithFull(r ImRect) {
	giLib.Call(_func_ImRect_ClipWithFull_, []interface{}{&rct, &r})
}

var _func_ImRect_Floor_ = &c.FuncPrototype{Name: "ImRect_Floor", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImRect_Floor(ImRect* self);
func (rct *ImRect) Floor() {
	giLib.Call(_func_ImRect_Floor_, []interface{}{&rct})
}

var _func_ImRect_IsInverted_ = &c.FuncPrototype{Name: "ImRect_IsInverted", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImRect_IsInverted(ImRect* self);
func (rct *ImRect) IsInverted() bool {
	return giLib.Call(_func_ImRect_IsInverted_, []interface{}{&rct}).BoolFree()
}

var _func_ImRect_ToVec4_ = &c.FuncPrototype{Name: "ImRect_ToVec4", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImRect_ToVec4(ImVec4* pOut,ImRect* self);
func (rct *ImRect) ToVec4(pOut *ImVec4) {
	giLib.Call(_func_ImRect_ToVec4_, []interface{}{&pOut, &rct})
}

// CIMGUI_API size_t igImBitArrayGetStorageSizeInBytes(int bitcount);
// CIMGUI_API void igImBitArrayClearAllBits(ImU32* arr,int bitcount);
// CIMGUI_API bool igImBitArrayTestBit(const ImU32* arr,int n);
// CIMGUI_API void igImBitArrayClearBit(ImU32* arr,int n);
// CIMGUI_API void igImBitArraySetBit(ImU32* arr,int n);
// CIMGUI_API void igImBitArraySetBitRange(ImU32* arr,int n,int n2);
// CIMGUI_API void ImBitVector_Create(ImBitVector* self,int sz);
// CIMGUI_API void ImBitVector_Clear(ImBitVector* self);
// CIMGUI_API bool ImBitVector_TestBit(ImBitVector* self,int n);
// CIMGUI_API void ImBitVector_SetBit(ImBitVector* self,int n);
// CIMGUI_API void ImBitVector_ClearBit(ImBitVector* self,int n);
// CIMGUI_API void ImGuiTextIndex_clear(ImGuiTextIndex* self);
// CIMGUI_API int ImGuiTextIndex_size(ImGuiTextIndex* self);
// CIMGUI_API const char* ImGuiTextIndex_get_line_begin(ImGuiTextIndex* self,const char* base,int n);
// CIMGUI_API const char* ImGuiTextIndex_get_line_end(ImGuiTextIndex* self,const char* base,int n);
// CIMGUI_API void ImGuiTextIndex_append(ImGuiTextIndex* self,const char* base,int old_size,int new_size);

var _func_ImDrawListSharedData_ImDrawListSharedData_ = &c.FuncPrototype{Name: "ImDrawListSharedData_ImDrawListSharedData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImDrawListSharedData* ImDrawListSharedData_ImDrawListSharedData(void);
func NewImDrawListSharedData() *ImDrawListSharedData {
	return (*ImDrawListSharedData)(giLib.Call(_func_ImDrawListSharedData_ImDrawListSharedData_, nil).PtrFree())
}

var _func_ImDrawListSharedData_destroy_ = &c.FuncPrototype{Name: "ImDrawListSharedData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawListSharedData_destroy(ImDrawListSharedData* self);
func (dsd *ImDrawListSharedData) Destroy() {
	giLib.Call(_func_ImDrawListSharedData_destroy_, []interface{}{&dsd})
}

var _func_ImDrawListSharedData_SetCircleTessellationMaxError_ = &c.FuncPrototype{Name: "ImDrawListSharedData_SetCircleTessellationMaxError", OutType: c.Void, InTypes: _typs_PF32}

// CIMGUI_API void ImDrawListSharedData_SetCircleTessellationMaxError(ImDrawListSharedData* self,float max_error);
func (dsd *ImDrawListSharedData) SetCircleTessellationMaxError(max_error float32) {
	giLib.Call(_func_ImDrawListSharedData_SetCircleTessellationMaxError_, []interface{}{&dsd, &max_error})
}

var _func_ImDrawDataBuilder_Clear_ = &c.FuncPrototype{Name: "ImDrawDataBuilder_Clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawDataBuilder_Clear(ImDrawDataBuilder* self);
func (ddb *ImDrawDataBuilder) Clear() {
	giLib.Call(_func_ImDrawDataBuilder_Clear_, []interface{}{&ddb})
}

var _func_ImDrawDataBuilder_ClearFreeMemory_ = &c.FuncPrototype{Name: "ImDrawDataBuilder_ClearFreeMemory", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawDataBuilder_ClearFreeMemory(ImDrawDataBuilder* self);
func (ddb *ImDrawDataBuilder) ClearFreeMemory() {
	giLib.Call(_func_ImDrawDataBuilder_ClearFreeMemory_, []interface{}{&ddb})
}

var _func_ImDrawDataBuilder_GetDrawListCount_ = &c.FuncPrototype{Name: "ImDrawDataBuilder_GetDrawListCount", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int ImDrawDataBuilder_GetDrawListCount(ImDrawDataBuilder* self);
func (ddb *ImDrawDataBuilder) GetDrawListCount() int32 {
	return giLib.Call(_func_ImDrawDataBuilder_GetDrawListCount_, []interface{}{&ddb}).I32Free()
}

var _func_ImDrawDataBuilder_FlattenIntoSingleLayer_ = &c.FuncPrototype{Name: "ImDrawDataBuilder_FlattenIntoSingleLayer", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImDrawDataBuilder_FlattenIntoSingleLayer(ImDrawDataBuilder* self);
func (ddb *ImDrawDataBuilder) FlattenIntoSingleLayer() {
	giLib.Call(_func_ImDrawDataBuilder_FlattenIntoSingleLayer_, []interface{}{&ddb})
}

var _func_ImGuiDataVarInfo_GetVarPtr_ = &c.FuncPrototype{Name: "ImGuiDataVarInfo_GetVarPtr", OutType: c.Pointer, InTypes: _typs_PP}

// CIMGUI_API void* ImGuiDataVarInfo_GetVarPtr(ImGuiDataVarInfo* self,void* parent);
func (dvi *ImGuiDataVarInfo) GetVarPtr(parent unsafe.Pointer) unsafe.Pointer {
	return giLib.Call(_func_ImGuiDataVarInfo_GetVarPtr_, []interface{}{&dvi, &parent}).PtrFree()
}

var _func_ImGuiStyleMod_ImGuiStyleMod_Int_ = &c.FuncPrototype{Name: "ImGuiStyleMod_ImGuiStyleMod_Int", OutType: c.Pointer, InTypes: _typs_U32I32}

// CIMGUI_API ImGuiStyleMod* ImGuiStyleMod_ImGuiStyleMod_Int(ImGuiStyleVar idx,int v);
func NewImGuiStyleMod_Int(idx ImGuiStyleVar, v int32) *ImGuiStyleMod {
	return (*ImGuiStyleMod)(giLib.Call(_func_ImGuiStyleMod_ImGuiStyleMod_Int_, []interface{}{&idx, &v}).PtrFree())
}

var _func_ImGuiStyleMod_destroy_ = &c.FuncPrototype{Name: "ImGuiStyleMod_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiStyleMod_destroy(ImGuiStyleMod* self);
func (sm *ImGuiStyleMod) Destroy() {
	giLib.Call(_func_ImGuiStyleMod_destroy_, []interface{}{&sm})
}

var _func_ImGuiStyleMod_ImGuiStyleMod_Float_ = &c.FuncPrototype{Name: "ImGuiStyleMod_ImGuiStyleMod_Float", OutType: c.Pointer, InTypes: _typs_U32F32}

// CIMGUI_API ImGuiStyleMod* ImGuiStyleMod_ImGuiStyleMod_Float(ImGuiStyleVar idx,float v);
func NewImGuiStyleModFloat(idx ImGuiStyleVar, v float32) *ImGuiStyleMod {
	return (*ImGuiStyleMod)(giLib.Call(_func_ImGuiStyleMod_ImGuiStyleMod_Float_, []interface{}{&idx, &v}).PtrFree())
}

var _func_ImGuiStyleMod_ImGuiStyleMod_Vec2_ = &c.FuncPrototype{Name: "ImGuiStyleMod_ImGuiStyleMod_Vec2", OutType: c.Pointer, InTypes: _typs_U32Vec2}

// CIMGUI_API ImGuiStyleMod* ImGuiStyleMod_ImGuiStyleMod_Vec2(ImGuiStyleVar idx,ImVec2 v);
func NewImGuiStyleModVec2(idx ImGuiStyleVar, v ImVec2) *ImGuiStyleMod {
	return (*ImGuiStyleMod)(giLib.Call(_func_ImGuiStyleMod_ImGuiStyleMod_Vec2_, []interface{}{&idx, &v}).PtrFree())
}

var _func_ImGuiComboPreviewData_ImGuiComboPreviewData_ = &c.FuncPrototype{Name: "ImGuiComboPreviewData_ImGuiComboPreviewData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiComboPreviewData* ImGuiComboPreviewData_ImGuiComboPreviewData(void);
func NewImGuiComboPreviewData() *ImGuiComboPreviewData {
	return (*ImGuiComboPreviewData)(giLib.Call(_func_ImGuiComboPreviewData_ImGuiComboPreviewData_, nil).PtrFree())
}

var _func_ImGuiComboPreviewData_destroy_ = &c.FuncPrototype{Name: "ImGuiComboPreviewData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiComboPreviewData_destroy(ImGuiComboPreviewData* self);
func (cpd *ImGuiComboPreviewData) Destroy() {
	giLib.Call(_func_ImGuiComboPreviewData_destroy_, []interface{}{&cpd})
}

var _func_ImGuiMenuColumns_ImGuiMenuColumns_ = &c.FuncPrototype{Name: "ImGuiMenuColumns_ImGuiMenuColumns", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiMenuColumns* ImGuiMenuColumns_ImGuiMenuColumns(void);
func NewImGuiMenuColumns() *ImGuiMenuColumns {
	return (*ImGuiMenuColumns)(giLib.Call(_func_ImGuiMenuColumns_ImGuiMenuColumns_, nil).PtrFree())
}

var _func_ImGuiMenuColumns_destroy_ = &c.FuncPrototype{Name: "ImGuiMenuColumns_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiMenuColumns_destroy(ImGuiMenuColumns* self);
func (mc *ImGuiMenuColumns) Destroy() {
	giLib.Call(_func_ImGuiMenuColumns_destroy_, []interface{}{&mc})
}

var _func_ImGuiMenuColumns_Update_ = &c.FuncPrototype{Name: "ImGuiMenuColumns_Update", OutType: c.Void, InTypes: _typs_PF32I32}

// CIMGUI_API void ImGuiMenuColumns_Update(ImGuiMenuColumns* self,float spacing,bool window_reappearing);
func (mc *ImGuiMenuColumns) Update(spacing float32, window_reappearing bool) {
	_window_reappearing := c.CBool(window_reappearing)
	giLib.Call(_func_ImGuiMenuColumns_Update_, []interface{}{&mc, &spacing, &_window_reappearing})
}

var _func_ImGuiMenuColumns_DeclColumns_ = &c.FuncPrototype{Name: "ImGuiMenuColumns_DeclColumns", OutType: c.F32, InTypes: _typs_PF32F32F32F32}

// CIMGUI_API float ImGuiMenuColumns_DeclColumns(ImGuiMenuColumns* self,float w_icon,float w_label,float w_shortcut,float w_mark);
func (mc *ImGuiMenuColumns) DeclColumns(w_icon, w_label, w_shortcut, w_mark float32) float32 {
	return giLib.Call(_func_ImGuiMenuColumns_DeclColumns_, []interface{}{&mc, &w_icon, &w_label, &w_shortcut, &w_mark}).F32Free()
}

var _func_ImGuiMenuColumns_CalcNextTotalWidth_ = &c.FuncPrototype{Name: "ImGuiMenuColumns_CalcNextTotalWidth", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void ImGuiMenuColumns_CalcNextTotalWidth(ImGuiMenuColumns* self,bool update_offsets);
func (mc *ImGuiMenuColumns) CalcNextTotalWidth(update_offsets bool) {
	_update_offsets := c.CBool(update_offsets)
	giLib.Call(_func_ImGuiMenuColumns_CalcNextTotalWidth_, []interface{}{&mc, &_update_offsets})
}

var _func_ImGuiInputTextDeactivatedState_ImGuiInputTextDeactivatedState_ = &c.FuncPrototype{Name: "ImGuiInputTextDeactivatedState_ImGuiInputTextDeactivatedState", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiInputTextDeactivatedState* ImGuiInputTextDeactivatedState_ImGuiInputTextDeactivatedState(void);
func NewImGuiInputTextDeactivatedState() *ImGuiInputTextDeactivatedState {
	return (*ImGuiInputTextDeactivatedState)(giLib.Call(_func_ImGuiInputTextDeactivatedState_ImGuiInputTextDeactivatedState_, nil).PtrFree())
}

var _func_ImGuiInputTextDeactivatedState_destroy_ = &c.FuncPrototype{Name: "ImGuiInputTextDeactivatedState_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextDeactivatedState_destroy(ImGuiInputTextDeactivatedState* self);
func (itds *ImGuiInputTextDeactivatedState) Destroy() {
	giLib.Call(_func_ImGuiInputTextDeactivatedState_destroy_, []interface{}{&itds})
}

var _func_ImGuiInputTextDeactivatedState_ClearFreeMemory_ = &c.FuncPrototype{Name: "ImGuiInputTextDeactivatedState_ClearFreeMemory", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextDeactivatedState_ClearFreeMemory(ImGuiInputTextDeactivatedState* self);
func (itds *ImGuiInputTextDeactivatedState) ClearFreeMemory() {
	giLib.Call(_func_ImGuiInputTextDeactivatedState_ClearFreeMemory_, []interface{}{&itds})
}

var _func_ImGuiInputTextState_ImGuiInputTextState_ = &c.FuncPrototype{Name: "ImGuiInputTextState_ImGuiInputTextState", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiInputTextState* ImGuiInputTextState_ImGuiInputTextState(void);
func NewImGuiInputTextState() *ImGuiInputTextState {
	return (*ImGuiInputTextState)(giLib.Call(_func_ImGuiInputTextState_ImGuiInputTextState_, nil).PtrFree())
}

var _func_ImGuiInputTextState_destroy_ = &c.FuncPrototype{Name: "ImGuiInputTextState_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextState_destroy(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) Destroy() {
	giLib.Call(_func_ImGuiInputTextState_destroy_, []interface{}{&its})
}

var _func_ImGuiInputTextState_ClearText_ = &c.FuncPrototype{Name: "ImGuiInputTextState_ClearText", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextState_ClearText(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) ClearText() {
	giLib.Call(_func_ImGuiInputTextState_ClearText_, []interface{}{&its})
}

var _func_ImGuiInputTextState_ClearFreeMemory_ = &c.FuncPrototype{Name: "ImGuiInputTextState_ClearFreeMemory", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextState_ClearFreeMemory(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) ClearFreeMemory() {
	giLib.Call(_func_ImGuiInputTextState_ClearFreeMemory_, []interface{}{&its})
}

var _func_ImGuiInputTextState_GetUndoAvailCount_ = &c.FuncPrototype{Name: "ImGuiInputTextState_GetUndoAvailCount", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int ImGuiInputTextState_GetUndoAvailCount(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) GetUndoAvailCount() int32 {
	return giLib.Call(_func_ImGuiInputTextState_GetUndoAvailCount_, []interface{}{&its}).I32Free()
}

var _func_ImGuiInputTextState_GetRedoAvailCount_ = &c.FuncPrototype{Name: "ImGuiInputTextState_GetRedoAvailCount", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int ImGuiInputTextState_GetRedoAvailCount(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) GetRedoAvailCount() int32 {
	return giLib.Call(_func_ImGuiInputTextState_GetRedoAvailCount_, []interface{}{&its}).I32Free()
}

var _func_ImGuiInputTextState_OnKeyPressed_ = &c.FuncPrototype{Name: "ImGuiInputTextState_OnKeyPressed", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void ImGuiInputTextState_OnKeyPressed(ImGuiInputTextState* self,int key);
func (its *ImGuiInputTextState) OnKeyPressed(key int32) {
	giLib.Call(_func_ImGuiInputTextState_OnKeyPressed_, []interface{}{&its, &key})
}

var _func_ImGuiInputTextState_CursorAnimReset_ = &c.FuncPrototype{Name: "ImGuiInputTextState_CursorAnimReset", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextState_CursorAnimReset(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) CursorAnimReset() {
	giLib.Call(_func_ImGuiInputTextState_CursorAnimReset_, []interface{}{&its})
}

var _func_ImGuiInputTextState_CursorClamp_ = &c.FuncPrototype{Name: "ImGuiInputTextState_CursorClamp", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextState_CursorClamp(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) CursorClamp() {
	giLib.Call(_func_ImGuiInputTextState_CursorClamp_, []interface{}{&its})
}

var _func_ImGuiInputTextState_HasSelection_ = &c.FuncPrototype{Name: "ImGuiInputTextState_HasSelection", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiInputTextState_HasSelection(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) HasSelection() bool {
	return giLib.Call(_func_ImGuiInputTextState_HasSelection_, []interface{}{&its}).BoolFree()
}

var _func_ImGuiInputTextState_ClearSelection_ = &c.FuncPrototype{Name: "ImGuiInputTextState_ClearSelection", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextState_ClearSelection(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) ClearSelection() {
	giLib.Call(_func_ImGuiInputTextState_ClearSelection_, []interface{}{&its})
}

var _func_ImGuiInputTextState_GetCursorPos_ = &c.FuncPrototype{Name: "ImGuiInputTextState_GetCursorPos", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int ImGuiInputTextState_GetCursorPos(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) GetCursorPos() int32 {
	return giLib.Call(_func_ImGuiInputTextState_GetCursorPos_, []interface{}{&its}).I32Free()
}

var _func_ImGuiInputTextState_GetSelectionStart_ = &c.FuncPrototype{Name: "ImGuiInputTextState_GetSelectionStart", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int ImGuiInputTextState_GetSelectionStart(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) GetSelectionStart() int32 {
	return giLib.Call(_func_ImGuiInputTextState_GetSelectionStart_, []interface{}{&its}).I32Free()
}

var _func_ImGuiInputTextState_GetSelectionEnd_ = &c.FuncPrototype{Name: "ImGuiInputTextState_GetSelectionEnd", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int ImGuiInputTextState_GetSelectionEnd(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) GetSelectionEnd() int32 {
	return giLib.Call(_func_ImGuiInputTextState_GetSelectionEnd_, []interface{}{&its}).I32Free()
}

var _func_ImGuiInputTextState_SelectAll_ = &c.FuncPrototype{Name: "ImGuiInputTextState_SelectAll", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputTextState_SelectAll(ImGuiInputTextState* self);
func (its *ImGuiInputTextState) SelectAll() {
	giLib.Call(_func_ImGuiInputTextState_SelectAll_, []interface{}{&its})
}

var _func_ImGuiPopupData_ImGuiPopupData_ = &c.FuncPrototype{Name: "ImGuiPopupData_ImGuiPopupData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiPopupData* ImGuiPopupData_ImGuiPopupData(void);
func NewImGuiPopupData() *ImGuiPopupData {
	return (*ImGuiPopupData)(giLib.Call(_func_ImGuiPopupData_ImGuiPopupData_, nil).PtrFree())
}

var _func_ImGuiPopupData_destroy_ = &c.FuncPrototype{Name: "ImGuiPopupData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiPopupData_destroy(ImGuiPopupData* self);
func (pd *ImGuiPopupData) Destroy() {
	giLib.Call(_func_ImGuiPopupData_destroy_, []interface{}{&pd})
}

var _func_ImGuiNextWindowData_ImGuiNextWindowData_ = &c.FuncPrototype{Name: "ImGuiNextWindowData_ImGuiNextWindowData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiNextWindowData* ImGuiNextWindowData_ImGuiNextWindowData(void);
func NewImGuiNextWindowData() *ImGuiNextWindowData {
	return (*ImGuiNextWindowData)(giLib.Call(_func_ImGuiNextWindowData_ImGuiNextWindowData_, nil).PtrFree())
}

var _func_ImGuiNextWindowData_destroy_ = &c.FuncPrototype{Name: "ImGuiNextWindowData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiNextWindowData_destroy(ImGuiNextWindowData* self);
func (nwd *ImGuiNextWindowData) Destroy() {
	giLib.Call(_func_ImGuiNextWindowData_destroy_, []interface{}{&nwd})
}

var _func_ImGuiNextWindowData_ClearFlags_ = &c.FuncPrototype{Name: "ImGuiNextWindowData_ClearFlags", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiNextWindowData_ClearFlags(ImGuiNextWindowData* self);
func (nwd *ImGuiNextWindowData) ClearFlags() {
	giLib.Call(_func_ImGuiNextWindowData_ClearFlags_, []interface{}{&nwd})
}

var _func_ImGuiNextItemData_ImGuiNextItemData_ = &c.FuncPrototype{Name: "ImGuiNextItemData_ImGuiNextItemData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiNextItemData* ImGuiNextItemData_ImGuiNextItemData(void);
func NewImGuiNextItemData() *ImGuiNextItemData {
	return (*ImGuiNextItemData)(giLib.Call(_func_ImGuiNextItemData_ImGuiNextItemData_, nil).PtrFree())
}

var _func_ImGuiNextItemData_destroy_ = &c.FuncPrototype{Name: "ImGuiNextItemData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiNextItemData_destroy(ImGuiNextItemData* self);
func (nid *ImGuiNextItemData) Destroy() {
	giLib.Call(_func_ImGuiNextItemData_destroy_, []interface{}{&nid})
}

var _func_ImGuiNextItemData_ClearFlags_ = &c.FuncPrototype{Name: "ImGuiNextItemData_ClearFlags", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiNextItemData_ClearFlags(ImGuiNextItemData* self);
func (nid *ImGuiNextItemData) ClearFlags() {
	giLib.Call(_func_ImGuiNextItemData_ClearFlags_, []interface{}{&nid})
}

var _func_ImGuiLastItemData_ImGuiLastItemData_ = &c.FuncPrototype{Name: "ImGuiLastItemData_ImGuiLastItemData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiLastItemData* ImGuiLastItemData_ImGuiLastItemData(void);
func NewImGuiLastItemData() *ImGuiLastItemData {
	return (*ImGuiLastItemData)(giLib.Call(_func_ImGuiLastItemData_ImGuiLastItemData_, nil).PtrFree())
}

var _func_ImGuiLastItemData_destroy_ = &c.FuncPrototype{Name: "ImGuiLastItemData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiLastItemData_destroy(ImGuiLastItemData* self);
func (lid *ImGuiLastItemData) Destroy() {
	giLib.Call(_func_ImGuiLastItemData_destroy_, []interface{}{&lid})
}

var _func_ImGuiStackSizes_ImGuiStackSizes_ = &c.FuncPrototype{Name: "ImGuiStackSizes_ImGuiStackSizes", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiStackSizes* ImGuiStackSizes_ImGuiStackSizes(void);
func NewImGuiStackSizes() *ImGuiStackSizes {
	return (*ImGuiStackSizes)(giLib.Call(_func_ImGuiStackSizes_ImGuiStackSizes_, nil).PtrFree())
}

var _func_ImGuiStackSizes_destroy_ = &c.FuncPrototype{Name: "ImGuiStackSizes_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiStackSizes_destroy(ImGuiStackSizes* self);
func (ss *ImGuiStackSizes) Destroy() {
	giLib.Call(_func_ImGuiStackSizes_destroy_, []interface{}{&ss})
}

var _func_ImGuiStackSizes_SetToContextState_ = &c.FuncPrototype{Name: "ImGuiStackSizes_SetToContextState", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiStackSizes_SetToContextState(ImGuiStackSizes* self,ImGuiContext* ctx);
func (ss *ImGuiStackSizes) SetToContextState(ctx *ImGuiContext) {
	giLib.Call(_func_ImGuiStackSizes_SetToContextState_, []interface{}{&ss, &ctx})
}

var _func_ImGuiStackSizes_CompareWithContextState_ = &c.FuncPrototype{Name: "ImGuiStackSizes_CompareWithContextState", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiStackSizes_CompareWithContextState(ImGuiStackSizes* self,ImGuiContext* ctx);
func (ss *ImGuiStackSizes) CompareWithContextState(ctx *ImGuiContext) {
	giLib.Call(_func_ImGuiStackSizes_CompareWithContextState_, []interface{}{&ss, &ctx})
}

var _func_ImGuiPtrOrIndex_ImGuiPtrOrIndex_Ptr_ = &c.FuncPrototype{Name: "ImGuiPtrOrIndex_ImGuiPtrOrIndex_Ptr", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiPtrOrIndex* ImGuiPtrOrIndex_ImGuiPtrOrIndex_Ptr(void* ptr);
func NewImGuiPtrOrIndex_Ptr(ptr unsafe.Pointer) *ImGuiPtrOrIndex {
	return (*ImGuiPtrOrIndex)(giLib.Call(_func_ImGuiPtrOrIndex_ImGuiPtrOrIndex_Ptr_, []interface{}{&ptr}).PtrFree())
}

var _func_ImGuiPtrOrIndex_destroy_ = &c.FuncPrototype{Name: "ImGuiPtrOrIndex_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiPtrOrIndex_destroy(ImGuiPtrOrIndex* self);
func (pi *ImGuiPtrOrIndex) Destroy() {
	giLib.Call(_func_ImGuiPtrOrIndex_destroy_, []interface{}{&pi})
}

var _func_ImGuiPtrOrIndex_ImGuiPtrOrIndex_Int_ = &c.FuncPrototype{Name: "ImGuiPtrOrIndex_ImGuiPtrOrIndex_Int", OutType: c.Pointer, InTypes: _typs_I32}

// CIMGUI_API ImGuiPtrOrIndex* ImGuiPtrOrIndex_ImGuiPtrOrIndex_Int(int index);
func NewImGuiPtrOrIndexInt(index int32) *ImGuiPtrOrIndex {
	return (*ImGuiPtrOrIndex)(giLib.Call(_func_ImGuiPtrOrIndex_ImGuiPtrOrIndex_Int_, []interface{}{&index}).PtrFree())
}

var _func_ImGuiInputEvent_ImGuiInputEvent_ = &c.FuncPrototype{Name: "ImGuiInputEvent_ImGuiInputEvent", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiInputEvent* ImGuiInputEvent_ImGuiInputEvent(void);
func NewImGuiInputEvent() *ImGuiInputEvent {
	return (*ImGuiInputEvent)(giLib.Call(_func_ImGuiInputEvent_ImGuiInputEvent_, nil).PtrFree())
}

var _func_ImGuiInputEvent_destroy_ = &c.FuncPrototype{Name: "ImGuiInputEvent_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiInputEvent_destroy(ImGuiInputEvent* self);
func (ie *ImGuiInputEvent) Destroy() {
	giLib.Call(_func_ImGuiInputEvent_destroy_, []interface{}{&ie})
}

var _func_ImGuiKeyRoutingData_ImGuiKeyRoutingData_ = &c.FuncPrototype{Name: "ImGuiKeyRoutingData_ImGuiKeyRoutingData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiKeyRoutingData* ImGuiKeyRoutingData_ImGuiKeyRoutingData(void);
func NewImGuiKeyRoutingData() *ImGuiKeyRoutingData {
	return (*ImGuiKeyRoutingData)(giLib.Call(_func_ImGuiKeyRoutingData_ImGuiKeyRoutingData_, nil).PtrFree())
}

var _func_ImGuiKeyRoutingData_destroy_ = &c.FuncPrototype{Name: "ImGuiKeyRoutingData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiKeyRoutingData_destroy(ImGuiKeyRoutingData* self);
func (krd *ImGuiKeyRoutingData) Destroy() {
	giLib.Call(_func_ImGuiKeyRoutingData_destroy_, []interface{}{&krd})
}

var _func_ImGuiKeyRoutingTable_ImGuiKeyRoutingTable_ = &c.FuncPrototype{Name: "ImGuiKeyRoutingTable_ImGuiKeyRoutingTable", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiKeyRoutingTable* ImGuiKeyRoutingTable_ImGuiKeyRoutingTable(void);
func NewImGuiKeyRoutingTable() *ImGuiKeyRoutingTable {
	return (*ImGuiKeyRoutingTable)(giLib.Call(_func_ImGuiKeyRoutingTable_ImGuiKeyRoutingTable_, nil).PtrFree())
}

var _func_ImGuiKeyRoutingTable_destroy_ = &c.FuncPrototype{Name: "ImGuiKeyRoutingTable_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiKeyRoutingTable_destroy(ImGuiKeyRoutingTable* self);
func (krt *ImGuiKeyRoutingTable) Destroy() {
	giLib.Call(_func_ImGuiKeyRoutingTable_destroy_, []interface{}{&krt})
}

var _func_ImGuiKeyRoutingTable_Clear_ = &c.FuncPrototype{Name: "ImGuiKeyRoutingTable_Clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiKeyRoutingTable_Clear(ImGuiKeyRoutingTable* self);
func (krt *ImGuiKeyRoutingTable) Clear() {
	giLib.Call(_func_ImGuiKeyRoutingTable_Clear_, []interface{}{&krt})
}

var _func_ImGuiKeyOwnerData_ImGuiKeyOwnerData_ = &c.FuncPrototype{Name: "ImGuiKeyOwnerData_ImGuiKeyOwnerData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiKeyOwnerData* ImGuiKeyOwnerData_ImGuiKeyOwnerData(void);
func NewImGuiKeyOwnerData() *ImGuiKeyOwnerData {
	return (*ImGuiKeyOwnerData)(giLib.Call(_func_ImGuiKeyOwnerData_ImGuiKeyOwnerData_, nil).PtrFree())
}

var _func_ImGuiKeyOwnerData_destroy_ = &c.FuncPrototype{Name: "ImGuiKeyOwnerData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiKeyOwnerData_destroy(ImGuiKeyOwnerData* self);
func (kod *ImGuiKeyOwnerData) Destroy(self *ImGuiKeyOwnerData) {
	giLib.Call(_func_ImGuiKeyOwnerData_destroy_, []interface{}{&kod})
}

// CIMGUI_API ImGuiListClipperRange ImGuiListClipperRange_FromIndices(int min,int max);
// CIMGUI_API ImGuiListClipperRange ImGuiListClipperRange_FromPositions(float y1,float y2,int off_min,int off_max);

var _func_ImGuiListClipperData_ImGuiListClipperData_ = &c.FuncPrototype{Name: "ImGuiListClipperData_ImGuiListClipperData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiListClipperData* ImGuiListClipperData_ImGuiListClipperData(void);
func NewImGuiListClipperData() *ImGuiListClipperData {
	return (*ImGuiListClipperData)(giLib.Call(_func_ImGuiListClipperData_ImGuiListClipperData_, nil).PtrFree())
}

var _func_ImGuiListClipperData_destroy_ = &c.FuncPrototype{Name: "ImGuiListClipperData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiListClipperData_destroy(ImGuiListClipperData* self);
func (lcd *ImGuiListClipperData) Destroy() {
	giLib.Call(_func_ImGuiListClipperData_destroy_, []interface{}{&lcd})
}

var _func_ImGuiListClipperData_Reset_ = &c.FuncPrototype{Name: "ImGuiListClipperData_Reset", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiListClipperData_Reset(ImGuiListClipperData* self,ImGuiListClipper* clipper);
func (lcd *ImGuiListClipperData) Reset(clipper *ImGuiListClipper) {
	giLib.Call(_func_ImGuiListClipperData_Reset_, []interface{}{&lcd, &clipper})
}

var _func_ImGuiNavItemData_ImGuiNavItemData_ = &c.FuncPrototype{Name: "ImGuiNavItemData_ImGuiNavItemData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiNavItemData* ImGuiNavItemData_ImGuiNavItemData(void);
func NewImGuiNavItemData() *ImGuiNavItemData {
	return (*ImGuiNavItemData)(giLib.Call(_func_ImGuiNavItemData_ImGuiNavItemData_, nil).PtrFree())
}

var _func_ImGuiNavItemData_destroy_ = &c.FuncPrototype{Name: "ImGuiNavItemData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiNavItemData_destroy(ImGuiNavItemData* self);
func (nid *ImGuiNavItemData) Destroy() {
	giLib.Call(_func_ImGuiNavItemData_destroy_, []interface{}{&nid})
}

var _func_ImGuiNavItemData_Clear_ = &c.FuncPrototype{Name: "ImGuiNavItemData_Clear", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiNavItemData_Clear(ImGuiNavItemData* self);
func (nid *ImGuiNavItemData) Clear() {
	giLib.Call(_func_ImGuiNavItemData_Clear_, []interface{}{&nid})
}

var _func_ImGuiOldColumnData_ImGuiOldColumnData_ = &c.FuncPrototype{Name: "ImGuiOldColumnData_ImGuiOldColumnData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiOldColumnData* ImGuiOldColumnData_ImGuiOldColumnData(void);
func NewImGuiOldColumnData() *ImGuiOldColumnData {
	return (*ImGuiOldColumnData)(giLib.Call(_func_ImGuiOldColumnData_ImGuiOldColumnData_, nil).PtrFree())
}

var _func_ImGuiOldColumnData_destroy_ = &c.FuncPrototype{Name: "ImGuiOldColumnData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiOldColumnData_destroy(ImGuiOldColumnData* self);
func (ocd *ImGuiOldColumnData) Destroy() {
	giLib.Call(_func_ImGuiOldColumnData_destroy_, []interface{}{&ocd})
}

var _func_ImGuiOldColumns_ImGuiOldColumns_ = &c.FuncPrototype{Name: "ImGuiOldColumns_ImGuiOldColumns", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiOldColumns* ImGuiOldColumns_ImGuiOldColumns(void);
func NewImGuiOldColumns() *ImGuiOldColumns {
	return (*ImGuiOldColumns)(giLib.Call(_func_ImGuiOldColumns_ImGuiOldColumns_, nil).PtrFree())
}

var _func_ImGuiOldColumns_destroy_ = &c.FuncPrototype{Name: "ImGuiOldColumns_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiOldColumns_destroy(ImGuiOldColumns* self);
func (oc *ImGuiOldColumns) Destroy() {
	giLib.Call(_func_ImGuiOldColumns_destroy_, []interface{}{&oc})
}

var _func_ImGuiDockNode_ImGuiDockNode_ = &c.FuncPrototype{Name: "ImGuiDockNode_ImGuiDockNode", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiDockNode* ImGuiDockNode_ImGuiDockNode(ImGuiID id);
func NewImGuiDockNode(id ImGuiID) *ImGuiDockNode {
	return (*ImGuiDockNode)(giLib.Call(_func_ImGuiDockNode_ImGuiDockNode_, []interface{}{&id}).PtrFree())
}

var _func_ImGuiDockNode_destroy_ = &c.FuncPrototype{Name: "ImGuiDockNode_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiDockNode_destroy(ImGuiDockNode* self);
func (dn *ImGuiDockNode) Destroy() {
	giLib.Call(_func_ImGuiDockNode_destroy_, []interface{}{&dn})
}

var _func_ImGuiDockNode_IsRootNode_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsRootNode", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsRootNode(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsRootNode() bool {
	return giLib.Call(_func_ImGuiDockNode_IsRootNode_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_IsDockSpace_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsDockSpace", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsDockSpace(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsDockSpace() bool {
	return giLib.Call(_func_ImGuiDockNode_IsDockSpace_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_IsFloatingNode_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsFloatingNode", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsFloatingNode(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsFloatingNode() bool {
	return giLib.Call(_func_ImGuiDockNode_IsFloatingNode_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_IsCentralNode_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsCentralNode", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsCentralNode(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsCentralNode() bool {
	return giLib.Call(_func_ImGuiDockNode_IsCentralNode_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_IsHiddenTabBar_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsHiddenTabBar", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsHiddenTabBar(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsHiddenTabBar() bool {
	return giLib.Call(_func_ImGuiDockNode_IsHiddenTabBar_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_IsNoTabBar_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsNoTabBar", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsNoTabBar(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsNoTabBar() bool {
	return giLib.Call(_func_ImGuiDockNode_IsNoTabBar_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_IsSplitNode_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsSplitNode", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsSplitNode(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsSplitNode() bool {
	return giLib.Call(_func_ImGuiDockNode_IsSplitNode_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_IsLeafNode_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsLeafNode", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsLeafNode(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsLeafNode() bool {
	return giLib.Call(_func_ImGuiDockNode_IsLeafNode_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_IsEmpty_ = &c.FuncPrototype{Name: "ImGuiDockNode_IsEmpty", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool ImGuiDockNode_IsEmpty(ImGuiDockNode* self);
func (dn *ImGuiDockNode) IsEmpty() bool {
	return giLib.Call(_func_ImGuiDockNode_IsEmpty_, []interface{}{&dn}).BoolFree()
}

var _func_ImGuiDockNode_Rect_ = &c.FuncPrototype{Name: "ImGuiDockNode_Rect", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiDockNode_Rect(ImRect* pOut,ImGuiDockNode* self);
func (dn *ImGuiDockNode) Rect(pOut *ImRect) {
	giLib.Call(_func_ImGuiDockNode_Rect_, []interface{}{&pOut, &dn})
}

var _func_ImGuiDockNode_SetLocalFlags_ = &c.FuncPrototype{Name: "ImGuiDockNode_SetLocalFlags", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void ImGuiDockNode_SetLocalFlags(ImGuiDockNode* self,ImGuiDockNodeFlags flags);
func (dn *ImGuiDockNode) SetLocalFlags(flags ImGuiDockNodeFlags) {
	giLib.Call(_func_ImGuiDockNode_SetLocalFlags_, []interface{}{&dn, &flags})
}

var _func_ImGuiDockNode_UpdateMergedFlags_ = &c.FuncPrototype{Name: "ImGuiDockNode_UpdateMergedFlags", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiDockNode_UpdateMergedFlags(ImGuiDockNode* self);
func (dn *ImGuiDockNode) UpdateMergedFlags() {
	giLib.Call(_func_ImGuiDockNode_UpdateMergedFlags_, []interface{}{&dn})
}

var _func_ImGuiDockContext_ImGuiDockContext_ = &c.FuncPrototype{Name: "ImGuiDockContext_ImGuiDockContext", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiDockContext* ImGuiDockContext_ImGuiDockContext(void);
func NewImGuiDockContext() *ImGuiDockContext {
	return (*ImGuiDockContext)(giLib.Call(_func_ImGuiDockContext_ImGuiDockContext_, nil).PtrFree())
}

var _func_ImGuiDockContext_destroy_ = &c.FuncPrototype{Name: "ImGuiDockContext_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiDockContext_destroy(ImGuiDockContext* self);
func (dc *ImGuiDockContext) Destroy() {
	giLib.Call(_func_ImGuiDockContext_destroy_, []interface{}{&dc})
}

var _func_ImGuiViewportP_ImGuiViewportP_ = &c.FuncPrototype{Name: "ImGuiViewportP_ImGuiViewportP", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiViewportP* ImGuiViewportP_ImGuiViewportP(void);
func NewImGuiViewportP() *ImGuiViewportP {
	return (*ImGuiViewportP)(giLib.Call(_func_ImGuiViewportP_ImGuiViewportP_, nil).PtrFree())
}

var _func_ImGuiViewportP_destroy_ = &c.FuncPrototype{Name: "ImGuiViewportP_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiViewportP_destroy(ImGuiViewportP* self);
func (vp *ImGuiViewportP) Destroy() {
	giLib.Call(_func_ImGuiViewportP_destroy_, []interface{}{&vp})
}

var _func_ImGuiViewportP_ClearRequestFlags_ = &c.FuncPrototype{Name: "ImGuiViewportP_ClearRequestFlags", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiViewportP_ClearRequestFlags(ImGuiViewportP* self);
func (vp *ImGuiViewportP) ClearRequestFlags() {
	giLib.Call(_func_ImGuiViewportP_ClearRequestFlags_, []interface{}{&vp})
}

var _func_ImGuiViewportP_CalcWorkRectPos_ = &c.FuncPrototype{Name: "ImGuiViewportP_CalcWorkRectPos", OutType: c.Void, InTypes: _typs_PPVec2}

// CIMGUI_API void ImGuiViewportP_CalcWorkRectPos(ImVec2* pOut,ImGuiViewportP* self,const ImVec2 off_min);
func (vp *ImGuiViewportP) CalcWorkRectPos(pOut *ImVec2, off_min ImVec2) {
	giLib.Call(_func_ImGuiViewportP_CalcWorkRectPos_, []interface{}{&pOut, &vp, &off_min})
}

var _func_ImGuiViewportP_CalcWorkRectSize_ = &c.FuncPrototype{Name: "ImGuiViewportP_CalcWorkRectSize", OutType: c.Void, InTypes: _typs_PPVec2Vec2}

// CIMGUI_API void ImGuiViewportP_CalcWorkRectSize(ImVec2* pOut,ImGuiViewportP* self,const ImVec2 off_min,const ImVec2 off_max);
func (vp *ImGuiViewportP) CalcWorkRectSize(pOut *ImVec2, off_min, off_max ImVec2) {
	giLib.Call(_func_ImGuiViewportP_CalcWorkRectSize_, []interface{}{&pOut, &vp, &off_min, &off_max})
}

var _func_ImGuiViewportP_UpdateWorkRect_ = &c.FuncPrototype{Name: "ImGuiViewportP_UpdateWorkRect", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiViewportP_UpdateWorkRect(ImGuiViewportP* self);
func (vp *ImGuiViewportP) UpdateWorkRect() {
	giLib.Call(_func_ImGuiViewportP_UpdateWorkRect_, []interface{}{&vp})
}

var _func_ImGuiViewportP_GetMainRect_ = &c.FuncPrototype{Name: "ImGuiViewportP_GetMainRect", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiViewportP_GetMainRect(ImRect* pOut,ImGuiViewportP* self);
func (vp *ImGuiViewportP) GetMainRect(pOut *ImRect) {
	giLib.Call(_func_ImGuiViewportP_GetMainRect_, []interface{}{&pOut, &vp})
}

var _func_ImGuiViewportP_GetWorkRect_ = &c.FuncPrototype{Name: "ImGuiViewportP_GetWorkRect", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiViewportP_GetWorkRect(ImRect* pOut,ImGuiViewportP* self);
func (vp *ImGuiViewportP) GetWorkRect(pOut *ImRect) {
	giLib.Call(_func_ImGuiViewportP_GetWorkRect_, []interface{}{&pOut, &vp})
}

var _func_ImGuiViewportP_GetBuildWorkRect_ = &c.FuncPrototype{Name: "ImGuiViewportP_GetBuildWorkRect", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiViewportP_GetBuildWorkRect(ImRect* pOut,ImGuiViewportP* self);
func (vp *ImGuiViewportP) GetBuildWorkRect(pOut *ImRect) {
	giLib.Call(_func_ImGuiViewportP_GetBuildWorkRect_, []interface{}{&pOut, &vp})
}

var _func_ImGuiWindowSettings_ImGuiWindowSettings_ = &c.FuncPrototype{Name: "ImGuiWindowSettings_ImGuiWindowSettings", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiWindowSettings* ImGuiWindowSettings_ImGuiWindowSettings(void);
func NewImGuiWindowSettings() *ImGuiWindowSettings {
	return (*ImGuiWindowSettings)(giLib.Call(_func_ImGuiWindowSettings_ImGuiWindowSettings_, nil).PtrFree())
}

var _func_ImGuiWindowSettings_destroy_ = &c.FuncPrototype{Name: "ImGuiWindowSettings_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiWindowSettings_destroy(ImGuiWindowSettings* self);
func (ws *ImGuiWindowSettings) Destroy() {
	giLib.Call(_func_ImGuiWindowSettings_destroy_, []interface{}{&ws})
}

var _func_ImGuiWindowSettings_GetName_ = &c.FuncPrototype{Name: "ImGuiWindowSettings_GetName", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API char* ImGuiWindowSettings_GetName(ImGuiWindowSettings* self);
func (ws *ImGuiWindowSettings) GetName() string {
	return giLib.Call(_func_ImGuiWindowSettings_GetName_, []interface{}{&ws}).StrFree()
}

var _func_ImGuiSettingsHandler_ImGuiSettingsHandler_ = &c.FuncPrototype{Name: "ImGuiSettingsHandler_ImGuiSettingsHandler", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiSettingsHandler* ImGuiSettingsHandler_ImGuiSettingsHandler(void);
func NewImGuiSettingsHandler() *ImGuiSettingsHandler {
	return (*ImGuiSettingsHandler)(giLib.Call(_func_ImGuiSettingsHandler_ImGuiSettingsHandler_, nil).PtrFree())
}

var _func_ImGuiSettingsHandler_destroy_ = &c.FuncPrototype{Name: "ImGuiSettingsHandler_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiSettingsHandler_destroy(ImGuiSettingsHandler* self);
func (sh *ImGuiSettingsHandler) Destroy() {
	giLib.Call(_func_ImGuiSettingsHandler_destroy_, []interface{}{&sh})
}

var _func_ImGuiStackLevelInfo_ImGuiStackLevelInfo_ = &c.FuncPrototype{Name: "ImGuiStackLevelInfo_ImGuiStackLevelInfo", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiStackLevelInfo* ImGuiStackLevelInfo_ImGuiStackLevelInfo(void);
func NewImGuiStackLevelInfo() *ImGuiStackLevelInfo {
	return (*ImGuiStackLevelInfo)(giLib.Call(_func_ImGuiStackLevelInfo_ImGuiStackLevelInfo_, nil).PtrFree())
}

var _func_ImGuiStackLevelInfo_destroy_ = &c.FuncPrototype{Name: "ImGuiStackLevelInfo_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiStackLevelInfo_destroy(ImGuiStackLevelInfo* self);
func (si *ImGuiStackLevelInfo) Destroy() {
	giLib.Call(_func_ImGuiStackLevelInfo_destroy_, []interface{}{&si})
}

var _func_ImGuiStackTool_ImGuiStackTool_ = &c.FuncPrototype{Name: "ImGuiStackTool_ImGuiStackTool", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiStackTool* ImGuiStackTool_ImGuiStackTool(void);
func NewImGuiStackTool() *ImGuiStackTool {
	return (*ImGuiStackTool)(giLib.Call(_func_ImGuiStackTool_ImGuiStackTool_, nil).PtrFree())
}

var _func_ImGuiStackTool_destroy_ = &c.FuncPrototype{Name: "ImGuiStackTool_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiStackTool_destroy(ImGuiStackTool* self);
func (st *ImGuiStackTool) Destroy() {
	giLib.Call(_func_ImGuiStackTool_destroy_, []interface{}{&st})
}

var _func_ImGuiContextHook_ImGuiContextHook_ = &c.FuncPrototype{Name: "ImGuiContextHook_ImGuiContextHook", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiContextHook* ImGuiContextHook_ImGuiContextHook(void);
func NewImGuiContextHook() *ImGuiContextHook {
	return (*ImGuiContextHook)(giLib.Call(_func_ImGuiContextHook_ImGuiContextHook_, nil).PtrFree())
}

var _func_ImGuiContextHook_destroy_ = &c.FuncPrototype{Name: "ImGuiContextHook_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiContextHook_destroy(ImGuiContextHook* self);
func (ch *ImGuiContextHook) Destroy() {
	giLib.Call(_func_ImGuiContextHook_destroy_, []interface{}{&ch})
}

var _func_ImGuiContext_ImGuiContext_ = &c.FuncPrototype{Name: "ImGuiContext_ImGuiContext", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiContext* ImGuiContext_ImGuiContext(ImFontAtlas* shared_font_atlas);
func NewImGuiContext(shared_font_atlas *ImFontAtlas) *ImGuiContext {
	return (*ImGuiContext)(giLib.Call(_func_ImGuiContext_ImGuiContext_, []interface{}{&shared_font_atlas}).PtrFree())
}

var _func_ImGuiContext_destroy_ = &c.FuncPrototype{Name: "ImGuiContext_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiContext_destroy(ImGuiContext* self);
func (ctx *ImGuiContext) Destroy() {
	giLib.Call(_func_ImGuiContext_destroy_, []interface{}{&ctx})
}

var _func_ImGuiWindow_ImGuiWindow_ = &c.FuncPrototype{Name: "ImGuiWindow_ImGuiWindow", OutType: c.Pointer, InTypes: _typs_PP}

// CIMGUI_API ImGuiWindow* ImGuiWindow_ImGuiWindow(ImGuiContext* context,const char* name);
func NewImGuiWindow(context *ImGuiContext, name string) *ImGuiWindow {
	_name := c.CStr(name)
	defer usf.Free(_name)
	return (*ImGuiWindow)(giLib.Call(_func_ImGuiWindow_ImGuiWindow_, []interface{}{&context, &_name}).PtrFree())
}

var _func_ImGuiWindow_destroy_ = &c.FuncPrototype{Name: "ImGuiWindow_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiWindow_destroy(ImGuiWindow* self);
func (win *ImGuiWindow) Destroy() {
	giLib.Call(_func_ImGuiWindow_destroy_, []interface{}{&win})
}

var _func_ImGuiWindow_GetID_Str_ = &c.FuncPrototype{Name: "ImGuiWindow_GetID_Str", OutType: c.U32, InTypes: _typs_PPP}

// CIMGUI_API ImGuiID ImGuiWindow_GetID_Str(ImGuiWindow* self,const char* str,const char* str_end);
func (win *ImGuiWindow) GetIDStr(str string, str_end string) ImGuiID {
	_str, _str_end := c.CStr(str), c.CStr(str_end)
	defer usf.Free(_str)
	defer usf.Free(_str_end)
	return (ImGuiID)(giLib.Call(_func_ImGuiWindow_GetID_Str_, []interface{}{&win, &_str, &_str_end}).U32Free())
}

var _func_ImGuiWindow_GetID_Ptr_ = &c.FuncPrototype{Name: "ImGuiWindow_GetID_Ptr", OutType: c.U32, InTypes: _typs_PP}

// CIMGUI_API ImGuiID ImGuiWindow_GetID_Ptr(ImGuiWindow* self,const void* ptr);
func (win *ImGuiWindow) GetIDPtr(ptr unsafe.Pointer) ImGuiID {

	return (ImGuiID)(giLib.Call(_func_ImGuiWindow_GetID_Ptr_, []interface{}{&win, &ptr}).U32Free())
}

var _func_ImGuiWindow_GetID_Int_ = &c.FuncPrototype{Name: "ImGuiWindow_GetID_Int", OutType: c.U32, InTypes: _typs_PI32}

// CIMGUI_API ImGuiID ImGuiWindow_GetID_Int(ImGuiWindow* self,int n);
func (win *ImGuiWindow) GetIDInt(n int32) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_ImGuiWindow_GetID_Int_, []interface{}{&win, &n}).U32Free())
}

var _func_ImGuiWindow_GetIDFromRectangle_ = &c.FuncPrototype{Name: "ImGuiWindow_GetIDFromRectangle", OutType: c.U32, InTypes: _typs_PRect}

// CIMGUI_API ImGuiID ImGuiWindow_GetIDFromRectangle(ImGuiWindow* self,const ImRect r_abs);
func (win *ImGuiWindow) GetIDFromRectangle(r_abs ImRect) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_ImGuiWindow_GetIDFromRectangle_, []interface{}{&win, &r_abs}).U32Free())
}

var _func_ImGuiWindow_Rect_ = &c.FuncPrototype{Name: "ImGuiWindow_Rect", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiWindow_Rect(ImRect* pOut,ImGuiWindow* self);
func (win *ImGuiWindow) Rect(pOut *ImRect) {
	giLib.Call(_func_ImGuiWindow_Rect_, []interface{}{&pOut, &win})
}

var _func_ImGuiWindow_CalcFontSize_ = &c.FuncPrototype{Name: "ImGuiWindow_CalcFontSize", OutType: c.F32, InTypes: _typs_P}

// CIMGUI_API float ImGuiWindow_CalcFontSize(ImGuiWindow* self);
func (win *ImGuiWindow) CalcFontSize() float32 {
	return giLib.Call(_func_ImGuiWindow_CalcFontSize_, []interface{}{&win}).F32Free()
}

var _func_ImGuiWindow_TitleBarHeight_ = &c.FuncPrototype{Name: "ImGuiWindow_TitleBarHeight", OutType: c.F32, InTypes: _typs_P}

// CIMGUI_API float ImGuiWindow_TitleBarHeight(ImGuiWindow* self);
func (win *ImGuiWindow) TitleBarHeight() float32 {
	return giLib.Call(_func_ImGuiWindow_TitleBarHeight_, []interface{}{&win}).F32Free()
}

var _func_ImGuiWindow_TitleBarRect_ = &c.FuncPrototype{Name: "ImGuiWindow_TitleBarRect", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiWindow_TitleBarRect(ImRect* pOut,ImGuiWindow* self);
func (win *ImGuiWindow) TitleBarRect(pOut *ImRect) {
	giLib.Call(_func_ImGuiWindow_TitleBarRect_, []interface{}{&pOut, &win})
}

var _func_ImGuiWindow_MenuBarHeight_ = &c.FuncPrototype{Name: "ImGuiWindow_MenuBarHeight", OutType: c.F32, InTypes: _typs_P}

// CIMGUI_API float ImGuiWindow_MenuBarHeight(ImGuiWindow* self);
func (win *ImGuiWindow) MenuBarHeight() float32 {
	return giLib.Call(_func_ImGuiWindow_MenuBarHeight_, []interface{}{&win}).F32Free()
}

var _func_ImGuiWindow_MenuBarRect_ = &c.FuncPrototype{Name: "ImGuiWindow_MenuBarRect", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void ImGuiWindow_MenuBarRect(ImRect* pOut,ImGuiWindow* self);
func (win *ImGuiWindow) MenuBarRect(pOut *ImRect) {
	giLib.Call(_func_ImGuiWindow_MenuBarRect_, []interface{}{&pOut, &win})
}

var _func_ImGuiTabItem_ImGuiTabItem_ = &c.FuncPrototype{Name: "ImGuiTabItem_ImGuiTabItem", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTabItem* ImGuiTabItem_ImGuiTabItem(void);
func NewImGuiTabItem() *ImGuiTabItem {
	return (*ImGuiTabItem)(giLib.Call(_func_ImGuiTabItem_ImGuiTabItem_, nil).PtrFree())
}

var _func_ImGuiTabItem_destroy_ = &c.FuncPrototype{Name: "ImGuiTabItem_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTabItem_destroy(ImGuiTabItem* self);
func (ti *ImGuiTabItem) Destroy() {
	giLib.Call(_func_ImGuiTabItem_destroy_, []interface{}{&ti})
}

var _func_ImGuiTabBar_ImGuiTabBar_ = &c.FuncPrototype{Name: "ImGuiTabBar_ImGuiTabBar", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTabBar* ImGuiTabBar_ImGuiTabBar(void);
func NewImGuiTabBar() *ImGuiTabBar {
	return (*ImGuiTabBar)(giLib.Call(_func_ImGuiTabBar_ImGuiTabBar_, nil).PtrFree())
}

var _func_ImGuiTabBar_destroy_ = &c.FuncPrototype{Name: "ImGuiTabBar_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTabBar_destroy(ImGuiTabBar* self);
func (tb *ImGuiTabBar) Destroy(self *ImGuiTabBar) {
	giLib.Call(_func_ImGuiTabBar_destroy_, []interface{}{&tb})
}

var _func_ImGuiTableColumn_ImGuiTableColumn_ = &c.FuncPrototype{Name: "ImGuiTableColumn_ImGuiTableColumn", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTableColumn* ImGuiTableColumn_ImGuiTableColumn(void);
func NewImGuiTableColumn() *ImGuiTableColumn {
	return (*ImGuiTableColumn)(giLib.Call(_func_ImGuiTableColumn_ImGuiTableColumn_, nil).PtrFree())
}

var _func_ImGuiTableColumn_destroy_ = &c.FuncPrototype{Name: "ImGuiTableColumn_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTableColumn_destroy(ImGuiTableColumn* self);
func (tc *ImGuiTableColumn) Destroy() {
	giLib.Call(_func_ImGuiTableColumn_destroy_, []interface{}{&tc})
}

var _func_ImGuiTableInstanceData_ImGuiTableInstanceData_ = &c.FuncPrototype{Name: "ImGuiTableInstanceData_ImGuiTableInstanceData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTableInstanceData* ImGuiTableInstanceData_ImGuiTableInstanceData(void);
func NewImGuiTableInstanceData() *ImGuiTableInstanceData {
	return (*ImGuiTableInstanceData)(giLib.Call(_func_ImGuiTableInstanceData_ImGuiTableInstanceData_, nil).PtrFree())
}

var _func_ImGuiTableInstanceData_destroy_ = &c.FuncPrototype{Name: "ImGuiTableInstanceData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTableInstanceData_destroy(ImGuiTableInstanceData* self);
func (tid *ImGuiTableInstanceData) Destroy() {
	giLib.Call(_func_ImGuiTableInstanceData_destroy_, []interface{}{&tid})
}

var _func_ImGuiTable_ImGuiTable_ = &c.FuncPrototype{Name: "ImGuiTable_ImGuiTable", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTable* ImGuiTable_ImGuiTable(void);
func NewImGuiTable() *ImGuiTable {
	return (*ImGuiTable)(giLib.Call(_func_ImGuiTable_ImGuiTable_, nil).PtrFree())
}

var _func_ImGuiTable_destroy_ = &c.FuncPrototype{Name: "ImGuiTable_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTable_destroy(ImGuiTable* self);
func (tb *ImGuiTable) Destroy() {
	giLib.Call(_func_ImGuiTable_destroy_, []interface{}{&tb})
}

var _func_ImGuiTableTempData_ImGuiTableTempData_ = &c.FuncPrototype{Name: "ImGuiTableTempData_ImGuiTableTempData", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTableTempData* ImGuiTableTempData_ImGuiTableTempData(void);
func NewImGuiTableTempData() *ImGuiTableTempData {
	return (*ImGuiTableTempData)(giLib.Call(_func_ImGuiTableTempData_ImGuiTableTempData_, nil).PtrFree())
}

var _func_ImGuiTableTempData_destroy_ = &c.FuncPrototype{Name: "ImGuiTableTempData_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTableTempData_destroy(ImGuiTableTempData* self);
func (ttd *ImGuiTableTempData) Destroy() {
	giLib.Call(_func_ImGuiTableTempData_destroy_, []interface{}{&ttd})
}

var _func_ImGuiTableColumnSettings_ImGuiTableColumnSettings_ = &c.FuncPrototype{Name: "ImGuiTableColumnSettings_ImGuiTableColumnSettings", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTableColumnSettings* ImGuiTableColumnSettings_ImGuiTableColumnSettings(void);
func NewImGuiTableColumnSettings() *ImGuiTableColumnSettings {
	return (*ImGuiTableColumnSettings)(giLib.Call(_func_ImGuiTableColumnSettings_ImGuiTableColumnSettings_, nil).PtrFree())
}

var _func_ImGuiTableColumnSettings_destroy_ = &c.FuncPrototype{Name: "ImGuiTableColumnSettings_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTableColumnSettings_destroy(ImGuiTableColumnSettings* self);
func (tcs *ImGuiTableColumnSettings) Destroy(self *ImGuiTableColumnSettings) {
	giLib.Call(_func_ImGuiTableColumnSettings_destroy_, []interface{}{&tcs})
}

var _func_ImGuiTableSettings_ImGuiTableSettings_ = &c.FuncPrototype{Name: "ImGuiTableSettings_ImGuiTableSettings", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTableSettings* ImGuiTableSettings_ImGuiTableSettings(void);
func NewImGuiTableSettings() *ImGuiTableSettings {
	return (*ImGuiTableSettings)(giLib.Call(_func_ImGuiTableSettings_ImGuiTableSettings_, nil).PtrFree())
}

var _func_ImGuiTableSettings_destroy_ = &c.FuncPrototype{Name: "ImGuiTableSettings_destroy", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void ImGuiTableSettings_destroy(ImGuiTableSettings* self);
func (ts *ImGuiTableSettings) Destroy() {
	giLib.Call(_func_ImGuiTableSettings_destroy_, []interface{}{&ts})
}

var _func_ImGuiTableSettings_GetColumnSettings_ = &c.FuncPrototype{Name: "ImGuiTableSettings_GetColumnSettings", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiTableColumnSettings* ImGuiTableSettings_GetColumnSettings(ImGuiTableSettings* self);
func NewGetColumnSettings(self *ImGuiTableSettings) *ImGuiTableColumnSettings {
	return (*ImGuiTableColumnSettings)(giLib.Call(_func_ImGuiTableSettings_GetColumnSettings_, []interface{}{&self}).PtrFree())
}

var _func_igGetCurrentWindowRead_ = &c.FuncPrototype{Name: "igGetCurrentWindowRead", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiWindow* igGetCurrentWindowRead(void);
func GetCurrentWindowRead() *ImGuiWindow {
	return (*ImGuiWindow)(giLib.Call(_func_igGetCurrentWindowRead_, nil).PtrFree())
}

var _func_igGetCurrentWindow_ = &c.FuncPrototype{Name: "igGetCurrentWindow", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiWindow* igGetCurrentWindow(void);
func GetCurrentWindow() *ImGuiWindow {
	return (*ImGuiWindow)(giLib.Call(_func_igGetCurrentWindow_, nil).PtrFree())
}

var _func_igFindWindowByID_ = &c.FuncPrototype{Name: "igFindWindowByID", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiWindow* igFindWindowByID(ImGuiID id);
func FindWindowByID(id ImGuiID) *ImGuiWindow {
	return (*ImGuiWindow)(giLib.Call(_func_igFindWindowByID_, []interface{}{&id}).PtrFree())
}

var _func_igFindWindowByName_ = &c.FuncPrototype{Name: "igFindWindowByName", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiWindow* igFindWindowByName(const char* name);
func FindWindowByName(name string) *ImGuiWindow {
	_name := c.CStr(name)
	defer usf.Free(_name)
	return (*ImGuiWindow)(giLib.Call(_func_igFindWindowByName_, []interface{}{&_name}).PtrFree())
}

var _func_igUpdateWindowParentAndRootLinks_ = &c.FuncPrototype{Name: "igUpdateWindowParentAndRootLinks", OutType: c.Void, InTypes: _typs_PU32P}

// CIMGUI_API void igUpdateWindowParentAndRootLinks(ImGuiWindow* window,ImGuiWindowFlags flags,ImGuiWindow* parent_window);
func (window *ImGuiWindow) UpdateWindowParentAndRootLinks(flags ImGuiWindowFlags, parent_window *ImGuiWindow) {
	giLib.Call(_func_igUpdateWindowParentAndRootLinks_, []interface{}{&window, &flags, &parent_window})
}

var _func_igCalcWindowNextAutoFitSize_ = &c.FuncPrototype{Name: "igCalcWindowNextAutoFitSize", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igCalcWindowNextAutoFitSize(ImVec2* pOut,ImGuiWindow* window);
func (window *ImGuiWindow) CalcWindowNextAutoFitSize(pOut *ImVec2) {
	giLib.Call(_func_igCalcWindowNextAutoFitSize_, []interface{}{&pOut, &window})
}

var _func_igIsWindowChildOf_ = &c.FuncPrototype{Name: "igIsWindowChildOf", OutType: c.I32, InTypes: _typs_PPI32I32}

// CIMGUI_API bool igIsWindowChildOf(ImGuiWindow* window,ImGuiWindow* potential_parent,bool popup_hierarchy,bool dock_hierarchy);
func (window *ImGuiWindow) IsWindowChildOf(potential_parent *ImGuiWindow, popup_hierarchy, dock_hierarchy bool) bool {
	_popup_hierarchy, _dock_hierarchy := c.CBool(popup_hierarchy), c.CBool(dock_hierarchy)
	return giLib.Call(_func_igIsWindowChildOf_, []interface{}{&window, &potential_parent, &_popup_hierarchy, &_dock_hierarchy}).BoolFree()
}

var _func_igIsWindowWithinBeginStackOf_ = &c.FuncPrototype{Name: "igIsWindowWithinBeginStackOf", OutType: c.I32, InTypes: _typs_PP}

// CIMGUI_API bool igIsWindowWithinBeginStackOf(ImGuiWindow* window,ImGuiWindow* potential_parent);
func (window *ImGuiWindow) IsWindowWithinBeginStackOf(potential_parent *ImGuiWindow) bool {
	return giLib.Call(_func_igIsWindowWithinBeginStackOf_, []interface{}{&window, &potential_parent}).BoolFree()
}

var _func_igIsWindowAbove_ = &c.FuncPrototype{Name: "igIsWindowAbove", OutType: c.I32, InTypes: _typs_PP}

// CIMGUI_API bool igIsWindowAbove(ImGuiWindow* potential_above,ImGuiWindow* potential_below);
func (window *ImGuiWindow) IsWindowAbove(potential_below *ImGuiWindow) bool {
	return giLib.Call(_func_igIsWindowAbove_, []interface{}{&window, &potential_below}).BoolFree()
}

// CIMGUI_API bool igIsWindowAbove(ImGuiWindow* potential_above,ImGuiWindow* potential_below);
func (window *ImGuiWindow) IsWindowBelow(potential_above *ImGuiWindow) bool {
	return giLib.Call(_func_igIsWindowAbove_, []interface{}{&potential_above, &window}).BoolFree()
}

var _func_igIsWindowNavFocusable_ = &c.FuncPrototype{Name: "igIsWindowNavFocusable", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool igIsWindowNavFocusable(ImGuiWindow* window);
func (window *ImGuiWindow) IsWindowNavFocusable() bool {
	return giLib.Call(_func_igIsWindowNavFocusable_, []interface{}{&window}).BoolFree()
}

var _func_igSetWindowPos_WindowPtr_ = &c.FuncPrototype{Name: "igSetWindowPos_WindowPtr", OutType: c.Void, InTypes: _typs_PVec2U32}

// CIMGUI_API void igSetWindowPos_WindowPtr(ImGuiWindow* window,const ImVec2 pos,ImGuiCond cond);
func (window *ImGuiWindow) SetWindowPos(pos ImVec2, cond ImGuiCond) {
	giLib.Call(_func_igSetWindowPos_WindowPtr_, []interface{}{&window, &pos, &cond})
}

var _func_igSetWindowSize_WindowPtr_ = &c.FuncPrototype{Name: "igSetWindowSize_WindowPtr", OutType: c.Void, InTypes: _typs_PVec2U32}

// CIMGUI_API void igSetWindowSize_WindowPtr(ImGuiWindow* window,const ImVec2 size,ImGuiCond cond);
func (window *ImGuiWindow) SetWindowSize(size ImVec2, cond ImGuiCond) {
	giLib.Call(_func_igSetWindowSize_WindowPtr_, []interface{}{&window, &size, &cond})
}

var _func_igSetWindowCollapsed_WindowPtr_ = &c.FuncPrototype{Name: "igSetWindowCollapsed_WindowPtr", OutType: c.Void, InTypes: _typs_PI32U32}

// CIMGUI_API void igSetWindowCollapsed_WindowPtr(ImGuiWindow* window,bool collapsed,ImGuiCond cond);
func (window *ImGuiWindow) SetWindowCollapsed(collapsed bool, cond ImGuiCond) {
	_collapsed := c.CBool(collapsed)
	giLib.Call(_func_igSetWindowCollapsed_WindowPtr_, []interface{}{&window, &_collapsed, &cond})
}

var _func_igSetWindowHitTestHole_ = &c.FuncPrototype{Name: "igSetWindowHitTestHole", OutType: c.Void, InTypes: _typs_PVec2Vec2}

// CIMGUI_API void igSetWindowHitTestHole(ImGuiWindow* window,const ImVec2 pos,const ImVec2 size);
func (window *ImGuiWindow) SetWindowHitTestHole(pos, size ImVec2) {
	giLib.Call(_func_igSetWindowHitTestHole_, []interface{}{&window, &pos, &size})
}

var _func_igSetWindowHiddendAndSkipItemsForCurrentFrame_ = &c.FuncPrototype{Name: "igSetWindowHiddendAndSkipItemsForCurrentFrame", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igSetWindowHiddendAndSkipItemsForCurrentFrame(ImGuiWindow* window);
func (window *ImGuiWindow) SetWindowHiddendAndSkipItemsForCurrentFrame() {
	giLib.Call(_func_igSetWindowHiddendAndSkipItemsForCurrentFrame_, []interface{}{&window})
}

var _func_igWindowRectAbsToRel_ = &c.FuncPrototype{Name: "igWindowRectAbsToRel", OutType: c.Void, InTypes: _typs_PPRect}

// CIMGUI_API void igWindowRectAbsToRel(ImRect* pOut,ImGuiWindow* window,const ImRect r);
func (window *ImGuiWindow) WindowRectAbsToRel(pOut *ImRect, r ImRect) {
	giLib.Call(_func_igWindowRectAbsToRel_, []interface{}{&pOut, &window, &r})
}

var _func_igWindowRectRelToAbs_ = &c.FuncPrototype{Name: "igWindowRectRelToAbs", OutType: c.Void, InTypes: _typs_PPRect}

// CIMGUI_API void igWindowRectRelToAbs(ImRect* pOut,ImGuiWindow* window,const ImRect r);
func (window *ImGuiWindow) WindowRectRelToAbs(pOut *ImRect, r ImRect) {
	giLib.Call(_func_igWindowRectRelToAbs_, []interface{}{&pOut, &window, &r})
}

var _func_igWindowPosRelToAbs_ = &c.FuncPrototype{Name: "igWindowPosRelToAbs", OutType: c.Void, InTypes: _typs_PPVec2}

// CIMGUI_API void igWindowPosRelToAbs(ImVec2* pOut,ImGuiWindow* window,const ImVec2 p);
func (window *ImGuiWindow) WindowPosRelToAbs(pOut *ImVec2, p ImVec2) {
	giLib.Call(_func_igWindowPosRelToAbs_, []interface{}{&pOut, &window, &p})
}

var _func_igFocusWindow_ = &c.FuncPrototype{Name: "igFocusWindow", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void igFocusWindow(ImGuiWindow* window,ImGuiFocusRequestFlags flags);
func (window *ImGuiWindow) FocusWindow(flags ImGuiFocusRequestFlags) {
	giLib.Call(_func_igFocusWindow_, []interface{}{&window, &flags})
}

var _func_igFocusTopMostWindowUnderOne_ = &c.FuncPrototype{Name: "igFocusTopMostWindowUnderOne", OutType: c.Void, InTypes: _typs_PPPU32}

// CIMGUI_API void igFocusTopMostWindowUnderOne(ImGuiWindow* under_this_window,ImGuiWindow* ignore_window,ImGuiViewport* filter_viewport,ImGuiFocusRequestFlags flags);
func (window *ImGuiWindow) FocusTopMostWindowUnderOne(ignore_window *ImGuiWindow, filter_viewport *ImGuiViewport, flags ImGuiFocusRequestFlags) {
	giLib.Call(_func_igFocusTopMostWindowUnderOne_, []interface{}{&window, &ignore_window, &filter_viewport, &flags})
}

var _func_igBringWindowToFocusFront_ = &c.FuncPrototype{Name: "igBringWindowToFocusFront", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igBringWindowToFocusFront(ImGuiWindow* window);
func (window *ImGuiWindow) BringWindowToFocusFront() {
	giLib.Call(_func_igBringWindowToFocusFront_, []interface{}{&window})
}

var _func_igBringWindowToDisplayFront_ = &c.FuncPrototype{Name: "igBringWindowToDisplayFront", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igBringWindowToDisplayFront(ImGuiWindow* window);
func (window *ImGuiWindow) BringWindowToDisplayFront() {
	giLib.Call(_func_igBringWindowToDisplayFront_, []interface{}{&window})
}

var _func_igBringWindowToDisplayBack_ = &c.FuncPrototype{Name: "igBringWindowToDisplayBack", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igBringWindowToDisplayBack(ImGuiWindow* window);
func (window *ImGuiWindow) BringWindowToDisplayBack() {
	giLib.Call(_func_igBringWindowToDisplayBack_, []interface{}{&window})
}

var _func_igBringWindowToDisplayBehind_ = &c.FuncPrototype{Name: "igBringWindowToDisplayBehind", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igBringWindowToDisplayBehind(ImGuiWindow* window,ImGuiWindow* above_window);
func (window *ImGuiWindow) BringWindowToDisplayBehind(above_window *ImGuiWindow) {
	giLib.Call(_func_igBringWindowToDisplayBehind_, []interface{}{&window, &above_window})
}

var _func_igFindWindowDisplayIndex_ = &c.FuncPrototype{Name: "igFindWindowDisplayIndex", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int igFindWindowDisplayIndex(ImGuiWindow* window);
func (window *ImGuiWindow) FindWindowDisplayIndex() int32 {
	return giLib.Call(_func_igFindWindowDisplayIndex_, []interface{}{&window}).I32Free()
}

var _func_igFindBottomMostVisibleWindowWithinBeginStack_ = &c.FuncPrototype{Name: "igFindBottomMostVisibleWindowWithinBeginStack", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiWindow* igFindBottomMostVisibleWindowWithinBeginStack(ImGuiWindow* window);
func (window *ImGuiWindow) FindBottomMostVisibleWindowWithinBeginStack() *ImGuiWindow {
	return (*ImGuiWindow)(giLib.Call(_func_igFindBottomMostVisibleWindowWithinBeginStack_, []interface{}{&window}).PtrFree())
}

var _func_igSetCurrentFont_ = &c.FuncPrototype{Name: "igSetCurrentFont", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igSetCurrentFont(ImFont* font);
func SetCurrentFont(font *ImFont) {
	giLib.Call(_func_igSetCurrentFont_, []interface{}{&font})
}

var _func_igGetDefaultFont_ = &c.FuncPrototype{Name: "igGetDefaultFont", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImFont* igGetDefaultFont(void);
func GetDefaultFont() *ImFont {
	return (*ImFont)(giLib.Call(_func_igGetDefaultFont_, nil).PtrFree())
}

var _func_igGetForegroundDrawList_WindowPtr_ = &c.FuncPrototype{Name: "igGetForegroundDrawList_WindowPtr", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImDrawList* igGetForegroundDrawList_WindowPtr(ImGuiWindow* window);
func (window *ImGuiWindow) GetForegroundDrawList() *ImDrawList {
	return (*ImDrawList)(giLib.Call(_func_igGetForegroundDrawList_WindowPtr_, []interface{}{&window}).PtrFree())
}

var _func_igInitialize_ = &c.FuncPrototype{Name: "igInitialize", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igInitialize(void);
func Initialize() {
	giLib.Call(_func_igInitialize_, nil)
}

var _func_igShutdown_ = &c.FuncPrototype{Name: "igShutdown", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igShutdown(void);
func Shutdown() {
	giLib.Call(_func_igShutdown_, nil)
}

var _func_igUpdateInputEvents_ = &c.FuncPrototype{Name: "igUpdateInputEvents", OutType: c.Void, InTypes: _typs_I32}

// CIMGUI_API void igUpdateInputEvents(bool trickle_fast_inputs);
func UpdateInputEvents(trickle_fast_inputs bool) {
	_trickle_fast_inputs := c.CBool(trickle_fast_inputs)
	giLib.Call(_func_igUpdateInputEvents_, []interface{}{&_trickle_fast_inputs})
}

var _func_igUpdateHoveredWindowAndCaptureFlags_ = &c.FuncPrototype{Name: "igUpdateHoveredWindowAndCaptureFlags", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igUpdateHoveredWindowAndCaptureFlags(void);
func UpdateHoveredWindowAndCaptureFlags() {
	giLib.Call(_func_igUpdateHoveredWindowAndCaptureFlags_, nil)
}

var _func_igStartMouseMovingWindow_ = &c.FuncPrototype{Name: "igStartMouseMovingWindow", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igStartMouseMovingWindow(ImGuiWindow* window);
func (window *ImGuiWindow) StartMouseMovingWindow() {
	giLib.Call(_func_igStartMouseMovingWindow_, []interface{}{&window})
}

var _func_igStartMouseMovingWindowOrNode_ = &c.FuncPrototype{Name: "igStartMouseMovingWindowOrNode", OutType: c.Void, InTypes: _typs_PPI32}

// CIMGUI_API void igStartMouseMovingWindowOrNode(ImGuiWindow* window,ImGuiDockNode* node,bool undock_floating_node);
func (window *ImGuiWindow) StartMouseMovingWindowOrNode(node *ImGuiDockNode, undock_floating_node bool) {
	_undock_floating_node := c.CBool(undock_floating_node)
	giLib.Call(_func_igStartMouseMovingWindowOrNode_, []interface{}{&window, &node, &_undock_floating_node})
}

var _func_igUpdateMouseMovingWindowNewFrame_ = &c.FuncPrototype{Name: "igUpdateMouseMovingWindowNewFrame", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igUpdateMouseMovingWindowNewFrame(void);
func UpdateMouseMovingWindowNewFrame() {
	giLib.Call(_func_igUpdateMouseMovingWindowNewFrame_, nil)
}

var _func_igUpdateMouseMovingWindowEndFrame_ = &c.FuncPrototype{Name: "igUpdateMouseMovingWindowEndFrame", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igUpdateMouseMovingWindowEndFrame(void);
func UpdateMouseMovingWindowEndFrame() {
	giLib.Call(_func_igUpdateMouseMovingWindowEndFrame_, nil)
}

var _func_igAddContextHook_ = &c.FuncPrototype{Name: "igAddContextHook", OutType: c.U32, InTypes: _typs_PP}

// CIMGUI_API ImGuiID igAddContextHook(ImGuiContext* context,const ImGuiContextHook* hook);
func AddContextHook(context *ImGuiContext, hook *ImGuiContextHook) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igAddContextHook_, []interface{}{&context, &hook}).U32Free())
}

var _func_igRemoveContextHook_ = &c.FuncPrototype{Name: "igRemoveContextHook", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void igRemoveContextHook(ImGuiContext* context,ImGuiID hook_to_remove);
func RemoveContextHook(context *ImGuiContext, hook_to_remove ImGuiID) {
	giLib.Call(_func_igRemoveContextHook_, []interface{}{&context, &hook_to_remove})
}

var _func_igCallContextHooks_ = &c.FuncPrototype{Name: "igCallContextHooks", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void igCallContextHooks(ImGuiContext* context,ImGuiContextHookType typ);
func CallContextHooks(context *ImGuiContext, typ ImGuiContextHookType) {
	giLib.Call(_func_igCallContextHooks_, []interface{}{&context, &typ})
}

var _func_igTranslateWindowsInViewport_ = &c.FuncPrototype{Name: "igTranslateWindowsInViewport", OutType: c.Void, InTypes: _typs_PVec2Vec2}

// CIMGUI_API void igTranslateWindowsInViewport(ImGuiViewportP* viewport,const ImVec2 old_pos,const ImVec2 new_pos);
func TranslateWindowsInViewport(viewport *ImGuiViewportP, old_pos ImVec2, new_pos ImVec2) {
	giLib.Call(_func_igTranslateWindowsInViewport_, []interface{}{&viewport, &old_pos, &new_pos})
}

var _func_igScaleWindowsInViewport_ = &c.FuncPrototype{Name: "igScaleWindowsInViewport", OutType: c.Void, InTypes: _typs_PF32}

// CIMGUI_API void igScaleWindowsInViewport(ImGuiViewportP* viewport,float scale);
func ScaleWindowsInViewport(viewport *ImGuiViewportP, scale float32) {
	giLib.Call(_func_igScaleWindowsInViewport_, []interface{}{&viewport, &scale})
}

var _func_igDestroyPlatformWindow_ = &c.FuncPrototype{Name: "igDestroyPlatformWindow", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDestroyPlatformWindow(ImGuiViewportP* viewport);
func DestroyPlatformWindow(viewport *ImGuiViewportP) {
	giLib.Call(_func_igDestroyPlatformWindow_, []interface{}{&viewport})
}

var _func_igSetWindowViewport_ = &c.FuncPrototype{Name: "igSetWindowViewport", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igSetWindowViewport(ImGuiWindow* window,ImGuiViewportP* viewport);
func SetWindowViewport(window *ImGuiWindow, viewport *ImGuiViewportP) {
	giLib.Call(_func_igSetWindowViewport_, []interface{}{&window, &viewport})
}

var _func_igSetCurrentViewport_ = &c.FuncPrototype{Name: "igSetCurrentViewport", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igSetCurrentViewport(ImGuiWindow* window,ImGuiViewportP* viewport);
func SetCurrentViewport(window *ImGuiWindow, viewport *ImGuiViewportP) {
	giLib.Call(_func_igSetCurrentViewport_, []interface{}{&window, &viewport})
}

var _func_igGetViewportPlatformMonitor_ = &c.FuncPrototype{Name: "igGetViewportPlatformMonitor", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API const ImGuiPlatformMonitor* igGetViewportPlatformMonitor(ImGuiViewport* viewport);
func GetViewportPlatformMonitor(viewport *ImGuiViewport) *ImGuiPlatformMonitor {
	return (*ImGuiPlatformMonitor)(giLib.Call(_func_igGetViewportPlatformMonitor_, []interface{}{&viewport}).PtrFree())
}

var _func_igFindHoveredViewportFromPlatformWindowStack_ = &c.FuncPrototype{Name: "igFindHoveredViewportFromPlatformWindowStack", OutType: c.Pointer, InTypes: _typs_Vec2}

// CIMGUI_API ImGuiViewportP* igFindHoveredViewportFromPlatformWindowStack(const ImVec2 mouse_platform_pos);
func FindHoveredViewportFromPlatformWindowStack(mouse_platform_pos ImVec2) *ImGuiViewportP {
	return (*ImGuiViewportP)(giLib.Call(_func_igFindHoveredViewportFromPlatformWindowStack_, []interface{}{&mouse_platform_pos}).PtrFree())
}

var _func_igMarkIniSettingsDirty_Nil_ = &c.FuncPrototype{Name: "igMarkIniSettingsDirty_Nil", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igMarkIniSettingsDirty_Nil(void);
func MarkIniSettingsDirty_Nil() {
	giLib.Call(_func_igMarkIniSettingsDirty_Nil_, nil)
}

var _func_igMarkIniSettingsDirty_WindowPtr_ = &c.FuncPrototype{Name: "igMarkIniSettingsDirty_WindowPtr", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igMarkIniSettingsDirty_WindowPtr(ImGuiWindow* window);
func MarkIniSettingsDirtyWindow(window *ImGuiWindow) {
	giLib.Call(_func_igMarkIniSettingsDirty_WindowPtr_, []interface{}{&window})
}

var _func_igClearIniSettings_ = &c.FuncPrototype{Name: "igClearIniSettings", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igClearIniSettings(void);
func ClearIniSettings() {
	giLib.Call(_func_igClearIniSettings_, nil)
}

var _func_igAddSettingsHandler_ = &c.FuncPrototype{Name: "igAddSettingsHandler", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igAddSettingsHandler(const ImGuiSettingsHandler* handler);
func AddSettingsHandler(handler *ImGuiSettingsHandler) {
	giLib.Call(_func_igAddSettingsHandler_, []interface{}{&handler})
}

var _func_igRemoveSettingsHandler_ = &c.FuncPrototype{Name: "igRemoveSettingsHandler", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igRemoveSettingsHandler(const char* type_name);
func RemoveSettingsHandler(type_name string) {
	_type_name := c.CStr(type_name)
	defer usf.Free(_type_name)
	giLib.Call(_func_igRemoveSettingsHandler_, []interface{}{&_type_name})
}

var _func_igFindSettingsHandler_ = &c.FuncPrototype{Name: "igFindSettingsHandler", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiSettingsHandler* igFindSettingsHandler(const char* type_name);
func FindSettingsHandler(type_name string) *ImGuiSettingsHandler {
	_type_name := c.CStr(type_name)
	defer usf.Free(_type_name)

	return (*ImGuiSettingsHandler)(giLib.Call(_func_igFindSettingsHandler_, []interface{}{&_type_name}).PtrFree())
}

var _func_igCreateNewWindowSettings_ = &c.FuncPrototype{Name: "igCreateNewWindowSettings", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiWindowSettings* igCreateNewWindowSettings(const char* name);
func CreateNewWindowSettings(name string) *ImGuiWindowSettings {
	_name := c.CStr(name)
	defer usf.Free(_name)

	return (*ImGuiWindowSettings)(giLib.Call(_func_igCreateNewWindowSettings_, []interface{}{&_name}).PtrFree())
}

var _func_igFindWindowSettingsByID_ = &c.FuncPrototype{Name: "igFindWindowSettingsByID", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiWindowSettings* igFindWindowSettingsByID(ImGuiID id);
func FindWindowSettingsByID(id ImGuiID) *ImGuiWindowSettings {

	return (*ImGuiWindowSettings)(giLib.Call(_func_igFindWindowSettingsByID_, []interface{}{&id}).PtrFree())
}

var _func_igFindWindowSettingsByWindow_ = &c.FuncPrototype{Name: "igFindWindowSettingsByWindow", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiWindowSettings* igFindWindowSettingsByWindow(ImGuiWindow* window);
func FindWindowSettingsByWindow(window *ImGuiWindow) *ImGuiWindowSettings {
	return (*ImGuiWindowSettings)(giLib.Call(_func_igFindWindowSettingsByWindow_, []interface{}{&window}).PtrFree())
}

var _func_igClearWindowSettings_ = &c.FuncPrototype{Name: "igClearWindowSettings", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igClearWindowSettings(const char* name);
func ClearWindowSettings(name string) {
	_name := c.CStr(name)
	defer usf.Free(_name)
	giLib.Call(_func_igClearWindowSettings_, []interface{}{&_name})
}

var _func_igLocalizeRegisterEntries_ = &c.FuncPrototype{Name: "igLocalizeRegisterEntries", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void igLocalizeRegisterEntries(const ImGuiLocEntry* entries,int count);
func LocalizeRegisterEntries(entries *ImGuiLocEntry, count int32) {
	giLib.Call(_func_igLocalizeRegisterEntries_, []interface{}{&entries, &count})
}

var _func_igLocalizeGetMsg_ = &c.FuncPrototype{Name: "igLocalizeGetMsg", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API const char* igLocalizeGetMsg(ImGuiLocKey key);
func LocalizeGetMsg(key ImGuiLocKey) string {
	return giLib.Call(_func_igLocalizeGetMsg_, []interface{}{&key}).StrFree()
}

var _func_igSetScrollX_WindowPtr_ = &c.FuncPrototype{Name: "igSetScrollX_WindowPtr", OutType: c.Void, InTypes: _typs_PF32}

// CIMGUI_API void igSetScrollX_WindowPtr(ImGuiWindow* window,float scroll_x);
func SetScrollXWindow(window *ImGuiWindow, scroll_x float32) {
	giLib.Call(_func_igSetScrollX_WindowPtr_, []interface{}{&window, &scroll_x})
}

var _func_igSetScrollY_WindowPtr_ = &c.FuncPrototype{Name: "igSetScrollY_WindowPtr", OutType: c.Void, InTypes: _typs_PF32}

// CIMGUI_API void igSetScrollY_WindowPtr(ImGuiWindow* window,float scroll_y);
func SetScrollYWindow(window *ImGuiWindow, scroll_y float32) {
	giLib.Call(_func_igSetScrollY_WindowPtr_, []interface{}{&window, &scroll_y})
}

var _func_igSetScrollFromPosX_WindowPtr_ = &c.FuncPrototype{Name: "igSetScrollFromPosX_WindowPtr", OutType: c.Void, InTypes: _typs_PF32F32}

// CIMGUI_API void igSetScrollFromPosX_WindowPtr(ImGuiWindow* window,float local_x,float center_x_ratio);
func SetScrollFromPosXWindow(window *ImGuiWindow, local_x float32, center_x_ratio float32) {
	giLib.Call(_func_igSetScrollFromPosX_WindowPtr_, []interface{}{&window, &local_x, &center_x_ratio})
}

var _func_igSetScrollFromPosY_WindowPtr_ = &c.FuncPrototype{Name: "igSetScrollFromPosY_WindowPtr", OutType: c.Void, InTypes: _typs_PF32F32}

// CIMGUI_API void igSetScrollFromPosY_WindowPtr(ImGuiWindow* window,float local_y,float center_y_ratio);
func SetScrollFromPosYWindow(window *ImGuiWindow, local_y float32, center_y_ratio float32) {
	giLib.Call(_func_igSetScrollFromPosY_WindowPtr_, []interface{}{&window, &local_y, &center_y_ratio})
}

var _func_igScrollToItem_ = &c.FuncPrototype{Name: "igScrollToItem", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igScrollToItem(ImGuiScrollFlags flags);
func ScrollToItem(flags ImGuiScrollFlags) {
	giLib.Call(_func_igScrollToItem_, []interface{}{&flags})
}

var _func_igScrollToRect_ = &c.FuncPrototype{Name: "igScrollToRect", OutType: c.Void, InTypes: _typs_PRectU32}

// CIMGUI_API void igScrollToRect(ImGuiWindow* window,const ImRect rect,ImGuiScrollFlags flags);
func ScrollToRect(window *ImGuiWindow, rect ImRect, flags ImGuiScrollFlags) {
	giLib.Call(_func_igScrollToRect_, []interface{}{&window, &rect, &flags})
}

var _func_igScrollToRectEx_ = &c.FuncPrototype{Name: "igScrollToRectEx", OutType: c.Void, InTypes: _typs_PPRectU32}

// CIMGUI_API void igScrollToRectEx(ImVec2* pOut,ImGuiWindow* window,const ImRect rect,ImGuiScrollFlags flags);
func ScrollToRectEx(pOut *ImVec2, window *ImGuiWindow, rect ImRect, flags ImGuiScrollFlags) {
	giLib.Call(_func_igScrollToRectEx_, []interface{}{&pOut, &window, &rect, &flags})
}

var _func_igScrollToBringRectIntoView_ = &c.FuncPrototype{Name: "igScrollToBringRectIntoView", OutType: c.Void, InTypes: _typs_PRect}

// CIMGUI_API void igScrollToBringRectIntoView(ImGuiWindow* window,const ImRect rect);
func ScrollToBringRectIntoView(window *ImGuiWindow, rect ImRect) {
	giLib.Call(_func_igScrollToBringRectIntoView_, []interface{}{&window, &rect})
}

var _func_igGetItemStatusFlags_ = &c.FuncPrototype{Name: "igGetItemStatusFlags", OutType: c.U32, InTypes: nil}

// CIMGUI_API ImGuiItemStatusFlags igGetItemStatusFlags(void);
func GetItemStatusFlags() ImGuiItemStatusFlags {
	return (ImGuiItemStatusFlags)(giLib.Call(_func_igGetItemStatusFlags_, nil).U32Free())
}

var _func_igGetItemFlags_ = &c.FuncPrototype{Name: "igGetItemFlags", OutType: c.U32, InTypes: nil}

// CIMGUI_API ImGuiItemFlags igGetItemFlags(void);
func GetItemFlags() ImGuiItemFlags {
	return (ImGuiItemFlags)(giLib.Call(_func_igGetItemFlags_, nil).U32Free())
}

var _func_igGetActiveID_ = &c.FuncPrototype{Name: "igGetActiveID", OutType: c.U32, InTypes: nil}

// CIMGUI_API ImGuiID igGetActiveID(void);
func GetActiveID() ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetActiveID_, nil).U32Free())
}

var _func_igGetFocusID_ = &c.FuncPrototype{Name: "igGetFocusID", OutType: c.U32, InTypes: nil}

// CIMGUI_API ImGuiID igGetFocusID(void);
func GetFocusID() ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetFocusID_, nil).U32Free())
}

var _func_igSetActiveID_ = &c.FuncPrototype{Name: "igSetActiveID", OutType: c.Void, InTypes: _typs_U32P}

// CIMGUI_API void igSetActiveID(ImGuiID id,ImGuiWindow* window);
func SetActiveID(id ImGuiID, window *ImGuiWindow) {
	giLib.Call(_func_igSetActiveID_, []interface{}{&id, &window})
}

var _func_igSetFocusID_ = &c.FuncPrototype{Name: "igSetFocusID", OutType: c.Void, InTypes: _typs_U32P}

// CIMGUI_API void igSetFocusID(ImGuiID id,ImGuiWindow* window);
func SetFocusID(id ImGuiID, window *ImGuiWindow) {
	giLib.Call(_func_igSetFocusID_, []interface{}{&id, &window})
}

var _func_igClearActiveID_ = &c.FuncPrototype{Name: "igClearActiveID", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igClearActiveID(void);
func ClearActiveID() {
	giLib.Call(_func_igClearActiveID_, nil)
}

var _func_igGetHoveredID_ = &c.FuncPrototype{Name: "igGetHoveredID", OutType: c.U32, InTypes: nil}

// CIMGUI_API ImGuiID igGetHoveredID(void);
func GetHoveredID() ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetHoveredID_, nil).U32Free())
}

var _func_igSetHoveredID_ = &c.FuncPrototype{Name: "igSetHoveredID", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igSetHoveredID(ImGuiID id);
func SetHoveredID(id ImGuiID) {
	giLib.Call(_func_igSetHoveredID_, []interface{}{&id})
}

var _func_igKeepAliveID_ = &c.FuncPrototype{Name: "igKeepAliveID", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igKeepAliveID(ImGuiID id);
func KeepAliveID(id ImGuiID) {
	giLib.Call(_func_igKeepAliveID_, []interface{}{&id})
}

var _func_igMarkItemEdited_ = &c.FuncPrototype{Name: "igMarkItemEdited", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igMarkItemEdited(ImGuiID id);
func MarkItemEdited(id ImGuiID) {
	giLib.Call(_func_igMarkItemEdited_, []interface{}{&id})
}

var _func_igPushOverrideID_ = &c.FuncPrototype{Name: "igPushOverrideID", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igPushOverrideID(ImGuiID id);
func PushOverrideID(id ImGuiID) {
	giLib.Call(_func_igPushOverrideID_, []interface{}{&id})
}

var _func_igGetIDWithSeed_Str_ = &c.FuncPrototype{Name: "igGetIDWithSeed_Str", OutType: c.U32, InTypes: _typs_PPU32}

// CIMGUI_API ImGuiID igGetIDWithSeed_Str(const char* str_id_begin,const char* str_id_end,ImGuiID seed);
func GetIDWithSeedStr(str_id_begin string, str_id_end string, seed ImGuiID) ImGuiID {
	_str_id_begin, _str_id_end := c.CStr(str_id_begin), c.CStr(str_id_end)
	defer usf.Free(_str_id_begin)
	defer usf.Free(_str_id_end)

	return (ImGuiID)(giLib.Call(_func_igGetIDWithSeed_Str_, []interface{}{&_str_id_begin, &_str_id_end, &seed}).U32Free())
}

var _func_igGetIDWithSeed_Int_ = &c.FuncPrototype{Name: "igGetIDWithSeed_Int", OutType: c.U32, InTypes: _typs_I32U32}

// CIMGUI_API ImGuiID igGetIDWithSeed_Int(int n,ImGuiID seed);
func GetIDWithSeed_Int(n int32, seed ImGuiID) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetIDWithSeed_Int_, []interface{}{&n, &seed}).U32Free())
}

var _func_igItemSize_Vec2_ = &c.FuncPrototype{Name: "igItemSize_Vec2", OutType: c.Void, InTypes: _typs_Vec2F32}

// CIMGUI_API void igItemSize_Vec2(const ImVec2 size,float text_baseline_y);
func ItemSizeVec2(size ImVec2, text_baseline_y float32) {
	giLib.Call(_func_igItemSize_Vec2_, []interface{}{&size, &text_baseline_y})
}

var _func_igItemSize_Rect_ = &c.FuncPrototype{Name: "igItemSize_Rect", OutType: c.Void, InTypes: _typs_RectF32}

// CIMGUI_API void igItemSize_Rect(const ImRect bb,float text_baseline_y);
func ItemSizeRect(bb ImRect, text_baseline_y float32) {
	giLib.Call(_func_igItemSize_Rect_, []interface{}{&bb, &text_baseline_y})
}

var _func_igItemAdd_ = &c.FuncPrototype{Name: "igItemAdd", OutType: c.I32, InTypes: _typs_RectU32PU32}

// CIMGUI_API bool igItemAdd(const ImRect bb,ImGuiID id,const ImRect* nav_bb,ImGuiItemFlags extra_flags);
func ItemAdd(bb ImRect, id ImGuiID, nav_bb *ImRect, extra_flags ImGuiItemFlags) bool {
	return giLib.Call(_func_igItemAdd_, []interface{}{&bb, &id, &nav_bb, &extra_flags}).BoolFree()
}

var _func_igItemHoverable_ = &c.FuncPrototype{Name: "igItemHoverable", OutType: c.I32, InTypes: _typs_RectU32U32}

// CIMGUI_API bool igItemHoverable(const ImRect bb,ImGuiID id,ImGuiItemFlags item_flags);
func ItemHoverable(bb ImRect, id ImGuiID, item_flags ImGuiItemFlags) bool {
	return giLib.Call(_func_igItemHoverable_, []interface{}{&bb, &id, &item_flags}).BoolFree()
}

var _func_igIsWindowContentHoverable_ = &c.FuncPrototype{Name: "igIsWindowContentHoverable", OutType: c.I32, InTypes: _typs_PU32}

// CIMGUI_API bool igIsWindowContentHoverable(ImGuiWindow* window,ImGuiHoveredFlags flags);
func IsWindowContentHoverable(window *ImGuiWindow, flags ImGuiHoveredFlags) bool {
	return giLib.Call(_func_igIsWindowContentHoverable_, []interface{}{&window, &flags}).BoolFree()
}

var _func_igIsClippedEx_ = &c.FuncPrototype{Name: "igIsClippedEx", OutType: c.I32, InTypes: _typs_RectU32}

// CIMGUI_API bool igIsClippedEx(const ImRect bb,ImGuiID id);
func IsClippedEx(bb ImRect, id ImGuiID) bool {
	return giLib.Call(_func_igIsClippedEx_, []interface{}{&bb, &id}).BoolFree()
}

var _func_igSetLastItemData_ = &c.FuncPrototype{Name: "igSetLastItemData", OutType: c.Void, InTypes: _typs_U32U32U32Rect}

// CIMGUI_API void igSetLastItemData(ImGuiID item_id,ImGuiItemFlags in_flags,ImGuiItemStatusFlags status_flags,const ImRect item_rect);
func SetLastItemData(item_id ImGuiID, in_flags ImGuiItemFlags, status_flags ImGuiItemStatusFlags, item_rect ImRect) {
	giLib.Call(_func_igSetLastItemData_, []interface{}{&item_id, &in_flags, &status_flags, &item_rect})
}

var _func_igCalcItemSize_ = &c.FuncPrototype{Name: "igCalcItemSize", OutType: c.Void, InTypes: _typs_PVec2F32F32}

// CIMGUI_API void igCalcItemSize(ImVec2* pOut,ImVec2 size,float default_w,float default_h);
func CalcItemSize(pOut *ImVec2, size ImVec2, default_w float32, default_h float32) {
	giLib.Call(_func_igCalcItemSize_, []interface{}{&pOut, &size, &default_w, &default_h})
}

var _func_igCalcWrapWidthForPos_ = &c.FuncPrototype{Name: "igCalcWrapWidthForPos", OutType: c.F32, InTypes: _typs_Vec2F32}

// CIMGUI_API float igCalcWrapWidthForPos(const ImVec2 pos,float wrap_pos_x);
func CalcWrapWidthForPos(pos ImVec2, wrap_pos_x float32) float32 {
	return giLib.Call(_func_igCalcWrapWidthForPos_, []interface{}{&pos, &wrap_pos_x}).F32Free()
}

var _func_igPushMultiItemsWidths_ = &c.FuncPrototype{Name: "igPushMultiItemsWidths", OutType: c.Void, InTypes: _typs_I32F32}

// CIMGUI_API void igPushMultiItemsWidths(int components,float width_full);
func PushMultiItemsWidths(components int32, width_full float32) {
	giLib.Call(_func_igPushMultiItemsWidths_, []interface{}{&components, &width_full})
}

var _func_igIsItemToggledSelection_ = &c.FuncPrototype{Name: "igIsItemToggledSelection", OutType: c.I32, InTypes: nil}

// CIMGUI_API bool igIsItemToggledSelection(void);
func IsItemToggledSelection() bool {
	return giLib.Call(_func_igIsItemToggledSelection_, nil).BoolFree()
}

var _func_igGetContentRegionMaxAbs_ = &c.FuncPrototype{Name: "igGetContentRegionMaxAbs", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igGetContentRegionMaxAbs(ImVec2* pOut);
func GetContentRegionMaxAbs(pOut *ImVec2) {
	giLib.Call(_func_igGetContentRegionMaxAbs_, []interface{}{&pOut})
}

var _func_igShrinkWidths_ = &c.FuncPrototype{Name: "igShrinkWidths", OutType: c.Void, InTypes: _typs_PI32F32}

// CIMGUI_API void igShrinkWidths(ImGuiShrinkWidthItem* items,int count,float width_excess);
func ShrinkWidths(items *ImGuiShrinkWidthItem, count int32, width_excess float32) {
	giLib.Call(_func_igShrinkWidths_, []interface{}{&items, &count, &width_excess})
}

var _func_igPushItemFlag_ = &c.FuncPrototype{Name: "igPushItemFlag", OutType: c.Void, InTypes: _typs_U32I32}

// CIMGUI_API void igPushItemFlag(ImGuiItemFlags option,bool enabled);
func PushItemFlag(option ImGuiItemFlags, enabled bool) {
	_enabled := c.CBool(enabled)
	giLib.Call(_func_igPushItemFlag_, []interface{}{&option, &_enabled})
}

var _func_igPopItemFlag_ = &c.FuncPrototype{Name: "igPopItemFlag", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igPopItemFlag(void);
func PopItemFlag() {
	giLib.Call(_func_igPopItemFlag_, nil)
}

var _func_igGetStyleVarInfo_ = &c.FuncPrototype{Name: "igGetStyleVarInfo", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API const ImGuiDataVarInfo* igGetStyleVarInfo(ImGuiStyleVar idx);
func GetStyleVarInfo(idx ImGuiStyleVar) *ImGuiDataVarInfo {
	return (*ImGuiDataVarInfo)(giLib.Call(_func_igGetStyleVarInfo_, []interface{}{&idx}).PtrFree())
}

var _func_igLogBegin_ = &c.FuncPrototype{Name: "igLogBegin", OutType: c.Void, InTypes: _typs_U32I32}

// CIMGUI_API void igLogBegin(ImGuiLogType typ,int auto_open_depth);
func LogBegin(typ ImGuiLogType, auto_open_depth int32) {
	giLib.Call(_func_igLogBegin_, []interface{}{&typ, &auto_open_depth})
}

var _func_igLogToBuffer_ = &c.FuncPrototype{Name: "igLogToBuffer", OutType: c.Void, InTypes: _typs_I32}

// CIMGUI_API void igLogToBuffer(int auto_open_depth);
func LogToBuffer(auto_open_depth int32) {
	giLib.Call(_func_igLogToBuffer_, []interface{}{&auto_open_depth})
}

var _func_igLogRenderedText_ = &c.FuncPrototype{Name: "igLogRenderedText", OutType: c.Void, InTypes: _typs_PPP}

// CIMGUI_API void igLogRenderedText(const ImVec2* ref_pos,const char* text,const char* text_end);
func LogRenderedText(ref_pos *ImVec2, text string, text_end string) {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	giLib.Call(_func_igLogRenderedText_, []interface{}{&ref_pos, &_text, &_text_end})
}

var _func_igLogSetNextTextDecoration_ = &c.FuncPrototype{Name: "igLogSetNextTextDecoration", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igLogSetNextTextDecoration(const char* prefix,const char* suffix);
func LogSetNextTextDecoration(prefix string, suffix string) {
	_prefix, _suffix := c.CStr(prefix), c.CStr(suffix)
	defer usf.Free(_prefix)
	defer usf.Free(_suffix)
	giLib.Call(_func_igLogSetNextTextDecoration_, []interface{}{&_prefix, &_suffix})
}

var _func_igBeginChildEx_ = &c.FuncPrototype{Name: "igBeginChildEx", OutType: c.I32, InTypes: _typs_PU32Vec2I32U32}

// CIMGUI_API bool igBeginChildEx(const char* name,ImGuiID id,const ImVec2 size_arg,bool border,ImGuiWindowFlags flags);
func BeginChildEx(name string, id ImGuiID, size_arg ImVec2, border bool, flags ImGuiWindowFlags) bool {
	_name, _border := c.CStr(name), c.CBool(border)
	defer usf.Free(_name)

	return giLib.Call(_func_igBeginChildEx_, []interface{}{&_name, &id, &size_arg, &_border, &flags}).BoolFree()
}

var _func_igOpenPopupEx_ = &c.FuncPrototype{Name: "igOpenPopupEx", OutType: c.Void, InTypes: _typs_U32U32}

// CIMGUI_API void igOpenPopupEx(ImGuiID id,ImGuiPopupFlags popup_flags);
func OpenPopupEx(id ImGuiID, popup_flags ImGuiPopupFlags) {
	giLib.Call(_func_igOpenPopupEx_, []interface{}{&id, &popup_flags})
}

var _func_igClosePopupToLevel_ = &c.FuncPrototype{Name: "igClosePopupToLevel", OutType: c.Void, InTypes: _typs_I32I32}

// CIMGUI_API void igClosePopupToLevel(int remaining,bool restore_focus_to_window_under_popup);
func ClosePopupToLevel(remaining int32, restore_focus_to_window_under_popup bool) {
	_restore_focus_to_window_under_popup := c.CBool(restore_focus_to_window_under_popup)
	giLib.Call(_func_igClosePopupToLevel_, []interface{}{&remaining, &_restore_focus_to_window_under_popup})
}

var _func_igClosePopupsOverWindow_ = &c.FuncPrototype{Name: "igClosePopupsOverWindow", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void igClosePopupsOverWindow(ImGuiWindow* ref_window,bool restore_focus_to_window_under_popup);
func ClosePopupsOverWindow(ref_window *ImGuiWindow, restore_focus_to_window_under_popup bool) {
	_restore_focus_to_window_under_popup := c.CBool(restore_focus_to_window_under_popup)
	giLib.Call(_func_igClosePopupsOverWindow_, []interface{}{&ref_window, &_restore_focus_to_window_under_popup})
}

var _func_igClosePopupsExceptModals_ = &c.FuncPrototype{Name: "igClosePopupsExceptModals", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igClosePopupsExceptModals(void);
func ClosePopupsExceptModals() {
	giLib.Call(_func_igClosePopupsExceptModals_, nil)
}

var _func_igIsPopupOpen_ID_ = &c.FuncPrototype{Name: "igIsPopupOpen_ID", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igIsPopupOpen_ID(ImGuiID id,ImGuiPopupFlags popup_flags);
func IsPopupOpenID(id ImGuiID, popup_flags ImGuiPopupFlags) bool {

	return giLib.Call(_func_igIsPopupOpen_ID_, []interface{}{&id, &popup_flags}).BoolFree()
}

var _func_igBeginPopupEx_ = &c.FuncPrototype{Name: "igBeginPopupEx", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igBeginPopupEx(ImGuiID id,ImGuiWindowFlags extra_flags);
func BeginPopupEx(id ImGuiID, extra_flags ImGuiWindowFlags) bool {
	return giLib.Call(_func_igBeginPopupEx_, []interface{}{&id, &extra_flags}).BoolFree()
}

var _func_igBeginTooltipEx_ = &c.FuncPrototype{Name: "igBeginTooltipEx", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igBeginTooltipEx(ImGuiTooltipFlags tooltip_flags,ImGuiWindowFlags extra_window_flags);
func BeginTooltipEx(tooltip_flags ImGuiTooltipFlags, extra_window_flags ImGuiWindowFlags) bool {
	return giLib.Call(_func_igBeginTooltipEx_, []interface{}{&tooltip_flags, &extra_window_flags}).BoolFree()
}

var _func_igGetPopupAllowedExtentRect_ = &c.FuncPrototype{Name: "igGetPopupAllowedExtentRect", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igGetPopupAllowedExtentRect(ImRect* pOut,ImGuiWindow* window);
func GetPopupAllowedExtentRect(pOut *ImRect, window *ImGuiWindow) {
	giLib.Call(_func_igGetPopupAllowedExtentRect_, []interface{}{&pOut, &window})
}

var _func_igGetTopMostPopupModal_ = &c.FuncPrototype{Name: "igGetTopMostPopupModal", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiWindow* igGetTopMostPopupModal(void);
func GetTopMostPopupModal() *ImGuiWindow {
	return (*ImGuiWindow)(giLib.Call(_func_igGetTopMostPopupModal_, nil).PtrFree())
}

var _func_igGetTopMostAndVisiblePopupModal_ = &c.FuncPrototype{Name: "igGetTopMostAndVisiblePopupModal", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiWindow* igGetTopMostAndVisiblePopupModal(void);
func GetTopMostAndVisiblePopupModal() *ImGuiWindow {
	return (*ImGuiWindow)(giLib.Call(_func_igGetTopMostAndVisiblePopupModal_, nil).PtrFree())
}

var _func_igFindBlockingModal_ = &c.FuncPrototype{Name: "igFindBlockingModal", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiWindow* igFindBlockingModal(ImGuiWindow* window);
func FindBlockingModal(window *ImGuiWindow) *ImGuiWindow {
	return (*ImGuiWindow)(giLib.Call(_func_igFindBlockingModal_, []interface{}{&window}).PtrFree())
}

var _func_igFindBestWindowPosForPopup_ = &c.FuncPrototype{Name: "igFindBestWindowPosForPopup", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igFindBestWindowPosForPopup(ImVec2* pOut,ImGuiWindow* window);
func FindBestWindowPosForPopup(pOut *ImVec2, window *ImGuiWindow) {
	giLib.Call(_func_igFindBestWindowPosForPopup_, []interface{}{&pOut, &window})
}

var _func_igFindBestWindowPosForPopupEx_ = &c.FuncPrototype{Name: "igFindBestWindowPosForPopupEx", OutType: c.Void, InTypes: _typs_PVec2Vec2PRectRectU32}

// CIMGUI_API void igFindBestWindowPosForPopupEx(ImVec2* pOut,const ImVec2 ref_pos,const ImVec2 size,ImGuiDir* last_dir,const ImRect r_outer,const ImRect r_avoid,ImGuiPopupPositionPolicy policy);
func FindBestWindowPosForPopupEx(pOut *ImVec2, ref_pos ImVec2, size ImVec2, last_dir *ImGuiDir, r_outer ImRect, r_avoid ImRect, policy ImGuiPopupPositionPolicy) {
	giLib.Call(_func_igFindBestWindowPosForPopupEx_, []interface{}{&pOut, &ref_pos, &size, &last_dir, &r_outer, &r_avoid, &policy})
}

var _func_igBeginViewportSideBar_ = &c.FuncPrototype{Name: "igBeginViewportSideBar", OutType: c.I32, InTypes: _typs_PPI32F32U32}

// CIMGUI_API bool igBeginViewportSideBar(const char* name,ImGuiViewport* viewport,ImGuiDir dir,float size,ImGuiWindowFlags window_flags);
func BeginViewportSideBar(name string, viewport *ImGuiViewport, dir ImGuiDir, size float32, window_flags ImGuiWindowFlags) bool {
	_name := c.CStr(name)
	defer usf.Free(_name)
	return giLib.Call(_func_igBeginViewportSideBar_, []interface{}{&_name, &viewport, &dir, &size, &window_flags}).BoolFree()
}

var _func_igBeginMenuEx_ = &c.FuncPrototype{Name: "igBeginMenuEx", OutType: c.I32, InTypes: _typs_PPI32}

// CIMGUI_API bool igBeginMenuEx(const char* label,const char* icon,bool enabled);
func BeginMenuEx(label string, icon string, enabled bool) bool {
	_label, _icon, _enabled := c.CStr(label), c.CStr(icon), c.CBool(enabled)
	defer usf.Free(_label)
	defer usf.Free(_icon)

	return giLib.Call(_func_igBeginMenuEx_, []interface{}{&_label, &_icon, &_enabled}).BoolFree()
}

var _func_igMenuItemEx_ = &c.FuncPrototype{Name: "igMenuItemEx", OutType: c.I32, InTypes: _typs_PPPI32I32}

// CIMGUI_API bool igMenuItemEx(const char* label,const char* icon,const char* shortcut,bool selected,bool enabled);
func MenuItemEx(label string, icon string, shortcut string, selected bool, enabled bool) bool {
	_label, _icon, _shortcut, _selected, _enabled := c.CStr(label), c.CStr(icon), c.CStr(shortcut), c.CBool(selected), c.CBool(enabled)
	defer usf.Free(_label)
	defer usf.Free(_icon)
	defer usf.Free(_shortcut)
	return giLib.Call(_func_igMenuItemEx_, []interface{}{&_label, &_icon, &_shortcut, &_selected, &_enabled}).BoolFree()
}

var _func_igBeginComboPopup_ = &c.FuncPrototype{Name: "igBeginComboPopup", OutType: c.I32, InTypes: _typs_U32RectU32}

// CIMGUI_API bool igBeginComboPopup(ImGuiID popup_id,const ImRect bb,ImGuiComboFlags flags);
func BeginComboPopup(popup_id ImGuiID, bb ImRect, flags ImGuiComboFlags) bool {
	return giLib.Call(_func_igBeginComboPopup_, []interface{}{&popup_id, &bb, &flags}).BoolFree()
}

var _func_igBeginComboPreview_ = &c.FuncPrototype{Name: "igBeginComboPreview", OutType: c.I32, InTypes: nil}

// CIMGUI_API bool igBeginComboPreview(void);
func BeginComboPreview() bool {
	return giLib.Call(_func_igBeginComboPreview_, nil).BoolFree()
}

var _func_igEndComboPreview_ = &c.FuncPrototype{Name: "igEndComboPreview", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igEndComboPreview(void);
func EndComboPreview() {
	giLib.Call(_func_igEndComboPreview_, nil)
}

var _func_igNavInitWindow_ = &c.FuncPrototype{Name: "igNavInitWindow", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void igNavInitWindow(ImGuiWindow* window,bool force_reinit);
func NavInitWindow(window *ImGuiWindow, force_reinit bool) {
	_force_reinit := c.CBool(force_reinit)
	giLib.Call(_func_igNavInitWindow_, []interface{}{&window, &_force_reinit})
}

var _func_igNavInitRequestApplyResult_ = &c.FuncPrototype{Name: "igNavInitRequestApplyResult", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igNavInitRequestApplyResult(void);
func NavInitRequestApplyResult() {
	giLib.Call(_func_igNavInitRequestApplyResult_, nil)
}

var _func_igNavMoveRequestButNoResultYet_ = &c.FuncPrototype{Name: "igNavMoveRequestButNoResultYet", OutType: c.I32, InTypes: nil}

// CIMGUI_API bool igNavMoveRequestButNoResultYet(void);
func NavMoveRequestButNoResultYet() bool {
	return giLib.Call(_func_igNavMoveRequestButNoResultYet_, nil).BoolFree()
}

var _func_igNavMoveRequestSubmit_ = &c.FuncPrototype{Name: "igNavMoveRequestSubmit", OutType: c.Void, InTypes: _typs_I32I32U32U32}

// CIMGUI_API void igNavMoveRequestSubmit(ImGuiDir move_dir,ImGuiDir clip_dir,ImGuiNavMoveFlags move_flags,ImGuiScrollFlags scroll_flags);
func NavMoveRequestSubmit(move_dir ImGuiDir, clip_dir ImGuiDir, move_flags ImGuiNavMoveFlags, scroll_flags ImGuiScrollFlags) {
	giLib.Call(_func_igNavMoveRequestSubmit_, []interface{}{&move_dir, &clip_dir, &move_flags, &scroll_flags})
}

var _func_igNavMoveRequestForward_ = &c.FuncPrototype{Name: "igNavMoveRequestForward", OutType: c.Void, InTypes: _typs_I32I32U32U32}

// CIMGUI_API void igNavMoveRequestForward(ImGuiDir move_dir,ImGuiDir clip_dir,ImGuiNavMoveFlags move_flags,ImGuiScrollFlags scroll_flags);
func NavMoveRequestForward(move_dir ImGuiDir, clip_dir ImGuiDir, move_flags ImGuiNavMoveFlags, scroll_flags ImGuiScrollFlags) {
	giLib.Call(_func_igNavMoveRequestForward_, []interface{}{&move_dir, &clip_dir, &move_flags, &scroll_flags})
}

var _func_igNavMoveRequestResolveWithLastItem_ = &c.FuncPrototype{Name: "igNavMoveRequestResolveWithLastItem", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igNavMoveRequestResolveWithLastItem(ImGuiNavItemData* result);
func NavMoveRequestResolveWithLastItem(result *ImGuiNavItemData) {
	giLib.Call(_func_igNavMoveRequestResolveWithLastItem_, []interface{}{&result})
}

var _func_igNavMoveRequestCancel_ = &c.FuncPrototype{Name: "igNavMoveRequestCancel", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igNavMoveRequestCancel(void);
func NavMoveRequestCancel() {
	giLib.Call(_func_igNavMoveRequestCancel_, nil)
}

var _func_igNavMoveRequestApplyResult_ = &c.FuncPrototype{Name: "igNavMoveRequestApplyResult", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igNavMoveRequestApplyResult(void);
func NavMoveRequestApplyResult() {
	giLib.Call(_func_igNavMoveRequestApplyResult_, nil)
}

var _func_igNavMoveRequestTryWrapping_ = &c.FuncPrototype{Name: "igNavMoveRequestTryWrapping", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void igNavMoveRequestTryWrapping(ImGuiWindow* window,ImGuiNavMoveFlags move_flags);
func NavMoveRequestTryWrapping(window *ImGuiWindow, move_flags ImGuiNavMoveFlags) {
	giLib.Call(_func_igNavMoveRequestTryWrapping_, []interface{}{&window, &move_flags})
}

var _func_igNavClearPreferredPosForAxis_ = &c.FuncPrototype{Name: "igNavClearPreferredPosForAxis", OutType: c.Void, InTypes: _typs_I32}

// CIMGUI_API void igNavClearPreferredPosForAxis(ImGuiAxis axis);
func NavClearPreferredPosForAxis(axis ImGuiAxis) {
	giLib.Call(_func_igNavClearPreferredPosForAxis_, []interface{}{&axis})
}

var _func_igNavUpdateCurrentWindowIsScrollPushableX_ = &c.FuncPrototype{Name: "igNavUpdateCurrentWindowIsScrollPushableX", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igNavUpdateCurrentWindowIsScrollPushableX(void);
func NavUpdateCurrentWindowIsScrollPushableX() {
	giLib.Call(_func_igNavUpdateCurrentWindowIsScrollPushableX_, nil)
}

var _func_igSetNavWindow_ = &c.FuncPrototype{Name: "igSetNavWindow", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igSetNavWindow(ImGuiWindow* window);
func SetNavWindow(window *ImGuiWindow) {
	giLib.Call(_func_igSetNavWindow_, []interface{}{&window})
}

var _func_igSetNavID_ = &c.FuncPrototype{Name: "igSetNavID", OutType: c.Void, InTypes: _typs_U32U32U32Rect}

// CIMGUI_API void igSetNavID(ImGuiID id,ImGuiNavLayer nav_layer,ImGuiID focus_scope_id,const ImRect rect_rel);
func SetNavID(id ImGuiID, nav_layer ImGuiNavLayer, focus_scope_id ImGuiID, rect_rel ImRect) {
	giLib.Call(_func_igSetNavID_, []interface{}{&id, &nav_layer, &focus_scope_id, &rect_rel})
}

var _func_igFocusItem_ = &c.FuncPrototype{Name: "igFocusItem", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igFocusItem(void);
func FocusItem() {
	giLib.Call(_func_igFocusItem_, nil)
}

var _func_igActivateItemByID_ = &c.FuncPrototype{Name: "igActivateItemByID", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igActivateItemByID(ImGuiID id);
func ActivateItemByID(id ImGuiID) {
	giLib.Call(_func_igActivateItemByID_, []interface{}{&id})
}

var _func_igIsNamedKey_ = &c.FuncPrototype{Name: "igIsNamedKey", OutType: c.I32, InTypes: _typs_U32}

// CIMGUI_API bool igIsNamedKey(ImGuiKey key);
func IsNamedKey(key ImGuiKey) bool {
	return giLib.Call(_func_igIsNamedKey_, []interface{}{&key}).BoolFree()
}

var _func_igIsNamedKeyOrModKey_ = &c.FuncPrototype{Name: "igIsNamedKeyOrModKey", OutType: c.I32, InTypes: _typs_U32}

// CIMGUI_API bool igIsNamedKeyOrModKey(ImGuiKey key);
func IsNamedKeyOrModKey(key ImGuiKey) bool {
	return giLib.Call(_func_igIsNamedKeyOrModKey_, []interface{}{&key}).BoolFree()
}

var _func_igIsLegacyKey_ = &c.FuncPrototype{Name: "igIsLegacyKey", OutType: c.I32, InTypes: _typs_U32}

// CIMGUI_API bool igIsLegacyKey(ImGuiKey key);
func IsLegacyKey(key ImGuiKey) bool {
	return giLib.Call(_func_igIsLegacyKey_, []interface{}{&key}).BoolFree()
}

var _func_igIsKeyboardKey_ = &c.FuncPrototype{Name: "igIsKeyboardKey", OutType: c.I32, InTypes: _typs_U32}

// CIMGUI_API bool igIsKeyboardKey(ImGuiKey key);
func IsKeyboardKey(key ImGuiKey) bool {
	return giLib.Call(_func_igIsKeyboardKey_, []interface{}{&key}).BoolFree()
}

var _func_igIsGamepadKey_ = &c.FuncPrototype{Name: "igIsGamepadKey", OutType: c.I32, InTypes: _typs_U32}

// CIMGUI_API bool igIsGamepadKey(ImGuiKey key);
func IsGamepadKey(key ImGuiKey) bool {
	return giLib.Call(_func_igIsGamepadKey_, []interface{}{&key}).BoolFree()
}

var _func_igIsMouseKey_ = &c.FuncPrototype{Name: "igIsMouseKey", OutType: c.I32, InTypes: _typs_U32}

// CIMGUI_API bool igIsMouseKey(ImGuiKey key);
func IsMouseKey(key ImGuiKey) bool {
	return giLib.Call(_func_igIsMouseKey_, []interface{}{&key}).BoolFree()
}

var _func_igIsAliasKey_ = &c.FuncPrototype{Name: "igIsAliasKey", OutType: c.I32, InTypes: _typs_U32}

// CIMGUI_API bool igIsAliasKey(ImGuiKey key);
func IsAliasKey(key ImGuiKey) bool {
	return giLib.Call(_func_igIsAliasKey_, []interface{}{&key}).BoolFree()
}

var _func_igConvertShortcutMod_ = &c.FuncPrototype{Name: "igConvertShortcutMod", OutType: c.I32, InTypes: _typs_I32}

// CIMGUI_API ImGuiKeyChord igConvertShortcutMod(ImGuiKeyChord key_chord);
func ConvertShortcutMod(key_chord ImGuiKeyChord) ImGuiKeyChord {
	return (ImGuiKeyChord)(giLib.Call(_func_igConvertShortcutMod_, []interface{}{&key_chord}).I32Free())
}

var _func_igConvertSingleModFlagToKey_ = &c.FuncPrototype{Name: "igConvertSingleModFlagToKey", OutType: c.U32, InTypes: _typs_PU32}

// CIMGUI_API ImGuiKey igConvertSingleModFlagToKey(ImGuiContext* ctx,ImGuiKey key);
func ConvertSingleModFlagToKey(ctx *ImGuiContext, key ImGuiKey) ImGuiKey {
	return (ImGuiKey)(giLib.Call(_func_igConvertSingleModFlagToKey_, []interface{}{&ctx, &key}).U32Free())
}

var _func_igGetKeyData_ContextPtr_ = &c.FuncPrototype{Name: "igGetKeyData_ContextPtr", OutType: c.Pointer, InTypes: _typs_PU32}

// CIMGUI_API ImGuiKeyData* igGetKeyData_ContextPtr(ImGuiContext* ctx,ImGuiKey key);
func GetKeyDataContextPtr(ctx *ImGuiContext, key ImGuiKey) *ImGuiKeyData {
	return (*ImGuiKeyData)(giLib.Call(_func_igGetKeyData_ContextPtr_, []interface{}{&ctx, &key}).PtrFree())
}

var _func_igGetKeyData_Key_ = &c.FuncPrototype{Name: "igGetKeyData_Key", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiKeyData* igGetKeyData_Key(ImGuiKey key);
func GetKeyData_Key(key ImGuiKey) *ImGuiKeyData {
	return (*ImGuiKeyData)(giLib.Call(_func_igGetKeyData_Key_, []interface{}{&key}).PtrFree())
}

var _func_igGetKeyChordName_ = &c.FuncPrototype{Name: "igGetKeyChordName", OutType: c.Void, InTypes: _typs_I32PI32}

// CIMGUI_API void igGetKeyChordName(ImGuiKeyChord key_chord,char* out_buf,int out_buf_size);
func GetKeyChordName(key_chord ImGuiKeyChord, out_buf string, out_buf_size int32) {
	_out_buf := c.CStr(out_buf)
	defer usf.Free(_out_buf)
	giLib.Call(_func_igGetKeyChordName_, []interface{}{&key_chord, &_out_buf, &out_buf_size})
}

var _func_igMouseButtonToKey_ = &c.FuncPrototype{Name: "igMouseButtonToKey", OutType: c.U32, InTypes: _typs_U32}

// CIMGUI_API ImGuiKey igMouseButtonToKey(ImGuiMouseButton button);
func MouseButtonToKey(button ImGuiMouseButton) ImGuiKey {
	return (ImGuiKey)(giLib.Call(_func_igMouseButtonToKey_, []interface{}{&button}).U32Free())
}

var _func_igIsMouseDragPastThreshold_ = &c.FuncPrototype{Name: "igIsMouseDragPastThreshold", OutType: c.I32, InTypes: _typs_U32F32}

// CIMGUI_API bool igIsMouseDragPastThreshold(ImGuiMouseButton button,float lock_threshold);
func IsMouseDragPastThreshold(button ImGuiMouseButton, lock_threshold float32) bool {
	return giLib.Call(_func_igIsMouseDragPastThreshold_, []interface{}{&button, &lock_threshold}).BoolFree()
}

var _func_igGetKeyMagnitude2d_ = &c.FuncPrototype{Name: "igGetKeyMagnitude2d", OutType: c.Void, InTypes: _typs_PU32U32U32U32}

// CIMGUI_API void igGetKeyMagnitude2d(ImVec2* pOut,ImGuiKey key_left,ImGuiKey key_right,ImGuiKey key_up,ImGuiKey key_down);
func GetKeyMagnitude2d(pOut *ImVec2, key_left ImGuiKey, key_right ImGuiKey, key_up ImGuiKey, key_down ImGuiKey) {
	giLib.Call(_func_igGetKeyMagnitude2d_, []interface{}{&pOut, &key_left, &key_right, &key_up, &key_down})
}

var _func_igGetNavTweakPressedAmount_ = &c.FuncPrototype{Name: "igGetNavTweakPressedAmount", OutType: c.F32, InTypes: _typs_I32}

// CIMGUI_API float igGetNavTweakPressedAmount(ImGuiAxis axis);
func GetNavTweakPressedAmount(axis ImGuiAxis) float32 {
	return giLib.Call(_func_igGetNavTweakPressedAmount_, []interface{}{&axis}).F32Free()
}

var _func_igCalcTypematicRepeatAmount_ = &c.FuncPrototype{Name: "igCalcTypematicRepeatAmount", OutType: c.I32, InTypes: _typs_F32F32F32F32}

// CIMGUI_API int igCalcTypematicRepeatAmount(float t0,float t1,float repeat_delay,float repeat_rate);
func CalcTypematicRepeatAmount(t0 float32, t1 float32, repeat_delay float32, repeat_rate float32) int32 {
	return giLib.Call(_func_igCalcTypematicRepeatAmount_, []interface{}{&t0, &t1, &repeat_delay, &repeat_rate}).I32Free()
}

var _func_igGetTypematicRepeatRate_ = &c.FuncPrototype{Name: "igGetTypematicRepeatRate", OutType: c.Void, InTypes: _typs_U32PP}

// CIMGUI_API void igGetTypematicRepeatRate(ImGuiInputFlags flags,float* repeat_delay,float* repeat_rate);
func GetTypematicRepeatRate(flags ImGuiInputFlags, repeat_delay *float32, repeat_rate *float32) {
	giLib.Call(_func_igGetTypematicRepeatRate_, []interface{}{&flags, &repeat_delay, &repeat_rate})
}

var _func_igSetActiveIdUsingAllKeyboardKeys_ = &c.FuncPrototype{Name: "igSetActiveIdUsingAllKeyboardKeys", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igSetActiveIdUsingAllKeyboardKeys(void);
func SetActiveIdUsingAllKeyboardKeys() {
	giLib.Call(_func_igSetActiveIdUsingAllKeyboardKeys_, nil)
}

var _func_igIsActiveIdUsingNavDir_ = &c.FuncPrototype{Name: "igIsActiveIdUsingNavDir", OutType: c.I32, InTypes: _typs_I32}

// CIMGUI_API bool igIsActiveIdUsingNavDir(ImGuiDir dir);
func IsActiveIdUsingNavDir(dir ImGuiDir) bool {
	return giLib.Call(_func_igIsActiveIdUsingNavDir_, []interface{}{&dir}).BoolFree()
}

var _func_igGetKeyOwner_ = &c.FuncPrototype{Name: "igGetKeyOwner", OutType: c.U32, InTypes: _typs_U32}

// CIMGUI_API ImGuiID igGetKeyOwner(ImGuiKey key);
func GetKeyOwner(key ImGuiKey) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetKeyOwner_, []interface{}{&key}).U32Free())
}

var _func_igSetKeyOwner_ = &c.FuncPrototype{Name: "igSetKeyOwner", OutType: c.Void, InTypes: _typs_U32U32U32}

// CIMGUI_API void igSetKeyOwner(ImGuiKey key,ImGuiID owner_id,ImGuiInputFlags flags);
func SetKeyOwner(key ImGuiKey, owner_id ImGuiID, flags ImGuiInputFlags) {
	giLib.Call(_func_igSetKeyOwner_, []interface{}{&key, &owner_id, &flags})
}

var _func_igSetKeyOwnersForKeyChord_ = &c.FuncPrototype{Name: "igSetKeyOwnersForKeyChord", OutType: c.Void, InTypes: _typs_I32U32U32}

// CIMGUI_API void igSetKeyOwnersForKeyChord(ImGuiKeyChord key,ImGuiID owner_id,ImGuiInputFlags flags);
func SetKeyOwnersForKeyChord(key ImGuiKeyChord, owner_id ImGuiID, flags ImGuiInputFlags) {
	giLib.Call(_func_igSetKeyOwnersForKeyChord_, []interface{}{&key, &owner_id, &flags})
}

var _func_igSetItemKeyOwner_ = &c.FuncPrototype{Name: "igSetItemKeyOwner", OutType: c.Void, InTypes: _typs_U32U32}

// CIMGUI_API void igSetItemKeyOwner(ImGuiKey key,ImGuiInputFlags flags);
func SetItemKeyOwner(key ImGuiKey, flags ImGuiInputFlags) {
	giLib.Call(_func_igSetItemKeyOwner_, []interface{}{&key, &flags})
}

var _func_igTestKeyOwner_ = &c.FuncPrototype{Name: "igTestKeyOwner", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igTestKeyOwner(ImGuiKey key,ImGuiID owner_id);
func TestKeyOwner(key ImGuiKey, owner_id ImGuiID) bool {
	return giLib.Call(_func_igTestKeyOwner_, []interface{}{&key, &owner_id}).BoolFree()
}

var _func_igGetKeyOwnerData_ = &c.FuncPrototype{Name: "igGetKeyOwnerData", OutType: c.Pointer, InTypes: _typs_PU32}

// CIMGUI_API ImGuiKeyOwnerData* igGetKeyOwnerData(ImGuiContext* ctx,ImGuiKey key);
func GetKeyOwnerData(ctx *ImGuiContext, key ImGuiKey) *ImGuiKeyOwnerData {
	return (*ImGuiKeyOwnerData)(giLib.Call(_func_igGetKeyOwnerData_, []interface{}{&ctx, &key}).PtrFree())
}

var _func_igIsKeyDown_ID_ = &c.FuncPrototype{Name: "igIsKeyDown_ID", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igIsKeyDown_ID(ImGuiKey key,ImGuiID owner_id);
func IsKeyDownID(key ImGuiKey, owner_id ImGuiID) bool {
	return giLib.Call(_func_igIsKeyDown_ID_, []interface{}{&key, &owner_id}).BoolFree()
}

var _func_igIsKeyPressed_ID_ = &c.FuncPrototype{Name: "igIsKeyPressed_ID", OutType: c.I32, InTypes: _typs_U32U32U32}

// CIMGUI_API bool igIsKeyPressed_ID(ImGuiKey key,ImGuiID owner_id,ImGuiInputFlags flags);
func IsKeyPressedID(key ImGuiKey, owner_id ImGuiID, flags ImGuiInputFlags) bool {
	return giLib.Call(_func_igIsKeyPressed_ID_, []interface{}{&key, &owner_id, &flags}).BoolFree()
}

var _func_igIsKeyReleased_ID_ = &c.FuncPrototype{Name: "igIsKeyReleased_ID", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igIsKeyReleased_ID(ImGuiKey key,ImGuiID owner_id);
func IsKeyReleasedID(key ImGuiKey, owner_id ImGuiID) bool {
	return giLib.Call(_func_igIsKeyReleased_ID_, []interface{}{&key, &owner_id}).BoolFree()
}

var _func_igIsMouseDown_ID_ = &c.FuncPrototype{Name: "igIsMouseDown_ID", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igIsMouseDown_ID(ImGuiMouseButton button,ImGuiID owner_id);
func IsMouseDownID(button ImGuiMouseButton, owner_id ImGuiID) bool {
	return giLib.Call(_func_igIsMouseDown_ID_, []interface{}{&button, &owner_id}).BoolFree()
}

var _func_igIsMouseClicked_ID_ = &c.FuncPrototype{Name: "igIsMouseClicked_ID", OutType: c.I32, InTypes: _typs_U32U32U32}

// CIMGUI_API bool igIsMouseClicked_ID(ImGuiMouseButton button,ImGuiID owner_id,ImGuiInputFlags flags);
func IsMouseClickedID(button ImGuiMouseButton, owner_id ImGuiID, flags ImGuiInputFlags) bool {
	return giLib.Call(_func_igIsMouseClicked_ID_, []interface{}{&button, &owner_id, &flags}).BoolFree()
}

var _func_igIsMouseReleased_ID_ = &c.FuncPrototype{Name: "igIsMouseReleased_ID", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igIsMouseReleased_ID(ImGuiMouseButton button,ImGuiID owner_id);
func IsMouseReleasedID(button ImGuiMouseButton, owner_id ImGuiID) bool {
	return giLib.Call(_func_igIsMouseReleased_ID_, []interface{}{&button, &owner_id}).BoolFree()
}

var _func_igShortcut_ = &c.FuncPrototype{Name: "igShortcut", OutType: c.I32, InTypes: _typs_I32U32U32}

// CIMGUI_API bool igShortcut(ImGuiKeyChord key_chord,ImGuiID owner_id,ImGuiInputFlags flags);
func Shortcut(key_chord ImGuiKeyChord, owner_id ImGuiID, flags ImGuiInputFlags) bool {
	return giLib.Call(_func_igShortcut_, []interface{}{&key_chord, &owner_id, &flags}).BoolFree()
}

var _func_igSetShortcutRouting_ = &c.FuncPrototype{Name: "igSetShortcutRouting", OutType: c.I32, InTypes: _typs_I32U32U32}

// CIMGUI_API bool igSetShortcutRouting(ImGuiKeyChord key_chord,ImGuiID owner_id,ImGuiInputFlags flags);
func SetShortcutRouting(key_chord ImGuiKeyChord, owner_id ImGuiID, flags ImGuiInputFlags) bool {
	return giLib.Call(_func_igSetShortcutRouting_, []interface{}{&key_chord, &owner_id, &flags}).BoolFree()
}

var _func_igTestShortcutRouting_ = &c.FuncPrototype{Name: "igTestShortcutRouting", OutType: c.I32, InTypes: _typs_I32U32}

// CIMGUI_API bool igTestShortcutRouting(ImGuiKeyChord key_chord,ImGuiID owner_id);
func TestShortcutRouting(key_chord ImGuiKeyChord, owner_id ImGuiID) bool {
	return giLib.Call(_func_igTestShortcutRouting_, []interface{}{&key_chord, &owner_id}).BoolFree()
}

var _func_igGetShortcutRoutingData_ = &c.FuncPrototype{Name: "igGetShortcutRoutingData", OutType: c.Pointer, InTypes: _typs_I32}

// CIMGUI_API ImGuiKeyRoutingData* igGetShortcutRoutingData(ImGuiKeyChord key_chord);
func GetShortcutRoutingData(key_chord ImGuiKeyChord) *ImGuiKeyRoutingData {
	return (*ImGuiKeyRoutingData)(giLib.Call(_func_igGetShortcutRoutingData_, []interface{}{&key_chord}).PtrFree())
}

var _func_igDockContextInitialize_ = &c.FuncPrototype{Name: "igDockContextInitialize", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDockContextInitialize(ImGuiContext* ctx);
func DockContextInitialize(ctx *ImGuiContext) {
	giLib.Call(_func_igDockContextInitialize_, []interface{}{&ctx})
}

var _func_igDockContextShutdown_ = &c.FuncPrototype{Name: "igDockContextShutdown", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDockContextShutdown(ImGuiContext* ctx);
func DockContextShutdown(ctx *ImGuiContext) {
	giLib.Call(_func_igDockContextShutdown_, []interface{}{&ctx})
}

var _func_igDockContextClearNodes_ = &c.FuncPrototype{Name: "igDockContextClearNodes", OutType: c.Void, InTypes: _typs_PU32I32}

// CIMGUI_API void igDockContextClearNodes(ImGuiContext* ctx,ImGuiID root_id,bool clear_settings_refs);
func DockContextClearNodes(ctx *ImGuiContext, root_id ImGuiID, clear_settings_refs bool) {
	_clear_settings_refs := c.CBool(clear_settings_refs)
	giLib.Call(_func_igDockContextClearNodes_, []interface{}{&ctx, &root_id, &_clear_settings_refs})
}

var _func_igDockContextRebuildNodes_ = &c.FuncPrototype{Name: "igDockContextRebuildNodes", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDockContextRebuildNodes(ImGuiContext* ctx);
func DockContextRebuildNodes(ctx *ImGuiContext) {
	giLib.Call(_func_igDockContextRebuildNodes_, []interface{}{&ctx})
}

var _func_igDockContextNewFrameUpdateUndocking_ = &c.FuncPrototype{Name: "igDockContextNewFrameUpdateUndocking", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDockContextNewFrameUpdateUndocking(ImGuiContext* ctx);
func DockContextNewFrameUpdateUndocking(ctx *ImGuiContext) {
	giLib.Call(_func_igDockContextNewFrameUpdateUndocking_, []interface{}{&ctx})
}

var _func_igDockContextNewFrameUpdateDocking_ = &c.FuncPrototype{Name: "igDockContextNewFrameUpdateDocking", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDockContextNewFrameUpdateDocking(ImGuiContext* ctx);
func DockContextNewFrameUpdateDocking(ctx *ImGuiContext) {
	giLib.Call(_func_igDockContextNewFrameUpdateDocking_, []interface{}{&ctx})
}

var _func_igDockContextEndFrame_ = &c.FuncPrototype{Name: "igDockContextEndFrame", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDockContextEndFrame(ImGuiContext* ctx);
func DockContextEndFrame(ctx *ImGuiContext) {
	giLib.Call(_func_igDockContextEndFrame_, []interface{}{&ctx})
}

var _func_igDockContextGenNodeID_ = &c.FuncPrototype{Name: "igDockContextGenNodeID", OutType: c.U32, InTypes: _typs_P}

// CIMGUI_API ImGuiID igDockContextGenNodeID(ImGuiContext* ctx);
func DockContextGenNodeID(ctx *ImGuiContext) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igDockContextGenNodeID_, []interface{}{&ctx}).U32Free())
}

var _func_igDockContextQueueDock_ = &c.FuncPrototype{Name: "igDockContextQueueDock", OutType: c.Void, InTypes: _typs_PPPPI32F32I32}

// CIMGUI_API void igDockContextQueueDock(ImGuiContext* ctx,ImGuiWindow* target,ImGuiDockNode* target_node,ImGuiWindow* payload,ImGuiDir split_dir,float split_ratio,bool split_outer);
func DockContextQueueDock(ctx *ImGuiContext, target *ImGuiWindow, target_node *ImGuiDockNode, payload *ImGuiWindow, split_dir ImGuiDir, split_ratio float32, split_outer bool) {
	_split_outer := c.CBool(split_outer)
	giLib.Call(_func_igDockContextQueueDock_, []interface{}{&ctx, &target, &target_node, &payload, &split_dir, &split_ratio, &_split_outer})
}

var _func_igDockContextQueueUndockWindow_ = &c.FuncPrototype{Name: "igDockContextQueueUndockWindow", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDockContextQueueUndockWindow(ImGuiContext* ctx,ImGuiWindow* window);
func DockContextQueueUndockWindow(ctx *ImGuiContext, window *ImGuiWindow) {
	giLib.Call(_func_igDockContextQueueUndockWindow_, []interface{}{&ctx, &window})
}

var _func_igDockContextQueueUndockNode_ = &c.FuncPrototype{Name: "igDockContextQueueUndockNode", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDockContextQueueUndockNode(ImGuiContext* ctx,ImGuiDockNode* node);
func DockContextQueueUndockNode(ctx *ImGuiContext, node *ImGuiDockNode) {
	giLib.Call(_func_igDockContextQueueUndockNode_, []interface{}{&ctx, &node})
}

var _func_igDockContextProcessUndockWindow_ = &c.FuncPrototype{Name: "igDockContextProcessUndockWindow", OutType: c.Void, InTypes: _typs_PPI32}

// CIMGUI_API void igDockContextProcessUndockWindow(ImGuiContext* ctx,ImGuiWindow* window,bool clear_persistent_docking_ref);
func DockContextProcessUndockWindow(ctx *ImGuiContext, window *ImGuiWindow, clear_persistent_docking_ref bool) {
	_clear_persistent_docking_ref := c.CBool(clear_persistent_docking_ref)
	giLib.Call(_func_igDockContextProcessUndockWindow_, []interface{}{&ctx, &window, &_clear_persistent_docking_ref})
}

var _func_igDockContextProcessUndockNode_ = &c.FuncPrototype{Name: "igDockContextProcessUndockNode", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDockContextProcessUndockNode(ImGuiContext* ctx,ImGuiDockNode* node);
func DockContextProcessUndockNode(ctx *ImGuiContext, node *ImGuiDockNode) {
	giLib.Call(_func_igDockContextProcessUndockNode_, []interface{}{&ctx, &node})
}

var _func_igDockContextCalcDropPosForDocking_ = &c.FuncPrototype{Name: "igDockContextCalcDropPosForDocking", OutType: c.I32, InTypes: _typs_PPPPI32I32P}

// CIMGUI_API bool igDockContextCalcDropPosForDocking(ImGuiWindow* target,ImGuiDockNode* target_node,ImGuiWindow* payload_window,ImGuiDockNode* payload_node,ImGuiDir split_dir,bool split_outer,ImVec2* out_pos);
func DockContextCalcDropPosForDocking(target *ImGuiWindow, target_node *ImGuiDockNode, payload_window *ImGuiWindow, payload_node *ImGuiDockNode, split_dir ImGuiDir, split_outer bool, out_pos *ImVec2) bool {
	_split_outer := c.CBool(split_outer)

	return giLib.Call(_func_igDockContextCalcDropPosForDocking_, []interface{}{&target, &target_node, &payload_window, &payload_node, &split_dir, &_split_outer, &out_pos}).BoolFree()
}

var _func_igDockContextFindNodeByID_ = &c.FuncPrototype{Name: "igDockContextFindNodeByID", OutType: c.Pointer, InTypes: _typs_PU32}

// CIMGUI_API ImGuiDockNode* igDockContextFindNodeByID(ImGuiContext* ctx,ImGuiID id);
func DockContextFindNodeByID(ctx *ImGuiContext, id ImGuiID) *ImGuiDockNode {
	return (*ImGuiDockNode)(giLib.Call(_func_igDockContextFindNodeByID_, []interface{}{&ctx, &id}).PtrFree())
}

var _func_igDockNodeWindowMenuHandler_Default_ = &c.FuncPrototype{Name: "igDockNodeWindowMenuHandler_Default", OutType: c.Void, InTypes: _typs_PPP}

// CIMGUI_API void igDockNodeWindowMenuHandler_Default(ImGuiContext* ctx,ImGuiDockNode* node,ImGuiTabBar* tab_bar);
func DockNodeWindowMenuHandler_Default(ctx *ImGuiContext, node *ImGuiDockNode, tab_bar *ImGuiTabBar) {
	giLib.Call(_func_igDockNodeWindowMenuHandler_Default_, []interface{}{&ctx, &node, &tab_bar})
}

var _func_igDockNodeBeginAmendTabBar_ = &c.FuncPrototype{Name: "igDockNodeBeginAmendTabBar", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool igDockNodeBeginAmendTabBar(ImGuiDockNode* node);
func DockNodeBeginAmendTabBar(node *ImGuiDockNode) bool {
	return giLib.Call(_func_igDockNodeBeginAmendTabBar_, []interface{}{&node}).BoolFree()
}

var _func_igDockNodeEndAmendTabBar_ = &c.FuncPrototype{Name: "igDockNodeEndAmendTabBar", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igDockNodeEndAmendTabBar(void);
func DockNodeEndAmendTabBar() {
	giLib.Call(_func_igDockNodeEndAmendTabBar_, nil)
}

var _func_igDockNodeGetRootNode_ = &c.FuncPrototype{Name: "igDockNodeGetRootNode", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiDockNode* igDockNodeGetRootNode(ImGuiDockNode* node);
func DockNodeGetRootNode(node *ImGuiDockNode) *ImGuiDockNode {
	return (*ImGuiDockNode)(giLib.Call(_func_igDockNodeGetRootNode_, []interface{}{&node}).PtrFree())
}

var _func_igDockNodeIsInHierarchyOf_ = &c.FuncPrototype{Name: "igDockNodeIsInHierarchyOf", OutType: c.I32, InTypes: _typs_PP}

// CIMGUI_API bool igDockNodeIsInHierarchyOf(ImGuiDockNode* node,ImGuiDockNode* parent);
func DockNodeIsInHierarchyOf(node *ImGuiDockNode, parent *ImGuiDockNode) bool {
	return giLib.Call(_func_igDockNodeIsInHierarchyOf_, []interface{}{&node, &parent}).BoolFree()
}

var _func_igDockNodeGetDepth_ = &c.FuncPrototype{Name: "igDockNodeGetDepth", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API int igDockNodeGetDepth(const ImGuiDockNode* node);
func DockNodeGetDepth(node *ImGuiDockNode) int32 {
	return giLib.Call(_func_igDockNodeGetDepth_, []interface{}{&node}).I32Free()
}

var _func_igDockNodeGetWindowMenuButtonId_ = &c.FuncPrototype{Name: "igDockNodeGetWindowMenuButtonId", OutType: c.U32, InTypes: _typs_P}

// CIMGUI_API ImGuiID igDockNodeGetWindowMenuButtonId(const ImGuiDockNode* node);
func DockNodeGetWindowMenuButtonId(node *ImGuiDockNode) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igDockNodeGetWindowMenuButtonId_, []interface{}{&node}).U32Free())
}

var _func_igGetWindowDockNode_ = &c.FuncPrototype{Name: "igGetWindowDockNode", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiDockNode* igGetWindowDockNode(void);
func GetWindowDockNode() *ImGuiDockNode {
	return (*ImGuiDockNode)(giLib.Call(_func_igGetWindowDockNode_, nil).PtrFree())
}

var _func_igGetWindowAlwaysWantOwnTabBar_ = &c.FuncPrototype{Name: "igGetWindowAlwaysWantOwnTabBar", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool igGetWindowAlwaysWantOwnTabBar(ImGuiWindow* window);
func GetWindowAlwaysWantOwnTabBar(window *ImGuiWindow) bool {
	return giLib.Call(_func_igGetWindowAlwaysWantOwnTabBar_, []interface{}{&window}).BoolFree()
}

var _func_igBeginDocked_ = &c.FuncPrototype{Name: "igBeginDocked", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igBeginDocked(ImGuiWindow* window,bool* p_open);
func BeginDocked(window *ImGuiWindow, p_open *bool) {
	_p_open := cboolPtr(p_open)
	giLib.Call(_func_igBeginDocked_, []interface{}{&window, &p_open})
	cboolGet(p_open, _p_open)
}

var _func_igBeginDockableDragDropSource_ = &c.FuncPrototype{Name: "igBeginDockableDragDropSource", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igBeginDockableDragDropSource(ImGuiWindow* window);
func BeginDockableDragDropSource(window *ImGuiWindow) {
	giLib.Call(_func_igBeginDockableDragDropSource_, []interface{}{&window})
}

var _func_igBeginDockableDragDropTarget_ = &c.FuncPrototype{Name: "igBeginDockableDragDropTarget", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igBeginDockableDragDropTarget(ImGuiWindow* window);
func BeginDockableDragDropTarget(window *ImGuiWindow) {
	giLib.Call(_func_igBeginDockableDragDropTarget_, []interface{}{&window})
}

var _func_igSetWindowDock_ = &c.FuncPrototype{Name: "igSetWindowDock", OutType: c.Void, InTypes: _typs_PU32U32}

// CIMGUI_API void igSetWindowDock(ImGuiWindow* window,ImGuiID dock_id,ImGuiCond cond);
func SetWindowDock(window *ImGuiWindow, dock_id ImGuiID, cond ImGuiCond) {
	giLib.Call(_func_igSetWindowDock_, []interface{}{&window, &dock_id, &cond})
}

var _func_igDockBuilderDockWindow_ = &c.FuncPrototype{Name: "igDockBuilderDockWindow", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void igDockBuilderDockWindow(const char* window_name,ImGuiID node_id);
func DockBuilderDockWindow(window_name string, node_id ImGuiID) {
	_window_name := c.CStr(window_name)
	defer usf.Free(_window_name)
	giLib.Call(_func_igDockBuilderDockWindow_, []interface{}{&_window_name, &node_id})
}

var _func_igDockBuilderGetNode_ = &c.FuncPrototype{Name: "igDockBuilderGetNode", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiDockNode* igDockBuilderGetNode(ImGuiID node_id);
func DockBuilderGetNode(node_id ImGuiID) *ImGuiDockNode {
	return (*ImGuiDockNode)(giLib.Call(_func_igDockBuilderGetNode_, []interface{}{&node_id}).PtrFree())
}

var _func_igDockBuilderGetCentralNode_ = &c.FuncPrototype{Name: "igDockBuilderGetCentralNode", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiDockNode* igDockBuilderGetCentralNode(ImGuiID node_id);
func DockBuilderGetCentralNode(node_id ImGuiID) *ImGuiDockNode {
	return (*ImGuiDockNode)(giLib.Call(_func_igDockBuilderGetCentralNode_, []interface{}{&node_id}).PtrFree())
}

var _func_igDockBuilderAddNode_ = &c.FuncPrototype{Name: "igDockBuilderAddNode", OutType: c.U32, InTypes: _typs_U32U32}

// CIMGUI_API ImGuiID igDockBuilderAddNode(ImGuiID node_id,ImGuiDockNodeFlags flags);
func DockBuilderAddNode(node_id ImGuiID, flags ImGuiDockNodeFlags) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igDockBuilderAddNode_, []interface{}{&node_id, &flags}).U32Free())
}

var _func_igDockBuilderRemoveNode_ = &c.FuncPrototype{Name: "igDockBuilderRemoveNode", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igDockBuilderRemoveNode(ImGuiID node_id);
func DockBuilderRemoveNode(node_id ImGuiID) {
	giLib.Call(_func_igDockBuilderRemoveNode_, []interface{}{&node_id})
}

var _func_igDockBuilderRemoveNodeDockedWindows_ = &c.FuncPrototype{Name: "igDockBuilderRemoveNodeDockedWindows", OutType: c.Void, InTypes: _typs_U32I32}

// CIMGUI_API void igDockBuilderRemoveNodeDockedWindows(ImGuiID node_id,bool clear_settings_refs);
func DockBuilderRemoveNodeDockedWindows(node_id ImGuiID, clear_settings_refs bool) {
	_clear_settings_refs := c.CBool(clear_settings_refs)
	giLib.Call(_func_igDockBuilderRemoveNodeDockedWindows_, []interface{}{&node_id, &_clear_settings_refs})
}

var _func_igDockBuilderRemoveNodeChildNodes_ = &c.FuncPrototype{Name: "igDockBuilderRemoveNodeChildNodes", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igDockBuilderRemoveNodeChildNodes(ImGuiID node_id);
func DockBuilderRemoveNodeChildNodes(node_id ImGuiID) {
	giLib.Call(_func_igDockBuilderRemoveNodeChildNodes_, []interface{}{&node_id})
}

var _func_igDockBuilderSetNodePos_ = &c.FuncPrototype{Name: "igDockBuilderSetNodePos", OutType: c.Void, InTypes: _typs_U32Vec2}

// CIMGUI_API void igDockBuilderSetNodePos(ImGuiID node_id,ImVec2 pos);
func DockBuilderSetNodePos(node_id ImGuiID, pos ImVec2) {
	giLib.Call(_func_igDockBuilderSetNodePos_, []interface{}{&node_id, &pos})
}

var _func_igDockBuilderSetNodeSize_ = &c.FuncPrototype{Name: "igDockBuilderSetNodeSize", OutType: c.Void, InTypes: _typs_U32Vec2}

// CIMGUI_API void igDockBuilderSetNodeSize(ImGuiID node_id,ImVec2 size);
func DockBuilderSetNodeSize(node_id ImGuiID, size ImVec2) {
	giLib.Call(_func_igDockBuilderSetNodeSize_, []interface{}{&node_id, &size})
}

var _func_igDockBuilderSplitNode_ = &c.FuncPrototype{Name: "igDockBuilderSplitNode", OutType: c.U32, InTypes: _typs_U32I32F32PP}

// CIMGUI_API ImGuiID igDockBuilderSplitNode(ImGuiID node_id,ImGuiDir split_dir,float size_ratio_for_node_at_dir,ImGuiID* out_id_at_dir,ImGuiID* out_id_at_opposite_dir);
func DockBuilderSplitNode(node_id ImGuiID, split_dir ImGuiDir, size_ratio_for_node_at_dir float32, out_id_at_dir *ImGuiID, out_id_at_opposite_dir *ImGuiID) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igDockBuilderSplitNode_, []interface{}{&node_id, &split_dir, &size_ratio_for_node_at_dir, &out_id_at_dir, &out_id_at_opposite_dir}).U32Free())
}

var _func_igDockBuilderCopyDockSpace_ = &c.FuncPrototype{Name: "igDockBuilderCopyDockSpace", OutType: c.Void, InTypes: _typs_U32U32P}

// CIMGUI_API void igDockBuilderCopyDockSpace(ImGuiID src_dockspace_id,ImGuiID dst_dockspace_id,ImVector_const_charPtr* in_window_remap_pairs);
func DockBuilderCopyDockSpace(src_dockspace_id ImGuiID, dst_dockspace_id ImGuiID, in_window_remap_pairs *ImVector_const_charPtr) {
	giLib.Call(_func_igDockBuilderCopyDockSpace_, []interface{}{&src_dockspace_id, &dst_dockspace_id, &in_window_remap_pairs})
}

var _func_igDockBuilderCopyNode_ = &c.FuncPrototype{Name: "igDockBuilderCopyNode", OutType: c.Void, InTypes: _typs_U32U32P}

// CIMGUI_API void igDockBuilderCopyNode(ImGuiID src_node_id,ImGuiID dst_node_id,ImVector_ImGuiID* out_node_remap_pairs);
func DockBuilderCopyNode(src_node_id ImGuiID, dst_node_id ImGuiID, out_node_remap_pairs *ImVector_ImGuiID) {
	giLib.Call(_func_igDockBuilderCopyNode_, []interface{}{&src_node_id, &dst_node_id, &out_node_remap_pairs})
}

var _func_igDockBuilderCopyWindowSettings_ = &c.FuncPrototype{Name: "igDockBuilderCopyWindowSettings", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDockBuilderCopyWindowSettings(const char* src_name,const char* dst_name);
func DockBuilderCopyWindowSettings(src_name string, dst_name string) {
	_src_name, _dst_name := c.CStr(src_name), c.CStr(dst_name)
	defer usf.Free(_src_name)
	defer usf.Free(_dst_name)
	giLib.Call(_func_igDockBuilderCopyWindowSettings_, []interface{}{&_src_name, &_dst_name})
}

var _func_igDockBuilderFinish_ = &c.FuncPrototype{Name: "igDockBuilderFinish", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igDockBuilderFinish(ImGuiID node_id);
func DockBuilderFinish(node_id ImGuiID) {
	giLib.Call(_func_igDockBuilderFinish_, []interface{}{&node_id})
}

var _func_igPushFocusScope_ = &c.FuncPrototype{Name: "igPushFocusScope", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igPushFocusScope(ImGuiID id);
func PushFocusScope(id ImGuiID) {
	giLib.Call(_func_igPushFocusScope_, []interface{}{&id})
}

var _func_igPopFocusScope_ = &c.FuncPrototype{Name: "igPopFocusScope", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igPopFocusScope(void);
func PopFocusScope() {
	giLib.Call(_func_igPopFocusScope_, nil)
}

var _func_igGetCurrentFocusScope_ = &c.FuncPrototype{Name: "igGetCurrentFocusScope", OutType: c.U32, InTypes: nil}

// CIMGUI_API ImGuiID igGetCurrentFocusScope(void);
func GetCurrentFocusScope() ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetCurrentFocusScope_, nil).U32Free())
}

var _func_igIsDragDropActive_ = &c.FuncPrototype{Name: "igIsDragDropActive", OutType: c.I32, InTypes: nil}

// CIMGUI_API bool igIsDragDropActive(void);
func IsDragDropActive() bool {
	return giLib.Call(_func_igIsDragDropActive_, nil).BoolFree()
}

var _func_igBeginDragDropTargetCustom_ = &c.FuncPrototype{Name: "igBeginDragDropTargetCustom", OutType: c.I32, InTypes: _typs_RectU32}

// CIMGUI_API bool igBeginDragDropTargetCustom(const ImRect bb,ImGuiID id);
func BeginDragDropTargetCustom(bb ImRect, id ImGuiID) bool {
	return giLib.Call(_func_igBeginDragDropTargetCustom_, []interface{}{&bb, &id}).BoolFree()
}

var _func_igClearDragDrop_ = &c.FuncPrototype{Name: "igClearDragDrop", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igClearDragDrop(void);
func ClearDragDrop() {
	giLib.Call(_func_igClearDragDrop_, nil)
}

var _func_igIsDragDropPayloadBeingAccepted_ = &c.FuncPrototype{Name: "igIsDragDropPayloadBeingAccepted", OutType: c.I32, InTypes: nil}

// CIMGUI_API bool igIsDragDropPayloadBeingAccepted(void);
func IsDragDropPayloadBeingAccepted() bool {
	return giLib.Call(_func_igIsDragDropPayloadBeingAccepted_, nil).BoolFree()
}

var _func_igRenderDragDropTargetRect_ = &c.FuncPrototype{Name: "igRenderDragDropTargetRect", OutType: c.Void, InTypes: _typs_Rect}

// CIMGUI_API void igRenderDragDropTargetRect(const ImRect bb);
func RenderDragDropTargetRect(bb ImRect) {
	giLib.Call(_func_igRenderDragDropTargetRect_, []interface{}{&bb})
}

var _func_igSetWindowClipRectBeforeSetChannel_ = &c.FuncPrototype{Name: "igSetWindowClipRectBeforeSetChannel", OutType: c.Void, InTypes: _typs_PRect}

// CIMGUI_API void igSetWindowClipRectBeforeSetChannel(ImGuiWindow* window,const ImRect clip_rect);
func SetWindowClipRectBeforeSetChannel(window *ImGuiWindow, clip_rect ImRect) {
	giLib.Call(_func_igSetWindowClipRectBeforeSetChannel_, []interface{}{&window, &clip_rect})
}

var _func_igBeginColumns_ = &c.FuncPrototype{Name: "igBeginColumns", OutType: c.Void, InTypes: _typs_PI32U32}

// CIMGUI_API void igBeginColumns(const char* str_id,int count,ImGuiOldColumnFlags flags);
func BeginColumns(str_id string, count int32, flags ImGuiOldColumnFlags) {
	_str_id := c.CStr(str_id)
	defer usf.Free(_str_id)
	giLib.Call(_func_igBeginColumns_, []interface{}{&_str_id, &count, &flags})
}

var _func_igEndColumns_ = &c.FuncPrototype{Name: "igEndColumns", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igEndColumns(void);
func EndColumns() {
	giLib.Call(_func_igEndColumns_, nil)
}

var _func_igPushColumnClipRect_ = &c.FuncPrototype{Name: "igPushColumnClipRect", OutType: c.Void, InTypes: _typs_I32}

// CIMGUI_API void igPushColumnClipRect(int column_index);
func PushColumnClipRect(column_index int32) {
	giLib.Call(_func_igPushColumnClipRect_, []interface{}{&column_index})
}

var _func_igPushColumnsBackground_ = &c.FuncPrototype{Name: "igPushColumnsBackground", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igPushColumnsBackground(void);
func PushColumnsBackground() {
	giLib.Call(_func_igPushColumnsBackground_, nil)
}

var _func_igPopColumnsBackground_ = &c.FuncPrototype{Name: "igPopColumnsBackground", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igPopColumnsBackground(void);
func PopColumnsBackground() {
	giLib.Call(_func_igPopColumnsBackground_, nil)
}

var _func_igGetColumnsID_ = &c.FuncPrototype{Name: "igGetColumnsID", OutType: c.U32, InTypes: _typs_PI32}

// CIMGUI_API ImGuiID igGetColumnsID(const char* str_id,int count);
func GetColumnsID(str_id string, count int32) ImGuiID {
	_str_id := c.CStr(str_id)
	defer usf.Free(_str_id)
	return (ImGuiID)(giLib.Call(_func_igGetColumnsID_, []interface{}{&_str_id, &count}).U32Free())
}

var _func_igFindOrCreateColumns_ = &c.FuncPrototype{Name: "igFindOrCreateColumns", OutType: c.Pointer, InTypes: _typs_PU32}

// CIMGUI_API ImGuiOldColumns* igFindOrCreateColumns(ImGuiWindow* window,ImGuiID id);
func FindOrCreateColumns(window *ImGuiWindow, id ImGuiID) *ImGuiOldColumns {
	return (*ImGuiOldColumns)(giLib.Call(_func_igFindOrCreateColumns_, []interface{}{&window, &id}).PtrFree())
}

var _func_igGetColumnOffsetFromNorm_ = &c.FuncPrototype{Name: "igGetColumnOffsetFromNorm", OutType: c.F32, InTypes: _typs_PF32}

// CIMGUI_API float igGetColumnOffsetFromNorm(const ImGuiOldColumns* columns,float offset_norm);
func GetColumnOffsetFromNorm(columns *ImGuiOldColumns, offset_norm float32) float32 {
	return giLib.Call(_func_igGetColumnOffsetFromNorm_, []interface{}{&columns, &offset_norm}).F32Free()
}

var _func_igGetColumnNormFromOffset_ = &c.FuncPrototype{Name: "igGetColumnNormFromOffset", OutType: c.F32, InTypes: _typs_PF32}

// CIMGUI_API float igGetColumnNormFromOffset(const ImGuiOldColumns* columns,float offset);
func GetColumnNormFromOffset(columns *ImGuiOldColumns, offset float32) float32 {
	return giLib.Call(_func_igGetColumnNormFromOffset_, []interface{}{&columns, &offset}).F32Free()
}

var _func_igTableOpenContextMenu_ = &c.FuncPrototype{Name: "igTableOpenContextMenu", OutType: c.Void, InTypes: _typs_I32}

// CIMGUI_API void igTableOpenContextMenu(int column_n);
func TableOpenContextMenu(column_n int32) {
	giLib.Call(_func_igTableOpenContextMenu_, []interface{}{&column_n})
}

var _func_igTableSetColumnWidth_ = &c.FuncPrototype{Name: "igTableSetColumnWidth", OutType: c.Void, InTypes: _typs_I32F32}

// CIMGUI_API void igTableSetColumnWidth(int column_n,float width);
func TableSetColumnWidth(column_n int32, width float32) {
	giLib.Call(_func_igTableSetColumnWidth_, []interface{}{&column_n, &width})
}

var _func_igTableSetColumnSortDirection_ = &c.FuncPrototype{Name: "igTableSetColumnSortDirection", OutType: c.Void, InTypes: _typs_I32U32I32}

// CIMGUI_API void igTableSetColumnSortDirection(int column_n,ImGuiSortDirection sort_direction,bool append_to_sort_specs);
func TableSetColumnSortDirection(column_n int32, sort_direction ImGuiSortDirection, append_to_sort_specs bool) {
	_append_to_sort_specs := c.CBool(append_to_sort_specs)
	giLib.Call(_func_igTableSetColumnSortDirection_, []interface{}{&column_n, &sort_direction, &_append_to_sort_specs})
}

var _func_igTableGetHoveredColumn_ = &c.FuncPrototype{Name: "igTableGetHoveredColumn", OutType: c.I32, InTypes: nil}

// CIMGUI_API int igTableGetHoveredColumn(void);
func TableGetHoveredColumn() int32 {
	return giLib.Call(_func_igTableGetHoveredColumn_, nil).I32Free()
}

var _func_igTableGetHeaderRowHeight_ = &c.FuncPrototype{Name: "igTableGetHeaderRowHeight", OutType: c.F32, InTypes: nil}

// CIMGUI_API float igTableGetHeaderRowHeight(void);
func TableGetHeaderRowHeight() float32 {
	return giLib.Call(_func_igTableGetHeaderRowHeight_, nil).F32Free()
}

var _func_igTablePushBackgroundChannel_ = &c.FuncPrototype{Name: "igTablePushBackgroundChannel", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igTablePushBackgroundChannel(void);
func TablePushBackgroundChannel() {
	giLib.Call(_func_igTablePushBackgroundChannel_, nil)
}

var _func_igTablePopBackgroundChannel_ = &c.FuncPrototype{Name: "igTablePopBackgroundChannel", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igTablePopBackgroundChannel(void);
func TablePopBackgroundChannel() {
	giLib.Call(_func_igTablePopBackgroundChannel_, nil)
}

var _func_igGetCurrentTable_ = &c.FuncPrototype{Name: "igGetCurrentTable", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTable* igGetCurrentTable(void);
func GetCurrentTable() *ImGuiTable {
	return (*ImGuiTable)(giLib.Call(_func_igGetCurrentTable_, nil).PtrFree())
}

var _func_igTableFindByID_ = &c.FuncPrototype{Name: "igTableFindByID", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiTable* igTableFindByID(ImGuiID id);
func TableFindByID(id ImGuiID) *ImGuiTable {
	return (*ImGuiTable)(giLib.Call(_func_igTableFindByID_, []interface{}{&id}).PtrFree())
}

var _func_igBeginTableEx_ = &c.FuncPrototype{Name: "igBeginTableEx", OutType: c.I32, InTypes: _typs_PU32I32U32Vec2F32}

// CIMGUI_API bool igBeginTableEx(const char* name,ImGuiID id,int columns_count,ImGuiTableFlags flags,const ImVec2 outer_size,float inner_width);
func BeginTableEx(name string, id ImGuiID, columns_count int32, flags ImGuiTableFlags, outer_size ImVec2, inner_width float32) bool {
	_name := c.CStr(name)
	defer usf.Free(_name)

	return giLib.Call(_func_igBeginTableEx_, []interface{}{&_name, &id, &columns_count, &flags, &outer_size, &inner_width}).BoolFree()
}

var _func_igTableBeginInitMemory_ = &c.FuncPrototype{Name: "igTableBeginInitMemory", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void igTableBeginInitMemory(ImGuiTable* table,int columns_count);
func TableBeginInitMemory(table *ImGuiTable, columns_count int32) {
	giLib.Call(_func_igTableBeginInitMemory_, []interface{}{&table, &columns_count})
}

var _func_igTableBeginApplyRequests_ = &c.FuncPrototype{Name: "igTableBeginApplyRequests", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableBeginApplyRequests(ImGuiTable* table);
func TableBeginApplyRequests(table *ImGuiTable) {
	giLib.Call(_func_igTableBeginApplyRequests_, []interface{}{&table})
}

var _func_igTableSetupDrawChannels_ = &c.FuncPrototype{Name: "igTableSetupDrawChannels", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableSetupDrawChannels(ImGuiTable* table);
func TableSetupDrawChannels(table *ImGuiTable) {
	giLib.Call(_func_igTableSetupDrawChannels_, []interface{}{&table})
}

var _func_igTableUpdateLayout_ = &c.FuncPrototype{Name: "igTableUpdateLayout", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableUpdateLayout(ImGuiTable* table);
func TableUpdateLayout(table *ImGuiTable) {
	giLib.Call(_func_igTableUpdateLayout_, []interface{}{&table})
}

var _func_igTableUpdateBorders_ = &c.FuncPrototype{Name: "igTableUpdateBorders", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableUpdateBorders(ImGuiTable* table);
func TableUpdateBorders(table *ImGuiTable) {
	giLib.Call(_func_igTableUpdateBorders_, []interface{}{&table})
}

var _func_igTableUpdateColumnsWeightFromWidth_ = &c.FuncPrototype{Name: "igTableUpdateColumnsWeightFromWidth", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableUpdateColumnsWeightFromWidth(ImGuiTable* table);
func TableUpdateColumnsWeightFromWidth(table *ImGuiTable) {
	giLib.Call(_func_igTableUpdateColumnsWeightFromWidth_, []interface{}{&table})
}

var _func_igTableDrawBorders_ = &c.FuncPrototype{Name: "igTableDrawBorders", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableDrawBorders(ImGuiTable* table);
func TableDrawBorders(table *ImGuiTable) {
	giLib.Call(_func_igTableDrawBorders_, []interface{}{&table})
}

var _func_igTableDrawContextMenu_ = &c.FuncPrototype{Name: "igTableDrawContextMenu", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableDrawContextMenu(ImGuiTable* table);
func TableDrawContextMenu(table *ImGuiTable) {
	giLib.Call(_func_igTableDrawContextMenu_, []interface{}{&table})
}

var _func_igTableBeginContextMenuPopup_ = &c.FuncPrototype{Name: "igTableBeginContextMenuPopup", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool igTableBeginContextMenuPopup(ImGuiTable* table);
func TableBeginContextMenuPopup(table *ImGuiTable) bool {
	return giLib.Call(_func_igTableBeginContextMenuPopup_, []interface{}{&table}).BoolFree()
}

var _func_igTableMergeDrawChannels_ = &c.FuncPrototype{Name: "igTableMergeDrawChannels", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableMergeDrawChannels(ImGuiTable* table);
func TableMergeDrawChannels(table *ImGuiTable) {
	giLib.Call(_func_igTableMergeDrawChannels_, []interface{}{&table})
}

var _func_igTableGetInstanceData_ = &c.FuncPrototype{Name: "igTableGetInstanceData", OutType: c.Pointer, InTypes: _typs_PI32}

// CIMGUI_API ImGuiTableInstanceData* igTableGetInstanceData(ImGuiTable* table,int instance_no);
func TableGetInstanceData(table *ImGuiTable, instance_no int32) *ImGuiTableInstanceData {
	return (*ImGuiTableInstanceData)(giLib.Call(_func_igTableGetInstanceData_, []interface{}{&table, &instance_no}).PtrFree())
}

var _func_igTableGetInstanceID_ = &c.FuncPrototype{Name: "igTableGetInstanceID", OutType: c.U32, InTypes: _typs_PI32}

// CIMGUI_API ImGuiID igTableGetInstanceID(ImGuiTable* table,int instance_no);
func TableGetInstanceID(table *ImGuiTable, instance_no int32) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igTableGetInstanceID_, []interface{}{&table, &instance_no}).U32Free())
}

var _func_igTableSortSpecsSanitize_ = &c.FuncPrototype{Name: "igTableSortSpecsSanitize", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableSortSpecsSanitize(ImGuiTable* table);
func TableSortSpecsSanitize(table *ImGuiTable) {
	giLib.Call(_func_igTableSortSpecsSanitize_, []interface{}{&table})
}

var _func_igTableSortSpecsBuild_ = &c.FuncPrototype{Name: "igTableSortSpecsBuild", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableSortSpecsBuild(ImGuiTable* table);
func TableSortSpecsBuild(table *ImGuiTable) {
	giLib.Call(_func_igTableSortSpecsBuild_, []interface{}{&table})
}

var _func_igTableGetColumnNextSortDirection_ = &c.FuncPrototype{Name: "igTableGetColumnNextSortDirection", OutType: c.U32, InTypes: _typs_P}

// CIMGUI_API ImGuiSortDirection igTableGetColumnNextSortDirection(ImGuiTableColumn* column);
func TableGetColumnNextSortDirection(column *ImGuiTableColumn) ImGuiSortDirection {
	return (ImGuiSortDirection)(giLib.Call(_func_igTableGetColumnNextSortDirection_, []interface{}{&column}).U32Free())
}

var _func_igTableFixColumnSortDirection_ = &c.FuncPrototype{Name: "igTableFixColumnSortDirection", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igTableFixColumnSortDirection(ImGuiTable* table,ImGuiTableColumn* column);
func TableFixColumnSortDirection(table *ImGuiTable, column *ImGuiTableColumn) {
	giLib.Call(_func_igTableFixColumnSortDirection_, []interface{}{&table, &column})
}

var _func_igTableGetColumnWidthAuto_ = &c.FuncPrototype{Name: "igTableGetColumnWidthAuto", OutType: c.F32, InTypes: _typs_PP}

// CIMGUI_API float igTableGetColumnWidthAuto(ImGuiTable* table,ImGuiTableColumn* column);
func TableGetColumnWidthAuto(table *ImGuiTable, column *ImGuiTableColumn) float32 {
	return giLib.Call(_func_igTableGetColumnWidthAuto_, []interface{}{&table, &column}).F32Free()
}

var _func_igTableBeginRow_ = &c.FuncPrototype{Name: "igTableBeginRow", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableBeginRow(ImGuiTable* table);
func TableBeginRow(table *ImGuiTable) {
	giLib.Call(_func_igTableBeginRow_, []interface{}{&table})
}

var _func_igTableEndRow_ = &c.FuncPrototype{Name: "igTableEndRow", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableEndRow(ImGuiTable* table);
func TableEndRow(table *ImGuiTable) {
	giLib.Call(_func_igTableEndRow_, []interface{}{&table})
}

var _func_igTableBeginCell_ = &c.FuncPrototype{Name: "igTableBeginCell", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void igTableBeginCell(ImGuiTable* table,int column_n);
func TableBeginCell(table *ImGuiTable, column_n int32) {
	giLib.Call(_func_igTableBeginCell_, []interface{}{&table, &column_n})
}

var _func_igTableEndCell_ = &c.FuncPrototype{Name: "igTableEndCell", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableEndCell(ImGuiTable* table);
func TableEndCell(table *ImGuiTable) {
	giLib.Call(_func_igTableEndCell_, []interface{}{&table})
}

var _func_igTableGetCellBgRect_ = &c.FuncPrototype{Name: "igTableGetCellBgRect", OutType: c.Void, InTypes: _typs_PPI32}

// CIMGUI_API void igTableGetCellBgRect(ImRect* pOut,const ImGuiTable* table,int column_n);
func TableGetCellBgRect(pOut *ImRect, table *ImGuiTable, column_n int32) {
	giLib.Call(_func_igTableGetCellBgRect_, []interface{}{&pOut, &table, &column_n})
}

var _func_igTableGetColumnName_TablePtr_ = &c.FuncPrototype{Name: "igTableGetColumnName_TablePtr", OutType: c.Pointer, InTypes: _typs_PI32}

// CIMGUI_API const char* igTableGetColumnName_TablePtr(const ImGuiTable* table,int column_n);
func TableGetColumnName_TablePtr(table *ImGuiTable, column_n int32) string {
	return giLib.Call(_func_igTableGetColumnName_TablePtr_, []interface{}{&table, &column_n}).StrFree()
}

var _func_igTableGetColumnResizeID_ = &c.FuncPrototype{Name: "igTableGetColumnResizeID", OutType: c.U32, InTypes: _typs_PI32I32}

// CIMGUI_API ImGuiID igTableGetColumnResizeID(ImGuiTable* table,int column_n,int instance_no);
func TableGetColumnResizeID(table *ImGuiTable, column_n int32, instance_no int32) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igTableGetColumnResizeID_, []interface{}{&table, &column_n, &instance_no}).U32Free())
}

var _func_igTableGetMaxColumnWidth_ = &c.FuncPrototype{Name: "igTableGetMaxColumnWidth", OutType: c.F32, InTypes: _typs_PI32}

// CIMGUI_API float igTableGetMaxColumnWidth(const ImGuiTable* table,int column_n);
func TableGetMaxColumnWidth(table *ImGuiTable, column_n int32) float32 {
	return giLib.Call(_func_igTableGetMaxColumnWidth_, []interface{}{&table, &column_n}).F32Free()
}

var _func_igTableSetColumnWidthAutoSingle_ = &c.FuncPrototype{Name: "igTableSetColumnWidthAutoSingle", OutType: c.Void, InTypes: _typs_PI32}

// CIMGUI_API void igTableSetColumnWidthAutoSingle(ImGuiTable* table,int column_n);
func TableSetColumnWidthAutoSingle(table *ImGuiTable, column_n int32) {
	giLib.Call(_func_igTableSetColumnWidthAutoSingle_, []interface{}{&table, &column_n})
}

var _func_igTableSetColumnWidthAutoAll_ = &c.FuncPrototype{Name: "igTableSetColumnWidthAutoAll", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableSetColumnWidthAutoAll(ImGuiTable* table);
func TableSetColumnWidthAutoAll(table *ImGuiTable) {
	giLib.Call(_func_igTableSetColumnWidthAutoAll_, []interface{}{&table})
}

var _func_igTableRemove_ = &c.FuncPrototype{Name: "igTableRemove", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableRemove(ImGuiTable* table);
func TableRemove(table *ImGuiTable) {
	giLib.Call(_func_igTableRemove_, []interface{}{&table})
}

var _func_igTableGcCompactTransientBuffers_TablePtr_ = &c.FuncPrototype{Name: "igTableGcCompactTransientBuffers_TablePtr", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableGcCompactTransientBuffers_TablePtr(ImGuiTable* table);
func TableGcCompactTransientBuffersByTablePtr(table *ImGuiTable) {
	giLib.Call(_func_igTableGcCompactTransientBuffers_TablePtr_, []interface{}{&table})
}

var _func_igTableGcCompactTransientBuffers_TableTempDataPtr_ = &c.FuncPrototype{Name: "igTableGcCompactTransientBuffers_TableTempDataPtr", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableGcCompactTransientBuffers_TableTempDataPtr(ImGuiTableTempData* table);
func TableGcCompactTransientBuffersByTableTempDataPtr(table *ImGuiTableTempData) {
	giLib.Call(_func_igTableGcCompactTransientBuffers_TableTempDataPtr_, []interface{}{&table})
}

var _func_igTableGcCompactSettings_ = &c.FuncPrototype{Name: "igTableGcCompactSettings", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igTableGcCompactSettings(void);
func TableGcCompactSettings() {
	giLib.Call(_func_igTableGcCompactSettings_, nil)
}

var _func_igTableLoadSettings_ = &c.FuncPrototype{Name: "igTableLoadSettings", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableLoadSettings(ImGuiTable* table);
func TableLoadSettings(table *ImGuiTable) {
	giLib.Call(_func_igTableLoadSettings_, []interface{}{&table})
}

var _func_igTableSaveSettings_ = &c.FuncPrototype{Name: "igTableSaveSettings", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableSaveSettings(ImGuiTable* table);
func TableSaveSettings(table *ImGuiTable) {
	giLib.Call(_func_igTableSaveSettings_, []interface{}{&table})
}

var _func_igTableResetSettings_ = &c.FuncPrototype{Name: "igTableResetSettings", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igTableResetSettings(ImGuiTable* table);
func TableResetSettings(table *ImGuiTable) {
	giLib.Call(_func_igTableResetSettings_, []interface{}{&table})
}

var _func_igTableGetBoundSettings_ = &c.FuncPrototype{Name: "igTableGetBoundSettings", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiTableSettings* igTableGetBoundSettings(ImGuiTable* table);
func TableGetBoundSettings(table *ImGuiTable) *ImGuiTableSettings {
	return (*ImGuiTableSettings)(giLib.Call(_func_igTableGetBoundSettings_, []interface{}{&table}).PtrFree())
}

var _func_igTableSettingsAddSettingsHandler_ = &c.FuncPrototype{Name: "igTableSettingsAddSettingsHandler", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igTableSettingsAddSettingsHandler(void);
func TableSettingsAddSettingsHandler() {
	giLib.Call(_func_igTableSettingsAddSettingsHandler_, nil)
}

var _func_igTableSettingsCreate_ = &c.FuncPrototype{Name: "igTableSettingsCreate", OutType: c.Pointer, InTypes: _typs_U32I32}

// CIMGUI_API ImGuiTableSettings* igTableSettingsCreate(ImGuiID id,int columns_count);
func TableSettingsCreate(id ImGuiID, columns_count int32) *ImGuiTableSettings {
	return (*ImGuiTableSettings)(giLib.Call(_func_igTableSettingsCreate_, []interface{}{&id, &columns_count}).PtrFree())
}

var _func_igTableSettingsFindByID_ = &c.FuncPrototype{Name: "igTableSettingsFindByID", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiTableSettings* igTableSettingsFindByID(ImGuiID id);
func TableSettingsFindByID(id ImGuiID) *ImGuiTableSettings {
	return (*ImGuiTableSettings)(giLib.Call(_func_igTableSettingsFindByID_, []interface{}{&id}).PtrFree())
}

var _func_igGetCurrentTabBar_ = &c.FuncPrototype{Name: "igGetCurrentTabBar", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API ImGuiTabBar* igGetCurrentTabBar(void);
func GetCurrentTabBar() *ImGuiTabBar {
	return (*ImGuiTabBar)(giLib.Call(_func_igGetCurrentTabBar_, nil).PtrFree())
}

var _func_igBeginTabBarEx_ = &c.FuncPrototype{Name: "igBeginTabBarEx", OutType: c.I32, InTypes: _typs_PRectU32P}

// CIMGUI_API bool igBeginTabBarEx(ImGuiTabBar* tab_bar,const ImRect bb,ImGuiTabBarFlags flags,ImGuiDockNode* dock_node);
func BeginTabBarEx(tab_bar *ImGuiTabBar, bb ImRect, flags ImGuiTabBarFlags, dock_node *ImGuiDockNode) bool {
	return giLib.Call(_func_igBeginTabBarEx_, []interface{}{&tab_bar, &bb, &flags, &dock_node}).BoolFree()
}

var _func_igTabBarFindTabByID_ = &c.FuncPrototype{Name: "igTabBarFindTabByID", OutType: c.Pointer, InTypes: _typs_PU32}

// CIMGUI_API ImGuiTabItem* igTabBarFindTabByID(ImGuiTabBar* tab_bar,ImGuiID tab_id);
func TabBarFindTabByID(tab_bar *ImGuiTabBar, tab_id ImGuiID) *ImGuiTabItem {
	return (*ImGuiTabItem)(giLib.Call(_func_igTabBarFindTabByID_, []interface{}{&tab_bar, &tab_id}).PtrFree())
}

var _func_igTabBarFindTabByOrder_ = &c.FuncPrototype{Name: "igTabBarFindTabByOrder", OutType: c.Pointer, InTypes: _typs_PI32}

// CIMGUI_API ImGuiTabItem* igTabBarFindTabByOrder(ImGuiTabBar* tab_bar,int order);
func TabBarFindTabByOrder(tab_bar *ImGuiTabBar, order int32) *ImGuiTabItem {
	return (*ImGuiTabItem)(giLib.Call(_func_igTabBarFindTabByOrder_, []interface{}{&tab_bar, &order}).PtrFree())
}

var _func_igTabBarFindMostRecentlySelectedTabForActiveWindow_ = &c.FuncPrototype{Name: "igTabBarFindMostRecentlySelectedTabForActiveWindow", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiTabItem* igTabBarFindMostRecentlySelectedTabForActiveWindow(ImGuiTabBar* tab_bar);
func TabBarFindMostRecentlySelectedTabForActiveWindow(tab_bar *ImGuiTabBar) *ImGuiTabItem {
	return (*ImGuiTabItem)(giLib.Call(_func_igTabBarFindMostRecentlySelectedTabForActiveWindow_, []interface{}{&tab_bar}).PtrFree())
}

var _func_igTabBarGetCurrentTab_ = &c.FuncPrototype{Name: "igTabBarGetCurrentTab", OutType: c.Pointer, InTypes: _typs_P}

// CIMGUI_API ImGuiTabItem* igTabBarGetCurrentTab(ImGuiTabBar* tab_bar);
func TabBarGetCurrentTab(tab_bar *ImGuiTabBar) *ImGuiTabItem {
	return (*ImGuiTabItem)(giLib.Call(_func_igTabBarGetCurrentTab_, []interface{}{&tab_bar}).PtrFree())
}

var _func_igTabBarGetTabOrder_ = &c.FuncPrototype{Name: "igTabBarGetTabOrder", OutType: c.I32, InTypes: _typs_PP}

// CIMGUI_API int igTabBarGetTabOrder(ImGuiTabBar* tab_bar,ImGuiTabItem* tab);
func TabBarGetTabOrder(tab_bar *ImGuiTabBar, tab *ImGuiTabItem) int32 {
	return giLib.Call(_func_igTabBarGetTabOrder_, []interface{}{&tab_bar, &tab}).I32Free()
}

var _func_igTabBarGetTabName_ = &c.FuncPrototype{Name: "igTabBarGetTabName", OutType: c.Pointer, InTypes: _typs_PP}

// CIMGUI_API const char* igTabBarGetTabName(ImGuiTabBar* tab_bar,ImGuiTabItem* tab);
func TabBarGetTabName(tab_bar *ImGuiTabBar, tab *ImGuiTabItem) string {
	return giLib.Call(_func_igTabBarGetTabName_, []interface{}{&tab_bar, &tab}).StrFree()
}

var _func_igTabBarAddTab_ = &c.FuncPrototype{Name: "igTabBarAddTab", OutType: c.Void, InTypes: _typs_PU32P}

// CIMGUI_API void igTabBarAddTab(ImGuiTabBar* tab_bar,ImGuiTabItemFlags tab_flags,ImGuiWindow* window);
func TabBarAddTab(tab_bar *ImGuiTabBar, tab_flags ImGuiTabItemFlags, window *ImGuiWindow) {
	giLib.Call(_func_igTabBarAddTab_, []interface{}{&tab_bar, &tab_flags, &window})
}

var _func_igTabBarRemoveTab_ = &c.FuncPrototype{Name: "igTabBarRemoveTab", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void igTabBarRemoveTab(ImGuiTabBar* tab_bar,ImGuiID tab_id);
func TabBarRemoveTab(tab_bar *ImGuiTabBar, tab_id ImGuiID) {
	giLib.Call(_func_igTabBarRemoveTab_, []interface{}{&tab_bar, &tab_id})
}

var _func_igTabBarCloseTab_ = &c.FuncPrototype{Name: "igTabBarCloseTab", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igTabBarCloseTab(ImGuiTabBar* tab_bar,ImGuiTabItem* tab);
func TabBarCloseTab(tab_bar *ImGuiTabBar, tab *ImGuiTabItem) {
	giLib.Call(_func_igTabBarCloseTab_, []interface{}{&tab_bar, &tab})
}

var _func_igTabBarQueueFocus_ = &c.FuncPrototype{Name: "igTabBarQueueFocus", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igTabBarQueueFocus(ImGuiTabBar* tab_bar,ImGuiTabItem* tab);
func TabBarQueueFocus(tab_bar *ImGuiTabBar, tab *ImGuiTabItem) {
	giLib.Call(_func_igTabBarQueueFocus_, []interface{}{&tab_bar, &tab})
}

var _func_igTabBarQueueReorder_ = &c.FuncPrototype{Name: "igTabBarQueueReorder", OutType: c.Void, InTypes: _typs_PPI32}

// CIMGUI_API void igTabBarQueueReorder(ImGuiTabBar* tab_bar,ImGuiTabItem* tab,int offset);
func TabBarQueueReorder(tab_bar *ImGuiTabBar, tab *ImGuiTabItem, offset int32) {
	giLib.Call(_func_igTabBarQueueReorder_, []interface{}{&tab_bar, &tab, &offset})
}

var _func_igTabBarQueueReorderFromMousePos_ = &c.FuncPrototype{Name: "igTabBarQueueReorderFromMousePos", OutType: c.Void, InTypes: _typs_PPVec2}

// CIMGUI_API void igTabBarQueueReorderFromMousePos(ImGuiTabBar* tab_bar,ImGuiTabItem* tab,ImVec2 mouse_pos);
func TabBarQueueReorderFromMousePos(tab_bar *ImGuiTabBar, tab *ImGuiTabItem, mouse_pos ImVec2) {
	giLib.Call(_func_igTabBarQueueReorderFromMousePos_, []interface{}{&tab_bar, &tab, &mouse_pos})
}

var _func_igTabBarProcessReorder_ = &c.FuncPrototype{Name: "igTabBarProcessReorder", OutType: c.I32, InTypes: _typs_P}

// CIMGUI_API bool igTabBarProcessReorder(ImGuiTabBar* tab_bar);
func TabBarProcessReorder(tab_bar *ImGuiTabBar) bool {
	return giLib.Call(_func_igTabBarProcessReorder_, []interface{}{&tab_bar}).BoolFree()
}

var _func_igTabItemEx_ = &c.FuncPrototype{Name: "igTabItemEx", OutType: c.I32, InTypes: _typs_PPPU32P}

// CIMGUI_API bool igTabItemEx(ImGuiTabBar* tab_bar,const char* label,bool* p_open,ImGuiTabItemFlags flags,ImGuiWindow* docked_window);
func TabItemEx(tab_bar *ImGuiTabBar, label string, p_open *bool, flags ImGuiTabItemFlags, docked_window *ImGuiWindow) bool {
	_label, _p_open := c.CStr(label), cboolPtr(p_open)
	defer usf.Free(_label)
	r := giLib.Call(_func_igTabItemEx_, []interface{}{&tab_bar, &_label, &p_open, &flags, &docked_window}).BoolFree()
	cboolGet(p_open, _p_open)
	return r
}

var _func_igTabItemCalcSize_Str_ = &c.FuncPrototype{Name: "igTabItemCalcSize_Str", OutType: c.Void, InTypes: _typs_PPI32}

// CIMGUI_API void igTabItemCalcSize_Str(ImVec2* pOut,const char* label,bool has_close_button_or_unsaved_marker);
func TabItemCalcSizeStr(pOut *ImVec2, label string, has_close_button_or_unsaved_marker bool) {
	_label, _has_close_button_or_unsaved_marker := c.CStr(label), c.CBool(has_close_button_or_unsaved_marker)
	defer usf.Free(_label)
	giLib.Call(_func_igTabItemCalcSize_Str_, []interface{}{&pOut, &_label, &_has_close_button_or_unsaved_marker})
}

var _func_igTabItemCalcSize_WindowPtr_ = &c.FuncPrototype{Name: "igTabItemCalcSize_WindowPtr", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igTabItemCalcSize_WindowPtr(ImVec2* pOut,ImGuiWindow* window);
func TabItemCalcSizeWindow(pOut *ImVec2, window *ImGuiWindow) {
	giLib.Call(_func_igTabItemCalcSize_WindowPtr_, []interface{}{&pOut, &window})
}

var _func_igTabItemBackground_ = &c.FuncPrototype{Name: "igTabItemBackground", OutType: c.Void, InTypes: _typs_PRectU32U32}

// CIMGUI_API void igTabItemBackground(ImDrawList* draw_list,const ImRect bb,ImGuiTabItemFlags flags,ImU32 col);
func TabItemBackground(draw_list *ImDrawList, bb ImRect, flags ImGuiTabItemFlags, col uint32) {
	giLib.Call(_func_igTabItemBackground_, []interface{}{&draw_list, &bb, &flags, &col})
}

var _func_igTabItemLabelAndCloseButton_ = &c.FuncPrototype{Name: "igTabItemLabelAndCloseButton", OutType: c.Void, InTypes: _typs_PRectU32Vec2PU32U32I32PP}

// CIMGUI_API void igTabItemLabelAndCloseButton(ImDrawList* draw_list,const ImRect bb,ImGuiTabItemFlags flags,ImVec2 frame_padding,const char* label,ImGuiID tab_id,ImGuiID close_button_id,bool is_contents_visible,bool* out_just_closed,bool* out_text_clipped);
func TabItemLabelAndCloseButton(draw_list *ImDrawList, bb ImRect, flags ImGuiTabItemFlags, frame_padding ImVec2, label string, tab_id ImGuiID, close_button_id ImGuiID, is_contents_visible bool, out_just_closed *bool, out_text_clipped *bool) {
	_label, _is_contents_visible, _out_just_closed, _out_text_clipped := c.CStr(label), c.CBool(is_contents_visible), cboolPtr(out_just_closed), cboolPtr(out_text_clipped)
	defer usf.Free(_label)
	giLib.Call(_func_igTabItemLabelAndCloseButton_, []interface{}{&draw_list, &bb, &flags, &frame_padding, &_label, &tab_id, &close_button_id, &_is_contents_visible, &out_just_closed, &out_text_clipped})
	cboolGet(out_just_closed, _out_just_closed)
	cboolGet(out_text_clipped, _out_text_clipped)
}

var _func_igRenderText_ = &c.FuncPrototype{Name: "igRenderText", OutType: c.Void, InTypes: _typs_Vec2PPI32}

// CIMGUI_API void igRenderText(ImVec2 pos,const char* text,const char* text_end,bool hide_text_after_hash);
func RenderText(pos ImVec2, text string, text_end string, hide_text_after_hash bool) {
	_text, _text_end, _hide_text_after_hash := c.CStr(text), c.CStr(text_end), c.CBool(hide_text_after_hash)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	giLib.Call(_func_igRenderText_, []interface{}{&pos, &_text, &_text_end, &_hide_text_after_hash})
}

var _func_igRenderTextWrapped_ = &c.FuncPrototype{Name: "igRenderTextWrapped", OutType: c.Void, InTypes: _typs_Vec2PPF32}

// CIMGUI_API void igRenderTextWrapped(ImVec2 pos,const char* text,const char* text_end,float wrap_width);
func RenderTextWrapped(pos ImVec2, text string, text_end string, wrap_width float32) {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	giLib.Call(_func_igRenderTextWrapped_, []interface{}{&pos, &_text, &_text_end, &wrap_width})
}

var _func_igRenderTextClipped_ = &c.FuncPrototype{Name: "igRenderTextClipped", OutType: c.Void, InTypes: _typs_Vec2Vec2PPPVec2P}

// CIMGUI_API void igRenderTextClipped(const ImVec2 pos_min,const ImVec2 pos_max,const char* text,const char* text_end,const ImVec2* text_size_if_known,const ImVec2 align,const ImRect* clip_rect);
func RenderTextClipped(pos_min ImVec2, pos_max ImVec2, text string, text_end string, text_size_if_known *ImVec2, align ImVec2, clip_rect *ImRect) {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	giLib.Call(_func_igRenderTextClipped_, []interface{}{&pos_min, &pos_max, &_text, &_text_end, &text_size_if_known, &align, &clip_rect})
}

var _func_igRenderTextClippedEx_ = &c.FuncPrototype{Name: "igRenderTextClippedEx", OutType: c.Void, InTypes: _typs_PVec2Vec2PPPVec2P}

// CIMGUI_API void igRenderTextClippedEx(ImDrawList* draw_list,const ImVec2 pos_min,const ImVec2 pos_max,const char* text,const char* text_end,const ImVec2* text_size_if_known,const ImVec2 align,const ImRect* clip_rect);
func RenderTextClippedEx(draw_list *ImDrawList, pos_min ImVec2, pos_max ImVec2, text string, text_end string, text_size_if_known *ImVec2, align ImVec2, clip_rect *ImRect) {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	giLib.Call(_func_igRenderTextClippedEx_, []interface{}{&draw_list, &pos_min, &pos_max, &_text, &_text_end, &text_size_if_known, &align, &clip_rect})
}

var _func_igRenderTextEllipsis_ = &c.FuncPrototype{Name: "igRenderTextEllipsis", OutType: c.Void, InTypes: _typs_PVec2Vec2F32F32PPP}

// CIMGUI_API void igRenderTextEllipsis(ImDrawList* draw_list,const ImVec2 pos_min,const ImVec2 pos_max,float clip_max_x,float ellipsis_max_x,const char* text,const char* text_end,const ImVec2* text_size_if_known);
func RenderTextEllipsis(draw_list *ImDrawList, pos_min ImVec2, pos_max ImVec2, clip_max_x float32, ellipsis_max_x float32, text string, text_end string, text_size_if_known *ImVec2) {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	giLib.Call(_func_igRenderTextEllipsis_, []interface{}{&draw_list, &pos_min, &pos_max, &clip_max_x, &ellipsis_max_x, &_text, &_text_end, &text_size_if_known})
}

var _func_igRenderFrame_ = &c.FuncPrototype{Name: "igRenderFrame", OutType: c.Void, InTypes: _typs_Vec2Vec2U32I32F32}

// CIMGUI_API void igRenderFrame(ImVec2 p_min,ImVec2 p_max,ImU32 fill_col,bool border,float rounding);
func RenderFrame(p_min ImVec2, p_max ImVec2, fill_col uint32, border bool, rounding float32) {
	_border := c.CBool(border)
	giLib.Call(_func_igRenderFrame_, []interface{}{&p_min, &p_max, &fill_col, &_border, &rounding})
}

var _func_igRenderFrameBorder_ = &c.FuncPrototype{Name: "igRenderFrameBorder", OutType: c.Void, InTypes: _typs_Vec2Vec2F32}

// CIMGUI_API void igRenderFrameBorder(ImVec2 p_min,ImVec2 p_max,float rounding);
func RenderFrameBorder(p_min ImVec2, p_max ImVec2, rounding float32) {
	giLib.Call(_func_igRenderFrameBorder_, []interface{}{&p_min, &p_max, &rounding})
}

var _func_igRenderColorRectWithAlphaCheckerboard_ = &c.FuncPrototype{Name: "igRenderColorRectWithAlphaCheckerboard", OutType: c.Void, InTypes: _typs_PVec2Vec2U32F32Vec2F32U32}

// CIMGUI_API void igRenderColorRectWithAlphaCheckerboard(ImDrawList* draw_list,ImVec2 p_min,ImVec2 p_max,ImU32 fill_col,float grid_step,ImVec2 grid_off,float rounding,ImDrawFlags flags);
func RenderColorRectWithAlphaCheckerboard(draw_list *ImDrawList, p_min ImVec2, p_max ImVec2, fill_col uint32, grid_step float32, grid_off ImVec2, rounding float32, flags ImDrawFlags) {
	giLib.Call(_func_igRenderColorRectWithAlphaCheckerboard_, []interface{}{&draw_list, &p_min, &p_max, &fill_col, &grid_step, &grid_off, &rounding, &flags})
}

var _func_igRenderNavHighlight_ = &c.FuncPrototype{Name: "igRenderNavHighlight", OutType: c.Void, InTypes: _typs_RectU32U32}

// CIMGUI_API void igRenderNavHighlight(const ImRect bb,ImGuiID id,ImGuiNavHighlightFlags flags);
func RenderNavHighlight(bb ImRect, id ImGuiID, flags ImGuiNavHighlightFlags) {
	giLib.Call(_func_igRenderNavHighlight_, []interface{}{&bb, &id, &flags})
}

var _func_igFindRenderedTextEnd_ = &c.FuncPrototype{Name: "igFindRenderedTextEnd", OutType: c.Pointer, InTypes: _typs_PP}

// CIMGUI_API const char* igFindRenderedTextEnd(const char* text,const char* text_end);
func FindRenderedTextEnd(text string, text_end string) string {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)

	return giLib.Call(_func_igFindRenderedTextEnd_, []interface{}{&_text, &_text_end}).StrFree()
}

var _func_igRenderMouseCursor_ = &c.FuncPrototype{Name: "igRenderMouseCursor", OutType: c.Void, InTypes: _typs_Vec2F32I32U32U32U32}

// CIMGUI_API void igRenderMouseCursor(ImVec2 pos,float scale,ImGuiMouseCursor mouse_cursor,ImU32 col_fill,ImU32 col_border,ImU32 col_shadow);
func RenderMouseCursor(pos ImVec2, scale float32, mouse_cursor ImGuiMouseCursor, col_fill uint32, col_border uint32, col_shadow uint32) {
	giLib.Call(_func_igRenderMouseCursor_, []interface{}{&pos, &scale, &mouse_cursor, &col_fill, &col_border, &col_shadow})
}

var _func_igRenderArrow_ = &c.FuncPrototype{Name: "igRenderArrow", OutType: c.Void, InTypes: _typs_PVec2U32I32F32}

// CIMGUI_API void igRenderArrow(ImDrawList* draw_list,ImVec2 pos,ImU32 col,ImGuiDir dir,float scale);
func RenderArrow(draw_list *ImDrawList, pos ImVec2, col uint32, dir ImGuiDir, scale float32) {
	giLib.Call(_func_igRenderArrow_, []interface{}{&draw_list, &pos, &col, &dir, &scale})
}

var _func_igRenderBullet_ = &c.FuncPrototype{Name: "igRenderBullet", OutType: c.Void, InTypes: _typs_PVec2U32}

// CIMGUI_API void igRenderBullet(ImDrawList* draw_list,ImVec2 pos,ImU32 col);
func RenderBullet(draw_list *ImDrawList, pos ImVec2, col uint32) {
	giLib.Call(_func_igRenderBullet_, []interface{}{&draw_list, &pos, &col})
}

var _func_igRenderCheckMark_ = &c.FuncPrototype{Name: "igRenderCheckMark", OutType: c.Void, InTypes: _typs_PVec2U32F32}

// CIMGUI_API void igRenderCheckMark(ImDrawList* draw_list,ImVec2 pos,ImU32 col,float sz);
func RenderCheckMark(draw_list *ImDrawList, pos ImVec2, col uint32, sz float32) {
	giLib.Call(_func_igRenderCheckMark_, []interface{}{&draw_list, &pos, &col, &sz})
}

var _func_igRenderArrowPointingAt_ = &c.FuncPrototype{Name: "igRenderArrowPointingAt", OutType: c.Void, InTypes: _typs_PVec2Vec2I32U32}

// CIMGUI_API void igRenderArrowPointingAt(ImDrawList* draw_list,ImVec2 pos,ImVec2 half_sz,ImGuiDir direction,ImU32 col);
func RenderArrowPointingAt(draw_list *ImDrawList, pos ImVec2, half_sz ImVec2, direction ImGuiDir, col uint32) {
	giLib.Call(_func_igRenderArrowPointingAt_, []interface{}{&draw_list, &pos, &half_sz, &direction, &col})
}

var _func_igRenderArrowDockMenu_ = &c.FuncPrototype{Name: "igRenderArrowDockMenu", OutType: c.Void, InTypes: _typs_PVec2F32U32}

// CIMGUI_API void igRenderArrowDockMenu(ImDrawList* draw_list,ImVec2 p_min,float sz,ImU32 col);
func RenderArrowDockMenu(draw_list *ImDrawList, p_min ImVec2, sz float32, col uint32) {
	giLib.Call(_func_igRenderArrowDockMenu_, []interface{}{&draw_list, &p_min, &sz, &col})
}

var _func_igRenderRectFilledRangeH_ = &c.FuncPrototype{Name: "igRenderRectFilledRangeH", OutType: c.Void, InTypes: _typs_PRectU32F32F32F32}

// CIMGUI_API void igRenderRectFilledRangeH(ImDrawList* draw_list,const ImRect rect,ImU32 col,float x_start_norm,float x_end_norm,float rounding);
func RenderRectFilledRangeH(draw_list *ImDrawList, rect ImRect, col uint32, x_start_norm float32, x_end_norm float32, rounding float32) {
	giLib.Call(_func_igRenderRectFilledRangeH_, []interface{}{&draw_list, &rect, &col, &x_start_norm, &x_end_norm, &rounding})
}

var _func_igRenderRectFilledWithHole_ = &c.FuncPrototype{Name: "igRenderRectFilledWithHole", OutType: c.Void, InTypes: _typs_PRectRectU32F32}

// CIMGUI_API void igRenderRectFilledWithHole(ImDrawList* draw_list,const ImRect outer,const ImRect inner,ImU32 col,float rounding);
func RenderRectFilledWithHole(draw_list *ImDrawList, outer ImRect, inner ImRect, col uint32, rounding float32) {
	giLib.Call(_func_igRenderRectFilledWithHole_, []interface{}{&draw_list, &outer, &inner, &col, &rounding})
}

var _func_igCalcRoundingFlagsForRectInRect_ = &c.FuncPrototype{Name: "igCalcRoundingFlagsForRectInRect", OutType: c.U32, InTypes: _typs_RectRectF32}

// CIMGUI_API ImDrawFlags igCalcRoundingFlagsForRectInRect(const ImRect r_in,const ImRect r_outer,float threshold);
func CalcRoundingFlagsForRectInRect(r_in ImRect, r_outer ImRect, threshold float32) ImDrawFlags {
	return (ImDrawFlags)(giLib.Call(_func_igCalcRoundingFlagsForRectInRect_, []interface{}{&r_in, &r_outer, &threshold}).U32Free())
}

var _func_igTextEx_ = &c.FuncPrototype{Name: "igTextEx", OutType: c.Void, InTypes: _typs_PPU32}

// CIMGUI_API void igTextEx(const char* text,const char* text_end,ImGuiTextFlags flags);
func TextEx(text string, text_end string, flags ImGuiTextFlags) {
	_text, _text_end := c.CStr(text), c.CStr(text_end)
	defer usf.Free(_text)
	defer usf.Free(_text_end)
	giLib.Call(_func_igTextEx_, []interface{}{&_text, &_text_end, &flags})
}

var _func_igButtonEx_ = &c.FuncPrototype{Name: "igButtonEx", OutType: c.I32, InTypes: _typs_PVec2U32}

// CIMGUI_API bool igButtonEx(const char* label,const ImVec2 size_arg,ImGuiButtonFlags flags);
func ButtonEx(label string, size_arg ImVec2, flags ImGuiButtonFlags) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)
	return giLib.Call(_func_igButtonEx_, []interface{}{&_label, &size_arg, &flags}).BoolFree()
}

var _func_igArrowButtonEx_ = &c.FuncPrototype{Name: "igArrowButtonEx", OutType: c.I32, InTypes: _typs_PI32Vec2U32}

// CIMGUI_API bool igArrowButtonEx(const char* str_id,ImGuiDir dir,ImVec2 size_arg,ImGuiButtonFlags flags);
func ArrowButtonEx(str_id string, dir ImGuiDir, size_arg ImVec2, flags ImGuiButtonFlags) bool {
	_str_id := c.CStr(str_id)
	defer usf.Free(_str_id)
	return giLib.Call(_func_igArrowButtonEx_, []interface{}{&_str_id, &dir, &size_arg, &flags}).BoolFree()
}

var _func_igImageButtonEx_ = &c.FuncPrototype{Name: "igImageButtonEx", OutType: c.I32, InTypes: _typs_U32PVec2Vec2Vec2Vec4Vec4U32}

// CIMGUI_API bool igImageButtonEx(ImGuiID id,ImTextureID texture_id,const ImVec2 size,const ImVec2 uv0,const ImVec2 uv1,const ImVec4 bg_col,const ImVec4 tint_col,ImGuiButtonFlags flags);
func ImageButtonEx(id ImGuiID, texture_id ImTextureID, size ImVec2, uv0 ImVec2, uv1 ImVec2, bg_col ImVec4, tint_col ImVec4, flags ImGuiButtonFlags) bool {
	return giLib.Call(_func_igImageButtonEx_, []interface{}{&id, &texture_id, &size, &uv0, &uv1, &bg_col, &tint_col, &flags}).BoolFree()
}

var _func_igSeparatorEx_ = &c.FuncPrototype{Name: "igSeparatorEx", OutType: c.Void, InTypes: _typs_U32F32}

// CIMGUI_API void igSeparatorEx(ImGuiSeparatorFlags flags,float thickness);
func SeparatorEx(flags ImGuiSeparatorFlags, thickness float32) {
	giLib.Call(_func_igSeparatorEx_, []interface{}{&flags, &thickness})
}

var _func_igSeparatorTextEx_ = &c.FuncPrototype{Name: "igSeparatorTextEx", OutType: c.Void, InTypes: _typs_U32PPF32}

// CIMGUI_API void igSeparatorTextEx(ImGuiID id,const char* label,const char* label_end,float extra_width);
func SeparatorTextEx(id ImGuiID, label string, label_end string, extra_width float32) {
	_label, _label_end := c.CStr(label), c.CStr(label_end)
	defer usf.Free(_label)
	defer usf.Free(_label_end)
	giLib.Call(_func_igSeparatorTextEx_, []interface{}{&id, &_label, &_label_end, &extra_width})
}

var _func_igCheckboxFlags_S64Ptr_ = &c.FuncPrototype{Name: "igCheckboxFlags_S64Ptr", OutType: c.I32, InTypes: _typs_PPI64}

// CIMGUI_API bool igCheckboxFlags_S64Ptr(const char* label,ImS64* flags,int64 flags_value);
func CheckboxFlagsI64Ptr(label string, flags *ImS64, flags_value int64) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)

	return giLib.Call(_func_igCheckboxFlags_S64Ptr_, []interface{}{&_label, &flags, &flags_value}).BoolFree()
}

var _func_igCheckboxFlags_U64Ptr_ = &c.FuncPrototype{Name: "igCheckboxFlags_U64Ptr", OutType: c.I32, InTypes: _typs_PPU64}

// CIMGUI_API bool igCheckboxFlags_U64Ptr(const char* label,ImU64* flags,uint64 flags_value);
func CheckboxFlagsU64Ptr(label string, flags *ImU64, flags_value uint64) bool {
	_label := c.CStr(label)
	defer usf.Free(_label)

	return giLib.Call(_func_igCheckboxFlags_U64Ptr_, []interface{}{&_label, &flags, &flags_value}).BoolFree()
}

var _func_igCloseButton_ = &c.FuncPrototype{Name: "igCloseButton", OutType: c.I32, InTypes: _typs_U32Vec2}

// CIMGUI_API bool igCloseButton(ImGuiID id,const ImVec2 pos);
func CloseButton(id ImGuiID, pos ImVec2) bool {
	return giLib.Call(_func_igCloseButton_, []interface{}{&id, &pos}).BoolFree()
}

var _func_igCollapseButton_ = &c.FuncPrototype{Name: "igCollapseButton", OutType: c.I32, InTypes: _typs_U32Vec2P}

// CIMGUI_API bool igCollapseButton(ImGuiID id,const ImVec2 pos,ImGuiDockNode* dock_node);
func CollapseButton(id ImGuiID, pos ImVec2, dock_node *ImGuiDockNode) bool {
	return giLib.Call(_func_igCollapseButton_, []interface{}{&id, &pos, &dock_node}).BoolFree()
}

var _func_igScrollbar_ = &c.FuncPrototype{Name: "igScrollbar", OutType: c.Void, InTypes: _typs_I32}

// CIMGUI_API void igScrollbar(ImGuiAxis axis);
func Scrollbar(axis ImGuiAxis) {
	giLib.Call(_func_igScrollbar_, []interface{}{&axis})
}

var _func_igScrollbarEx_ = &c.FuncPrototype{Name: "igScrollbarEx", OutType: c.I32, InTypes: _typs_RectU32I32PI64I64U32}

// CIMGUI_API bool igScrollbarEx(const ImRect bb,ImGuiID id,ImGuiAxis axis,int64* p_scroll_v,int64 avail_v,int64 contents_v,ImDrawFlags flags);
func ScrollbarEx(bb ImRect, id ImGuiID, axis ImGuiAxis, p_scroll_v *int64, avail_v int64, contents_v int64, flags ImDrawFlags) bool {
	return giLib.Call(_func_igScrollbarEx_, []interface{}{&bb, &id, &axis, &p_scroll_v, &avail_v, &contents_v, &flags}).BoolFree()
}

var _func_igGetWindowScrollbarRect_ = &c.FuncPrototype{Name: "igGetWindowScrollbarRect", OutType: c.Void, InTypes: _typs_PPI32}

// CIMGUI_API void igGetWindowScrollbarRect(ImRect* pOut,ImGuiWindow* window,ImGuiAxis axis);
func GetWindowScrollbarRect(pOut *ImRect, window *ImGuiWindow, axis ImGuiAxis) {
	giLib.Call(_func_igGetWindowScrollbarRect_, []interface{}{&pOut, &window, &axis})
}

var _func_igGetWindowScrollbarID_ = &c.FuncPrototype{Name: "igGetWindowScrollbarID", OutType: c.U32, InTypes: _typs_PI32}

// CIMGUI_API ImGuiID igGetWindowScrollbarID(ImGuiWindow* window,ImGuiAxis axis);
func GetWindowScrollbarID(window *ImGuiWindow, axis ImGuiAxis) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetWindowScrollbarID_, []interface{}{&window, &axis}).U32Free())
}

var _func_igGetWindowResizeCornerID_ = &c.FuncPrototype{Name: "igGetWindowResizeCornerID", OutType: c.U32, InTypes: _typs_PI32}

// CIMGUI_API ImGuiID igGetWindowResizeCornerID(ImGuiWindow* window,int n);
func GetWindowResizeCornerID(window *ImGuiWindow, n int32) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetWindowResizeCornerID_, []interface{}{&window, &n}).U32Free())
}

var _func_igGetWindowResizeBorderID_ = &c.FuncPrototype{Name: "igGetWindowResizeBorderID", OutType: c.U32, InTypes: _typs_PI32}

// CIMGUI_API ImGuiID igGetWindowResizeBorderID(ImGuiWindow* window,ImGuiDir dir);
func GetWindowResizeBorderID(window *ImGuiWindow, dir ImGuiDir) ImGuiID {
	return (ImGuiID)(giLib.Call(_func_igGetWindowResizeBorderID_, []interface{}{&window, &dir}).U32Free())
}

var _func_igButtonBehavior_ = &c.FuncPrototype{Name: "igButtonBehavior", OutType: c.I32, InTypes: _typs_RectU32PPU32}

// CIMGUI_API bool igButtonBehavior(const ImRect bb,ImGuiID id,bool* out_hovered,bool* out_held,ImGuiButtonFlags flags);
func ButtonBehavior(bb ImRect, id ImGuiID, out_hovered *bool, out_held *bool, flags ImGuiButtonFlags) bool {
	_out_hovered, _out_held := cboolPtr(out_hovered), cboolPtr(out_held)
	r := giLib.Call(_func_igButtonBehavior_, []interface{}{&bb, &id, &out_hovered, &out_held, &flags}).BoolFree()
	cboolGet(out_hovered, _out_hovered)
	cboolGet(out_held, _out_held)
	return r
}

var _func_igDragBehavior_ = &c.FuncPrototype{Name: "igDragBehavior", OutType: c.I32, InTypes: _typs_U32U32PF32PPPU32}

// CIMGUI_API bool igDragBehavior(ImGuiID id,ImGuiDataType data_type,void* p_v,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags);
func DragBehavior(id ImGuiID, data_type ImGuiDataType, p_v unsafe.Pointer, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	_format := c.CStr(format)
	defer usf.Free(_format)
	return giLib.Call(_func_igDragBehavior_, []interface{}{&id, &data_type, &p_v, &v_speed, &p_min, &p_max, &_format, &flags}).BoolFree()
}

var _func_igSliderBehavior_ = &c.FuncPrototype{Name: "igSliderBehavior", OutType: c.I32, InTypes: _typs_RectU32U32PPPPU32P}

// CIMGUI_API bool igSliderBehavior(const ImRect bb,ImGuiID id,ImGuiDataType data_type,void* p_v,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags,ImRect* out_grab_bb);
func SliderBehavior(bb ImRect, id ImGuiID, data_type ImGuiDataType, p_v unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags, out_grab_bb *ImRect) bool {
	_format := c.CStr(format)
	defer usf.Free(_format)

	return giLib.Call(_func_igSliderBehavior_, []interface{}{&bb, &id, &data_type, &p_v, &p_min, &p_max, &_format, &flags, &out_grab_bb}).BoolFree()
}

var _func_igSplitterBehavior_ = &c.FuncPrototype{Name: "igSplitterBehavior", OutType: c.I32, InTypes: _typs_RectU32I32PPF32F32F32F32U32}

// CIMGUI_API bool igSplitterBehavior(const ImRect bb,ImGuiID id,ImGuiAxis axis,float* size1,float* size2,float min_size1,float min_size2,float hover_extend,float hover_visibility_delay,ImU32 bg_col);
func SplitterBehavior(bb ImRect, id ImGuiID, axis ImGuiAxis, size1 *float32, size2 *float32, min_size1 float32, min_size2 float32, hover_extend float32, hover_visibility_delay float32, bg_col uint32) bool {
	return giLib.Call(_func_igSplitterBehavior_, []interface{}{&bb, &id, &axis, &size1, &size2, &min_size1, &min_size2, &hover_extend, &hover_visibility_delay, &bg_col}).BoolFree()
}

var _func_igTreeNodeBehavior_ = &c.FuncPrototype{Name: "igTreeNodeBehavior", OutType: c.I32, InTypes: _typs_U32U32PP}

// CIMGUI_API bool igTreeNodeBehavior(ImGuiID id,ImGuiTreeNodeFlags flags,const char* label,const char* label_end);
func TreeNodeBehavior(id ImGuiID, flags ImGuiTreeNodeFlags, label string, label_end string) bool {
	_label, _label_end := c.CStr(label), c.CStr(label_end)
	defer usf.Free(_label)
	defer usf.Free(_label_end)
	return giLib.Call(_func_igTreeNodeBehavior_, []interface{}{&id, &flags, &_label, &_label_end}).BoolFree()
}

var _func_igTreePushOverrideID_ = &c.FuncPrototype{Name: "igTreePushOverrideID", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igTreePushOverrideID(ImGuiID id);
func TreePushOverrideID(id ImGuiID) {
	giLib.Call(_func_igTreePushOverrideID_, []interface{}{&id})
}

var _func_igTreeNodeSetOpen_ = &c.FuncPrototype{Name: "igTreeNodeSetOpen", OutType: c.Void, InTypes: _typs_U32I32}

// CIMGUI_API void igTreeNodeSetOpen(ImGuiID id,bool open);
func TreeNodeSetOpen(id ImGuiID, open bool) {
	_open := c.CBool(open)
	giLib.Call(_func_igTreeNodeSetOpen_, []interface{}{&id, &_open})
}

var _func_igTreeNodeUpdateNextOpen_ = &c.FuncPrototype{Name: "igTreeNodeUpdateNextOpen", OutType: c.I32, InTypes: _typs_U32U32}

// CIMGUI_API bool igTreeNodeUpdateNextOpen(ImGuiID id,ImGuiTreeNodeFlags flags);
func TreeNodeUpdateNextOpen(id ImGuiID, flags ImGuiTreeNodeFlags) bool {
	return giLib.Call(_func_igTreeNodeUpdateNextOpen_, []interface{}{&id, &flags}).BoolFree()
}

var _func_igDataTypeGetInfo_ = &c.FuncPrototype{Name: "igDataTypeGetInfo", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API const ImGuiDataTypeInfo* igDataTypeGetInfo(ImGuiDataType data_type);
func DataTypeGetInfo(data_type ImGuiDataType) *ImGuiDataTypeInfo {
	return (*ImGuiDataTypeInfo)(giLib.Call(_func_igDataTypeGetInfo_, []interface{}{&data_type}).PtrFree())
}

var _func_igDataTypeFormatString_ = &c.FuncPrototype{Name: "igDataTypeFormatString", OutType: c.I32, InTypes: _typs_PI32U32PP}

// CIMGUI_API int igDataTypeFormatString(char* buf,int buf_size,ImGuiDataType data_type,const void* p_data,const char* format);
func DataTypeFormatString(buf string, buf_size int32, data_type ImGuiDataType, p_data unsafe.Pointer, format string) int32 {
	_buf, _format := c.CStr(buf), c.CStr(format)
	defer usf.Free(_buf)
	defer usf.Free(_format)

	return giLib.Call(_func_igDataTypeFormatString_, []interface{}{&_buf, &buf_size, &data_type, &p_data, &_format}).I32Free()
}

var _func_igDataTypeApplyOp_ = &c.FuncPrototype{Name: "igDataTypeApplyOp", OutType: c.Void, InTypes: _typs_U32I32PPP}

// CIMGUI_API void igDataTypeApplyOp(ImGuiDataType data_type,int op,void* output,const void* arg_1,const void* arg_2);
func DataTypeApplyOp(data_type ImGuiDataType, op int32, output, arg_1, arg_2 unsafe.Pointer) {
	giLib.Call(_func_igDataTypeApplyOp_, []interface{}{&data_type, &op, &output, &arg_1, &arg_2})
}

var _func_igDataTypeApplyFromText_ = &c.FuncPrototype{Name: "igDataTypeApplyFromText", OutType: c.I32, InTypes: _typs_PU32PP}

// CIMGUI_API bool igDataTypeApplyFromText(const char* buf,ImGuiDataType data_type,void* p_data,const char* format);
func DataTypeApplyFromText(buf string, data_type ImGuiDataType, p_data unsafe.Pointer, format string) bool {
	_buf, _format := c.CStr(buf), c.CStr(format)
	defer usf.Free(_buf)
	defer usf.Free(_format)

	return giLib.Call(_func_igDataTypeApplyFromText_, []interface{}{&_buf, &data_type, &p_data, &_format}).BoolFree()
}

var _func_igDataTypeCompare_ = &c.FuncPrototype{Name: "igDataTypeCompare", OutType: c.I32, InTypes: _typs_U32PP}

// CIMGUI_API int igDataTypeCompare(ImGuiDataType data_type,const void* arg_1,const void* arg_2);
func DataTypeCompare(data_type ImGuiDataType, arg_1, arg_2 unsafe.Pointer) int32 {
	return giLib.Call(_func_igDataTypeCompare_, []interface{}{&data_type, &arg_1, &arg_2}).I32Free()
}

var _func_igDataTypeClamp_ = &c.FuncPrototype{Name: "igDataTypeClamp", OutType: c.I32, InTypes: _typs_U32PPP}

// CIMGUI_API bool igDataTypeClamp(ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max);
func DataTypeClamp(data_type ImGuiDataType, p_data, p_min, p_max unsafe.Pointer) bool {
	return giLib.Call(_func_igDataTypeClamp_, []interface{}{&data_type, &p_data, &p_min, &p_max}).BoolFree()
}

// CIMGUI_API bool igInputTextEx(const char* label,const char* hint,char* buf,int buf_size,const ImVec2 size_arg,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data);

var _func_igInputTextDeactivateHook_ = &c.FuncPrototype{Name: "igInputTextDeactivateHook", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igInputTextDeactivateHook(ImGuiID id);
func InputTextDeactivateHook(id ImGuiID) {
	giLib.Call(_func_igInputTextDeactivateHook_, []interface{}{&id})
}

var _func_igTempInputText_ = &c.FuncPrototype{Name: "igTempInputText", OutType: c.I32, InTypes: _typs_RectU32PPI32U32}

// CIMGUI_API bool igTempInputText(const ImRect bb,ImGuiID id,const char* label,char* buf,int buf_size,ImGuiInputTextFlags flags);
func TempInputText(bb ImRect, id ImGuiID, label string, buf string, buf_size int32, flags ImGuiInputTextFlags) bool {
	_label, _buf := c.CStr(label), c.CStr(buf)
	defer usf.Free(_label)
	defer usf.Free(_buf)

	return giLib.Call(_func_igTempInputText_, []interface{}{&bb, &id, &_label, &_buf, &buf_size, &flags}).BoolFree()
}

var _func_igTempInputScalar_ = &c.FuncPrototype{Name: "igTempInputScalar", OutType: c.I32, InTypes: _typs_RectU32PU32PPPP}

// CIMGUI_API bool igTempInputScalar(const ImRect bb,ImGuiID id,const char* label,ImGuiDataType data_type,void* p_data,const char* format,const void* p_clamp_min,const void* p_clamp_max);
func TempInputScalar(bb ImRect, id ImGuiID, label string, data_type ImGuiDataType, p_data unsafe.Pointer, format string, p_clamp_min unsafe.Pointer, p_clamp_max unsafe.Pointer) bool {
	_label, _format := c.CStr(label), c.CStr(format)
	defer usf.Free(_label)
	defer usf.Free(_format)

	return giLib.Call(_func_igTempInputScalar_, []interface{}{&bb, &id, &_label, &data_type, &p_data, &_format, &p_clamp_min, &p_clamp_max}).BoolFree()
}

var _func_igTempInputIsActive_ = &c.FuncPrototype{Name: "igTempInputIsActive", OutType: c.I32, InTypes: _typs_U32}

// CIMGUI_API bool igTempInputIsActive(ImGuiID id);
func TempInputIsActive(id ImGuiID) bool {
	return giLib.Call(_func_igTempInputIsActive_, []interface{}{&id}).BoolFree()
}

var _func_igGetInputTextState_ = &c.FuncPrototype{Name: "igGetInputTextState", OutType: c.Pointer, InTypes: _typs_U32}

// CIMGUI_API ImGuiInputTextState* igGetInputTextState(ImGuiID id);
func GetInputTextState(id ImGuiID) *ImGuiInputTextState {
	return (*ImGuiInputTextState)(giLib.Call(_func_igGetInputTextState_, []interface{}{&id}).PtrFree())
}

var _func_igColorTooltip_ = &c.FuncPrototype{Name: "igColorTooltip", OutType: c.Void, InTypes: _typs_PPU32}

// CIMGUI_API void igColorTooltip(const char* text,const float* col,ImGuiColorEditFlags flags);
func ColorTooltip(text string, col *float32, flags ImGuiColorEditFlags) {
	_text := c.CStr(text)
	defer usf.Free(_text)
	giLib.Call(_func_igColorTooltip_, []interface{}{&_text, &col, &flags})
}

var _func_igColorEditOptionsPopup_ = &c.FuncPrototype{Name: "igColorEditOptionsPopup", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void igColorEditOptionsPopup(const float* col,ImGuiColorEditFlags flags);
func ColorEditOptionsPopup(col *float32, flags ImGuiColorEditFlags) {
	giLib.Call(_func_igColorEditOptionsPopup_, []interface{}{&col, &flags})
}

var _func_igColorPickerOptionsPopup_ = &c.FuncPrototype{Name: "igColorPickerOptionsPopup", OutType: c.Void, InTypes: _typs_PU32}

// CIMGUI_API void igColorPickerOptionsPopup(const float* ref_col,ImGuiColorEditFlags flags);
func ColorPickerOptionsPopup(ref_col *float32, flags ImGuiColorEditFlags) {
	giLib.Call(_func_igColorPickerOptionsPopup_, []interface{}{&ref_col, &flags})
}

// CIMGUI_API int igPlotEx(ImGuiPlotType plot_type,const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,const ImVec2 size_arg);

var _func_igShadeVertsLinearColorGradientKeepAlpha_ = &c.FuncPrototype{Name: "igShadeVertsLinearColorGradientKeepAlpha", OutType: c.Void, InTypes: _typs_PI32I32Vec2Vec2U32U32}

// CIMGUI_API void igShadeVertsLinearColorGradientKeepAlpha(ImDrawList* draw_list,int vert_start_idx,int vert_end_idx,ImVec2 gradient_p0,ImVec2 gradient_p1,ImU32 col0,ImU32 col1);
func ShadeVertsLinearColorGradientKeepAlpha(draw_list *ImDrawList, vert_start_idx int32, vert_end_idx int32, gradient_p0 ImVec2, gradient_p1 ImVec2, col0 uint32, col1 uint32) {
	giLib.Call(_func_igShadeVertsLinearColorGradientKeepAlpha_, []interface{}{&draw_list, &vert_start_idx, &vert_end_idx, &gradient_p0, &gradient_p1, &col0, &col1})
}

var _func_igShadeVertsLinearUV_ = &c.FuncPrototype{Name: "igShadeVertsLinearUV", OutType: c.Void, InTypes: _typs_PI32I32Vec2Vec2Vec2Vec2I32}

// CIMGUI_API void igShadeVertsLinearUV(ImDrawList* draw_list,int vert_start_idx,int vert_end_idx,const ImVec2 a,const ImVec2 b,const ImVec2 uv_a,const ImVec2 uv_b,bool clamp);
func ShadeVertsLinearUV(draw_list *ImDrawList, vert_start_idx int32, vert_end_idx int32, a ImVec2, b ImVec2, uv_a ImVec2, uv_b ImVec2, clamp bool) {
	_clamp := c.CBool(clamp)
	giLib.Call(_func_igShadeVertsLinearUV_, []interface{}{&draw_list, &vert_start_idx, &vert_end_idx, &a, &b, &uv_a, &uv_b, &_clamp})
}

var _func_igGcCompactTransientMiscBuffers_ = &c.FuncPrototype{Name: "igGcCompactTransientMiscBuffers", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igGcCompactTransientMiscBuffers(void);
func GcCompactTransientMiscBuffers() {
	giLib.Call(_func_igGcCompactTransientMiscBuffers_, nil)
}

var _func_igGcCompactTransientWindowBuffers_ = &c.FuncPrototype{Name: "igGcCompactTransientWindowBuffers", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igGcCompactTransientWindowBuffers(ImGuiWindow* window);
func GcCompactTransientWindowBuffers(window *ImGuiWindow) {
	giLib.Call(_func_igGcCompactTransientWindowBuffers_, []interface{}{&window})
}

var _func_igGcAwakeTransientWindowBuffers_ = &c.FuncPrototype{Name: "igGcAwakeTransientWindowBuffers", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igGcAwakeTransientWindowBuffers(ImGuiWindow* window);
func GcAwakeTransientWindowBuffers(window *ImGuiWindow) {
	giLib.Call(_func_igGcAwakeTransientWindowBuffers_, []interface{}{&window})
}

// CIMGUI_API void igDebugLog(const char* fmt,...);
// CIMGUI_API void igDebugLogV(const char* fmt,va_list args);
// CIMGUI_API void igErrorCheckEndFrameRecover(ImGuiErrorLogCallback log_callback,void* user_data);
// CIMGUI_API void igErrorCheckEndWindowRecover(ImGuiErrorLogCallback log_callback,void* user_data);
// CIMGUI_API void igErrorCheckUsingSetCursorPosToExtendParentBoundaries(void);

var _func_igDebugLocateItem_ = &c.FuncPrototype{Name: "igDebugLocateItem", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igDebugLocateItem(ImGuiID target_id);
func DebugLocateItem(target_id ImGuiID) {
	giLib.Call(_func_igDebugLocateItem_, []interface{}{&target_id})
}

var _func_igDebugLocateItemOnHover_ = &c.FuncPrototype{Name: "igDebugLocateItemOnHover", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igDebugLocateItemOnHover(ImGuiID target_id);
func DebugLocateItemOnHover(target_id ImGuiID) {
	giLib.Call(_func_igDebugLocateItemOnHover_, []interface{}{&target_id})
}

var _func_igDebugLocateItemResolveWithLastItem_ = &c.FuncPrototype{Name: "igDebugLocateItemResolveWithLastItem", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igDebugLocateItemResolveWithLastItem(void);
func DebugLocateItemResolveWithLastItem() {
	giLib.Call(_func_igDebugLocateItemResolveWithLastItem_, nil)
}

var _func_igDebugDrawItemRect_ = &c.FuncPrototype{Name: "igDebugDrawItemRect", OutType: c.Void, InTypes: _typs_U32}

// CIMGUI_API void igDebugDrawItemRect(ImU32 col);
func DebugDrawItemRect(col uint32) {
	giLib.Call(_func_igDebugDrawItemRect_, []interface{}{&col})
}

var _func_igDebugStartItemPicker_ = &c.FuncPrototype{Name: "igDebugStartItemPicker", OutType: c.Void, InTypes: nil}

// CIMGUI_API void igDebugStartItemPicker(void);
func DebugStartItemPicker() {
	giLib.Call(_func_igDebugStartItemPicker_, nil)
}

var _func_igShowFontAtlas_ = &c.FuncPrototype{Name: "igShowFontAtlas", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igShowFontAtlas(ImFontAtlas* atlas);
func ShowFontAtlas(atlas *ImFontAtlas) {
	giLib.Call(_func_igShowFontAtlas_, []interface{}{&atlas})
}

// CIMGUI_API void igDebugHookIdInfo(ImGuiID id,ImGuiDataType data_type,const void* data_id,const void* data_id_end);

var _func_igDebugNodeColumns_ = &c.FuncPrototype{Name: "igDebugNodeColumns", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDebugNodeColumns(ImGuiOldColumns* columns);
func DebugNodeColumns(columns *ImGuiOldColumns) {
	giLib.Call(_func_igDebugNodeColumns_, []interface{}{&columns})
}

var _func_igDebugNodeDockNode_ = &c.FuncPrototype{Name: "igDebugNodeDockNode", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDebugNodeDockNode(ImGuiDockNode* node,const char* label);
func DebugNodeDockNode(node *ImGuiDockNode, label string) {
	_label := c.CStr(label)
	defer usf.Free(_label)
	giLib.Call(_func_igDebugNodeDockNode_, []interface{}{&node, &_label})
}

var _func_igDebugNodeDrawList_ = &c.FuncPrototype{Name: "igDebugNodeDrawList", OutType: c.Void, InTypes: _typs_PPPP}

// CIMGUI_API void igDebugNodeDrawList(ImGuiWindow* window,ImGuiViewportP* viewport,const ImDrawList* draw_list,const char* label);
func DebugNodeDrawList(window *ImGuiWindow, viewport *ImGuiViewportP, draw_list *ImDrawList, label string) {
	_label := c.CStr(label)
	defer usf.Free(_label)
	giLib.Call(_func_igDebugNodeDrawList_, []interface{}{&window, &viewport, &draw_list, &_label})
}

var _func_igDebugNodeDrawCmdShowMeshAndBoundingBox_ = &c.FuncPrototype{Name: "igDebugNodeDrawCmdShowMeshAndBoundingBox", OutType: c.Void, InTypes: _typs_PPPI32I32}

// CIMGUI_API void igDebugNodeDrawCmdShowMeshAndBoundingBox(ImDrawList* out_draw_list,const ImDrawList* draw_list,const ImDrawCmd* draw_cmd,bool show_mesh,bool show_aabb);
func DebugNodeDrawCmdShowMeshAndBoundingBox(out_draw_list *ImDrawList, draw_list *ImDrawList, draw_cmd *ImDrawCmd, show_mesh bool, show_aabb bool) {
	_show_mesh, _show_aabb := c.CBool(show_mesh), c.CBool(show_aabb)
	giLib.Call(_func_igDebugNodeDrawCmdShowMeshAndBoundingBox_, []interface{}{&out_draw_list, &draw_list, &draw_cmd, &_show_mesh, &_show_aabb})
}

var _func_igDebugNodeFont_ = &c.FuncPrototype{Name: "igDebugNodeFont", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDebugNodeFont(ImFont* font);
func DebugNodeFont(font *ImFont) {
	giLib.Call(_func_igDebugNodeFont_, []interface{}{&font})
}

var _func_igDebugNodeFontGlyph_ = &c.FuncPrototype{Name: "igDebugNodeFontGlyph", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDebugNodeFontGlyph(ImFont* font,const ImFontGlyph* glyph);
func DebugNodeFontGlyph(font *ImFont, glyph *ImFontGlyph) {
	giLib.Call(_func_igDebugNodeFontGlyph_, []interface{}{&font, &glyph})
}

var _func_igDebugNodeStorage_ = &c.FuncPrototype{Name: "igDebugNodeStorage", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDebugNodeStorage(ImGuiStorage* storage,const char* label);
func DebugNodeStorage(storage *ImGuiStorage, label string) {
	_label := c.CStr(label)
	defer usf.Free(_label)
	giLib.Call(_func_igDebugNodeStorage_, []interface{}{&storage, &_label})
}

var _func_igDebugNodeTabBar_ = &c.FuncPrototype{Name: "igDebugNodeTabBar", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDebugNodeTabBar(ImGuiTabBar* tab_bar,const char* label);
func DebugNodeTabBar(tab_bar *ImGuiTabBar, label string) {
	_label := c.CStr(label)
	defer usf.Free(_label)
	giLib.Call(_func_igDebugNodeTabBar_, []interface{}{&tab_bar, &_label})
}

var _func_igDebugNodeTable_ = &c.FuncPrototype{Name: "igDebugNodeTable", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDebugNodeTable(ImGuiTable* table);
func DebugNodeTable(table *ImGuiTable) {
	giLib.Call(_func_igDebugNodeTable_, []interface{}{&table})
}

var _func_igDebugNodeTableSettings_ = &c.FuncPrototype{Name: "igDebugNodeTableSettings", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDebugNodeTableSettings(ImGuiTableSettings* settings);
func DebugNodeTableSettings(settings *ImGuiTableSettings) {
	giLib.Call(_func_igDebugNodeTableSettings_, []interface{}{&settings})
}

var _func_igDebugNodeInputTextState_ = &c.FuncPrototype{Name: "igDebugNodeInputTextState", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDebugNodeInputTextState(ImGuiInputTextState* state);
func DebugNodeInputTextState(state *ImGuiInputTextState) {
	giLib.Call(_func_igDebugNodeInputTextState_, []interface{}{&state})
}

var _func_igDebugNodeWindow_ = &c.FuncPrototype{Name: "igDebugNodeWindow", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDebugNodeWindow(ImGuiWindow* window,const char* label);
func DebugNodeWindow(window *ImGuiWindow, label string) {
	_label := c.CStr(label)
	defer usf.Free(_label)
	giLib.Call(_func_igDebugNodeWindow_, []interface{}{&window, &_label})
}

var _func_igDebugNodeWindowSettings_ = &c.FuncPrototype{Name: "igDebugNodeWindowSettings", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDebugNodeWindowSettings(ImGuiWindowSettings* settings);
func DebugNodeWindowSettings(settings *ImGuiWindowSettings) {
	giLib.Call(_func_igDebugNodeWindowSettings_, []interface{}{&settings})
}

var _func_igDebugNodeWindowsList_ = &c.FuncPrototype{Name: "igDebugNodeWindowsList", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igDebugNodeWindowsList(ImVector_ImGuiWindowPtr* windows,const char* label);
func DebugNodeWindowsList(windows *ImVector_ImGuiWindowPtr, label string) {
	_label := c.CStr(label)
	defer usf.Free(_label)
	giLib.Call(_func_igDebugNodeWindowsList_, []interface{}{&windows, &_label})
}

var _func_igDebugNodeWindowsListByBeginStackParent_ = &c.FuncPrototype{Name: "igDebugNodeWindowsListByBeginStackParent", OutType: c.Void, InTypes: _typs_PI32P}

// CIMGUI_API void igDebugNodeWindowsListByBeginStackParent(ImGuiWindow** windows,int windows_size,ImGuiWindow* parent_in_begin_stack);
func DebugNodeWindowsListByBeginStackParent(windows *ImGuiWindow, windows_size int32, parent_in_begin_stack *ImGuiWindow) {
	giLib.Call(_func_igDebugNodeWindowsListByBeginStackParent_, []interface{}{&windows, &windows_size, &parent_in_begin_stack})
}

var _func_igDebugNodeViewport_ = &c.FuncPrototype{Name: "igDebugNodeViewport", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDebugNodeViewport(ImGuiViewportP* viewport);
func DebugNodeViewport(viewport *ImGuiViewportP) {
	giLib.Call(_func_igDebugNodeViewport_, []interface{}{&viewport})
}

var _func_igDebugRenderKeyboardPreview_ = &c.FuncPrototype{Name: "igDebugRenderKeyboardPreview", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igDebugRenderKeyboardPreview(ImDrawList* draw_list);
func DebugRenderKeyboardPreview(draw_list *ImDrawList) {
	giLib.Call(_func_igDebugRenderKeyboardPreview_, []interface{}{&draw_list})
}

var _func_igDebugRenderViewportThumbnail_ = &c.FuncPrototype{Name: "igDebugRenderViewportThumbnail", OutType: c.Void, InTypes: _typs_PPRect}

// CIMGUI_API void igDebugRenderViewportThumbnail(ImDrawList* draw_list,ImGuiViewportP* viewport,const ImRect bb);
func DebugRenderViewportThumbnail(draw_list *ImDrawList, viewport *ImGuiViewportP, bb ImRect) {
	giLib.Call(_func_igDebugRenderViewportThumbnail_, []interface{}{&draw_list, &viewport, &bb})
}

var _func_igIsKeyPressedMap_ = &c.FuncPrototype{Name: "igIsKeyPressedMap", OutType: c.I32, InTypes: _typs_U32I32}

// CIMGUI_API bool igIsKeyPressedMap(ImGuiKey key,bool repeat);
func IsKeyPressedMap(key ImGuiKey, repeat bool) bool {
	_repeat := c.CBool(repeat)

	return giLib.Call(_func_igIsKeyPressedMap_, []interface{}{&key, &_repeat}).BoolFree()
}

var _func_igImFontAtlasGetBuilderForStbTruetype_ = &c.FuncPrototype{Name: "igImFontAtlasGetBuilderForStbTruetype", OutType: c.Pointer, InTypes: nil}

// CIMGUI_API const ImFontBuilderIO* igImFontAtlasGetBuilderForStbTruetype(void);
func ImFontAtlasGetBuilderForStbTruetype() *ImFontBuilderIO {
	return (*ImFontBuilderIO)(giLib.Call(_func_igImFontAtlasGetBuilderForStbTruetype_, nil).PtrFree())
}

var _func_igImFontAtlasBuildInit_ = &c.FuncPrototype{Name: "igImFontAtlasBuildInit", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igImFontAtlasBuildInit(ImFontAtlas* atlas);
func ImFontAtlasBuildInit(atlas *ImFontAtlas) {
	giLib.Call(_func_igImFontAtlasBuildInit_, []interface{}{&atlas})
}

var _func_igImFontAtlasBuildSetupFont_ = &c.FuncPrototype{Name: "igImFontAtlasBuildSetupFont", OutType: c.Void, InTypes: _typs_PPPF32F32}

// CIMGUI_API void igImFontAtlasBuildSetupFont(ImFontAtlas* atlas,ImFont* font,ImFontConfig* font_config,float ascent,float descent);
func ImFontAtlasBuildSetupFont(atlas *ImFontAtlas, font *ImFont, font_config *ImFontConfig, ascent float32, descent float32) {
	giLib.Call(_func_igImFontAtlasBuildSetupFont_, []interface{}{&atlas, &font, &font_config, &ascent, &descent})
}

var _func_igImFontAtlasBuildPackCustomRects_ = &c.FuncPrototype{Name: "igImFontAtlasBuildPackCustomRects", OutType: c.Void, InTypes: _typs_PP}

// CIMGUI_API void igImFontAtlasBuildPackCustomRects(ImFontAtlas* atlas,void* stbrp_context_opaque);
func ImFontAtlasBuildPackCustomRects(atlas *ImFontAtlas, stbrp_context_opaque unsafe.Pointer) {
	giLib.Call(_func_igImFontAtlasBuildPackCustomRects_, []interface{}{&atlas, &stbrp_context_opaque})
}

var _func_igImFontAtlasBuildFinish_ = &c.FuncPrototype{Name: "igImFontAtlasBuildFinish", OutType: c.Void, InTypes: _typs_P}

// CIMGUI_API void igImFontAtlasBuildFinish(ImFontAtlas* atlas);
func ImFontAtlasBuildFinish(atlas *ImFontAtlas) {
	giLib.Call(_func_igImFontAtlasBuildFinish_, []interface{}{&atlas})
}

var _func_igImFontAtlasBuildRender8bppRectFromString_ = &c.FuncPrototype{Name: "igImFontAtlasBuildRender8bppRectFromString", OutType: c.Void, InTypes: _typs_PI32I32I32I32PU8U8}

// CIMGUI_API void igImFontAtlasBuildRender8bppRectFromString(ImFontAtlas* atlas,int x,int y,int w,int h,const char* in_str,char in_marker_char,uint8 in_marker_pixel_value);
func ImFontAtlasBuildRender8bppRectFromString(atlas *ImFontAtlas, x int32, y int32, w int32, h int32, in_str string, in_marker_char byte, in_marker_pixel_value byte) {
	_in_str := c.CStr(in_str)
	defer usf.Free(_in_str)
	giLib.Call(_func_igImFontAtlasBuildRender8bppRectFromString_, []interface{}{&atlas, &x, &y, &w, &h, &_in_str, &in_marker_char, &in_marker_pixel_value})
}

var _func_igImFontAtlasBuildRender32bppRectFromString_ = &c.FuncPrototype{Name: "igImFontAtlasBuildRender32bppRectFromString", OutType: c.Void, InTypes: _typs_PI32I32I32I32PU8U32}

// CIMGUI_API void igImFontAtlasBuildRender32bppRectFromString(ImFontAtlas* atlas,int x,int y,int w,int h,const char* in_str,char in_marker_char,uint in_marker_pixel_value);
func ImFontAtlasBuildRender32bppRectFromString(atlas *ImFontAtlas, x int32, y int32, w int32, h int32, in_str string, in_marker_char byte, in_marker_pixel_value uint32) {
	_in_str := c.CStr(in_str)
	defer usf.Free(_in_str)
	giLib.Call(_func_igImFontAtlasBuildRender32bppRectFromString_, []interface{}{&atlas, &x, &y, &w, &h, &_in_str, &in_marker_char, &in_marker_pixel_value})
}

// CIMGUI_API void igImFontAtlasBuildMultiplyCalcLookupTable(unsigned char out_table[256],float in_multiply_factor);
// CIMGUI_API void igImFontAtlasBuildMultiplyRectAlpha8(const unsigned char table[256],unsigned char* pixels,int x,int y,int w,int h,int stride);
