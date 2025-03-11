package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// Register обрабатывает регистрацию пользователя.
func Register(c *fiber.Ctx, db *pgxpool.Pool) error {
	type Request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Хешируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	// Сохраняем пользователя в базу данных
	_, err = db.Exec(c.Context(), "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", req.Username, req.Email, string(hash))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

// Login обрабатывает вход пользователя.
func Login(c *fiber.Ctx, db *pgxpool.Pool) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Получаем пользователя из базы данных
	var user struct {
		ID           int    `db:"id"`
		Username     string `db:"username"`
		Email        string `db:"email"`
		PasswordHash string `db:"password_hash"`
	}

	err := db.QueryRow(c.Context(), "SELECT id, username, email, password_hash FROM users WHERE email = $1", req.Email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Генерация токена (здесь можно использовать JWT или другую библиотеку)
	token, err := generateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

// Logout обрабатывает выход пользователя.
func Logout(c *fiber.Ctx) error {
	// Здесь можно добавить логику для инвалидации токена (если используется JWT)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out successfully"})
}

// RefreshToken обновляет токен пользователя.
func RefreshToken(c *fiber.Ctx) error {
	// Здесь можно добавить логику для обновления токена (если используется JWT)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Token refreshed successfully"})
}

// generateToken генерирует JWT токен.
func generateToken(userID int) (string, error) {
	// Здесь можно использовать библиотеку для генерации JWT, например, github.com/golang-jwt/jwt
	// Пример:
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//     "user_id": userID,
	//     "exp":     time.Now().Add(time.Hour * 24).Unix(),
	// })
	// return token.SignedString([]byte("your-secret-key"))
	return "dummy-token", nil
}
