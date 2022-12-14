package models

import "time"

// 👈 Cloumn is used to sent data to update board-colums from database

type CloumnUpdate struct {
	Columns []struct {
		ID      string   `json:"id" bson:"id"`
		Name    string   `json:"name" bson:"name"`
		CardIds []string `json:"cardIds" bson:"cardIds"`
	} `json:"columns" bson:"columns"`
}

// 👈 ColumnOrder is used to sent data to update board-column-order from database

type ColumnOrderUpdate struct {
	ColumnOrder []string `json:"columnOrder" bson:"columnOrder"`
}

// 👈 Board is used to sent data to rename columns board from database
type ColumnRename struct {
	ColumnID string `json:"id" bson:"id" binding:"required"`
	NewName  string `json:"name" bson:"name" binding:"required"`
}

// 👈 Kanban is used to sent data to database

type CloumnResponse struct {
	Columns struct {
		ID      string   `json:"id" bson:"id"`
		Name    string   `json:"name" bson:"name"`
		CardIds []string `json:"cardIds" bson:"cardIds"`
	} `json:"columns" bson:"columns"`
}

// 👈 Kanban is used to sent data to add card from database

type AddCardInput struct {
	Card struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Assignee    []struct {
			ID     string `json:"id"`
			Avatar string `json:"avatar"`
			Name   string `json:"name"`
		} `json:"assignee"`
		Due         []int64       `json:"due"`
		Attachments []interface{} `json:"attachments"`
		Comments    []struct {
			ID          string    `json:"id"`
			Avatar      string    `json:"avatar"`
			Name        string    `json:"name"`
			CreatedAt   time.Time `json:"createdAt"`
			MessageType string    `json:"messageType"`
			Message     string    `json:"message"`
		} `json:"comments"`
		Completed bool `json:"completed"`
	} `json:"card"`
	ColumnID string `json:"columnId" bson:"columnId" binding:"required"`
}

// 👈 DBResponse is used to sent data to database

type KanbanInput struct {
	Board struct {
		Cards []struct {
			ID          string `json:"id" bson:"id"`
			Name        string `json:"name" bson:"name"`
			Description string `json:"description" bson:"description"`
			Assignee    []struct {
				ID     string `json:"id" bson:"id"`
				Avatar string `json:"avatar" bson:"avatar"`
				Name   string `json:"name" bson:"name"`
			} `json:"assignee" bson:"assignee"`
			Due         []int64       `json:"due" bson:"due"`
			Attachments []interface{} `json:"attachments" bson:"attachments"`
			Comments    []struct {
				ID          string    `json:"id" bson:"id"`
				Avatar      string    `json:"avatar" bson:"avatar"`
				Name        string    `json:"name" bson:"name"`
				CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
				MessageType string    `json:"messageType" bson:"messageType"`
				Message     string    `json:"message" bson:"message"`
			} `json:"comments" bson:"comments"`
			Completed bool `json:"completed" bson:"completed"`
		} `json:"cards" bson:"cards"`
		Columns []struct {
			ID      string   `json:"id" bson:"id"`
			Name    string   `json:"name" bson:"name"`
			CardIds []string `json:"cardIds" bson:"cardIds"`
		} `json:"columns" bson:"columns"`
		ColumnOrder []string `json:"columnOrder" bson:"columnOrder"`
	} `json:"board" bson:"board"`
}

// 👈 Kanban is used to sent data to database markdone task

type MarkDoneInput struct {
	CardID    string `json:"id" bson:"id" binding:"required"`
	Completed bool   `json:"completed" bson:"completed"`
}

// 👈  Kanban is used to sent data to database markdone task
type DeleteCardInput struct {
	Cards struct {
		ID          string `json:"id" bson:"id" binding:"required"`
		Name        string `json:"name" bson:"name"`
		Description string `json:"description" bson:"description"`
		Assignee    []struct {
			ID     string `json:"id" bson:"id"`
			Avatar string `json:"avatar" bson:"avatar"`
			Name   string `json:"name" bson:"name"`
		} `json:"assignee" bson:"assignee"`
		Due         []int64       `json:"due" bson:"due"`
		Attachments []interface{} `json:"attachments" bson:"attachments"`
		Comments    []struct {
			ID          string    `json:"id" bson:"id"`
			Avatar      string    `json:"avatar" bson:"avatar"`
			Name        string    `json:"name" bson:"name"`
			CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
			MessageType string    `json:"messageType" bson:"messageType"`
			Message     string    `json:"message" bson:"message"`
		} `json:"comments" bson:"comments"`
		Completed bool `json:"completed" bson:"completed"`
	} `json:"card" bson:"card"`
	ColumnID string `json:"columnId" bson:"columnId"`
}

// 👈 Kanban is used to create new columns and sent data to database

type AddNewColumnInput struct {
	Column struct {
		ID      string   `json:"id" bson:"id"`
		Name    string   `json:"name" bson:"name"`
		CardIds []string `json:"cardIds" bson:"cardIds"`
	} `json:"column" bson:"column"`
}

// 👈 Kanban is used to Delete columns and sent data to database
type DeleteColumnInput struct {
	Column struct {
		ID string `json:"id" bson:"id"`
	} `json:"column" bson:"column"`
}

// 👈 Kanban is used to Delete columns and sent data to database
type AddAttachmentInput struct {
	Data struct {
		ID  string `json:"id" bson:"id"`
		URL string `json:"url" bson:"url"`
	}
}

// 👈 Kanban is used to rename card and sent data to database
type RenameCard struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// 👈 Kanban is used to change description card and sent data to database
type ChangeDescriptionCard struct {
	ID             string `json:"id" bson:"id"`
	NewDescription string `json:"description" bson:"description"`
}

type NewComment struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Avatar      string `json:"avatar" bson:"avatar"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	MessageType string `json:"messageType" bson:"messageType"`
	Message     string `repeats:"message" bson:"message"`
}

type AddNewComment struct {
	CardID     string `json:"id" bson:"id"`
	NewComment `json:"comment" bson:"comment"`
}
