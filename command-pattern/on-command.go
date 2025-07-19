package commandpattern

type OnCommand struct {
	device Device
}

func NewOnCommand(device Device) Command {
	return &OnCommand{
		device: device,
	}
}

func (oc *OnCommand) execute() {
	oc.device.On()
}
