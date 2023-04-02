package Pet_project_ToDoApp // объявим структуру http сервера, которую будем использовать для запуска http сервера

import (
	"context"
	"net/http"
	"time"
)

type Server struct { // небольшая абстракция над структурой сервера из пакета http и имеет всего одно поле - указатель на эту структуру
	httpServer *http.Server
}

// у сервера два метода - запуск и остановка работы
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{ //инкапсулируем значения для
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe() // лиснэндсерв под капотом запускает бесконечный цикл фор и слушает все входящие запросы для последующей обработки
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
