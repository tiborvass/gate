// Code generated by internal/event-type-gen. DO NOT EDIT.

package event

func (x *FailNetwork) EventName() string    { return Event_Type_name[x.EventType()] }
func (x *FailProtocol) EventName() string   { return Event_Type_name[x.EventType()] }
func (x *FailRequest) EventName() string    { return Event_Type_name[x.EventType()] }
func (x *InstanceAttach) EventName() string { return Event_Type_name[x.EventType()] }
func (x *InstanceCreate) EventName() string { return Event_Type_name[x.EventType()] }
func (x *InstanceDelete) EventName() string { return Event_Type_name[x.EventType()] }
func (x *InstanceDetach) EventName() string { return Event_Type_name[x.EventType()] }
func (x *ProgramCheck) EventName() string   { return Event_Type_name[x.EventType()] }
func (x *ProgramCreate) EventName() string  { return Event_Type_name[x.EventType()] }
func (x *ProgramLoad) EventName() string    { return Event_Type_name[x.EventType()] }
func (x *ServerAccess) EventName() string   { return Event_Type_name[x.EventType()] }

func (*FailNetwork) EventType() int32    { return int32(Event_FailNetwork) }
func (*FailProtocol) EventType() int32   { return int32(Event_FailProtocol) }
func (*FailRequest) EventType() int32    { return int32(Event_FailRequest) }
func (*InstanceAttach) EventType() int32 { return int32(Event_InstanceAttach) }
func (*InstanceCreate) EventType() int32 { return int32(Event_InstanceCreate) }
func (*InstanceDelete) EventType() int32 { return int32(Event_InstanceDelete) }
func (*InstanceDetach) EventType() int32 { return int32(Event_InstanceDetach) }
func (*ProgramCheck) EventType() int32   { return int32(Event_ProgramCheck) }
func (*ProgramCreate) EventType() int32  { return int32(Event_ProgramCreate) }
func (*ProgramLoad) EventType() int32    { return int32(Event_ProgramLoad) }
func (*ServerAccess) EventType() int32   { return int32(Event_ServerAccess) }
