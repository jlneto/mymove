package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-openapi/strfmt"

	authcontext "github.com/transcom/mymove/pkg/auth/context"
	uploadop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/uploads"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

type FakeS3 struct {
	putFiles []*s3.PutObjectInput
}

func (fake *FakeS3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	fake.putFiles = append(fake.putFiles, input)
	buf := []byte{}
	_, err := input.Body.Read(buf)
	if err != nil {
		panic(err)
	}
	return nil, nil
}

func (suite *HandlerSuite) TestCreateUploadsHandler() {
	t := suite.T()

	move, err := testdatagen.MakeMove(suite.db)
	if err != nil {
		t.Fatalf("could not create move: %s", err)
	}

	document, err := testdatagen.MakeDocument(suite.db, &move)
	if err != nil {
		t.Fatalf("could not create document: %s", err)
	}
	fakeS3 := &FakeS3{}

	userID := move.UserID

	params := uploadop.NewCreateUploadParams()
	params.MoveID = strfmt.UUID(move.ID.String())
	params.DocumentID = strfmt.UUID(document.ID.String())
	params.File = suite.fixture("test.pdf")

	ctx := authcontext.PopulateAuthContext(context.Background(), userID, "fake token")
	params.HTTPRequest = (&http.Request{}).WithContext(ctx)

	context := NewHandlerContext(suite.db, suite.logger)
	handler := CreateUploadHandler(NewS3HandlerContext(context, fakeS3))
	response := handler.Handle(params)

	createdResponse, ok := response.(*uploadop.CreateUploadCreated)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}

	uploadPayload := createdResponse.Payload
	upload := models.Upload{}
	err = suite.db.Find(&upload, uploadPayload.ID)
	if err != nil {
		t.Fatalf("Couldn't find expected upload.")
	}

	expectedChecksum := "nOE6HwzyE4VEDXn67ULeeA=="
	if upload.Checksum != expectedChecksum {
		t.Errorf("Did not calculate the correct MD5: expected %s, got %s", expectedChecksum, upload.Checksum)
	}

	if len(fakeS3.putFiles) != 1 {
		t.Errorf("Wrong number of putFiles: expected 1, got %d", len(fakeS3.putFiles))
	}

	key := fmt.Sprintf("dev/moves/%s/documents/%s/uploads/%s", move.ID, document.ID, upload.ID)
	if *fakeS3.putFiles[0].Key != key {
		t.Errorf("Wrong key name: expected %s, got %s", key, *fakeS3.putFiles[0].Key)
	}

	pos, err := (*fakeS3.putFiles[0]).Body.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatalf("Could't check position in uploaded file: %s", err)
	}

	if pos != 0 {
		t.Errorf("Wrong file position: expected 0, got %d", pos)
	}

	// TODO verify Body
}
