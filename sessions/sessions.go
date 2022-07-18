package sessions

// Just like on websites, you may want to store some variables in a session while building a whatsapp chatbot.
// While it is easy doing this on web applications, it’s not for chatbots because these variables don’t
// persist in the session as the user continually interacts with your chatbot.
// For example, if a user sends a message containing his name and you don't store that in a session, by the
// time you are sending a reply, everything in the session is lost!
// So what’s the solution? Database!
// https://medium.com/@mmuoDev/how-to-manage-session-variables-while-building-a-whatsapp-facebook-chatbot-fa20131c8c95

//SessionManager represents methods needed for database systems e.g. Mongo, MySQL needs to implement
//in order to manage sessions for a chatbot
//sessionId should be the whatsapp number
//data to be stored should be a map of data type and it's value.
type SessionManager interface {
	SessionExists(sessionId string) bool
	StartSession(sessionId string, data map[string]interface{}) error
	UpdateSession(sessionId, key string, data interface{}) error
	RetrieveData(sessionId, key string) interface{}
	DeleteData(sessionId, key string) error
	EndSession(sessionId string) error
}
