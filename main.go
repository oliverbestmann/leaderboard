package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

func main() {
	db := sqlx.MustConnect("sqlite", "db/highscore.db")

	db.MustExec("CREATE TABLE IF NOT EXISTS highscore (game text, player text, score int)")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/games/:game/highscore", func(c echo.Context) error {
		ctx := c.Request().Context()

		var query struct {
			Player    string `query:"player"`
			Highscore int    `query:"score"`
		}

		if err := (&echo.DefaultBinder{}).BindQueryParams(c, &query); err != nil {
			return err
		}

		game := c.Param("game")

		tx, err := db.Beginx()
		if err != nil {
			return err
		}

		defer tx.Rollback()

		if _, err := tx.ExecContext(ctx, "INSERT INTO highscore (game, player, score) VALUES (?, ?, ?)", game, query.Player, query.Highscore); err != nil {
			return err
		}

		var rows []struct {
			Player string `db:"player" json:"player"`
			Score  int    `db:"score" json:"score"`
		}

		if err := tx.SelectContext(ctx, &rows, "SELECT player, MAX(score) as score FROM highscore WHERE game=$1 GROUP BY 1 ORDER BY 2 DESC LIMIT 100", game); err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, rows)
	})

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
