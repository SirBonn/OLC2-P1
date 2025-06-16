package repl

import (
	"fmt"
	"log"

	parser "compiler/parser"
	"compiler/value"

	"github.com/antlr4-go/antlr/v4"
)

type DclVisitor struct {
	parser.BaseVlangVisitor
	ScopeTrace  *ScopeTrace
	ErrorTable  *ErrorTable
	StructNames []string
}

func NewDclVisitor(errorTable *ErrorTable) *DclVisitor {
	return &DclVisitor{
		ScopeTrace:  NewScopeTrace(),
		ErrorTable:  errorTable,
		StructNames: []string{},
	}
}

func (v *DclVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *DclVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return nil
}

func (v *DclVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {

	if ctx.Func_dcl() != nil {
		v.Visit(ctx.Func_dcl())
	} else if ctx.Struct_dcl() != nil {
		v.Visit(ctx.Struct_dcl())
	}

	return nil
}

func (v *DclVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {

	// Entorno -> Global
	// Entorno -> Actual
	if v.ScopeTrace.CurrentScope != v.ScopeTrace.GlobalScope {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Las funciones solo pueden ser declaradas en el scope global")
	}

	funcName := ctx.ID().GetText()

	args := make([]*Param, 0)

	if ctx.Arg_list() != nil {
		args = v.Visit(ctx.Arg_list()).([]*Param)
	}

	if len(args) > 0 {

		// ParamType() -> IVOR_NIL | IVOR_INT | IVOR_FLOAT | IVOR_STRING | IVOR_BOOL
		baseParamType := args[0].ParamType()

		for _, param := range args {
			if param.ParamType() != baseParamType {
				v.ErrorTable.NewSemanticError(param.Token, "Todos los parametros de la funcion deben ser del mismo tipo")
				return nil
			}
		}

		returnType := value.IVOR_NIL
		var returnTypeToken antlr.Token = nil

		body := ctx.AllStmt()

		function := &Function{ // pointer ?
			Name:            funcName,
			Param:           args,
			ReturnType:      returnType,
			Body:            body,
			DeclScope:       v.ScopeTrace.CurrentScope,
			ReturnTypeToken: returnTypeToken,
			Token:           ctx.GetStart(),
		}

		ok, msg := v.ScopeTrace.AddFunction(funcName, function)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}

		return nil
	}
	return nil
}
func (v *DclVisitor) VisitArgList(ctx *parser.ArgListContext) interface{} {

	if ctx == nil {
		// fmt.Println("DEBUG: ctx es nil en VisitArgList")
		return nil
	}
	// fmt.Printf("DEBUG: ArgList texto: '%s'\n", ctx.GetText())
	// fmt.Printf("DEBUG: ArgList hijos: %d\n", ctx.GetChildCount())

	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		// fmt.Printf("DEBUG: Hijo %d: %T = '%s'\n", i, child, child)

		if paramCtx, ok := child.(*parser.ParametrosContext); ok {
			// fmt.Printf("DEBUG: Procesando parámetro %d\n", i)

			// fmt.Println("DEBUG: *** A PUNTO DE LLAMAR v.Visit(paramCtx) ***")
			result := v.Visit(paramCtx)
			// fmt.Printf("DEBUG: Resultado: %v (tipo: %T)\n", result, result)

			// AQUÍ ES LA LÍNEA 126 DEL ERROR
			if result == nil {
				// fmt.Println("DEBUG: ¡AQUÍ ES DONDE result ES NIL!")
				continue
			}

			// Esta es probablemente la línea que falla
			// fmt.Println("DEBUG: A punto de hacer conversión de tipo...")
			// param, ok := result.(*Param)
			if !ok {
				fmt.Printf("DEBUG: Conversión falló. Tipo real: %T\n", result)
				continue
			}

			// fmt.Printf("DEBUG: Conversión exitosa: %+v\n", param)
		}
	}

	return []*Param{} // Retornar algo válido por ahora
}
