package increment

import (
	"errors"
	"golang.org/x/net/context"
	"sync/atomic"
)

var increments map[string]*uint64

func init() {
	increments = make(map[string]*uint64)
}

type IncrServer struct{}

func (serv *IncrServer) GetIncrId(ctx context.Context, in *GetIncrIdRequest) (*GetIncrIdReply, error) {
	v, ok := increments[in.Name]
	if !ok {
		return &GetIncrIdReply{}, errors.New("Key Not Exist")
	}
	nId := atomic.AddUint64(v, 1)
	return &GetIncrIdReply{IncId: nId}, nil
}

func (serv *IncrServer) CheckIncrKeyExist(ctx context.Context, in *GetIncrIdRequest) (*CheckIncrKeyExistReply, error) {
	_, ok := increments[in.Name]
	if !ok {
		return &CheckIncrKeyExistReply{Exist: false}, nil
	}
	return &CheckIncrKeyExistReply{Exist: true}, nil
}

func (serv *IncrServer) CreateIncrKey(ctx context.Context, in *CreateIncrKeyRequest) (*NoneReply, error) {
	// 其实此方法应加读写锁，前2个方法属于读方法，此方法属于写方法
	_, ok := increments[in.Name]
	if ok {
		return &NoneReply{}, errors.New("Key Already Exist")
	}
	nId := in.InitialValue
	increments[in.Name] = &nId
	return &NoneReply{}, nil
}
