package main

import (
	"conferencia/goClient/models"
	confproto "conferencia/goClient/proto"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// get method with the id parameter
	app.Get("/get-appointments-grpc", func(c *fiber.Ctx) error {

		id := c.Query("id")

		start := time.Now()

		// get the appointments
		data, size := getAppointmentsgRPCBasic(id)

		elapsed := time.Since(start)

		// convert the size to MB
		sizeFloat := float64(size) / 1024 / 1024

		log.Printf("getAppointmentsgRPCBasic took %s and the size is %f MB", elapsed, sizeFloat)

		return c.JSON(data)
	})

	// get method with the id parameter
	app.Get("/get-appointments-http", func(c *fiber.Ctx) error {

		id := c.Query("id")

		start := time.Now()

		// get the appointments
		data, size := getAppointmentsHttp(id)

		elapsed := time.Since(start)

		sizeMB := float64(size) / 1024.0 / 1024.0

		log.Printf("getAppointmentsHttp took %s and the size is %f MB", elapsed, sizeMB)

		return c.JSON(data)

	})

	app.Listen(":3000")
}

// get appointments using gRPC the basic form
func getAppointmentsgRPCBasic(id string) ([]*confproto.Appointment, int) {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := confproto.NewAppointmentsServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}

	}(conn)

	// get the appointments
	response, err := cl.GetAppointments(context.Background(), &confproto.GetAppointmentsRequest{
		DoctorId: id,
	})

	if err != nil {
		log.Fatalln(err)
	}

	size := proto.Size(response)

	return response.Appointments, size

}

// get appointments using Http
func getAppointmentsHttp(id string) ([]models.Appointment, int) {

	resp, err := http.Post("http://localhost:3002", "application/json", strings.NewReader(`{"id":"`+id+`"}`))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var appointments []models.Appointment
	err = json.Unmarshal(body, &appointments)
	if err != nil {
		log.Fatalln(err)
	}

	return appointments, len(body)
}
