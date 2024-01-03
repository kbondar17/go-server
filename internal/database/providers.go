package database

func ProvideDatabase() (*Database, bool) {
	cars := []Car{
		{ID: "id1", Name: "Renault Logan", Brand: "Renault"},
		{ID: "id2", Name: "Renault Duster", Brand: "Renault"},
		{ID: "id3", Name: "BMW X6", Brand: "BMW"},
		{ID: "id4", Name: "BMW M5", Brand: "BMW"},
		{ID: "id5", Name: "VW Passat", Brand: "VW"},
		{ID: "id6", Name: "VW Jetta", Brand: "VW"},
		{ID: "id7", Name: "Audi A4", Brand: "Audi"},
		{ID: "id8", Name: "Audi Q7", Brand: "Audi"},
	}

	db := Database{Data: make(map[string]Car)}
	// add cars to database
	for _, car := range cars {
		db.Create(car)
	}
	return &db, true

}
