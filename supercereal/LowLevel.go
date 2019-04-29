package supercereal

import (
	"fmt"
)

type JSONStream struct {
	buffer []byte
	comma bool
}

func (p *JSONStream) writeArray(b []byte) {
	p.buffer = append(p.buffer, b...)
}

func (p *JSONStream) write(b byte) {
	p.buffer = append(p.buffer, b)
}

func (p *JSONStream) reset() {
	p.buffer = p.buffer[:0]
	p.comma = false
}

func NewJSONStream() *JSONStream {
	js := &JSONStream{buffer: make([]byte, 0, 1024 * 1024)}
	js.reset()
	return js
}

func (p *JSONStream) end() []byte {
	return p.buffer
}

func (p *JSONStream) consumeComma() {
	if p.comma {
		p.write(',')
		p.comma = false
	}
}

func (p *JSONStream) putComma() {
	if p.comma {
		p.write(',')
	}
}

func (p *JSONStream) OpenArray() {
	p.consumeComma()
	p.write('[')
}

func (p *JSONStream) OpenObject() {
	p.consumeComma()
	p.write('{')
}

func (p *JSONStream) CloseArray() {
	p.write(']')
	p.comma = true
}

func (p *JSONStream) CloseObject() {
	p.write('}')
	p.comma = true
}

func (p *JSONStream) PutKey(key []byte) {
	p.consumeComma()
	p.write('"')
	p.escapedCopy(key);
	p.write('"')
	p.write(':')
}

func (p *JSONStream) PutInt(value int) {
	p.putComma()
	p.comma = true
	p.writeArray([]byte(fmt.Sprintf("%d", value)))
}

func (p *JSONStream) PutFloat64(value float64) {
	p.putComma()
	p.comma = true
	p.writeArray([]byte(fmt.Sprintf("%f", value)))
}

func (p *JSONStream) PutNull() {
	p.putComma()
	p.comma = true
	p.writeArray([]byte("null"))
}

func (p *JSONStream) PutBoolean(value bool) {
	p.putComma()
	p.comma = true
	if value {
		p.writeArray([]byte("true"))
	} else {
		p.writeArray([]byte("false"))
	}
}

func (p *JSONStream) escapedCopy(value []byte) {
	for i := 0; i < len(value); i++ {
		if value[i] != '\\' && (value[i] > '"' || value[i] == ' ') {
			p.write(value[i])
		} else if value[i] == '"' {
			p.write('\\')
			p.write('"')
		} else if value[i] == '\\' {
			p.write('\\')
			p.write('"')
		} else if value[i] == '\n' {
			p.write('\\')
			p.write('n')
		} else if value[i] == '\r' {
			p.write('\\')
			p.write('r')
		} else if value[i] == '\t' {
			p.write('\\')
			p.write('t')
		} else if value[i] == '\f' {
			p.write('\\')
			p.write('f')
		} else if value[i] == '\b' {
			p.write('\\')
			p.write('b')
		} else {
			p.write(value[i])
		}
	}
}

func (p *JSONStream) PutString(value []byte) {
	p.putComma()
	p.comma = true
	p.write('"')
	p.escapedCopy(value)
	p.write('"')
}