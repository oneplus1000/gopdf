package core

type ValueRecord struct {
	XPlacement int  // Horizontal adjustment for placement-in design units
	YPlacement int  // Vertical adjustment for placement, in design units
	XAdvance   int  // Horizontal adjustment for advance, in design units (only used for horizontal writing)
	YAdvance   int  // Vertical adjustment for advance, in design units (only used for vertical writing)
	XPlaDevice uint // Offset to Device table (non-variable font) / VariationIndex table (variable font) for horizontal placement, from beginning of PosTable (may be NULL)
	YPlaDevice uint // Offset to Device table (non-variable font) / VariationIndex table (variable font) for vertical placement, from beginning of PosTable (may be NULL)
	XAdvDevice uint // Offset to Device table (non-variable font) / VariationIndex table (variable font) for horizontal advance, from beginning of PosTable (may be NULL)
	YAdvDevice uint // Offset to Device table (non-variable font) / VariationIndex table (variable font) for vertical advance, from beginning of PosTable (may be NULL)
}
