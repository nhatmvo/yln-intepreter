package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLIGAL = "ILLIGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	EXCLAMATION = "!"
	ASTERISK    = "*"
	SLASH       = "/"

	EQ     = "=="
	NOT_EQ = "!="
	GT     = ">"
	LT     = "<"

	COMMA      = ","
	SEMILCOLON = ";"

	LPREN  = "("
	RPREN  = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}
