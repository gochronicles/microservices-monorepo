package queries

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go-fiber-template/internal/domain/models"
	"go-fiber-template/pkg/utils"
	"go-fiber-template/platform/database"
)

type Domain models.Domain

func (p *Domain) Create(tx neo4j.Transaction) (interface{}, error) {
	domain := utils.ConvertToMap(*p)
	params := map[string]interface{}{"params": domain}
	records, err := tx.Run(`CALL apoc.create.node(["Domain"],$params) yield node as n return apoc.convert.toMap(n) as domain`, params)
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

func (p *Domain) GetById(id string) (interface{}, error) {
	sess := database.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer sess.Close()
	result, err := sess.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		records, err := tx.Run("MATCH (d:Domain{id:$id}) RETURN apoc.convert.toMap(d) as domain", map[string]interface{}{"id": id})
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

func (p *Domain) GetAll(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("match (d:Domain)  with {domains:collect(apoc.convert.toMap(d))} as domains return domains", nil)
	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	if err != nil {
		return nil, err
	}
	return &record.Values[0], nil
}
