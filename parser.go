package typml

type Expr struct {
	left  Token
	right Token
}

type Func struct {
	name string
	args []Token
}
