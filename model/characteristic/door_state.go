package characteristic

import (
	"github.com/brutella/hc/model"
)

type DoorState struct {
	*ByteCharacteristic
}

func NewDoorState(current model.DoorStateType, CharacteristicType CharacteristicType, permissions []string) *DoorState {
	s := DoorState{NewByteCharacteristic(byte(current), permissions)}
	s.Type = CharacteristicType

	return &s
}

func NewCurrentDoorState(current model.DoorStateType) *DoorState {
	return NewDoorState(current, TypeCurrentDoorState, PermsRead())
}

func NewTargetDoorState(current model.DoorStateType) *DoorState {
	return NewDoorState(current, TypeTargetDoorState, PermsAll())
}

func (s *DoorState) SetDoorState(state model.DoorStateType) {
	s.SetByte(byte(state))
}

func (s *DoorState) DoorState() model.DoorStateType {
	return model.DoorStateType(s.Byte())
}