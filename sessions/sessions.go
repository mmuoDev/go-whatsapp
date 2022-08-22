/**
Just like on websites, you may want to store some variables in a session while building a whatsapp chatbot.
While it is easy doing this on web applications, it’s not for chatbots because these variables don’t
persist in the session as the user continually interacts with your chatbot.
For example, if a user sends a message containing his name and you don't store that in a session, by the
time you are sending a reply, everything in the session is lost!
Reference - https://medium.com/@mmuoDev/how-to-manage-session-variables-while-building-a-whatsapp-facebook-chatbot-fa20131c8c95
**/
package sessions

//SessionManager represents methods database systems e.g. Mongo, MySQL need to implement
//in order to manage sessions for a chatbot.
//sessionId should be the user's whatsapp number
//data to be stored should be a map of data type and it's value.
type SessionManager interface {
	StartSession(sessionId string, data interface{}) error
	SessionExists(sessionId string) (bool, error)
	RetrieveData(sessionId string, result interface{})
	UpdateSession(sessionId string, data interface{}) error
	EndSession(sessionId string) error
}
