package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	mock_repo "url_shortener/internal/repo/mocks"
	"url_shortener/pkg/logging"
)

func Init(ctl *gomock.Controller) (context.Context, logging.Logger, *mock_repo.MockRepository) {
	ctx := context.Background()
	l := logging.GetLogger()
	repo := mock_repo.NewMockRepository(ctl)

	return ctx, l, repo
}

func TestGetUrlCorrect(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx, l, repo := Init(ctl)
	in := "luvHykg_po"

	expUrl := "http://localhost"
	repo.EXPECT().GetUrl(ctx, in).Return(expUrl, nil).Times(1)

	service := NewService(repo, l)
	_, err := service.GetUrl(ctx, in)
	require.NoError(t, err)
}

func TestGetUrlError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx, l, repo := Init(ctl)
	in := ""

	expUrl := ""
	repo.EXPECT().GetUrl(ctx, in).Return(expUrl, errors.New("url doesn't exist"))

	service := NewService(repo, l)
	_, err := service.GetUrl(ctx, in)
	require.Error(t, err)
}

func TestAddUrlCorrect(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx, l, repo := Init(ctl)
	in := "http://localhost"

	expAlias
}
