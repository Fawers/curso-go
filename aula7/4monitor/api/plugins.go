package api

type Plugin interface {
	Init()
	Collect() (err error)
	Describe() string
}
