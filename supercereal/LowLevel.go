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
	js := &JSONStream{buffer: make([]byte, 1024*1024*10)} // 1mb buffer by default
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
	//copy(p.buffer[p.front:], value)
	//p.front += len(value)

	for i := 0; i < len(value); i++ {
		if value[i] != '\\' && (value[i] > '"' || value[i] == ' ') {
			p.buffer[p.front] = value[i]
			p.front++
		} else if value[i] == '"' {
			p.buffer[p.front] = '\\'
			p.buffer[p.front+1] = '"'
			p.front += 2
		} else if value[i] == '\\' {
			p.buffer[p.front] = '\\'
			p.buffer[p.front+1] = '\\'
			p.front += 2
		} else if value[i] == '\n' {
			p.buffer[p.front] = '\\'
			p.buffer[p.front+1] = 'n'
			p.front += 2
		} else if value[i] == '\r' {
			p.buffer[p.front] = '\\'
			p.buffer[p.front+1] = 'r'
			p.front += 2
		} else if value[i] == '\t' {
			p.buffer[p.front] = '\\'
			p.buffer[p.front+1] = 't'
			p.front += 2
		} else if value[i] == '\f' {
			p.buffer[p.front] = '\\'
			p.buffer[p.front+1] = 'f'
			p.front += 2
		} else if value[i] == '\b' {
			p.buffer[p.front] = '\\'
			p.buffer[p.front+1] = 'b'
			p.front += 2
		} else {
			p.buffer[p.front] = value[i]
			p.front++
		}
	}

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
