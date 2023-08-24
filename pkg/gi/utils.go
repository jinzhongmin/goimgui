package gi

import (
	"image"
	"log"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/goimgui/pkg/gl"
	"github.com/jinzhongmin/goimgui/pkg/glfw"
)

func InitLaunch(glfwlib, cimguilib string, mod c.LibMode) {
	glfw.InitLib(glfwlib, mod)
	glfw.SwapInterval(0)
	glfw.SetErrorCallback(func(code int32, des string) {
		log.Printf("error: %d, %s \n", code, des)
	})

	InitLib("cimgui.dll", mod)

	glfw.DefaultWindowHints()
	glfw.InitGL()
}

func QuickLaunch(win *glfw.Window, loopFun func(w, h int32) bool) {
	defer glfw.Terminate()
	defer win.DestroyWindow()

	ctx := CreateContext(nil)
	defer DestroyContext(ctx)

	ImGui_ImplGlfw_InitForOpenGL(win.CPtr(), true)
	ImGui_ImplOpenGL3_Init(GLSLVer_3_2_Plus)

	//for limit fps
	chanWait := make(chan bool, 1)
	go func() {
		for {
			time.Sleep(time.Second / 30)
			chanWait <- true
		}
	}()
	open := true

	//load chinese font
	ms_cn_font_path := `C:\Windows\Fonts\msyh.ttc`
	_, err := os.Stat(ms_cn_font_path)
	var font *ImFont
	if err == nil {
		fontCfg := NewImFontConfig()
		io := GetIO()
		io.ConfigFlags |= ImGuiConfigFlags_DockingEnable
		io.Fonts.AddFontFromFileTTF(ms_cn_font_path, 20, fontCfg, NewImFontAtlas().GetGlyphRangesChineseSimplifiedCommon())
	}

	var zx, zy, w, h int32
	zx, zy = 0, 0
	var r, g, b, a float32 = 0.45, 0.55, 0.60, 1.00
	runtime.GC()
	for !(!open || win.WindowShouldClose()) {

		glfw.PollEvents()
		ImGui_ImplOpenGL3_NewFrame()
		ImGui_ImplGlfw_NewFrame()
		NewFrame()

		win.GetFramebufferSizeRef(&w, &h)

		if font != nil {
			PushFont(font)
			open = loopFun(w, h)
			PopFont()
		} else {
			open = loopFun(w, h)
		}

		Render()
		gl.Viewport(zx, zy, w, h)
		gl.ClearColor(r, g, b, a)
		gl.Clear(gl.GLbitfield_COLOR_BUFFER_BIT)

		ImGui_ImplOpenGL3_RenderDrawData(GetDrawData())
		win.SwapBuffers()
		<-chanWait
	}
	ImGui_ImplOpenGL3_Shutdown()
	ImGui_ImplGlfw_Shutdown()
}

type ImageGlTexure struct {
	tex   uint32
	texID ImTextureID
	img   *image.RGBA
}

func NewImageGlTexureFromRGBA(rgba *image.RGBA) *ImageGlTexure {
	tex := gl.GenTextures(1)[0]
	gl.BindTexture(gl.GLTexTag_2D, tex)

	GL_LINEAR := int32(0x2601)
	GL_CLAMP_TO_EDGE := int32(0x812F)
	gl.TexParameteri(gl.GLTexTag_2D, gl.GLTexPName_MIN_FILTER, GL_LINEAR)
	gl.TexParameteri(gl.GLTexTag_2D, gl.GLTexPName_MAG_FILTER, GL_LINEAR)
	gl.TexParameteri(gl.GLTexTag_2D, gl.GLTexPName_WRAP_S, GL_CLAMP_TO_EDGE)
	gl.TexParameteri(gl.GLTexTag_2D, gl.GLTexPName_WRAP_T, GL_CLAMP_TO_EDGE)
	gl.TexImage2D(gl.GLTexTag_2D, 0, gl.GLTexInterFmt_RGBA,
		uint32(rgba.Rect.Dx()), uint32(rgba.Rect.Dy()), 0, gl.GLTexDataFmt_RGBA, gl.GLTexDataType_UNSIGNED_BYTE, unsafe.Pointer(&rgba.Pix[0]))

	_tex := uint64(tex)
	return &ImageGlTexure{tex: tex, texID: *(*ImTextureID)(unsafe.Pointer(&_tex)), img: rgba}
}
func (imTex *ImageGlTexure) Image(w, h float32) {
	Image(imTex.texID, ImVec2{X: w, Y: h}, ImVec2{0, 0}, ImVec2{1, 1},
		ImVec4{1, 1, 1, 1}, ImVec4{0, 0, 0, 0})
}
func (imTex *ImageGlTexure) ImageButton(label string, w, h float32) bool {
	return ImageButton(label, imTex.texID, ImVec2{X: w, Y: h}, ImVec2{0, 0}, ImVec2{1, 1},
		ImVec4{0, 0, 0, 0}, ImVec4{1, 1, 1, 1})
}
func (imTex *ImageGlTexure) GetTexture() ImTextureID { return imTex.texID }
func (imTex *ImageGlTexure) DeleteTexture()          { gl.DeleteTextures([]uint32{imTex.tex}) }
