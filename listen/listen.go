package listen

type Listener interface {
	Text (string) bool 
	Intent (interface{}) (interface{}, error) 
	Attachments (interface{}) (interface{}, error) //image,video,documents
	Location (interface{}) (interface{}, error)
}

type Listen struct {}

func NewListener() *Listen {
	return &Listen{}
}

func (l *Listen) Text (string) bool {
	return true
}

func (l *Listen) Intent (interface{}) (interface{}, error)  {
	return nil, nil
}

func (l *Listen) Attachments (interface{}) (interface{}, error) {
	return nil, nil 
}

func (l *Listen) Location (interface{}) (interface{}, error) {
	return nil, nil 
}

