package sessions

//Don't use this!!!!
//This shows how to implement the sessions methods using im-memory storage i.e. map
//For real life example, you should use a DBMS e.g. MySQL, Mongo, etc

type data map[string]interface{}

type inmemory struct {
	memory map[string]data
}

//SessionExists checks if session already exists
func (i *inmemory) SessionExists(sessionId string) bool {
	if _, ok := i.memory[sessionId]; ok {
		return true
	}
	return false
}

func (i *inmemory) StartSession(sessionId string, data map[string]interface{}) error {
	i.memory[sessionId] = data
	return nil
}

func (i *inmemory) UpdateSession(sessionId, key string, data interface{}) error {
	d := i.memory[sessionId]
	d[key] = data
	i.memory[sessionId] = d
	return nil
}

func (i *inmemory) RetrieveData(sessionId, key string) interface{} {
	d := i.memory[sessionId]
	return d[key]
}

func (i *inmemory) DeleteData(sessionId, key string) error {
	d := i.memory[sessionId]
	delete(d, key)
	return nil
}

func (i *inmemory) EndSession(sessionId string) error {
	delete(i.memory, sessionId)
	return nil 
}
