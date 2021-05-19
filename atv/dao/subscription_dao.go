package dao

import (
	"RedSaludAtv/atv/entites"
	"database/sql"
)

type SubscriptionDao struct {
	Db *sql.DB
}

func (subscriptionDao SubscriptionDao) Create(subscription entites.Subscription) error {
	result, err := subscriptionDao.Db.Exec("INSERT INTO asegurado_suscripcion(fecha_nacimiento, telefono, tipo, tc_datos, idplan, frecuencia_pago, tipo_afiliacion) values(STR_TO_DATE(? ,'%d-%m-%Y'),?,?,?,?,?,?)", subscription.FechaNacimiento, subscription.CelNumber, 3, subscription.TcDatos, subscription.IdPlan, subscription.FrecuenciaPago, subscription.TipoAfiliacion)
	if err != nil {
		return err
	} else {
		subscription.Id, err = result.LastInsertId()
		return nil
	}
}

func (subscriptionDao SubscriptionDao) FindAll() ([]entites.Subscription, error) {
	var subscriptionsArray []entites.Subscription
	rows, err := subscriptionDao.Db.Query("select idasegurado_suscripcion, nro_documento, apellido_parteno, apellido_materno, nombre1, nombre2 from asegurado_suscripcion	")

	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var subscription entites.Subscription

			err = rows.Scan(&subscription.Id, &subscription.NumDoc, &subscription.Ape1, &subscription.Ape2, &subscription.Nom1, &subscription.Nom2)

			if err != nil {
				return nil, err
			} else {
				subscriptionsArray = append(subscriptionsArray, subscription)
			}
		}
	}

	return subscriptionsArray, err
}
