package main

import (
	appointmentProto "conferencia/goClient/appointmentProto"
	benchmarkProto "conferencia/goClient/benchmarkProto"
	"conferencia/goClient/models"
	"context"
	"encoding/json"
	"math"

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

	var benchmark = &benchmarkProto.BenchmarkRequest{
		Http: &benchmarkProto.Http{
			Time:   []float32{},
			Weight: []float32{},
		},
		Grpc: &benchmarkProto.Grpc{
			Time:   []float32{},
			Weight: []float32{},
		},
	}

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

		// add the time and the size to the benchmark
		time := float64(elapsed.Milliseconds())
		time = math.Round(time*100) / 100 // Round to 2 decimal places
		benchmark.Grpc.Time = append(benchmark.Grpc.Time, float32(time))

		sizeMB := float64(size) / 1024.0 / 1024.0
		sizeMB = math.Round(sizeMB*100) / 100 // Round to 2 decimal places
		benchmark.Grpc.Weight = append(benchmark.Grpc.Weight, float32(sizeMB))

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

		// add the time and the size to the benchmark
		time := float64(elapsed.Milliseconds())
		time = math.Round(time*100) / 100 // Round to 2 decimal places
		benchmark.Http.Time = append(benchmark.Http.Time, float32(time))

		sizeMB = math.Round(sizeMB*100) / 100 // Round to 2 decimal places
		benchmark.Http.Weight = append(benchmark.Http.Weight, float32(sizeMB))

		return c.JSON(data)

	})

	// send the benchmark results to the server
	app.Get("/send-benchmark", func(c *fiber.Ctx) error {

		sendBenchmarkResults(benchmark)

		return c.SendString("Benchmark sent")
	})

	app.Listen(":3000")
}

// get appointments using gRPC the basic form
func getAppointmentsgRPCBasic(id string) ([]*appointmentProto.Appointment, int) {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := appointmentProto.NewAppointmentsServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}

	}(conn)

	// get the appointments
	response, err := cl.GetAppointments(context.Background(), &appointmentProto.GetAppointmentsRequest{
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

// Send the benchmark results to the server (python)
func sendBenchmarkResults(data *benchmarkProto.BenchmarkRequest) {
	conn, err := grpc.Dial("localhost:3003", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := benchmarkProto.NewBenchmarkClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}

	}(conn)

	// send the benchmark
	_, err = cl.Benchmark(context.Background(), data)
	if err != nil {
		log.Fatalln(err)
	}
}
