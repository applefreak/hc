package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type garageDoorOpener struct {
	*Accessory

	garageDoorOpener *service.GarageDoorOpener
}

// NewGarageDoorOpener returns a garageDoorOpener which implements model.GarageDoorOpener.
func NewGarageDoorOpener(info model.Info) *garageDoorOpener {
	accessory := New(info)
	o := service.NewGarageDoorOpener(info.Name, model.DoorStateClosed)

	accessory.AddService(o.Service)

	return &garageDoorOpener{accessory, o}
}

func (o *garageDoorOpener) State() model.DoorStateType {
	return o.garageDoorOpener.State.DoorState()
}

func (o *garageDoorOpener) SetTargetState(value model.DoorStateType) {
	o.garageDoorOpener.TargetState.SetDoorState(value)
}

func (o *garageDoorOpener) SetState(value model.DoorStateType) {
	o.garageDoorOpener.State.SetDoorState(value)
}

func (o *garageDoorOpener) SetLockTargetState(value model.LockStateType) {
	o.garageDoorOpener.LockTargetState.SetLockState(value)
}

func (o *garageDoorOpener) OnTargetStateChange(fn func(model.DoorStateType)) {
	o.garageDoorOpener.TargetState.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(model.DoorStateType(new.(byte)))
	})
}

func (o *garageDoorOpener) OnLockTargetStateChange(fn func(model.LockStateType)) {
	o.garageDoorOpener.LockTargetState.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(model.LockStateType(new.(byte)))
	})
}
