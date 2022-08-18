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
result := howlongtobeat.Search("doom")
fmt.Println(result)
```

* Search response example:

```go
[
   {
      "image":"https://howlongtobeat.com/games/79790_Stray_(2021).jpg",
      "title":"Stray",
      "main":"5",
      "extra":"6",
      "completionist":"9"
   }{
      "image":"https://howlongtobeat.com/games/103825_Xenoblade_Chronicles_3.jpg",
      "title":"Xenoblade Chronicles 3",
      "main":"56½",
      "extra":"82",
      "completionist":"161"
   }
]
```

### Missing features
    Single-Player
    Solo
    Co-Op
    Vs.


## License

DO WHAT THE FUCK YOU WANT