module start/hello

go 1.21.6

require (
	rsc.io/quote v1.5.2
	start/greetings v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace start/greetings => ../greetings
