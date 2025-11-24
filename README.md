# leaderboard

This leaderboard service allows you to track scores for your games, making it ideal for game jams. The API provides a single endpoint where you can post a player's name and score to add a new entry. The response includes all current leaderboard entries.

It is publicly available at `https://highscore.narf.zone`.

Replace GAME with your game's unique identifier. Add a `-dev` suffix if you want to keep development scores separate.

```sh
curl -X POST 'https://highscore.narf.zone/games/GAME/highscore?player=Oliver&score=9000'
```

Response:

```json
[
   { "player": "Oliver", "score": 9000 },
   { "player": "Foo", "score": 8467 },
   { "player": "Bar", "score": 1337 }
]
```

## Hosting

You can either use the publicly available instance of the service, or run it yourself.
There is a docker container available.

```sh
docker run -d --name=leaderboard -p 8080:8080 -v leaderboard-db:/db ghcr.io/oliverbestmann/leaderboard
```
