package message

const (
	TYPE_REQUEST  = "request"
	TYPE_RESPONSE = "response"
	TYPE_ERROR    = "error"
	TYPE_INFO     = "info"
)

// IMessageClient ...
type IMessageClient interface {
	connectService(config *OptionsMessageCLient) error
	PublishMessage(routing string, params *MessageParam) error
	ReceiveMessage(routing string) (<-chan MessageParam, error)
	ReceiveOneMessage(routing string) (*MessageParam, error)
}

type OptionsMessageCLient struct {
	URL      string                 `json:"url"`
	Password string                 `json:"password"`
	Args     map[string]interface{} `json:"args"`
	Driver   string                 `json:"driver"`
}

type MessageParam struct {
	ID     string `json:"id"`
	Body   []byte `json:"body"`
	Method string `json:"method"`
}

func (o *OptionsMessageCLient) ConfigureMessageQueue() (*IMessageClient, error) {
	var iMessage IMessageClient
	switch o.Driver {
	case "redis":

		red := &redisMessage{}

		errRed := red.connectService(o)
		if errRed != nil {
			return nil, errRed

		}
		iMessage = red

	}
	return &iMessage, nil
}
