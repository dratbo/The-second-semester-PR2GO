package main

import (
	"context"
	"log"
	"time"

	"example.com/pz2-grpc/gen/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("Не удалось подключиться к серверу:", err)
	}
	defer conn.Close()

	client := studentpb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pingResp, err := client.Ping(ctx, &studentpb.PingRequest{
		Message: "hello grpc",
	})
	if err != nil {
		log.Fatal("Ошибка Ping:", err)
	}
	log.Println("Ping response:", pingResp.GetMessage())

	studentResp, err := client.GetStudentByID(ctx, &studentpb.GetStudentRequest{
		Id: 1,
	})
	if err != nil {
		log.Println("Ошибка GetStudentByID (ожидаемо, т.к. студента нет):", err)
	} else {
		st := studentResp.GetStudent()
		log.Printf("Student: id=%d, full_name=%s, group=%s, email=%s\n",
			st.GetId(),
			st.GetFullName(),
			st.GetGroup(),
			st.GetEmail(),
		)
	}

	// 3. Получение списка всех студентов
	listResp, err := client.ListStudents(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal("Ошибка ListStudents:", err)
	}
	log.Println("Список всех студентов:")
	for _, st := range listResp.GetStudents() {
		log.Printf("ID: %d, Name: %s, Group: %s, Email: %s\n",
			st.GetId(),
			st.GetFullName(),
			st.GetGroup(),
			st.GetEmail(),
		)
	}
}
