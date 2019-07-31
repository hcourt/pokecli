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
$ pokecli
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

### search - Search for entities
Search takes a type `(-t)` and one or more search strings, and prints a list of
matching entities.

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
$ pokecli search -t pokemon foo leo
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

### show - Show an Entity
Show takes a type `(-t)` and a single name or ID number, and prints a simple
summary of the entity.

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

### effect - Check the Type Effectiveness of a Move
Effect takes a pokemon `(-p)` and a move `(-m)`, and prints a message about the
effectiveness of that move's type on the defending pokemon.  If the move is
non-damaging it will return that information instead.

<details> <summary>Examples</summary>


```console
$ pokecli effect -m rock-slide -p charizard
If a rock move attacks a [flying fire] pokemon, the damage is double super effective.
```

```console
$ pokecli effect -m flamethrower -p bulbasaur
If a fire move attacks a [poison grass] pokemon, the damage is super effective.
```

```console
$ pokecli effect -m shadow-ball -p beedrill
  If a ghost move attacks a [poison bug] pokemon, the damage is effective.
```

```console
$ pokecli effect -m body-slam -p steelix
If a normal move attacks a [ground steel] pokemon, the damage is not very effective (50%).
```

```console
$ pokecli effect -m solar-beam -p dialga
If a grass move attacks a [dragon steel] pokemon, the damage is not very effective (25%).
```

```console
$ pokecli effect -m thunder -p geodude
If a electric move attacks a [ground rock] pokemon, the damage is not effective.
```

Non-damaging moves:
```console
$ pokecli effect -m hypnosis -p snorlax
Move is a status move and will not cause typed damage.
```

</details>