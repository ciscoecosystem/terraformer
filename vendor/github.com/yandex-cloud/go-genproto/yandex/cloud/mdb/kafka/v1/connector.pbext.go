// Code generated by protoc-gen-goext. DO NOT EDIT.

package kafka

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type ConnectorSpec_ConnectorConfig = isConnectorSpec_ConnectorConfig

func (m *ConnectorSpec) SetConnectorConfig(v ConnectorSpec_ConnectorConfig) {
	m.ConnectorConfig = v
}

func (m *ConnectorSpec) SetName(v string) {
	m.Name = v
}

func (m *ConnectorSpec) SetTasksMax(v *wrapperspb.Int64Value) {
	m.TasksMax = v
}

func (m *ConnectorSpec) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *ConnectorSpec) SetConnectorConfigMirrormaker(v *ConnectorConfigMirrorMakerSpec) {
	m.ConnectorConfig = &ConnectorSpec_ConnectorConfigMirrormaker{
		ConnectorConfigMirrormaker: v,
	}
}

type UpdateConnectorSpec_ConnectorConfig = isUpdateConnectorSpec_ConnectorConfig

func (m *UpdateConnectorSpec) SetConnectorConfig(v UpdateConnectorSpec_ConnectorConfig) {
	m.ConnectorConfig = v
}

func (m *UpdateConnectorSpec) SetTasksMax(v *wrapperspb.Int64Value) {
	m.TasksMax = v
}

func (m *UpdateConnectorSpec) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *UpdateConnectorSpec) SetConnectorConfigMirrormaker(v *ConnectorConfigMirrorMakerSpec) {
	m.ConnectorConfig = &UpdateConnectorSpec_ConnectorConfigMirrormaker{
		ConnectorConfigMirrormaker: v,
	}
}

func (m *ConnectorConfigMirrorMakerSpec) SetSourceCluster(v *ClusterConnectionSpec) {
	m.SourceCluster = v
}

func (m *ConnectorConfigMirrorMakerSpec) SetTargetCluster(v *ClusterConnectionSpec) {
	m.TargetCluster = v
}

func (m *ConnectorConfigMirrorMakerSpec) SetTopics(v string) {
	m.Topics = v
}

func (m *ConnectorConfigMirrorMakerSpec) SetReplicationFactor(v *wrapperspb.Int64Value) {
	m.ReplicationFactor = v
}

type ClusterConnectionSpec_ClusterConnection = isClusterConnectionSpec_ClusterConnection

func (m *ClusterConnectionSpec) SetClusterConnection(v ClusterConnectionSpec_ClusterConnection) {
	m.ClusterConnection = v
}

func (m *ClusterConnectionSpec) SetAlias(v string) {
	m.Alias = v
}

func (m *ClusterConnectionSpec) SetThisCluster(v *ThisClusterSpec) {
	m.ClusterConnection = &ClusterConnectionSpec_ThisCluster{
		ThisCluster: v,
	}
}

func (m *ClusterConnectionSpec) SetExternalCluster(v *ExternalClusterConnectionSpec) {
	m.ClusterConnection = &ClusterConnectionSpec_ExternalCluster{
		ExternalCluster: v,
	}
}

func (m *ExternalClusterConnectionSpec) SetBootstrapServers(v string) {
	m.BootstrapServers = v
}

func (m *ExternalClusterConnectionSpec) SetSaslUsername(v string) {
	m.SaslUsername = v
}

func (m *ExternalClusterConnectionSpec) SetSaslPassword(v string) {
	m.SaslPassword = v
}

func (m *ExternalClusterConnectionSpec) SetSaslMechanism(v string) {
	m.SaslMechanism = v
}

func (m *ExternalClusterConnectionSpec) SetSecurityProtocol(v string) {
	m.SecurityProtocol = v
}

type Connector_ConnectorConfig = isConnector_ConnectorConfig

func (m *Connector) SetConnectorConfig(v Connector_ConnectorConfig) {
	m.ConnectorConfig = v
}

func (m *Connector) SetName(v string) {
	m.Name = v
}

func (m *Connector) SetTasksMax(v *wrapperspb.Int64Value) {
	m.TasksMax = v
}

func (m *Connector) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *Connector) SetHealth(v Connector_Health) {
	m.Health = v
}

func (m *Connector) SetStatus(v Connector_Status) {
	m.Status = v
}

func (m *Connector) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Connector) SetConnectorConfigMirrormaker(v *ConnectorConfigMirrorMaker) {
	m.ConnectorConfig = &Connector_ConnectorConfigMirrormaker{
		ConnectorConfigMirrormaker: v,
	}
}

func (m *ConnectorConfigMirrorMaker) SetSourceCluster(v *ClusterConnection) {
	m.SourceCluster = v
}

func (m *ConnectorConfigMirrorMaker) SetTargetCluster(v *ClusterConnection) {
	m.TargetCluster = v
}

func (m *ConnectorConfigMirrorMaker) SetTopics(v string) {
	m.Topics = v
}

func (m *ConnectorConfigMirrorMaker) SetReplicationFactor(v *wrapperspb.Int64Value) {
	m.ReplicationFactor = v
}

type ClusterConnection_ClusterConnection = isClusterConnection_ClusterConnection

func (m *ClusterConnection) SetClusterConnection(v ClusterConnection_ClusterConnection) {
	m.ClusterConnection = v
}

func (m *ClusterConnection) SetAlias(v string) {
	m.Alias = v
}

func (m *ClusterConnection) SetThisCluster(v *ThisCluster) {
	m.ClusterConnection = &ClusterConnection_ThisCluster{
		ThisCluster: v,
	}
}

func (m *ClusterConnection) SetExternalCluster(v *ExternalClusterConnection) {
	m.ClusterConnection = &ClusterConnection_ExternalCluster{
		ExternalCluster: v,
	}
}

func (m *ExternalClusterConnection) SetBootstrapServers(v string) {
	m.BootstrapServers = v
}

func (m *ExternalClusterConnection) SetSaslUsername(v string) {
	m.SaslUsername = v
}

func (m *ExternalClusterConnection) SetSaslMechanism(v string) {
	m.SaslMechanism = v
}

func (m *ExternalClusterConnection) SetSecurityProtocol(v string) {
	m.SecurityProtocol = v
}
