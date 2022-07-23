package sessions

//Don't use this!!!!
//This shows how to implement the sessions methods using im-memory storage i.e. map
//For real life example, you should use a DBMS e.g. MySQL, Mongo, etc

type Data map[string]interface{}

type Inmemory struct {
	Memory map[string]Data
}

//SessionExists checks if session already exists
func (i *Inmemory) SessionExists(sessionId string) bool {
	if _, ok := i.Memory[sessionId]; ok {
		return true
	}
	return false
}

func (i *Inmemory) StartSession(sessionId string, data map[string]interface{}) error {
	i.Memory[sessionId] = data
	return nil
}

func (i *Inmemory) UpdateSession(sessionId, key string, data interface{}) error {
	d := i.Memory[sessionId]
	d[key] = data
	i.Memory[sessionId] = d
	return nil
}

func (i *Inmemory) RetrieveData(sessionId, key string) interface{} {
	d := i.Memory[sessionId]
	return d[key]
}

func (i *Inmemory) DeleteData(sessionId, key string) error {
	d := i.Memory[sessionId]
	delete(d, key)
	return nil
}

func (i *Inmemory) EndSession(sessionId string) error {
	delete(i.Memory, sessionId)
	return nil 
}
