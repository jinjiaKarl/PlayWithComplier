package main

import (
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
type TokenType int


type Token struct {
	TokenType string
	TokenText string
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
func Tokenize(code string)  {
	var state DfaState
	state = Initial //初始化状态

	for i := 0; i < len(code); i++{
		switch state {
		case Initial:
			state = InitToken(code[i]) // 确定后续状态
		case Identifier:
			if IsAlpha(code[i]) || IsDigit(code[i]){
				TokenText += string(code[i])  //保持标识符状态
			}else{
				state = InitToken(code[i]) //退出标识符状态，并保存Token
			}
		case IntLiteral:
			if IsDigit(code[i]){
				TokenText += string(code[i])
			}else{
				state = InitToken(code[i])
			}
		case GT:
			if code[i] == '='{
				NowToken.TokenType = "GE"
				TokenText += string(code[i])
				state = GE
			}else{
				state = InitToken(code[i]) //退出GT状态，并保存Token
			}
		case GE:
			state = InitToken(code[i]) //退出GE状态，并保存Token
		case Identifier_int1:
			if code[i] == 'n'{
				state = Identifier_int2
				TokenText += string(code[i])
			}else if IsAlpha(code[i]) || IsDigit(code[i]){
				state = Identifier //切换会Identifier
				TokenText += string(TokenText)
			}else{
				state = InitToken(code[i])
			}
		case Identifier_int2:
			if code[i] == 't'{
				state = Identifier_int3
				TokenText += string(code[i])
			}else if IsAlpha(code[i]) || IsDigit(code[i]){
				state = Identifier //切换会Identifier
				TokenText += string(code[i])
			}else{
				state = InitToken(code[i])
			}
		case Identifier_int3:
			if IsBlank(code[i]){
				NowToken.TokenType = "Int"
				state = InitToken(code[i])
			}else{
				state = Identifier //切换会Identifier
				TokenText += string(code[i])
			}
		case Assignment:
			state = InitToken(code[i])
		case SemiColon:
			state = InitToken(code[i])
		case LeftParen:
			state = InitToken(code[i])
		case RightParen:
			state = InitToken(code[i])
		case Plus:
			state = InitToken(code[i])
		case Minus:
			state = InitToken(code[i])
		case Star:
			state = InitToken(code[i])
		case Slash:
			state = InitToken(code[i])
		default:
			
		}
	}
	//把最后一个token加进来
	if len(TokenText) > 0{
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
	fmt.Println(strings.Repeat("-",20))

	Tokens = Tokens[:0] //清空切片
	script = "int age = 45;"
	fmt.Println("parse: " + script)
	Tokenize(script)
	for i := 0; i < len(Tokens); i++{
		fmt.Println(Tokens[i])
	}
	fmt.Println(strings.Repeat("-",20))

	Tokens = Tokens[:0]
	script = "2 * 3 = 6;"
	fmt.Println("parse: " + script)
	Tokenize(script)
	for i := 0; i < len(Tokens); i++{
		fmt.Println(Tokens[i])
	}
}
