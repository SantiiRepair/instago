package ig

import (
	"net/url"
	"strings"
)

// ChallengeGET performs an unsigned GET private request for a challenge path (strips optional api/v1/ prefix).
func (c *Client) ChallengeGET(apiPath string, params url.Values) (map[string]any, error) {
	ep := strings.TrimPrefix(apiPath, "/")
	if after, ok :=strings.CutPrefix(ep, "api/v1/"); ok  {
		ep = after
	}
	return c.PrivateRequest(PrivateRequestOpts{
		Endpoint:      ep,
		Params:        params,
		WithSignature: false,
	})
}

// ChallengePOST posts signed form data to a challenge API path.
func (c *Client) ChallengePOST(apiPath string, data map[string]any) (map[string]any, error) {
	ep := strings.TrimPrefix(apiPath, "/")
	if after, ok :=strings.CutPrefix(ep, "api/v1/"); ok  {
		ep = after
	}
	return c.PrivateRequest(PrivateRequestOpts{
		Endpoint:      ep,
		Data:          data,
		WithSignature: true,
	})
}
