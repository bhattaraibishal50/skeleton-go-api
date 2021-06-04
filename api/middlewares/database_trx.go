package middlewares

import (
	"behealth-api/constants"
	"behealth-api/infrastructure"
	"behealth-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//DBTransactionMiddleware -> struct for transaction
type DBTransactionMiddleware struct {
	handler infrastructure.RequestHandler
	logger  infrastructure.Logger
	db      infrastructure.Database
}

//NewDBTransactionMiddleware -> new instance of transaction
func NewDBTransactionMiddleware(
	handler infrastructure.RequestHandler,
	logger infrastructure.Logger,
	db infrastructure.Database,
) DBTransactionMiddleware {
	return DBTransactionMiddleware{
		handler: handler,
		logger:  logger,
		db:      db,
	}
}

//Handle -> It setup the database transaction middleware
func (m DBTransactionMiddleware) Handle() gin.HandlerFunc {
	m.logger.Zap.Info("setting up database transaction middleware")

	return func(c *gin.Context) {
		txHandle := m.db.DB.Begin()
		m.logger.Zap.Info("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				if err := txHandle.Error; err != nil {
					m.logger.Zap.Error("trx commit error: ", err)
				}
				txHandle.Rollback()
			}
		}()

		c.Set(constants.DBTransaction, txHandle)
		c.Next()

		if utils.StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			m.logger.Zap.Info("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				m.logger.Zap.Error("trx commit error: ", err)
			}
		} else {
			m.logger.Zap.Info("rolling back transaction due to status code: ", c.Writer.Status())
			txHandle.Rollback()
		}
	}
}
