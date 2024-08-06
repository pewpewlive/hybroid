package lsp

import (
	"context"

	"github.com/sourcegraph/jsonrpc2"
)

func (h *langHandler) handleShutdown(_ context.Context, conn *jsonrpc2.Conn, _ *jsonrpc2.Request) (result any, err error) {
	close(h.request)
	return nil, conn.Close()
}
