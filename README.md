# pokecli

Command line interface (CLI) for [PokéAPI](https://pokeapi.co), written in Go using [pokeapi-go](https://github.com/mtslzr/pokeapi-go). *Supports PokeAPI v2.*

## Documentation

Full API documentation can be found at [pokeapi.co](https://pokeapi.co/docs/v2.html).

## Getting Started

Build the `pokecli` binary from source:
```console
$ make
```
Then run:
```console
$ ./bin/pokecli
```

Add it to your path:
```console
$ export PATH=$PATH:$(pwd)/bin
```

## Usage

```console
$ pokecli -h
A simple command line interface wrapper for PokéAPI written by
                hcourt in Go using pokeapi-go.
                Complete documentation is available at https://pokeapi.co

Usage:
  pokecli [command]

Available Commands:
  help        Help about any command
  search      Search for entities by type and name
  show        Show information about an entity

Flags:
  -h, --help      help for pokecli
  -v, --verbose   enable logging

Use "pokecli [command] --help" for more information about a command.

```

## Operations

### Search for entities
Search takes a type `(-t)` and one or more search strings.

<details> <summary>Examples</summary>

```console
$ pokecli search -t pokemon saur
bulbasaur
ivysaur
venusaur
venusaur-mega
```

```console
$ pokecli search -t move thunder
thunder-punch
thunder-shock
thunderbolt
thunder-wave
thunder
thunder-fang
10-000-000-volt-thunderbolt
```

With multiple search strings:

```console
pokecli search -t pokemon foo leo
foongus
mienfoo
charmeleon
kecleon
sealeo
empoleon
litleo
solgaleo
```

</details>

### Show
Show takes a type `(-t)` and a single name or ID number.

<details> <summary>Examples</summary>

```console
$ pokecli show -t pokemon bulbasaur
bulbasaur (#1) [poison grass]
```

```console
$ pokecli show -t move flamethrower
flamethrower (class: special, type: fire, power: 90, accuracy: 100)
```

Using an ID number:
```console
$ pokecli show -t pokemon 100
voltorb (#100) [electric]
```
</details>
