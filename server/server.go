package server

import (
    "myapp/database"
    "myapp/models"
    "net/http"

    "github.com/labstack/echo/v4"
)

type Server struct {
    Echo      *echo.Echo
    DB        *database.Database
}

func NewServer() *Server {
    db, err := database.NewDatabase("root:amir1379@tcp(127.0.0.1:3306)/myapp")
    if err != nil {
        panic(err)
    }

    s := &Server{
        Echo: echo.New(),
        DB:   db,
    }

    s.routes()
    return s
}

func (s *Server) routes() {
    s.Echo.GET("/users", s.getUsers)
}

func (s *Server) getUsers(c echo.Context) error {
    rows, err := s.DB.Connection.Query("SELECT id, name FROM users")
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error querying database")
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name); err != nil {
            return c.String(http.StatusInternalServerError, "Error scanning rows")
        }
        users = append(users, user)
    }

    return c.JSON(http.StatusOK, users)
}

func (s *Server) Start() {
    defer s.DB.Close()
    s.Echo.Logger.Fatal(s.Echo.Start(":8080"))
}
