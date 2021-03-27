package gogl

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/go-gl/gl/v3.3-core/gl"
	"fmt"
	"strings"
)

func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func MakeShader(shaderSource string, shaderType uint32) uint32{
	shaderId := gl.CreateShader(shaderType)
	shaderSource = shaderSource + "\x00"
	csource, free := gl.Strs(shaderSource)
	gl.ShaderSource(shaderId, 1, csource, nil)
	free()
	gl.CompileShader(shaderId)
	var status int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderId, logLength, nil, gl.Str(log))
		panic("Failed to compile shader: \n" + log)
	}
	return shaderId
}