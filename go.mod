module github.com/cdvelop/timeserver

go 1.20

require (
	github.com/cdvelop/strings v0.0.7
	github.com/cdvelop/timetools v0.0.26
)

require github.com/cdvelop/model v0.0.77 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timetools => ../timetools
