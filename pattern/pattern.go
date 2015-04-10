// patterns is a package that provides a way to create
// musical scores similar to the SuperCollider patterns API
// http://doc.sccode.org/Tutorials/Streams-Patterns-Events1.html.
// Each pattern type implements the Pattern interface, which means that
// it can transform something that emits uint64 ticks to something
// that emits interface{} values.
package pattern

type Ticks chan uint64

type Values chan interface{}

type Pattern interface {
	Stream(ticks Ticks) Values
}
