package queries

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go-fiber-template/internal/patient/models"
	"go-fiber-template/pkg/utils"
	"go-fiber-template/platform/database"
)

type Patient models.Patient

func (p *Patient) Create(tx neo4j.Transaction) (interface{}, error) {
	patient := utils.ConvertToMap(*p)
	params := map[string]interface{}{"params": patient}
	records, err := tx.Run(`CALL apoc.create.node(["Patient"],$params) yield node as n return apoc.convert.toMap(n) as patient`, params)
	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	fmt.Println(record)
	if err != nil {
		return nil, err
	}
	return &record.Values[0], nil
}

func (p *Patient) GetByMRN(mrn string) (interface{}, error) {
	sess := database.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer sess.Close()
	result, err := sess.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		records, err := tx.Run("MATCH (p:Patient{patientMRN:$mrn}) RETURN apoc.convert.toMap(p) as patient", map[string]interface{}{"mrn": mrn})
		if err != nil {
			return nil, err
		}
		record, err := records.Single()
		if err != nil {
			return nil, err
		}
		return &record.Values[0], nil

	})
	return &result, err
}

func (p *Patient) GetAll(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("match (n:Patient)  with {patients:collect(apoc.convert.toMap(n))} as patients return patients", nil)
	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	if err != nil {
		return nil, err
	}
	return &record.Values[0], nil
}
