package registration

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseJson(t *testing.T) {
	login := "123"
	password := "1234"
	user_type := "simple_user"

	json := fmt.Sprintf(`{ "body" : {"login" : "%s", "password" : "%s", "user_type" : "%s"}}`, login, password, user_type)

	form, err := ParseJSON(strings.NewReader(json))

	if err != nil {
		t.Fatalf("Failed on right json!")
	}

	is_login_correct := form.Body_.Login == login
	is_password_correct := form.Body_.Password == password
	is_user_type_correct := form.Body_.UserType == user_type

	if !is_login_correct || !is_password_correct || !is_user_type_correct {
		t.Fatalf(`Wrong parsing! Required: login: %s, password: %s,user_type: %s.
				Get: login: %s, password: %s, user_type: %s.`,
			login, password, user_type, form.Body_.Login, form.Body_.Password, form.Body_.UserType)
	}
}
