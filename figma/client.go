package figma

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const DefaultMCPEndpoint = "http://127.0.0.1:3845/mcp"

// MCPClient calls the Figma desktop app's local MCP server.
// Enable it in the Figma desktop app via Dev Mode (Shift+D) → Enable MCP server.
type MCPClient struct {
	Endpoint   string
	HTTPClient *http.Client
	sessionID  string
	nextID     int
}

func NewMCPClient(endpoint string) *MCPClient {
	if endpoint == "" {
		endpoint = DefaultMCPEndpoint
	}
	return &MCPClient{
		Endpoint:   endpoint,
		HTTPClient: &http.Client{},
		nextID:     1,
	}
}

type mcpRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
}

// mcpNotification has no ID field — used for one-way messages like notifications/initialized.
type mcpNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
}

type mcpRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *mcpRPCError    `json:"error,omitempty"`
}

type mcpRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type toolCallResult struct {
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	IsError bool `json:"isError"`
}

// ListTools returns the raw JSON tool list from the MCP server.
func (c *MCPClient) ListTools() (string, error) {
	if err := c.handshake(); err != nil {
		return "", err
	}
	result, err := c.call("tools/list", nil)
	if err != nil {
		return "", err
	}
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, result, "", "  "); err != nil {
		return string(result), nil
	}
	return pretty.String(), nil
}

// RunScript performs the MCP handshake then executes the Figma Plugin API
// JavaScript via the use_figma tool.
func (c *MCPClient) RunScript(script string) error {
	if err := c.handshake(); err != nil {
		return err
	}

	result, err := c.call("tools/call", map[string]any{
		"name":      "use_figma",
		"arguments": map[string]any{"code": script},
	})
	if err != nil {
		return err
	}

	var tcr toolCallResult
	if err := json.Unmarshal(result, &tcr); err != nil {
		return fmt.Errorf("parse result: %w", err)
	}
	if tcr.IsError {
		var msgs []string
		for _, item := range tcr.Content {
			msgs = append(msgs, item.Text)
		}
		return fmt.Errorf("figma plugin error: %s", strings.Join(msgs, "; "))
	}
	return nil
}

// handshake sends initialize and then the notifications/initialized notification.
func (c *MCPClient) handshake() error {
	if _, err := c.call("initialize", map[string]any{
		"protocolVersion": "2024-11-05",
		"capabilities":    map[string]any{},
		"clientInfo":      map[string]any{"name": "go-design-tokens", "version": "1.0.0"},
	}); err != nil {
		return fmt.Errorf("initialize: %w", err)
	}
	return c.notify("notifications/initialized", nil)
}

// notify sends a JSON-RPC notification (no ID, no response expected).
func (c *MCPClient) notify(method string, params any) error {
	data, err := json.Marshal(mcpNotification{JSONRPC: "2.0", Method: method, Params: params})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.Endpoint, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.sessionID != "" {
		req.Header.Set("Mcp-Session-Id", c.sessionID)
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func (c *MCPClient) call(method string, params any) (json.RawMessage, error) {
	id := c.nextID
	c.nextID++

	data, err := json.Marshal(mcpRPCRequest{JSONRPC: "2.0", ID: id, Method: method, Params: params})
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.Endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json, text/event-stream")
	if c.sessionID != "" {
		httpReq.Header.Set("Mcp-Session-Id", c.sessionID)
	}

	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("cannot reach MCP server at %s: %w\n(Is the Figma desktop app open with Dev Mode enabled?)", c.Endpoint, err)
	}
	defer httpResp.Body.Close()

	// Capture session ID for subsequent requests.
	if sid := httpResp.Header.Get("Mcp-Session-Id"); sid != "" {
		c.sessionID = sid
	}

	if strings.Contains(httpResp.Header.Get("Content-Type"), "text/event-stream") {
		return readSSEResponse(httpResp.Body, id)
	}

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	var resp mcpRPCResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("MCP error %d: %s", resp.Error.Code, resp.Error.Message)
	}
	return resp.Result, nil
}

func readSSEResponse(r io.Reader, targetID int) (json.RawMessage, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		data := strings.TrimPrefix(line, "data: ")
		if data == "[DONE]" {
			break
		}
		var resp mcpRPCResponse
		if err := json.Unmarshal([]byte(data), &resp); err != nil {
			continue
		}
		if resp.ID != targetID {
			continue
		}
		if resp.Error != nil {
			return nil, fmt.Errorf("MCP error %d: %s", resp.Error.Code, resp.Error.Message)
		}
		return resp.Result, nil
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("SSE read: %w", err)
	}
	return nil, fmt.Errorf("no response received for request ID %d", targetID)
}
