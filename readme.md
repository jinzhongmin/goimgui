# goimgui

## example

``` go
package main

import (
	"runtime"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/goimgui/pkg/gi"
)

func main() {
	runtime.LockOSThread()
	gi.QuickInitLibraryGlfw("glfw3.dll", "cimgui.dll", c.ModeLazy)

	_, Str := gi.NewMyStrs()
	gi.QuickLaunchGlfw("hello", 900, 600, func(io *gi.QuickLaunchGlfwIO) {
		gi.Begin(Str("你好 imgui"), nil, 0)
		{
			if gi.Button(Str("关闭主窗口(close main window)"), gi.ImVec2XY(240, 35)) {
				io.Close()
			}
			gi.End()
		}
	})
}

```