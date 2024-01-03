package database

type Car struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

type Database struct {
	Data map[string]Car
}

// get all cars
func (db *Database) GetAll() (map[string]Car, bool) {
	return db.Data, true
}

// create new car
func (db *Database) Create(car Car) (Car, bool) {
	db.Data[car.ID] = car
	return car, true
}

// update a car
func (db *Database) Update(car Car) (Car, bool) {
	if _, ok := db.Data[car.ID]; !ok {
		panic("Car with this ID does not exist")
	}
	db.Data[car.ID] = car
	return car, true
}

func (db *Database) GetByBrand(brand string) ([]Car, bool) {
	var result []Car
	for _, car := range db.Data {
		if car.Brand == brand {
			result = append(result, car)
		}
	}
	if len(result) == 0 {
		return nil, false
	}
	return result, true

}
