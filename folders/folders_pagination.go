package folders

import (
	"encoding/base64"

	"github.com/gofrs/uuid"
)

/*
Short Explanation:
The GetPaginatedFolders function gets a set number of folders for a given organisation. It uses a token to remember where it left off.
If there's a token, it decodes it to find the starting point. The FetchPaginatedFoldersByOrgID function filters folders by organisation ID
and returns the requested number of folders. If there are more folders to fetch, it creates a new token from the last folder's ID.
The encodeToken and decodeToken functions handle converting folder IDs to tokens and back.
This method uses token-based pagination for fetching folders. I chose this approach because it is more efficient,
especially if the query was on an actual database. Tokens keep track of the current position, making it faster to fetch the next
set of results without recalculating offsets.
*/

// Retrieves all folders for a given organisation with pagination.
func GetPaginatedFolders(req *FolderPaginationRequest) (*PaginatedFolderResponse, error) {
	// Fetch all folders by organisation ID with pagination
	folders, err := FetchPaginatedFoldersByOrgID(req.OrgID, req.PerPage, req.Token)
	if err != nil {
		return nil, err
	}

	// Calculate the next token
	var nextToken string
	if len(folders) == req.PerPage {
		lastFolder := folders[len(folders)-1]
		nextToken = encodeToken(lastFolder.Id.String())
	}

	// Create a PaginatedFolderResponse containing the folders and the next token
	return &PaginatedFolderResponse{Folders: folders, Token: nextToken}, nil
}

// Fetches all folders for a given orgID with pagination.
func FetchPaginatedFoldersByOrgID(orgID uuid.UUID, perPage int, token string) ([]*Folder, error) {
	folders := GetSampleData() // Get sample data for demonstration

	// Filter folders by organisation ID
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	// Decode the token to get the start index
	start := 0
	if token != "" {
		decodedToken, err := decodeToken(token)
		if err != nil {
			return nil, err
		}
		startIndex, err := uuid.FromString(decodedToken)
		if err != nil {
			return nil, err
		}
		// Find the start index in the filtered list
		for i, folder := range resFolder {
			if folder.Id == startIndex {
				start = i + 1
				break
			}
		}
	}

	// Apply pagination
	end := start + perPage
	if start >= len(resFolder) {
		return []*Folder{}, nil
	}

	if end > len(resFolder) {
		end = len(resFolder)
	}

	return resFolder[start:end], nil
}

// Encodes the given UUID into a base64 token.
func encodeToken(id string) string {
	return base64.StdEncoding.EncodeToString([]byte(id))
}

// Decodes the base64 token into a UUID string.
func decodeToken(token string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
