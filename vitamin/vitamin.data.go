package vitamin

import (
	"VitaminApp/database"
	"VitaminApp/graph/model"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Vitamin struct {
	VitaminId   int
	VitaminType string
	Benefits    string
}

func GetVitaminList() ([]*model.Vitamin, error) {
	query := fmt.Sprintf("SELECT VitaminID, VitaminType, Benefits FROM vitamin")
	context, cancel := context.WithTimeout(context.Background(), 8000*time.Millisecond)
	defer cancel()
	results, err := database.DbConn.QueryContext(context, query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	vitamins := make([]*model.Vitamin, 0)
	for results.Next() {
		var vitamin model.Vitamin
		results.Scan(&vitamin.VitaminID, &vitamin.VitaminID, &vitamin.Benefits)

		vitamins = append(vitamins, &vitamin)
	}

	return vitamins, nil
}

func getVitaminById(vitaminId int) (*Vitamin, error) {
	query := fmt.Sprintf("SELECT VitaminID, VitaminType, Benefits FROM Vitamin WHERE VitaminID = %d", vitaminId)
	result := database.DbConn.QueryRow(query)
	vitamin := &Vitamin{}
	err := result.Scan(&vitamin.VitaminId, &vitamin.VitaminType, &vitamin.Benefits)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return vitamin, nil
}

func AddVitamin(vitamin model.NewVitamin) error {
	command := fmt.Sprintf("INSERT INTO Vitamin (VitaminType, Benefits) VALUES ('%s', '%s')", vitamin.VitaminType, vitamin.Benefits)
	_, err := database.DbConn.Exec(command)
	fmt.Println(command)
	if err != nil {
		return err
	}
	return nil
}

func updateVitamin(vitamin Vitamin) error {
	command := fmt.Sprintf("UPDATE Vitamin SET VitaminType = '%s', Benefits = '%s' WHERE VitaminID = %d", vitamin.VitaminType, vitamin.Benefits, vitamin.VitaminId)
	_, err := database.DbConn.Exec(command)
	if err != nil {
		return err
	}
	return nil
}

func deleteVitamin(vitaminId int) error {
	command := fmt.Sprintf("DELETE FROM Vitamin WHERE VitaminID = %d", vitaminId)
	_, err := database.DbConn.Exec(command)
	if err != nil {
		return err
	}
	return nil
}
