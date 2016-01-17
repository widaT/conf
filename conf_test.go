package conf

import (
    "testing"
)

var (
    conf config
)

func init() {
    conf = NewConfig("conf.conf")
}


func TestGetInt( t *testing.T) {
    if conf.GetInt("mysql","port") != 3306 {
        t.Errorf("(expected) %d != %d (actual)",conf.GetInt("mysql","port"),3600)
    }
}

func TestGetInt64 (t *testing.T){
    if conf.GetInt64("test","int64") != 9223372036854775807 {
        t.Errorf("(expected) %d != %d (actual)",conf.GetInt64("test","int64"),9223372036854775807)
    }
}


func TestGetFloat64(t *testing.T) {
    if conf.GetFloat64("test","float") !=  3.40282e+038 {
        t.Errorf("(expected) %f != %f (actual)",conf.GetFloat64("test","float"), 3.40282e+038)
    }
}


func TestGetArray(t *testing.T) {
    arr := conf.GetArray("test","arr",",")
    arr2 := []string{"1","2","3","3"}
    for i,v :=  range arr {
        if arr2[i] != v {
            t.Errorf("(expected) %s != %s (actual)",arr2[i], v)
        }
    }
}