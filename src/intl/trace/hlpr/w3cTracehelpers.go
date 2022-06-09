package hlpr

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	INVALID_VALUE_COUNT = "Invalid number of values"
)


func GenerateRandomBytes(n int) ([]byte, error) {
  b := make([]byte, n)
  if _, err := rand.Read(b); err != nil {
    return nil, err
  } else {
    return b, nil
  }
}

func GenerateTraceIdRaw() ([]byte, error) {
  if raw, err := GenerateRandomBytes(16); err != nil {
    return nil, fmt.Errorf("Failed to generate bytes %w", err)
  } else {
    return raw, nil
  }
}

func GenerateTraceId() (string, error) {
  if raw, err := GenerateTraceIdRaw(); err != nil {
    return "", fmt.Errorf("Failed to generate trace Id %w", err)
  } else {
    return base64.StdEncoding.EncodeToString(raw), nil
  }
}

func GenerateParentIdRaw() ([]byte, error) {
  if raw, err := GenerateRandomBytes(8); err != nil {
    return nil, fmt.Errorf("Failed to generate butes %w", err)
  } else {
    return raw, nil
  }
}
// Counting on go compiler to inline these plsplspls :)
func GenerateParentId() (string, error) {
  if raw, err := GenerateParentIdRaw(); err != nil {
    return "", fmt.Errorf("Failed to generate parent Id %w", err)
  } else {
    return base64.StdEncoding.EncodeToString(raw), nil
  }
}

func GenerateNewTraceparent(sampled bool) (string, error) {
  var flag string
  if sampled {
    flag = "01"
  } else {
    flag = "00"
  }
  tid, err := GenerateTraceId()
  if err != nil {
    return "", fmt.Errorf("Failed to generate traceId %w", err)
  }
  pid, err := GenerateParentId()
  if err != nil {
    return "", fmt.Errorf("Failed to generate parentId %w", err)
  }
  return fmt.Sprintf(
    "00-%s-%s-%s",
    tid,
    pid,
    flag,
  ), nil
}

func ParseTraceparentRaw(
  traceparent string,
) ([]byte, []byte, []byte, []byte, error) {
  values := strings.Split(traceparent, "-")
  if len(values) != 4 {
  	return nil, nil, nil, nil, fmt.Errorf(INVALID_VALUE_COUNT)
  }
 	
 	vers, err := base64.StdEncoding.DecodeString(values[0])
 	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to decode version %w", err)
 	}
	tid, err := base64.StdEncoding.DecodeString(values[1])
 	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to decode traceId %w", err)
 	}
	pid, err := base64.StdEncoding.DecodeString(values[2])
 	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to decode parentId %w", err)
 	}
	flg, err := base64.StdEncoding.DecodeString(values[3])
 	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to decode flag %w", err)
 	}
 	return vers, tid, pid, flg, nil
}

func ValidateTraceId(traceId string) error {
  return nil
}

func ValidateParentId(parentId string) error {
  return nil
}
