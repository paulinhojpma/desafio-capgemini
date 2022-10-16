package routes

import (
	"defafio-cap/sequence-validator/message"
)

type Manager struct {
	routes map[string]func(*message.MessageParam) *message.MessageParam
}

// AddRoute ...
func (m *Manager) AddRoute(name string, f func(*message.MessageParam) *message.MessageParam) {
	m.routes[name] = f
}

// CallService ..
func (m *Manager) CallService(route string, msg *message.MessageParam) *message.MessageParam {
	newMsg := m.routes[route](msg)
	return newMsg
}

// NewManager ...
func NewManager() *Manager {
	manager := &Manager{}
	manager.routes = make(map[string]func(*message.MessageParam) *message.MessageParam)
	return manager
}

func (m *Manager) ManagerMessage(client message.IMessageClient, msg *message.MessageParam) error {

	if msg != nil {
		return client.PublishMessage(msg.ID, msg)
	}
	return nil

}
