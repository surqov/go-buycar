package handler_test

func TestLogin(t *testing.T) {
    app := fiber.New()
    app.Post("/login", handler.Login)

    tests := []struct {
        name       string
        body       string
        wantStatus int
    }{
        {"valid", `{"email":"a@b.com","password":"123456"}`, http.StatusOK},
        {"invalid json", `{email:"a@b.com"}`, http.StatusBadRequest},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tc.body))
            req.Header.Set("Content-Type", "application/json")

            resp, err := app.Test(req)
            if err != nil {
                t.Fatalf("Request failed: %v", err)
            }
            if resp.StatusCode != tc.wantStatus {
                t.Errorf("Expected status %d, got %d", tc.wantStatus, resp.StatusCode)
            }
        })
    }
}
