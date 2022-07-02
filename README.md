# go-wordwrap

Because sometimes the formatting of the source text is good, but you also need
it to look good when printed.

## Install

```shell
go get github.com/p-se/go-wordwrap
```

## Examples

```go
package main

import (
	"fmt"

	"github.com/p-se/go-wordwrap"
)

func main() {
	s := `
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

	fmt.Printf("Original:\n\n%s\n\n", s)
	fmt.Printf("Prepared:\n\n%s\n\n", wordwrap.Prepare(s))
	fmt.Printf("Wrapped (60 runes width):\n\n%s\n\n", wordwrap.WrapText(s, 60))
	fmt.Printf("Wrapped and indented (60 runes width):\n\n%s\n\n", wordwrap.Wrap(s, 60, "    "))
	fmt.Printf("Wrapped and indented (40 runes width):\n\n%s\n\n", wordwrap.Wrap(s, 40, "    "))
}
```

```
Original:


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
	

Prepared:

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Interdum posuere lorem ipsum dolor sit amet consectetur adipiscing elit. Massa tincidunt dui ut ornare lectus sit amet est.

Aliquet enim tortor at auctor urna nunc id cursus. Neque vitae tempus quam pellentesque nec nam. Malesuada fames ac turpis egestas integer eget aliquet nibh. Facilisi etiam dignissim diam quis. Congue quisque egestas diam in arcu. Turpis massa sed elementum tempus egestas sed sed risus. Orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor. Nibh nisl condimentum id venenatis a condimentum.

Quis viverra nibh cras pulvinar mattis nunc sed blandit. Aliquet risus feugiat in ante. Lacus sed viverra tellus in hac habitasse platea dictumst vestibulum. A diam maecenas sed enim ut sem. Sed cras ornare arcu dui vivamus arcu felis bibendum. Imperdiet massa tincidunt nunc pulvinar sapien.

Wrapped (60 runes width):

Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna
aliqua. Interdum posuere lorem ipsum dolor sit amet
consectetur adipiscing elit. Massa tincidunt dui ut ornare
lectus sit amet est.

Aliquet enim tortor at auctor urna nunc id cursus. Neque
vitae tempus quam pellentesque nec nam. Malesuada fames ac
turpis egestas integer eget aliquet nibh. Facilisi etiam
dignissim diam quis. Congue quisque egestas diam in arcu.
Turpis massa sed elementum tempus egestas sed sed risus.
Orci phasellus egestas tellus rutrum tellus pellentesque eu
tincidunt tortor. Nibh nisl condimentum id venenatis a
condimentum.

Quis viverra nibh cras pulvinar mattis nunc sed blandit.
Aliquet risus feugiat in ante. Lacus sed viverra tellus in
hac habitasse platea dictumst vestibulum. A diam maecenas
sed enim ut sem. Sed cras ornare arcu dui vivamus arcu
felis bibendum. Imperdiet massa tincidunt nunc pulvinar
sapien.

Wrapped and indented (60 runes width):

    Lorem ipsum dolor sit amet, consectetur adipiscing
    elit, sed do eiusmod tempor incididunt ut labore et
    dolore magna aliqua. Interdum posuere lorem ipsum dolor
    sit amet consectetur adipiscing elit. Massa tincidunt
    dui ut ornare lectus sit amet est.

    Aliquet enim tortor at auctor urna nunc id cursus.
    Neque vitae tempus quam pellentesque nec nam. Malesuada
    fames ac turpis egestas integer eget aliquet nibh.
    Facilisi etiam dignissim diam quis. Congue quisque
    egestas diam in arcu. Turpis massa sed elementum tempus
    egestas sed sed risus. Orci phasellus egestas tellus
    rutrum tellus pellentesque eu tincidunt tortor. Nibh
    nisl condimentum id venenatis a condimentum.

    Quis viverra nibh cras pulvinar mattis nunc sed
    blandit. Aliquet risus feugiat in ante. Lacus sed
    viverra tellus in hac habitasse platea dictumst
    vestibulum. A diam maecenas sed enim ut sem. Sed cras
    ornare arcu dui vivamus arcu felis bibendum. Imperdiet
    massa tincidunt nunc pulvinar sapien.

Wrapped and indented (40 runes width):

    Lorem ipsum dolor sit amet,
    consectetur adipiscing elit, sed do
    eiusmod tempor incididunt ut labore
    et dolore magna aliqua. Interdum
    posuere lorem ipsum dolor sit amet
    consectetur adipiscing elit. Massa
    tincidunt dui ut ornare lectus sit
    amet est.

    Aliquet enim tortor at auctor urna
    nunc id cursus. Neque vitae tempus
    quam pellentesque nec nam.
    Malesuada fames ac turpis egestas
    integer eget aliquet nibh. Facilisi
    etiam dignissim diam quis. Congue
    quisque egestas diam in arcu.
    Turpis massa sed elementum tempus
    egestas sed sed risus. Orci
    phasellus egestas tellus rutrum
    tellus pellentesque eu tincidunt
    tortor. Nibh nisl condimentum id
    venenatis a condimentum.

    Quis viverra nibh cras pulvinar
    mattis nunc sed blandit. Aliquet
    risus feugiat in ante. Lacus sed
    viverra tellus in hac habitasse
    platea dictumst vestibulum. A diam
    maecenas sed enim ut sem. Sed cras
    ornare arcu dui vivamus arcu felis
    bibendum. Imperdiet massa tincidunt
    nunc pulvinar sapien.

```
