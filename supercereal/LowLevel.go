package supercereal

import (
	"fmt"
)

type JSONStream struct {
	cb func(json []byte)

	front  int
	buffer []byte

	comma bool
}

func NewJSONStream() *JSONStream {
	js := &JSONStream{buffer: make([]byte, 1024*1024)} // 1mb buffer by default
	js.Reset()
	return js
}

func (p *JSONStream) OnJSON(cb func(json []byte)) {
	p.cb = cb
}

func (p *JSONStream) consumeComma() {
	if p.comma {
		p.buffer[p.front] = ','
		p.front++
		p.comma = false
	}
}

func (p *JSONStream) putComma() {
	if p.comma {
		p.buffer[p.front] = ','
		p.front++
	}
}

func (p *JSONStream) OpenArray() {
	p.consumeComma()
	p.buffer[p.front] = '['
	p.front++
}

func (p *JSONStream) OpenObject() {
	p.consumeComma()
	p.buffer[p.front] = '{'
	p.front++
}

func (p *JSONStream) CloseArray() {
	p.buffer[p.front] = ']'
	p.front++
	p.comma = true
}

func (p *JSONStream) CloseObject() {
	p.buffer[p.front] = '}'
	p.front++
	p.comma = true
}

func (p *JSONStream) PutKey(key []byte) {
	p.consumeComma()
	p.buffer[p.front] = '"'
	p.front++
	copy(p.buffer[p.front:], key)
	p.front += len(key)
	p.buffer[p.front] = '"'
	p.front++
	p.buffer[p.front] = ':'
	p.front++
}

func (p *JSONStream) PutInt(value int) {
	p.putComma()
	p.comma = true

	byteRep := []byte(fmt.Sprintf("%d", value))
	copy(p.buffer[p.front:], byteRep)
	p.front += len(byteRep)
}

func (p *JSONStream) PutNull() {
	p.putComma()
	p.comma = true

	byteRep := []byte("null")
	copy(p.buffer[p.front:], byteRep)
	p.front += len(byteRep)
}

func (p *JSONStream) PutBoolean(value bool) {
	p.putComma()
	p.comma = true

	byteRep := "false"
	if value {
		byteRep = "true"
	}
	copy(p.buffer[p.front:], byteRep)
	p.front += len(byteRep)
}

func (p *JSONStream) PutString(value []byte) {
	p.putComma()
	p.comma = true

	p.buffer[p.front] = '"'
	p.front++
	copy(p.buffer[p.front:], value)
	p.front += len(value)
	p.buffer[p.front] = '"'
	p.front++
}

func (p *JSONStream) Reset() {
	p.front = 0
	p.comma = false
}

func (p *JSONStream) End() {
	p.cb(p.buffer[:p.front])
}
