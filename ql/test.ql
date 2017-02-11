#!/usr/bin/env qlang

sprintln = fn(args...) {
	for b = bytes.buffer(); true; _ {
		fprintln(b, args...)
		return b.String()
	}
}

sSort = fn(s) {
	for i, n = 0, len(s); i < n; i++ {
		for j = i + 1; j < n; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

test1 = fn() {
	pkglst = meta.pkgs()
	println("go pkgs:")
	println(pkglst)
	println()

	mFnlist = make(map[string]bool)
	for _, k = range meta.fnlist() {
		mFnlist[k] = false
	}

	for _, k = range pkglst {
		delete(mFnlist, k)
	}

	for _, k = range meta.dir(math) {
		delete(mFnlist, k)
	}

	fnlst = make([]string, 0, len(mFnlist))
	for k = range mFnlist {
		fnlst = append(fnlst, k)
	}

	println("builtin fnlist (except math):")
	println(sSort(fnlst))
	println()
}

test2 = fn() {
	mf = fn(v) {
		return fn() { return v }
	}

	a = 1
	f1 = mf(a)
	defer println("expect '1', f1():", f1())

	a = 2
	print("a: ", a, ", f1(): ", f1())
	println()

	f2 = fn {
		v = a
		return fn() { return v }
	}
	defer println("expect '2', f2():", f2())

	a = 3
	print("a: ", a, ", f2(): ", f2())
	println()
}

test3 = fn() {
	s = sprintln("Hello, world!")
	print(s)
}

main {
	test1()
	test2()
	test3()
}
