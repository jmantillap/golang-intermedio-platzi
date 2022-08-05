package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Job represents a job to be executed, with a name and a number and a delay
type Job struct {
	Name   string        // name of the job
	Delay  time.Duration // delay between each job
	Number int           // number to calculate on the fibonacci sequence
}

// Worker will be our concurrency-friendly worker
type Worker struct {
	Id         int           // id of the worker
	JobQueue   chan Job      // Jobs to be processed
	WorkerPool chan chan Job // Pool of workers
	QuitChan   chan bool     // Quit worker
}

// Dispatcher is a dispatcher that will dispatch jobs to workers
type Dispatcher struct {
	WorkerPool chan chan Job // Pool of workers
	MaxWorkers int           // Maximum number of workers
	JobQueue   chan Job      // Jobs to be processed
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,              //Se asigna un id
		WorkerPool: workerPool,      //Se le indica el canal donde tiene quie agregar su canal de tareas
		JobQueue:   make(chan Job),  //Canal de tareas del worker
		QuitChan:   make(chan bool), //Canal para parar al worker
	}
}

//El dispatcher cuenta con el el canal global de jobs y un canal de todos los canales de los workers

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {

	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

func (d *Dispatcher) Dispatch() {
	//Inicia de manera indefinidad a mandar jobs a los canales que se van recibiendo en el canal de caneles de jobs
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				//Si se recibe un job en el canal de tareas del worker se ejecuta
				fmt.Printf("Worker with id %d Started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finishes with result %d\n", w.Id, fib)

			case <-w.QuitChan:
				//Si se recibe un job en el canal de salida se para el worker (lo sca del ciclo)
				fmt.Printf("Worker with id %d Stopped\n", w.Id)
				return
			}
		}
	}()

}

//La funcion stop manda un true al canl de salida del worker
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" { // GET, PUT, DELETE
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()
	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
