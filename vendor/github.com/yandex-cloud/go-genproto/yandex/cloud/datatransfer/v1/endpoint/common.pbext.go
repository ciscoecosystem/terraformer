// Code generated by protoc-gen-goext. DO NOT EDIT.

package endpoint

import (
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Secret_Value = isSecret_Value

func (m *Secret) SetValue(v Secret_Value) {
	m.Value = v
}

func (m *Secret) SetRaw(v string) {
	m.Value = &Secret_Raw{
		Raw: v,
	}
}

type TLSMode_TlsMode = isTLSMode_TlsMode

func (m *TLSMode) SetTlsMode(v TLSMode_TlsMode) {
	m.TlsMode = v
}

func (m *TLSMode) SetDisabled(v *emptypb.Empty) {
	m.TlsMode = &TLSMode_Disabled{
		Disabled: v,
	}
}

func (m *TLSMode) SetEnabled(v *TLSConfig) {
	m.TlsMode = &TLSMode_Enabled{
		Enabled: v,
	}
}

func (m *TLSConfig) SetCaCertificate(v string) {
	m.CaCertificate = v
}
