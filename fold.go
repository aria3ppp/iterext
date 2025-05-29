package iterext

import "iter"

func Fold[Item any, F any](it iter.Seq[Item], f F, fnc func(F, Item) F) F {
	for nv := range it {
		f = fnc(f, nv)
	}
	return f
}
