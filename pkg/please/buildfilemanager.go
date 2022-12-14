package please

import (
	"errors"
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

type BuildFileManager struct {
	buildFiles map[string]*BuildFile

	mux *sync.Mutex
}

func NewBuildFileManager() *BuildFileManager {
	return &BuildFileManager{
		buildFiles: map[string]*BuildFile{},
		mux:        &sync.Mutex{},
	}
}

func (m *BuildFileManager) GetBuildFileForTarget(target *Target) (*BuildFile, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if bf, ok := m.buildFiles[target.BuildFilePath()]; ok {
		return bf, nil
	}

	var bf *BuildFile
	if _, err := os.Stat(target.BuildFilePath()); errors.Is(err, os.ErrNotExist) {
		var err error
		bf, err = NewBuildFile(target.BuildFilePath())
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		bf, err = LoadBuildFileFromFile(target.BuildFilePath())
		if err != nil {
			return nil, err
		}
	}

	m.buildFiles[target.BuildFilePath()] = bf

	return bf, nil
}

func (m *BuildFileManager) SaveAll() error {
	files := []string{}
	for fName, bf := range m.buildFiles {
		if err := bf.Save(); err != nil {
			return err
		}
		files = append(files, fName)
	}

	log.Info().
		Strs("files", files).
		Int("amount", len(files)).
		Msg("saved BUILD files")

	return nil
}
