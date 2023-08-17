package main

import (
	"bytes"
	"fmt"
	// "strings"
)

func isValid(s string) bool {
	b := false
	i := 0
	p := 0
	type bracket struct {
		x int
		y byte
	}
	bro := bracket{0, '0'}
	sis := bracket{0, '0'}
	// 1. Get string
	// fmt.Println("Input:", s)
	// 2. Is string bigger than 0
	// 2.1. If false -> kill funcion, return bool
	for {
		if len(s) == 0 {
			return b
		}
		// 2.2. If true -> continue
		// 3. Search for bracket
		i = 0
		g1 := func() bool {
			b = false
			// fmt.Printf("%T", s)
			for ; i < len(s); i++ {
				p = i
				b = bytes.ContainsAny([]byte(s), "()[]{}")
				if b == true {
					return b
				}
			}
			return b
		}
		b = g1()
		if b == false {
			return b
		}
		// 4. Is bracket opener -> True / False
		// 4.1. If true -> Save bracket 1 and position
		o1 := func() bool {
			b = false
			switch {
			case s[p] == '(':
				bro = bracket{p, '('}
				b = true
			case s[p] == '{':
				bro = bracket{p, '{'}
				b = true
			case s[p] == '[':
				bro = bracket{p, '['}
				b = true
			}
			return b
		}
		b = o1()
		// 4.2. If false -> kill function, return false
		if b == false {
			return b
		}
		// 5. Search for bracket after position of bracket 1
		l := 0
		for {
			// fmt.Println("Loop round:", l)
			l = l + 1
			i = p + 1
			b = g1()
			if b == false {
				return b
			}
			// 6. Is bracket opener -> True / False
			b = o1()
			// 6.1. If false -> GoTo Step 7.
			if b == false {
				break
			}
			// 6.2. If true -> overwrite bracket 1 and position -> GoTo Step 5.
		}
		// 7. Is bracket closer to bracket 1
		o2 := func() bool {
			b = false
			switch {
			case s[p] == ')':
				sis = bracket{p, '('}
				b = true
			case s[p] == '}':
				sis = bracket{p, '{'}
				b = true
			case s[p] == ']':
				sis = bracket{p, '['}
				b = true
			}
			return b
		}
		b = o2()
		// 7.1. If true -> Delete both, go back so step 1.
		del := func() {
			brocopy := s
			brocopy = brocopy[:(bro.x)]
			// fmt.Println(brocopy)
			siscopy := s
			siscopy = siscopy[(sis.x + 1):]
			// fmt.Println(siscopy)
			s = brocopy + siscopy
		}
		b = false
		if bro.y == sis.y {
			b = true
			del()
		}
		// 7.2. If false -> kill function, return false
		if bro.y != sis.y {
			b = false
			break
		}
	}
	return b
}

func main() {
	fmt.Println("START")
	fmt.Println("Is Valid?:", isValid("()"))
	fmt.Println("Is Valid?:", isValid("()[]{}"))
	fmt.Println("Is Valid?:", isValid("(])"))
	fmt.Println("Is Valid?:", isValid("([)]"))
	fmt.Println("Is Valid?:", isValid("{[]}"))
	fmt.Println("STOP")
}
