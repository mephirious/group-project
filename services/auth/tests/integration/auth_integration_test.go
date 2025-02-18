package integration

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	db "github.com/mephirious/group-project/services/auth/db/mongo"
	"github.com/mephirious/group-project/services/auth/db/mongo/repository"
	"github.com/mephirious/group-project/services/auth/domain"
	"github.com/mephirious/group-project/services/auth/service"
)

var dbInstance *repository.DB

func TestMain(m *testing.M) {
	// Load environment variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("⚠️ Warning: Could not load .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ✅ Use correct database connection function from connection.go
	var err error
	dbInstance, err = db.NewDB(ctx)
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	// Run tests
	code := m.Run()

	// Cleanup database connection
	if err := dbInstance.Close(ctx); err != nil {
		log.Printf("⚠️ Database close error: %v", err)
	}

	os.Exit(code)
}

// ✅ **Test: Database Connection**
func TestDatabaseConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := dbInstance.Client.Ping(ctx, nil); err != nil {
		t.Fatalf("❌ Database ping failed: %v", err)
	}

	log.Println("✅ Database connected succesfully")
}

// ✅ **Test: Create and Fetch Customer**
func TestCreateAndFetchCustomer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := repository.CreateCustomerInput{
		Email:    "test@example.com",
		Password: "securepassword",
		Role:     "user",
	}

	_, err := dbInstance.CreateCustomer(ctx, input)
	if err != nil {
		t.Fatalf("❌ Failed to create customer: %v", err)
	}

	// ✅ Convert string to pointer
	email := input.Email
	fetchedCustomer, err := dbInstance.GetCustomersOne(ctx, repository.GetCustomersInput{Email: &email})
	if err != nil || fetchedCustomer == nil {
		t.Fatalf("❌ Failed to fetch customer: %v", err)
	}

	// Validate email
	if fetchedCustomer.Email != input.Email {
		t.Errorf("❌ Expected email %s, got %s", input.Email, fetchedCustomer.Email)
	}

	log.Println("✅ Customer created and retrieved successfully:", fetchedCustomer.ID)
}

// ✅ **Test: Create and Fetch Session**
func TestCreateAndFetchSession(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	email := "test@example.com"
	// Fetch an existing user
	existingCustomer, err := dbInstance.GetCustomersOne(ctx, repository.GetCustomersInput{Email: &email})
	if err != nil {
		t.Fatalf("❌ No customer found: %v", err)
	}

	// Use existing CreateSession function
	sessionInput := repository.CreateSessionInput{
		UserID:    existingCustomer.ID,
		UserAgent: "Test-Agent",
	}

	// ✅ Create a session and check for errors
	_, err = dbInstance.CreateSession(ctx, sessionInput) // ⚠️ Removed unused variable
	if err != nil {
		t.Fatalf("❌ Failed to create session: %v", err)
	}

	// Fetch the session
	fetchedSession, err := dbInstance.GetSessionOne(ctx, repository.GetSessionsInput{UserID: &existingCustomer.ID})
	if err != nil || fetchedSession == nil {
		t.Fatalf("❌ Failed to fetch session: %v", err)
	}

	log.Println("✅ Session created and retrieved successfully:", fetchedSession.ID)
}

func TestLogout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ✅ Fetch existing user
	email := "test@example.com"
	existingCustomer, err := dbInstance.GetCustomersOne(ctx, repository.GetCustomersInput{Email: &email})
	if err != nil {
		t.Fatalf("❌ No customer found: %v", err)
	}

	// ✅ Create a session to simulate login
	sessionInput := repository.CreateSessionInput{
		UserID:    existingCustomer.ID,
		UserAgent: "Test-Agent",
	}
	session, err := dbInstance.CreateSession(ctx, sessionInput)
	if err != nil {
		t.Fatalf("❌ Failed to create session: %v", err)
	}

	log.Printf("🔍 Created session ID: %s\n", session.ID)

	// ✅ Delete session
	err = dbInstance.DeleteSession(ctx, session.ID)
	if err != nil {
		t.Fatalf("❌ Logout failed: %v", err)
	}

	// ✅ Fetch session again
	fetchedSession, err := dbInstance.GetSessionOne(ctx, repository.GetSessionsInput{ID: &session.ID})
	if err == nil && fetchedSession != nil {
		t.Fatalf("❌ Session was not deleted after logout: %s", fetchedSession.ID)
	}

	log.Println("✅ Logout test passed: Session successfully deleted")
}

func TestRefreshToken(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ✅ Initialize auth service
	authService := service.NewAuthService(dbInstance)

	// ✅ Register the user (to ensure the password is hashed)
	registerInput := domain.RegisterInput{
		Email:     "test1@example.com",
		Password:  "securepassword", // ✅ This will be hashed by Register()
		UserAgent: "Test-Agent",
	}
	_, err := authService.Register(ctx, registerInput)
	if err != nil {
		t.Fatalf("❌ Registration failed: %v", err)
	}

	// ✅ Use Login() to get the correct refresh token
	loginInput := domain.LoginInput{
		Email:     "test1@example.com",
		Password:  "securepassword",
		UserAgent: "Test-Agent",
	}
	loginResponse, err := authService.Login(ctx, loginInput)
	if err != nil {
		t.Fatalf("❌ Login failed: %v", err)
	}

	// ✅ Now use the correct refresh token
	refreshToken := loginResponse.RefreshToken
	log.Printf("🔍 Refresh Token BEFORE: %s\n", refreshToken)

	// ✅ Use Refresh Token via the Service Layer
	refreshInput := domain.RefreshInput{
		RefreshToken: refreshToken,
	}
	response, err := authService.RefreshUserAccessToken(ctx, refreshInput)
	if err != nil {
		t.Fatalf("❌ Failed to refresh token: %v", err)
	}

	// ✅ Debug log for the generated access token
	log.Printf("✅ Generated New Access Token: %s\n", response.AccessToken)

	// ✅ Ensure a new access token is returned
	if response.AccessToken == "" {
		t.Fatalf("❌ Refresh token did not generate a new access token")
	}

	log.Println("✅ Refresh token test passed: New access token generated successfully")
}
