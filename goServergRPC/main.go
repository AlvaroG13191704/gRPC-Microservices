package main

import (
	confproto "conferencia/goClientgRPC/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type AppointmentServer struct {
	confproto.UnimplementedAppointmentsServiceServer
}

// GetAppointments is the method that will be called by the client
func (s *AppointmentServer) GetAppointments(ctx context.Context, in *confproto.GetAppointmentsRequest) (*confproto.GetAppointmentsResponse, error) {

	fmt.Println("GetAppointments called")

	fmt.Println("Doctor ID: ", in.DoctorId)

	return &confproto.GetAppointmentsResponse{
		Appointments: []*confproto.Appointment{
			{
				Id:          "1",
				DoctorId:    in.DoctorId,
				PatientName: "John Doe",
				Description: "Checkup",
			},
		},
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":3001")

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	// register the service
	confproto.RegisterAppointmentsServiceServer(server, &AppointmentServer{})

	fmt.Println("Server running on port :3001")
	// start the server
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
