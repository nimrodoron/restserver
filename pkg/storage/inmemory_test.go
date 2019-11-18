package storage

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInMemoryStorage_Persist(t *testing.T) {
	storage := CreateInMemoryStorage()
	storage.Persist("test", "test")
	v := storage.m["test"]
	if v != "test" {
		t.Errorf("Persist failed. Expected: %v, got: %v", "test", v)
	} else {
		t.Logf("Persist success. Expected: %v, got: %v", "test", v)
	}
}

func TestInMemoryStorage_Persist_Override(t *testing.T) {
	storage := CreateInMemoryStorage()
	storage.m["test"] = "oldtest"
	storage.Persist("test", "test")
	v := storage.m["test"]
	if v != "test" {
		t.Errorf("Persist failed. Expected: %v, got: %v", "test", v)
	} else {
		t.Logf("Persist success. Expected: %v, got: %v", "test", v)
	}
}

func TestInMemoryStorage_Retrieve(t *testing.T) {
	storage := CreateInMemoryStorage()
	storage.m["test"] = "test"
	v, e := storage.Retrieve("test")
	if v != "test" {
		t.Errorf("Retrieve failed. Expected: %v, got: %v", "test", v)
	} else {
		t.Logf("Retrieve success. Expected: %v, got: %v", "test", v)
	}
	if e != nil {
		t.Errorf("Retrieve failed. Expected error: %v, got: %v", nil, e)
	} else {
		t.Logf("Retrieve success. Expected error: %v, got: %v", nil, e)
	}
}

func TestInMemoryStorage_Retrieve_NotExist(t *testing.T) {
	storage := CreateInMemoryStorage()
	v, e := storage.Retrieve("test")
	if v != "" {
		t.Errorf("Retrieve failed. Expected: %v, got: %v", "", v)
	} else {
		t.Logf("Retrieve success. Expected: %v, got: %v", "", v)
	}
	if e != nil {
		t.Logf("Retrieve success. Expected error not : %v, got: %v", nil, e)
	} else {
		t.Errorf("Persist failed. Expected error: %v, got: %v", nil, e)
	}
}

var _= Describe("InMemoryStorage", func() {
	var (
		storage *InMemoryStorage
	)

	BeforeEach(func() {
		storage = CreateInMemoryStorage()
	})

	Describe("Persist in storage", func() {
		Context("Resource with same name not exist", func() {
			It("Resource saved successfully", func() {
				storage.Persist("test", "test")
				Expect(storage.m["test"]).To(Equal("test"))
			})
		})
		Context("Resource with same name already exist", func() {
			It("Resource override successfully", func() {
				storage.m["test"] = "oldtest"
				storage.Persist("test", "test")
				Expect(storage.m["test"]).To(Equal("test"))
			})
		})
	})

	Describe("Retrieve from storage", func() {
		var (
			v string
			e error
		)
		BeforeEach(func() {
			storage.Persist("test", "test")
		})
		Context("Resource exist in storage", func() {
			BeforeEach(func() {
				v,e = storage.Retrieve("test")
			})
			It("Resource returned successfully", func() {
				Expect(v).To(Equal("test"))
			})
			It("Should not error", func() {
				Expect(e).NotTo(HaveOccurred())
			})
		})
		Context("Resource not exist in storage", func() {
			BeforeEach(func() {
				v,e = storage.Retrieve("test2")
			})
			It("Resource returned successfully", func() {
				Expect(v).To(Equal(""))
			})
			It("Should not error", func() {
				Expect(e).To(HaveOccurred())
			})
		})
	})
})