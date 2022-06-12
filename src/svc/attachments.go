package svc

import (
	"github.com/betalixt/bloggo/clnt"
	"github.com/betalixt/bloggo/clnt/models"
	"github.com/betalixt/bloggo/util/txcontext"
)

type AttachmentService struct {
	fsclnt *clnt.FileServiceClient
}

func (svc *AttachmentService) GetAttachmentMeta(
	ctx *txcontext.TransactionContext,
	fileId string,
) (*models.FileMeta, error) {
  return svc.fsclnt.GetFileMeta(ctx, fileId)
}

func NewAttachmentService(
  fsclnt *clnt.FileServiceClient,
) *AttachmentService {
  return &AttachmentService{
    fsclnt: fsclnt,
  }
}
