package dockerlang

import (
	"testing"
)

func TestTokenizeNumber(t *testing.T) {
	conf := &Config{SrcFileName: "test/test.doc"}
	compt := NewCompterpreter(conf)

	err := compt.LoadSourceCode()
	if err != nil {
		t.Error(err)
	}

	// advance ptr to first character
	compt.Advance()

	compt.TokenizeNumber(compt.CurrentChar)

	if compt.CurrentToken.Value != "1234" {
		t.Error("incorrect token!")
	}
}

func TestGetNextToken(t *testing.T) {
	conf := &Config{SrcFileName: "test/test.doc"}
	compt := NewCompterpreter(conf)

	err := compt.LoadSourceCode()
	if err != nil {
		t.Error(err)
	}

	// advance ptr to first character
	compt.Advance()

	s, err := compt.GetNextToken()
	if err != nil {
		t.Error(err)
	}
	if s.Value != "1234" {
		t.Errorf("incorrect first token! Expected '1234' got '%s'", s)
	}

	s, err = compt.GetNextToken()
	if err != nil {
		t.Error(err)
	}
	if s.Value != "5678" {
		t.Errorf("incorrect second token! Expected '5678' got '%s'", s)
	}
}

func TestIsOperator(t *testing.T) {
	conf := &Config{SrcFileName: "test/test.doc"}
	compt := NewCompterpreter(conf)
	for _, operator := range []rune{'+', '†', '*', '‡', '%'} {
		ok := compt.IsOperator(operator)
		if !ok {
			t.Error("not an operator! but it should be!")
		}
	}
	for _, operator := range []rune{'q', '!', '❧', '0', ' '} {
		ok := compt.IsOperator(operator)
		if ok {
			t.Error("that was an operator! but it shouldn't be!")
		}
	}
}

func TestIsPunctuation(t *testing.T) {
	conf := &Config{SrcFileName: "test/test.doc"}
	compt := NewCompterpreter(conf)
	for _, operator := range []rune{'(', ')', '(', ')'} {
		ok := compt.IsPunctuation(operator)
		if !ok {
			t.Error("not punctuation! but it should be!")
		}
	}
	for _, operator := range []rune{'q', '!', '❧', '0', ' '} {
		ok := compt.IsPunctuation(operator)
		if ok {
			t.Error("that was puncuation! but it shouldn't be!")
		}
	}
}

func TestTokenizeOperator(t *testing.T) {
	conf := &Config{SrcFileName: "test/test-operators.doc"}
	compt := NewCompterpreter(conf)

	err := compt.LoadSourceCode()
	if err != nil {
		t.Error(err)
	}

	compt.Advance()
	// advance ptr to first character
	for _, op := range []string{"‡", "*", "+", "%", "†"} {
		compt.CurrentToken = Token{}
		compt.TokenizeOperator(compt.CurrentChar)
		if string(compt.CurrentChar) == "EOF" {
			break
		}
		if compt.CurrentToken.Value != op {
			t.Error("incorrect token")
		}
	}
}

func TestLex(t *testing.T) {
	conf := &Config{SrcFileName: "test/test_tokenize.doc"}
	compt := NewCompterpreter(conf)

	err := compt.LoadSourceCode()
	if err != nil {
		t.Error(err)
	}

	err = compt.Lex()
	if err != nil {
		t.Error(err)
	}

	expectedTokens := []string{
		"\n", "123", "†", "3", "*", "2", "‡", "45787894357893", "\n", "0", "+", "00", "+", "1", "\n",
	}
	for idx, token := range compt.Tokens {
		if token.Value != expectedTokens[idx] {
			t.Error("Incorrect token! Try harder")
		}
	}
}