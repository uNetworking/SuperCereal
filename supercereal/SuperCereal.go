package supercereal

import (
	"fmt"
)

type JSONStream struct {
	cb func(data []byte)

	front  int
	buffer []byte

	// ha en stack av vad som är öppet?
	currentScopeIsEmpty bool
}

func NewJSONStream(cb func(data []byte)) *JSONStream {
	js := &JSONStream{cb: cb, buffer: make([]byte, 1024*1024), front: 0} // 1mb buffer by default
	js.OpenObject()
	return js
}

func (p *JSONStream) OpenObject() {
	p.currentScopeIsEmpty = true
	p.buffer[p.front] = '{'
	p.front++
}

func (p *JSONStream) PutKey(key []byte) {

	// put comma if this scope is not empty!
	if !p.currentScopeIsEmpty {
		p.buffer[p.front] = ','
		p.front++
	}
	p.currentScopeIsEmpty = false

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
	byteRep := []byte(fmt.Sprintf("%d", value))
	copy(p.buffer[p.front:], byteRep)
	p.front += len(byteRep)
}

func (p *JSONStream) PutString(value []byte) {
	p.buffer[p.front] = '"'
	p.front++
	copy(p.buffer[p.front:], value)
	p.front += len(value)
	p.buffer[p.front] = '"'
	p.front++
}

func (p *JSONStream) CloseObject() {
	p.currentScopeIsEmpty = false
	p.buffer[p.front] = '}'
	p.front++
}

func (p *JSONStream) Reset() {
	p.front = 0
	p.OpenObject()
}

func (p *JSONStream) End() {
	p.CloseObject()
	p.cb(p.buffer[:p.front])
}
