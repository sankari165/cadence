// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/shared/v1/error.proto

package sharedv1

var yarpcFileDescriptorClosure3688ca0fd170c8f9 = [][]byte{
	// uber/cadence/shared/v1/error.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0x6f, 0xd3, 0x30,
		0x14, 0xc5, 0xd5, 0xf2, 0x47, 0xab, 0x3b, 0x01, 0x8b, 0xa6, 0x51, 0x3a, 0x09, 0xaa, 0x80, 0xd0,
		0xe0, 0x21, 0xa1, 0x43, 0x88, 0x07, 0x9e, 0xe8, 0x28, 0x5a, 0xd1, 0x06, 0x52, 0x3a, 0x86, 0xc4,
		0x4b, 0xe4, 0xda, 0x97, 0xc6, 0x6a, 0x73, 0x5d, 0xd9, 0xb7, 0x2d, 0xf9, 0x16, 0x3c, 0xf3, 0x69,
		0x91, 0xed, 0x94, 0xad, 0x7b, 0xdb, 0x5b, 0x72, 0xef, 0xef, 0x9c, 0x9c, 0x1c, 0xd9, 0x2c, 0x5e,
		0x4e, 0xc0, 0xa4, 0x82, 0x4b, 0x40, 0x01, 0xa9, 0x2d, 0xb8, 0x01, 0x99, 0xae, 0xfa, 0x29, 0x18,
		0xa3, 0x4d, 0xb2, 0x30, 0x9a, 0x74, 0x74, 0xe0, 0x98, 0xa4, 0x66, 0x92, 0xc0, 0x24, 0xab, 0x7e,
		0xb7, 0xb7, 0xa5, 0xe5, 0x0b, 0xe5, 0x84, 0x42, 0x97, 0xa5, 0xc6, 0xa0, 0xec, 0x3e, 0xdf, 0x26,
		0x64, 0xa9, 0xd0, 0x31, 0x85, 0xb2, 0xa4, 0x4d, 0x15, 0xa0, 0xf8, 0x9c, 0x3d, 0x39, 0x59, 0x1a,
		0x03, 0x48, 0x03, 0xc3, 0x51, 0x14, 0x27, 0x05, 0xc7, 0x29, 0xc8, 0xa1, 0x4b, 0x10, 0xbd, 0x61,
		0xfb, 0x22, 0x2c, 0xf3, 0x89, 0xdf, 0xe6, 0xa4, 0x67, 0x80, 0x9d, 0x46, 0xaf, 0x71, 0xb4, 0x9b,
		0x45, 0xe2, 0xba, 0xf0, 0xc2, 0x6d, 0xe2, 0x1e, 0x7b, 0x3a, 0x42, 0x02, 0x83, 0x7c, 0xfe, 0x89,
		0x13, 0x1f, 0xa1, 0xd0, 0x68, 0x95, 0x25, 0x40, 0x51, 0x79, 0xcf, 0xb8, 0xcb, 0x3a, 0xc3, 0x15,
		0x20, 0x7d, 0x9c, 0x1b, 0xe0, 0xb2, 0x1a, 0x13, 0x37, 0x54, 0x7f, 0x2f, 0xee, 0xb0, 0x83, 0x0c,
		0x4a, 0x4d, 0x30, 0xae, 0x50, 0x9c, 0x73, 0x12, 0xc5, 0x66, 0xf3, 0xb7, 0xc9, 0x1e, 0x65, 0x40,
		0xa6, 0xba, 0xe0, 0x76, 0x76, 0x79, 0x1c, 0xe2, 0x1d, 0xb2, 0x96, 0xd4, 0x25, 0x57, 0x98, 0x2b,
		0xe9, 0x33, 0xb5, 0xb2, 0x9d, 0x30, 0x18, 0xc9, 0xe8, 0x3b, 0x8b, 0xd6, 0xda, 0xcc, 0x7e, 0xcd,
		0xf5, 0x3a, 0x87, 0xdf, 0x20, 0x96, 0xa4, 0x34, 0x76, 0x9a, 0xbd, 0xc6, 0x51, 0xfb, 0xf8, 0x65,
		0xb2, 0x55, 0x2a, 0x5f, 0xa8, 0x64, 0xd5, 0x4f, 0x7e, 0xd4, 0xf8, 0x70, 0x43, 0x67, 0x7b, 0xeb,
		0x9b, 0xa3, 0xe8, 0x0b, 0x6b, 0x5b, 0x17, 0x39, 0x07, 0xf7, 0x13, 0x9d, 0x3b, 0xde, 0xef, 0xd5,
		0x0d, 0x3f, 0x57, 0xb5, 0x73, 0xbc, 0x04, 0x63, 0x95, 0xc6, 0xd3, 0xd0, 0xf8, 0x88, 0xa0, 0xcc,
		0x98, 0x57, 0xfb, 0x06, 0xa2, 0xcf, 0xac, 0x05, 0x28, 0x6b, 0xa7, 0xbb, 0xb7, 0x75, 0xda, 0x01,
		0x94, 0xde, 0x27, 0x4e, 0xd9, 0xe3, 0x71, 0xc1, 0x8d, 0xfc, 0xb6, 0x46, 0x30, 0xb6, 0x50, 0x8b,
		0x33, 0x6d, 0x29, 0x54, 0xb4, 0xcf, 0xee, 0x69, 0x37, 0xad, 0xeb, 0x09, 0x2f, 0xf1, 0x9f, 0x06,
		0x3b, 0x74, 0x45, 0x9e, 0x29, 0x4b, 0x5f, 0x35, 0x39, 0x9d, 0x1c, 0x54, 0xa7, 0xff, 0x55, 0xaf,
		0xd9, 0x9e, 0x03, 0x65, 0x3e, 0xa9, 0x72, 0x25, 0x01, 0x49, 0x51, 0x55, 0x3b, 0x3c, 0xd4, 0x01,
		0x1e, 0xd5, 0xe3, 0xe8, 0x19, 0x6b, 0x97, 0xd7, 0xa8, 0xa6, 0xa7, 0x58, 0x79, 0x05, 0xbc, 0x60,
		0x0f, 0x88, 0xdb, 0x59, 0x3e, 0x57, 0x96, 0x72, 0xe4, 0x25, 0xf8, 0xd2, 0x5a, 0xd9, 0x2e, 0x6d,
		0x12, 0xf0, 0x12, 0x06, 0xef, 0x7f, 0xbe, 0x9b, 0x2a, 0x2a, 0x96, 0x93, 0x44, 0xe8, 0x32, 0xdd,
		0x3a, 0xb9, 0xc9, 0x14, 0x30, 0xf5, 0xa7, 0xf5, 0xea, 0x8a, 0x7c, 0x08, 0x4f, 0xab, 0xfe, 0xe4,
		0xbe, 0xdf, 0xbc, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x28, 0x15, 0x25, 0x92, 0x4c, 0x03, 0x00,
		0x00,
	},
	// uber/cadence/api/v1/common.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x51, 0x73, 0xdb, 0xc4,
		0x13, 0xff, 0x2b, 0x8e, 0x9d, 0x64, 0xed, 0x26, 0xfe, 0x5f, 0x48, 0xe2, 0xa4, 0x04, 0x52, 0xcd,
		0x30, 0x0d, 0x1d, 0x90, 0x27, 0xee, 0x4b, 0x87, 0x4e, 0x01, 0x27, 0x76, 0x12, 0xb5, 0xc1, 0x36,
		0xb2, 0x69, 0xa6, 0x30, 0x83, 0xe6, 0x2c, 0x9d, 0xdc, 0xc3, 0xd2, 0x9d, 0x38, 0x9d, 0x9c, 0xf8,
		0x85, 0xe1, 0x93, 0xf0, 0xc0, 0xd7, 0xe1, 0x91, 0x2f, 0xc4, 0x48, 0x3a, 0xc5, 0x76, 0x71, 0xa6,
		0x3c, 0x30, 0xbc, 0xdd, 0xed, 0xef, 0xb7, 0xbb, 0xbf, 0x3b, 0xed, 0xae, 0x0e, 0x8e, 0xe2, 0x21,
		0x11, 0x75, 0x07, 0xbb, 0x84, 0x39, 0xa4, 0x8e, 0x43, 0x5a, 0x9f, 0x9c, 0xd4, 0x1d, 0x1e, 0x04,
		0x9c, 0x19, 0xa1, 0xe0, 0x92, 0xa3, 0xed, 0x84, 0x61, 0x28, 0x86, 0x81, 0x43, 0x6a, 0x4c, 0x4e,
		0x0e, 0x3e, 0x1a, 0x71, 0x3e, 0xf2, 0x49, 0x3d, 0xa5, 0x0c, 0x63, 0xaf, 0xee, 0xc6, 0x02, 0x4b,
		0x9a, 0x3b, 0xe9, 0xaf, 0xe0, 0xff, 0xd7, 0x5c, 0x8c, 0x3d, 0x9f, 0xdf, 0xb4, 0x6f, 0x89, 0x13,
		0x27, 0x10, 0xfa, 0x18, 0xca, 0x37, 0xca, 0x68, 0x53, 0xb7, 0xa6, 0x1d, 0x69, 0xc7, 0x1b, 0x16,
		0xe4, 0x26, 0xd3, 0x45, 0x3b, 0x50, 0x12, 0x31, 0x4b, 0xb0, 0x95, 0x14, 0x2b, 0x8a, 0x98, 0x99,
		0xae, 0xae, 0x43, 0x25, 0x0f, 0x36, 0x98, 0x86, 0x04, 0x21, 0x58, 0x65, 0x38, 0x20, 0x2a, 0x40,
		0xba, 0x4e, 0x38, 0x4d, 0x47, 0xd2, 0x09, 0x95, 0xd3, 0x7b, 0x39, 0x87, 0xb0, 0xd6, 0xc3, 0x53,
		0x9f, 0x63, 0x37, 0x81, 0x5d, 0x2c, 0x71, 0x0a, 0x57, 0xac, 0x74, 0xad, 0x3f, 0x87, 0xb5, 0x73,
		0x4c, 0xfd, 0x58, 0x10, 0xb4, 0x0b, 0x25, 0x41, 0x70, 0xc4, 0x99, 0xf2, 0x57, 0x3b, 0x54, 0x83,
		0x35, 0x97, 0x48, 0x4c, 0xfd, 0x28, 0x55, 0x58, 0xb1, 0xf2, 0xad, 0xfe, 0x9b, 0x06, 0xab, 0xdf,
		0x90, 0x80, 0xa3, 0x17, 0x50, 0xf2, 0x28, 0xf1, 0xdd, 0xa8, 0xa6, 0x1d, 0x15, 0x8e, 0xcb, 0x8d,
		0x4f, 0x8c, 0x25, 0xf7, 0x67, 0x24, 0x54, 0xe3, 0x3c, 0xe5, 0xb5, 0x99, 0x14, 0x53, 0x4b, 0x39,
		0x1d, 0x5c, 0x43, 0x79, 0xce, 0x8c, 0xaa, 0x50, 0x18, 0x93, 0xa9, 0x52, 0x91, 0x2c, 0x51, 0x03,
		0x8a, 0x13, 0xec, 0xc7, 0x24, 0x15, 0x50, 0x6e, 0x7c, 0xb8, 0x34, 0xbc, 0x3a, 0xa6, 0x95, 0x51,
		0xbf, 0x58, 0x79, 0xa6, 0xe9, 0xbf, 0x6b, 0x50, 0xba, 0x24, 0xd8, 0x25, 0x02, 0x7d, 0xf5, 0x8e,
		0xc4, 0xc7, 0x4b, 0x63, 0x64, 0xe4, 0xff, 0x56, 0xe4, 0x9f, 0x1a, 0x54, 0xfb, 0x04, 0x0b, 0xe7,
		0x6d, 0x53, 0x4a, 0x41, 0x87, 0xb1, 0x24, 0x11, 0xb2, 0x61, 0x93, 0x32, 0x97, 0xdc, 0x12, 0xd7,
		0x5e, 0x90, 0xfd, 0x6c, 0x69, 0xd4, 0x77, 0xdd, 0x0d, 0x33, 0xf3, 0x9d, 0x3f, 0xc7, 0x03, 0x3a,
		0x6f, 0x3b, 0xf8, 0x11, 0xd0, 0xdf, 0x49, 0xff, 0xe2, 0xa9, 0x3c, 0x58, 0x6f, 0x61, 0x89, 0x4f,
		0x7d, 0x3e, 0x44, 0xe7, 0xf0, 0x80, 0x30, 0x87, 0xbb, 0x94, 0x8d, 0x6c, 0x39, 0x0d, 0xb3, 0x02,
		0xdd, 0x6c, 0x3c, 0x5a, 0x1a, 0xab, 0xad, 0x98, 0x49, 0x45, 0x5b, 0x15, 0x32, 0xb7, 0xbb, 0x2b,
		0xe0, 0x95, 0xb9, 0x02, 0xee, 0x65, 0x4d, 0x47, 0xc4, 0x6b, 0x22, 0x22, 0xca, 0x99, 0xc9, 0x3c,
		0x9e, 0x10, 0x69, 0x10, 0xfa, 0x79, 0x23, 0x24, 0x6b, 0xf4, 0x18, 0xb6, 0x3c, 0x82, 0x65, 0x2c,
		0x88, 0x3d, 0xc9, 0xa8, 0xaa, 0xe1, 0x36, 0x95, 0x59, 0x05, 0xd0, 0x5f, 0xc1, 0x5e, 0x3f, 0x0e,
		0x43, 0x2e, 0x24, 0x71, 0xcf, 0x7c, 0x4a, 0x98, 0x54, 0x48, 0x94, 0xf4, 0xea, 0x88, 0xdb, 0x91,
		0x3b, 0x56, 0x91, 0x8b, 0x23, 0xde, 0x77, 0xc7, 0x68, 0x1f, 0xd6, 0x7f, 0xc2, 0x13, 0x9c, 0x02,
		0x59, 0xcc, 0xb5, 0x64, 0xdf, 0x77, 0xc7, 0xfa, 0xaf, 0x05, 0x28, 0x5b, 0x44, 0x8a, 0x69, 0x8f,
		0xfb, 0xd4, 0x99, 0xa2, 0x16, 0x54, 0x29, 0xa3, 0x92, 0x62, 0xdf, 0xa6, 0x4c, 0x12, 0x31, 0xc1,
		0x99, 0xca, 0x72, 0x63, 0xdf, 0xc8, 0xc6, 0x8b, 0x91, 0x8f, 0x17, 0xa3, 0xa5, 0xc6, 0x8b, 0xb5,
		0xa5, 0x5c, 0x4c, 0xe5, 0x81, 0xea, 0xb0, 0x3d, 0xc4, 0xce, 0x98, 0x7b, 0x9e, 0xed, 0x70, 0xe2,
		0x79, 0xd4, 0x49, 0x64, 0xa6, 0xb9, 0x35, 0x0b, 0x29, 0xe8, 0x6c, 0x86, 0x24, 0x69, 0x03, 0x7c,
		0x4b, 0x83, 0x38, 0x98, 0xa5, 0x2d, 0xbc, 0x37, 0xad, 0x72, 0xb9, 0x4b, 0xfb, 0xe9, 0x2c, 0x0a,
		0x96, 0x92, 0x04, 0xa1, 0x8c, 0x6a, 0xab, 0x47, 0xda, 0x71, 0xf1, 0x8e, 0xda, 0x54, 0x66, 0xf4,
		0x02, 0x1e, 0x32, 0xce, 0x6c, 0x91, 0x1c, 0x1d, 0x0f, 0x7d, 0x62, 0x13, 0x21, 0xb8, 0xb0, 0xb3,
		0x91, 0x12, 0xd5, 0x8a, 0x47, 0x85, 0xe3, 0x0d, 0xab, 0xc6, 0x38, 0xb3, 0x72, 0x46, 0x3b, 0x21,
		0x58, 0x19, 0x8e, 0x5e, 0xc2, 0x36, 0xb9, 0x0d, 0x69, 0x26, 0x64, 0x26, 0xb9, 0xf4, 0x3e, 0xc9,
		0x68, 0xe6, 0x95, 0xab, 0xd6, 0x03, 0xd8, 0x33, 0x23, 0xee, 0xa7, 0xc6, 0x0b, 0xc1, 0xe3, 0xb0,
		0x87, 0x85, 0xa4, 0xe9, 0x70, 0x5e, 0x32, 0x30, 0xd1, 0x97, 0x50, 0x8c, 0x24, 0x96, 0x59, 0xc1,
		0x6f, 0x36, 0x8e, 0x97, 0x16, 0xe9, 0x62, 0xc0, 0x7e, 0xc2, 0xb7, 0x32, 0x37, 0x7d, 0x02, 0x0f,
		0x17, 0xd1, 0x33, 0xce, 0x3c, 0x3a, 0x52, 0x0a, 0xd1, 0x35, 0x54, 0x69, 0x0e, 0xdb, 0xa3, 0x04,
		0xcf, 0x5b, 0xfb, 0xb3, 0x7f, 0x90, 0xe9, 0x4e, 0xba, 0xb5, 0x45, 0x17, 0x80, 0x48, 0xff, 0x43,
		0x83, 0x83, 0x66, 0x34, 0x65, 0x4e, 0xfe, 0xdb, 0x58, 0xcc, 0x5b, 0x83, 0x35, 0xc2, 0x92, 0x7b,
		0xce, 0xfe, 0x41, 0xeb, 0x56, 0xbe, 0x45, 0x0d, 0xd8, 0x09, 0x05, 0x71, 0x89, 0x47, 0x19, 0x71,
		0xed, 0x9f, 0x63, 0x12, 0x13, 0x3b, 0xbd, 0x95, 0xac, 0x94, 0xb7, 0x67, 0xe0, 0xb7, 0x09, 0xd6,
		0x49, 0x2e, 0xe9, 0x10, 0x20, 0x23, 0xa6, 0xed, 0x5c, 0x48, 0x89, 0x1b, 0xa9, 0x25, 0x6d, 0xd4,
		0xaf, 0xa1, 0x92, 0xc1, 0x4e, 0xaa, 0x21, 0x2d, 0x92, 0x72, 0xe3, 0x70, 0xe9, 0x01, 0xf3, 0x29,
		0x61, 0x95, 0x53, 0x97, 0x4c, 0xf5, 0x93, 0x1b, 0xa8, 0xcc, 0x0f, 0x02, 0xb4, 0x0f, 0x3b, 0xed,
		0xce, 0x59, 0xb7, 0x65, 0x76, 0x2e, 0xec, 0xc1, 0x9b, 0x5e, 0xdb, 0x36, 0x3b, 0xaf, 0x9b, 0x57,
		0x66, 0xab, 0xfa, 0x3f, 0x74, 0x00, 0xbb, 0x8b, 0xd0, 0xe0, 0xd2, 0x32, 0xcf, 0x07, 0xd6, 0x75,
		0x55, 0x43, 0xbb, 0x80, 0x16, 0xb1, 0x97, 0xfd, 0x6e, 0xa7, 0xba, 0x82, 0x6a, 0xf0, 0xc1, 0xa2,
		0xbd, 0x67, 0x75, 0x07, 0xdd, 0xa7, 0xd5, 0xc2, 0x93, 0x5f, 0x60, 0x7b, 0xc9, 0xc7, 0x45, 0x8f,
		0xe0, 0xd0, 0xec, 0x77, 0xaf, 0x9a, 0x03, 0xb3, 0xdb, 0xb1, 0x2f, 0xac, 0xee, 0x77, 0x3d, 0xbb,
		0x3f, 0x68, 0x0e, 0xe6, 0x75, 0xdc, 0x4b, 0xb9, 0x6c, 0x37, 0xaf, 0x06, 0x97, 0x6f, 0xaa, 0xda,
		0xfd, 0x94, 0x96, 0xd5, 0x34, 0x3b, 0xed, 0x56, 0x75, 0xe5, 0xf4, 0x07, 0xd8, 0x73, 0x78, 0xb0,
		0xec, 0xa6, 0x4e, 0xcb, 0x67, 0xe9, 0x13, 0xa5, 0x97, 0x54, 0x7d, 0x4f, 0xfb, 0xfe, 0x64, 0x44,
		0xe5, 0xdb, 0x78, 0x68, 0x38, 0x3c, 0xa8, 0xcf, 0x3f, 0x68, 0x3e, 0xa7, 0xae, 0x5f, 0x1f, 0xf1,
		0xec, 0x99, 0xa2, 0x5e, 0x37, 0xcf, 0x71, 0x48, 0x27, 0x27, 0xc3, 0x52, 0x6a, 0x7b, 0xfa, 0x57,
		0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0xd9, 0xb2, 0xe0, 0x01, 0x09, 0x00, 0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
	// uber/cadence/admin/v1/history.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x4d, 0x4a, 0x2d,
		0xd2, 0x4f, 0x4e, 0x4c, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f,
		0x33, 0xd4, 0xcf, 0xc8, 0x2c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
		0x12, 0x05, 0x29, 0xd2, 0x83, 0x2a, 0xd2, 0x03, 0x2b, 0xd2, 0x2b, 0x33, 0x54, 0xf2, 0xe4, 0x12,
		0x0a, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xf3, 0x80, 0x28, 0xf7, 0x2c, 0x49, 0xcd, 0x15, 0x92,
		0xe4, 0xe2, 0x48, 0x2d, 0x4b, 0xcd, 0x2b, 0x89, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
		0x0e, 0x62, 0x07, 0xf3, 0x3d, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0xcb, 0x20, 0x1a, 0x24, 0x98, 0x20,
		0x32, 0x50, 0xae, 0x52, 0x09, 0x17, 0x1f, 0xaa, 0x51, 0x42, 0x8a, 0x5c, 0x3c, 0x49, 0x45, 0x89,
		0x79, 0xc9, 0x19, 0xf1, 0x25, 0xf9, 0xd9, 0xa9, 0x79, 0x60, 0xa3, 0x78, 0x82, 0xb8, 0x21, 0x62,
		0x21, 0x20, 0x21, 0x21, 0x7b, 0x2e, 0xd6, 0xcc, 0x92, 0xd4, 0xdc, 0x62, 0x09, 0x26, 0x05, 0x66,
		0x0d, 0x6e, 0x23, 0x4d, 0x3d, 0xac, 0xce, 0xd4, 0xc3, 0x74, 0x63, 0x10, 0x44, 0x9f, 0x93, 0x79,
		0x94, 0x69, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x72, 0x48, 0xe8,
		0x66, 0xa6, 0xe4, 0xe8, 0xa7, 0xe7, 0xeb, 0x83, 0xfd, 0x0f, 0x0f, 0x16, 0x6b, 0x30, 0xa3, 0xcc,
		0x30, 0x89, 0x0d, 0x2c, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x44, 0x14, 0xd7, 0xd4, 0x3e,
		0x01, 0x00, 0x00,
	},
}
