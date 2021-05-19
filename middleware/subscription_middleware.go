package middleware

import (
	"RedSaludAtv/atv/dao"
	"RedSaludAtv/config"
	"RedSaludAtv/utils"
	"net/http"
)

func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := config.GetMySQLDB()
	database := dao.SubscriptionDao{
		Db: db,
	}
	subscriptions, err := database.FindAll()

	if err == nil {
		utils.RespondWithSuccess(subscriptions, w)
	} else {
		utils.RespondWithError(err, w)
	}
}

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

}
