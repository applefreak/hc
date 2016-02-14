package model

type DoorStateType byte

const (
	DoorStateOpen    DoorStateType = 0x00
	DoorStateClosed  DoorStateType = 0x01
	DoorStateOpening DoorStateType = 0x02
	DoorStateClosing DoorStateType = 0x03
    DoorStateStopped DoorStateType = 0x04
)

type LockStateType byte

const (
	LockStateUnsecured    LockStateType = 0x00
	LockStateSecured      LockStateType = 0x01
)

// A Garage Door Opener is a device that allows you to control a motor used
// to open a garage door. Unfortunately, it doesn't have a Stop characteristics. 
type GarageDoorOpener interface {
	// SetState sets the current state
	SetState(DoorStateType)

	// State returns the current state
	State() DoorStateType

	// SetTargetState sets the target state
	SetTargetState(DoorStateType)

	// TargetState returns the target state
	TargetState() DoorStateType

    // SetObstructionDetected sets the obstruction detection
    SetObstructionDetected(bool)

    // ObstructionDetected returns the obstruction detection
    ObstructionDetected() bool

    SetLockTargetState(DoorStateType)

    OnTargetStateChange(func(DoorStateType))

    OnLockTargetStateChange(func(LockStateType))
}