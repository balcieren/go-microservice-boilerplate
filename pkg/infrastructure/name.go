package infrastructure

type AppName string

func (an AppName) String() string {
	return string(an)
}
