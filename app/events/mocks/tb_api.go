// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	tbapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
)

// TbAPIMock is a mock implementation of events.TbAPI.
//
//	func TestSomethingThatUsesTbAPI(t *testing.T) {
//
//		// make and configure a mocked events.TbAPI
//		mockedTbAPI := &TbAPIMock{
//			GetChatFunc: func(config tbapi.ChatInfoConfig) (tbapi.Chat, error) {
//				panic("mock out the GetChat method")
//			},
//			GetUpdatesChanFunc: func(config tbapi.UpdateConfig) tbapi.UpdatesChannel {
//				panic("mock out the GetUpdatesChan method")
//			},
//			RequestFunc: func(c tbapi.Chattable) (*tbapi.APIResponse, error) {
//				panic("mock out the Request method")
//			},
//			SendFunc: func(c tbapi.Chattable) (tbapi.Message, error) {
//				panic("mock out the Send method")
//			},
//		}
//
//		// use mockedTbAPI in code that requires events.TbAPI
//		// and then make assertions.
//
//	}
type TbAPIMock struct {
	// GetChatFunc mocks the GetChat method.
	GetChatFunc func(config tbapi.ChatInfoConfig) (tbapi.Chat, error)

	// GetUpdatesChanFunc mocks the GetUpdatesChan method.
	GetUpdatesChanFunc func(config tbapi.UpdateConfig) tbapi.UpdatesChannel

	// RequestFunc mocks the Request method.
	RequestFunc func(c tbapi.Chattable) (*tbapi.APIResponse, error)

	// SendFunc mocks the Send method.
	SendFunc func(c tbapi.Chattable) (tbapi.Message, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetChat holds details about calls to the GetChat method.
		GetChat []struct {
			// Config is the config argument value.
			Config tbapi.ChatInfoConfig
		}
		// GetUpdatesChan holds details about calls to the GetUpdatesChan method.
		GetUpdatesChan []struct {
			// Config is the config argument value.
			Config tbapi.UpdateConfig
		}
		// Request holds details about calls to the Request method.
		Request []struct {
			// C is the c argument value.
			C tbapi.Chattable
		}
		// Send holds details about calls to the Send method.
		Send []struct {
			// C is the c argument value.
			C tbapi.Chattable
		}
	}
	lockGetChat        sync.RWMutex
	lockGetUpdatesChan sync.RWMutex
	lockRequest        sync.RWMutex
	lockSend           sync.RWMutex
}

// GetChat calls GetChatFunc.
func (mock *TbAPIMock) GetChat(config tbapi.ChatInfoConfig) (tbapi.Chat, error) {
	if mock.GetChatFunc == nil {
		panic("TbAPIMock.GetChatFunc: method is nil but TbAPI.GetChat was just called")
	}
	callInfo := struct {
		Config tbapi.ChatInfoConfig
	}{
		Config: config,
	}
	mock.lockGetChat.Lock()
	mock.calls.GetChat = append(mock.calls.GetChat, callInfo)
	mock.lockGetChat.Unlock()
	return mock.GetChatFunc(config)
}

// GetChatCalls gets all the calls that were made to GetChat.
// Check the length with:
//
//	len(mockedTbAPI.GetChatCalls())
func (mock *TbAPIMock) GetChatCalls() []struct {
	Config tbapi.ChatInfoConfig
} {
	var calls []struct {
		Config tbapi.ChatInfoConfig
	}
	mock.lockGetChat.RLock()
	calls = mock.calls.GetChat
	mock.lockGetChat.RUnlock()
	return calls
}

// GetUpdatesChan calls GetUpdatesChanFunc.
func (mock *TbAPIMock) GetUpdatesChan(config tbapi.UpdateConfig) tbapi.UpdatesChannel {
	if mock.GetUpdatesChanFunc == nil {
		panic("TbAPIMock.GetUpdatesChanFunc: method is nil but TbAPI.GetUpdatesChan was just called")
	}
	callInfo := struct {
		Config tbapi.UpdateConfig
	}{
		Config: config,
	}
	mock.lockGetUpdatesChan.Lock()
	mock.calls.GetUpdatesChan = append(mock.calls.GetUpdatesChan, callInfo)
	mock.lockGetUpdatesChan.Unlock()
	return mock.GetUpdatesChanFunc(config)
}

// GetUpdatesChanCalls gets all the calls that were made to GetUpdatesChan.
// Check the length with:
//
//	len(mockedTbAPI.GetUpdatesChanCalls())
func (mock *TbAPIMock) GetUpdatesChanCalls() []struct {
	Config tbapi.UpdateConfig
} {
	var calls []struct {
		Config tbapi.UpdateConfig
	}
	mock.lockGetUpdatesChan.RLock()
	calls = mock.calls.GetUpdatesChan
	mock.lockGetUpdatesChan.RUnlock()
	return calls
}

// Request calls RequestFunc.
func (mock *TbAPIMock) Request(c tbapi.Chattable) (*tbapi.APIResponse, error) {
	if mock.RequestFunc == nil {
		panic("TbAPIMock.RequestFunc: method is nil but TbAPI.Request was just called")
	}
	callInfo := struct {
		C tbapi.Chattable
	}{
		C: c,
	}
	mock.lockRequest.Lock()
	mock.calls.Request = append(mock.calls.Request, callInfo)
	mock.lockRequest.Unlock()
	return mock.RequestFunc(c)
}

// RequestCalls gets all the calls that were made to Request.
// Check the length with:
//
//	len(mockedTbAPI.RequestCalls())
func (mock *TbAPIMock) RequestCalls() []struct {
	C tbapi.Chattable
} {
	var calls []struct {
		C tbapi.Chattable
	}
	mock.lockRequest.RLock()
	calls = mock.calls.Request
	mock.lockRequest.RUnlock()
	return calls
}

// Send calls SendFunc.
func (mock *TbAPIMock) Send(c tbapi.Chattable) (tbapi.Message, error) {
	if mock.SendFunc == nil {
		panic("TbAPIMock.SendFunc: method is nil but TbAPI.Send was just called")
	}
	callInfo := struct {
		C tbapi.Chattable
	}{
		C: c,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	return mock.SendFunc(c)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//
//	len(mockedTbAPI.SendCalls())
func (mock *TbAPIMock) SendCalls() []struct {
	C tbapi.Chattable
} {
	var calls []struct {
		C tbapi.Chattable
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}