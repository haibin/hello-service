package tests

import (
	"bytes"
	"encoding/json"
	"github.com/haibin/hello-service/business/validate"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/haibin/hello-service/app/hello-api/handlers"
	"github.com/haibin/hello-service/business/data/user"
	"github.com/haibin/hello-service/business/tests"
)

// UserTests holds methods for each user subtest. This type allows passing
// dependencies for tests while still providing a convenient syntax when
// subtests are registered.
type UserTests struct {
	app        http.Handler
	kid        string
	userToken  string
	adminToken string
}

// TestUsers is the entry point for testing user management functions.
func TestUsers(t *testing.T) {
	test := tests.NewIntegration(t)
	t.Cleanup(test.Teardown)

	shutdown := make(chan os.Signal, 1)
	tests := UserTests{
		app: handlers.API(shutdown, test.Log, test.DB),
		kid: test.KID,
		//userToken:  test.Token("user@example.com", "gophers"),
		//adminToken: test.Token("admin@example.com", "gophers"),
	}

	t.Run("postUser400", tests.postUser400)
}

// postUser400 validates a user can't be created with the endpoint
// unless a valid user document is submitted.
func (ut *UserTests) postUser400(t *testing.T) {
	body, err := json.Marshal(&user.NewUser{})
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	r.Header.Set("Authorization", "Bearer "+ut.adminToken)
	ut.app.ServeHTTP(w, r)

	t.Log("Given the need to validate a new user can't be created with an invalid document.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen using an incomplete user value.", testID)
		{
			if w.Code != http.StatusBadRequest {
				t.Fatalf("\t%s\tTest %d:\tShould receive a status code of 400 for the response : %v", tests.Failed, testID, w.Code)
			}
			t.Logf("\t%s\tTest %d:\tShould receive a status code of 400 for the response.", tests.Success, testID)

			var got validate.ErrorResponse
			if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to unmarshal the response to an error type : %v", tests.Failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to unmarshal the response to an error type.", tests.Success, testID)

			fields := validate.FieldErrors{
				{Field: "name", Error: "name is a required field"},
				{Field: "email", Error: "email is a required field"},
				{Field: "roles", Error: "roles is a required field"},
				{Field: "password", Error: "password is a required field"},
			}
			exp := validate.ErrorResponse{
				Error:  "data validation error",
				Fields: fields.Error(),
			}

			// We can't rely on the order of the field errors so they have to be
			// sorted. Tell the cmp package how to sort them.
			sorter := cmpopts.SortSlices(func(a, b validate.FieldError) bool {
				return a.Field < b.Field
			})

			if diff := cmp.Diff(got, exp, sorter); diff != "" {
				t.Fatalf("\t%s\tTest %d:\tShould get the expected result. Diff:\n%s", tests.Failed, testID, diff)
			}
			t.Logf("\t%s\tTest %d:\tShould get the expected result.", tests.Success, testID)
		}
	}
}
