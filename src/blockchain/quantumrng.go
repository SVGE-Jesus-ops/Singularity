// Package quantumrng provides functions to fetch truly random numbers from the ANU Quantum Random Number Generator
package quantumrng

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// QRNGResponse matches the format returned by the ANU Quantum Random Numbers JSON API
type QRNGResponse struct {
	Type    string   `json:"type"`
	Length  int      `json:"length"`
	Size    int      `json:"size"`
	Data    []string `json:"data"`
	Success bool     `json:"success"`
}

// FetchQuantumHex32 fetches quantum random hex strings (32 bits each) from ANU QRNG
// Returns a slice of hex strings, or an error
func FetchQuantumHex32(count int) ([]string, error) {
	url := fmt.Sprintf("https://qrng.anu.edu.au/API/jsonI.php?length=%d&type=hex16&size=32", count)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch quantum random numbers: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var qrngResp QRNGResponse
	if err := json.Unmarshal(body, &qrngResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	
	if !qrngResp.Success {
		return nil, fmt.Errorf("QRNG API request failed")
	}
	
	return qrngResp.Data, nil
}

// FetchQuantumUint8 fetches quantum random uint8 values from ANU QRNG
func FetchQuantumUint8(count int) ([]uint8, error) {
	url := fmt.Sprintf("https://qrng.anu.edu.au/API/jsonI.php?length=%d&type=uint8", count)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch quantum random numbers: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result struct {
		Type    string  `json:"type"`
		Length  int     `json:"length"`
		Data    []uint8 `json:"data"`
		Success bool    `json:"success"`
	}
	
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	
	if !result.Success {
		return nil, fmt.Errorf("QRNG API request failed")
	}
	
	return result.Data, nil
}
