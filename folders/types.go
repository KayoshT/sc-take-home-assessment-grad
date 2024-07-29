package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

// Request for fetching folders with pagination.
type FolderPaginationRequest struct {
	OrgID   uuid.UUID
	PerPage int
	Token   string // Token for pagination
}

// Response for fetching folders with pagination.
type PaginatedFolderResponse struct {
	Folders []*Folder
	Token   string // Token for the next set of results
}
