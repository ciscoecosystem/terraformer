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

// GetByName implements devices.DeviceServiceClient
func (c *DeviceServiceClient) GetByName(ctx context.Context, in *devices.GetByNameDeviceRequest, opts ...grpc.CallOption) (*devices.Device, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).GetByName(ctx, in, opts...)
}

// List implements devices.DeviceServiceClient
func (c *DeviceServiceClient) List(ctx context.Context, in *devices.ListDevicesRequest, opts ...grpc.CallOption) (*devices.ListDevicesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewDeviceServiceClient(conn).List(ctx, in, opts...)
}

type DeviceIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *DeviceServiceClient
	request *devices.ListDevicesRequest

	items []*devices.Device
}

func (c *DeviceServiceClient) DeviceIterator(ctx context.Context, req *devices.ListDevicesRequest, opts ...grpc.CallOption) *DeviceIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &DeviceIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *DeviceIterator) Next() bool {
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

	it.items = response.Devices
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *DeviceIterator) Take(size int64) ([]*devices.Device, error) {
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

	var result []*devices.Device

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *DeviceIterator) TakeAll() ([]*devices.Device, error) {
	return it.Take(0)
}

func (it *DeviceIterator) Value() *devices.Device {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *DeviceIterator) Error() error {
	return it.err
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

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *DeviceServiceClient
	request *devices.ListDeviceCertificatesRequest

	items []*devices.DeviceCertificate
}

func (c *DeviceServiceClient) DeviceCertificatesIterator(ctx context.Context, req *devices.ListDeviceCertificatesRequest, opts ...grpc.CallOption) *DeviceCertificatesIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &DeviceCertificatesIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
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

func (it *DeviceCertificatesIterator) Take(size int64) ([]*devices.DeviceCertificate, error) {
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

	var result []*devices.DeviceCertificate

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *DeviceCertificatesIterator) TakeAll() ([]*devices.DeviceCertificate, error) {
	return it.Take(0)
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

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *DeviceServiceClient
	request *devices.ListDeviceOperationsRequest

	items []*operation.Operation
}

func (c *DeviceServiceClient) DeviceOperationsIterator(ctx context.Context, req *devices.ListDeviceOperationsRequest, opts ...grpc.CallOption) *DeviceOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &DeviceOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
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

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *DeviceOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *DeviceOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
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

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *DeviceServiceClient
	request *devices.ListDevicePasswordsRequest

	items []*devices.DevicePassword
}

func (c *DeviceServiceClient) DevicePasswordsIterator(ctx context.Context, req *devices.ListDevicePasswordsRequest, opts ...grpc.CallOption) *DevicePasswordsIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &DevicePasswordsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
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

func (it *DevicePasswordsIterator) Take(size int64) ([]*devices.DevicePassword, error) {
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

	var result []*devices.DevicePassword

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *DevicePasswordsIterator) TakeAll() ([]*devices.DevicePassword, error) {
	return it.Take(0)
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
