// Code generated by sdkgen. DO NOT EDIT.

//nolint
package dataproc

import (
	"context"

	"google.golang.org/grpc"

	dataproc "github.com/yandex-cloud/go-genproto/yandex/cloud/dataproc/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// SubclusterServiceClient is a dataproc.SubclusterServiceClient with
// lazy GRPC connection initialization.
type SubclusterServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements dataproc.SubclusterServiceClient
func (c *SubclusterServiceClient) Create(ctx context.Context, in *dataproc.CreateSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return dataproc.NewSubclusterServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements dataproc.SubclusterServiceClient
func (c *SubclusterServiceClient) Delete(ctx context.Context, in *dataproc.DeleteSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return dataproc.NewSubclusterServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements dataproc.SubclusterServiceClient
func (c *SubclusterServiceClient) Get(ctx context.Context, in *dataproc.GetSubclusterRequest, opts ...grpc.CallOption) (*dataproc.Subcluster, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return dataproc.NewSubclusterServiceClient(conn).Get(ctx, in, opts...)
}

// List implements dataproc.SubclusterServiceClient
func (c *SubclusterServiceClient) List(ctx context.Context, in *dataproc.ListSubclustersRequest, opts ...grpc.CallOption) (*dataproc.ListSubclustersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return dataproc.NewSubclusterServiceClient(conn).List(ctx, in, opts...)
}

type SubclusterIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SubclusterServiceClient
	request *dataproc.ListSubclustersRequest

	items []*dataproc.Subcluster
}

func (c *SubclusterServiceClient) SubclusterIterator(ctx context.Context, req *dataproc.ListSubclustersRequest, opts ...grpc.CallOption) *SubclusterIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SubclusterIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SubclusterIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Subclusters
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SubclusterIterator) Take(size int64) ([]*dataproc.Subcluster, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*dataproc.Subcluster

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *SubclusterIterator) TakeAll() ([]*dataproc.Subcluster, error) {
	return it.Take(0)
}

func (it *SubclusterIterator) Value() *dataproc.Subcluster {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SubclusterIterator) Error() error {
	return it.err
}

// Update implements dataproc.SubclusterServiceClient
func (c *SubclusterServiceClient) Update(ctx context.Context, in *dataproc.UpdateSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return dataproc.NewSubclusterServiceClient(conn).Update(ctx, in, opts...)
}
