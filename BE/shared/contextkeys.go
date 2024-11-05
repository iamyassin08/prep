package shagreen

type contextKey int

const (
	ContextKeyRequestId contextKey = iota
	ContextKeyClaims    contextKey = iota
)
