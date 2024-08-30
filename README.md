# PokeAPI Simplificado

Este programa em Go busca dados dos primeiros N Pokémon da [PokeAPI](https://pokeapi.co) e os salva em um arquivo JSON.

## Funcionalidades

- Busca dados dos Pokémon, incluindo nome, tipos, stats e imagem.
- Salva os dados buscados em um arquivo JSON.

## Dependências

- Go 1.16 ou superior

## Instalação

1. Clone o repositório:
    ```sh
    git clone https://github.com/wfabi0/pokeapi-simplify.git
    ```

2. Instale as dependências:
    ```sh
    go mod tidy
    ```

## Uso

1. Execute o programa:
    ```sh
    go run main.go
    ```

2. O programa buscará dados dos primeiros 50 Pokémon e os salvará em `pokemons.json`.

## Exemplo

Exemplo do output JSON:
```json
[
  {
    "name": "bulbasaur",
    "types": ["grass", "poison"],
    "stats": [45, 49, 49, 65, 65, 45],
    "image": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png"
  },
]