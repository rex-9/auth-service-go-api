package helper

import (
    "auth-service-go-api/dto/response"
    "net/http"

    "github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Execute request handlers
        c.Next()

        // If there are any errors
        if len(c.Errors) > 0 {
            err := c.Errors.Last()
            
            var errorResponse response.ErrorResponse
            
            switch e := err.Err.(type) {
            case *response.ErrorResponse:
                errorResponse = response.ErrorResponse{
                    Code:    e.Code,
                    Message: e.Message,
                    Details: e.Details,
                }
                c.JSON(e.Code, response.ToErrorResponseWrapper[any](errorResponse))
            default:
                errorResponse = response.ErrorResponse{
                    Code:   500,
                    Message: "INTERNAL SERVER OCCUR",
                    Details: "An unexpected error occurred",
                }
                c.JSON(http.StatusInternalServerError, response.ToErrorResponseWrapper[any](errorResponse))
            }
        }
    }
}