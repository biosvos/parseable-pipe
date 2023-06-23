package main

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	sprintf := fmt.Sprintf("%v:%v", "admin", "admin")
	toString := base64.StdEncoding.EncodeToString([]byte(sprintf))
	require.Equal(t, "YWRtaW46YWRtaW4K", toString)
}
