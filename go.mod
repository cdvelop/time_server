module github.com/cdvelop/timeserver

go 1.20

require (
	github.com/cdvelop/strings v0.0.8
	github.com/cdvelop/timetools v0.0.28
)

require github.com/cdvelop/model v0.0.78 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timetools => ../timetools
