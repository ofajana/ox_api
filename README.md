I've written the api so it can handle multiple conccurent games.
I opted to use the gorrila mux router because it's battle tested and production ready.
Data integrity is maintained with the use of Mutex locks on individual games, so if both players i.e X and O for a game
are accessing the api at the same time, neither would get dirty data 

Usage:

1. Download and Install App
        go get github.com/ofajana/ox_api
        go install github.com/ofajana/ox_api

2. Start Api Server
        ox_api [-p portNo]    NB: -p flag is optional, server will default to port 8080 if a port isn't provided
   
3. Start a Game PS: I'll assume app was started on default port i.e. 8080

        api endpoint:   localhost:8080

        expected response

        {"GameId":"TEWeq","Topic":"Start Game","Message":"New Game Successfuly Initiated"}

4. Play a turn as either player X or O
NB: Each box on the game board has a number i.e. 1-9

        1  2  3
        4  5  6
        7  8  9

        each play command should specify the following attributes
        1. Game id
        2. player i.e X or O
        3. Box number i.e. 1-9

        api endpoint http://localhost:8080/play/{gameId}/{player}/{boxId}
        e.g. http://localhost:8080/play/TEWeq/X/1
        example responses
        {"GameId":"TEWeq","Topic":"Play Turn","Message":"Move Successfully Recorded"}
        {"GameId":"TEWeq","Topic":"Play Turn","Message":"Invalid Play, Try another Square"}
        {"GameId":"TEWeq","Topic":"Play Turn","Message":"Not Your Turn"}
        
 5.  Check Current State of Game
        each check game state command should specify the following attributes
        1. Game id
        
        api endpoint http://localhost:8080/games/state/{gameId}
        e.g. http://localhost:8080/games/state/TEWeq
        
        example responses
        {"GameId":"YptfZ","Topic":"Game's Current State","Message":"No Game With Id Provided"}
        {
          GameId: "TEWeq",
          Topic: "Game's Current State",
          Message: "Game is still in Play, and there are 8 empty boxes on the Game Board 'X' has had 1 turn(s) 'Y' has had 0 turn(s) Next to play 'O'"
        }
        
  6.   Check Result of Game
        each check game result command should specify the following attributes
        1. Game id
        
        api endpoint http://localhost:8080/games/result/{gameId}
        e.g. http://localhost:8080/games/result/TEWeq
        
        example responses
        {"GameId":"TEWeq","Topic":"Game Result","Message":"Game Is Still In Play, No Result Yet"}
        {"GameId":"TEWeq","Topic":"Game Result","Message":"Game Ended in a Draw"}
        {"GameId":"TEWeq","Topic":"Game Result","Message":"Game was Won By 'X'"}


        
