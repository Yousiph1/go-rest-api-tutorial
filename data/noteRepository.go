package data

import (
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "toonji.com/m/models"
)


type NoteRepository struct {
    C *mgo.Collection
}

func (n *NoteRepository) Create(note *models.TaskNote) error {
  id := bson.NewObjectId()
  note.ID = id
  note.CreatedOn = time.Now()
  err := n.C.Insert(note)
  return err
}

func (n *NoteRepository) GetNotes() []models.TaskNote {
    notes := []models.TaskNote{}
    iter := n.C.Find(nil).Iter()
    note := models.TaskNote{}
    for iter.Next(&note) {
      notes = append(notes,note)
    }
    return notes
}

func (n *NoteRepository) GeNote(id string) (note models.TaskNote, err error) {
    objectID := bson.ObjectIdHex(id)
    note = models.TaskNote{}
    err  = n.C.Find(bson.M{"_id":objectID}).One(&note)
    if err != nil {
       note = models.TaskNote{}
       return
    }
    return note,nil
}

func (n *NoteRepository) Update(id string, note models.TaskNote) error {
     ID := bson.ObjectIdHex(id)
     err := n.C.Update(bson.M{"_id":ID},bson.M{"$set":bson.M{
                         "taskid":note.TaskID,
                         "description": note.Description,
                      }})
    return err
}

func (n *NoteRepository) Delete(id string) error {
   ID := bson.ObjectIdHex(id)
   err := n.C.Remove(bson.M{"_id":ID})
   return err
}

func (n *NoteRepository) GetNotesByTask(id string) []models.TaskNote {
  notes := []models.TaskNote{}
  taskID := bson.ObjectId(id)
  iter := n.C.Find(bson.M{"taskid":taskID}).Iter()
  note := models.TaskNote{}
  for iter.Next(&note) {
    notes = append(notes,note)
  }
  return notes
}
