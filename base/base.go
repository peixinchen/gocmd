package base

type Command interface {
    GetName() string

    PreExecute()

    Execute()

    PostExecute()
}

type BaseCommand struct {
    Name    string
    Short   string
    Long    string
}

func (c BaseCommand) GetName() string {
    return c.Name
}

func (c BaseCommand) PreExecute() {
}

func (c BaseCommand) PostExecute() {
}
