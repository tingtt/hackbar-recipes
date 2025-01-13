// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type GlassType struct {
	Name        string  `json:"name"`
	ImageURL    *string `json:"imageURL,omitempty"`
	Description *string `json:"description,omitempty"`
}

type InputAsMenuArgs struct {
	Flavor *string `json:"flavor,omitempty"`
}

type InputAsMenuItemArgs struct {
	ImageURL *string `json:"imageURL,omitempty"`
	Price    int     `json:"price"`
}

type InputGlassType struct {
	Name        string  `json:"name"`
	ImageURL    *string `json:"imageURL,omitempty"`
	Description *string `json:"description,omitempty"`
	Save        *bool   `json:"save,omitempty"`
}

type InputRecipe struct {
	Name       string               `json:"name"`
	RecipeType *InputRecipeType     `json:"recipeType,omitempty"`
	GlassType  *InputGlassType      `json:"glassType,omitempty"`
	Steps      []*InputStep         `json:"steps,omitempty"`
	AsMenu     *InputAsMenuItemArgs `json:"asMenu,omitempty"`
}

type InputRecipeGroup struct {
	Name     string           `json:"name"`
	ImageURL *string          `json:"imageURL,omitempty"`
	Recipes  []*InputRecipe   `json:"recipes,omitempty"`
	AsMenu   *InputAsMenuArgs `json:"asMenu,omitempty"`
}

type InputRecipeType struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Save        *bool   `json:"save,omitempty"`
}

type InputStep struct {
	Material    *string `json:"material,omitempty"`
	Amount      *string `json:"amount,omitempty"`
	Description *string `json:"description,omitempty"`
}

type MenuGroup struct {
	Name        string      `json:"name"`
	ImageURL    *string     `json:"imageURL,omitempty"`
	Flavor      *string     `json:"flavor,omitempty"`
	Items       []*MenuItem `json:"items,omitempty"`
	MinPriceYen int         `json:"minPriceYen"`
}

type MenuItem struct {
	Name       string   `json:"name"`
	ImageURL   *string  `json:"imageURL,omitempty"`
	Materials  []string `json:"materials,omitempty"`
	OutOfStock bool     `json:"outOfStock"`
	PriceYen   int      `json:"priceYen"`
	Recipe     *Recipe  `json:"recipe,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type Recipe struct {
	Name  string      `json:"name"`
	Type  *RecipeType `json:"type,omitempty"`
	Glass *GlassType  `json:"glass,omitempty"`
	Steps []*Step     `json:"steps,omitempty"`
}

type RecipeGroup struct {
	Name     string    `json:"name"`
	ImageURL *string   `json:"imageURL,omitempty"`
	Recipes  []*Recipe `json:"recipes,omitempty"`
}

type RecipeType struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type Step struct {
	Material    *string `json:"material,omitempty"`
	Amount      *string `json:"amount,omitempty"`
	Description *string `json:"description,omitempty"`
}
