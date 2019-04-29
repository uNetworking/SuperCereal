package supercereal

type Array JSONStream
type Object JSONStream

func (p *JSONStream) routeValueType(value interface{}) {
	switch v := value.(type) {
	case string:
		p.PutString([]byte(v))
	case int:
		p.PutInt(v)
	case float64:
		p.PutFloat64(v)
	case bool:
		p.PutBoolean(v)
	case func(array *Array):
		p.OpenArray()
		v((*Array)(p))
		p.CloseArray()
	case func(object *Object):
		p.OpenObject()
		v((*Object)(p))
		p.CloseObject()
	default:
		p.PutNull()
	}
}

func (p *Array) Put(value interface{}) {
	(*JSONStream)(p).routeValueType(value)
}

func (p *Object) Put(key string, value interface{}) {
	(*JSONStream)(p).PutKey([]byte(key))
	(*JSONStream)(p).routeValueType(value)
}