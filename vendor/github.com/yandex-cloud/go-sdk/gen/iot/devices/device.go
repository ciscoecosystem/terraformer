// Code generated by sdkgen. DO NOT EDIT.

//nolint
package devices

import (
	"context"

	"google.golang.org/grpc"

	devices "github.com/yandex-cloud/go-genproto/yandex/cloud/iot/devices/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// DeviceServiceClient is a devices.DeviceServiceClient with
// lazy GRPC connection initialization.
type DeviceServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddCertificate implements devices.DeviceServiceClient
func (c *DeviceServiceClient) AddCertificate(ctx context.Context, in *devices.AddDeviceCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).AddCertificate(ctx, in, opts...)
}

// AddPassword implements devices.DeviceServiceClient
func (c *DeviceServiceClient) AddPassword(ctx context.Context, in *devices.AddDevicePasswordRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).AddPassword(ctx, in, opts...)
}

// Create implements devices.DeviceServiceClient
func (c *DeviceServiceClient) Create(ctx context.Context, in *devices.CreateDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements devices.DeviceServiceClient
func (c *DeviceServiceClient) Delete(ctx context.Context, in *devices.DeleteDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).Delete(ctx, in, opts...)
}

// DeleteCertificate implements devices.DeviceServiceClient
func (c *DeviceServiceClient) DeleteCertificate(ctx context.Context, in *devices.DeleteDeviceCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).DeleteCertificate(ctx, in, opts...)
}

// DeletePassword implements devices.DeviceServiceClient
func (c *DeviceServiceClient) DeletePassword(ctx context.Context, in *devices.DeleteDevicePasswordRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).DeletePassword(ctx, in, opts...)
}

// Get implements devices.DeviceServiceClient
func (c *DeviceServiceClient) Get(ctx context.Context, in *devices.GetDeviceRequest, opts ...grpc.CallOption) (*devices.Device, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).Get(ctx, in, opts...)
}

// List implements devices.DeviceServiceClient
func (c *DeviceServiceClient) List(ctx context.Context, in *devices.ListDevicesRequest, opts ...grpc.CallOption) (*devices.ListDevicesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).List(ctx, in, opts...)
}

// ListCertificates implements devices.DeviceServiceClient
func (c *DeviceServiceClient) ListCertificates(ctx context.Context, in *devices.ListDeviceCertificatesRequest, opts ...grpc.CallOption) (*devices.ListDeviceCertificatesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).ListCertificates(ctx, in, opts...)
}

type DeviceCertificatesIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *DeviceServiceClient
	request *devices.ListDeviceCertificatesRequest

	items []*devices.DeviceCertificate
}

func (c *DeviceServiceClient) DeviceCertificatesIterator(ctx context.Context, deviceId string, opts ...grpc.CallOption) *DeviceCertificatesIterator {
	return &DeviceCertificatesIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &devices.ListDeviceCertificatesRequest{
			DeviceId: deviceId,
		},
	}
}

func (it *DeviceCertificatesIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started {
		return false
	}
	it.started = true

	response, err := it.client.ListCertificates(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Certificates
	return len(it.items) > 0
}

func (it *DeviceCertificatesIterator) Value() *devices.DeviceCertificate {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *DeviceCertificatesIterator) Error() error {
	return it.err
}

// ListOperations implements devices.DeviceServiceClient
func (c *DeviceServiceClient) ListOperations(ctx context.Context, in *devices.ListDeviceOperationsRequest, opts ...grpc.CallOption) (*devices.ListDeviceOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).ListOperations(ctx, in, opts...)
}

type DeviceOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *DeviceServiceClient
	request *devices.ListDeviceOperationsRequest

	items []*operation.Operation
}

func (c *DeviceServiceClient) DeviceOperationsIterator(ctx context.Context, deviceId string, opts ...grpc.CallOption) *DeviceOperationsIterator {
	return &DeviceOperationsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &devices.ListDeviceOperationsRequest{
			DeviceId: deviceId,
			PageSize: 1000,
		},
	}
}

func (it *DeviceOperationsIterator) Next() bool {
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

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *DeviceOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *DeviceOperationsIterator) Error() error {
	return it.err
}

// ListPasswords implements devices.DeviceServiceClient
func (c *DeviceServiceClient) ListPasswords(ctx context.Context, in *devices.ListDevicePasswordsRequest, opts ...grpc.CallOption) (*devices.ListDevicePasswordsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).ListPasswords(ctx, in, opts...)
}

type DevicePasswordsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *DeviceServiceClient
	request *devices.ListDevicePasswordsRequest

	items []*devices.DevicePassword
}

func (c *DeviceServiceClient) DevicePasswordsIterator(ctx context.Context, deviceId string, opts ...grpc.CallOption) *DevicePasswordsIterator {
	return &DevicePasswordsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &devices.ListDevicePasswordsRequest{
			DeviceId: deviceId,
		},
	}
}

func (it *DevicePasswordsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started {
		return false
	}
	it.started = true

	response, err := it.client.ListPasswords(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Passwords
	return len(it.items) > 0
}

func (it *DevicePasswordsIterator) Value() *devices.DevicePassword {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *DevicePasswordsIterator) Error() error {
	return it.err
}

// Update implements devices.DeviceServiceClient
func (c *DeviceServiceClient) Update(ctx context.Context, in *devices.UpdateDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).Update(ctx, in, opts...)
}