package m3u8

import "fmt"

// SessionKeyItem represents a set of EXT-X-SESSION-KEY attributes
type SessionKeyItem struct {
	Encryptable *Encryptable
}

func NewSessionKeyItem(text string) (*SessionKeyItem, error) {
	attributes := ParseAttributes(text)
	return &SessionKeyItem{
		Encryptable: NewEncryptable(attributes),
	}, nil
}

func (ski *SessionKeyItem) String() string {
	return fmt.Sprintf("%s:%v", SessionKeyItemTag, ski.Encryptable.String())
}
