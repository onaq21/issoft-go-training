package db 

func SetupDatabase() ([]string, error) {
	db, err := Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err = CreateTable(db); err != nil {
		return nil, err
	}

	errorsList, err := FillTable(db)
	if err != nil {
		return nil, err
	}

	return errorsList, nil
}