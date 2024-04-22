package lexier

type Tag uint32

const (
	EQ Tag = iota + 256 //==
	Begin
	IF
	Then
	While
	DO
	END
	RETURN
	Plus
	Sub //-
	Mul //*
	Div ///
	LT
	LE
	NE
	GT
	GE
	SEMI          //;
	Left_Bracket  //(
	Right_Bracket //)
	LMB           //[
	RMB           //]
	LGB           //{
	RGB           //}
	ID            //identifier
	NUM           //number
	ERROR         //error
)

var token_map = make(map[Tag]string)

func init() {
	token_map[EQ] = "=="
	token_map[Begin] = "begin"
	token_map[IF] = "if"
	token_map[Then] = "then"
	token_map[While] = "while"
	token_map[DO] = "do"
	token_map[END] = "end"
	token_map[RETURN] = "="
	token_map[Plus] = "+"
	token_map[Sub] = "-"
	token_map[Mul] = "*"
	token_map[Div] = "/"
	token_map[LT] = "<"
	token_map[LE] = "<="
	token_map[NE] = "!="
	token_map[GT] = ">"
	token_map[GE] = ">="
	token_map[SEMI] = ";"
	token_map[Right_Bracket] = "("
	token_map[Left_Bracket] = ")"
	token_map[LMB] = "["
	token_map[RMB] = "]"
	token_map[LGB] = "{"
	token_map[RGB] = "}"
	token_map[ID] = "identifier"
	token_map[NUM] = "number"
	token_map[ERROR] = "error"
}

type Token struct {
	Tag Tag
}

func (t *Token) ToString() string {
	return token_map[t.Tag]
}

func NewToken(tag Tag) Token {
	return Token{
		Tag: tag,
	}
}
