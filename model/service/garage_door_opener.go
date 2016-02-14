package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

// GarageDoorOpener is service to represent a thermostat.
type GarageDoorOpener struct {
	*Service

	Name *characteristic.Name
	State       *characteristic.DoorState
	TargetState *characteristic.DoorState
	ObstructionDetected *characteristic.ObstructionDetected

	LockTargetState *characteristic.LockState

	targetStateChange func(model.DoorStateType)
	targetLockStateChange func(model.LockStateType)
}

// NewGarageDoorOpener returns a thermostat service.
func NewGarageDoorOpener(name string, currentState model.DoorStateType) *GarageDoorOpener {
	nameChar := characteristic.NewName(name)
	state := characteristic.NewCurrentDoorState(currentState)
	targetState := characteristic.NewTargetDoorState(model.DoorStateClosed)
	obstructionDetected := characteristic.NewObstructionDetected(false)
	lockTargetState := characteristic.NewLockTargetState(model.LockStateSecured)

	svc := New()
	svc.Type = typeGarageDoorOpener
	svc.AddCharacteristic(nameChar.Characteristic)
	svc.AddCharacteristic(state.Characteristic)
	svc.AddCharacteristic(targetState.Characteristic)
	svc.AddCharacteristic(obstructionDetected.Characteristic)
	svc.AddCharacteristic(lockTargetState.Characteristic)

	return &GarageDoorOpener{svc, nameChar, state, targetState, obstructionDetected, lockTargetState, nil, nil}
}
