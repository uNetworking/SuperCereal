package supercereal

type JSONArray JSONStream
type JSONObject JSONStream

func (p *JSONStream) routeValueType(value interface{}) {
	switch v := value.(type) {
	case string:
		p.PutString([]byte(v))
	case int:
		p.PutInt(v)
	case bool:
		p.PutBoolean(v)
	case func(array *JSONArray):
		p.OpenArray()
		v((*JSONArray)(p))
		p.CloseArray()
	case func(object *JSONObject):
		p.OpenObject()
		v((*JSONObject)(p))
		p.CloseObject()
	default:
		p.PutNull()
	}
}

func (p *JSONStream) Serialize(value interface{}) {
	p.Reset()
	p.routeValueType(value)
	p.End()
}

func (p *JSONArray) Put(value interface{}) {
	(*JSONStream)(p).routeValueType(value)
}

func (p *JSONObject) Put(key string, value interface{}) {
	(*JSONStream)(p).PutKey([]byte(key))
	(*JSONStream)(p).routeValueType(value)
}