//go:build ((!gles && !arm && !arm64) || darwin) && !js && !wasm && !test_web_driver
// +build !gles,!arm,!arm64 darwin
// +build !js
// +build !wasm
// +build !test_web_driver

package glfw

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func initWindowHints() {
	fmt.Println("Initialising glfw")
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)

	//glfw.WindowHint(glfw.CocoaGraphicsSwitching, glfw.True)
}
