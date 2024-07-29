package folders

import (
	"github.com/gofrs/uuid"
)

/*
Deletions:
- Removed unused variables
- Removed two unnecessary 'for' loops as we can work with the slice of pointers that are returned by the above function.
Retrieves all folders for a given organization.
*/
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	folders, err := FetchAllFoldersByOrgID(req.OrgID) // Fetch all folders by organization ID
	// Adding error handling in case future iterations of FetchAllFoldersbyOrgID can create errors.
	if err != nil {
		return nil, err
	}
	return &FetchFolderResponse{Folders: folders}, nil // Create a FetchFolderResponse containing the folders
}

/*
Assumption: Fetching all folders including the ones that are flagged as deleted.
Suggestion: To keep "OrgId" consistent in the codebase I would suggest to change the
function names and values to "OrgId" or in the folder type declaration
change "OrgId" to "OrgID" .
Fetches all folders for a given organization ID.
*/
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() // Get sample data for demonstration
	// Filter folders by organization ID
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder) // Append the matching folder to the result slice
		}
	}
	return resFolder, nil
}

/*
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Remove Unused Variables
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	f := []Folder{}
	// Need to retrieve er
	r, _ := FetchAllFoldersByOrgID(req.OrgID)

	// two 'for' loops unnecessary as we can work with the slice of pointers that are returned by the above function.
	for k, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}
*/
