//go:build required
// +build required

package imgui

import (
	_ "github.com/elokore/cimgui-go-vulkan/lib/linux/x64"
	_ "github.com/elokore/cimgui-go-vulkanulkan/lib/macos/arm64"
	_ "github.com/elokore/cimgui-go-vulkanulkan/lib/macos/x64"
	_ "github.com/elokore/cimgui-go-vulkanulkan/lib/windows/x64"
)
