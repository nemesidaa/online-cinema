package aggregator

import "sync/atomic"

type Aggregator struct {
	acnt atomic.Int32
	recv <-chan []byte
}

func New(recvChan <-chan []byte) *Aggregator {
	return &Aggregator{
		acnt: atomic.Int32{},
		recv: recvChan,
	}
}

// Задача: принять данные из fileStorage в виде <-chan, и образоавть массив(целый). Использовать chan, данные не менять. Проверять на закрытие канала
// Сигнатура в boilerplate/aggregator.go
