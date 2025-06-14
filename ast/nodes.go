// ast/nodes.go
package ast

// Interfaces base
type Node interface {
	Accept(visitor Visitor) interface{}
	GetLine() int
	GetColumn() int
}

type Expression interface {
	Node
	IsExpression()
}

type Statement interface {
	Node
	IsStatement()
}

// Visitor interface
type Visitor interface {
	VisitProgram(node *Program) interface{}
	VisitBinaryExpr(node *BinaryExpr) interface{}
	VisitUnaryExpr(node *UnaryExpr) interface{}
	VisitLiteral(node *Literal) interface{}
	VisitIdentifier(node *Identifier) interface{}
	VisitPrintStmt(node *PrintStmt) interface{}
	VisitVarDecl(node *VarDecl) interface{}
	VisitAssignment(node *Assignment) interface{}
	VisitIfStmt(node *IfStmt) interface{}
	VisitForStmt(node *ForStmt) interface{}
	VisitWhileStmt(node *WhileStmt) interface{}
	VisitFuncDecl(node *FuncDecl) interface{}
	VisitFuncCall(node *FuncCall) interface{}
	VisitReturn(node *Return) interface{}
	VisitBreak(node *Break) interface{}
	VisitContinue(node *Continue) interface{}
	VisitStructDecl(node *StructDecl) interface{}
	VisitStructInstance(node *StructInstance) interface{}
}

// === PROGRAMA ===
type Program struct {
	Statements []Statement
}

func (p *Program) Accept(v Visitor) interface{} { return v.VisitProgram(p) }
func (p *Program) GetLine() int                 { return 0 }
func (p *Program) GetColumn() int               { return 0 }

// === EXPRESIONES ===
type BinaryExpr struct {
	Left     Expression
	Operator string
	Right    Expression
	Line     int
	Column   int
}

func (b *BinaryExpr) Accept(v Visitor) interface{} { return v.VisitBinaryExpr(b) }
func (b *BinaryExpr) IsExpression()                {}
func (b *BinaryExpr) GetLine() int                 { return b.Line }
func (b *BinaryExpr) GetColumn() int               { return b.Column }

type UnaryExpr struct {
	Operator string
	Operand  Expression
	Line     int
	Column   int
}

func (u *UnaryExpr) Accept(v Visitor) interface{} { return v.VisitUnaryExpr(u) }
func (u *UnaryExpr) IsExpression()                {}
func (u *UnaryExpr) GetLine() int                 { return u.Line }
func (u *UnaryExpr) GetColumn() int               { return u.Column }

type Literal struct {
	Value  interface{}
	Type   string // "int", "float64", "string", "bool", "char"
	Line   int
	Column int
}

func (l *Literal) Accept(v Visitor) interface{} { return v.VisitLiteral(l) }
func (l *Literal) IsExpression()                {}
func (l *Literal) GetLine() int                 { return l.Line }
func (l *Literal) GetColumn() int               { return l.Column }

type Identifier struct {
	Name   string
	Line   int
	Column int
}

func (i *Identifier) Accept(v Visitor) interface{} { return v.VisitIdentifier(i) }
func (i *Identifier) IsExpression()                {}
func (i *Identifier) GetLine() int                 { return i.Line }
func (i *Identifier) GetColumn() int               { return i.Column }

type FuncCall struct {
	Name      string
	Arguments []Expression
	Line      int
	Column    int
}

func (f *FuncCall) Accept(v Visitor) interface{} { return v.VisitFuncCall(f) }
func (f *FuncCall) IsExpression()                {}
func (f *FuncCall) GetLine() int                 { return f.Line }
func (f *FuncCall) GetColumn() int               { return f.Column }

// === SENTENCIAS ===
type PrintStmt struct {
	Arguments []Expression
	NewLine   bool // true para println, false para print
	Line      int
	Column    int
}

func (p *PrintStmt) Accept(v Visitor) interface{} { return v.VisitPrintStmt(p) }
func (p *PrintStmt) IsStatement()                 {}
func (p *PrintStmt) GetLine() int                 { return p.Line }
func (p *PrintStmt) GetColumn() int               { return p.Column }

type VarDecl struct {
	Name      string
	Type      string // Tipo explícito o inferido
	Value     Expression
	IsMutable bool
	Line      int
	Column    int
}

func (v *VarDecl) Accept(vi Visitor) interface{} { return vi.VisitVarDecl(v) }
func (v *VarDecl) IsStatement()                  {}
func (v *VarDecl) GetLine() int                  { return v.Line }
func (v *VarDecl) GetColumn() int                { return v.Column }

type Assignment struct {
	Target Expression // Identifier o acceso a campo
	Value  Expression
	Line   int
	Column int
}

func (a *Assignment) Accept(v Visitor) interface{} { return v.VisitAssignment(a) }
func (a *Assignment) IsStatement()                 {}
func (a *Assignment) GetLine() int                 { return a.Line }
func (a *Assignment) GetColumn() int               { return a.Column }

type IfStmt struct {
	Condition  Expression
	ThenBranch []Statement
	ElseBranch []Statement // puede ser vacío
	Line       int
	Column     int
}

func (i *IfStmt) Accept(v Visitor) interface{} { return v.VisitIfStmt(i) }
func (i *IfStmt) IsStatement()                 {}
func (i *IfStmt) GetLine() int                 { return i.Line }
func (i *IfStmt) GetColumn() int               { return i.Column }

type ForStmt struct {
	Variable string
	Iterable Expression
	Body     []Statement
	Line     int
	Column   int
}

func (f *ForStmt) Accept(v Visitor) interface{} { return v.VisitForStmt(f) }
func (f *ForStmt) IsStatement()                 {}
func (f *ForStmt) GetLine() int                 { return f.Line }
func (f *ForStmt) GetColumn() int               { return f.Column }

type WhileStmt struct {
	Condition Expression
	Body      []Statement
	Line      int
	Column    int
}

func (w *WhileStmt) Accept(v Visitor) interface{} { return v.VisitWhileStmt(w) }
func (w *WhileStmt) IsStatement()                 {}
func (w *WhileStmt) GetLine() int                 { return w.Line }
func (w *WhileStmt) GetColumn() int               { return w.Column }

type FuncDecl struct {
	Name       string
	Parameters []Parameter
	ReturnType string
	Body       []Statement
	Line       int
	Column     int
}

type Parameter struct {
	Name string
	Type string
}

func (f *FuncDecl) Accept(v Visitor) interface{} { return v.VisitFuncDecl(f) }
func (f *FuncDecl) IsStatement()                 {}
func (f *FuncDecl) GetLine() int                 { return f.Line }
func (f *FuncDecl) GetColumn() int               { return f.Column }

type Return struct {
	Value  Expression // puede ser nil
	Line   int
	Column int
}

func (r *Return) Accept(v Visitor) interface{} { return v.VisitReturn(r) }
func (r *Return) IsStatement()                 {}
func (r *Return) GetLine() int                 { return r.Line }
func (r *Return) GetColumn() int               { return r.Column }

type Break struct {
	Line   int
	Column int
}

func (b *Break) Accept(v Visitor) interface{} { return v.VisitBreak(b) }
func (b *Break) IsStatement()                 {}
func (b *Break) GetLine() int                 { return b.Line }
func (b *Break) GetColumn() int               { return b.Column }

type Continue struct {
	Line   int
	Column int
}

func (c *Continue) Accept(v Visitor) interface{} { return v.VisitContinue(c) }
func (c *Continue) IsStatement()                 {}
func (c *Continue) GetLine() int                 { return c.Line }
func (c *Continue) GetColumn() int               { return c.Column }

// === STRUCTS ===
type StructDecl struct {
	Name   string
	Fields []Field
	Line   int
	Column int
}

type Field struct {
	Name string
	Type string
}

func (s *StructDecl) Accept(v Visitor) interface{} { return v.VisitStructDecl(s) }
func (s *StructDecl) IsStatement()                 {}
func (s *StructDecl) GetLine() int                 { return s.Line }
func (s *StructDecl) GetColumn() int               { return s.Column }

type StructInstance struct {
	StructName string
	Fields     map[string]Expression
	Line       int
	Column     int
}

func (s *StructInstance) Accept(v Visitor) interface{} { return v.VisitStructInstance(s) }
func (s *StructInstance) IsExpression()                {}
func (s *StructInstance) GetLine() int                 { return s.Line }
func (s *StructInstance) GetColumn() int               { return s.Column }
