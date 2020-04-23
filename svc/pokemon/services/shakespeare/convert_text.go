package shakespeare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type convertTextResponse struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Content struct {
		Translated  string `json:"translated"`
		Text        string `json:"text"`
		Translation string `json:"translation"`
	} `json:"contents"`
}

func (s *service) ConvertText(ctx context.Context, text string) (string, error) {
	// Ensure our request doesn't hang indefinitely
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	form := url.Values{}
	form.Set("text", text)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprint(baseURL, "/translate/shakespeare"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Shakespeare returned: %v", resp.Status)
	}

	translation := &convertTextResponse{}
	err = json.NewDecoder(resp.Body).Decode(translation)
	if err != nil {
		return "", err
	}

	return translation.Content.Translated, nil
}
