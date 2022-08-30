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
      "main":"52 Hours",
      "extra":"98 Hours",
      "completionist":"131 Hours"
   }{
      "image":"https://howlongtobeat.com/games/108888_Elden_Ring_GB.jpg",
      "title":"Elden Ring GB",
      "main":"21 Mins",
      "extra":"29 Mins",
      "completionist":"--"
   }
]
```

## License

DO WHAT THE FUCK YOU WANT