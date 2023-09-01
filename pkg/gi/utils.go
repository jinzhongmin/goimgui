package gi

import (
	"image"
	"log"
	"os"
	"time"
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/goimgui/pkg/gl"
	"github.com/jinzhongmin/goimgui/pkg/glfw"
)

func QuickInitLibraryGlfw(glfwlib, cimguilib string, mod c.LibMode) {
	glfw.InitLib(glfwlib, mod)
	glfw.DefaultWindowHints()
	glfw.InitGL()
	glfw.SetErrorCallback(func(code int32, des string) { log.Printf("error: %d, %s \n", code, des) })
	InitLib("cimgui.dll", mod)
}

type QuickLaunchGlfwIO struct {
	currWinW           int32
	currWinH           int32
	bgR, bgG, bgB, bgA float32

	shouldClose bool
}

func (io QuickLaunchGlfwIO) GetWinWH() (w, h int32) {
	w = io.currWinW
	h = io.currWinH
	return
}

func (io QuickLaunchGlfwIO) GetBgColor() (r, g, b, a float32) {
	r, g, b, a = io.bgR, io.bgG, io.bgB, io.bgA
	return
}

// r,g,b,a range 0-1.0
func (io *QuickLaunchGlfwIO) SetBgColor(r, g, b, a float32) {
	io.bgR, io.bgG, io.bgB, io.bgA = r, g, b, a
}

func (io *QuickLaunchGlfwIO) Close() {
	io.shouldClose = true
}

func QuickLaunchGlfw(title string, width, height int, loop func(io *QuickLaunchGlfwIO)) {
	win := glfw.CreateWindow(int32(width), int32(height), title, nil, nil)
	win.MakeContextCurrent()
	defer win.DestroyWindow()

	ctx := CreateContext(nil)
	defer DestroyContext(ctx)

	ImGui_ImplGlfw_InitForOpenGL(win.CPtr(), true)
	ImGui_ImplOpenGL3_Init(GLSLVer_3_2_Plus)

	chanWait := make(chan bool, 1)
	{ //for limit fps
		go func() {
			for {
				time.Sleep(time.Second / 30)
				chanWait <- true
				ok := <-chanWait
				if !ok {
					return
				}
			}
		}()
	}

	{ //for windows add chhinese font
		ms_cn_font_path := `C:\Windows\Fonts\msyh.ttc`
		if _, err := os.Stat(ms_cn_font_path); err == nil {
			fontCfg := NewImFontConfig()
			io := GetIO()
			io.Fonts.AddFontFromFileTTF(ms_cn_font_path, 20, fontCfg, NewImFontAtlas().GetGlyphRangesChineseSimplifiedCommon())
		}
	}

	io := new(QuickLaunchGlfwIO)
	io.bgR, io.bgG, io.bgB, io.bgA = 0.45, 0.55, 0.60, 1.00
	for !win.WindowShouldClose() {

		glfw.PollEvents()
		ImGui_ImplOpenGL3_NewFrame()
		ImGui_ImplGlfw_NewFrame()
		NewFrame()

		win.GetFramebufferSizeRef(&io.currWinW, &io.currWinH)

		loop(io)

		Render()
		gl.Viewport(0, 0, io.currWinW, io.currWinH)
		gl.ClearColor(io.bgR, io.bgG, io.bgB, io.bgA)
		gl.Clear(gl.GLbitfield_COLOR_BUFFER_BIT)

		ImGui_ImplOpenGL3_RenderDrawData(GetDrawData())
		win.SwapBuffers()
		if io.shouldClose {
			<-chanWait
			chanWait <- true
			break
		}
		<-chanWait
		chanWait <- true
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
func (imTex *ImageGlTexure) Image(size ImVec2) {
	ImageDefault(imTex.texID, size)
}
func (imTex *ImageGlTexure) ImageButton(label *c.Str, size ImVec2) bool {
	return ImageButtonDefault(label, imTex.texID, size)
}
func (imTex *ImageGlTexure) GetTexture() ImTextureID { return imTex.texID }
func (imTex *ImageGlTexure) DeleteTexture()          { gl.DeleteTextures([]uint32{imTex.tex}) }

type MyStrs map[string]*c.Str

type MyStoreOpera byte

const (
	OpDel MyStoreOpera = 1
)

func NewMyStrs() (strs MyStrs, getter func(str string, op ...MyStoreOpera) *c.Str) {
	strs = make(map[string]*c.Str)
	getter = strs.Opera
	return
}
func (ms MyStrs) Opera(str string, op ...MyStoreOpera) *c.Str {
	if op != nil {
		switch op[0] {
		case OpDel:
			delete(ms, str)
		}
		return nil
	}
	if v, ok := ms[str]; ok {
		return v
	}
	ms[str] = c.NewStr(str)
	return ms[str]
}

// type MyVec2s map[string]*ImVec2
// func NewMyVec2s() (vec MyVec2s, getter func(name string, op ...MyStoreOpera) *ImVec2) {
// 	vec = make(map[string]*ImVec2)
// 	getter = vec.Opera
// 	return
// }
// func (vec MyVec2s) Opera(name string, op ...MyStoreOpera) *ImVec2 {
// 	if op != nil {
// 		switch op[0] {
// 		case OpDel:
// 			delete(vec, name)
// 		}
// 		return nil
// 	}
// 	if v, ok := vec[name]; ok {
// 		return v
// 	}
// 	vec[name] = NewImVec2()
// 	return vec[name]
// }
