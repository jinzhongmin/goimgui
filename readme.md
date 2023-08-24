# goimgui

## example

``` go
package main

import (
	"runtime"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/goimgui/pkg/gi"
	"github.com/jinzhongmin/goimgui/pkg/glfw"
)

func main() {
	runtime.LockOSThread()

    //get from  github.com/jinzhongmin/libgoimgui
	gi.InitLaunch("glfw3.dll", "cimgui.dll", c.ModeLazy)
    
	win := glfw.CreateWindow(900, 600, "hello", nil, nil)
	win.MakeContextCurrent()

	world := "你好世界"

	gi.QuickLaunch(win, func(w, h int32) bool {
		gi.Begin(world, nil, 0)

		gi.End()
		return true
	})
}

```