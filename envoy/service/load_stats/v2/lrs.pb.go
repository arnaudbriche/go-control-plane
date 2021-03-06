// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/load_stats/v2/lrs.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
import types "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A load report Envoy sends to the management server.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type LoadStatsRequest struct {
	// Node identifier for Envoy instance.
	Node *core.Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	// A list of load stats to report.
	ClusterStats         []*endpoint.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lrs_7926936dce4e8a94, []int{0}
}
func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(dst, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return m.Size()
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*endpoint.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

// The management server sends envoy a LoadStatsResponse with all clusters it
// is interested in learning load stats about.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type LoadStatsResponse struct {
	// Clusters to report stats for.
	Clusters []string `protobuf:"bytes,1,rep,name=clusters" json:"clusters,omitempty"`
	// The minimum interval of time to collect stats over. This is only a minimum for two reasons:
	// 1. There may be some delay from when the timer fires until stats sampling occurs.
	// 2. For clusters that were already feature in the previous *LoadStatsResponse*, any traffic
	//    that is observed in between the corresponding previous *LoadStatsRequest* and this
	//    *LoadStatsResponse* will also be accumulated and billed to the cluster. This avoids a period
	//    of inobservability that might otherwise exists between the messages. New clusters are not
	//    subject to this consideration.
	LoadReportingInterval *types.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval" json:"load_reporting_interval,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}        `json:"-"`
	XXX_unrecognized      []byte          `json:"-"`
	XXX_sizecache         int32           `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lrs_7926936dce4e8a94, []int{1}
}
func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(dst, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return m.Size()
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *types.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v2.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v2.LoadStatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoadReportingService service

type LoadReportingServiceClient interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//    capacity configuration. The capacity configuration definition is
	//    outside of the scope of this document.
	// 2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//    to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	// 1. Once a connection establishes, the management server publishes a
	//    LoadStatsResponse for all clusters it is interested in learning load
	//    stats about.
	// 2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//    based on per-zone weights and/or per-instance weights (if specified)
	//    based on intra-zone LbPolicy. This information comes from the above
	//    {Stream,Fetch}Endpoints.
	// 3. When upstream hosts reply, they optionally add header <define header
	//    name> with ASCII representation of EndpointLoadMetricStats.
	// 4. Envoy aggregates load reports over the period of time given to it in
	//    LoadStatsResponse.load_reporting_interval. This includes aggregation
	//    stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//    well as load metrics from upstream hosts.
	// 5. When the timer of load_reporting_interval expires, Envoy sends new
	//    LoadStatsRequest filled with load reports for each cluster.
	// 6. The management server uses the load reports from all reported Envoys
	//    from around the world, computes global assignment and prepares traffic
	//    assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v2.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for LoadReportingService service

type LoadReportingServiceServer interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//    capacity configuration. The capacity configuration definition is
	//    outside of the scope of this document.
	// 2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//    to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	// 1. Once a connection establishes, the management server publishes a
	//    LoadStatsResponse for all clusters it is interested in learning load
	//    stats about.
	// 2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//    based on per-zone weights and/or per-instance weights (if specified)
	//    based on intra-zone LbPolicy. This information comes from the above
	//    {Stream,Fetch}Endpoints.
	// 3. When upstream hosts reply, they optionally add header <define header
	//    name> with ASCII representation of EndpointLoadMetricStats.
	// 4. Envoy aggregates load reports over the period of time given to it in
	//    LoadStatsResponse.load_reporting_interval. This includes aggregation
	//    stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//    well as load metrics from upstream hosts.
	// 5. When the timer of load_reporting_interval expires, Envoy sends new
	//    LoadStatsRequest filled with load reports for each cluster.
	// 6. The management server uses the load reports from all reported Envoys
	//    from around the world, computes global assignment and prepares traffic
	//    assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v2.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v2/lrs.proto",
}

func (m *LoadStatsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadStatsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLrs(dAtA, i, uint64(m.Node.Size()))
		n1, err := m.Node.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.ClusterStats) > 0 {
		for _, msg := range m.ClusterStats {
			dAtA[i] = 0x12
			i++
			i = encodeVarintLrs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LoadStatsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadStatsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Clusters) > 0 {
		for _, s := range m.Clusters {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.LoadReportingInterval != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLrs(dAtA, i, uint64(m.LoadReportingInterval.Size()))
		n2, err := m.LoadReportingInterval.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintLrs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LoadStatsRequest) Size() (n int) {
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovLrs(uint64(l))
	}
	if len(m.ClusterStats) > 0 {
		for _, e := range m.ClusterStats {
			l = e.Size()
			n += 1 + l + sovLrs(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoadStatsResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Clusters) > 0 {
		for _, s := range m.Clusters {
			l = len(s)
			n += 1 + l + sovLrs(uint64(l))
		}
	}
	if m.LoadReportingInterval != nil {
		l = m.LoadReportingInterval.Size()
		n += 1 + l + sovLrs(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLrs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLrs(x uint64) (n int) {
	return sovLrs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoadStatsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadStatsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadStatsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &core.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterStats = append(m.ClusterStats, &endpoint.ClusterStats{})
			if err := m.ClusterStats[len(m.ClusterStats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLrs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoadStatsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadStatsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadStatsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLrs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadReportingInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LoadReportingInterval == nil {
				m.LoadReportingInterval = &types.Duration{}
			}
			if err := m.LoadReportingInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLrs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLrs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLrs
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
					return 0, ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLrs
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLrs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLrs
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLrs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLrs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLrs   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/service/load_stats/v2/lrs.proto", fileDescriptor_lrs_7926936dce4e8a94)
}

var fileDescriptor_lrs_7926936dce4e8a94 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x71, 0x0a, 0xa8, 0xb8, 0x20, 0x20, 0x80, 0x1a, 0x0a, 0x8a, 0xaa, 0x22, 0x20, 0x12,
	0xc2, 0x46, 0xe1, 0x0d, 0x0a, 0x07, 0x90, 0x2a, 0x24, 0xd2, 0xdb, 0x2e, 0x95, 0x9b, 0xfc, 0x57,
	0x59, 0xca, 0xec, 0xcc, 0x76, 0x2c, 0xed, 0x0d, 0xb6, 0x4b, 0x0f, 0x7b, 0x9c, 0x9d, 0x76, 0xdc,
	0x71, 0x8f, 0x30, 0xf5, 0xb6, 0xb7, 0x98, 0x12, 0x27, 0x5d, 0xba, 0xc3, 0xb4, 0x5b, 0xac, 0xef,
	0xf7, 0xc5, 0xdf, 0xf7, 0x19, 0x7f, 0x06, 0x61, 0xe5, 0x11, 0xd5, 0xa0, 0x2c, 0x4f, 0x81, 0xe6,
	0x92, 0x65, 0x0b, 0x6d, 0x98, 0xd1, 0xd4, 0xc6, 0x34, 0x57, 0x9a, 0x14, 0x4a, 0x1a, 0xe9, 0x7f,
	0xa8, 0x31, 0xd2, 0x60, 0xe4, 0x16, 0x23, 0x36, 0x1e, 0x7d, 0x74, 0xff, 0x60, 0x05, 0xaf, 0x4c,
	0xa9, 0x54, 0x40, 0x97, 0x4c, 0x83, 0xb3, 0x8e, 0xbe, 0xee, 0xa8, 0x20, 0xb2, 0x42, 0x72, 0x61,
	0xdc, 0x4d, 0x0a, 0x0a, 0xa9, 0x4c, 0x03, 0x86, 0x2b, 0x29, 0x57, 0x39, 0xd0, 0xfa, 0xb4, 0x2c,
	0xf7, 0x69, 0x56, 0x2a, 0x66, 0xb8, 0x14, 0x8d, 0x3e, 0xb4, 0x2c, 0xe7, 0x19, 0x33, 0x40, 0xdb,
	0x0f, 0x27, 0x4c, 0x4e, 0x10, 0x7e, 0x35, 0x93, 0x2c, 0x9b, 0x57, 0x81, 0x12, 0x38, 0x2c, 0x41,
	0x1b, 0xff, 0x1b, 0x7e, 0x2c, 0x64, 0x06, 0x01, 0x1a, 0xa3, 0x68, 0x10, 0x0f, 0x89, 0x2b, 0xc0,
	0x0a, 0x4e, 0x6c, 0x4c, 0xaa, 0x8c, 0xe4, 0x9f, 0xcc, 0x20, 0xa9, 0x21, 0xff, 0x0f, 0x7e, 0x91,
	0xe6, 0xa5, 0x36, 0xa0, 0x5c, 0xab, 0xc0, 0x1b, 0xf7, 0xa2, 0x41, 0xfc, 0x69, 0xd7, 0xd5, 0x66,
	0x27, 0xbf, 0x1c, 0xeb, 0xee, 0x7b, 0x9e, 0x76, 0x4e, 0x93, 0x35, 0xc2, 0xaf, 0x3b, 0x59, 0x74,
	0x21, 0x85, 0x06, 0xff, 0x0b, 0xee, 0x37, 0x94, 0x0e, 0xd0, 0xb8, 0x17, 0x3d, 0x9b, 0xe2, 0xb3,
	0xeb, 0xf3, 0xde, 0x93, 0x53, 0xe4, 0xf5, 0x51, 0xb2, 0xd5, 0xfc, 0xff, 0x78, 0xd8, 0xd9, 0x85,
	0x8b, 0xd5, 0x82, 0x0b, 0x03, 0xca, 0xb2, 0x3c, 0xf0, 0xea, 0x1e, 0xef, 0x89, 0x1b, 0x89, 0xb4,
	0x23, 0x91, 0xdf, 0xcd, 0x48, 0xc9, 0xbb, 0xca, 0x99, 0xb4, 0xc6, 0xbf, 0x8d, 0x2f, 0x5e, 0x23,
	0xfc, 0x76, 0xd6, 0x55, 0xe6, 0xee, 0x0d, 0x7d, 0x8b, 0x5f, 0xce, 0x8d, 0x02, 0x76, 0xb0, 0x8d,
	0xeb, 0x7f, 0x27, 0xf7, 0x3c, 0x33, 0xb9, 0x3b, 0xf1, 0x88, 0x3c, 0x14, 0x77, 0x2b, 0x4c, 0x1e,
	0x45, 0xe8, 0x07, 0x9a, 0xbe, 0xb9, 0xd8, 0x84, 0xe8, 0x72, 0x13, 0xa2, 0xab, 0x4d, 0x88, 0xf6,
	0x3c, 0x1b, 0x1f, 0x23, 0xb4, 0x7c, 0x5a, 0xf7, 0xf9, 0x79, 0x13, 0x00, 0x00, 0xff, 0xff, 0x84,
	0x9e, 0x59, 0xb4, 0x8f, 0x02, 0x00, 0x00,
}
