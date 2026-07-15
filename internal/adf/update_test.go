package adf

import "testing"

func TestVersionEqual(t *testing.T) {
	if !versionEqual("v0.8.0", "0.8.0") {
		t.Fatal("normalização falhou")
	}
	if versionEqual("0.8.0", "0.7.0") {
		t.Fatal("versões diferentes")
	}
}

func TestReleaseAssetName(t *testing.T) {
	name, err := releaseAssetName()
	if err != nil {
		t.Fatal(err)
	}
	if name == "" {
		t.Fatal("asset vazio")
	}
}
