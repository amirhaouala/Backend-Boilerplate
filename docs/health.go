//lint:file-ignore U1000 ignore go-swagger template
package docs

import "github.com/supabase/auth/internal/api"

// swagger:route GET /scratchpad scratchpad scratchpad
// The scratchpadcheck endpoint for gotrue. Returns the current gotrue version.
// responses:
//   200: scratchpadCheckResponse

// swagger:response scratchpadCheckResponse
type scratchpadCheckResponseWrapper struct {
	// in:body
	Body api.ScratchpadResponse
}
