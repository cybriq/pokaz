// Package color provides the material design color palette.
//
// Source: https://material.google.com/style/color.html
// from: https://github.com/leobcn/materialcolors/blob/master/colors.go
package color

import (
	"image/color"
)

type Color struct {
	color.NRGBA
}

func getColor(r, g, b, a uint8) Color {
	return Color{
		NRGBA: color.NRGBA{R: r, G: g, B: b, A: a},
	}
}

// Material design color definitions.
var (
	Red50          = getColor(0xff, 0xeb, 0xee, 0xff)
	Red100         = getColor(0xff, 0xcd, 0xd2, 0xff)
	Red200         = getColor(0xef, 0x9a, 0x9a, 0xff)
	Red300         = getColor(0xe5, 0x73, 0x73, 0xff)
	Red400         = getColor(0xef, 0x53, 0x50, 0xff)
	Red500         = getColor(0xf4, 0x43, 0x36, 0xff)
	Red600         = getColor(0xe5, 0x39, 0x35, 0xff)
	Red700         = getColor(0xd3, 0x2f, 0x2f, 0xff)
	Red800         = getColor(0xc6, 0x28, 0x28, 0xff)
	Red900         = getColor(0xb7, 0x1c, 0x1c, 0xff)
	RedA100        = getColor(0xff, 0x8a, 0x80, 0xff)
	RedA200        = getColor(0xff, 0x52, 0x52, 0xff)
	RedA400        = getColor(0xff, 0x17, 0x44, 0xff)
	RedA700        = getColor(0xd5, 0x00, 0x00, 0xff)
	Pink50         = getColor(0xfc, 0xe4, 0xec, 0xff)
	Pink100        = getColor(0xf8, 0xbb, 0xd0, 0xff)
	Pink200        = getColor(0xf4, 0x8f, 0xb1, 0xff)
	Pink300        = getColor(0xf0, 0x62, 0x92, 0xff)
	Pink400        = getColor(0xec, 0x40, 0x7a, 0xff)
	Pink500        = getColor(0xe9, 0x1e, 0x63, 0xff)
	Pink600        = getColor(0xd8, 0x1b, 0x60, 0xff)
	Pink700        = getColor(0xc2, 0x18, 0x5b, 0xff)
	Pink800        = getColor(0xad, 0x14, 0x57, 0xff)
	Pink900        = getColor(0x88, 0x0e, 0x4f, 0xff)
	PinkA100       = getColor(0xff, 0x80, 0xab, 0xff)
	PinkA200       = getColor(0xff, 0x40, 0x81, 0xff)
	PinkA400       = getColor(0xf5, 0x00, 0x57, 0xff)
	PinkA700       = getColor(0xc5, 0x11, 0x62, 0xff)
	Purple50       = getColor(0xf3, 0xe5, 0xf5, 0xff)
	Purple100      = getColor(0xe1, 0xbe, 0xe7, 0xff)
	Purple200      = getColor(0xce, 0x93, 0xd8, 0xff)
	Purple300      = getColor(0xba, 0x68, 0xc8, 0xff)
	Purple400      = getColor(0xab, 0x47, 0xbc, 0xff)
	Purple500      = getColor(0x9c, 0x27, 0xb0, 0xff)
	Purple600      = getColor(0x8e, 0x24, 0xaa, 0xff)
	Purple700      = getColor(0x7b, 0x1f, 0xa2, 0xff)
	Purple800      = getColor(0x6a, 0x1b, 0x9a, 0xff)
	Purple900      = getColor(0x4a, 0x14, 0x8c, 0xff)
	PurpleA100     = getColor(0xea, 0x80, 0xfc, 0xff)
	PurpleA200     = getColor(0xe0, 0x40, 0xfb, 0xff)
	PurpleA400     = getColor(0xd5, 0x00, 0xf9, 0xff)
	PurpleA700     = getColor(0xaa, 0x00, 0xff, 0xff)
	DeepPurple50   = getColor(0xed, 0xe7, 0xf6, 0xff)
	DeepPurple100  = getColor(0xd1, 0xc4, 0xe9, 0xff)
	DeepPurple200  = getColor(0xb3, 0x9d, 0xdb, 0xff)
	DeepPurple300  = getColor(0x95, 0x75, 0xcd, 0xff)
	DeepPurple400  = getColor(0x7e, 0x57, 0xc2, 0xff)
	DeepPurple500  = getColor(0x67, 0x3a, 0xb7, 0xff)
	DeepPurple600  = getColor(0x5e, 0x35, 0xb1, 0xff)
	DeepPurple700  = getColor(0x51, 0x2d, 0xa8, 0xff)
	DeepPurple800  = getColor(0x45, 0x27, 0xa0, 0xff)
	DeepPurple900  = getColor(0x31, 0x1b, 0x92, 0xff)
	DeepPurpleA100 = getColor(0xb3, 0x88, 0xff, 0xff)
	DeepPurpleA200 = getColor(0x7c, 0x4d, 0xff, 0xff)
	DeepPurpleA400 = getColor(0x65, 0x1f, 0xff, 0xff)
	DeepPurpleA700 = getColor(0x62, 0x00, 0xea, 0xff)
	Indigo50       = getColor(0xe8, 0xea, 0xf6, 0xff)
	Indigo100      = getColor(0xc5, 0xca, 0xe9, 0xff)
	Indigo200      = getColor(0x9f, 0xa8, 0xda, 0xff)
	Indigo300      = getColor(0x79, 0x86, 0xcb, 0xff)
	Indigo400      = getColor(0x5c, 0x6b, 0xc0, 0xff)
	Indigo500      = getColor(0x3f, 0x51, 0xb5, 0xff)
	Indigo600      = getColor(0x39, 0x49, 0xab, 0xff)
	Indigo700      = getColor(0x30, 0x3f, 0x9f, 0xff)
	Indigo800      = getColor(0x28, 0x35, 0x93, 0xff)
	Indigo900      = getColor(0x1a, 0x23, 0x7e, 0xff)
	IndigoA100     = getColor(0x8c, 0x9e, 0xff, 0xff)
	IndigoA200     = getColor(0x53, 0x6d, 0xfe, 0xff)
	IndigoA400     = getColor(0x3d, 0x5a, 0xfe, 0xff)
	IndigoA700     = getColor(0x30, 0x4f, 0xfe, 0xff)
	Blue50         = getColor(0xe3, 0xf2, 0xfd, 0xff)
	Blue100        = getColor(0xbb, 0xde, 0xfb, 0xff)
	Blue200        = getColor(0x90, 0xca, 0xf9, 0xff)
	Blue300        = getColor(0x64, 0xb5, 0xf6, 0xff)
	Blue400        = getColor(0x42, 0xa5, 0xf5, 0xff)
	Blue500        = getColor(0x21, 0x96, 0xf3, 0xff)
	Blue600        = getColor(0x1e, 0x88, 0xe5, 0xff)
	Blue700        = getColor(0x19, 0x76, 0xd2, 0xff)
	Blue800        = getColor(0x15, 0x65, 0xc0, 0xff)
	Blue900        = getColor(0x0d, 0x47, 0xa1, 0xff)
	BlueA100       = getColor(0x82, 0xb1, 0xff, 0xff)
	BlueA200       = getColor(0x44, 0x8a, 0xff, 0xff)
	BlueA400       = getColor(0x29, 0x79, 0xff, 0xff)
	BlueA700       = getColor(0x29, 0x62, 0xff, 0xff)
	LightBlue50    = getColor(0xe1, 0xf5, 0xfe, 0xff)
	LightBlue100   = getColor(0xb3, 0xe5, 0xfc, 0xff)
	LightBlue200   = getColor(0x81, 0xd4, 0xfa, 0xff)
	LightBlue300   = getColor(0x4f, 0xc3, 0xf7, 0xff)
	LightBlue400   = getColor(0x29, 0xb6, 0xf6, 0xff)
	LightBlue500   = getColor(0x03, 0xa9, 0xf4, 0xff)
	LightBlue600   = getColor(0x03, 0x9b, 0xe5, 0xff)
	LightBlue700   = getColor(0x02, 0x88, 0xd1, 0xff)
	LightBlue800   = getColor(0x02, 0x77, 0xbd, 0xff)
	LightBlue900   = getColor(0x01, 0x57, 0x9b, 0xff)
	LightBlueA100  = getColor(0x80, 0xd8, 0xff, 0xff)
	LightBlueA200  = getColor(0x40, 0xc4, 0xff, 0xff)
	LightBlueA400  = getColor(0x00, 0xb0, 0xff, 0xff)
	LightBlueA700  = getColor(0x00, 0x91, 0xea, 0xff)
	Cyan50         = getColor(0xe0, 0xf7, 0xfa, 0xff)
	Cyan100        = getColor(0xb2, 0xeb, 0xf2, 0xff)
	Cyan200        = getColor(0x80, 0xde, 0xea, 0xff)
	Cyan300        = getColor(0x4d, 0xd0, 0xe1, 0xff)
	Cyan400        = getColor(0x26, 0xc6, 0xda, 0xff)
	Cyan500        = getColor(0x00, 0xbc, 0xd4, 0xff)
	Cyan600        = getColor(0x00, 0xac, 0xc1, 0xff)
	Cyan700        = getColor(0x00, 0x97, 0xa7, 0xff)
	Cyan800        = getColor(0x00, 0x83, 0x8f, 0xff)
	Cyan900        = getColor(0x00, 0x60, 0x64, 0xff)
	CyanA100       = getColor(0x84, 0xff, 0xff, 0xff)
	CyanA200       = getColor(0x18, 0xff, 0xff, 0xff)
	CyanA400       = getColor(0x00, 0xe5, 0xff, 0xff)
	CyanA700       = getColor(0x00, 0xb8, 0xd4, 0xff)
	Teal50         = getColor(0xe0, 0xf2, 0xf1, 0xff)
	Teal100        = getColor(0xb2, 0xdf, 0xdb, 0xff)
	Teal200        = getColor(0x80, 0xcb, 0xc4, 0xff)
	Teal300        = getColor(0x4d, 0xb6, 0xac, 0xff)
	Teal400        = getColor(0x26, 0xa6, 0x9a, 0xff)
	Teal500        = getColor(0x00, 0x96, 0x88, 0xff)
	Teal600        = getColor(0x00, 0x89, 0x7b, 0xff)
	Teal700        = getColor(0x00, 0x79, 0x6b, 0xff)
	Teal800        = getColor(0x00, 0x69, 0x5c, 0xff)
	Teal900        = getColor(0x00, 0x4d, 0x40, 0xff)
	TealA100       = getColor(0xa7, 0xff, 0xeb, 0xff)
	TealA200       = getColor(0x64, 0xff, 0xda, 0xff)
	TealA400       = getColor(0x1d, 0xe9, 0xb6, 0xff)
	TealA700       = getColor(0x00, 0xbf, 0xa5, 0xff)
	Green50        = getColor(0xe8, 0xf5, 0xe9, 0xff)
	Green100       = getColor(0xc8, 0xe6, 0xc9, 0xff)
	Green200       = getColor(0xa5, 0xd6, 0xa7, 0xff)
	Green300       = getColor(0x81, 0xc7, 0x84, 0xff)
	Green400       = getColor(0x66, 0xbb, 0x6a, 0xff)
	Green500       = getColor(0x4c, 0xaf, 0x50, 0xff)
	Green600       = getColor(0x43, 0xa0, 0x47, 0xff)
	Green700       = getColor(0x38, 0x8e, 0x3c, 0xff)
	Green800       = getColor(0x2e, 0x7d, 0x32, 0xff)
	Green900       = getColor(0x1b, 0x5e, 0x20, 0xff)
	GreenA100      = getColor(0xb9, 0xf6, 0xca, 0xff)
	GreenA200      = getColor(0x69, 0xf0, 0xae, 0xff)
	GreenA400      = getColor(0x00, 0xe6, 0x76, 0xff)
	GreenA700      = getColor(0x00, 0xc8, 0x53, 0xff)
	LightGreen50   = getColor(0xf1, 0xf8, 0xe9, 0xff)
	LightGreen100  = getColor(0xdc, 0xed, 0xc8, 0xff)
	LightGreen200  = getColor(0xc5, 0xe1, 0xa5, 0xff)
	LightGreen300  = getColor(0xae, 0xd5, 0x81, 0xff)
	LightGreen400  = getColor(0x9c, 0xcc, 0x65, 0xff)
	LightGreen500  = getColor(0x8b, 0xc3, 0x4a, 0xff)
	LightGreen600  = getColor(0x7c, 0xb3, 0x42, 0xff)
	LightGreen700  = getColor(0x68, 0x9f, 0x38, 0xff)
	LightGreen800  = getColor(0x55, 0x8b, 0x2f, 0xff)
	LightGreen900  = getColor(0x33, 0x69, 0x1e, 0xff)
	LightGreenA100 = getColor(0xcc, 0xff, 0x90, 0xff)
	LightGreenA200 = getColor(0xb2, 0xff, 0x59, 0xff)
	LightGreenA400 = getColor(0x76, 0xff, 0x03, 0xff)
	LightGreenA700 = getColor(0x64, 0xdd, 0x17, 0xff)
	Lime50         = getColor(0xf9, 0xfb, 0xe7, 0xff)
	Lime100        = getColor(0xf0, 0xf4, 0xc3, 0xff)
	Lime200        = getColor(0xe6, 0xee, 0x9c, 0xff)
	Lime300        = getColor(0xdc, 0xe7, 0x75, 0xff)
	Lime400        = getColor(0xd4, 0xe1, 0x57, 0xff)
	Lime500        = getColor(0xcd, 0xdc, 0x39, 0xff)
	Lime600        = getColor(0xc0, 0xca, 0x33, 0xff)
	Lime700        = getColor(0xaf, 0xb4, 0x2b, 0xff)
	Lime800        = getColor(0x9e, 0x9d, 0x24, 0xff)
	Lime900        = getColor(0x82, 0x77, 0x17, 0xff)
	LimeA100       = getColor(0xf4, 0xff, 0x81, 0xff)
	LimeA200       = getColor(0xee, 0xff, 0x41, 0xff)
	LimeA400       = getColor(0xc6, 0xff, 0x00, 0xff)
	LimeA700       = getColor(0xae, 0xea, 0x00, 0xff)
	Yellow50       = getColor(0xff, 0xfd, 0xe7, 0xff)
	Yellow100      = getColor(0xff, 0xf9, 0xc4, 0xff)
	Yellow200      = getColor(0xff, 0xf5, 0x9d, 0xff)
	Yellow300      = getColor(0xff, 0xf1, 0x76, 0xff)
	Yellow400      = getColor(0xff, 0xee, 0x58, 0xff)
	Yellow500      = getColor(0xff, 0xeb, 0x3b, 0xff)
	Yellow600      = getColor(0xfd, 0xd8, 0x35, 0xff)
	Yellow700      = getColor(0xfb, 0xc0, 0x2d, 0xff)
	Yellow800      = getColor(0xf9, 0xa8, 0x25, 0xff)
	Yellow900      = getColor(0xf5, 0x7f, 0x17, 0xff)
	YellowA100     = getColor(0xff, 0xff, 0x8d, 0xff)
	YellowA200     = getColor(0xff, 0xff, 0x00, 0xff)
	YellowA400     = getColor(0xff, 0xea, 0x00, 0xff)
	YellowA700     = getColor(0xff, 0xd6, 0x00, 0xff)
	Amber50        = getColor(0xff, 0xf8, 0xe1, 0xff)
	Amber100       = getColor(0xff, 0xec, 0xb3, 0xff)
	Amber200       = getColor(0xff, 0xe0, 0x82, 0xff)
	Amber300       = getColor(0xff, 0xd5, 0x4f, 0xff)
	Amber400       = getColor(0xff, 0xca, 0x28, 0xff)
	Amber500       = getColor(0xff, 0xc1, 0x07, 0xff)
	Amber600       = getColor(0xff, 0xb3, 0x00, 0xff)
	Amber700       = getColor(0xff, 0xa0, 0x00, 0xff)
	Amber800       = getColor(0xff, 0x8f, 0x00, 0xff)
	Amber900       = getColor(0xff, 0x6f, 0x00, 0xff)
	AmberA100      = getColor(0xff, 0xe5, 0x7f, 0xff)
	AmberA200      = getColor(0xff, 0xd7, 0x40, 0xff)
	AmberA400      = getColor(0xff, 0xc4, 0x00, 0xff)
	AmberA700      = getColor(0xff, 0xab, 0x00, 0xff)
	Orange50       = getColor(0xff, 0xf3, 0xe0, 0xff)
	Orange100      = getColor(0xff, 0xe0, 0xb2, 0xff)
	Orange200      = getColor(0xff, 0xcc, 0x80, 0xff)
	Orange300      = getColor(0xff, 0xb7, 0x4d, 0xff)
	Orange400      = getColor(0xff, 0xa7, 0x26, 0xff)
	Orange500      = getColor(0xff, 0x98, 0x00, 0xff)
	Orange600      = getColor(0xfb, 0x8c, 0x00, 0xff)
	Orange700      = getColor(0xf5, 0x7c, 0x00, 0xff)
	Orange800      = getColor(0xef, 0x6c, 0x00, 0xff)
	Orange900      = getColor(0xe6, 0x51, 0x00, 0xff)
	OrangeA100     = getColor(0xff, 0xd1, 0x80, 0xff)
	OrangeA200     = getColor(0xff, 0xab, 0x40, 0xff)
	OrangeA400     = getColor(0xff, 0x91, 0x00, 0xff)
	OrangeA700     = getColor(0xff, 0x6d, 0x00, 0xff)
	DeepOrange50   = getColor(0xfb, 0xe9, 0xe7, 0xff)
	DeepOrange100  = getColor(0xff, 0xcc, 0xbc, 0xff)
	DeepOrange200  = getColor(0xff, 0xab, 0x91, 0xff)
	DeepOrange300  = getColor(0xff, 0x8a, 0x65, 0xff)
	DeepOrange400  = getColor(0xff, 0x70, 0x43, 0xff)
	DeepOrange500  = getColor(0xff, 0x57, 0x22, 0xff)
	DeepOrange600  = getColor(0xf4, 0x51, 0x1e, 0xff)
	DeepOrange700  = getColor(0xe6, 0x4a, 0x19, 0xff)
	DeepOrange800  = getColor(0xd8, 0x43, 0x15, 0xff)
	DeepOrange900  = getColor(0xbf, 0x36, 0x0c, 0xff)
	DeepOrangeA100 = getColor(0xff, 0x9e, 0x80, 0xff)
	DeepOrangeA200 = getColor(0xff, 0x6e, 0x40, 0xff)
	DeepOrangeA400 = getColor(0xff, 0x3d, 0x00, 0xff)
	DeepOrangeA700 = getColor(0xdd, 0x2c, 0x00, 0xff)
	Brown50        = getColor(0xef, 0xeb, 0xe9, 0xff)
	Brown100       = getColor(0xd7, 0xcc, 0xc8, 0xff)
	Brown200       = getColor(0xbc, 0xaa, 0xa4, 0xff)
	Brown300       = getColor(0xa1, 0x88, 0x7f, 0xff)
	Brown400       = getColor(0x8d, 0x6e, 0x63, 0xff)
	Brown500       = getColor(0x79, 0x55, 0x48, 0xff)
	Brown600       = getColor(0x6d, 0x4c, 0x41, 0xff)
	Brown700       = getColor(0x5d, 0x40, 0x37, 0xff)
	Brown800       = getColor(0x4e, 0x34, 0x2e, 0xff)
	Brown900       = getColor(0x3e, 0x27, 0x23, 0xff)
	Grey50         = getColor(0xfa, 0xfa, 0xfa, 0xff)
	Grey100        = getColor(0xf5, 0xf5, 0xf5, 0xff)
	Grey200        = getColor(0xee, 0xee, 0xee, 0xff)
	Grey300        = getColor(0xe0, 0xe0, 0xe0, 0xff)
	Grey400        = getColor(0xbd, 0xbd, 0xbd, 0xff)
	Grey500        = getColor(0x9e, 0x9e, 0x9e, 0xff)
	Grey600        = getColor(0x75, 0x75, 0x75, 0xff)
	Grey700        = getColor(0x61, 0x61, 0x61, 0xff)
	Grey800        = getColor(0x42, 0x42, 0x42, 0xff)
	Grey900        = getColor(0x21, 0x21, 0x21, 0xff)
	BlueGrey50     = getColor(0xec, 0xef, 0xf1, 0xff)
	BlueGrey100    = getColor(0xcf, 0xd8, 0xdc, 0xff)
	BlueGrey200    = getColor(0xb0, 0xbe, 0xc5, 0xff)
	BlueGrey300    = getColor(0x90, 0xa4, 0xae, 0xff)
	BlueGrey400    = getColor(0x78, 0x90, 0x9c, 0xff)
	BlueGrey500    = getColor(0x60, 0x7d, 0x8b, 0xff)
	BlueGrey600    = getColor(0x54, 0x6e, 0x7a, 0xff)
	BlueGrey700    = getColor(0x45, 0x5a, 0x64, 0xff)
	BlueGrey800    = getColor(0x37, 0x47, 0x4f, 0xff)
	BlueGrey900    = getColor(0x26, 0x32, 0x38, 0xff)
	Black          = getColor(0x00, 0x00, 0x00, 0xff)
	White          = getColor(0xff, 0xff, 0xff, 0xff)
)

// Palette is the material design color palette.
var Palette = []Color{
	Red50,
	Red100,
	Red200,
	Red300,
	Red400,
	Red500,
	Red600,
	Red700,
	Red800,
	Red900,
	RedA100,
	RedA200,
	RedA400,
	RedA700,
	Pink50,
	Pink100,
	Pink200,
	Pink300,
	Pink400,
	Pink500,
	Pink600,
	Pink700,
	Pink800,
	Pink900,
	PinkA100,
	PinkA200,
	PinkA400,
	PinkA700,
	Purple50,
	Purple100,
	Purple200,
	Purple300,
	Purple400,
	Purple500,
	Purple600,
	Purple700,
	Purple800,
	Purple900,
	PurpleA100,
	PurpleA200,
	PurpleA400,
	PurpleA700,
	DeepPurple50,
	DeepPurple100,
	DeepPurple200,
	DeepPurple300,
	DeepPurple400,
	DeepPurple500,
	DeepPurple600,
	DeepPurple700,
	DeepPurple800,
	DeepPurple900,
	DeepPurpleA100,
	DeepPurpleA200,
	DeepPurpleA400,
	DeepPurpleA700,
	Indigo50,
	Indigo100,
	Indigo200,
	Indigo300,
	Indigo400,
	Indigo500,
	Indigo600,
	Indigo700,
	Indigo800,
	Indigo900,
	IndigoA100,
	IndigoA200,
	IndigoA400,
	IndigoA700,
	Blue50,
	Blue100,
	Blue200,
	Blue300,
	Blue400,
	Blue500,
	Blue600,
	Blue700,
	Blue800,
	Blue900,
	BlueA100,
	BlueA200,
	BlueA400,
	BlueA700,
	LightBlue50,
	LightBlue100,
	LightBlue200,
	LightBlue300,
	LightBlue400,
	LightBlue500,
	LightBlue600,
	LightBlue700,
	LightBlue800,
	LightBlue900,
	LightBlueA100,
	LightBlueA200,
	LightBlueA400,
	LightBlueA700,
	Cyan50,
	Cyan100,
	Cyan200,
	Cyan300,
	Cyan400,
	Cyan500,
	Cyan600,
	Cyan700,
	Cyan800,
	Cyan900,
	CyanA100,
	CyanA200,
	CyanA400,
	CyanA700,
	Teal50,
	Teal100,
	Teal200,
	Teal300,
	Teal400,
	Teal500,
	Teal600,
	Teal700,
	Teal800,
	Teal900,
	TealA100,
	TealA200,
	TealA400,
	TealA700,
	Green50,
	Green100,
	Green200,
	Green300,
	Green400,
	Green500,
	Green600,
	Green700,
	Green800,
	Green900,
	GreenA100,
	GreenA200,
	GreenA400,
	GreenA700,
	LightGreen50,
	LightGreen100,
	LightGreen200,
	LightGreen300,
	LightGreen400,
	LightGreen500,
	LightGreen600,
	LightGreen700,
	LightGreen800,
	LightGreen900,
	LightGreenA100,
	LightGreenA200,
	LightGreenA400,
	LightGreenA700,
	Lime50,
	Lime100,
	Lime200,
	Lime300,
	Lime400,
	Lime500,
	Lime600,
	Lime700,
	Lime800,
	Lime900,
	LimeA100,
	LimeA200,
	LimeA400,
	LimeA700,
	Yellow50,
	Yellow100,
	Yellow200,
	Yellow300,
	Yellow400,
	Yellow500,
	Yellow600,
	Yellow700,
	Yellow800,
	Yellow900,
	YellowA100,
	YellowA200,
	YellowA400,
	YellowA700,
	Amber50,
	Amber100,
	Amber200,
	Amber300,
	Amber400,
	Amber500,
	Amber600,
	Amber700,
	Amber800,
	Amber900,
	AmberA100,
	AmberA200,
	AmberA400,
	AmberA700,
	Orange50,
	Orange100,
	Orange200,
	Orange300,
	Orange400,
	Orange500,
	Orange600,
	Orange700,
	Orange800,
	Orange900,
	OrangeA100,
	OrangeA200,
	OrangeA400,
	OrangeA700,
	DeepOrange50,
	DeepOrange100,
	DeepOrange200,
	DeepOrange300,
	DeepOrange400,
	DeepOrange500,
	DeepOrange600,
	DeepOrange700,
	DeepOrange800,
	DeepOrange900,
	DeepOrangeA100,
	DeepOrangeA200,
	DeepOrangeA400,
	DeepOrangeA700,
	Brown50,
	Brown100,
	Brown200,
	Brown300,
	Brown400,
	Brown500,
	Brown600,
	Brown700,
	Brown800,
	Brown900,
	Grey50,
	Grey100,
	Grey200,
	Grey300,
	Grey400,
	Grey500,
	Grey600,
	Grey700,
	Grey800,
	Grey900,
	BlueGrey50,
	BlueGrey100,
	BlueGrey200,
	BlueGrey300,
	BlueGrey400,
	BlueGrey500,
	BlueGrey600,
	BlueGrey700,
	BlueGrey800,
	BlueGrey900,
	Black,
	White,
}

// ByName is color names to color values map.
var ByName = map[string]Color{
	"Red50":          Red50,
	"Red100":         Red100,
	"Red200":         Red200,
	"Red300":         Red300,
	"Red400":         Red400,
	"Red500":         Red500,
	"Red600":         Red600,
	"Red700":         Red700,
	"Red800":         Red800,
	"Red900":         Red900,
	"RedA100":        RedA100,
	"RedA200":        RedA200,
	"RedA400":        RedA400,
	"RedA700":        RedA700,
	"Pink50":         Pink50,
	"Pink100":        Pink100,
	"Pink200":        Pink200,
	"Pink300":        Pink300,
	"Pink400":        Pink400,
	"Pink500":        Pink500,
	"Pink600":        Pink600,
	"Pink700":        Pink700,
	"Pink800":        Pink800,
	"Pink900":        Pink900,
	"PinkA100":       PinkA100,
	"PinkA200":       PinkA200,
	"PinkA400":       PinkA400,
	"PinkA700":       PinkA700,
	"Purple50":       Purple50,
	"Purple100":      Purple100,
	"Purple200":      Purple200,
	"Purple300":      Purple300,
	"Purple400":      Purple400,
	"Purple500":      Purple500,
	"Purple600":      Purple600,
	"Purple700":      Purple700,
	"Purple800":      Purple800,
	"Purple900":      Purple900,
	"PurpleA100":     PurpleA100,
	"PurpleA200":     PurpleA200,
	"PurpleA400":     PurpleA400,
	"PurpleA700":     PurpleA700,
	"DeepPurple50":   DeepPurple50,
	"DeepPurple100":  DeepPurple100,
	"DeepPurple200":  DeepPurple200,
	"DeepPurple300":  DeepPurple300,
	"DeepPurple400":  DeepPurple400,
	"DeepPurple500":  DeepPurple500,
	"DeepPurple600":  DeepPurple600,
	"DeepPurple700":  DeepPurple700,
	"DeepPurple800":  DeepPurple800,
	"DeepPurple900":  DeepPurple900,
	"DeepPurpleA100": DeepPurpleA100,
	"DeepPurpleA200": DeepPurpleA200,
	"DeepPurpleA400": DeepPurpleA400,
	"DeepPurpleA700": DeepPurpleA700,
	"Indigo50":       Indigo50,
	"Indigo100":      Indigo100,
	"Indigo200":      Indigo200,
	"Indigo300":      Indigo300,
	"Indigo400":      Indigo400,
	"Indigo500":      Indigo500,
	"Indigo600":      Indigo600,
	"Indigo700":      Indigo700,
	"Indigo800":      Indigo800,
	"Indigo900":      Indigo900,
	"IndigoA100":     IndigoA100,
	"IndigoA200":     IndigoA200,
	"IndigoA400":     IndigoA400,
	"IndigoA700":     IndigoA700,
	"Blue50":         Blue50,
	"Blue100":        Blue100,
	"Blue200":        Blue200,
	"Blue300":        Blue300,
	"Blue400":        Blue400,
	"Blue500":        Blue500,
	"Blue600":        Blue600,
	"Blue700":        Blue700,
	"Blue800":        Blue800,
	"Blue900":        Blue900,
	"BlueA100":       BlueA100,
	"BlueA200":       BlueA200,
	"BlueA400":       BlueA400,
	"BlueA700":       BlueA700,
	"LightBlue50":    LightBlue50,
	"LightBlue100":   LightBlue100,
	"LightBlue200":   LightBlue200,
	"LightBlue300":   LightBlue300,
	"LightBlue400":   LightBlue400,
	"LightBlue500":   LightBlue500,
	"LightBlue600":   LightBlue600,
	"LightBlue700":   LightBlue700,
	"LightBlue800":   LightBlue800,
	"LightBlue900":   LightBlue900,
	"LightBlueA100":  LightBlueA100,
	"LightBlueA200":  LightBlueA200,
	"LightBlueA400":  LightBlueA400,
	"LightBlueA700":  LightBlueA700,
	"Cyan50":         Cyan50,
	"Cyan100":        Cyan100,
	"Cyan200":        Cyan200,
	"Cyan300":        Cyan300,
	"Cyan400":        Cyan400,
	"Cyan500":        Cyan500,
	"Cyan600":        Cyan600,
	"Cyan700":        Cyan700,
	"Cyan800":        Cyan800,
	"Cyan900":        Cyan900,
	"CyanA100":       CyanA100,
	"CyanA200":       CyanA200,
	"CyanA400":       CyanA400,
	"CyanA700":       CyanA700,
	"Teal50":         Teal50,
	"Teal100":        Teal100,
	"Teal200":        Teal200,
	"Teal300":        Teal300,
	"Teal400":        Teal400,
	"Teal500":        Teal500,
	"Teal600":        Teal600,
	"Teal700":        Teal700,
	"Teal800":        Teal800,
	"Teal900":        Teal900,
	"TealA100":       TealA100,
	"TealA200":       TealA200,
	"TealA400":       TealA400,
	"TealA700":       TealA700,
	"Green50":        Green50,
	"Green100":       Green100,
	"Green200":       Green200,
	"Green300":       Green300,
	"Green400":       Green400,
	"Green500":       Green500,
	"Green600":       Green600,
	"Green700":       Green700,
	"Green800":       Green800,
	"Green900":       Green900,
	"GreenA100":      GreenA100,
	"GreenA200":      GreenA200,
	"GreenA400":      GreenA400,
	"GreenA700":      GreenA700,
	"LightGreen50":   LightGreen50,
	"LightGreen100":  LightGreen100,
	"LightGreen200":  LightGreen200,
	"LightGreen300":  LightGreen300,
	"LightGreen400":  LightGreen400,
	"LightGreen500":  LightGreen500,
	"LightGreen600":  LightGreen600,
	"LightGreen700":  LightGreen700,
	"LightGreen800":  LightGreen800,
	"LightGreen900":  LightGreen900,
	"LightGreenA100": LightGreenA100,
	"LightGreenA200": LightGreenA200,
	"LightGreenA400": LightGreenA400,
	"LightGreenA700": LightGreenA700,
	"Lime50":         Lime50,
	"Lime100":        Lime100,
	"Lime200":        Lime200,
	"Lime300":        Lime300,
	"Lime400":        Lime400,
	"Lime500":        Lime500,
	"Lime600":        Lime600,
	"Lime700":        Lime700,
	"Lime800":        Lime800,
	"Lime900":        Lime900,
	"LimeA100":       LimeA100,
	"LimeA200":       LimeA200,
	"LimeA400":       LimeA400,
	"LimeA700":       LimeA700,
	"Yellow50":       Yellow50,
	"Yellow100":      Yellow100,
	"Yellow200":      Yellow200,
	"Yellow300":      Yellow300,
	"Yellow400":      Yellow400,
	"Yellow500":      Yellow500,
	"Yellow600":      Yellow600,
	"Yellow700":      Yellow700,
	"Yellow800":      Yellow800,
	"Yellow900":      Yellow900,
	"YellowA100":     YellowA100,
	"YellowA200":     YellowA200,
	"YellowA400":     YellowA400,
	"YellowA700":     YellowA700,
	"Amber50":        Amber50,
	"Amber100":       Amber100,
	"Amber200":       Amber200,
	"Amber300":       Amber300,
	"Amber400":       Amber400,
	"Amber500":       Amber500,
	"Amber600":       Amber600,
	"Amber700":       Amber700,
	"Amber800":       Amber800,
	"Amber900":       Amber900,
	"AmberA100":      AmberA100,
	"AmberA200":      AmberA200,
	"AmberA400":      AmberA400,
	"AmberA700":      AmberA700,
	"Orange50":       Orange50,
	"Orange100":      Orange100,
	"Orange200":      Orange200,
	"Orange300":      Orange300,
	"Orange400":      Orange400,
	"Orange500":      Orange500,
	"Orange600":      Orange600,
	"Orange700":      Orange700,
	"Orange800":      Orange800,
	"Orange900":      Orange900,
	"OrangeA100":     OrangeA100,
	"OrangeA200":     OrangeA200,
	"OrangeA400":     OrangeA400,
	"OrangeA700":     OrangeA700,
	"DeepOrange50":   DeepOrange50,
	"DeepOrange100":  DeepOrange100,
	"DeepOrange200":  DeepOrange200,
	"DeepOrange300":  DeepOrange300,
	"DeepOrange400":  DeepOrange400,
	"DeepOrange500":  DeepOrange500,
	"DeepOrange600":  DeepOrange600,
	"DeepOrange700":  DeepOrange700,
	"DeepOrange800":  DeepOrange800,
	"DeepOrange900":  DeepOrange900,
	"DeepOrangeA100": DeepOrangeA100,
	"DeepOrangeA200": DeepOrangeA200,
	"DeepOrangeA400": DeepOrangeA400,
	"DeepOrangeA700": DeepOrangeA700,
	"Brown50":        Brown50,
	"Brown100":       Brown100,
	"Brown200":       Brown200,
	"Brown300":       Brown300,
	"Brown400":       Brown400,
	"Brown500":       Brown500,
	"Brown600":       Brown600,
	"Brown700":       Brown700,
	"Brown800":       Brown800,
	"Brown900":       Brown900,
	"Grey50":         Grey50,
	"Grey100":        Grey100,
	"Grey200":        Grey200,
	"Grey300":        Grey300,
	"Grey400":        Grey400,
	"Grey500":        Grey500,
	"Grey600":        Grey600,
	"Grey700":        Grey700,
	"Grey800":        Grey800,
	"Grey900":        Grey900,
	"BlueGrey50":     BlueGrey50,
	"BlueGrey100":    BlueGrey100,
	"BlueGrey200":    BlueGrey200,
	"BlueGrey300":    BlueGrey300,
	"BlueGrey400":    BlueGrey400,
	"BlueGrey500":    BlueGrey500,
	"BlueGrey600":    BlueGrey600,
	"BlueGrey700":    BlueGrey700,
	"BlueGrey800":    BlueGrey800,
	"BlueGrey900":    BlueGrey900,
	"Black":          Black,
	"White":          White,
}
