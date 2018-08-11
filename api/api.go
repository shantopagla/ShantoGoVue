package api

import (
	"github.com/shantopagla/ShantoGoVue/models"
)

// API -
type API struct {
	users  *models.UserManager
	quotes *models.QuoteManager
	views  *models.ViewManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	usermgr, _ := models.NewUserManager(db)
	quotemgr, _ := models.NewQuoteManager(db)
	viewmgr, _ := models.NewViewManager(db)

	return &API{
		users:  usermgr,
		quotes: quotemgr,
		views: viewmgr,
	}
}
