package characteristic

import (
	"github.com/brutella/hc/model"
)

type LockState struct {
	*ByteCharacteristic
}

func NewLockState(current model.LockStateType, CharacteristicType CharacteristicType, permissions []string) *LockState {
	s := LockState{NewByteCharacteristic(byte(current), permissions)}
	s.Type = CharacteristicType

	return &s
}

func NewLockCurrentState(current model.LockStateType) *LockState {
	return NewLockState(current, TypeLockMechanismCurrentState, PermsRead())
}

func NewLockTargetState(current model.LockStateType) *LockState {
	return NewLockState(current, TypeLockMechanismTargetState, PermsAll())
}

func (s *LockState) SetLockState(state model.LockStateType) {
	s.SetByte(byte(state))
}

func (s *LockState) LockState() model.LockStateType {
	return model.LockStateType(s.Byte())
}