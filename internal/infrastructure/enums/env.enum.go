package enums

// EnvEnum: Enumeraciones para los environments
type EnvEnum string

const (
	EnvLocal EnvEnum = "local"
	EnvDev   EnvEnum = "dev"
	EnvQa    EnvEnum = "qa"
	EnvProd  EnvEnum = "pdn"
)
