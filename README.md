# Wordle game for terminal

Who said you can't have fun on a terminal?

Let's play [WORDLE](https://www.nytimes.com/games/wordle/index.html) then! 

## Screenshots

![Screenshot 2023-03-31 at 12 28 25](https://user-images.githubusercontent.com/5497475/229344196-b1a7bd89-60fe-4a05-a7f6-218df50b802c.png)


![Screenshot 2023-03-31 at 12 30 47](https://user-images.githubusercontent.com/5497475/229344202-93be0052-ed8a-4f56-9dff-631ca9a78f67.png)

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
# Technologies
The project is written in golang 1.20

# Development
To run the app use the command
```
make run
```
to build the binary
```
make build
```
to run the tests
```
make test
```

# Authors
[Vincenzo Trapani](https://github.com/trapvincenzo)
