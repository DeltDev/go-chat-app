package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Handler struct {
	Service
	secretKey string 
}

func NewHandler(s Service, secretKey string) *Handler {
	return &Handler{
		Service:   s,
		secretKey: secretKey,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var user LoginUserReq

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", u.AccessToken, 3600, "/", "", false, true)

	res := &LoginUserRes{
		Username: u.Username,
		ID:       u.ID,
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        tokenString, err := c.Cookie("jwt")
        if err != nil {

            authHeader := c.GetHeader("Authorization")
            if authHeader == "" {
                c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token found"})
                return
            }

            tokenString = strings.TrimPrefix(authHeader, "Bearer ")
        }

        token, err := jwt.ParseWithClaims(tokenString, &MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {

            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }

            return []byte(h.secretKey), nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
            return
        }

        if claims, ok := token.Claims.(*MyJWTClaims); ok {
            c.Set("id", claims.ID)
            c.Set("username", claims.Username)
        }

        c.Next()
    }
}

func (h *Handler) CheckAuth(c *gin.Context) {

    id, _ := c.Get("id")
    username, _ := c.Get("username")

    c.JSON(200, gin.H{
        "status": "authenticated",
        "id":     id,
        "user":   username,
    })
}