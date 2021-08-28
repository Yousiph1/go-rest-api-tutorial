package data

import(
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

  "toonji.com/m/models"
)

type TaskRepository struct {
   C *mgo.Collection
}

func (r *TaskRepository) Create(task *models.Task) error {
   object_id := bson.NewObjectId()
   task.ID = object_id
   task.CreatedOn = time.Now()
   task.Status = "Created"
   err := r.C.Insert(task)
   return err
}

func (r *TaskRepository) Update(task *models.Task) error {

   err := r.C.Update(bson.M{"_id":task.ID},
                     bson.M{"$set":bson.M {
                           "name": task.Name,
                            "description": task.Description,
                            "due": task.Due,
                            "status": task.Status,
                            "tags": task.Tags,
                        }})
  return err
}

func (r *TaskRepository) Delete(id string) error {
    _id := bson.ObjectIdHex(id)
    err := r.C.Remove(bson.M{"_id":_id})
    return err
}

func (r *TaskRepository) GetAll() []models.Task{
     tasks := []models.Task{}
     iter := r.C.Find(nil).Iter()
     result := models.Task{}
     for iter.Next(&result) {
       tasks = append(tasks,result)
     }
     return tasks
}

func (r *TaskRepository) GetById(id string) (task models.Task, err error) {
    object_id := bson.ObjectIdHex(id)
    err = r.C.FindId(object_id).One(&task)
    if err != nil {
      task = models.Task{}
      return
    }
  return
}

func (r *TaskRepository) GetByUser(user string) []models.Task {
    var tasks []models.Task
    iter := r.C.Find(bson.M{"createdby": user}).Iter()
    result := models.Task{}
    for iter.Next(&result) {
       tasks = append(tasks, result)
    }
    return tasks
}
