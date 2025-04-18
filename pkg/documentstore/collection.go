package documentstore

import (
	log "github.com/sirupsen/logrus"
	"lesson4/pkg/err"
)

type Collection struct {
	Documents map[string]Document `json:"documents,omitempty"`
	Config    CollectionConfig    `json:"config"`
}

type CollectionConfig struct {
	PrimaryKey string `json:"cgg"`
}

func (s *Collection) Put(doc Document) error {
	// Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`

	keyFilds, ok := doc.Fields[s.Config.PrimaryKey]
	if !ok {
		log.Error("error: Document must contain a key field")
		return err.ErrUnsupportedDocumentField
	}

	if keyFilds.Type != DocumentFieldTypeString {
		log.Error("error: Key field must be of type string")
		return err.ErrUnsupportedDocumentField
	}
	if s.Documents == nil {
		s.Documents = map[string]Document{}
	}
	s.Documents[s.Config.PrimaryKey] = doc
	log.Info("document added - ", s.Config.PrimaryKey)
	return nil
}

func (s *Collection) Get(key string) (*Document, error) {
	if doc, exists := s.Documents[key]; exists {
		return &doc, nil
	}
	log.Infof("document not found - %s", key)
	return nil, err.ErrDocumentNotFound
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.Documents[key]; exists {
		delete(s.Documents, key)
		log.Info("document delete - ", key)
		return true
	}
	return false
}

func (s *Collection) List() []Document {
	sList := make([]Document, 0, len(s.Documents))
	for _, v := range s.Documents {
		sList = append(sList, v)
	}
	return sList
}
