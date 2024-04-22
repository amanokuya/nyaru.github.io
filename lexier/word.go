package lexier

type Word struct {
	lexeme string
	Tag    Token
}

func NewWordToken(s string, tag Tag) Word {
	return Word{
		lexeme: s,
		Tag:    NewToken(tag),
	}
}

func GetKeyWords() []Word {
	key_words := []Word{}
	key_words = append(key_words, NewWordToken("==", EQ))
	key_words = append(key_words, NewWordToken("begin", Begin))
	key_words = append(key_words, NewWordToken("if", IF))
	key_words = append(key_words, NewWordToken("then", Then))
	key_words = append(key_words, NewWordToken("while", While))
	key_words = append(key_words, NewWordToken("do", DO))
	key_words = append(key_words, NewWordToken("end", END))
	key_words = append(key_words, NewWordToken("=", RETURN))
	key_words = append(key_words, NewWordToken("+", Plus))
	key_words = append(key_words, NewWordToken("-", Sub))
	key_words = append(key_words, NewWordToken("*", Mul))
	key_words = append(key_words, NewWordToken("/", Div))
	key_words = append(key_words, NewWordToken("<", LT))
	key_words = append(key_words, NewWordToken("<=", LE))
	key_words = append(key_words, NewWordToken("!=", NE))
	key_words = append(key_words, NewWordToken(">", GT))
	key_words = append(key_words, NewWordToken(">=", GE))
	key_words = append(key_words, NewWordToken(";", SEMI))
	key_words = append(key_words, NewWordToken("(", Left_Bracket))
	key_words = append(key_words, NewWordToken(")", Right_Bracket))
	key_words = append(key_words, NewWordToken("[", LMB))
	key_words = append(key_words, NewWordToken("]", RMB))
	key_words = append(key_words, NewWordToken("{", LGB))
	key_words = append(key_words, NewWordToken("}", RGB))
	key_words = append(key_words, NewWordToken("identifier", ID))
	key_words = append(key_words, NewWordToken("number", NUM))

	return key_words
}

func (w *Word) ToString() string {
	return w.lexeme
}
