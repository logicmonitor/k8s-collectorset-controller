// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/prediction_service.proto

package automl

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for
// [PredictionService.Predict][google.cloud.automl.v1.PredictionService.Predict].
type PredictRequest struct {
	// Name of the model requested to serve the prediction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Payload to perform a prediction on. The payload must match the
	// problem type that the model was trained to solve.
	Payload *ExamplePayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific parameters, any string must be up to 25000
	// characters long.
	//
	// *  For Image Classification:
	//
	//    `score_threshold` - (float) A value from 0.0 to 1.0. When the model
	//     makes predictions for an image, it will only produce results that have
	//     at least this confidence score. The default is 0.5.
	//
	//  *  For Image Object Detection:
	//    `score_threshold` - (float) When Model detects objects on the image,
	//        it will only produce bounding boxes which have at least this
	//        confidence score. Value in 0 to 1 range, default is 0.5.
	//    `max_bounding_box_count` - (int64) No more than this number of bounding
	//        boxes will be returned in the response. Default is 100, the
	//        requested value may be limited by server.
	Params               map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictRequest) Reset()         { *m = PredictRequest{} }
func (m *PredictRequest) String() string { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()    {}
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a60105ec759f54a4, []int{0}
}

func (m *PredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictRequest.Unmarshal(m, b)
}
func (m *PredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictRequest.Marshal(b, m, deterministic)
}
func (m *PredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictRequest.Merge(m, src)
}
func (m *PredictRequest) XXX_Size() int {
	return xxx_messageInfo_PredictRequest.Size(m)
}
func (m *PredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictRequest proto.InternalMessageInfo

func (m *PredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictRequest) GetPayload() *ExamplePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Response message for
// [PredictionService.Predict][google.cloud.automl.v1.PredictionService.Predict].
type PredictResponse struct {
	// Prediction result.
	// Translation and Text Sentiment will return precisely one payload.
	Payload []*AnnotationPayload `protobuf:"bytes,1,rep,name=payload,proto3" json:"payload,omitempty"`
	// The preprocessed example that AutoML actually makes prediction on.
	// Empty if AutoML does not preprocess the input example.
	// * For Text Extraction:
	//   If the input is a .pdf file, the OCR'ed text will be provided in
	//   [document_text][google.cloud.automl.v1.Document.document_text].
	//
	// * For Text Classification:
	//   If the input is a .pdf file, the OCR'ed trucated text will be provided in
	//   [document_text][google.cloud.automl.v1.Document.document_text].
	//
	// * For Text Sentiment:
	//   If the input is a .pdf file, the OCR'ed trucated text will be provided in
	//   [document_text][google.cloud.automl.v1.Document.document_text].
	PreprocessedInput *ExamplePayload `protobuf:"bytes,3,opt,name=preprocessed_input,json=preprocessedInput,proto3" json:"preprocessed_input,omitempty"`
	// Additional domain-specific prediction response metadata.
	//
	// * For Image Object Detection:
	//  `max_bounding_box_count` - (int64) At most that many bounding boxes per
	//      image could have been returned.
	//
	// * For Text Sentiment:
	//  `sentiment_score` - (float, deprecated) A value between -1 and 1,
	//      -1 maps to least positive sentiment, while 1 maps to the most positive
	//      one and the higher the score, the more positive the sentiment in the
	//      document is. Yet these values are relative to the training data, so
	//      e.g. if all data was positive then -1 will be also positive (though
	//      the least).
	//      The sentiment_score shouldn't be confused with "score" or "magnitude"
	//      from the previous Natural Language Sentiment Analysis API.
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictResponse) Reset()         { *m = PredictResponse{} }
func (m *PredictResponse) String() string { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()    {}
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a60105ec759f54a4, []int{1}
}

func (m *PredictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictResponse.Unmarshal(m, b)
}
func (m *PredictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictResponse.Marshal(b, m, deterministic)
}
func (m *PredictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictResponse.Merge(m, src)
}
func (m *PredictResponse) XXX_Size() int {
	return xxx_messageInfo_PredictResponse.Size(m)
}
func (m *PredictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictResponse proto.InternalMessageInfo

func (m *PredictResponse) GetPayload() []*AnnotationPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictResponse) GetPreprocessedInput() *ExamplePayload {
	if m != nil {
		return m.PreprocessedInput
	}
	return nil
}

func (m *PredictResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Request message for
// [PredictionService.BatchPredict][google.cloud.automl.v1.PredictionService.BatchPredict].
type BatchPredictRequest struct {
	// Name of the model requested to serve the batch prediction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The input configuration for batch prediction.
	InputConfig *BatchPredictInputConfig `protobuf:"bytes,3,opt,name=input_config,json=inputConfig,proto3" json:"input_config,omitempty"`
	// Required. The Configuration specifying where output predictions should
	// be written.
	OutputConfig *BatchPredictOutputConfig `protobuf:"bytes,4,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	// Additional domain-specific parameters for the predictions, any string must
	// be up to 25000 characters long.
	//
	// *  For Text Classification:
	//
	//    `score_threshold` - (float) A value from 0.0 to 1.0. When the model
	//         makes predictions for a text snippet, it will only produce results
	//         that have at least this confidence score. The default is 0.5.
	//
	// *  For Image Classification:
	//
	//    `score_threshold` - (float) A value from 0.0 to 1.0. When the model
	//         makes predictions for an image, it will only produce results that
	//         have at least this confidence score. The default is 0.5.
	//
	// *  For Image Object Detection:
	//
	//    `score_threshold` - (float) When Model detects objects on the image,
	//        it will only produce bounding boxes which have at least this
	//        confidence score. Value in 0 to 1 range, default is 0.5.
	//    `max_bounding_box_count` - (int64) No more than this number of bounding
	//        boxes will be produced per image. Default is 100, the
	//        requested value may be limited by server.
	Params               map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BatchPredictRequest) Reset()         { *m = BatchPredictRequest{} }
func (m *BatchPredictRequest) String() string { return proto.CompactTextString(m) }
func (*BatchPredictRequest) ProtoMessage()    {}
func (*BatchPredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a60105ec759f54a4, []int{2}
}

func (m *BatchPredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPredictRequest.Unmarshal(m, b)
}
func (m *BatchPredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPredictRequest.Marshal(b, m, deterministic)
}
func (m *BatchPredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPredictRequest.Merge(m, src)
}
func (m *BatchPredictRequest) XXX_Size() int {
	return xxx_messageInfo_BatchPredictRequest.Size(m)
}
func (m *BatchPredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPredictRequest proto.InternalMessageInfo

func (m *BatchPredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BatchPredictRequest) GetInputConfig() *BatchPredictInputConfig {
	if m != nil {
		return m.InputConfig
	}
	return nil
}

func (m *BatchPredictRequest) GetOutputConfig() *BatchPredictOutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

func (m *BatchPredictRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Result of the Batch Predict. This message is returned in
// [response][google.longrunning.Operation.response] of the operation returned
// by the
// [PredictionService.BatchPredict][google.cloud.automl.v1.PredictionService.BatchPredict].
type BatchPredictResult struct {
	// Additional domain-specific prediction response metadata.
	//
	// *  For Image Object Detection:
	//  `max_bounding_box_count` - (int64) At most that many bounding boxes per
	//      image could have been returned.
	Metadata             map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BatchPredictResult) Reset()         { *m = BatchPredictResult{} }
func (m *BatchPredictResult) String() string { return proto.CompactTextString(m) }
func (*BatchPredictResult) ProtoMessage()    {}
func (*BatchPredictResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a60105ec759f54a4, []int{3}
}

func (m *BatchPredictResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPredictResult.Unmarshal(m, b)
}
func (m *BatchPredictResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPredictResult.Marshal(b, m, deterministic)
}
func (m *BatchPredictResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPredictResult.Merge(m, src)
}
func (m *BatchPredictResult) XXX_Size() int {
	return xxx_messageInfo_BatchPredictResult.Size(m)
}
func (m *BatchPredictResult) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPredictResult.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPredictResult proto.InternalMessageInfo

func (m *BatchPredictResult) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "google.cloud.automl.v1.PredictRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.PredictRequest.ParamsEntry")
	proto.RegisterType((*PredictResponse)(nil), "google.cloud.automl.v1.PredictResponse")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.PredictResponse.MetadataEntry")
	proto.RegisterType((*BatchPredictRequest)(nil), "google.cloud.automl.v1.BatchPredictRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.BatchPredictRequest.ParamsEntry")
	proto.RegisterType((*BatchPredictResult)(nil), "google.cloud.automl.v1.BatchPredictResult")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.BatchPredictResult.MetadataEntry")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/prediction_service.proto", fileDescriptor_a60105ec759f54a4)
}

var fileDescriptor_a60105ec759f54a4 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4f, 0x4f, 0x13, 0x41,
	0x14, 0xcf, 0x6e, 0xf9, 0x23, 0x53, 0x50, 0x19, 0x15, 0xcb, 0x46, 0x23, 0xa9, 0x09, 0x60, 0x8d,
	0x3b, 0xb6, 0xc6, 0x80, 0x0b, 0x24, 0xb6, 0x84, 0x18, 0x8c, 0x84, 0x5a, 0x85, 0x83, 0x21, 0x69,
	0x86, 0xed, 0xb0, 0xac, 0xee, 0xce, 0x8c, 0xbb, 0xb3, 0x45, 0x62, 0xbc, 0xf8, 0x15, 0x3c, 0x19,
	0x13, 0xaf, 0x26, 0x7e, 0x07, 0x2f, 0x1e, 0x39, 0xea, 0x17, 0xf0, 0xe0, 0xc9, 0xc4, 0xef, 0x60,
	0x76, 0x66, 0xdb, 0x6e, 0xc1, 0x86, 0xf6, 0xc0, 0x6d, 0x66, 0xdf, 0xfb, 0xfd, 0xde, 0xef, 0xbd,
	0x79, 0xfb, 0x1e, 0x40, 0x0e, 0x63, 0x8e, 0x47, 0x90, 0xed, 0xb1, 0xa8, 0x81, 0x70, 0x24, 0x98,
	0xef, 0xa1, 0x66, 0x11, 0xf1, 0x80, 0x34, 0x5c, 0x5b, 0xb8, 0x8c, 0xd6, 0x43, 0x12, 0x34, 0x5d,
	0x9b, 0x98, 0x3c, 0x60, 0x82, 0xc1, 0x29, 0x05, 0x30, 0x25, 0xc0, 0x54, 0x00, 0xb3, 0x59, 0x34,
	0xae, 0x25, 0x44, 0x98, 0xbb, 0x08, 0x53, 0xca, 0x04, 0x8e, 0xc1, 0xa1, 0x42, 0x19, 0x57, 0x53,
	0x56, 0xdb, 0x73, 0x09, 0x15, 0x89, 0x61, 0x3a, 0x65, 0x08, 0x48, 0xc8, 0xa2, 0xa0, 0x15, 0xc9,
	0xe8, 0x25, 0xad, 0xc3, 0x5e, 0xe7, 0xf8, 0xd0, 0x63, 0xb8, 0x91, 0x00, 0xe6, 0x7a, 0x00, 0x1a,
	0x58, 0xe0, 0xba, 0x2b, 0x88, 0xdf, 0x52, 0x73, 0xa3, 0x87, 0xa3, 0xcb, 0x4e, 0x61, 0x62, 0x9c,
	0x04, 0x5d, 0x79, 0xdd, 0x4c, 0x1c, 0x3d, 0x46, 0x9d, 0x20, 0xa2, 0xd4, 0xa5, 0xce, 0x09, 0xa7,
	0xfc, 0x5f, 0x0d, 0x9c, 0xaf, 0xaa, 0x7a, 0xd6, 0xc8, 0xeb, 0x88, 0x84, 0x02, 0x42, 0x30, 0x44,
	0xb1, 0x4f, 0x72, 0xda, 0x8c, 0x36, 0x3f, 0x56, 0x93, 0x67, 0xf8, 0x10, 0x8c, 0x26, 0xf9, 0xe4,
	0xf4, 0x19, 0x6d, 0x3e, 0x5b, 0x9a, 0x35, 0xff, 0x5f, 0x6b, 0x73, 0xed, 0x0d, 0xf6, 0xb9, 0x47,
	0xaa, 0xca, 0xbb, 0xd6, 0x82, 0xc1, 0xc7, 0x60, 0x84, 0xe3, 0x00, 0xfb, 0x61, 0x2e, 0x33, 0x93,
	0x99, 0xcf, 0x96, 0x4a, 0xbd, 0x08, 0xba, 0xd5, 0x98, 0x55, 0x09, 0x5a, 0xa3, 0x22, 0x38, 0xac,
	0x25, 0x0c, 0xc6, 0x03, 0x90, 0x4d, 0x7d, 0x86, 0x17, 0x41, 0xe6, 0x15, 0x39, 0x4c, 0xf4, 0xc6,
	0x47, 0x78, 0x19, 0x0c, 0x37, 0xb1, 0x17, 0x11, 0x29, 0x76, 0xac, 0xa6, 0x2e, 0x96, 0xbe, 0xa8,
	0xe5, 0xbf, 0xe9, 0xe0, 0x42, 0x3b, 0x42, 0xc8, 0x19, 0x0d, 0x09, 0x5c, 0xed, 0x24, 0xa7, 0x49,
	0x6d, 0xb7, 0x7a, 0x69, 0x2b, 0xb7, 0x9f, 0xf7, 0x44, 0x7e, 0x5b, 0x00, 0xf2, 0x80, 0xf0, 0x80,
	0xd9, 0x24, 0x0c, 0x49, 0xa3, 0xee, 0x52, 0x1e, 0x89, 0x5c, 0x66, 0xa0, 0x62, 0x4d, 0xa6, 0x19,
	0xd6, 0x63, 0x02, 0xf8, 0x14, 0x9c, 0xf3, 0x89, 0xc0, 0x71, 0x9b, 0xe4, 0x74, 0x29, 0xee, 0xfe,
	0xa9, 0x85, 0x53, 0x69, 0x99, 0x1b, 0x09, 0x4e, 0xd5, 0xae, 0x4d, 0x63, 0x2c, 0x81, 0x89, 0x2e,
	0xd3, 0x40, 0xf5, 0xfb, 0xa5, 0x83, 0x4b, 0x15, 0x2c, 0xec, 0xfd, 0x3e, 0x9a, 0xa6, 0x06, 0xc6,
	0x65, 0x15, 0xea, 0x36, 0xa3, 0x7b, 0xae, 0x93, 0x14, 0x03, 0xf5, 0xd2, 0x9f, 0xa6, 0x95, 0xc9,
	0xaf, 0x4a, 0x58, 0x2d, 0xeb, 0x76, 0x2e, 0x70, 0x0b, 0x4c, 0xb0, 0x48, 0xa4, 0x48, 0x87, 0x24,
	0xe9, 0xdd, 0x7e, 0x48, 0x37, 0x25, 0x30, 0x61, 0x1d, 0x67, 0xa9, 0x1b, 0xdc, 0x6c, 0x77, 0xe7,
	0xb0, 0x2c, 0xf2, 0x42, 0x3f, 0x7c, 0x67, 0xd4, 0xa2, 0x5f, 0x34, 0x00, 0xbb, 0xc3, 0x84, 0x91,
	0x27, 0xe0, 0xf3, 0x54, 0x27, 0xa8, 0x36, 0x5d, 0xec, 0x4f, 0x64, 0x8c, 0x3e, 0x93, 0x66, 0x28,
	0x7d, 0xca, 0x80, 0xc9, 0x6a, 0x7b, 0x18, 0x3f, 0x53, 0xb3, 0x18, 0x7e, 0xd4, 0xc0, 0x68, 0xf2,
	0x15, 0xce, 0xf6, 0xf7, 0x97, 0x1b, 0x73, 0x7d, 0x36, 0x75, 0x7e, 0xe5, 0xfd, 0xcf, 0xdf, 0x1f,
	0xf4, 0x85, 0x7c, 0x29, 0x1e, 0x79, 0x6f, 0xe3, 0x36, 0x5b, 0xe1, 0x01, 0x7b, 0x49, 0x6c, 0x11,
	0xa2, 0x02, 0xf2, 0x98, 0xad, 0xa6, 0x1b, 0x2a, 0x20, 0x9f, 0x35, 0x88, 0x17, 0xa2, 0xc2, 0x3b,
	0x2b, 0xd9, 0x17, 0x96, 0x56, 0x80, 0x9f, 0x35, 0x30, 0x9e, 0xae, 0x0e, 0xbc, 0x3d, 0xc0, 0x43,
	0x1b, 0xd7, 0x5b, 0xce, 0xa9, 0x91, 0x6a, 0x6e, 0xb6, 0x46, 0x6a, 0xbe, 0x22, 0xb5, 0x2d, 0xe7,
	0x17, 0x06, 0xd0, 0xb6, 0x9b, 0x0a, 0x63, 0x69, 0x05, 0x63, 0xfd, 0xa8, 0x7c, 0x25, 0x11, 0xa1,
	0x62, 0x61, 0xee, 0x86, 0xa6, 0xcd, 0xfc, 0x1f, 0x65, 0x73, 0x5f, 0x08, 0x1e, 0x5a, 0x08, 0x1d,
	0x1c, 0x1c, 0x1c, 0x33, 0xc6, 0x6b, 0x60, 0x5f, 0x6d, 0x84, 0x3b, 0xdc, 0xc3, 0x62, 0x8f, 0x05,
	0x7e, 0xe5, 0xbb, 0x06, 0x0c, 0x9b, 0xf9, 0x3d, 0x12, 0xac, 0x4c, 0x9d, 0x78, 0xb9, 0x6a, 0xbc,
	0x11, 0xaa, 0xda, 0x8b, 0xe5, 0x04, 0xe1, 0x30, 0x0f, 0x53, 0xc7, 0x64, 0x81, 0x83, 0x1c, 0x42,
	0xe5, 0xbe, 0x40, 0x9d, 0xb8, 0xc7, 0x17, 0xd0, 0x92, 0x3a, 0x7d, 0xd5, 0xa7, 0x1e, 0x29, 0xf8,
	0xaa, 0x0c, 0x58, 0x8e, 0x04, 0xdb, 0x78, 0x62, 0x6e, 0x17, 0x8f, 0x5a, 0x86, 0x1d, 0x69, 0xd8,
	0x91, 0x06, 0x6f, 0x67, 0xbb, 0xf8, 0x47, 0x9f, 0x56, 0x06, 0xcb, 0x92, 0x16, 0xcb, 0x52, 0x18,
	0xcb, 0xda, 0x2e, 0xee, 0x8e, 0xc8, 0xb0, 0xf7, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x91,
	0x49, 0xdb, 0x0c, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Perform an online prediction. The prediction result will be directly
	// returned in the response.
	// Available for following ML problems, and their expected request payloads:
	// * Image Classification - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                          up to 30MB.
	// * Image Object Detection - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                            up to 30MB.
	// * Text Classification - TextSnippet, content up to 60,000 characters,
	//                         UTF-8 encoded.
	// * Text Extraction - TextSnippet, content up to 30,000 characters,
	//                     UTF-8 NFC encoded.
	// * Translation - TextSnippet, content up to 25,000 characters, UTF-8
	//                 encoded.
	// * Text Sentiment - TextSnippet, content up 500 characters, UTF-8
	//                     encoded.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	// Perform a batch prediction. Unlike the online
	// [Predict][google.cloud.automl.v1.PredictionService.Predict], batch
	// prediction result won't be immediately available in the response. Instead,
	// a long running operation object is returned. User can poll the operation
	// result via [GetOperation][google.longrunning.Operations.GetOperation]
	// method. Once the operation is done,
	// [BatchPredictResult][google.cloud.automl.v1.BatchPredictResult] is returned
	// in the [response][google.longrunning.Operation.response] field. Available
	// for following ML problems:
	// * Image Classification
	// * Image Object Detection
	// * Text Extraction
	BatchPredict(ctx context.Context, in *BatchPredictRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.automl.v1.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) BatchPredict(ctx context.Context, in *BatchPredictRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.automl.v1.PredictionService/BatchPredict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Perform an online prediction. The prediction result will be directly
	// returned in the response.
	// Available for following ML problems, and their expected request payloads:
	// * Image Classification - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                          up to 30MB.
	// * Image Object Detection - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                            up to 30MB.
	// * Text Classification - TextSnippet, content up to 60,000 characters,
	//                         UTF-8 encoded.
	// * Text Extraction - TextSnippet, content up to 30,000 characters,
	//                     UTF-8 NFC encoded.
	// * Translation - TextSnippet, content up to 25,000 characters, UTF-8
	//                 encoded.
	// * Text Sentiment - TextSnippet, content up 500 characters, UTF-8
	//                     encoded.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	// Perform a batch prediction. Unlike the online
	// [Predict][google.cloud.automl.v1.PredictionService.Predict], batch
	// prediction result won't be immediately available in the response. Instead,
	// a long running operation object is returned. User can poll the operation
	// result via [GetOperation][google.longrunning.Operations.GetOperation]
	// method. Once the operation is done,
	// [BatchPredictResult][google.cloud.automl.v1.BatchPredictResult] is returned
	// in the [response][google.longrunning.Operation.response] field. Available
	// for following ML problems:
	// * Image Classification
	// * Image Object Detection
	// * Text Extraction
	BatchPredict(context.Context, *BatchPredictRequest) (*longrunning.Operation, error)
}

// UnimplementedPredictionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (*UnimplementedPredictionServiceServer) Predict(ctx context.Context, req *PredictRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (*UnimplementedPredictionServiceServer) BatchPredict(ctx context.Context, req *BatchPredictRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchPredict not implemented")
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.automl.v1.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_BatchPredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchPredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).BatchPredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.automl.v1.PredictionService/BatchPredict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).BatchPredict(ctx, req.(*BatchPredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.automl.v1.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
		{
			MethodName: "BatchPredict",
			Handler:    _PredictionService_BatchPredict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/automl/v1/prediction_service.proto",
}
