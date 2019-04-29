package supercereal

import (
	"fmt"
)

type JSONStream struct {
	front  int
	buffer []byte

	comma bool
}

func (p *JSONStream) write(b byte) {
	p.buffer[p.front] = b
	p.front++
}

func (p *JSONStream) reset() {
	p.front = 0
	p.comma = false
}

func NewJSONStream() *JSONStream {

	// cap should be 1 mb by default, reset should not reset it only the length
	js := &JSONStream{buffer: make([]byte, 1024*1024*10)} // 1mb buffer by default
	js.reset()
	return js
}

func (p *JSONStream) end() []byte {
	return p.buffer[:p.front]
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

	// todo: write
	byteRep := []byte(fmt.Sprintf("%d", value))
	copy(p.buffer[p.front:], byteRep)
	p.front += len(byteRep)
}

func (p *JSONStream) PutFloat64(value float64) {
	p.putComma()
	p.comma = true

	// todo: write
	byteRep := []byte(fmt.Sprintf("%f", value))
	copy(p.buffer[p.front:], byteRep)
	p.front += len(byteRep)
}

func (p *JSONStream) PutNull() {
	p.putComma()
	p.comma = true

	// todo: write
	byteRep := []byte("null")
	copy(p.buffer[p.front:], byteRep)
	p.front += len(byteRep)
}

func (p *JSONStream) PutBoolean(value bool) {
	p.putComma()
	p.comma = true

	// todo: write
	byteRep := "false"
	if value {
		byteRep = "true"
	}
	copy(p.buffer[p.front:], byteRep)
	p.front += len(byteRep)
}

func (p *JSONStream) escapedCopy(value []byte) {
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
}

func (p *JSONStream) PutString(value []byte) {
	p.putComma()
	p.comma = true
	p.write('"')
	p.escapedCopy(value)
	p.write('"')
}