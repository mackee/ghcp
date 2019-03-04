// Package usecases provides use cases of this application.
package usecases

import (
	"context"
	"log"

	"github.com/int128/ghcp/adaptors/interfaces"
	"github.com/int128/ghcp/git"
	"github.com/int128/ghcp/usecases/interfaces"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func NewPush(i Push) usecases.Push {
	return &i
}

// Push performs commit and push files to the repository.
type Push struct {
	dig.In
	FileSystem adaptors.FileSystem
	GitHub     adaptors.GitHub
}

func (u *Push) Do(ctx context.Context, in usecases.PushIn) error {
	filenames, err := u.FileSystem.FindFiles(in.Paths)
	if err != nil {
		return errors.Wrapf(err, "error while finding files")
	}

	out, err := u.GitHub.GetRepository(ctx, adaptors.GetRepositoryIn{
		Repository: git.RepositoryID{Owner: in.Repository.Owner, Name: in.Repository.Name},
	})
	if err != nil {
		return errors.Wrapf(err, "error while getting the repository")
	}
	log.Printf("Logged in as %s", out.CurrentUserName)

	gitFiles := make([]git.File, len(filenames))
	for i, filename := range filenames {
		content, err := u.FileSystem.ReadAsBase64EncodedContent(filename)
		if err != nil {
			return errors.Wrapf(err, "error while reading file %s", filename)
		}
		blobSHA, err := u.GitHub.CreateBlob(ctx, adaptors.NewBlob{
			Repository: out.Repository,
			Content:    content,
		})
		if err != nil {
			return errors.Wrapf(err, "error while creating a blob for %s", filename)
		}
		gitFiles[i] = git.File{
			Filename: filename,
			BlobSHA:  blobSHA,
			//TODO: Executable
		}
		log.Printf("Uploaded %s as blob %s", filename, blobSHA)
	}

	treeSHA, err := u.GitHub.CreateTree(ctx, adaptors.NewTree{
		Repository:  out.Repository,
		BaseTreeSHA: out.DefaultBranchTreeSHA,
		Files:       gitFiles,
	})
	if err != nil {
		return errors.Wrapf(err, "error while creating a tree")
	}
	log.Printf("Created tree %s", treeSHA)

	commitSHA, err := u.GitHub.CreateCommit(ctx, adaptors.NewCommit{
		Repository:      out.Repository,
		Message:         in.CommitMessage,
		ParentCommitSHA: out.DefaultBranchCommitSHA,
		TreeSHA:         treeSHA,
	})
	if err != nil {
		return errors.Wrapf(err, "error while creating a commit")
	}
	log.Printf("Created commit %s", commitSHA)

	if err := u.GitHub.UpdateBranch(ctx, adaptors.NewBranch{
		Repository: out.Repository,
		BranchName: out.DefaultBranchName,
		CommitSHA:  commitSHA,
	}, false); err != nil {
		return errors.Wrapf(err, "error while creating a branch %s", out.DefaultBranchName)
	}
	log.Printf("Updated branch %s", out.DefaultBranchName)

	return nil
}