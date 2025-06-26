package codegen

import (
	"compiler/codegen/arm64"
	"fmt"
)

// GeneratorFactory crea generadores para diferentes plataformas
type GeneratorFactory struct{}

// NewGeneratorFactory crea una nueva fábrica de generadores
func NewGeneratorFactory() *GeneratorFactory {
	return &GeneratorFactory{}
}

// CreateGenerator crea un generador para la plataforma especificada
func (f *GeneratorFactory) CreateGenerator(platform TargetPlatform) (CodeGenerator, error) {
	switch platform {
	case ARM64:
		return arm64.NewARM64Generator(), nil
	case X86_64:
		return nil, fmt.Errorf("x86_64 generator not yet implemented")
	case WASM:
		return nil, fmt.Errorf("WASM generator not yet implemented")
	default:
		return nil, fmt.Errorf("unsupported platform: %s", platform)
	}
}
