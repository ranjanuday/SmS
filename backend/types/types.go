package types

import "time"

// =========================
// SMS STORE INTERFACE
// =========================

	type SMSStore interface {
		CreateUser(User) error
		GetUserByEmail(string) (*User, error)
		GetUserByID(int) (*User, error)
	}

// =========================
// USER TYPES
// =========================

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// =========================
// STUDENT TYPES
// =========================

type Student struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userID"`
	StudentID  string    `json:"studentID"`
	Roll       int       `json:"roll"`
	Name       string    `json:"name"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"createdAt"`
}

type CreateStudentPayload struct {
	StudentID  string `json:"studentID" validate:"required"`
	Roll       int    `json:"roll" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Department string `json:"department" validate:"required"`
}

// =========================
// MARKS TYPES
// =========================

type Marks struct {
	ID        int       `json:"id"`
	StudentID int       `json:"studentID"`
	Subject   string    `json:"subject"`
	Marks     float64   `json:"marks"`
	Grade     string    `json:"grade"`
	CreatedAt time.Time `json:"createdAt"`
}

type AddMarksPayload struct {
	StudentID int     `json:"studentID" validate:"required"`
	Subject   string  `json:"subject" validate:"required"`
	Marks     float64 `json:"marks" validate:"required"`
	Grade     string  `json:"grade"`
}

// =========================
// NOTICE TYPES
// =========================

type Notice struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateNoticePayload struct {
	Title   string `json:"title" validate:"required"`
	Message string `json:"message" validate:"required"`
}

// =========================
// PLACEMENT TYPES
// =========================

type Placement struct {
	ID          int       `json:"id"`
	StudentID   int       `json:"studentID"`
	CompanyName string    `json:"companyName"`
	Role        string    `json:"role"`
	LPA         float64   `json:"lpa"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreatePlacementPayload struct {
	StudentID   int     `json:"studentID" validate:"required"`
	CompanyName string  `json:"companyName" validate:"required"`
	Role        string  `json:"role" validate:"required"`
	LPA         float64 `json:"lpa" validate:"required"`
}

// =========================
// ASSIGNMENT TYPES
// =========================

type Assignment struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     string    `json:"dueDate"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateAssignmentPayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	DueDate     string `json:"dueDate"`
}

// =========================
// QUERY TYPES
// =========================

type Query struct {
	ID        int       `json:"id"`
	StudentID int       `json:"studentID"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type RaiseQueryPayload struct {
	StudentID int    `json:"studentID" validate:"required"`
	Subject   string `json:"subject" validate:"required"`
	Message   string `json:"message" validate:"required"`
}