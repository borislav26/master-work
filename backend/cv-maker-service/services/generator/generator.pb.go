// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: generator.proto

package generator

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

type GeneratePDFFromTemplateAndDataProgrammingLanguage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Percentage int64  `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *GeneratePDFFromTemplateAndDataProgrammingLanguage) Reset() {
	*x = GeneratePDFFromTemplateAndDataProgrammingLanguage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePDFFromTemplateAndDataProgrammingLanguage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePDFFromTemplateAndDataProgrammingLanguage) ProtoMessage() {}

func (x *GeneratePDFFromTemplateAndDataProgrammingLanguage) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePDFFromTemplateAndDataProgrammingLanguage.ProtoReflect.Descriptor instead.
func (*GeneratePDFFromTemplateAndDataProgrammingLanguage) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{0}
}

func (x *GeneratePDFFromTemplateAndDataProgrammingLanguage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GeneratePDFFromTemplateAndDataProgrammingLanguage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataProgrammingLanguage) GetPercentage() int64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

type GeneratePDFFromTemplateAndDataEducation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Faculty string `protobuf:"bytes,1,opt,name=faculty,proto3" json:"faculty,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	From    string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To      string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *GeneratePDFFromTemplateAndDataEducation) Reset() {
	*x = GeneratePDFFromTemplateAndDataEducation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePDFFromTemplateAndDataEducation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePDFFromTemplateAndDataEducation) ProtoMessage() {}

func (x *GeneratePDFFromTemplateAndDataEducation) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePDFFromTemplateAndDataEducation.ProtoReflect.Descriptor instead.
func (*GeneratePDFFromTemplateAndDataEducation) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{1}
}

func (x *GeneratePDFFromTemplateAndDataEducation) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataEducation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataEducation) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataEducation) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type GeneratePDFFromTemplateAndDataExperience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company     string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	JobTitle    string `protobuf:"bytes,2,opt,name=jobTitle,proto3" json:"jobTitle,omitempty"`
	From        string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To          string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GeneratePDFFromTemplateAndDataExperience) Reset() {
	*x = GeneratePDFFromTemplateAndDataExperience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePDFFromTemplateAndDataExperience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePDFFromTemplateAndDataExperience) ProtoMessage() {}

func (x *GeneratePDFFromTemplateAndDataExperience) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePDFFromTemplateAndDataExperience.ProtoReflect.Descriptor instead.
func (*GeneratePDFFromTemplateAndDataExperience) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{2}
}

func (x *GeneratePDFFromTemplateAndDataExperience) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataExperience) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataExperience) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataExperience) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataExperience) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GeneratePDFFromTemplateAndDataLanguage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GeneratePDFFromTemplateAndDataLanguage) Reset() {
	*x = GeneratePDFFromTemplateAndDataLanguage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePDFFromTemplateAndDataLanguage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePDFFromTemplateAndDataLanguage) ProtoMessage() {}

func (x *GeneratePDFFromTemplateAndDataLanguage) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePDFFromTemplateAndDataLanguage.ProtoReflect.Descriptor instead.
func (*GeneratePDFFromTemplateAndDataLanguage) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{3}
}

func (x *GeneratePDFFromTemplateAndDataLanguage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GeneratePDFFromTemplateAndDataUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string                                               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string                                               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string                                               `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	ProgrammingLanguages []*GeneratePDFFromTemplateAndDataProgrammingLanguage `protobuf:"bytes,4,rep,name=programmingLanguages,proto3" json:"programmingLanguages,omitempty"`
	Description          string                                               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Educations           []*GeneratePDFFromTemplateAndDataEducation           `protobuf:"bytes,6,rep,name=educations,proto3" json:"educations,omitempty"`
	Experiences          []*GeneratePDFFromTemplateAndDataExperience          `protobuf:"bytes,7,rep,name=experiences,proto3" json:"experiences,omitempty"`
	Languages            []*GeneratePDFFromTemplateAndDataLanguage            `protobuf:"bytes,8,rep,name=languages,proto3" json:"languages,omitempty"`
	Website              string                                               `protobuf:"bytes,9,opt,name=website,proto3" json:"website,omitempty"`
	LinkedinProfile      string                                               `protobuf:"bytes,10,opt,name=linkedinProfile,proto3" json:"linkedinProfile,omitempty"`
	GithubProfile        string                                               `protobuf:"bytes,11,opt,name=githubProfile,proto3" json:"githubProfile,omitempty"`
	ProfilePhotoBytes    string                                               `protobuf:"bytes,12,opt,name=profilePhotoBytes,proto3" json:"profilePhotoBytes,omitempty"`
}

func (x *GeneratePDFFromTemplateAndDataUser) Reset() {
	*x = GeneratePDFFromTemplateAndDataUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePDFFromTemplateAndDataUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePDFFromTemplateAndDataUser) ProtoMessage() {}

func (x *GeneratePDFFromTemplateAndDataUser) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePDFFromTemplateAndDataUser.ProtoReflect.Descriptor instead.
func (*GeneratePDFFromTemplateAndDataUser) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{4}
}

func (x *GeneratePDFFromTemplateAndDataUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataUser) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataUser) GetProgrammingLanguages() []*GeneratePDFFromTemplateAndDataProgrammingLanguage {
	if x != nil {
		return x.ProgrammingLanguages
	}
	return nil
}

func (x *GeneratePDFFromTemplateAndDataUser) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataUser) GetEducations() []*GeneratePDFFromTemplateAndDataEducation {
	if x != nil {
		return x.Educations
	}
	return nil
}

func (x *GeneratePDFFromTemplateAndDataUser) GetExperiences() []*GeneratePDFFromTemplateAndDataExperience {
	if x != nil {
		return x.Experiences
	}
	return nil
}

func (x *GeneratePDFFromTemplateAndDataUser) GetLanguages() []*GeneratePDFFromTemplateAndDataLanguage {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *GeneratePDFFromTemplateAndDataUser) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataUser) GetLinkedinProfile() string {
	if x != nil {
		return x.LinkedinProfile
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataUser) GetGithubProfile() string {
	if x != nil {
		return x.GithubProfile
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataUser) GetProfilePhotoBytes() string {
	if x != nil {
		return x.ProfilePhotoBytes
	}
	return ""
}

type GeneratePDFFromTemplateAndDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template string                              `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	UserData *GeneratePDFFromTemplateAndDataUser `protobuf:"bytes,6,opt,name=userData,proto3" json:"userData,omitempty"`
}

func (x *GeneratePDFFromTemplateAndDataRequest) Reset() {
	*x = GeneratePDFFromTemplateAndDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePDFFromTemplateAndDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePDFFromTemplateAndDataRequest) ProtoMessage() {}

func (x *GeneratePDFFromTemplateAndDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePDFFromTemplateAndDataRequest.ProtoReflect.Descriptor instead.
func (*GeneratePDFFromTemplateAndDataRequest) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{5}
}

func (x *GeneratePDFFromTemplateAndDataRequest) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *GeneratePDFFromTemplateAndDataRequest) GetUserData() *GeneratePDFFromTemplateAndDataUser {
	if x != nil {
		return x.UserData
	}
	return nil
}

type GeneratePDFFromTemplateAndDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	PdfBytes []byte `protobuf:"bytes,2,opt,name=pdfBytes,proto3" json:"pdfBytes,omitempty"`
}

func (x *GeneratePDFFromTemplateAndDataResponse) Reset() {
	*x = GeneratePDFFromTemplateAndDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePDFFromTemplateAndDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePDFFromTemplateAndDataResponse) ProtoMessage() {}

func (x *GeneratePDFFromTemplateAndDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_generator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePDFFromTemplateAndDataResponse.ProtoReflect.Descriptor instead.
func (*GeneratePDFFromTemplateAndDataResponse) Descriptor() ([]byte, []int) {
	return file_generator_proto_rawDescGZIP(), []int{6}
}

func (x *GeneratePDFFromTemplateAndDataResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *GeneratePDFFromTemplateAndDataResponse) GetPdfBytes() []byte {
	if x != nil {
		return x.PdfBytes
	}
	return nil
}

var File_generator_proto protoreflect.FileDescriptor

var file_generator_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x77, 0x0a, 0x31,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x0a, 0x27, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x28, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6a,
	0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a,
	0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a,
	0x26, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x98, 0x05, 0x0a, 0x22,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x70,
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69,
	0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x0a, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72,
	0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x55, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x4f, 0x0a,
	0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x25, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x26, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x64, 0x66,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x64, 0x66,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x1e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x30, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_generator_proto_rawDescOnce sync.Once
	file_generator_proto_rawDescData = file_generator_proto_rawDesc
)

func file_generator_proto_rawDescGZIP() []byte {
	file_generator_proto_rawDescOnce.Do(func() {
		file_generator_proto_rawDescData = protoimpl.X.CompressGZIP(file_generator_proto_rawDescData)
	})
	return file_generator_proto_rawDescData
}

var file_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_generator_proto_goTypes = []interface{}{
	(*GeneratePDFFromTemplateAndDataProgrammingLanguage)(nil), // 0: generator.GeneratePDFFromTemplateAndDataProgrammingLanguage
	(*GeneratePDFFromTemplateAndDataEducation)(nil),           // 1: generator.GeneratePDFFromTemplateAndDataEducation
	(*GeneratePDFFromTemplateAndDataExperience)(nil),          // 2: generator.GeneratePDFFromTemplateAndDataExperience
	(*GeneratePDFFromTemplateAndDataLanguage)(nil),            // 3: generator.GeneratePDFFromTemplateAndDataLanguage
	(*GeneratePDFFromTemplateAndDataUser)(nil),                // 4: generator.GeneratePDFFromTemplateAndDataUser
	(*GeneratePDFFromTemplateAndDataRequest)(nil),             // 5: generator.GeneratePDFFromTemplateAndDataRequest
	(*GeneratePDFFromTemplateAndDataResponse)(nil),            // 6: generator.GeneratePDFFromTemplateAndDataResponse
}
var file_generator_proto_depIdxs = []int32{
	0, // 0: generator.GeneratePDFFromTemplateAndDataUser.programmingLanguages:type_name -> generator.GeneratePDFFromTemplateAndDataProgrammingLanguage
	1, // 1: generator.GeneratePDFFromTemplateAndDataUser.educations:type_name -> generator.GeneratePDFFromTemplateAndDataEducation
	2, // 2: generator.GeneratePDFFromTemplateAndDataUser.experiences:type_name -> generator.GeneratePDFFromTemplateAndDataExperience
	3, // 3: generator.GeneratePDFFromTemplateAndDataUser.languages:type_name -> generator.GeneratePDFFromTemplateAndDataLanguage
	4, // 4: generator.GeneratePDFFromTemplateAndDataRequest.userData:type_name -> generator.GeneratePDFFromTemplateAndDataUser
	5, // 5: generator.GeneratorService.GeneratePDFFromTemplateAndData:input_type -> generator.GeneratePDFFromTemplateAndDataRequest
	6, // 6: generator.GeneratorService.GeneratePDFFromTemplateAndData:output_type -> generator.GeneratePDFFromTemplateAndDataResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_generator_proto_init() }
func file_generator_proto_init() {
	if File_generator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_generator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePDFFromTemplateAndDataProgrammingLanguage); i {
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
		file_generator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePDFFromTemplateAndDataEducation); i {
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
		file_generator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePDFFromTemplateAndDataExperience); i {
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
		file_generator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePDFFromTemplateAndDataLanguage); i {
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
		file_generator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePDFFromTemplateAndDataUser); i {
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
		file_generator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePDFFromTemplateAndDataRequest); i {
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
		file_generator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePDFFromTemplateAndDataResponse); i {
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
			RawDescriptor: file_generator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_generator_proto_goTypes,
		DependencyIndexes: file_generator_proto_depIdxs,
		MessageInfos:      file_generator_proto_msgTypes,
	}.Build()
	File_generator_proto = out.File
	file_generator_proto_rawDesc = nil
	file_generator_proto_goTypes = nil
	file_generator_proto_depIdxs = nil
}