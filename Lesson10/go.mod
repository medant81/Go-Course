module Lesson10

go 1.22.1

replace (
	github.com/medant81/hello_old v1.1.1 => github.com/medant81/hello v1.0.0
)

require (
	github.com/medant81/hello v1.1.0 // indirect
	github.com/medant81/hello_new v1.1.1 // indirect
	github.com/medant81/hello_old v1.1.1 // indirect
)
