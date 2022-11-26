module github.com/Bitspark/go-vyze

go 1.19

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20221120182715-47415e33c366
	github.com/gorilla/websocket v1.5.0
	github.com/iancoleman/strcase v0.2.0
	github.com/jxskiss/base62 v1.1.0
	github.com/xeipuuv/gojsonschema v1.2.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
)

replace github.com/Bitspark/go-vyze/core => ./core

replace github.com/Bitspark/go-vyze/lang => ./lang

replace github.com/Bitspark/go-vyze/lang/parser => ./lang/parser

replace github.com/Bitspark/go-vyze/service => ./service

replace github.com/Bitspark/go-vyze/state => ./state

replace github.com/Bitspark/go-vyze/state/parser => ./state/parser

replace github.com/Bitspark/go-vyze/system => ./system

replace github.com/Bitspark/go-vyze/util => ./util
