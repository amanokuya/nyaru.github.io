package lexier

import (
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	peek      byte             //read in
	line      int              //in which line
	reader    *bufio.Reader    //read 字节流
	key_words map[string]Token //储存keyword
}

func NewLexer(source string) Lexer {
	str := strings.NewReader(source)
	source_reader := bufio.NewReaderSize(str, len(source))
	lexer := Lexer{
		line:      1,
		reader:    source_reader,
		key_words: make(map[string]Token),
	}

	lexer.reserve()

	return lexer
}

func (l *Lexer) reserve() {
	key_words := GetKeyWords()
	for _, key_word := range key_words {
		l.key_words[key_word.ToString()] = key_word.Tag
	}
}

func (l *Lexer) Readch() error {
	char, err := l.reader.ReadByte()
	l.peek = char
	return err
}

func (l *Lexer) ReadCharacter(c byte) (bool, error) {
	err := l.Readch()
	if err != nil {
		return false, err
	}

	if l.peek != c {
		return false, nil
	}

	l.peek = ' '
	return true, nil
}

func (l *Lexer) Scan() (Token, error) {
	for {
		err := l.Readch()
		if err != nil {
			return NewToken(ERROR), err
		}
		if l.peek == ' ' || l.peek == '\t' {
			continue
		} else if l.peek == '\n' {
			break
		}
	}

	switch l.peek {
	case '+':
		return NewToken(Plus), nil
	case '-':
		return NewToken(Sub), nil
	case '*':
		return NewToken(Mul), nil
	case '/':
		return NewToken(Div), nil
	case ';':
		return NewToken(SEMI), nil
	case '(':
		return NewToken(Right_Bracket), nil
	case ')':
		return NewToken(Left_Bracket), nil
	case '[':
		return NewToken(RMB), nil
	case ']':
		return NewToken(LMB), nil
	case '{':
		return NewToken(RGB), nil
	case '}':
		return NewToken(LGB), nil
	case '!':
		return NewToken(NE), nil
	case '=':
		if ok, _ := l.ReadCharacter('='); ok {
			word := NewWordToken("==", EQ)
			return word.Tag, nil
		} else {
			return NewToken(RETURN), nil
		}
	case '<':
		if ok, _ := l.ReadCharacter('='); ok {
			word := NewWordToken("<=", LE)
			return word.Tag, nil
		} else {
			return NewToken(LT), nil
		}
	case '>':
		if ok, _ := l.ReadCharacter('='); ok {
			word := NewWordToken(">=", GE)
			return word.Tag, nil
		} else {
			return NewToken(GT), nil
		}
	}

	if unicode.IsNumber(rune(l.peek)) {
		var v int
		var err error
		for {
			num, err := strconv.Atoi(string(l.peek))
			if err != nil {
				break
			}
			v = v*10 + num
			l.Readch()
		}
		return NewToken(NUM), err
	}

	if unicode.IsLetter(rune(l.peek)) {
		var buffer []byte
		for {
			buffer = append(buffer, l.peek)
			l.Readch()
			if !unicode.IsLetter(rune(l.peek)) {
				break
			}
		}

		token, ok := l.key_words[string(buffer)]
		if ok {
			return token, nil
		}
		return NewToken(ID), nil
	}
	return NewToken(ERROR), nil
}
