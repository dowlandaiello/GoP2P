package unicodeh

import "unicode"

// Unicode property "NFC_Quick_Check" (known as "NFC_QC", "NFC_Quick_Check").
// Kind of property: "Enumerated".
// Based on file "DerivedNormalizationProps.txt".
var (
	NFCQuickCheckMaybe = nFCQuickCheckMaybe // Value "Maybe" (known as "M", "Maybe").
	NFCQuickCheckNo    = nFCQuickCheckNo    // Value "No" (known as "N", "No").
	NFCQuickCheckYes   = nFCQuickCheckYes   // Value "Yes" (known as "Y", "Yes").
)

var (
	nFCQuickCheckMaybe = &unicode.RangeTable{[]unicode.Range16{{0x300, 0x304, 0x1}, {0x306, 0x30c, 0x1}, {0x30f, 0x313, 0x2}, {0x314, 0x31b, 0x7}, {0x323, 0x328, 0x1}, {0x32d, 0x32e, 0x1}, {0x330, 0x331, 0x1}, {0x338, 0x342, 0xa}, {0x345, 0x653, 0x30e}, {0x654, 0x655, 0x1}, {0x93c, 0x9be, 0x82}, {0x9d7, 0xb3e, 0x167}, {0xb56, 0xb57, 0x1}, {0xbbe, 0xbd7, 0x19}, {0xc56, 0xcc2, 0x6c}, {0xcd5, 0xcd6, 0x1}, {0xd3e, 0xd57, 0x19}, {0xdca, 0xdcf, 0x5}, {0xddf, 0x102e, 0x24f}, {0x1161, 0x1175, 0x1}, {0x11a8, 0x11c2, 0x1}, {0x1b35, 0x3099, 0x1564}, {0x309a, 0x309a, 0x1}}, []unicode.Range32{{0x110ba, 0x11127, 0x6d}, {0x1133e, 0x11357, 0x19}, {0x114b0, 0x114ba, 0xa}, {0x114bd, 0x115af, 0xf2}}, 0}
	nFCQuickCheckNo    = &unicode.RangeTable{[]unicode.Range16{{0x340, 0x341, 0x1}, {0x343, 0x344, 0x1}, {0x374, 0x37e, 0xa}, {0x387, 0x958, 0x5d1}, {0x959, 0x95f, 0x1}, {0x9dc, 0x9dd, 0x1}, {0x9df, 0xa33, 0x54}, {0xa36, 0xa59, 0x23}, {0xa5a, 0xa5b, 0x1}, {0xa5e, 0xb5c, 0xfe}, {0xb5d, 0xf43, 0x3e6}, {0xf4d, 0xf5c, 0x5}, {0xf69, 0xf73, 0xa}, {0xf75, 0xf76, 0x1}, {0xf78, 0xf81, 0x9}, {0xf93, 0xf9d, 0xa}, {0xfa2, 0xfac, 0x5}, {0xfb9, 0x1f71, 0xfb8}, {0x1f73, 0x1f7d, 0x2}, {0x1fbb, 0x1fbe, 0x3}, {0x1fc9, 0x1fcb, 0x2}, {0x1fd3, 0x1feb, 0x8}, {0x1fee, 0x1fef, 0x1}, {0x1ff9, 0x1ffd, 0x2}, {0x2000, 0x2001, 0x1}, {0x2126, 0x212a, 0x4}, {0x212b, 0x2329, 0x1fe}, {0x232a, 0x2adc, 0x7b2}, {0xf900, 0xfa0d, 0x1}, {0xfa10, 0xfa12, 0x2}, {0xfa15, 0xfa1e, 0x1}, {0xfa20, 0xfa22, 0x2}, {0xfa25, 0xfa26, 0x1}, {0xfa2a, 0xfa6d, 0x1}, {0xfa70, 0xfad9, 0x1}, {0xfb1d, 0xfb1f, 0x2}, {0xfb2a, 0xfb36, 0x1}, {0xfb38, 0xfb3c, 0x1}, {0xfb3e, 0xfb40, 0x2}, {0xfb41, 0xfb43, 0x2}, {0xfb44, 0xfb46, 0x2}, {0xfb47, 0xfb4e, 0x1}}, []unicode.Range32{{0x1d15e, 0x1d164, 0x1}, {0x1d1bb, 0x1d1c0, 0x1}, {0x2f800, 0x2fa1d, 0x1}}, 0}
	nFCQuickCheckYes   = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x2ff, 0x1}, {0x305, 0x30d, 0x8}, {0x30e, 0x312, 0x2}, {0x315, 0x31a, 0x1}, {0x31c, 0x322, 0x1}, {0x329, 0x32c, 0x1}, {0x32f, 0x332, 0x3}, {0x333, 0x337, 0x1}, {0x339, 0x33f, 0x1}, {0x346, 0x373, 0x1}, {0x375, 0x37d, 0x1}, {0x37f, 0x386, 0x1}, {0x388, 0x652, 0x1}, {0x656, 0x93b, 0x1}, {0x93d, 0x957, 0x1}, {0x960, 0x9bd, 0x1}, {0x9bf, 0x9d6, 0x1}, {0x9d8, 0x9db, 0x1}, {0x9de, 0x9e0, 0x2}, {0x9e1, 0xa32, 0x1}, {0xa34, 0xa35, 0x1}, {0xa37, 0xa58, 0x1}, {0xa5c, 0xa5d, 0x1}, {0xa5f, 0xb3d, 0x1}, {0xb3f, 0xb55, 0x1}, {0xb58, 0xb5b, 0x1}, {0xb5e, 0xbbd, 0x1}, {0xbbf, 0xbd6, 0x1}, {0xbd8, 0xc55, 0x1}, {0xc57, 0xcc1, 0x1}, {0xcc3, 0xcd4, 0x1}, {0xcd7, 0xd3d, 0x1}, {0xd3f, 0xd56, 0x1}, {0xd58, 0xdc9, 0x1}, {0xdcb, 0xdce, 0x1}, {0xdd0, 0xdde, 0x1}, {0xde0, 0xf42, 0x1}, {0xf44, 0xf4c, 0x1}, {0xf4e, 0xf51, 0x1}, {0xf53, 0xf56, 0x1}, {0xf58, 0xf5b, 0x1}, {0xf5d, 0xf68, 0x1}, {0xf6a, 0xf72, 0x1}, {0xf74, 0xf77, 0x3}, {0xf79, 0xf80, 0x1}, {0xf82, 0xf92, 0x1}, {0xf94, 0xf9c, 0x1}, {0xf9e, 0xfa1, 0x1}, {0xfa3, 0xfa6, 0x1}, {0xfa8, 0xfab, 0x1}, {0xfad, 0xfb8, 0x1}, {0xfba, 0x102d, 0x1}, {0x102f, 0x1160, 0x1}, {0x1176, 0x11a7, 0x1}, {0x11c3, 0x1b34, 0x1}, {0x1b36, 0x1f70, 0x1}, {0x1f72, 0x1f7e, 0x2}, {0x1f7f, 0x1fba, 0x1}, {0x1fbc, 0x1fbd, 0x1}, {0x1fbf, 0x1fc8, 0x1}, {0x1fca, 0x1fcc, 0x2}, {0x1fcd, 0x1fd2, 0x1}, {0x1fd4, 0x1fda, 0x1}, {0x1fdc, 0x1fe2, 0x1}, {0x1fe4, 0x1fea, 0x1}, {0x1fec, 0x1fed, 0x1}, {0x1ff0, 0x1ff8, 0x1}, {0x1ffa, 0x1ffe, 0x2}, {0x1fff, 0x2002, 0x3}, {0x2003, 0x2125, 0x1}, {0x2127, 0x2129, 0x1}, {0x212c, 0x2328, 0x1}, {0x232b, 0x2adb, 0x1}, {0x2add, 0x3098, 0x1}, {0x309b, 0xf8ff, 0x1}, {0xfa0e, 0xfa0f, 0x1}, {0xfa11, 0xfa13, 0x2}, {0xfa14, 0xfa1f, 0xb}, {0xfa21, 0xfa23, 0x2}, {0xfa24, 0xfa27, 0x3}, {0xfa28, 0xfa29, 0x1}, {0xfa6e, 0xfa6f, 0x1}, {0xfada, 0xfb1c, 0x1}, {0xfb1e, 0xfb20, 0x2}, {0xfb21, 0xfb29, 0x1}, {0xfb37, 0xfb3d, 0x6}, {0xfb3f, 0xfb45, 0x3}, {0xfb4f, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x110b9, 0x1}, {0x110bb, 0x11126, 0x1}, {0x11128, 0x1133d, 0x1}, {0x1133f, 0x11356, 0x1}, {0x11358, 0x114af, 0x1}, {0x114b1, 0x114b9, 0x1}, {0x114bb, 0x114bc, 0x1}, {0x114be, 0x115ae, 0x1}, {0x115b0, 0x1d15d, 0x1}, {0x1d165, 0x1d1ba, 0x1}, {0x1d1c1, 0x2f7ff, 0x1}, {0x2fa1e, 0x10ffff, 0x1}}, 0}
)
