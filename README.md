# Howlongtobeat API

## About & Credits

[How long to beat](https://howlongtobeat.com/) provides information and data about games and how long it will take to finish them.

This library is a simple wrapper api to fetch data from [How long to beat](https://howlongtobeat.com/) (search and detail).
It is an awesome website and a great service, also heavily living from community data. Please check the website and [support](https://howlongtobeat.com/donate.php) if you like what they are doing.

## Usage

### Install the dependency

```go
go get github.com/saturnavt/howlongtobeat
```
or
```go
go get github.com/saturnavt/howlongtobeat@latest
```

### Use in code

#### Add imports


```go
import "github.com/saturnavt/howlongtobeat"
```


#### Searching for a game

```go
result := howlongtobeat.Search("Elden Ring")
fmt.Println(result)
```

* Search response example:

```go
[
   {
      "image":"https://howlongtobeat.com/games/68151_Elden_Ring.jpg",
      "title":"Elden Ring",
      "main":"52",
      "extra":"98",
      "completionist":"131"
   }{
      "image":"https://howlongtobeat.com/games/108888_Elden_Ring_GB.jpg",
      "title":"Elden Ring GB",
      "main":"21",
      "extra":"29",
      "completionist":"--"
   }
]
```

### Missing features
    Single-Player
    Solo
    Co-Op
    Vs.

### Why missing features
    To get the hours to complete a online game is really hard becouse it depends on how long you will play it or how many updates will the game have in the future. Not even in howlongtobeat page is accurated for example take a look at (Valorant) it says 23½ Hours ive play it and it take way more than that, there are many online games that shows the wrong time.

## License

DO WHAT THE FUCK YOU WANT