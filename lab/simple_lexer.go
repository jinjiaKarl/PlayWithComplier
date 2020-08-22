package main

import (
	"bytes"
	"fmt"
	"strings"
)

//定义状态
type DfaState int
const(
	Initial = iota
	Identifier
	Int
	Identifier_int1
	Identifier_int2
	Identifier_int3
	GT // >
	GE // >=
	Assignment
	IntLiteral  //数字字面量

	Plus
	Minus
	Star
	Slash

	SemiColon
	LeftParen
	RightParen
)


type Token struct {
	TokenType string
	TokenText string
}

func (t Token)String() string {
	return fmt.Sprintf("[TokenType: %s,  TokenText: %s]\n", t.TokenType, t.TokenText)
}
var TokenText string //临时保存token文本
var Tokens []Token //保存解析出来的token
var NowToken Token //正在解析的token

func IsAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' //'a' 97 'A'63
}
func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
func IsBlank(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func InitToken(ch byte) DfaState {
	if len(TokenText) > 0{
		NowToken.TokenText = TokenText
		Tokens = append(Tokens, NowToken)

		NowToken = Token{}
		TokenText = ""
	}
	var NewState DfaState
	if IsAlpha(ch){
		if ch == 'i'{
			NewState = Identifier_int1
		}else{
			NewState = Identifier
		}
		NowToken.TokenType = "Identifier"
		TokenText += string(ch)
	}else if IsDigit(ch){
		NewState = IntLiteral
		NowToken.TokenType = "IntLiteral"
		TokenText += string(ch)
	}else if ch == '>'{
		NewState = GT
		NowToken.TokenType = "GT"
		TokenText += string(ch)
	}else if ch == '='{
		NewState = Assignment
		NowToken.TokenType = "Assignment"
		TokenText += string(ch)
	}else if IsBlank(ch){
		NewState = Initial
	}else if ch == ';'{
		NewState = SemiColon
		NowToken.TokenType = "SemiColon"
		TokenText += string(ch)
	}else if ch == '('{
		NewState = LeftParen
		NowToken.TokenType = "LeftParen"
		TokenText += string(ch)
	}else if ch == ')'{
		NewState = RightParen
		NowToken.TokenType = "RightParen"
		TokenText += string(ch)
	}else if ch == '+'{
		NewState = Plus
		NowToken.TokenType = "Plus"
		TokenText += string(ch)
	}else if ch == '-'{
		NewState = Minus
		NowToken.TokenType = "Minus"
		TokenText += string(ch)
	}else if ch == '*'{
		NewState = Star
		NowToken.TokenType = "Star"
		TokenText += string(ch)
	}else if ch == '/'{
		NewState = Slash
		NowToken.TokenType = "Slash"
		TokenText += string(ch)
	}else{
		NewState = Initial  //跳过其他所有的情况
	}
	return NewState
}
//解析字符串，在不同的状态中迁移
func Tokenize(code string) {
	//将string放在一个缓冲区中
	buf := bytes.NewBufferString(code)
	var state DfaState
	state = Initial //初始化状态

	for {
		ch, err := buf.ReadByte()
		if err != nil {
			break
		}
		switch state {
		case Initial:
			state = InitToken(ch) // 确定后续状态
		case Identifier:
			if IsAlpha(ch) || IsDigit(ch) {
				TokenText += string(ch) //保持标识符状态
			} else {
				state = InitToken(ch) //退出标识符状态，并保存Token
			}
		case IntLiteral:
			if IsDigit(ch) {
				TokenText += string(ch)
			} else {
				state = InitToken(ch)
			}
		case GT:
			if ch == '=' {
				NowToken.TokenType = "GE"
				TokenText += string(ch)
				state = GE
			} else {
				state = InitToken(ch) //退出GT状态，并保存Token
			}
		case GE:
			state = InitToken(ch) //退出GE状态，并保存Token
		case Identifier_int1:
			if ch == 'n' {
				state = Identifier_int2
				TokenText += string(ch)
			} else if IsAlpha(ch) || IsDigit(ch) {
				state = Identifier //切换会Identifier
				TokenText += string(TokenText)
			} else {
				state = InitToken(ch)
			}
		case Identifier_int2:
			if ch == 't' {
				state = Identifier_int3
				TokenText += string(ch)
			} else if IsAlpha(ch) || IsDigit(ch) {
				state = Identifier //切换会Identifier
				TokenText += string(ch)
			} else {
				state = InitToken(ch)
			}
		case Identifier_int3:
			if IsBlank(ch) {
				NowToken.TokenType = "Int"
				state = InitToken(ch)
			} else {
				state = Identifier //切换会Identifier
				TokenText += string(ch)
			}
		case Assignment:
			state = InitToken(ch)
		case SemiColon:
			state = InitToken(ch)
		case LeftParen:
			state = InitToken(ch)
		case RightParen:
			state = InitToken(ch)
		case Plus:
			state = InitToken(ch)
		case Minus:
			state = InitToken(ch)
		case Star:
			state = InitToken(ch)
		case Slash:
			state = InitToken(ch)
		default:
		}
	}
	//把最后一个token加进来
	if len(TokenText) > 0 {
		InitToken(' ')
	}
}
func main()  {
	//遇到空格就要init
	script := "age >= 45"
	fmt.Println("parse: " + script)
	Tokenize(script)
	for i := 0; i < len(Tokens); i++{
		fmt.Println(Tokens[i])
	}
	fmt.Println(strings.Repeat("-",100))

	Tokens = Tokens[:0] //清空切片
	script = "int age = 45;"
	fmt.Println("parse: " + script)
	Tokenize(script)
	for i := 0; i < len(Tokens); i++{
		fmt.Println(Tokens[i])
	}
	fmt.Println(strings.Repeat("-",100))

	Tokens = Tokens[:0]
	script = "2 * 3 = 6;"
	fmt.Println("parse: " + script)
	Tokenize(script)
	for i := 0; i < len(Tokens); i++{
		fmt.Println(Tokens[i])
	}
}
