package commandpattern

type OffCommand struct {
	device Device
}

func NewOffCommand(device Device) Command {
	return &OffCommand{
		device: device,
	}
}

func (oc *OffCommand) execute() {
	oc.device.Off()
}
