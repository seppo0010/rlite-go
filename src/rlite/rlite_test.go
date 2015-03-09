package rlite

import (
	"reflect"
	"testing"
)

func TestOpen(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Error("Unable to open database")
	}
	Close(db)
}

func TestStatus(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Error("Unable to open database")
	}
	reply, err := Command(db, []string{"SET", "key", "value"})
    if err != nil {
        t.Error("Got error")
    }
    if reply != "OK" {
        t.Error("Got invalid reply")
    }
}

func TestString(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Error("Unable to open database")
	}
	Command(db, []string{"SET", "key", "value"})

	reply, err := Command(db, []string{"GET", "key"})
    if err != nil {
        t.Error("Got error")
    }
    if reply != "value" {
        t.Error("Got invalid reply")
    }
}

func TestInteger(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Error("Unable to open database")
	}
	reply, err := Command(db, []string{"RPUSH", "key", "value", "value2"})

    if err != nil {
        t.Error("Got error")
    }
    if reply != 2 {
        t.Error("Got invalid reply")
    }
}

func TestError(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Error("Unable to open database")
	}
	_, err = Command(db, []string{"SET", "key"})

    if err == nil {
        t.Error("Got no error")
    }
}

func TestNil(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Error("Unable to open database")
	}
	reply, err := Command(db, []string{"GET", "key"})

    if err != nil {
        t.Error("Got error")
    }
    if reply != nil {
        t.Error("Got reply")
    }
}

func TestArray(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Error("Unable to open database")
	}
	Command(db, []string{"MSET", "key", "value", "key2", "value2"})
	reply, err := Command(db, []string{"MGET", "key", "key2", "key3"})

    if err != nil {
        t.Error("Got error")
    }

    if reflect.TypeOf(reply).Kind() != reflect.Slice {
        t.Error("Expected slice")
    }

    slice := reflect.ValueOf(reply)

    if slice.Len() != 3 {
        t.Error("Unexpected length")
    }

    if slice.Index(0).Interface() != "value" {
        t.Error("Unexpected value on position 0")
    }

    if slice.Index(1).Interface() != "value2" {
        t.Error("Unexpected value on position 1")
    }

    if slice.Index(2).Interface() != nil {
        t.Error("Unexpected value on position 2")
    }
}
