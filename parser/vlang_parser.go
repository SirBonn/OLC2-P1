// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VlangParser struct {
	*antlr.BaseParser
}

var VlangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.LiteralNames = []string{
		"", "'println'", "'print'", "'range'", "'fn'", "'struct'", "'int'",
		"'float'", "'f64'", "'string'", "'bool'", "'char'", "'void'", "'len'",
		"'cap'", "'append'", "'mut'", "", "", "", "", "", "", "", "'if'", "'else'",
		"'while'", "'for'", "'in'", "'return'", "'break'", "'continue'", "'switch'",
		"'case'", "'default'", "'fallthrough'", "", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'!'", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='", "'>'",
		"'>='", "'='", "'++'", "'--'", "'+='", "'-='", "'*='", "'/='", "'%='",
		"':'", "';'", "'('", "')'", "'['", "'{'", "']'", "'}'", "'.'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "LEN", "CAP", "APPEND",
		"MUT", "BOOLEANO", "ENTERO", "FLOAT", "DECIMAL", "CADENA", "CARACTER",
		"STRING_INTERPOLATION", "IF_KW", "ELSE_KW", "WHILE_KW", "FOR_KW", "IN_KW",
		"RETURN_KW", "BREAK_KW", "CONTINUE_KW", "SWITCH_KW", "CASE_KW", "DEFAULT_KW",
		"FALLTHROUGH_KW", "ID", "PLUS", "MINUS", "MUL", "DIV", "MOD", "NOT",
		"OR", "AND", "EQ", "NEQ", "LT", "LE", "GT", "GE", "ASSIGN", "INC", "DEC",
		"PLUS_ASSIGN", "MINUS_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN",
		"COLON", "SEMICOLON", "LPAREN", "RPAREN", "LBRACK", "LCOR", "RBRACK",
		"RCOR", "DOT", "COMMA", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"programa", "stmt", "if_stmt", "if_chain", "else_stmt", "println", "print",
		"while_stmt", "for_stmt", "for_enhanced_stmt", "for_init", "for_update",
		"switch_stmt", "switch_case", "case_values", "default_case", "range",
		"func_call", "func_dcl", "struct_dcl", "struct_field", "transfer_stmt",
		"assign_stmt", "decl_stmt", "array_literal", "id_pattern", "expresion",
		"parametros", "func_param", "arg_list", "func_arg", "valores", "valor",
		"strct_instancia", "struct_prop", "var_type", "struct_vector", "incredecre",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 564, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 1, 0, 4, 0, 78, 8, 0, 11, 0, 12, 0, 79, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 98, 8, 1, 1, 2, 1, 2, 1, 2, 5, 2, 103, 8, 2, 10, 2, 12, 2,
		106, 9, 2, 1, 2, 3, 2, 109, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 115, 8,
		3, 10, 3, 12, 3, 118, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 125, 8,
		4, 10, 4, 12, 4, 128, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5,
		5, 137, 8, 5, 10, 5, 12, 5, 140, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 5, 6, 149, 8, 6, 10, 6, 12, 6, 152, 9, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 160, 8, 7, 10, 7, 12, 7, 163, 9, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 172, 8, 8, 1, 8, 1, 8, 5, 8, 176, 8, 8,
		10, 8, 12, 8, 179, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 187,
		8, 9, 10, 9, 12, 9, 190, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 196, 8, 9,
		1, 9, 1, 9, 3, 9, 200, 8, 9, 1, 9, 1, 9, 3, 9, 204, 8, 9, 1, 9, 1, 9, 5,
		9, 208, 8, 9, 10, 9, 12, 9, 211, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 5, 9, 222, 8, 9, 10, 9, 12, 9, 225, 9, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 5, 9, 232, 8, 9, 10, 9, 12, 9, 235, 9, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 247, 8, 9, 10,
		9, 12, 9, 250, 9, 9, 1, 9, 1, 9, 3, 9, 254, 8, 9, 1, 10, 1, 10, 3, 10,
		258, 8, 10, 1, 11, 1, 11, 3, 11, 262, 8, 11, 1, 12, 1, 12, 3, 12, 266,
		8, 12, 1, 12, 1, 12, 5, 12, 270, 8, 12, 10, 12, 12, 12, 273, 9, 12, 1,
		12, 3, 12, 276, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13,
		284, 8, 13, 10, 13, 12, 13, 287, 9, 13, 1, 14, 1, 14, 1, 14, 5, 14, 292,
		8, 14, 10, 14, 12, 14, 295, 9, 14, 1, 15, 1, 15, 1, 15, 5, 15, 300, 8,
		15, 10, 15, 12, 15, 303, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 3, 17, 314, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 322, 8, 18, 1, 18, 1, 18, 3, 18, 326, 8, 18, 1, 18, 1,
		18, 5, 18, 330, 8, 18, 10, 18, 12, 18, 333, 9, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 5, 19, 341, 8, 19, 10, 19, 12, 19, 344, 9, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 3, 21, 353, 8, 21, 1, 21,
		1, 21, 3, 21, 357, 8, 21, 1, 21, 1, 21, 3, 21, 361, 8, 21, 1, 21, 3, 21,
		364, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 390, 8, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 408, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 5, 24, 417, 8, 24, 10, 24, 12, 24, 420, 9, 24, 3, 24,
		422, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 5, 25, 429, 8, 25, 10, 25,
		12, 25, 432, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 446, 8, 26, 10, 26, 12, 26, 449,
		9, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 466, 8, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 486, 8, 26, 10, 26, 12, 26,
		489, 9, 26, 1, 27, 1, 27, 1, 27, 5, 27, 494, 8, 27, 10, 27, 12, 27, 497,
		9, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 5, 29, 504, 8, 29, 10, 29, 12,
		29, 507, 9, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 521, 8, 32, 1, 33, 1, 33, 1, 33, 1,
		33, 5, 33, 527, 8, 33, 10, 33, 12, 33, 530, 9, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 1, 34, 3, 34, 538, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 550, 8, 35, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 562, 8,
		37, 1, 37, 0, 1, 52, 38, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 0, 5, 2, 0, 38, 38, 42, 42, 1, 0, 39, 41, 1,
		0, 37, 38, 1, 0, 47, 50, 1, 0, 45, 46, 625, 0, 77, 1, 0, 0, 0, 2, 97, 1,
		0, 0, 0, 4, 99, 1, 0, 0, 0, 6, 110, 1, 0, 0, 0, 8, 121, 1, 0, 0, 0, 10,
		131, 1, 0, 0, 0, 12, 143, 1, 0, 0, 0, 14, 155, 1, 0, 0, 0, 16, 166, 1,
		0, 0, 0, 18, 253, 1, 0, 0, 0, 20, 257, 1, 0, 0, 0, 22, 261, 1, 0, 0, 0,
		24, 263, 1, 0, 0, 0, 26, 279, 1, 0, 0, 0, 28, 288, 1, 0, 0, 0, 30, 296,
		1, 0, 0, 0, 32, 304, 1, 0, 0, 0, 34, 310, 1, 0, 0, 0, 36, 317, 1, 0, 0,
		0, 38, 336, 1, 0, 0, 0, 40, 347, 1, 0, 0, 0, 42, 363, 1, 0, 0, 0, 44, 389,
		1, 0, 0, 0, 46, 407, 1, 0, 0, 0, 48, 409, 1, 0, 0, 0, 50, 425, 1, 0, 0,
		0, 52, 465, 1, 0, 0, 0, 54, 490, 1, 0, 0, 0, 56, 498, 1, 0, 0, 0, 58, 500,
		1, 0, 0, 0, 60, 508, 1, 0, 0, 0, 62, 511, 1, 0, 0, 0, 64, 520, 1, 0, 0,
		0, 66, 522, 1, 0, 0, 0, 68, 533, 1, 0, 0, 0, 70, 549, 1, 0, 0, 0, 72, 551,
		1, 0, 0, 0, 74, 561, 1, 0, 0, 0, 76, 78, 3, 2, 1, 0, 77, 76, 1, 0, 0, 0,
		78, 79, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 1,
		0, 0, 0, 81, 82, 5, 0, 0, 1, 82, 1, 1, 0, 0, 0, 83, 98, 3, 46, 23, 0, 84,
		98, 3, 44, 22, 0, 85, 98, 3, 4, 2, 0, 86, 98, 3, 14, 7, 0, 87, 98, 3, 16,
		8, 0, 88, 98, 3, 18, 9, 0, 89, 98, 3, 24, 12, 0, 90, 98, 3, 34, 17, 0,
		91, 98, 3, 36, 18, 0, 92, 98, 3, 38, 19, 0, 93, 98, 3, 66, 33, 0, 94, 98,
		3, 42, 21, 0, 95, 98, 3, 10, 5, 0, 96, 98, 3, 12, 6, 0, 97, 83, 1, 0, 0,
		0, 97, 84, 1, 0, 0, 0, 97, 85, 1, 0, 0, 0, 97, 86, 1, 0, 0, 0, 97, 87,
		1, 0, 0, 0, 97, 88, 1, 0, 0, 0, 97, 89, 1, 0, 0, 0, 97, 90, 1, 0, 0, 0,
		97, 91, 1, 0, 0, 0, 97, 92, 1, 0, 0, 0, 97, 93, 1, 0, 0, 0, 97, 94, 1,
		0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 3, 1, 0, 0, 0, 99,
		104, 3, 6, 3, 0, 100, 101, 5, 25, 0, 0, 101, 103, 3, 6, 3, 0, 102, 100,
		1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0,
		0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 109, 3, 8, 4, 0,
		108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 5, 1, 0, 0, 0, 110, 111,
		5, 24, 0, 0, 111, 112, 3, 52, 26, 0, 112, 116, 5, 64, 0, 0, 113, 115, 3,
		2, 1, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0,
		0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119,
		120, 5, 66, 0, 0, 120, 7, 1, 0, 0, 0, 121, 122, 5, 25, 0, 0, 122, 126,
		5, 64, 0, 0, 123, 125, 3, 2, 1, 0, 124, 123, 1, 0, 0, 0, 125, 128, 1, 0,
		0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 129, 1, 0, 0, 0,
		128, 126, 1, 0, 0, 0, 129, 130, 5, 66, 0, 0, 130, 9, 1, 0, 0, 0, 131, 132,
		5, 1, 0, 0, 132, 133, 5, 61, 0, 0, 133, 138, 3, 52, 26, 0, 134, 135, 5,
		68, 0, 0, 135, 137, 3, 52, 26, 0, 136, 134, 1, 0, 0, 0, 137, 140, 1, 0,
		0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0,
		140, 138, 1, 0, 0, 0, 141, 142, 5, 62, 0, 0, 142, 11, 1, 0, 0, 0, 143,
		144, 5, 2, 0, 0, 144, 145, 5, 61, 0, 0, 145, 150, 3, 52, 26, 0, 146, 147,
		5, 68, 0, 0, 147, 149, 3, 52, 26, 0, 148, 146, 1, 0, 0, 0, 149, 152, 1,
		0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0,
		0, 152, 150, 1, 0, 0, 0, 153, 154, 5, 62, 0, 0, 154, 13, 1, 0, 0, 0, 155,
		156, 5, 26, 0, 0, 156, 157, 3, 52, 26, 0, 157, 161, 5, 64, 0, 0, 158, 160,
		3, 2, 1, 0, 159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0,
		0, 0, 161, 162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0,
		164, 165, 5, 66, 0, 0, 165, 15, 1, 0, 0, 0, 166, 167, 5, 27, 0, 0, 167,
		168, 5, 36, 0, 0, 168, 171, 5, 28, 0, 0, 169, 172, 3, 52, 26, 0, 170, 172,
		3, 32, 16, 0, 171, 169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 173, 1,
		0, 0, 0, 173, 177, 5, 64, 0, 0, 174, 176, 3, 2, 1, 0, 175, 174, 1, 0, 0,
		0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178,
		180, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 181, 5, 66, 0, 0, 181, 17,
		1, 0, 0, 0, 182, 183, 5, 27, 0, 0, 183, 184, 3, 52, 26, 0, 184, 188, 5,
		64, 0, 0, 185, 187, 3, 2, 1, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0,
		0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0, 190,
		188, 1, 0, 0, 0, 191, 192, 5, 66, 0, 0, 192, 254, 1, 0, 0, 0, 193, 195,
		5, 27, 0, 0, 194, 196, 3, 20, 10, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1,
		0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 5, 60, 0, 0, 198, 200, 3, 52,
		26, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 203, 5, 60, 0, 0, 202, 204, 3, 22, 11, 0, 203, 202, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 209, 5, 64, 0, 0, 206, 208,
		3, 2, 1, 0, 207, 206, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0,
		0, 0, 209, 210, 1, 0, 0, 0, 210, 212, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0,
		212, 254, 5, 66, 0, 0, 213, 214, 5, 27, 0, 0, 214, 215, 5, 36, 0, 0, 215,
		216, 5, 68, 0, 0, 216, 217, 5, 36, 0, 0, 217, 218, 5, 28, 0, 0, 218, 219,
		3, 52, 26, 0, 219, 223, 5, 64, 0, 0, 220, 222, 3, 2, 1, 0, 221, 220, 1,
		0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0,
		0, 224, 226, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 227, 5, 66, 0, 0, 227,
		254, 1, 0, 0, 0, 228, 229, 5, 27, 0, 0, 229, 233, 5, 64, 0, 0, 230, 232,
		3, 2, 1, 0, 231, 230, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0,
		0, 0, 233, 234, 1, 0, 0, 0, 234, 236, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0,
		236, 254, 5, 66, 0, 0, 237, 238, 5, 27, 0, 0, 238, 239, 5, 36, 0, 0, 239,
		240, 5, 68, 0, 0, 240, 241, 5, 36, 0, 0, 241, 242, 5, 51, 0, 0, 242, 243,
		5, 3, 0, 0, 243, 244, 3, 52, 26, 0, 244, 248, 5, 64, 0, 0, 245, 247, 3,
		2, 1, 0, 246, 245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0,
		0, 248, 249, 1, 0, 0, 0, 249, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251,
		252, 5, 66, 0, 0, 252, 254, 1, 0, 0, 0, 253, 182, 1, 0, 0, 0, 253, 193,
		1, 0, 0, 0, 253, 213, 1, 0, 0, 0, 253, 228, 1, 0, 0, 0, 253, 237, 1, 0,
		0, 0, 254, 19, 1, 0, 0, 0, 255, 258, 3, 46, 23, 0, 256, 258, 3, 44, 22,
		0, 257, 255, 1, 0, 0, 0, 257, 256, 1, 0, 0, 0, 258, 21, 1, 0, 0, 0, 259,
		262, 3, 44, 22, 0, 260, 262, 3, 74, 37, 0, 261, 259, 1, 0, 0, 0, 261, 260,
		1, 0, 0, 0, 262, 23, 1, 0, 0, 0, 263, 265, 5, 32, 0, 0, 264, 266, 3, 52,
		26, 0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0,
		267, 271, 5, 64, 0, 0, 268, 270, 3, 26, 13, 0, 269, 268, 1, 0, 0, 0, 270,
		273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 275,
		1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 276, 3, 30, 15, 0, 275, 274, 1,
		0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 5, 66, 0,
		0, 278, 25, 1, 0, 0, 0, 279, 280, 5, 33, 0, 0, 280, 281, 3, 28, 14, 0,
		281, 285, 5, 59, 0, 0, 282, 284, 3, 2, 1, 0, 283, 282, 1, 0, 0, 0, 284,
		287, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 27, 1,
		0, 0, 0, 287, 285, 1, 0, 0, 0, 288, 293, 3, 52, 26, 0, 289, 290, 5, 68,
		0, 0, 290, 292, 3, 52, 26, 0, 291, 289, 1, 0, 0, 0, 292, 295, 1, 0, 0,
		0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 29, 1, 0, 0, 0, 295,
		293, 1, 0, 0, 0, 296, 297, 5, 34, 0, 0, 297, 301, 5, 59, 0, 0, 298, 300,
		3, 2, 1, 0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0,
		0, 0, 301, 302, 1, 0, 0, 0, 302, 31, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0,
		304, 305, 3, 52, 26, 0, 305, 306, 5, 67, 0, 0, 306, 307, 5, 67, 0, 0, 307,
		308, 5, 67, 0, 0, 308, 309, 3, 52, 26, 0, 309, 33, 1, 0, 0, 0, 310, 311,
		3, 50, 25, 0, 311, 313, 5, 61, 0, 0, 312, 314, 3, 54, 27, 0, 313, 312,
		1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316, 5, 62,
		0, 0, 316, 35, 1, 0, 0, 0, 317, 318, 5, 4, 0, 0, 318, 319, 5, 36, 0, 0,
		319, 321, 5, 61, 0, 0, 320, 322, 3, 58, 29, 0, 321, 320, 1, 0, 0, 0, 321,
		322, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 325, 5, 62, 0, 0, 324, 326,
		3, 70, 35, 0, 325, 324, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 1,
		0, 0, 0, 327, 331, 5, 64, 0, 0, 328, 330, 3, 2, 1, 0, 329, 328, 1, 0, 0,
		0, 330, 333, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332,
		334, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334, 335, 5, 66, 0, 0, 335, 37,
		1, 0, 0, 0, 336, 337, 5, 5, 0, 0, 337, 338, 5, 36, 0, 0, 338, 342, 5, 64,
		0, 0, 339, 341, 3, 40, 20, 0, 340, 339, 1, 0, 0, 0, 341, 344, 1, 0, 0,
		0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 345, 1, 0, 0, 0, 344,
		342, 1, 0, 0, 0, 345, 346, 5, 66, 0, 0, 346, 39, 1, 0, 0, 0, 347, 348,
		3, 70, 35, 0, 348, 349, 5, 36, 0, 0, 349, 41, 1, 0, 0, 0, 350, 352, 5,
		29, 0, 0, 351, 353, 3, 52, 26, 0, 352, 351, 1, 0, 0, 0, 352, 353, 1, 0,
		0, 0, 353, 364, 1, 0, 0, 0, 354, 356, 5, 30, 0, 0, 355, 357, 5, 36, 0,
		0, 356, 355, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 364, 1, 0, 0, 0, 358,
		360, 5, 31, 0, 0, 359, 361, 5, 36, 0, 0, 360, 359, 1, 0, 0, 0, 360, 361,
		1, 0, 0, 0, 361, 364, 1, 0, 0, 0, 362, 364, 5, 35, 0, 0, 363, 350, 1, 0,
		0, 0, 363, 354, 1, 0, 0, 0, 363, 358, 1, 0, 0, 0, 363, 362, 1, 0, 0, 0,
		364, 43, 1, 0, 0, 0, 365, 366, 3, 50, 25, 0, 366, 367, 5, 51, 0, 0, 367,
		368, 3, 52, 26, 0, 368, 390, 1, 0, 0, 0, 369, 370, 3, 50, 25, 0, 370, 371,
		5, 54, 0, 0, 371, 372, 3, 52, 26, 0, 372, 390, 1, 0, 0, 0, 373, 374, 3,
		50, 25, 0, 374, 375, 5, 55, 0, 0, 375, 376, 3, 52, 26, 0, 376, 390, 1,
		0, 0, 0, 377, 378, 3, 50, 25, 0, 378, 379, 5, 56, 0, 0, 379, 380, 3, 52,
		26, 0, 380, 390, 1, 0, 0, 0, 381, 382, 3, 50, 25, 0, 382, 383, 5, 57, 0,
		0, 383, 384, 3, 52, 26, 0, 384, 390, 1, 0, 0, 0, 385, 386, 3, 50, 25, 0,
		386, 387, 5, 58, 0, 0, 387, 388, 3, 52, 26, 0, 388, 390, 1, 0, 0, 0, 389,
		365, 1, 0, 0, 0, 389, 369, 1, 0, 0, 0, 389, 373, 1, 0, 0, 0, 389, 377,
		1, 0, 0, 0, 389, 381, 1, 0, 0, 0, 389, 385, 1, 0, 0, 0, 390, 45, 1, 0,
		0, 0, 391, 392, 5, 16, 0, 0, 392, 393, 5, 36, 0, 0, 393, 394, 5, 51, 0,
		0, 394, 408, 3, 52, 26, 0, 395, 396, 5, 16, 0, 0, 396, 397, 5, 36, 0, 0,
		397, 398, 3, 70, 35, 0, 398, 399, 5, 51, 0, 0, 399, 400, 3, 52, 26, 0,
		400, 408, 1, 0, 0, 0, 401, 402, 5, 16, 0, 0, 402, 403, 5, 36, 0, 0, 403,
		408, 3, 70, 35, 0, 404, 405, 5, 36, 0, 0, 405, 406, 5, 51, 0, 0, 406, 408,
		3, 48, 24, 0, 407, 391, 1, 0, 0, 0, 407, 395, 1, 0, 0, 0, 407, 401, 1,
		0, 0, 0, 407, 404, 1, 0, 0, 0, 408, 47, 1, 0, 0, 0, 409, 410, 5, 63, 0,
		0, 410, 411, 5, 65, 0, 0, 411, 412, 3, 70, 35, 0, 412, 421, 5, 64, 0, 0,
		413, 418, 3, 52, 26, 0, 414, 415, 5, 68, 0, 0, 415, 417, 3, 52, 26, 0,
		416, 414, 1, 0, 0, 0, 417, 420, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418,
		419, 1, 0, 0, 0, 419, 422, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 421, 413,
		1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 5, 66,
		0, 0, 424, 49, 1, 0, 0, 0, 425, 430, 5, 36, 0, 0, 426, 427, 5, 67, 0, 0,
		427, 429, 5, 36, 0, 0, 428, 426, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430,
		428, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 51, 1, 0, 0, 0, 432, 430, 1,
		0, 0, 0, 433, 434, 6, 26, -1, 0, 434, 466, 3, 64, 32, 0, 435, 466, 3, 34,
		17, 0, 436, 466, 3, 44, 22, 0, 437, 438, 5, 61, 0, 0, 438, 439, 3, 52,
		26, 0, 439, 440, 5, 62, 0, 0, 440, 466, 1, 0, 0, 0, 441, 442, 5, 63, 0,
		0, 442, 447, 3, 52, 26, 0, 443, 444, 5, 68, 0, 0, 444, 446, 3, 52, 26,
		0, 445, 443, 1, 0, 0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 447,
		448, 1, 0, 0, 0, 448, 450, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 450, 451,
		5, 65, 0, 0, 451, 466, 1, 0, 0, 0, 452, 453, 7, 0, 0, 0, 453, 466, 3, 52,
		26, 12, 454, 466, 5, 36, 0, 0, 455, 466, 3, 74, 37, 0, 456, 457, 5, 36,
		0, 0, 457, 458, 5, 67, 0, 0, 458, 466, 5, 36, 0, 0, 459, 460, 5, 36, 0,
		0, 460, 461, 5, 67, 0, 0, 461, 466, 3, 52, 26, 2, 462, 463, 5, 36, 0, 0,
		463, 464, 5, 51, 0, 0, 464, 466, 3, 52, 26, 1, 465, 433, 1, 0, 0, 0, 465,
		435, 1, 0, 0, 0, 465, 436, 1, 0, 0, 0, 465, 437, 1, 0, 0, 0, 465, 441,
		1, 0, 0, 0, 465, 452, 1, 0, 0, 0, 465, 454, 1, 0, 0, 0, 465, 455, 1, 0,
		0, 0, 465, 456, 1, 0, 0, 0, 465, 459, 1, 0, 0, 0, 465, 462, 1, 0, 0, 0,
		466, 487, 1, 0, 0, 0, 467, 468, 10, 11, 0, 0, 468, 469, 7, 1, 0, 0, 469,
		486, 3, 52, 26, 12, 470, 471, 10, 10, 0, 0, 471, 472, 7, 2, 0, 0, 472,
		486, 3, 52, 26, 11, 473, 474, 10, 9, 0, 0, 474, 475, 7, 3, 0, 0, 475, 486,
		3, 52, 26, 10, 476, 477, 10, 8, 0, 0, 477, 478, 7, 4, 0, 0, 478, 486, 3,
		52, 26, 9, 479, 480, 10, 7, 0, 0, 480, 481, 5, 44, 0, 0, 481, 486, 3, 52,
		26, 8, 482, 483, 10, 6, 0, 0, 483, 484, 5, 43, 0, 0, 484, 486, 3, 52, 26,
		7, 485, 467, 1, 0, 0, 0, 485, 470, 1, 0, 0, 0, 485, 473, 1, 0, 0, 0, 485,
		476, 1, 0, 0, 0, 485, 479, 1, 0, 0, 0, 485, 482, 1, 0, 0, 0, 486, 489,
		1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 53, 1, 0,
		0, 0, 489, 487, 1, 0, 0, 0, 490, 495, 3, 56, 28, 0, 491, 492, 5, 68, 0,
		0, 492, 494, 3, 56, 28, 0, 493, 491, 1, 0, 0, 0, 494, 497, 1, 0, 0, 0,
		495, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 55, 1, 0, 0, 0, 497, 495,
		1, 0, 0, 0, 498, 499, 3, 52, 26, 0, 499, 57, 1, 0, 0, 0, 500, 505, 3, 60,
		30, 0, 501, 502, 5, 68, 0, 0, 502, 504, 3, 60, 30, 0, 503, 501, 1, 0, 0,
		0, 504, 507, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 505, 506, 1, 0, 0, 0, 506,
		59, 1, 0, 0, 0, 507, 505, 1, 0, 0, 0, 508, 509, 5, 36, 0, 0, 509, 510,
		3, 70, 35, 0, 510, 61, 1, 0, 0, 0, 511, 512, 3, 64, 32, 0, 512, 63, 1,
		0, 0, 0, 513, 521, 5, 18, 0, 0, 514, 521, 5, 20, 0, 0, 515, 521, 5, 21,
		0, 0, 516, 521, 5, 17, 0, 0, 517, 521, 5, 22, 0, 0, 518, 521, 5, 19, 0,
		0, 519, 521, 5, 23, 0, 0, 520, 513, 1, 0, 0, 0, 520, 514, 1, 0, 0, 0, 520,
		515, 1, 0, 0, 0, 520, 516, 1, 0, 0, 0, 520, 517, 1, 0, 0, 0, 520, 518,
		1, 0, 0, 0, 520, 519, 1, 0, 0, 0, 521, 65, 1, 0, 0, 0, 522, 523, 5, 5,
		0, 0, 523, 524, 5, 36, 0, 0, 524, 528, 5, 64, 0, 0, 525, 527, 3, 68, 34,
		0, 526, 525, 1, 0, 0, 0, 527, 530, 1, 0, 0, 0, 528, 526, 1, 0, 0, 0, 528,
		529, 1, 0, 0, 0, 529, 531, 1, 0, 0, 0, 530, 528, 1, 0, 0, 0, 531, 532,
		5, 66, 0, 0, 532, 67, 1, 0, 0, 0, 533, 534, 5, 36, 0, 0, 534, 535, 5, 59,
		0, 0, 535, 537, 3, 52, 26, 0, 536, 538, 5, 68, 0, 0, 537, 536, 1, 0, 0,
		0, 537, 538, 1, 0, 0, 0, 538, 69, 1, 0, 0, 0, 539, 550, 5, 6, 0, 0, 540,
		550, 5, 7, 0, 0, 541, 550, 5, 8, 0, 0, 542, 550, 5, 9, 0, 0, 543, 550,
		5, 10, 0, 0, 544, 550, 5, 11, 0, 0, 545, 550, 5, 12, 0, 0, 546, 547, 5,
		63, 0, 0, 547, 548, 5, 65, 0, 0, 548, 550, 3, 70, 35, 0, 549, 539, 1, 0,
		0, 0, 549, 540, 1, 0, 0, 0, 549, 541, 1, 0, 0, 0, 549, 542, 1, 0, 0, 0,
		549, 543, 1, 0, 0, 0, 549, 544, 1, 0, 0, 0, 549, 545, 1, 0, 0, 0, 549,
		546, 1, 0, 0, 0, 550, 71, 1, 0, 0, 0, 551, 552, 5, 63, 0, 0, 552, 553,
		5, 36, 0, 0, 553, 554, 5, 65, 0, 0, 554, 555, 5, 61, 0, 0, 555, 556, 5,
		62, 0, 0, 556, 73, 1, 0, 0, 0, 557, 558, 5, 36, 0, 0, 558, 562, 5, 52,
		0, 0, 559, 560, 5, 36, 0, 0, 560, 562, 5, 53, 0, 0, 561, 557, 1, 0, 0,
		0, 561, 559, 1, 0, 0, 0, 562, 75, 1, 0, 0, 0, 53, 79, 97, 104, 108, 116,
		126, 138, 150, 161, 171, 177, 188, 195, 199, 203, 209, 223, 233, 248, 253,
		257, 261, 265, 271, 275, 285, 293, 301, 313, 321, 325, 331, 342, 352, 356,
		360, 363, 389, 407, 418, 421, 430, 447, 465, 485, 487, 495, 505, 520, 528,
		537, 549, 561,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VlangParserInit initializes any static state used to implement VlangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVlangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.once.Do(vlangParserInit)
}

// NewVlangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVlangParser(input antlr.TokenStream) *VlangParser {
	VlangParserInit()
	this := new(VlangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VlangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Vlang.g4"

	return this
}

// VlangParser tokens.
const (
	VlangParserEOF                  = antlr.TokenEOF
	VlangParserT__0                 = 1
	VlangParserT__1                 = 2
	VlangParserT__2                 = 3
	VlangParserT__3                 = 4
	VlangParserT__4                 = 5
	VlangParserT__5                 = 6
	VlangParserT__6                 = 7
	VlangParserT__7                 = 8
	VlangParserT__8                 = 9
	VlangParserT__9                 = 10
	VlangParserT__10                = 11
	VlangParserT__11                = 12
	VlangParserLEN                  = 13
	VlangParserCAP                  = 14
	VlangParserAPPEND               = 15
	VlangParserMUT                  = 16
	VlangParserBOOLEANO             = 17
	VlangParserENTERO               = 18
	VlangParserFLOAT                = 19
	VlangParserDECIMAL              = 20
	VlangParserCADENA               = 21
	VlangParserCARACTER             = 22
	VlangParserSTRING_INTERPOLATION = 23
	VlangParserIF_KW                = 24
	VlangParserELSE_KW              = 25
	VlangParserWHILE_KW             = 26
	VlangParserFOR_KW               = 27
	VlangParserIN_KW                = 28
	VlangParserRETURN_KW            = 29
	VlangParserBREAK_KW             = 30
	VlangParserCONTINUE_KW          = 31
	VlangParserSWITCH_KW            = 32
	VlangParserCASE_KW              = 33
	VlangParserDEFAULT_KW           = 34
	VlangParserFALLTHROUGH_KW       = 35
	VlangParserID                   = 36
	VlangParserPLUS                 = 37
	VlangParserMINUS                = 38
	VlangParserMUL                  = 39
	VlangParserDIV                  = 40
	VlangParserMOD                  = 41
	VlangParserNOT                  = 42
	VlangParserOR                   = 43
	VlangParserAND                  = 44
	VlangParserEQ                   = 45
	VlangParserNEQ                  = 46
	VlangParserLT                   = 47
	VlangParserLE                   = 48
	VlangParserGT                   = 49
	VlangParserGE                   = 50
	VlangParserASSIGN               = 51
	VlangParserINC                  = 52
	VlangParserDEC                  = 53
	VlangParserPLUS_ASSIGN          = 54
	VlangParserMINUS_ASSIGN         = 55
	VlangParserMUL_ASSIGN           = 56
	VlangParserDIV_ASSIGN           = 57
	VlangParserMOD_ASSIGN           = 58
	VlangParserCOLON                = 59
	VlangParserSEMICOLON            = 60
	VlangParserLPAREN               = 61
	VlangParserRPAREN               = 62
	VlangParserLBRACK               = 63
	VlangParserLCOR                 = 64
	VlangParserRBRACK               = 65
	VlangParserRCOR                 = 66
	VlangParserDOT                  = 67
	VlangParserCOMMA                = 68
	VlangParserWS                   = 69
	VlangParserLINE_COMMENT         = 70
	VlangParserBLOCK_COMMENT        = 71
)

// VlangParser rules.
const (
	VlangParserRULE_programa          = 0
	VlangParserRULE_stmt              = 1
	VlangParserRULE_if_stmt           = 2
	VlangParserRULE_if_chain          = 3
	VlangParserRULE_else_stmt         = 4
	VlangParserRULE_println           = 5
	VlangParserRULE_print             = 6
	VlangParserRULE_while_stmt        = 7
	VlangParserRULE_for_stmt          = 8
	VlangParserRULE_for_enhanced_stmt = 9
	VlangParserRULE_for_init          = 10
	VlangParserRULE_for_update        = 11
	VlangParserRULE_switch_stmt       = 12
	VlangParserRULE_switch_case       = 13
	VlangParserRULE_case_values       = 14
	VlangParserRULE_default_case      = 15
	VlangParserRULE_range             = 16
	VlangParserRULE_func_call         = 17
	VlangParserRULE_func_dcl          = 18
	VlangParserRULE_struct_dcl        = 19
	VlangParserRULE_struct_field      = 20
	VlangParserRULE_transfer_stmt     = 21
	VlangParserRULE_assign_stmt       = 22
	VlangParserRULE_decl_stmt         = 23
	VlangParserRULE_array_literal     = 24
	VlangParserRULE_id_pattern        = 25
	VlangParserRULE_expresion         = 26
	VlangParserRULE_parametros        = 27
	VlangParserRULE_func_param        = 28
	VlangParserRULE_arg_list          = 29
	VlangParserRULE_func_arg          = 30
	VlangParserRULE_valores           = 31
	VlangParserRULE_valor             = 32
	VlangParserRULE_strct_instancia   = 33
	VlangParserRULE_struct_prop       = 34
	VlangParserRULE_var_type          = 35
	VlangParserRULE_struct_vector     = 36
	VlangParserRULE_incredecre        = 37
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(VlangParserEOF, 0)
}

func (s *ProgramaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VlangParserRULE_programa)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0) {
		{
			p.SetState(76)
			p.Stmt()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(VlangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Assign_stmt() IAssign_stmtContext
	If_stmt() IIf_stmtContext
	While_stmt() IWhile_stmtContext
	For_stmt() IFor_stmtContext
	For_enhanced_stmt() IFor_enhanced_stmtContext
	Switch_stmt() ISwitch_stmtContext
	Func_call() IFunc_callContext
	Func_dcl() IFunc_dclContext
	Struct_dcl() IStruct_dclContext
	Strct_instancia() IStrct_instanciaContext
	Transfer_stmt() ITransfer_stmtContext
	Println_() IPrintlnContext
	Print_() IPrintContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Decl_stmt() IDecl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_stmtContext)
}

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) For_enhanced_stmt() IFor_enhanced_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_enhanced_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_enhanced_stmtContext)
}

func (s *StmtContext) Switch_stmt() ISwitch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StmtContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *StmtContext) Func_dcl() IFunc_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_dclContext)
}

func (s *StmtContext) Struct_dcl() IStruct_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_dclContext)
}

func (s *StmtContext) Strct_instancia() IStrct_instanciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrct_instanciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrct_instanciaContext)
}

func (s *StmtContext) Transfer_stmt() ITransfer_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransfer_stmtContext)
}

func (s *StmtContext) Println_() IPrintlnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintlnContext)
}

func (s *StmtContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VlangParserRULE_stmt)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Decl_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Assign_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.If_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(86)
			p.While_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(87)
			p.For_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(88)
			p.For_enhanced_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(89)
			p.Switch_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(90)
			p.Func_call()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(91)
			p.Func_dcl()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(92)
			p.Struct_dcl()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(93)
			p.Strct_instancia()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(94)
			p.Transfer_stmt()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(95)
			p.Println_()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(96)
			p.Print_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) CopyAll(ctx *If_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStmtContext struct {
	If_stmtContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	InitEmptyIf_stmtContext(&p.If_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_stmtContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) AllIf_chain() []IIf_chainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_chainContext); ok {
			len++
		}
	}

	tst := make([]IIf_chainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_chainContext); ok {
			tst[i] = t.(IIf_chainContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) If_chain(i int) IIf_chainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_chainContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_chainContext)
}

func (s *IfStmtContext) AllELSE_KW() []antlr.TerminalNode {
	return s.GetTokens(VlangParserELSE_KW)
}

func (s *IfStmtContext) ELSE_KW(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserELSE_KW, i)
}

func (s *IfStmtContext) Else_stmt() IElse_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_stmtContext)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VlangParserRULE_if_stmt)
	var _la int

	var _alt int

	localctx = NewIfStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.If_chain()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(100)
				p.Match(VlangParserELSE_KW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(101)
				p.If_chain()
			}

		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserELSE_KW {
		{
			p.SetState(107)
			p.Else_stmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_chainContext is an interface to support dynamic dispatch.
type IIf_chainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_chainContext differentiates from other interfaces.
	IsIf_chainContext()
}

type If_chainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_chainContext() *If_chainContext {
	var p = new(If_chainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_if_chain
	return p
}

func InitEmptyIf_chainContext(p *If_chainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_if_chain
}

func (*If_chainContext) IsIf_chainContext() {}

func NewIf_chainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_chainContext {
	var p = new(If_chainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_if_chain

	return p
}

func (s *If_chainContext) GetParser() antlr.Parser { return s.parser }

func (s *If_chainContext) CopyAll(ctx *If_chainContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_chainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_chainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfChainContext struct {
	If_chainContext
}

func NewIfChainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfChainContext {
	var p = new(IfChainContext)

	InitEmptyIf_chainContext(&p.If_chainContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_chainContext))

	return p
}

func (s *IfChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfChainContext) IF_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserIF_KW, 0)
}

func (s *IfChainContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IfChainContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *IfChainContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *IfChainContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfChainContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *IfChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIfChain(s)
	}
}

func (s *IfChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIfChain(s)
	}
}

func (s *IfChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIfChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) If_chain() (localctx IIf_chainContext) {
	localctx = NewIf_chainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VlangParserRULE_if_chain)
	var _la int

	localctx = NewIfChainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(VlangParserIF_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.expresion(0)
	}
	{
		p.SetState(112)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
		{
			p.SetState(113)
			p.Stmt()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(119)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_stmtContext is an interface to support dynamic dispatch.
type IElse_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElse_stmtContext differentiates from other interfaces.
	IsElse_stmtContext()
}

type Else_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_stmtContext() *Else_stmtContext {
	var p = new(Else_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_else_stmt
	return p
}

func InitEmptyElse_stmtContext(p *Else_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_else_stmt
}

func (*Else_stmtContext) IsElse_stmtContext() {}

func NewElse_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_stmtContext {
	var p = new(Else_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_else_stmt

	return p
}

func (s *Else_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_stmtContext) CopyAll(ctx *Else_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Else_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseStmtContext struct {
	Else_stmtContext
}

func NewElseStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseStmtContext {
	var p = new(ElseStmtContext)

	InitEmptyElse_stmtContext(&p.Else_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_stmtContext))

	return p
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ELSE_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserELSE_KW, 0)
}

func (s *ElseStmtContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ElseStmtContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ElseStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElseStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Else_stmt() (localctx IElse_stmtContext) {
	localctx = NewElse_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VlangParserRULE_else_stmt)
	var _la int

	localctx = NewElseStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(VlangParserELSE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
		{
			p.SetState(123)
			p.Stmt()
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintlnContext is an interface to support dynamic dispatch.
type IPrintlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrintlnContext differentiates from other interfaces.
	IsPrintlnContext()
}

type PrintlnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintlnContext() *PrintlnContext {
	var p = new(PrintlnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_println
	return p
}

func InitEmptyPrintlnContext(p *PrintlnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_println
}

func (*PrintlnContext) IsPrintlnContext() {}

func NewPrintlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnContext {
	var p = new(PrintlnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_println

	return p
}

func (s *PrintlnContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnContext) CopyAll(ctx *PrintlnContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrintlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintlnStmtContext struct {
	PrintlnContext
}

func NewPrintlnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnStmtContext {
	var p = new(PrintlnStmtContext)

	InitEmptyPrintlnContext(&p.PrintlnContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrintlnContext))

	return p
}

func (s *PrintlnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *PrintlnStmtContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *PrintlnStmtContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *PrintlnStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *PrintlnStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *PrintlnStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *PrintlnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrintlnStmt(s)
	}
}

func (s *PrintlnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrintlnStmt(s)
	}
}

func (s *PrintlnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrintlnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Println_() (localctx IPrintlnContext) {
	localctx = NewPrintlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VlangParserRULE_println)
	var _la int

	localctx = NewPrintlnStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(VlangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.expresion(0)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(134)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.expresion(0)
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) CopyAll(ctx *PrintContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintStmtContext struct {
	PrintContext
}

func NewPrintStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStmtContext {
	var p = new(PrintStmtContext)

	InitEmptyPrintContext(&p.PrintContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrintContext))

	return p
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *PrintStmtContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *PrintStmtContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *PrintStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *PrintStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *PrintStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VlangParserRULE_print)
	var _la int

	localctx = NewPrintStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(VlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.expresion(0)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(146)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.expresion(0)
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) CopyAll(ctx *While_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStmtContext struct {
	While_stmtContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	InitEmptyWhile_stmtContext(&p.While_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_stmtContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) WHILE_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserWHILE_KW, 0)
}

func (s *WhileStmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *WhileStmtContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *WhileStmtContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *WhileStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *WhileStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VlangParserRULE_while_stmt)
	var _la int

	localctx = NewWhileStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(VlangParserWHILE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.expresion(0)
	}
	{
		p.SetState(157)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
		{
			p.SetState(158)
			p.Stmt()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) CopyAll(ctx *For_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForStmtContext struct {
	For_stmtContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR_KW, 0)
}

func (s *ForStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ForStmtContext) IN_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserIN_KW, 0)
}

func (s *ForStmtContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ForStmtContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ForStmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForStmtContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VlangParserRULE_for_stmt)
	var _la int

	localctx = NewForStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(VlangParserFOR_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(VlangParserIN_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(169)
			p.expresion(0)
		}

	case 2:
		{
			p.SetState(170)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(173)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
		{
			p.SetState(174)
			p.Stmt()
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(180)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_enhanced_stmtContext is an interface to support dynamic dispatch.
type IFor_enhanced_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_enhanced_stmtContext differentiates from other interfaces.
	IsFor_enhanced_stmtContext()
}

type For_enhanced_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_enhanced_stmtContext() *For_enhanced_stmtContext {
	var p = new(For_enhanced_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_enhanced_stmt
	return p
}

func InitEmptyFor_enhanced_stmtContext(p *For_enhanced_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_enhanced_stmt
}

func (*For_enhanced_stmtContext) IsFor_enhanced_stmtContext() {}

func NewFor_enhanced_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_enhanced_stmtContext {
	var p = new(For_enhanced_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_for_enhanced_stmt

	return p
}

func (s *For_enhanced_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_enhanced_stmtContext) CopyAll(ctx *For_enhanced_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_enhanced_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_enhanced_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForClassicContext struct {
	For_enhanced_stmtContext
}

func NewForClassicContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForClassicContext {
	var p = new(ForClassicContext)

	InitEmptyFor_enhanced_stmtContext(&p.For_enhanced_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_enhanced_stmtContext))

	return p
}

func (s *ForClassicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClassicContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR_KW, 0)
}

func (s *ForClassicContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(VlangParserSEMICOLON)
}

func (s *ForClassicContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserSEMICOLON, i)
}

func (s *ForClassicContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ForClassicContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ForClassicContext) For_init() IFor_initContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_initContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_initContext)
}

func (s *ForClassicContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForClassicContext) For_update() IFor_updateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_updateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_updateContext)
}

func (s *ForClassicContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForClassicContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForClassicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForClassic(s)
	}
}

func (s *ForClassicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForClassic(s)
	}
}

func (s *ForClassicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForClassic(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForInfiniteContext struct {
	For_enhanced_stmtContext
}

func NewForInfiniteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInfiniteContext {
	var p = new(ForInfiniteContext)

	InitEmptyFor_enhanced_stmtContext(&p.For_enhanced_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_enhanced_stmtContext))

	return p
}

func (s *ForInfiniteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInfiniteContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR_KW, 0)
}

func (s *ForInfiniteContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ForInfiniteContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ForInfiniteContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForInfiniteContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForInfiniteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForInfinite(s)
	}
}

func (s *ForInfiniteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForInfinite(s)
	}
}

func (s *ForInfiniteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForInfinite(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForRangeContext struct {
	For_enhanced_stmtContext
}

func NewForRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeContext {
	var p = new(ForRangeContext)

	InitEmptyFor_enhanced_stmtContext(&p.For_enhanced_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_enhanced_stmtContext))

	return p
}

func (s *ForRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR_KW, 0)
}

func (s *ForRangeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *ForRangeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *ForRangeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *ForRangeContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *ForRangeContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForRangeContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ForRangeContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ForRangeContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForRangeContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForRange(s)
	}
}

func (s *ForRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForRange(s)
	}
}

func (s *ForRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForRange(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForIndexValueContext struct {
	For_enhanced_stmtContext
}

func NewForIndexValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForIndexValueContext {
	var p = new(ForIndexValueContext)

	InitEmptyFor_enhanced_stmtContext(&p.For_enhanced_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_enhanced_stmtContext))

	return p
}

func (s *ForIndexValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForIndexValueContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR_KW, 0)
}

func (s *ForIndexValueContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *ForIndexValueContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *ForIndexValueContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *ForIndexValueContext) IN_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserIN_KW, 0)
}

func (s *ForIndexValueContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForIndexValueContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ForIndexValueContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ForIndexValueContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForIndexValueContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForIndexValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForIndexValue(s)
	}
}

func (s *ForIndexValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForIndexValue(s)
	}
}

func (s *ForIndexValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForIndexValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForConditionContext struct {
	For_enhanced_stmtContext
}

func NewForConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForConditionContext {
	var p = new(ForConditionContext)

	InitEmptyFor_enhanced_stmtContext(&p.For_enhanced_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_enhanced_stmtContext))

	return p
}

func (s *ForConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForConditionContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR_KW, 0)
}

func (s *ForConditionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForConditionContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ForConditionContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ForConditionContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForConditionContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForCondition(s)
	}
}

func (s *ForConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForCondition(s)
	}
}

func (s *ForConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) For_enhanced_stmt() (localctx IFor_enhanced_stmtContext) {
	localctx = NewFor_enhanced_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VlangParserRULE_for_enhanced_stmt)
	var _la int

	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Match(VlangParserFOR_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.expresion(0)
		}
		{
			p.SetState(184)
			p.Match(VlangParserLCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
			{
				p.SetState(185)
				p.Stmt()
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(191)
			p.Match(VlangParserRCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewForClassicContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(VlangParserFOR_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VlangParserMUT || _la == VlangParserID {
			{
				p.SetState(194)
				p.For_init()
			}

		}
		{
			p.SetState(197)
			p.Match(VlangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6917524285980540928) != 0 {
			{
				p.SetState(198)
				p.expresion(0)
			}

		}
		{
			p.SetState(201)
			p.Match(VlangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VlangParserID {
			{
				p.SetState(202)
				p.For_update()
			}

		}
		{
			p.SetState(205)
			p.Match(VlangParserLCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
			{
				p.SetState(206)
				p.Stmt()
			}

			p.SetState(211)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(212)
			p.Match(VlangParserRCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewForIndexValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(213)
			p.Match(VlangParserFOR_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(VlangParserIN_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.expresion(0)
		}
		{
			p.SetState(219)
			p.Match(VlangParserLCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
			{
				p.SetState(220)
				p.Stmt()
			}

			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(226)
			p.Match(VlangParserRCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewForInfiniteContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(228)
			p.Match(VlangParserFOR_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(VlangParserLCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
			{
				p.SetState(230)
				p.Stmt()
			}

			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(236)
			p.Match(VlangParserRCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(237)
			p.Match(VlangParserFOR_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(VlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.expresion(0)
		}
		{
			p.SetState(244)
			p.Match(VlangParserLCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
			{
				p.SetState(245)
				p.Stmt()
			}

			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(251)
			p.Match(VlangParserRCOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_initContext is an interface to support dynamic dispatch.
type IFor_initContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Assign_stmt() IAssign_stmtContext

	// IsFor_initContext differentiates from other interfaces.
	IsFor_initContext()
}

type For_initContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_initContext() *For_initContext {
	var p = new(For_initContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_init
	return p
}

func InitEmptyFor_initContext(p *For_initContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_init
}

func (*For_initContext) IsFor_initContext() {}

func NewFor_initContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_initContext {
	var p = new(For_initContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_for_init

	return p
}

func (s *For_initContext) GetParser() antlr.Parser { return s.parser }

func (s *For_initContext) Decl_stmt() IDecl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_stmtContext)
}

func (s *For_initContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *For_initContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_initContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_initContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFor_init(s)
	}
}

func (s *For_initContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFor_init(s)
	}
}

func (s *For_initContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFor_init(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) For_init() (localctx IFor_initContext) {
	localctx = NewFor_initContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VlangParserRULE_for_init)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(255)
			p.Decl_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.Assign_stmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_updateContext is an interface to support dynamic dispatch.
type IFor_updateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign_stmt() IAssign_stmtContext
	Incredecre() IIncredecreContext

	// IsFor_updateContext differentiates from other interfaces.
	IsFor_updateContext()
}

type For_updateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_updateContext() *For_updateContext {
	var p = new(For_updateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_update
	return p
}

func InitEmptyFor_updateContext(p *For_updateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_update
}

func (*For_updateContext) IsFor_updateContext() {}

func NewFor_updateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_updateContext {
	var p = new(For_updateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_for_update

	return p
}

func (s *For_updateContext) GetParser() antlr.Parser { return s.parser }

func (s *For_updateContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *For_updateContext) Incredecre() IIncredecreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncredecreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncredecreContext)
}

func (s *For_updateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_updateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_updateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFor_update(s)
	}
}

func (s *For_updateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFor_update(s)
	}
}

func (s *For_updateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFor_update(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) For_update() (localctx IFor_updateContext) {
	localctx = NewFor_updateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VlangParserRULE_for_update)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Assign_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Incredecre()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_switch_stmt
	return p
}

func InitEmptySwitch_stmtContext(p *Switch_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_switch_stmt
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) CopyAll(ctx *Switch_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStmtContext struct {
	Switch_stmtContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	InitEmptySwitch_stmtContext(&p.Switch_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_stmtContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) SWITCH_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserSWITCH_KW, 0)
}

func (s *SwitchStmtContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *SwitchStmtContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *SwitchStmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SwitchStmtContext) AllSwitch_case() []ISwitch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			len++
		}
	}

	tst := make([]ISwitch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitch_caseContext); ok {
			tst[i] = t.(ISwitch_caseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Switch_case(i int) ISwitch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_caseContext)
}

func (s *SwitchStmtContext) Default_case() IDefault_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_caseContext)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VlangParserRULE_switch_stmt)
	var _la int

	localctx = NewSwitchStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(VlangParserSWITCH_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6917524285980540928) != 0 {
		{
			p.SetState(264)
			p.expresion(0)
		}

	}
	{
		p.SetState(267)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCASE_KW {
		{
			p.SetState(268)
			p.Switch_case()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserDEFAULT_KW {
		{
			p.SetState(274)
			p.Default_case()
		}

	}
	{
		p.SetState(277)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) CopyAll(ctx *Switch_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchCaseContext struct {
	Switch_caseContext
}

func NewSwitchCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	InitEmptySwitch_caseContext(&p.Switch_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_caseContext))

	return p
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) CASE_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserCASE_KW, 0)
}

func (s *SwitchCaseContext) Case_values() ICase_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICase_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICase_valuesContext)
}

func (s *SwitchCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *SwitchCaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *SwitchCaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VlangParserRULE_switch_case)
	var _la int

	localctx = NewSwitchCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(VlangParserCASE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.Case_values()
	}
	{
		p.SetState(281)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
		{
			p.SetState(282)
			p.Stmt()
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICase_valuesContext is an interface to support dynamic dispatch.
type ICase_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCase_valuesContext differentiates from other interfaces.
	IsCase_valuesContext()
}

type Case_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_valuesContext() *Case_valuesContext {
	var p = new(Case_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_case_values
	return p
}

func InitEmptyCase_valuesContext(p *Case_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_case_values
}

func (*Case_valuesContext) IsCase_valuesContext() {}

func NewCase_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_valuesContext {
	var p = new(Case_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_case_values

	return p
}

func (s *Case_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_valuesContext) CopyAll(ctx *Case_valuesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Case_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CaseValueListContext struct {
	Case_valuesContext
}

func NewCaseValueListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseValueListContext {
	var p = new(CaseValueListContext)

	InitEmptyCase_valuesContext(&p.Case_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Case_valuesContext))

	return p
}

func (s *CaseValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseValueListContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *CaseValueListContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CaseValueListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *CaseValueListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *CaseValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCaseValueList(s)
	}
}

func (s *CaseValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCaseValueList(s)
	}
}

func (s *CaseValueListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCaseValueList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Case_values() (localctx ICase_valuesContext) {
	localctx = NewCase_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VlangParserRULE_case_values)
	var _la int

	localctx = NewCaseValueListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.expresion(0)
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(289)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.expresion(0)
		}

		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefault_caseContext is an interface to support dynamic dispatch.
type IDefault_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefault_caseContext differentiates from other interfaces.
	IsDefault_caseContext()
}

type Default_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_caseContext() *Default_caseContext {
	var p = new(Default_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_default_case
	return p
}

func InitEmptyDefault_caseContext(p *Default_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_default_case
}

func (*Default_caseContext) IsDefault_caseContext() {}

func NewDefault_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_caseContext {
	var p = new(Default_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_default_case

	return p
}

func (s *Default_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_caseContext) CopyAll(ctx *Default_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultCaseContext struct {
	Default_caseContext
}

func NewDefaultCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	InitEmptyDefault_caseContext(&p.Default_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Default_caseContext))

	return p
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) DEFAULT_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserDEFAULT_KW, 0)
}

func (s *DefaultCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *DefaultCaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *DefaultCaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VlangParserRULE_default_case)
	var _la int

	localctx = NewDefaultCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(VlangParserDEFAULT_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
		{
			p.SetState(298)
			p.Stmt()
		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) CopyAll(ctx *RangeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumericRangeContext struct {
	RangeContext
}

func NewNumericRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericRangeContext {
	var p = new(NumericRangeContext)

	InitEmptyRangeContext(&p.RangeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RangeContext))

	return p
}

func (s *NumericRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericRangeContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *NumericRangeContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *NumericRangeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(VlangParserDOT)
}

func (s *NumericRangeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, i)
}

func (s *NumericRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterNumericRange(s)
	}
}

func (s *NumericRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitNumericRange(s)
	}
}

func (s *NumericRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitNumericRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VlangParserRULE_range)
	localctx = NewNumericRangeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.expresion(0)
	}
	{
		p.SetState(305)
		p.Match(VlangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(VlangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(VlangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) CopyAll(ctx *Func_callContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncCallContext struct {
	Func_callContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	InitEmptyFunc_callContext(&p.Func_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_callContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *FuncCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *FuncCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *FuncCallContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VlangParserRULE_func_call)
	var _la int

	localctx = NewFuncCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Id_pattern()
	}
	{
		p.SetState(311)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6917524285980540928) != 0 {
		{
			p.SetState(312)
			p.Parametros()
		}

	}
	{
		p.SetState(315)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_dclContext is an interface to support dynamic dispatch.
type IFunc_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_dclContext differentiates from other interfaces.
	IsFunc_dclContext()
}

type Func_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_dclContext() *Func_dclContext {
	var p = new(Func_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_dcl
	return p
}

func InitEmptyFunc_dclContext(p *Func_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_dcl
}

func (*Func_dclContext) IsFunc_dclContext() {}

func NewFunc_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_dclContext {
	var p = new(Func_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_func_dcl

	return p
}

func (s *Func_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_dclContext) CopyAll(ctx *Func_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncDeclContext struct {
	Func_dclContext
}

func NewFuncDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncDeclContext {
	var p = new(FuncDeclContext)

	InitEmptyFunc_dclContext(&p.Func_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_dclContext))

	return p
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *FuncDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *FuncDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *FuncDeclContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *FuncDeclContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *FuncDeclContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *FuncDeclContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *FuncDeclContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncDeclContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Func_dcl() (localctx IFunc_dclContext) {
	localctx = NewFunc_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VlangParserRULE_func_dcl)
	var _la int

	localctx = NewFuncDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(VlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserID {
		{
			p.SetState(320)
			p.Arg_list()
		}

	}
	{
		p.SetState(323)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9223372036854767680) != 0 {
		{
			p.SetState(324)
			p.Var_type()
		}

	}
	{
		p.SetState(327)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&111350448182) != 0 {
		{
			p.SetState(328)
			p.Stmt()
		}

		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(334)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_dclContext is an interface to support dynamic dispatch.
type IStruct_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_dclContext differentiates from other interfaces.
	IsStruct_dclContext()
}

type Struct_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_dclContext() *Struct_dclContext {
	var p = new(Struct_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_dcl
	return p
}

func InitEmptyStruct_dclContext(p *Struct_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_dcl
}

func (*Struct_dclContext) IsStruct_dclContext() {}

func NewStruct_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_dclContext {
	var p = new(Struct_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_struct_dcl

	return p
}

func (s *Struct_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_dclContext) CopyAll(ctx *Struct_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclContext struct {
	Struct_dclContext
}

func NewStructDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclContext {
	var p = new(StructDeclContext)

	InitEmptyStruct_dclContext(&p.Struct_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_dclContext))

	return p
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructDeclContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *StructDeclContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *StructDeclContext) AllStruct_field() []IStruct_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_fieldContext); ok {
			len++
		}
	}

	tst := make([]IStruct_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_fieldContext); ok {
			tst[i] = t.(IStruct_fieldContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclContext) Struct_field(i int) IStruct_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_fieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_fieldContext)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Struct_dcl() (localctx IStruct_dclContext) {
	localctx = NewStruct_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VlangParserRULE_struct_dcl)
	var _la int

	localctx = NewStructDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		p.Match(VlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9223372036854767680) != 0 {
		{
			p.SetState(339)
			p.Struct_field()
		}

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(345)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_fieldContext is an interface to support dynamic dispatch.
type IStruct_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_fieldContext differentiates from other interfaces.
	IsStruct_fieldContext()
}

type Struct_fieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_fieldContext() *Struct_fieldContext {
	var p = new(Struct_fieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_field
	return p
}

func InitEmptyStruct_fieldContext(p *Struct_fieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_field
}

func (*Struct_fieldContext) IsStruct_fieldContext() {}

func NewStruct_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_fieldContext {
	var p = new(Struct_fieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_struct_field

	return p
}

func (s *Struct_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_fieldContext) CopyAll(ctx *Struct_fieldContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructFieldContext struct {
	Struct_fieldContext
}

func NewStructFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructFieldContext {
	var p = new(StructFieldContext)

	InitEmptyStruct_fieldContext(&p.Struct_fieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_fieldContext))

	return p
}

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *StructFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructField(s)
	}
}

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructField(s)
	}
}

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Struct_field() (localctx IStruct_fieldContext) {
	localctx = NewStruct_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VlangParserRULE_struct_field)
	localctx = NewStructFieldContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Var_type()
	}
	{
		p.SetState(348)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransfer_stmtContext is an interface to support dynamic dispatch.
type ITransfer_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransfer_stmtContext differentiates from other interfaces.
	IsTransfer_stmtContext()
}

type Transfer_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransfer_stmtContext() *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_transfer_stmt
	return p
}

func InitEmptyTransfer_stmtContext(p *Transfer_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_transfer_stmt
}

func (*Transfer_stmtContext) IsTransfer_stmtContext() {}

func NewTransfer_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_transfer_stmt

	return p
}

func (s *Transfer_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Transfer_stmtContext) CopyAll(ctx *Transfer_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Transfer_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transfer_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	Transfer_stmtContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) CONTINUE_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserCONTINUE_KW, 0)
}

func (s *ContinueStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type FallthroughStmtContext struct {
	Transfer_stmtContext
}

func NewFallthroughStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FallthroughStmtContext {
	var p = new(FallthroughStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *FallthroughStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FallthroughStmtContext) FALLTHROUGH_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserFALLTHROUGH_KW, 0)
}

func (s *FallthroughStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFallthroughStmt(s)
	}
}

func (s *FallthroughStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFallthroughStmt(s)
	}
}

func (s *FallthroughStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFallthroughStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	Transfer_stmtContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) BREAK_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserBREAK_KW, 0)
}

func (s *BreakStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	Transfer_stmtContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RETURN_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserRETURN_KW, 0)
}

func (s *ReturnStmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Transfer_stmt() (localctx ITransfer_stmtContext) {
	localctx = NewTransfer_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VlangParserRULE_transfer_stmt)
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserRETURN_KW:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.Match(VlangParserRETURN_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(351)
				p.expresion(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case VlangParserBREAK_KW:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(VlangParserBREAK_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(355)
				p.Match(VlangParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case VlangParserCONTINUE_KW:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(358)
			p.Match(VlangParserCONTINUE_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(359)
				p.Match(VlangParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case VlangParserFALLTHROUGH_KW:
		localctx = NewFallthroughStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(362)
			p.Match(VlangParserFALLTHROUGH_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_assign_stmt
	return p
}

func InitEmptyAssign_stmtContext(p *Assign_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_assign_stmt
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) CopyAll(ctx *Assign_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MinusAssignContext struct {
	Assign_stmtContext
}

func NewMinusAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusAssignContext {
	var p = new(MinusAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *MinusAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *MinusAssignContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS_ASSIGN, 0)
}

func (s *MinusAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *MinusAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterMinusAssign(s)
	}
}

func (s *MinusAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitMinusAssign(s)
	}
}

func (s *MinusAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitMinusAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusAssignContext struct {
	Assign_stmtContext
}

func NewPlusAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusAssignContext {
	var p = new(PlusAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *PlusAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *PlusAssignContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserPLUS_ASSIGN, 0)
}

func (s *PlusAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *PlusAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPlusAssign(s)
	}
}

func (s *PlusAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPlusAssign(s)
	}
}

func (s *PlusAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPlusAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type DirectAssignContext struct {
	Assign_stmtContext
}

func NewDirectAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DirectAssignContext {
	var p = new(DirectAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *DirectAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *DirectAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *DirectAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DirectAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDirectAssign(s)
	}
}

func (s *DirectAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDirectAssign(s)
	}
}

func (s *DirectAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDirectAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulAssignContext struct {
	Assign_stmtContext
}

func NewMulAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulAssignContext {
	var p = new(MulAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *MulAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *MulAssignContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserMUL_ASSIGN, 0)
}

func (s *MulAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *MulAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterMulAssign(s)
	}
}

func (s *MulAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitMulAssign(s)
	}
}

func (s *MulAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitMulAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivAssignContext struct {
	Assign_stmtContext
}

func NewDivAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivAssignContext {
	var p = new(DivAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *DivAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *DivAssignContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserDIV_ASSIGN, 0)
}

func (s *DivAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DivAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDivAssign(s)
	}
}

func (s *DivAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDivAssign(s)
	}
}

func (s *DivAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDivAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModAssignContext struct {
	Assign_stmtContext
}

func NewModAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModAssignContext {
	var p = new(ModAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *ModAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *ModAssignContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserMOD_ASSIGN, 0)
}

func (s *ModAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ModAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterModAssign(s)
	}
}

func (s *ModAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitModAssign(s)
	}
}

func (s *ModAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitModAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VlangParserRULE_assign_stmt)
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDirectAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(365)
			p.Id_pattern()
		}
		{
			p.SetState(366)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.expresion(0)
		}

	case 2:
		localctx = NewPlusAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(369)
			p.Id_pattern()
		}
		{
			p.SetState(370)
			p.Match(VlangParserPLUS_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.expresion(0)
		}

	case 3:
		localctx = NewMinusAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.Id_pattern()
		}
		{
			p.SetState(374)
			p.Match(VlangParserMINUS_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.expresion(0)
		}

	case 4:
		localctx = NewMulAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(377)
			p.Id_pattern()
		}
		{
			p.SetState(378)
			p.Match(VlangParserMUL_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.expresion(0)
		}

	case 5:
		localctx = NewDivAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(381)
			p.Id_pattern()
		}
		{
			p.SetState(382)
			p.Match(VlangParserDIV_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.expresion(0)
		}

	case 6:
		localctx = NewModAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(385)
			p.Id_pattern()
		}
		{
			p.SetState(386)
			p.Match(VlangParserMOD_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecl_stmtContext is an interface to support dynamic dispatch.
type IDecl_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDecl_stmtContext differentiates from other interfaces.
	IsDecl_stmtContext()
}

type Decl_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_stmtContext() *Decl_stmtContext {
	var p = new(Decl_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_decl_stmt
	return p
}

func InitEmptyDecl_stmtContext(p *Decl_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_decl_stmt
}

func (*Decl_stmtContext) IsDecl_stmtContext() {}

func NewDecl_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_stmtContext {
	var p = new(Decl_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_decl_stmt

	return p
}

func (s *Decl_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_stmtContext) CopyAll(ctx *Decl_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Decl_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayAssignContext struct {
	Decl_stmtContext
}

func NewArrayAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAssignContext {
	var p = new(ArrayAssignContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *ArrayAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ArrayAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *ArrayAssignContext) Array_literal() IArray_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_literalContext)
}

func (s *ArrayAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterArrayAssign(s)
	}
}

func (s *ArrayAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitArrayAssign(s)
	}
}

func (s *ArrayAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitArrayAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclAssignContext struct {
	Decl_stmtContext
}

func NewDeclAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAssignContext {
	var p = new(DeclAssignContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclAssignContext) MUT() antlr.TerminalNode {
	return s.GetToken(VlangParserMUT, 0)
}

func (s *DeclAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *DeclAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *DeclAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDeclAssign(s)
	}
}

func (s *DeclAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDeclAssign(s)
	}
}

func (s *DeclAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDeclAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclAssignTypeContext struct {
	Decl_stmtContext
}

func NewDeclAssignTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAssignTypeContext {
	var p = new(DeclAssignTypeContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclAssignTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclAssignTypeContext) MUT() antlr.TerminalNode {
	return s.GetToken(VlangParserMUT, 0)
}

func (s *DeclAssignTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *DeclAssignTypeContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *DeclAssignTypeContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *DeclAssignTypeContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclAssignTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDeclAssignType(s)
	}
}

func (s *DeclAssignTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDeclAssignType(s)
	}
}

func (s *DeclAssignTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDeclAssignType(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclTypeContext struct {
	Decl_stmtContext
}

func NewDeclTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeContext {
	var p = new(DeclTypeContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeContext) MUT() antlr.TerminalNode {
	return s.GetToken(VlangParserMUT, 0)
}

func (s *DeclTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *DeclTypeContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *DeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDeclType(s)
	}
}

func (s *DeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDeclType(s)
	}
}

func (s *DeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Decl_stmt() (localctx IDecl_stmtContext) {
	localctx = NewDecl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VlangParserRULE_decl_stmt)
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.Match(VlangParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.expresion(0)
		}

	case 2:
		localctx = NewDeclAssignTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.Match(VlangParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Var_type()
		}
		{
			p.SetState(398)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.expresion(0)
		}

	case 3:
		localctx = NewDeclTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(401)
			p.Match(VlangParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Var_type()
		}

	case 4:
		localctx = NewArrayAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(404)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Array_literal()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_literalContext is an interface to support dynamic dispatch.
type IArray_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArray_literalContext differentiates from other interfaces.
	IsArray_literalContext()
}

type Array_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_literalContext() *Array_literalContext {
	var p = new(Array_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_array_literal
	return p
}

func InitEmptyArray_literalContext(p *Array_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_array_literal
}

func (*Array_literalContext) IsArray_literalContext() {}

func NewArray_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_literalContext {
	var p = new(Array_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_array_literal

	return p
}

func (s *Array_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_literalContext) CopyAll(ctx *Array_literalContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Array_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayLiteralContext struct {
	Array_literalContext
}

func NewArrayLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	InitEmptyArray_literalContext(&p.Array_literalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Array_literalContext))

	return p
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *ArrayLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *ArrayLiteralContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *ArrayLiteralContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ArrayLiteralContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ArrayLiteralContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ArrayLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ArrayLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Array_literal() (localctx IArray_literalContext) {
	localctx = NewArray_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VlangParserRULE_array_literal)
	var _la int

	localctx = NewArrayLiteralContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.Match(VlangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(410)
		p.Match(VlangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(411)
		p.Var_type()
	}
	{
		p.SetState(412)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6917524285980540928) != 0 {
		{
			p.SetState(413)
			p.expresion(0)
		}
		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VlangParserCOMMA {
			{
				p.SetState(414)
				p.Match(VlangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(415)
				p.expresion(0)
			}

			p.SetState(420)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(423)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IId_patternContext is an interface to support dynamic dispatch.
type IId_patternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsId_patternContext differentiates from other interfaces.
	IsId_patternContext()
}

type Id_patternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_patternContext() *Id_patternContext {
	var p = new(Id_patternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_id_pattern
	return p
}

func InitEmptyId_patternContext(p *Id_patternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_id_pattern
}

func (*Id_patternContext) IsId_patternContext() {}

func NewId_patternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_patternContext {
	var p = new(Id_patternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_id_pattern

	return p
}

func (s *Id_patternContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_patternContext) CopyAll(ctx *Id_patternContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Id_patternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_patternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdPatternContext struct {
	Id_patternContext
}

func NewIdPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdPatternContext {
	var p = new(IdPatternContext)

	InitEmptyId_patternContext(&p.Id_patternContext)
	p.parser = parser
	p.CopyAll(ctx.(*Id_patternContext))

	return p
}

func (s *IdPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdPatternContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *IdPatternContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *IdPatternContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(VlangParserDOT)
}

func (s *IdPatternContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, i)
}

func (s *IdPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIdPattern(s)
	}
}

func (s *IdPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIdPattern(s)
	}
}

func (s *IdPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIdPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Id_pattern() (localctx IId_patternContext) {
	localctx = NewId_patternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VlangParserRULE_id_pattern)
	var _la int

	localctx = NewIdPatternContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserDOT {
		{
			p.SetState(426)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionexpreContext struct {
	ExpresionContext
}

func NewFuncionexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionexpreContext {
	var p = new(FuncionexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *FuncionexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionexpreContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *FuncionexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncionexpre(s)
	}
}

func (s *FuncionexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncionexpre(s)
	}
}

func (s *FuncionexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncionexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarioContext struct {
	ExpresionContext
	op antlr.Token
}

func NewUnarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarioContext {
	var p = new(UnarioContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *UnarioContext) GetOp() antlr.Token { return s.op }

func (s *UnarioContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarioContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *UnarioContext) NOT() antlr.TerminalNode {
	return s.GetToken(VlangParserNOT, 0)
}

func (s *UnarioContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *UnarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterUnario(s)
	}
}

func (s *UnarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitUnario(s)
	}
}

func (s *UnarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitUnario(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionExprContext struct {
	ExpresionContext
}

func NewAsignacionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionExprContext {
	var p = new(AsignacionExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionExprContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *AsignacionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionExpr(s)
	}
}

func (s *AsignacionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionExpr(s)
	}
}

func (s *AsignacionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentesisexpreContext struct {
	ExpresionContext
}

func NewParentesisexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentesisexpreContext {
	var p = new(ParentesisexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ParentesisexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentesisexpreContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *ParentesisexpreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ParentesisexpreContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *ParentesisexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParentesisexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncredecrContext struct {
	ExpresionContext
}

func NewIncredecrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncredecrContext {
	var p = new(IncredecrContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IncredecrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncredecrContext) Incredecre() IIncredecreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncredecreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncredecreContext)
}

func (s *IncredecrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIncredecr(s)
	}
}

func (s *IncredecrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIncredecr(s)
	}
}

func (s *IncredecrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIncredecr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorexpresionContext struct {
	ExpresionContext
}

func NewValorexpresionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorexpresionContext {
	var p = new(ValorexpresionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ValorexpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorexpresionContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValorexpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorexpresion(s)
	}
}

func (s *ValorexpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorexpresion(s)
	}
}

func (s *ValorexpresionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorexpresion(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayexpreContext struct {
	ExpresionContext
}

func NewArrayexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayexpreContext {
	var p = new(ArrayexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ArrayexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayexpreContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *ArrayexpreContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayexpreContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ArrayexpreContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *ArrayexpreContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ArrayexpreContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ArrayexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterArrayexpre(s)
	}
}

func (s *ArrayexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitArrayexpre(s)
	}
}

func (s *ArrayexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitArrayexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionforContext struct {
	ExpresionContext
}

func NewAsignacionforContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionforContext {
	var p = new(AsignacionforContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionforContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionforContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *AsignacionforContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionforContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionfor(s)
	}
}

func (s *AsignacionforContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionfor(s)
	}
}

func (s *AsignacionforContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionfor(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExpresionContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expdotexp1Context struct {
	ExpresionContext
}

func NewExpdotexp1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expdotexp1Context {
	var p = new(Expdotexp1Context)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expdotexp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expdotexp1Context) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *Expdotexp1Context) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *Expdotexp1Context) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *Expdotexp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpdotexp1(s)
	}
}

func (s *Expdotexp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpdotexp1(s)
	}
}

func (s *Expdotexp1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpdotexp1(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpdotexpContext struct {
	ExpresionContext
}

func NewExpdotexpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpdotexpContext {
	var p = new(ExpdotexpContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExpdotexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpdotexpContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ExpdotexpContext) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *ExpdotexpContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpdotexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpdotexp(s)
	}
}

func (s *ExpdotexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpdotexp(s)
	}
}

func (s *ExpdotexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpdotexp(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewBinaryExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpContext {
	var p = new(BinaryExpContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *BinaryExpContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExpContext) GetLeft() IExpresionContext { return s.left }

func (s *BinaryExpContext) GetRight() IExpresionContext { return s.right }

func (s *BinaryExpContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *BinaryExpContext) SetRight(v IExpresionContext) { s.right = v }

func (s *BinaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *BinaryExpContext) MUL() antlr.TerminalNode {
	return s.GetToken(VlangParserMUL, 0)
}

func (s *BinaryExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(VlangParserDIV, 0)
}

func (s *BinaryExpContext) MOD() antlr.TerminalNode {
	return s.GetToken(VlangParserMOD, 0)
}

func (s *BinaryExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VlangParserPLUS, 0)
}

func (s *BinaryExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *BinaryExpContext) LT() antlr.TerminalNode {
	return s.GetToken(VlangParserLT, 0)
}

func (s *BinaryExpContext) LE() antlr.TerminalNode {
	return s.GetToken(VlangParserLE, 0)
}

func (s *BinaryExpContext) GT() antlr.TerminalNode {
	return s.GetToken(VlangParserGT, 0)
}

func (s *BinaryExpContext) GE() antlr.TerminalNode {
	return s.GetToken(VlangParserGE, 0)
}

func (s *BinaryExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(VlangParserEQ, 0)
}

func (s *BinaryExpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(VlangParserNEQ, 0)
}

func (s *BinaryExpContext) AND() antlr.TerminalNode {
	return s.GetToken(VlangParserAND, 0)
}

func (s *BinaryExpContext) OR() antlr.TerminalNode {
	return s.GetToken(VlangParserOR, 0)
}

func (s *BinaryExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBinaryExp(s)
	}
}

func (s *BinaryExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBinaryExp(s)
	}
}

func (s *BinaryExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBinaryExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *VlangParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, VlangParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValorexpresionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(434)
			p.Valor()
		}

	case 2:
		localctx = NewFuncionexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(435)
			p.Func_call()
		}

	case 3:
		localctx = NewAsignacionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(436)
			p.Assign_stmt()
		}

	case 4:
		localctx = NewParentesisexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(437)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.expresion(0)
		}
		{
			p.SetState(439)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewArrayexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(441)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.expresion(0)
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VlangParserCOMMA {
			{
				p.SetState(443)
				p.Match(VlangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(444)
				p.expresion(0)
			}

			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(450)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewUnarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(452)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnarioContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VlangParserMINUS || _la == VlangParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnarioContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(453)
			p.expresion(12)
		}

	case 7:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(454)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewIncredecrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(455)
			p.Incredecre()
		}

	case 9:
		localctx = NewExpdotexp1Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(456)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewExpdotexpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(459)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.expresion(2)
		}

	case 11:
		localctx = NewAsignacionforContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(462)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(464)
			p.expresion(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(485)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(467)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(468)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3848290697216) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(469)

					var _x = p.expresion(12)

					localctx.(*BinaryExpContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(470)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(471)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserPLUS || _la == VlangParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(472)

					var _x = p.expresion(11)

					localctx.(*BinaryExpContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(473)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(474)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2111062325329920) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(475)

					var _x = p.expresion(10)

					localctx.(*BinaryExpContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(476)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(477)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserEQ || _la == VlangParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(478)

					var _x = p.expresion(9)

					localctx.(*BinaryExpContext).right = _x
				}

			case 5:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(479)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(480)

					var _m = p.Match(VlangParserAND)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(481)

					var _x = p.expresion(8)

					localctx.(*BinaryExpContext).right = _x
				}

			case 6:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(482)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(483)

					var _m = p.Match(VlangParserOR)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(484)

					var _x = p.expresion(7)

					localctx.(*BinaryExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) CopyAll(ctx *ParametrosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamListContext struct {
	ParametrosContext
}

func NewParamListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamListContext {
	var p = new(ParamListContext)

	InitEmptyParametrosContext(&p.ParametrosContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParametrosContext))

	return p
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) AllFunc_param() []IFunc_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_paramContext); ok {
			len++
		}
	}

	tst := make([]IFunc_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_paramContext); ok {
			tst[i] = t.(IFunc_paramContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Func_param(i int) IFunc_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_paramContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_paramContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VlangParserRULE_parametros)
	var _la int

	localctx = NewParamListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.Func_param()
	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(491)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)
			p.Func_param()
		}

		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_paramContext is an interface to support dynamic dispatch.
type IFunc_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_paramContext differentiates from other interfaces.
	IsFunc_paramContext()
}

type Func_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_paramContext() *Func_paramContext {
	var p = new(Func_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_param
	return p
}

func InitEmptyFunc_paramContext(p *Func_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_param
}

func (*Func_paramContext) IsFunc_paramContext() {}

func NewFunc_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_paramContext {
	var p = new(Func_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_func_param

	return p
}

func (s *Func_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_paramContext) CopyAll(ctx *Func_paramContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncParamContext struct {
	Func_paramContext
}

func NewFuncParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParamContext {
	var p = new(FuncParamContext)

	InitEmptyFunc_paramContext(&p.Func_paramContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_paramContext))

	return p
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *FuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncParam(s)
	}
}

func (s *FuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncParam(s)
	}
}

func (s *FuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Func_param() (localctx IFunc_paramContext) {
	localctx = NewFunc_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VlangParserRULE_func_param)
	localctx = NewFuncParamContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) CopyAll(ctx *Arg_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgListContext struct {
	Arg_listContext
}

func NewArgListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgListContext {
	var p = new(ArgListContext)

	InitEmptyArg_listContext(&p.Arg_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arg_listContext))

	return p
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) AllFunc_arg() []IFunc_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_argContext); ok {
			len++
		}
	}

	tst := make([]IFunc_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_argContext); ok {
			tst[i] = t.(IFunc_argContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Func_arg(i int) IFunc_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_argContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_argContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, VlangParserRULE_arg_list)
	var _la int

	localctx = NewArgListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		p.Func_arg()
	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(501)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(502)
			p.Func_arg()
		}

		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_argContext is an interface to support dynamic dispatch.
type IFunc_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_argContext differentiates from other interfaces.
	IsFunc_argContext()
}

type Func_argContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_argContext() *Func_argContext {
	var p = new(Func_argContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_arg
	return p
}

func InitEmptyFunc_argContext(p *Func_argContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_arg
}

func (*Func_argContext) IsFunc_argContext() {}

func NewFunc_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_argContext {
	var p = new(Func_argContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_func_arg

	return p
}

func (s *Func_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_argContext) CopyAll(ctx *Func_argContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncArgContext struct {
	Func_argContext
}

func NewFuncArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArgContext {
	var p = new(FuncArgContext)

	InitEmptyFunc_argContext(&p.Func_argContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_argContext))

	return p
}

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *FuncArgContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *FuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncArg(s)
	}
}

func (s *FuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncArg(s)
	}
}

func (s *FuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Func_arg() (localctx IFunc_argContext) {
	localctx = NewFunc_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, VlangParserRULE_func_arg)
	localctx = NewFuncArgContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(508)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(509)
		p.Var_type()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValoresContext is an interface to support dynamic dispatch.
type IValoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Valor() IValorContext

	// IsValoresContext differentiates from other interfaces.
	IsValoresContext()
}

type ValoresContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValoresContext() *ValoresContext {
	var p = new(ValoresContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valores
	return p
}

func InitEmptyValoresContext(p *ValoresContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valores
}

func (*ValoresContext) IsValoresContext() {}

func NewValoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValoresContext {
	var p = new(ValoresContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valores

	return p
}

func (s *ValoresContext) GetParser() antlr.Parser { return s.parser }

func (s *ValoresContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValores(s)
	}
}

func (s *ValoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValores(s)
	}
}

func (s *ValoresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValores(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valores() (localctx IValoresContext) {
	localctx = NewValoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, VlangParserRULE_valores)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(511)
		p.Valor()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValorContext is an interface to support dynamic dispatch.
type IValorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValorContext differentiates from other interfaces.
	IsValorContext()
}

type ValorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValorContext() *ValorContext {
	var p = new(ValorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
	return p
}

func InitEmptyValorContext(p *ValorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
}

func (*ValorContext) IsValorContext() {}

func NewValorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValorContext {
	var p = new(ValorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valor

	return p
}

func (s *ValorContext) GetParser() antlr.Parser { return s.parser }

func (s *ValorContext) CopyAll(ctx *ValorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValorDecimalContext struct {
	ValorContext
}

func NewValorDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorDecimalContext {
	var p = new(ValorDecimalContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorDecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorDecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(VlangParserDECIMAL, 0)
}

func (s *ValorDecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorDecimal(s)
	}
}

func (s *ValorDecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorDecimal(s)
	}
}

func (s *ValorDecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorDecimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorEnteroContext struct {
	ValorContext
}

func NewValorEnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorEnteroContext {
	var p = new(ValorEnteroContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorEnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorEnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(VlangParserENTERO, 0)
}

func (s *ValorEnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorEntero(s)
	}
}

func (s *ValorEnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorEntero(s)
	}
}

func (s *ValorEnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorEntero(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorStringInterpolationContext struct {
	ValorContext
}

func NewValorStringInterpolationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorStringInterpolationContext {
	var p = new(ValorStringInterpolationContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorStringInterpolationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorStringInterpolationContext) STRING_INTERPOLATION() antlr.TerminalNode {
	return s.GetToken(VlangParserSTRING_INTERPOLATION, 0)
}

func (s *ValorStringInterpolationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorStringInterpolation(s)
	}
}

func (s *ValorStringInterpolationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorStringInterpolation(s)
	}
}

func (s *ValorStringInterpolationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorStringInterpolation(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorBooleanoContext struct {
	ValorContext
}

func NewValorBooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorBooleanoContext {
	var p = new(ValorBooleanoContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorBooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorBooleanoContext) BOOLEANO() antlr.TerminalNode {
	return s.GetToken(VlangParserBOOLEANO, 0)
}

func (s *ValorBooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorBooleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCaracterContext struct {
	ValorContext
}

func NewValorCaracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCaracterContext {
	var p = new(ValorCaracterContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCaracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCaracterContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(VlangParserCARACTER, 0)
}

func (s *ValorCaracterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCaracter(s)
	}
}

func (s *ValorCaracterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCaracter(s)
	}
}

func (s *ValorCaracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCaracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCadenaContext struct {
	ValorContext
}

func NewValorCadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCadenaContext {
	var p = new(ValorCadenaContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCadenaContext) CADENA() antlr.TerminalNode {
	return s.GetToken(VlangParserCADENA, 0)
}

func (s *ValorCadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCadena(s)
	}
}

func (s *ValorCadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCadena(s)
	}
}

func (s *ValorCadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCadena(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorFloatContext struct {
	ValorContext
}

func NewValorFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorFloatContext {
	var p = new(ValorFloatContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorFloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(VlangParserFLOAT, 0)
}

func (s *ValorFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorFloat(s)
	}
}

func (s *ValorFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorFloat(s)
	}
}

func (s *ValorFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valor() (localctx IValorContext) {
	localctx = NewValorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, VlangParserRULE_valor)
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserENTERO:
		localctx = NewValorEnteroContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(513)
			p.Match(VlangParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserDECIMAL:
		localctx = NewValorDecimalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(514)
			p.Match(VlangParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCADENA:
		localctx = NewValorCadenaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(515)
			p.Match(VlangParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserBOOLEANO:
		localctx = NewValorBooleanoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(516)
			p.Match(VlangParserBOOLEANO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCARACTER:
		localctx = NewValorCaracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(517)
			p.Match(VlangParserCARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserFLOAT:
		localctx = NewValorFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(518)
			p.Match(VlangParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserSTRING_INTERPOLATION:
		localctx = NewValorStringInterpolationContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(519)
			p.Match(VlangParserSTRING_INTERPOLATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrct_instanciaContext is an interface to support dynamic dispatch.
type IStrct_instanciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStrct_instanciaContext differentiates from other interfaces.
	IsStrct_instanciaContext()
}

type Strct_instanciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrct_instanciaContext() *Strct_instanciaContext {
	var p = new(Strct_instanciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_strct_instancia
	return p
}

func InitEmptyStrct_instanciaContext(p *Strct_instanciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_strct_instancia
}

func (*Strct_instanciaContext) IsStrct_instanciaContext() {}

func NewStrct_instanciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Strct_instanciaContext {
	var p = new(Strct_instanciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_strct_instancia

	return p
}

func (s *Strct_instanciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Strct_instanciaContext) CopyAll(ctx *Strct_instanciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Strct_instanciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Strct_instanciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructInstanciaContext struct {
	Strct_instanciaContext
}

func NewStructInstanciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInstanciaContext {
	var p = new(StructInstanciaContext)

	InitEmptyStrct_instanciaContext(&p.Strct_instanciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Strct_instanciaContext))

	return p
}

func (s *StructInstanciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstanciaContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructInstanciaContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *StructInstanciaContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *StructInstanciaContext) AllStruct_prop() []IStruct_propContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_propContext); ok {
			len++
		}
	}

	tst := make([]IStruct_propContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_propContext); ok {
			tst[i] = t.(IStruct_propContext)
			i++
		}
	}

	return tst
}

func (s *StructInstanciaContext) Struct_prop(i int) IStruct_propContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_propContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_propContext)
}

func (s *StructInstanciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructInstancia(s)
	}
}

func (s *StructInstanciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructInstancia(s)
	}
}

func (s *StructInstanciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructInstancia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Strct_instancia() (localctx IStrct_instanciaContext) {
	localctx = NewStrct_instanciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, VlangParserRULE_strct_instancia)
	var _la int

	localctx = NewStructInstanciaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		p.Match(VlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(523)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(524)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserID {
		{
			p.SetState(525)
			p.Struct_prop()
		}

		p.SetState(530)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(531)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_propContext is an interface to support dynamic dispatch.
type IStruct_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_propContext differentiates from other interfaces.
	IsStruct_propContext()
}

type Struct_propContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_propContext() *Struct_propContext {
	var p = new(Struct_propContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_prop
	return p
}

func InitEmptyStruct_propContext(p *Struct_propContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_prop
}

func (*Struct_propContext) IsStruct_propContext() {}

func NewStruct_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_propContext {
	var p = new(Struct_propContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_struct_prop

	return p
}

func (s *Struct_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_propContext) CopyAll(ctx *Struct_propContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructAttrContext struct {
	Struct_propContext
}

func NewStructAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAttrContext {
	var p = new(StructAttrContext)

	InitEmptyStruct_propContext(&p.Struct_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_propContext))

	return p
}

func (s *StructAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttrContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructAttrContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *StructAttrContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *StructAttrContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *StructAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructAttr(s)
	}
}

func (s *StructAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructAttr(s)
	}
}

func (s *StructAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Struct_prop() (localctx IStruct_propContext) {
	localctx = NewStruct_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, VlangParserRULE_struct_prop)
	var _la int

	localctx = NewStructAttrContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(533)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(534)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(535)
		p.expresion(0)
	}
	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserCOMMA {
		{
			p.SetState(536)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) CopyAll(ctx *Var_typeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type F64TypeContext struct {
	Var_typeContext
}

func NewF64TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *F64TypeContext {
	var p = new(F64TypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *F64TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F64TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterF64Type(s)
	}
}

func (s *F64TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitF64Type(s)
	}
}

func (s *F64TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitF64Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type VoidTypeContext struct {
	Var_typeContext
}

func NewVoidTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidTypeContext {
	var p = new(VoidTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *VoidTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterVoidType(s)
	}
}

func (s *VoidTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitVoidType(s)
	}
}

func (s *VoidTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitVoidType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayTypeContext struct {
	Var_typeContext
}

func NewArrayTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *ArrayTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *ArrayTypeContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolTypeContext struct {
	Var_typeContext
}

func NewBoolTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolTypeContext {
	var p = new(BoolTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBoolType(s)
	}
}

func (s *BoolTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBoolType(s)
	}
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringTypeContext struct {
	Var_typeContext
}

func NewStringTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringTypeContext {
	var p = new(StringTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

type CharTypeContext struct {
	Var_typeContext
}

func NewCharTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharTypeContext {
	var p = new(CharTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *CharTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCharType(s)
	}
}

func (s *CharTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCharType(s)
	}
}

func (s *CharTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCharType(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntTypeContext struct {
	Var_typeContext
}

func NewIntTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntTypeContext {
	var p = new(IntTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatTypeContext struct {
	Var_typeContext
}

func NewFloatTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatTypeContext {
	var p = new(FloatTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *FloatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFloatType(s)
	}
}

func (s *FloatTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFloatType(s)
	}
}

func (s *FloatTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFloatType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, VlangParserRULE_var_type)
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserT__5:
		localctx = NewIntTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(539)
			p.Match(VlangParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__6:
		localctx = NewFloatTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(540)
			p.Match(VlangParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__7:
		localctx = NewF64TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(541)
			p.Match(VlangParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__8:
		localctx = NewStringTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(542)
			p.Match(VlangParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__9:
		localctx = NewBoolTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(543)
			p.Match(VlangParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__10:
		localctx = NewCharTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(544)
			p.Match(VlangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__11:
		localctx = NewVoidTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(545)
			p.Match(VlangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserLBRACK:
		localctx = NewArrayTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(546)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(548)
			p.Var_type()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_vectorContext is an interface to support dynamic dispatch.
type IStruct_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_vectorContext differentiates from other interfaces.
	IsStruct_vectorContext()
}

type Struct_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_vectorContext() *Struct_vectorContext {
	var p = new(Struct_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_vector
	return p
}

func InitEmptyStruct_vectorContext(p *Struct_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_vector
}

func (*Struct_vectorContext) IsStruct_vectorContext() {}

func NewStruct_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_vectorContext {
	var p = new(Struct_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_struct_vector

	return p
}

func (s *Struct_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_vectorContext) CopyAll(ctx *Struct_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructVectorContext struct {
	Struct_vectorContext
}

func NewStructVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructVectorContext {
	var p = new(StructVectorContext)

	InitEmptyStruct_vectorContext(&p.Struct_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_vectorContext))

	return p
}

func (s *StructVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructVectorContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *StructVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructVectorContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *StructVectorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *StructVectorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *StructVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructVector(s)
	}
}

func (s *StructVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructVector(s)
	}
}

func (s *StructVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Struct_vector() (localctx IStruct_vectorContext) {
	localctx = NewStruct_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, VlangParserRULE_struct_vector)
	localctx = NewStructVectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(551)
		p.Match(VlangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(552)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(553)
		p.Match(VlangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(554)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(555)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncredecreContext is an interface to support dynamic dispatch.
type IIncredecreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIncredecreContext differentiates from other interfaces.
	IsIncredecreContext()
}

type IncredecreContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncredecreContext() *IncredecreContext {
	var p = new(IncredecreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_incredecre
	return p
}

func InitEmptyIncredecreContext(p *IncredecreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_incredecre
}

func (*IncredecreContext) IsIncredecreContext() {}

func NewIncredecreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncredecreContext {
	var p = new(IncredecreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_incredecre

	return p
}

func (s *IncredecreContext) GetParser() antlr.Parser { return s.parser }

func (s *IncredecreContext) CopyAll(ctx *IncredecreContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IncredecreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncredecreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IncrementoContext struct {
	IncredecreContext
}

func NewIncrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementoContext {
	var p = new(IncrementoContext)

	InitEmptyIncredecreContext(&p.IncredecreContext)
	p.parser = parser
	p.CopyAll(ctx.(*IncredecreContext))

	return p
}

func (s *IncrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IncrementoContext) INC() antlr.TerminalNode {
	return s.GetToken(VlangParserINC, 0)
}

func (s *IncrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIncremento(s)
	}
}

func (s *IncrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIncremento(s)
	}
}

func (s *IncrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIncremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecrementoContext struct {
	IncredecreContext
}

func NewDecrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecrementoContext {
	var p = new(DecrementoContext)

	InitEmptyIncredecreContext(&p.IncredecreContext)
	p.parser = parser
	p.CopyAll(ctx.(*IncredecreContext))

	return p
}

func (s *DecrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *DecrementoContext) DEC() antlr.TerminalNode {
	return s.GetToken(VlangParserDEC, 0)
}

func (s *DecrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDecremento(s)
	}
}

func (s *DecrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDecremento(s)
	}
}

func (s *DecrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDecremento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Incredecre() (localctx IIncredecreContext) {
	localctx = NewIncredecreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, VlangParserRULE_incredecre)
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(557)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)
			p.Match(VlangParserINC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(559)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)
			p.Match(VlangParserDEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 26:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VlangParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
