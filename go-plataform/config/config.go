package config

type Configuration interface {
	GetString(name string)(configValue string, found bool)
	GetInt(name string)(configValue int, found bool)
	GetBool(name string) (configValue bool, found bool)
	GetFloat(name string)(configValue float64, found bool)

	GetStringDefault(name string, defVal string )(configValue string)
	GetIntDefault(name string, defVal int)(configValue string)
	GetBoolDefault(name string, defVal bool)(configValue string)
	GetFloatDefault(name string, defVal float64)(configValue string)

	GetSection(sectionName string)(section Configuration, found bool)
}