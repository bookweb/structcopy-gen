package structcopy

type InputOption struct {
	SkipFieldsMap       map[string]bool
	MatchFieldsMap      map[string]string
	MatchMethodsMap     map[string]string
	ConvertersMap       map[string]string
	StructConverterFunc string
}
