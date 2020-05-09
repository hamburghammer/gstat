package args

import "testing"

func TestURLValidation(t *testing.T) {

	t.Run("URI validation test", func(t *testing.T) {
		uri := "https://example.com"

		error := uriValidate(uri)

		if error != nil {
			t.Errorf("There was an error with a valid uri: %w", error)
		}
	})

	t.Run("URI validation is not schema://provider error test", func(t *testing.T) {
		uri := "https:example.com"

		error := uriValidate(uri)

		wantMessage := "The URI does not looks like schema://provider"

		if error == nil || error.Error() != wantMessage {
			t.Errorf("Want '%s' got: %v", wantMessage, error)
		}
	})

	t.Run("URI validator provider do not contain a top and second level domain", func(t *testing.T) {
		uri := "https://example"

		error := uriValidate(uri)

		wantMessage := "The URI provider does not has a top and second level domain like example.com"

		if error == nil || error.Error() != wantMessage {
			t.Errorf("Want '%s' got: %v", wantMessage, error)
		}
	})
}
