package gl

import (
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/usf"
)

type (
	Bitfield        int32
	TexTag          int32
	TexParam        int32
	TexLevel        int32
	TexFormat       int32
	TexDataType     int32
	PixelStoreParam int32
)

const (
	BufferBitColor   Bitfield = 0x00004000
	BufferBitDepth   Bitfield = 0x00000100
	BufferBitStencil Bitfield = 0x00000400

	TexTag1D                 TexTag = 0x0DE0
	TexTag2D                 TexTag = 0x0DE1
	TexTag3D                 TexTag = 0x806F
	TexTag1DArray            TexTag = 0x8C18
	TexTag2DArray            TexTag = 0x8C1A
	TexTagRectangle          TexTag = 0x84F5
	TexTagCubeMap            TexTag = 0x8513
	TexTagBuffer             TexTag = 0x8C2A
	TexTag2DMultisample      TexTag = 0x9100
	TexTag2DMultisampleArray TexTag = 0x9102

	TexParamBaseLevel   TexParam = 0x813C
	TexParamCompareFunc TexParam = 0x884D
	TexParamCompareMode TexParam = 0x884C
	TexParamLodBias     TexParam = 0x8501
	TexParamMinFilter   TexParam = 0x2801
	TexParamMagFilter   TexParam = 0x2800
	TexParamMinLod      TexParam = 0x813A
	TexParamMaxLod      TexParam = 0x813B
	TexParamMaxLevel    TexParam = 0x813D
	TexParamSwizzleR    TexParam = 0x8E42
	TexParamSwizzleG    TexParam = 0x8E43
	TexParamSwizzleB    TexParam = 0x8E44
	TexParamSwizzleA    TexParam = 0x8E45
	TexParamWrapS       TexParam = 0x2802
	TexParamWrapT       TexParam = 0x2803
	TexParamWrapR       TexParam = 0x8072

	TexLevelTextureRectangle      TexLevel = 0x84F5
	TexLevelProxyTextureRectangle TexLevel = 0x84F7

	TexFormatAlpha               TexFormat = 0x1906
	TexFormatAlpha4              TexFormat = 0x803B
	TexFormatAlpha8              TexFormat = 0x803C
	TexFormatAlpha12             TexFormat = 0x803D
	TexFormatAlpha16             TexFormat = 0x803E
	TexFormatLuminance           TexFormat = 0x1909
	TexFormatLuminance4          TexFormat = 0x803F
	TexFormatLuminance8          TexFormat = 0x8040
	TexFormatLuminance12         TexFormat = 0x8041
	TexFormatLuminance16         TexFormat = 0x8042
	TexFormatLuminanceAlpha      TexFormat = 0x190A
	TexFormatLuminance4_Alpha4   TexFormat = 0x8043
	TexFormatLuminance6_Alpha2   TexFormat = 0x8044
	TexFormatLuminance8_Alpha8   TexFormat = 0x8045
	TexFormatLuminance12_Alpha4  TexFormat = 0x8046
	TexFormatLuminance12_Alpha12 TexFormat = 0x8047
	TexFormatLuminance16_Alpha16 TexFormat = 0x8048
	TexFormatIntensity           TexFormat = 0x8049
	TexFormatIntensity4          TexFormat = 0x804A
	TexFormatIntensity8          TexFormat = 0x804B
	TexFormatIntensity12         TexFormat = 0x804C
	TexFormatIntensity16         TexFormat = 0x804D
	TexFormatR3_G3_B2            TexFormat = 0x2A10
	TexFormatRGB                 TexFormat = 0x1907
	TexFormatRGB4                TexFormat = 0x804F
	TexFormatRGB5                TexFormat = 0x8050
	TexFormatRGB8                TexFormat = 0x8051
	TexFormatRGB10               TexFormat = 0x8052
	TexFormatRGB12               TexFormat = 0x8053
	TexFormatRGB16               TexFormat = 0x8054
	TexFormatRGBA                TexFormat = 0x1908
	TexFormatRGBA2               TexFormat = 0x8055
	TexFormatRGBA4               TexFormat = 0x8056
	TexFormatRGB5_A1             TexFormat = 0x8057
	TexFormatRGBA8               TexFormat = 0x8058
	TexFormatRGB10_A2            TexFormat = 0x8059
	TexFormatRGBA12              TexFormat = 0x805A
	TexFormatRGBA16              TexFormat = 0x805B
	TexFormatRED                 TexFormat = 0x1903
	TexFormatRG                  TexFormat = 0x8227
	TexFormatBGR                 TexFormat = 0x80E0
	TexFormatBGRA                TexFormat = 0x80E1
	TexFormatREDInteger          TexFormat = 0x8D94
	TexFormatRGInteger           TexFormat = 0x8228
	TexFormatRGBInteger          TexFormat = 0x8D98
	TexFormatBGRInteger          TexFormat = 0x8D9A
	TexFormatRGBAInteger         TexFormat = 0x8D99
	TexFormatBGRAInteger         TexFormat = 0x8D9B
	TexFormatStencilIndex        TexFormat = 0x1901
	TexFormatDepthComponent      TexFormat = 0x1902
	TexFormatDepthStencil        TexFormat = 0x84F9

	TexDataTypeUByte                   TexDataType = 0x1401
	TexDataTypeByte                    TexDataType = 0x1400
	TexDataTypeUShort                  TexDataType = 0x1403
	TexDataTypeShort                   TexDataType = 0x1402
	TexDataTypeUInt                    TexDataType = 0x1405
	TexDataTypeInt                     TexDataType = 0x1404
	TexDataTypeHalfFloat               TexDataType = 0x140B
	TexDataTypeFloat                   TexDataType = 0x1406
	TexDataTypeUByte332                TexDataType = 0x8032
	TexDataTypeUByte233Rev             TexDataType = 0x8362
	TexDataTypeUShort565               TexDataType = 0x8363
	TexDataTypeUShort565Rev            TexDataType = 0x8364
	TexDataTypeUShort4444              TexDataType = 0x8033
	TexDataTypeUShort4444Rev           TexDataType = 0x8365
	TexDataTypeUShort5551              TexDataType = 0x8034
	TexDataTypeUShort1555Rev           TexDataType = 0x8366
	TexDataTypeUInt8888                TexDataType = 0x8035
	TexDataTypeUInt8888Rev             TexDataType = 0x8367
	TexDataTypeUInt10_10_10_10_2       TexDataType = 0x8036
	TexDataTypeUInt_INT_2_10_10_10_Rev TexDataType = 0x8368

	ParamsPackSwapBytes    PixelStoreParam = 0x0D00
	ParamPackLsbFirst      PixelStoreParam = 0x0D01
	ParamPackRowLength     PixelStoreParam = 0x0D02
	ParamPackImageHeight   PixelStoreParam = 0x806C
	ParamPackSkipPixels    PixelStoreParam = 0x0D04
	ParamPackSkipRows      PixelStoreParam = 0x0D03
	ParamPackSkipImages    PixelStoreParam = 0x806B
	ParamPackAlignment     PixelStoreParam = 0x0D05
	ParamUnpackSwapBytes   PixelStoreParam = 0x0CF0
	ParamUnpackLsbFirst    PixelStoreParam = 0x0CF1
	ParamUnpackRowLength   PixelStoreParam = 0x0CF2
	ParamUnpackImageHeight PixelStoreParam = 0x806E
	ParamUnpackSkipPixels  PixelStoreParam = 0x0CF4
	ParamUnpackSkipRows    PixelStoreParam = 0x0CF3
	ParamUnpackSkipImages  PixelStoreParam = 0x806D
	ParamUnpackAlignment   PixelStoreParam = 0x0CF5
)

func InitProc(lib *c.Lib) {
	_GLLib = lib
}

var _GLLib *c.Lib

var _func_glfwGetProcAddress_ = &c.FuncPrototype{Name: "glfwGetProcAddress", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

func _addr(fp *c.FuncPrototype) *c.FuncPrototype {
	if fp.Ptr != nil {
		return fp
	}

	n := c.CStr(fp.Name)
	defer usf.Free(n)

	fp.Ptr = _GLLib.Call(_func_glfwGetProcAddress_, []interface{}{&n}).PtrFree()
	return fp
}

var _func_glViewport_ = &c.FuncPrototype{Name: "glViewport", OutType: c.Void, InTypes: []c.Type{c.I32, c.I32, c.I32, c.I32}}

func Viewport(x, y, width, height int32) {
	_GLLib.Call(_addr(_func_glViewport_), []interface{}{&x, &y, &width, &height})
}

var _func_glClear_ = &c.FuncPrototype{Name: "glClear", OutType: c.Void, InTypes: []c.Type{c.I32}}

func Clear(mask Bitfield) {
	_GLLib.Call(_addr(_func_glClear_), []interface{}{&mask})
}

var _func_glClearColor_ = &c.FuncPrototype{Name: "glClearColor", OutType: c.Void, InTypes: []c.Type{c.F32, c.F32, c.F32, c.F32}}

func ClearColor(r, g, b, a float32) {
	_GLLib.Call(_addr(_func_glClearColor_), []interface{}{&r, &g, &b, &a})
}

var _func_glGenTextures_ = &c.FuncPrototype{Name: "glGenTextures", OutType: c.Void, InTypes: []c.Type{c.I32, c.Pointer}}

func GenTextures(num int32) (textIds []uint32) {
	textIds = make([]uint32, num)
	ptr := &textIds[0]
	_GLLib.Call(_addr(_func_glGenTextures_), []interface{}{&num, &ptr})
	return
}

var _func_glDeleteTextures_ = &c.FuncPrototype{Name: "glDeleteTextures", OutType: c.Void, InTypes: []c.Type{c.I32, c.Pointer}}

func DeleteTextures(textIds []uint32) {
	num := int32(len(textIds))
	ptr := &textIds[0]
	_GLLib.Call(_addr(_func_glDeleteTextures_), []interface{}{&num, &ptr})
}

var _func_glBindTexture_ = &c.FuncPrototype{Name: "glBindTexture", OutType: c.Void, InTypes: []c.Type{c.I32, c.U32}}

func BindTexture(target TexTag, texId uint32) {
	_GLLib.Call(_addr(_func_glBindTexture_), []interface{}{&target, &texId})
}

var _func_glTexParameteri_ = &c.FuncPrototype{Name: "glTexParameteri", OutType: c.Void, InTypes: []c.Type{c.I32, c.I32, c.I32}}

// void glTexParameteri(GLenum target, GLenum pname, GLint param);
func TexParameteri(tag TexTag, pname TexParam, param int32) {
	_GLLib.Call(_addr(_func_glTexParameteri_), []interface{}{&tag, &pname, &param})
}

var _func_glTexImage2D_ = &c.FuncPrototype{Name: "glTexImage2D", OutType: c.Void, InTypes: []c.Type{c.I32, c.I32, c.I32, c.U32, c.U32, c.I32, c.I32, c.I32, c.Pointer}}

// void glTexImage2D(  GLenum target, GLint level, GLint internalFormat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid * data);
func TexImage2D(target TexTag, level TexLevel, internalFormat TexFormat,
	width, height uint32, border int32, dataFmt TexFormat, dataTyp TexDataType, data unsafe.Pointer) {
	_GLLib.Call(_addr(_func_glTexImage2D_),
		[]interface{}{&target, &level, &internalFormat, &width, &height, &border, &dataFmt, &dataTyp, &data})
}

var _func_glTexSubImage2D_ = &c.FuncPrototype{Name: "glTexSubImage2D", OutType: c.Void, InTypes: []c.Type{c.I32, c.I32, c.I32, c.I32, c.U32, c.U32, c.I32, c.I32, c.Pointer}}

func TexSubImage2D(target TexTag, level TexLevel, xOffset, yOffset int32,
	width, height uint32, dataFmt TexFormat, dataTyp TexDataType, data unsafe.Pointer) {
	_GLLib.Call(_addr(_func_glTexSubImage2D_),
		[]interface{}{&target, &level, &xOffset, &yOffset, &width, &height, &dataFmt, &dataTyp, &data})
}

var _func_glGenerateMipmap_ = &c.FuncPrototype{Name: "glGenerateMipmap", OutType: c.Void, InTypes: []c.Type{c.I32}}

func GenerateMipmap(target TexTag) {
	_GLLib.Call(_addr(_func_glGenerateMipmap_), []interface{}{&target})
}

var _func_glPixelStoref_ = &c.FuncPrototype{Name: "glPixelStoref", OutType: c.Void, InTypes: []c.Type{c.I32, c.F32}}

func PixelStoref(pname PixelStoreParam, param float32) {
	_GLLib.Call(_addr(_func_glPixelStoref_), []interface{}{&pname, &param})
}

var _func_glPixelStorei_ = &c.FuncPrototype{Name: "glPixelStorei", OutType: c.Void, InTypes: []c.Type{c.I32, c.I32}}

func PixelStorei(pname PixelStoreParam, param int32) {
	_GLLib.Call(_addr(_func_glPixelStorei_), []interface{}{&pname, &param})
}
