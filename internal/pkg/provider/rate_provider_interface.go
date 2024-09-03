package provider

type IRateProvider interface {
	GetRate() (*float64, error)
}
