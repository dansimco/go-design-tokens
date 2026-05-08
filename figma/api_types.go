package figma

type PublishVariablesBody struct {
	VariableCollections []VariableCollectionChange `json:"variableCollections,omitempty"`
	VariableModes       []VariableModeChange       `json:"variableModes,omitempty"`
	Variables           []VariableChange           `json:"variables,omitempty"`
	VariableModeValues  []VariableModeValue        `json:"variableModeValues,omitempty"`
}

type VariableCollectionChange struct {
	Action        string `json:"action"`
	ID            string `json:"id"`
	Name          string `json:"name,omitempty"`
	InitialModeID string `json:"initialModeId,omitempty"`
}

type VariableModeChange struct {
	Action               string `json:"action"`
	ID                   string `json:"id"`
	Name                 string `json:"name,omitempty"`
	VariableCollectionID string `json:"variableCollectionId,omitempty"`
}

type VariableChange struct {
	Action               string `json:"action"`
	ID                   string `json:"id"`
	Name                 string `json:"name,omitempty"`
	VariableCollectionID string `json:"variableCollectionId,omitempty"`
	ResolvedType         string `json:"resolvedType,omitempty"`
	Description          string `json:"description,omitempty"`
}

type VariableModeValue struct {
	VariableID string      `json:"variableId"`
	ModeID     string      `json:"modeId"`
	Value      interface{} `json:"value"`
}

type RGBAValue struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}
