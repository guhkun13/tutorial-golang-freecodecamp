package main

import (
	"context"
	"log"
	"time"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoInstance struct {
	Client 	*mongo.Client
	DB			*mongo.Database
}

var mg MongoInstance

const DBName = "gofiber-mg"
const mongoURI = "mongodb://localhost:27017/" + DBName

type Employee struct {
	ID 			string 		`json:"id,omitempty" bson:"_id,omitempty"`
	Name 		string 		`json:"name"`
	Salary 	float64  	`json:"salary"`
	Age  		float64 	`json:"age"`
}

func Connect() error {
	client, _ := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	
	defer cancel() 
	
	err := client.Connect(ctx)
	db := client.Database(DBName)
	
	if err != nil {
		return err 
		// panic("Failed to Connect to Database")
	}
	
	mg = MongoInstance{
		Client: client,
		DB: db,
	}
	return nil
}

func getEmployees(c *fiber.Ctx) error {
	
	query := bson.D{{}}
	
	cursor, err := mg.DB.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	
	var employees []Employee = make([]Employee, 0)
	
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	
	return c.JSON(employees)
}

func getEmployee(c *fiber.Ctx) error {
	idParam := c.Params("id")
	
	objID, err := primitive.ObjectIDFromHex(idParam)
	
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	
	query := bson.D{{Key: "_id", Value: objID}}
	
	var employee Employee
	collection := mg.DB.Collection("employees")
	result := collection.FindOne(c.Context(), query)
	
	fmt.Printf("result : %v \n", result)
	
	result.Decode(&employee)
	
	fmt.Printf("found employee : %v \n", employee)
	if employee.ID == "" {
		return c.Status(404).SendString("Data not found")
	}

	employee.ID = idParam
	
	return c.Status(200).JSON(employee)
}


func createEmployee(c *fiber.Ctx) error {
	
	collection := mg.DB.Collection("employees")
	
	employee := new(Employee)
	
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	
	employee.ID = ""
	
	insertionResult, err := collection.InsertOne(c.Context(), employee)
	
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)
	
	createdEmployee := &Employee{}
	createdRecord.Decode(createdEmployee)
	
	return c.Status(200).JSON(createdEmployee)
	
}

func updateEmployee(c *fiber.Ctx) error {
	fmt.Println("--- updateEmployee ---")
	
	idParam := c.Params("id")
	
	fmt.Printf("idParam: %v \n",idParam)
	
	employeeID, err := primitive.ObjectIDFromHex(idParam)
	
	fmt.Printf("employeeID: %v \n",employeeID)
	
	if err != nil {
		c.Status(400).SendString(err.Error())
	}
	
	employee := new(Employee)
	// var employee Employee
	
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	
	// build query: find the exact record on the db
	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}
	
	collection := mg.DB.Collection("employees")
	err = collection.FindOneAndUpdate(c.Context(), query, update).Err()
	
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(400).SendString(err.Error())
		}
		return c.SendStatus(500)
	}
	
	employee.ID = idParam
	
	return c.Status(200).JSON(employee)
}

func deleteEmployee(c *fiber.Ctx) error {
	fmt.Println("--- deleteEmployee ---")
	
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
	fmt.Printf("employeeID: %v \n", employeeID)
	
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	
	query := bson.D{
		{
			Key: "_id", Value: employeeID,
		},
	}
	collection := mg.DB.Collection("employees")
	
	result, err := collection.DeleteOne(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	
	if result.DeletedCount < 1 {
		return c.Status(404).SendString("Data not found")
	}
	
	return c.SendStatus(200)	
}

func setupRoutes(app *fiber.App) {
	app.Get("/employees", getEmployees)
	app.Post("/employee", createEmployee)
	app.Get("/employee/:id", getEmployee)
	app.Delete("/employee/:id", deleteEmployee)
	app.Put("/employee/:id", updateEmployee)
}

func main(){
	
	if err := Connect(); err != nil{
		log.Fatal(err)
	}
	app := fiber.New()
	
	setupRoutes(app)
	
	log.Fatal(app.Listen(":8010"))
	
}