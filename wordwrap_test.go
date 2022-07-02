package wordwrap

import (
	"testing"
)

func TestDedent(t *testing.T) {
	const source = `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
		Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit.
		Massa tincidunt dui ut ornare lectus sit amet est.

		Aliquet enim tortor at auctor urna nunc id cursus.
		Neque vitae tempus quam pellentesque nec nam.
		Malesuada fames ac turpis egestas integer eget aliquet nibh.
		Facilisi etiam dignissim diam quis. Congue quisque egestas diam in arcu.
		Turpis massa sed elementum tempus egestas sed sed risus.
		Orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor.
		Nibh nisl condimentum id venenatis a condimentum.

		Quis viverra nibh cras pulvinar mattis nunc sed blandit.
		Aliquet risus feugiat in ante.
		Lacus sed viverra tellus in hac habitasse platea dictumst vestibulum.
		A diam maecenas sed enim ut sem. Sed cras ornare arcu dui vivamus arcu felis bibendum.
		Imperdiet massa tincidunt nunc pulvinar sapien.
`
	const expected = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit.
Massa tincidunt dui ut ornare lectus sit amet est.

Aliquet enim tortor at auctor urna nunc id cursus.
Neque vitae tempus quam pellentesque nec nam.
Malesuada fames ac turpis egestas integer eget aliquet nibh.
Facilisi etiam dignissim diam quis. Congue quisque egestas diam in arcu.
Turpis massa sed elementum tempus egestas sed sed risus.
Orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor.
Nibh nisl condimentum id venenatis a condimentum.

Quis viverra nibh cras pulvinar mattis nunc sed blandit.
Aliquet risus feugiat in ante.
Lacus sed viverra tellus in hac habitasse platea dictumst vestibulum.
A diam maecenas sed enim ut sem. Sed cras ornare arcu dui vivamus arcu felis bibendum.
Imperdiet massa tincidunt nunc pulvinar sapien.
`

	if dedented := Dedent(source); dedented != expected {
		t.Fail()
	}
}

func TestJoin(t *testing.T) {
	const source = `Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit.
Massa tincidunt dui ut ornare lectus sit amet est.`
	const expected = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit. Massa tincidunt dui ut ornare lectus sit amet est.`

	if joined := Join(Dedent(source)); joined != expected {
		t.Fail()
	}
}

func TestPrepare(t *testing.T) {
	const source = `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
		Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit.
		Massa tincidunt dui ut ornare lectus sit amet est.
	`
	const expected = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit. Massa tincidunt dui ut ornare lectus sit amet est.`

	if prepared := Prepare(source); prepared != expected {
		t.Fail()
	}
}

func TestWrapParagraph(t *testing.T) {
	const source = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit. Massa tincidunt dui ut ornare lectus sit amet est.`
	expected := `Lorem ipsum dolor sit amet,
consectetur adipiscing elit,
sed do eiusmod tempor
incididunt ut labore et
dolore magna aliqua. Interdum
posuere lorem ipsum dolor sit
amet consectetur adipiscing
elit. Massa tincidunt dui ut
ornare lectus sit amet est.`
	if WordWrapParagraph(source, 30) != expected {
		t.Fail()
	}
}

func TestWrapText(t *testing.T) {
	const source = `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
		Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit.
		Massa tincidunt dui ut ornare lectus sit amet est.

		Aliquet enim tortor at auctor urna nunc id cursus.
		Neque vitae tempus quam pellentesque nec nam.
		Malesuada fames ac turpis egestas integer eget aliquet nibh.
		Facilisi etiam dignissim diam quis. Congue quisque egestas diam in arcu.
		Turpis massa sed elementum tempus egestas sed sed risus.
		Orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor.
		Nibh nisl condimentum id venenatis a condimentum.

		Quis viverra nibh cras pulvinar mattis nunc sed blandit.
		Aliquet risus feugiat in ante.
		Lacus sed viverra tellus in hac habitasse platea dictumst vestibulum.
		A diam maecenas sed enim ut sem. Sed cras ornare arcu dui vivamus arcu felis bibendum.
		Imperdiet massa tincidunt nunc pulvinar sapien.
	`
	expected := `Lorem ipsum dolor sit amet,
consectetur adipiscing elit,
sed do eiusmod tempor
incididunt ut labore et
dolore magna aliqua. Interdum
posuere lorem ipsum dolor sit
amet consectetur adipiscing
elit. Massa tincidunt dui ut
ornare lectus sit amet est.

Aliquet enim tortor at auctor
urna nunc id cursus. Neque
vitae tempus quam
pellentesque nec nam.
Malesuada fames ac turpis
egestas integer eget aliquet
nibh. Facilisi etiam
dignissim diam quis. Congue
quisque egestas diam in arcu.
Turpis massa sed elementum
tempus egestas sed sed risus.
Orci phasellus egestas tellus
rutrum tellus pellentesque eu
tincidunt tortor. Nibh nisl
condimentum id venenatis a
condimentum.

Quis viverra nibh cras
pulvinar mattis nunc sed
blandit. Aliquet risus
feugiat in ante. Lacus sed
viverra tellus in hac
habitasse platea dictumst
vestibulum. A diam maecenas
sed enim ut sem. Sed cras
ornare arcu dui vivamus arcu
felis bibendum. Imperdiet
massa tincidunt nunc pulvinar
sapien.`
	if WrapText(source, 30) != expected {
		t.Fail()
	}
}

func TestWrap(t *testing.T) {
	const source = `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
		Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit.
		Massa tincidunt dui ut ornare lectus sit amet est.

		Aliquet enim tortor at auctor urna nunc id cursus.
		Neque vitae tempus quam pellentesque nec nam.
		Malesuada fames ac turpis egestas integer eget aliquet nibh.
		Facilisi etiam dignissim diam quis. Congue quisque egestas diam in arcu.
		Turpis massa sed elementum tempus egestas sed sed risus.
		Orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor.
		Nibh nisl condimentum id venenatis a condimentum.

		Quis viverra nibh cras pulvinar mattis nunc sed blandit.
		Aliquet risus feugiat in ante.
		Lacus sed viverra tellus in hac habitasse platea dictumst vestibulum.
		A diam maecenas sed enim ut sem. Sed cras ornare arcu dui vivamus arcu felis bibendum.
		Imperdiet massa tincidunt nunc pulvinar sapien.
`
	const expected = `	Lorem ipsum dolor sit amet, consectetur adipiscing elit,
	sed do eiusmod tempor incididunt ut labore et dolore magna
	aliqua. Interdum posuere lorem ipsum dolor sit amet
	consectetur adipiscing elit. Massa tincidunt dui ut ornare
	lectus sit amet est.

	Aliquet enim tortor at auctor urna nunc id cursus. Neque
	vitae tempus quam pellentesque nec nam. Malesuada fames ac
	turpis egestas integer eget aliquet nibh. Facilisi etiam
	dignissim diam quis. Congue quisque egestas diam in arcu.
	Turpis massa sed elementum tempus egestas sed sed risus.
	Orci phasellus egestas tellus rutrum tellus pellentesque
	eu tincidunt tortor. Nibh nisl condimentum id venenatis a
	condimentum.

	Quis viverra nibh cras pulvinar mattis nunc sed blandit.
	Aliquet risus feugiat in ante. Lacus sed viverra tellus in
	hac habitasse platea dictumst vestibulum. A diam maecenas
	sed enim ut sem. Sed cras ornare arcu dui vivamus arcu
	felis bibendum. Imperdiet massa tincidunt nunc pulvinar
	sapien.`

	if Wrap(source, 60, "\t") != expected {
		t.Fail()
	}
}
