// Code generated by protoc-gen-go.
// source: replication.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	replication.proto

It has these top-level messages:
	ServiceLinkTicket
	ServiceLinkRequest
	ServiceLinkResponse
	AuthGroup
	AuthSecret
	AuthIPWhitelist
	AuthIPWhitelistAssignment
	AuthDB
	AuthDBRevision
	ChangeNotification
	ReplicationPushRequest
	ReplicationPushResponse
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Status codes.
type ServiceLinkResponse_Status int32

const (
	// The service is now linked and primary will be pushing updates to it.
	ServiceLinkResponse_SUCCESS ServiceLinkResponse_Status = 0
	// Primary do not replies.
	ServiceLinkResponse_TRANSPORT_ERROR ServiceLinkResponse_Status = 1
	// Linking ticket is invalid or expired.
	ServiceLinkResponse_BAD_TICKET ServiceLinkResponse_Status = 2
	// Linking ticket was generated for another app, not the calling one.
	ServiceLinkResponse_AUTH_ERROR ServiceLinkResponse_Status = 3
)

var ServiceLinkResponse_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "TRANSPORT_ERROR",
	2: "BAD_TICKET",
	3: "AUTH_ERROR",
}
var ServiceLinkResponse_Status_value = map[string]int32{
	"SUCCESS":         0,
	"TRANSPORT_ERROR": 1,
	"BAD_TICKET":      2,
	"AUTH_ERROR":      3,
}

func (x ServiceLinkResponse_Status) Enum() *ServiceLinkResponse_Status {
	p := new(ServiceLinkResponse_Status)
	*p = x
	return p
}
func (x ServiceLinkResponse_Status) String() string {
	return proto.EnumName(ServiceLinkResponse_Status_name, int32(x))
}
func (x *ServiceLinkResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServiceLinkResponse_Status_value, data, "ServiceLinkResponse_Status")
	if err != nil {
		return err
	}
	*x = ServiceLinkResponse_Status(value)
	return nil
}
func (ServiceLinkResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

// Overall status of the operation.
type ReplicationPushResponse_Status int32

const (
	// Replica accepted the push request and updated its copy of auth db.
	ReplicationPushResponse_APPLIED ReplicationPushResponse_Status = 0
	// Replica has a newer version of AuthDB, the push request is skipped.
	ReplicationPushResponse_SKIPPED ReplicationPushResponse_Status = 1
	// Non fatal error happened, the push request may be retried.
	ReplicationPushResponse_TRANSIENT_ERROR ReplicationPushResponse_Status = 2
	// Fatal error happened, the push request must not be retried.
	ReplicationPushResponse_FATAL_ERROR ReplicationPushResponse_Status = 3
)

var ReplicationPushResponse_Status_name = map[int32]string{
	0: "APPLIED",
	1: "SKIPPED",
	2: "TRANSIENT_ERROR",
	3: "FATAL_ERROR",
}
var ReplicationPushResponse_Status_value = map[string]int32{
	"APPLIED":         0,
	"SKIPPED":         1,
	"TRANSIENT_ERROR": 2,
	"FATAL_ERROR":     3,
}

func (x ReplicationPushResponse_Status) Enum() *ReplicationPushResponse_Status {
	p := new(ReplicationPushResponse_Status)
	*p = x
	return p
}
func (x ReplicationPushResponse_Status) String() string {
	return proto.EnumName(ReplicationPushResponse_Status_name, int32(x))
}
func (x *ReplicationPushResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReplicationPushResponse_Status_value, data, "ReplicationPushResponse_Status")
	if err != nil {
		return err
	}
	*x = ReplicationPushResponse_Status(value)
	return nil
}
func (ReplicationPushResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11, 0}
}

// Error codes, for TRANSIENT_ERROR and FATAL_ERROR statuses.
type ReplicationPushResponse_ErrorCode int32

const (
	// Trying to push an update to service that is not a replica.
	ReplicationPushResponse_NOT_A_REPLICA ReplicationPushResponse_ErrorCode = 1
	// Replica doesn't know about the service that pushing the update.
	ReplicationPushResponse_FORBIDDEN ReplicationPushResponse_ErrorCode = 2
	// Signature headers are missing.
	ReplicationPushResponse_MISSING_SIGNATURE ReplicationPushResponse_ErrorCode = 3
	// Signature is not valid.
	ReplicationPushResponse_BAD_SIGNATURE ReplicationPushResponse_ErrorCode = 4
	// Format of the request is not valid.
	ReplicationPushResponse_BAD_REQUEST ReplicationPushResponse_ErrorCode = 5
)

var ReplicationPushResponse_ErrorCode_name = map[int32]string{
	1: "NOT_A_REPLICA",
	2: "FORBIDDEN",
	3: "MISSING_SIGNATURE",
	4: "BAD_SIGNATURE",
	5: "BAD_REQUEST",
}
var ReplicationPushResponse_ErrorCode_value = map[string]int32{
	"NOT_A_REPLICA":     1,
	"FORBIDDEN":         2,
	"MISSING_SIGNATURE": 3,
	"BAD_SIGNATURE":     4,
	"BAD_REQUEST":       5,
}

func (x ReplicationPushResponse_ErrorCode) Enum() *ReplicationPushResponse_ErrorCode {
	p := new(ReplicationPushResponse_ErrorCode)
	*p = x
	return p
}
func (x ReplicationPushResponse_ErrorCode) String() string {
	return proto.EnumName(ReplicationPushResponse_ErrorCode_name, int32(x))
}
func (x *ReplicationPushResponse_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReplicationPushResponse_ErrorCode_value, data, "ReplicationPushResponse_ErrorCode")
	if err != nil {
		return err
	}
	*x = ReplicationPushResponse_ErrorCode(value)
	return nil
}
func (ReplicationPushResponse_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11, 1}
}

// Generated by Primary, passed to Replica to initiate linking process.
type ServiceLinkTicket struct {
	// GAE application ID of Primary that generated this ticket. Replica will send
	// ServiceLinkRequest to this service when it processes the ticket.
	PrimaryId *string `protobuf:"bytes,1,req,name=primary_id" json:"primary_id,omitempty"`
	// URL to the root page of a primary service, i.e. https://<...>.appspot.com.
	// Useful when testing on dev appserver and on non-default version.
	PrimaryUrl *string `protobuf:"bytes,2,req,name=primary_url" json:"primary_url,omitempty"`
	// Identity of a user that generated this ticket.
	GeneratedBy *string `protobuf:"bytes,3,req,name=generated_by" json:"generated_by,omitempty"`
	// Opaque blob passed back to Primary in ServiceLinkRequest. Its exact
	// structure is an implementation detail of Primary. It contains app_id of
	// a replica this ticket is intended for, timestamp and HMAC tag.
	Ticket           []byte `protobuf:"bytes,4,req,name=ticket" json:"ticket,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ServiceLinkTicket) Reset()                    { *m = ServiceLinkTicket{} }
func (m *ServiceLinkTicket) String() string            { return proto.CompactTextString(m) }
func (*ServiceLinkTicket) ProtoMessage()               {}
func (*ServiceLinkTicket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceLinkTicket) GetPrimaryId() string {
	if m != nil && m.PrimaryId != nil {
		return *m.PrimaryId
	}
	return ""
}

func (m *ServiceLinkTicket) GetPrimaryUrl() string {
	if m != nil && m.PrimaryUrl != nil {
		return *m.PrimaryUrl
	}
	return ""
}

func (m *ServiceLinkTicket) GetGeneratedBy() string {
	if m != nil && m.GeneratedBy != nil {
		return *m.GeneratedBy
	}
	return ""
}

func (m *ServiceLinkTicket) GetTicket() []byte {
	if m != nil {
		return m.Ticket
	}
	return nil
}

// Sent from Replica to Primary via direct serivce <-> service HTTP call,
// replicas app_id would be available via X-Appengine-Inbound-Appid header.
type ServiceLinkRequest struct {
	// Same ticket that was passed to Replica via ServiceLinkTicket.
	Ticket []byte `protobuf:"bytes,1,req,name=ticket" json:"ticket,omitempty"`
	// URL to use when making requests to Replica from Primary.
	ReplicaUrl *string `protobuf:"bytes,2,req,name=replica_url" json:"replica_url,omitempty"`
	// Identity of a user that accepted the ticket and initiated this request.
	InitiatedBy      *string `protobuf:"bytes,3,req,name=initiated_by" json:"initiated_by,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServiceLinkRequest) Reset()                    { *m = ServiceLinkRequest{} }
func (m *ServiceLinkRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceLinkRequest) ProtoMessage()               {}
func (*ServiceLinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceLinkRequest) GetTicket() []byte {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func (m *ServiceLinkRequest) GetReplicaUrl() string {
	if m != nil && m.ReplicaUrl != nil {
		return *m.ReplicaUrl
	}
	return ""
}

func (m *ServiceLinkRequest) GetInitiatedBy() string {
	if m != nil && m.InitiatedBy != nil {
		return *m.InitiatedBy
	}
	return ""
}

// Primary's response to ServiceLinkRequest. Always returned with HTTP code 200.
type ServiceLinkResponse struct {
	Status           *ServiceLinkResponse_Status `protobuf:"varint,1,req,name=status,enum=protocol.ServiceLinkResponse_Status" json:"status,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *ServiceLinkResponse) Reset()                    { *m = ServiceLinkResponse{} }
func (m *ServiceLinkResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceLinkResponse) ProtoMessage()               {}
func (*ServiceLinkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceLinkResponse) GetStatus() ServiceLinkResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ServiceLinkResponse_SUCCESS
}

// Some user group. Corresponds to AuthGroup entity in model.py.
type AuthGroup struct {
	// Name of the group.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// List of members that are explicitly in this group.
	Members []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
	// List of identity-glob expressions (like 'user:*@example.com').
	Globs []string `protobuf:"bytes,3,rep,name=globs" json:"globs,omitempty"`
	// List of nested group names.
	Nested []string `protobuf:"bytes,4,rep,name=nested" json:"nested,omitempty"`
	// Human readable description.
	Description *string `protobuf:"bytes,5,req,name=description" json:"description,omitempty"`
	// When the group was created. Microseconds since epoch.
	CreatedTs *int64 `protobuf:"varint,6,req,name=created_ts" json:"created_ts,omitempty"`
	// Who created the group.
	CreatedBy *string `protobuf:"bytes,7,req,name=created_by" json:"created_by,omitempty"`
	// When the group was modified last time. Microseconds since epoch.
	ModifiedTs *int64 `protobuf:"varint,8,req,name=modified_ts" json:"modified_ts,omitempty"`
	// Who modified the group last time.
	ModifiedBy *string `protobuf:"bytes,9,req,name=modified_by" json:"modified_by,omitempty"`
	// A name of the group that can modify or delete this group.
	Owners           *string `protobuf:"bytes,10,opt,name=owners" json:"owners,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthGroup) Reset()                    { *m = AuthGroup{} }
func (m *AuthGroup) String() string            { return proto.CompactTextString(m) }
func (*AuthGroup) ProtoMessage()               {}
func (*AuthGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AuthGroup) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AuthGroup) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *AuthGroup) GetGlobs() []string {
	if m != nil {
		return m.Globs
	}
	return nil
}

func (m *AuthGroup) GetNested() []string {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *AuthGroup) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *AuthGroup) GetCreatedTs() int64 {
	if m != nil && m.CreatedTs != nil {
		return *m.CreatedTs
	}
	return 0
}

func (m *AuthGroup) GetCreatedBy() string {
	if m != nil && m.CreatedBy != nil {
		return *m.CreatedBy
	}
	return ""
}

func (m *AuthGroup) GetModifiedTs() int64 {
	if m != nil && m.ModifiedTs != nil {
		return *m.ModifiedTs
	}
	return 0
}

func (m *AuthGroup) GetModifiedBy() string {
	if m != nil && m.ModifiedBy != nil {
		return *m.ModifiedBy
	}
	return ""
}

func (m *AuthGroup) GetOwners() string {
	if m != nil && m.Owners != nil {
		return *m.Owners
	}
	return ""
}

// Some secret blob. Corresponds to AuthSecret entity in model.py.
type AuthSecret struct {
	// Name of the secret.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Last several values of a secret, with current value in front.
	Values [][]byte `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	// When secret was modified last time. Microseconds since epoch.
	ModifiedTs *int64 `protobuf:"varint,3,req,name=modified_ts" json:"modified_ts,omitempty"`
	// Who modified the secret last time.
	ModifiedBy       *string `protobuf:"bytes,4,req,name=modified_by" json:"modified_by,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthSecret) Reset()                    { *m = AuthSecret{} }
func (m *AuthSecret) String() string            { return proto.CompactTextString(m) }
func (*AuthSecret) ProtoMessage()               {}
func (*AuthSecret) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AuthSecret) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AuthSecret) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *AuthSecret) GetModifiedTs() int64 {
	if m != nil && m.ModifiedTs != nil {
		return *m.ModifiedTs
	}
	return 0
}

func (m *AuthSecret) GetModifiedBy() string {
	if m != nil && m.ModifiedBy != nil {
		return *m.ModifiedBy
	}
	return ""
}

// A named set of whitelisted IP addresses. Corresponds to AuthIPWhitelist
// entity in model.py.
type AuthIPWhitelist struct {
	// Name of the IP whitelist.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The list of IP subnets.
	Subnets []string `protobuf:"bytes,2,rep,name=subnets" json:"subnets,omitempty"`
	// Human readable description.
	Description *string `protobuf:"bytes,3,req,name=description" json:"description,omitempty"`
	// When the list was created. Microseconds since epoch.
	CreatedTs *int64 `protobuf:"varint,4,req,name=created_ts" json:"created_ts,omitempty"`
	// Who created the list.
	CreatedBy *string `protobuf:"bytes,5,req,name=created_by" json:"created_by,omitempty"`
	// When the list was modified. Microseconds since epoch.
	ModifiedTs *int64 `protobuf:"varint,6,req,name=modified_ts" json:"modified_ts,omitempty"`
	// Who modified the list the last time.
	ModifiedBy       *string `protobuf:"bytes,7,req,name=modified_by" json:"modified_by,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthIPWhitelist) Reset()                    { *m = AuthIPWhitelist{} }
func (m *AuthIPWhitelist) String() string            { return proto.CompactTextString(m) }
func (*AuthIPWhitelist) ProtoMessage()               {}
func (*AuthIPWhitelist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AuthIPWhitelist) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AuthIPWhitelist) GetSubnets() []string {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *AuthIPWhitelist) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *AuthIPWhitelist) GetCreatedTs() int64 {
	if m != nil && m.CreatedTs != nil {
		return *m.CreatedTs
	}
	return 0
}

func (m *AuthIPWhitelist) GetCreatedBy() string {
	if m != nil && m.CreatedBy != nil {
		return *m.CreatedBy
	}
	return ""
}

func (m *AuthIPWhitelist) GetModifiedTs() int64 {
	if m != nil && m.ModifiedTs != nil {
		return *m.ModifiedTs
	}
	return 0
}

func (m *AuthIPWhitelist) GetModifiedBy() string {
	if m != nil && m.ModifiedBy != nil {
		return *m.ModifiedBy
	}
	return ""
}

// A pair (identity, IP whitelist name) plus some metadata. Corresponds to
// AuthIPWhitelistAssignments.Assignment model in model.py.
type AuthIPWhitelistAssignment struct {
	// Identity name to limit by IP whitelist.
	Identity *string `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`
	// Name of IP whitelist to use (see AuthIPWhitelist).
	IpWhitelist *string `protobuf:"bytes,2,req,name=ip_whitelist" json:"ip_whitelist,omitempty"`
	// Why the assignment was created.
	Comment *string `protobuf:"bytes,3,req,name=comment" json:"comment,omitempty"`
	// When the assignment was created. Microseconds since epoch.
	CreatedTs *int64 `protobuf:"varint,4,req,name=created_ts" json:"created_ts,omitempty"`
	// Who created the assignment.
	CreatedBy        *string `protobuf:"bytes,5,req,name=created_by" json:"created_by,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthIPWhitelistAssignment) Reset()                    { *m = AuthIPWhitelistAssignment{} }
func (m *AuthIPWhitelistAssignment) String() string            { return proto.CompactTextString(m) }
func (*AuthIPWhitelistAssignment) ProtoMessage()               {}
func (*AuthIPWhitelistAssignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AuthIPWhitelistAssignment) GetIdentity() string {
	if m != nil && m.Identity != nil {
		return *m.Identity
	}
	return ""
}

func (m *AuthIPWhitelistAssignment) GetIpWhitelist() string {
	if m != nil && m.IpWhitelist != nil {
		return *m.IpWhitelist
	}
	return ""
}

func (m *AuthIPWhitelistAssignment) GetComment() string {
	if m != nil && m.Comment != nil {
		return *m.Comment
	}
	return ""
}

func (m *AuthIPWhitelistAssignment) GetCreatedTs() int64 {
	if m != nil && m.CreatedTs != nil {
		return *m.CreatedTs
	}
	return 0
}

func (m *AuthIPWhitelistAssignment) GetCreatedBy() string {
	if m != nil && m.CreatedBy != nil {
		return *m.CreatedBy
	}
	return ""
}

// An entire database of auth configuration that is being replicated.
// Corresponds to AuthGlobalConfig entity in model.py, plus a list of all groups
// and a list of global secrets.
type AuthDB struct {
	// OAuth2 client_id to use to mint new OAuth2 tokens.
	OauthClientId *string `protobuf:"bytes,1,req,name=oauth_client_id" json:"oauth_client_id,omitempty"`
	// OAuth2 client secret. Not so secret really, since it's passed to clients.
	OauthClientSecret *string `protobuf:"bytes,2,req,name=oauth_client_secret" json:"oauth_client_secret,omitempty"`
	// Additional OAuth2 client_ids allowed to access the services.
	OauthAdditionalClientIds []string `protobuf:"bytes,3,rep,name=oauth_additional_client_ids" json:"oauth_additional_client_ids,omitempty"`
	// All groups.
	Groups []*AuthGroup `protobuf:"bytes,4,rep,name=groups" json:"groups,omitempty"`
	// Global secrets shared between services.
	Secrets []*AuthSecret `protobuf:"bytes,5,rep,name=secrets" json:"secrets,omitempty"`
	// All IP whitelists.
	IpWhitelists []*AuthIPWhitelist `protobuf:"bytes,6,rep,name=ip_whitelists" json:"ip_whitelists,omitempty"`
	// Mapping 'account -> IP whitlist to use for that account'.
	IpWhitelistAssignments []*AuthIPWhitelistAssignment `protobuf:"bytes,7,rep,name=ip_whitelist_assignments" json:"ip_whitelist_assignments,omitempty"`
	XXX_unrecognized       []byte                       `json:"-"`
}

func (m *AuthDB) Reset()                    { *m = AuthDB{} }
func (m *AuthDB) String() string            { return proto.CompactTextString(m) }
func (*AuthDB) ProtoMessage()               {}
func (*AuthDB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AuthDB) GetOauthClientId() string {
	if m != nil && m.OauthClientId != nil {
		return *m.OauthClientId
	}
	return ""
}

func (m *AuthDB) GetOauthClientSecret() string {
	if m != nil && m.OauthClientSecret != nil {
		return *m.OauthClientSecret
	}
	return ""
}

func (m *AuthDB) GetOauthAdditionalClientIds() []string {
	if m != nil {
		return m.OauthAdditionalClientIds
	}
	return nil
}

func (m *AuthDB) GetGroups() []*AuthGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *AuthDB) GetSecrets() []*AuthSecret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *AuthDB) GetIpWhitelists() []*AuthIPWhitelist {
	if m != nil {
		return m.IpWhitelists
	}
	return nil
}

func (m *AuthDB) GetIpWhitelistAssignments() []*AuthIPWhitelistAssignment {
	if m != nil {
		return m.IpWhitelistAssignments
	}
	return nil
}

// Information about some particular revision of auth DB.
type AuthDBRevision struct {
	// GAE App ID of a service holding primary copy of Auth DB.
	PrimaryId *string `protobuf:"bytes,1,req,name=primary_id" json:"primary_id,omitempty"`
	// Revision of Auth DB being pushed.
	AuthDbRev *int64 `protobuf:"varint,2,req,name=auth_db_rev" json:"auth_db_rev,omitempty"`
	// Timestamp of that revision by Primary's clock, microseconds since epoch.
	ModifiedTs       *int64 `protobuf:"varint,3,req,name=modified_ts" json:"modified_ts,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AuthDBRevision) Reset()                    { *m = AuthDBRevision{} }
func (m *AuthDBRevision) String() string            { return proto.CompactTextString(m) }
func (*AuthDBRevision) ProtoMessage()               {}
func (*AuthDBRevision) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AuthDBRevision) GetPrimaryId() string {
	if m != nil && m.PrimaryId != nil {
		return *m.PrimaryId
	}
	return ""
}

func (m *AuthDBRevision) GetAuthDbRev() int64 {
	if m != nil && m.AuthDbRev != nil {
		return *m.AuthDbRev
	}
	return 0
}

func (m *AuthDBRevision) GetModifiedTs() int64 {
	if m != nil && m.ModifiedTs != nil {
		return *m.ModifiedTs
	}
	return 0
}

// Published by Primary into 'auth-db-changed' PubSub topic. The body of the
// message is base64 encoded serialized ChangeNotification. Additional
// attributes are:
//  X-AuthDB-SigKey-v1: <id of a public key>
//  X-AuthDB-SigVal-v1: <base64 encoded RSA-SHA256(blob) signature>
type ChangeNotification struct {
	// New revision of the AuthDB.
	Revision         *AuthDBRevision `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ChangeNotification) Reset()                    { *m = ChangeNotification{} }
func (m *ChangeNotification) String() string            { return proto.CompactTextString(m) }
func (*ChangeNotification) ProtoMessage()               {}
func (*ChangeNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ChangeNotification) GetRevision() *AuthDBRevision {
	if m != nil {
		return m.Revision
	}
	return nil
}

// Sent from Primary to Replica to update Replica's AuthDB.
// Primary signs the entire serialized message with its private key and appends
// two headers to HTTP request that carries the blob:
//  X-AuthDB-SigKey-v1: <id of a public key>
//  X-AuthDB-SigVal-v1: <base64 encoded RSA-SHA256(SHA512(blob)) signature>
type ReplicationPushRequest struct {
	// Revision that is being pushed.
	Revision *AuthDBRevision `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
	// An entire database of auth configuration for specific revision.
	AuthDb *AuthDB `protobuf:"bytes,2,opt,name=auth_db" json:"auth_db,omitempty"`
	// Version of 'auth' component on Primary, see components/auth/version.py.
	AuthCodeVersion  *string `protobuf:"bytes,3,opt,name=auth_code_version" json:"auth_code_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReplicationPushRequest) Reset()                    { *m = ReplicationPushRequest{} }
func (m *ReplicationPushRequest) String() string            { return proto.CompactTextString(m) }
func (*ReplicationPushRequest) ProtoMessage()               {}
func (*ReplicationPushRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReplicationPushRequest) GetRevision() *AuthDBRevision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *ReplicationPushRequest) GetAuthDb() *AuthDB {
	if m != nil {
		return m.AuthDb
	}
	return nil
}

func (m *ReplicationPushRequest) GetAuthCodeVersion() string {
	if m != nil && m.AuthCodeVersion != nil {
		return *m.AuthCodeVersion
	}
	return ""
}

// Replica's response to ReplicationPushRequest.
type ReplicationPushResponse struct {
	// Overall status of the operation.
	Status *ReplicationPushResponse_Status `protobuf:"varint,1,req,name=status,enum=protocol.ReplicationPushResponse_Status" json:"status,omitempty"`
	// Revision known by Replica (set for APPLIED and SKIPPED statuses).
	CurrentRevision *AuthDBRevision `protobuf:"bytes,2,opt,name=current_revision" json:"current_revision,omitempty"`
	// Present for TRANSIENT_ERROR and FATAL_ERROR statuses.
	ErrorCode *ReplicationPushResponse_ErrorCode `protobuf:"varint,3,opt,name=error_code,enum=protocol.ReplicationPushResponse_ErrorCode" json:"error_code,omitempty"`
	// Version of 'auth' component on Replica, see components/auth/version.py.
	AuthCodeVersion  *string `protobuf:"bytes,4,opt,name=auth_code_version" json:"auth_code_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReplicationPushResponse) Reset()                    { *m = ReplicationPushResponse{} }
func (m *ReplicationPushResponse) String() string            { return proto.CompactTextString(m) }
func (*ReplicationPushResponse) ProtoMessage()               {}
func (*ReplicationPushResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReplicationPushResponse) GetStatus() ReplicationPushResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ReplicationPushResponse_APPLIED
}

func (m *ReplicationPushResponse) GetCurrentRevision() *AuthDBRevision {
	if m != nil {
		return m.CurrentRevision
	}
	return nil
}

func (m *ReplicationPushResponse) GetErrorCode() ReplicationPushResponse_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return ReplicationPushResponse_NOT_A_REPLICA
}

func (m *ReplicationPushResponse) GetAuthCodeVersion() string {
	if m != nil && m.AuthCodeVersion != nil {
		return *m.AuthCodeVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceLinkTicket)(nil), "protocol.ServiceLinkTicket")
	proto.RegisterType((*ServiceLinkRequest)(nil), "protocol.ServiceLinkRequest")
	proto.RegisterType((*ServiceLinkResponse)(nil), "protocol.ServiceLinkResponse")
	proto.RegisterType((*AuthGroup)(nil), "protocol.AuthGroup")
	proto.RegisterType((*AuthSecret)(nil), "protocol.AuthSecret")
	proto.RegisterType((*AuthIPWhitelist)(nil), "protocol.AuthIPWhitelist")
	proto.RegisterType((*AuthIPWhitelistAssignment)(nil), "protocol.AuthIPWhitelistAssignment")
	proto.RegisterType((*AuthDB)(nil), "protocol.AuthDB")
	proto.RegisterType((*AuthDBRevision)(nil), "protocol.AuthDBRevision")
	proto.RegisterType((*ChangeNotification)(nil), "protocol.ChangeNotification")
	proto.RegisterType((*ReplicationPushRequest)(nil), "protocol.ReplicationPushRequest")
	proto.RegisterType((*ReplicationPushResponse)(nil), "protocol.ReplicationPushResponse")
	proto.RegisterEnum("protocol.ServiceLinkResponse_Status", ServiceLinkResponse_Status_name, ServiceLinkResponse_Status_value)
	proto.RegisterEnum("protocol.ReplicationPushResponse_Status", ReplicationPushResponse_Status_name, ReplicationPushResponse_Status_value)
	proto.RegisterEnum("protocol.ReplicationPushResponse_ErrorCode", ReplicationPushResponse_ErrorCode_name, ReplicationPushResponse_ErrorCode_value)
}

var fileDescriptor0 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xae, 0x44, 0x59, 0xb6, 0xc6, 0xb2, 0x2d, 0xd1, 0x69, 0x23, 0x23, 0x97, 0x94, 0x6e, 0x81,
	0xa0, 0x05, 0x8c, 0xc2, 0xe8, 0xa1, 0xb7, 0x96, 0x96, 0x18, 0x87, 0x8d, 0x2b, 0xa9, 0x24, 0x8d,
	0x1e, 0x09, 0x8a, 0xdc, 0x4a, 0x8b, 0x50, 0xa4, 0xba, 0xbb, 0x54, 0x60, 0xa0, 0xa7, 0xbe, 0x44,
	0x5e, 0xa1, 0x8f, 0xd1, 0x47, 0xeb, 0xec, 0x92, 0x34, 0x29, 0x85, 0x46, 0x90, 0x93, 0xb8, 0xb3,
	0xf3, 0xf3, 0x7d, 0xdf, 0xcc, 0x8e, 0x60, 0xc8, 0xc8, 0x26, 0xa6, 0x61, 0x20, 0x68, 0x9a, 0x5c,
	0x6d, 0x58, 0x2a, 0x52, 0xfd, 0x48, 0xfd, 0x84, 0x69, 0x6c, 0x2c, 0x60, 0xe8, 0x12, 0xb6, 0xa5,
	0x21, 0xb9, 0xa3, 0xc9, 0x3b, 0x8f, 0x86, 0xef, 0x88, 0xd0, 0x75, 0x80, 0x0d, 0xa3, 0xeb, 0x80,
	0x3d, 0xf8, 0x34, 0x1a, 0xb5, 0x5e, 0xb6, 0x5f, 0xf5, 0xf4, 0x73, 0x38, 0x2e, 0x6d, 0x19, 0x8b,
	0x47, 0x6d, 0x65, 0x7c, 0x06, 0xfd, 0x25, 0x49, 0x08, 0x0b, 0x04, 0x89, 0xfc, 0xc5, 0xc3, 0x48,
	0x53, 0xd6, 0x53, 0xe8, 0x0a, 0x95, 0x68, 0xd4, 0xc1, 0x73, 0xdf, 0x98, 0x81, 0x5e, 0xab, 0xe1,
	0x90, 0xbf, 0x32, 0xc2, 0x45, 0xcd, 0x4b, 0x16, 0xe8, 0xcb, 0x02, 0x05, 0xd0, 0xdd, 0x02, 0x34,
	0xa1, 0x82, 0xee, 0x14, 0x30, 0x3e, 0xb4, 0xe0, 0x7c, 0x27, 0x23, 0xdf, 0xa4, 0x09, 0x27, 0xfa,
	0x8f, 0xd0, 0xe5, 0x22, 0x10, 0x19, 0x57, 0x29, 0x4f, 0xaf, 0xbf, 0xb9, 0x2a, 0x79, 0x5e, 0x35,
	0xb8, 0x5f, 0xb9, 0xca, 0xd7, 0xf8, 0x15, 0xba, 0xf9, 0x97, 0x7e, 0x0c, 0x87, 0xee, 0xfd, 0x78,
	0x6c, 0xb9, 0xee, 0xe0, 0x0b, 0xc4, 0x73, 0xe6, 0x39, 0xe6, 0xd4, 0x9d, 0xcf, 0x1c, 0xcf, 0xb7,
	0x1c, 0x67, 0xe6, 0x0c, 0x5a, 0x08, 0x1a, 0x6e, 0xcc, 0x89, 0xef, 0xd9, 0xe3, 0xb7, 0x96, 0x37,
	0x68, 0xcb, 0xb3, 0x79, 0xef, 0xbd, 0x29, 0xee, 0x35, 0xe3, 0xbf, 0x16, 0xf4, 0xcc, 0x4c, 0xac,
	0x6e, 0x59, 0x9a, 0x6d, 0xf4, 0x3e, 0x74, 0x92, 0x60, 0x4d, 0x0a, 0x05, 0xcf, 0xe0, 0x70, 0x4d,
	0xd6, 0x0b, 0xc2, 0x38, 0x92, 0xd3, 0xd0, 0x70, 0x02, 0x07, 0xcb, 0x38, 0x5d, 0x70, 0x64, 0xa5,
	0xe5, 0xb2, 0x25, 0x28, 0x0c, 0x89, 0x50, 0x36, 0x2d, 0x57, 0x3c, 0x22, 0x3c, 0x64, 0x74, 0x23,
	0x3b, 0x37, 0x3a, 0x50, 0x49, 0xb0, 0x35, 0x21, 0x23, 0x4a, 0x0e, 0xc1, 0x47, 0x5d, 0xb4, 0x69,
	0x75, 0x1b, 0x4a, 0x74, 0x58, 0xb6, 0x6b, 0x9d, 0x46, 0xf4, 0x4f, 0x9a, 0x3b, 0x1e, 0x29, 0xc7,
	0xba, 0x11, 0x3d, 0x7b, 0x65, 0xb7, 0xd2, 0xf7, 0x89, 0x44, 0x05, 0x2f, 0x5b, 0x28, 0xae, 0x87,
	0x94, 0x90, 0x81, 0x4b, 0x30, 0xa7, 0xd8, 0xa3, 0x80, 0xbe, 0xdb, 0x20, 0xc6, 0xf6, 0x29, 0x06,
	0xfd, 0xfd, 0x2a, 0x5a, 0x53, 0x95, 0x4e, 0xd9, 0xb2, 0x33, 0x99, 0xd6, 0x9e, 0xff, 0xb1, 0xa2,
	0x82, 0xc4, 0x94, 0x8b, 0x8f, 0xe5, 0xe1, 0xd9, 0x22, 0x21, 0xa2, 0x94, 0x67, 0x8f, 0xbf, 0xd6,
	0xc0, 0xbf, 0xd3, 0xc0, 0xff, 0xa0, 0x89, 0x7f, 0xb7, 0x09, 0x99, 0x52, 0xca, 0xf8, 0x1b, 0x2e,
	0xf6, 0x80, 0x99, 0x9c, 0xd3, 0x65, 0xb2, 0x26, 0x89, 0xd0, 0x07, 0x70, 0x44, 0x23, 0xfc, 0xa0,
	0xe2, 0xa1, 0x80, 0x29, 0x27, 0x72, 0xe3, 0xbf, 0x2f, 0x7d, 0x8b, 0x39, 0x45, 0xf0, 0x61, 0xba,
	0x96, 0x21, 0x9f, 0x87, 0xd3, 0xf8, 0xb7, 0x0d, 0x5d, 0x59, 0x7e, 0x72, 0xa3, 0x3f, 0x87, 0xb3,
	0x34, 0xc0, 0x4f, 0x3f, 0x8c, 0x29, 0x26, 0xaa, 0x9e, 0xde, 0x0b, 0x38, 0xdf, 0xb9, 0xe0, 0xaa,
	0x35, 0x45, 0xe5, 0x4b, 0x78, 0x91, 0x5f, 0x06, 0x51, 0x44, 0xa5, 0x50, 0x41, 0x5c, 0x25, 0x28,
	0x47, 0xeb, 0x12, 0xba, 0x4b, 0x39, 0x91, 0x5c, 0x8d, 0xd6, 0xf1, 0xf5, 0x79, 0xf5, 0x30, 0xaa,
	0x69, 0xfd, 0x16, 0x1b, 0xa0, 0x32, 0x73, 0xc4, 0x26, 0xbd, 0x9e, 0xed, 0x7a, 0x15, 0x13, 0xf1,
	0x03, 0x9c, 0xd4, 0x05, 0x90, 0xda, 0x4a, 0xe7, 0x8b, 0x5d, 0xe7, 0x7a, 0x9f, 0x2d, 0x18, 0xd5,
	0x23, 0xfc, 0xe0, 0x51, 0x5f, 0x8e, 0x3d, 0x90, 0xc1, 0x97, 0x4f, 0x06, 0x57, 0xbd, 0x30, 0xa6,
	0x70, 0x9a, 0x2b, 0xe5, 0x90, 0x2d, 0xe5, 0xc8, 0xf4, 0xa9, 0x3d, 0xa5, 0xe4, 0x88, 0x16, 0x3e,
	0x23, 0x5b, 0x25, 0x92, 0xd6, 0x38, 0xa7, 0xc6, 0x2f, 0xa0, 0x8f, 0x57, 0x41, 0xb2, 0x24, 0xd3,
	0x54, 0xe0, 0x55, 0xbe, 0x20, 0xf5, 0xef, 0xe0, 0x88, 0x15, 0xf9, 0x31, 0x63, 0x0b, 0xc1, 0x8d,
	0x76, 0xc1, 0x55, 0xf5, 0x8d, 0x7f, 0x5a, 0xf0, 0x95, 0x53, 0x2d, 0xd7, 0x79, 0xc6, 0x57, 0xe5,
	0x76, 0xfb, 0x8c, 0x34, 0xfa, 0xd7, 0x70, 0x58, 0x40, 0x46, 0xb8, 0xd2, 0x75, 0xb0, 0xef, 0xaa,
	0x5f, 0xc0, 0x30, 0x9f, 0x80, 0x34, 0x22, 0xfe, 0x16, 0x1f, 0x6b, 0xfe, 0x22, 0xe4, 0x7b, 0xfd,
	0xa0, 0xc1, 0xf3, 0x8f, 0x40, 0x14, 0x0b, 0xf1, 0xa7, 0xbd, 0x85, 0xf8, 0xaa, 0x4a, 0xfc, 0x44,
	0x48, 0xb1, 0x14, 0xf5, 0x6b, 0x18, 0x84, 0x19, 0x63, 0x72, 0x8c, 0x1e, 0x79, 0xb4, 0x3f, 0xc1,
	0xe3, 0x67, 0x00, 0xc2, 0x58, 0xca, 0x14, 0x4a, 0x85, 0xee, 0xf4, 0xfa, 0xfb, 0x4f, 0x57, 0xb4,
	0x64, 0xcc, 0x18, 0x43, 0x9a, 0x59, 0x76, 0x14, 0xcb, 0x37, 0xf5, 0x25, 0x6d, 0xce, 0xe7, 0x77,
	0xb6, 0x35, 0xc1, 0x25, 0x2d, 0x37, 0xf6, 0x5b, 0x7b, 0x3e, 0xc7, 0x43, 0xeb, 0x71, 0x63, 0xdb,
	0xd6, 0xb4, 0xdc, 0xd8, 0x6d, 0x7c, 0x99, 0xc7, 0xaf, 0x4d, 0xcf, 0xbc, 0x7b, 0x5c, 0xd1, 0x2b,
	0xe8, 0x55, 0x15, 0x87, 0x70, 0x32, 0x9d, 0x79, 0xbe, 0xe9, 0x3b, 0x16, 0xa6, 0x1c, 0x9b, 0x98,
	0xe5, 0x04, 0x7a, 0xaf, 0x67, 0xce, 0x8d, 0x3d, 0x99, 0x58, 0x53, 0x8c, 0xff, 0x12, 0x86, 0xbf,
	0xd9, 0xae, 0x6b, 0x4f, 0x6f, 0x7d, 0xd7, 0xbe, 0x9d, 0x9a, 0xde, 0xbd, 0x63, 0x0d, 0x34, 0x19,
	0x28, 0xff, 0x08, 0x2a, 0x53, 0x47, 0x56, 0x92, 0x26, 0xc7, 0xfa, 0xfd, 0xde, 0x72, 0xbd, 0xc1,
	0xc1, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x40, 0x66, 0xe1, 0x79, 0x07, 0x00, 0x00,
}
