package config

import "testing"

func TestGoodConfig(t *testing.T) {
    _, err := EnsureConfig("profile", "bucket", "queue")

    if err != nil {
        t.Errorf("`EnsureConfig` failed with error %s: ", err)
    }
}
