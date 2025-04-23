package common

func Pa(a, b *[]int) {
	*a = append([]int{(*b)[0]}, *a...)
	*b = (*b)[1:]
}

func Pb(a, b *[]int) {
	*b = append([]int{(*a)[0]}, *b...)
	*a = (*a)[1:]
}

func Sa(a []int) {
	a[0], a[1] = a[1], a[0]
}

func Ra(a []int) {
	for i := 0; i < len(a)-1; i++ {
		a[i], a[i+1] = a[i+1], a[i]
	}
}

func Rra(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
}

/******************** Ussed only for checker *******************/

func Sb(b []int) {
	b[0], b[1] = b[1], b[0]
}

func Ss(a, b []int) {
	Sa(a)
	Sb(b)
}

func Rb(b []int) {
	for i := 0; i < len(b)-1; i++ {
		b[i], b[i+1] = b[i+1], b[i]
	}
}

func Rrb(b []int) {
	for i := len(b) - 1; i > 0; i-- {
		b[i], b[i-1] = b[i-1], b[i]
	}
}

func Rrr(a, b []int) {
	Rra(a)
	Rrb(b)
}

func Rr(a, b []int) {
	Ra(a)
	Rb(b)
}
