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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 308,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 53, 10, 3, 12, 3, 14, 3, 56,
	11, 3, 3, 3, 7, 3, 59, 10, 3, 12, 3, 14, 3, 62, 11, 3, 3, 3, 3, 3, 7, 3,
	66, 10, 3, 12, 3, 14, 3, 69, 11, 3, 3, 3, 3, 3, 6, 3, 73, 10, 3, 13, 3,
	14, 3, 74, 3, 4, 5, 4, 78, 10, 4, 3, 4, 6, 4, 81, 10, 4, 13, 4, 14, 4,
	82, 3, 5, 7, 5, 86, 10, 5, 12, 5, 14, 5, 89, 11, 5, 3, 5, 7, 5, 92, 10,
	5, 12, 5, 14, 5, 95, 11, 5, 3, 5, 3, 5, 7, 5, 99, 10, 5, 12, 5, 14, 5,
	102, 11, 5, 3, 5, 5, 5, 105, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 116, 10, 6, 12, 6, 14, 6, 119, 11, 6, 3, 7, 7,
	7, 122, 10, 7, 12, 7, 14, 7, 125, 11, 7, 3, 7, 7, 7, 128, 10, 7, 12, 7,
	14, 7, 131, 11, 7, 3, 7, 7, 7, 134, 10, 7, 12, 7, 14, 7, 137, 11, 7, 3,
	7, 3, 7, 7, 7, 141, 10, 7, 12, 7, 14, 7, 144, 11, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 7, 8, 151, 10, 8, 12, 8, 14, 8, 154, 11, 8, 3, 8, 3, 8, 3,
	8, 5, 8, 159, 10, 8, 3, 8, 7, 8, 162, 10, 8, 12, 8, 14, 8, 165, 11, 8,
	3, 8, 3, 8, 3, 9, 7, 9, 170, 10, 9, 12, 9, 14, 9, 173, 11, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 7, 11, 180, 10, 11, 12, 11, 14, 11, 183, 11, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 7, 12, 189, 10, 12, 12, 12, 14, 12, 192, 11,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 7, 13, 198, 10, 13, 12, 13, 14, 13, 201,
	11, 13, 3, 13, 3, 13, 3, 13, 3, 14, 7, 14, 207, 10, 14, 12, 14, 14, 14,
	210, 11, 14, 3, 14, 3, 14, 3, 14, 3, 15, 7, 15, 216, 10, 15, 12, 15, 14,
	15, 219, 11, 15, 3, 15, 3, 15, 3, 15, 3, 16, 7, 16, 225, 10, 16, 12, 16,
	14, 16, 228, 11, 16, 3, 16, 3, 16, 7, 16, 232, 10, 16, 12, 16, 14, 16,
	235, 11, 16, 3, 17, 3, 17, 6, 17, 239, 10, 17, 13, 17, 14, 17, 240, 3,
	17, 5, 17, 244, 10, 17, 3, 18, 3, 18, 3, 18, 7, 18, 249, 10, 18, 12, 18,
	14, 18, 252, 11, 18, 3, 19, 7, 19, 255, 10, 19, 12, 19, 14, 19, 258, 11,
	19, 3, 19, 3, 19, 6, 19, 262, 10, 19, 13, 19, 14, 19, 263, 3, 19, 6, 19,
	267, 10, 19, 13, 19, 14, 19, 268, 3, 19, 5, 19, 272, 10, 19, 3, 20, 7,
	20, 275, 10, 20, 12, 20, 14, 20, 278, 11, 20, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22, 289, 10, 22, 12, 22, 14, 22,
	292, 11, 22, 3, 23, 3, 23, 7, 23, 296, 10, 23, 12, 23, 14, 23, 299, 11,
	23, 3, 24, 3, 24, 7, 24, 303, 10, 24, 12, 24, 14, 24, 306, 11, 24, 3, 24,
	3, 171, 2, 25, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 2, 8, 3, 2, 13, 14, 3, 3, 14, 14, 4, 2,
	16, 16, 20, 20, 4, 2, 13, 13, 16, 20, 5, 2, 13, 13, 15, 18, 20, 20, 4,
	2, 13, 13, 15, 20, 2, 327, 2, 48, 3, 2, 2, 2, 4, 54, 3, 2, 2, 2, 6, 77,
	3, 2, 2, 2, 8, 87, 3, 2, 2, 2, 10, 117, 3, 2, 2, 2, 12, 123, 3, 2, 2, 2,
	14, 152, 3, 2, 2, 2, 16, 171, 3, 2, 2, 2, 18, 174, 3, 2, 2, 2, 20, 181,
	3, 2, 2, 2, 22, 190, 3, 2, 2, 2, 24, 199, 3, 2, 2, 2, 26, 208, 3, 2, 2,
	2, 28, 217, 3, 2, 2, 2, 30, 226, 3, 2, 2, 2, 32, 236, 3, 2, 2, 2, 34, 250,
	3, 2, 2, 2, 36, 256, 3, 2, 2, 2, 38, 276, 3, 2, 2, 2, 40, 282, 3, 2, 2,
	2, 42, 286, 3, 2, 2, 2, 44, 293, 3, 2, 2, 2, 46, 300, 3, 2, 2, 2, 48, 49,
	5, 4, 3, 2, 49, 50, 5, 6, 4, 2, 50, 3, 3, 2, 2, 2, 51, 53, 9, 2, 2, 2,
	52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3,
	2, 2, 2, 55, 60, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 59, 5, 14, 8, 2, 58,
	57, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2,
	2, 61, 63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 67, 7, 12, 2, 2, 64, 66,
	7, 13, 2, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2,
	67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 72, 5,
	46, 24, 2, 71, 73, 7, 14, 2, 2, 72, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 5, 3, 2, 2, 2, 76, 78, 5, 8,
	5, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 81,
	5, 12, 7, 2, 80, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2,
	82, 83, 3, 2, 2, 2, 83, 7, 3, 2, 2, 2, 84, 86, 9, 2, 2, 2, 85, 84, 3, 2,
	2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 93,
	3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 92, 5, 14, 8, 2, 91, 90, 3, 2, 2, 2,
	92, 95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 96, 3,
	2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 100, 7, 10, 2, 2, 97, 99, 7, 13, 2, 2,
	98, 97, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101,
	3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 5, 46,
	24, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2,
	106, 107, 9, 3, 2, 2, 107, 108, 5, 10, 6, 2, 108, 109, 9, 3, 2, 2, 109,
	9, 3, 2, 2, 2, 110, 116, 5, 20, 11, 2, 111, 116, 5, 22, 12, 2, 112, 116,
	5, 24, 13, 2, 113, 116, 5, 28, 15, 2, 114, 116, 5, 26, 14, 2, 115, 110,
	3, 2, 2, 2, 115, 111, 3, 2, 2, 2, 115, 112, 3, 2, 2, 2, 115, 113, 3, 2,
	2, 2, 115, 114, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2,
	117, 118, 3, 2, 2, 2, 118, 11, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 122,
	9, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121, 3, 2,
	2, 2, 123, 124, 3, 2, 2, 2, 124, 129, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2,
	126, 128, 5, 14, 8, 2, 127, 126, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129,
	127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 135, 3, 2, 2, 2, 131, 129,
	3, 2, 2, 2, 132, 134, 7, 13, 2, 2, 133, 132, 3, 2, 2, 2, 134, 137, 3, 2,
	2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 138, 3, 2, 2, 2,
	137, 135, 3, 2, 2, 2, 138, 142, 7, 11, 2, 2, 139, 141, 7, 13, 2, 2, 140,
	139, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143,
	3, 2, 2, 2, 143, 145, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 5, 46,
	24, 2, 146, 147, 9, 3, 2, 2, 147, 148, 5, 10, 6, 2, 148, 13, 3, 2, 2, 2,
	149, 151, 9, 2, 2, 2, 150, 149, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152,
	150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 155, 3, 2, 2, 2, 154, 152,
	3, 2, 2, 2, 155, 156, 7, 18, 2, 2, 156, 158, 5, 16, 9, 2, 157, 159, 5,
	18, 10, 2, 158, 157, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 163, 3, 2,
	2, 2, 160, 162, 7, 13, 2, 2, 161, 160, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2,
	163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165,
	163, 3, 2, 2, 2, 166, 167, 7, 14, 2, 2, 167, 15, 3, 2, 2, 2, 168, 170,
	11, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171, 172, 3, 2,
	2, 2, 171, 169, 3, 2, 2, 2, 172, 17, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2,
	174, 175, 7, 16, 2, 2, 175, 176, 5, 46, 24, 2, 176, 177, 7, 17, 2, 2, 177,
	19, 3, 2, 2, 2, 178, 180, 9, 2, 2, 2, 179, 178, 3, 2, 2, 2, 180, 183, 3,
	2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 184, 3, 2, 2,
	2, 183, 181, 3, 2, 2, 2, 184, 185, 7, 7, 2, 2, 185, 186, 5, 30, 16, 2,
	186, 21, 3, 2, 2, 2, 187, 189, 9, 2, 2, 2, 188, 187, 3, 2, 2, 2, 189, 192,
	3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 193, 3, 2,
	2, 2, 192, 190, 3, 2, 2, 2, 193, 194, 7, 8, 2, 2, 194, 195, 5, 30, 16,
	2, 195, 23, 3, 2, 2, 2, 196, 198, 9, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198,
	201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 202,
	3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202, 203, 7, 6, 2, 2, 203, 204, 5, 30,
	16, 2, 204, 25, 3, 2, 2, 2, 205, 207, 9, 2, 2, 2, 206, 205, 3, 2, 2, 2,
	207, 210, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209,
	211, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211, 212, 7, 5, 2, 2, 212, 213,
	5, 30, 16, 2, 213, 27, 3, 2, 2, 2, 214, 216, 9, 2, 2, 2, 215, 214, 3, 2,
	2, 2, 216, 219, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2,
	218, 220, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 220, 221, 7, 9, 2, 2, 221,
	222, 5, 30, 16, 2, 222, 29, 3, 2, 2, 2, 223, 225, 7, 13, 2, 2, 224, 223,
	3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2,
	2, 2, 227, 229, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229, 233, 5, 32, 17,
	2, 230, 232, 5, 36, 19, 2, 231, 230, 3, 2, 2, 2, 232, 235, 3, 2, 2, 2,
	233, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 31, 3, 2, 2, 2, 235, 233,
	3, 2, 2, 2, 236, 243, 5, 34, 18, 2, 237, 239, 7, 14, 2, 2, 238, 237, 3,
	2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2,
	2, 241, 244, 3, 2, 2, 2, 242, 244, 7, 2, 2, 3, 243, 238, 3, 2, 2, 2, 243,
	242, 3, 2, 2, 2, 244, 33, 3, 2, 2, 2, 245, 249, 5, 42, 22, 2, 246, 249,
	7, 13, 2, 2, 247, 249, 5, 40, 21, 2, 248, 245, 3, 2, 2, 2, 248, 246, 3,
	2, 2, 2, 248, 247, 3, 2, 2, 2, 249, 252, 3, 2, 2, 2, 250, 248, 3, 2, 2,
	2, 250, 251, 3, 2, 2, 2, 251, 35, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 253,
	255, 7, 13, 2, 2, 254, 253, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 254,
	3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 259, 3, 2, 2, 2, 258, 256, 3, 2,
	2, 2, 259, 261, 7, 19, 2, 2, 260, 262, 5, 38, 20, 2, 261, 260, 3, 2, 2,
	2, 262, 263, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264,
	271, 3, 2, 2, 2, 265, 267, 7, 14, 2, 2, 266, 265, 3, 2, 2, 2, 267, 268,
	3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 272, 3, 2,
	2, 2, 270, 272, 7, 2, 2, 3, 271, 266, 3, 2, 2, 2, 271, 270, 3, 2, 2, 2,
	272, 37, 3, 2, 2, 2, 273, 275, 7, 13, 2, 2, 274, 273, 3, 2, 2, 2, 275,
	278, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 279,
	3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 279, 280, 5, 44, 23, 2, 280, 281, 7,
	19, 2, 2, 281, 39, 3, 2, 2, 2, 282, 283, 7, 15, 2, 2, 283, 284, 5, 16,
	9, 2, 284, 285, 7, 15, 2, 2, 285, 41, 3, 2, 2, 2, 286, 290, 9, 4, 2, 2,
	287, 289, 9, 5, 2, 2, 288, 287, 3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290,
	288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 43, 3, 2, 2, 2, 292, 290, 3,
	2, 2, 2, 293, 297, 9, 4, 2, 2, 294, 296, 9, 6, 2, 2, 295, 294, 3, 2, 2,
	2, 296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298,
	45, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 300, 304, 9, 4, 2, 2, 301, 303, 9,
	7, 2, 2, 302, 301, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2, 304, 302, 3, 2, 2,
	2, 304, 305, 3, 2, 2, 2, 305, 47, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 41,
	54, 60, 67, 74, 77, 82, 87, 93, 100, 104, 115, 117, 123, 129, 135, 142,
	152, 158, 163, 171, 181, 190, 199, 208, 217, 226, 233, 240, 243, 248, 250,
	256, 263, 268, 271, 276, 290, 297, 304,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "'\"'", "'('", "')'",
	"'@'", "'|'",
}
var symbolicNames = []string{
	"", "Comment", "EmptyLine", "And", "Or", "Given", "When", "Then", "Background",
	"Scenario", "Feature", "Space", "NewLine", "Quote", "LBracket", "RBracket",
	"At", "Pipe", "Char",
}

var ruleNames = []string{
	"feature", "featureHeader", "featureBody", "background", "blockBody", "scenario",
	"tags", "anyText", "value", "given", "when", "or", "and", "then", "step",
	"stepContent", "stepText", "row", "cell", "parameter", "contentNoQuotes",
	"contentNoPipes", "content",
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
	FeatureParserComment    = 1
	FeatureParserEmptyLine  = 2
	FeatureParserAnd        = 3
	FeatureParserOr         = 4
	FeatureParserGiven      = 5
	FeatureParserWhen       = 6
	FeatureParserThen       = 7
	FeatureParserBackground = 8
	FeatureParserScenario   = 9
	FeatureParserFeature    = 10
	FeatureParserSpace      = 11
	FeatureParserNewLine    = 12
	FeatureParserQuote      = 13
	FeatureParserLBracket   = 14
	FeatureParserRBracket   = 15
	FeatureParserAt         = 16
	FeatureParserPipe       = 17
	FeatureParserChar       = 18
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
	FeatureParserRULE_step            = 14
	FeatureParserRULE_stepContent     = 15
	FeatureParserRULE_stepText        = 16
	FeatureParserRULE_row             = 17
	FeatureParserRULE_cell            = 18
	FeatureParserRULE_parameter       = 19
	FeatureParserRULE_contentNoQuotes = 20
	FeatureParserRULE_contentNoPipes  = 21
	FeatureParserRULE_content         = 22
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.FeatureHeader()
	}
	{
		p.SetState(47)
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
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(49)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserNewLine)|(1<<FeatureParserAt))) != 0 {
		{
			p.SetState(55)
			p.Tags()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
		p.Match(FeatureParserFeature)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(62)
			p.Match(FeatureParserSpace)
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Content()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(69)
				p.Match(FeatureParserNewLine)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(74)
			p.Background()
		}

	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserScenario)|(1<<FeatureParserSpace)|(1<<FeatureParserNewLine)|(1<<FeatureParserAt))) != 0) {
		{
			p.SetState(77)
			p.Scenario()
		}

		p.SetState(80)
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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(82)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserNewLine)|(1<<FeatureParserAt))) != 0 {
		{
			p.SetState(88)
			p.Tags()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(FeatureParserBackground)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(95)
			p.Match(FeatureParserSpace)
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FeatureParserLBracket || _la == FeatureParserChar {
		{
			p.SetState(101)
			p.Content()
		}

	}
	{
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserEOF || _la == FeatureParserNewLine) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(105)
		p.BlockBody()
	}
	{
		p.SetState(106)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(108)
					p.Given()
				}

			case 2:
				{
					p.SetState(109)
					p.When()
				}

			case 3:
				{
					p.SetState(110)
					p.Or()
				}

			case 4:
				{
					p.SetState(111)
					p.Then()
				}

			case 5:
				{
					p.SetState(112)
					p.And()
				}

			}

		}
		p.SetState(117)
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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(118)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(124)
				p.Tags()
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(130)
			p.Match(FeatureParserSpace)
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(FeatureParserScenario)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(137)
			p.Match(FeatureParserSpace)
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Content()
	}
	{
		p.SetState(144)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserEOF || _la == FeatureParserNewLine) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(145)
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(147)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(FeatureParserAt)
	}
	{
		p.SetState(154)
		p.AnyText()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FeatureParserLBracket {
		{
			p.SetState(155)
			p.Value()
		}

	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(158)
			p.Match(FeatureParserSpace)
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
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
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(166)
			p.MatchWildcard()

		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
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
		p.SetState(172)
		p.Match(FeatureParserLBracket)
	}
	{
		p.SetState(173)
		p.Content()
	}
	{
		p.SetState(174)
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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(176)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(182)
		p.Match(FeatureParserGiven)
	}
	{
		p.SetState(183)
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
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(185)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(191)
		p.Match(FeatureParserWhen)
	}
	{
		p.SetState(192)
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
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(194)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(200)
		p.Match(FeatureParserOr)
	}
	{
		p.SetState(201)
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
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(203)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(209)
		p.Match(FeatureParserAnd)
	}
	{
		p.SetState(210)
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
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace || _la == FeatureParserNewLine {
		{
			p.SetState(212)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FeatureParserSpace || _la == FeatureParserNewLine) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(218)
		p.Match(FeatureParserThen)
	}
	{
		p.SetState(219)
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

func (s *StepContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *StepContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
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
	p.EnterRule(localctx, 28, FeatureParserRULE_step)

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
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(221)
				p.Match(FeatureParserSpace)
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	{
		p.SetState(227)
		p.StepContent()
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(228)
				p.Row()
			}

		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 30, FeatureParserRULE_stepContent)

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
		p.SetState(234)
		p.StepText()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FeatureParserNewLine:
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(235)
					p.Match(FeatureParserNewLine)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
		}

	case FeatureParserEOF:
		{
			p.SetState(240)
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
	p.EnterRule(localctx, 32, FeatureParserRULE_stepText)
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
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserQuote)|(1<<FeatureParserLBracket)|(1<<FeatureParserChar))) != 0 {
		p.SetState(246)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FeatureParserLBracket, FeatureParserChar:
			{
				p.SetState(243)
				p.ContentNoQuotes()
			}

		case FeatureParserSpace:
			{
				p.SetState(244)
				p.Match(FeatureParserSpace)
			}

		case FeatureParserQuote:
			{
				p.SetState(245)
				p.Parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 34, FeatureParserRULE_row)
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
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(251)
			p.Match(FeatureParserSpace)
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(257)
		p.Match(FeatureParserPipe)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserLBracket)|(1<<FeatureParserChar))) != 0) {
		{
			p.SetState(258)
			p.Cell()
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FeatureParserNewLine:
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(263)
					p.Match(FeatureParserNewLine)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(266)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
		}

	case FeatureParserEOF:
		{
			p.SetState(268)
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
	p.EnterRule(localctx, 36, FeatureParserRULE_cell)
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
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FeatureParserSpace {
		{
			p.SetState(271)
			p.Match(FeatureParserSpace)
		}

		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(277)
		p.ContentNoPipes()
	}
	{
		p.SetState(278)
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
	p.EnterRule(localctx, 38, FeatureParserRULE_parameter)

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
		p.Match(FeatureParserQuote)
	}
	{
		p.SetState(281)
		p.AnyText()
	}
	{
		p.SetState(282)
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
	p.EnterRule(localctx, 40, FeatureParserRULE_contentNoQuotes)
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
		p.SetState(284)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserLBracket || _la == FeatureParserChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(285)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserLBracket)|(1<<FeatureParserRBracket)|(1<<FeatureParserAt)|(1<<FeatureParserPipe)|(1<<FeatureParserChar))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 42, FeatureParserRULE_contentNoPipes)
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
		p.SetState(291)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserLBracket || _la == FeatureParserChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserQuote)|(1<<FeatureParserLBracket)|(1<<FeatureParserRBracket)|(1<<FeatureParserAt)|(1<<FeatureParserChar))) != 0 {
		{
			p.SetState(292)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserQuote)|(1<<FeatureParserLBracket)|(1<<FeatureParserRBracket)|(1<<FeatureParserAt)|(1<<FeatureParserChar))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(297)
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
	p.EnterRule(localctx, 44, FeatureParserRULE_content)
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
		p.SetState(298)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FeatureParserLBracket || _la == FeatureParserChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(299)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FeatureParserSpace)|(1<<FeatureParserQuote)|(1<<FeatureParserLBracket)|(1<<FeatureParserRBracket)|(1<<FeatureParserAt)|(1<<FeatureParserPipe)|(1<<FeatureParserChar))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}

	return localctx
}
