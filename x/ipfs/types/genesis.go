package types

// GenesisState - all track state that must be provided at genesis
type GenesisState struct {
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
	}
}

// ValidateGenesis validates the track genesis parameters
func ValidateGenesis(data GenesisState) error {
	return nil
}