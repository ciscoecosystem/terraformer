// Code generated by protoc-gen-goext. DO NOT EDIT.

package vision

func (m *FaceAnnotation) SetFaces(v []*Face) {
	m.Faces = v
}

func (m *Face) SetBoundingBox(v *Polygon) {
	m.BoundingBox = v
}
