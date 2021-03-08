# Description
CLI-Snake is a simple command-line implementation of the [snake game](https://en.wikipedia.org/wiki/Snake_(video_game_genre)), written in GO. The game state is saved by writing game data to ```temp/gamestate.json``` in the root directory of the project. The ```gamestate.json``` is deleted in a "Game Over" condition or if the game is manually ended.

# Libraries used
For simplicity, the project used the [urfave](https://github.com/urfave/cli) library to scaffold a golang command line application.

# Requirements
- Golang v.1.16

# Usage
### Build (OSX/Linux)
Navigate to the root directory of the project and run the following command:
```make build```
A binary will be built and output into the ```./bin``` folder

The binary can be run by navigating into the ```./bin``` directory and running the following command:
```./clisnake -h```

### Compile only
Navigate to the root directory of the project and run the following command:
```go run cmd/main.go -h```

# Gameplay
All the commands for playing the game can be viewed by running the built-in CLI help with either of the two commands described in **Usage** above.

### Starting a new game
Navigate to the project's root directory and run the following command to start a new game:
```go run cmd/main.go newgame```
_This wil start a new game with a default board size of **10 rows and 22 columns**._

To start a new game with a custom board size, enter the following command:
```go run cmd/main.go newgame --size 15x30```
_This wil start a new game with a board size of **15 rows and 30 columns**._

### Playing a game
Once a new game has been started, the snake's head will be displayed on the board by a **◯** and food will be displayed by a ◈. The snake's head can then be moved towards the food (one position at a time) by entering the following commands:

- Move snake head up => ```go run cmd/main.go play --move up```
- Move snake head down => ```go run cmd/main.go play --move down```
- Move snake head left => ```go run cmd/main.go play --move left```
- Move snake head right => ```go run cmd/main.go play --move right```

Once a game has been started (and is in progress), the game can be rendered at any time by entering the following command:
```go run cmd/main.go info```

### Game Over conditions
The game will end automatically if one of the following conditions are met.
- Snake head position touches any of the board's borders
- Snake head touched the snake body (see **Known issues** below)

### Ending a game
A game can be manually ended by entering the following command:
```go run cmd/main.go endgame```

# Known issues
The game does not enter "game over" mode when a snake is of length 1 and the user moves the snakes head onto its body.