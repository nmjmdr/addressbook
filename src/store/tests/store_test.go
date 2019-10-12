package tests

import (
	"testing"
	"addressbook/store"
	mocks "addressbook/store/mock_store"
	. "github.com/golang/mock/gomock"
	"addressbook/models"
)

func Test_Given_That_Item_Is_Present_In_Cache_Against_Id_It_Should_Fetch_It_From_Cache(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	cache := mocks.NewMockCache(ctrl)
	proxy := mocks.NewMockAutopilotProxy(ctrl)

	id := "1"
	contact := models.Contact{
		ID: id,
		Email: "1@1.com",
	}

	cache.
		EXPECT().
		Get(Any()).
		Return(contact, nil).
		Times(1)

	// expect that proxy should not be called
	proxy.
		EXPECT().
		Get(Any()).
		Return([]models.Contact{ contact }, nil).
		Times(0)

	st := store.NewStore(cache, proxy)

	returned, err := st.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if len(returned) == 0 || returned[0].ID != id {
		t.Fatal("Unexpected contact value returned")
	}
}

func Test_Given_That_Item_Is_Present_In_Cache_Against_Email_It_Should_Fetch_It_From_Cache(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	cache := mocks.NewMockCache(ctrl)
	proxy := mocks.NewMockAutopilotProxy(ctrl)

	id := "1"
	contact := models.Contact{
		ID: id,
		Email: "1@1.com",
	}

	cache.
		EXPECT().
		Get(Any()).
		Return(contact, nil).
		Times(1)

	// expect that proxy should not be called
	proxy.
		EXPECT().
		Get(Any()).
		Return([]models.Contact{ contact }, nil).
		Times(0)

	st := store.NewStore(cache, proxy)

	returned, err := st.Get(contact.Email)
	if err != nil {
		t.Fatal(err)
	}

	if len(returned) == 0 || returned[0].ID != id {
		t.Fatal("Unexpected contact value returned")
	}
}

func Test_Given_That_Item_Is_Not_Present_In_Cache_It_Should_Fetch_It_From_API_And_Store_It(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	cache := mocks.NewMockCache(ctrl)
	proxy := mocks.NewMockAutopilotProxy(ctrl)

	id := "1"
	contact := models.Contact{
		ID: id,
		Email: "1@1.com",
	}

	cache.
		EXPECT().
		Get(Any()).
		Return(nil, nil).
		Times(1)

	// expect that proxy should not be called
	proxy.
		EXPECT().
		Get(Any()).
		Return([]models.Contact{ contact }, nil).
		Times(1)

	cache.
		EXPECT().
		Add(Eq(contact.ID), Any()).
		Return(nil).
		Times(1)
	
	cache.
		EXPECT().
		Add(Eq(contact.Email), Any()).
		Return(nil).
		Times(1)

	st := store.NewStore(cache, proxy)

	returned, err := st.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if len(returned) == 0 || returned[0].ID != id {
		t.Fatal("Unexpected contact value returned")
	}
}

func Test_When_Item_Is_Updated_It_Should_It_Should_Invoke_Upsert_API_And_Remove_From_Cache(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	cache := mocks.NewMockCache(ctrl)
	proxy := mocks.NewMockAutopilotProxy(ctrl)

	id := "1"
	contact := models.Contact{
		ID: id,
		Email: "1@1.com",
	}

	cache.
		EXPECT().
		Del(Eq(contact.ID)).
		Return(nil).
		Times(1)
	
	cache.
		EXPECT().
		Del(Eq(contact.Email)).
		Return(nil).
		Times(1)

	// expect that proxy should not be called
	proxy.
		EXPECT().
		Upsert(Any()).
		Return(nil).
		Times(1)


	st := store.NewStore(cache, proxy)

	err := st.Upsert(contact)
	if err != nil {
		t.Fatal(err)
	}
}