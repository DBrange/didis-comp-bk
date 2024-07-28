package services

import (
	"fmt"
	"strings"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/config"
	ports "github.com/DBrange/didis-comp-bk/domains/control_plane/ports/drivens"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ControlPlaneService struct {
	controlPlaneQuerier ports.ForQueryingControlPlane
}

func NewControlPlaneService(controlPlaneQuerier ports.ForQueryingControlPlane) *ControlPlaneService {
	return &ControlPlaneService{
		controlPlaneQuerier: controlPlaneQuerier,
	}
}

func (s *ControlPlaneService) AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			s.permissionDenied(c, "when reading authorization header")
			return
		}

		token, err := s.validateToken(authHeader)
		if err != nil || !token.Valid {
			s.permissionDenied(c, "invalid token")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["sub"].(string)
			if !ok {
				s.permissionDenied(c, "invalid user ID claim")
				return
			}
			c.Set("userID", userID)

			roles, ok := claims["roles"]
			if !ok {
				s.permissionDenied(c, "invalid user ID claim")
				return
			}
			c.Set("roles", roles)

		} else {
			s.permissionDenied(c, "invalid token claims")
			return
		}

		c.Next()
	}
}

func (s *ControlPlaneService) AuthorizationMiddleware(requiredRole []models.ROLE) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesAny, exists := c.Get("roles")
		if !exists {
			s.permissionDenied(c, "unauthorized")
			return
		}

		rolesInterface, ok := rolesAny.([]interface{})
		if !ok {
			s.permissionDenied(c, "invalid roles format")
			return
		}

		roles := make([]models.ROLE, len(rolesInterface))
		for i, v := range rolesInterface {
			roleStr, ok := v.(string)
			if !ok {
				s.permissionDenied(c, "invalid role type")
				return
			}

			roles[i] = models.ROLE(roleStr)
		}

		roleMap := make(map[models.ROLE]struct{})
		for _, rr := range requiredRole {
			roleMap[rr] = struct{}{}
		}

		for _, r := range roles {
			if _, ok := roleMap[r]; ok {
				c.Next()
				return
			}
		}

		s.permissionDenied(c, "forbidden")
	}
}

func (s *ControlPlaneService) validateToken(authHeader string) (*jwt.Token, error) {
	token := strings.TrimPrefix(authHeader, "Bearer ")
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("%w: unexpected signing method: %v", customerrors.ErrNotFound, token.Header["alg"])
			return nil, customerrors.HandleErrMsg(err, "auth", "unexpected signing method")
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func (s *ControlPlaneService) permissionDenied(c *gin.Context, text string) {
	err := fmt.Errorf("error authorization: %w", customerrors.ErrAuthorization)
	errMsgTemplate := fmt.Sprintf("error %s", text)
	customerrors.HandleErrMsg(err, "auth", errMsgTemplate)
	customerrors.ErrorResponse(err, c)
	// c.JSON(http.StatusInternalServerError, gin.H{"status": customerrors.ErrAuthorization, "error": err})
	c.Abort()
}
