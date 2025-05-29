package iterext

import "iter"

func Map[From any, To any](it iter.Seq[From], mapFnc func(From) To) iter.Seq[To] {
	return func(yield func(To) bool) {
		for v := range it {
			if !yield(mapFnc(v)) {
				return
			}
		}
	}
}
