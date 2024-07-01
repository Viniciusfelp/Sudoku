// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/sudoku.proto

package sudoku

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SudokuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grid []int32 `protobuf:"varint,1,rep,packed,name=grid,proto3" json:"grid,omitempty"`
}

func (x *SudokuRequest) Reset() {
	*x = SudokuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sudoku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SudokuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SudokuRequest) ProtoMessage() {}

func (x *SudokuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sudoku_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SudokuRequest.ProtoReflect.Descriptor instead.
func (*SudokuRequest) Descriptor() ([]byte, []int) {
	return file_proto_sudoku_proto_rawDescGZIP(), []int{0}
}

func (x *SudokuRequest) GetGrid() []int32 {
	if x != nil {
		return x.Grid
	}
	return nil
}

type SudokuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SolvedGrid []int32 `protobuf:"varint,1,rep,packed,name=solvedGrid,proto3" json:"solvedGrid,omitempty"`
	Success    bool    `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SudokuResponse) Reset() {
	*x = SudokuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sudoku_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SudokuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SudokuResponse) ProtoMessage() {}

func (x *SudokuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sudoku_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SudokuResponse.ProtoReflect.Descriptor instead.
func (*SudokuResponse) Descriptor() ([]byte, []int) {
	return file_proto_sudoku_proto_rawDescGZIP(), []int{1}
}

func (x *SudokuResponse) GetSolvedGrid() []int32 {
	if x != nil {
		return x.SolvedGrid
	}
	return nil
}

func (x *SudokuResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_sudoku_proto protoreflect.FileDescriptor

var file_proto_sudoku_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x22, 0x23, 0x0a, 0x0d,
	0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x67, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x67, 0x72, 0x69,
	0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x47, 0x72, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x47,
	0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x4c, 0x0a,
	0x0c, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x3c, 0x0a,
	0x0b, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x12, 0x15, 0x2e, 0x73,
	0x75, 0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x53, 0x75, 0x64,
	0x6f, 0x6b, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_sudoku_proto_rawDescOnce sync.Once
	file_proto_sudoku_proto_rawDescData = file_proto_sudoku_proto_rawDesc
)

func file_proto_sudoku_proto_rawDescGZIP() []byte {
	file_proto_sudoku_proto_rawDescOnce.Do(func() {
		file_proto_sudoku_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sudoku_proto_rawDescData)
	})
	return file_proto_sudoku_proto_rawDescData
}

var file_proto_sudoku_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_sudoku_proto_goTypes = []any{
	(*SudokuRequest)(nil),  // 0: sudoku.SudokuRequest
	(*SudokuResponse)(nil), // 1: sudoku.SudokuResponse
}
var file_proto_sudoku_proto_depIdxs = []int32{
	0, // 0: sudoku.SudokuSolver.SolveSudoku:input_type -> sudoku.SudokuRequest
	1, // 1: sudoku.SudokuSolver.SolveSudoku:output_type -> sudoku.SudokuResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sudoku_proto_init() }
func file_proto_sudoku_proto_init() {
	if File_proto_sudoku_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sudoku_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SudokuRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sudoku_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SudokuResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_sudoku_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sudoku_proto_goTypes,
		DependencyIndexes: file_proto_sudoku_proto_depIdxs,
		MessageInfos:      file_proto_sudoku_proto_msgTypes,
	}.Build()
	File_proto_sudoku_proto = out.File
	file_proto_sudoku_proto_rawDesc = nil
	file_proto_sudoku_proto_goTypes = nil
	file_proto_sudoku_proto_depIdxs = nil
}
