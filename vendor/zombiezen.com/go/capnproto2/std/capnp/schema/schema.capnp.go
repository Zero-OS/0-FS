package schema

// AUTO GENERATED - DO NOT EDIT

import (
	math "math"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in schema.capnp.
const (
	Field_noDiscriminant = uint16(65535)
)

type Node struct{ capnp.Struct }
type Node_structNode Node
type Node_enum Node
type Node_interface Node
type Node_const Node
type Node_annotation Node
type Node_Which uint16

const (
	Node_Which_file       Node_Which = 0
	Node_Which_structNode Node_Which = 1
	Node_Which_enum       Node_Which = 2
	Node_Which_interface  Node_Which = 3
	Node_Which_const      Node_Which = 4
	Node_Which_annotation Node_Which = 5
)

func (w Node_Which) String() string {
	const s = "filestructNodeenuminterfaceconstannotation"
	switch w {
	case Node_Which_file:
		return s[0:4]
	case Node_Which_structNode:
		return s[4:14]
	case Node_Which_enum:
		return s[14:18]
	case Node_Which_interface:
		return s[18:27]
	case Node_Which_const:
		return s[27:32]
	case Node_Which_annotation:
		return s[32:42]

	}
	return "Node_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Node_TypeID is the unique identifier for the type Node.
const Node_TypeID = 0xe682ab4cf923a417

func NewNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6})
	return Node{st}, err
}

func NewRootNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6})
	return Node{st}, err
}

func ReadRootNode(msg *capnp.Message) (Node, error) {
	root, err := msg.RootPtr()
	return Node{root.Struct()}, err
}

func (s Node) Which() Node_Which {
	return Node_Which(s.Struct.Uint16(12))
}
func (s Node) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Node) DisplayName() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Node) HasDisplayName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node) DisplayNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node) SetDisplayName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Node) DisplayNamePrefixLength() uint32 {
	return s.Struct.Uint32(8)
}

func (s Node) SetDisplayNamePrefixLength(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s Node) ScopeId() uint64 {
	return s.Struct.Uint64(16)
}

func (s Node) SetScopeId(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Node) Parameters() (Node_Parameter_List, error) {
	p, err := s.Struct.Ptr(5)
	return Node_Parameter_List{List: p.List()}, err
}

func (s Node) HasParameters() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetParameters(v Node_Parameter_List) error {
	return s.Struct.SetPtr(5, v.List.ToPtr())
}

// NewParameters sets the parameters field to a newly
// allocated Node_Parameter_List, preferring placement in s's segment.
func (s Node) NewParameters(n int32) (Node_Parameter_List, error) {
	l, err := NewNode_Parameter_List(s.Struct.Segment(), n)
	if err != nil {
		return Node_Parameter_List{}, err
	}
	err = s.Struct.SetPtr(5, l.List.ToPtr())
	return l, err
}

func (s Node) IsGeneric() bool {
	return s.Struct.Bit(288)
}

func (s Node) SetIsGeneric(v bool) {
	s.Struct.SetBit(288, v)
}

func (s Node) NestedNodes() (Node_NestedNode_List, error) {
	p, err := s.Struct.Ptr(1)
	return Node_NestedNode_List{List: p.List()}, err
}

func (s Node) HasNestedNodes() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Node) SetNestedNodes(v Node_NestedNode_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewNestedNodes sets the nestedNodes field to a newly
// allocated Node_NestedNode_List, preferring placement in s's segment.
func (s Node) NewNestedNodes(n int32) (Node_NestedNode_List, error) {
	l, err := NewNode_NestedNode_List(s.Struct.Segment(), n)
	if err != nil {
		return Node_NestedNode_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Node) Annotations() (Annotation_List, error) {
	p, err := s.Struct.Ptr(2)
	return Annotation_List{List: p.List()}, err
}

func (s Node) HasAnnotations() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Node) SetAnnotations(v Annotation_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Node) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(s.Struct.Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Node) SetFile() {
	s.Struct.SetUint16(12, 0)

}

func (s Node) StructNode() Node_structNode { return Node_structNode(s) }

func (s Node) SetStructNode() {
	s.Struct.SetUint16(12, 1)
}

func (s Node_structNode) DataWordCount() uint16 {
	return s.Struct.Uint16(14)
}

func (s Node_structNode) SetDataWordCount(v uint16) {
	s.Struct.SetUint16(14, v)
}

func (s Node_structNode) PointerCount() uint16 {
	return s.Struct.Uint16(24)
}

func (s Node_structNode) SetPointerCount(v uint16) {
	s.Struct.SetUint16(24, v)
}

func (s Node_structNode) PreferredListEncoding() ElementSize {
	return ElementSize(s.Struct.Uint16(26))
}

func (s Node_structNode) SetPreferredListEncoding(v ElementSize) {
	s.Struct.SetUint16(26, uint16(v))
}

func (s Node_structNode) IsGroup() bool {
	return s.Struct.Bit(224)
}

func (s Node_structNode) SetIsGroup(v bool) {
	s.Struct.SetBit(224, v)
}

func (s Node_structNode) DiscriminantCount() uint16 {
	return s.Struct.Uint16(30)
}

func (s Node_structNode) SetDiscriminantCount(v uint16) {
	s.Struct.SetUint16(30, v)
}

func (s Node_structNode) DiscriminantOffset() uint32 {
	return s.Struct.Uint32(32)
}

func (s Node_structNode) SetDiscriminantOffset(v uint32) {
	s.Struct.SetUint32(32, v)
}

func (s Node_structNode) Fields() (Field_List, error) {
	p, err := s.Struct.Ptr(3)
	return Field_List{List: p.List()}, err
}

func (s Node_structNode) HasFields() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node_structNode) SetFields(v Field_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewFields sets the fields field to a newly
// allocated Field_List, preferring placement in s's segment.
func (s Node_structNode) NewFields(n int32) (Field_List, error) {
	l, err := NewField_List(s.Struct.Segment(), n)
	if err != nil {
		return Field_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s Node) Enum() Node_enum { return Node_enum(s) }

func (s Node) SetEnum() {
	s.Struct.SetUint16(12, 2)
}

func (s Node_enum) Enumerants() (Enumerant_List, error) {
	p, err := s.Struct.Ptr(3)
	return Enumerant_List{List: p.List()}, err
}

func (s Node_enum) HasEnumerants() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node_enum) SetEnumerants(v Enumerant_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewEnumerants sets the enumerants field to a newly
// allocated Enumerant_List, preferring placement in s's segment.
func (s Node_enum) NewEnumerants(n int32) (Enumerant_List, error) {
	l, err := NewEnumerant_List(s.Struct.Segment(), n)
	if err != nil {
		return Enumerant_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s Node) Interface() Node_interface { return Node_interface(s) }

func (s Node) SetInterface() {
	s.Struct.SetUint16(12, 3)
}

func (s Node_interface) Methods() (Method_List, error) {
	p, err := s.Struct.Ptr(3)
	return Method_List{List: p.List()}, err
}

func (s Node_interface) HasMethods() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node_interface) SetMethods(v Method_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewMethods sets the methods field to a newly
// allocated Method_List, preferring placement in s's segment.
func (s Node_interface) NewMethods(n int32) (Method_List, error) {
	l, err := NewMethod_List(s.Struct.Segment(), n)
	if err != nil {
		return Method_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s Node_interface) Superclasses() (Superclass_List, error) {
	p, err := s.Struct.Ptr(4)
	return Superclass_List{List: p.List()}, err
}

func (s Node_interface) HasSuperclasses() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Node_interface) SetSuperclasses(v Superclass_List) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewSuperclasses sets the superclasses field to a newly
// allocated Superclass_List, preferring placement in s's segment.
func (s Node_interface) NewSuperclasses(n int32) (Superclass_List, error) {
	l, err := NewSuperclass_List(s.Struct.Segment(), n)
	if err != nil {
		return Superclass_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s Node) Const() Node_const { return Node_const(s) }

func (s Node) SetConst() {
	s.Struct.SetUint16(12, 4)
}

func (s Node_const) Type() (Type, error) {
	p, err := s.Struct.Ptr(3)
	return Type{Struct: p.Struct()}, err
}

func (s Node_const) HasType() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node_const) SetType(v Type) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Node_const) NewType() (Type, error) {
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node_const) Value() (Value, error) {
	p, err := s.Struct.Ptr(4)
	return Value{Struct: p.Struct()}, err
}

func (s Node_const) HasValue() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Node_const) SetValue(v Value) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewValue sets the value field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Node_const) NewValue() (Value, error) {
	ss, err := NewValue(s.Struct.Segment())
	if err != nil {
		return Value{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Annotation() Node_annotation { return Node_annotation(s) }

func (s Node) SetAnnotation() {
	s.Struct.SetUint16(12, 5)
}

func (s Node_annotation) Type() (Type, error) {
	p, err := s.Struct.Ptr(3)
	return Type{Struct: p.Struct()}, err
}

func (s Node_annotation) HasType() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node_annotation) SetType(v Type) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Node_annotation) NewType() (Type, error) {
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node_annotation) TargetsFile() bool {
	return s.Struct.Bit(112)
}

func (s Node_annotation) SetTargetsFile(v bool) {
	s.Struct.SetBit(112, v)
}

func (s Node_annotation) TargetsConst() bool {
	return s.Struct.Bit(113)
}

func (s Node_annotation) SetTargetsConst(v bool) {
	s.Struct.SetBit(113, v)
}

func (s Node_annotation) TargetsEnum() bool {
	return s.Struct.Bit(114)
}

func (s Node_annotation) SetTargetsEnum(v bool) {
	s.Struct.SetBit(114, v)
}

func (s Node_annotation) TargetsEnumerant() bool {
	return s.Struct.Bit(115)
}

func (s Node_annotation) SetTargetsEnumerant(v bool) {
	s.Struct.SetBit(115, v)
}

func (s Node_annotation) TargetsStruct() bool {
	return s.Struct.Bit(116)
}

func (s Node_annotation) SetTargetsStruct(v bool) {
	s.Struct.SetBit(116, v)
}

func (s Node_annotation) TargetsField() bool {
	return s.Struct.Bit(117)
}

func (s Node_annotation) SetTargetsField(v bool) {
	s.Struct.SetBit(117, v)
}

func (s Node_annotation) TargetsUnion() bool {
	return s.Struct.Bit(118)
}

func (s Node_annotation) SetTargetsUnion(v bool) {
	s.Struct.SetBit(118, v)
}

func (s Node_annotation) TargetsGroup() bool {
	return s.Struct.Bit(119)
}

func (s Node_annotation) SetTargetsGroup(v bool) {
	s.Struct.SetBit(119, v)
}

func (s Node_annotation) TargetsInterface() bool {
	return s.Struct.Bit(120)
}

func (s Node_annotation) SetTargetsInterface(v bool) {
	s.Struct.SetBit(120, v)
}

func (s Node_annotation) TargetsMethod() bool {
	return s.Struct.Bit(121)
}

func (s Node_annotation) SetTargetsMethod(v bool) {
	s.Struct.SetBit(121, v)
}

func (s Node_annotation) TargetsParam() bool {
	return s.Struct.Bit(122)
}

func (s Node_annotation) SetTargetsParam(v bool) {
	s.Struct.SetBit(122, v)
}

func (s Node_annotation) TargetsAnnotation() bool {
	return s.Struct.Bit(123)
}

func (s Node_annotation) SetTargetsAnnotation(v bool) {
	s.Struct.SetBit(123, v)
}

// Node_List is a list of Node.
type Node_List struct{ capnp.List }

// NewNode creates a new list of Node.
func NewNode_List(s *capnp.Segment, sz int32) (Node_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6}, sz)
	return Node_List{l}, err
}

func (s Node_List) At(i int) Node { return Node{s.List.Struct(i)} }

func (s Node_List) Set(i int, v Node) error { return s.List.SetStruct(i, v.Struct) }

// Node_Promise is a wrapper for a Node promised by a client call.
type Node_Promise struct{ *capnp.Pipeline }

func (p Node_Promise) Struct() (Node, error) {
	s, err := p.Pipeline.Struct()
	return Node{s}, err
}

func (p Node_Promise) StructNode() Node_structNode_Promise { return Node_structNode_Promise{p.Pipeline} }

// Node_structNode_Promise is a wrapper for a Node_structNode promised by a client call.
type Node_structNode_Promise struct{ *capnp.Pipeline }

func (p Node_structNode_Promise) Struct() (Node_structNode, error) {
	s, err := p.Pipeline.Struct()
	return Node_structNode{s}, err
}

func (p Node_Promise) Enum() Node_enum_Promise { return Node_enum_Promise{p.Pipeline} }

// Node_enum_Promise is a wrapper for a Node_enum promised by a client call.
type Node_enum_Promise struct{ *capnp.Pipeline }

func (p Node_enum_Promise) Struct() (Node_enum, error) {
	s, err := p.Pipeline.Struct()
	return Node_enum{s}, err
}

func (p Node_Promise) Interface() Node_interface_Promise { return Node_interface_Promise{p.Pipeline} }

// Node_interface_Promise is a wrapper for a Node_interface promised by a client call.
type Node_interface_Promise struct{ *capnp.Pipeline }

func (p Node_interface_Promise) Struct() (Node_interface, error) {
	s, err := p.Pipeline.Struct()
	return Node_interface{s}, err
}

func (p Node_Promise) Const() Node_const_Promise { return Node_const_Promise{p.Pipeline} }

// Node_const_Promise is a wrapper for a Node_const promised by a client call.
type Node_const_Promise struct{ *capnp.Pipeline }

func (p Node_const_Promise) Struct() (Node_const, error) {
	s, err := p.Pipeline.Struct()
	return Node_const{s}, err
}

func (p Node_const_Promise) Type() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_const_Promise) Value() Value_Promise {
	return Value_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

func (p Node_Promise) Annotation() Node_annotation_Promise { return Node_annotation_Promise{p.Pipeline} }

// Node_annotation_Promise is a wrapper for a Node_annotation promised by a client call.
type Node_annotation_Promise struct{ *capnp.Pipeline }

func (p Node_annotation_Promise) Struct() (Node_annotation, error) {
	s, err := p.Pipeline.Struct()
	return Node_annotation{s}, err
}

func (p Node_annotation_Promise) Type() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type Node_Parameter struct{ capnp.Struct }

// Node_Parameter_TypeID is the unique identifier for the type Node_Parameter.
const Node_Parameter_TypeID = 0xb9521bccf10fa3b1

func NewNode_Parameter(s *capnp.Segment) (Node_Parameter, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Node_Parameter{st}, err
}

func NewRootNode_Parameter(s *capnp.Segment) (Node_Parameter, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Node_Parameter{st}, err
}

func ReadRootNode_Parameter(msg *capnp.Message) (Node_Parameter, error) {
	root, err := msg.RootPtr()
	return Node_Parameter{root.Struct()}, err
}

func (s Node_Parameter) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Node_Parameter) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node_Parameter) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node_Parameter) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

// Node_Parameter_List is a list of Node_Parameter.
type Node_Parameter_List struct{ capnp.List }

// NewNode_Parameter creates a new list of Node_Parameter.
func NewNode_Parameter_List(s *capnp.Segment, sz int32) (Node_Parameter_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Node_Parameter_List{l}, err
}

func (s Node_Parameter_List) At(i int) Node_Parameter { return Node_Parameter{s.List.Struct(i)} }

func (s Node_Parameter_List) Set(i int, v Node_Parameter) error { return s.List.SetStruct(i, v.Struct) }

// Node_Parameter_Promise is a wrapper for a Node_Parameter promised by a client call.
type Node_Parameter_Promise struct{ *capnp.Pipeline }

func (p Node_Parameter_Promise) Struct() (Node_Parameter, error) {
	s, err := p.Pipeline.Struct()
	return Node_Parameter{s}, err
}

type Node_NestedNode struct{ capnp.Struct }

// Node_NestedNode_TypeID is the unique identifier for the type Node_NestedNode.
const Node_NestedNode_TypeID = 0xdebf55bbfa0fc242

func NewNode_NestedNode(s *capnp.Segment) (Node_NestedNode, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Node_NestedNode{st}, err
}

func NewRootNode_NestedNode(s *capnp.Segment) (Node_NestedNode, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Node_NestedNode{st}, err
}

func ReadRootNode_NestedNode(msg *capnp.Message) (Node_NestedNode, error) {
	root, err := msg.RootPtr()
	return Node_NestedNode{root.Struct()}, err
}

func (s Node_NestedNode) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Node_NestedNode) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node_NestedNode) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node_NestedNode) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Node_NestedNode) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node_NestedNode) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Node_NestedNode_List is a list of Node_NestedNode.
type Node_NestedNode_List struct{ capnp.List }

// NewNode_NestedNode creates a new list of Node_NestedNode.
func NewNode_NestedNode_List(s *capnp.Segment, sz int32) (Node_NestedNode_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Node_NestedNode_List{l}, err
}

func (s Node_NestedNode_List) At(i int) Node_NestedNode { return Node_NestedNode{s.List.Struct(i)} }

func (s Node_NestedNode_List) Set(i int, v Node_NestedNode) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_NestedNode_Promise is a wrapper for a Node_NestedNode promised by a client call.
type Node_NestedNode_Promise struct{ *capnp.Pipeline }

func (p Node_NestedNode_Promise) Struct() (Node_NestedNode, error) {
	s, err := p.Pipeline.Struct()
	return Node_NestedNode{s}, err
}

type Field struct{ capnp.Struct }
type Field_slot Field
type Field_group Field
type Field_ordinal Field
type Field_Which uint16

const (
	Field_Which_slot  Field_Which = 0
	Field_Which_group Field_Which = 1
)

func (w Field_Which) String() string {
	const s = "slotgroup"
	switch w {
	case Field_Which_slot:
		return s[0:4]
	case Field_Which_group:
		return s[4:9]

	}
	return "Field_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type Field_ordinal_Which uint16

const (
	Field_ordinal_Which_implicit Field_ordinal_Which = 0
	Field_ordinal_Which_explicit Field_ordinal_Which = 1
)

func (w Field_ordinal_Which) String() string {
	const s = "implicitexplicit"
	switch w {
	case Field_ordinal_Which_implicit:
		return s[0:8]
	case Field_ordinal_Which_explicit:
		return s[8:16]

	}
	return "Field_ordinal_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Field_TypeID is the unique identifier for the type Field.
const Field_TypeID = 0x9aad50a41f4af45f

func NewField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return Field{st}, err
}

func NewRootField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return Field{st}, err
}

func ReadRootField(msg *capnp.Message) (Field, error) {
	root, err := msg.RootPtr()
	return Field{root.Struct()}, err
}

func (s Field) Which() Field_Which {
	return Field_Which(s.Struct.Uint16(8))
}
func (s Field) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Field) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Field) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Field) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Field) CodeOrder() uint16 {
	return s.Struct.Uint16(0)
}

func (s Field) SetCodeOrder(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s Field) Annotations() (Annotation_List, error) {
	p, err := s.Struct.Ptr(1)
	return Annotation_List{List: p.List()}, err
}

func (s Field) HasAnnotations() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Field) SetAnnotations(v Annotation_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Field) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(s.Struct.Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Field) DiscriminantValue() uint16 {
	return s.Struct.Uint16(2) ^ 65535
}

func (s Field) SetDiscriminantValue(v uint16) {
	s.Struct.SetUint16(2, v^65535)
}

func (s Field) Slot() Field_slot { return Field_slot(s) }

func (s Field) SetSlot() {
	s.Struct.SetUint16(8, 0)
}

func (s Field_slot) Offset() uint32 {
	return s.Struct.Uint32(4)
}

func (s Field_slot) SetOffset(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s Field_slot) Type() (Type, error) {
	p, err := s.Struct.Ptr(2)
	return Type{Struct: p.Struct()}, err
}

func (s Field_slot) HasType() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Field_slot) SetType(v Type) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Field_slot) NewType() (Type, error) {
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Field_slot) DefaultValue() (Value, error) {
	p, err := s.Struct.Ptr(3)
	return Value{Struct: p.Struct()}, err
}

func (s Field_slot) HasDefaultValue() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Field_slot) SetDefaultValue(v Value) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDefaultValue sets the defaultValue field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Field_slot) NewDefaultValue() (Value, error) {
	ss, err := NewValue(s.Struct.Segment())
	if err != nil {
		return Value{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Field_slot) HadExplicitDefault() bool {
	return s.Struct.Bit(128)
}

func (s Field_slot) SetHadExplicitDefault(v bool) {
	s.Struct.SetBit(128, v)
}

func (s Field) Group() Field_group { return Field_group(s) }

func (s Field) SetGroup() {
	s.Struct.SetUint16(8, 1)
}

func (s Field_group) TypeId() uint64 {
	return s.Struct.Uint64(16)
}

func (s Field_group) SetTypeId(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Field) Ordinal() Field_ordinal { return Field_ordinal(s) }

func (s Field_ordinal) Which() Field_ordinal_Which {
	return Field_ordinal_Which(s.Struct.Uint16(10))
}
func (s Field_ordinal) SetImplicit() {
	s.Struct.SetUint16(10, 0)

}

func (s Field_ordinal) Explicit() uint16 {
	return s.Struct.Uint16(12)
}

func (s Field_ordinal) SetExplicit(v uint16) {
	s.Struct.SetUint16(10, 1)
	s.Struct.SetUint16(12, v)
}

// Field_List is a list of Field.
type Field_List struct{ capnp.List }

// NewField creates a new list of Field.
func NewField_List(s *capnp.Segment, sz int32) (Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4}, sz)
	return Field_List{l}, err
}

func (s Field_List) At(i int) Field { return Field{s.List.Struct(i)} }

func (s Field_List) Set(i int, v Field) error { return s.List.SetStruct(i, v.Struct) }

// Field_Promise is a wrapper for a Field promised by a client call.
type Field_Promise struct{ *capnp.Pipeline }

func (p Field_Promise) Struct() (Field, error) {
	s, err := p.Pipeline.Struct()
	return Field{s}, err
}

func (p Field_Promise) Slot() Field_slot_Promise { return Field_slot_Promise{p.Pipeline} }

// Field_slot_Promise is a wrapper for a Field_slot promised by a client call.
type Field_slot_Promise struct{ *capnp.Pipeline }

func (p Field_slot_Promise) Struct() (Field_slot, error) {
	s, err := p.Pipeline.Struct()
	return Field_slot{s}, err
}

func (p Field_slot_Promise) Type() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Field_slot_Promise) DefaultValue() Value_Promise {
	return Value_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Field_Promise) Group() Field_group_Promise { return Field_group_Promise{p.Pipeline} }

// Field_group_Promise is a wrapper for a Field_group promised by a client call.
type Field_group_Promise struct{ *capnp.Pipeline }

func (p Field_group_Promise) Struct() (Field_group, error) {
	s, err := p.Pipeline.Struct()
	return Field_group{s}, err
}

func (p Field_Promise) Ordinal() Field_ordinal_Promise { return Field_ordinal_Promise{p.Pipeline} }

// Field_ordinal_Promise is a wrapper for a Field_ordinal promised by a client call.
type Field_ordinal_Promise struct{ *capnp.Pipeline }

func (p Field_ordinal_Promise) Struct() (Field_ordinal, error) {
	s, err := p.Pipeline.Struct()
	return Field_ordinal{s}, err
}

type Enumerant struct{ capnp.Struct }

// Enumerant_TypeID is the unique identifier for the type Enumerant.
const Enumerant_TypeID = 0x978a7cebdc549a4d

func NewEnumerant(s *capnp.Segment) (Enumerant, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Enumerant{st}, err
}

func NewRootEnumerant(s *capnp.Segment) (Enumerant, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Enumerant{st}, err
}

func ReadRootEnumerant(msg *capnp.Message) (Enumerant, error) {
	root, err := msg.RootPtr()
	return Enumerant{root.Struct()}, err
}

func (s Enumerant) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Enumerant) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Enumerant) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Enumerant) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Enumerant) CodeOrder() uint16 {
	return s.Struct.Uint16(0)
}

func (s Enumerant) SetCodeOrder(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s Enumerant) Annotations() (Annotation_List, error) {
	p, err := s.Struct.Ptr(1)
	return Annotation_List{List: p.List()}, err
}

func (s Enumerant) HasAnnotations() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Enumerant) SetAnnotations(v Annotation_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Enumerant) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(s.Struct.Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Enumerant_List is a list of Enumerant.
type Enumerant_List struct{ capnp.List }

// NewEnumerant creates a new list of Enumerant.
func NewEnumerant_List(s *capnp.Segment, sz int32) (Enumerant_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Enumerant_List{l}, err
}

func (s Enumerant_List) At(i int) Enumerant { return Enumerant{s.List.Struct(i)} }

func (s Enumerant_List) Set(i int, v Enumerant) error { return s.List.SetStruct(i, v.Struct) }

// Enumerant_Promise is a wrapper for a Enumerant promised by a client call.
type Enumerant_Promise struct{ *capnp.Pipeline }

func (p Enumerant_Promise) Struct() (Enumerant, error) {
	s, err := p.Pipeline.Struct()
	return Enumerant{s}, err
}

type Superclass struct{ capnp.Struct }

// Superclass_TypeID is the unique identifier for the type Superclass.
const Superclass_TypeID = 0xa9962a9ed0a4d7f8

func NewSuperclass(s *capnp.Segment) (Superclass, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Superclass{st}, err
}

func NewRootSuperclass(s *capnp.Segment) (Superclass, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Superclass{st}, err
}

func ReadRootSuperclass(msg *capnp.Message) (Superclass, error) {
	root, err := msg.RootPtr()
	return Superclass{root.Struct()}, err
}

func (s Superclass) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s Superclass) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Superclass) Brand() (Brand, error) {
	p, err := s.Struct.Ptr(0)
	return Brand{Struct: p.Struct()}, err
}

func (s Superclass) HasBrand() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Superclass) SetBrand(v Brand) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Superclass) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Superclass_List is a list of Superclass.
type Superclass_List struct{ capnp.List }

// NewSuperclass creates a new list of Superclass.
func NewSuperclass_List(s *capnp.Segment, sz int32) (Superclass_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Superclass_List{l}, err
}

func (s Superclass_List) At(i int) Superclass { return Superclass{s.List.Struct(i)} }

func (s Superclass_List) Set(i int, v Superclass) error { return s.List.SetStruct(i, v.Struct) }

// Superclass_Promise is a wrapper for a Superclass promised by a client call.
type Superclass_Promise struct{ *capnp.Pipeline }

func (p Superclass_Promise) Struct() (Superclass, error) {
	s, err := p.Pipeline.Struct()
	return Superclass{s}, err
}

func (p Superclass_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Method struct{ capnp.Struct }

// Method_TypeID is the unique identifier for the type Method.
const Method_TypeID = 0x9500cce23b334d80

func NewMethod(s *capnp.Segment) (Method, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5})
	return Method{st}, err
}

func NewRootMethod(s *capnp.Segment) (Method, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5})
	return Method{st}, err
}

func ReadRootMethod(msg *capnp.Message) (Method, error) {
	root, err := msg.RootPtr()
	return Method{root.Struct()}, err
}

func (s Method) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Method) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Method) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Method) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Method) CodeOrder() uint16 {
	return s.Struct.Uint16(0)
}

func (s Method) SetCodeOrder(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s Method) ImplicitParameters() (Node_Parameter_List, error) {
	p, err := s.Struct.Ptr(4)
	return Node_Parameter_List{List: p.List()}, err
}

func (s Method) HasImplicitParameters() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Method) SetImplicitParameters(v Node_Parameter_List) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewImplicitParameters sets the implicitParameters field to a newly
// allocated Node_Parameter_List, preferring placement in s's segment.
func (s Method) NewImplicitParameters(n int32) (Node_Parameter_List, error) {
	l, err := NewNode_Parameter_List(s.Struct.Segment(), n)
	if err != nil {
		return Node_Parameter_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s Method) ParamStructType() uint64 {
	return s.Struct.Uint64(8)
}

func (s Method) SetParamStructType(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Method) ParamBrand() (Brand, error) {
	p, err := s.Struct.Ptr(2)
	return Brand{Struct: p.Struct()}, err
}

func (s Method) HasParamBrand() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Method) SetParamBrand(v Brand) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewParamBrand sets the paramBrand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Method) NewParamBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Method) ResultStructType() uint64 {
	return s.Struct.Uint64(16)
}

func (s Method) SetResultStructType(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Method) ResultBrand() (Brand, error) {
	p, err := s.Struct.Ptr(3)
	return Brand{Struct: p.Struct()}, err
}

func (s Method) HasResultBrand() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Method) SetResultBrand(v Brand) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewResultBrand sets the resultBrand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Method) NewResultBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Method) Annotations() (Annotation_List, error) {
	p, err := s.Struct.Ptr(1)
	return Annotation_List{List: p.List()}, err
}

func (s Method) HasAnnotations() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Method) SetAnnotations(v Annotation_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Method) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(s.Struct.Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Method_List is a list of Method.
type Method_List struct{ capnp.List }

// NewMethod creates a new list of Method.
func NewMethod_List(s *capnp.Segment, sz int32) (Method_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5}, sz)
	return Method_List{l}, err
}

func (s Method_List) At(i int) Method { return Method{s.List.Struct(i)} }

func (s Method_List) Set(i int, v Method) error { return s.List.SetStruct(i, v.Struct) }

// Method_Promise is a wrapper for a Method promised by a client call.
type Method_Promise struct{ *capnp.Pipeline }

func (p Method_Promise) Struct() (Method, error) {
	s, err := p.Pipeline.Struct()
	return Method{s}, err
}

func (p Method_Promise) ParamBrand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Method_Promise) ResultBrand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type Type struct{ capnp.Struct }
type Type_list Type
type Type_enum Type
type Type_structType Type
type Type_interface Type
type Type_anyPointer Type
type Type_anyPointer_unconstrained Type
type Type_anyPointer_parameter Type
type Type_anyPointer_implicitMethodParameter Type
type Type_Which uint16

const (
	Type_Which_void       Type_Which = 0
	Type_Which_bool       Type_Which = 1
	Type_Which_int8       Type_Which = 2
	Type_Which_int16      Type_Which = 3
	Type_Which_int32      Type_Which = 4
	Type_Which_int64      Type_Which = 5
	Type_Which_uint8      Type_Which = 6
	Type_Which_uint16     Type_Which = 7
	Type_Which_uint32     Type_Which = 8
	Type_Which_uint64     Type_Which = 9
	Type_Which_float32    Type_Which = 10
	Type_Which_float64    Type_Which = 11
	Type_Which_text       Type_Which = 12
	Type_Which_data       Type_Which = 13
	Type_Which_list       Type_Which = 14
	Type_Which_enum       Type_Which = 15
	Type_Which_structType Type_Which = 16
	Type_Which_interface  Type_Which = 17
	Type_Which_anyPointer Type_Which = 18
)

func (w Type_Which) String() string {
	const s = "voidboolint8int16int32int64uint8uint16uint32uint64float32float64textdatalistenumstructTypeinterfaceanyPointer"
	switch w {
	case Type_Which_void:
		return s[0:4]
	case Type_Which_bool:
		return s[4:8]
	case Type_Which_int8:
		return s[8:12]
	case Type_Which_int16:
		return s[12:17]
	case Type_Which_int32:
		return s[17:22]
	case Type_Which_int64:
		return s[22:27]
	case Type_Which_uint8:
		return s[27:32]
	case Type_Which_uint16:
		return s[32:38]
	case Type_Which_uint32:
		return s[38:44]
	case Type_Which_uint64:
		return s[44:50]
	case Type_Which_float32:
		return s[50:57]
	case Type_Which_float64:
		return s[57:64]
	case Type_Which_text:
		return s[64:68]
	case Type_Which_data:
		return s[68:72]
	case Type_Which_list:
		return s[72:76]
	case Type_Which_enum:
		return s[76:80]
	case Type_Which_structType:
		return s[80:90]
	case Type_Which_interface:
		return s[90:99]
	case Type_Which_anyPointer:
		return s[99:109]

	}
	return "Type_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type Type_anyPointer_Which uint16

const (
	Type_anyPointer_Which_unconstrained           Type_anyPointer_Which = 0
	Type_anyPointer_Which_parameter               Type_anyPointer_Which = 1
	Type_anyPointer_Which_implicitMethodParameter Type_anyPointer_Which = 2
)

func (w Type_anyPointer_Which) String() string {
	const s = "unconstrainedparameterimplicitMethodParameter"
	switch w {
	case Type_anyPointer_Which_unconstrained:
		return s[0:13]
	case Type_anyPointer_Which_parameter:
		return s[13:22]
	case Type_anyPointer_Which_implicitMethodParameter:
		return s[22:45]

	}
	return "Type_anyPointer_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type Type_anyPointer_unconstrained_Which uint16

const (
	Type_anyPointer_unconstrained_Which_anyKind    Type_anyPointer_unconstrained_Which = 0
	Type_anyPointer_unconstrained_Which_struct     Type_anyPointer_unconstrained_Which = 1
	Type_anyPointer_unconstrained_Which_list       Type_anyPointer_unconstrained_Which = 2
	Type_anyPointer_unconstrained_Which_capability Type_anyPointer_unconstrained_Which = 3
)

func (w Type_anyPointer_unconstrained_Which) String() string {
	const s = "anyKindstructlistcapability"
	switch w {
	case Type_anyPointer_unconstrained_Which_anyKind:
		return s[0:7]
	case Type_anyPointer_unconstrained_Which_struct:
		return s[7:13]
	case Type_anyPointer_unconstrained_Which_list:
		return s[13:17]
	case Type_anyPointer_unconstrained_Which_capability:
		return s[17:27]

	}
	return "Type_anyPointer_unconstrained_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Type_TypeID is the unique identifier for the type Type.
const Type_TypeID = 0xd07378ede1f9cc60

func NewType(s *capnp.Segment) (Type, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return Type{st}, err
}

func NewRootType(s *capnp.Segment) (Type, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return Type{st}, err
}

func ReadRootType(msg *capnp.Message) (Type, error) {
	root, err := msg.RootPtr()
	return Type{root.Struct()}, err
}

func (s Type) Which() Type_Which {
	return Type_Which(s.Struct.Uint16(0))
}
func (s Type) SetVoid() {
	s.Struct.SetUint16(0, 0)

}

func (s Type) SetBool() {
	s.Struct.SetUint16(0, 1)

}

func (s Type) SetInt8() {
	s.Struct.SetUint16(0, 2)

}

func (s Type) SetInt16() {
	s.Struct.SetUint16(0, 3)

}

func (s Type) SetInt32() {
	s.Struct.SetUint16(0, 4)

}

func (s Type) SetInt64() {
	s.Struct.SetUint16(0, 5)

}

func (s Type) SetUint8() {
	s.Struct.SetUint16(0, 6)

}

func (s Type) SetUint16() {
	s.Struct.SetUint16(0, 7)

}

func (s Type) SetUint32() {
	s.Struct.SetUint16(0, 8)

}

func (s Type) SetUint64() {
	s.Struct.SetUint16(0, 9)

}

func (s Type) SetFloat32() {
	s.Struct.SetUint16(0, 10)

}

func (s Type) SetFloat64() {
	s.Struct.SetUint16(0, 11)

}

func (s Type) SetText() {
	s.Struct.SetUint16(0, 12)

}

func (s Type) SetData() {
	s.Struct.SetUint16(0, 13)

}

func (s Type) List() Type_list { return Type_list(s) }

func (s Type) SetList() {
	s.Struct.SetUint16(0, 14)
}

func (s Type_list) ElementType() (Type, error) {
	p, err := s.Struct.Ptr(0)
	return Type{Struct: p.Struct()}, err
}

func (s Type_list) HasElementType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Type_list) SetElementType(v Type) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewElementType sets the elementType field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Type_list) NewElementType() (Type, error) {
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Type) Enum() Type_enum { return Type_enum(s) }

func (s Type) SetEnum() {
	s.Struct.SetUint16(0, 15)
}

func (s Type_enum) TypeId() uint64 {
	return s.Struct.Uint64(8)
}

func (s Type_enum) SetTypeId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Type_enum) Brand() (Brand, error) {
	p, err := s.Struct.Ptr(0)
	return Brand{Struct: p.Struct()}, err
}

func (s Type_enum) HasBrand() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Type_enum) SetBrand(v Brand) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Type_enum) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Type) StructType() Type_structType { return Type_structType(s) }

func (s Type) SetStructType() {
	s.Struct.SetUint16(0, 16)
}

func (s Type_structType) TypeId() uint64 {
	return s.Struct.Uint64(8)
}

func (s Type_structType) SetTypeId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Type_structType) Brand() (Brand, error) {
	p, err := s.Struct.Ptr(0)
	return Brand{Struct: p.Struct()}, err
}

func (s Type_structType) HasBrand() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Type_structType) SetBrand(v Brand) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Type_structType) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Type) Interface() Type_interface { return Type_interface(s) }

func (s Type) SetInterface() {
	s.Struct.SetUint16(0, 17)
}

func (s Type_interface) TypeId() uint64 {
	return s.Struct.Uint64(8)
}

func (s Type_interface) SetTypeId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Type_interface) Brand() (Brand, error) {
	p, err := s.Struct.Ptr(0)
	return Brand{Struct: p.Struct()}, err
}

func (s Type_interface) HasBrand() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Type_interface) SetBrand(v Brand) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Type_interface) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Type) AnyPointer() Type_anyPointer { return Type_anyPointer(s) }

func (s Type) SetAnyPointer() {
	s.Struct.SetUint16(0, 18)
}

func (s Type_anyPointer) Which() Type_anyPointer_Which {
	return Type_anyPointer_Which(s.Struct.Uint16(8))
}
func (s Type_anyPointer) Unconstrained() Type_anyPointer_unconstrained {
	return Type_anyPointer_unconstrained(s)
}

func (s Type_anyPointer) SetUnconstrained() {
	s.Struct.SetUint16(8, 0)
}

func (s Type_anyPointer_unconstrained) Which() Type_anyPointer_unconstrained_Which {
	return Type_anyPointer_unconstrained_Which(s.Struct.Uint16(10))
}
func (s Type_anyPointer_unconstrained) SetAnyKind() {
	s.Struct.SetUint16(10, 0)

}

func (s Type_anyPointer_unconstrained) SetStruct() {
	s.Struct.SetUint16(10, 1)

}

func (s Type_anyPointer_unconstrained) SetList() {
	s.Struct.SetUint16(10, 2)

}

func (s Type_anyPointer_unconstrained) SetCapability() {
	s.Struct.SetUint16(10, 3)

}

func (s Type_anyPointer) Parameter() Type_anyPointer_parameter { return Type_anyPointer_parameter(s) }

func (s Type_anyPointer) SetParameter() {
	s.Struct.SetUint16(8, 1)
}

func (s Type_anyPointer_parameter) ScopeId() uint64 {
	return s.Struct.Uint64(16)
}

func (s Type_anyPointer_parameter) SetScopeId(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Type_anyPointer_parameter) ParameterIndex() uint16 {
	return s.Struct.Uint16(10)
}

func (s Type_anyPointer_parameter) SetParameterIndex(v uint16) {
	s.Struct.SetUint16(10, v)
}

func (s Type_anyPointer) ImplicitMethodParameter() Type_anyPointer_implicitMethodParameter {
	return Type_anyPointer_implicitMethodParameter(s)
}

func (s Type_anyPointer) SetImplicitMethodParameter() {
	s.Struct.SetUint16(8, 2)
}

func (s Type_anyPointer_implicitMethodParameter) ParameterIndex() uint16 {
	return s.Struct.Uint16(10)
}

func (s Type_anyPointer_implicitMethodParameter) SetParameterIndex(v uint16) {
	s.Struct.SetUint16(10, v)
}

// Type_List is a list of Type.
type Type_List struct{ capnp.List }

// NewType creates a new list of Type.
func NewType_List(s *capnp.Segment, sz int32) (Type_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1}, sz)
	return Type_List{l}, err
}

func (s Type_List) At(i int) Type { return Type{s.List.Struct(i)} }

func (s Type_List) Set(i int, v Type) error { return s.List.SetStruct(i, v.Struct) }

// Type_Promise is a wrapper for a Type promised by a client call.
type Type_Promise struct{ *capnp.Pipeline }

func (p Type_Promise) Struct() (Type, error) {
	s, err := p.Pipeline.Struct()
	return Type{s}, err
}

func (p Type_Promise) List() Type_list_Promise { return Type_list_Promise{p.Pipeline} }

// Type_list_Promise is a wrapper for a Type_list promised by a client call.
type Type_list_Promise struct{ *capnp.Pipeline }

func (p Type_list_Promise) Struct() (Type_list, error) {
	s, err := p.Pipeline.Struct()
	return Type_list{s}, err
}

func (p Type_list_Promise) ElementType() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Type_Promise) Enum() Type_enum_Promise { return Type_enum_Promise{p.Pipeline} }

// Type_enum_Promise is a wrapper for a Type_enum promised by a client call.
type Type_enum_Promise struct{ *capnp.Pipeline }

func (p Type_enum_Promise) Struct() (Type_enum, error) {
	s, err := p.Pipeline.Struct()
	return Type_enum{s}, err
}

func (p Type_enum_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Type_Promise) StructType() Type_structType_Promise { return Type_structType_Promise{p.Pipeline} }

// Type_structType_Promise is a wrapper for a Type_structType promised by a client call.
type Type_structType_Promise struct{ *capnp.Pipeline }

func (p Type_structType_Promise) Struct() (Type_structType, error) {
	s, err := p.Pipeline.Struct()
	return Type_structType{s}, err
}

func (p Type_structType_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Type_Promise) Interface() Type_interface_Promise { return Type_interface_Promise{p.Pipeline} }

// Type_interface_Promise is a wrapper for a Type_interface promised by a client call.
type Type_interface_Promise struct{ *capnp.Pipeline }

func (p Type_interface_Promise) Struct() (Type_interface, error) {
	s, err := p.Pipeline.Struct()
	return Type_interface{s}, err
}

func (p Type_interface_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Type_Promise) AnyPointer() Type_anyPointer_Promise { return Type_anyPointer_Promise{p.Pipeline} }

// Type_anyPointer_Promise is a wrapper for a Type_anyPointer promised by a client call.
type Type_anyPointer_Promise struct{ *capnp.Pipeline }

func (p Type_anyPointer_Promise) Struct() (Type_anyPointer, error) {
	s, err := p.Pipeline.Struct()
	return Type_anyPointer{s}, err
}

func (p Type_anyPointer_Promise) Unconstrained() Type_anyPointer_unconstrained_Promise {
	return Type_anyPointer_unconstrained_Promise{p.Pipeline}
}

// Type_anyPointer_unconstrained_Promise is a wrapper for a Type_anyPointer_unconstrained promised by a client call.
type Type_anyPointer_unconstrained_Promise struct{ *capnp.Pipeline }

func (p Type_anyPointer_unconstrained_Promise) Struct() (Type_anyPointer_unconstrained, error) {
	s, err := p.Pipeline.Struct()
	return Type_anyPointer_unconstrained{s}, err
}

func (p Type_anyPointer_Promise) Parameter() Type_anyPointer_parameter_Promise {
	return Type_anyPointer_parameter_Promise{p.Pipeline}
}

// Type_anyPointer_parameter_Promise is a wrapper for a Type_anyPointer_parameter promised by a client call.
type Type_anyPointer_parameter_Promise struct{ *capnp.Pipeline }

func (p Type_anyPointer_parameter_Promise) Struct() (Type_anyPointer_parameter, error) {
	s, err := p.Pipeline.Struct()
	return Type_anyPointer_parameter{s}, err
}

func (p Type_anyPointer_Promise) ImplicitMethodParameter() Type_anyPointer_implicitMethodParameter_Promise {
	return Type_anyPointer_implicitMethodParameter_Promise{p.Pipeline}
}

// Type_anyPointer_implicitMethodParameter_Promise is a wrapper for a Type_anyPointer_implicitMethodParameter promised by a client call.
type Type_anyPointer_implicitMethodParameter_Promise struct{ *capnp.Pipeline }

func (p Type_anyPointer_implicitMethodParameter_Promise) Struct() (Type_anyPointer_implicitMethodParameter, error) {
	s, err := p.Pipeline.Struct()
	return Type_anyPointer_implicitMethodParameter{s}, err
}

type Brand struct{ capnp.Struct }

// Brand_TypeID is the unique identifier for the type Brand.
const Brand_TypeID = 0x903455f06065422b

func NewBrand(s *capnp.Segment) (Brand, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Brand{st}, err
}

func NewRootBrand(s *capnp.Segment) (Brand, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Brand{st}, err
}

func ReadRootBrand(msg *capnp.Message) (Brand, error) {
	root, err := msg.RootPtr()
	return Brand{root.Struct()}, err
}

func (s Brand) Scopes() (Brand_Scope_List, error) {
	p, err := s.Struct.Ptr(0)
	return Brand_Scope_List{List: p.List()}, err
}

func (s Brand) HasScopes() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Brand) SetScopes(v Brand_Scope_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewScopes sets the scopes field to a newly
// allocated Brand_Scope_List, preferring placement in s's segment.
func (s Brand) NewScopes(n int32) (Brand_Scope_List, error) {
	l, err := NewBrand_Scope_List(s.Struct.Segment(), n)
	if err != nil {
		return Brand_Scope_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Brand_List is a list of Brand.
type Brand_List struct{ capnp.List }

// NewBrand creates a new list of Brand.
func NewBrand_List(s *capnp.Segment, sz int32) (Brand_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Brand_List{l}, err
}

func (s Brand_List) At(i int) Brand { return Brand{s.List.Struct(i)} }

func (s Brand_List) Set(i int, v Brand) error { return s.List.SetStruct(i, v.Struct) }

// Brand_Promise is a wrapper for a Brand promised by a client call.
type Brand_Promise struct{ *capnp.Pipeline }

func (p Brand_Promise) Struct() (Brand, error) {
	s, err := p.Pipeline.Struct()
	return Brand{s}, err
}

type Brand_Scope struct{ capnp.Struct }
type Brand_Scope_Which uint16

const (
	Brand_Scope_Which_bind    Brand_Scope_Which = 0
	Brand_Scope_Which_inherit Brand_Scope_Which = 1
)

func (w Brand_Scope_Which) String() string {
	const s = "bindinherit"
	switch w {
	case Brand_Scope_Which_bind:
		return s[0:4]
	case Brand_Scope_Which_inherit:
		return s[4:11]

	}
	return "Brand_Scope_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Brand_Scope_TypeID is the unique identifier for the type Brand_Scope.
const Brand_Scope_TypeID = 0xabd73485a9636bc9

func NewBrand_Scope(s *capnp.Segment) (Brand_Scope, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Brand_Scope{st}, err
}

func NewRootBrand_Scope(s *capnp.Segment) (Brand_Scope, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Brand_Scope{st}, err
}

func ReadRootBrand_Scope(msg *capnp.Message) (Brand_Scope, error) {
	root, err := msg.RootPtr()
	return Brand_Scope{root.Struct()}, err
}

func (s Brand_Scope) Which() Brand_Scope_Which {
	return Brand_Scope_Which(s.Struct.Uint16(8))
}
func (s Brand_Scope) ScopeId() uint64 {
	return s.Struct.Uint64(0)
}

func (s Brand_Scope) SetScopeId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Brand_Scope) Bind() (Brand_Binding_List, error) {
	p, err := s.Struct.Ptr(0)
	return Brand_Binding_List{List: p.List()}, err
}

func (s Brand_Scope) HasBind() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Brand_Scope) SetBind(v Brand_Binding_List) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewBind sets the bind field to a newly
// allocated Brand_Binding_List, preferring placement in s's segment.
func (s Brand_Scope) NewBind(n int32) (Brand_Binding_List, error) {
	s.Struct.SetUint16(8, 0)
	l, err := NewBrand_Binding_List(s.Struct.Segment(), n)
	if err != nil {
		return Brand_Binding_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Brand_Scope) SetInherit() {
	s.Struct.SetUint16(8, 1)

}

// Brand_Scope_List is a list of Brand_Scope.
type Brand_Scope_List struct{ capnp.List }

// NewBrand_Scope creates a new list of Brand_Scope.
func NewBrand_Scope_List(s *capnp.Segment, sz int32) (Brand_Scope_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Brand_Scope_List{l}, err
}

func (s Brand_Scope_List) At(i int) Brand_Scope { return Brand_Scope{s.List.Struct(i)} }

func (s Brand_Scope_List) Set(i int, v Brand_Scope) error { return s.List.SetStruct(i, v.Struct) }

// Brand_Scope_Promise is a wrapper for a Brand_Scope promised by a client call.
type Brand_Scope_Promise struct{ *capnp.Pipeline }

func (p Brand_Scope_Promise) Struct() (Brand_Scope, error) {
	s, err := p.Pipeline.Struct()
	return Brand_Scope{s}, err
}

type Brand_Binding struct{ capnp.Struct }
type Brand_Binding_Which uint16

const (
	Brand_Binding_Which_unbound Brand_Binding_Which = 0
	Brand_Binding_Which_type    Brand_Binding_Which = 1
)

func (w Brand_Binding_Which) String() string {
	const s = "unboundtype"
	switch w {
	case Brand_Binding_Which_unbound:
		return s[0:7]
	case Brand_Binding_Which_type:
		return s[7:11]

	}
	return "Brand_Binding_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Brand_Binding_TypeID is the unique identifier for the type Brand_Binding.
const Brand_Binding_TypeID = 0xc863cd16969ee7fc

func NewBrand_Binding(s *capnp.Segment) (Brand_Binding, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Brand_Binding{st}, err
}

func NewRootBrand_Binding(s *capnp.Segment) (Brand_Binding, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Brand_Binding{st}, err
}

func ReadRootBrand_Binding(msg *capnp.Message) (Brand_Binding, error) {
	root, err := msg.RootPtr()
	return Brand_Binding{root.Struct()}, err
}

func (s Brand_Binding) Which() Brand_Binding_Which {
	return Brand_Binding_Which(s.Struct.Uint16(0))
}
func (s Brand_Binding) SetUnbound() {
	s.Struct.SetUint16(0, 0)

}

func (s Brand_Binding) Type() (Type, error) {
	p, err := s.Struct.Ptr(0)
	return Type{Struct: p.Struct()}, err
}

func (s Brand_Binding) HasType() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Brand_Binding) SetType(v Type) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Brand_Binding) NewType() (Type, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Brand_Binding_List is a list of Brand_Binding.
type Brand_Binding_List struct{ capnp.List }

// NewBrand_Binding creates a new list of Brand_Binding.
func NewBrand_Binding_List(s *capnp.Segment, sz int32) (Brand_Binding_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Brand_Binding_List{l}, err
}

func (s Brand_Binding_List) At(i int) Brand_Binding { return Brand_Binding{s.List.Struct(i)} }

func (s Brand_Binding_List) Set(i int, v Brand_Binding) error { return s.List.SetStruct(i, v.Struct) }

// Brand_Binding_Promise is a wrapper for a Brand_Binding promised by a client call.
type Brand_Binding_Promise struct{ *capnp.Pipeline }

func (p Brand_Binding_Promise) Struct() (Brand_Binding, error) {
	s, err := p.Pipeline.Struct()
	return Brand_Binding{s}, err
}

func (p Brand_Binding_Promise) Type() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Value struct{ capnp.Struct }
type Value_Which uint16

const (
	Value_Which_void        Value_Which = 0
	Value_Which_bool        Value_Which = 1
	Value_Which_int8        Value_Which = 2
	Value_Which_int16       Value_Which = 3
	Value_Which_int32       Value_Which = 4
	Value_Which_int64       Value_Which = 5
	Value_Which_uint8       Value_Which = 6
	Value_Which_uint16      Value_Which = 7
	Value_Which_uint32      Value_Which = 8
	Value_Which_uint64      Value_Which = 9
	Value_Which_float32     Value_Which = 10
	Value_Which_float64     Value_Which = 11
	Value_Which_text        Value_Which = 12
	Value_Which_data        Value_Which = 13
	Value_Which_list        Value_Which = 14
	Value_Which_enum        Value_Which = 15
	Value_Which_structValue Value_Which = 16
	Value_Which_interface   Value_Which = 17
	Value_Which_anyPointer  Value_Which = 18
)

func (w Value_Which) String() string {
	const s = "voidboolint8int16int32int64uint8uint16uint32uint64float32float64textdatalistenumstructValueinterfaceanyPointer"
	switch w {
	case Value_Which_void:
		return s[0:4]
	case Value_Which_bool:
		return s[4:8]
	case Value_Which_int8:
		return s[8:12]
	case Value_Which_int16:
		return s[12:17]
	case Value_Which_int32:
		return s[17:22]
	case Value_Which_int64:
		return s[22:27]
	case Value_Which_uint8:
		return s[27:32]
	case Value_Which_uint16:
		return s[32:38]
	case Value_Which_uint32:
		return s[38:44]
	case Value_Which_uint64:
		return s[44:50]
	case Value_Which_float32:
		return s[50:57]
	case Value_Which_float64:
		return s[57:64]
	case Value_Which_text:
		return s[64:68]
	case Value_Which_data:
		return s[68:72]
	case Value_Which_list:
		return s[72:76]
	case Value_Which_enum:
		return s[76:80]
	case Value_Which_structValue:
		return s[80:91]
	case Value_Which_interface:
		return s[91:100]
	case Value_Which_anyPointer:
		return s[100:110]

	}
	return "Value_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Value_TypeID is the unique identifier for the type Value.
const Value_TypeID = 0xce23dcd2d7b00c9b

func NewValue(s *capnp.Segment) (Value, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Value{st}, err
}

func NewRootValue(s *capnp.Segment) (Value, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Value{st}, err
}

func ReadRootValue(msg *capnp.Message) (Value, error) {
	root, err := msg.RootPtr()
	return Value{root.Struct()}, err
}

func (s Value) Which() Value_Which {
	return Value_Which(s.Struct.Uint16(0))
}
func (s Value) SetVoid() {
	s.Struct.SetUint16(0, 0)

}

func (s Value) Bool() bool {
	return s.Struct.Bit(16)
}

func (s Value) SetBool(v bool) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetBit(16, v)
}

func (s Value) Int8() int8 {
	return int8(s.Struct.Uint8(2))
}

func (s Value) SetInt8(v int8) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint8(2, uint8(v))
}

func (s Value) Int16() int16 {
	return int16(s.Struct.Uint16(2))
}

func (s Value) SetInt16(v int16) {
	s.Struct.SetUint16(0, 3)
	s.Struct.SetUint16(2, uint16(v))
}

func (s Value) Int32() int32 {
	return int32(s.Struct.Uint32(4))
}

func (s Value) SetInt32(v int32) {
	s.Struct.SetUint16(0, 4)
	s.Struct.SetUint32(4, uint32(v))
}

func (s Value) Int64() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Value) SetInt64(v int64) {
	s.Struct.SetUint16(0, 5)
	s.Struct.SetUint64(8, uint64(v))
}

func (s Value) Uint8() uint8 {
	return s.Struct.Uint8(2)
}

func (s Value) SetUint8(v uint8) {
	s.Struct.SetUint16(0, 6)
	s.Struct.SetUint8(2, v)
}

func (s Value) Uint16() uint16 {
	return s.Struct.Uint16(2)
}

func (s Value) SetUint16(v uint16) {
	s.Struct.SetUint16(0, 7)
	s.Struct.SetUint16(2, v)
}

func (s Value) Uint32() uint32 {
	return s.Struct.Uint32(4)
}

func (s Value) SetUint32(v uint32) {
	s.Struct.SetUint16(0, 8)
	s.Struct.SetUint32(4, v)
}

func (s Value) Uint64() uint64 {
	return s.Struct.Uint64(8)
}

func (s Value) SetUint64(v uint64) {
	s.Struct.SetUint16(0, 9)
	s.Struct.SetUint64(8, v)
}

func (s Value) Float32() float32 {
	return math.Float32frombits(s.Struct.Uint32(4))
}

func (s Value) SetFloat32(v float32) {
	s.Struct.SetUint16(0, 10)
	s.Struct.SetUint32(4, math.Float32bits(v))
}

func (s Value) Float64() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Value) SetFloat64(v float64) {
	s.Struct.SetUint16(0, 11)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Value) Text() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Value) HasText() bool {
	if s.Struct.Uint16(0) != 12 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Value) SetText(v string) error {
	s.Struct.SetUint16(0, 12)
	return s.Struct.SetText(0, v)
}

func (s Value) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Value) HasData() bool {
	if s.Struct.Uint16(0) != 13 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) SetData(v []byte) error {
	s.Struct.SetUint16(0, 13)
	return s.Struct.SetData(0, v)
}

func (s Value) List() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Value) HasList() bool {
	if s.Struct.Uint16(0) != 14 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) ListPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Value) SetList(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 14)
	return s.Struct.SetPointer(0, v)
}

func (s Value) SetListPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 14)
	return s.Struct.SetPtr(0, v)
}

func (s Value) Enum() uint16 {
	return s.Struct.Uint16(2)
}

func (s Value) SetEnum(v uint16) {
	s.Struct.SetUint16(0, 15)
	s.Struct.SetUint16(2, v)
}

func (s Value) StructValue() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Value) HasStructValue() bool {
	if s.Struct.Uint16(0) != 16 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) StructValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Value) SetStructValue(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 16)
	return s.Struct.SetPointer(0, v)
}

func (s Value) SetStructValuePtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 16)
	return s.Struct.SetPtr(0, v)
}

func (s Value) SetInterface() {
	s.Struct.SetUint16(0, 17)

}

func (s Value) AnyPointer() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Value) HasAnyPointer() bool {
	if s.Struct.Uint16(0) != 18 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) AnyPointerPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Value) SetAnyPointer(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 18)
	return s.Struct.SetPointer(0, v)
}

func (s Value) SetAnyPointerPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 18)
	return s.Struct.SetPtr(0, v)
}

// Value_List is a list of Value.
type Value_List struct{ capnp.List }

// NewValue creates a new list of Value.
func NewValue_List(s *capnp.Segment, sz int32) (Value_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Value_List{l}, err
}

func (s Value_List) At(i int) Value { return Value{s.List.Struct(i)} }

func (s Value_List) Set(i int, v Value) error { return s.List.SetStruct(i, v.Struct) }

// Value_Promise is a wrapper for a Value promised by a client call.
type Value_Promise struct{ *capnp.Pipeline }

func (p Value_Promise) Struct() (Value, error) {
	s, err := p.Pipeline.Struct()
	return Value{s}, err
}

func (p Value_Promise) List() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p Value_Promise) StructValue() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p Value_Promise) AnyPointer() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Annotation struct{ capnp.Struct }

// Annotation_TypeID is the unique identifier for the type Annotation.
const Annotation_TypeID = 0xf1c8950dab257542

func NewAnnotation(s *capnp.Segment) (Annotation, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Annotation{st}, err
}

func NewRootAnnotation(s *capnp.Segment) (Annotation, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Annotation{st}, err
}

func ReadRootAnnotation(msg *capnp.Message) (Annotation, error) {
	root, err := msg.RootPtr()
	return Annotation{root.Struct()}, err
}

func (s Annotation) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s Annotation) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Annotation) Brand() (Brand, error) {
	p, err := s.Struct.Ptr(1)
	return Brand{Struct: p.Struct()}, err
}

func (s Annotation) HasBrand() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Annotation) SetBrand(v Brand) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Annotation) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Annotation) Value() (Value, error) {
	p, err := s.Struct.Ptr(0)
	return Value{Struct: p.Struct()}, err
}

func (s Annotation) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Annotation) SetValue(v Value) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewValue sets the value field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Annotation) NewValue() (Value, error) {
	ss, err := NewValue(s.Struct.Segment())
	if err != nil {
		return Value{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Annotation_List is a list of Annotation.
type Annotation_List struct{ capnp.List }

// NewAnnotation creates a new list of Annotation.
func NewAnnotation_List(s *capnp.Segment, sz int32) (Annotation_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Annotation_List{l}, err
}

func (s Annotation_List) At(i int) Annotation { return Annotation{s.List.Struct(i)} }

func (s Annotation_List) Set(i int, v Annotation) error { return s.List.SetStruct(i, v.Struct) }

// Annotation_Promise is a wrapper for a Annotation promised by a client call.
type Annotation_Promise struct{ *capnp.Pipeline }

func (p Annotation_Promise) Struct() (Annotation, error) {
	s, err := p.Pipeline.Struct()
	return Annotation{s}, err
}

func (p Annotation_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Annotation_Promise) Value() Value_Promise {
	return Value_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type ElementSize uint16

// Values of ElementSize.
const (
	ElementSize_empty           ElementSize = 0
	ElementSize_bit             ElementSize = 1
	ElementSize_byte            ElementSize = 2
	ElementSize_twoBytes        ElementSize = 3
	ElementSize_fourBytes       ElementSize = 4
	ElementSize_eightBytes      ElementSize = 5
	ElementSize_pointer         ElementSize = 6
	ElementSize_inlineComposite ElementSize = 7
)

// String returns the enum's constant name.
func (c ElementSize) String() string {
	switch c {
	case ElementSize_empty:
		return "empty"
	case ElementSize_bit:
		return "bit"
	case ElementSize_byte:
		return "byte"
	case ElementSize_twoBytes:
		return "twoBytes"
	case ElementSize_fourBytes:
		return "fourBytes"
	case ElementSize_eightBytes:
		return "eightBytes"
	case ElementSize_pointer:
		return "pointer"
	case ElementSize_inlineComposite:
		return "inlineComposite"

	default:
		return ""
	}
}

// ElementSizeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ElementSizeFromString(c string) ElementSize {
	switch c {
	case "empty":
		return ElementSize_empty
	case "bit":
		return ElementSize_bit
	case "byte":
		return ElementSize_byte
	case "twoBytes":
		return ElementSize_twoBytes
	case "fourBytes":
		return ElementSize_fourBytes
	case "eightBytes":
		return ElementSize_eightBytes
	case "pointer":
		return ElementSize_pointer
	case "inlineComposite":
		return ElementSize_inlineComposite

	default:
		return 0
	}
}

type ElementSize_List struct{ capnp.List }

func NewElementSize_List(s *capnp.Segment, sz int32) (ElementSize_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return ElementSize_List{l.List}, err
}

func (l ElementSize_List) At(i int) ElementSize {
	ul := capnp.UInt16List{List: l.List}
	return ElementSize(ul.At(i))
}

func (l ElementSize_List) Set(i int, v ElementSize) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type CodeGeneratorRequest struct{ capnp.Struct }

// CodeGeneratorRequest_TypeID is the unique identifier for the type CodeGeneratorRequest.
const CodeGeneratorRequest_TypeID = 0xbfc546f6210ad7ce

func NewCodeGeneratorRequest(s *capnp.Segment) (CodeGeneratorRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return CodeGeneratorRequest{st}, err
}

func NewRootCodeGeneratorRequest(s *capnp.Segment) (CodeGeneratorRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return CodeGeneratorRequest{st}, err
}

func ReadRootCodeGeneratorRequest(msg *capnp.Message) (CodeGeneratorRequest, error) {
	root, err := msg.RootPtr()
	return CodeGeneratorRequest{root.Struct()}, err
}

func (s CodeGeneratorRequest) Nodes() (Node_List, error) {
	p, err := s.Struct.Ptr(0)
	return Node_List{List: p.List()}, err
}

func (s CodeGeneratorRequest) HasNodes() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequest) SetNodes(v Node_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewNodes sets the nodes field to a newly
// allocated Node_List, preferring placement in s's segment.
func (s CodeGeneratorRequest) NewNodes(n int32) (Node_List, error) {
	l, err := NewNode_List(s.Struct.Segment(), n)
	if err != nil {
		return Node_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s CodeGeneratorRequest) RequestedFiles() (CodeGeneratorRequest_RequestedFile_List, error) {
	p, err := s.Struct.Ptr(1)
	return CodeGeneratorRequest_RequestedFile_List{List: p.List()}, err
}

func (s CodeGeneratorRequest) HasRequestedFiles() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequest) SetRequestedFiles(v CodeGeneratorRequest_RequestedFile_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequestedFiles sets the requestedFiles field to a newly
// allocated CodeGeneratorRequest_RequestedFile_List, preferring placement in s's segment.
func (s CodeGeneratorRequest) NewRequestedFiles(n int32) (CodeGeneratorRequest_RequestedFile_List, error) {
	l, err := NewCodeGeneratorRequest_RequestedFile_List(s.Struct.Segment(), n)
	if err != nil {
		return CodeGeneratorRequest_RequestedFile_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// CodeGeneratorRequest_List is a list of CodeGeneratorRequest.
type CodeGeneratorRequest_List struct{ capnp.List }

// NewCodeGeneratorRequest creates a new list of CodeGeneratorRequest.
func NewCodeGeneratorRequest_List(s *capnp.Segment, sz int32) (CodeGeneratorRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return CodeGeneratorRequest_List{l}, err
}

func (s CodeGeneratorRequest_List) At(i int) CodeGeneratorRequest {
	return CodeGeneratorRequest{s.List.Struct(i)}
}

func (s CodeGeneratorRequest_List) Set(i int, v CodeGeneratorRequest) error {
	return s.List.SetStruct(i, v.Struct)
}

// CodeGeneratorRequest_Promise is a wrapper for a CodeGeneratorRequest promised by a client call.
type CodeGeneratorRequest_Promise struct{ *capnp.Pipeline }

func (p CodeGeneratorRequest_Promise) Struct() (CodeGeneratorRequest, error) {
	s, err := p.Pipeline.Struct()
	return CodeGeneratorRequest{s}, err
}

type CodeGeneratorRequest_RequestedFile struct{ capnp.Struct }

// CodeGeneratorRequest_RequestedFile_TypeID is the unique identifier for the type CodeGeneratorRequest_RequestedFile.
const CodeGeneratorRequest_RequestedFile_TypeID = 0xcfea0eb02e810062

func NewCodeGeneratorRequest_RequestedFile(s *capnp.Segment) (CodeGeneratorRequest_RequestedFile, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return CodeGeneratorRequest_RequestedFile{st}, err
}

func NewRootCodeGeneratorRequest_RequestedFile(s *capnp.Segment) (CodeGeneratorRequest_RequestedFile, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return CodeGeneratorRequest_RequestedFile{st}, err
}

func ReadRootCodeGeneratorRequest_RequestedFile(msg *capnp.Message) (CodeGeneratorRequest_RequestedFile, error) {
	root, err := msg.RootPtr()
	return CodeGeneratorRequest_RequestedFile{root.Struct()}, err
}

func (s CodeGeneratorRequest_RequestedFile) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s CodeGeneratorRequest_RequestedFile) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s CodeGeneratorRequest_RequestedFile) Filename() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s CodeGeneratorRequest_RequestedFile) HasFilename() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequest_RequestedFile) FilenameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s CodeGeneratorRequest_RequestedFile) SetFilename(v string) error {
	return s.Struct.SetText(0, v)
}

func (s CodeGeneratorRequest_RequestedFile) Imports() (CodeGeneratorRequest_RequestedFile_Import_List, error) {
	p, err := s.Struct.Ptr(1)
	return CodeGeneratorRequest_RequestedFile_Import_List{List: p.List()}, err
}

func (s CodeGeneratorRequest_RequestedFile) HasImports() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequest_RequestedFile) SetImports(v CodeGeneratorRequest_RequestedFile_Import_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewImports sets the imports field to a newly
// allocated CodeGeneratorRequest_RequestedFile_Import_List, preferring placement in s's segment.
func (s CodeGeneratorRequest_RequestedFile) NewImports(n int32) (CodeGeneratorRequest_RequestedFile_Import_List, error) {
	l, err := NewCodeGeneratorRequest_RequestedFile_Import_List(s.Struct.Segment(), n)
	if err != nil {
		return CodeGeneratorRequest_RequestedFile_Import_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// CodeGeneratorRequest_RequestedFile_List is a list of CodeGeneratorRequest_RequestedFile.
type CodeGeneratorRequest_RequestedFile_List struct{ capnp.List }

// NewCodeGeneratorRequest_RequestedFile creates a new list of CodeGeneratorRequest_RequestedFile.
func NewCodeGeneratorRequest_RequestedFile_List(s *capnp.Segment, sz int32) (CodeGeneratorRequest_RequestedFile_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return CodeGeneratorRequest_RequestedFile_List{l}, err
}

func (s CodeGeneratorRequest_RequestedFile_List) At(i int) CodeGeneratorRequest_RequestedFile {
	return CodeGeneratorRequest_RequestedFile{s.List.Struct(i)}
}

func (s CodeGeneratorRequest_RequestedFile_List) Set(i int, v CodeGeneratorRequest_RequestedFile) error {
	return s.List.SetStruct(i, v.Struct)
}

// CodeGeneratorRequest_RequestedFile_Promise is a wrapper for a CodeGeneratorRequest_RequestedFile promised by a client call.
type CodeGeneratorRequest_RequestedFile_Promise struct{ *capnp.Pipeline }

func (p CodeGeneratorRequest_RequestedFile_Promise) Struct() (CodeGeneratorRequest_RequestedFile, error) {
	s, err := p.Pipeline.Struct()
	return CodeGeneratorRequest_RequestedFile{s}, err
}

type CodeGeneratorRequest_RequestedFile_Import struct{ capnp.Struct }

// CodeGeneratorRequest_RequestedFile_Import_TypeID is the unique identifier for the type CodeGeneratorRequest_RequestedFile_Import.
const CodeGeneratorRequest_RequestedFile_Import_TypeID = 0xae504193122357e5

func NewCodeGeneratorRequest_RequestedFile_Import(s *capnp.Segment) (CodeGeneratorRequest_RequestedFile_Import, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CodeGeneratorRequest_RequestedFile_Import{st}, err
}

func NewRootCodeGeneratorRequest_RequestedFile_Import(s *capnp.Segment) (CodeGeneratorRequest_RequestedFile_Import, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CodeGeneratorRequest_RequestedFile_Import{st}, err
}

func ReadRootCodeGeneratorRequest_RequestedFile_Import(msg *capnp.Message) (CodeGeneratorRequest_RequestedFile_Import, error) {
	root, err := msg.RootPtr()
	return CodeGeneratorRequest_RequestedFile_Import{root.Struct()}, err
}

func (s CodeGeneratorRequest_RequestedFile_Import) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s CodeGeneratorRequest_RequestedFile_Import) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s CodeGeneratorRequest_RequestedFile_Import) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s CodeGeneratorRequest_RequestedFile_Import) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequest_RequestedFile_Import) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s CodeGeneratorRequest_RequestedFile_Import) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

// CodeGeneratorRequest_RequestedFile_Import_List is a list of CodeGeneratorRequest_RequestedFile_Import.
type CodeGeneratorRequest_RequestedFile_Import_List struct{ capnp.List }

// NewCodeGeneratorRequest_RequestedFile_Import creates a new list of CodeGeneratorRequest_RequestedFile_Import.
func NewCodeGeneratorRequest_RequestedFile_Import_List(s *capnp.Segment, sz int32) (CodeGeneratorRequest_RequestedFile_Import_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return CodeGeneratorRequest_RequestedFile_Import_List{l}, err
}

func (s CodeGeneratorRequest_RequestedFile_Import_List) At(i int) CodeGeneratorRequest_RequestedFile_Import {
	return CodeGeneratorRequest_RequestedFile_Import{s.List.Struct(i)}
}

func (s CodeGeneratorRequest_RequestedFile_Import_List) Set(i int, v CodeGeneratorRequest_RequestedFile_Import) error {
	return s.List.SetStruct(i, v.Struct)
}

// CodeGeneratorRequest_RequestedFile_Import_Promise is a wrapper for a CodeGeneratorRequest_RequestedFile_Import promised by a client call.
type CodeGeneratorRequest_RequestedFile_Import_Promise struct{ *capnp.Pipeline }

func (p CodeGeneratorRequest_RequestedFile_Import_Promise) Struct() (CodeGeneratorRequest_RequestedFile_Import, error) {
	s, err := p.Pipeline.Struct()
	return CodeGeneratorRequest_RequestedFile_Import{s}, err
}

const schema_a93fc509624c72d9 = "x\xda\xacY}\x94\x14\xd5\x95\xbf\xf7U\x7f\xccG\x17" +
	"\xdd5\xaf\x10\x8784\xa8\xa8\x10\x9d\xc0\x0c\xb08\x91" +
	"\x1d\x18\x18\x0c,\x90)\x1aP8\xeb\x095\xd35L" +
	"\x99\x9e\xea\xa6\xbaZi\x16\xcf\x80'\x1e\xb3\xee\xba~" +
	"\xacD\x97$\x9el\x96\x9c\x13W\xdd\xc0\xae\x9c\x13\x89" +
	"Y\xcd\x1cY\x81\x83\x89\xee\x9acHt7xBt" +
	"\xd9e\x85]P\x91\x8f\xdas_UW\xd74=\xd1" +
	"\x9c\xf8WW\xdf\xfb\xea\xbd\xfb\xee\xbd\xbf\xfbU\xb3>" +
	"\xdf\xbc\x90\xcd\x8e\xa6'\x00h\xbb\xa31\xf7\xf1\x13\x1b" +
	"\x9b\xa6\xdf\xfc\xee\xfd\xa0\xb5\xa0\xe4n<r\xee\xd8\xc9" +
	"-\xc5\xd7`\"\xc6\x11\x80\x9f\x8d\xed\x07\xe4gc\xdd" +
	"\x80\xee\xbay\xeb/\x95\xbf\xf2\xc5\xbf\x02\xad\x0d%\xf7" +
	"\xf4\xb2?\xfe\xee{\xdd\xb7\x8d\xc2Z\x8cc\x04\xa3\x9d" +
	"\xad\xf1\x0d\x08\xc8\xa7\xc7\xdf\x05t?\xdfcl<\xb5" +
	"v\xce\xc3\xa0\xc8\xe8\x1e\xb5W\xf47\x1e\xe8~\x0a\xa2" +
	"\x18\x07\xe0\x17\xe3\xbbx\xb4\xe1z\x00\xde\xda\xd0\x0d\xf8" +
	"\xe2\xf6\x95\x9d_|\xe7\xc8NMF)\xb44JK" +
	"\x975|\x97k\x0dq\x80\xce\x95\x0d\xaf \xa0\xbbr" +
	"\xd7\x9a\xb7\xfek\xdb\x03\x8f\x83&cxcF\xab\xd7" +
	"6\xed\xe7w4\xd1\xd3\xfa\xa6\x1f\x00\xba-\xaf\\\xda" +
	"\xf6\xcf+\xf6>\x0e\x0a\x8f\xb8_9\xb3<\xbd\xbb\xef" +
	"\xd9]\x00\xd8\xd9\xd8\xdc\x82\xbc\xb59\x0e\x90Q\x9b%" +
	"\xccLmf\x08P]2V\x94\xdeH\x9ca\x84+" +
	"\xcd\xbbxk\xf3$\x80\xce\xe9\xcd\x0f\x91,\xf7-\xd7" +
	"\xcf\\\xfb\xe1\xeb\xdf\xae\xd1\x87\xa7\xb9\xcer\xa2\x8b\xd4" +
	"\xb1#q7\xa0\xfb\xd4\xfd#\x937o\x99\xf0d}" +
	"%\x1fK\x90\x92\x8f\x89\x95s\x97\x9c\xff\xa3o\xee\xfd" +
	"\x8eX\x19u'\xed\xbe\xe6\xdc\x8a\xa7\xef\xfd-L\x8c" +
	"\x89\x95\x0b\xe4\xc3\x80\x9d\x8bdq\xfeGo\xee~\xed" +
	"\xc9\x99\xdfx\xaaV\x17B\xc9\xe7&\x8crL\x0au" +
	"O\xa0}\x0f}u\xe0\xa9\xfb\xe6\xbc\xf94h\x1cY" +
	"\xd5<\xbd(\xee\xb6>y\x98\x1bb\xb5\x9e$\xcd\xfd" +
	"\xdb\xbf\xacx\x7fc\xbe\xeb\x99\xfa\xf26\xa6\x0e\x03r" +
	"9E\xfb\x1e\xbf\xed\x9a\x96\xbf^\xd4\xf7\x0f\xa0\xcdB" +
	"\xbc\xd8\xbf\xa3}\xcf\x84\x13?\x13\"t\x9a\xa9\xfd\xc8" +
	"w\xa4h\xd7{\xc4\xda\xa9\x8f\xc9[\x9f\xf9\xde\x03{" +
	"\xeb\xdf\xedhj\x14\x90\x1f\x15+\x9f8\xdb\xb9x\xde" +
	"?-\xdfW\x7f\xe5\\\x85\xf45W!\xa7\xdc\xfbw" +
	"\xc9\xd3G>\xb7\xfayPZ\xb0\xba\xd0\xd3\x81\xa9\xbc" +
	"\xc3K\x0a=m\x16k\x9dw\xd6%Z\x0e\xbd\xbf\xbf" +
	"\xbe\xc1\x1eQ\xfe\x9e\x0c\xf6\xb7b\xe9o\x9b\xef\xbf\x7f" +
	"\xf4\xe7\x0f\xff\x88\xd4%U\x1dcm$\x8e\x0c\xa3\xfc" +
	"\x80\xf2K@~H!Y\x7f\xfaf\xd3\xb4\x0f\x96\x1e" +
	"x\xb1\xc6\xd5\xc9#;ojiA\xbe\xa0e\x12\xb9" +
	"r\x0b-\x0e\x0e\x1d\xabX\xc2\x90\x84\x11~\xa8\xe5\x04" +
	" \x7f\xb5\x85l\x90w\x9e\xfb\xea\xad\xd1k^\xae\x11" +
	"abD\xe8`\x1a'mM\xe3\x04\xb6\x0b\xef>\xf9" +
	"\x8d+^\x1d8H+q\xacm\x018\xaa\xbf\xe4\xb2" +
	"JZhTI\x04\xa5\xedWC\xbfz\xf5\xc2\xe1\xfa" +
	"\xfb\x1a*\xd9\xd6TI\x09\xdfL\xecy\xf3_\xdf\xba" +
	"\xe6\xa7\xe4`,\x84\x07\x8cs\x00\xfe\xe7\xea.\xfe\x08" +
	"\xed\xdb\xf9\xa0\xfa\x85\x08\x04\xc6\xd7\xae\xc6\x90R<=" +
	"Ll\xbd\x17\xf9\x8cV\xd2\xc3\xecV\xba\\p\xf3\x1a" +
	"\xa8y[\xbf\xdd\xfa(?\xdeJ/\x1ek\xa5\xad\xdd" +
	"\xebZW\xef\xbf\xe7\xa1\x9d\xaf\x83\"\x87\x04\x01\xe4o" +
	"\xb7\x1d\xe6\xff\xd9F\x97;\xde\xf6\x0a\xa0\xdb3\x9a\xfc" +
	"\xf8Gk_\xfcw\xd2\xefe\xfe\xb0~\xca\x09nL" +
	"\x11^>\x854\x11\xb05\x19\xa3!!b\xf1\x18\xc6" +
	"\xf8\x8f\xa7<\xca\x0fL\xb9\x9e\x84\x982I\x02t\x1f" +
	"\x9a6z\xeag\x99\xeb\xdf\xab\xef\x94\x17\xa7\xbd\x03\xc8" +
	"\xf1j\xda\xf8a\xd6\xb4\xf0\x8d\xd6+\xfe\xbb\xfe\xca\x1d" +
	"W\x9f\x00\xec\xbc\xef\xea\xff`\x80\xee\x8b\x89\xf3\x1f\x9a" +
	"\x87\xff\xe2d}\xa4\xed\x9dN\x9b\xee\x9bN\x9b\xf6\x94" +
	"\xa6?-\xef<x\xban\xe8S\xae\x1b\xe5\xad\xd7\xd1" +
	"\xd3\xc4\xeb~\x00\xab\xdc\xe2\xc0\x901\xac\xb7\x0f\xa0^" +
	"\xb0\x0a]k\xca\x85n\xa3=g\x16\x1d-\"E\x00" +
	"T\x9c\x00\xa0\xc8\xfd\x00ZBB\xedJ\x86\xae\x913" +
	"\x86\x0d\xcbY\x03\xf1r\xc1\xc0TU\x12@L\x01\x06" +
	"\x1bF*\x1b\x1a\xed\xbaU\xee\xcb\x9b\x96c\xd8\xed%" +
	"k o\x15\x1d[7-\xc9\xc8j))\x92p]" +
	"\x15[\x00\x14\xbd\x07@\xfbS\x09\xb5!\x862^r" +
	"Ul\x05P\x8c.\x00m\xa3\x84Z\x8e\xa1\xcc.\xba" +
	"*N\x06P\xcc\x99\x00ZVB\xad\xc0P\x96.\xb8" +
	"*~\x0e@\x19\xde\x00\xa0\xe5$\xd4\xb60\x1c\xd1\xad" +
	"\xf2\x9f\x98V\x16b\xddE\xc7.\x0d8\x10K\xd2\xbd" +
	" \xe6\x0e\xe8\x05\xbd\xdf\xcc\x99 9e\x88\xd5h\xa0" +
	"\xc7\xd6%+\xab5`(\x12*\x8d\x1dU\xe8(\xd1" +
	"\x9etf _0FzL+kZ\x9b<ME" +
	"\x90\x14E\xc26H\xa8]\xcb\xb0\xbbH\x8b\x8a8\x01" +
	"\xb0OBLU\xb7\x03$b\xcd\xb9+\x0d'>\x94" +
	"\xcf\xf6!jS\x83\xfd^\xa7k\x1e\x91P\xfb\x05C" +
	"D\x15\x89\xf6\xf3\xd5\x00\xda\x1b\x12j\xbff\xa8H\xa8" +
	"\"\x03P\xde\xbe\x17@{KB\xed=\x86J\x94\xa9" +
	"(\x01(\xc7\x1f\x00\xd0\xde\x93P;\xc3P\x89\xa3\x8a" +
	"\x11\x00\xe54\xd9\xf2\x94\x84\x99\x042T\"L\xc5(" +
	"\x01\x1e7\x00d\x1aP\xc2\x8cJ\xf4\x98\xa4b\x8c\xbc" +
	"\x05\xfb\x012)\xa2_Et\x16Q\x05@Zq\x17" +
	"@\xe6*\xa2\xcfG\x86IK\x1f60\x01\x0c\x13\x80" +
	"\xee@>k|\xd9\xce\x1a\x806\xc6\x81a\x1c\xd0-" +
	"\xe8\xb6>\x9cql,\x0d8\xe4\x13\x80\x8d\xc0\xb0\x11" +
	"\xd0\xb5\x8db)\xe7d\x1c\xb4+\xac*O\xb7\xac\xbc" +
	"\xa3;&\xc4\xf3VH\x93\x81\x83\xfb\x9a\x14\x9b\xf7\xd8" +
	":HV\x16S\xd5\xb8\xe6{\xa4wB\x8f\x0dq\xbd" +
	".\xdf\x1c.\xe4\xcc\x01\xd3\xc1>\xda\xc7p\x0c\xc9\x0e" +
	"\x1d\x16\xe4\x8d\xbaf\xeb\xb5J\xdd\xc3\x86\xad[\x0eY" +
	".\x11X\xae\x97,\xb7PBmE\xd5r\xcb\xc8r" +
	"_\x92P[C\x9a\xf4-\xa7\x91=\xfa<\xff\xfed" +
	"5~J\x85Td\x94\x84\x8cKM#\x97m\xb7\xf2" +
	"K\xcc\xe2\x80m\x0e\x9b\x96n!\x89K\xbb\xcaq\xd7" +
	"\xbd\xecRKMC\xcae\xb5\x08\x86+#\xdc\xeaV" +
	"\xb6\x80n\xb1\x89\xa3]\x15\xdcw\x1f\xddw\x8f\x84\xda" +
	"\x0b\xd5\xfb>O\xf7\xfd\xa1\x84\xda\xcb\xa1\xfb\xfe\x84\xee" +
	"\xfb\x92\x84\xda[\xbe\xfbJ\x88\xca\xd1G\xab\xee+G" +
	"\\\x17CIM9>\x13\x98\x1c\xbdD\xc4 #)" +
	"\xafv\x00\xc3\x18\x86\xf2\xaf\xb2\xaf\x07\xd8g\xa6\xc1\xac" +
	"\xaf+\xa4{\xae\xd3s%4\xaa\xeaJ\x16sy'" +
	"\xbd\xc9\xce\x97\x0a#y;kZz\xaeF\xe5\xb5a" +
	"\xaf\xa0\xdb\xdd\xc2\xb5l\xadA\x8a\xa4\x98\x8a\x1c@\x99" +
	"A\x11\xefZ\x09\xb5Y\x0c\x15\x8c\xaa\xa8\x02(7m" +
	"\x05\xd0n\x94P\x9b\xcfpD\xc4\x90e\xd9\x00\x10\x05" +
	"\xdfC\xa1\xdb^fe\x8d-\xc1\xb5\xea\x85q\xc3*" +
	"\x0d\x8b\xe3P\xc5$\x1d\xd7U=\x8el4\x91N\xeb" +
	"\x00\xd0n\x90P\x9b\xc3\xb0\xdb)\x87\x0fK\xf7\xdb\xf5" +
	"\xf1R9\x8b\x89\xb3V\xe5\xb3F\xbb\x1fb\x81<\"" +
	"\x15\x171B\xd9g\x03h\xcfI\xa8\xbdD\xb7K\xa8" +
	"\xd8\x00\xa0\xfc\xf8N\x00\xed\x05\x09\xb5\x83\xe4\x13\xb2\x8a" +
	"\x8d\x00\xca\x81\x7f\x04\xd0\x0eJ\xa8\xbdA>\xf1k\x15" +
	"\x9b(\xf6\xf5Tc\x9f\x12I\xaa\xd8L\xc1\x8f\x1c\xe5" +
	"\x17\x12j\xbf\xa18\xd7\xa0b\x02@9\xb6\x0b@\xfb" +
	"\x8d\x84\xda)?r\xc9\x00\xca\xc9./\xf8e\"\xc8" +
	"\xd0\xcd\xea\x8e~[\xde\xceBzq\xbed9\xd5\xb0" +
	"\xe4\xd9g1$\xc7\x92mc\xd0\xb0m\x03\xb3+\xcc" +
	"\xa2\xd3k\x0d\xa4\xf3\x14\xe91Y\xad/\x001\x098" +
	"b\x16o%7@\x04\x86X\xeb7t\x16Vw\x1d" +
	"\xc3\xfb\xf2\xe0`Q2\x1cl\x00\x86\x0d\x80\xdd\x83\x04" +
	"\xd2\x90C\x86\xba\x901\x90\xf6\x94\x9e)\x15\x0c{ " +
	"\xa7\x17\x8b@q\xa7!\xc0\xe1\x8c\xc9cm\x8c56" +
	"\x96\xcc\xdf\xd7\xbe=\xb4\xaa]d<\x80\x9a \xd7S" +
	"\x0dr2R\x1e\x17a\x8eB\xc1\x12\x09\xb5\x8d\x94\xb1" +
	")\x8f\x13\xee\xef\xa0\xb5\xb7K\xa8e/w\xebd\xbf" +
	"ie\xabW\x0fR\xadw\xf5\x11\xd3\x1a2l\xd3\x09" +
	"\xa5jV\x05Y\xe0y\xbe\xa3\xa7\xea8\xfa\x15\x7f\xa0" +
	"\xa3G\xc5q\x8b\xf3Y\xe3V\xc32l\xdd\xc9\xdb\xab" +
	"\x8d\xcd%\xa3\xe8\xb4\xfb\xbfFv\xa9\x993\xda\xbb\x97" +
	"\x0d\x17\xf2\xb6\xf3)L2\xb3\xaeI\xc6\x86\xaf:8" +
	"\x13\xc5\x93wW\xc9\xbf\xeb\xccp\x0c\x89\xa8\xa8\xd4\\" +
	"6\xe9\xd4\xad\xd3\xd2w\xe9\xb9\x12\xd1\x83\xda\xbd\xe6\xd2" +
	"X9\xb5\x12I\"\xde\xa1\xa2 \xdc\xe0\x17\x847P" +
	"Ah\x95D\x06\x04\xc9\x09\xf9o\xd0\x7f\xd7\xf5_q" +
	"\x19/\xdd\xc6\x1d\xc3&\x85\x85\xaa\xa8\x99~\x15\xa5\xb2" +
	"q4\x12\xad\x1bd+y|\xa5\xe1P%E\xbb'" +
	"E\xc8%\xc9\xa3*^I{o\x0d\x95\xb2\x9f\x18M" +
	"Y(\x7f\x8a@/\xe99R\xbe\xa8Y)\x9c\xcdX" +
	"^Ut\x1b^rc^@\x9bM\xe4Y\x12j\xb7" +
	"\xb0jy\x01\x10w\x8d-\x95g\xb8\xec0i</" +
	"C\x87Rq\xa5kR\xd0v+^\x07i\xe1wa" +
	"g\xeb\xf0\xdda!\xb9\x83\xefm\x0b\xe8\xd2\xb7H\xa8" +
	"\xdd\xce0m\xe5\xb3\xe1\xaa4\xe89|;\xd9\x95\xad" +
	"\xbb\xc5\xd6\xd5\x95\x95\xf3\xeb\xda3d\x8a$\xd9\x82\x82" +
	"DB\xe4\xf1`\x16\xa4\xf4\xda\xc0\xa8\xae\xc7\xd0@D" +
	"\x99\xbd\x1a\x18\x95\xf5\x18j\xba\x95\xb6Q`n\xa5S" +
	"\x80\xb4nZF\xb6j,\xb4\x03\x95FC\xa6&K" +
	"CE\xaa\x90\xdd(W\x035\x1a)\xbf\x02\xd6\xbb\xaa" +
	"}\x86\x82^\x01\xac\x183\xabm\x86\xc2\xbc\xeaW1" +
	")Y\x0dI\xa89\x94\x97\xb6{yi3e\x1bG" +
	"Bm;\xc3\xee\xfc\xe0`\xb1\x1a\xc4\xc7\x81\x9a\x9b5" +
	"\x06\xf5R\xceY\x07\xc9q07\xa4g{\xc93p" +
	"\xc0t\x96\xd0b)\xe7\x04)\xa5^8\x16\xad\x87d" +
	"m\xf2#M\xc2u=\xeb\x87\x0a\x0a\xaf\x85\xaa\x0d6" +
	"#%\xab?_\xa2\xb6h<a\xeb\xf9\xbe(u\x00" +
	"\x04\x90\x98\x97\xdd\x83VG\xbd,\x9e\xd6\x84\x90uz" +
	"N*\x19$\xe9\x8a@R\x1ee3\x01V3\xeaC" +
	"\x98\x07\x9d\x94\x10\x967\x12#\x13!N\x8a8\xec\xa2" +
	"\xcbD\x06\xe1\xb2\xe04\x10G%\x8et\xc1E\xd1\xe8" +
	"p\x85u\x00d\x12\xc4\xb9\x928\x91\xf3\xaegk>" +
	"QpR\xc4\xb9\x8a8\xd1\x8f\x89C\x1dO\xab\xe0\xa8" +
	"\xc4\x99J\x9c\xd89:\x87z\x9e6\xc1\xb9\x928\xd7" +
	"\x12'\xfe\x11\xbdC]\xcf4\xd6E]\x0fqn " +
	"N\xc3\x87\xc4i\x00\xe0\xd3\x05g*qn$N\xe3" +
	"\x07\xc4i\x04\xe03\x04\xe7Z\xe2\xcc\"N\xd3Y\xe2" +
	"4\x01\xf0\x9bX\x0f@\xe6\x06\xe2\xcc!N\xf3\x19\xe2" +
	"4\x03\xf0\xd9\x82s#q\xe63\x86r\xe2\xff\\Q" +
	"\xe9\xf0\xb9B\x05\xb3\x88q\x0b1\xe4\xffuE\xb5\xc3" +
	"o\x16\x8c9\xc4XH\x8c\x09\xa7]\x11\xa9\xf9\x02\xc1" +
	"\x98O\x8c%tH\xf2\x94\xebU\x83|\x91\xe0\xdcB" +
	"\x9c/\xd1+\xa9\xf7]\x91Qx/\xeb\xe2\xbd,\x9d" +
	"\x19\"\x96C,\xe5\x7f\\\x91W\xf8f\xb6\x1a S" +
	" \xc66b\xb4\x9c\xf4zw^f\xd4?n!\xc6" +
	"\xd7\x18\xc3\xe4]y\x93\\\xac?\x9f\xcfU\xfc8i" +
	"Z\xce|d\xc0\x90\x01\xa6M\xcb\x99=\x0f%`(" +
	"y\xff:;0\x02\x0c#\xde\xbfys0\x0a\x0c\xa3" +
	"\x80\xe9\x92x/F\xe5>`w\xc9{\xd1\x8f\x9d\xe2" +
	"ogGPD\x95\xbcW}?\x1c\x19\xcc\xe5ub" +
	"7\x01\xc3\xa6\xca\xffys\xb0\x19\x186\x13\\\x8d-" +
	"N%\xb7$\xa9DD\x19\x18\xca\x80bR\x80-\xc0" +
	"\xb0\x050I\x09.8\xd1\xab7\xb4\x082\xf7\xf4\x83" +
	"_\x98\xd4\xb2\xf1\xf9Q\xd0\"\x0c\x17\xa9(\x8aQ\xec" +
	"w\xbd%\xebt\x88\xe7J\xd4\xcd\xfa\xfb\xb8\"K\x0d" +
	"\xea\x03\x80\x06\xc4\xdcJ\xe2\x02\xc9\xb0\x83%\xbfo\xc5" +
	"\x01\xa2I\x0bF\xab\x0av\xf9%H\xb8J\x9b\\\xa7" +
	"\x15]^\xaf\x15\xa5\xe8\xb1\xc2\x0b\x8c\xa1\xa2\xc4\x1d4" +
	"s\x06\xa5aJZ\xbe\xb6FLqJ(\x8f\x042" +
	"\xd4m\x93\xd7\x94\x0b\xf8\xbbC@\x10\xaej#\x807" +
	"\xf4\xa9\x13\x00\xbc\xb9O\x1d\xfc\xcb\x91\xf3n]\xf8\xcb" +
	"\xd1\x8f\xdd\xba\xe8\x97c\xe7\xdc\xba\xe0\x97\xe3\x1f\xb9u" +
	"\xb1/7|\xe8\xd6\x85\xbe\xdc\xf8\x81[\x17\xf9r\xd3" +
	"Y\xb7.\xf0\xe5\xe63\xeeg\x88{\xc4\xea\x97\x1a\x01" +
	"\x7f&\x13\xe8C\x9f\x16\xf8tAM\xbdO\xd4`\x80" +
	"\xcfe\xd6\xc5e\x96\x06\x81w\x0c\x0d\x1cy\x1b\xa3," +
	"\xddr\x92\xa8\xc1\\\x9a7\xb2\x0d\xc0\xc2H\x87\x98\xc0" +
	"8\xc4<t{\xbf\x9d\x1d\xde\xef\xbc9\x10\xf3\xb0\x0c" +
	"1\x1f\xc5\xfe\x03\xad\xf0\x91\x0b\xb1\x0af+OD\x13" +
	"8\x85\x98@\xa8?\xc5\x13\x98\xfcd,n\xf0\xb1\xb8" +
	"\xa6\x0c\x12\xb5/a\x08\x8e\x01`M\xd6\xeb\xf5f\x9c" +
	"\x19s\xab\xdf\xf3L\x15\x00Y\xdfA\xde\xadhW\x03" +
	" \x13}\x0eJ\xca\xa2\xe5\x00\x18Q\x16\xac\x06\xc0\xa8" +
	"r\xf3\x06\x00\x8c)s{\xa8\xc0Sn\xba\x17 m" +
	"\x0c\x17\x9cr\xbc\xdft\x92\xfde\xc7p\x9d\xbb\xf3=" +
	"e\xc7(\x02\x80;\x98/\xd9\xf4\x07\xb0\xe8\x1a\xe6\xa6" +
	"!\xa7\xa7L\xe2\x14G\xfc\xbe\xd45\xad\x9ci\x19\x8b" +
	"\xf3\x91\xe1B\xbeh:\xc6\xd8\x02GT\xd1\xabD0" +
	"H\xd2sM\xdf1\xb3N\xdf19\xd4\x18\x84K\xeb" +
	"0\xde/\xeb\x00\xd0\x10\xf3\xd0\xea\x17\x94\xc6\xd5\xd5\xf9" +
	"\xb9\xd2\xb8\xc1\xed\x0b\xd5f\x9e<\xab\xf2 e\x0dm" +
	"~E\x1a^\xc6\xc9\x00\x19\x07%\xccl\xc7@ ~" +
	"\x8f\x180n#\xf2\xd7\xc5\x80\xd1O\xf6\xf7\xe1(@" +
	"\xe6\xebD\x7f\x8c\xe8\x927\xd3\xe4\x8f A\xe4/\x89" +
	"\xfe\x84\x18T\xfa\x89~\xa7\xd8\xe71\xa2\xef!z\xdc" +
	"\x1fl>+\xe8\xcf\x10\xfd \x12l]\x1f\xe8\x07\x90" +
	"\x90\xf3\x121\x8e\x10\xa3Q\x14\xa6\xc1G5~\x08\xbb" +
	"\xf8!$44\x89\xea4\xf8\xd0\xc4\x8f\"\xa1\xa7\xf9" +
	"\x02Q\x83I?\xff\x09\x12F\x12\xe7\x89\x1a|\xbe\xe2" +
	"\xcfb\x070Y\xfe\x98\xa8\xc1\xac\x9f\xef\xc4\x0d\xc0\x94" +
	"HT\xc5\xa9\x00|\x87\x98\xbdn'I\xbe\x85\x0cg" +
	"G\xa7\xa2\x8a\xd3\x00\xf8\xdf e\xdb'\x88\xb1\x1b\xc7" +
	"\x86\xe4\xacY,\xe4\xf4\xf2*\x88\x87\xdb\xa3\x0a\x95\xe9" +
	"\xc3F\x9fm\x0c\x9a[V\x18\xd6&g\x08*\x09\xf2" +
	"\xb2y\x92U1V|LS\x10\x18\xd7\x0f\xe6\x9fn" +
	"h\x96\xa4D\x11\xcc\xdb?\x05,='\x01\x10h\x0e" +
	"c3-J\xff\xea\xb1R\xde\xaa\xd6\xff\xbfs:k" +
	"\x16E\xc24\x01\x07\xc6)\xa0\x05h\xbc\xb3\xe2\xfa\x80" +
	"Q\xe9\xa5\x93\xe1\xf2y\xa1\xdfK\xa7\xa9y\xba\xb3\xda" +
	"<\x8d\x0c\x8b~\xa3z~\xe5\x93t%\xe9\xf9c\x19" +
	"H\x16\x8ba}\x06\x1f`\xc7o\x86\xbd\xbb&\x1d3" +
	"ois<\x91\xa8\xb0\xbaC\xb8\xe9\xed\xe4\x03Y\xf2" +
	"k,\x88\xd1!\xd7\x85_o$zN\xe0f\xb3\x98" +
	"\x1er\x13\xef\x04\xc8\x0c\x11\xdd\x11\xb8\xb1\xc5\x9c\x8fo" +
	"\x16\xeb\x0bD\xdf&\x06\xffE1\x16\xe1e|`\x0c" +
	"\xfe\xa2\x8e\x8a\x93\x04\xfel\x80\xcc\xd7\x88\xfe\xb0\xc0Y" +
	"I\xf4\xcf\xfcA\xb1\x7f\x15\x7f\xf1\xbb\xc4g\x1a\xbeS" +
	"\xd0\x05\xfe\xbeC\xf4\x86\xbb\xc5\x87\x1a\xfemA\xff\x16" +
	"\xd1\xbfO\xf4\xc6-\xe2S\x0d\xff\x9e8\xf7\xfbD\x7f" +
	"\x8e\xe8Me\x15\xaf\x02\xe0{\xc5\xb9{\x88\xfe\x02\xd1" +
	"\x9b\xb7\xaa\xd8\x06\xc0\x9f\x17\xfb\xfc\x90\xe8/\x13=\xf1" +
	"g*N\x01\x82\xdd\xa3\x00\x99\x97\x89\xfe\x1a\x8e;\xec" +
	"p\x1d\xdd\xded8\xc5\xa5\x107sF\xe0\x18>u" +
	"1$\xc9\xddj\xc9\xbd\x10\xa7\xc2\xaf\x96\x8a\xfe\xb8\xc3" +
	"\xa1\xbd\xc7\xf22\x90\x16~]K_\x0aIj\xa3j" +
	"\xc9k!i\x99y\xab\x96|+$\xc7\x8c\x14}\xf2" +
	"2\xf4\xf1a\\~\xf0JH\x0b\xc7\xac\xa5\xf7A\x92" +
	"PSK^\x84>\xb0\xf2h\x8d\x83\x12\xd1\xd0\xd7\xa0" +
	"\x04\xbd\xe1R\xedtm\xd2g2F^T\x11\xc9\xaa" +
	"\x1d2\x86\xcbW\xe6\x97\xaf\x1d\xfe\x88\xb1\xcf\x9frP" +
	"v^\xd9Q\xadi\xc3c\xceq\x06]\xe3\xc9\xf5\xff" +
	"\x01\x00\x00\xff\xff\\\xb8\xbc\xd5"

func init() {
	schemas.Register(schema_a93fc509624c72d9,
		0x87e739250a60ea97,
		0x8e3b5f79fe593656,
		0x903455f06065422b,
		0x9500cce23b334d80,
		0x978a7cebdc549a4d,
		0x97b14cbe7cfec712,
		0x9aad50a41f4af45f,
		0x9dd1f724f4614a85,
		0x9e0e78711a7f87a9,
		0x9ea0b19b37fb4435,
		0xa9962a9ed0a4d7f8,
		0xabd73485a9636bc9,
		0xac3a6f60ef4cc6d3,
		0xae504193122357e5,
		0xb18aa5ac7a0d9420,
		0xb54ab3364333f598,
		0xb9521bccf10fa3b1,
		0xbaefc9120c56e274,
		0xbb90d5c287870be6,
		0xbfc546f6210ad7ce,
		0xc2573fe8a23e49f1,
		0xc42305476bb4746f,
		0xc863cd16969ee7fc,
		0xcafccddb68db1d11,
		0xce23dcd2d7b00c9b,
		0xcfea0eb02e810062,
		0xd07378ede1f9cc60,
		0xd1958f7dba521926,
		0xdebf55bbfa0fc242,
		0xe682ab4cf923a417,
		0xe82753cff0c2218f,
		0xec1619d4400a0290,
		0xed8bca69f7fb0cbf,
		0xf1c8950dab257542)
}