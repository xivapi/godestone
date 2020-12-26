// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package pack

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MinionTable struct {
	_tab flatbuffers.Table
}

func GetRootAsMinionTable(buf []byte, offset flatbuffers.UOffsetT) *MinionTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MinionTable{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MinionTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MinionTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MinionTable) Minions(obj *Minion, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MinionTable) MinionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MinionTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MinionTableAddMinions(builder *flatbuffers.Builder, Minions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Minions), 0)
}
func MinionTableStartMinionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MinionTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Minion struct {
	_tab flatbuffers.Table
}

func GetRootAsMinion(buf []byte, offset flatbuffers.UOffsetT) *Minion {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Minion{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Minion) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Minion) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Minion) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Minion) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Minion) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Minion) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Minion) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Minion) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MinionStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func MinionAddId(builder *flatbuffers.Builder, Id uint32) {
	builder.PrependUint32Slot(0, Id, 0)
}
func MinionAddNameEn(builder *flatbuffers.Builder, NameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(NameEn), 0)
}
func MinionAddNameFr(builder *flatbuffers.Builder, NameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(NameFr), 0)
}
func MinionAddNameDe(builder *flatbuffers.Builder, NameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(NameDe), 0)
}
func MinionAddNameJa(builder *flatbuffers.Builder, NameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(NameJa), 0)
}
func MinionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
