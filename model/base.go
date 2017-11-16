package model

import (
	"context"

	"github.com/benjamw/golibs/db"
	"google.golang.org/appengine/datastore"
)

type Base struct {
	key *datastore.Key `datastore:"-"`
}

func (b *Base) EntityType() string {
	return "BASE"
}

func (b *Base) GetKey() *datastore.Key {
	return b.key
}

func (b *Base) SetKey(key *datastore.Key) error {
	b.key = key
	return nil
}

func (b *Base) PreSave(ctx context.Context) error {
	return nil
}

func (b *Base) PostSave(ctx context.Context) error {
	return nil
}

func (b *Base) PostLoad(ctx context.Context) error {
	fullKey := b.GetKey()
	if fullKey == nil {
		err := &db.MissingKeyError{}
		return err
	}

	return nil
}

func (b *Base) Transform(ctx context.Context, pl datastore.PropertyList) error {
	return nil
}

func (b *Base) PreDelete(ctx context.Context) error {
	return nil
}

// Prepare gets a properly sized []db.Model ready for use in db.LoadMultiX
func (b *Base) Prepare(num int) []db.Model {
	return nil

	/* concrete versions should look like the following:
	// replace 'Foo' with the concrete model name

	l := make([]*Foo, n)
	ml := make([]db.Model, n)
	for k := range l {
		v := new(Foo)
		ml[k] = db.Model(v)
	}

	return ml

	// also, see PrepSplit below

	*/
}

/*

A related function to go along with the b.Prepare above
but not actually part of the model because it acts on FooList ( []Foo ):
Replace 'Foo' with the concrete model name

type FooList []Foo

// PrepSplit converts the []db.Model slice returned by LoadMultiX to a FooList
func (l *FooList) PrepSplit(ml []db.Model) {
	*l = make(FooList, 0)
	for k := range ml {
		v, ok := ml[k].(*Foo)
		if !ok {
			continue
		}

		*l = append(*l, *v)
	}

	return
}

*/
