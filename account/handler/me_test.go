package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"

	"github.com/thesirtaylor/memrizr/mocks"
	"github.com/thesirtaylor/memrizr/utils"

	"github.com/google/uuid"
	"github.com/thesirtaylor/memrizr/model"
)


func TestMe(t *testing.T){
	gin.SetMode(gin.TestMode);

	t.Run(("Success"), func(t *testing.T) {
		uid, _ := uuid.NewRandom();

		mockUserResp := &model.User{
			UID: uid,
			Email: "ajayiolatunji17@gmail.com",
			Name: "Ajayi Olatunji",
		}

		mockUserService := new(mocks.MockUserService);
		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil);

		responseRecord := httptest.NewRecorder();

		//use middleware to set context for test

		router := gin.Default();
		router.Use(func(ctx *gin.Context) {
			ctx.Set("user", &model.User{
				UID: uid,
			});
		})

		NewHandler(&Config{
			R: router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil);
		assert.NoError(t, err);

		router.ServeHTTP(responseRecord, request);

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})

		assert.NoError(t, err);
		// assert.Equal(t, http.StatusOK, responseRecord.Code);
		assert.Equal(t, respBody, responseRecord.Body.Bytes());
		mockUserService.AssertExpectations(t);
	})

	t.Run("NoContextUser", func(t *testing.T){
		mockUserService := new(mocks.MockUserService);
		mockUserService.On(("Get"), mock.Anything, mock.Anything).Return(nil, nil);

		responseRecord := httptest.NewRecorder();

		router := gin.Default();
		NewHandler(&Config{
			R: router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil);
		assert.NoError(t, err);

		router.ServeHTTP(responseRecord, request);

		assert.Equal(t, http.StatusInternalServerError, responseRecord.Code);
		mockUserService.AssertNotCalled(t, "Get", mock.Anything);
	})

	t.Run(("NotFound"), func(t *testing.T) {
		uid, _ := uuid.NewRandom();

		mockUserService := new(mocks.MockUserService);
		mockUserService.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down call chain"));

		responseRecord := httptest.NewRecorder();

		router := gin.Default();
		router.Use(func(ctx *gin.Context) {
			ctx.Set("user", &model.User{
				UID: uid,
			});
		})

		NewHandler(&Config{
			R: router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil);
		assert.NoError(t, err);

		router.ServeHTTP(responseRecord, request);

		respError := utils.NewNotFound("user", uid.String());

		respBody, err := json.Marshal(gin.H{
			"error": respError,
		})

		assert.NoError(t, err);

		assert.Equal(t, respError.Message, responseRecord.Code);
		assert.Equal(t, respBody, responseRecord.Body.Bytes());
		mockUserService.AssertExpectations(t);
	})
}