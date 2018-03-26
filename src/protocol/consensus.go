package protocol

import (
	"fmt"
	"time"
	"IOS/src/iosbase"
)

type View interface {
	GetPrimaryID() string
	GetBackupID() []string
	isPrimary(ID string) bool
	isBackup(ID string) bool
}

type Character int

const (
	Primary Character = iota
	Backup
	Idle
)

type Phase int

const (
	StartPhase      Phase = iota
	PrePreparePhase
	PreparePhase
	CommitPhase
	PanicPhase
	EndPhase
)

const (
	Port       = 12306
	ExpireTime = 1 * time.Minute
)

type Consensus struct {
	iosbase.Member
	Replica
	Validator
	NetworkFilter

	phase     Phase
	isRunning bool
}

func (c *Consensus) Init(bc iosbase.BlockChain, sp iosbase.StatePool, network iosbase.Network) error {
	err := c.replicaInit(bc, sp)
	if err != nil {
		return err
	}
	err = c.networkFilterInit(network)
	return err
}

func (c *Consensus) Run() {
	c.phase = StartPhase
	var req iosbase.Request
	var err error = nil
	c.isRunning = true

	for c.isRunning {
		switch c.phase {
		case StartPhase:
			c.phase, err = c.onNewView()
		case PrePreparePhase:
			pp := PrePrepare{}
			pp.Unmarshal(req.Body)
			c.phase, err = c.onPrePrepare(pp)
		case PreparePhase:
			p := Prepare{}
			p.Unmarshal(req.Body)
			c.phase, err = c.onPrepare(p)
		case CommitPhase:
			cm := Commit{}
			cm.Unmarshal(req.Body)
			c.phase, err = c.onCommit(cm)
		case PanicPhase:
		case EndPhase:
			return
		}

		if err != nil {
			fmt.Println(err)
		}

		to := time.NewTimer(1 * time.Minute)
		select {
		case <-c.receiveChan:
			req = <-c.receiveChan
			if ! to.Stop() {
				<-to.C
			}
			to.Reset(ExpireTime)
		case <-to.C:
			c.phase, err = c.onTimeOut()
			if err != nil {
				return
			}
			to.Reset(ExpireTime)
		}
	}
}

func (c *Consensus) Stop() {
	c.isRunning = false
}
func (c *Consensus) PublishTx(tx iosbase.Tx) error {

	return nil
}
func (c *Consensus) CheckTx(tx iosbase.Tx) (iosbase.TxStatus, error) {
	return iosbase.POOL, nil
}
func (c *Consensus) GetStatus() (iosbase.BlockChain, iosbase.StatePool, error) {
	return c.blockChain, c.statePool, nil
}
func (c *Consensus) GetCachedStatus() (iosbase.BlockChain, iosbase.StatePool, error) {
	return c.blockChain, c.statePool, nil
}
