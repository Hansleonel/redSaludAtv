package dao

import (
	"RedSaludAtv/atv/entites/payments"
)

func (subscriptionDao SubscriptionDao) FindClients(id int64) ([]payments.PaymentClient, error) {
	var paymentClientArray []payments.PaymentClient

	rows, err := subscriptionDao.Db.Query("SELECT idasegurado_suscripcion FROM asegurado_suscripcion WHERE idcontratante = ?", id)
	subscriptionDao.Db.Close()
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var paymentClient payments.PaymentClient
			err = rows.Scan(&paymentClient.IdClient)

			if err != nil {
				return nil, err
			} else {
				paymentClientArray = append(paymentClientArray, paymentClient)
			}

		}
	}

	return paymentClientArray, err
}
