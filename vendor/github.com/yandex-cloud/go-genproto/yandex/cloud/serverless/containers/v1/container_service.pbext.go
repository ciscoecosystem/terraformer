// Code generated by protoc-gen-goext. DO NOT EDIT.

package containers

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetContainerRequest) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *ListContainersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListContainersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListContainersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListContainersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListContainersResponse) SetContainers(v []*Container) {
	m.Containers = v
}

func (m *ListContainersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateContainerRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateContainerRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateContainerRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateContainerRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateContainerMetadata) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *UpdateContainerRequest) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *UpdateContainerRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateContainerRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateContainerRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateContainerRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateContainerMetadata) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *DeleteContainerRequest) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *DeleteContainerMetadata) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *GetContainerRevisionRequest) SetContainerRevisionId(v string) {
	m.ContainerRevisionId = v
}

type ListContainersRevisionsRequest_Id = isListContainersRevisionsRequest_Id

func (m *ListContainersRevisionsRequest) SetId(v ListContainersRevisionsRequest_Id) {
	m.Id = v
}

func (m *ListContainersRevisionsRequest) SetFolderId(v string) {
	m.Id = &ListContainersRevisionsRequest_FolderId{
		FolderId: v,
	}
}

func (m *ListContainersRevisionsRequest) SetContainerId(v string) {
	m.Id = &ListContainersRevisionsRequest_ContainerId{
		ContainerId: v,
	}
}

func (m *ListContainersRevisionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListContainersRevisionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListContainersRevisionsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListContainersRevisionsResponse) SetRevisions(v []*Revision) {
	m.Revisions = v
}

func (m *ListContainersRevisionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *DeployContainerRevisionRequest) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *DeployContainerRevisionRequest) SetDescription(v string) {
	m.Description = v
}

func (m *DeployContainerRevisionRequest) SetResources(v *Resources) {
	m.Resources = v
}

func (m *DeployContainerRevisionRequest) SetExecutionTimeout(v *durationpb.Duration) {
	m.ExecutionTimeout = v
}

func (m *DeployContainerRevisionRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *DeployContainerRevisionRequest) SetImageSpec(v *ImageSpec) {
	m.ImageSpec = v
}

func (m *DeployContainerRevisionRequest) SetConcurrency(v int64) {
	m.Concurrency = v
}

func (m *DeployContainerRevisionRequest) SetSecrets(v []*Secret) {
	m.Secrets = v
}

func (m *ImageSpec) SetImageUrl(v string) {
	m.ImageUrl = v
}

func (m *ImageSpec) SetCommand(v *Command) {
	m.Command = v
}

func (m *ImageSpec) SetArgs(v *Args) {
	m.Args = v
}

func (m *ImageSpec) SetEnvironment(v map[string]string) {
	m.Environment = v
}

func (m *ImageSpec) SetWorkingDir(v string) {
	m.WorkingDir = v
}

func (m *DeployContainerRevisionMetadata) SetContainerRevisionId(v string) {
	m.ContainerRevisionId = v
}

func (m *ListContainerOperationsRequest) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *ListContainerOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListContainerOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListContainerOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListContainerOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListContainerOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
