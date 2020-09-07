// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package attest

import (
	"fmt"
	"testing"
	"time"

	"github.com/lf-edge/eve/pkg/pillar/base"
	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/sirupsen/logrus"
)

type VerifierMock struct{}

var (
	simulateControllerReqFailure = false
	simulateNonceMismatch        = false
	simulateQuoteMismatch        = false
	simulateNoCertToValidate     = false
	simulateITokenMismatch       = false
	simulateTpmAgentDown         = false
	simulateNoVerifier           = false
	simulateNoEscrowData         = false
)

func (server *VerifierMock) SendNonceRequest(ctx *Context) error {
	switch {
	case simulateControllerReqFailure:
		fmt.Printf("Simulating Controller being down\n")
		return ErrControllerReqFailed
	case simulateNoVerifier:
		fmt.Printf("Simulating no verifier support in Controller\n")
		return ErrNoVerifier
	}
	return nil
}

func (server *VerifierMock) SendAttestQuote(ctx *Context) error {
	switch {
	case simulateControllerReqFailure:
		fmt.Printf("Simulating Controller being down\n")
		return ErrControllerReqFailed

	case simulateNonceMismatch:
		fmt.Printf("Simulating Nonce mismatch\n")
		return ErrNonceMismatch

	case simulateQuoteMismatch:
		fmt.Printf("Simulating Quote mismatch\n")
		return ErrQuoteMismatch

	case simulateNoCertToValidate:
		fmt.Printf("Simulating No quote cert in Controller\n")
		return ErrNoCertYet
	}
	return nil
}

func (server *VerifierMock) SendAttestEscrow(ctx *Context) error {
	switch {
	case simulateControllerReqFailure:
		fmt.Printf("Simulating Controller being down\n")
		return ErrControllerReqFailed

	case simulateITokenMismatch:
		fmt.Printf("Simulating Integrity Token mismatch\n")
		return ErrITokenMismatch
	}

	return nil
}

type TpmAgentMock struct{}

func (agent *TpmAgentMock) SendInternalQuoteRequest(ctx *Context) error {
	if simulateTpmAgentDown {
		return ErrTpmAgentUnavailable
	}
	return nil
}

func initTest() *Context {
	tpmAgent = &TpmAgentMock{}
	verifier = &VerifierMock{}
	simulateControllerReqFailure = false
	simulateNonceMismatch = false
	simulateQuoteMismatch = false
	simulateNoCertToValidate = false
	simulateITokenMismatch = false
	simulateTpmAgentDown = false
	simulateNoVerifier = false
	simulateNoEscrowData = false

	logger := logrus.StandardLogger()
	log := base.NewSourceLogObject(logger, "test", 1234)
	ps := pubsub.New(&pubsub.EmptyDriver{}, logger, log)
	ctx := &Context{
		PubSub:       ps,
		log:          log,
		event:        EventInitialize,
		state:        StateNone,
		restartTimer: time.NewTimer(1 * time.Second),
		eventTrigger: make(chan Event),
		retryTime:    1,
	}
	ctx.restartTimer.Stop()
	return ctx
}

func TestGoodPath(t *testing.T) {
	fmt.Println("------TestGoodPath-------")
	ctx := initTest()
	stopTrigger := make(chan int)

	go func() {
		ctx.eventTrigger <- EventInitialize
		for ctx.state != StateInternalQuoteWait {
			time.Sleep(1 * time.Second)
		}
		ctx.eventTrigger <- EventInternalQuoteRecvd
		time.Sleep(10 * time.Second)
		stopTrigger <- 1
	}()
	for {
		select {
		case trigger := <-ctx.eventTrigger:
			if err := despatchEvent(trigger, ctx.state, ctx); err != nil {
				t.Errorf("%v", err)
			}
		case <-stopTrigger:
			if ctx.state != StateComplete {
				t.Errorf("Expected %s, Got %s", StateComplete.String(), ctx.state.String())
			}
			return
		}
	}
}

func TestNonceMismatch(t *testing.T) {
	fmt.Println("-------TestNonceMismatch---")
	ctx := initTest()
	simulateNonceMismatch = true

	go func() {
		ctx.eventTrigger <- EventInitialize
		for ctx.state != StateInternalQuoteWait {
			time.Sleep(1 * time.Second)
		}
		ctx.eventTrigger <- EventInternalQuoteRecvd
	}()
	for {
		select {
		case trigger := <-ctx.eventTrigger:
			if err := despatchEvent(trigger, ctx.state, ctx); err != nil {
				t.Errorf("%v", err)
			}
		case <-ctx.restartTimer.C:
			if ctx.state != StateRestartWait {
				t.Errorf("Expected %s, Got %s", StateRestartWait.String(), ctx.state.String())
			}
			return
		}
	}
}

func TestQuoteMismatch(t *testing.T) {
	fmt.Println("-------TestQuoteMismatch------")
	ctx := initTest()
	simulateQuoteMismatch = true

	go func() {
		ctx.eventTrigger <- EventInitialize
		for ctx.state != StateInternalQuoteWait {
			time.Sleep(1 * time.Second)
		}
		ctx.eventTrigger <- EventInternalQuoteRecvd
	}()
	for {
		select {
		case trigger := <-ctx.eventTrigger:
			if err := despatchEvent(trigger, ctx.state, ctx); err != nil {
				t.Errorf("%v", err)
			}
		case <-ctx.restartTimer.C:
			if ctx.state != StateRestartWait {
				t.Errorf("Expected %s, Got %s", StateRestartWait.String(), ctx.state.String())
			}
			return
		}
	}
}

func TestNoCertInController(t *testing.T) {
	fmt.Println("-------TestNoCertInController------")
	ctx := initTest()
	simulateNoCertToValidate = true

	go func() {
		ctx.eventTrigger <- EventInitialize
		for ctx.state != StateInternalQuoteWait {
			time.Sleep(1 * time.Second)
		}
		ctx.eventTrigger <- EventInternalQuoteRecvd
	}()
	for {
		select {
		case trigger := <-ctx.eventTrigger:
			if err := despatchEvent(trigger, ctx.state, ctx); err != nil {
				t.Errorf("%v", err)
			}
		case <-ctx.restartTimer.C:
			if ctx.state != StateRestartWait {
				t.Errorf("Expected %s, Got %s", StateRestartWait.String(), ctx.state.String())
			}
			return
		}
	}
}

func TestControllerNotAvbleInNonceWait(t *testing.T) {
	fmt.Println("-----TestControllerNotAvbleInNonceWait--")
	ctx := initTest()
	simulateControllerReqFailure = true

	go func() {
		ctx.eventTrigger <- EventInitialize
		for ctx.state != StateInternalQuoteWait {
			time.Sleep(1 * time.Second)
		}
		ctx.eventTrigger <- EventInternalQuoteRecvd
	}()
	for {
		select {
		case trigger := <-ctx.eventTrigger:
			if err := despatchEvent(trigger, ctx.state, ctx); err != nil {
				t.Errorf("%v", err)
			}
		case <-ctx.restartTimer.C:
			if ctx.state != StateNonceWait {
				t.Errorf("Expected %s, Got %s", StateNonceWait.String(), ctx.state.String())
			}
			return
		}
	}
}

func TestControllerNotAvbleInAttestWait(t *testing.T) {
	fmt.Println("--------TestControllerNotAvbleInAttestWait----")
	ctx := initTest()

	go func() {
		ctx.state = StateInternalQuoteWait
		simulateControllerReqFailure = true
		ctx.eventTrigger <- EventInternalQuoteRecvd
	}()
	for {
		select {
		case trigger := <-ctx.eventTrigger:
			if err := despatchEvent(trigger, ctx.state, ctx); err != nil {
				t.Errorf("%v", err)
			}
		case <-ctx.restartTimer.C:
			if ctx.state != StateAttestWait {
				t.Errorf("Expected %s, Got %s", StateAttestWait.String(), ctx.state.String())
			}
			return
		}
	}
}

func TestControllerNotAvbleInAttestEscrowWait(t *testing.T) {
	fmt.Println("--------TestControllerNotAvbleInAttestEscrowWait----")
	ctx := initTest()

	go func() {
		ctx.state = StateAttestWait
		simulateControllerReqFailure = true
		ctx.eventTrigger <- EventAttestSuccessful
	}()
	for {
		select {
		case trigger := <-ctx.eventTrigger:
			if err := despatchEvent(trigger, ctx.state, ctx); err != nil {
				t.Errorf("%v", err)
			}
		case <-ctx.restartTimer.C:
			if ctx.state != StateAttestEscrowWait {
				t.Errorf("Expected %s, Got %s", StateAttestEscrowWait.String(), ctx.state.String())
			}
			return
		}
	}
}

func TestNoVerifierInNonceWait(t *testing.T) {
	fmt.Println("--------TestNoVerifierInNonceWait----")
	ctx := initTest()
	stopTrigger := make(chan int)

	go func() {
		ctx.state = StateNone
		simulateNoVerifier = true
		ctx.eventTrigger <- EventInitialize
		time.Sleep(1 * time.Second)
		stopTrigger <- 1
	}()
	for {
		select {
		case trigger := <-ctx.eventTrigger:
			if err := despatchEvent(trigger, ctx.state, ctx); err != nil {
				t.Errorf("%v", err)
			}
		case <-stopTrigger:
			if ctx.state != StateNonceWait {
				t.Errorf("Expected %s, Got %s", StateNonceWait.String(), ctx.state.String())
			}
			return
		}
	}
}
