package main //golang-tdd

import (
	"testing"
	"bytes"
	"image/png"
	"errors"
)


func TestGenerateQRCodeGeneratesPNG(t *testing.T) {

	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368")
	
	if buffer.Len() == 0 {
		t.Errorf("Generated QR code lenght is zero")
	}
	
	_, err := png.Decode(buffer)
	
	if err != nil {
		t.Errorf("Generated QR Code is not a PNG: %s", err)
	}
}

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w,"555-2368")
	
	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error not propagated correctly, got %v", err)
	}  
}