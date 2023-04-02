# Wordle game for terminal

Who said you can't have fun on a terminal?

Let's play [WORDLE](https://www.nytimes.com/games/wordle/index.html) then! 

## Screenshots


## Fork it and have fun
The repo comes with a static word to be guessed but you can implement your own source! Just implement the **Storage** interface and you are done!

```go
type Storage interface {
	GetWord() string
}

```

You wanna render it differently or you wanna create different implementaions for different terminals? Implement your own **Renderer**

```go
type Renderer interface {
	Render(*game.Game)
	Print(string)
}
```