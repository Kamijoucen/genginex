// Code generated from ../gengine.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 58, 518,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34,
	3, 34, 5, 34, 251, 10, 34, 3, 34, 6, 34, 254, 10, 34, 13, 34, 14, 34, 255,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40,
	3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	52, 3, 52, 3, 52, 3, 52, 3, 53, 6, 53, 357, 10, 53, 13, 53, 14, 53, 358,
	3, 53, 7, 53, 362, 10, 53, 12, 53, 14, 53, 365, 11, 53, 3, 54, 6, 54, 368,
	10, 54, 13, 54, 14, 54, 369, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57,
	3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3,
	62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66,
	3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3,
	70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74,
	3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77, 3, 78, 3, 78, 3, 79, 3,
	79, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 7, 80, 437, 10, 80, 12, 80,
	14, 80, 440, 11, 80, 3, 80, 3, 80, 3, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3,
	82, 3, 82, 3, 82, 3, 82, 3, 82, 3, 83, 6, 83, 455, 10, 83, 13, 83, 14,
	83, 456, 5, 83, 459, 10, 83, 3, 83, 3, 83, 6, 83, 463, 10, 83, 13, 83,
	14, 83, 464, 3, 83, 6, 83, 468, 10, 83, 13, 83, 14, 83, 469, 3, 83, 3,
	83, 3, 83, 3, 83, 6, 83, 476, 10, 83, 13, 83, 14, 83, 477, 5, 83, 480,
	10, 83, 3, 83, 3, 83, 6, 83, 484, 10, 83, 13, 83, 14, 83, 485, 3, 83, 3,
	83, 3, 83, 6, 83, 491, 10, 83, 13, 83, 14, 83, 492, 3, 83, 3, 83, 5, 83,
	497, 10, 83, 3, 84, 3, 84, 3, 84, 3, 84, 7, 84, 503, 10, 84, 12, 84, 14,
	84, 506, 11, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 85, 6, 85, 513, 10, 85,
	13, 85, 14, 85, 514, 3, 85, 3, 85, 3, 504, 2, 86, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2,
	31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 2, 51,
	2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2, 69, 8, 71, 9,
	73, 10, 75, 11, 77, 12, 79, 13, 81, 14, 83, 15, 85, 16, 87, 17, 89, 18,
	91, 19, 93, 20, 95, 21, 97, 22, 99, 23, 101, 24, 103, 25, 105, 26, 107,
	27, 109, 28, 111, 29, 113, 30, 115, 31, 117, 32, 119, 33, 121, 34, 123,
	35, 125, 36, 127, 37, 129, 38, 131, 39, 133, 40, 135, 41, 137, 42, 139,
	43, 141, 44, 143, 45, 145, 46, 147, 47, 149, 48, 151, 49, 153, 50, 155,
	51, 157, 52, 159, 53, 161, 54, 163, 55, 165, 56, 167, 57, 169, 58, 3, 2,
	33, 3, 2, 50, 59, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69,
	69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72,
	72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75,
	75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78,
	78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81,
	81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84,
	84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87,
	87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90,
	90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 5, 2, 67,
	92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 36, 36,
	94, 94, 5, 2, 11, 12, 15, 15, 34, 34, 2, 510, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 69,
	3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2,
	77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2,
	2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2,
	2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2,
	2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107,
	3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2,
	2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3,
	2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2,
	129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2,
	2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143,
	3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2,
	2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3,
	2, 2, 2, 2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2,
	165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 3, 171, 3, 2,
	2, 2, 5, 173, 3, 2, 2, 2, 7, 179, 3, 2, 2, 2, 9, 183, 3, 2, 2, 2, 11, 189,
	3, 2, 2, 2, 13, 194, 3, 2, 2, 2, 15, 196, 3, 2, 2, 2, 17, 198, 3, 2, 2,
	2, 19, 200, 3, 2, 2, 2, 21, 202, 3, 2, 2, 2, 23, 204, 3, 2, 2, 2, 25, 206,
	3, 2, 2, 2, 27, 208, 3, 2, 2, 2, 29, 210, 3, 2, 2, 2, 31, 212, 3, 2, 2,
	2, 33, 214, 3, 2, 2, 2, 35, 216, 3, 2, 2, 2, 37, 218, 3, 2, 2, 2, 39, 220,
	3, 2, 2, 2, 41, 222, 3, 2, 2, 2, 43, 224, 3, 2, 2, 2, 45, 226, 3, 2, 2,
	2, 47, 228, 3, 2, 2, 2, 49, 230, 3, 2, 2, 2, 51, 232, 3, 2, 2, 2, 53, 234,
	3, 2, 2, 2, 55, 236, 3, 2, 2, 2, 57, 238, 3, 2, 2, 2, 59, 240, 3, 2, 2,
	2, 61, 242, 3, 2, 2, 2, 63, 244, 3, 2, 2, 2, 65, 246, 3, 2, 2, 2, 67, 248,
	3, 2, 2, 2, 69, 257, 3, 2, 2, 2, 71, 261, 3, 2, 2, 2, 73, 266, 3, 2, 2,
	2, 75, 269, 3, 2, 2, 2, 77, 272, 3, 2, 2, 2, 79, 277, 3, 2, 2, 2, 81, 280,
	3, 2, 2, 2, 83, 285, 3, 2, 2, 2, 85, 292, 3, 2, 2, 2, 87, 296, 3, 2, 2,
	2, 89, 302, 3, 2, 2, 2, 91, 311, 3, 2, 2, 2, 93, 320, 3, 2, 2, 2, 95, 325,
	3, 2, 2, 2, 97, 331, 3, 2, 2, 2, 99, 336, 3, 2, 2, 2, 101, 345, 3, 2, 2,
	2, 103, 351, 3, 2, 2, 2, 105, 356, 3, 2, 2, 2, 107, 367, 3, 2, 2, 2, 109,
	371, 3, 2, 2, 2, 111, 373, 3, 2, 2, 2, 113, 375, 3, 2, 2, 2, 115, 377,
	3, 2, 2, 2, 117, 379, 3, 2, 2, 2, 119, 382, 3, 2, 2, 2, 121, 384, 3, 2,
	2, 2, 123, 386, 3, 2, 2, 2, 125, 389, 3, 2, 2, 2, 127, 392, 3, 2, 2, 2,
	129, 395, 3, 2, 2, 2, 131, 397, 3, 2, 2, 2, 133, 400, 3, 2, 2, 2, 135,
	402, 3, 2, 2, 2, 137, 405, 3, 2, 2, 2, 139, 408, 3, 2, 2, 2, 141, 411,
	3, 2, 2, 2, 143, 414, 3, 2, 2, 2, 145, 416, 3, 2, 2, 2, 147, 418, 3, 2,
	2, 2, 149, 420, 3, 2, 2, 2, 151, 422, 3, 2, 2, 2, 153, 424, 3, 2, 2, 2,
	155, 426, 3, 2, 2, 2, 157, 428, 3, 2, 2, 2, 159, 430, 3, 2, 2, 2, 161,
	443, 3, 2, 2, 2, 163, 447, 3, 2, 2, 2, 165, 496, 3, 2, 2, 2, 167, 498,
	3, 2, 2, 2, 169, 512, 3, 2, 2, 2, 171, 172, 7, 46, 2, 2, 172, 4, 3, 2,
	2, 2, 173, 174, 7, 66, 2, 2, 174, 175, 7, 112, 2, 2, 175, 176, 7, 99, 2,
	2, 176, 177, 7, 111, 2, 2, 177, 178, 7, 103, 2, 2, 178, 6, 3, 2, 2, 2,
	179, 180, 7, 66, 2, 2, 180, 181, 7, 107, 2, 2, 181, 182, 7, 102, 2, 2,
	182, 8, 3, 2, 2, 2, 183, 184, 7, 66, 2, 2, 184, 185, 7, 102, 2, 2, 185,
	186, 7, 103, 2, 2, 186, 187, 7, 117, 2, 2, 187, 188, 7, 101, 2, 2, 188,
	10, 3, 2, 2, 2, 189, 190, 7, 66, 2, 2, 190, 191, 7, 117, 2, 2, 191, 192,
	7, 99, 2, 2, 192, 193, 7, 110, 2, 2, 193, 12, 3, 2, 2, 2, 194, 195, 9,
	2, 2, 2, 195, 14, 3, 2, 2, 2, 196, 197, 9, 3, 2, 2, 197, 16, 3, 2, 2, 2,
	198, 199, 9, 4, 2, 2, 199, 18, 3, 2, 2, 2, 200, 201, 9, 5, 2, 2, 201, 20,
	3, 2, 2, 2, 202, 203, 9, 6, 2, 2, 203, 22, 3, 2, 2, 2, 204, 205, 9, 7,
	2, 2, 205, 24, 3, 2, 2, 2, 206, 207, 9, 8, 2, 2, 207, 26, 3, 2, 2, 2, 208,
	209, 9, 9, 2, 2, 209, 28, 3, 2, 2, 2, 210, 211, 9, 10, 2, 2, 211, 30, 3,
	2, 2, 2, 212, 213, 9, 11, 2, 2, 213, 32, 3, 2, 2, 2, 214, 215, 9, 12, 2,
	2, 215, 34, 3, 2, 2, 2, 216, 217, 9, 13, 2, 2, 217, 36, 3, 2, 2, 2, 218,
	219, 9, 14, 2, 2, 219, 38, 3, 2, 2, 2, 220, 221, 9, 15, 2, 2, 221, 40,
	3, 2, 2, 2, 222, 223, 9, 16, 2, 2, 223, 42, 3, 2, 2, 2, 224, 225, 9, 17,
	2, 2, 225, 44, 3, 2, 2, 2, 226, 227, 9, 18, 2, 2, 227, 46, 3, 2, 2, 2,
	228, 229, 9, 19, 2, 2, 229, 48, 3, 2, 2, 2, 230, 231, 9, 20, 2, 2, 231,
	50, 3, 2, 2, 2, 232, 233, 9, 21, 2, 2, 233, 52, 3, 2, 2, 2, 234, 235, 9,
	22, 2, 2, 235, 54, 3, 2, 2, 2, 236, 237, 9, 23, 2, 2, 237, 56, 3, 2, 2,
	2, 238, 239, 9, 24, 2, 2, 239, 58, 3, 2, 2, 2, 240, 241, 9, 25, 2, 2, 241,
	60, 3, 2, 2, 2, 242, 243, 9, 26, 2, 2, 243, 62, 3, 2, 2, 2, 244, 245, 9,
	27, 2, 2, 245, 64, 3, 2, 2, 2, 246, 247, 9, 28, 2, 2, 247, 66, 3, 2, 2,
	2, 248, 250, 9, 7, 2, 2, 249, 251, 7, 47, 2, 2, 250, 249, 3, 2, 2, 2, 250,
	251, 3, 2, 2, 2, 251, 253, 3, 2, 2, 2, 252, 254, 5, 13, 7, 2, 253, 252,
	3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2,
	2, 2, 256, 68, 3, 2, 2, 2, 257, 258, 5, 41, 21, 2, 258, 259, 5, 31, 16,
	2, 259, 260, 5, 37, 19, 2, 260, 70, 3, 2, 2, 2, 261, 262, 5, 49, 25, 2,
	262, 263, 5, 55, 28, 2, 263, 264, 5, 37, 19, 2, 264, 265, 5, 23, 12, 2,
	265, 72, 3, 2, 2, 2, 266, 267, 7, 40, 2, 2, 267, 268, 7, 40, 2, 2, 268,
	74, 3, 2, 2, 2, 269, 270, 7, 126, 2, 2, 270, 271, 7, 126, 2, 2, 271, 76,
	3, 2, 2, 2, 272, 273, 5, 19, 10, 2, 273, 274, 5, 43, 22, 2, 274, 275, 5,
	41, 21, 2, 275, 276, 5, 19, 10, 2, 276, 78, 3, 2, 2, 2, 277, 278, 5, 31,
	16, 2, 278, 279, 5, 25, 13, 2, 279, 80, 3, 2, 2, 2, 280, 281, 5, 23, 12,
	2, 281, 282, 5, 37, 19, 2, 282, 283, 5, 51, 26, 2, 283, 284, 5, 23, 12,
	2, 284, 82, 3, 2, 2, 2, 285, 286, 5, 49, 25, 2, 286, 287, 5, 23, 12, 2,
	287, 288, 5, 53, 27, 2, 288, 289, 5, 55, 28, 2, 289, 290, 5, 49, 25, 2,
	290, 291, 5, 41, 21, 2, 291, 84, 3, 2, 2, 2, 292, 293, 5, 25, 13, 2, 293,
	294, 5, 43, 22, 2, 294, 295, 5, 49, 25, 2, 295, 86, 3, 2, 2, 2, 296, 297,
	5, 17, 9, 2, 297, 298, 5, 49, 25, 2, 298, 299, 5, 23, 12, 2, 299, 300,
	5, 15, 8, 2, 300, 301, 5, 35, 18, 2, 301, 88, 3, 2, 2, 2, 302, 303, 5,
	25, 13, 2, 303, 304, 5, 43, 22, 2, 304, 305, 5, 49, 25, 2, 305, 306, 5,
	49, 25, 2, 306, 307, 5, 15, 8, 2, 307, 308, 5, 41, 21, 2, 308, 309, 5,
	27, 14, 2, 309, 310, 5, 23, 12, 2, 310, 90, 3, 2, 2, 2, 311, 312, 5, 19,
	10, 2, 312, 313, 5, 43, 22, 2, 313, 314, 5, 41, 21, 2, 314, 315, 5, 53,
	27, 2, 315, 316, 5, 31, 16, 2, 316, 317, 5, 41, 21, 2, 317, 318, 5, 55,
	28, 2, 318, 319, 5, 23, 12, 2, 319, 92, 3, 2, 2, 2, 320, 321, 5, 53, 27,
	2, 321, 322, 5, 49, 25, 2, 322, 323, 5, 55, 28, 2, 323, 324, 5, 23, 12,
	2, 324, 94, 3, 2, 2, 2, 325, 326, 5, 25, 13, 2, 326, 327, 5, 15, 8, 2,
	327, 328, 5, 37, 19, 2, 328, 329, 5, 51, 26, 2, 329, 330, 5, 23, 12, 2,
	330, 96, 3, 2, 2, 2, 331, 332, 5, 41, 21, 2, 332, 333, 5, 55, 28, 2, 333,
	334, 5, 37, 19, 2, 334, 335, 5, 37, 19, 2, 335, 98, 3, 2, 2, 2, 336, 337,
	5, 51, 26, 2, 337, 338, 5, 15, 8, 2, 338, 339, 5, 37, 19, 2, 339, 340,
	5, 31, 16, 2, 340, 341, 5, 23, 12, 2, 341, 342, 5, 41, 21, 2, 342, 343,
	5, 19, 10, 2, 343, 344, 5, 23, 12, 2, 344, 100, 3, 2, 2, 2, 345, 346, 5,
	17, 9, 2, 346, 347, 5, 23, 12, 2, 347, 348, 5, 27, 14, 2, 348, 349, 5,
	31, 16, 2, 349, 350, 5, 41, 21, 2, 350, 102, 3, 2, 2, 2, 351, 352, 5, 23,
	12, 2, 352, 353, 5, 41, 21, 2, 353, 354, 5, 21, 11, 2, 354, 104, 3, 2,
	2, 2, 355, 357, 9, 29, 2, 2, 356, 355, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2,
	358, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 363, 3, 2, 2, 2, 360,
	362, 9, 30, 2, 2, 361, 360, 3, 2, 2, 2, 362, 365, 3, 2, 2, 2, 363, 361,
	3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 106, 3, 2, 2, 2, 365, 363, 3, 2,
	2, 2, 366, 368, 4, 50, 59, 2, 367, 366, 3, 2, 2, 2, 368, 369, 3, 2, 2,
	2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 108, 3, 2, 2, 2, 371,
	372, 7, 45, 2, 2, 372, 110, 3, 2, 2, 2, 373, 374, 7, 47, 2, 2, 374, 112,
	3, 2, 2, 2, 375, 376, 7, 49, 2, 2, 376, 114, 3, 2, 2, 2, 377, 378, 7, 44,
	2, 2, 378, 116, 3, 2, 2, 2, 379, 380, 7, 63, 2, 2, 380, 381, 7, 63, 2,
	2, 381, 118, 3, 2, 2, 2, 382, 383, 7, 64, 2, 2, 383, 120, 3, 2, 2, 2, 384,
	385, 7, 62, 2, 2, 385, 122, 3, 2, 2, 2, 386, 387, 7, 64, 2, 2, 387, 388,
	7, 63, 2, 2, 388, 124, 3, 2, 2, 2, 389, 390, 7, 62, 2, 2, 390, 391, 7,
	63, 2, 2, 391, 126, 3, 2, 2, 2, 392, 393, 7, 35, 2, 2, 393, 394, 7, 63,
	2, 2, 394, 128, 3, 2, 2, 2, 395, 396, 7, 35, 2, 2, 396, 130, 3, 2, 2, 2,
	397, 398, 7, 60, 2, 2, 398, 399, 7, 63, 2, 2, 399, 132, 3, 2, 2, 2, 400,
	401, 7, 63, 2, 2, 401, 134, 3, 2, 2, 2, 402, 403, 7, 45, 2, 2, 403, 404,
	7, 63, 2, 2, 404, 136, 3, 2, 2, 2, 405, 406, 7, 47, 2, 2, 406, 407, 7,
	63, 2, 2, 407, 138, 3, 2, 2, 2, 408, 409, 7, 44, 2, 2, 409, 410, 7, 63,
	2, 2, 410, 140, 3, 2, 2, 2, 411, 412, 7, 49, 2, 2, 412, 413, 7, 63, 2,
	2, 413, 142, 3, 2, 2, 2, 414, 415, 7, 93, 2, 2, 415, 144, 3, 2, 2, 2, 416,
	417, 7, 95, 2, 2, 417, 146, 3, 2, 2, 2, 418, 419, 7, 61, 2, 2, 419, 148,
	3, 2, 2, 2, 420, 421, 7, 125, 2, 2, 421, 150, 3, 2, 2, 2, 422, 423, 7,
	127, 2, 2, 423, 152, 3, 2, 2, 2, 424, 425, 7, 42, 2, 2, 425, 154, 3, 2,
	2, 2, 426, 427, 7, 43, 2, 2, 427, 156, 3, 2, 2, 2, 428, 429, 7, 48, 2,
	2, 429, 158, 3, 2, 2, 2, 430, 438, 7, 36, 2, 2, 431, 432, 7, 94, 2, 2,
	432, 437, 11, 2, 2, 2, 433, 434, 7, 36, 2, 2, 434, 437, 7, 36, 2, 2, 435,
	437, 10, 31, 2, 2, 436, 431, 3, 2, 2, 2, 436, 433, 3, 2, 2, 2, 436, 435,
	3, 2, 2, 2, 437, 440, 3, 2, 2, 2, 438, 436, 3, 2, 2, 2, 438, 439, 3, 2,
	2, 2, 439, 441, 3, 2, 2, 2, 440, 438, 3, 2, 2, 2, 441, 442, 7, 36, 2, 2,
	442, 160, 3, 2, 2, 2, 443, 444, 5, 105, 53, 2, 444, 445, 5, 157, 79, 2,
	445, 446, 5, 105, 53, 2, 446, 162, 3, 2, 2, 2, 447, 448, 5, 105, 53, 2,
	448, 449, 5, 157, 79, 2, 449, 450, 5, 105, 53, 2, 450, 451, 5, 157, 79,
	2, 451, 452, 5, 105, 53, 2, 452, 164, 3, 2, 2, 2, 453, 455, 5, 13, 7, 2,
	454, 453, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 454, 3, 2, 2, 2, 456,
	457, 3, 2, 2, 2, 457, 459, 3, 2, 2, 2, 458, 454, 3, 2, 2, 2, 458, 459,
	3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460, 462, 7, 48, 2, 2, 461, 463, 5, 13,
	7, 2, 462, 461, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2,
	464, 465, 3, 2, 2, 2, 465, 497, 3, 2, 2, 2, 466, 468, 5, 13, 7, 2, 467,
	466, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2, 469, 470,
	3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 472, 7, 48, 2, 2, 472, 473, 5, 67,
	34, 2, 473, 497, 3, 2, 2, 2, 474, 476, 5, 13, 7, 2, 475, 474, 3, 2, 2,
	2, 476, 477, 3, 2, 2, 2, 477, 475, 3, 2, 2, 2, 477, 478, 3, 2, 2, 2, 478,
	480, 3, 2, 2, 2, 479, 475, 3, 2, 2, 2, 479, 480, 3, 2, 2, 2, 480, 481,
	3, 2, 2, 2, 481, 483, 7, 48, 2, 2, 482, 484, 5, 13, 7, 2, 483, 482, 3,
	2, 2, 2, 484, 485, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2, 485, 486, 3, 2, 2,
	2, 486, 487, 3, 2, 2, 2, 487, 488, 5, 67, 34, 2, 488, 497, 3, 2, 2, 2,
	489, 491, 5, 13, 7, 2, 490, 489, 3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492,
	490, 3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494, 495,
	5, 67, 34, 2, 495, 497, 3, 2, 2, 2, 496, 458, 3, 2, 2, 2, 496, 467, 3,
	2, 2, 2, 496, 479, 3, 2, 2, 2, 496, 490, 3, 2, 2, 2, 497, 166, 3, 2, 2,
	2, 498, 499, 7, 49, 2, 2, 499, 500, 7, 49, 2, 2, 500, 504, 3, 2, 2, 2,
	501, 503, 11, 2, 2, 2, 502, 501, 3, 2, 2, 2, 503, 506, 3, 2, 2, 2, 504,
	505, 3, 2, 2, 2, 504, 502, 3, 2, 2, 2, 505, 507, 3, 2, 2, 2, 506, 504,
	3, 2, 2, 2, 507, 508, 7, 12, 2, 2, 508, 509, 3, 2, 2, 2, 509, 510, 8, 84,
	2, 2, 510, 168, 3, 2, 2, 2, 511, 513, 9, 32, 2, 2, 512, 511, 3, 2, 2, 2,
	513, 514, 3, 2, 2, 2, 514, 512, 3, 2, 2, 2, 514, 515, 3, 2, 2, 2, 515,
	516, 3, 2, 2, 2, 516, 517, 8, 85, 2, 2, 517, 170, 3, 2, 2, 2, 22, 2, 250,
	255, 358, 361, 363, 369, 436, 438, 456, 458, 464, 469, 477, 479, 485, 492,
	496, 504, 514, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'@name'", "'@id'", "'@desc'", "'@sal'", "", "", "'&&'", "'||'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'+'",
	"'-'", "'/'", "'*'", "'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'!'",
	"':='", "'='", "'+='", "'-='", "'*='", "'/='", "'['", "']'", "';'", "'{'",
	"'}'", "'('", "')'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "CONC", "IF", "ELSE",
	"RETURN", "FOR", "BREAK", "FORRANGE", "CONTINUE", "TRUE", "FALSE", "NULL_LITERAL",
	"SALIENCE", "BEGIN", "END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV",
	"MUL", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN",
	"SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE",
	"SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT",
	"DQUOTA_STRING", "DOTTEDNAME", "DOUBLEDOTTEDNAME", "REAL_LITERAL", "SL_COMMENT",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "DEC_DIGIT", "A", "B", "C", "D",
	"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
	"T", "U", "V", "W", "X", "Y", "Z", "EXPONENT_NUM_PART", "NIL", "RULE",
	"AND", "OR", "CONC", "IF", "ELSE", "RETURN", "FOR", "BREAK", "FORRANGE",
	"CONTINUE", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END",
	"SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT",
	"GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL",
	"MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE",
	"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME",
	"DOUBLEDOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
}

type gengineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewgengineLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *gengineLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgengineLexer(input antlr.CharStream) *gengineLexer {
	l := new(gengineLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "gengine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gengineLexer tokens.
const (
	gengineLexerT__0             = 1
	gengineLexerT__1             = 2
	gengineLexerT__2             = 3
	gengineLexerT__3             = 4
	gengineLexerT__4             = 5
	gengineLexerNIL              = 6
	gengineLexerRULE             = 7
	gengineLexerAND              = 8
	gengineLexerOR               = 9
	gengineLexerCONC             = 10
	gengineLexerIF               = 11
	gengineLexerELSE             = 12
	gengineLexerRETURN           = 13
	gengineLexerFOR              = 14
	gengineLexerBREAK            = 15
	gengineLexerFORRANGE         = 16
	gengineLexerCONTINUE         = 17
	gengineLexerTRUE             = 18
	gengineLexerFALSE            = 19
	gengineLexerNULL_LITERAL     = 20
	gengineLexerSALIENCE         = 21
	gengineLexerBEGIN            = 22
	gengineLexerEND              = 23
	gengineLexerSIMPLENAME       = 24
	gengineLexerINT              = 25
	gengineLexerPLUS             = 26
	gengineLexerMINUS            = 27
	gengineLexerDIV              = 28
	gengineLexerMUL              = 29
	gengineLexerEQUALS           = 30
	gengineLexerGT               = 31
	gengineLexerLT               = 32
	gengineLexerGTE              = 33
	gengineLexerLTE              = 34
	gengineLexerNOTEQUALS        = 35
	gengineLexerNOT              = 36
	gengineLexerASSIGN           = 37
	gengineLexerSET              = 38
	gengineLexerPLUSEQUAL        = 39
	gengineLexerMINUSEQUAL       = 40
	gengineLexerMULTIEQUAL       = 41
	gengineLexerDIVEQUAL         = 42
	gengineLexerLSQARE           = 43
	gengineLexerRSQARE           = 44
	gengineLexerSEMICOLON        = 45
	gengineLexerLR_BRACE         = 46
	gengineLexerRR_BRACE         = 47
	gengineLexerLR_BRACKET       = 48
	gengineLexerRR_BRACKET       = 49
	gengineLexerDOT              = 50
	gengineLexerDQUOTA_STRING    = 51
	gengineLexerDOTTEDNAME       = 52
	gengineLexerDOUBLEDOTTEDNAME = 53
	gengineLexerREAL_LITERAL     = 54
	gengineLexerSL_COMMENT       = 55
	gengineLexerWS               = 56
)