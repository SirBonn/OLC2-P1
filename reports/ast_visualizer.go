package reports

import (
	"compiler/ast"
	"fmt"
	"strings"
)

// ASTVisualizer genera una representación visual del AST
type ASTVisualizer struct {
	ast.Visitor
	nodeCounter int
	dotBuilder  strings.Builder
	edges       []string
}

// NewASTVisualizer crea un nuevo visualizador
func NewASTVisualizer() *ASTVisualizer {
	return &ASTVisualizer{
		nodeCounter: 0,
		edges:       make([]string, 0),
	}
}

// GenerateDOT genera el código DOT para visualizar el AST
func (v *ASTVisualizer) GenerateDOT(root *ast.Program) string {
	v.dotBuilder.WriteString("digraph AST {\n")
	v.dotBuilder.WriteString("  rankdir=TD;\n")
	v.dotBuilder.WriteString("  node [shape=box, style=rounded, fontname=\"Arial\"];\n")
	v.dotBuilder.WriteString("  edge [fontname=\"Arial\"];\n\n")

	// Visitar el árbol
	root.Accept(v)

	// Agregar todas las aristas
	for _, edge := range v.edges {
		v.dotBuilder.WriteString(edge)
	}

	v.dotBuilder.WriteString("}\n")
	return v.dotBuilder.String()
}

// GenerateHTML genera un reporte HTML del AST
func GenerateASTHTML(root *ast.Program) string {
	visualizer := NewASTVisualizer()
	dot := visualizer.GenerateDOT(root)

	html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Reporte AST - V-Lang Cherry</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/viz.js/2.1.2/viz.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/viz.js/2.1.2/full.render.js"></script>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
            background-color: #f5f5f5;
        }
        h1 {
            color: #333;
            text-align: center;
        }
        #graph-container {
            background: white;
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 20px;
            margin: 20px auto;
            max-width: 95%;
            overflow: auto;
        }
        #graph {
            text-align: center;
        }
        .info {
            background: #e8f4f8;
            border-left: 4px solid #2196F3;
            padding: 10px;
            margin: 20px 0;
        }
    </style>
</head>
<body>
    <h1>🌳 Árbol de Sintaxis Abstracta (AST)</h1>
    <div class="info">
        <strong>Información:</strong> Este reporte muestra la estructura del AST generado a partir del código V-Lang Cherry.
        Los nodos representan las diferentes construcciones del lenguaje y las aristas muestran las relaciones padre-hijo.
    </div>
    <div id="graph-container">
        <div id="graph"></div>
    </div>
    <script>
        var viz = new Viz();
        var dotString = ` + "`" + dot + "`" + `;
        
        viz.renderSVGElement(dotString)
        .then(function(element) {
            document.getElementById('graph').appendChild(element);
        })
        .catch(error => {
            console.error(error);
            document.getElementById('graph').innerHTML = '<p style="color: red;">Error al generar el gráfico</p>';
        });
    </script>
</body>
</html>`

	return html
}

// Helper para crear un ID único para cada nodo
func (v *ASTVisualizer) newNodeID() string {
	v.nodeCounter++
	return fmt.Sprintf("node%d", v.nodeCounter)
}

// Helper para agregar un nodo al gráfico
func (v *ASTVisualizer) addNode(id, label, color string) {
	v.dotBuilder.WriteString(fmt.Sprintf("  %s [label=\"%s\", fillcolor=\"%s\", style=filled];\n",
		id, label, color))
}

// Helper para agregar una arista
func (v *ASTVisualizer) addEdge(from, to string) {
	v.edges = append(v.edges, fmt.Sprintf("  %s -> %s;\n", from, to))
}

// === VISITOR METHODS ===

func (v *ASTVisualizer) VisitProgram(node *ast.Program) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, "Program", "#FFE082")

	for _, stmt := range node.Statements {
		childID := stmt.Accept(v).(string)
		v.addEdge(nodeID, childID)
	}

	return nodeID
}

func (v *ASTVisualizer) VisitBinaryExpr(node *ast.BinaryExpr) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, fmt.Sprintf("BinaryOp\\n%s", node.Operator), "#81C784")

	leftID := node.Left.Accept(v).(string)
	rightID := node.Right.Accept(v).(string)

	v.addEdge(nodeID, leftID)
	v.addEdge(nodeID, rightID)

	return nodeID
}

func (v *ASTVisualizer) VisitUnaryExpr(node *ast.UnaryExpr) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, fmt.Sprintf("UnaryOp\\n%s", node.Operator), "#4FC3F7")

	operandID := node.Operand.Accept(v).(string)
	v.addEdge(nodeID, operandID)

	return nodeID
}

func (v *ASTVisualizer) VisitLiteral(node *ast.Literal) interface{} {
	nodeID := v.newNodeID()
	label := fmt.Sprintf("%s\\n%v", node.Type, node.Value)
	v.addNode(nodeID, label, "#F06292")
	return nodeID
}

func (v *ASTVisualizer) VisitIdentifier(node *ast.Identifier) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, fmt.Sprintf("ID\\n%s", node.Name), "#BA68C8")
	return nodeID
}

func (v *ASTVisualizer) VisitPrintStmt(node *ast.PrintStmt) interface{} {
	nodeID := v.newNodeID()
	printType := "print"
	if node.NewLine {
		printType = "println"
	}
	v.addNode(nodeID, printType, "#64B5F6")

	for _, arg := range node.Arguments {
		argID := arg.Accept(v).(string)
		v.addEdge(nodeID, argID)
	}

	return nodeID
}

func (v *ASTVisualizer) VisitVarDecl(node *ast.VarDecl) interface{} {
	nodeID := v.newNodeID()
	label := fmt.Sprintf("VarDecl\\n%s", node.Name)
	if node.Type != "" {
		label += fmt.Sprintf("\\n:%s", node.Type)
	}
	if node.IsMutable {
		label = "mut " + label
	}
	v.addNode(nodeID, label, "#FFB74D")

	if node.Value != nil {
		valueID := node.Value.Accept(v).(string)
		v.addEdge(nodeID, valueID)
	}

	return nodeID
}

func (v *ASTVisualizer) VisitAssignment(node *ast.Assignment) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, "Assignment\\n=", "#A1887F")

	targetID := node.Target.Accept(v).(string)
	valueID := node.Value.Accept(v).(string)

	v.addEdge(nodeID, targetID)
	v.addEdge(nodeID, valueID)

	return nodeID
}

func (v *ASTVisualizer) VisitIfStmt(node *ast.IfStmt) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, "If", "#7986CB")

	// Condición
	condID := node.Condition.Accept(v).(string)
	v.addEdge(nodeID, condID)

	// Then branch
	if len(node.ThenBranch) > 0 {
		thenID := v.newNodeID()
		v.addNode(thenID, "Then", "#9CCC65")
		v.addEdge(nodeID, thenID)

		for _, stmt := range node.ThenBranch {
			stmtID := stmt.Accept(v).(string)
			v.addEdge(thenID, stmtID)
		}
	}

	// Else branch
	if len(node.ElseBranch) > 0 {
		elseID := v.newNodeID()
		v.addNode(elseID, "Else", "#EF5350")
		v.addEdge(nodeID, elseID)

		for _, stmt := range node.ElseBranch {
			stmtID := stmt.Accept(v).(string)
			v.addEdge(elseID, stmtID)
		}
	}

	return nodeID
}

func (v *ASTVisualizer) VisitForStmt(node *ast.ForStmt) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, fmt.Sprintf("For\\n%s", node.Variable), "#26A69A")

	// Iterable
	if node.Iterable != nil {
		iterID := node.Iterable.Accept(v).(string)
		v.addEdge(nodeID, iterID)
	}

	// Body
	if len(node.Body) > 0 {
		bodyID := v.newNodeID()
		v.addNode(bodyID, "Body", "#66BB6A")
		v.addEdge(nodeID, bodyID)

		for _, stmt := range node.Body {
			stmtID := stmt.Accept(v).(string)
			v.addEdge(bodyID, stmtID)
		}
	}

	return nodeID
}

func (v *ASTVisualizer) VisitWhileStmt(node *ast.WhileStmt) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, "While", "#42A5F5")

	// Condición
	condID := node.Condition.Accept(v).(string)
	v.addEdge(nodeID, condID)

	// Body
	if len(node.Body) > 0 {
		bodyID := v.newNodeID()
		v.addNode(bodyID, "Body", "#66BB6A")
		v.addEdge(nodeID, bodyID)

		for _, stmt := range node.Body {
			stmtID := stmt.Accept(v).(string)
			v.addEdge(bodyID, stmtID)
		}
	}

	return nodeID
}

func (v *ASTVisualizer) VisitFuncDecl(node *ast.FuncDecl) interface{} {
	nodeID := v.newNodeID()
	label := fmt.Sprintf("FuncDecl\\n%s", node.Name)
	if node.ReturnType != "" {
		label += fmt.Sprintf("\\n→ %s", node.ReturnType)
	}
	v.addNode(nodeID, label, "#AB47BC")

	// Parámetros
	if len(node.Parameters) > 0 {
		paramsID := v.newNodeID()
		v.addNode(paramsID, "Parameters", "#CE93D8")
		v.addEdge(nodeID, paramsID)

		for _, param := range node.Parameters {
			paramID := v.newNodeID()
			v.addNode(paramID, fmt.Sprintf("%s: %s", param.Name, param.Type), "#E1BEE7")
			v.addEdge(paramsID, paramID)
		}
	}

	// Body
	if len(node.Body) > 0 {
		bodyID := v.newNodeID()
		v.addNode(bodyID, "Body", "#66BB6A")
		v.addEdge(nodeID, bodyID)

		for _, stmt := range node.Body {
			stmtID := stmt.Accept(v).(string)
			v.addEdge(bodyID, stmtID)
		}
	}

	return nodeID
}

func (v *ASTVisualizer) VisitFuncCall(node *ast.FuncCall) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, fmt.Sprintf("Call\\n%s", node.Name), "#FF7043")

	for _, arg := range node.Arguments {
		argID := arg.Accept(v).(string)
		v.addEdge(nodeID, argID)
	}

	return nodeID
}

func (v *ASTVisualizer) VisitReturn(node *ast.Return) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, "Return", "#8D6E63")

	if node.Value != nil {
		valueID := node.Value.Accept(v).(string)
		v.addEdge(nodeID, valueID)
	}

	return nodeID
}

func (v *ASTVisualizer) VisitBreak(node *ast.Break) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, "Break", "#E57373")
	return nodeID
}

func (v *ASTVisualizer) VisitContinue(node *ast.Continue) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, "Continue", "#FFB300")
	return nodeID
}

func (v *ASTVisualizer) VisitStructDecl(node *ast.StructDecl) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, fmt.Sprintf("Struct\\n%s", node.Name), "#4DB6AC")

	// Fields
	if len(node.Fields) > 0 {
		fieldsID := v.newNodeID()
		v.addNode(fieldsID, "Fields", "#80CBC4")
		v.addEdge(nodeID, fieldsID)

		for _, field := range node.Fields {
			fieldID := v.newNodeID()
			v.addNode(fieldID, fmt.Sprintf("%s: %s", field.Name, field.Type), "#B2DFDB")
			v.addEdge(fieldsID, fieldID)
		}
	}

	return nodeID
}

func (v *ASTVisualizer) VisitStructInstance(node *ast.StructInstance) interface{} {
	nodeID := v.newNodeID()
	v.addNode(nodeID, fmt.Sprintf("New\\n%s", node.StructName), "#00ACC1")

	// Fields initialization
	for name, expr := range node.Fields {
		fieldID := v.newNodeID()
		v.addNode(fieldID, name, "#4DD0E1")
		v.addEdge(nodeID, fieldID)

		if expr != nil {
			exprID := expr.Accept(v).(string)
			v.addEdge(fieldID, exprID)
		}
	}

	return nodeID
}
