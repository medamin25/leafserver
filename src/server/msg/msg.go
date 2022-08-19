package msg

import (
	"github.com/name5566/leaf/network/json"
	//"github.com/name5566/leaf/network/protobuf"
)

// Create a JSON Processor（or protobuf if you like）
var Processor = json.NewProcessor()

// One struct for one message
// Contains a string member
type Hello struct {
    Name string
}

func init() {
    // Register message Hello
    Processor.Register(&Hello{})
}
