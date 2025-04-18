package documentstore

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"lesson4/pkg/err"
	"os"
)

type Store struct {
	Collections map[string]*Collection `json:"collections,omitempty"`
}

func NewStore() *Store {
	return &Store{
		Collections: make(map[string]*Collection),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (error, *Collection) {
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створеня то повертаємо `false` та nil
	if _, exists := s.Collections[name]; exists {
		log.Info("collection already exists")
		return err.ErrCollectionAlreadyExists, nil
	}
	coll := &Collection{
		Documents: make(map[string]Document),
		Config:    *cfg}
	s.Collections[name] = coll
	log.Info("collection added - ", coll.Config.PrimaryKey)
	return nil, coll
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	if colect, ok := s.Collections[name]; ok {
		return colect, nil
	}
	log.Infof("collection not found - %s", name)
	return nil, err.ErrCollectionNotFound
}

func (s *Store) DeleteCollection(name string) bool {
	if _, ok := s.Collections[name]; ok {
		delete(s.Collections, name)
		log.Info("collection delete - ", name)
		return true
	}
	return false
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	// Функція повинна створити та проініціалізувати новий `Store`
	// зі всіма колекціями та даними з вхідного дампу.
	var s Store
	if err := json.Unmarshal(dump, &s); err != nil {
		log.Warning("")
		return nil, err
	}
	if len(s.Collections) == 0 {
		log.Warning("collection not added")
		return nil, err.ErrNotFound
	}
	return &s, nil
}

func (s *Store) Dump() ([]byte, error) {
	// Методи повинен віддати дамп нашого стору в який включені дані про колекції та документ
	sToJson, err := json.MarshalIndent(s, " ", "")
	if err != nil {
		return nil, err
	}
	return sToJson, nil
}

//
//// Значення яке повертає метод `store.Dump()` має без помилок оброблятись функцією `NewStoreFromDump`
//

func NewStoreFromFile(filename string) (*Store, error) {
	// Робить те ж саме що і функція `NewStoreFromDump`, але сам дамп має діставатись з файлу
	f := fmt.Sprintf("%s.json", filename)
	dump, err := os.ReadFile(f)
	if err != nil {
		log.Error("file not read")
		return nil, err
	}
	log.Infof("file read successfully %s", f)
	var s Store
	if err := json.Unmarshal(dump, &s); err != nil {

		return nil, err
	}
	if len(s.Collections) == 0 {
		log.Warning("no collections found in store from file")
		return nil, fmt.Errorf("no collections in store")
	}
	return &s, nil
}

func (s *Store) DumpToFile(filename string) error {
	// Робить те ж саме що і метод  `Dump`, але записує у файл замість того щоб повертати сам дамп
	sDump, err := s.Dump()
	if err != nil {

		fmt.Println(err)
	}
	f := fmt.Sprintf("%s.json", filename)

	return os.WriteFile(f, sDump, 0644)
}
