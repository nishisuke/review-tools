// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func dupSubExpr(m dsl.Matcher) {
	m.Match(`$x || $x`,
		`$x && $x`,
		`$x | $x`,
		`$x & $x`).
		Where(m["x"].Pure).
		Report(`suspicious identical LHS and RHS`)
}

func boolExprSimplify(m dsl.Matcher) {
	m.Match(`!($x != $y)`).Suggest(`$x == $y`)
	m.Match(`!($x == $y)`).Suggest(`$x != $y`)
}

func exposedMutex(m dsl.Matcher) {
	isExported := func(v dsl.Var) bool {
		return v.Text.Matches(`^\p{Lu}`)
	}

	m.Match(`type $name struct { $*_; sync.Mutex; $*_ }`).
		Where(isExported(m["name"])).
		Report("do not embed sync.Mutex")

	m.Match(`type $name struct { $*_; sync.RWMutex; $*_ }`).
		Where(isExported(m["name"])).
		Report("don not embed sync.RWMutex")
}
