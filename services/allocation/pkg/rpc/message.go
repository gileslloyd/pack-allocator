package rpc

type Message struct {
	data map[string]interface{}
}

func NewMessage(data map[string]interface{}) *Message {
	return &Message{
		data: data,
	}
}

func (m Message) Get(key string, def string) string {
	val := m.data[key]

	if val == nil {
		return def
	}

	return val.(string)
}
