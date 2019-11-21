// Code generated from Feature.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Feature

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 289,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 5, 2, 54, 10, 2,
	3, 2, 7, 2, 57, 10, 2, 12, 2, 14, 2, 60, 11, 2, 3, 2, 7, 2, 63, 10, 2,
	12, 2, 14, 2, 66, 11, 2, 3, 2, 5, 2, 69, 10, 2, 3, 2, 7, 2, 72, 10, 2,
	12, 2, 14, 2, 75, 11, 2, 3, 2, 7, 2, 78, 10, 2, 12, 2, 14, 2, 81, 11, 2,
	3, 2, 3, 2, 7, 2, 85, 10, 2, 12, 2, 14, 2, 88, 11, 2, 3, 2, 3, 2, 6, 2,
	92, 10, 2, 13, 2, 14, 2, 93, 3, 2, 3, 2, 3, 2, 7, 2, 99, 10, 2, 12, 2,
	14, 2, 102, 11, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 109, 10, 3, 12,
	3, 14, 3, 112, 11, 3, 3, 4, 5, 4, 115, 10, 4, 3, 4, 5, 4, 118, 10, 4, 3,
	4, 3, 4, 7, 4, 122, 10, 4, 12, 4, 14, 4, 125, 11, 4, 3, 4, 3, 4, 6, 4,
	129, 10, 4, 13, 4, 14, 4, 130, 3, 4, 3, 4, 3, 5, 5, 5, 136, 10, 5, 3, 5,
	5, 5, 139, 10, 5, 3, 5, 3, 5, 7, 5, 143, 10, 5, 12, 5, 14, 5, 146, 11,
	5, 3, 5, 3, 5, 6, 5, 150, 10, 5, 13, 5, 14, 5, 151, 3, 5, 3, 5, 3, 5, 3,
	6, 7, 6, 158, 10, 6, 12, 6, 14, 6, 161, 11, 6, 3, 7, 3, 7, 7, 7, 165, 10,
	7, 12, 7, 14, 7, 168, 11, 7, 3, 7, 3, 7, 6, 7, 172, 10, 7, 13, 7, 14, 7,
	173, 3, 7, 5, 7, 177, 10, 7, 3, 8, 6, 8, 180, 10, 8, 13, 8, 14, 8, 181,
	3, 9, 3, 9, 7, 9, 186, 10, 9, 12, 9, 14, 9, 189, 11, 9, 3, 9, 3, 9, 6,
	9, 193, 10, 9, 13, 9, 14, 9, 194, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6, 11,
	202, 10, 11, 13, 11, 14, 11, 203, 3, 12, 7, 12, 207, 10, 12, 12, 12, 14,
	12, 210, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 216, 10, 12, 13, 12,
	14, 12, 217, 3, 12, 7, 12, 221, 10, 12, 12, 12, 14, 12, 224, 11, 12, 3,
	12, 6, 12, 227, 10, 12, 13, 12, 14, 12, 228, 3, 13, 3, 13, 7, 13, 233,
	10, 13, 12, 13, 14, 13, 236, 11, 13, 3, 14, 3, 14, 3, 14, 7, 14, 241, 10,
	14, 12, 14, 14, 14, 244, 11, 14, 3, 14, 6, 14, 247, 10, 14, 13, 14, 14,
	14, 248, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 7, 16, 257, 10, 16,
	12, 16, 14, 16, 260, 11, 16, 3, 17, 3, 17, 3, 17, 3, 18, 7, 18, 266, 10,
	18, 12, 18, 14, 18, 269, 11, 18, 3, 19, 3, 19, 5, 19, 273, 10, 19, 3, 20,
	3, 20, 5, 20, 277, 10, 20, 3, 21, 3, 21, 5, 21, 281, 10, 21, 3, 22, 3,
	22, 3, 23, 3, 23, 5, 23, 287, 10, 23, 3, 23, 2, 2, 24, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2,
	5, 4, 2, 3, 3, 17, 17, 3, 2, 17, 17, 3, 2, 10, 14, 2, 307, 2, 49, 3, 2,
	2, 2, 4, 110, 3, 2, 2, 2, 6, 114, 3, 2, 2, 2, 8, 135, 3, 2, 2, 2, 10, 159,
	3, 2, 2, 2, 12, 162, 3, 2, 2, 2, 14, 179, 3, 2, 2, 2, 16, 183, 3, 2, 2,
	2, 18, 198, 3, 2, 2, 2, 20, 201, 3, 2, 2, 2, 22, 208, 3, 2, 2, 2, 24, 234,
	3, 2, 2, 2, 26, 237, 3, 2, 2, 2, 28, 250, 3, 2, 2, 2, 30, 258, 3, 2, 2,
	2, 32, 261, 3, 2, 2, 2, 34, 267, 3, 2, 2, 2, 36, 270, 3, 2, 2, 2, 38, 274,
	3, 2, 2, 2, 40, 278, 3, 2, 2, 2, 42, 282, 3, 2, 2, 2, 44, 284, 3, 2, 2,
	2, 46, 48, 7, 17, 2, 2, 47, 46, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47,
	3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2,
	52, 54, 5, 30, 16, 2, 53, 52, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 58, 3,
	2, 2, 2, 55, 57, 7, 17, 2, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58,
	56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 64, 3, 2, 2, 2, 60, 58, 3, 2, 2,
	2, 61, 63, 7, 18, 2, 2, 62, 61, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2,
	67, 69, 5, 26, 14, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 73, 3,
	2, 2, 2, 70, 72, 7, 17, 2, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73,
	71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 79, 3, 2, 2, 2, 75, 73, 3, 2, 2,
	2, 76, 78, 7, 18, 2, 2, 77, 76, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77,
	3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2,
	82, 86, 5, 36, 19, 2, 83, 85, 7, 18, 2, 2, 84, 83, 3, 2, 2, 2, 85, 88,
	3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2,
	88, 86, 3, 2, 2, 2, 89, 91, 5, 34, 18, 2, 90, 92, 7, 17, 2, 2, 91, 90,
	3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2,
	94, 100, 3, 2, 2, 2, 95, 96, 5, 4, 3, 2, 96, 97, 11, 2, 2, 2, 97, 99, 3,
	2, 2, 2, 98, 95, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2,
	100, 101, 3, 2, 2, 2, 101, 103, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103,
	104, 5, 4, 3, 2, 104, 105, 7, 2, 2, 3, 105, 3, 3, 2, 2, 2, 106, 109, 5,
	8, 5, 2, 107, 109, 5, 6, 4, 2, 108, 106, 3, 2, 2, 2, 108, 107, 3, 2, 2,
	2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	5, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 115, 5, 30, 16, 2, 114, 113,
	3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 117, 3, 2, 2, 2, 116, 118, 5, 26,
	14, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2,
	119, 123, 5, 38, 20, 2, 120, 122, 7, 18, 2, 2, 121, 120, 3, 2, 2, 2, 122,
	125, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126,
	3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 126, 128, 5, 34, 18, 2, 127, 129, 7,
	17, 2, 2, 128, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 128, 3, 2, 2,
	2, 130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 5, 10, 6, 2, 133,
	7, 3, 2, 2, 2, 134, 136, 5, 30, 16, 2, 135, 134, 3, 2, 2, 2, 135, 136,
	3, 2, 2, 2, 136, 138, 3, 2, 2, 2, 137, 139, 5, 26, 14, 2, 138, 137, 3,
	2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 144, 5, 40, 21,
	2, 141, 143, 7, 18, 2, 2, 142, 141, 3, 2, 2, 2, 143, 146, 3, 2, 2, 2, 144,
	142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 147, 3, 2, 2, 2, 146, 144,
	3, 2, 2, 2, 147, 149, 5, 34, 18, 2, 148, 150, 7, 17, 2, 2, 149, 148, 3,
	2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2,
	2, 152, 153, 3, 2, 2, 2, 153, 154, 5, 10, 6, 2, 154, 155, 5, 14, 8, 2,
	155, 9, 3, 2, 2, 2, 156, 158, 5, 12, 7, 2, 157, 156, 3, 2, 2, 2, 158, 161,
	3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 11, 3, 2,
	2, 2, 161, 159, 3, 2, 2, 2, 162, 166, 5, 42, 22, 2, 163, 165, 7, 18, 2,
	2, 164, 163, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166,
	167, 3, 2, 2, 2, 167, 169, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 171,
	5, 34, 18, 2, 170, 172, 7, 17, 2, 2, 171, 170, 3, 2, 2, 2, 172, 173, 3,
	2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 176, 3, 2, 2,
	2, 175, 177, 5, 18, 10, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2,
	177, 13, 3, 2, 2, 2, 178, 180, 5, 16, 9, 2, 179, 178, 3, 2, 2, 2, 180,
	181, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 15, 3,
	2, 2, 2, 183, 187, 5, 44, 23, 2, 184, 186, 7, 18, 2, 2, 185, 184, 3, 2,
	2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2,
	188, 190, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 192, 5, 34, 18, 2, 191,
	193, 7, 17, 2, 2, 192, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 192,
	3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 197, 5, 20,
	11, 2, 197, 17, 3, 2, 2, 2, 198, 199, 5, 20, 11, 2, 199, 19, 3, 2, 2, 2,
	200, 202, 5, 22, 12, 2, 201, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203,
	201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 21, 3, 2, 2, 2, 205, 207, 7,
	18, 2, 2, 206, 205, 3, 2, 2, 2, 207, 210, 3, 2, 2, 2, 208, 206, 3, 2, 2,
	2, 208, 209, 3, 2, 2, 2, 209, 211, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211,
	215, 7, 3, 2, 2, 212, 213, 5, 24, 13, 2, 213, 214, 7, 3, 2, 2, 214, 216,
	3, 2, 2, 2, 215, 212, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 215, 3, 2,
	2, 2, 217, 218, 3, 2, 2, 2, 218, 222, 3, 2, 2, 2, 219, 221, 7, 18, 2, 2,
	220, 219, 3, 2, 2, 2, 221, 224, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222,
	223, 3, 2, 2, 2, 223, 226, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 227,
	7, 17, 2, 2, 226, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 226, 3, 2,
	2, 2, 228, 229, 3, 2, 2, 2, 229, 23, 3, 2, 2, 2, 230, 231, 10, 2, 2, 2,
	231, 233, 11, 2, 2, 2, 232, 230, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234,
	232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 25, 3, 2, 2, 2, 236, 234, 3,
	2, 2, 2, 237, 242, 5, 28, 15, 2, 238, 239, 7, 18, 2, 2, 239, 241, 5, 28,
	15, 2, 240, 238, 3, 2, 2, 2, 241, 244, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2,
	242, 243, 3, 2, 2, 2, 243, 246, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 245,
	247, 7, 17, 2, 2, 246, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 246,
	3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 27, 3, 2, 2, 2, 250, 251, 7, 4,
	2, 2, 251, 252, 7, 16, 2, 2, 252, 29, 3, 2, 2, 2, 253, 254, 5, 32, 17,
	2, 254, 255, 7, 17, 2, 2, 255, 257, 3, 2, 2, 2, 256, 253, 3, 2, 2, 2, 257,
	260, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 31, 3,
	2, 2, 2, 260, 258, 3, 2, 2, 2, 261, 262, 7, 5, 2, 2, 262, 263, 5, 34, 18,
	2, 263, 33, 3, 2, 2, 2, 264, 266, 10, 3, 2, 2, 265, 264, 3, 2, 2, 2, 266,
	269, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 35, 3,
	2, 2, 2, 269, 267, 3, 2, 2, 2, 270, 272, 7, 6, 2, 2, 271, 273, 7, 7, 2,
	2, 272, 271, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 37, 3, 2, 2, 2, 274,
	276, 7, 8, 2, 2, 275, 277, 7, 7, 2, 2, 276, 275, 3, 2, 2, 2, 276, 277,
	3, 2, 2, 2, 277, 39, 3, 2, 2, 2, 278, 280, 7, 9, 2, 2, 279, 281, 7, 7,
	2, 2, 280, 279, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 41, 3, 2, 2, 2,
	282, 283, 9, 4, 2, 2, 283, 43, 3, 2, 2, 2, 284, 286, 7, 15, 2, 2, 285,
	287, 7, 7, 2, 2, 286, 285, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 45, 3,
	2, 2, 2, 43, 49, 53, 58, 64, 68, 73, 79, 86, 93, 100, 108, 110, 114, 117,
	123, 130, 135, 138, 144, 151, 159, 166, 173, 176, 181, 187, 194, 203, 208,
	217, 222, 228, 234, 242, 248, 258, 267, 272, 276, 280, 286,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'|'", "'@'", "'#'", "'Feature'", "':'", "'Scenario'", "'Scenario Outline'",
	"'Given'", "'When'", "'Then'", "'And'", "'But'", "'Examples'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "NEWLINE",
	"SPACE",
}

var ruleNames = []string{
	"compilationUnit", "feature_elements", "scenario", "scenario_outline",
	"steps", "step", "examples_sections", "examples", "multiline_arg", "table",
	"table_row", "cell", "tags", "tag", "comment", "comment_line", "line_to_eol",
	"feature_keyword", "scenario_keyword", "scenario_outline_keyword", "step_keyword",
	"examples_keyword",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FeatureParser struct {
	*antlr.BaseParser
}

func NewFeatureParser(input antlr.TokenStream) *FeatureParser {
	this := new(FeatureParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Feature.g4"

	return this
}

// FeatureParser tokens.
const (
	FeatureParserEOF     = antlr.TokenEOF
	FeatureParserT__0    = 1
	FeatureParserT__1    = 2
	FeatureParserT__2    = 3
	FeatureParserT__3    = 4
	FeatureParserT__4    = 5
	FeatureParserT__5    = 6
	FeatureParserT__6    = 7
	FeatureParserT__7    = 8
	FeatureParserT__8    = 9
	FeatureParserT__9    = 10
	FeatureParserT__10   = 11
	FeatureParserT__11   = 12
	FeatureParserT__12   = 13
	FeatureParserID      = 14
	FeatureParserNEWLINE = 15
	FeatureParserSPACE   = 16
)

// FeatureParser rules.
const (
	FeatureParserRULE_compilationUnit          = 0
	FeatureParserRULE_feature_elements         = 1
	FeatureParserRULE_scenario                 = 2
	FeatureParserRULE_scenario_outline         = 3
	FeatureParserRULE_steps                    = 4
	FeatureParserRULE_step                     = 5
	FeatureParserRULE_examples_sections        = 6
	FeatureParserRULE_examples                 = 7
	FeatureParserRULE_multiline_arg            = 8
	FeatureParserRULE_table                    = 9
	FeatureParserRULE_table_row                = 10
	FeatureParserRULE_cell                     = 11
	FeatureParserRULE_tags                     = 12
	FeatureParserRULE_tag                      = 13
	FeatureParserRULE_comment                  = 14
	FeatureParserRULE_comment_line             = 15
	FeatureParserRULE_line_to_eol              = 16
	FeatureParserRULE_feature_keyword          = 17
	FeatureParserRULE_scenario_keyword         = 18
	FeatureParserRULE_scenario_outline_keyword = 19
	FeatureParserRULE_step_keyword             = 20
	FeatureParserRULE_examples_keyword         = 21
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) Feature_keyword() IFeature_keywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFeature_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFeature_keywordContext)
}

func (s *CompilationUnitContext) Line_to_eol() ILine_to_eolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILine_to_eolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILine_to_eolContext)
}

func (s *CompilationUnitContext) AllFeature_elements() []IFeature_elementsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFeature_elementsContext)(nil)).Elem())
	var tst = make([]IFeature_elementsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFeature_elementsContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) Feature_elements(i int) IFeature_elementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFeature_elementsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFeature_elementsContext)
}

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(FeatureParserEOF, 0)
}

func (s *CompilationUnitContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *CompilationUnitContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *CompilationUnitContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *CompilationUnitContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSPACE)
}

func (s *CompilationUnitContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSPACE, i)
}

func (s *CompilationUnitContext) Tags() ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FeatureParserRULE_compilationUnit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(44)
				p.Match(FeatureParserNEWLINE)
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(50)
			p.Comment()
		}

	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(53)
				p.Match(FeatureParserNEWLINE)
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(59)
				p.Match(FeatureParserSPACE)
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FeatureParserT__1 {
		{
			p.SetState(65)
			p.Tags()
		}

	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserNEWLINE {
		{
			p.SetState(68)
			p.Match(FeatureParserNEWLINE)
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSPACE {
		{
			p.SetState(74)
			p.Match(FeatureParserSPACE)
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Feature_keyword()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(81)
				p.Match(FeatureParserSPACE)
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	{
		p.SetState(87)
		p.Line_to_eol()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(88)
				p.Match(FeatureParserNEWLINE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(93)
				p.Feature_elements()
			}
			p.SetState(94)
			p.MatchWildcard()

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	{
		p.SetState(101)
		p.Feature_elements()
	}
	{
		p.SetState(102)
		p.Match(FeatureParserEOF)
	}

	return localctx
}

// IFeature_elementsContext is an interface to support dynamic dispatch.
type IFeature_elementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFeature_elementsContext differentiates from other interfaces.
	IsFeature_elementsContext()
}

type Feature_elementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeature_elementsContext() *Feature_elementsContext {
	var p = new(Feature_elementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_feature_elements
	return p
}

func (*Feature_elementsContext) IsFeature_elementsContext() {}

func NewFeature_elementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Feature_elementsContext {
	var p = new(Feature_elementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_feature_elements

	return p
}

func (s *Feature_elementsContext) GetParser() antlr.Parser { return s.parser }

func (s *Feature_elementsContext) AllScenario_outline() []IScenario_outlineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScenario_outlineContext)(nil)).Elem())
	var tst = make([]IScenario_outlineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScenario_outlineContext)
		}
	}

	return tst
}

func (s *Feature_elementsContext) Scenario_outline(i int) IScenario_outlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScenario_outlineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScenario_outlineContext)
}

func (s *Feature_elementsContext) AllScenario() []IScenarioContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScenarioContext)(nil)).Elem())
	var tst = make([]IScenarioContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScenarioContext)
		}
	}

	return tst
}

func (s *Feature_elementsContext) Scenario(i int) IScenarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScenarioContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScenarioContext)
}

func (s *Feature_elementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Feature_elementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Feature_elementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterFeature_elements(s)
	}
}

func (s *Feature_elementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitFeature_elements(s)
	}
}

func (s *Feature_elementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitFeature_elements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Feature_elements() (localctx IFeature_elementsContext) {
	localctx = NewFeature_elementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FeatureParserRULE_feature_elements)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(104)
					p.Scenario_outline()
				}

			case 2:
				{
					p.SetState(105)
					p.Scenario()
				}

			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IScenarioContext is an interface to support dynamic dispatch.
type IScenarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScenarioContext differentiates from other interfaces.
	IsScenarioContext()
}

type ScenarioContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScenarioContext() *ScenarioContext {
	var p = new(ScenarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_scenario
	return p
}

func (*ScenarioContext) IsScenarioContext() {}

func NewScenarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScenarioContext {
	var p = new(ScenarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_scenario

	return p
}

func (s *ScenarioContext) GetParser() antlr.Parser { return s.parser }

func (s *ScenarioContext) Scenario_keyword() IScenario_keywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScenario_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScenario_keywordContext)
}

func (s *ScenarioContext) Line_to_eol() ILine_to_eolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILine_to_eolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILine_to_eolContext)
}

func (s *ScenarioContext) Steps() IStepsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepsContext)
}

func (s *ScenarioContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *ScenarioContext) Tags() ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *ScenarioContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSPACE)
}

func (s *ScenarioContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSPACE, i)
}

func (s *ScenarioContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *ScenarioContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *ScenarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScenarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScenarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterScenario(s)
	}
}

func (s *ScenarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitScenario(s)
	}
}

func (s *ScenarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitScenario(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Scenario() (localctx IScenarioContext) {
	localctx = NewScenarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FeatureParserRULE_scenario)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(111)
			p.Comment()
		}

	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FeatureParserT__1 {
		{
			p.SetState(114)
			p.Tags()
		}

	}
	{
		p.SetState(117)
		p.Scenario_keyword()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(118)
				p.Match(FeatureParserSPACE)
			}

		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	{
		p.SetState(124)
		p.Line_to_eol()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(125)
				p.Match(FeatureParserNEWLINE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	{
		p.SetState(130)
		p.Steps()
	}

	return localctx
}

// IScenario_outlineContext is an interface to support dynamic dispatch.
type IScenario_outlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScenario_outlineContext differentiates from other interfaces.
	IsScenario_outlineContext()
}

type Scenario_outlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScenario_outlineContext() *Scenario_outlineContext {
	var p = new(Scenario_outlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_scenario_outline
	return p
}

func (*Scenario_outlineContext) IsScenario_outlineContext() {}

func NewScenario_outlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scenario_outlineContext {
	var p = new(Scenario_outlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_scenario_outline

	return p
}

func (s *Scenario_outlineContext) GetParser() antlr.Parser { return s.parser }

func (s *Scenario_outlineContext) Scenario_outline_keyword() IScenario_outline_keywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScenario_outline_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScenario_outline_keywordContext)
}

func (s *Scenario_outlineContext) Line_to_eol() ILine_to_eolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILine_to_eolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILine_to_eolContext)
}

func (s *Scenario_outlineContext) Steps() IStepsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepsContext)
}

func (s *Scenario_outlineContext) Examples_sections() IExamples_sectionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExamples_sectionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExamples_sectionsContext)
}

func (s *Scenario_outlineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *Scenario_outlineContext) Tags() ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *Scenario_outlineContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSPACE)
}

func (s *Scenario_outlineContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSPACE, i)
}

func (s *Scenario_outlineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *Scenario_outlineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *Scenario_outlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scenario_outlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scenario_outlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterScenario_outline(s)
	}
}

func (s *Scenario_outlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitScenario_outline(s)
	}
}

func (s *Scenario_outlineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitScenario_outline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Scenario_outline() (localctx IScenario_outlineContext) {
	localctx = NewScenario_outlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FeatureParserRULE_scenario_outline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(132)
			p.Comment()
		}

	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FeatureParserT__1 {
		{
			p.SetState(135)
			p.Tags()
		}

	}
	{
		p.SetState(138)
		p.Scenario_outline_keyword()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(139)
				p.Match(FeatureParserSPACE)
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}
	{
		p.SetState(145)
		p.Line_to_eol()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FeatureParserNEWLINE {
		{
			p.SetState(146)
			p.Match(FeatureParserNEWLINE)
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
		p.Steps()
	}
	{
		p.SetState(152)
		p.Examples_sections()
	}

	return localctx
}

// IStepsContext is an interface to support dynamic dispatch.
type IStepsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepsContext differentiates from other interfaces.
	IsStepsContext()
}

type StepsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepsContext() *StepsContext {
	var p = new(StepsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_steps
	return p
}

func (*StepsContext) IsStepsContext() {}

func NewStepsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepsContext {
	var p = new(StepsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_steps

	return p
}

func (s *StepsContext) GetParser() antlr.Parser { return s.parser }

func (s *StepsContext) AllStep() []IStepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStepContext)(nil)).Elem())
	var tst = make([]IStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStepContext)
		}
	}

	return tst
}

func (s *StepsContext) Step(i int) IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *StepsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterSteps(s)
	}
}

func (s *StepsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitSteps(s)
	}
}

func (s *StepsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitSteps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Steps() (localctx IStepsContext) {
	localctx = NewStepsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FeatureParserRULE_steps)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(154)
				p.Step()
			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IStepContext is an interface to support dynamic dispatch.
type IStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepContext differentiates from other interfaces.
	IsStepContext()
}

type StepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepContext() *StepContext {
	var p = new(StepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_step
	return p
}

func (*StepContext) IsStepContext() {}

func NewStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepContext {
	var p = new(StepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_step

	return p
}

func (s *StepContext) GetParser() antlr.Parser { return s.parser }

func (s *StepContext) Step_keyword() IStep_keywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStep_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStep_keywordContext)
}

func (s *StepContext) Line_to_eol() ILine_to_eolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILine_to_eolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILine_to_eolContext)
}

func (s *StepContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSPACE)
}

func (s *StepContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSPACE, i)
}

func (s *StepContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *StepContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *StepContext) Multiline_arg() IMultiline_argContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiline_argContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiline_argContext)
}

func (s *StepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterStep(s)
	}
}

func (s *StepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitStep(s)
	}
}

func (s *StepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitStep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Step() (localctx IStepContext) {
	localctx = NewStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FeatureParserRULE_step)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Step_keyword()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(161)
				p.Match(FeatureParserSPACE)
			}

		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	{
		p.SetState(167)
		p.Line_to_eol()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(168)
				p.Match(FeatureParserNEWLINE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(173)
			p.Multiline_arg()
		}

	}

	return localctx
}

// IExamples_sectionsContext is an interface to support dynamic dispatch.
type IExamples_sectionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExamples_sectionsContext differentiates from other interfaces.
	IsExamples_sectionsContext()
}

type Examples_sectionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExamples_sectionsContext() *Examples_sectionsContext {
	var p = new(Examples_sectionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_examples_sections
	return p
}

func (*Examples_sectionsContext) IsExamples_sectionsContext() {}

func NewExamples_sectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Examples_sectionsContext {
	var p = new(Examples_sectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_examples_sections

	return p
}

func (s *Examples_sectionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Examples_sectionsContext) AllExamples() []IExamplesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExamplesContext)(nil)).Elem())
	var tst = make([]IExamplesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExamplesContext)
		}
	}

	return tst
}

func (s *Examples_sectionsContext) Examples(i int) IExamplesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExamplesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExamplesContext)
}

func (s *Examples_sectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Examples_sectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Examples_sectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterExamples_sections(s)
	}
}

func (s *Examples_sectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitExamples_sections(s)
	}
}

func (s *Examples_sectionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitExamples_sections(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Examples_sections() (localctx IExamples_sectionsContext) {
	localctx = NewExamples_sectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FeatureParserRULE_examples_sections)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(176)
				p.Examples()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IExamplesContext is an interface to support dynamic dispatch.
type IExamplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExamplesContext differentiates from other interfaces.
	IsExamplesContext()
}

type ExamplesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExamplesContext() *ExamplesContext {
	var p = new(ExamplesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_examples
	return p
}

func (*ExamplesContext) IsExamplesContext() {}

func NewExamplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExamplesContext {
	var p = new(ExamplesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_examples

	return p
}

func (s *ExamplesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExamplesContext) Examples_keyword() IExamples_keywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExamples_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExamples_keywordContext)
}

func (s *ExamplesContext) Line_to_eol() ILine_to_eolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILine_to_eolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILine_to_eolContext)
}

func (s *ExamplesContext) Table() ITableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *ExamplesContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSPACE)
}

func (s *ExamplesContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSPACE, i)
}

func (s *ExamplesContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *ExamplesContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *ExamplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExamplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExamplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterExamples(s)
	}
}

func (s *ExamplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitExamples(s)
	}
}

func (s *ExamplesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitExamples(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Examples() (localctx IExamplesContext) {
	localctx = NewExamplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FeatureParserRULE_examples)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Examples_keyword()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(182)
				p.Match(FeatureParserSPACE)
			}

		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	{
		p.SetState(188)
		p.Line_to_eol()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FeatureParserNEWLINE {
		{
			p.SetState(189)
			p.Match(FeatureParserNEWLINE)
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Table()
	}

	return localctx
}

// IMultiline_argContext is an interface to support dynamic dispatch.
type IMultiline_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiline_argContext differentiates from other interfaces.
	IsMultiline_argContext()
}

type Multiline_argContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiline_argContext() *Multiline_argContext {
	var p = new(Multiline_argContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_multiline_arg
	return p
}

func (*Multiline_argContext) IsMultiline_argContext() {}

func NewMultiline_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Multiline_argContext {
	var p = new(Multiline_argContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_multiline_arg

	return p
}

func (s *Multiline_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Multiline_argContext) Table() ITableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *Multiline_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Multiline_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Multiline_argContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterMultiline_arg(s)
	}
}

func (s *Multiline_argContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitMultiline_arg(s)
	}
}

func (s *Multiline_argContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitMultiline_arg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Multiline_arg() (localctx IMultiline_argContext) {
	localctx = NewMultiline_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FeatureParserRULE_multiline_arg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Table()
	}

	return localctx
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_table
	return p
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) AllTable_row() []ITable_rowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITable_rowContext)(nil)).Elem())
	var tst = make([]ITable_rowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITable_rowContext)
		}
	}

	return tst
}

func (s *TableContext) Table_row(i int) ITable_rowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_rowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITable_rowContext)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitTable(s)
	}
}

func (s *TableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FeatureParserRULE_table)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(198)
				p.Table_row()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// ITable_rowContext is an interface to support dynamic dispatch.
type ITable_rowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_rowContext differentiates from other interfaces.
	IsTable_rowContext()
}

type Table_rowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_rowContext() *Table_rowContext {
	var p = new(Table_rowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_table_row
	return p
}

func (*Table_rowContext) IsTable_rowContext() {}

func NewTable_rowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_rowContext {
	var p = new(Table_rowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_table_row

	return p
}

func (s *Table_rowContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_rowContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSPACE)
}

func (s *Table_rowContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSPACE, i)
}

func (s *Table_rowContext) AllCell() []ICellContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICellContext)(nil)).Elem())
	var tst = make([]ICellContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICellContext)
		}
	}

	return tst
}

func (s *Table_rowContext) Cell(i int) ICellContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICellContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICellContext)
}

func (s *Table_rowContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *Table_rowContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *Table_rowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_rowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_rowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterTable_row(s)
	}
}

func (s *Table_rowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitTable_row(s)
	}
}

func (s *Table_rowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitTable_row(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Table_row() (localctx ITable_rowContext) {
	localctx = NewTable_rowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FeatureParserRULE_table_row)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSPACE {
		{
			p.SetState(203)
			p.Match(FeatureParserSPACE)
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(209)
		p.Match(FeatureParserT__0)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(210)
				p.Cell()
			}
			{
				p.SetState(211)
				p.Match(FeatureParserT__0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSPACE {
		{
			p.SetState(217)
			p.Match(FeatureParserSPACE)
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(223)
				p.Match(FeatureParserNEWLINE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// ICellContext is an interface to support dynamic dispatch.
type ICellContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCellContext differentiates from other interfaces.
	IsCellContext()
}

type CellContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCellContext() *CellContext {
	var p = new(CellContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_cell
	return p
}

func (*CellContext) IsCellContext() {}

func NewCellContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CellContext {
	var p = new(CellContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_cell

	return p
}

func (s *CellContext) GetParser() antlr.Parser { return s.parser }

func (s *CellContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *CellContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *CellContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CellContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CellContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterCell(s)
	}
}

func (s *CellContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitCell(s)
	}
}

func (s *CellContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitCell(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Cell() (localctx ICellContext) {
	localctx = NewCellContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FeatureParserRULE_cell)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserT__1)|(1<<FeatureParserT__2)|(1<<FeatureParserT__3)|(1<<FeatureParserT__4)|(1<<FeatureParserT__5)|(1<<FeatureParserT__6)|(1<<FeatureParserT__7)|(1<<FeatureParserT__8)|(1<<FeatureParserT__9)|(1<<FeatureParserT__10)|(1<<FeatureParserT__11)|(1<<FeatureParserT__12)|(1<<FeatureParserID)|(1<<FeatureParserSPACE))) != 0 {
		{
			p.SetState(228)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == FeatureParserT__0 || _la == FeatureParserNEWLINE {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(229)
		p.MatchWildcard()

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITagsContext is an interface to support dynamic dispatch.
type ITagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagsContext differentiates from other interfaces.
	IsTagsContext()
}

type TagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagsContext() *TagsContext {
	var p = new(TagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_tags
	return p
}

func (*TagsContext) IsTagsContext() {}

func NewTagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagsContext {
	var p = new(TagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_tags

	return p
}

func (s *TagsContext) GetParser() antlr.Parser { return s.parser }

func (s *TagsContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *TagsContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *TagsContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSPACE)
}

func (s *TagsContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSPACE, i)
}

func (s *TagsContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *TagsContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *TagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterTags(s)
	}
}

func (s *TagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitTags(s)
	}
}

func (s *TagsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitTags(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Tags() (localctx ITagsContext) {
	localctx = NewTagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FeatureParserRULE_tags)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Tag()
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSPACE {
		{
			p.SetState(236)
			p.Match(FeatureParserSPACE)
		}
		{
			p.SetState(237)
			p.Tag()
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(243)
				p.Match(FeatureParserNEWLINE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTag_name returns the tag_name token.
	GetTag_name() antlr.Token

	// SetTag_name sets the tag_name token.
	SetTag_name(antlr.Token)

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	tag_name antlr.Token
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) GetTag_name() antlr.Token { return s.tag_name }

func (s *TagContext) SetTag_name(v antlr.Token) { s.tag_name = v }

func (s *TagContext) ID() antlr.TerminalNode {
	return s.GetToken(FeatureParserID, 0)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitTag(s)
	}
}

func (s *TagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FeatureParserRULE_tag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(FeatureParserT__1)
	}
	{
		p.SetState(249)

		var _m = p.Match(FeatureParserID)

		localctx.(*TagContext).tag_name = _m
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) AllComment_line() []IComment_lineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComment_lineContext)(nil)).Elem())
	var tst = make([]IComment_lineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComment_lineContext)
		}
	}

	return tst
}

func (s *CommentContext) Comment_line(i int) IComment_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComment_lineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComment_lineContext)
}

func (s *CommentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *CommentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FeatureParserRULE_comment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserT__2 {
		{
			p.SetState(251)
			p.Comment_line()
		}
		{
			p.SetState(252)
			p.Match(FeatureParserNEWLINE)
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComment_lineContext is an interface to support dynamic dispatch.
type IComment_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComment_lineContext differentiates from other interfaces.
	IsComment_lineContext()
}

type Comment_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComment_lineContext() *Comment_lineContext {
	var p = new(Comment_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_comment_line
	return p
}

func (*Comment_lineContext) IsComment_lineContext() {}

func NewComment_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comment_lineContext {
	var p = new(Comment_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_comment_line

	return p
}

func (s *Comment_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Comment_lineContext) Line_to_eol() ILine_to_eolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILine_to_eolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILine_to_eolContext)
}

func (s *Comment_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comment_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comment_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterComment_line(s)
	}
}

func (s *Comment_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitComment_line(s)
	}
}

func (s *Comment_lineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitComment_line(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Comment_line() (localctx IComment_lineContext) {
	localctx = NewComment_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FeatureParserRULE_comment_line)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(FeatureParserT__2)
	}
	{
		p.SetState(260)
		p.Line_to_eol()
	}

	return localctx
}

// ILine_to_eolContext is an interface to support dynamic dispatch.
type ILine_to_eolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLine_to_eolContext differentiates from other interfaces.
	IsLine_to_eolContext()
}

type Line_to_eolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_to_eolContext() *Line_to_eolContext {
	var p = new(Line_to_eolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_line_to_eol
	return p
}

func (*Line_to_eolContext) IsLine_to_eolContext() {}

func NewLine_to_eolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_to_eolContext {
	var p = new(Line_to_eolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_line_to_eol

	return p
}

func (s *Line_to_eolContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_to_eolContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNEWLINE)
}

func (s *Line_to_eolContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNEWLINE, i)
}

func (s *Line_to_eolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_to_eolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_to_eolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterLine_to_eol(s)
	}
}

func (s *Line_to_eolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitLine_to_eol(s)
	}
}

func (s *Line_to_eolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitLine_to_eol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Line_to_eol() (localctx ILine_to_eolContext) {
	localctx = NewLine_to_eolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FeatureParserRULE_line_to_eol)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserT__0)|(1<<FeatureParserT__1)|(1<<FeatureParserT__2)|(1<<FeatureParserT__3)|(1<<FeatureParserT__4)|(1<<FeatureParserT__5)|(1<<FeatureParserT__6)|(1<<FeatureParserT__7)|(1<<FeatureParserT__8)|(1<<FeatureParserT__9)|(1<<FeatureParserT__10)|(1<<FeatureParserT__11)|(1<<FeatureParserT__12)|(1<<FeatureParserID)|(1<<FeatureParserSPACE))) != 0 {
		{
			p.SetState(262)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == FeatureParserNEWLINE {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFeature_keywordContext is an interface to support dynamic dispatch.
type IFeature_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFeature_keywordContext differentiates from other interfaces.
	IsFeature_keywordContext()
}

type Feature_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeature_keywordContext() *Feature_keywordContext {
	var p = new(Feature_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_feature_keyword
	return p
}

func (*Feature_keywordContext) IsFeature_keywordContext() {}

func NewFeature_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Feature_keywordContext {
	var p = new(Feature_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_feature_keyword

	return p
}

func (s *Feature_keywordContext) GetParser() antlr.Parser { return s.parser }
func (s *Feature_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Feature_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Feature_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterFeature_keyword(s)
	}
}

func (s *Feature_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitFeature_keyword(s)
	}
}

func (s *Feature_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitFeature_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Feature_keyword() (localctx IFeature_keywordContext) {
	localctx = NewFeature_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FeatureParserRULE_feature_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(FeatureParserT__3)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(269)
			p.Match(FeatureParserT__4)
		}

	}

	return localctx
}

// IScenario_keywordContext is an interface to support dynamic dispatch.
type IScenario_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScenario_keywordContext differentiates from other interfaces.
	IsScenario_keywordContext()
}

type Scenario_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScenario_keywordContext() *Scenario_keywordContext {
	var p = new(Scenario_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_scenario_keyword
	return p
}

func (*Scenario_keywordContext) IsScenario_keywordContext() {}

func NewScenario_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scenario_keywordContext {
	var p = new(Scenario_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_scenario_keyword

	return p
}

func (s *Scenario_keywordContext) GetParser() antlr.Parser { return s.parser }
func (s *Scenario_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scenario_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scenario_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterScenario_keyword(s)
	}
}

func (s *Scenario_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitScenario_keyword(s)
	}
}

func (s *Scenario_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitScenario_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Scenario_keyword() (localctx IScenario_keywordContext) {
	localctx = NewScenario_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FeatureParserRULE_scenario_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(FeatureParserT__5)
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(273)
			p.Match(FeatureParserT__4)
		}

	}

	return localctx
}

// IScenario_outline_keywordContext is an interface to support dynamic dispatch.
type IScenario_outline_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScenario_outline_keywordContext differentiates from other interfaces.
	IsScenario_outline_keywordContext()
}

type Scenario_outline_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScenario_outline_keywordContext() *Scenario_outline_keywordContext {
	var p = new(Scenario_outline_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_scenario_outline_keyword
	return p
}

func (*Scenario_outline_keywordContext) IsScenario_outline_keywordContext() {}

func NewScenario_outline_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scenario_outline_keywordContext {
	var p = new(Scenario_outline_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_scenario_outline_keyword

	return p
}

func (s *Scenario_outline_keywordContext) GetParser() antlr.Parser { return s.parser }
func (s *Scenario_outline_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scenario_outline_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scenario_outline_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterScenario_outline_keyword(s)
	}
}

func (s *Scenario_outline_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitScenario_outline_keyword(s)
	}
}

func (s *Scenario_outline_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitScenario_outline_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Scenario_outline_keyword() (localctx IScenario_outline_keywordContext) {
	localctx = NewScenario_outline_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FeatureParserRULE_scenario_outline_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(FeatureParserT__6)
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(277)
			p.Match(FeatureParserT__4)
		}

	}

	return localctx
}

// IStep_keywordContext is an interface to support dynamic dispatch.
type IStep_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStep_keywordContext differentiates from other interfaces.
	IsStep_keywordContext()
}

type Step_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStep_keywordContext() *Step_keywordContext {
	var p = new(Step_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_step_keyword
	return p
}

func (*Step_keywordContext) IsStep_keywordContext() {}

func NewStep_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Step_keywordContext {
	var p = new(Step_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_step_keyword

	return p
}

func (s *Step_keywordContext) GetParser() antlr.Parser { return s.parser }
func (s *Step_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Step_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Step_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterStep_keyword(s)
	}
}

func (s *Step_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitStep_keyword(s)
	}
}

func (s *Step_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitStep_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Step_keyword() (localctx IStep_keywordContext) {
	localctx = NewStep_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FeatureParserRULE_step_keyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserT__7)|(1<<FeatureParserT__8)|(1<<FeatureParserT__9)|(1<<FeatureParserT__10)|(1<<FeatureParserT__11))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExamples_keywordContext is an interface to support dynamic dispatch.
type IExamples_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExamples_keywordContext differentiates from other interfaces.
	IsExamples_keywordContext()
}

type Examples_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExamples_keywordContext() *Examples_keywordContext {
	var p = new(Examples_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_examples_keyword
	return p
}

func (*Examples_keywordContext) IsExamples_keywordContext() {}

func NewExamples_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Examples_keywordContext {
	var p = new(Examples_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_examples_keyword

	return p
}

func (s *Examples_keywordContext) GetParser() antlr.Parser { return s.parser }
func (s *Examples_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Examples_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Examples_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterExamples_keyword(s)
	}
}

func (s *Examples_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitExamples_keyword(s)
	}
}

func (s *Examples_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitExamples_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Examples_keyword() (localctx IExamples_keywordContext) {
	localctx = NewExamples_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FeatureParserRULE_examples_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(FeatureParserT__12)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(283)
			p.Match(FeatureParserT__4)
		}

	}

	return localctx
}
