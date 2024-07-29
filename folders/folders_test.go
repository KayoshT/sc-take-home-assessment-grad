package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("Fetch with one folder for an OrgID", func(t *testing.T) {
		orgID, err := uuid.FromString("9727c9a2-52ec-4787-9d70-2125c0d77db4")
		assert.NoError(t, err)

		req := &folders.FetchFolderRequest{OrgID: orgID}
		expectedFolders := 1
		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, expectedFolders, len(res.Folders))
	})

	t.Run("Fetch with multiple folders for the same OrgID", func(t *testing.T) {
		orgID, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		assert.NoError(t, err)

		req := &folders.FetchFolderRequest{OrgID: orgID}
		expectedFolders := 666

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, expectedFolders, len(res.Folders))
	})

	t.Run("Fetch with non-existing OrgID", func(t *testing.T) {
		orgID, err := uuid.FromString("00000000-0000-0000-0000-000000000000")
		assert.NoError(t, err)

		req := &folders.FetchFolderRequest{OrgID: orgID}
		expectedFolders := 0

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, expectedFolders, len(res.Folders))
	})

	t.Run("Fetch with empty OrgID", func(t *testing.T) {
		orgID := uuid.Nil
		req := &folders.FetchFolderRequest{OrgID: orgID}

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.Equal(t, []*folders.Folder{}, res.Folders)
	})

	t.Run("Fetch with no folders in organisation", func(t *testing.T) {
		orgID, err := uuid.NewV4()
		assert.NoError(t, err)

		req := &folders.FetchFolderRequest{OrgID: orgID}
		expectedFolders := 0

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, expectedFolders, len(res.Folders))
	})

}

func Test_GetPaginatedFolders(t *testing.T) {
	t.Run("Fetch first page of folders for an OrgID", func(t *testing.T) {
		orgID, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		assert.NoError(t, err)

		req := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: 1,
			Token:   "",
		}
		res, err := folders.GetPaginatedFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Folders, 1)
		assert.NotEmpty(t, res.Token)
	})

	t.Run("Fetch second page of folders for the same OrgID", func(t *testing.T) {
		orgID, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		assert.NoError(t, err)

		// Assume that the first page has been fetched and token is retrieved
		firstPageReq := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: 1,
			Token:   "",
		}
		firstPageRes, err := folders.GetPaginatedFolders(firstPageReq)
		assert.NoError(t, err)
		assert.NotNil(t, firstPageRes)
		assert.Len(t, firstPageRes.Folders, 1)
		assert.NotEmpty(t, firstPageRes.Token)

		// Fetch second page using the token from the first page response
		secondPageReq := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: 1,
			Token:   firstPageRes.Token,
		}
		secondPageRes, err := folders.GetPaginatedFolders(secondPageReq)

		assert.NoError(t, err)
		assert.NotNil(t, secondPageRes)
		assert.Len(t, secondPageRes.Folders, 1)
	})

	t.Run("Fetch first page of folders for an OrgID with multiple per page", func(t *testing.T) {
		orgID, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		assert.NoError(t, err)

		req := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: 5,
			Token:   "",
		}
		res, err := folders.GetPaginatedFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Folders, 5)
		assert.NotEmpty(t, res.Token)
	})

	t.Run("Fetch second page of folders for the same OrgID with multiple per page", func(t *testing.T) {
		orgID, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		assert.NoError(t, err)

		// Assume that the first page has been fetched and token is retrieved
		firstPageReq := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: 5,
			Token:   "",
		}
		firstPageRes, err := folders.GetPaginatedFolders(firstPageReq)
		assert.NoError(t, err)
		assert.NotNil(t, firstPageRes)
		assert.Len(t, firstPageRes.Folders, 5)
		assert.NotEmpty(t, firstPageRes.Token)

		// Fetch second page using the token from the first page response
		secondPageReq := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: 5,
			Token:   firstPageRes.Token,
		}
		secondPageRes, err := folders.GetPaginatedFolders(secondPageReq)

		assert.NoError(t, err)
		assert.NotNil(t, secondPageRes)
		assert.Len(t, secondPageRes.Folders, 5)
	})

	t.Run("Fetch with non-existing OrgID", func(t *testing.T) {
		orgID, err := uuid.FromString("00000000-0000-0000-0000-000000000000")
		assert.NoError(t, err)

		req := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: 1,
			Token:   "",
		}
		res, err := folders.GetPaginatedFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Empty(t, res.Folders)
		assert.Empty(t, res.Token)
	})

	t.Run("Invalid OrgID format", func(t *testing.T) {
		_, err := uuid.FromString("invalid-uuid")
		assert.Error(t, err)
	})

	t.Run("Fetch with a larger page size than available folders", func(t *testing.T) {
		orgID, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		assert.NoError(t, err)

		req := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: 1000,
			Token:   "",
		}
		res, err := folders.GetPaginatedFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.True(t, len(res.Folders) <= 1000)
		assert.Empty(t, res.Token) // No next token as we've fetched all available folders
	})
}

func Test_FoldersPaginationIntegration(t *testing.T) {
	orgID, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	assert.NoError(t, err)

	// Fetch all folders using GetAllFolders
	reqAll := &folders.FetchFolderRequest{OrgID: orgID}
	allFoldersRes, err := folders.GetAllFolders(reqAll)
	assert.NoError(t, err)
	assert.NotNil(t, allFoldersRes)
	totalFolders := len(allFoldersRes.Folders)

	// Fetch folders using pagination
	perPage := 20
	var paginatedFolders []folders.Folder
	token := ""
	for {
		reqPage := &folders.FolderPaginationRequest{
			OrgID:   orgID,
			PerPage: perPage,
			Token:   token,
		}
		pageRes, err := folders.GetPaginatedFolders(reqPage)
		assert.NoError(t, err)
		assert.NotNil(t, pageRes)

		// Append paginated folders
		for _, folder := range pageRes.Folders {
			paginatedFolders = append(paginatedFolders, *folder)
		}
		token = pageRes.Token

		if token == "" {
			break
		}
	}
	// Assert that the total number of folders matches
	assert.Equal(t, totalFolders, len(paginatedFolders))
}
