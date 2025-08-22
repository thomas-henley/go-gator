package config

import "testing"

func TestConfigRead(t *testing.T) {
	expected_db := "postgres://example"
	config, err := Read()
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if config.DbUrl != expected_db {
		t.Errorf("wanted %v got %v", expected_db, config.DbUrl)
	}
}

func TestConfigSetUser(t *testing.T) {
	expected_username := "Test123"
	config := Config{}
	config.SetUser(expected_username)

	actual_config, err := Read()
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if actual_config.CurrentUserName != expected_username {
		t.Errorf("wanted %v got %v", expected_username, actual_config.CurrentUserName)
	}
}
