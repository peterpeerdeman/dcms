package resty

var messageChannel chan message

func Init() {
	messageChannel = make(chan message)
	go DocumentProcessor(messageChannel)

	documentTypes = make(map[string] documentType)
}
