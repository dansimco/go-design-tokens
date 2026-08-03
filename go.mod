module github.com/dansimco/go-design-tokens

go 1.25.0

// The v1.0.x tags were published in error and have been deleted from the
// repository, but remain cached in the module proxy. Without these
// retractions `go get` resolves to v1.0.4 and then fails to download it,
// because the underlying revision no longer exists. Development continues
// on v0.x; see v0.3.0.
retract (
	v1.0.0 // Tagged in error; tag deleted, content unfetchable.
	v1.0.1 // Tagged in error; tag deleted, content unfetchable.
	v1.0.2 // Tagged in error; tag deleted, content unfetchable.
	v1.0.3 // Tagged in error; tag deleted, content unfetchable.
	v1.0.4 // Tagged in error; tag deleted, content unfetchable.
	v1.0.5 // Contains only these retractions.
)
