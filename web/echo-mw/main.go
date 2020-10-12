package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: make(map[string]int),
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func MyMw1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println(&c)
		println("start mw1")
		var n = next(c)
		println("end mw1")
		return n
	}

}

func MyMw2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println(&c)
		println("start  mw2")

		var n = next(c)
		println("end mw2")
		return n
	}

}

func main() {
	e := echo.New()
	// Debug mode
	e.Debug = true
	//-------------------
	// Custom middleware
	//-------------------
	// Stats
	s := NewStats()
	e.Use(s.Process)
	e.GET("/stats", s.Handle) // Endpoint to get stats
	// Server header
	e.Use(ServerHeader)
	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/order", func(c echo.Context) error {
		println("heel")
		return c.String(http.StatusOK, "Hello, World!")
	}, MyMw1, MyMw2)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
