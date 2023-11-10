package gi

import (
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
)

// IMGUI_IMPL_API bool     ImGui_ImplGlfw_InitForOpenGL(GLFWwindow* window, bool install_callbacks);
func ImGui_ImplGlfw_InitForOpenGL(glfwWinPtr unsafe.Pointer, install_callbacks bool) bool {
	b := c.CBool(install_callbacks)
	return giLib.Call(_func_ImGui_ImplGlfw_InitForOpenGL_, []interface{}{&glfwWinPtr, &b}).BoolFree()
}

// IMGUI_IMPL_API bool     ImGui_ImplGlfw_InitForVulkan(GLFWwindow* window, bool install_callbacks);
// IMGUI_IMPL_API bool     ImGui_ImplGlfw_InitForOther(GLFWwindow* window, bool install_callbacks);

// IMGUI_IMPL_API void     ImGui_ImplGlfw_Shutdown();
func ImGui_ImplGlfw_Shutdown() {
	giLib.Call(_func_ImGui_ImplGlfw_Shutdown_, nil)
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_NewFrame();
func ImGui_ImplGlfw_NewFrame() {
	giLib.Call(_func_ImGui_ImplGlfw_NewFrame_, nil)
}

// // GLFW callbacks install
// // - When calling Init with 'install_callbacks=true': ImGui_ImplGlfw_InstallCallbacks() is called. GLFW callbacks will be installed for you. They will chain-call user's previously installed callbacks, if any.
// // - When calling Init with 'install_callbacks=false': GLFW callbacks won't be installed. You will need to call individual function yourself from your own GLFW callbacks.

// IMGUI_IMPL_API void     ImGui_ImplGlfw_InstallCallbacks(GLFWwindow* window);
func ImGui_ImplGlfw_InstallCallbacks(glfwWinPtr unsafe.Pointer) {
	giLib.Call(_func_ImGui_ImplGlfw_InitForOpenGL_, nil)
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_RestoreCallbacks(GLFWwindow* window);
func ImGui_ImplGlfw_RestoreCallbacks(glfwWinPtr unsafe.Pointer) {
	giLib.Call(_func_ImGui_ImplGlfw_RestoreCallbacks_, nil)
}

// // GFLW callbacks options:
// // - Set 'chain_for_all_windows=true' to enable chaining callbacks for all windows (including secondary viewports created by backends or by user)
// IMGUI_IMPL_API void     ImGui_ImplGlfw_SetCallbacksChainForAllWindows(bool chain_for_all_windows);
func ImGui_ImplGlfw_SetCallbacksChainForAllWindows(chain_for_all_windows bool) {
	b := c.CBool(chain_for_all_windows)
	giLib.Call(_func_ImGui_ImplGlfw_SetCallbacksChainForAllWindows_, []interface{}{&b})
}

// // GLFW callbacks (individual callbacks to call yourself if you didn't install callbacks)

// IMGUI_IMPL_API void     ImGui_ImplGlfw_WindowFocusCallback(GLFWwindow* window, int focused);        // Since 1.84
func ImGui_ImplGlfw_WindowFocusCallback(glfwWinPtr unsafe.Pointer, focused bool) {
	b := c.CBool(focused)
	giLib.Call(_func_ImGui_ImplGlfw_WindowFocusCallback_, []interface{}{&glfwWinPtr, &b})
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_CursorEnterCallback(GLFWwindow* window, int entered);        // Since 1.84
func ImGui_ImplGlfw_CursorEnterCallback(glfwWinPtr unsafe.Pointer, entered bool) {
	b := c.CBool(entered)
	giLib.Call(_func_ImGui_ImplGlfw_CursorEnterCallback_, []interface{}{&glfwWinPtr, &b})
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_CursorPosCallback(GLFWwindow* window, double x, double y);   // Since 1.87
func ImGui_ImplGlfw_CursorPosCallback(glfwWinPtr unsafe.Pointer, x, y float64) {
	giLib.Call(_func_ImGui_ImplGlfw_CursorPosCallback_, []interface{}{&glfwWinPtr, &x, &y})
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_MouseButtonCallback(GLFWwindow* window, int button, int action, int mods);
func ImGui_ImplGlfw_MouseButtonCallback(glfwWinPtr unsafe.Pointer, button, action, mods int32) {
	giLib.Call(_func_ImGui_ImplGlfw_MouseButtonCallback_, []interface{}{&glfwWinPtr, &button, &action, &mods})
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_ScrollCallback(GLFWwindow* window, double xoffset, double yoffset);
func ImGui_ImplGlfw_ScrollCallback(glfwWinPtr unsafe.Pointer, xoffset, yoffset float64) {
	giLib.Call(_func_ImGui_ImplGlfw_ScrollCallback_, []interface{}{&glfwWinPtr, &xoffset, &yoffset})
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_KeyCallback(GLFWwindow* window, int key, int scancode, int action, int mods);
func ImGui_ImplGlfw_KeyCallback(glfwWinPtr unsafe.Pointer, key, scancode, action, mods int32) {
	giLib.Call(_func_ImGui_ImplGlfw_KeyCallback_, []interface{}{&glfwWinPtr, &key, &scancode, &action, &mods})
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_CharCallback(GLFWwindow* window, unsigned int c);
func ImGui_ImplGlfw_CharCallback(glfwWinPtr unsafe.Pointer, ch uint32) {
	giLib.Call(_func_ImGui_ImplGlfw_CharCallback_, []interface{}{&glfwWinPtr, &ch})
}

// IMGUI_IMPL_API void     ImGui_ImplGlfw_MonitorCallback(GLFWmonitor* monitor, int event);
func ImGui_ImplGlfw_MonitorCallback(glfwMonitor unsafe.Pointer, event int32) {
	giLib.Call(_func_ImGui_ImplGlfw_MonitorCallback_, []interface{}{&glfwMonitor, &event})
}

const (
	// GL ES 2.0 + GLSL 100
	GLSLVer_2_0_Plus string = "#version 100"
	//GL 3.2 + GLSL 150
	GLSLVer_3_2_Plus string = "#version 150"
	// GL 3.0 + GLSL 130
	GLSLVer_3_0_Plus string = "#version 130"
)

// IMGUI_IMPL_API bool     ImGui_ImplOpenGL3_Init(const char* glsl_version = nullptr);
func ImGui_ImplOpenGL3_Init(glsl_version string) bool {
	v := c.CStr(glsl_version)
	return giLib.Call(_func_ImGui_ImplOpenGL3_Init_, []interface{}{&v}).BoolFree()
}

// IMGUI_IMPL_API void     ImGui_ImplOpenGL3_Shutdown();
func ImGui_ImplOpenGL3_Shutdown() {
	giLib.Call(_func_ImGui_ImplOpenGL3_Shutdown_, nil)
}

// IMGUI_IMPL_API void     ImGui_ImplOpenGL3_NewFrame();
func ImGui_ImplOpenGL3_NewFrame() {
	giLib.Call(_func_ImGui_ImplOpenGL3_NewFrame_, nil)
}

// IMGUI_IMPL_API void     ImGui_ImplOpenGL3_RenderDrawData(ImDrawData* draw_data);
func ImGui_ImplOpenGL3_RenderDrawData(data *ImDrawData) {
	giLib.Call(_func_ImGui_ImplOpenGL3_RenderDrawData_, []interface{}{&data})
}

// // (Optional) Called by Init/NewFrame/Shutdown
// IMGUI_IMPL_API bool     ImGui_ImplOpenGL3_CreateFontsTexture();
func ImGui_ImplOpenGL3_CreateFontsTexture() bool {
	return giLib.Call(_func_ImGui_ImplOpenGL3_CreateFontsTexture_, nil).BoolFree()
}

// IMGUI_IMPL_API void     ImGui_ImplOpenGL3_DestroyFontsTexture();
func ImGui_ImplOpenGL3_DestroyFontsTexture() {
	giLib.Call(_func_ImGui_ImplOpenGL3_DestroyFontsTexture_, nil)
}

// IMGUI_IMPL_API bool     ImGui_ImplOpenGL3_CreateDeviceObjects();
func ImGui_ImplOpenGL3_CreateDeviceObjects() bool {
	return giLib.Call(_func_ImGui_ImplOpenGL3_CreateDeviceObjects_, nil).BoolFree()
}

// IMGUI_IMPL_API void     ImGui_ImplOpenGL3_DestroyDeviceObjects();
func ImGui_ImplOpenGL3_DestroyDeviceObjects() {
	giLib.Call(_func_ImGui_ImplOpenGL3_DestroyDeviceObjects_, nil)
}
