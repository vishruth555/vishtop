# vishtop

vishtop is a CLI tool that provides a summarized graph UI of your system performance directly in your terminal.

## Features
- Real-time, terminal-based graphical display of CPU and memory usage
- Customizable refresh rate for performance polling
- Simple and intuitive interface

## Installation

To install vishtop globally, run:

```
go install github.com/vishruth555/vishtop@latest
```


## Usage

Run vishtop in your terminal:

```
vishtop
```

You can specify the refresh rate (polling interval) in milliseconds using the `-poll` argument (default is 500ms):

```
vishtop -poll 1000
```

## Example Output

```
 30.78 ┤                                          ╭╮
 27.94 ┤                                          ││
 25.10 ┤                                          ││
 22.27 ┤                                          ││
 19.43 ┤                                          ││
 16.59 ┤                                         ╭╯│
 13.75 ┤                              ╭╮         │ │              ╭╮
 10.92 ┤                           ╭──╯│      ╭╮ │ │        ╭╮    ││               ╭╮           ╭╮
  8.08 ┤                          ╭╯   ╰──╮   ││ │ │ ╭─────╮││   ╭╯│               ││           │╰╮
  5.24 ┤   ╭╮        ╭╮          ╭╯       ╰╮╭╮││╭╯ ╰─╯     ╰╯╰───╯ ╰──────╮╭──────╮│╰─╮╭──╮╭────╯ ╰─╮ ╭╮
  2.41 ┼───╯╰────────╯╰──────────╯         ╰╯╰╯╰╯                         ╰╯      ╰╯  ╰╯  ╰╯        ╰─╯╰──
                                                   CPU Usage (%)
 60.81 ┤                                          ╭─╮
 60.70 ┤                                          │ ╰──╮
 60.58 ┤                                          │    ╰─╮
 60.47 ┤                                          │      ╰──╮    ╭╮
 60.36 ┤                                          │         ╰────╯│
 60.25 ┤                                          │               │
 60.13 ┤                                         ╭╯               ╰╮
 60.02 ┤                                         │                 ╰───╮                          ╭╮╭╮
 59.91 ┤                             ╭──╮ ╭╮     │                     ╰────╮                   ╭─╯╰╯╰────
 59.80 ┤             ╭───────────────╯  ╰─╯╰─────╯                          ╰──╮   ╭────────────╯
 59.69 ┼─────────────╯                                                         ╰───╯
                                                 Memory Usage (%)
 622 ┤                            ╭╮          ╭╮                      ╭╮
 560 ┤                            ││          ││                      ││
 497 ┤             ╭╮             ││          ││                      ││
 435 ┤             ││             ││          ││       ╭╮             ││
 373 ┤             ││             ││          ││       ││             ││           ╭╮         ╭╮
 311 ┤             ││             ││          ││       ││             ││          ╭╯│         ││
 249 ┤      ╭╮     ││             ││          ││       ││             ││          │ │         ││
 187 ┤      ││     ││             ││          ││       ││             ││          │ │         ││
 124 ┤      │╰╮    ││             ││          ││       ││             ││          │ │         ││
  62 ┤   ╭╮ │ │    ││        ╭╮   ││   ╭╮     ││  ╭╮   ││   ╭╮        ││╭╮      ╭╮│ │     ╭╮  ││    ╭╮
   0 ┼───╯╰─╯ ╰────╯╰────────╯╰───╯╰───╯╰─────╯╰──╯╰───╯╰───╯╰────────╯╰╯╰──────╯╰╯ ╰─────╯╰──╯╰────╯╰──
```

## License

MIT License
