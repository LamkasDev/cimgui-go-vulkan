//go:build required
// +build required

package imgui

import (
	_ "github.com/LamkasDev/cimgui-go-vulkan/lib/linux/x64"
	_ "github.com/LamkasDev/cimgui-go-vulkan/lib/macos/arm64"
	_ "github.com/LamkasDev/cimgui-go-vulkan/lib/macos/x64"
	_ "github.com/LamkasDev/cimgui-go-vulkan/lib/windows/x64"
)
