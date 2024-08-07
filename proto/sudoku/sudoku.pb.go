// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: sudoku.proto

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

	Grid []*Row `protobuf:"bytes,1,rep,name=grid,proto3" json:"grid,omitempty"`
}

func (x *SudokuRequest) Reset() {
	*x = SudokuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudoku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SudokuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SudokuRequest) ProtoMessage() {}

func (x *SudokuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sudoku_proto_msgTypes[0]
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
	return file_sudoku_proto_rawDescGZIP(), []int{0}
}

func (x *SudokuRequest) GetGrid() []*Row {
	if x != nil {
		return x.Grid
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells []int32 `protobuf:"varint,1,rep,packed,name=cells,proto3" json:"cells,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudoku_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_sudoku_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_sudoku_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetCells() []int32 {
	if x != nil {
		return x.Cells
	}
	return nil
}

type SudokuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SolvedGrid []*Row `protobuf:"bytes,1,rep,name=solvedGrid,proto3" json:"solvedGrid,omitempty"`
	Success    bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SudokuResponse) Reset() {
	*x = SudokuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudoku_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SudokuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SudokuResponse) ProtoMessage() {}

func (x *SudokuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sudoku_proto_msgTypes[2]
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
	return file_sudoku_proto_rawDescGZIP(), []int{2}
}

func (x *SudokuResponse) GetSolvedGrid() []*Row {
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

var File_sudoku_proto protoreflect.FileDescriptor

var file_sudoku_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x22, 0x30, 0x0a, 0x0d, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x67, 0x72, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x52,
	0x6f, 0x77, 0x52, 0x04, 0x67, 0x72, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x65, 0x6c, 0x6c, 0x73, 0x22, 0x57, 0x0a, 0x0e, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x64, 0x47, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x75,
	0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64,
	0x47, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x96,
	0x01, 0x0a, 0x0c, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12,
	0x3c, 0x0a, 0x0b, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x12, 0x15,
	0x2e, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x53,
	0x75, 0x64, 0x6f, 0x6b, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x17, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x43, 0x6f, 0x6e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x12, 0x15, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x6b,
	0x75, 0x2e, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x2e, 0x53, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x75, 0x64, 0x6f, 0x6b, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sudoku_proto_rawDescOnce sync.Once
	file_sudoku_proto_rawDescData = file_sudoku_proto_rawDesc
)

func file_sudoku_proto_rawDescGZIP() []byte {
	file_sudoku_proto_rawDescOnce.Do(func() {
		file_sudoku_proto_rawDescData = protoimpl.X.CompressGZIP(file_sudoku_proto_rawDescData)
	})
	return file_sudoku_proto_rawDescData
}

var file_sudoku_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sudoku_proto_goTypes = []any{
	(*SudokuRequest)(nil),  // 0: sudoku.SudokuRequest
	(*Row)(nil),            // 1: sudoku.Row
	(*SudokuResponse)(nil), // 2: sudoku.SudokuResponse
}
var file_sudoku_proto_depIdxs = []int32{
	1, // 0: sudoku.SudokuRequest.grid:type_name -> sudoku.Row
	1, // 1: sudoku.SudokuResponse.solvedGrid:type_name -> sudoku.Row
	0, // 2: sudoku.SudokuSolver.SolveSudoku:input_type -> sudoku.SudokuRequest
	0, // 3: sudoku.SudokuSolver.SolveSudokuConcurrently:input_type -> sudoku.SudokuRequest
	2, // 4: sudoku.SudokuSolver.SolveSudoku:output_type -> sudoku.SudokuResponse
	2, // 5: sudoku.SudokuSolver.SolveSudokuConcurrently:output_type -> sudoku.SudokuResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sudoku_proto_init() }
func file_sudoku_proto_init() {
	if File_sudoku_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sudoku_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_sudoku_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Row); i {
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
		file_sudoku_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
			RawDescriptor: file_sudoku_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sudoku_proto_goTypes,
		DependencyIndexes: file_sudoku_proto_depIdxs,
		MessageInfos:      file_sudoku_proto_msgTypes,
	}.Build()
	File_sudoku_proto = out.File
	file_sudoku_proto_rawDesc = nil
	file_sudoku_proto_goTypes = nil
	file_sudoku_proto_depIdxs = nil
}
