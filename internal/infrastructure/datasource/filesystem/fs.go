package filesystem

import (
	"hackbar-copilot/internal/domain/recipe"
)

type Filesystem interface {
	Recipe() recipe.Repository
	SavePersistently() error
}

// NewRepository creates a new instance that implements
// filesystem.Filesystem and recipes.Repository.
//
// Example usage:
//
//	repository := filesystem.NewRepository("/path/to/data-dir")
//	err := repository.Save(...)
//	// handle error ...
//	err = repository.SavePersistently()
//	// handle error ...
//
// It create files to baseDir follow the structure:
//
//	```
//	/path/to/data-dir
//	├── 1_recipe_groups.toml
//	├── 2_recipe_types.toml
//	└── 3_glass_types.toml
//	```
func NewRepository(baseDir string) (Filesystem, error) {
	fsR := newFSR(baseDir)
	fsW := newFSW(baseDir)
	data, err := loadData(fsR)
	if err != nil {
		return nil, err
	}

	fs := filesystem{read: fsR, write: fsW, data: data}

	if fs.data.isEmpty() {
		// check base dir is writable
		err = fs.SavePersistently()
		if err != nil {
			return nil, err
		}
	}

	return &fs, nil
}

type filesystem struct {
	read  fsR
	write fsW
	data  data
}

type data struct {
	recipeGroups []recipe.RecipeGroup
	recipeTypes  map[string]recipe.RecipeType
	glassTypes   map[string]recipe.GlassType
}

func (data data) isEmpty() bool {
	return len(data.recipeGroups) == 0 && len(data.recipeTypes) == 0 && len(data.glassTypes) == 0
}

// Recipe implements Filesystem.
func (f *filesystem) Recipe() recipe.Repository {
	return &recipeRepository{f}
}
