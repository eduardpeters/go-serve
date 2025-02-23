# go-serve
The server project in [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)

## Specs

We have been tasked with creating a web server where users can track how many games players have won:

- `GET /players/{name}` should return a number indicating the total number of wins

- `POST /players/{name}` should record a win for that name, incrementing for every subsequent POST

- `GET /league` should return a list of all players in JSON format:
```json
[
   {
      "Name":"Bill",
      "Wins":10
   },
   {
      "Name":"Alice",
      "Wins":15
   }
]
```
