# PokedexCLI

A command-line interface (CLI) tool for fetching and exploring Pokémon data from the [PokéAPI](https://pokeapi.co/).

## Features
- Fetch details of a Pokémon by name or ID
- List Pokémon locations
- Explore the Pokémon world
- Inspect Pokémon details
- Caching mechanism for optimized performance

## Learning Goals
- Learn how to parse JSON in Go
- Practice making HTTP requests in Go
- Learn how to build a CLI tool that makes interacting with a back-end server easier
- Get hands-on practice with local Go development and tooling
- Learn about caching and how to use it to improve performance

## Installation
### Prerequisites
- Go (1.18 or later)

### Clone the repository
```sh
git clone https://github.com/sabrek15/pokedexcli.git
cd pokedexcli
```
### Build the Project
```sh
go build -o pokedexcli
```

## Usage
Run the CLI tool:
```sh
./pokedexcli
```

### Available Commands
| Command   | Description |
|-----------|-------------|
| `catch`   | Catch a Pokémon |
| `exit`    | Exit the CLI |
| `explore` | Explore Pokémon world |
| `help`    | Display help information |
| `inspect` | Inspect a Pokémon |
| `map`     | View the map |
| `pokedex` | View caught Pokémon |

## Project Structure
```sh
pokedexcli/
│── internal/
│   ├── pokeapi/
│   │   ├── client.go
│   │   ├── location_get.go
│   │   ├── location_list.go
│   │   ├── pokeapi.go
│   │   ├── pokemon_get.go
│   │   ├── type_locations.go
│   │   ├── type_pokemon.go
│   ├── pokecache/
│   │   ├── pokecache.go
│   │   ├── pokecache_test.go
│── .gitignore
│── command_catch.go
│── command_exit.go
│── command_explore.go
│── command_help.go
│── command_inspect.go
│── command_map.go
│── command_pokedex.go
│── go.mod
│── main.go
│── main.sh
│── README.md
│── repl.go
│── repl_test.go
│── repl.log
│── test.sh
```

## REPL(Read-Eval-Print Loop)
REPL stands for "read-eval-print loop." It's a common type of program that allows for interactivity. You type in a command, and the program evaluates it and prints the result. You can then type in another command, and so on.

## Testing
Run tests using:
```sh
./test.sh
```

## Contributing
Feel free to submit pull request and issues to improve this project.