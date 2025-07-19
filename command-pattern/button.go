package commandpattern

type Button struct {
	command Command
}

func (b *Button) SetCommand(command Command) {
	b.command = command
}

func (b *Button) Press() {
	b.command.execute()
}
