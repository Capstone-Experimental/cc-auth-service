package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		if err := c.Next(); err != nil {
			return err
		}

		end := time.Now()

		latency := end.Sub(start)

		logMessage := fmt.Sprintf("[%s] %s - %s %s %s %d - Latency: %v",
			end.Format("2006-01-02 15:04:05"),
			c.IP(),
			c.Method(),
			c.OriginalURL(),
			c.Protocol(),
			c.Response().StatusCode(),
			latency,
		)

		// Cetak log ke konsol atau file log
		fmt.Println(logMessage)

		// Kembalikan response ke client
		return nil
	}
}
