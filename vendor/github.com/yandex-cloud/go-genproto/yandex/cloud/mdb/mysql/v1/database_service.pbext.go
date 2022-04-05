// Code generated by protoc-gen-goext. DO NOT EDIT.

package mysql

func (m *GetDatabaseRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetDatabaseRequest) SetDatabaseName(v string) {
	m.DatabaseName = v
}

func (m *ListDatabasesRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListDatabasesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDatabasesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDatabasesResponse) SetDatabases(v []*Database) {
	m.Databases = v
}

func (m *ListDatabasesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateDatabaseRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateDatabaseRequest) SetDatabaseSpec(v *DatabaseSpec) {
	m.DatabaseSpec = v
}

func (m *CreateDatabaseMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateDatabaseMetadata) SetDatabaseName(v string) {
	m.DatabaseName = v
}

func (m *DeleteDatabaseRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteDatabaseRequest) SetDatabaseName(v string) {
	m.DatabaseName = v
}

func (m *DeleteDatabaseMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteDatabaseMetadata) SetDatabaseName(v string) {
	m.DatabaseName = v
}
