package optional

type Map struct {
	isSet bool
	value map[any]any
}

func NewMap(value map[any]any) Map {
	return Map{
		true,
		value,
	}
}

// EmptyMap returns a new Map that does not have a value set.
func EmptyMap() Map {
	return Map{
		false,
		nil,
	}
}

func (i Map) IsSet() bool {
	return i.isSet
}

func (i Map) Value() map[any]any {
	return i.value
}

func (i Map) Default(defaultValue map[any]any) map[any]any {
	if i.isSet {
		return i.value
	}
	return defaultValue
}
