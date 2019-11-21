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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 380,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 3, 2, 7, 2, 60, 10, 2, 12, 2, 14, 2, 63, 11, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 7, 3, 69, 10, 3, 12, 3, 14, 3, 72, 11, 3, 3, 3, 7, 3, 75, 10,
	3, 12, 3, 14, 3, 78, 11, 3, 3, 3, 3, 3, 7, 3, 82, 10, 3, 12, 3, 14, 3,
	85, 11, 3, 3, 3, 3, 3, 6, 3, 89, 10, 3, 13, 3, 14, 3, 90, 3, 4, 5, 4, 94,
	10, 4, 3, 4, 6, 4, 97, 10, 4, 13, 4, 14, 4, 98, 3, 5, 7, 5, 102, 10, 5,
	12, 5, 14, 5, 105, 11, 5, 3, 5, 7, 5, 108, 10, 5, 12, 5, 14, 5, 111, 11,
	5, 3, 5, 3, 5, 7, 5, 115, 10, 5, 12, 5, 14, 5, 118, 11, 5, 3, 5, 5, 5,
	121, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	7, 6, 133, 10, 6, 12, 6, 14, 6, 136, 11, 6, 3, 7, 7, 7, 139, 10, 7, 12,
	7, 14, 7, 142, 11, 7, 3, 7, 7, 7, 145, 10, 7, 12, 7, 14, 7, 148, 11, 7,
	3, 7, 7, 7, 151, 10, 7, 12, 7, 14, 7, 154, 11, 7, 3, 7, 3, 7, 7, 7, 158,
	10, 7, 12, 7, 14, 7, 161, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 7, 8, 168,
	10, 8, 12, 8, 14, 8, 171, 11, 8, 3, 8, 3, 8, 3, 8, 5, 8, 176, 10, 8, 3,
	8, 7, 8, 179, 10, 8, 12, 8, 14, 8, 182, 11, 8, 3, 8, 3, 8, 3, 9, 7, 9,
	187, 10, 9, 12, 9, 14, 9, 190, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11,
	7, 11, 197, 10, 11, 12, 11, 14, 11, 200, 11, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 7, 12, 206, 10, 12, 12, 12, 14, 12, 209, 11, 12, 3, 12, 3, 12, 3, 12,
	3, 13, 7, 13, 215, 10, 13, 12, 13, 14, 13, 218, 11, 13, 3, 13, 3, 13, 3,
	13, 3, 14, 7, 14, 224, 10, 14, 12, 14, 14, 14, 227, 11, 14, 3, 14, 3, 14,
	3, 14, 3, 15, 7, 15, 233, 10, 15, 12, 15, 14, 15, 236, 11, 15, 3, 15, 3,
	15, 3, 15, 3, 16, 7, 16, 242, 10, 16, 12, 16, 14, 16, 245, 11, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 7, 17, 251, 10, 17, 12, 17, 14, 17, 254, 11, 17, 3,
	17, 3, 17, 5, 17, 258, 10, 17, 3, 18, 3, 18, 6, 18, 262, 10, 18, 13, 18,
	14, 18, 263, 3, 18, 5, 18, 267, 10, 18, 3, 19, 3, 19, 3, 19, 7, 19, 272,
	10, 19, 12, 19, 14, 19, 275, 11, 19, 3, 20, 3, 20, 7, 20, 279, 10, 20,
	12, 20, 14, 20, 282, 11, 20, 3, 21, 7, 21, 285, 10, 21, 12, 21, 14, 21,
	288, 11, 21, 3, 21, 3, 21, 6, 21, 292, 10, 21, 13, 21, 14, 21, 293, 3,
	21, 6, 21, 297, 10, 21, 13, 21, 14, 21, 298, 3, 21, 5, 21, 302, 10, 21,
	3, 22, 7, 22, 305, 10, 22, 12, 22, 14, 22, 308, 11, 22, 3, 22, 3, 22, 6,
	22, 312, 10, 22, 13, 22, 14, 22, 313, 3, 22, 6, 22, 317, 10, 22, 13, 22,
	14, 22, 318, 3, 22, 5, 22, 322, 10, 22, 3, 23, 7, 23, 325, 10, 23, 12,
	23, 14, 23, 328, 11, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 7, 25, 339, 10, 25, 12, 25, 14, 25, 342, 11, 25, 3, 26, 3,
	26, 7, 26, 346, 10, 26, 12, 26, 14, 26, 349, 11, 26, 3, 27, 3, 27, 7, 27,
	353, 10, 27, 12, 27, 14, 27, 356, 11, 27, 3, 28, 7, 28, 359, 10, 28, 12,
	28, 14, 28, 362, 11, 28, 3, 28, 3, 28, 7, 28, 366, 10, 28, 12, 28, 14,
	28, 369, 11, 28, 3, 28, 3, 28, 3, 28, 3, 29, 7, 29, 375, 10, 29, 12, 29,
	14, 29, 378, 11, 29, 3, 29, 4, 188, 376, 2, 30, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 2, 8, 3, 2, 15, 16, 3, 3, 16, 16, 4, 2, 18, 18, 22, 22, 4,
	2, 15, 15, 18, 22, 5, 2, 15, 15, 17, 20, 22, 22, 4, 2, 15, 15, 17, 22,
	2, 405, 2, 61, 3, 2, 2, 2, 4, 70, 3, 2, 2, 2, 6, 93, 3, 2, 2, 2, 8, 103,
	3, 2, 2, 2, 10, 134, 3, 2, 2, 2, 12, 140, 3, 2, 2, 2, 14, 169, 3, 2, 2,
	2, 16, 188, 3, 2, 2, 2, 18, 191, 3, 2, 2, 2, 20, 198, 3, 2, 2, 2, 22, 207,
	3, 2, 2, 2, 24, 216, 3, 2, 2, 2, 26, 225, 3, 2, 2, 2, 28, 234, 3, 2, 2,
	2, 30, 243, 3, 2, 2, 2, 32, 252, 3, 2, 2, 2, 34, 259, 3, 2, 2, 2, 36, 273,
	3, 2, 2, 2, 38, 276, 3, 2, 2, 2, 40, 286, 3, 2, 2, 2, 42, 306, 3, 2, 2,
	2, 44, 326, 3, 2, 2, 2, 46, 332, 3, 2, 2, 2, 48, 336, 3, 2, 2, 2, 50, 343,
	3, 2, 2, 2, 52, 350, 3, 2, 2, 2, 54, 360, 3, 2, 2, 2, 56, 376, 3, 2, 2,
	2, 58, 60, 5, 54, 28, 2, 59, 58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59,
	3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2,
	64, 65, 5, 4, 3, 2, 65, 66, 5, 6, 4, 2, 66, 3, 3, 2, 2, 2, 67, 69, 9, 2,
	2, 2, 68, 67, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71,
	3, 2, 2, 2, 71, 76, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 75, 5, 14, 8, 2,
	74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3,
	2, 2, 2, 77, 79, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 79, 83, 7, 14, 2, 2, 80,
	82, 7, 15, 2, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2,
	2, 2, 83, 84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 88,
	5, 52, 27, 2, 87, 89, 7, 16, 2, 2, 88, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2,
	2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 5, 3, 2, 2, 2, 92, 94, 5,
	8, 5, 2, 93, 92, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 96, 3, 2, 2, 2, 95,
	97, 5, 12, 7, 2, 96, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 96, 3, 2,
	2, 2, 98, 99, 3, 2, 2, 2, 99, 7, 3, 2, 2, 2, 100, 102, 9, 2, 2, 2, 101,
	100, 3, 2, 2, 2, 102, 105, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104,
	3, 2, 2, 2, 104, 109, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 106, 108, 5, 14,
	8, 2, 107, 106, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2,
	109, 110, 3, 2, 2, 2, 110, 112, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112,
	116, 7, 12, 2, 2, 113, 115, 7, 15, 2, 2, 114, 113, 3, 2, 2, 2, 115, 118,
	3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 120, 3, 2,
	2, 2, 118, 116, 3, 2, 2, 2, 119, 121, 5, 52, 27, 2, 120, 119, 3, 2, 2,
	2, 120, 121, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 9, 3, 2, 2, 123,
	124, 5, 10, 6, 2, 124, 125, 9, 3, 2, 2, 125, 9, 3, 2, 2, 2, 126, 133, 5,
	20, 11, 2, 127, 133, 5, 22, 12, 2, 128, 133, 5, 24, 13, 2, 129, 133, 5,
	28, 15, 2, 130, 133, 5, 26, 14, 2, 131, 133, 5, 30, 16, 2, 132, 126, 3,
	2, 2, 2, 132, 127, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2, 132, 129, 3, 2, 2,
	2, 132, 130, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134,
	132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 11, 3, 2, 2, 2, 136, 134, 3,
	2, 2, 2, 137, 139, 9, 2, 2, 2, 138, 137, 3, 2, 2, 2, 139, 142, 3, 2, 2,
	2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 146, 3, 2, 2, 2, 142,
	140, 3, 2, 2, 2, 143, 145, 5, 14, 8, 2, 144, 143, 3, 2, 2, 2, 145, 148,
	3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 152, 3, 2,
	2, 2, 148, 146, 3, 2, 2, 2, 149, 151, 7, 15, 2, 2, 150, 149, 3, 2, 2, 2,
	151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153,
	155, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 159, 7, 13, 2, 2, 156, 158,
	7, 15, 2, 2, 157, 156, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2,
	2, 2, 159, 160, 3, 2, 2, 2, 160, 162, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2,
	162, 163, 5, 52, 27, 2, 163, 164, 9, 3, 2, 2, 164, 165, 5, 10, 6, 2, 165,
	13, 3, 2, 2, 2, 166, 168, 9, 2, 2, 2, 167, 166, 3, 2, 2, 2, 168, 171, 3,
	2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 172, 3, 2, 2,
	2, 171, 169, 3, 2, 2, 2, 172, 173, 7, 20, 2, 2, 173, 175, 5, 16, 9, 2,
	174, 176, 5, 18, 10, 2, 175, 174, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176,
	180, 3, 2, 2, 2, 177, 179, 7, 15, 2, 2, 178, 177, 3, 2, 2, 2, 179, 182,
	3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2,
	2, 2, 182, 180, 3, 2, 2, 2, 183, 184, 7, 16, 2, 2, 184, 15, 3, 2, 2, 2,
	185, 187, 11, 2, 2, 2, 186, 185, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188,
	189, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 17, 3, 2, 2, 2, 190, 188, 3,
	2, 2, 2, 191, 192, 7, 18, 2, 2, 192, 193, 5, 52, 27, 2, 193, 194, 7, 19,
	2, 2, 194, 19, 3, 2, 2, 2, 195, 197, 9, 2, 2, 2, 196, 195, 3, 2, 2, 2,
	197, 200, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199,
	201, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 201, 202, 7, 8, 2, 2, 202, 203,
	5, 32, 17, 2, 203, 21, 3, 2, 2, 2, 204, 206, 9, 2, 2, 2, 205, 204, 3, 2,
	2, 2, 206, 209, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2,
	208, 210, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 210, 211, 7, 9, 2, 2, 211,
	212, 5, 32, 17, 2, 212, 23, 3, 2, 2, 2, 213, 215, 9, 2, 2, 2, 214, 213,
	3, 2, 2, 2, 215, 218, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216, 217, 3, 2,
	2, 2, 217, 219, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 219, 220, 7, 7, 2, 2,
	220, 221, 5, 32, 17, 2, 221, 25, 3, 2, 2, 2, 222, 224, 9, 2, 2, 2, 223,
	222, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226,
	3, 2, 2, 2, 226, 228, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 229, 7, 6,
	2, 2, 229, 230, 5, 32, 17, 2, 230, 27, 3, 2, 2, 2, 231, 233, 9, 2, 2, 2,
	232, 231, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234,
	235, 3, 2, 2, 2, 235, 237, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 238,
	7, 10, 2, 2, 238, 239, 5, 32, 17, 2, 239, 29, 3, 2, 2, 2, 240, 242, 9,
	2, 2, 2, 241, 240, 3, 2, 2, 2, 242, 245, 3, 2, 2, 2, 243, 241, 3, 2, 2,
	2, 243, 244, 3, 2, 2, 2, 244, 246, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246,
	247, 7, 11, 2, 2, 247, 248, 5, 32, 17, 2, 248, 31, 3, 2, 2, 2, 249, 251,
	7, 15, 2, 2, 250, 249, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 250, 3, 2,
	2, 2, 252, 253, 3, 2, 2, 2, 253, 255, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2,
	255, 257, 5, 34, 18, 2, 256, 258, 5, 38, 20, 2, 257, 256, 3, 2, 2, 2, 257,
	258, 3, 2, 2, 2, 258, 33, 3, 2, 2, 2, 259, 266, 5, 36, 19, 2, 260, 262,
	7, 16, 2, 2, 261, 260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 261, 3, 2,
	2, 2, 263, 264, 3, 2, 2, 2, 264, 267, 3, 2, 2, 2, 265, 267, 7, 2, 2, 3,
	266, 261, 3, 2, 2, 2, 266, 265, 3, 2, 2, 2, 267, 35, 3, 2, 2, 2, 268, 272,
	5, 48, 25, 2, 269, 272, 7, 15, 2, 2, 270, 272, 5, 46, 24, 2, 271, 268,
	3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 270, 3, 2, 2, 2, 272, 275, 3, 2,
	2, 2, 273, 271, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 37, 3, 2, 2, 2,
	275, 273, 3, 2, 2, 2, 276, 280, 5, 40, 21, 2, 277, 279, 5, 42, 22, 2, 278,
	277, 3, 2, 2, 2, 279, 282, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 280, 281,
	3, 2, 2, 2, 281, 39, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 283, 285, 7, 15,
	2, 2, 284, 283, 3, 2, 2, 2, 285, 288, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2,
	286, 287, 3, 2, 2, 2, 287, 289, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 289,
	291, 7, 21, 2, 2, 290, 292, 5, 44, 23, 2, 291, 290, 3, 2, 2, 2, 292, 293,
	3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 301, 3, 2,
	2, 2, 295, 297, 7, 16, 2, 2, 296, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2,
	298, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 302, 3, 2, 2, 2, 300,
	302, 7, 2, 2, 3, 301, 296, 3, 2, 2, 2, 301, 300, 3, 2, 2, 2, 302, 41, 3,
	2, 2, 2, 303, 305, 7, 15, 2, 2, 304, 303, 3, 2, 2, 2, 305, 308, 3, 2, 2,
	2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 309, 3, 2, 2, 2, 308,
	306, 3, 2, 2, 2, 309, 311, 7, 21, 2, 2, 310, 312, 5, 44, 23, 2, 311, 310,
	3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2,
	2, 2, 314, 321, 3, 2, 2, 2, 315, 317, 7, 16, 2, 2, 316, 315, 3, 2, 2, 2,
	317, 318, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319,
	322, 3, 2, 2, 2, 320, 322, 7, 2, 2, 3, 321, 316, 3, 2, 2, 2, 321, 320,
	3, 2, 2, 2, 322, 43, 3, 2, 2, 2, 323, 325, 7, 15, 2, 2, 324, 323, 3, 2,
	2, 2, 325, 328, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2,
	327, 329, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 329, 330, 5, 50, 26, 2, 330,
	331, 7, 21, 2, 2, 331, 45, 3, 2, 2, 2, 332, 333, 7, 17, 2, 2, 333, 334,
	5, 16, 9, 2, 334, 335, 7, 17, 2, 2, 335, 47, 3, 2, 2, 2, 336, 340, 9, 4,
	2, 2, 337, 339, 9, 5, 2, 2, 338, 337, 3, 2, 2, 2, 339, 342, 3, 2, 2, 2,
	340, 338, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 49, 3, 2, 2, 2, 342, 340,
	3, 2, 2, 2, 343, 347, 9, 4, 2, 2, 344, 346, 9, 6, 2, 2, 345, 344, 3, 2,
	2, 2, 346, 349, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2,
	348, 51, 3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 350, 354, 9, 4, 2, 2, 351, 353,
	9, 7, 2, 2, 352, 351, 3, 2, 2, 2, 353, 356, 3, 2, 2, 2, 354, 352, 3, 2,
	2, 2, 354, 355, 3, 2, 2, 2, 355, 53, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2,
	357, 359, 7, 15, 2, 2, 358, 357, 3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360,
	358, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 363, 3, 2, 2, 2, 362, 360,
	3, 2, 2, 2, 363, 367, 7, 3, 2, 2, 364, 366, 7, 15, 2, 2, 365, 364, 3, 2,
	2, 2, 366, 369, 3, 2, 2, 2, 367, 365, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2,
	368, 370, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 370, 371, 5, 56, 29, 2, 371,
	372, 7, 16, 2, 2, 372, 55, 3, 2, 2, 2, 373, 375, 11, 2, 2, 2, 374, 373,
	3, 2, 2, 2, 375, 378, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 376, 374, 3, 2,
	2, 2, 377, 57, 3, 2, 2, 2, 378, 376, 3, 2, 2, 2, 51, 61, 70, 76, 83, 90,
	93, 98, 103, 109, 116, 120, 132, 134, 140, 146, 152, 159, 169, 175, 180,
	188, 198, 207, 216, 225, 234, 243, 252, 257, 263, 266, 271, 273, 280, 286,
	293, 298, 301, 306, 313, 318, 321, 326, 340, 347, 354, 360, 367, 376,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'#'", "", "", "", "", "", "", "", "", "", "", "", "", "", "'\"'",
	"'('", "')'", "'@'", "'|'",
}
var symbolicNames = []string{
	"", "", "IDENTIFIER", "EmptyLine", "And", "Or", "Given", "When", "Then",
	"Examples", "Background", "Scenario", "Feature", "Space", "NewLine", "Quote",
	"LBracket", "RBracket", "At", "Pipe", "Char",
}

var ruleNames = []string{
	"feature", "featureHeader", "featureBody", "background", "blockBody", "scenario",
	"tags", "anyText", "value", "given", "when", "or", "and", "then", "example",
	"step", "stepContent", "stepText", "table", "tableHeader", "row", "cell",
	"parameter", "contentNoQuotes", "contentNoPipes", "content", "comment",
	"commentText",
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
	FeatureParserEOF        = antlr.TokenEOF
	FeatureParserT__0       = 1
	FeatureParserIDENTIFIER = 2
	FeatureParserEmptyLine  = 3
	FeatureParserAnd        = 4
	FeatureParserOr         = 5
	FeatureParserGiven      = 6
	FeatureParserWhen       = 7
	FeatureParserThen       = 8
	FeatureParserExamples   = 9
	FeatureParserBackground = 10
	FeatureParserScenario   = 11
	FeatureParserFeature    = 12
	FeatureParserSpace      = 13
	FeatureParserNewLine    = 14
	FeatureParserQuote      = 15
	FeatureParserLBracket   = 16
	FeatureParserRBracket   = 17
	FeatureParserAt         = 18
	FeatureParserPipe       = 19
	FeatureParserChar       = 20
)

// FeatureParser rules.
const (
	FeatureParserRULE_feature         = 0
	FeatureParserRULE_featureHeader   = 1
	FeatureParserRULE_featureBody     = 2
	FeatureParserRULE_background      = 3
	FeatureParserRULE_blockBody       = 4
	FeatureParserRULE_scenario        = 5
	FeatureParserRULE_tags            = 6
	FeatureParserRULE_anyText         = 7
	FeatureParserRULE_value           = 8
	FeatureParserRULE_given           = 9
	FeatureParserRULE_when            = 10
	FeatureParserRULE_or              = 11
	FeatureParserRULE_and             = 12
	FeatureParserRULE_then            = 13
	FeatureParserRULE_example         = 14
	FeatureParserRULE_step            = 15
	FeatureParserRULE_stepContent     = 16
	FeatureParserRULE_stepText        = 17
	FeatureParserRULE_table           = 18
	FeatureParserRULE_tableHeader     = 19
	FeatureParserRULE_row             = 20
	FeatureParserRULE_cell            = 21
	FeatureParserRULE_parameter       = 22
	FeatureParserRULE_contentNoQuotes = 23
	FeatureParserRULE_contentNoPipes  = 24
	FeatureParserRULE_content         = 25
	FeatureParserRULE_comment         = 26
	FeatureParserRULE_commentText     = 27
)

// IFeatureContext is an interface to support dynamic dispatch.
type IFeatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFeatureContext differentiates from other interfaces.
	IsFeatureContext()
}

type FeatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeatureContext() *FeatureContext {
	var p = new(FeatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_feature
	return p
}

func (*FeatureContext) IsFeatureContext() {}

func NewFeatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeatureContext {
	var p = new(FeatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_feature

	return p
}

func (s *FeatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FeatureContext) FeatureHeader() IFeatureHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFeatureHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFeatureHeaderContext)
}

func (s *FeatureContext) FeatureBody() IFeatureBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFeatureBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFeatureBodyContext)
}

func (s *FeatureContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *FeatureContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *FeatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterFeature(s)
	}
}

func (s *FeatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitFeature(s)
	}
}

func (s *FeatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitFeature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Feature() (localctx IFeatureContext) {
	localctx = NewFeatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FeatureParserRULE_feature)

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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(56)
				p.Comment()
			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(62)
		p.FeatureHeader()
	}
	{
		p.SetState(63)
		p.FeatureBody()
	}

	return localctx
}

// IFeatureHeaderContext is an interface to support dynamic dispatch.
type IFeatureHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFeatureHeaderContext differentiates from other interfaces.
	IsFeatureHeaderContext()
}

type FeatureHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeatureHeaderContext() *FeatureHeaderContext {
	var p = new(FeatureHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_featureHeader
	return p
}

func (*FeatureHeaderContext) IsFeatureHeaderContext() {}

func NewFeatureHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeatureHeaderContext {
	var p = new(FeatureHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_featureHeader

	return p
}

func (s *FeatureHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *FeatureHeaderContext) Feature() antlr.TerminalNode {
	return s.GetToken(FeatureParserFeature, 0)
}

func (s *FeatureHeaderContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *FeatureHeaderContext) AllTags() []ITagsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagsContext)(nil)).Elem())
	var tst = make([]ITagsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagsContext)
		}
	}

	return tst
}

func (s *FeatureHeaderContext) Tags(i int) ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *FeatureHeaderContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *FeatureHeaderContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *FeatureHeaderContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *FeatureHeaderContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *FeatureHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeatureHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeatureHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterFeatureHeader(s)
	}
}

func (s *FeatureHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitFeatureHeader(s)
	}
}

func (s *FeatureHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitFeatureHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) FeatureHeader() (localctx IFeatureHeaderContext) {
	localctx = NewFeatureHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FeatureParserRULE_featureHeader)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(65)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserNewLine)|(1<<FeatureParserAt))) != 0 {
		{
			p.SetState(71)
			p.Tags()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(FeatureParserFeature)
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(78)
			p.Match(FeatureParserSpace)
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(84)
		p.Content()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(85)
				p.Match(FeatureParserNewLine)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IFeatureBodyContext is an interface to support dynamic dispatch.
type IFeatureBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFeatureBodyContext differentiates from other interfaces.
	IsFeatureBodyContext()
}

type FeatureBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeatureBodyContext() *FeatureBodyContext {
	var p = new(FeatureBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_featureBody
	return p
}

func (*FeatureBodyContext) IsFeatureBodyContext() {}

func NewFeatureBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeatureBodyContext {
	var p = new(FeatureBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_featureBody

	return p
}

func (s *FeatureBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FeatureBodyContext) Background() IBackgroundContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBackgroundContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBackgroundContext)
}

func (s *FeatureBodyContext) AllScenario() []IScenarioContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScenarioContext)(nil)).Elem())
	var tst = make([]IScenarioContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScenarioContext)
		}
	}

	return tst
}

func (s *FeatureBodyContext) Scenario(i int) IScenarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScenarioContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScenarioContext)
}

func (s *FeatureBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeatureBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeatureBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterFeatureBody(s)
	}
}

func (s *FeatureBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitFeatureBody(s)
	}
}

func (s *FeatureBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitFeatureBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) FeatureBody() (localctx IFeatureBodyContext) {
	localctx = NewFeatureBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FeatureParserRULE_featureBody)
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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(90)
			p.Background()
		}

	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserScenario)|(1<<FeatureParserSpace)|(1<<FeatureParserNewLine)|(1<<FeatureParserAt))) != 0) {
		{
			p.SetState(93)
			p.Scenario()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBackgroundContext is an interface to support dynamic dispatch.
type IBackgroundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBackgroundContext differentiates from other interfaces.
	IsBackgroundContext()
}

type BackgroundContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBackgroundContext() *BackgroundContext {
	var p = new(BackgroundContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_background
	return p
}

func (*BackgroundContext) IsBackgroundContext() {}

func NewBackgroundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BackgroundContext {
	var p = new(BackgroundContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_background

	return p
}

func (s *BackgroundContext) GetParser() antlr.Parser { return s.parser }

func (s *BackgroundContext) Background() antlr.TerminalNode {
	return s.GetToken(FeatureParserBackground, 0)
}

func (s *BackgroundContext) BlockBody() IBlockBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *BackgroundContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *BackgroundContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *BackgroundContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserEOF)
}

func (s *BackgroundContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserEOF, i)
}

func (s *BackgroundContext) AllTags() []ITagsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagsContext)(nil)).Elem())
	var tst = make([]ITagsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagsContext)
		}
	}

	return tst
}

func (s *BackgroundContext) Tags(i int) ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *BackgroundContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *BackgroundContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *BackgroundContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *BackgroundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BackgroundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BackgroundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterBackground(s)
	}
}

func (s *BackgroundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitBackground(s)
	}
}

func (s *BackgroundContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitBackground(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Background() (localctx IBackgroundContext) {
	localctx = NewBackgroundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FeatureParserRULE_background)
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(98)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserNewLine)|(1<<FeatureParserAt))) != 0 {
		{
			p.SetState(104)
			p.Tags()
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)
		p.Match(FeatureParserBackground)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(111)
			p.Match(FeatureParserSpace)
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FeatureParserLBracket || _la == FeatureParserChar {
		{
			p.SetState(117)
			p.Content()
		}

	}
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserEOF || _la == FeatureParserNewLine) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(121)
		p.BlockBody()
	}
	{
		p.SetState(122)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserEOF || _la == FeatureParserNewLine) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBlockBodyContext is an interface to support dynamic dispatch.
type IBlockBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockBodyContext differentiates from other interfaces.
	IsBlockBodyContext()
}

type BlockBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockBodyContext() *BlockBodyContext {
	var p = new(BlockBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_blockBody
	return p
}

func (*BlockBodyContext) IsBlockBodyContext() {}

func NewBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockBodyContext {
	var p = new(BlockBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_blockBody

	return p
}

func (s *BlockBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockBodyContext) AllGiven() []IGivenContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGivenContext)(nil)).Elem())
	var tst = make([]IGivenContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGivenContext)
		}
	}

	return tst
}

func (s *BlockBodyContext) Given(i int) IGivenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGivenContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGivenContext)
}

func (s *BlockBodyContext) AllWhen() []IWhenContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhenContext)(nil)).Elem())
	var tst = make([]IWhenContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhenContext)
		}
	}

	return tst
}

func (s *BlockBodyContext) When(i int) IWhenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhenContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhenContext)
}

func (s *BlockBodyContext) AllOr() []IOrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrContext)(nil)).Elem())
	var tst = make([]IOrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrContext)
		}
	}

	return tst
}

func (s *BlockBodyContext) Or(i int) IOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrContext)
}

func (s *BlockBodyContext) AllThen() []IThenContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThenContext)(nil)).Elem())
	var tst = make([]IThenContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThenContext)
		}
	}

	return tst
}

func (s *BlockBodyContext) Then(i int) IThenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThenContext)
}

func (s *BlockBodyContext) AllAnd() []IAndContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAndContext)(nil)).Elem())
	var tst = make([]IAndContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAndContext)
		}
	}

	return tst
}

func (s *BlockBodyContext) And(i int) IAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAndContext)
}

func (s *BlockBodyContext) AllExample() []IExampleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExampleContext)(nil)).Elem())
	var tst = make([]IExampleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExampleContext)
		}
	}

	return tst
}

func (s *BlockBodyContext) Example(i int) IExampleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExampleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExampleContext)
}

func (s *BlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterBlockBody(s)
	}
}

func (s *BlockBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitBlockBody(s)
	}
}

func (s *BlockBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitBlockBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) BlockBody() (localctx IBlockBodyContext) {
	localctx = NewBlockBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FeatureParserRULE_blockBody)

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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(124)
					p.Given()
				}

			case 2:
				{
					p.SetState(125)
					p.When()
				}

			case 3:
				{
					p.SetState(126)
					p.Or()
				}

			case 4:
				{
					p.SetState(127)
					p.Then()
				}

			case 5:
				{
					p.SetState(128)
					p.And()
				}

			case 6:
				{
					p.SetState(129)
					p.Example()
				}

			}

		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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

func (s *ScenarioContext) Scenario() antlr.TerminalNode {
	return s.GetToken(FeatureParserScenario, 0)
}

func (s *ScenarioContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *ScenarioContext) BlockBody() IBlockBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *ScenarioContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *ScenarioContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *ScenarioContext) EOF() antlr.TerminalNode {
	return s.GetToken(FeatureParserEOF, 0)
}

func (s *ScenarioContext) AllTags() []ITagsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagsContext)(nil)).Elem())
	var tst = make([]ITagsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagsContext)
		}
	}

	return tst
}

func (s *ScenarioContext) Tags(i int) ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *ScenarioContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *ScenarioContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
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
	p.EnterRule(localctx, 10, FeatureParserRULE_scenario)
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
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(135)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(141)
				p.Tags()
			}

		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(147)
			p.Match(FeatureParserSpace)
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(FeatureParserScenario)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(154)
			p.Match(FeatureParserSpace)
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(160)
		p.Content()
	}
	{
		p.SetState(161)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserEOF || _la == FeatureParserNewLine) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(162)
		p.BlockBody()
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

func (s *TagsContext) At() antlr.TerminalNode {
	return s.GetToken(FeatureParserAt, 0)
}

func (s *TagsContext) AnyText() IAnyTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnyTextContext)
}

func (s *TagsContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *TagsContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *TagsContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *TagsContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *TagsContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
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
	p.EnterRule(localctx, 12, FeatureParserRULE_tags)
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
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(164)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(170)
		p.Match(FeatureParserAt)
	}
	{
		p.SetState(171)
		p.AnyText()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FeatureParserLBracket {
		{
			p.SetState(172)
			p.Value()
		}

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(175)
			p.Match(FeatureParserSpace)
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(181)
		p.Match(FeatureParserNewLine)
	}

	return localctx
}

// IAnyTextContext is an interface to support dynamic dispatch.
type IAnyTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnyTextContext differentiates from other interfaces.
	IsAnyTextContext()
}

type AnyTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyTextContext() *AnyTextContext {
	var p = new(AnyTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_anyText
	return p
}

func (*AnyTextContext) IsAnyTextContext() {}

func NewAnyTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyTextContext {
	var p = new(AnyTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_anyText

	return p
}

func (s *AnyTextContext) GetParser() antlr.Parser { return s.parser }
func (s *AnyTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterAnyText(s)
	}
}

func (s *AnyTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitAnyText(s)
	}
}

func (s *AnyTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitAnyText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) AnyText() (localctx IAnyTextContext) {
	localctx = NewAnyTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FeatureParserRULE_anyText)

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
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(183)
			p.MatchWildcard()

		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) LBracket() antlr.TerminalNode {
	return s.GetToken(FeatureParserLBracket, 0)
}

func (s *ValueContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *ValueContext) RBracket() antlr.TerminalNode {
	return s.GetToken(FeatureParserRBracket, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FeatureParserRULE_value)

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
		p.SetState(189)
		p.Match(FeatureParserLBracket)
	}
	{
		p.SetState(190)
		p.Content()
	}
	{
		p.SetState(191)
		p.Match(FeatureParserRBracket)
	}

	return localctx
}

// IGivenContext is an interface to support dynamic dispatch.
type IGivenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGivenContext differentiates from other interfaces.
	IsGivenContext()
}

type GivenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGivenContext() *GivenContext {
	var p = new(GivenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_given
	return p
}

func (*GivenContext) IsGivenContext() {}

func NewGivenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GivenContext {
	var p = new(GivenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_given

	return p
}

func (s *GivenContext) GetParser() antlr.Parser { return s.parser }

func (s *GivenContext) Given() antlr.TerminalNode {
	return s.GetToken(FeatureParserGiven, 0)
}

func (s *GivenContext) Step() IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *GivenContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *GivenContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *GivenContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *GivenContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *GivenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GivenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GivenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterGiven(s)
	}
}

func (s *GivenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitGiven(s)
	}
}

func (s *GivenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitGiven(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Given() (localctx IGivenContext) {
	localctx = NewGivenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FeatureParserRULE_given)
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
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(193)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(199)
		p.Match(FeatureParserGiven)
	}
	{
		p.SetState(200)
		p.Step()
	}

	return localctx
}

// IWhenContext is an interface to support dynamic dispatch.
type IWhenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhenContext differentiates from other interfaces.
	IsWhenContext()
}

type WhenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhenContext() *WhenContext {
	var p = new(WhenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_when
	return p
}

func (*WhenContext) IsWhenContext() {}

func NewWhenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenContext {
	var p = new(WhenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_when

	return p
}

func (s *WhenContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenContext) When() antlr.TerminalNode {
	return s.GetToken(FeatureParserWhen, 0)
}

func (s *WhenContext) Step() IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *WhenContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *WhenContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *WhenContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *WhenContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *WhenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterWhen(s)
	}
}

func (s *WhenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitWhen(s)
	}
}

func (s *WhenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitWhen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) When() (localctx IWhenContext) {
	localctx = NewWhenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FeatureParserRULE_when)
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
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(202)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(208)
		p.Match(FeatureParserWhen)
	}
	{
		p.SetState(209)
		p.Step()
	}

	return localctx
}

// IOrContext is an interface to support dynamic dispatch.
type IOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrContext differentiates from other interfaces.
	IsOrContext()
}

type OrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrContext() *OrContext {
	var p = new(OrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_or
	return p
}

func (*OrContext) IsOrContext() {}

func NewOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrContext {
	var p = new(OrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_or

	return p
}

func (s *OrContext) GetParser() antlr.Parser { return s.parser }

func (s *OrContext) Or() antlr.TerminalNode {
	return s.GetToken(FeatureParserOr, 0)
}

func (s *OrContext) Step() IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *OrContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *OrContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *OrContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *OrContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Or() (localctx IOrContext) {
	localctx = NewOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FeatureParserRULE_or)
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
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(211)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(217)
		p.Match(FeatureParserOr)
	}
	{
		p.SetState(218)
		p.Step()
	}

	return localctx
}

// IAndContext is an interface to support dynamic dispatch.
type IAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndContext differentiates from other interfaces.
	IsAndContext()
}

type AndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndContext() *AndContext {
	var p = new(AndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_and
	return p
}

func (*AndContext) IsAndContext() {}

func NewAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndContext {
	var p = new(AndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_and

	return p
}

func (s *AndContext) GetParser() antlr.Parser { return s.parser }

func (s *AndContext) And() antlr.TerminalNode {
	return s.GetToken(FeatureParserAnd, 0)
}

func (s *AndContext) Step() IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *AndContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *AndContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *AndContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *AndContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) And() (localctx IAndContext) {
	localctx = NewAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FeatureParserRULE_and)
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
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(220)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(226)
		p.Match(FeatureParserAnd)
	}
	{
		p.SetState(227)
		p.Step()
	}

	return localctx
}

// IThenContext is an interface to support dynamic dispatch.
type IThenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenContext differentiates from other interfaces.
	IsThenContext()
}

type ThenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenContext() *ThenContext {
	var p = new(ThenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_then
	return p
}

func (*ThenContext) IsThenContext() {}

func NewThenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenContext {
	var p = new(ThenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_then

	return p
}

func (s *ThenContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenContext) Then() antlr.TerminalNode {
	return s.GetToken(FeatureParserThen, 0)
}

func (s *ThenContext) Step() IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *ThenContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *ThenContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *ThenContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *ThenContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *ThenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterThen(s)
	}
}

func (s *ThenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitThen(s)
	}
}

func (s *ThenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitThen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Then() (localctx IThenContext) {
	localctx = NewThenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FeatureParserRULE_then)
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

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(229)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(235)
		p.Match(FeatureParserThen)
	}
	{
		p.SetState(236)
		p.Step()
	}

	return localctx
}

// IExampleContext is an interface to support dynamic dispatch.
type IExampleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExampleContext differentiates from other interfaces.
	IsExampleContext()
}

type ExampleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExampleContext() *ExampleContext {
	var p = new(ExampleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_example
	return p
}

func (*ExampleContext) IsExampleContext() {}

func NewExampleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExampleContext {
	var p = new(ExampleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_example

	return p
}

func (s *ExampleContext) GetParser() antlr.Parser { return s.parser }

func (s *ExampleContext) Examples() antlr.TerminalNode {
	return s.GetToken(FeatureParserExamples, 0)
}

func (s *ExampleContext) Step() IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *ExampleContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *ExampleContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *ExampleContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *ExampleContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *ExampleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExampleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExampleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterExample(s)
	}
}

func (s *ExampleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitExample(s)
	}
}

func (s *ExampleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitExample(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Example() (localctx IExampleContext) {
	localctx = NewExampleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FeatureParserRULE_example)
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
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(238)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(244)
		p.Match(FeatureParserExamples)
	}
	{
		p.SetState(245)
		p.Step()
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

func (s *StepContext) StepContent() IStepContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepContentContext)
}

func (s *StepContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *StepContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *StepContext) Table() ITableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableContext)
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
	p.EnterRule(localctx, 30, FeatureParserRULE_step)

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
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(247)
				p.Match(FeatureParserSpace)
			}

		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	{
		p.SetState(253)
		p.StepContent()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(254)
			p.Table()
		}

	}

	return localctx
}

// IStepContentContext is an interface to support dynamic dispatch.
type IStepContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepContentContext differentiates from other interfaces.
	IsStepContentContext()
}

type StepContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepContentContext() *StepContentContext {
	var p = new(StepContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_stepContent
	return p
}

func (*StepContentContext) IsStepContentContext() {}

func NewStepContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepContentContext {
	var p = new(StepContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_stepContent

	return p
}

func (s *StepContentContext) GetParser() antlr.Parser { return s.parser }

func (s *StepContentContext) StepText() IStepTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepTextContext)
}

func (s *StepContentContext) EOF() antlr.TerminalNode {
	return s.GetToken(FeatureParserEOF, 0)
}

func (s *StepContentContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *StepContentContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *StepContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterStepContent(s)
	}
}

func (s *StepContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitStepContent(s)
	}
}

func (s *StepContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitStepContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) StepContent() (localctx IStepContentContext) {
	localctx = NewStepContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FeatureParserRULE_stepContent)

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
		p.SetState(257)
		p.StepText()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FeatureParserNewLine:
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(258)
					p.Match(FeatureParserNewLine)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(261)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
		}

	case FeatureParserEOF:
		{
			p.SetState(263)
			p.Match(FeatureParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStepTextContext is an interface to support dynamic dispatch.
type IStepTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepTextContext differentiates from other interfaces.
	IsStepTextContext()
}

type StepTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepTextContext() *StepTextContext {
	var p = new(StepTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_stepText
	return p
}

func (*StepTextContext) IsStepTextContext() {}

func NewStepTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepTextContext {
	var p = new(StepTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_stepText

	return p
}

func (s *StepTextContext) GetParser() antlr.Parser { return s.parser }

func (s *StepTextContext) AllContentNoQuotes() []IContentNoQuotesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IContentNoQuotesContext)(nil)).Elem())
	var tst = make([]IContentNoQuotesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IContentNoQuotesContext)
		}
	}

	return tst
}

func (s *StepTextContext) ContentNoQuotes(i int) IContentNoQuotesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentNoQuotesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IContentNoQuotesContext)
}

func (s *StepTextContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *StepTextContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *StepTextContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *StepTextContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *StepTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterStepText(s)
	}
}

func (s *StepTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitStepText(s)
	}
}

func (s *StepTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitStepText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) StepText() (localctx IStepTextContext) {
	localctx = NewStepTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FeatureParserRULE_stepText)
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
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserQuote)|(1<<FeatureParserLBracket)|(1<<FeatureParserChar))) != 0 {
		p.SetState(269)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FeatureParserLBracket, FeatureParserChar:
			{
				p.SetState(266)
				p.ContentNoQuotes()
			}

		case FeatureParserSpace:
			{
				p.SetState(267)
				p.Match(FeatureParserSpace)
			}

		case FeatureParserQuote:
			{
				p.SetState(268)
				p.Parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *TableContext) TableHeader() ITableHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableHeaderContext)
}

func (s *TableContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *TableContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
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
	p.EnterRule(localctx, 36, FeatureParserRULE_table)

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
		p.SetState(274)
		p.TableHeader()
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(275)
				p.Row()
			}

		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// ITableHeaderContext is an interface to support dynamic dispatch.
type ITableHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableHeaderContext differentiates from other interfaces.
	IsTableHeaderContext()
}

type TableHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableHeaderContext() *TableHeaderContext {
	var p = new(TableHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_tableHeader
	return p
}

func (*TableHeaderContext) IsTableHeaderContext() {}

func NewTableHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableHeaderContext {
	var p = new(TableHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_tableHeader

	return p
}

func (s *TableHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *TableHeaderContext) Pipe() antlr.TerminalNode {
	return s.GetToken(FeatureParserPipe, 0)
}

func (s *TableHeaderContext) EOF() antlr.TerminalNode {
	return s.GetToken(FeatureParserEOF, 0)
}

func (s *TableHeaderContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *TableHeaderContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *TableHeaderContext) AllCell() []ICellContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICellContext)(nil)).Elem())
	var tst = make([]ICellContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICellContext)
		}
	}

	return tst
}

func (s *TableHeaderContext) Cell(i int) ICellContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICellContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICellContext)
}

func (s *TableHeaderContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *TableHeaderContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *TableHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterTableHeader(s)
	}
}

func (s *TableHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitTableHeader(s)
	}
}

func (s *TableHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitTableHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) TableHeader() (localctx ITableHeaderContext) {
	localctx = NewTableHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FeatureParserRULE_tableHeader)
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
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(281)
			p.Match(FeatureParserSpace)
		}

		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(287)
		p.Match(FeatureParserPipe)
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserLBracket)|(1<<FeatureParserChar))) != 0) {
		{
			p.SetState(288)
			p.Cell()
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FeatureParserNewLine:
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(293)
					p.Match(FeatureParserNewLine)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
		}

	case FeatureParserEOF:
		{
			p.SetState(298)
			p.Match(FeatureParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) Pipe() antlr.TerminalNode {
	return s.GetToken(FeatureParserPipe, 0)
}

func (s *RowContext) EOF() antlr.TerminalNode {
	return s.GetToken(FeatureParserEOF, 0)
}

func (s *RowContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *RowContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *RowContext) AllCell() []ICellContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICellContext)(nil)).Elem())
	var tst = make([]ICellContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICellContext)
		}
	}

	return tst
}

func (s *RowContext) Cell(i int) ICellContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICellContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICellContext)
}

func (s *RowContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserNewLine)
}

func (s *RowContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, i)
}

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterRow(s)
	}
}

func (s *RowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitRow(s)
	}
}

func (s *RowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitRow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Row() (localctx IRowContext) {
	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FeatureParserRULE_row)
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
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(301)
			p.Match(FeatureParserSpace)
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
		p.Match(FeatureParserPipe)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserLBracket)|(1<<FeatureParserChar))) != 0) {
		{
			p.SetState(308)
			p.Cell()
		}

		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FeatureParserNewLine:
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(313)
					p.Match(FeatureParserNewLine)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
		}

	case FeatureParserEOF:
		{
			p.SetState(318)
			p.Match(FeatureParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *CellContext) ContentNoPipes() IContentNoPipesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentNoPipesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentNoPipesContext)
}

func (s *CellContext) Pipe() antlr.TerminalNode {
	return s.GetToken(FeatureParserPipe, 0)
}

func (s *CellContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *CellContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
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
	p.EnterRule(localctx, 42, FeatureParserRULE_cell)
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
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(321)
			p.Match(FeatureParserSpace)
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(327)
		p.ContentNoPipes()
	}
	{
		p.SetState(328)
		p.Match(FeatureParserPipe)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) AllQuote() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserQuote)
}

func (s *ParameterContext) Quote(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserQuote, i)
}

func (s *ParameterContext) AnyText() IAnyTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnyTextContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FeatureParserRULE_parameter)

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
		p.SetState(330)
		p.Match(FeatureParserQuote)
	}
	{
		p.SetState(331)
		p.AnyText()
	}
	{
		p.SetState(332)
		p.Match(FeatureParserQuote)
	}

	return localctx
}

// IContentNoQuotesContext is an interface to support dynamic dispatch.
type IContentNoQuotesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentNoQuotesContext differentiates from other interfaces.
	IsContentNoQuotesContext()
}

type ContentNoQuotesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentNoQuotesContext() *ContentNoQuotesContext {
	var p = new(ContentNoQuotesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_contentNoQuotes
	return p
}

func (*ContentNoQuotesContext) IsContentNoQuotesContext() {}

func NewContentNoQuotesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentNoQuotesContext {
	var p = new(ContentNoQuotesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_contentNoQuotes

	return p
}

func (s *ContentNoQuotesContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentNoQuotesContext) AllChar() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserChar)
}

func (s *ContentNoQuotesContext) Char(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserChar, i)
}

func (s *ContentNoQuotesContext) AllLBracket() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserLBracket)
}

func (s *ContentNoQuotesContext) LBracket(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserLBracket, i)
}

func (s *ContentNoQuotesContext) AllRBracket() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserRBracket)
}

func (s *ContentNoQuotesContext) RBracket(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserRBracket, i)
}

func (s *ContentNoQuotesContext) AllAt() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserAt)
}

func (s *ContentNoQuotesContext) At(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserAt, i)
}

func (s *ContentNoQuotesContext) AllPipe() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserPipe)
}

func (s *ContentNoQuotesContext) Pipe(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserPipe, i)
}

func (s *ContentNoQuotesContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *ContentNoQuotesContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *ContentNoQuotesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentNoQuotesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentNoQuotesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterContentNoQuotes(s)
	}
}

func (s *ContentNoQuotesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitContentNoQuotes(s)
	}
}

func (s *ContentNoQuotesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitContentNoQuotes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) ContentNoQuotes() (localctx IContentNoQuotesContext) {
	localctx = NewContentNoQuotesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FeatureParserRULE_contentNoQuotes)
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
		p.SetState(334)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserLBracket || _la == FeatureParserChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(335)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserLBracket)|(1<<FeatureParserRBracket)|(1<<FeatureParserAt)|(1<<FeatureParserPipe)|(1<<FeatureParserChar))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// IContentNoPipesContext is an interface to support dynamic dispatch.
type IContentNoPipesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentNoPipesContext differentiates from other interfaces.
	IsContentNoPipesContext()
}

type ContentNoPipesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentNoPipesContext() *ContentNoPipesContext {
	var p = new(ContentNoPipesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_contentNoPipes
	return p
}

func (*ContentNoPipesContext) IsContentNoPipesContext() {}

func NewContentNoPipesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentNoPipesContext {
	var p = new(ContentNoPipesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_contentNoPipes

	return p
}

func (s *ContentNoPipesContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentNoPipesContext) AllChar() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserChar)
}

func (s *ContentNoPipesContext) Char(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserChar, i)
}

func (s *ContentNoPipesContext) AllLBracket() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserLBracket)
}

func (s *ContentNoPipesContext) LBracket(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserLBracket, i)
}

func (s *ContentNoPipesContext) AllRBracket() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserRBracket)
}

func (s *ContentNoPipesContext) RBracket(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserRBracket, i)
}

func (s *ContentNoPipesContext) AllAt() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserAt)
}

func (s *ContentNoPipesContext) At(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserAt, i)
}

func (s *ContentNoPipesContext) AllQuote() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserQuote)
}

func (s *ContentNoPipesContext) Quote(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserQuote, i)
}

func (s *ContentNoPipesContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *ContentNoPipesContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *ContentNoPipesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentNoPipesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentNoPipesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterContentNoPipes(s)
	}
}

func (s *ContentNoPipesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitContentNoPipes(s)
	}
}

func (s *ContentNoPipesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitContentNoPipes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) ContentNoPipes() (localctx IContentNoPipesContext) {
	localctx = NewContentNoPipesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FeatureParserRULE_contentNoPipes)
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
		p.SetState(341)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserLBracket || _la == FeatureParserChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserQuote)|(1<<FeatureParserLBracket)|(1<<FeatureParserRBracket)|(1<<FeatureParserAt)|(1<<FeatureParserChar))) != 0 {
		{
			p.SetState(342)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserQuote)|(1<<FeatureParserLBracket)|(1<<FeatureParserRBracket)|(1<<FeatureParserAt)|(1<<FeatureParserChar))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) AllChar() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserChar)
}

func (s *ContentContext) Char(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserChar, i)
}

func (s *ContentContext) AllLBracket() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserLBracket)
}

func (s *ContentContext) LBracket(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserLBracket, i)
}

func (s *ContentContext) AllRBracket() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserRBracket)
}

func (s *ContentContext) RBracket(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserRBracket, i)
}

func (s *ContentContext) AllAt() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserAt)
}

func (s *ContentContext) At(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserAt, i)
}

func (s *ContentContext) AllQuote() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserQuote)
}

func (s *ContentContext) Quote(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserQuote, i)
}

func (s *ContentContext) AllPipe() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserPipe)
}

func (s *ContentContext) Pipe(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserPipe, i)
}

func (s *ContentContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *ContentContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitContent(s)
	}
}

func (s *ContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FeatureParserRULE_content)
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
		p.SetState(348)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserLBracket || _la == FeatureParserChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(349)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserQuote)|(1<<FeatureParserLBracket)|(1<<FeatureParserRBracket)|(1<<FeatureParserAt)|(1<<FeatureParserPipe)|(1<<FeatureParserChar))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
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

func (s *CommentContext) CommentText() ICommentTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentTextContext)
}

func (s *CommentContext) NewLine() antlr.TerminalNode {
	return s.GetToken(FeatureParserNewLine, 0)
}

func (s *CommentContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(FeatureParserSpace)
}

func (s *CommentContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(FeatureParserSpace, i)
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
	p.EnterRule(localctx, 52, FeatureParserRULE_comment)
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
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(355)
			p.Match(FeatureParserSpace)
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(361)
		p.Match(FeatureParserT__0)
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(362)
				p.Match(FeatureParserSpace)
			}

		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
	}
	{
		p.SetState(368)
		p.CommentText()
	}
	{
		p.SetState(369)
		p.Match(FeatureParserNewLine)
	}

	return localctx
}

// ICommentTextContext is an interface to support dynamic dispatch.
type ICommentTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentTextContext differentiates from other interfaces.
	IsCommentTextContext()
}

type CommentTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentTextContext() *CommentTextContext {
	var p = new(CommentTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FeatureParserRULE_commentText
	return p
}

func (*CommentTextContext) IsCommentTextContext() {}

func NewCommentTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentTextContext {
	var p = new(CommentTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FeatureParserRULE_commentText

	return p
}

func (s *CommentTextContext) GetParser() antlr.Parser { return s.parser }
func (s *CommentTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.EnterCommentText(s)
	}
}

func (s *CommentTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FeatureListener); ok {
		listenerT.ExitCommentText(s)
	}
}

func (s *CommentTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FeatureVisitor:
		return t.VisitCommentText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FeatureParser) CommentText() (localctx ICommentTextContext) {
	localctx = NewCommentTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FeatureParserRULE_commentText)

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
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(371)
			p.MatchWildcard()

		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}

	return localctx
}
