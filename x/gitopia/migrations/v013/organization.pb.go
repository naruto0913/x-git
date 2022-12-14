// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/organization.proto

package v013

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type OrganizationMember_Role int32

const (
	OrganizationMember_MEMBER OrganizationMember_Role = 0
	OrganizationMember_OWNER  OrganizationMember_Role = 1
)

var OrganizationMember_Role_name = map[int32]string{
	0: "MEMBER",
	1: "OWNER",
}

var OrganizationMember_Role_value = map[string]int32{
	"MEMBER": 0,
	"OWNER":  1,
}

func (x OrganizationMember_Role) String() string {
	return proto.EnumName(OrganizationMember_Role_name, int32(x))
}

func (OrganizationMember_Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2866e48a11710e51, []int{2, 0}
}

type Organization struct {
	Creator      string                    `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id           uint64                    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Address      string                    `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Name         string                    `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	AvatarUrl    string                    `protobuf:"bytes,5,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	Followers    []string                  `protobuf:"bytes,6,rep,name=followers,proto3" json:"followers,omitempty"`
	Following    []string                  `protobuf:"bytes,7,rep,name=following,proto3" json:"following,omitempty"`
	Repositories []*OrganizationRepository `protobuf:"bytes,8,rep,name=repositories,proto3" json:"repositories,omitempty"`
	Teams        []uint64                  `protobuf:"varint,9,rep,packed,name=teams,proto3" json:"teams,omitempty"`
	Members      []*OrganizationMember     `protobuf:"bytes,10,rep,name=members,proto3" json:"members,omitempty"`
	Location     string                    `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	Website      string                    `protobuf:"bytes,12,opt,name=website,proto3" json:"website,omitempty"`
	Verified     bool                      `protobuf:"varint,13,opt,name=verified,proto3" json:"verified,omitempty"`
	Description  string                    `protobuf:"bytes,14,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt    int64                     `protobuf:"varint,15,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    int64                     `protobuf:"varint,16,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (m *Organization) Reset()         { *m = Organization{} }
func (m *Organization) String() string { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()    {}
func (*Organization) Descriptor() ([]byte, []int) {
	return fileDescriptor_2866e48a11710e51, []int{0}
}
func (m *Organization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Organization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Organization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Organization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organization.Merge(m, src)
}
func (m *Organization) XXX_Size() int {
	return m.Size()
}
func (m *Organization) XXX_DiscardUnknown() {
	xxx_messageInfo_Organization.DiscardUnknown(m)
}

var xxx_messageInfo_Organization proto.InternalMessageInfo

func (m *Organization) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Organization) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Organization) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Organization) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Organization) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *Organization) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *Organization) GetFollowing() []string {
	if m != nil {
		return m.Following
	}
	return nil
}

func (m *Organization) GetRepositories() []*OrganizationRepository {
	if m != nil {
		return m.Repositories
	}
	return nil
}

func (m *Organization) GetTeams() []uint64 {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *Organization) GetMembers() []*OrganizationMember {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Organization) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Organization) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Organization) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *Organization) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Organization) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Organization) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type OrganizationRepository struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *OrganizationRepository) Reset()         { *m = OrganizationRepository{} }
func (m *OrganizationRepository) String() string { return proto.CompactTextString(m) }
func (*OrganizationRepository) ProtoMessage()    {}
func (*OrganizationRepository) Descriptor() ([]byte, []int) {
	return fileDescriptor_2866e48a11710e51, []int{1}
}
func (m *OrganizationRepository) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrganizationRepository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrganizationRepository.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrganizationRepository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationRepository.Merge(m, src)
}
func (m *OrganizationRepository) XXX_Size() int {
	return m.Size()
}
func (m *OrganizationRepository) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationRepository.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationRepository proto.InternalMessageInfo

func (m *OrganizationRepository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrganizationRepository) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type OrganizationMember struct {
	Id   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role OrganizationMember_Role `protobuf:"varint,2,opt,name=role,proto3,enum=gitopia.gitopia.gitopia.OrganizationMember_Role" json:"role,omitempty"`
}

func (m *OrganizationMember) Reset()         { *m = OrganizationMember{} }
func (m *OrganizationMember) String() string { return proto.CompactTextString(m) }
func (*OrganizationMember) ProtoMessage()    {}
func (*OrganizationMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_2866e48a11710e51, []int{2}
}
func (m *OrganizationMember) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrganizationMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrganizationMember.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrganizationMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationMember.Merge(m, src)
}
func (m *OrganizationMember) XXX_Size() int {
	return m.Size()
}
func (m *OrganizationMember) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationMember.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationMember proto.InternalMessageInfo

func (m *OrganizationMember) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrganizationMember) GetRole() OrganizationMember_Role {
	if m != nil {
		return m.Role
	}
	return OrganizationMember_MEMBER
}

func init() {
	proto.RegisterEnum("gitopia.gitopia.gitopia.OrganizationMember_Role", OrganizationMember_Role_name, OrganizationMember_Role_value)
	proto.RegisterType((*Organization)(nil), "gitopia.gitopia.gitopia.Organization")
	proto.RegisterType((*OrganizationRepository)(nil), "gitopia.gitopia.gitopia.OrganizationRepository")
	proto.RegisterType((*OrganizationMember)(nil), "gitopia.gitopia.gitopia.OrganizationMember")
}

func init() { proto.RegisterFile("gitopia/organization.proto", fileDescriptor_2866e48a11710e51) }

var fileDescriptor_2866e48a11710e51 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0xad, 0x9b, 0xf4, 0x75, 0x5b, 0x4a, 0x65, 0x8d, 0xc0, 0xaa, 0x20, 0x8a, 0xba, 0x8a, 0x40,
	0x4a, 0xd1, 0xb0, 0x65, 0xc3, 0x68, 0xba, 0x2c, 0x23, 0x19, 0x21, 0x24, 0x76, 0x6e, 0xe3, 0x09,
	0x96, 0x92, 0x38, 0x72, 0x3c, 0x33, 0x0c, 0x7f, 0xc0, 0x8e, 0x6f, 0xe0, 0x6b, 0x58, 0xce, 0x92,
	0x25, 0x6a, 0x7f, 0x04, 0xd9, 0x79, 0x34, 0x14, 0x90, 0x98, 0x95, 0x7d, 0xce, 0xb9, 0xe7, 0x3a,
	0x39, 0xd7, 0x86, 0x79, 0x2c, 0xb4, 0xcc, 0x05, 0x5b, 0x4a, 0x15, 0xb3, 0x4c, 0x7c, 0x66, 0x5a,
	0xc8, 0x2c, 0xcc, 0x95, 0xd4, 0x12, 0x3f, 0xae, 0xb4, 0xf0, 0x68, 0x9d, 0x9f, 0xc4, 0x32, 0x96,
	0xb6, 0x66, 0x69, 0x76, 0x65, 0xf9, 0xe2, 0x9b, 0x0b, 0x93, 0x8b, 0x56, 0x17, 0x4c, 0x60, 0xb0,
	0x55, 0x9c, 0x69, 0xa9, 0x08, 0xf2, 0x51, 0x30, 0xa2, 0x35, 0xc4, 0x53, 0xe8, 0x8a, 0x88, 0x74,
	0x7d, 0x14, 0xb8, 0xb4, 0x2b, 0x22, 0x53, 0xc9, 0xa2, 0x48, 0xf1, 0xa2, 0x20, 0x4e, 0x59, 0x59,
	0x41, 0x8c, 0xc1, 0xcd, 0x58, 0xca, 0x89, 0x6b, 0x69, 0xbb, 0xc7, 0x4f, 0x60, 0xc4, 0xae, 0x99,
	0x66, 0xea, 0x9d, 0x4a, 0x48, 0xcf, 0x0a, 0x07, 0xc2, 0xa8, 0x97, 0x32, 0x49, 0xe4, 0x0d, 0x57,
	0x05, 0xe9, 0xfb, 0x8e, 0x51, 0x1b, 0xe2, 0xa0, 0x8a, 0x2c, 0x26, 0x83, 0xb6, 0x2a, 0xb2, 0x18,
	0xbf, 0x85, 0x89, 0xe2, 0xb9, 0x2c, 0x84, 0x96, 0x4a, 0xf0, 0x82, 0x0c, 0x7d, 0x27, 0x18, 0x9f,
	0x2e, 0xc3, 0x7f, 0x04, 0x11, 0xb6, 0x7f, 0x97, 0xd6, 0xc6, 0x5b, 0xfa, 0x5b, 0x13, 0x7c, 0x02,
	0x3d, 0xcd, 0x59, 0x5a, 0x90, 0x91, 0xef, 0x04, 0x2e, 0x2d, 0x01, 0x5e, 0xc1, 0x20, 0xe5, 0xe9,
	0xc6, 0x7c, 0x24, 0xd8, 0x53, 0x9e, 0xff, 0xd7, 0x29, 0x6b, 0xeb, 0xa1, 0xb5, 0x17, 0xcf, 0x61,
	0x98, 0xc8, 0xad, 0x95, 0xc8, 0xd8, 0x46, 0xd1, 0x60, 0x93, 0xea, 0x0d, 0xdf, 0x14, 0x42, 0x73,
	0x32, 0x29, 0x53, 0xad, 0xa0, 0x71, 0x5d, 0x73, 0x25, 0x2e, 0x05, 0x8f, 0xc8, 0x03, 0x1f, 0x05,
	0x43, 0xda, 0x60, 0xec, 0xc3, 0x38, 0xe2, 0xc5, 0x56, 0x89, 0xdc, 0x36, 0x9d, 0x5a, 0x67, 0x9b,
	0x32, 0x19, 0xda, 0x41, 0xf2, 0xe8, 0xb5, 0x26, 0x0f, 0x7d, 0x14, 0x38, 0xf4, 0x40, 0x18, 0xf5,
	0x2a, 0x8f, 0x2a, 0x75, 0x56, 0xaa, 0x0d, 0xb1, 0x78, 0x05, 0x8f, 0xfe, 0x1e, 0x5a, 0x33, 0x69,
	0xd4, 0x9a, 0xf4, 0xd1, 0x3d, 0x59, 0x7c, 0x41, 0x80, 0xff, 0x4c, 0xa3, 0x2a, 0x2b, 0x8d, 0xe6,
	0x3a, 0x9d, 0x83, 0xab, 0x64, 0xc2, 0xad, 0x71, 0x7a, 0xfa, 0xe2, 0x1e, 0xc1, 0x86, 0x54, 0x26,
	0x9c, 0x5a, 0xf7, 0xe2, 0x29, 0xb8, 0x06, 0x61, 0x80, 0xfe, 0x7a, 0xb5, 0x3e, 0x5b, 0xd1, 0x59,
	0x07, 0x8f, 0xa0, 0x77, 0xf1, 0xfe, 0xcd, 0x8a, 0xce, 0xd0, 0xd9, 0xf9, 0xf7, 0x9d, 0x87, 0xee,
	0x76, 0x1e, 0xfa, 0xb9, 0xf3, 0xd0, 0xd7, 0xbd, 0xd7, 0xb9, 0xdb, 0x7b, 0x9d, 0x1f, 0x7b, 0xaf,
	0xf3, 0xe1, 0x59, 0x2c, 0xf4, 0xc7, 0xab, 0x4d, 0xb8, 0x95, 0xe9, 0xb2, 0x7e, 0x5e, 0xf5, 0xfa,
	0xa9, 0xd9, 0xe9, 0xdb, 0x9c, 0x17, 0x9b, 0xbe, 0x7d, 0x3b, 0x2f, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x00, 0x08, 0x59, 0x66, 0x88, 0x03, 0x00, 0x00,
}

func (m *Organization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Organization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Organization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != 0 {
		i = encodeVarintOrganization(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.CreatedAt != 0 {
		i = encodeVarintOrganization(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x78
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x72
	}
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOrganization(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Teams) > 0 {
		dAtA2 := make([]byte, len(m.Teams)*10)
		var j1 int
		for _, num := range m.Teams {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintOrganization(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Repositories) > 0 {
		for iNdEx := len(m.Repositories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Repositories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOrganization(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Following) > 0 {
		for iNdEx := len(m.Following) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Following[iNdEx])
			copy(dAtA[i:], m.Following[iNdEx])
			i = encodeVarintOrganization(dAtA, i, uint64(len(m.Following[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Followers) > 0 {
		for iNdEx := len(m.Followers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Followers[iNdEx])
			copy(dAtA[i:], m.Followers[iNdEx])
			i = encodeVarintOrganization(dAtA, i, uint64(len(m.Followers[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.AvatarUrl) > 0 {
		i -= len(m.AvatarUrl)
		copy(dAtA[i:], m.AvatarUrl)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.AvatarUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintOrganization(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OrganizationRepository) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrganizationRepository) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrganizationRepository) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintOrganization(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OrganizationMember) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrganizationMember) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrganizationMember) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Role != 0 {
		i = encodeVarintOrganization(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintOrganization(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrganization(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrganization(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Organization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovOrganization(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	l = len(m.AvatarUrl)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	if len(m.Followers) > 0 {
		for _, s := range m.Followers {
			l = len(s)
			n += 1 + l + sovOrganization(uint64(l))
		}
	}
	if len(m.Following) > 0 {
		for _, s := range m.Following {
			l = len(s)
			n += 1 + l + sovOrganization(uint64(l))
		}
	}
	if len(m.Repositories) > 0 {
		for _, e := range m.Repositories {
			l = e.Size()
			n += 1 + l + sovOrganization(uint64(l))
		}
	}
	if len(m.Teams) > 0 {
		l = 0
		for _, e := range m.Teams {
			l += sovOrganization(uint64(e))
		}
		n += 1 + sovOrganization(uint64(l)) + l
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovOrganization(uint64(l))
		}
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	if m.Verified {
		n += 2
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovOrganization(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 2 + sovOrganization(uint64(m.UpdatedAt))
	}
	return n
}

func (m *OrganizationRepository) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovOrganization(uint64(m.Id))
	}
	return n
}

func (m *OrganizationMember) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovOrganization(uint64(l))
	}
	if m.Role != 0 {
		n += 1 + sovOrganization(uint64(m.Role))
	}
	return n
}

func sovOrganization(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrganization(x uint64) (n int) {
	return sovOrganization(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Organization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrganization
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Organization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Organization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvatarUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AvatarUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Followers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Followers = append(m.Followers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Following", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Following = append(m.Following, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repositories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repositories = append(m.Repositories, &OrganizationRepository{})
			if err := m.Repositories[len(m.Repositories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrganization
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Teams = append(m.Teams, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrganization
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthOrganization
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthOrganization
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Teams) == 0 {
					m.Teams = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOrganization
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Teams = append(m.Teams, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Teams", wireType)
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &OrganizationMember{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Verified = bool(v != 0)
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrganization(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrganization
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OrganizationRepository) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrganization
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OrganizationRepository: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrganizationRepository: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrganization(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrganization
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OrganizationMember) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrganization
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OrganizationMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrganizationMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrganization
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrganization
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= OrganizationMember_Role(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrganization(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrganization
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOrganization(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrganization
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrganization
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOrganization
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrganization
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrganization
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrganization        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrganization          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrganization = fmt.Errorf("proto: unexpected end of group")
)
