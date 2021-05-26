package dao

import (
	"RedSaludAtv/atv/entites"
	"database/sql"
	"fmt"
)

type SubscriptionDao struct {
	Db *sql.DB
}

func (subscriptionDao SubscriptionDao) Create(subscription *entites.Subscription) error {
	fmt.Println(subscription.FechaNacimiento)
	result, err := subscriptionDao.Db.Exec("INSERT INTO asegurado_suscripcion(idtipodocumento, nro_documento, apellido_parteno, apellido_materno, nombre1, nombre2, fecha_nacimiento, correo, telefono, tipo, idcontratante, tc_datos, tc_comunicaciones, tc_pago, estado, tipo_procesamiento, idplan, frecuencia_pago, tipo_afiliacion) values(?,?,?,?,?,?,STR_TO_DATE(? ,'%d-%m-%Y'),?,?,?,?,?,?,?,?,?,?,?,?)", subscription.TypeDocument, subscription.NumDoc, subscription.Ape1, subscription.Ape2, subscription.Nom1, subscription.Nom2, subscription.FechaNacimiento, subscription.Mail, subscription.CelNumber, subscription.Type, subscription.IdContratante, subscription.TcDatos, subscription.TcComunicaciones, subscription.TcPagos, subscription.Estado, subscription.TypeProcesamiento, subscription.IdPlan, subscription.FrecuenciaPago, subscription.TipoAfiliacion)
	if err != nil {
		return err
	} else {
		subscription.Id, err = result.LastInsertId()
		return nil
	}
}

func (subscriptionDao SubscriptionDao) CreateStepOne(subscriptionStepOne *entites.SubscriptionStepOne) error {
	result, err := subscriptionDao.Db.Exec("INSERT INTO asegurado_suscripcion(fecha_nacimiento, telefono, tipo, tc_datos, idplan, frecuencia_pago, tipo_afiliacion) values(STR_TO_DATE(? ,'%d-%m-%Y'),?,?,?,?,?,?)", subscriptionStepOne.FechaNacimiento, subscriptionStepOne.CelNumber, subscriptionStepOne.Type, subscriptionStepOne.TcDatos, subscriptionStepOne.IdPlan, subscriptionStepOne.FrecuenciaPago, subscriptionStepOne.TipoAfiliacion)
	if err != nil {
		return err
	} else {
		subscriptionStepOne.Id, err = result.LastInsertId()
		return nil
	}
}

func (subscriptionDao SubscriptionDao) Update(subscriptionStepTwo entites.SubscriptionStepTwo) (int64, error) {
	result, err := subscriptionDao.Db.Exec("UPDATE asegurado_suscripcion set nro_documento = ?, nombre1 = ?, nombre2 = ?, apellido_parteno = ?, apellido_materno = ?, correo = ? WHERE idasegurado_suscripcion = ?", subscriptionStepTwo.NumDoc, subscriptionStepTwo.Nom1, subscriptionStepTwo.Nom2, subscriptionStepTwo.Ape1, subscriptionStepTwo.Ape2, subscriptionStepTwo.Mail, subscriptionStepTwo.Id)
	if err != nil {
		return 0, err
	} else {
		// subscriptionStepTwo.Id, err = result.LastInsertId()
		return result.RowsAffected()
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
