package gl

import (
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/usf"
)

type (
	GLbitfield    int32
	GLTexTag      int32
	GLTexPName    int32
	GLTexLevel    int32
	GLTexInterFmt int32
	GLTexDataFmt  int32
	GLTexDataType int32
)

const (
	GLbitfield_COLOR_BUFFER_BIT   GLbitfield = 0x00004000
	GLbitfield_DEPTH_BUFFER_BIT   GLbitfield = 0x00000100
	GLbitfield_STENCIL_BUFFER_BIT GLbitfield = 0x00000400

	GLTexTag_1D                   GLTexTag = 0x0DE0
	GLTexTag_2D                   GLTexTag = 0x0DE1
	GLTexTag_3D                   GLTexTag = 0x806F
	GLTexTag_1D_ARRAY             GLTexTag = 0x8C18
	GLTexTag_2D_ARRAY             GLTexTag = 0x8C1A
	GLTexTag_RECTANGLE            GLTexTag = 0x84F5
	GLTexTag_CUBE_MAP             GLTexTag = 0x8513
	GLTexTag_BUFFER               GLTexTag = 0x8C2A
	GLTexTag_2D_MULTISAMPLE       GLTexTag = 0x9100
	GLTexTag_2D_MULTISAMPLE_ARRAY GLTexTag = 0x9102

	GLTexPName_BASE_LEVEL   GLTexPName = 0x813C
	GLTexPName_COMPARE_FUNC GLTexPName = 0x884D
	GLTexPName_COMPARE_MODE GLTexPName = 0x884C
	GLTexPName_LOD_BIAS     GLTexPName = 0x8501
	GLTexPName_MIN_FILTER   GLTexPName = 0x2801
	GLTexPName_MAG_FILTER   GLTexPName = 0x2800
	GLTexPName_MIN_LOD      GLTexPName = 0x813A
	GLTexPName_MAX_LOD      GLTexPName = 0x813B
	GLTexPName_MAX_LEVEL    GLTexPName = 0x813D
	GLTexPName_SWIZZLE_R    GLTexPName = 0x8E42
	GLTexPName_SWIZZLE_G    GLTexPName = 0x8E43
	GLTexPName_SWIZZLE_B    GLTexPName = 0x8E44
	GLTexPName_SWIZZLE_A    GLTexPName = 0x8E45
	GLTexPName_WRAP_S       GLTexPName = 0x2802
	GLTexPName_WRAP_T       GLTexPName = 0x2803
	GLTexPName_WRAP_R       GLTexPName = 0x8072

	GLTexLevel_TEXTURE_RECTANGLE       GLTexLevel = 0x84F5
	GLTexLevel_PROXY_TEXTURE_RECTANGLE GLTexLevel = 0x84F7

	GLTexInterFmt_ALPHA               GLTexInterFmt = 0x1906
	GLTexInterFmt_ALPHA4              GLTexInterFmt = 0x803B
	GLTexInterFmt_ALPHA8              GLTexInterFmt = 0x803C
	GLTexInterFmt_ALPHA12             GLTexInterFmt = 0x803D
	GLTexInterFmt_ALPHA16             GLTexInterFmt = 0x803E
	GLTexInterFmt_LUMINANCE           GLTexInterFmt = 0x1909
	GLTexInterFmt_LUMINANCE4          GLTexInterFmt = 0x803F
	GLTexInterFmt_LUMINANCE8          GLTexInterFmt = 0x8040
	GLTexInterFmt_LUMINANCE12         GLTexInterFmt = 0x8041
	GLTexInterFmt_LUMINANCE16         GLTexInterFmt = 0x8042
	GLTexInterFmt_LUMINANCE_ALPHA     GLTexInterFmt = 0x190A
	GLTexInterFmt_LUMINANCE4_ALPHA4   GLTexInterFmt = 0x8043
	GLTexInterFmt_LUMINANCE6_ALPHA2   GLTexInterFmt = 0x8044
	GLTexInterFmt_LUMINANCE8_ALPHA8   GLTexInterFmt = 0x8045
	GLTexInterFmt_LUMINANCE12_ALPHA4  GLTexInterFmt = 0x8046
	GLTexInterFmt_LUMINANCE12_ALPHA12 GLTexInterFmt = 0x8047
	GLTexInterFmt_LUMINANCE16_ALPHA16 GLTexInterFmt = 0x8048
	GLTexInterFmt_INTENSITY           GLTexInterFmt = 0x8049
	GLTexInterFmt_INTENSITY4          GLTexInterFmt = 0x804A
	GLTexInterFmt_INTENSITY8          GLTexInterFmt = 0x804B
	GLTexInterFmt_INTENSITY12         GLTexInterFmt = 0x804C
	GLTexInterFmt_INTENSITY16         GLTexInterFmt = 0x804D
	GLTexInterFmt_R3_G3_B2            GLTexInterFmt = 0x2A10
	GLTexInterFmt_RGB                 GLTexInterFmt = 0x1907
	GLTexInterFmt_RGB4                GLTexInterFmt = 0x804F
	GLTexInterFmt_RGB5                GLTexInterFmt = 0x8050
	GLTexInterFmt_RGB8                GLTexInterFmt = 0x8051
	GLTexInterFmt_RGB10               GLTexInterFmt = 0x8052
	GLTexInterFmt_RGB12               GLTexInterFmt = 0x8053
	GLTexInterFmt_RGB16               GLTexInterFmt = 0x8054
	GLTexInterFmt_RGBA                GLTexInterFmt = 0x1908
	GLTexInterFmt_RGBA2               GLTexInterFmt = 0x8055
	GLTexInterFmt_RGBA4               GLTexInterFmt = 0x8056
	GLTexInterFmt_RGB5_A1             GLTexInterFmt = 0x8057
	GLTexInterFmt_RGBA8               GLTexInterFmt = 0x8058
	GLTexInterFmt_RGB10_A2            GLTexInterFmt = 0x8059
	GLTexInterFmt_RGBA12              GLTexInterFmt = 0x805A
	GLTexInterFmt_RGBA16              GLTexInterFmt = 0x805B

	GLTexDataFmt_RED             GLTexDataFmt = 0x1903
	GLTexDataFmt_RG              GLTexDataFmt = 0x8227
	GLTexDataFmt_RGB             GLTexDataFmt = 0x1907
	GLTexDataFmt_BGR             GLTexDataFmt = 0x80E0
	GLTexDataFmt_RGBA            GLTexDataFmt = 0x1908
	GLTexDataFmt_BGRA            GLTexDataFmt = 0x80E1
	GLTexDataFmt_RED_INTEGER     GLTexDataFmt = 0x8D94
	GLTexDataFmt_RG_INTEGER      GLTexDataFmt = 0x8228
	GLTexDataFmt_RGB_INTEGER     GLTexDataFmt = 0x8D98
	GLTexDataFmt_BGR_INTEGER     GLTexDataFmt = 0x8D9A
	GLTexDataFmt_RGBA_INTEGER    GLTexDataFmt = 0x8D99
	GLTexDataFmt_BGRA_INTEGER    GLTexDataFmt = 0x8D9B
	GLTexDataFmt_STENCIL_INDEX   GLTexDataFmt = 0x1901
	GLTexDataFmt_DEPTH_COMPONENT GLTexDataFmt = 0x1902
	GLTexDataFmt_DEPTH_STENCIL   GLTexDataFmt = 0x84F9

	GLTexDataType_UNSIGNED_BYTE               GLTexDataType = 0x1401
	GLTexDataType_BYTE                        GLTexDataType = 0x1400
	GLTexDataType_UNSIGNED_SHORT              GLTexDataType = 0x1403
	GLTexDataType_SHORT                       GLTexDataType = 0x1402
	GLTexDataType_UNSIGNED_INT                GLTexDataType = 0x1405
	GLTexDataType_INT                         GLTexDataType = 0x1404
	GLTexDataType_HALF_FLOAT                  GLTexDataType = 0x140B
	GLTexDataType_FLOAT                       GLTexDataType = 0x1406
	GLTexDataType_UNSIGNED_BYTE_3_3_2         GLTexDataType = 0x8032
	GLTexDataType_UNSIGNED_BYTE_2_3_3_REV     GLTexDataType = 0x8362
	GLTexDataType_UNSIGNED_SHORT_5_6_5        GLTexDataType = 0x8363
	GLTexDataType_UNSIGNED_SHORT_5_6_5_REV    GLTexDataType = 0x8364
	GLTexDataType_UNSIGNED_SHORT_4_4_4_4      GLTexDataType = 0x8033
	GLTexDataType_UNSIGNED_SHORT_4_4_4_4_REV  GLTexDataType = 0x8365
	GLTexDataType_UNSIGNED_SHORT_5_5_5_1      GLTexDataType = 0x8034
	GLTexDataType_UNSIGNED_SHORT_1_5_5_5_REV  GLTexDataType = 0x8366
	GLTexDataType_UNSIGNED_INT_8_8_8_8        GLTexDataType = 0x8035
	GLTexDataType_UNSIGNED_INT_8_8_8_8_REV    GLTexDataType = 0x8367
	GLTexDataType_UNSIGNED_INT_10_10_10_2     GLTexDataType = 0x8036
	GLTexDataType_UNSIGNED_INT_2_10_10_10_REV GLTexDataType = 0x8368

	_fn_glfwGetProcAddress_ = "glfwGetProcAddress"
	_fn_glViewport_         = "glViewport"
	_fn_glClear_            = "glClear"
	_fn_glClearColor_       = "glClearColor"
	_fn_glGenTextures_      = "glGenTextures"
	_fn_glDeleteTextures_   = "glDeleteTextures"
	_fn_glBindTexture_      = "glBindTexture"
	_fn_glTexParameteri_    = "glTexParameteri"
	_fn_glTexImage2D_       = "glTexImage2D"
	_fn_glGenerateMipmap_   = "glGenerateMipmap"
)

var (
	_typs_P                         = []c.Type{c.Pointer}
	_typs_I32                       = []c.Type{c.I32}
	_typs_I32P                      = []c.Type{c.I32, c.Pointer}
	_typs_I32U32                    = []c.Type{c.I32, c.U32}
	_typs_I32I32I32                 = []c.Type{c.I32, c.I32, c.I32}
	_typs_I32I32I32I32              = []c.Type{c.I32, c.I32, c.I32, c.I32}
	_typs_F32F32F32F32              = []c.Type{c.F32, c.F32, c.F32, c.F32}
	_typs_I32I32I32U32U32I32I32I32P = []c.Type{c.I32, c.I32, c.I32, c.U32, c.U32, c.I32, c.I32, c.I32, c.Pointer}
)

func InitProc(lib *c.Lib) {
	_GLLib = lib
}

var (
	_GLLib *c.Lib

	_func_glfwGetProcAddress_ = &c.FuncPrototype{Name: _fn_glfwGetProcAddress_, OutType: c.Pointer, InTypes: _typs_P}
	_func_glViewport_         = &c.FuncPrototype{Name: _fn_glViewport_, OutType: c.Void, InTypes: _typs_I32I32I32I32}
	_func_glClear_            = &c.FuncPrototype{Name: _fn_glClear_, OutType: c.Void, InTypes: _typs_I32}
	_func_glClearColor_       = &c.FuncPrototype{Name: _fn_glClearColor_, OutType: c.Void, InTypes: _typs_F32F32F32F32}
	_func_glGenTextures_      = &c.FuncPrototype{Name: _fn_glGenTextures_, OutType: c.Void, InTypes: _typs_I32P}
	_func_glDeleteTextures_   = &c.FuncPrototype{Name: _fn_glDeleteTextures_, OutType: c.Void, InTypes: _typs_I32P}
	_func_glBindTexture_      = &c.FuncPrototype{Name: _fn_glBindTexture_, OutType: c.Void, InTypes: _typs_I32U32}
	_func_glTexParameteri_    = &c.FuncPrototype{Name: _fn_glTexParameteri_, OutType: c.Void, InTypes: _typs_I32I32I32}
	_func_glTexImage2D_       = &c.FuncPrototype{Name: _fn_glTexImage2D_, OutType: c.Void, InTypes: _typs_I32I32I32U32U32I32I32I32P}
	_func_glGenerateMipmap_   = &c.FuncPrototype{Name: _fn_glGenerateMipmap_, OutType: c.Void, InTypes: _typs_I32}
)

func _addr(fp *c.FuncPrototype) *c.FuncPrototype {
	if fp.Ptr != nil {
		return fp
	}

	n := c.CStr(fp.Name)
	defer usf.Free(n)

	fp.Ptr = _GLLib.Call(_func_glfwGetProcAddress_, []interface{}{&n}).PtrFree()
	return fp
}
func Viewport(x, y, width, height int32) {
	_GLLib.Call(_addr(_func_glViewport_), []interface{}{&x, &y, &width, &height})
}
func Clear(mask GLbitfield) {
	_GLLib.Call(_addr(_func_glClear_), []interface{}{&mask})
}
func ClearColor(r, g, b, a float32) {
	_GLLib.Call(_addr(_func_glClearColor_), []interface{}{&r, &g, &b, &a})
}
func GenTextures(num int32) (textIds []uint32) {
	textIds = make([]uint32, num)
	ptr := &textIds[0]
	_GLLib.Call(_addr(_func_glGenTextures_), []interface{}{&num, &ptr})
	return
}
func DeleteTextures(textIds []uint32) {
	num := int32(len(textIds))
	ptr := &textIds[0]
	_GLLib.Call(_addr(_func_glDeleteTextures_), []interface{}{&num, &ptr})
}

func BindTexture(target GLTexTag, texId uint32) {
	_GLLib.Call(_addr(_func_glBindTexture_), []interface{}{&target, &texId})
}

// void glTexParameteri(GLenum target, GLenum pname, GLint param);
func TexParameteri(tag GLTexTag, pname GLTexPName, param int32) {
	_GLLib.Call(_addr(_func_glTexParameteri_), []interface{}{&tag, &pname, &param})
}

// void glTexImage2D(  GLenum target, GLint level, GLint internalFormat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid * data);
func TexImage2D(target GLTexTag, level GLTexLevel, internalFormat GLTexInterFmt,
	width, height uint32, border int32, dataFmt GLTexDataFmt, dataTyp GLTexDataType, data unsafe.Pointer) {
	_GLLib.Call(_addr(_func_glTexImage2D_),
		[]interface{}{&target, &level, &internalFormat, &width, &height, &border, &dataFmt, &dataTyp, &data})
}

func GenerateMipmap(target GLTexTag) {
	_GLLib.Call(_addr(_func_glGenerateMipmap_), []interface{}{&target})
}
