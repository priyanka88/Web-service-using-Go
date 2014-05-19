package main

import (

    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
)
type Date int64


func main() {

    handler := rest.ResourceHandler{
                EnableRelaxedContentType: true,
        }
    handler.SetRoutes(
        &rest.Route{"GET", "/task/list", GetAllTasks},
        &rest.Route{"POST", "/task/add", PostTask},
        &rest.Route{"GET", "/task/:taskdesc", GetTask},
		&rest.Route{"POST","/task/:taskdesc",UpdateTask},
		&rest.Route{"DELETE", "/task/:taskdesc", DeleteTask},
    )
    http.ListenAndServe(":8080", &handler)
}

type Task struct {
    Taskdesc string
    Date string
	Completed string
}

var store = map[string]*Task{}

func GetTask(w rest.ResponseWriter, r *rest.Request) {
    taskdesc := r.PathParam("taskdesc")
    task := store[taskdesc]
    if task == nil {
        rest.NotFound(w, r)
        return
    }
    w.WriteJson(&task)
}

func GetAllTasks(w rest.ResponseWriter, r *rest.Request) {
    tasks := make([]*Task, len(store))
    i := 0
    for _, task := range store {
        tasks[i] = task
        i++
    }
    w.WriteJson(&tasks)
}

func PostTask(w rest.ResponseWriter, r *rest.Request) {
    task := Task{}
    err := r.DecodeJsonPayload(&task)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    store[task.Taskdesc] = &task
    w.WriteJson(&task)
}

func UpdateTask(w rest.ResponseWriter, r *rest.Request) {
    taskdesc := r.PathParam("taskdesc")
	task := Task{}
    err := r.DecodeJsonPayload(&task)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
   
    store[taskdesc] = &task
    w.WriteJson(&task)
    
}
func DeleteTask(w rest.ResponseWriter, r *rest.Request) {
    taskdesc := r.PathParam("taskdesc")
    delete(store, taskdesc)
}
