package primes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= sqrt(num); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func sqrt(start int) int {
	var t int = start
	var root int = start >> 1

	for (t - root) > 0 {
		t = root
		root = (t + (start / t)) >> 1
	}

	return root
}

func CheckPrimes(c *fiber.Ctx) error {
	var nums []interface{}
	if err := c.BodyParser(&nums); err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	result := make([]bool, len(nums))
	for i, num := range nums {
		switch val := num.(type) {
		case float64:
			if isPrime(int(val)) {
				result[i] = true
			} else {
				result[i] = false
			}
		default:
			return c.Status(400).JSON(map[string]interface{}{
				"error": fmt.Sprintf("Element on index %d is not valid", i),
			})
		}
	}

	return c.JSON(result)
}
