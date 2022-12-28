package environment

type Environment int

const (
	Local Environment = iota
	Development
	QA
	Staging
	Sandbox
	Production
)

func (e Environment) String() string {
	return [...]string{
		"Local",
		"Development",
		"QA",
		"Staging",
		"Sandbox",
		"Production",
	}[e]
}
