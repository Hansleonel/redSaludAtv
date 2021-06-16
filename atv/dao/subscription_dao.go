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
	subscriptionDao.Db.Close()
	if err != nil {
		return err
	} else {
		subscription.Id, err = result.LastInsertId()
		return nil
	}
}

func (subscriptionDao SubscriptionDao) CreateStepOne(subscriptionStepOne *entites.SubscriptionStepOne) error {
	result, err := subscriptionDao.Db.Exec("INSERT INTO asegurado_suscripcion(idtipodocumento, nro_documento, telefono, tipo, tc_datos, idplan, frecuencia_pago, tipo_afiliacion, importe, sexo) values(?,?,?,?,?,?,?,?,?,?)", subscriptionStepOne.TypeDoc, subscriptionStepOne.NumDoc, subscriptionStepOne.CelNumber, subscriptionStepOne.Type, subscriptionStepOne.TcDatos, subscriptionStepOne.IdPlan, subscriptionStepOne.FrecuenciaPago, subscriptionStepOne.TipoAfiliacion, subscriptionStepOne.Importe, subscriptionStepOne.Sexo)
	subscriptionDao.Db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		subscriptionStepOne.Id, err = result.LastInsertId()
		return nil
	}
}

func (subscriptionDao SubscriptionDao) CreateFamiliar(subscriptionFamiliar *entites.SubsFamiliar) error {
	result, err := subscriptionDao.Db.Exec("INSERT INTO asegurado_suscripcion(idtipodocumento, nro_documento, tipo, tc_datos, idplan, frecuencia_pago, tipo_Afiliacion, fecha_nacimiento, nombre1, nombre2, apellido_paterno, apellido_materno, idcontratante, importe) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", subscriptionFamiliar.TypeDoc, subscriptionFamiliar.NumDoc, subscriptionFamiliar.Type, subscriptionFamiliar.TcDatos, subscriptionFamiliar.IdPlan, subscriptionFamiliar.FrecuenciaPago, subscriptionFamiliar.TipoAfiliacion, subscriptionFamiliar.FechaNacimiento, subscriptionFamiliar.Nom1, subscriptionFamiliar.Nom2, subscriptionFamiliar.Ape1, subscriptionFamiliar.Ape2, subscriptionFamiliar.IdContratante, subscriptionFamiliar.Importe)
	subscriptionDao.Db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		subscriptionFamiliar.Id, err = result.LastInsertId()
		return nil
	}
}

func (subscriptionDao SubscriptionDao) Update(subscriptionStepTwo entites.SubscriptionStepTwo) (int64, error) {
	result, err := subscriptionDao.Db.Exec("UPDATE asegurado_suscripcion SET fecha_nacimiento = ?, nombre1 = ?, nombre2 = ?, apellido_paterno = ?, apellido_materno = ?, correo = ?, idcontratante = ? WHERE idasegurado_suscripcion = ?", subscriptionStepTwo.FechaNacimiento, subscriptionStepTwo.Nom1, subscriptionStepTwo.Nom2, subscriptionStepTwo.Ape1, subscriptionStepTwo.Ape2, subscriptionStepTwo.Mail, subscriptionStepTwo.IdContratante, subscriptionStepTwo.Id)
	subscriptionDao.Db.Close()
	if err != nil {
		return 0, err
	} else {
		// subscriptionStepTwo.Id, err = result.LastInsertId()
		return result.RowsAffected()
	}
}

func (subscriptionDao SubscriptionDao) FindAll() ([]entites.Subscription, error) {
	var subscriptionsArray []entites.Subscription
	rows, err := subscriptionDao.Db.Query("select idasegurado_suscripcion, nro_documento, apellido_parteno, apellido_materno, nombre1, nombre2 from asegurado_suscripcion limit 5")
	subscriptionDao.Db.Close()
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var subscription entites.Subscription
			var numDoc interface{}
			err = rows.Scan(&subscription.Id, &numDoc, &subscription.Ape1, &subscription.Ape2, &subscription.Nom1, &subscription.Nom2)

			if err != nil {
				return nil, err
			} else {
				subscription.NumDoc = fmt.Sprintf("%s", numDoc)
				subscriptionsArray = append(subscriptionsArray, subscription)
			}
		}
	}

	return subscriptionsArray, err
}

func (subscriptionDao SubscriptionDao) ValidationDni(validation entites.ValidationPerson) (int8, error) {
	result, err := subscriptionDao.Db.Query("select count(a.idasegurado) as cant from certificado c inner join certificado_asegurado ca on ca.idcertificado=c.idcertificado inner join asegurado a on a.idasegurado = ca.idasegurado and a.nro_documento=? where c.idplan=1 and estado_afiliacion=1", validation.DniValidation)
	subscriptionDao.Db.Close()
	var cant int8
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	} else {
		for result.Next() {
			err = result.Scan(&cant)
			fmt.Println(err)
			fmt.Println(cant)
		}
		return cant, err
	}
}

func (subscriptionDao SubscriptionDao) UpdateDeclaration(subscriptionStepThree entites.SubscriptionStepThree) (int64, error) {
	result, err := subscriptionDao.Db.Exec("UPDATE asegurado_suscripcion SET declaracion_jurada = ?, pregunta1 = ?, pregunta2 = ?, pregunta3 = ? WHERE idasegurado_suscripcion = ?", subscriptionStepThree.DecJur, subscriptionStepThree.QuestionFirst, subscriptionStepThree.QuestionSecond, subscriptionStepThree.QuestionThird, subscriptionStepThree.Id)
	subscriptionDao.Db.Close()
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (subscriptionDao SubscriptionDao) CreateQuestionDetail(questionDetail *entites.SubsQuestions) error {
	result, err := subscriptionDao.Db.Exec("INSERT INTO suscripcion_detalle(idasegurado_suscripcion, pregunta, descripcion) values(?,?,?)", questionDetail.IdSubscription, questionDetail.Question, questionDetail.Description)
	subscriptionDao.Db.Close()
	if err != nil {
		return err
	} else {
		questionDetail.IdQuestion, err = result.LastInsertId()
		return nil
	}
}

func (subscriptionDao SubscriptionDao) CreateDeclaration(questions []entites.SubsDeclaration) error {
	for i := 0; i < len(questions); i++ {
		result, err := subscriptionDao.Db.Exec("INSERT INTO suscripcion_detalle(idasegurado_suscripcion, pregunta, descripcion) VALUES (?,?,?)", questions[i].IdSubscription, questions[i].Question, questions[i].Description)
		if err != nil {
			return err
		} else {
			questions[i].Id, err = result.LastInsertId()
		}
	}
	subscriptionDao.Db.Close()

	return nil
}
