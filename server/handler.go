package main

import (
	"context"
	"fmt"
	"sync"

	pb "dice/proto"

	"google.golang.org/grpc/stats"
)

type connCtxKey struct{}
type rpcCtxKey struct{}

type gotData struct {
	ctx    context.Context
	client bool
	s      any
}

// 참고: https://github.com/grpc/grpc-go/blob/2c00469782f1dd8c7456dcc7238a957781246e84/stats/stats_test.go#L785
type MyHandler struct {
	mu      sync.Mutex
	gotRPC  []*gotData
	gotConn []*gotData
}

func (h *MyHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return context.WithValue(ctx, connCtxKey{}, info)
}

func (h *MyHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return context.WithValue(ctx, rpcCtxKey{}, info)
}

// HandleConn processes the Conn stats.
func (h *MyHandler) HandleConn(ctx context.Context, s stats.ConnStats) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.gotConn = append(h.gotConn, &gotData{ctx, s.IsClient(), s})
}

// https://pkg.go.dev/google.golang.org/grpc/stats#OutPayload
// payload: 압축까지 진행된 데이터
// Length: 압축되기 전의 데이터 길이. framing 데이터 포함 안함
// CompressedLength: 압축 후의 payload 데이터 길이. framing 데이터 포함 안함. 압축 안할 시 length와 동일
// WireLength: framing 데이터 포함한 전체 길이.
func (h *MyHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {
	// h.mu.Lock()
	// defer h.mu.Unlock()
	// h.gotRPC = append(h.gotRPC, &gotData{ctx, s.IsClient(), s})
	switch ss := s.(type) {
	case *stats.InPayload:
		fmt.Println(ss.CompressedLength)
		switch k := ss.Payload.(type) {
		case *pb.GreetRequest:
			fmt.Println(k.GetName())
			fmt.Println(k.GetMsg())
			fmt.Println(k.GetAge())
			v := ctx.Value(rpcCtxKey{})
			switch v.(type) {
			case *stats.RPCTagInfo:
				fmt.Println(v)
			}
		}
	}
}
