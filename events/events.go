package events

//Attachment represents an attachment details
//Attachment may include but not limited to contacts, videos, images, etc. 
//You can filter by the content type. 
//Contacts have a content type of text/vcard
type Attachment struct {
	MediaURL  string
	MediaName string
	MediaType string
}

//Location represents a location details
type Location struct {
	Longitude string
	Latitude  string
}
