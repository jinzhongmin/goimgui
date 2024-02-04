package gi

import (
	"log"
	"os"
	"sync"
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
	InitLib(cimguilib, mod)
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

type GlfwQuickLauncher struct {
	win *glfw.Window
	ctx *ImGuiContext

	FPS      int
	FontPath string
	FontSize float32
	BgColor  []float32
}

func NewGlfwQuickLauncher(title string, width, height int) *GlfwQuickLauncher {
	win := glfw.CreateWindow(int32(width), int32(height), title, nil, nil)
	win.MakeContextCurrent()
	glfw.SwapInterval(1)

	ctx := CreateContext(nil)
	ImGui_ImplGlfw_InitForOpenGL(win.CPtr(), true)
	ImGui_ImplOpenGL3_Init(GLSLVer_3_2_Plus)

	launcher := &GlfwQuickLauncher{win: win, ctx: ctx}

	return launcher
}
func (launcher GlfwQuickLauncher) Run(loop func(io *QuickLaunchGlfwIO)) {

	win := launcher.win
	ctx := launcher.ctx
	defer win.DestroyWindow()
	defer DestroyContext(ctx)

	{ //for windows add chinese font
		font_path := `C:\Windows\Fonts\msyh.ttc`
		font_size := float32(20)
		if launcher.FontPath != "" {
			font_path = launcher.FontPath
		}
		if launcher.FontSize > 0 {
			font_size = launcher.FontSize
		}
		if _, err := os.Stat(font_path); err == nil {
			fontCfg := NewImFontConfig()
			io := GetIO()
			atlas := NewImFontAtlas()
			glyphRanges := atlas.GetGlyphRangesChineseFull()
			atlas.Destroy()
			io.Fonts.Clear()
			io.Fonts.AddFontFromFileTTF(font_path, font_size, fontCfg, glyphRanges)
		}
	}

	io := new(QuickLaunchGlfwIO)
	io.bgR, io.bgG, io.bgB, io.bgA = 0.45, 0.55, 0.60, 1.00

	if launcher.BgColor != nil && len(launcher.BgColor) < 4 {
		clr := launcher.BgColor
		io.bgR, io.bgG, io.bgB, io.bgA = clr[0], clr[1], clr[2], clr[3]
	}

	chanWait := make(chan bool, 1)
	{ //for limit fps
		fps := 30
		if launcher.FPS > 0 {
			fps = launcher.FPS
		}
		go func() {
			for {
				time.Sleep(time.Second / time.Duration(fps))
				chanWait <- true
				ok := <-chanWait
				if !ok {
					return
				}
			}
		}()
	}

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
		gl.Clear(gl.BufferBitColor)

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
}

//  var imagePtrHolder = make(map[unsafe.Pointer]interface{})
//
//	func NewImageGlTexureFromRGBA(rgba *image.RGBA) *ImageGlTexure {
//		tex := gl.GenTextures(1)[0]
//		gl.BindTexture(gl.GLTexTag_2D, tex)
//		ptr := unsafe.Pointer(&rgba.Pix[0])
//		imagePtrHolder[ptr] = rgba
//		GL_LINEAR := int32(0x2601)
//		GL_CLAMP_TO_EDGE := int32(0x812F)
//		gl.TexParameteri(gl.GLTexTag_2D, gl.GLTexPName_MIN_FILTER, GL_LINEAR)
//		gl.TexParameteri(gl.GLTexTag_2D, gl.GLTexPName_MAG_FILTER, GL_LINEAR)
//		gl.TexParameteri(gl.GLTexTag_2D, gl.GLTexPName_WRAP_S, GL_CLAMP_TO_EDGE)
//		gl.TexParameteri(gl.GLTexTag_2D, gl.GLTexPName_WRAP_T, GL_CLAMP_TO_EDGE)
//		gl.TexImage2D(gl.GLTexTag_2D, 0, gl.GLTexInterFmt_RGBA,
//			uint32(rgba.Rect.Dx()), uint32(rgba.Rect.Dy()), 0, gl.GLTexDataFmt_RGBA, gl.GLTexDataType_UNSIGNED_BYTE, ptr)
//			_tex := uint64(tex)
//			return &ImageGlTexure{tex: tex, texID: *(*ImTextureID)(unsafe.Pointer(&_tex)), ptr: ptr}
//		}

var imageGlTexureConfs = make(map[*ImageGlTexure]imageGlTexureConf)

type imageGlTexureConf struct {
	ptr           unsafe.Pointer
	width, height uint32
	border        int32
	interFmt      gl.TexFormat
	dataFmt       gl.TexFormat
	dataType      gl.TexDataType
}

var bindTextureLock = sync.Mutex{}

func NewImageGlTexure(ptr unsafe.Pointer, width, height uint32, border bool,
	interFmt gl.TexFormat, dataFmt gl.TexFormat, dataType gl.TexDataType) *ImageGlTexure {
	tex := gl.GenTextures(1)[0]
	bindTextureLock.Lock()
	gl.BindTexture(gl.TexTag2D, tex)

	_border := int32(0)
	if border {
		_border = 1
	}
	GL_LINEAR := int32(0x2601)
	GL_CLAMP_TO_EDGE := int32(0x812F)
	gl.TexParameteri(gl.TexTag2D, gl.TexParamMinFilter, GL_LINEAR)
	gl.TexParameteri(gl.TexTag2D, gl.TexParamMagFilter, GL_LINEAR)
	gl.TexParameteri(gl.TexTag2D, gl.TexParamWrapS, GL_CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TexTag2D, gl.TexParamWrapT, GL_CLAMP_TO_EDGE)
	gl.PixelStorei(gl.ParamUnpackAlignment, 1)
	gl.TexImage2D(gl.TexTag2D, 0, interFmt, width, height, _border, dataFmt, dataType, ptr)

	bindTextureLock.Unlock()
	_tex := uint64(tex)
	img := &ImageGlTexure{tex: tex, texID: *(*ImTextureID)(unsafe.Pointer(&_tex))}
	imageGlTexureConfs[img] = imageGlTexureConf{
		ptr: ptr, width: width, height: height, border: _border,
		interFmt: interFmt, dataFmt: dataFmt, dataType: dataType}
	return img
}
func (imTex *ImageGlTexure) Image(size ImVec2) {
	ImageDefault(imTex.texID, size)
}
func (imTex *ImageGlTexure) ImageWithBorder(size ImVec2) {
	Image(imTex.texID, size, *_vec2Zero(), *_vec2Ones(), *_vec4Ones(), *_vec4Ones())
}
func (imTex *ImageGlTexure) ImageCustom(size ImVec2, lt, rb ImVec2, tint, border ImVec4) {
	Image(imTex.texID, size, lt, rb, tint, border)
}
func (imTex *ImageGlTexure) ImageButton(label *c.Str, size ImVec2) bool {
	return ImageButtonDefault(label, imTex.texID, size)
}
func (imTex *ImageGlTexure) GetTexture() ImTextureID { return imTex.texID }
func (imTex *ImageGlTexure) DeleteTexture() {
	if imTex == nil {
		return
	}
	gl.DeleteTextures([]uint32{imTex.tex})
	delete(imageGlTexureConfs, imTex)
	imTex = nil
}
func (imTex *ImageGlTexure) UpdateTexureByLastConf(ptr unsafe.Pointer) {
	bindTextureLock.Lock()
	defer bindTextureLock.Unlock()

	gl.BindTexture(gl.TexTag2D, imTex.tex)
	conf := imageGlTexureConfs[imTex]
	if ptr != nil {
		conf.ptr = ptr
	}

	GL_LINEAR := int32(0x2601)
	gl.TexParameteri(gl.TexTag2D, gl.TexParamMinFilter, GL_LINEAR)
	gl.TexParameteri(gl.TexTag2D, gl.TexParamMagFilter, GL_LINEAR)
	gl.PixelStorei(gl.ParamUnpackAlignment, 1)
	gl.TexImage2D(gl.TexTag2D, 0, conf.interFmt, conf.width, conf.height,
		conf.border, conf.dataFmt, conf.dataType, conf.ptr)
}
func (imTex *ImageGlTexure) UpdateTexure(ptr unsafe.Pointer, width, height uint32, border bool,
	interFmt gl.TexFormat, dataFmt gl.TexFormat, dataType gl.TexDataType) {

	bindTextureLock.Lock()
	gl.BindTexture(gl.TexTag2D, imTex.tex)

	_border := int32(0)
	if border {
		_border = 1
	}
	GL_LINEAR := int32(0x2601)
	gl.TexParameteri(gl.TexTag2D, gl.TexParamMinFilter, GL_LINEAR)
	gl.TexParameteri(gl.TexTag2D, gl.TexParamMagFilter, GL_LINEAR)
	gl.PixelStorei(gl.ParamUnpackAlignment, 1)
	gl.TexImage2D(gl.TexTag2D, 0, interFmt, width, height, _border, dataFmt, dataType, ptr)
	bindTextureLock.Unlock()

	imageGlTexureConfs[imTex] = imageGlTexureConf{
		ptr: ptr, width: width, height: height, border: _border,
		interFmt: interFmt, dataFmt: dataFmt, dataType: dataType}
}

type Strs map[string]*c.Str

type StrsOpera byte

const (
	OpDel StrsOpera = 1
)

func NewStrs() (strs Strs, getter func(str string, op ...StrsOpera) *c.Str) {
	strs = make(map[string]*c.Str)
	getter = strs.Opera
	return
}
func (ms Strs) Opera(str string, op ...StrsOpera) *c.Str {
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
func (ms Strs) Clear() {
	for k, v := range ms {
		v.Free()
		delete(ms, k)
	}
}
func (ms Strs) Close() {
	ms.Clear()
	ms = nil
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
