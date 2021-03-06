package pob

import (
	"errors"
	"github.com/iost-official/Go-IOS-Protocol/core/blockcache"
	"github.com/iost-official/Go-IOS-Protocol/core/tx"
	"sync"
)

var ErrIllegalTx = errors.New("illegal tx")

type Recorder interface {
	Record(tx2 tx.Tx) error
	Pop() tx.Tx
	Listen()
	Close()
	Exist(tx2 tx.Tx) bool
}

func NewRecorder() Recorder {
	return &RecorderImpl{
		TxHeap: NewTxHeap(),
	}
}

type RecorderImpl struct {
	TxHeap
	storage     map[string]bool
	isClose     bool
	isListening sync.Mutex
}

func (p *RecorderImpl) Record(tx2 tx.Tx) error {
	p.isListening.Lock()
	defer p.isListening.Unlock()
	if blockcache.VerifyTxSig(tx2) {
		tx.RecordTx(tx2, tx.Data.Self())
		p.storage[string(tx2.Hash())] = true
		p.Push(tx2)
		return nil
	} else {
		return ErrIllegalTx
	}
}

func (p *RecorderImpl) Pop() tx.Tx {
	tx2 := p.TxHeap.Pop()
	delete(p.storage, string(tx2.Hash()))
	return tx2
}

func (p *RecorderImpl) Exist(tx2 tx.Tx) bool {
	_, ok := p.storage[string(tx2.Hash())]
	return ok
}

func (p *RecorderImpl) Listen() {
	if p.isClose == true {
		p.isClose = false
	} else {
		return
	}
	p.isListening.Unlock()
}

func (p *RecorderImpl) Close() {
	if p.isClose == false {
		p.isClose = true
	} else {
		return
	}
	p.isListening.Lock()
}
