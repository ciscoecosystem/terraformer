// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetSnapshotRequest) SetSnapshotId(v string) {
	m.SnapshotId = v
}

func (m *ListSnapshotsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListSnapshotsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListSnapshotsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListSnapshotsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListSnapshotsResponse) SetSnapshots(v []*Snapshot) {
	m.Snapshots = v
}

func (m *ListSnapshotsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateSnapshotRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateSnapshotRequest) SetDiskId(v string) {
	m.DiskId = v
}

func (m *CreateSnapshotRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateSnapshotRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateSnapshotRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateSnapshotMetadata) SetSnapshotId(v string) {
	m.SnapshotId = v
}

func (m *CreateSnapshotMetadata) SetDiskId(v string) {
	m.DiskId = v
}

func (m *UpdateSnapshotRequest) SetSnapshotId(v string) {
	m.SnapshotId = v
}

func (m *UpdateSnapshotRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateSnapshotRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateSnapshotRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateSnapshotRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateSnapshotMetadata) SetSnapshotId(v string) {
	m.SnapshotId = v
}

func (m *DeleteSnapshotRequest) SetSnapshotId(v string) {
	m.SnapshotId = v
}

func (m *DeleteSnapshotMetadata) SetSnapshotId(v string) {
	m.SnapshotId = v
}

func (m *ListSnapshotOperationsRequest) SetSnapshotId(v string) {
	m.SnapshotId = v
}

func (m *ListSnapshotOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListSnapshotOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListSnapshotOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListSnapshotOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
