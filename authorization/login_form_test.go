package authorization

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseJson(t *testing.T) {
	login := "123"
	password := "1234"

	json := fmt.Sprintf(`{ "body" : {"login" : "%s", "password" : "%s"}}`, login, password)

	form, err := ParseJSON(strings.NewReader(json))

	if err != nil {
		t.Fatalf("Failed on right json!")
	}

	is_login_correct := form.Body_.Login == login
	is_password_correct := form.Body_.Password == password

	if !is_login_correct || !is_password_correct {
		t.Fatalf(`Wrong parsing! Required: login: %s, password: %s.
				Get: login: %s, password: %s.`,
			login, password, form.Body_.Login, form.Body_.Password)
	}
}
