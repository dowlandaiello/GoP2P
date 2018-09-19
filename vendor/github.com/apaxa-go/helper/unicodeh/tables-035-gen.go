package unicodeh

import "unicode"

// Unicode property "Changes_When_NFKC_Casefolded" (known as "CWKCF", "Changes_When_NFKC_Casefolded").
// Kind of property: "Binary".
// Based on file "DerivedNormalizationProps.txt".
var (
	ChangesWhenNFKCCasefoldedNo  = changesWhenNFKCCasefoldedNo  // Value "No" (known as "N", "No", "F", "False").
	ChangesWhenNFKCCasefoldedYes = changesWhenNFKCCasefoldedYes // Value "Yes" (known as "Y", "Yes", "T", "True").
)

var (
	changesWhenNFKCCasefoldedNo  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x40, 0x1}, {0x5b, 0x9f, 0x1}, {0xa1, 0xa7, 0x1}, {0xa9, 0xab, 0x2}, {0xac, 0xb0, 0x2}, {0xb1, 0xb6, 0x5}, {0xb7, 0xbf, 0x4}, {0xd7, 0xe0, 0x9}, {0xe1, 0xff, 0x1}, {0x101, 0x131, 0x2}, {0x135, 0x137, 0x2}, {0x138, 0x13e, 0x2}, {0x142, 0x148, 0x2}, {0x14b, 0x177, 0x2}, {0x17a, 0x180, 0x2}, {0x183, 0x185, 0x2}, {0x188, 0x18c, 0x4}, {0x18d, 0x192, 0x5}, {0x195, 0x199, 0x4}, {0x19a, 0x19b, 0x1}, {0x19e, 0x1a1, 0x3}, {0x1a3, 0x1a5, 0x2}, {0x1a8, 0x1aa, 0x2}, {0x1ab, 0x1ad, 0x2}, {0x1b0, 0x1b4, 0x4}, {0x1b6, 0x1b9, 0x3}, {0x1ba, 0x1bb, 0x1}, {0x1bd, 0x1c3, 0x1}, {0x1ce, 0x1dc, 0x2}, {0x1dd, 0x1ef, 0x2}, {0x1f0, 0x1f5, 0x5}, {0x1f9, 0x233, 0x2}, {0x234, 0x239, 0x1}, {0x23c, 0x23f, 0x3}, {0x240, 0x242, 0x2}, {0x247, 0x24f, 0x2}, {0x250, 0x2af, 0x1}, {0x2b9, 0x2d7, 0x1}, {0x2de, 0x2df, 0x1}, {0x2e5, 0x33f, 0x1}, {0x342, 0x346, 0x4}, {0x347, 0x34e, 0x1}, {0x350, 0x36f, 0x1}, {0x371, 0x377, 0x2}, {0x378, 0x379, 0x1}, {0x37b, 0x37d, 0x1}, {0x380, 0x383, 0x1}, {0x38b, 0x38d, 0x2}, {0x390, 0x3a2, 0x12}, {0x3ac, 0x3c1, 0x1}, {0x3c3, 0x3ce, 0x1}, {0x3d7, 0x3ef, 0x2}, {0x3f3, 0x3f6, 0x3}, {0x3f8, 0x3fb, 0x3}, {0x3fc, 0x430, 0x34}, {0x431, 0x45f, 0x1}, {0x461, 0x481, 0x2}, {0x482, 0x489, 0x1}, {0x48b, 0x4bf, 0x2}, {0x4c2, 0x4ce, 0x2}, {0x4cf, 0x52f, 0x2}, {0x530, 0x557, 0x27}, {0x558, 0x586, 0x1}, {0x588, 0x61b, 0x1}, {0x61d, 0x674, 0x1}, {0x679, 0x957, 0x1}, {0x960, 0x9db, 0x1}, {0x9de, 0x9e0, 0x2}, {0x9e1, 0xa32, 0x1}, {0xa34, 0xa35, 0x1}, {0xa37, 0xa58, 0x1}, {0xa5c, 0xa5d, 0x1}, {0xa5f, 0xb5b, 0x1}, {0xb5e, 0xe32, 0x1}, {0xe34, 0xeb2, 0x1}, {0xeb4, 0xedb, 0x1}, {0xede, 0xf0b, 0x1}, {0xf0d, 0xf42, 0x1}, {0xf44, 0xf4c, 0x1}, {0xf4e, 0xf51, 0x1}, {0xf53, 0xf56, 0x1}, {0xf58, 0xf5b, 0x1}, {0xf5d, 0xf68, 0x1}, {0xf6a, 0xf72, 0x1}, {0xf74, 0xf7a, 0x6}, {0xf7b, 0xf80, 0x1}, {0xf82, 0xf92, 0x1}, {0xf94, 0xf9c, 0x1}, {0xf9e, 0xfa1, 0x1}, {0xfa3, 0xfa6, 0x1}, {0xfa8, 0xfab, 0x1}, {0xfad, 0xfb8, 0x1}, {0xfba, 0x109f, 0x1}, {0x10c6, 0x10c8, 0x2}, {0x10c9, 0x10cc, 0x1}, {0x10ce, 0x10fb, 0x1}, {0x10fd, 0x115e, 0x1}, {0x1161, 0x13f7, 0x1}, {0x13fe, 0x17b3, 0x1}, {0x17b6, 0x180a, 0x1}, {0x180f, 0x1c7f, 0x1}, {0x1c89, 0x1d2b, 0x1}, {0x1d2f, 0x1d3b, 0xc}, {0x1d4e, 0x1d6b, 0x1d}, {0x1d6c, 0x1d77, 0x1}, {0x1d79, 0x1d9a, 0x1}, {0x1dc0, 0x1dff, 0x1}, {0x1e01, 0x1e95, 0x2}, {0x1e96, 0x1e99, 0x1}, {0x1e9c, 0x1e9d, 0x1}, {0x1e9f, 0x1eff, 0x2}, {0x1f00, 0x1f07, 0x1}, {0x1f10, 0x1f17, 0x1}, {0x1f1e, 0x1f27, 0x1}, {0x1f30, 0x1f37, 0x1}, {0x1f40, 0x1f47, 0x1}, {0x1f4e, 0x1f58, 0x1}, {0x1f5a, 0x1f60, 0x2}, {0x1f61, 0x1f67, 0x1}, {0x1f70, 0x1f7e, 0x2}, {0x1f7f, 0x1fb0, 0x31}, {0x1fb1, 0x1fb5, 0x4}, {0x1fb6, 0x1fc5, 0xf}, {0x1fc6, 0x1fd0, 0xa}, {0x1fd1, 0x1fd2, 0x1}, {0x1fd4, 0x1fd7, 0x1}, {0x1fdc, 0x1fe0, 0x4}, {0x1fe1, 0x1fe2, 0x1}, {0x1fe4, 0x1fe7, 0x1}, {0x1ff0, 0x1ff1, 0x1}, {0x1ff5, 0x1ff6, 0x1}, {0x1fff, 0x2010, 0x11}, {0x2012, 0x2016, 0x1}, {0x2018, 0x2023, 0x1}, {0x2027, 0x2029, 0x1}, {0x2030, 0x2032, 0x1}, {0x2035, 0x2038, 0x3}, {0x2039, 0x203b, 0x1}, {0x203d, 0x203f, 0x2}, {0x2040, 0x2046, 0x1}, {0x204a, 0x2056, 0x1}, {0x2058, 0x205e, 0x1}, {0x2072, 0x2073, 0x1}, {0x208f, 0x209d, 0xe}, {0x209e, 0x20a7, 0x1}, {0x20a9, 0x20ff, 0x1}, {0x2104, 0x2108, 0x4}, {0x2114, 0x2117, 0x3}, {0x2118, 0x211e, 0x6}, {0x211f, 0x2123, 0x4}, {0x2125, 0x2129, 0x2}, {0x212e, 0x213a, 0xc}, {0x2141, 0x2144, 0x1}, {0x214a, 0x214f, 0x1}, {0x2180, 0x2182, 0x1}, {0x2184, 0x2188, 0x1}, {0x218a, 0x222b, 0x1}, {0x222e, 0x2231, 0x3}, {0x2232, 0x2328, 0x1}, {0x232b, 0x245f, 0x1}, {0x24eb, 0x2a0b, 0x1}, {0x2a0d, 0x2a73, 0x1}, {0x2a77, 0x2adb, 0x1}, {0x2add, 0x2bff, 0x1}, {0x2c2f, 0x2c5f, 0x1}, {0x2c61, 0x2c65, 0x4}, {0x2c66, 0x2c6c, 0x2}, {0x2c71, 0x2c73, 0x2}, {0x2c74, 0x2c76, 0x2}, {0x2c77, 0x2c7b, 0x1}, {0x2c81, 0x2ce3, 0x2}, {0x2ce4, 0x2cea, 0x1}, {0x2cec, 0x2cee, 0x2}, {0x2cef, 0x2cf1, 0x1}, {0x2cf3, 0x2d6e, 0x1}, {0x2d70, 0x2e9e, 0x1}, {0x2ea0, 0x2ef2, 0x1}, {0x2ef4, 0x2eff, 0x1}, {0x2fd6, 0x2fff, 0x1}, {0x3001, 0x3035, 0x1}, {0x3037, 0x303b, 0x4}, {0x303c, 0x309a, 0x1}, {0x309d, 0x309e, 0x1}, {0x30a0, 0x30fe, 0x1}, {0x3100, 0x3130, 0x1}, {0x318f, 0x3191, 0x1}, {0x31a0, 0x31ff, 0x1}, {0x321f, 0x3248, 0x29}, {0x3249, 0x324f, 0x1}, {0x327f, 0x32ff, 0x80}, {0x3400, 0xa63f, 0x1}, {0xa641, 0xa66d, 0x2}, {0xa66e, 0xa67f, 0x1}, {0xa681, 0xa69b, 0x2}, {0xa69e, 0xa721, 0x1}, {0xa723, 0xa72f, 0x2}, {0xa730, 0xa731, 0x1}, {0xa733, 0xa771, 0x2}, {0xa772, 0xa778, 0x1}, {0xa77a, 0xa77c, 0x2}, {0xa77f, 0xa787, 0x2}, {0xa788, 0xa78a, 0x1}, {0xa78c, 0xa78e, 0x2}, {0xa78f, 0xa793, 0x2}, {0xa794, 0xa795, 0x1}, {0xa797, 0xa7a9, 0x2}, {0xa7af, 0xa7b5, 0x6}, {0xa7b7, 0xa7f7, 0x1}, {0xa7fa, 0xab5b, 0x1}, {0xab60, 0xab6f, 0x1}, {0xabc0, 0xf8ff, 0x1}, {0xfa0e, 0xfa0f, 0x1}, {0xfa11, 0xfa13, 0x2}, {0xfa14, 0xfa1f, 0xb}, {0xfa21, 0xfa23, 0x2}, {0xfa24, 0xfa27, 0x3}, {0xfa28, 0xfa29, 0x1}, {0xfa6e, 0xfa6f, 0x1}, {0xfada, 0xfaff, 0x1}, {0xfb07, 0xfb12, 0x1}, {0xfb18, 0xfb1c, 0x1}, {0xfb1e, 0xfb37, 0x19}, {0xfb3d, 0xfb3f, 0x2}, {0xfb42, 0xfb45, 0x3}, {0xfbb2, 0xfbd2, 0x1}, {0xfd3e, 0xfd4f, 0x1}, {0xfd90, 0xfd91, 0x1}, {0xfdc8, 0xfdef, 0x1}, {0xfdfd, 0xfdff, 0x1}, {0xfe1a, 0xfe2f, 0x1}, {0xfe45, 0xfe46, 0x1}, {0xfe53, 0xfe67, 0x14}, {0xfe6c, 0xfe6f, 0x1}, {0xfe73, 0xfe75, 0x2}, {0xfefd, 0xfefe, 0x1}, {0xff00, 0xffbf, 0xbf}, {0xffc0, 0xffc1, 0x1}, {0xffc8, 0xffc9, 0x1}, {0xffd0, 0xffd1, 0x1}, {0xffd8, 0xffd9, 0x1}, {0xffdd, 0xffdf, 0x1}, {0xffe7, 0xffef, 0x8}, {0xfff9, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x103ff, 0x1}, {0x10428, 0x104af, 0x1}, {0x104d4, 0x10c7f, 0x1}, {0x10cb3, 0x1189f, 0x1}, {0x118c0, 0x1bc9f, 0x1}, {0x1bca4, 0x1d15d, 0x1}, {0x1d165, 0x1d172, 0x1}, {0x1d17b, 0x1d1ba, 0x1}, {0x1d1c1, 0x1d3ff, 0x1}, {0x1d455, 0x1d49d, 0x48}, {0x1d4a0, 0x1d4a1, 0x1}, {0x1d4a3, 0x1d4a4, 0x1}, {0x1d4a7, 0x1d4a8, 0x1}, {0x1d4ad, 0x1d4ba, 0xd}, {0x1d4bc, 0x1d4c4, 0x8}, {0x1d506, 0x1d50b, 0x5}, {0x1d50c, 0x1d515, 0x9}, {0x1d51d, 0x1d53a, 0x1d}, {0x1d53f, 0x1d545, 0x6}, {0x1d547, 0x1d549, 0x1}, {0x1d551, 0x1d6a6, 0x155}, {0x1d6a7, 0x1d7cc, 0x125}, {0x1d7cd, 0x1d800, 0x33}, {0x1d801, 0x1e8ff, 0x1}, {0x1e922, 0x1edff, 0x1}, {0x1ee04, 0x1ee20, 0x1c}, {0x1ee23, 0x1ee25, 0x2}, {0x1ee26, 0x1ee28, 0x2}, {0x1ee33, 0x1ee38, 0x5}, {0x1ee3a, 0x1ee3c, 0x2}, {0x1ee3d, 0x1ee41, 0x1}, {0x1ee43, 0x1ee46, 0x1}, {0x1ee48, 0x1ee4c, 0x2}, {0x1ee50, 0x1ee53, 0x3}, {0x1ee55, 0x1ee56, 0x1}, {0x1ee58, 0x1ee60, 0x2}, {0x1ee63, 0x1ee65, 0x2}, {0x1ee66, 0x1ee6b, 0x5}, {0x1ee73, 0x1ee7d, 0x5}, {0x1ee7f, 0x1ee8a, 0xb}, {0x1ee9c, 0x1eea0, 0x1}, {0x1eea4, 0x1eeaa, 0x6}, {0x1eebc, 0x1f0ff, 0x1}, {0x1f10b, 0x1f10f, 0x1}, {0x1f12f, 0x1f150, 0x21}, {0x1f151, 0x1f169, 0x1}, {0x1f16c, 0x1f18f, 0x1}, {0x1f191, 0x1f1ff, 0x1}, {0x1f203, 0x1f20f, 0x1}, {0x1f23c, 0x1f23f, 0x1}, {0x1f249, 0x1f24f, 0x1}, {0x1f252, 0x2f7ff, 0x1}, {0x2fa1e, 0xdffff, 0x1}, {0xe1000, 0x10ffff, 0x1}}, 9}
	changesWhenNFKCCasefoldedYes = &unicode.RangeTable{[]unicode.Range16{{0x41, 0x5a, 0x1}, {0xa0, 0xa8, 0x8}, {0xaa, 0xad, 0x3}, {0xaf, 0xb2, 0x3}, {0xb3, 0xb5, 0x1}, {0xb8, 0xba, 0x1}, {0xbc, 0xbe, 0x1}, {0xc0, 0xd6, 0x1}, {0xd8, 0xdf, 0x1}, {0x100, 0x132, 0x2}, {0x133, 0x134, 0x1}, {0x136, 0x139, 0x3}, {0x13b, 0x13f, 0x2}, {0x140, 0x141, 0x1}, {0x143, 0x149, 0x2}, {0x14a, 0x178, 0x2}, {0x179, 0x181, 0x2}, {0x182, 0x186, 0x2}, {0x187, 0x189, 0x2}, {0x18a, 0x18b, 0x1}, {0x18e, 0x191, 0x1}, {0x193, 0x194, 0x1}, {0x196, 0x198, 0x1}, {0x19c, 0x19d, 0x1}, {0x19f, 0x1a0, 0x1}, {0x1a2, 0x1a6, 0x2}, {0x1a7, 0x1a9, 0x2}, {0x1ac, 0x1ae, 0x2}, {0x1af, 0x1b1, 0x2}, {0x1b2, 0x1b3, 0x1}, {0x1b5, 0x1b7, 0x2}, {0x1b8, 0x1bc, 0x4}, {0x1c4, 0x1cd, 0x1}, {0x1cf, 0x1db, 0x2}, {0x1de, 0x1ee, 0x2}, {0x1f1, 0x1f4, 0x1}, {0x1f6, 0x1f8, 0x1}, {0x1fa, 0x232, 0x2}, {0x23a, 0x23b, 0x1}, {0x23d, 0x23e, 0x1}, {0x241, 0x243, 0x2}, {0x244, 0x246, 0x1}, {0x248, 0x24e, 0x2}, {0x2b0, 0x2b8, 0x1}, {0x2d8, 0x2dd, 0x1}, {0x2e0, 0x2e4, 0x1}, {0x340, 0x341, 0x1}, {0x343, 0x345, 0x1}, {0x34f, 0x370, 0x21}, {0x372, 0x376, 0x2}, {0x37a, 0x37e, 0x4}, {0x37f, 0x384, 0x5}, {0x385, 0x38a, 0x1}, {0x38c, 0x38e, 0x2}, {0x38f, 0x391, 0x2}, {0x392, 0x3a1, 0x1}, {0x3a3, 0x3ab, 0x1}, {0x3c2, 0x3cf, 0xd}, {0x3d0, 0x3d6, 0x1}, {0x3d8, 0x3f0, 0x2}, {0x3f1, 0x3f2, 0x1}, {0x3f4, 0x3f5, 0x1}, {0x3f7, 0x3f9, 0x2}, {0x3fa, 0x3fd, 0x3}, {0x3fe, 0x42f, 0x1}, {0x460, 0x480, 0x2}, {0x48a, 0x4c0, 0x2}, {0x4c1, 0x4cd, 0x2}, {0x4d0, 0x52e, 0x2}, {0x531, 0x556, 0x1}, {0x587, 0x61c, 0x95}, {0x675, 0x678, 0x1}, {0x958, 0x95f, 0x1}, {0x9dc, 0x9dd, 0x1}, {0x9df, 0xa33, 0x54}, {0xa36, 0xa59, 0x23}, {0xa5a, 0xa5b, 0x1}, {0xa5e, 0xb5c, 0xfe}, {0xb5d, 0xe33, 0x2d6}, {0xeb3, 0xedc, 0x29}, {0xedd, 0xf0c, 0x2f}, {0xf43, 0xf4d, 0xa}, {0xf52, 0xf5c, 0x5}, {0xf69, 0xf73, 0xa}, {0xf75, 0xf79, 0x1}, {0xf81, 0xf93, 0x12}, {0xf9d, 0xfac, 0x5}, {0xfb9, 0x10a0, 0xe7}, {0x10a1, 0x10c5, 0x1}, {0x10c7, 0x10cd, 0x6}, {0x10fc, 0x115f, 0x63}, {0x1160, 0x13f8, 0x298}, {0x13f9, 0x13fd, 0x1}, {0x17b4, 0x17b5, 0x1}, {0x180b, 0x180e, 0x1}, {0x1c80, 0x1c88, 0x1}, {0x1d2c, 0x1d2e, 0x1}, {0x1d30, 0x1d3a, 0x1}, {0x1d3c, 0x1d4d, 0x1}, {0x1d4f, 0x1d6a, 0x1}, {0x1d78, 0x1d9b, 0x23}, {0x1d9c, 0x1dbf, 0x1}, {0x1e00, 0x1e94, 0x2}, {0x1e9a, 0x1e9b, 0x1}, {0x1e9e, 0x1efe, 0x2}, {0x1f08, 0x1f0f, 0x1}, {0x1f18, 0x1f1d, 0x1}, {0x1f28, 0x1f2f, 0x1}, {0x1f38, 0x1f3f, 0x1}, {0x1f48, 0x1f4d, 0x1}, {0x1f59, 0x1f5f, 0x2}, {0x1f68, 0x1f6f, 0x1}, {0x1f71, 0x1f7d, 0x2}, {0x1f80, 0x1faf, 0x1}, {0x1fb2, 0x1fb4, 0x1}, {0x1fb7, 0x1fc4, 0x1}, {0x1fc7, 0x1fcf, 0x1}, {0x1fd3, 0x1fd8, 0x5}, {0x1fd9, 0x1fdb, 0x1}, {0x1fdd, 0x1fdf, 0x1}, {0x1fe3, 0x1fe8, 0x5}, {0x1fe9, 0x1fef, 0x1}, {0x1ff2, 0x1ff4, 0x1}, {0x1ff7, 0x1ffe, 0x1}, {0x2000, 0x200f, 0x1}, {0x2011, 0x2017, 0x6}, {0x2024, 0x2026, 0x1}, {0x202a, 0x202f, 0x1}, {0x2033, 0x2034, 0x1}, {0x2036, 0x2037, 0x1}, {0x203c, 0x203e, 0x2}, {0x2047, 0x2049, 0x1}, {0x2057, 0x205f, 0x8}, {0x2060, 0x2071, 0x1}, {0x2074, 0x208e, 0x1}, {0x2090, 0x209c, 0x1}, {0x20a8, 0x2100, 0x58}, {0x2101, 0x2103, 0x1}, {0x2105, 0x2107, 0x1}, {0x2109, 0x2113, 0x1}, {0x2115, 0x2116, 0x1}, {0x2119, 0x211d, 0x1}, {0x2120, 0x2122, 0x1}, {0x2124, 0x212a, 0x2}, {0x212b, 0x212d, 0x1}, {0x212f, 0x2139, 0x1}, {0x213b, 0x2140, 0x1}, {0x2145, 0x2149, 0x1}, {0x2150, 0x217f, 0x1}, {0x2183, 0x2189, 0x6}, {0x222c, 0x222d, 0x1}, {0x222f, 0x2230, 0x1}, {0x2329, 0x232a, 0x1}, {0x2460, 0x24ea, 0x1}, {0x2a0c, 0x2a74, 0x68}, {0x2a75, 0x2a76, 0x1}, {0x2adc, 0x2c00, 0x124}, {0x2c01, 0x2c2e, 0x1}, {0x2c60, 0x2c62, 0x2}, {0x2c63, 0x2c64, 0x1}, {0x2c67, 0x2c6d, 0x2}, {0x2c6e, 0x2c70, 0x1}, {0x2c72, 0x2c75, 0x3}, {0x2c7c, 0x2c80, 0x1}, {0x2c82, 0x2ce2, 0x2}, {0x2ceb, 0x2ced, 0x2}, {0x2cf2, 0x2d6f, 0x7d}, {0x2e9f, 0x2ef3, 0x54}, {0x2f00, 0x2fd5, 0x1}, {0x3000, 0x3036, 0x36}, {0x3038, 0x303a, 0x1}, {0x309b, 0x309c, 0x1}, {0x309f, 0x30ff, 0x60}, {0x3131, 0x318e, 0x1}, {0x3192, 0x319f, 0x1}, {0x3200, 0x321e, 0x1}, {0x3220, 0x3247, 0x1}, {0x3250, 0x327e, 0x1}, {0x3280, 0x32fe, 0x1}, {0x3300, 0x33ff, 0x1}, {0xa640, 0xa66c, 0x2}, {0xa680, 0xa69c, 0x2}, {0xa69d, 0xa722, 0x85}, {0xa724, 0xa72e, 0x2}, {0xa732, 0xa770, 0x2}, {0xa779, 0xa77d, 0x2}, {0xa77e, 0xa786, 0x2}, {0xa78b, 0xa78d, 0x2}, {0xa790, 0xa792, 0x2}, {0xa796, 0xa7aa, 0x2}, {0xa7ab, 0xa7ae, 0x1}, {0xa7b0, 0xa7b4, 0x1}, {0xa7b6, 0xa7f8, 0x42}, {0xa7f9, 0xab5c, 0x363}, {0xab5d, 0xab5f, 0x1}, {0xab70, 0xabbf, 0x1}, {0xf900, 0xfa0d, 0x1}, {0xfa10, 0xfa12, 0x2}, {0xfa15, 0xfa1e, 0x1}, {0xfa20, 0xfa22, 0x2}, {0xfa25, 0xfa26, 0x1}, {0xfa2a, 0xfa6d, 0x1}, {0xfa70, 0xfad9, 0x1}, {0xfb00, 0xfb06, 0x1}, {0xfb13, 0xfb17, 0x1}, {0xfb1d, 0xfb1f, 0x2}, {0xfb20, 0xfb36, 0x1}, {0xfb38, 0xfb3c, 0x1}, {0xfb3e, 0xfb40, 0x2}, {0xfb41, 0xfb43, 0x2}, {0xfb44, 0xfb46, 0x2}, {0xfb47, 0xfbb1, 0x1}, {0xfbd3, 0xfd3d, 0x1}, {0xfd50, 0xfd8f, 0x1}, {0xfd92, 0xfdc7, 0x1}, {0xfdf0, 0xfdfc, 0x1}, {0xfe00, 0xfe19, 0x1}, {0xfe30, 0xfe44, 0x1}, {0xfe47, 0xfe52, 0x1}, {0xfe54, 0xfe66, 0x1}, {0xfe68, 0xfe6b, 0x1}, {0xfe70, 0xfe72, 0x1}, {0xfe74, 0xfe76, 0x2}, {0xfe77, 0xfefc, 0x1}, {0xfeff, 0xff01, 0x2}, {0xff02, 0xffbe, 0x1}, {0xffc2, 0xffc7, 0x1}, {0xffca, 0xffcf, 0x1}, {0xffd2, 0xffd7, 0x1}, {0xffda, 0xffdc, 0x1}, {0xffe0, 0xffe6, 0x1}, {0xffe8, 0xffee, 0x1}, {0xfff0, 0xfff8, 0x1}}, []unicode.Range32{{0x10400, 0x10427, 0x1}, {0x104b0, 0x104d3, 0x1}, {0x10c80, 0x10cb2, 0x1}, {0x118a0, 0x118bf, 0x1}, {0x1bca0, 0x1bca3, 0x1}, {0x1d15e, 0x1d164, 0x1}, {0x1d173, 0x1d17a, 0x1}, {0x1d1bb, 0x1d1c0, 0x1}, {0x1d400, 0x1d454, 0x1}, {0x1d456, 0x1d49c, 0x1}, {0x1d49e, 0x1d49f, 0x1}, {0x1d4a2, 0x1d4a5, 0x3}, {0x1d4a6, 0x1d4a9, 0x3}, {0x1d4aa, 0x1d4ac, 0x1}, {0x1d4ae, 0x1d4b9, 0x1}, {0x1d4bb, 0x1d4bd, 0x2}, {0x1d4be, 0x1d4c3, 0x1}, {0x1d4c5, 0x1d505, 0x1}, {0x1d507, 0x1d50a, 0x1}, {0x1d50d, 0x1d514, 0x1}, {0x1d516, 0x1d51c, 0x1}, {0x1d51e, 0x1d539, 0x1}, {0x1d53b, 0x1d53e, 0x1}, {0x1d540, 0x1d544, 0x1}, {0x1d546, 0x1d54a, 0x4}, {0x1d54b, 0x1d550, 0x1}, {0x1d552, 0x1d6a5, 0x1}, {0x1d6a8, 0x1d7cb, 0x1}, {0x1d7ce, 0x1d7ff, 0x1}, {0x1e900, 0x1e921, 0x1}, {0x1ee00, 0x1ee03, 0x1}, {0x1ee05, 0x1ee1f, 0x1}, {0x1ee21, 0x1ee22, 0x1}, {0x1ee24, 0x1ee27, 0x3}, {0x1ee29, 0x1ee32, 0x1}, {0x1ee34, 0x1ee37, 0x1}, {0x1ee39, 0x1ee3b, 0x2}, {0x1ee42, 0x1ee47, 0x5}, {0x1ee49, 0x1ee4d, 0x2}, {0x1ee4e, 0x1ee4f, 0x1}, {0x1ee51, 0x1ee52, 0x1}, {0x1ee54, 0x1ee57, 0x3}, {0x1ee59, 0x1ee61, 0x2}, {0x1ee62, 0x1ee64, 0x2}, {0x1ee67, 0x1ee6a, 0x1}, {0x1ee6c, 0x1ee72, 0x1}, {0x1ee74, 0x1ee77, 0x1}, {0x1ee79, 0x1ee7c, 0x1}, {0x1ee7e, 0x1ee80, 0x2}, {0x1ee81, 0x1ee89, 0x1}, {0x1ee8b, 0x1ee9b, 0x1}, {0x1eea1, 0x1eea3, 0x1}, {0x1eea5, 0x1eea9, 0x1}, {0x1eeab, 0x1eebb, 0x1}, {0x1f100, 0x1f10a, 0x1}, {0x1f110, 0x1f12e, 0x1}, {0x1f130, 0x1f14f, 0x1}, {0x1f16a, 0x1f16b, 0x1}, {0x1f190, 0x1f200, 0x70}, {0x1f201, 0x1f202, 0x1}, {0x1f210, 0x1f23b, 0x1}, {0x1f240, 0x1f248, 0x1}, {0x1f250, 0x1f251, 0x1}, {0x2f800, 0x2fa1d, 0x1}, {0xe0000, 0xe0fff, 0x1}}, 9}
)
