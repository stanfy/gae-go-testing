package appenginetesting

import (
	"testing"

	"appengine/datastore"
	"appengine/memcache"
	"appengine/user"
)

type Entity struct {
	Foo, Bar string
}

func TestContext(t *testing.T) {
	c, err := NewContext(nil)
	if err != nil {
		t.Fatalf("NewContext: %v", err)
	}
	defer c.Close()

	_, err = memcache.Get(c, "foo")
	if err != memcache.ErrCacheMiss {
		t.Fatalf("Get err = %v; want ErrCacheMiss", err)
	}

	it := &memcache.Item{
		Key:   "foo",
		Value: []byte("value"),
	}
	err = memcache.Set(c, it)
	if err != nil {
		t.Fatalf("Set err = %v", err)
	}
	it, err = memcache.Get(c, "foo")
	if err != nil {
		t.Fatalf("Get err = %v; want no error", err)
	}
	if string(it.Value) != "value" {
		t.Errorf("got Item.Value = %q; want %q", string(it.Value), "value")
	}

	e := &Entity{Foo: "foo", Bar: "bar"}
	k := datastore.NewKey(c, "Entity", "", 1, nil)
	_, err = datastore.Put(c, k, e)
	if err != nil {
		t.Fatalf("datastore.Put: %v", err)
	}
	u := user.Current(c)
	if u != nil {
		t.Fatalf("User should not be not logged in!")
	}
	c.Login("user@host.com", false)
	u = user.Current(c)
	if u == nil {
		t.Fatalf("User should be logged in!")
	}
	id1 := u.ID
	c.Logout()
	u = user.Current(c)
	if u != nil {
		t.Fatalf("User should not be not logged in!")
	}
	c.Login("differentuser@host.com", false)
	u = user.Current(c)
	if u == nil {
		t.Fatalf("User should be logged in!")
	}
	if id1 == u.ID {
		t.Fatalf("User IDs should be unique")
	}
}
