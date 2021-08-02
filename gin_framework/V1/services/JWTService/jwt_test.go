package JWTService

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const KEY = "super_secret_Key"
const ValidToken = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im0uZGVoZ2hhbnBvdXJAZ21haWwuY29tIiwibGFzdG5hbWUiOiJkZWhnaGFucG91ciIsIm5hbWUiOiJtb2hhbW1hZCJ9.YCSi04IrIE5k8RRxBPv4XHtYHcspRCPzx2_eFQhG9Rb_vZ0s9ObUSy7q_mOzn7HTmBXtTGn9N75dQ1Wjih017A"

func TestGenerator(t *testing.T) {
	tests := []struct {
		name           string
		input          map[string]interface{}
		testDependency bool
		expectedToken  string
		expectedErr    error
	}{
		{"ok", map[string]interface{}{"email": "m.dehghanpour@gmail.com", "lastname": "dehghanpour", "name": "mohammad"}, false, ValidToken, nil},
		{"jwt dependency Err", map[string]interface{}{"email": "m.dehghanpour@gmail.com", "lastname": "dehghanpour", "name": "mohammad"}, true, "", errors.New("dependency Err")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jwt := NewJWTService(KEY)
			if test.testDependency {
				jwt.SetDependencyTest()
			}
			generator, err := jwt.Generator(test.input)

			assert.Equal(t, generator, test.expectedToken)
			if err == nil {
				assert.Nil(t, test.expectedErr)
			} else {
				assert.EqualError(t, err, test.expectedErr.Error())
			}
		})

	}
}

func TestValidator(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		dependencyTest bool
		expectedErr    error
		expectedOk     bool
		expectedData   map[string]interface{}
	}{
		{"ok", ValidToken, false, nil, true, map[string]interface{}{"email": "m.dehghanpour@gmail.com", "lastname": "dehghanpour", "name": "mohammad"}},
		{"invalid", "invalidToken", false, errors.New("error"), false, make(map[string]interface{})},
		{"dependency Err", ValidToken, true, errors.New("error"), false, make(map[string]interface{})},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jwt := NewJWTService(KEY)
			if test.dependencyTest {
				jwt.SetDependencyTest()
			}
			ok, data, err := jwt.Validator(test.input)
			if err == nil {
				assert.Nil(t, test.expectedErr)
			} else {
				assert.NotNil(t, test.expectedErr)
			}
			assert.Equal(t, test.expectedOk, ok)
			assert.Equal(t, test.expectedData, data)
		})
	}

}
