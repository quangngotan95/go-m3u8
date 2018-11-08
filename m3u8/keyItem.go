package m3u8

import "fmt"

// KeyItem represents a set of EXT-X-KEY attributes
type KeyItem struct {
	Encryptable *Encryptable
}

// NewKeyItem parses a text line and returns a *KeyItem
func NewKeyItem(text string) (*KeyItem, error) {
	attributes := ParseAttributes(text)
	return &KeyItem{
		Encryptable: NewEncryptable(attributes),
	}, nil
}

func (ki *KeyItem) String() string {
	return fmt.Sprintf("%s:%v", KeyItemTag, ki.Encryptable.String())
}
